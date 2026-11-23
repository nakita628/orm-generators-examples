# Loads a Rails app's db/schema.rb into an empty database with Active Record alone, the way
# `rails db:schema:load` does, without the app: the tables, indexes, foreign keys, check
# constraints, enums and extensions Active Record writes from the app's schema definition.
#
#   ruby load_schema.rb <database url> <schema.rb>
#
# Run by pull.mjs in the image of ../../activerecord (Ruby, Active Record and the pg, mysql2 and
# sqlite3 drivers). An app's schema.rb may also use the DSL of a gem the image does not have; the
# common ones are defined here the way those gems define them: scenic's create_view and fx's
# create_function and create_trigger run the SQL the schema holds.
require "active_record"

url, file = ARGV
abort "usage: ruby load_schema.rb <database url> <schema.rb>" unless url && file

module PullSchemaShims
  # scenic
  def create_view(name, version: nil, sql_definition: nil, materialized: false)
    raise ArgumentError, "create_view #{name}: no sql_definition" unless sql_definition

    execute("CREATE #{materialized ? 'MATERIALIZED ' : ''}VIEW #{quote_table_name(name)} AS #{sql_definition}")
  end

  # fx
  def create_function(_name, sql_definition: nil, **)
    execute(sql_definition) if sql_definition
  end

  # fx's create_trigger takes the SQL; hairtrigger's is a chain that ends in a block giving the
  # trigger's body: create_trigger("name", generated: true).on("table").after(:insert)
  # .for_each(:row) { "SQL" }.
  def create_trigger(name, sql_definition: nil, **)
    return execute(sql_definition) if sql_definition

    HairTrigger.new(self, name)
  end

  # What hairtrigger writes on PostgreSQL: a plpgsql function of the trigger's name around the
  # body, and the trigger that runs it. The block with the body goes to whichever link of the
  # chain comes last (`for_each(:row)`, `declare(...)`, ...); that link writes the trigger.
  class HairTrigger
    def initialize(connection, name)
      @connection = connection
      @name = name
      @events = []
      @level = :row
    end

    def on(table, &) = step(-> { @table = table }, &)
    def name(name, &) = step(-> { @name = name }, &)
    def before(*events, &) = step(-> { @timing, @events = "BEFORE", events }, &)
    def after(*events, &) = step(-> { @timing, @events = "AFTER", events }, &)
    def of(*columns, &) = step(-> { @columns = columns }, &)
    def where(condition, &) = step(-> { @where = condition }, &)
    def declare(declarations, &) = step(-> { @declare = declarations }, &)
    def for_each(level, &) = step(-> { @level = level }, &)
    def security(*, &) = step(-> {}, &)
    def nowrap(*, &) = step(-> {}, &)

    private

    def step(set)
      set.call
      return self unless block_given?

      write(yield)
    end

    def write(body)
      events = @events.map { |e| e.to_s.upcase }
      events = events.map { |e| e == "UPDATE" && @columns ? "UPDATE OF #{@columns.join(', ')}" : e }
      @connection.execute(<<~SQL)
        CREATE OR REPLACE FUNCTION #{@name}() RETURNS TRIGGER LANGUAGE plpgsql AS $$
        #{@declare ? "DECLARE #{@declare};" : ''}
        BEGIN
          #{body}
          RETURN #{@timing == 'BEFORE' ? 'NEW' : 'NULL'};
        END;
        $$
      SQL
      @connection.execute(
        "CREATE TRIGGER #{@name} #{@timing} #{events.join(' OR ')} ON #{@connection.quote_table_name(@table)} " \
        "FOR EACH #{@level.to_s.upcase}#{@where ? " WHEN (#{@where})" : ''} EXECUTE FUNCTION #{@name}()",
      )
    end
  end
end

ActiveRecord::Base.establish_connection(url)
ActiveRecord::Schema.verbose = false
connection = ActiveRecord::Base.connection
connection.singleton_class.include(PullSchemaShims)

if connection.adapter_name == "PostgreSQL"
  # neighbor: `t.vector "embedding", limit: 1536` is pgvector's vector(1536).
  ActiveRecord::ConnectionAdapters::PostgreSQL::TableDefinition.class_eval do
    def vector(*names, limit: nil, **options)
      names.each { |name| column(name, limit ? "vector(#{limit})" : "vector", **options) }
    end
  end
end
# A schema.rb of a newer Rails than the image's (`ActiveRecord::Schema[8.2]` on Rails main) loads
# under the newest version this Active Record knows.
source = File.read(file)
current = ActiveRecord::Migration.current_version.to_s
source = source.sub(/ActiveRecord::Schema\[(\d+\.\d+)\]/) do |match|
  Gem::Version.new(Regexp.last_match(1)) > Gem::Version.new(current) ? "ActiveRecord::Schema[#{current}]" : match
end
eval(source, TOPLEVEL_BINDING, File.expand_path(file)) # rubocop:disable Security/Eval
puts connection.tables.sort
