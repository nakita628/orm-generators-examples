use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AttachmentDocument")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "attachmentSourceId", column_type = "Text")]
    pub attachment_source_id: String,
    #[sea_orm(column_name = "fileId", column_type = "Text")]
    pub file_id: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_name = "mimeType", column_type = "Text")]
    pub mime_type: String,
    #[sea_orm(column_name = "modifiedAt", column_type = "custom(\"timestamp(3)\")")]
    pub modified_at: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub summary: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub content: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub metadata: Option<Json>,
    #[sea_orm(column_name = "indexedAt", column_type = "custom(\"timestamp(3)\")")]
    pub indexed_at: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub error: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::attachment_source::Entity",
        from = "Column::AttachmentSourceId",
        to = "super::attachment_source::Column::Id",
        on_delete = "Cascade",
        fk_name = "AttachmentDocument_attachmentSourceId_fkey"
    )]
    AttachmentSource,
}

impl Related<super::attachment_source::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AttachmentSource.def()
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
