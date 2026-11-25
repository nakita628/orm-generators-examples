use super::workflow_run_log_status::WorkflowRunLogStatus;
use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "WorkflowRunLog")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "runId", column_type = "Text")]
    pub run_id: String,
    pub sequence: i32,
    #[sea_orm(column_name = "stepId", column_type = "Text")]
    pub step_id: String,
    #[sea_orm(column_name = "stepType", column_type = "Text")]
    pub step_type: String,
    #[sea_orm(default_value = "pending")]
    pub status: WorkflowRunLogStatus,
    /// [WorkflowRunLogInput]
    #[sea_orm(column_type = "JsonBinary", default_value = "{}")]
    pub input: Json,
    /// [WorkflowRunLogOutput]
    #[sea_orm(column_type = "JsonBinary", default_value = "{}")]
    pub output: Json,
    #[sea_orm(column_type = "Text")]
    pub error: Option<String>,
    #[sea_orm(column_name = "startedAt", column_type = "custom(\"timestamp(3)\")")]
    pub started_at: Option<DateTime>,
    #[sea_orm(column_name = "finishedAt", column_type = "custom(\"timestamp(3)\")")]
    pub finished_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::workflow_run::Entity",
        from = "Column::RunId",
        to = "super::workflow_run::Column::Id",
        on_delete = "Cascade",
        fk_name = "WorkflowRunLog_runId_fkey"
    )]
    Run,
}

impl Related<super::workflow_run::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Run.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
