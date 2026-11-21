# A way back in (D05 - see specs/identity/sso-onboarding-tiers.feature).
#
# One named person who may still sign in without the identity provider
# while their organization's sign-in belongs to one, and the date on which
# that stops. Activation asks whether such a person exists, because the
# failure it prevents is the only one that cannot be recovered from inside
# the product: a misconfigured connection goes live, every member is handed
# to an identity provider that will not authenticate them, and nobody is
# left who can turn it off.
#
# Rows are immutable except for `supersededAt` and `warnedDays`. A renewal
# writes a NEW row naming the one it replaced, so the date a way in
# previously ended is still readable in the history - which is what makes a
# renewal auditable rather than a column somebody moved.
#
# No relation to Organization or User, matching the identity heads above:
# the binding is the record of who could get in, and a cascade would erase
# exactly the history a lockout post-mortem reads.
#
# userId:: Who holds the way in.
# grantedByUserId:: Who granted it. A way back in is never self-served.
# supersededAt:: Set when a renewal replaced this row; the row itself is never edited.
# renewedFromId:: The binding this one renewed, so the previous end date stays readable.
# warnedDays:: Which of the fourteen, seven and one day warnings have been sent, so a
#   sweep that runs twice in a day never says the same thing twice.
class SsoBreakGlassBinding < ApplicationRecord
  self.table_name = "SsoBreakGlassBinding"

  attribute :warnedDays, default: -> { [] }
  alias_attribute :created_at, :createdAt

  validates :organizationId, presence: true
  validates :userId, presence: true
  validates :grantedByUserId, presence: true
  validates :grantedAt, presence: true
  validates :expiresAt, presence: true
end
