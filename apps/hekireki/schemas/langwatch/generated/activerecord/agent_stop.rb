# The `agent_stop` aggregate's Postgres projection: one row per kill switch stop, its status and
# every target's outcome, folded from the stop's events.
class AgentStop < ApplicationRecord
  self.table_name = "AgentStop"

  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :organizationId, presence: true
  validates :agentId, presence: true
  validates :agentName, presence: true
  validates :status, presence: true
  validates :severity, presence: true
  validates :approval, presence: true
  validates :trigger, presence: true
  validates :reason, presence: true
  validates :targets, exclusion: { in: [ nil ] }
  validates :requestedAt, presence: true
  validates :lastAppliedEventId, presence: true
  validates :lastAcceptedAt, presence: true
  validates :lastOccurredAt, presence: true
  validates :projectionVersion, presence: true
end
