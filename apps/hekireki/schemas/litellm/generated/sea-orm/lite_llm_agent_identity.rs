use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LiteLLM_AgentIdentity")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub agent_id: String,
    #[sea_orm(default_value = true)]
    pub active: bool,
    #[sea_orm(column_type = "Text")]
    pub provider: String,
    #[sea_orm(column_type = "Text")]
    pub issuer: String,
    #[sea_orm(column_type = "Text")]
    pub tenant_id: String,
    #[sea_orm(column_type = "Text")]
    pub client_id: String,
    #[sea_orm(column_type = "Text")]
    pub service_principal_id: Option<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub required_roles: Vec<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub required_scopes: Vec<String>,
    #[sea_orm(column_type = "Text")]
    pub revision: String,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub last_authenticated_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::lite_llm_agents_table::Entity",
        from = "Column::AgentId",
        to = "super::lite_llm_agents_table::Column::AgentId",
        on_delete = "Cascade",
        fk_name = "LiteLLM_AgentIdentity_agent_id_fkey"
    )]
    Agent,
}

impl Related<super::lite_llm_agents_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Agent.def()
    }
}

impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            revision: Set(uuid::Uuid::new_v4().to_string()),
            ..ActiveModelTrait::default()
        }
    }
}
