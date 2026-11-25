use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LiteLLM_ManagedObjectTable")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub unified_object_id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub model_object_id: String,
    #[sea_orm(column_type = "JsonBinary")]
    pub file_object: Json,
    #[sea_orm(column_type = "Text")]
    pub file_purpose: String,
    #[sea_orm(column_type = "Text")]
    pub status: Option<String>,
    #[sea_orm(default_value = false)]
    pub batch_processed: bool,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub created_by: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub team_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub org_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub api_key: Option<String>,
    #[sea_orm(column_type = "JsonBinary", default_value = "[]")]
    pub request_tags: Option<Json>,
    #[sea_orm(column_type = "Text")]
    pub end_user_id: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub creator_billing_context: Option<Json>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
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
            id: Set(uuid::Uuid::new_v4().to_string()),
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
