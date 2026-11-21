# One agent in the organisation's governance registry, from any source governance reads
# (enterprise/modules/governance/specs/agent-registry.feature). `entry` is the latest
# machine-readable snapshot; `version` grows by one on every material change.
#
# sourceKey:: Where the agent came from and its id there, e.g. `langwatch:<agentId>`.
class AgentRegistryEntry < ApplicationRecord
  self.table_name = "AgentRegistryEntry"
  self.implicit_order_column = "createdAt"

  attribute :id, default: -> { Nanoid.generate }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :organizationId, presence: true, uniqueness: { scope: :sourceKey }
  validates :sourceKey, presence: true
  validates :source, presence: true
  validates :status, presence: true
  validates :version, presence: true
  validates :fingerprint, presence: true
  validates :entry, exclusion: { in: [ nil ] }
  validates :firstSeenAt, presence: true
end
