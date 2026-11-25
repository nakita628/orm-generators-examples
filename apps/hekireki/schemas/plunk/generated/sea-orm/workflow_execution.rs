use super::workflow_execution_status::WorkflowExecutionStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "workflow_executions")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "workflowId", column_type = "Text")]
    pub workflow_id: String,
    #[sea_orm(column_name = "contactId", column_type = "Text")]
    pub contact_id: String,
    #[sea_orm(default_value = "RUNNING")]
    pub status: WorkflowExecutionStatus,
    #[sea_orm(column_name = "currentStepId", column_type = "Text")]
    pub current_step_id: Option<String>,
    #[sea_orm(column_name = "exitReason", column_type = "Text")]
    pub exit_reason: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub context: Option<Json>,
    #[sea_orm(
        column_name = "startedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub started_at: DateTime,
    #[sea_orm(column_name = "completedAt", column_type = "custom(\"timestamp(3)\")")]
    pub completed_at: Option<DateTime>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::workflow::Entity",
        from = "Column::WorkflowId",
        to = "super::workflow::Column::Id",
        on_delete = "Cascade",
        fk_name = "workflow_executions_workflowId_fkey"
    )]
    Workflow,
    #[sea_orm(
        belongs_to = "super::contact::Entity",
        from = "Column::ContactId",
        to = "super::contact::Column::Id",
        on_delete = "Cascade",
        fk_name = "workflow_executions_contactId_fkey"
    )]
    Contact,
    #[sea_orm(
        belongs_to = "super::workflow_step::Entity",
        from = "Column::CurrentStepId",
        to = "super::workflow_step::Column::Id",
        fk_name = "workflow_executions_currentStepId_fkey"
    )]
    CurrentStep,
    #[sea_orm(has_many = "super::workflow_step_execution::Entity")]
    StepExecutions,
    #[sea_orm(has_many = "super::email::Entity")]
    Emails,
}

impl Related<super::workflow::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Workflow.def()
    }
}

impl Related<super::contact::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Contact.def()
    }
}

impl Related<super::workflow_step::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CurrentStep.def()
    }
}

impl Related<super::workflow_step_execution::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::StepExecutions.def()
    }
}

impl Related<super::email::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Emails.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            id: Set(uuid::Uuid::new_v4().to_string()),
            ..ActiveModelTrait::default()
        }
    }

    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.started_at.is_not_set() {
            self.started_at = Set(now.naive_utc());
        }
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.naive_utc());
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
