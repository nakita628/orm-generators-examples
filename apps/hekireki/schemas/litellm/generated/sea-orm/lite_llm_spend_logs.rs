use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "LiteLLM_SpendLogs")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub request_id: String,
    #[sea_orm(column_type = "Text")]
    pub call_type: String,
    #[sea_orm(column_type = "Text", default_value = "")]
    pub api_key: String,
    #[sea_orm(default_value = 0)]
    pub spend: f64,
    #[sea_orm(default_value = 0)]
    pub total_tokens: i32,
    #[sea_orm(default_value = 0)]
    pub prompt_tokens: i32,
    #[sea_orm(default_value = 0)]
    pub completion_tokens: i32,
    #[sea_orm(column_name = "startTime", column_type = "custom(\"timestamp(3)\")")]
    pub start_time: DateTime,
    #[sea_orm(column_name = "endTime", column_type = "custom(\"timestamp(3)\")")]
    pub end_time: DateTime,
    pub request_duration_ms: Option<i32>,
    #[sea_orm(
        column_name = "completionStartTime",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub completion_start_time: Option<DateTime>,
    #[sea_orm(column_type = "Text", default_value = "")]
    pub model: String,
    #[sea_orm(column_type = "Text", default_value = "")]
    pub model_id: Option<String>,
    #[sea_orm(column_type = "Text", default_value = "")]
    pub model_group: Option<String>,
    #[sea_orm(column_type = "Text", default_value = "")]
    pub custom_llm_provider: Option<String>,
    #[sea_orm(column_type = "Text", default_value = "")]
    pub api_base: Option<String>,
    #[sea_orm(column_type = "Text", default_value = "")]
    pub user: Option<String>,
    #[sea_orm(column_type = "JsonBinary", default_value = "{}")]
    pub metadata: Option<Json>,
    #[sea_orm(column_type = "Text", default_value = "")]
    pub cache_hit: Option<String>,
    #[sea_orm(column_type = "Text", default_value = "")]
    pub cache_key: Option<String>,
    #[sea_orm(column_type = "JsonBinary", default_value = "[]")]
    pub request_tags: Option<Json>,
    #[sea_orm(column_type = "Text")]
    pub team_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub organization_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub end_user: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub requester_ip_address: Option<String>,
    #[sea_orm(column_type = "JsonBinary", default_value = "{}")]
    pub messages: Option<Json>,
    #[sea_orm(column_type = "JsonBinary", default_value = "{}")]
    pub response: Option<Json>,
    #[sea_orm(column_type = "Text")]
    pub session_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub status: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub mcp_namespaced_tool_name: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub billing_agent_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub agent_id: Option<String>,
    #[sea_orm(column_type = "JsonBinary", default_value = "{}")]
    pub proxy_server_request: Option<Json>,
    #[sea_orm(column_type = "Text")]
    pub litellm_call_id: Option<String>,
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
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

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
