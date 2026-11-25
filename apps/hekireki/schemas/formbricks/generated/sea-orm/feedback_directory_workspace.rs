use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Links feedback directories to workspaces (workspaces).
/// Manages which workspaces can access a given directory.
///
/// @property feedbackDirectory - The directory being accessed
/// @property workspace - The workspace receiving access
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "FeedbackDirectoryWorkspace")]
pub struct Model {
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "feedbackDirectoryId", column_type = "Text")]
    pub feedback_directory_id: String,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "workspaceId", column_type = "Text")]
    pub workspace_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::feedback_directory::Entity",
        from = "Column::FeedbackDirectoryId",
        to = "super::feedback_directory::Column::Id",
        on_delete = "Cascade",
        fk_name = "FeedbackDirectoryWorkspace_feedbackDirectoryId_fkey"
    )]
    FeedbackDirectory,
    #[sea_orm(
        belongs_to = "super::workspace::Entity",
        from = "Column::WorkspaceId",
        to = "super::workspace::Column::Id",
        on_delete = "Cascade",
        fk_name = "FeedbackDirectoryWorkspace_workspaceId_fkey"
    )]
    Workspace,
    #[sea_orm(has_many = "super::feedback_source::Entity")]
    FeedbackSources,
}

impl Related<super::feedback_directory::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::FeedbackDirectory.def()
    }
}

impl Related<super::workspace::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Workspace.def()
    }
}

impl Related<super::feedback_source::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::FeedbackSources.def()
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
