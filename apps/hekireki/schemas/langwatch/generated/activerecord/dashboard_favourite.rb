# A member's star, in their own order within a project: a board or a From
# LangWatch template. Personal: one member's stars never show to another. No
# foreign key (per the database rule): the dashboard service removes a board's
# favourites when it deletes the board, so a star never outlives its dashboard.
#
# dashboardId:: A DashboardsV2 board's id (copied boards kept their Dashboard id); null when
#   the star is a template.
# templateId:: The starred From LangWatch template id; null when the star is a board.
# position:: The member's drag order among their own stars in the project.
class DashboardFavourite < ApplicationRecord
  self.table_name = "DashboardFavourite"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  alias_attribute :created_at, :createdAt

  validates :userId, presence: true, uniqueness: { scope: :dashboardId, conditions: -> { where.not(dashboardId: nil) } }
  validates :projectId, presence: true,
                        uniqueness: { scope: [ :userId, :templateId ], conditions: -> { where.not(templateId: nil) } }
  validates :position, presence: true
end
