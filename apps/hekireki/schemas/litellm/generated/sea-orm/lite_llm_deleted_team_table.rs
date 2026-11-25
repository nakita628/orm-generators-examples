use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LiteLLM_DeletedTeamTable")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
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
    pub model_id: Option<i32>,
    #[sea_orm(default_value = false)]
    pub allow_team_guardrail_config: bool,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub created_at: Option<DateTime>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: Option<DateTime>,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub deleted_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub deleted_by: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub deleted_by_api_key: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub litellm_changed_by: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            id: Set(uuid::Uuid::new_v4().to_string()),
            ..ActiveModelTrait::default()
        }
    }

    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.deleted_at.is_not_set() {
            self.deleted_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
