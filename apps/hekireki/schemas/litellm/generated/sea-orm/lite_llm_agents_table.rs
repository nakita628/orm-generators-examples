use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LiteLLM_AgentsTable")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub agent_id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub agent_name: String,
    #[sea_orm(column_type = "JsonBinary")]
    pub litellm_params: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub agent_card_params: Json,
    #[sea_orm(column_type = "JsonBinary", default_value = "{}")]
    pub static_headers: Option<Json>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub extra_headers: Vec<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub kill_switch: Option<Json>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub agent_access_groups: Vec<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub access_group_ids: Vec<String>,
    #[sea_orm(column_type = "Text")]
    pub object_permission_id: Option<String>,
    #[sea_orm(default_value = 0)]
    pub spend: f64,
    #[sea_orm(default_value = false)]
    pub identity_managed: bool,
    #[sea_orm(default_value = true)]
    pub enabled: bool,
    #[sea_orm(column_type = "Text", default_value = "autonomous")]
    pub execution_mode: String,
    pub tpm_limit: Option<i32>,
    pub rpm_limit: Option<i32>,
    pub session_tpm_limit: Option<i32>,
    pub session_rpm_limit: Option<i32>,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub created_by: String,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub updated_by: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::lite_llm_object_permission_table::Entity",
        from = "Column::ObjectPermissionId",
        to = "super::lite_llm_object_permission_table::Column::ObjectPermissionId",
        fk_name = "LiteLLM_AgentsTable_object_permission_id_fkey"
    )]
    ObjectPermission,
    #[sea_orm(has_many = "super::lite_llm_retired_agent_identity::Entity")]
    RetiredIdentities,
    #[sea_orm(has_one = "super::lite_llm_agent_identity::Entity")]
    Identity,
}

impl Related<super::lite_llm_object_permission_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ObjectPermission.def()
    }
}

impl Related<super::lite_llm_retired_agent_identity::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::RetiredIdentities.def()
    }
}

impl Related<super::lite_llm_agent_identity::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Identity.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            agent_id: Set(uuid::Uuid::new_v4().to_string()),
            ..ActiveModelTrait::default()
        }
    }

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
