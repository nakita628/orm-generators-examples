use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "TaskRunWaitpoint")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "taskRunId", column_type = "Text")]
    pub task_run_id: String,
    #[sea_orm(column_name = "waitpointId", column_type = "Text")]
    pub waitpoint_id: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    /// This span id is completed when the waitpoint is completed. This is used with cached runs (idempotent)
    #[sea_orm(column_name = "spanIdToComplete", column_type = "Text")]
    pub span_id_to_complete: Option<String>,
    #[sea_orm(column_name = "batchId", column_type = "Text")]
    pub batch_id: Option<String>,
    #[sea_orm(column_name = "batchIndex")]
    pub batch_index: Option<i32>,
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
        belongs_to = "super::task_run::Entity",
        from = "Column::TaskRunId",
        to = "super::task_run::Column::Id",
        fk_name = "TaskRunWaitpoint_taskRunId_fkey"
    )]
    TaskRun,
    #[sea_orm(
        belongs_to = "super::waitpoint::Entity",
        from = "Column::WaitpointId",
        to = "super::waitpoint::Column::Id",
        fk_name = "TaskRunWaitpoint_waitpointId_fkey"
    )]
    Waitpoint,
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "TaskRunWaitpoint_projectId_fkey"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::batch_task_run::Entity",
        from = "Column::BatchId",
        to = "super::batch_task_run::Column::Id",
        fk_name = "TaskRunWaitpoint_batchId_fkey"
    )]
    Batch,
}

impl Related<super::task_run::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::TaskRun.def()
    }
}

impl Related<super::waitpoint::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Waitpoint.def()
    }
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::batch_task_run::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Batch.def()
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
