use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

/// Completion ledger for computed-update outbox tasks (lineage / latency
/// tracing). Successful tasks are hard-deleted from the outbox on markDone,
/// so this table keeps a retention-bounded trail per completed task: the
/// run/stage chain (runId, stageDepth, predecessorTaskId) and the timeline
/// from source mutation to converged write (sourceChangedAt → completedAt).
/// Observability only — the worker never reads it; rows are pruned after a
/// retention window.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "computed_update_run_history")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub task_id: String,
    #[sea_orm(column_type = "Text")]
    pub base_id: String,
    #[sea_orm(column_type = "Text")]
    pub seed_table_id: String,
    #[sea_orm(column_type = "Text")]
    pub change_type: String,
    #[sea_orm(column_type = "Text")]
    pub run_id: String,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub origin_run_ids: Vec<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub steps: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub edges: Option<Json>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub affected_table_ids: Vec<String>,
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub affected_field_ids: Vec<String>,
    /// Source (changed) field ids; populated for seed tasks, where
    /// affected_field_ids carries the trigger fields rather than targets.
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub source_field_ids: Vec<String>,
    #[sea_orm(default_value = 0)]
    pub seed_record_count: i32,
    #[sea_orm(default_value = 0)]
    pub stage_depth: i32,
    #[sea_orm(column_type = "Text")]
    pub predecessor_task_id: Option<String>,
    #[sea_orm(default_value = 0)]
    pub run_total_steps: i32,
    #[sea_orm(default_value = 0)]
    pub run_completed_steps_before: i32,
    pub sync_max_level: Option<i32>,
    #[sea_orm(default_value = 0)]
    pub estimated_complexity: i32,
    #[sea_orm(default_value = 0)]
    pub attempts: i32,
    #[sea_orm(column_type = "Text")]
    pub outcome: String,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub source_changed_at: Option<DateTime>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub enqueued_at: DateTime,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub started_at: Option<DateTime>,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub completed_at: DateTime,
    #[sea_orm(default_value = 0)]
    pub duration_ms: i32,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {}

impl ActiveModelBehavior for ActiveModel {}
