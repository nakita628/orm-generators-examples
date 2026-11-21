# Defines org-scoped data store overrides (e.g. dedicated ClickHouse for HIPAA orgs).
# Multiple organizations can share a single data store row via organizationIds.
#
# key:: Human-readable unique key (e.g. "hipaa-clickhouse-us-east")
# organizationIds:: Organization IDs that use this data store
# config:: Versioned config JSON. Structure is discriminated by the top-level `version` field.
class OrganizationDataStore < ApplicationRecord
  self.table_name = "OrganizationDataStore"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  attribute :organizationIds, default: -> { [] }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :kind, { clickhouse: "CLICKHOUSE" }, validate: true

  validates :key, presence: true, uniqueness: true
  validates :config, exclusion: { in: [ nil ] }
end
