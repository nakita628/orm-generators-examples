# frozen_string_literal: true

# Runs the plan through the Active Record models hekireki wrote, against the database Prisma
# created.
#
#   ruby harness.rb crud <plan.json> <models dir> <url>
#
# Active Record reads the columns from the database, so every column has an attribute; what is
# held to the plan is the model's table, its attribute defaults, its type casts (what a value
# reads back as) and its associations.
require "base64"
require "bigdecimal"
require "json"
require "bundler/setup"
# What a Rails application does at boot: every gem of the Gemfile loaded, so that the generated
# defaults find Cuid, Cuid2, Nanoid and ULID.
Bundler.require
require "active_record"

mode, plan_path, models_dir, url = ARGV
raise "crud only" unless mode == "crud"

plan = JSON.parse(File.read(plan_path))

ActiveRecord::Base.logger = nil
ActiveRecord.default_timezone = :utc
config =
  case plan["dialect"]
  when "sqlite" then { adapter: "sqlite3", database: url.delete_prefix("file:") }
  when "mysql" then url.sub(%r{\Amysql://}, "mysql2://")
  else url
  end
ActiveRecord::Base.establish_connection(config)

# hekireki-activerecord writes application_record.rb only into a directory named `models`.
base = File.join(models_dir, "application_record.rb")
if File.exist?(base)
  require base
else
  class ApplicationRecord < ActiveRecord::Base
    primary_abstract_class
  end
end
files = Dir[File.join(models_dir, "*.rb")].sort - [File.expand_path(base)]
support, models_only = files.partition { |f| File.basename(f) == "prisma_date_time.rb" }
(support + models_only).each { |f| require f }

models = ApplicationRecord.descendants.reject(&:abstract_class?).to_h { |m| [m.table_name, m] }

def message(e)
  "#{e.class}: #{e.message}"[0, 600]
end

def to_ruby(field, value, column_type)
  return nil if value.nil?
  return value.map { |v| to_ruby(field.merge("list" => false), v, column_type) } if field["list"]

  case field["type"]
  when "DateTime" then Time.iso8601(value).utc
  when "BigInt" then Integer(value)
  when "Decimal" then BigDecimal(value)
  when "Bytes" then Base64.strict_decode64(value)
  else value
  end
end

def normal(value, type = nil)
  case value
  when nil then nil
  # A time column reads back on 2000-01-01, Rails' dummy date: only the time of day is the value.
  when Time then type == :time ? value.utc.strftime("1970-01-01T%H:%M:%S.%LZ") : value.utc.strftime("%Y-%m-%dT%H:%M:%S.%LZ")
  when Time, DateTime then value.to_time.utc.strftime("%Y-%m-%dT%H:%M:%S.%LZ")
  when Date then value.strftime("%Y-%m-%dT00:00:00.000Z")
  when BigDecimal then value.to_s("F").sub(/\.0\z/, "")
  when Array then value.map { |v| normal(v, type) }
  when Hash then JSON.parse(JSON.generate(value))
  when String
    if type == :binary || value.encoding == Encoding::BINARY && !value.dup.force_encoding("UTF-8").valid_encoding?
      Base64.strict_encode64(value)
    else
      value
    end
  when Integer then value.abs > 2**53 ? value.to_s : value
  else
    value.respond_to?(:iso8601) ? normal(value.to_time) : value
  end
end

# An enum attribute reads back as the enum's key (smtp_auth); the plan holds the database value
# (SMTP_AUTH) the key maps to.
def db_value(model, record, name)
  value = record.read_attribute(name)
  mapping = model.defined_enums[name] || model.defined_enums[model.attribute_aliases.key(name)]
  return value unless mapping

  value.is_a?(Array) ? value.map { |v| mapping.fetch(v, v) } : mapping.fetch(value, value)
end

def key_of(model, record)
  pk = Array(model.primary_key)
  pk.size == 1 ? record.read_attribute(pk.first) : pk.map { |c| record.read_attribute(c) }
end

def find(model, key)
  pk = Array(model.primary_key)
  conditions = pk.size == 1 ? { pk.first => key } : pk.zip(key).to_h
  model.find_by(conditions)
end

out = { "orm" => "activerecord", "models" => [] }
rows = {}
state = {}

plan["models"].each do |m|
  r = { "model" => m["name"] }
  out["models"] << r
  model = models[m["table"]]
  unless model
    r["missing"] = true
    next
  end
  begin
    columns = model.columns_hash
  rescue StandardError => e
    r["read"] = { "ok" => false, "error" => message(e) }
    next
  end
  r["unmapped"] = m["fields"].map { |f| f["column"] }.reject { |c| columns.key?(c) }
  unless m["insertable"]
    r["skipped"] = "a required key closes a cycle"
    next
  end
  record = nil
  begin
    record = model.new
    m["fields"].each do |f|
      next unless columns.key?(f["column"])

      case f["write"]
      when "value"
        record.write_attribute(f["column"], to_ruby(f, f["value"], columns[f["column"]].type))
      when "parent"
        parent = f["fk"]["available"] ? rows[f["fk"]["model"]] : nil
        record.write_attribute(f["column"], parent ? to_ruby(f, parent[f["fk"]["column"]], columns[f["column"]].type) : nil)
      end
    end
    record.save!
    r["insert"] = { "ok" => true }
  rescue StandardError => e
    r["insert"] = { "ok" => false, "error" => message(e) }
    next
  end
  key = key_of(model, record)
  r["pk"] = Array(model.primary_key).to_h { |c| [c, normal(record.read_attribute(c))] }
  begin
    loaded = find(model, key) or raise "no row came back for the key the insert returned"
    values = columns.to_h { |name, col| [name, normal(db_value(model, loaded, name), col.type)] }
    r["read"] = { "ok" => true, "values" => values }
    rows[m["name"]] = values
  rescue StandardError => e
    r["read"] = { "ok" => false, "error" => message(e) }
    next
  end
  if m["update"]
    begin
      col = m["update"]["column"]
      find(model, key).update!(col => m["update"]["value"])
      r["update"] = { "ok" => true, "value" => normal(db_value(model, find(model, key), col)) }
    rescue StandardError => e
      r["update"] = { "ok" => false, "error" => message(e) }
    end
  end
  state[m["name"]] = [model, key]
end

squash = ->(s) { s.to_s.downcase.delete("_") }
by_model = out["models"].to_h { |r| [r["model"], r] }
plan["models"].each do |m|
  next unless state[m["name"]]

  model, key = state[m["name"]]
  by_model[m["name"]]["relations"] = m["relations"].map do |rel|
    result = { "name" => rel["name"], "kind" => rel["kind"] }
    next result.merge("skipped" => "no row on the other side") unless state[rel["model"]]

    assoc = model.reflect_on_all_associations.find { |a| squash.(a.name) == squash.(rel["name"]) }
    next result.merge("found" => false) unless assoc

    begin
      value = find(model, key).public_send(assoc.name)
      count = value.respond_to?(:to_a) && !value.is_a?(ActiveRecord::Base) ? value.to_a.size : (value.nil? ? 0 : 1)
      result.merge("found" => true, "loaded" => count)
    rescue StandardError => e
      result.merge("found" => true, "error" => message(e))
    end
  end
end

plan["models"].reverse_each do |m|
  next unless state[m["name"]]

  model, key = state[m["name"]]
  r = by_model[m["name"]]
  begin
    find(model, key).destroy!
    r["delete"] = find(model, key).nil? ? { "ok" => true } : { "ok" => false, "error" => "the row is still there" }
  rescue StandardError => e
    r["delete"] = { "ok" => false, "error" => message(e) }
  end
end

$stdout.write(JSON.generate(out))
