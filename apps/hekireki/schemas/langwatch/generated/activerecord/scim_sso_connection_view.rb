# SCIM's own copy of an organization's SSO connections: a peer fold over
# identity's connection facts (ARCHITECTURE §9), folded with identity-contract's
# reducer, one row per connection. SCIM reads this instead of asking identity.
# Rebuilt from identity's event log by a projection replay, so no relation:
# a foreign key would make a replay depend on state the log does not carry.
#
# id:: The aggregate: identity's connection id.
# folded:: The connection as the reducer folded it (SsoConnectionState + LastEventOccurredAt).
# appliedEventIds:: The event ids of the last delivery folded in, so a redelivery folds once.
# projectionVersion:: A row of another version reads as undecodable and is re-folded from the log.
# createdAt:: Event-derived, so a replayed row equals the row it rebuilds.
class ScimSsoConnectionView < ApplicationRecord
  self.table_name = "ScimSsoConnectionView"

  attribute :appliedEventIds, default: -> { [] }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :organizationId, presence: true
  validates :folded, exclusion: { in: [ nil ] }
  validates :projectionVersion, presence: true
end
