use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// The `agent_stop` aggregate's Postgres projection: one row per kill switch stop, its status and
/// every target's outcome, folded from the stop's events.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AgentStop")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_name = "agentId", column_type = "Text")]
    pub agent_id: String,
    #[sea_orm(column_name = "agentName", column_type = "Text")]
    pub agent_name: String,
    #[sea_orm(column_type = "Text")]
    pub status: String,
    #[sea_orm(column_type = "Text")]
    pub severity: String,
    #[sea_orm(column_type = "Text")]
    pub approval: String,
    #[sea_orm(column_type = "Text")]
    pub trigger: String,
    #[sea_orm(column_type = "Text")]
    pub reason: String,
    #[sea_orm(column_name = "requestedByUserId", column_type = "Text")]
    pub requested_by_user_id: Option<String>,
    #[sea_orm(column_name = "requestedByApiKeyId", column_type = "Text")]
    pub requested_by_api_key_id: Option<String>,
    #[sea_orm(column_name = "decidedByUserId", column_type = "Text")]
    pub decided_by_user_id: Option<String>,
    #[sea_orm(column_name = "decidedByApiKeyId", column_type = "Text")]
    pub decided_by_api_key_id: Option<String>,
    #[sea_orm(column_name = "restoredByUserId", column_type = "Text")]
    pub restored_by_user_id: Option<String>,
    #[sea_orm(column_name = "restoredByApiKeyId", column_type = "Text")]
    pub restored_by_api_key_id: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub targets: Json,
    #[sea_orm(column_name = "requestedAt", column_type = "custom(\"timestamp(3)\")")]
    pub requested_at: DateTime,
    #[sea_orm(column_name = "decidedAt", column_type = "custom(\"timestamp(3)\")")]
    pub decided_at: Option<DateTime>,
    #[sea_orm(
        column_name = "quarantinedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub quarantined_at: Option<DateTime>,
    #[sea_orm(
        column_name = "restoreRequestedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub restore_requested_at: Option<DateTime>,
    #[sea_orm(column_name = "restoredAt", column_type = "custom(\"timestamp(3)\")")]
    pub restored_at: Option<DateTime>,
    #[sea_orm(column_name = "lastAppliedEventId", column_type = "Text")]
    pub last_applied_event_id: String,
    #[sea_orm(column_name = "lastAcceptedAt")]
    pub last_accepted_at: i64,
    #[sea_orm(column_name = "lastOccurredAt")]
    pub last_occurred_at: i64,
    #[sea_orm(column_name = "projectionVersion", column_type = "Text")]
    pub projection_version: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(
        column_name = "updatedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

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
