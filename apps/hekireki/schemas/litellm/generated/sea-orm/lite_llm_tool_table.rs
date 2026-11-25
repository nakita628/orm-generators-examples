use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LiteLLM_ToolTable")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub tool_id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub tool_name: String,
    #[sea_orm(column_type = "Text")]
    pub origin: Option<String>,
    #[sea_orm(column_type = "Text", default_value = "untrusted")]
    pub input_policy: String,
    #[sea_orm(column_type = "Text", default_value = "untrusted")]
    pub output_policy: String,
    #[sea_orm(default_value = 0)]
    pub call_count: i32,
    #[sea_orm(column_type = "JsonBinary", default_value = "{}")]
    pub assignments: Option<Json>,
    #[sea_orm(column_type = "Text")]
    pub key_hash: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub team_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub key_alias: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub user_agent: Option<String>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub last_used_at: Option<DateTime>,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub created_by: Option<String>,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub updated_by: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            tool_id: Set(uuid::Uuid::new_v4().to_string()),
            ..ActiveModelTrait::default()
        }
    }

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
