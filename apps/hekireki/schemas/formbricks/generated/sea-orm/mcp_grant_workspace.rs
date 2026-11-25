use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// A workspace a grant with `workspaces` access reaches.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "McpGrantWorkspace")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "grantId", column_type = "Text")]
    pub grant_id: String,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "workspaceId", column_type = "Text")]
    pub workspace_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::mcp_grant::Entity",
        from = "Column::GrantId",
        to = "super::mcp_grant::Column::Id",
        on_delete = "Cascade",
        fk_name = "McpGrantWorkspace_grantId_fkey"
    )]
    Grant,
    #[sea_orm(
        belongs_to = "super::workspace::Entity",
        from = "Column::WorkspaceId",
        to = "super::workspace::Column::Id",
        on_delete = "Cascade",
        fk_name = "McpGrantWorkspace_workspaceId_fkey"
    )]
    Workspace,
}

impl Related<super::mcp_grant::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Grant.def()
    }
}

impl Related<super::workspace::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Workspace.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
