use super::integration_type::IntegrationType;
use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Represents third-party service integrations.
/// Enables data flow between Formbricks and external services.
///
/// @property id - Unique identifier for the integration
/// @property type - The service being integrated (Google Sheets, Notion, etc.)
/// @property workspace - The workspace this integration belongs to
/// @property config - Service-specific configuration details
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Integration")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    pub r#type: IntegrationType,
    /// [IntegrationConfig]
    #[sea_orm(column_type = "JsonBinary")]
    pub config: Json,
    #[sea_orm(column_name = "workspaceId", column_type = "Text")]
    pub workspace_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::workspace::Entity",
        from = "Column::WorkspaceId",
        to = "super::workspace::Column::Id",
        on_delete = "Cascade",
        fk_name = "Integration_workspaceId_fkey"
    )]
    Workspace,
}

impl Related<super::workspace::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Workspace.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
