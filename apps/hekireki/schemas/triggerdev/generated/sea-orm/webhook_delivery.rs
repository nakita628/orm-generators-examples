use super::runtime_environment_type::RuntimeEnvironmentType;
use super::webhook_delivery_status::WebhookDeliveryStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "WebhookDelivery")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "friendlyId", column_type = "Text")]
    pub friendly_id: String,
    #[sea_orm(column_name = "webhookEndpointId", column_type = "Text")]
    pub webhook_endpoint_id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_name = "runtimeEnvironmentId", column_type = "Text")]
    pub runtime_environment_id: String,
    #[sea_orm(column_name = "environmentType")]
    pub environment_type: RuntimeEnvironmentType,
    #[sea_orm(column_name = "externalDeliveryId", column_type = "Text")]
    pub external_delivery_id: String,
    #[sea_orm(column_name = "idempotencyKey", column_type = "Text")]
    pub idempotency_key: String,
    #[sea_orm(column_name = "runId", column_type = "Text")]
    pub run_id: Option<String>,
    #[sea_orm(default_value = "PENDING")]
    pub status: WebhookDeliveryStatus,
    /// Set from the x-trigger-test ingress header; marks console/test-send deliveries so the list can filter them.
    #[sea_orm(column_name = "isTest", default_value = false)]
    pub is_test: bool,
    #[sea_orm(column_name = "parsedEvent", column_type = "JsonBinary")]
    pub parsed_event: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub headers: Option<Json>,
    #[sea_orm(column_name = "rawBodyHash", column_type = "Text")]
    pub raw_body_hash: Option<String>,
    #[sea_orm(column_name = "errorMessage", column_type = "Text")]
    pub error_message: Option<String>,
    #[sea_orm(column_name = "filterReason", column_type = "Text")]
    pub filter_reason: Option<String>,
    /// Per-target results: WebhookDeliveryTargetResult[] ({ id, type, status, reason?, runId?, error? }).
    #[sea_orm(column_type = "JsonBinary", default_value = "[]")]
    pub targets: Json,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "processedAt", column_type = "custom(\"timestamp(3)\")")]
    pub processed_at: Option<DateTime>,
    /// Retention class in days: the LIST partition key, stamped at ingest from the org's plan and encoded in the id.
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "retentionDays", default_value = 30)]
    pub retention_days: i32,
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
