use super::audit_log_record_type::AuditLogRecordType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "audit_logs")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
    #[sea_orm(default_value = "USER")]
    pub r#type: AuditLogRecordType,
    #[sea_orm(column_type = "Text")]
    pub api_key_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub user_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub org_id: String,
    #[sea_orm(column_type = "Text")]
    pub user_org_role: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub project_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub user_project_role: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub resource_type: String,
    #[sea_orm(column_type = "Text")]
    pub resource_id: String,
    #[sea_orm(column_type = "Text")]
    pub action: String,
    #[sea_orm(column_type = "Text")]
    pub before: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub after: Option<String>,
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
