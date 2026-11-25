use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "DocumentUpload")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "documentId", column_type = "Text")]
    pub document_id: String,
    #[sea_orm(column_name = "teamId", column_type = "Text")]
    pub team_id: String,
    #[sea_orm(column_name = "viewerId", column_type = "Text")]
    pub viewer_id: Option<String>,
    #[sea_orm(column_name = "viewId", column_type = "Text")]
    pub view_id: Option<String>,
    #[sea_orm(column_name = "linkId", column_type = "Text")]
    pub link_id: String,
    #[sea_orm(column_name = "dataroomId", column_type = "Text")]
    pub dataroom_id: Option<String>,
    #[sea_orm(column_name = "dataroomDocumentId", column_type = "Text")]
    pub dataroom_document_id: Option<String>,
    #[sea_orm(column_name = "taskId", column_type = "Text")]
    pub task_id: Option<String>,
    #[sea_orm(column_name = "originalFilename", column_type = "Text")]
    pub original_filename: Option<String>,
    #[sea_orm(column_name = "fileSize")]
    pub file_size: Option<i64>,
    #[sea_orm(column_name = "numPages")]
    pub num_pages: Option<i32>,
    #[sea_orm(column_name = "mimeType", column_type = "Text")]
    pub mime_type: Option<String>,
    #[sea_orm(
        column_name = "uploadedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub uploaded_at: DateTime,
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
        belongs_to = "super::document::Entity",
        from = "Column::DocumentId",
        to = "super::document::Column::Id",
        on_delete = "Cascade",
        fk_name = "DocumentUpload_documentId_fkey"
    )]
    Document,
    #[sea_orm(
        belongs_to = "super::team::Entity",
        from = "Column::TeamId",
        to = "super::team::Column::Id",
        on_delete = "Cascade",
        fk_name = "DocumentUpload_teamId_fkey"
    )]
    Team,
    #[sea_orm(
        belongs_to = "super::viewer::Entity",
        from = "Column::ViewerId",
        to = "super::viewer::Column::Id",
        on_delete = "SetNull",
        fk_name = "DocumentUpload_viewerId_fkey"
    )]
    Viewer,
    #[sea_orm(
        belongs_to = "super::view::Entity",
        from = "Column::ViewId",
        to = "super::view::Column::Id",
        on_delete = "SetNull",
        fk_name = "DocumentUpload_viewId_fkey"
    )]
    View,
    #[sea_orm(
        belongs_to = "super::link::Entity",
        from = "Column::LinkId",
        to = "super::link::Column::Id",
        on_delete = "Cascade",
        fk_name = "DocumentUpload_linkId_fkey"
    )]
    Link,
    #[sea_orm(
        belongs_to = "super::dataroom::Entity",
        from = "Column::DataroomId",
        to = "super::dataroom::Column::Id",
        on_delete = "SetNull",
        fk_name = "DocumentUpload_dataroomId_fkey"
    )]
    Dataroom,
    #[sea_orm(
        belongs_to = "super::dataroom_document::Entity",
        from = "Column::DataroomDocumentId",
        to = "super::dataroom_document::Column::Id",
        on_delete = "SetNull",
        fk_name = "DocumentUpload_dataroomDocumentId_fkey"
    )]
    DataroomDocument,
    #[sea_orm(
        belongs_to = "super::task::Entity",
        from = "Column::TaskId",
        to = "super::task::Column::Id",
        on_delete = "SetNull",
        fk_name = "DocumentUpload_taskId_fkey"
    )]
    Task,
}

impl Related<super::document::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Document.def()
    }
}

impl Related<super::team::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Team.def()
    }
}

impl Related<super::viewer::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Viewer.def()
    }
}

impl Related<super::view::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::View.def()
    }
}

impl Related<super::link::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Link.def()
    }
}

impl Related<super::dataroom::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Dataroom.def()
    }
}

impl Related<super::dataroom_document::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::DataroomDocument.def()
    }
}

impl Related<super::task::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Task.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.uploaded_at.is_not_set() {
            self.uploaded_at = Set(now.naive_utc());
        }
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.naive_utc());
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
