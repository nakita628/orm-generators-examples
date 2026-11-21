# The SSO connection pipeline's projection (ADR-117 §5, D04 - see
# specs/identity/sso-connection-lifecycle.feature). A pure event-truth head
# like `Identifier`: every column is fold-written, rows are never deleted
# (TORN_DOWN is a tombstone), and a replay rebuilds it whole-row from the
# event log.
#
# The fold cursor lives on the row rather than in a sibling table, because
# one row IS one aggregate here - a connection - so there is nothing for a
# separate cursor table to key on that this row does not already key on.
#
# No relation to Organization on purpose, matching `Identifier`: a
# projection is rebuilt from events, and a foreign key would make a replay
# depend on relational state the log does not carry. `idpMetadata` holds
# endpoints and credential REFERENCES; no secret is ever stored here or in
# any event.
#
# type:: oidc | saml
# state:: DRAFT | CLAIMED | APPROVED | REJECTED | DISCARDED |
#   VERIFICATION_PENDING | VERIFIED | ACTIVE | SUSPENDED |
#   TEARDOWN_PENDING | TORN_DOWN
# domainClaims:: [{ domain, state, claimedAtMs, claimedByActorId, decidedAtMs,
#   decidedByActorId, authority, waitedMs, note }] - every claim this
#   connection has made and how long each waited to be decided. The tier-3
#   approval queue is a read of this column, and the epic's Open Q2 wants
#   that latency measured from the day the queue exists.
# verifiedDomains:: The only domains that ever route, and only while state is ACTIVE.
# lapsedDomains:: The subset of `verifiedDomains` whose published record stayed missing
#   through its grace window (ADR-123). They still ROUTE - everyone already
#   here signs in unchanged - and they stop vouching for anybody NEW: no
#   provisioning on first sign-in, no walking in by domain. A column rather
#   than a read of the JSON below, because both questions are asked on a
#   sign-in path and have to be one indexed predicate.
# domainVerifications:: [{ domain, method, actorId, verifiedAtMs, proofState, firstAbsentAtMs,
#   graceEndsAtMs }] - what proved each verified domain, who, and whether
#   that evidence is still there. Kept on the row rather than only in the
#   log so that an operator-attested domain can never be read as one the
#   customer proved, and a wavering one never as one nothing is doubting.
# pendingVerification:: { domain, method, tokenHash } while a ceremony is in flight.
# arrivalPolicy:: admit | request | refuse - what this connection does with somebody who
#   signs in through it and is not a member yet (ADR-117 3). Stated at
#   registration and changed by the setup journey, so it is never absent
#   and no reader has to decide what absence means. This was two fields
#   for a while - a boolean `allowsJit` beside it, kept in step by the
#   fold - and they disagreed about the middle answer.
# allowsJit:: DEAD, AND STAYING FOR ONE RELEASE. Nothing reads or writes this; the
#   routing repository derives the same answer from `arrivalPolicy`. It is
#   declared so the schema still describes the table, and left in place so
#   the PREVIOUS release's pods — which select it by name on every
#   connection read — keep working through the rollout. Dropping it in the
#   release that stopped reading it would stop single sign-on routing and
#   the connection projection for the length of the deploy. Dropped by
#   `20260827120000_drop_sso_connection_allows_jit` in the next one.
# arrivalPolicyDecidedAt:: When somebody CHOSE the policy above, and null while the registration
#   default stands. A different fact from the policy rather than a second
#   copy of it: a connection always has a behaviour, and separately
#   somebody has or has not decided on it. Going live waits for the
#   deciding - "turn everybody away" is a decision too.
# source:: self-serve | legacy-grandfathered
# replacesConnectionId:: The grandfathered connection this direct connection replaces. Null for
#   ordinary connections and for the grandfathered predecessor itself.
#
#   Partial uniqueness for nonterminal replacements is installed by the
#   flattened migration; Prisma cannot express that predicate here.
# migrationPhase:: SETUP | GRACE_LEGACY | GRACE_DIRECT | FINALIZING | FINALIZED.
#   Null means this connection is not a legacy-to-direct replacement.
# routeChangedAt:: When an administrator last selected the normal route. This is audit
#   evidence only: routing is derived from `migrationPhase`, never recency.
# rejection:: { domain, note } - the last ops rejection, read back on a re-claim.
# tearDownAfter:: While TEARDOWN_PENDING: when the process manager's wake completes it.
# createdAt:: Event-derived, not database-managed: `@default(now())` / `@updatedAt`
#   would make a replayed row differ from the row it rebuilds, which is
#   exactly the parity this projection has to keep.
class SsoConnection < ApplicationRecord
  self.table_name = "SsoConnection"
  self.inheritance_column = nil

  attribute :claimedDomains, default: -> { [] }
  attribute :domainClaims, default: -> { JSON.parse("[]") }
  attribute :approvedDomains, default: -> { [] }
  attribute :verifiedDomains, default: -> { [] }
  attribute :lapsedDomains, default: -> { [] }
  attribute :domainVerifications, default: -> { JSON.parse("[]") }
  attribute :arrivalPolicy, default: "refuse"
  attribute :allowsJit, default: false
  alias_attribute :created_at, :createdAt
  alias_attribute :updated_at, :updatedAt

  enum :migrationPhase, { setup: "SETUP", grace_legacy: "GRACE_LEGACY", grace_direct: "GRACE_DIRECT", finalizing: "FINALIZING", finalized: "FINALIZED" },
       validate: { allow_nil: true }

  validates :organizationId, presence: true
  validates :type, presence: true
  validates :state, presence: true
  validates :idpMetadata, exclusion: { in: [ nil ] }
  validates :source, presence: true
  validates :occurredAt, presence: true
  validates :lastEventId, presence: true
  validates :acceptedAt, presence: true
  validates :projectionVersion, presence: true

  has_one :reproof_cursor, class_name: "SsoConnectionReproofCursor", foreign_key: "connectionId",
                           inverse_of: :connection, dependent: :destroy
end
