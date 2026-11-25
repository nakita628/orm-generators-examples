use super::experiment_type::ExperimentType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Experiment")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub name: Option<String>,
    pub r#type: ExperimentType,
    #[sea_orm(column_type = "Text")]
    pub slug: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_name = "workflowId", column_type = "Text")]
    pub workflow_id: Option<String>,
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
    #[sea_orm(column_name = "archivedAt", column_type = "custom(\"timestamp(3)\")")]
    pub archived_at: Option<DateTime>,
    #[sea_orm(column_name = "workbenchState", column_type = "JsonBinary")]
    pub workbench_state: Option<Json>,
    #[sea_orm(column_name = "workbenchVersion", default_value = 0)]
    pub workbench_version: i32,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        fk_name = "Experiment_projectId_fkey"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::workflow::Entity",
        from = "Column::WorkflowId",
        to = "super::workflow::Column::Id",
        fk_name = "Experiment_workflowId_fkey"
    )]
    Workflow,
    #[sea_orm(has_many = "super::batch_evaluation::Entity")]
    BatchEvaluations,
    #[sea_orm(has_one = "super::monitor::Entity")]
    Monitor,
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

impl Related<super::batch_evaluation::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::BatchEvaluations.def()
    }
}

impl Related<super::monitor::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Monitor.def()
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
