use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "DocumentRedaction")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "jobId", column_type = "Text")]
    pub job_id: String,
    #[sea_orm(column_name = "pageNumber")]
    pub page_number: i32,
    pub x: f64,
    pub y: f64,
    pub width: f64,
    pub height: f64,
    #[sea_orm(column_name = "textPreview", column_type = "Text")]
    pub text_preview: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub category: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub confidence: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub reason: Option<String>,
    #[sea_orm(column_type = "Text", default_value = "AI")]
    pub source: String,
    #[sea_orm(column_type = "Text", default_value = "PENDING")]
    pub status: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::document_redaction_job::Entity",
        from = "Column::JobId",
        to = "super::document_redaction_job::Column::Id",
        on_delete = "Cascade",
        fk_name = "DocumentRedaction_jobId_fkey"
    )]
    Job,
}

impl Related<super::document_redaction_job::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Job.def()
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
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
