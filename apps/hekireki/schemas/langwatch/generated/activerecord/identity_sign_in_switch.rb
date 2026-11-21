# Whether a user signs in on the identifier path: identity's own switch, set when the
# identifier backfill proves the user and pinned off by an operator's rollback. No row is the
# legacy path. Spec: modules/identity/specs/sign-in-switch.feature. Column refs only, no FK.
class IdentitySignInSwitch < ApplicationRecord
  self.table_name = "IdentitySignInSwitch"
  self.primary_key = "userId"

  attribute :pinned, default: false
  alias_attribute :updated_at, :updatedAt

  validates :onIdentifiers, inclusion: { in: [ true, false ] }
  validates :switchedAt, presence: true
end
