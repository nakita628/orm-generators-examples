# query:: The TSQL query text that was executed
# scope:: The scope of the query (determines which tenant IDs were used)
# stats:: Query execution statistics from ClickHouse
# title:: AI-generated title summarizing the query
# source:: Where the query originated from
# user:: Optional user who executed the query (null for API calls)
# filterPeriod:: Time filter settings used when the query was executed
#   Period like "7d", "24h", etc.
# filterFrom:: Custom start date (ISO 8601)
# filterTo:: Custom end date (ISO 8601)
class CustomerQuery < ApplicationRecord
  self.table_name = "CustomerQuery"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Cuid.generate }
  alias_attribute :created_at, :createdAt

  enum :scope, { organization: "ORGANIZATION", project: "PROJECT", environment: "ENVIRONMENT" }, validate: true
  enum :source, { dashboard: "DASHBOARD", api: "API" }, default: :dashboard, validate: true

  validates :query, presence: true
  validates :stats, exclusion: { in: [ nil ] }

  belongs_to :organization, foreign_key: "organizationId", inverse_of: :customer_queries
  belongs_to :project, foreign_key: "projectId", optional: true, inverse_of: :customer_queries
  belongs_to :environment, class_name: "RuntimeEnvironment", foreign_key: "environmentId", optional: true,
                           inverse_of: :customer_queries
  belongs_to :user, foreign_key: "userId", optional: true, inverse_of: :customer_queries
end
