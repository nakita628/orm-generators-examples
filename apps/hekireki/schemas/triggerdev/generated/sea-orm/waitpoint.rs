use super::waitpoint_status::WaitpointStatus;
use super::waitpoint_type::WaitpointType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// A Waitpoint blocks a run from continuing until it's completed
/// If there's a waitpoint blocking a run, it shouldn't be in the queue
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Waitpoint")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "friendlyId", column_type = "Text")]
    pub friendly_id: String,
    pub r#type: WaitpointType,
    #[sea_orm(default_value = "PENDING")]
    pub status: WaitpointStatus,
    #[sea_orm(column_name = "completedAt", column_type = "custom(\"timestamp(3)\")")]
    pub completed_at: Option<DateTime>,
    /// If it's an Event type waitpoint, this is the event. It can also be provided for the DATETIME type
    #[sea_orm(column_name = "idempotencyKey", column_type = "Text")]
    pub idempotency_key: String,
    /// If this is true then we can show it in the dashboard/return it from the SDK
    #[sea_orm(column_name = "userProvidedIdempotencyKey")]
    pub user_provided_idempotency_key: bool,
    /// If there's a user provided idempotency key, this is the time it expires at
    #[sea_orm(
        column_name = "idempotencyKeyExpiresAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub idempotency_key_expires_at: Option<DateTime>,
    /// If an idempotencyKey is no longer active, we store it here and generate a new one for the idempotencyKey field.
    /// Clearing an idempotencyKey is useful for debounce or cancelling child runs.
    /// This is a workaround because Prisma doesn't support partial indexes.
    #[sea_orm(column_name = "inactiveIdempotencyKey", column_type = "Text")]
    pub inactive_idempotency_key: Option<String>,
    /// If it's a RUN type waitpoint, this is the associated run
    #[sea_orm(unique)]
    #[sea_orm(column_name = "completedByTaskRunId", column_type = "Text")]
    pub completed_by_task_run_id: Option<String>,
    /// If it's a DATETIME type waitpoint, this is the date.
    /// If it's a MANUAL waitpoint, this can be set as the `timeout`.
    #[sea_orm(
        column_name = "completedAfter",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub completed_after: Option<DateTime>,
    /// If it's a BATCH type waitpoint, this is the associated batch
    #[sea_orm(column_name = "completedByBatchId", column_type = "Text")]
    pub completed_by_batch_id: Option<String>,
    /// When completed, an output can be stored here
    #[sea_orm(column_type = "Text")]
    pub output: Option<String>,
    #[sea_orm(
        column_name = "outputType",
        column_type = "Text",
        default_value = "application/json"
    )]
    pub output_type: String,
    #[sea_orm(column_name = "outputIsError", default_value = false)]
    pub output_is_error: bool,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_name = "environmentId", column_type = "Text")]
    pub environment_id: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    /// Denormized column that holds the raw tags
    /// Denormalized column that holds the raw tags
    #[sea_orm(column_type = "Array(RcOrArc::new(ColumnType::Text))")]
    pub tags: Vec<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::task_run::Entity",
        from = "Column::CompletedByTaskRunId",
        to = "super::task_run::Column::Id",
        on_delete = "SetNull",
        fk_name = "Waitpoint_completedByTaskRunId_fkey"
    )]
    CompletedByTaskRun,
    #[sea_orm(
        belongs_to = "super::batch_task_run::Entity",
        from = "Column::CompletedByBatchId",
        to = "super::batch_task_run::Column::Id",
        on_delete = "SetNull",
        fk_name = "Waitpoint_completedByBatchId_fkey"
    )]
    CompletedByBatch,
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "Waitpoint_projectId_fkey"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::runtime_environment::Entity",
        from = "Column::EnvironmentId",
        to = "super::runtime_environment::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "Waitpoint_environmentId_fkey"
    )]
    Environment,
    #[sea_orm(has_many = "super::task_run_waitpoint::Entity")]
    BlockingTaskRuns,
}

impl Related<super::task_run::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CompletedByTaskRun.def()
    }
}

impl Related<super::batch_task_run::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CompletedByBatch.def()
    }
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::runtime_environment::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Environment.def()
    }
}

impl Related<super::task_run_waitpoint::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::BlockingTaskRuns.def()
    }
}

impl Related<super::task_run_execution_snapshot::Entity> for Entity {
    fn to() -> RelationDef {
        super::completed_waitpoints::Relation::TaskRunExecutionSnapshot.def()
    }
    fn via() -> Option<RelationDef> {
        Some(super::completed_waitpoints::Relation::Waitpoint.def().rev())
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
