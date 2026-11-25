use super::execution_status::ExecutionStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "WorkflowExecution")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "workflowId", column_type = "Text")]
    pub workflow_id: String,
    #[sea_orm(column_name = "visitorEmail", column_type = "Text")]
    pub visitor_email: Option<String>,
    #[sea_orm(column_name = "visitorIp", column_type = "Text")]
    pub visitor_ip: Option<String>,
    pub status: ExecutionStatus,
    #[sea_orm(
        column_name = "startedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub started_at: DateTime,
    #[sea_orm(column_name = "completedAt", column_type = "custom(\"timestamp(3)\")")]
    pub completed_at: Option<DateTime>,
    #[sea_orm(column_type = "JsonBinary")]
    pub result: Option<Json>,
    #[sea_orm(column_type = "JsonBinary")]
    pub metadata: Option<Json>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::workflow::Entity",
        from = "Column::WorkflowId",
        to = "super::workflow::Column::Id",
        on_delete = "Cascade",
        fk_name = "WorkflowExecution_workflowId_fkey"
    )]
    Workflow,
    #[sea_orm(has_many = "super::workflow_step_log::Entity")]
    StepLogs,
}

impl Related<super::workflow::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Workflow.def()
    }
}

impl Related<super::workflow_step_log::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::StepLogs.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.started_at.is_not_set() {
            self.started_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
