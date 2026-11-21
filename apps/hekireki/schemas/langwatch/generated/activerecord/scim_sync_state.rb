# D08 — one connection's directory sync, as the `scim-sync` pipeline's fold
# writes it.
#
# Pure event truth: every column is fold-written, rows are never deleted,
# and REVOKED is a tombstone rather than an absence — a dead letter that
# vanished when a connection was torn down would be a removal nobody could
# check afterwards.
#
# No relation to Organization or SsoConnection, matching `SsoConnection`
# itself: a projection is rebuilt from events, and a foreign key would make
# a replay depend on relational state the log does not carry.
#
# id:: The aggregate: the connection's id, which IS its sync's id.
# state:: TOKEN_ISSUED | SYNCING | ERROR | REVOKED
# lastFailure:: The failure standing right now: { op, errorCode, attempts, retiredAtMs,
#   userId, occurredAtMs }. A reason CODE and ids — never a token, never a
#   secret, never an internal hostname.
# deadLetters:: Every failure retired without ever being applied, newest last. The dead
#   letters an administrator has to see; retiring never drops one.
# revokedCause:: revoke | teardown, once the sync has ended.
# createdAt:: Event-derived, not database-managed: `@default(now())` / `@updatedAt`
#   would make a replayed row differ from the row it rebuilds, which is
#   exactly the parity this projection has to keep.
class ScimSyncState < ApplicationRecord
  self.table_name = "ScimSyncState"

  attribute :deadLetters, default: -> { JSON.parse("[]") }
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  validates :connectionId, presence: true
  validates :organizationId, presence: true
  validates :state, presence: true
  validates :occurredAt, presence: true
  validates :lastEventId, presence: true
  validates :acceptedAt, presence: true
  validates :projectionVersion, presence: true
end
