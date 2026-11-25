use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LiteLLM_ObjectPermissionTable")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub object_permission_id: String,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub mcp_servers: Vec<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub mcp_access_groups: Vec<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub mcp_tool_permissions: Option<Json>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub vector_stores: Vec<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub agents: Vec<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub agent_access_groups: Vec<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub models: Vec<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub blocked_tools: Vec<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub mcp_toolsets: Vec<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub search_tools: Vec<String>,
    pub mcp_tool_search_enabled: Option<bool>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub skills: Vec<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub apps: Vec<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(has_many = "super::lite_llm_team_table::Entity")]
    Teams,
    #[sea_orm(has_many = "super::lite_llm_project_table::Entity")]
    Projects,
    #[sea_orm(has_many = "super::lite_llm_verification_token::Entity")]
    VerificationTokens,
    #[sea_orm(has_many = "super::lite_llm_organization_table::Entity")]
    Organizations,
    #[sea_orm(has_many = "super::lite_llm_user_table::Entity")]
    Users,
    #[sea_orm(has_many = "super::lite_llm_end_user_table::Entity")]
    EndUsers,
    #[sea_orm(has_many = "super::lite_llm_agents_table::Entity")]
    AgentsTable,
}

impl Related<super::lite_llm_team_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Teams.def()
    }
}

impl Related<super::lite_llm_project_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Projects.def()
    }
}

impl Related<super::lite_llm_verification_token::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::VerificationTokens.def()
    }
}

impl Related<super::lite_llm_organization_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organizations.def()
    }
}

impl Related<super::lite_llm_user_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Users.def()
    }
}

impl Related<super::lite_llm_end_user_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::EndUsers.def()
    }
}

impl Related<super::lite_llm_agents_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AgentsTable.def()
    }
}

impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            object_permission_id: Set(uuid::Uuid::new_v4().to_string()),
            ..ActiveModelTrait::default()
        }
    }
}
