# projectId:: The project that owns the board; only there can it be edited.
# order:: Legacy board sort position. The Dashboards v2 app orders boards by a
#   member's favourites and never writes this; kept because the legacy
#   Analytics reports section and the REST reorder endpoint still write it,
#   and a rolling deploy's older image still reads it.
# description:: Free text under the board's name; null until a member writes one.
#   Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
# createdById:: The member who created it; null for boards that predate the column.
#   Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
# scope:: Who sees it. Only the author changes it; a board with no author stays PROJECT.
#   Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
# organizationId:: The owning project's organization, stamped when the board is first set to
#   ORGANIZATION; null before. No relation, per the database rule.
#   Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
class Dashboard < ApplicationRecord
  self.table_name = "Dashboard"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  attribute :order, default: 0
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :scope, { private: "PRIVATE", project: "PROJECT", organization: "ORGANIZATION" }, default: :project, prefix: true,
                                                                                         validate: true

  validates :name, presence: true

  belongs_to :project, foreign_key: "projectId", inverse_of: :dashboards
  has_many :graphs, class_name: "CustomGraph", foreign_key: "dashboardId", inverse_of: :dashboard, dependent: :destroy
end
