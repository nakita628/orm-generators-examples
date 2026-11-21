# The registry's append-only change feed: one row per registered, versioned, removed or
# tool-deprecated fact. `sequence` is the pull cursor; `deliveryRecordedAt` marks the webhook
# fan-out as recorded in governance's outbox.
class AgentRegistryEvent < ApplicationRecord
  self.table_name = "AgentRegistryEvent"
  self.inheritance_column = nil

  attribute :id, default: -> { Nanoid.generate }

  validates :organizationId, presence: true, uniqueness: { scope: :sequence }
  validates :type, presence: true
  validates :entryId, presence: true
  validates :payload, exclusion: { in: [ nil ] }
  validates :occurredAt, presence: true
end
