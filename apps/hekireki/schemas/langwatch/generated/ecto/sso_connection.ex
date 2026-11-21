defmodule LangWatch.SsoConnection do
  @moduledoc """
  The SSO connection pipeline's projection (ADR-117 §5, D04 - see
  specs/identity/sso-connection-lifecycle.feature). A pure event-truth head
  like `Identifier`: every column is fold-written, rows are never deleted
  (TORN_DOWN is a tombstone), and a replay rebuilds it whole-row from the
  event log.

  The fold cursor lives on the row rather than in a sibling table, because
  one row IS one aggregate here - a connection - so there is nothing for a
  separate cursor table to key on that this row does not already key on.

  No relation to Organization on purpose, matching `Identifier`: a
  projection is rebuilt from events, and a foreign key would make a replay
  depend on relational state the log does not carry. `idpMetadata` holds
  endpoints and credential REFERENCES; no secret is ever stored here or in
  any event.
  """

  use Ecto.Schema

  @primary_key false

  @type t :: %__MODULE__{
          id: String.t(),
          organization_id: String.t(),
          type: String.t(),
          state: String.t(),
          claimed_domains: [String.t()],
          domain_claims: map(),
          approved_domains: [String.t()],
          verified_domains: [String.t()],
          lapsed_domains: [String.t()],
          domain_verifications: map(),
          pending_verification: map() | nil,
          idp_metadata: map(),
          arrival_policy: String.t(),
          allows_jit: boolean(),
          arrival_policy_decided_at: DateTime.t() | nil,
          source: String.t(),
          test_login_account_id: String.t() | nil,
          replaces_connection_id: String.t() | nil,
          migration_phase: atom() | nil,
          grace_started_at: DateTime.t() | nil,
          route_changed_at: DateTime.t() | nil,
          finalization_requested_at: DateTime.t() | nil,
          finalized_at: DateTime.t() | nil,
          rejection: map() | nil,
          created_by: String.t() | nil,
          tear_down_after: DateTime.t() | nil,
          occurred_at: DateTime.t(),
          last_event_id: String.t(),
          accepted_at: DateTime.t(),
          projection_version: String.t(),
          created_at: DateTime.t(),
          updated_at: DateTime.t(),
          reproof_cursor: LangWatch.SsoConnectionReproofCursor.t() | nil
        }

  schema "SsoConnection" do
    field(:id, :string, primary_key: true)
    field(:organization_id, :string, source: :organizationId)
    # oidc | saml
    field(:type, :string)
    # DRAFT | CLAIMED | APPROVED | REJECTED | DISCARDED |
    # VERIFICATION_PENDING | VERIFIED | ACTIVE | SUSPENDED |
    # TEARDOWN_PENDING | TORN_DOWN
    field(:state, :string)
    field(:claimed_domains, {:array, :string}, source: :claimedDomains)
    # [{ domain, state, claimedAtMs, claimedByActorId, decidedAtMs,
    # decidedByActorId, authority, waitedMs, note }] - every claim this
    # connection has made and how long each waited to be decided. The tier-3
    # approval queue is a read of this column, and the epic's Open Q2 wants
    # that latency measured from the day the queue exists.
    field(:domain_claims, :map, source: :domainClaims)
    field(:approved_domains, {:array, :string}, source: :approvedDomains)
    # The only domains that ever route, and only while state is ACTIVE.
    field(:verified_domains, {:array, :string}, source: :verifiedDomains)
    # The subset of `verifiedDomains` whose published record stayed missing
    # through its grace window (ADR-123). They still ROUTE - everyone already
    # here signs in unchanged - and they stop vouching for anybody NEW: no
    # provisioning on first sign-in, no walking in by domain. A column rather
    # than a read of the JSON below, because both questions are asked on a
    # sign-in path and have to be one indexed predicate.
    field(:lapsed_domains, {:array, :string}, source: :lapsedDomains)
    # [{ domain, method, actorId, verifiedAtMs, proofState, firstAbsentAtMs,
    # graceEndsAtMs }] - what proved each verified domain, who, and whether
    # that evidence is still there. Kept on the row rather than only in the
    # log so that an operator-attested domain can never be read as one the
    # customer proved, and a wavering one never as one nothing is doubting.
    field(:domain_verifications, :map, source: :domainVerifications)
    # { domain, method, tokenHash } while a ceremony is in flight.
    field(:pending_verification, :map, source: :pendingVerification)
    field(:idp_metadata, :map, source: :idpMetadata)
    # admit | request | refuse - what this connection does with somebody who
    # signs in through it and is not a member yet (ADR-117 3). Stated at
    # registration and changed by the setup journey, so it is never absent
    # and no reader has to decide what absence means. This was two fields
    # for a while - a boolean `allowsJit` beside it, kept in step by the
    # fold - and they disagreed about the middle answer.
    field(:arrival_policy, :string, default: "refuse", source: :arrivalPolicy)
    # DEAD, AND STAYING FOR ONE RELEASE. Nothing reads or writes this; the
    # routing repository derives the same answer from `arrivalPolicy`. It is
    # declared so the schema still describes the table, and left in place so
    # the PREVIOUS release's pods — which select it by name on every
    # connection read — keep working through the rollout. Dropping it in the
    # release that stopped reading it would stop single sign-on routing and
    # the connection projection for the length of the deploy. Dropped by
    # `20260827120000_drop_sso_connection_allows_jit` in the next one.
    field(:allows_jit, :boolean, default: false, source: :allowsJit)
    # When somebody CHOSE the policy above, and null while the registration
    # default stands. A different fact from the policy rather than a second
    # copy of it: a connection always has a behaviour, and separately
    # somebody has or has not decided on it. Going live waits for the
    # deciding - "turn everybody away" is a decision too.
    field(:arrival_policy_decided_at, LangWatch.PrismaDateTime, source: :arrivalPolicyDecidedAt)
    # self-serve | legacy-grandfathered
    field(:source, :string)
    field(:test_login_account_id, :string, source: :testLoginAccountId)
    # The grandfathered connection this direct connection replaces. Null for
    # ordinary connections and for the grandfathered predecessor itself.
    #
    # Partial uniqueness for nonterminal replacements is installed by the
    # flattened migration; Prisma cannot express that predicate here.
    field(:replaces_connection_id, :string, source: :replacesConnectionId)
    # SETUP | GRACE_LEGACY | GRACE_DIRECT | FINALIZING | FINALIZED.
    # Null means this connection is not a legacy-to-direct replacement.
    field(:migration_phase, Ecto.Enum, values: [:SETUP, :GRACE_LEGACY, :GRACE_DIRECT, :FINALIZING, :FINALIZED], source: :migrationPhase)
    field(:grace_started_at, LangWatch.PrismaDateTime, source: :graceStartedAt)
    # When an administrator last selected the normal route. This is audit
    # evidence only: routing is derived from `migrationPhase`, never recency.
    field(:route_changed_at, LangWatch.PrismaDateTime, source: :routeChangedAt)
    field(:finalization_requested_at, LangWatch.PrismaDateTime, source: :finalizationRequestedAt)
    field(:finalized_at, LangWatch.PrismaDateTime, source: :finalizedAt)
    # { domain, note } - the last ops rejection, read back on a re-claim.
    field(:rejection, :map)
    field(:created_by, :string, source: :createdBy)
    # While TEARDOWN_PENDING: when the process manager's wake completes it.
    field(:tear_down_after, LangWatch.PrismaDateTime, source: :tearDownAfter)
    field(:occurred_at, LangWatch.PrismaDateTime, source: :occurredAt)
    field(:last_event_id, :string, source: :lastEventId)
    field(:accepted_at, LangWatch.PrismaDateTime, source: :acceptedAt)
    field(:projection_version, :string, source: :projectionVersion)
    # Event-derived, not database-managed: `@default(now())` / `@updatedAt`
    # would make a replayed row differ from the row it rebuilds, which is
    # exactly the parity this projection has to keep.
    field(:created_at, LangWatch.PrismaDateTime, source: :createdAt)
    field(:updated_at, LangWatch.PrismaDateTime, source: :updatedAt)
    has_one(:reproof_cursor, LangWatch.SsoConnectionReproofCursor, foreign_key: :connection_id, references: :id)
  end
end