use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "GatewayA2ATarget")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_name = "agentName", column_type = "Text")]
    pub agent_name: String,
    #[sea_orm(column_name = "cardUrl", column_type = "Text")]
    pub card_url: String,
    #[sea_orm(column_name = "upstreamUrl", column_type = "Text")]
    pub upstream_url: String,
    #[sea_orm(column_name = "encryptedUpstreamHeaders", column_type = "Text")]
    pub encrypted_upstream_headers: Option<String>,
    #[sea_orm(column_name = "restrictToAllowList", default_value = false)]
    pub restrict_to_allow_list: bool,
    #[sea_orm(
        column_name = "allowedVirtualKeyIds",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub allowed_virtual_key_ids: Vec<String>,
    #[sea_orm(default_value = false)]
    pub blocked: bool,
    #[sea_orm(column_name = "blockReason", column_type = "Text")]
    pub block_reason: Option<String>,
    #[sea_orm(column_name = "createdById", column_type = "Text")]
    pub created_by_id: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "archivedAt", column_type = "custom(\"timestamp(3)\")")]
    pub archived_at: Option<DateTime>,
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
