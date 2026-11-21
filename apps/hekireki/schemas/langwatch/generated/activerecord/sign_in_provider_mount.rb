# id:: Fixed, so this table holds one row: licensing's copy of sso's sign-in
#   provider fact (ADR-179).
# provider:: The configured provider's name; null when the deployment signs in by email.
# lastAppliedEventId:: The last sso event applied: the row's idempotency guard and its version.
class SignInProviderMount < ApplicationRecord
  self.table_name = "SignInProviderMount"

  attribute :id, default: "self"
  alias_attribute :updated_at, :updatedAt

  validates :mounted, inclusion: { in: [ true, false ] }
  validates :lastAppliedEventId, presence: true
end
