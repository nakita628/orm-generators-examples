use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LiteLLM_AutoRouterBaselineComparison")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub scope: String,
    #[sea_orm(column_type = "Text")]
    pub api_key: String,
    #[sea_orm(column_type = "Text")]
    pub session_id: String,
    #[sea_orm(column_type = "Text")]
    pub router_name: String,
    pub initial_equivalent: bool,
    #[sea_orm(default_value = "0")]
    pub revision: i64,
    #[sea_orm(default_value = "0")]
    pub published_revision: i64,
    #[sea_orm(column_type = "Text")]
    pub history: Option<String>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub attempted_at: Option<DateTime>,
    #[sea_orm(default_value = false)]
    pub retired: bool,
    #[sea_orm(
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
        if insert && self.updated_at.is_not_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
