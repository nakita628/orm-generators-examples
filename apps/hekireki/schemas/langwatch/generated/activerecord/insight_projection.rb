# One insight in its owner's inbox in a project, folded from its `insight` aggregate
# (modules/insight). A read model: rebuildable by replaying the event log. Epochs in
# milliseconds.
#
# id:: The insight id, which is also its aggregate id.
# tone:: bad | watch | good
# renewedAt:: Set by a later run that found the insight still true.
# boardId:: The board and widget it came from. A pointer only, never a relation: the row reads the
#   same once either is deleted, so the names are kept as filed.
# filedVia:: chat | run
# replayStart:: The fixed window `lwql` replays over (epoch milliseconds, half-open) and its step.
# replayContext:: The values in force when it was filed: `{ period, parameters }`.
# ownerUserId:: Whose insight it is: the one person who reads it. NULL on a row folded before the
#   column existed, which is owned by whoever filed it (`filedByUserId`).
# findingKey:: The stable key a daily run matches a repeated finding by, so it renews the insight
#   and files no twin. NULL for an insight saved from a chat and on a row folded before
#   the column existed.
class InsightProjection < ApplicationRecord
  self.table_name = "InsightProjection"

  attribute :filedVia, default: "chat"

  validates :projectId, presence: true
  validates :title, presence: true
  validates :body, presence: true
  validates :tone, presence: true
  validates :validDays, presence: true
  validates :filedAt, presence: true
  validates :createdAt, presence: true
  validates :updatedAt, presence: true
  validates :occurredAt, presence: true
  validates :acceptedAt, presence: true
  validates :lastEventId, presence: true
  validates :projectionVersion, presence: true
end
