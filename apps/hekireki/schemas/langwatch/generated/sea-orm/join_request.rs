use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// The join-request pipeline's projection (ADR-117, D12 - see
/// specs/identity/join-requests.feature). A pure event-truth head like
/// `Identifier` and `SsoConnection`: every column is fold-written, rows are
/// never deleted (a terminal state is a tombstone the admin panel stops
/// listing, not an absence), and a replay rebuilds it whole-row from the
/// event log.
///
/// The fold cursor lives on the row rather than in a sibling table, because
/// one row IS one aggregate here - a request.
///
/// No relation to Organization or User on purpose, matching the two heads
/// above: a projection is rebuilt from events, and a foreign key would make a
/// replay depend on relational state the log does not carry.
///
/// Membership is NOT here. An approval dispatches an attach on the grants
/// ledger, exactly as accepting an invitation does, and that ledger fact
/// carries `source: "join-request"` as its provenance - which is what puts a
/// join approval on the customer's audit page beside every other membership
/// change rather than in a history of its own.
///
/// The requester's address is not here either: the DOMAIN is the fact, and it
/// is the only part of an address a join decision may depend on.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "JoinRequest")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    /// The normalized domain the match was made on.
    #[sea_orm(column_type = "Text")]
    pub domain: String,
    /// PENDING | APPROVED | REJECTED | EXPIRED | WITHDRAWN
    #[sea_orm(column_type = "Text")]
    pub state: String,
    /// How the domain was matched: `verified-identifier-domain` or `sso-connection-domain`.
    #[sea_orm(column_name = "matchedVia", column_type = "Text")]
    pub matched_via: String,
    /// Where the request was made: `web` or `cli` (ADR-171 v6). A request from
    /// the terminal lands a Developer seat whatever the joiner seat says.
    #[sea_orm(column_type = "Text", default_value = "web")]
    pub origin: String,
    /// While PENDING: when the day-14 wake lapses it. Null once resolved.
    #[sea_orm(column_name = "expiresAt", column_type = "custom(\"timestamp(3)\")")]
    pub expires_at: Option<DateTime>,
    #[sea_orm(column_name = "resolvedAt", column_type = "custom(\"timestamp(3)\")")]
    pub resolved_at: Option<DateTime>,
    /// user | policy | invite - who ended it. `policy` is domain auto-join's
    /// own principal, deliberately not the one SSO auto-join uses.
    #[sea_orm(column_name = "resolvedByType", column_type = "Text")]
    pub resolved_by_type: Option<String>,
    /// The admin's user id, `domain-auto`, or the invitation id.
    #[sea_orm(column_name = "resolvedById", column_type = "Text")]
    pub resolved_by_id: Option<String>,
    /// user | invite-accepted, on a withdrawal.
    #[sea_orm(column_name = "withdrawalCause", column_type = "Text")]
    pub withdrawal_cause: Option<String>,
    /// The SSO connection a single sign-on arrival was admitted through; null otherwise.
    #[sea_orm(column_name = "connectionId", column_type = "Text")]
    pub connection_id: Option<String>,
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
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
