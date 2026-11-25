use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Task")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub title: String,
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    #[sea_orm(column_type = "Text", default_value = "TODO")]
    pub r#type: String,
    #[sea_orm(column_type = "Text", default_value = "OPEN")]
    pub status: String,
    #[sea_orm(column_type = "Text")]
    pub category: Option<String>,
    #[sea_orm(column_name = "dueDate", column_type = "custom(\"timestamp(3)\")")]
    pub due_date: Option<DateTime>,
    #[sea_orm(column_name = "orderIndex")]
    pub order_index: Option<i32>,
    #[sea_orm(column_name = "taskListId", column_type = "Text")]
    pub task_list_id: String,
    #[sea_orm(column_name = "dataroomId", column_type = "Text")]
    pub dataroom_id: String,
    #[sea_orm(column_name = "teamId", column_type = "Text")]
    pub team_id: String,
    #[sea_orm(column_name = "uploadFolderId", column_type = "Text")]
    pub upload_folder_id: Option<String>,
    #[sea_orm(column_name = "createdByUserId", column_type = "Text")]
    pub created_by_user_id: Option<String>,
    #[sea_orm(
        column_name = "lastActivityNotifiedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub last_activity_notified_at: Option<DateTime>,
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
        belongs_to = "super::task_list::Entity",
        from = "Column::TaskListId",
        to = "super::task_list::Column::Id",
        on_delete = "Cascade",
        fk_name = "Task_taskListId_fkey"
    )]
    TaskList,
    #[sea_orm(
        belongs_to = "super::dataroom::Entity",
        from = "Column::DataroomId",
        to = "super::dataroom::Column::Id",
        on_delete = "Cascade",
        fk_name = "Task_dataroomId_fkey"
    )]
    Dataroom,
    #[sea_orm(
        belongs_to = "super::team::Entity",
        from = "Column::TeamId",
        to = "super::team::Column::Id",
        on_delete = "Cascade",
        fk_name = "Task_teamId_fkey"
    )]
    Team,
    #[sea_orm(
        belongs_to = "super::dataroom_folder::Entity",
        from = "Column::UploadFolderId",
        to = "super::dataroom_folder::Column::Id",
        on_delete = "SetNull",
        fk_name = "Task_uploadFolderId_fkey"
    )]
    UploadFolder,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CreatedByUserId",
        to = "super::user::Column::Id",
        on_delete = "SetNull",
        fk_name = "Task_createdByUserId_fkey"
    )]
    CreatedByUser,
    #[sea_orm(has_many = "super::task_assignment::Entity")]
    Assignments,
    #[sea_orm(has_many = "super::task_activity::Entity")]
    Activities,
    #[sea_orm(has_many = "super::document_upload::Entity")]
    Uploads,
}

impl Related<super::task_list::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::TaskList.def()
    }
}

impl Related<super::dataroom::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Dataroom.def()
    }
}

impl Related<super::team::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Team.def()
    }
}

impl Related<super::dataroom_folder::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::UploadFolder.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CreatedByUser.def()
    }
}

impl Related<super::task_assignment::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Assignments.def()
    }
}

impl Related<super::task_activity::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Activities.def()
    }
}

impl Related<super::document_upload::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Uploads.def()
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
