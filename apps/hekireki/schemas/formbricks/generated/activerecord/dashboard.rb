# Represents a dashboard containing multiple charts.
# Dashboards aggregate analytics insights at the workspace level.
#
# @property id - Unique identifier for the dashboard
# @property name - Display name of the dashboard
# @property workspace - The workspace this dashboard belongs to
# @property widgets - Charts on this dashboard
# @property filters - The dashboard filters pinned to its top bar, in order: `{ id, field, operator, label }`,
# at most 10 (ADR-0026). Each viewer's picks live in their own browser, not here.
# @property createdBy - User who created the dashboard
class Dashboard < ApplicationRecord
  self.table_name = "Dashboard"
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }
  attribute :filters, default: -> { JSON.parse("[]") }

  validates :name, presence: true, uniqueness: { scope: :workspaceId }

  belongs_to :workspace, foreign_key: "workspaceId", inverse_of: :dashboards
  belongs_to :creator, class_name: "User", foreign_key: "createdBy", optional: true, inverse_of: :dashboards
  has_many :widgets, class_name: "DashboardWidget", foreign_key: "dashboardId", inverse_of: :dashboard,
                     dependent: :destroy
end
