use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Nurturing's own trace counts from trace's span_received and trace_failed (ARCHITECTURE §9), a
/// root span being one trace, one row per project and UTC day of append time, incremented in place
/// by each batch. A redelivered batch can count a trace twice; accepted (issue 9378). Kept 30 days.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "NurturingProjectTraceDay")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "projectId")]
    pub project_id: String,
    /// UTC `YYYY-MM-DD` of the root spans' append time.
    #[sea_orm(primary_key, auto_increment = false)]
    pub day: String,
    #[sea_orm(column_name = "traceCount", default_value = 0)]
    pub trace_count: i32,
    /// Traces trace recorded failed (lw.trace.trace_failed).
    #[sea_orm(column_name = "errorCount", default_value = 0)]
    pub error_count: i32,
    /// Append time of the newest root span.
    #[sea_orm(column_name = "lastTraceAt")]
    pub last_trace_at: Option<DateTime>,
    /// Append time of the newest trace_failed.
    #[sea_orm(column_name = "lastFailedTraceAt")]
    pub last_failed_trace_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
