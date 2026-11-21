# Represents a chart/visualization that can be used in multiple dashboards.
# Charts are reusable components that query analytics data.
#
# @property id - Unique identifier for the chart
# @property name - Display name of the chart
# @property type - Type of visualization (bar, line, pie, etc.)
# @property workspace - The workspace this chart belongs to
# @property query - Cube.js query configuration (JSON)
# @property config - Chart-specific configuration (colors, labels, etc.)
# @property createdBy - User who created the chart
# @property dashboards - Dashboards that use this chart
#
# query:: [ChartQuery] - Cube.js query configuration
# config:: [ChartConfig] - Visualization configuration (colors, labels, formatting)
class Chart < ApplicationRecord
  self.table_name = "Chart"
  self.inheritance_column = nil
  self.implicit_order_column = "created_at"

  attribute :id, default: -> { Cuid.generate }
  attribute :query, default: -> { JSON.parse("{}") }
  attribute :config, default: -> { JSON.parse("{}") }

  enum :type, { area: "area", bar: "bar", pie: "pie", big_number: "big_number", matrix: "matrix" }, validate: true

  validates :name, presence: true, uniqueness: { scope: :workspaceId }

  belongs_to :workspace, foreign_key: "workspaceId", inverse_of: :charts
  belongs_to :creator, class_name: "User", foreign_key: "createdBy", optional: true, inverse_of: :charts
  belongs_to :feedback_directory, foreign_key: "feedbackDirectoryId", inverse_of: :charts
  has_many :widgets, class_name: "DashboardWidget", foreign_key: "chartId", inverse_of: :chart, dependent: :destroy
end
