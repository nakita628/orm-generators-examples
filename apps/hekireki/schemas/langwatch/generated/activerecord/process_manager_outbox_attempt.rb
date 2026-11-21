# Why an outbox delivery failed, one row per FAILED attempt
# (specs/ops/dead-letter-recovery.feature). Successes write nothing, so the
# table grows with trouble, not with traffic. Rows die with their message:
# the cascade delete is the retention story, riding the existing sweep.
#
# attempt:: 1-based attempt number, as the dispatcher counts it.
# errorMessage:: The safe failure diagnostic the dispatcher already computes for its
#   span — never a raw provider body.
class ProcessManagerOutboxAttempt < ApplicationRecord
  self.table_name = "ProcessManagerOutboxAttempt"

  attribute :id, default: -> { Nanoid.generate }

  enum :outcome, { retry_scheduled: "retry_scheduled", dead: "dead" }, validate: true

  validates :projectId, presence: true
  validates :attempt, presence: true
  validates :occurredAt, presence: true
  validates :errorType, presence: true
  validates :errorMessage, presence: true

  belongs_to :outbox, class_name: "ProcessManagerOutbox", foreign_key: "outboxId", inverse_of: :attempt_log
end
