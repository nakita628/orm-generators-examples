use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// One agent in the organisation's governance registry, from any source governance reads
/// (enterprise/modules/governance/specs/agent-registry.feature). `entry` is the latest
/// machine-readable snapshot; `version` grows by one on every material change.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AgentRegistryEntry")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    /// Where the agent came from and its id there, e.g. `langwatch:<agentId>`.
    #[sea_orm(column_name = "sourceKey")]
    pub source_key: String,
    pub source: String,
    pub status: String,
    pub version: i32,
    pub fingerprint: String,
    pub entry: Json,
    #[sea_orm(column_name = "firstSeenAt")]
    pub first_seen_at: DateTime,
    #[sea_orm(column_name = "lastSeenAt")]
    pub last_seen_at: Option<DateTime>,
    #[sea_orm(column_name = "removedAt")]
    pub removed_at: Option<DateTime>,
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
