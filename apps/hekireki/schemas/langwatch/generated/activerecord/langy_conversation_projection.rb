# Rebuildable operational state folded from Langy's canonical event stream.
# The source of truth remains the ClickHouse event log.
#
# userId:: Owner; first-writer-wins from the first message's userId.
# titleSource:: derived | auto | user — governs auto-title precedence (user is sticky).
# origin:: Who started the conversation: `interactive` for the person in the panel, `run` for a
#   module acting for them with nobody at the keyboard. A row folded before the column
#   existed reads as interactive.
# currentTurnId:: The turn in flight, or null when idle (busy guard reads this).
# lastModel:: The model the latest accepted turn ran on; reopening seeds the picker from it.
# pendingHandoffToken:: ADR-048 shutdown-handoff resume token (opaque, worker-authored).
# runToken:: Per-conversation HMAC key for worker stream frames; server-side only.
# occurredAt:: Business time of the latest applied event.
# acceptedAt:: Event-log acceptance time. Together with lastEventId this is the cursor.
class LangyConversationProjection < ApplicationRecord
  self.table_name = "LangyConversationProjection"

  attribute :id, default: -> { Nanoid.generate }
  attribute :origin, default: "interactive"
  attribute :isShared, default: false
  attribute :messageCount, default: 0

  enum :titleSource, { derived: "derived", auto: "auto", user: "user" }, validate: true

  validates :projectId, presence: true, uniqueness: { scope: :conversationId }
  validates :conversationId, presence: true
  validates :userId, presence: true
  validates :status, presence: true
  validates :createdAt, presence: true
  validates :updatedAt, presence: true
  validates :occurredAt, presence: true
  validates :acceptedAt, presence: true
  validates :lastEventId, presence: true
  validates :projectionVersion, presence: true
end
