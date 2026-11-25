use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LiteLLM_MCPServerTable")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub server_id: String,
    #[sea_orm(column_type = "Text")]
    pub server_name: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub alias: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub instructions: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub url: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub spec_path: Option<String>,
    #[sea_orm(column_type = "Text", default_value = "sse")]
    pub transport: String,
    #[sea_orm(column_type = "Text")]
    pub auth_type: Option<String>,
    #[sea_orm(column_type = "JsonBinary", default_value = "{}")]
    pub credentials: Option<Json>,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub created_by: Option<String>,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub updated_by: Option<String>,
    #[sea_orm(column_type = "JsonBinary", default_value = "{}")]
    pub mcp_info: Option<Json>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub mcp_access_groups: Vec<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub allowed_tools: Vec<String>,
    #[sea_orm(column_type = "JsonBinary", default_value = "{}")]
    pub tool_name_to_display_name: Option<Json>,
    #[sea_orm(column_type = "JsonBinary", default_value = "{}")]
    pub tool_name_to_description: Option<Json>,
    #[sea_orm(column_type = "JsonBinary", default_value = "{}")]
    pub pinned_tools: Option<Json>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub extra_headers: Vec<String>,
    #[sea_orm(column_type = "JsonBinary", default_value = "{}")]
    pub static_headers: Option<Json>,
    #[sea_orm(column_type = "JsonBinary", default_value = "[]")]
    pub env_vars: Option<Json>,
    #[sea_orm(column_type = "Text", default_value = "unknown")]
    pub status: Option<String>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub last_health_check: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub health_check_error: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub command: Option<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub args: Vec<String>,
    #[sea_orm(column_type = "JsonBinary", default_value = "{}")]
    pub env: Option<Json>,
    #[sea_orm(column_type = "Text")]
    pub issuer: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub authorization_url: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub token_url: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub registration_url: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub oauth2_flow: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub token_exchange_endpoint: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub audience: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub subject_token_type: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub token_exchange_profile: Option<String>,
    #[sea_orm(default_value = false)]
    pub allow_all_keys: bool,
    #[sea_orm(default_value = true)]
    pub available_on_public_internet: bool,
    #[sea_orm(default_value = false)]
    pub delegate_auth_to_upstream: bool,
    #[sea_orm(default_value = false)]
    pub oauth_passthrough: bool,
    pub dcr_bridge: Option<bool>,
    #[sea_orm(default_value = false)]
    pub per_server_oauth_discovery: bool,
    #[sea_orm(default_value = false)]
    pub is_byok: bool,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub byok_description: Vec<String>,
    #[sea_orm(column_type = "Text")]
    pub byok_api_key_help_url: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub source_url: Option<String>,
    pub timeout: Option<f64>,
    pub max_concurrent_requests: Option<i32>,
    pub rpm: Option<i32>,
    #[sea_orm(column_type = "JsonBinary")]
    pub events: Option<Json>,
    #[sea_orm(column_type = "Text", default_value = "active")]
    pub approval_status: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub submitted_by: Option<String>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub submitted_at: Option<DateTime>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub reviewed_at: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub review_notes: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            server_id: Set(uuid::Uuid::new_v4().to_string()),
            ..ActiveModelTrait::default()
        }
    }

    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(Some(now.naive_utc()));
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(Some(now.naive_utc()));
        }
        Ok(self)
    }
}
