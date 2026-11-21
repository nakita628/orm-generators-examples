# A Dashboards v2 widget: a srcdoc definition placed on a DashboardsV2 board, or
# unplaced. Widgets made before the split were copied from CustomGraph with their
# ids. No relation, per the database rule.
#
# dashboardId:: The DashboardsV2 board it sits on; null while unplaced.
# definition:: The widget's srcdoc definition.
class DashboardWidget < ApplicationRecord
  self.table_name = "DashboardWidget"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  attribute :gridColumn, default: 0
  attribute :gridRow, default: 0
  attribute :colSpan, default: 1
  attribute :rowSpan, default: 1
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :projectId, presence: true
  validates :name, presence: true
  validates :definition, exclusion: { in: [ nil ] }
end
