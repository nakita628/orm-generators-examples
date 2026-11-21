use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// D08 — one connection's directory sync, as the `scim-sync` pipeline's fold
/// writes it.
///
/// Pure event truth: every column is fold-written, rows are never deleted,
/// and REVOKED is a tombstone rather than an absence — a dead letter that
/// vanished when a connection was torn down would be a removal nobody could
/// check afterwards.
///
/// No relation to Organization or SsoConnection, matching `SsoConnection`
/// itself: a projection is rebuilt from events, and a foreign key would make
/// a replay depend on relational state the log does not carry.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ScimSyncState")]
pub struct Model {
    /// The aggregate: the connection's id, which IS its sync's id.
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "connectionId")]
    pub connection_id: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    /// TOKEN_ISSUED | SYNCING | ERROR | REVOKED
    pub state: String,
    #[sea_orm(column_name = "lastPushedAt")]
    pub last_pushed_at: Option<DateTime>,
    /// The failure standing right now: { op, errorCode, attempts, retiredAtMs,
    /// userId, occurredAtMs }. A reason CODE and ids — never a token, never a
    /// secret, never an internal hostname.
    #[sea_orm(column_name = "lastFailure")]
    pub last_failure: Option<Json>,
    /// Every failure retired without ever being applied, newest last. The dead
    /// letters an administrator has to see; retiring never drops one.
    #[sea_orm(column_name = "deadLetters", default_value = "[]")]
    pub dead_letters: Json,
    /// revoke | teardown, once the sync has ended.
    #[sea_orm(column_name = "revokedCause")]
    pub revoked_cause: Option<String>,
    #[sea_orm(column_name = "occurredAt")]
    pub occurred_at: DateTime,
    #[sea_orm(column_name = "lastEventId")]
    pub last_event_id: String,
    #[sea_orm(column_name = "acceptedAt")]
    pub accepted_at: DateTime,
    #[sea_orm(column_name = "projectionVersion")]
    pub projection_version: String,
    /// Event-derived, not database-managed: `@default(now())` / `@updatedAt`
    /// would make a replayed row differ from the row it rebuilds, which is
    /// exactly the parity this projection has to keep.
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
