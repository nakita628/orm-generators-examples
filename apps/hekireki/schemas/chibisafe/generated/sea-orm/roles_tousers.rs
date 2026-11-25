use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "_rolesTousers")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false, column_name = "A")]
    pub roles_id: i32,
    #[sea_orm(primary_key, auto_increment = false, column_name = "B")]
    pub users_id: i32,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::roles::Entity",
        from = "Column::RolesId",
        to = "super::roles::Column::Id",
        fk_name = "_rolesTousers_A_fkey"
    )]
    Roles,
    #[sea_orm(
        belongs_to = "super::users::Entity",
        from = "Column::UsersId",
        to = "super::users::Column::Id",
        fk_name = "_rolesTousers_B_fkey"
    )]
    Users,
}

impl ActiveModelBehavior for ActiveModel {}
