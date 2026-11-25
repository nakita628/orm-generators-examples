use super::batch_task_run_status::BatchTaskRunStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "BatchTaskRun")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "friendlyId", column_type = "Text")]
    pub friendly_id: String,
    #[sea_orm(column_name = "idempotencyKey", column_type = "Text")]
    pub idempotency_key: Option<String>,
    #[sea_orm(
        column_name = "idempotencyKeyExpiresAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub idempotency_key_expires_at: Option<DateTime>,
    #[sea_orm(default_value = "PENDING")]
    pub status: BatchTaskRunStatus,
    #[sea_orm(column_name = "runtimeEnvironmentId", column_type = "Text")]
    pub runtime_environment_id: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    /// Friendly IDs
    #[sea_orm(
        column_name = "runIds",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub run_ids: Vec<String>,
    #[sea_orm(column_name = "runCount", default_value = 0)]
    pub run_count: i32,
    #[sea_orm(column_type = "Text")]
    pub payload: Option<String>,
    #[sea_orm(
        column_name = "payloadType",
        column_type = "Text",
        default_value = "application/json"
    )]
    pub payload_type: String,
    #[sea_orm(column_type = "JsonBinary")]
    pub options: Option<Json>,
    #[sea_orm(
        column_name = "batchVersion",
        column_type = "Text",
        default_value = "v1"
    )]
    pub batch_version: String,
    /// sealed is set to true once no more items can be added to the batch
    #[sea_orm(default_value = false)]
    pub sealed: bool,
    #[sea_orm(column_name = "sealedAt", column_type = "custom(\"timestamp(3)\")")]
    pub sealed_at: Option<DateTime>,
    /// this is the expected number of items in the batch
    #[sea_orm(column_name = "expectedCount", default_value = 0)]
    pub expected_count: i32,
    /// this is the completed number of items in the batch. once this reaches expectedCount, and the batch is sealed, the batch is considered completed
    #[sea_orm(column_name = "completedCount", default_value = 0)]
    pub completed_count: i32,
    #[sea_orm(column_name = "completedAt", column_type = "custom(\"timestamp(3)\")")]
    pub completed_at: Option<DateTime>,
    #[sea_orm(column_name = "resumedAt", column_type = "custom(\"timestamp(3)\")")]
    pub resumed_at: Option<DateTime>,
    /// this is used to be able to "seal" this BatchTaskRun when all of the runs have been triggered asynchronously, and using the "parallel" processing strategy
    #[sea_orm(column_name = "processingJobsCount", default_value = 0)]
    pub processing_jobs_count: i32,
    #[sea_orm(column_name = "processingJobsExpectedCount", default_value = 0)]
    pub processing_jobs_expected_count: i32,
    /// optional token that can be used to authenticate the task run
    #[sea_orm(unique)]
    #[sea_orm(column_name = "oneTimeUseToken", column_type = "Text")]
    pub one_time_use_token: Option<String>,
    /// When processing started (status changed to PROCESSING)
    #[sea_orm(
        column_name = "processingStartedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub processing_started_at: Option<DateTime>,
    /// When processing completed (all items processed)
    #[sea_orm(
        column_name = "processingCompletedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub processing_completed_at: Option<DateTime>,
    /// Count of successfully created runs
    #[sea_orm(column_name = "successfulRunCount")]
    pub successful_run_count: Option<i32>,
    /// Count of failed run creations
    #[sea_orm(column_name = "failedRunCount")]
    pub failed_run_count: Option<i32>,
    #[sea_orm(column_name = "taskIdentifier", column_type = "Text")]
    pub task_identifier: Option<String>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "checkpointEventId", column_type = "Text")]
    pub checkpoint_event_id: Option<String>,
    #[sea_orm(column_name = "dependentTaskAttemptId", column_type = "Text")]
    pub dependent_task_attempt_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::checkpoint_restore_event::Entity",
        from = "Column::CheckpointEventId",
        to = "super::checkpoint_restore_event::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "BatchTaskRun_checkpointEventId_fkey"
    )]
    CheckpointEvent,
    #[sea_orm(
        belongs_to = "super::task_run_attempt::Entity",
        from = "Column::DependentTaskAttemptId",
        to = "super::task_run_attempt::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "BatchTaskRun_dependentTaskAttemptId_fkey"
    )]
    DependentTaskAttempt,
    #[sea_orm(has_many = "super::task_run::Entity")]
    Runs,
    #[sea_orm(has_many = "super::task_run_waitpoint::Entity")]
    RunsBlocked,
    #[sea_orm(has_many = "super::waitpoint::Entity")]
    Waitpoints,
    #[sea_orm(has_many = "super::batch_task_run_error::Entity")]
    Errors,
    #[sea_orm(has_many = "super::batch_task_run_item::Entity")]
    Items,
    #[sea_orm(has_many = "super::task_run_dependency::Entity")]
    RunDependencies,
}

impl Related<super::checkpoint_restore_event::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CheckpointEvent.def()
    }
}

impl Related<super::task_run_attempt::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::DependentTaskAttempt.def()
    }
}

impl Related<super::task_run::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Runs.def()
    }
}

impl Related<super::task_run_waitpoint::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::RunsBlocked.def()
    }
}

impl Related<super::waitpoint::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Waitpoints.def()
    }
}

impl Related<super::batch_task_run_error::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Errors.def()
    }
}

impl Related<super::batch_task_run_item::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Items.def()
    }
}

impl Related<super::task_run_dependency::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::RunDependencies.def()
    }
}

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
