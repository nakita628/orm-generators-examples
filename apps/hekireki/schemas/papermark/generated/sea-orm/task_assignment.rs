use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "TaskAssignment")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "taskId", column_type = "Text")]
    pub task_id: String,
    #[sea_orm(column_name = "viewerId", column_type = "Text")]
    pub viewer_id: Option<String>,
    #[sea_orm(column_name = "groupId", column_type = "Text")]
    pub group_id: Option<String>,
    #[sea_orm(column_name = "linkId", column_type = "Text")]
    pub link_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub email: Option<String>,
    #[sea_orm(column_name = "notifiedAt", column_type = "custom(\"timestamp(3)\")")]
    pub notified_at: Option<DateTime>,
    #[sea_orm(
        column_name = "lastReminderAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub last_reminder_at: Option<DateTime>,
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
        belongs_to = "super::task::Entity",
        from = "Column::TaskId",
        to = "super::task::Column::Id",
        on_delete = "Cascade",
        fk_name = "TaskAssignment_taskId_fkey"
    )]
    Task,
    #[sea_orm(
        belongs_to = "super::viewer::Entity",
        from = "Column::ViewerId",
        to = "super::viewer::Column::Id",
        on_delete = "Cascade",
        fk_name = "TaskAssignment_viewerId_fkey"
    )]
    Viewer,
    #[sea_orm(
        belongs_to = "super::viewer_group::Entity",
        from = "Column::GroupId",
        to = "super::viewer_group::Column::Id",
        on_delete = "Cascade",
        fk_name = "TaskAssignment_groupId_fkey"
    )]
    Group,
    #[sea_orm(
        belongs_to = "super::link::Entity",
        from = "Column::LinkId",
        to = "super::link::Column::Id",
        on_delete = "Cascade",
        fk_name = "TaskAssignment_linkId_fkey"
    )]
    Link,
}

impl Related<super::task::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Task.def()
    }
}

impl Related<super::viewer::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Viewer.def()
    }
}

impl Related<super::viewer_group::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Group.def()
    }
}

impl Related<super::link::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Link.def()
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
