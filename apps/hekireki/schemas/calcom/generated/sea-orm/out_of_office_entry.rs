use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "OutOfOfficeEntry")]
pub struct Model {
    #[sea_orm(primary_key)]
    pub id: i32,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub uuid: String,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub start: DateTime,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub end: DateTime,
    #[sea_orm(column_type = "Text")]
    pub notes: Option<String>,
    #[sea_orm(column_name = "showNotePublicly", default_value = false)]
    pub show_note_publicly: bool,
    #[sea_orm(column_name = "userId")]
    pub user_id: i32,
    #[sea_orm(column_name = "toUserId")]
    pub to_user_id: Option<i32>,
    #[sea_orm(column_name = "reasonId")]
    pub reason_id: Option<i32>,
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
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "OutOfOfficeEntry_userId_fkey"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::ToUserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "OutOfOfficeEntry_toUserId_fkey"
    )]
    ToUser,
    #[sea_orm(
        belongs_to = "super::out_of_office_reason::Entity",
        from = "Column::ReasonId",
        to = "super::out_of_office_reason::Column::Id",
        on_delete = "SetNull",
        fk_name = "OutOfOfficeEntry_reasonId_fkey"
    )]
    Reason,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::out_of_office_reason::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Reason.def()
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
