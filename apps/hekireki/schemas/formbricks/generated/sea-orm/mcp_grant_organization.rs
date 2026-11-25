use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// An organization a grant with `organizations` access reaches, workspaces created there later included.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "McpGrantOrganization")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "grantId", column_type = "Text")]
    pub grant_id: String,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::mcp_grant::Entity",
        from = "Column::GrantId",
        to = "super::mcp_grant::Column::Id",
        on_delete = "Cascade",
        fk_name = "McpGrantOrganization_grantId_fkey"
    )]
    Grant,
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_delete = "Cascade",
        fk_name = "McpGrantOrganization_organizationId_fkey"
    )]
    Organization,
}

impl Related<super::mcp_grant::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Grant.def()
    }
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
