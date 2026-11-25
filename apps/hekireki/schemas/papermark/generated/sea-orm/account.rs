use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Account")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    #[sea_orm(column_type = "Text")]
    pub r#type: String,
    #[sea_orm(column_type = "Text")]
    pub provider: String,
    #[sea_orm(column_name = "providerAccountId", column_type = "Text")]
    pub provider_account_id: String,
    #[sea_orm(column_type = "Text")]
    pub refresh_token: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub access_token: Option<String>,
    pub expires_at: Option<i32>,
    #[sea_orm(column_type = "Text")]
    pub token_type: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub scope: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub id_token: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub session_state: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "Account_userId_fkey"
    )]
    User,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
