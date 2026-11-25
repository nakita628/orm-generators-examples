use super::run_engine_version::RunEngineVersion;
use super::runtime_environment_type::RuntimeEnvironmentType;
use super::task_run_status::TaskRunStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "TaskRun")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(default_value = 0)]
    pub number: i32,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "friendlyId", column_type = "Text")]
    pub friendly_id: String,
    #[sea_orm(default_value = "V1")]
    pub engine: RunEngineVersion,
    #[sea_orm(default_value = "PENDING")]
    pub status: TaskRunStatus,
    #[sea_orm(column_name = "statusReason", column_type = "Text")]
    pub status_reason: Option<String>,
    #[sea_orm(column_name = "idempotencyKey", column_type = "Text")]
    pub idempotency_key: Option<String>,
    #[sea_orm(
        column_name = "idempotencyKeyExpiresAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub idempotency_key_expires_at: Option<DateTime>,
    /// Stores the user-provided key and scope: { key: string, scope: "run" | "attempt" | "global" }
    #[sea_orm(column_name = "idempotencyKeyOptions", column_type = "JsonBinary")]
    pub idempotency_key_options: Option<Json>,
    /// Debounce options: { key: string, delay: string, createdAt: Date }
    #[sea_orm(column_type = "JsonBinary")]
    pub debounce: Option<Json>,
    #[sea_orm(column_name = "taskIdentifier", column_type = "Text")]
    pub task_identifier: String,
    #[sea_orm(column_name = "isTest", default_value = false)]
    pub is_test: bool,
    #[sea_orm(column_type = "Text")]
    pub payload: String,
    #[sea_orm(
        column_name = "payloadType",
        column_type = "Text",
        default_value = "application/json"
    )]
    pub payload_type: String,
    #[sea_orm(column_type = "JsonBinary")]
    pub context: Option<Json>,
    #[sea_orm(column_name = "traceContext", column_type = "JsonBinary")]
    pub trace_context: Option<Json>,
    #[sea_orm(column_name = "traceId", column_type = "Text")]
    pub trace_id: String,
    #[sea_orm(column_name = "spanId", column_type = "Text")]
    pub span_id: String,
    #[sea_orm(column_name = "runtimeEnvironmentId", column_type = "Text")]
    pub runtime_environment_id: String,
    #[sea_orm(column_name = "environmentType")]
    pub environment_type: Option<RuntimeEnvironmentType>,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub queue: String,
    #[sea_orm(column_name = "lockedQueueId", column_type = "Text")]
    pub locked_queue_id: Option<String>,
    /// The main queue that this run is part of
    #[sea_orm(
        column_name = "masterQueue",
        column_type = "Text",
        default_value = "main"
    )]
    pub worker_queue: String,
    /// User-facing geo region, stamped at trigger; workerQueue is where it actually ran.
    #[sea_orm(column_type = "Text")]
    pub region: Option<String>,
    /// @deprecated
    #[sea_orm(column_name = "secondaryMasterQueue", column_type = "Text")]
    pub secondary_master_queue: Option<String>,
    /// From engine v2+ this will be defined after a run has been dequeued (starting at 1)
    #[sea_orm(column_name = "attemptNumber")]
    pub attempt_number: Option<i32>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    /// Denormized column that holds the raw tags
    #[sea_orm(
        column_name = "runTags",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub run_tags: Vec<String>,
    /// Denormalized version of the background worker task
    #[sea_orm(column_name = "taskVersion", column_type = "Text")]
    pub task_version: Option<String>,
    #[sea_orm(column_name = "sdkVersion", column_type = "Text")]
    pub sdk_version: Option<String>,
    #[sea_orm(column_name = "cliVersion", column_type = "Text")]
    pub cli_version: Option<String>,
    /// startedAt marks the point at which a run is dequeued from MarQS
    #[sea_orm(column_name = "startedAt", column_type = "custom(\"timestamp(3)\")")]
    pub started_at: Option<DateTime>,
    /// executedAt is set when the first attempt is about to execute
    #[sea_orm(column_name = "executedAt", column_type = "custom(\"timestamp(3)\")")]
    pub executed_at: Option<DateTime>,
    #[sea_orm(column_name = "completedAt", column_type = "custom(\"timestamp(3)\")")]
    pub completed_at: Option<DateTime>,
    #[sea_orm(column_name = "machinePreset", column_type = "Text")]
    pub machine_preset: Option<String>,
    #[sea_orm(column_name = "usageDurationMs", default_value = 0)]
    pub usage_duration_ms: i32,
    #[sea_orm(column_name = "costInCents", default_value = 0)]
    pub cost_in_cents: f64,
    #[sea_orm(column_name = "baseCostInCents", default_value = 0)]
    pub base_cost_in_cents: f64,
    #[sea_orm(column_name = "lockedAt", column_type = "custom(\"timestamp(3)\")")]
    pub locked_at: Option<DateTime>,
    #[sea_orm(column_name = "lockedById", column_type = "Text")]
    pub locked_by_id: Option<String>,
    #[sea_orm(column_name = "lockedToVersionId", column_type = "Text")]
    pub locked_to_version_id: Option<String>,
    /// The "priority" of the run. This is just a negative offset in ms for the queue timestamp
    /// E.g. a value of 60_000 would put the run into the queue 60s ago.
    #[sea_orm(column_name = "priorityMs", default_value = 0)]
    pub priority_ms: i32,
    #[sea_orm(column_name = "concurrencyKey", column_type = "Text")]
    pub concurrency_key: Option<String>,
    /// Gates for this run: other queues it must also hold a concurrency slot in while
    /// executing. Shape: [{ queue: string, concurrencyKey?: string }]
    #[sea_orm(column_type = "JsonBinary")]
    pub gates: Option<Json>,
    #[sea_orm(column_name = "delayUntil", column_type = "custom(\"timestamp(3)\")")]
    pub delay_until: Option<DateTime>,
    #[sea_orm(column_name = "queuedAt", column_type = "custom(\"timestamp(3)\")")]
    pub queued_at: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub ttl: Option<String>,
    #[sea_orm(column_name = "expiredAt", column_type = "custom(\"timestamp(3)\")")]
    pub expired_at: Option<DateTime>,
    #[sea_orm(column_name = "maxAttempts")]
    pub max_attempts: Option<i32>,
    #[sea_orm(column_name = "lockedRetryConfig", column_type = "JsonBinary")]
    pub locked_retry_config: Option<Json>,
    /// optional token that can be used to authenticate the task run
    #[sea_orm(unique)]
    #[sea_orm(column_name = "oneTimeUseToken", column_type = "Text")]
    pub one_time_use_token: Option<String>,
    /// Where the logs are stored
    #[sea_orm(
        column_name = "taskEventStore",
        column_type = "Text",
        default_value = "taskEvent"
    )]
    pub task_event_store: String,
    #[sea_orm(
        column_name = "queueTimestamp",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub queue_timestamp: Option<DateTime>,
    #[sea_orm(column_name = "scheduleInstanceId", column_type = "Text")]
    pub schedule_instance_id: Option<String>,
    #[sea_orm(column_name = "scheduleId", column_type = "Text")]
    pub schedule_id: Option<String>,
    #[sea_orm(
        column_name = "bulkActionGroupIds",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub bulk_action_group_ids: Vec<String>,
    #[sea_orm(
        column_name = "logsDeletedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub logs_deleted_at: Option<DateTime>,
    #[sea_orm(column_name = "replayedFromTaskRunFriendlyId", column_type = "Text")]
    pub replayed_from_task_run_friendly_id: Option<String>,
    #[sea_orm(column_name = "rootTaskRunId", column_type = "Text")]
    pub root_task_run_id: Option<String>,
    #[sea_orm(column_name = "parentTaskRunId", column_type = "Text")]
    pub parent_task_run_id: Option<String>,
    #[sea_orm(column_name = "parentTaskRunAttemptId", column_type = "Text")]
    pub parent_task_run_attempt_id: Option<String>,
    #[sea_orm(column_name = "batchId", column_type = "Text")]
    pub batch_id: Option<String>,
    /// whether or not the task run was created because of a triggerAndWait for batchTriggerAndWait
    #[sea_orm(column_name = "resumeParentOnCompletion", default_value = false)]
    pub resume_parent_on_completion: bool,
    /// The depth of this task run in the task run hierarchy
    #[sea_orm(default_value = 0)]
    pub depth: i32,
    /// The span ID of the "trigger" span in the parent task run
    #[sea_orm(column_name = "parentSpanId", column_type = "Text")]
    pub parent_span_id: Option<String>,
    /// Holds the state of the run chain for deadlock detection
    #[sea_orm(column_name = "runChainState", column_type = "JsonBinary")]
    pub run_chain_state: Option<Json>,
    /// seed run metadata
    #[sea_orm(column_name = "seedMetadata", column_type = "Text")]
    pub seed_metadata: Option<String>,
    #[sea_orm(
        column_name = "seedMetadataType",
        column_type = "Text",
        default_value = "application/json"
    )]
    pub seed_metadata_type: String,
    /// Run metadata
    #[sea_orm(column_type = "Text")]
    pub metadata: Option<String>,
    #[sea_orm(
        column_name = "metadataType",
        column_type = "Text",
        default_value = "application/json"
    )]
    pub metadata_type: String,
    #[sea_orm(column_name = "metadataVersion", default_value = 1)]
    pub metadata_version: i32,
    /// Structured annotations: triggerSource, triggerAction, rootTriggerSource, rootScheduleId
    #[sea_orm(column_type = "JsonBinary")]
    pub annotations: Option<Json>,
    /// Whether the latest attempt was a warm start. Null until first attempt starts.
    #[sea_orm(column_name = "isWarmStart")]
    pub is_warm_start: Option<bool>,
    /// Run output
    #[sea_orm(column_type = "Text")]
    pub output: Option<String>,
    #[sea_orm(
        column_name = "outputType",
        column_type = "Text",
        default_value = "application/json"
    )]
    pub output_type: String,
    /// Run error
    #[sea_orm(column_type = "JsonBinary")]
    pub error: Option<Json>,
    /// Organization's billing plan type (cached for fallback when billing API fails)
    #[sea_orm(column_name = "planType", column_type = "Text")]
    pub plan_type: Option<String>,
    #[sea_orm(column_name = "maxDurationInSeconds")]
    pub max_duration_in_seconds: Option<i32>,
    /// The version of the realtime streams implementation used by the run
    #[sea_orm(
        column_name = "realtimeStreamsVersion",
        column_type = "Text",
        default_value = "v1"
    )]
    pub realtime_streams_version: String,
    /// Store the stream keys that are being used by the run
    #[sea_orm(
        column_name = "realtimeStreams",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub realtime_streams: Vec<String>,
    /// S2 basin where this run's realtime streams live. Stamped at create
    /// time from `Organization.streamBasinName` so reads can resolve the
    /// basin without joining org. Null when the org has no per-org basin
    /// (OSS, or pre-backfill); reads fall back to the global basin.
    #[sea_orm(column_name = "streamBasinName", column_type = "Text")]
    pub stream_basin_name: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::runtime_environment::Entity",
        from = "Column::RuntimeEnvironmentId",
        to = "super::runtime_environment::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "TaskRun_runtimeEnvironmentId_fkey"
    )]
    RuntimeEnvironment,
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "TaskRun_projectId_fkey"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::background_worker_task::Entity",
        from = "Column::LockedById",
        to = "super::background_worker_task::Column::Id",
        fk_name = "TaskRun_lockedById_fkey"
    )]
    LockedBy,
    #[sea_orm(
        belongs_to = "super::background_worker::Entity",
        from = "Column::LockedToVersionId",
        to = "super::background_worker::Column::Id",
        fk_name = "TaskRun_lockedToVersionId_fkey"
    )]
    LockedToVersion,
    #[sea_orm(
        belongs_to = "super::task_run::Entity",
        from = "Column::RootTaskRunId",
        to = "super::task_run::Column::Id",
        on_update = "NoAction",
        on_delete = "SetNull",
        fk_name = "TaskRun_rootTaskRunId_fkey"
    )]
    RootTaskRun,
    #[sea_orm(
        belongs_to = "super::task_run::Entity",
        from = "Column::ParentTaskRunId",
        to = "super::task_run::Column::Id",
        on_update = "NoAction",
        on_delete = "SetNull",
        fk_name = "TaskRun_parentTaskRunId_fkey"
    )]
    ParentTaskRun,
    #[sea_orm(
        belongs_to = "super::task_run_attempt::Entity",
        from = "Column::ParentTaskRunAttemptId",
        to = "super::task_run_attempt::Column::Id",
        on_update = "NoAction",
        on_delete = "SetNull",
        fk_name = "TaskRun_parentTaskRunAttemptId_fkey"
    )]
    ParentTaskRunAttempt,
    #[sea_orm(
        belongs_to = "super::batch_task_run::Entity",
        from = "Column::BatchId",
        to = "super::batch_task_run::Column::Id",
        on_update = "NoAction",
        on_delete = "SetNull",
        fk_name = "TaskRun_batchId_fkey"
    )]
    Batch,
    #[sea_orm(
        has_many = "super::task_run_attempt::Entity",
        from = "Column::Id",
        to = "super::task_run_attempt::Column::TaskRunId"
    )]
    Attempts,
    #[sea_orm(has_many = "super::checkpoint::Entity")]
    Checkpoints,
    #[sea_orm(has_many = "super::task_run_waitpoint::Entity")]
    BlockedByWaitpoints,
    #[sea_orm(has_many = "super::batch_task_run_item::Entity")]
    BatchItems,
    #[sea_orm(has_many = "super::checkpoint_restore_event::Entity")]
    CheckpointRestoreEvent,
    #[sea_orm(has_many = "super::task_run_execution_snapshot::Entity")]
    ExecutionSnapshots,
    #[sea_orm(
        has_many = "super::bulk_action_item::Entity",
        from = "Column::Id",
        to = "super::bulk_action_item::Column::SourceRunId"
    )]
    SourceBulkActionItems,
    #[sea_orm(
        has_many = "super::bulk_action_item::Entity",
        from = "Column::Id",
        to = "super::bulk_action_item::Column::DestinationRunId"
    )]
    DestinationBulkActionItems,
    #[sea_orm(
        has_many = "super::task_run::Entity",
        from = "Column::Id",
        to = "super::task_run::Column::RootTaskRunId"
    )]
    DescendantRuns,
    #[sea_orm(
        has_many = "super::task_run::Entity",
        from = "Column::Id",
        to = "super::task_run::Column::ParentTaskRunId"
    )]
    ChildRuns,
    #[sea_orm(has_many = "super::playground_conversation::Entity")]
    PlaygroundConversations,
    #[sea_orm(has_one = "super::waitpoint::Entity")]
    AssociatedWaitpoint,
    #[sea_orm(has_one = "super::task_run_dependency::Entity")]
    Dependency,
}

impl Related<super::runtime_environment::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::RuntimeEnvironment.def()
    }
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::background_worker_task::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::LockedBy.def()
    }
}

impl Related<super::background_worker::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::LockedToVersion.def()
    }
}

impl Related<super::task_run::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::RootTaskRun.def()
    }
}

impl Related<super::task_run_attempt::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ParentTaskRunAttempt.def()
    }
}

impl Related<super::batch_task_run::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Batch.def()
    }
}

impl Related<super::checkpoint::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Checkpoints.def()
    }
}

impl Related<super::task_run_waitpoint::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::BlockedByWaitpoints.def()
    }
}

impl Related<super::batch_task_run_item::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::BatchItems.def()
    }
}

impl Related<super::checkpoint_restore_event::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CheckpointRestoreEvent.def()
    }
}

impl Related<super::task_run_execution_snapshot::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ExecutionSnapshots.def()
    }
}

impl Related<super::bulk_action_item::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::SourceBulkActionItems.def()
    }
}

impl Related<super::playground_conversation::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::PlaygroundConversations.def()
    }
}

impl Related<super::waitpoint::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AssociatedWaitpoint.def()
    }
}

impl Related<super::task_run_dependency::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Dependency.def()
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
