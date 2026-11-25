use super::runtime_environment_type::RuntimeEnvironmentType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Durable, typed, bidirectional I/O primitive. Owns two S2 streams (.out / .in).
/// The row is essentially static — no status, no counters, no pointers. No
/// foreign keys: project/runtimeEnvironment/organization ids are plain
/// scalar columns (matches TaskRun pattern). List-style queries are served
/// from ClickHouse, not Postgres, so only point-lookup indexes live here.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Session")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "friendlyId", column_type = "Text")]
    pub friendly_id: String,
    /// User-supplied identifier scoped to the environment. Used for
    /// idempotent upsert and for resolving sessions via the public API.
    #[sea_orm(column_name = "externalId", column_type = "Text")]
    pub external_id: Option<String>,
    /// Plain string — intentionally not an enum.
    #[sea_orm(column_type = "Text")]
    pub r#type: String,
    /// Denormalized scoping columns — no FK relations.
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_name = "runtimeEnvironmentId", column_type = "Text")]
    pub runtime_environment_id: String,
    #[sea_orm(column_name = "environmentType")]
    pub environment_type: RuntimeEnvironmentType,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    /// Task this session triggers runs against. Required — Sessions are
    /// task-bound: creating a session also triggers its first run, and
    /// every subsequent re-trigger uses this same identifier.
    #[sea_orm(column_name = "taskIdentifier", column_type = "Text")]
    pub task_identifier: String,
    /// Trigger config used for every run this session schedules. Shape
    /// (validated at the route layer, opaque to the DB):
    /// { basePayload: object, machine?: string, queue?: string,
    /// tags?: string[], maxAttempts?: number,
    /// idleTimeoutInSeconds?: number }
    /// `basePayload` carries the customer's client-data; runtime fields
    /// (chatId, messages, trigger) are merged at trigger time.
    #[sea_orm(column_name = "triggerConfig", column_type = "JsonBinary")]
    pub trigger_config: Json,
    /// Whether this session was created from the Test/playground UI rather
    /// than by a real chat.agent() trigger. Mirrors TaskRun.isTest.
    #[sea_orm(column_name = "isTest", default_value = false)]
    pub is_test: bool,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub tags: Vec<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub metadata: Option<Json>,
    /// Live run pointer — non-FK so run deletion never cascades. Can lag
    /// reality; the `.in/append` handler re-checks the snapshot status
    /// before reusing it.
    #[sea_orm(column_name = "currentRunId", column_type = "Text")]
    pub current_run_id: Option<String>,
    /// Monotonic counter used for optimistic locking on `currentRunId`
    /// swaps. Bumped atomically alongside any update that changes
    /// `currentRunId`.
    #[sea_orm(column_name = "currentRunVersion", default_value = 0)]
    pub current_run_version: i32,
    /// Terminal markers — written once, never flipped back.
    #[sea_orm(column_name = "closedAt", column_type = "custom(\"timestamp(3)\")")]
    pub closed_at: Option<DateTime>,
    #[sea_orm(column_name = "closedReason", column_type = "Text")]
    pub closed_reason: Option<String>,
    #[sea_orm(column_name = "expiresAt", column_type = "custom(\"timestamp(3)\")")]
    pub expires_at: Option<DateTime>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    /// S2 basin where this session's stream pair lives. Stamped at create
    /// time from `Organization.streamBasinName` so reads can resolve the
    /// basin without joining org. Null when the org has no per-org basin
    /// (OSS, or pre-backfill); reads fall back to the global basin.
    #[sea_orm(column_name = "streamBasinName", column_type = "Text")]
    pub stream_basin_name: Option<String>,
    /// Storage URI (with protocol prefix) for this session's chat.agent
    /// snapshot blob. Set on first snapshot write. Null = pre-column session,
    /// fall back to computed default path.
    #[sea_orm(column_name = "chatSnapshotStoragePath", column_type = "Text")]
    pub chat_snapshot_storage_path: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(has_many = "super::session_run::Entity")]
    Runs,
}

impl Related<super::session_run::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Runs.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.naive_utc());
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
