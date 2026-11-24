use super::workflow_run_status::WorkflowRunStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "WorkflowRun")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    pub created_at: DateTime,
    pub updated_at: DateTime,
    #[sea_orm(column_name = "workflowId")]
    pub workflow_id: String,
    #[sea_orm(column_name = "workspaceId")]
    pub workspace_id: String,
    #[sea_orm(column_name = "workflowVersionId")]
    pub workflow_version_id: Option<String>,
    #[sea_orm(column_name = "responseId")]
    pub response_id: Option<String>,
    #[sea_orm(default_value = "queued")]
    pub status: WorkflowRunStatus,
    #[sea_orm(column_name = "triggerType")]
    pub trigger_type: String,
    #[sea_orm(column_name = "surveyId")]
    pub survey_id: Option<String>,
    #[sea_orm(column_name = "isDryRun", default_value = false)]
    pub is_dry_run: bool,
    #[sea_orm(column_name = "idempotencyKey")]
    pub idempotency_key: Option<String>,
    #[sea_orm(default_value = 0)]
    pub attempt: i32,
    #[sea_orm(column_name = "nextAttemptAt")]
    pub next_attempt_at: Option<DateTime>,
    #[sea_orm(column_name = "lastErrorAt")]
    pub last_error_at: Option<DateTime>,
    /// [WorkflowTriggerRunPayload]
    #[sea_orm(column_name = "triggerPayload")]
    pub trigger_payload: Json,
    /// [WorkflowRunData]
    #[sea_orm(default_value = "{}")]
    pub data: Json,
    pub error: Option<String>,
    #[sea_orm(column_name = "dispatchedAt")]
    pub dispatched_at: Option<DateTime>,
    #[sea_orm(column_name = "startedAt")]
    pub started_at: Option<DateTime>,
    #[sea_orm(column_name = "finishedAt")]
    pub finished_at: Option<DateTime>,
    /// The caller's IP for a run started by a Webhook Secret URL call; shown in run detail.
    #[sea_orm(column_name = "callerIp")]
    pub caller_ip: Option<String>,
    /// When the run retention job cleared the run's content (ADR 0044): payload contents, run data, step
    /// input, output and errors, the run error and the caller IP. Null while the content is kept.
    #[sea_orm(column_name = "contentDeletedAt")]
    pub content_deleted_at: Option<DateTime>,
    /// The retention period in days that applied when the content was cleared ("Details deleted after N days").
    #[sea_orm(column_name = "contentRetentionDays")]
    pub content_retention_days: Option<i32>,
    /// The draft definition a Test run executed (ADR 0049), so its graph can be drawn after the draft
    /// changes. Null for live runs (they point at `workflowVersionId`) and for Test runs saved before it
    /// existed. Configuration like a version's definition, not run content: retention leaves it alone.
    /// [WorkflowExecutableDefinition]
    #[sea_orm(column_name = "testRunDefinition")]
    pub test_run_definition: Option<Json>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::workflow::Entity",
        from = "(Column::WorkflowId, Column::WorkspaceId)",
        to = "(super::workflow::Column::Id, super::workflow::Column::WorkspaceId)",
        on_delete = "Cascade"
    )]
    Workflow,
    #[sea_orm(
        belongs_to = "super::workspace::Entity",
        from = "Column::WorkspaceId",
        to = "super::workspace::Column::Id",
        on_delete = "Cascade"
    )]
    Workspace,
    #[sea_orm(
        belongs_to = "super::workflow_version::Entity",
        from = "Column::WorkflowVersionId",
        to = "super::workflow_version::Column::Id",
        on_delete = "SetNull"
    )]
    WorkflowVersion,
    #[sea_orm(
        belongs_to = "super::response::Entity",
        from = "Column::ResponseId",
        to = "super::response::Column::Id",
        on_delete = "Cascade"
    )]
    Response,
    #[sea_orm(has_many = "super::workflow_run_log::Entity")]
    Logs,
}

impl Related<super::workflow::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Workflow.def()
    }
}

impl Related<super::workspace::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Workspace.def()
    }
}

impl Related<super::workflow_version::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::WorkflowVersion.def()
    }
}

impl Related<super::response::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Response.def()
    }
}

impl Related<super::workflow_run_log::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Logs.def()
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
