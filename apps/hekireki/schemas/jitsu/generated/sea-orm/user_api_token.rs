use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "UserApiToken")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    /// @zod.custom(z.coerce.date())
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    /// @zod.custom(z.coerce.date())
    #[sea_orm(
        column_name = "updatedAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub updated_at: DateTime,
    /// @zod.custom(z.coerce.date())
    #[sea_orm(column_name = "lastUsed", column_type = "custom(\"timestamp(3)\")")]
    pub last_used: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub hint: String,
    #[sea_orm(column_type = "Text")]
    pub hash: String,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    #[sea_orm(column_type = "Text")]
    pub r#type: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub name: Option<String>,
    /// @zod.custom(z.coerce.date())
    #[sea_orm(column_name = "expiresAt", column_type = "custom(\"timestamp(3)\")")]
    pub expires_at: Option<DateTime>,
    #[sea_orm(column_name = "workspaceId", column_type = "Text")]
    pub workspace_id: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub role: Option<String>,
    #[sea_orm(column_name = "oauthClientId", column_type = "Text")]
    pub oauth_client_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user_profile::Entity",
        from = "Column::UserId",
        to = "super::user_profile::Column::Id",
        fk_name = "UserApiToken_userId_fkey"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::oauth_client::Entity",
        from = "Column::OauthClientId",
        to = "super::oauth_client::Column::Id",
        fk_name = "UserApiToken_oauthClientId_fkey"
    )]
    OauthClient,
    #[sea_orm(has_many = "super::oauth_access_token::Entity")]
    AccessTokens,
}

impl Related<super::user_profile::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::oauth_client::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::OauthClient.def()
    }
}

impl Related<super::oauth_access_token::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AccessTokens.def()
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
