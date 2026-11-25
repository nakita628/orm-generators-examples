use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "OAuthAuthorization")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_name = "oauthAppId", column_type = "Text")]
    pub oauth_app_id: String,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    #[sea_orm(column_name = "organizationId", column_type = "Text")]
    pub organization_id: String,
    #[sea_orm(column_name = "accessToken", column_type = "Text")]
    pub access_token: Option<String>,
    #[sea_orm(column_name = "authorizationCode", column_type = "Text")]
    pub authorization_code: Option<String>,
    #[sea_orm(
        column_name = "codeExpiresAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub code_expires_at: Option<DateTime>,
    #[sea_orm(column_name = "codeChallenge", column_type = "Text")]
    pub code_challenge: Option<String>,
    #[sea_orm(column_name = "codeChallengeMethod", column_type = "Text")]
    pub code_challenge_method: Option<String>,
    #[sea_orm(column_name = "redirectUri", column_type = "Text")]
    pub redirect_uri: Option<String>,
    #[sea_orm(column_name = "revokedAt", column_type = "custom(\"timestamp(3)\")")]
    pub revoked_at: Option<DateTime>,
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
        belongs_to = "super::oauth_app::Entity",
        from = "Column::OauthAppId",
        to = "super::oauth_app::Column::Id",
        fk_name = "OAuthAuthorization_oauthAppId_fkey"
    )]
    OauthApp,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        fk_name = "OAuthAuthorization_userId_fkey"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        fk_name = "OAuthAuthorization_organizationId_fkey"
    )]
    Organization,
}

impl Related<super::oauth_app::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::OauthApp.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            id: Set(uuid::Uuid::new_v4().to_string()),
            ..ActiveModelTrait::default()
        }
    }

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
