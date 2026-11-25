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
    #[sea_orm(column_name = "workflowId", column_type = "Text")]
    pub workflow_id: String,
    #[sea_orm(column_name = "workspaceId", column_type = "Text")]
    pub workspace_id: String,
    pub version: i32,
    /// [WorkflowExecutableDefinition]
    #[sea_orm(column_type = "JsonBinary")]
    pub definition: Json,
    #[sea_orm(
        column_name = "publishedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub published_at: DateTime,
    #[sea_orm(column_name = "publishedBy", column_type = "Text")]
    pub published_by: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::workflow::Entity",
        from = "(Column::WorkflowId, Column::WorkspaceId)",
        to = "(super::workflow::Column::Id, super::workflow::Column::WorkspaceId)",
        on_delete = "Cascade",
        fk_name = "WorkflowVersion_workflowId_workspaceId_fkey"
    )]
    Workflow,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::PublishedBy",
        to = "super::user::Column::Id",
        on_delete = "SetNull",
        fk_name = "WorkflowVersion_publishedBy_fkey"
    )]
    Publisher,
    #[sea_orm(has_many = "super::workflow_run::Entity")]
    Runs,
}

impl Related<super::workflow::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Workflow.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Publisher.def()
    }
}

impl Related<super::workflow_run::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Runs.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    async fn before_save<C>(mut self, _db: &C, insert: bool) -> Result<Self, DbErr>
    where
        C: ConnectionTrait,
    {
        let now = chrono::Utc::now().trunc_subsecs(3);
        if insert && self.published_at.is_not_set() {
            self.published_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
