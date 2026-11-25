use super::step_execution_status::StepExecutionStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "workflow_step_executions")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "executionId", column_type = "Text")]
    pub execution_id: String,
    #[sea_orm(column_name = "stepId", column_type = "Text")]
    pub step_id: String,
    #[sea_orm(default_value = "PENDING")]
    pub status: StepExecutionStatus,
    #[sea_orm(column_name = "scheduledFor", column_type = "custom(\"timestamp(3)\")")]
    pub scheduled_for: Option<DateTime>,
    #[sea_orm(column_name = "executeAfter", column_type = "custom(\"timestamp(3)\")")]
    pub execute_after: Option<DateTime>,
    #[sea_orm(column_type = "JsonBinary")]
    pub output: Option<Json>,
    #[sea_orm(column_type = "Text")]
    pub error: Option<String>,
    #[sea_orm(column_name = "startedAt", column_type = "custom(\"timestamp(3)\")")]
    pub started_at: Option<DateTime>,
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
        belongs_to = "super::workflow_execution::Entity",
        from = "Column::ExecutionId",
        to = "super::workflow_execution::Column::Id",
        on_delete = "Cascade",
        fk_name = "workflow_step_executions_executionId_fkey"
    )]
    Execution,
    #[sea_orm(
        belongs_to = "super::workflow_step::Entity",
        from = "Column::StepId",
        to = "super::workflow_step::Column::Id",
        on_delete = "Cascade",
        fk_name = "workflow_step_executions_stepId_fkey"
    )]
    Step,
    #[sea_orm(has_many = "super::email::Entity")]
    Emails,
}

impl Related<super::workflow_execution::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Execution.def()
    }
}

impl Related<super::workflow_step::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Step.def()
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
        if insert && self.created_at.is_not_set() {
            self.created_at = Set(now.naive_utc());
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
