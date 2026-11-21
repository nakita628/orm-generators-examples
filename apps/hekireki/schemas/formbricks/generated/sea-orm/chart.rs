use super::chart_type::ChartType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Represents a chart/visualization that can be used in multiple dashboards.
/// Charts are reusable components that query analytics data.
///
/// @property id - Unique identifier for the chart
/// @property name - Display name of the chart
/// @property type - Type of visualization (bar, line, pie, etc.)
/// @property workspace - The workspace this chart belongs to
/// @property query - Cube.js query configuration (JSON)
/// @property config - Chart-specific configuration (colors, labels, etc.)
/// @property createdBy - User who created the chart
/// @property dashboards - Dashboards that use this chart
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Chart")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    pub created_at: DateTime,
    pub updated_at: DateTime,
    pub name: String,
    pub r#type: ChartType,
    #[sea_orm(column_name = "workspaceId")]
    pub workspace_id: String,
    /// [ChartQuery] - Cube.js query configuration
    #[sea_orm(default_value = "{}")]
    pub query: Json,
    /// [ChartConfig] - Visualization configuration (colors, labels, formatting)
    #[sea_orm(default_value = "{}")]
    pub config: Json,
    #[sea_orm(column_name = "createdBy")]
    pub created_by: Option<String>,
    #[sea_orm(column_name = "feedbackDirectoryId")]
    pub feedback_directory_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::workspace::Entity",
        from = "Column::WorkspaceId",
        to = "super::workspace::Column::Id",
        on_delete = "Cascade"
    )]
    Workspace,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CreatedBy",
        to = "super::user::Column::Id",
        on_delete = "SetNull"
    )]
    Creator,
    #[sea_orm(
        belongs_to = "super::feedback_directory::Entity",
        from = "Column::FeedbackDirectoryId",
        to = "super::feedback_directory::Column::Id"
    )]
    FeedbackDirectory,
    #[sea_orm(has_many = "super::dashboard_widget::Entity")]
    Widgets,
}

impl Related<super::workspace::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Workspace.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Creator.def()
    }
}

impl Related<super::feedback_directory::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::FeedbackDirectory.def()
    }
}

impl Related<super::dashboard_widget::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Widgets.def()
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
