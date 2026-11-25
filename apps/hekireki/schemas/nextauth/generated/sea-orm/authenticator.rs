use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Authenticator")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(unique)]
    #[sea_orm(column_name = "credentialID", column_type = "Text")]
    pub credential_id: String,
    #[sea_orm(column_name = "providerAccountId", column_type = "Text")]
    pub provider_account_id: String,
    #[sea_orm(column_name = "credentialPublicKey", column_type = "Text")]
    pub credential_public_key: String,
    pub counter: i32,
    #[sea_orm(column_name = "credentialDeviceType", column_type = "Text")]
    pub credential_device_type: String,
    #[sea_orm(column_name = "credentialBackedUp")]
    pub credential_backed_up: bool,
    #[sea_orm(column_type = "Text")]
    pub transports: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "Authenticator_userId_fkey"
    )]
    User,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
