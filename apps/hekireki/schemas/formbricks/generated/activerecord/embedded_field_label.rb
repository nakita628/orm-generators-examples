# A workspace's display name for an imported (embedded-data) key of its Feedback Datasets, shown in
# place of the key in chart legends, headers and the chart builder's field pickers. One per key and
# workspace; clearing the name deletes the row, and the key shows as collected again. Built-in fields
# are never renamed here: they keep their translated names.
#
# @property key - The embedded-data key exactly as collected, case included
# @property label - The name shown for the key across the workspace
class EmbeddedFieldLabel < ApplicationRecord
  self.table_name = "EmbeddedFieldLabel"
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }

  validates :key, presence: true, uniqueness: { scope: :workspaceId }
  validates :label, presence: true

  belongs_to :workspace, foreign_key: "workspaceId", inverse_of: :embedded_field_labels
end
