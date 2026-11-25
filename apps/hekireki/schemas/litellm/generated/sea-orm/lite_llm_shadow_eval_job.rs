use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LiteLLM_ShadowEvalJob")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub group_id: String,
    #[sea_orm(column_type = "Text", default_value = "key")]
    pub target_type: String,
    #[sea_orm(column_type = "Text")]
    pub target_id: String,
    #[sea_orm(column_type = "Text")]
    pub router_name: String,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub router_names: Vec<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub models: Vec<String>,
    #[sea_orm(column_type = "Text", default_value = "forward")]
    pub direction: String,
    #[sea_orm(column_type = "Text")]
    pub baseline_model: Option<String>,
    #[sea_orm(default_value = false)]
    pub skip_same_model: bool,
    #[sea_orm(column_type = "Text")]
    pub judge_model: String,
    pub shadow_percentage: f64,
    pub max_turns: i32,
    pub max_budget: Option<f64>,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub created_by: Option<String>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub ends_at: DateTime,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub stopped_at: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub stopped_by: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub rubric: Option<String>,
    pub rubric_version: Option<i32>,
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
