use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "AnnotationQueueItem")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "annotationQueueId", column_type = "Text")]
    pub annotation_queue_id: Option<String>,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: Option<String>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(
        column_name = "updatedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "createdByUserId", column_type = "Text")]
    pub created_by_user_id: Option<String>,
    #[sea_orm(column_name = "traceId", column_type = "Text")]
    pub trace_id: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_name = "doneAt", column_type = "custom(\"timestamp(3)\")")]
    pub done_at: Option<DateTime>,
    /// Set while the annotator wants this item's trace added to a dataset once
    /// the queue walk is finished. Persisted rather than kept in the browser so
    /// the selection survives a refresh and stays visible to teammates.
    #[sea_orm(
        column_name = "markedForDatasetAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub marked_for_dataset_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::annotation_queue::Entity",
        from = "Column::AnnotationQueueId",
        to = "super::annotation_queue::Column::Id",
        fk_name = "AnnotationQueueItem_annotationQueueId_fkey"
    )]
    AnnotationQueue,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        fk_name = "AnnotationQueueItem_userId_fkey"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CreatedByUserId",
        to = "super::user::Column::Id",
        fk_name = "AnnotationQueueItem_createdByUserId_fkey"
    )]
    CreatedByUser,
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        fk_name = "AnnotationQueueItem_projectId_fkey"
    )]
    Project,
}

impl Related<super::annotation_queue::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AnnotationQueue.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
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
