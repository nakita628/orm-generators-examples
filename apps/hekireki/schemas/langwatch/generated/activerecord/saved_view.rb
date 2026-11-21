# kind:: Discriminator for the two storage shapes that share this table:
#   - "v1-traces-filter" (default) - legacy /messages page filter map saved views
#   - "v2-traces-lens"             - new traces v2 lens configs (columns, addons, grouping, sort, filterText) encoded inside `filters` JSON
#   Each client filters by its own kind on read so neither sees the other's rows.
class SavedView < ApplicationRecord
  self.table_name = "SavedView"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  attribute :filters, default: -> { JSON.parse("{}") }
  attribute :order, default: 0
  attribute :kind, default: "v1-traces-filter"
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :name, presence: true

  belongs_to :project, foreign_key: "projectId", inverse_of: :saved_views
  belongs_to :user, foreign_key: "userId", optional: true, inverse_of: :saved_views
end
