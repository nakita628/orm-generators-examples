use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "_UserToUserLink")]
pub struct Model {
    #[sea_orm(
        primary_key,
        auto_increment = false,
        column_name = "A",
        column_type = "String(StringLen::N(191))"
    )]
    pub user_id: String,
    #[sea_orm(
        primary_key,
        auto_increment = false,
        column_name = "B",
        column_type = "String(StringLen::N(191))"
    )]
    pub user_link_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        fk_name = "_UserToUserLink_A_fkey"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::user_link::Entity",
        from = "Column::UserLinkId",
        to = "super::user_link::Column::Id",
        fk_name = "_UserToUserLink_B_fkey"
    )]
    UserLink,
}

impl ActiveModelBehavior for ActiveModel {}
