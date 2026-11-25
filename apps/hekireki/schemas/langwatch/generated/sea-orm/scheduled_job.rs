use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "ScheduledJob")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_name = "targetType", column_type = "Text")]
    pub target_type: String,
    #[sea_orm(column_name = "targetId", column_type = "Text")]
    pub target_id: String,
    #[sea_orm(column_type = "Text")]
    pub cron: String,
    #[sea_orm(column_type = "Text")]
    pub timezone: String,
    #[sea_orm(column_name = "nextRunAt", column_type = "custom(\"timestamp(3)\")")]
    pub next_run_at: DateTime,
    #[sea_orm(column_name = "lastSlot", column_type = "custom(\"timestamp(3)\")")]
    pub last_slot: Option<DateTime>,
    #[sea_orm(column_name = "currentSlot", column_type = "custom(\"timestamp(3)\")")]
    pub current_slot: Option<DateTime>,
    #[sea_orm(default_value = 0)]
    pub attempts: i32,
    #[sea_orm(column_name = "lastError", column_type = "Text")]
    pub last_error: Option<String>,
    #[sea_orm(default_value = true)]
    pub active: bool,
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
