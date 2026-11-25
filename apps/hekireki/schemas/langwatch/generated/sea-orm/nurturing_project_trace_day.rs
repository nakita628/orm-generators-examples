use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// retired: unread since v3.28.0. Replaced by NurturingProjectTraceDayCount, which the trace-day
/// peer folds write and a replay rebuilds; an operator step may drop it once old writers are gone.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "NurturingProjectTraceDay")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    /// UTC `YYYY-MM-DD` of the root spans' append time.
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub day: String,
    #[sea_orm(column_name = "traceCount", default_value = 0)]
    pub trace_count: i32,
    /// Traces trace recorded failed (lw.trace.trace_failed).
    #[sea_orm(column_name = "errorCount", default_value = 0)]
    pub error_count: i32,
    /// Append time of the newest root span.
    #[sea_orm(column_name = "lastTraceAt", column_type = "custom(\"timestamp(3)\")")]
    pub last_trace_at: Option<DateTime>,
    /// Append time of the newest trace_failed.
    #[sea_orm(
        column_name = "lastFailedTraceAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub last_failed_trace_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
