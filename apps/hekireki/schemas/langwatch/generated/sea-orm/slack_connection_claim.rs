use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// An automation holding a Slack connection in use (ARCHITECTURE.md §3, the
/// claims ruling): a claimed connection is never deleted. `projectId` is the
/// claimant's project; `organizationId` is the connection's tenancy anchor.
/// No foreign key: slack's service refuses to delete a claimed connection.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "slack_connection_claim")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "connectionId")]
    pub connection_id: String,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "claimantId")]
    pub claimant_id: String,
    #[sea_orm(column_name = "claimantLabel")]
    pub claimant_label: String,
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    #[sea_orm(column_name = "projectId")]
    pub project_id: String,
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
