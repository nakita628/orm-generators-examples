use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "mastra_scorers")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "scorerId", column_type = "Text")]
    pub scorer_id: String,
    #[sea_orm(column_name = "traceId", column_type = "Text")]
    pub trace_id: Option<String>,
    #[sea_orm(column_name = "runId", column_type = "Text")]
    pub run_id: String,
    #[sea_orm(column_type = "JsonBinary")]
    pub scorer: Json,
    #[sea_orm(column_name = "preprocessStepResult", column_type = "JsonBinary")]
    pub preprocess_step_result: Option<Json>,
    #[sea_orm(column_name = "extractStepResult", column_type = "JsonBinary")]
    pub extract_step_result: Option<Json>,
    #[sea_orm(column_name = "analyzeStepResult", column_type = "JsonBinary")]
    pub analyze_step_result: Option<Json>,
    pub score: f64,
    #[sea_orm(column_type = "Text")]
    pub reason: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub metadata: Option<Json>,
    #[sea_orm(column_name = "preprocessPrompt", column_type = "Text")]
    pub preprocess_prompt: Option<String>,
    #[sea_orm(column_name = "extractPrompt", column_type = "Text")]
    pub extract_prompt: Option<String>,
    #[sea_orm(column_name = "generateScorePrompt", column_type = "Text")]
    pub generate_score_prompt: Option<String>,
    #[sea_orm(column_name = "generateReasonPrompt", column_type = "Text")]
    pub generate_reason_prompt: Option<String>,
    #[sea_orm(column_name = "analyzePrompt", column_type = "Text")]
    pub analyze_prompt: Option<String>,
    #[sea_orm(column_name = "reasonPrompt", column_type = "Text")]
    pub reason_prompt: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub input: Json,
    #[sea_orm(column_type = "JsonBinary")]
    pub output: Json,
    #[sea_orm(column_name = "additionalContext", column_type = "JsonBinary")]
    pub additional_context: Option<Json>,
    #[sea_orm(column_name = "runtimeContext", column_type = "JsonBinary")]
    pub runtime_context: Option<Json>,
    #[sea_orm(column_name = "entityType", column_type = "Text")]
    pub entity_type: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub entity: Option<Json>,
    #[sea_orm(column_name = "entityId", column_type = "Text")]
    pub entity_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub source: String,
    #[sea_orm(column_name = "resourceId", column_type = "Text")]
    pub resource_id: Option<String>,
    #[sea_orm(column_name = "threadId", column_type = "Text")]
    pub thread_id: Option<String>,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
    #[sea_orm(
        column_name = "createdAtZ",
        column_type = "TimestampWithTimeZone",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at_z: Option<DateTimeWithTimeZone>,
    #[sea_orm(
        column_name = "updatedAtZ",
        column_type = "TimestampWithTimeZone",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at_z: Option<DateTimeWithTimeZone>,
    #[sea_orm(column_name = "spanId", column_type = "Text")]
    pub span_id: Option<String>,
    #[sea_orm(column_name = "requestContext", column_type = "JsonBinary")]
    pub request_context: Option<Json>,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: Option<String>,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: Option<String>,
    #[sea_orm(column_name = "batchId", column_type = "Text")]
    pub batch_id: Option<String>,
    #[sea_orm(column_name = "datasetId", column_type = "Text")]
    pub dataset_id: Option<String>,
    #[sea_orm(column_name = "datasetItemId", column_type = "Text")]
    pub dataset_item_id: Option<String>,
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
        if insert && self.created_at_z.is_not_set() {
            self.created_at_z = Set(Some(now.fixed_offset()));
        }
        if insert && self.updated_at_z.is_not_set() {
            self.updated_at_z = Set(Some(now.fixed_offset()));
        }
        Ok(self)
    }
}
