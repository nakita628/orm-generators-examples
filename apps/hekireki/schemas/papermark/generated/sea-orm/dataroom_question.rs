use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "DataroomQuestion")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text")]
    pub content: String,
    #[sea_orm(column_type = "Text")]
    pub category: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub priority: Option<String>,
    #[sea_orm(column_type = "Text", default_value = "unanswered")]
    pub status: String,
    #[sea_orm(column_name = "orderIndex")]
    pub order_index: Option<i32>,
    #[sea_orm(column_name = "importBatchId", column_type = "Text")]
    pub import_batch_id: Option<String>,
    #[sea_orm(column_name = "dataroomId", column_type = "Text")]
    pub dataroom_id: String,
    #[sea_orm(column_name = "teamId", column_type = "Text")]
    pub team_id: String,
    #[sea_orm(column_name = "createdByUserId", column_type = "Text")]
    pub created_by_user_id: Option<String>,
    #[sea_orm(column_name = "dataroomDocumentId", column_type = "Text")]
    pub dataroom_document_id: Option<String>,
    #[sea_orm(column_name = "documentPageNumber")]
    pub document_page_number: Option<i32>,
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
        belongs_to = "super::dataroom::Entity",
        from = "Column::DataroomId",
        to = "super::dataroom::Column::Id",
        on_delete = "Cascade",
        fk_name = "DataroomQuestion_dataroomId_fkey"
    )]
    Dataroom,
    #[sea_orm(
        belongs_to = "super::team::Entity",
        from = "Column::TeamId",
        to = "super::team::Column::Id",
        on_delete = "Cascade",
        fk_name = "DataroomQuestion_teamId_fkey"
    )]
    Team,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::CreatedByUserId",
        to = "super::user::Column::Id",
        on_delete = "SetNull",
        fk_name = "DataroomQuestion_createdByUserId_fkey"
    )]
    CreatedByUser,
    #[sea_orm(
        belongs_to = "super::dataroom_document::Entity",
        from = "Column::DataroomDocumentId",
        to = "super::dataroom_document::Column::Id",
        on_delete = "SetNull",
        fk_name = "DataroomQuestion_dataroomDocumentId_fkey"
    )]
    DataroomDocument,
    #[sea_orm(has_many = "super::dataroom_question_assignment::Entity")]
    Assignments,
    #[sea_orm(has_many = "super::conversation::Entity")]
    Conversations,
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

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CreatedByUser.def()
    }
}

impl Related<super::dataroom_document::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::DataroomDocument.def()
    }
}

impl Related<super::dataroom_question_assignment::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Assignments.def()
    }
}

impl Related<super::conversation::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Conversations.def()
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
