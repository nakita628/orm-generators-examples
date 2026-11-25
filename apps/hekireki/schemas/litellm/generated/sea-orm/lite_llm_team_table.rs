use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LiteLLM_TeamTable")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub team_id: String,
    #[sea_orm(column_type = "Text")]
    pub team_alias: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub organization_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub object_permission_id: Option<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub admins: Vec<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub members: Vec<String>,
    #[sea_orm(column_type = "JsonBinary", default_value = "{}")]
    pub members_with_roles: Json,
    #[sea_orm(column_type = "JsonBinary", default_value = "{}")]
    pub metadata: Json,
    pub max_budget: Option<f64>,
    pub soft_budget: Option<f64>,
    #[sea_orm(default_value = 0)]
    pub spend: f64,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub models: Vec<String>,
    pub max_parallel_requests: Option<i32>,
    pub tpm_limit: Option<i64>,
    pub rpm_limit: Option<i64>,
    pub tpd_limit: Option<i64>,
    #[sea_orm(column_type = "Text")]
    pub budget_duration: Option<String>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub budget_reset_at: Option<DateTime>,
    #[sea_orm(default_value = false)]
    pub blocked: bool,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
    #[sea_orm(column_type = "JsonBinary", default_value = "{}")]
    pub model_spend: Json,
    #[sea_orm(column_type = "JsonBinary", default_value = "{}")]
    pub model_max_budget: Json,
    #[sea_orm(column_type = "JsonBinary", default_value = "{}")]
    pub router_settings: Option<Json>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub team_member_permissions: Vec<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub access_group_ids: Vec<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub policies: Vec<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub default_team_member_models: Vec<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub budget_limits: Option<Json>,
    #[sea_orm(unique)]
    pub model_id: Option<i32>,
    #[sea_orm(default_value = false)]
    pub allow_team_guardrail_config: bool,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::lite_llm_organization_table::Entity",
        from = "Column::OrganizationId",
        to = "super::lite_llm_organization_table::Column::OrganizationId",
        fk_name = "LiteLLM_TeamTable_organization_id_fkey"
    )]
    LitellmOrganizationTable,
    #[sea_orm(
        belongs_to = "super::lite_llm_model_table::Entity",
        from = "Column::ModelId",
        to = "super::lite_llm_model_table::Column::Id",
        fk_name = "LiteLLM_TeamTable_model_id_fkey"
    )]
    LitellmModelTable,
    #[sea_orm(
        belongs_to = "super::lite_llm_object_permission_table::Entity",
        from = "Column::ObjectPermissionId",
        to = "super::lite_llm_object_permission_table::Column::ObjectPermissionId",
        fk_name = "LiteLLM_TeamTable_object_permission_id_fkey"
    )]
    ObjectPermission,
    #[sea_orm(has_many = "super::lite_llm_project_table::Entity")]
    Projects,
}

impl Related<super::lite_llm_organization_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::LitellmOrganizationTable.def()
    }
}

impl Related<super::lite_llm_model_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::LitellmModelTable.def()
    }
}

impl Related<super::lite_llm_object_permission_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ObjectPermission.def()
    }
}

impl Related<super::lite_llm_project_table::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Projects.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            team_id: Set(uuid::Uuid::new_v4().to_string()),
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
