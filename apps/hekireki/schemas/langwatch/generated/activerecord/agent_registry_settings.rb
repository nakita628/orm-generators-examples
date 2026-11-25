# An organisation's agent registry settings. No row reads as the defaults.
#
# dormantAfterDays:: Days without traffic after which a listed agent reads as dormant.
class AgentRegistrySettings < ApplicationRecord
  self.table_name = "AgentRegistrySettings"
  self.primary_key = "organizationId"

  attribute :dormantAfterDays, default: 30
  alias_attribute :updated_at, :updatedAt
end
