use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LiteLLM_HealthCheckTable")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub health_check_id: String,
    #[sea_orm(column_type = "Text")]
    pub model_name: String,
    #[sea_orm(column_type = "Text")]
    pub model_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub status: String,
    #[sea_orm(default_value = 0)]
    pub healthy_count: i32,
    #[sea_orm(default_value = 0)]
    pub unhealthy_count: i32,
    #[sea_orm(column_type = "Text")]
    pub error_message: Option<String>,
    pub response_time_ms: Option<f64>,
    #[sea_orm(column_type = "JsonBinary")]
    pub details: Option<Json>,
    #[sea_orm(column_type = "Text")]
    pub checked_by: Option<String>,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub checked_at: DateTime,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            health_check_id: Set(uuid::Uuid::new_v4().to_string()),
            ..ActiveModelTrait::default()
        }
    }

    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.checked_at.is_not_set() {
            self.checked_at = Set(now.naive_utc());
        }
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.naive_utc());
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
