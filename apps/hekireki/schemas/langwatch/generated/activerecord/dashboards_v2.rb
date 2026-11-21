# A Dashboards v2 board, apart from the Analytics Reports boards in Dashboard. Boards
# made before the split were copied here with their ids, so stars and insight
# pointers still name them. No relation, per the database rule.
#
# projectId:: The project that owns the board; only there can it be edited.
# description:: Free text under the board's name; null until a member writes one.
# createdById:: The member who created it; null for boards that predate authorship.
# scope:: Who sees it. Only the author changes it; a board with no author stays PROJECT.
# organizationId:: The owning project's organization, stamped when the board is first set to
#   ORGANIZATION; null before.
# excludedOrigins:: Deprecated: unused since each widget query names its own sources (dashboards-v2 AC200,
#   2026-10-28); dropped in a later migration, once no serving release reads it.
class DashboardsV2 < ApplicationRecord
  self.table_name = "DashboardsV2"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :scope, { private: "PRIVATE", project: "PROJECT", organization: "ORGANIZATION" }, default: :project, prefix: true,
                                                                                         validate: true

  validates :projectId, presence: true
  validates :name, presence: true
end
