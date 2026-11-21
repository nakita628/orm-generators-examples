# An organisation's registry baseline: its first reconcile, written as already pushed so the
# agents it already had are not announced. Every reconcile after it pushes its changes.
class AgentRegistryBaseline < ApplicationRecord
  self.table_name = "AgentRegistryBaseline"
  self.primary_key = "organizationId"

  validates :baselinedAt, presence: true
end
