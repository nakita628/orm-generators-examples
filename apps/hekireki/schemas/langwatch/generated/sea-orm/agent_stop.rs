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
    pub id: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    #[sea_orm(column_name = "agentId")]
    pub agent_id: String,
    #[sea_orm(column_name = "agentName")]
    pub agent_name: String,
    pub status: String,
    pub severity: String,
    pub approval: String,
    pub trigger: String,
    pub reason: String,
    #[sea_orm(column_name = "requestedByUserId")]
    pub requested_by_user_id: Option<String>,
    #[sea_orm(column_name = "requestedByApiKeyId")]
    pub requested_by_api_key_id: Option<String>,
    #[sea_orm(column_name = "decidedByUserId")]
    pub decided_by_user_id: Option<String>,
    #[sea_orm(column_name = "decidedByApiKeyId")]
    pub decided_by_api_key_id: Option<String>,
    #[sea_orm(column_name = "restoredByUserId")]
    pub restored_by_user_id: Option<String>,
    #[sea_orm(column_name = "restoredByApiKeyId")]
    pub restored_by_api_key_id: Option<String>,
    pub targets: Json,
    #[sea_orm(column_name = "requestedAt")]
    pub requested_at: DateTime,
    #[sea_orm(column_name = "decidedAt")]
    pub decided_at: Option<DateTime>,
    #[sea_orm(column_name = "quarantinedAt")]
    pub quarantined_at: Option<DateTime>,
    #[sea_orm(column_name = "restoreRequestedAt")]
    pub restore_requested_at: Option<DateTime>,
    #[sea_orm(column_name = "restoredAt")]
    pub restored_at: Option<DateTime>,
    #[sea_orm(column_name = "lastAppliedEventId")]
    pub last_applied_event_id: String,
    #[sea_orm(column_name = "lastAcceptedAt")]
    pub last_accepted_at: i64,
    #[sea_orm(column_name = "lastOccurredAt")]
    pub last_occurred_at: i64,
    #[sea_orm(column_name = "projectionVersion")]
    pub projection_version: String,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
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
