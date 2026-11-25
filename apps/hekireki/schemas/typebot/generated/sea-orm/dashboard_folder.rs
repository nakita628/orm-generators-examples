use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "DashboardFolder")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
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
    #[sea_orm(column_type = "Text")]
    pub name: String,
    #[sea_orm(column_name = "parentFolderId", column_type = "Text")]
    pub parent_folder_id: Option<String>,
    #[sea_orm(column_name = "workspaceId", column_type = "Text")]
    pub workspace_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::dashboard_folder::Entity",
        from = "Column::ParentFolderId",
        to = "super::dashboard_folder::Column::Id",
        fk_name = "DashboardFolder_parentFolderId_fkey"
    )]
    ParentFolder,
    #[sea_orm(
        belongs_to = "super::workspace::Entity",
        from = "Column::WorkspaceId",
        to = "super::workspace::Column::Id",
        on_delete = "Cascade",
        fk_name = "DashboardFolder_workspaceId_fkey"
    )]
    Workspace,
    #[sea_orm(
        has_many = "super::dashboard_folder::Entity",
        from = "Column::Id",
        to = "super::dashboard_folder::Column::ParentFolderId"
    )]
    ChildrenFolder,
    #[sea_orm(has_many = "super::typebot::Entity")]
    Typebots,
}

impl Related<super::dashboard_folder::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ParentFolder.def()
    }
}

impl Related<super::workspace::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Workspace.def()
    }
}

impl Related<super::typebot::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Typebots.def()
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
