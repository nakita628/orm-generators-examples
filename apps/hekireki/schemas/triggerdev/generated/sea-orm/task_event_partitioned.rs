use super::runtime_environment_type::RuntimeEnvironmentType;
use super::task_event_kind::TaskEventKind;
use super::task_event_level::TaskEventLevel;
use super::task_event_status::TaskEventStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// This is the unified otel span/log model that will eventually be replaced by clickhouse
#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "TaskEventPartitioned")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    /// This matches the span name for a trace event, or the log body for a log event
    #[sea_orm(column_type = "Text")]
    pub message: String,
    #[sea_orm(column_name = "traceId", column_type = "Text")]
    pub trace_id: String,
    #[sea_orm(column_name = "spanId", column_type = "Text")]
    pub span_id: String,
    #[sea_orm(column_name = "parentId", column_type = "Text")]
    pub parent_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub tracestate: Option<String>,
    #[sea_orm(column_name = "isError", default_value = false)]
    pub is_error: bool,
    #[sea_orm(column_name = "isPartial", default_value = false)]
    pub is_partial: bool,
    #[sea_orm(column_name = "isCancelled", default_value = false)]
    pub is_cancelled: bool,
    #[sea_orm(column_name = "serviceName", column_type = "Text")]
    pub service_name: String,
    #[sea_orm(column_name = "serviceNamespace", column_type = "Text")]
    pub service_namespace: String,
    #[sea_orm(default_value = "TRACE")]
    pub level: TaskEventLevel,
    #[sea_orm(default_value = "INTERNAL")]
    pub kind: TaskEventKind,
    #[sea_orm(default_value = "UNSET")]
    pub status: TaskEventStatus,
    #[sea_orm(column_type = "JsonBinary")]
    pub links: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub events: Option<Json>,
    /// This is the time the event started in nanoseconds since the epoch
    #[sea_orm(column_name = "startTime")]
    pub start_time: i64,
    /// This is the duration of the event in nanoseconds
    #[sea_orm(default_value = "0")]
    pub duration: i64,
    #[sea_orm(column_name = "attemptId", column_type = "Text")]
    pub attempt_id: Option<String>,
    #[sea_orm(column_name = "attemptNumber")]
    pub attempt_number: Option<i32>,
    #[sea_orm(column_name = "environmentId", column_type = "Text")]
    pub environment_id: String,
    #[sea_orm(column_name = "environmentType")]
    pub environment_type: RuntimeEnvironmentType,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_name = "projectRef", column_type = "Text")]
    pub project_ref: String,
    #[sea_orm(column_name = "runId", column_type = "Text")]
    pub run_id: String,
    #[sea_orm(column_name = "runIsTest", default_value = false)]
    pub run_is_test: bool,
    #[sea_orm(column_name = "idempotencyKey", column_type = "Text")]
    pub idempotency_key: Option<String>,
    #[sea_orm(column_name = "taskSlug", column_type = "Text")]
    pub task_slug: String,
    #[sea_orm(column_name = "taskPath", column_type = "Text")]
    pub task_path: Option<String>,
    #[sea_orm(column_name = "taskExportName", column_type = "Text")]
    pub task_export_name: Option<String>,
    #[sea_orm(column_name = "workerId", column_type = "Text")]
    pub worker_id: Option<String>,
    #[sea_orm(column_name = "workerVersion", column_type = "Text")]
    pub worker_version: Option<String>,
    #[sea_orm(column_name = "queueId", column_type = "Text")]
    pub queue_id: Option<String>,
    #[sea_orm(column_name = "queueName", column_type = "Text")]
    pub queue_name: Option<String>,
    #[sea_orm(column_name = "batchId", column_type = "Text")]
    pub batch_id: Option<String>,
    /// This represents all the span attributes available, like http.status_code, and special attributes like $style.icon, $output, $metadata.payload.userId, as it's used for searching and filtering
    #[sea_orm(column_type = "JsonBinary")]
    pub properties: Json,
    /// This represents all span attributes in the $metadata namespace, like $metadata.payload
    #[sea_orm(column_type = "JsonBinary")]
    pub metadata: Option<Json>,
    /// This represents all span attributes in the $style namespace, like $style
    #[sea_orm(column_type = "JsonBinary")]
    pub style: Option<Json>,
    /// This represents all span attributes in the $output namespace, like $output
    #[sea_orm(column_type = "JsonBinary")]
    pub output: Option<Json>,
    /// This represents the mimetype of the output, such as application/json or application/super+json
    #[sea_orm(column_name = "outputType", column_type = "Text")]
    pub output_type: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub payload: Option<Json>,
    #[sea_orm(column_name = "payloadType", column_type = "Text")]
    pub payload_type: Option<String>,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "usageDurationMs", default_value = 0)]
    pub usage_duration_ms: i32,
    #[sea_orm(column_name = "usageCostInCents", default_value = 0)]
    pub usage_cost_in_cents: f64,
    #[sea_orm(column_name = "machinePreset", column_type = "Text")]
    pub machine_preset: Option<String>,
    #[sea_orm(column_name = "machinePresetCpu")]
    pub machine_preset_cpu: Option<f64>,
    #[sea_orm(column_name = "machinePresetMemory")]
    pub machine_preset_memory: Option<f64>,
    #[sea_orm(column_name = "machinePresetCentsPerMs")]
    pub machine_preset_cents_per_ms: Option<f64>,
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
        Ok(self)
    }
}
