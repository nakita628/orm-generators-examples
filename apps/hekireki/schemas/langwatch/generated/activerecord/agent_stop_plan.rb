# One agent's stop plan version: bumped by every save, so a save over a plan someone else saved
# since it was read is refused. No row means the plan was never saved (version 0).
#
# agentId:: The AgentRegistryEntry the plan belongs to.
class AgentStopPlan < ApplicationRecord
  self.table_name = "AgentStopPlan"
  self.primary_key = %w[organizationId agentId]

  alias_attribute :updated_at, :updatedAt

  validates :version, presence: true
end
