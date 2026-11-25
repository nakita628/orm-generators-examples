use super::sso_connection_migration_phase::SsoConnectionMigrationPhase;
use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// The SSO connection pipeline's projection (ADR-117 §5, D04 - see
/// specs/identity/sso-connection-lifecycle.feature). A pure event-truth head
/// like `Identifier`: every column is fold-written, rows are never deleted
/// (TORN_DOWN is a tombstone), and a replay rebuilds it whole-row from the
/// event log.
///
/// The fold cursor lives on the row rather than in a sibling table, because
/// one row IS one aggregate here - a connection - so there is nothing for a
/// separate cursor table to key on that this row does not already key on.
///
/// No relation to Organization on purpose, matching `Identifier`: a
/// projection is rebuilt from events, and a foreign key would make a replay
/// depend on relational state the log does not carry. `idpMetadata` holds
/// endpoints and credential REFERENCES; no secret is ever stored here or in
/// any event.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "SsoConnection")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    /// oidc | saml
    #[sea_orm(column_type = "Text")]
    pub r#type: String,
    /// DRAFT | CLAIMED | APPROVED | REJECTED | DISCARDED |
    /// VERIFICATION_PENDING | VERIFIED | ACTIVE | SUSPENDED |
    /// TEARDOWN_PENDING | TORN_DOWN
    #[sea_orm(column_type = "Text")]
    pub state: String,
    #[sea_orm(
        column_name = "claimedDomains",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub claimed_domains: Vec<String>,
    /// [{ domain, state, claimedAtMs, claimedByActorId, decidedAtMs,
    /// decidedByActorId, authority, waitedMs, note }] - every claim this
    /// connection has made and how long each waited to be decided. The tier-3
    /// approval queue is a read of this column, and the epic's Open Q2 wants
    /// that latency measured from the day the queue exists.
    #[sea_orm(
        column_name = "domainClaims",
        column_type = "JsonBinary",
        default_value = "[]"
    )]
    pub domain_claims: Json,
    #[sea_orm(
        column_name = "approvedDomains",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub approved_domains: Vec<String>,
    /// The only domains that ever route, and only while state is ACTIVE.
    #[sea_orm(
        column_name = "verifiedDomains",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub verified_domains: Vec<String>,
    /// The subset of `verifiedDomains` whose published record stayed missing
    /// through its grace window (ADR-123). They still ROUTE - everyone already
    /// here signs in unchanged - and they stop vouching for anybody NEW: no
    /// provisioning on first sign-in, no walking in by domain. A column rather
    /// than a read of the JSON below, because both questions are asked on a
    /// sign-in path and have to be one indexed predicate.
    #[sea_orm(
        column_name = "lapsedDomains",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub lapsed_domains: Vec<String>,
    /// [{ domain, method, actorId, verifiedAtMs, proofState, firstAbsentAtMs,
    /// graceEndsAtMs }] - what proved each verified domain, who, and whether
    /// that evidence is still there. Kept on the row rather than only in the
    /// log so that an operator-attested domain can never be read as one the
    /// customer proved, and a wavering one never as one nothing is doubting.
    #[sea_orm(
        column_name = "domainVerifications",
        column_type = "JsonBinary",
        default_value = "[]"
    )]
    pub domain_verifications: Json,
    /// { domain, method, tokenHash } while a ceremony is in flight.
    #[sea_orm(column_name = "pendingVerification", column_type = "JsonBinary")]
    pub pending_verification: Option<Json>,
    #[sea_orm(column_name = "idpMetadata", column_type = "JsonBinary")]
    pub idp_metadata: Json,
    /// admit | request | refuse - what this connection does with somebody who
    /// signs in through it and is not a member yet (ADR-117 3). Stated at
    /// registration and changed by the setup journey, so it is never absent
    /// and no reader has to decide what absence means. This was two fields
    /// for a while - a boolean `allowsJit` beside it, kept in step by the
    /// fold - and they disagreed about the middle answer.
    #[sea_orm(
        column_name = "arrivalPolicy",
        column_type = "Text",
        default_value = "refuse"
    )]
    pub arrival_policy: String,
    /// DEAD, AND STAYING FOR ONE RELEASE. Nothing reads or writes this; the
    /// routing repository derives the same answer from `arrivalPolicy`. It is
    /// declared so the schema still describes the table, and left in place so
    /// the PREVIOUS release's pods — which select it by name on every
    /// connection read — keep working through the rollout. Dropping it in the
    /// release that stopped reading it would stop single sign-on routing and
    /// the connection projection for the length of the deploy. Dropped by
    /// `20260827120000_drop_sso_connection_allows_jit` in the next one.
    #[sea_orm(column_name = "allowsJit", default_value = false)]
    pub allows_jit: bool,
    /// When somebody CHOSE the policy above, and null while the registration
    /// default stands. A different fact from the policy rather than a second
    /// copy of it: a connection always has a behaviour, and separately
    /// somebody has or has not decided on it. Going live waits for the
    /// deciding - "turn everybody away" is a decision too.
    #[sea_orm(
        column_name = "arrivalPolicyDecidedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub arrival_policy_decided_at: Option<DateTime>,
    /// self-serve | legacy-grandfathered
    #[sea_orm(column_type = "Text")]
    pub source: String,
    #[sea_orm(column_name = "testLoginAccountId", column_type = "Text")]
    pub test_login_account_id: Option<String>,
    /// The grandfathered connection this direct connection replaces. Null for
    /// ordinary connections and for the grandfathered predecessor itself.
    ///
    /// Partial uniqueness for nonterminal replacements is installed by the
    /// flattened migration; Prisma cannot express that predicate here.
    #[sea_orm(column_name = "replacesConnectionId", column_type = "Text")]
    pub replaces_connection_id: Option<String>,
    /// SETUP | GRACE_LEGACY | GRACE_DIRECT | FINALIZING | FINALIZED.
    /// Null means this connection is not a legacy-to-direct replacement.
    #[sea_orm(column_name = "migrationPhase")]
    pub migration_phase: Option<SsoConnectionMigrationPhase>,
    #[sea_orm(
        column_name = "graceStartedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub grace_started_at: Option<DateTime>,
    /// When an administrator last selected the normal route. This is audit
    /// evidence only: routing is derived from `migrationPhase`, never recency.
    #[sea_orm(
        column_name = "routeChangedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub route_changed_at: Option<DateTime>,
    #[sea_orm(
        column_name = "finalizationRequestedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub finalization_requested_at: Option<DateTime>,
    #[sea_orm(column_name = "finalizedAt", column_type = "custom(\"timestamp(3)\")")]
    pub finalized_at: Option<DateTime>,
    /// { domain, note } - the last ops rejection, read back on a re-claim.
    #[sea_orm(column_type = "JsonBinary")]
    pub rejection: Option<Json>,
    #[sea_orm(column_name = "createdBy", column_type = "Text")]
    pub created_by: Option<String>,
    /// While TEARDOWN_PENDING: when the process manager's wake completes it.
    #[sea_orm(
        column_name = "tearDownAfter",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub tear_down_after: Option<DateTime>,
    #[sea_orm(column_name = "occurredAt", column_type = "custom(\"timestamp(3)\")")]
    pub occurred_at: DateTime,
    #[sea_orm(column_name = "lastEventId", column_type = "Text")]
    pub last_event_id: String,
    #[sea_orm(column_name = "acceptedAt", column_type = "custom(\"timestamp(3)\")")]
    pub accepted_at: DateTime,
    #[sea_orm(column_name = "projectionVersion", column_type = "Text")]
    pub projection_version: String,
    /// Event-derived, not database-managed: `@default(now())` / `@updatedAt`
    /// would make a replayed row differ from the row it rebuilds, which is
    /// exactly the parity this projection has to keep.
    #[sea_orm(column_name = "createdAt", column_type = "custom(\"timestamp(3)\")")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(has_one = "super::sso_connection_reproof_cursor::Entity")]
    ReproofCursor,
}

impl Related<super::sso_connection_reproof_cursor::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ReproofCursor.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
