use super::gateway_realtime_session_status::GatewayRealtimeSessionStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "GatewayRealtimeSession")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_name = "virtualKeyId", column_type = "Text")]
    pub virtual_key_id: String,
    #[sea_orm(column_name = "modelProviderId", column_type = "Text")]
    pub model_provider_id: String,
    #[sea_orm(column_type = "Text")]
    pub vendor: String,
    #[sea_orm(column_name = "agentId", column_type = "Text")]
    pub agent_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub model: String,
    #[sea_orm(column_name = "requestedModel", column_type = "Text")]
    pub requested_model: Option<String>,
    #[sea_orm(column_name = "vendorConversationId", column_type = "Text")]
    pub vendor_conversation_id: Option<String>,
    #[sea_orm(default_value = "OPEN")]
    pub status: GatewayRealtimeSessionStatus,
    #[sea_orm(
        column_name = "mintedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub minted_at: DateTime,
    #[sea_orm(column_name = "closedAt", column_type = "custom(\"timestamp(3)\")")]
    pub closed_at: Option<DateTime>,
    #[sea_orm(column_name = "closeReason", column_type = "Text")]
    pub close_reason: Option<String>,
    #[sea_orm(column_name = "traceId", column_type = "Text")]
    pub trace_id: Option<String>,
    #[sea_orm(column_name = "vendorCostRaw", column_type = "JsonBinary")]
    pub vendor_cost_raw: Option<Json>,
    #[sea_orm(column_type = "Text")]
    pub kind: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub metering: Option<String>,
    #[sea_orm(
        column_name = "credentialExpiresAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub credential_expires_at: Option<DateTime>,
    #[sea_orm(column_name = "transcriptionModel", column_type = "Text")]
    pub transcription_model: Option<String>,
    #[sea_orm(column_name = "endUserId", column_type = "Text")]
    pub end_user_id: Option<String>,
    #[sea_orm(column_name = "lastReportAt", column_type = "custom(\"timestamp(3)\")")]
    pub last_report_at: Option<DateTime>,
    #[sea_orm(column_name = "reportedCostNanoUsd", default_value = "0")]
    pub reported_cost_nano_usd: i64,
    #[sea_orm(column_name = "reportCount", default_value = 0)]
    pub report_count: i32,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(
        column_name = "updatedAt",
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
        if insert && self.minted_at.is_not_set() {
            self.minted_at = Set(now.naive_utc());
        }
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.naive_utc());
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
