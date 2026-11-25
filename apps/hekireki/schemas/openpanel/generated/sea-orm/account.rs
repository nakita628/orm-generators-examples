use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "accounts")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(
        column_type = "Text",
        default_expr = "Expr::cust(\"gen_random_uuid()\")"
    )]
    pub id: String,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    #[sea_orm(column_type = "Text")]
    pub email: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub provider: String,
    #[sea_orm(column_name = "providerId", column_type = "Text")]
    pub provider_id: Option<String>,
    #[sea_orm(column_name = "accountId", column_type = "Text")]
    pub account_id: Option<String>,
    #[sea_orm(column_name = "idToken", column_type = "Text")]
    pub id_token: Option<String>,
    #[sea_orm(column_name = "accessToken", column_type = "Text")]
    pub access_token: Option<String>,
    #[sea_orm(column_name = "refreshToken", column_type = "Text")]
    pub refresh_token: Option<String>,
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
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(
        column_name = "updatedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "accounts_userId_fkey"
    )]
    User,
    #[sea_orm(has_many = "super::reset_password::Entity")]
    ResetPasswords,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::reset_password::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ResetPasswords.def()
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
