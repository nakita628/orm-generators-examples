# One person's daily insights run on one board in a project, folded from its
# `insight_daily_schedule` aggregate (modules/insight). A read model: rebuildable by
# replaying the event log. Epochs in milliseconds.
#
# id:: The schedule id, which is also its aggregate id: one per project, person and board.
# boardKind:: dashboard | template. A pointer only, never a relation: a run checks the board is there.
# state:: on | off | undecided.
# offReason:: board_deleted | skipped_days | member_removed | project_archived. Why the system turned
#   it off. NULL while it is on, and when the person turned it off.
# lastRunId:: The last run that recorded an outcome. NULL until one did.
# lastRunOutcome:: filed | nothing | failed | skipped
# lastRunRenewed:: How many insights the run found still true and renewed.
# consecutiveSkips:: The scheduled days in a row whose run ended skipped, and the slot of the newest counted.
class InsightDailyScheduleProjection < ApplicationRecord
  self.table_name = "InsightDailyScheduleProjection"

  attribute :state, default: "undecided"
  attribute :consecutiveSkips, default: 0

  validates :projectId, presence: true, uniqueness: { scope: [ :userId, :boardKind, :boardId ] }
  validates :userId, presence: true
  validates :boardKind, presence: true
  validates :boardId, presence: true
  validates :boardName, presence: true
  validates :createdAt, presence: true
  validates :updatedAt, presence: true
  validates :occurredAt, presence: true
  validates :acceptedAt, presence: true
  validates :lastEventId, presence: true
  validates :projectionVersion, presence: true
end
