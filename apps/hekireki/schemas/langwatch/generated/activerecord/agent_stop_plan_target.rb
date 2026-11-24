# One target of a registry agent's kill switch stop plan (enterprise/modules/governance/specs/
# agent-kill-switch.feature): a Databricks serving endpoint, a Bedrock agent alias or a gateway
# virtual key. A plan write replaces every row of the agent at once.
#
# agentId:: The AgentRegistryEntry the plan belongs to.
class AgentStopPlanTarget < ApplicationRecord
  self.table_name = "AgentStopPlanTarget"

  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :organizationId, presence: true
  validates :agentId, presence: true
  validates :position, presence: true
  validates :label, presence: true
  validates :target, exclusion: { in: [ nil ] }
end
