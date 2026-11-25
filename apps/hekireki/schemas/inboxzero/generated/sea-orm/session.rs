use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Session")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "sessionToken", column_type = "Text")]
    pub session_token: String,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    #[sea_orm(column_type = "custom(\"timestamp(3)\")")]
    pub expires: DateTime,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "emailOtp", default_value = false)]
    pub email_otp: bool,
    #[sea_orm(column_name = "emailOtpVersion", default_value = 0)]
    pub email_otp_version: i32,
    #[sea_orm(column_name = "ipAddress", column_type = "Text")]
    pub ip_address: Option<String>,
    #[sea_orm(column_name = "userAgent", column_type = "Text")]
    pub user_agent: Option<String>,
    #[sea_orm(column_name = "activeOrganizationId", column_type = "Text")]
    pub active_organization_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "Session_userId_fkey"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::ActiveOrganizationId",
        to = "super::organization::Column::Id",
        on_delete = "SetNull",
        fk_name = "Session_activeOrganizationId_fkey"
    )]
    ActiveOrganization,
    #[sea_orm(has_many = "super::oauth_refresh_token::Entity")]
    OauthRefreshTokenRecords,
    #[sea_orm(has_many = "super::oauth_access_token::Entity")]
    OauthAccessTokenRecords,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ActiveOrganization.def()
    }
}

impl Related<super::oauth_refresh_token::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::OauthRefreshTokenRecords.def()
    }
}

impl Related<super::oauth_access_token::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::OauthAccessTokenRecords.def()
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
