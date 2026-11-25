use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Evaluator")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_type = "Text")]
    pub slug: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub r#type: String,
    #[sea_orm(column_type = "JsonBinary")]
    pub config: Json,
    #[sea_orm(column_name = "workflowId", column_type = "Text")]
    pub workflow_id: Option<String>,
    #[sea_orm(column_name = "copiedFromEvaluatorId", column_type = "Text")]
    pub copied_from_evaluator_id: Option<String>,
    #[sea_orm(column_name = "archivedAt", column_type = "custom(\"timestamp(3)\")")]
    pub archived_at: Option<DateTime>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(
        column_name = "updatedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "createdById", column_type = "Text")]
    pub created_by_id: Option<String>,
    #[sea_orm(column_name = "updatedById", column_type = "Text")]
    pub updated_by_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        fk_name = "Evaluator_projectId_fkey"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::workflow::Entity",
        from = "Column::WorkflowId",
        to = "super::workflow::Column::Id",
        fk_name = "Evaluator_workflowId_fkey"
    )]
    Workflow,
    #[sea_orm(
        belongs_to = "super::evaluator::Entity",
        from = "Column::CopiedFromEvaluatorId",
        to = "super::evaluator::Column::Id",
        on_update = "Restrict",
        on_delete = "Restrict",
        fk_name = "Evaluator_copiedFromEvaluatorId_fkey"
    )]
    CopiedFrom,
    #[sea_orm(
        has_many = "super::evaluator::Entity",
        from = "Column::Id",
        to = "super::evaluator::Column::CopiedFromEvaluatorId"
    )]
    CopiedEvaluators,
    #[sea_orm(has_many = "super::monitor::Entity")]
    Monitors,
    #[sea_orm(has_many = "super::gateway_guardrail::Entity")]
    GatewayGuardrails,
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::workflow::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Workflow.def()
    }
}

impl Related<super::evaluator::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CopiedFrom.def()
    }
}

impl Related<super::monitor::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Monitors.def()
    }
}

impl Related<super::gateway_guardrail::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::GatewayGuardrails.def()
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
