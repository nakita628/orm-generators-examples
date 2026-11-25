use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "DataroomDocument")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "dataroomId", column_type = "Text")]
    pub dataroom_id: String,
    #[sea_orm(column_name = "documentId", column_type = "Text")]
    pub document_id: String,
    #[sea_orm(column_name = "folderId", column_type = "Text")]
    pub folder_id: Option<String>,
    #[sea_orm(column_name = "orderIndex")]
    pub order_index: Option<i32>,
    #[sea_orm(column_name = "hierarchicalIndex", column_type = "Text")]
    pub hierarchical_index: Option<String>,
    #[sea_orm(column_name = "vectorStoreFileId", column_type = "Text")]
    pub vector_store_file_id: Option<String>,
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
        belongs_to = "super::dataroom::Entity",
        from = "Column::DataroomId",
        to = "super::dataroom::Column::Id",
        on_delete = "Cascade",
        fk_name = "DataroomDocument_dataroomId_fkey"
    )]
    Dataroom,
    #[sea_orm(
        belongs_to = "super::document::Entity",
        from = "Column::DocumentId",
        to = "super::document::Column::Id",
        on_delete = "Cascade",
        fk_name = "DataroomDocument_documentId_fkey"
    )]
    Document,
    #[sea_orm(
        belongs_to = "super::dataroom_folder::Entity",
        from = "Column::FolderId",
        to = "super::dataroom_folder::Column::Id",
        on_delete = "SetNull",
        fk_name = "DataroomDocument_folderId_fkey"
    )]
    Folder,
    #[sea_orm(has_many = "super::conversation::Entity")]
    Conversations,
    #[sea_orm(has_many = "super::dataroom_question::Entity")]
    Questions,
    #[sea_orm(has_many = "super::dataroom_faq_item::Entity")]
    DataroomFAQs,
    #[sea_orm(has_many = "super::document_upload::Entity")]
    UploadedDocuments,
}

impl Related<super::dataroom::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Dataroom.def()
    }
}

impl Related<super::document::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Document.def()
    }
}

impl Related<super::dataroom_folder::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Folder.def()
    }
}

impl Related<super::conversation::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Conversations.def()
    }
}

impl Related<super::dataroom_question::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Questions.def()
    }
}

impl Related<super::dataroom_faq_item::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::DataroomFAQs.def()
    }
}

impl Related<super::document_upload::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::UploadedDocuments.def()
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
