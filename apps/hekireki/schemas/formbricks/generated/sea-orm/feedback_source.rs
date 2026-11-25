use super::feedback_source_element_scope::FeedbackSourceElementScope;
use super::feedback_source_import_mode::FeedbackSourceImportMode;
use super::feedback_source_status::FeedbackSourceStatus;
use super::feedback_source_type::FeedbackSourceType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// A feedback source — connects external data into Unify Feedback for record creation.
/// Naming note: "FeedbackSource" is the Unify Feedback domain term; the word "Connector"
/// is reserved for the Workflows domain (see docs/development/standards/practices/naming-conventions.mdx).
///
/// @property id - Unique identifier for the feedback source
/// @property name - Display name for the feedback source
/// @property type - Type of feedback source (formbricks_survey, csv, …)
/// @property status - Current state of the feedback source (active, paused, error)
/// @property workspace - The workspace this feedback source belongs to
/// @property feedbackDirectory - Hub directory the feedback source writes into
/// @property formbricksMappings - Element mappings for Formbricks feedback sources
/// @property fieldMappings - Field mappings for other feedback source types
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "FeedbackSource")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    pub r#type: FeedbackSourceType,
    #[sea_orm(default_value = "active")]
    pub status: FeedbackSourceStatus,
    /// Defaults to completedOnly so a new source agrees with the live ingestion path, which only
    /// fires on responseFinished. Existing rows take the default; nothing is re-imported.
    #[sea_orm(default_value = "completedOnly")]
    pub import_mode: FeedbackSourceImportMode,
    #[sea_orm(column_name = "elementScope", default_value = "specific")]
    pub element_scope: FeedbackSourceElementScope,
    #[sea_orm(column_name = "workspaceId", column_type = "Text")]
    pub workspace_id: String,
    #[sea_orm(column_name = "feedbackDirectoryId", column_type = "Text")]
    pub feedback_directory_id: String,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub last_sync_at: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub created_by: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::workspace::Entity",
        from = "Column::WorkspaceId",
        to = "super::workspace::Column::Id",
        on_delete = "Cascade",
        fk_name = "FeedbackSource_workspaceId_fkey"
    )]
    Workspace,
    #[sea_orm(
        belongs_to = "super::feedback_directory::Entity",
        from = "Column::FeedbackDirectoryId",
        to = "super::feedback_directory::Column::Id",
        on_delete = "Cascade",
        fk_name = "FeedbackSource_feedbackDirectoryId_fkey"
    )]
    FeedbackDirectory,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CreatedBy",
        to = "super::user::Column::Id",
        on_delete = "SetNull",
        fk_name = "FeedbackSource_created_by_fkey"
    )]
    Creator,
    #[sea_orm(
        belongs_to = "super::feedback_directory_workspace::Entity",
        from = "(Column::FeedbackDirectoryId, Column::WorkspaceId)",
        to = "(super::feedback_directory_workspace::Column::FeedbackDirectoryId, super::feedback_directory_workspace::Column::WorkspaceId)",
        on_update = "Cascade",
        on_delete = "NoAction",
        fk_name = "FeedbackSource_feedbackDirectoryId_workspaceId_fkey"
    )]
    FeedbackDirectoryWorkspace,
    #[sea_orm(has_many = "super::feedback_source_formbricks_mapping::Entity")]
    FormbricksMappings,
    #[sea_orm(has_many = "super::feedback_source_field_mapping::Entity")]
    FieldMappings,
}

impl Related<super::workspace::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Workspace.def()
    }
}

impl Related<super::feedback_directory::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::FeedbackDirectory.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Creator.def()
    }
}

impl Related<super::feedback_directory_workspace::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::FeedbackDirectoryWorkspace.def()
    }
}

impl Related<super::feedback_source_formbricks_mapping::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::FormbricksMappings.def()
    }
}

impl Related<super::feedback_source_field_mapping::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::FieldMappings.def()
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
