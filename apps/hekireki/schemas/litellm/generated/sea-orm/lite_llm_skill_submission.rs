use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LiteLLM_SkillSubmission")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub plugin_name: String,
    #[sea_orm(column_type = "Text")]
    pub operation: String,
    #[sea_orm(column_type = "Text", default_value = "pending_review")]
    pub status: String,
    #[sea_orm(column_type = "Text")]
    pub proposal_json: String,
    pub archive_content: Option<Vec<u8>>,
    #[sea_orm(column_type = "Text")]
    pub submitted_by: String,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub base_plugin_id: Option<String>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub base_updated_at: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub reviewed_by: Option<String>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub reviewed_at: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub rejection_reason: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub scan_json: Option<String>,
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
        Ok(self)
    }
}
