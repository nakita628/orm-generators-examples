use sea_orm::entity::prelude::*;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "account")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "accountId", column_type = "Text")]
    pub account_id: String,
    #[sea_orm(column_name = "providerId", column_type = "Text")]
    pub provider_id: String,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    #[sea_orm(column_name = "accessToken", column_type = "Text")]
    pub access_token: Option<String>,
    #[sea_orm(column_name = "refreshToken", column_type = "Text")]
    pub refresh_token: Option<String>,
    #[sea_orm(column_name = "idToken", column_type = "Text")]
    pub id_token: Option<String>,
    #[sea_orm(
        column_name = "accessTokenExpiresAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub access_token_expires_at: Option<DateTime>,
    #[sea_orm(
        column_name = "refreshTokenExpiresAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub refresh_token_expires_at: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub scope: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub password: Option<String>,
    #[sea_orm(column_name = "createdAt", column_type = "custom(\"timestamp(3)\")")]
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
        fk_name = "account_userId_fkey"
    )]
    User,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl ActiveModelBehavior for ActiveModel {}
