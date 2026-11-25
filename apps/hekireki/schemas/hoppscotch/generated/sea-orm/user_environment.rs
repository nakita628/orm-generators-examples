use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "UserEnvironment")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "userUid", column_type = "Text")]
    pub user_uid: String,
    #[sea_orm(column_type = "Text")]
    pub name: Option<String>,
    #[sea_orm(column_type = "JsonBinary")]
    pub variables: Json,
    #[sea_orm(column_name = "isGlobal")]
    pub is_global: bool,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserUid",
        to = "super::user::Column::Uid",
        on_delete = "Cascade",
        fk_name = "UserEnvironment_userUid_fkey"
    )]
    User,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
