use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "WorkflowVersion")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub version: String,
    #[sea_orm(column_name = "commitMessage", column_type = "Text")]
    pub commit_message: String,
    #[sea_orm(column_name = "authorId", column_type = "Text")]
    pub author_id: String,
    #[sea_orm(column_name = "projectId", column_type = "Text")]
    pub project_id: String,
    #[sea_orm(column_name = "workflowId", column_type = "Text")]
    pub workflow_id: String,
    #[sea_orm(column_name = "parentId", column_type = "Text")]
    pub parent_id: Option<String>,
    #[sea_orm(column_name = "autoSaved", default_value = false)]
    pub auto_saved: bool,
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
    #[sea_orm(column_type = "JsonBinary")]
    pub dsl: Json,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::AuthorId",
        to = "super::user::Column::Id",
        fk_name = "WorkflowVersion_authorId_fkey"
    )]
    Author,
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        fk_name = "WorkflowVersion_projectId_fkey"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::workflow::Entity",
        from = "Column::WorkflowId",
        to = "super::workflow::Column::Id",
        fk_name = "WorkflowVersion_workflowId_fkey"
    )]
    Workflow,
    #[sea_orm(
        belongs_to = "super::workflow_version::Entity",
        from = "Column::ParentId",
        to = "super::workflow_version::Column::Id",
        on_update = "Restrict",
        on_delete = "Restrict",
        fk_name = "WorkflowVersion_parentId_fkey"
    )]
    Parent,
    #[sea_orm(
        has_many = "super::workflow_version::Entity",
        from = "Column::Id",
        to = "super::workflow_version::Column::ParentId"
    )]
    Children,
    #[sea_orm(
        has_many = "super::workflow::Entity",
        from = "Column::Id",
        to = "super::workflow::Column::LatestVersionId"
    )]
    WorkflowAsLatestVersion,
    #[sea_orm(
        has_many = "super::workflow::Entity",
        from = "Column::Id",
        to = "super::workflow::Column::CurrentVersionId"
    )]
    WorkflowAsCurrentVersion,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Author.def()
    }
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

impl Related<super::workflow_version::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Parent.def()
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
