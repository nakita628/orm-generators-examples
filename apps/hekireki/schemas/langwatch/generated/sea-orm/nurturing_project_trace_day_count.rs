use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Nurturing's own trace counts (ARCHITECTURE §9), one row per project, UTC day of append time and
/// source: `spans` from trace's span_received peer fold (a customer root span is one trace),
/// `failures` from trace_failed's. Each fold owns its rows, so the two never overwrite each other;
/// readers sum a day's sources. Rebuilt by a projection replay. Kept 30 days.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "NurturingProjectTraceDayCount")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    /// UTC `YYYY-MM-DD` of the facts' append time.
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub day: String,
    /// The peer fold that wrote the row: `spans` or `failures`.
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub source: String,
    #[sea_orm(column_name = "traceCount", default_value = 0)]
    pub trace_count: i32,
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
    /// The event ids of the last delivery folded in, so a redelivery folds once.
    #[sea_orm(
        column_name = "appliedEventIds",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub applied_event_ids: Vec<String>,
    /// A row of another version reads as undecodable and is left out of the read.
    #[sea_orm(column_name = "projectionVersion", column_type = "Text")]
    pub projection_version: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
