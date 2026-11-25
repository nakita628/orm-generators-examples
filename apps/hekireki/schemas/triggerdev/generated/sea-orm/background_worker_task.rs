use super::task_execution_mode::TaskExecutionMode;
use super::task_trigger_source::TaskTriggerSource;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "BackgroundWorkerTask")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub slug: String,
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "friendlyId", column_type = "Text")]
    pub friendly_id: String,
    #[sea_orm(column_name = "filePath", column_type = "Text")]
    pub file_path: String,
    #[sea_orm(column_name = "exportName", column_type = "Text")]
    pub export_name: Option<String>,
    #[sea_orm(column_name = "workerId", column_type = "Text")]
    pub worker_id: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_name = "fileId", column_type = "Text")]
    pub file_id: Option<String>,
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
    #[sea_orm(column_name = "queueConfig", column_type = "JsonBinary")]
    pub queue_config: Option<Json>,
    #[sea_orm(column_name = "retryConfig", column_type = "JsonBinary")]
    pub retry_config: Option<Json>,
    #[sea_orm(column_name = "machineConfig", column_type = "JsonBinary")]
    pub machine_config: Option<Json>,
    /// Gates declared on the task: other queues its runs must also hold a concurrency
    /// slot in while executing. Shape: [{ queue: string, concurrencyKey?: string }]
    #[sea_orm(column_type = "JsonBinary")]
    pub gates: Option<Json>,
    #[sea_orm(column_name = "queueId", column_type = "Text")]
    pub queue_id: Option<String>,
    #[sea_orm(column_name = "maxDurationInSeconds")]
    pub max_duration_in_seconds: Option<i32>,
    #[sea_orm(column_type = "Text")]
    pub ttl: Option<String>,
    #[sea_orm(column_name = "triggerSource", default_value = "STANDARD")]
    pub trigger_source: TaskTriggerSource,
    /// Extra task configuration JSON. Shape depends on triggerSource.
    /// AGENT: { type: "ai-sdk-chat" }
    #[sea_orm(column_type = "JsonBinary")]
    pub config: Option<Json>,
    #[sea_orm(column_name = "executionMode", default_value = "STANDARD")]
    pub execution_mode: TaskExecutionMode,
    /// RunExecutionPolicy from @trigger.dev/core. Null means the standard policy.
    #[sea_orm(column_name = "executionPolicy", column_type = "JsonBinary")]
    pub execution_policy: Option<Json>,
    #[sea_orm(column_name = "payloadSchema", column_type = "JsonBinary")]
    pub payload_schema: Option<Json>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::background_worker::Entity",
        from = "Column::WorkerId",
        to = "super::background_worker::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "BackgroundWorkerTask_workerId_fkey"
    )]
    Worker,
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "BackgroundWorkerTask_projectId_fkey"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::background_worker_file::Entity",
        from = "Column::FileId",
        to = "super::background_worker_file::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "BackgroundWorkerTask_fileId_fkey"
    )]
    File,
    #[sea_orm(
        belongs_to = "super::runtime_environment::Entity",
        from = "Column::RuntimeEnvironmentId",
        to = "super::runtime_environment::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "BackgroundWorkerTask_runtimeEnvironmentId_fkey"
    )]
    RuntimeEnvironment,
    #[sea_orm(
        belongs_to = "super::task_queue::Entity",
        from = "Column::QueueId",
        to = "super::task_queue::Column::Id",
        on_update = "Cascade",
        on_delete = "SetNull",
        fk_name = "BackgroundWorkerTask_queueId_fkey"
    )]
    Queue,
    #[sea_orm(has_many = "super::task_run_attempt::Entity")]
    Attempts,
    #[sea_orm(has_many = "super::task_run::Entity")]
    Runs,
}

impl Related<super::background_worker::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Worker.def()
    }
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::background_worker_file::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::File.def()
    }
}

impl Related<super::runtime_environment::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::RuntimeEnvironment.def()
    }
}

impl Related<super::task_queue::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Queue.def()
    }
}

impl Related<super::task_run_attempt::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Attempts.def()
    }
}

impl Related<super::task_run::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Runs.def()
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
