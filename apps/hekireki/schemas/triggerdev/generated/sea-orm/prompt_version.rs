use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "prompt_versions")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    pub version: i32,
    #[sea_orm(column_name = "promptId", column_type = "Text")]
    pub prompt_id: String,
    #[sea_orm(column_name = "textContent", column_type = "Text")]
    pub text_content: Option<String>,
    #[sea_orm(column_name = "chatContent", column_type = "JsonBinary")]
    pub chat_content: Option<Json>,
    #[sea_orm(column_type = "Text")]
    pub model: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub config: Option<Json>,
    #[sea_orm(column_type = "Text")]
    pub source: String,
    #[sea_orm(column_name = "commitMessage", column_type = "Text")]
    pub commit_message: Option<String>,
    #[sea_orm(column_name = "contentHash", column_type = "Text")]
    pub content_hash: String,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub labels: Vec<String>,
    #[sea_orm(column_name = "workerId", column_type = "Text")]
    pub worker_id: Option<String>,
    #[sea_orm(column_name = "createdBy", column_type = "Text")]
    pub created_by: Option<String>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::prompt::Entity",
        from = "Column::PromptId",
        to = "super::prompt::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "prompt_versions_promptId_fkey"
    )]
    Prompt,
    #[sea_orm(
        belongs_to = "super::background_worker::Entity",
        from = "Column::WorkerId",
        to = "super::background_worker::Column::Id",
        fk_name = "prompt_versions_workerId_fkey"
    )]
    Worker,
}

impl Related<super::prompt::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Prompt.def()
    }
}

impl Related<super::background_worker::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Worker.def()
    }
}

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
