use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Workflow")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_type = "Text")]
    pub icon: String,
    #[sea_orm(column_type = "Text")]
    pub description: String,
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
    #[sea_orm(column_name = "latestVersionId", column_type = "Text")]
    pub latest_version_id: Option<String>,
    #[sea_orm(column_name = "currentVersionId", column_type = "Text")]
    pub current_version_id: Option<String>,
    #[sea_orm(column_name = "publishedId", column_type = "Text")]
    pub published_id: Option<String>,
    #[sea_orm(column_name = "publishedById", column_type = "Text")]
    pub published_by_id: Option<String>,
    #[sea_orm(column_name = "archivedAt", column_type = "custom(\"timestamp(3)\")")]
    pub archived_at: Option<DateTime>,
    #[sea_orm(column_name = "isEvaluator", default_value = false)]
    pub is_evaluator: bool,
    #[sea_orm(column_name = "isComponent", default_value = false)]
    pub is_component: bool,
    #[sea_orm(column_name = "copiedFromWorkflowId", column_type = "Text")]
    pub copied_from_workflow_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        fk_name = "Workflow_projectId_fkey"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::workflow_version::Entity",
        from = "Column::LatestVersionId",
        to = "super::workflow_version::Column::Id",
        on_update = "Restrict",
        on_delete = "Restrict",
        fk_name = "Workflow_latestVersionId_fkey"
    )]
    LatestVersion,
    #[sea_orm(
        belongs_to = "super::workflow_version::Entity",
        from = "Column::CurrentVersionId",
        to = "super::workflow_version::Column::Id",
        on_update = "Restrict",
        on_delete = "Restrict",
        fk_name = "Workflow_currentVersionId_fkey"
    )]
    CurrentVersion,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::PublishedById",
        to = "super::user::Column::Id",
        fk_name = "Workflow_publishedById_fkey"
    )]
    PublishedBy,
    #[sea_orm(
        belongs_to = "super::workflow::Entity",
        from = "Column::CopiedFromWorkflowId",
        to = "super::workflow::Column::Id",
        on_update = "Restrict",
        on_delete = "Restrict",
        fk_name = "Workflow_copiedFromWorkflowId_fkey"
    )]
    CopiedFrom,
    #[sea_orm(
        has_many = "super::workflow_version::Entity",
        from = "Column::Id",
        to = "super::workflow_version::Column::WorkflowId"
    )]
    Versions,
    #[sea_orm(has_many = "super::experiment::Entity")]
    Experiments,
    #[sea_orm(
        has_many = "super::workflow::Entity",
        from = "Column::Id",
        to = "super::workflow::Column::CopiedFromWorkflowId"
    )]
    CopiedWorkflows,
    #[sea_orm(has_many = "super::agent::Entity")]
    Agents,
    #[sea_orm(has_many = "super::evaluator::Entity")]
    Evaluators,
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::workflow_version::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::LatestVersion.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::PublishedBy.def()
    }
}

impl Related<super::workflow::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CopiedFrom.def()
    }
}

impl Related<super::experiment::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Experiments.def()
    }
}

impl Related<super::agent::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Agents.def()
    }
}

impl Related<super::evaluator::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Evaluators.def()
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
