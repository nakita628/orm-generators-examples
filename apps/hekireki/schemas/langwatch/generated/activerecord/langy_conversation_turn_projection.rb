# Rebuildable per-turn operational render document.
#
# status:: LangyConversationTurnStatus. Deliberately TEXT, not a Postgres enum: this
#   is a PROJECTION — one writer (the fold), rebuildable from the event log,
#   and nothing queries or indexes on it. An enum bought no selectivity here,
#   only typo protection, and charged a migration for every new status that
#   Postgres then could not remove. The one definition lives in
#   `LANGY_CONVERSATION_TURN_STATUS`; the repository validates against it.
# questionParts:: LangyMessagePart[] — the user question that opened the turn.
# answerParts:: LangyMessagePart[] — the agent's whole final answer.
# toolCalls:: LangyTurnToolCall[] — tool-call lifecycle audit in initiation order.
# plan:: LangyPlanItemData[] | null — null means the turn never kept a plan.
class LangyConversationTurnProjection < ApplicationRecord
  self.table_name = "LangyConversationTurnProjection"

  attribute :id, default: -> { Nanoid.generate }

  validates :projectId, presence: true, uniqueness: { scope: [ :conversationId, :turnId ] }
  validates :conversationId, presence: true
  validates :turnId, presence: true
  validates :status, presence: true
  validates :questionParts, exclusion: { in: [ nil ] }
  validates :answerParts, exclusion: { in: [ nil ] }
  validates :toolCalls, exclusion: { in: [ nil ] }
  validates :createdAt, presence: true
  validates :updatedAt, presence: true
  validates :occurredAt, presence: true
  validates :acceptedAt, presence: true
  validates :lastEventId, presence: true
  validates :projectionVersion, presence: true
end
