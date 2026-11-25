use super::run_engine_version::RunEngineVersion;
use super::runtime_environment_type::RuntimeEnvironmentType;
use super::task_run_execution_status::TaskRunExecutionStatus;
use super::task_run_status::TaskRunStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Used by the RunEngine during TaskRun execution
/// It has the required information to transactionally progress a run through states,
/// and prevent side effects like heartbeats failing a run that has progressed.
/// It is optimised for performance and is designed to be cleared at some point,
/// so there are no cascading relationships to other models.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "TaskRunExecutionSnapshot")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    /// This should always be 2+ (V1 didn't use the run engine or snapshots)
    #[sea_orm(default_value = "V2")]
    pub engine: RunEngineVersion,
    /// The execution status
    #[sea_orm(column_name = "executionStatus")]
    pub execution_status: TaskRunExecutionStatus,
    /// For debugging
    #[sea_orm(column_type = "Text")]
    pub description: String,
    /// We store invalid snapshots as a record of the run state when we tried to move
    #[sea_orm(column_name = "isValid", default_value = true)]
    pub is_valid: bool,
    #[sea_orm(column_type = "Text")]
    pub error: Option<String>,
    /// The previous snapshot ID
    #[sea_orm(column_name = "previousSnapshotId", column_type = "Text")]
    pub previous_snapshot_id: Option<String>,
    /// Run
    #[sea_orm(column_name = "runId", column_type = "Text")]
    pub run_id: String,
    #[sea_orm(column_name = "runStatus")]
    pub run_status: TaskRunStatus,
    #[sea_orm(column_name = "batchId", column_type = "Text")]
    pub batch_id: Option<String>,
    /// This is the current run attempt number. Users can define how many attempts they want for a run.
    #[sea_orm(column_name = "attemptNumber")]
    pub attempt_number: Option<i32>,
    /// Environment
    #[sea_orm(column_name = "environmentId", column_type = "Text")]
    pub environment_id: String,
    #[sea_orm(column_name = "environmentType")]
    pub environment_type: RuntimeEnvironmentType,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    /// An array of waitpoint IDs in the correct order, used for batches
    #[sea_orm(
        column_name = "completedWaitpointOrder",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub completed_waitpoint_order: Vec<String>,
    /// Checkpoint
    #[sea_orm(column_name = "checkpointId", column_type = "Text")]
    pub checkpoint_id: Option<String>,
    /// Worker
    #[sea_orm(column_name = "workerId", column_type = "Text")]
    pub worker_id: Option<String>,
    #[sea_orm(column_name = "runnerId", column_type = "Text")]
    pub runner_id: Option<String>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(
        column_name = "lastHeartbeatAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub last_heartbeat_at: Option<DateTime>,
    /// Metadata used by various systems in the run engine
    #[sea_orm(column_type = "JsonBinary")]
    pub metadata: Option<Json>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::task_run::Entity",
        from = "Column::RunId",
        to = "super::task_run::Column::Id",
        fk_name = "TaskRunExecutionSnapshot_runId_fkey"
    )]
    Run,
    #[sea_orm(
        belongs_to = "super::task_run_checkpoint::Entity",
        from = "Column::CheckpointId",
        to = "super::task_run_checkpoint::Column::Id",
        fk_name = "TaskRunExecutionSnapshot_checkpointId_fkey"
    )]
    Checkpoint,
}

impl Related<super::task_run::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Run.def()
    }
}

impl Related<super::task_run_checkpoint::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Checkpoint.def()
    }
}

impl Related<super::waitpoint::Entity> for Entity {
    fn to() -> RelationDef {
        super::completed_waitpoints::Relation::Waitpoint.def()
    }
    fn via() -> Option<RelationDef> {
        Some(
            super::completed_waitpoints::Relation::TaskRunExecutionSnapshot
                .def()
                .rev(),
        )
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
