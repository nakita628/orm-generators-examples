use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

/// Per-call data kept beside a workflow (ADR 0048): one row per source, `trigger` for the Webhook
/// trigger's last payload (plus the last call's time and IP), or an HTTP request step id for its last
/// Test request response. Written on every call, so it lives here and never bumps `Workflow.updatedAt`.
#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "WorkflowSample")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "workflowId", column_type = "Text")]
    pub workflow_id: String,
    #[sea_orm(column_name = "workspaceId", column_type = "Text")]
    pub workspace_id: String,
    #[sea_orm(column_type = "Text")]
    pub source: String,
    /// [WorkflowSampleData]
    #[sea_orm(column_type = "JsonBinary")]
    pub sample: Json,
    #[sea_orm(column_name = "lastCalledAt", column_type = "custom(\"timestamp(3)\")")]
    pub last_called_at: Option<DateTime>,
    #[sea_orm(column_name = "lastCallerIp", column_type = "Text")]
    pub last_caller_ip: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::workflow::Entity",
        from = "(Column::WorkflowId, Column::WorkspaceId)",
        to = "(super::workflow::Column::Id, super::workflow::Column::WorkspaceId)",
        on_delete = "Cascade",
        fk_name = "WorkflowSample_workflowId_workspaceId_fkey"
    )]
    Workflow,
}

impl Related<super::workflow::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Workflow.def()
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
