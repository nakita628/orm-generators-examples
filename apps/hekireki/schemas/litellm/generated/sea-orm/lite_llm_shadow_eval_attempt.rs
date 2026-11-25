use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LiteLLM_ShadowEvalAttempt")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub job_id: String,
    #[sea_orm(column_type = "Text")]
    pub request_id: String,
    #[sea_orm(column_type = "Text")]
    pub outcome: String,
    #[sea_orm(column_type = "Text")]
    pub router_name: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub tier: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub real_model: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub shadow_model: Option<String>,
    pub confidence: Option<f64>,
    #[sea_orm(default_value = 0)]
    pub judge_cost: f64,
    #[sea_orm(default_value = 0)]
    pub shadow_cost: f64,
    pub real_cost: Option<f64>,
    #[sea_orm(default_value = 0)]
    pub real_classifier_cost: f64,
    #[sea_orm(default_value = 0)]
    pub shadow_classifier_cost: f64,
    #[sea_orm(default_value = false)]
    pub real_cache_hit: bool,
    pub real_grade: Option<i32>,
    pub shadow_grade: Option<i32>,
    #[sea_orm(column_type = "Text")]
    pub error: Option<String>,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
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
        Ok(self)
    }
}
