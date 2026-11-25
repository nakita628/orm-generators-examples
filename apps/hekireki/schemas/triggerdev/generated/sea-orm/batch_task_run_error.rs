use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Track individual run creation failures in batch processing (Run Engine v2)
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "BatchTaskRunError")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "batchTaskRunId", column_type = "Text")]
    pub batch_task_run_id: String,
    /// Which item in the batch (0-based index)
    pub index: i32,
    /// The task identifier that was being triggered
    #[sea_orm(column_name = "taskIdentifier", column_type = "Text")]
    pub task_identifier: String,
    /// The payload that failed (JSON, may be truncated)
    #[sea_orm(column_type = "Text")]
    pub payload: Option<String>,
    /// The options that were used
    #[sea_orm(column_type = "JsonBinary")]
    pub options: Option<Json>,
    /// Error message
    #[sea_orm(column_type = "Text")]
    pub error: String,
    /// Error code if available
    #[sea_orm(column_name = "errorCode", column_type = "Text")]
    pub error_code: Option<String>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::batch_task_run::Entity",
        from = "Column::BatchTaskRunId",
        to = "super::batch_task_run::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "BatchTaskRunError_batchTaskRunId_fkey"
    )]
    BatchTaskRun,
}

impl Related<super::batch_task_run::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::BatchTaskRun.def()
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
        Ok(self)
    }
}
