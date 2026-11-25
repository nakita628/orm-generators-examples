use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "TaskActivity")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "taskId", column_type = "Text")]
    pub task_id: String,
    #[sea_orm(column_type = "Text")]
    pub r#type: String,
    #[sea_orm(column_type = "Text")]
    pub comment: Option<String>,
    #[sea_orm(column_name = "fromStatus", column_type = "Text")]
    pub from_status: Option<String>,
    #[sea_orm(column_name = "toStatus", column_type = "Text")]
    pub to_status: Option<String>,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: Option<String>,
    #[sea_orm(column_name = "viewerId", column_type = "Text")]
    pub viewer_id: Option<String>,
    #[sea_orm(column_name = "viewId", column_type = "Text")]
    pub view_id: Option<String>,
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
        belongs_to = "super::task::Entity",
        from = "Column::TaskId",
        to = "super::task::Column::Id",
        on_delete = "Cascade",
        fk_name = "TaskActivity_taskId_fkey"
    )]
    Task,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "SetNull",
        fk_name = "TaskActivity_userId_fkey"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::viewer::Entity",
        from = "Column::ViewerId",
        to = "super::viewer::Column::Id",
        on_delete = "SetNull",
        fk_name = "TaskActivity_viewerId_fkey"
    )]
    Viewer,
    #[sea_orm(
        belongs_to = "super::view::Entity",
        from = "Column::ViewId",
        to = "super::view::Column::Id",
        on_delete = "SetNull",
        fk_name = "TaskActivity_viewId_fkey"
    )]
    View,
}

impl Related<super::task::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Task.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::viewer::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Viewer.def()
    }
}

impl Related<super::view::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::View.def()
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
