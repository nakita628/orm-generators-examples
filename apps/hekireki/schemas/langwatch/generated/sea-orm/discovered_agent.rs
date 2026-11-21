use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// The same shape as DiscoveredPerson for non-human actors - Genie spaces,
/// Copilot bots, OpenAI projects - plus whatever provider-native metadata the
/// source published. Discovering an agent grants it nothing either.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "DiscoveredAgent")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    pub provider: String,
    /// Genie space / Copilot bot / OpenAI project id, verbatim.
    #[sea_orm(column_name = "rawAgentId")]
    pub raw_agent_id: String,
    #[sea_orm(column_name = "displayText")]
    pub display_text: String,
    /// Workspace, plan, and kin - provider-native descriptive fields only. Never
    /// credentials: this row is read by admin screens, not by any pull path.
    pub metadata: Json,
    #[sea_orm(column_name = "firstSeenAt")]
    pub first_seen_at: DateTime,
    #[sea_orm(column_name = "lastSeenAt")]
    pub last_seen_at: DateTime,
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
