# Whether an organization's permission checks run on the authorization engine: authz's own
# switch, set when the engine migration proves the organization and pinned off by an operator's
# rollback. No row is the legacy path. Spec: modules/authz/specs/engine-switch.feature.
class AuthzEngineSwitch < ApplicationRecord
  self.table_name = "AuthzEngineSwitch"
  self.primary_key = "organizationId"

  attribute :pinned, default: false
  alias_attribute :updated_at, :updatedAt

  validates :onEngine, inclusion: { in: [ true, false ] }
  validates :switchedAt, presence: true
end
