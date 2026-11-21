# The join-request pipeline's projection (ADR-117, D12 - see
# specs/identity/join-requests.feature). A pure event-truth head like
# `Identifier` and `SsoConnection`: every column is fold-written, rows are
# never deleted (a terminal state is a tombstone the admin panel stops
# listing, not an absence), and a replay rebuilds it whole-row from the
# event log.
#
# The fold cursor lives on the row rather than in a sibling table, because
# one row IS one aggregate here - a request.
#
# No relation to Organization or User on purpose, matching the two heads
# above: a projection is rebuilt from events, and a foreign key would make a
# replay depend on relational state the log does not carry.
#
# Membership is NOT here. An approval dispatches an attach on the grants
# ledger, exactly as accepting an invitation does, and that ledger fact
# carries `source: "join-request"` as its provenance - which is what puts a
# join approval on the customer's audit page beside every other membership
# change rather than in a history of its own.
#
# The requester's address is not here either: the DOMAIN is the fact, and it
# is the only part of an address a join decision may depend on.
#
# domain:: The normalized domain the match was made on.
# state:: PENDING | APPROVED | REJECTED | EXPIRED | WITHDRAWN
# matchedVia:: How the domain was matched: `verified-identifier-domain` or `sso-connection-domain`.
# origin:: Where the request was made: `web` or `cli` (ADR-171 v6). A request from
#   the terminal lands a Developer seat whatever the joiner seat says.
# expiresAt:: While PENDING: when the day-14 wake lapses it. Null once resolved.
# resolvedByType:: user | policy | invite - who ended it. `policy` is domain auto-join's
#   own principal, deliberately not the one SSO auto-join uses.
# resolvedById:: The admin's user id, `domain-auto`, or the invitation id.
# withdrawalCause:: user | invite-accepted, on a withdrawal.
# connectionId:: The SSO connection a single sign-on arrival was admitted through; null otherwise.
# createdAt:: Event-derived, not database-managed: `@default(now())` / `@updatedAt`
#   would make a replayed row differ from the row it rebuilds, which is
#   exactly the parity this projection has to keep.
class JoinRequest < ApplicationRecord
  self.table_name = "JoinRequest"

  attribute :origin, default: "web"
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :userId, presence: true
  validates :organizationId, presence: true
  validates :domain, presence: true
  validates :state, presence: true
  validates :matchedVia, presence: true
  validates :occurredAt, presence: true
  validates :lastEventId, presence: true
  validates :acceptedAt, presence: true
  validates :projectionVersion, presence: true
end
