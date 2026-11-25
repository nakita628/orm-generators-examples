use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "OAuthCode")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub id: String,
    #[sea_orm(column_name = "clientId", column_type = "String(StringLen::N(191))")]
    pub client_id: String,
    #[sea_orm(column_name = "userId", column_type = "String(StringLen::N(191))")]
    pub user_id: String,
    #[sea_orm(column_name = "projectId", column_type = "String(StringLen::N(191))")]
    pub project_id: String,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub code: String,
    #[sea_orm(column_type = "Text")]
    pub scopes: Option<String>,
    #[sea_orm(column_name = "redirectUri", column_type = "String(StringLen::N(191))")]
    pub redirect_uri: String,
    #[sea_orm(
        column_name = "codeChallenge",
        column_type = "String(StringLen::N(191))"
    )]
    pub code_challenge: Option<String>,
    #[sea_orm(
        column_name = "codeChallengeMethod",
        column_type = "String(StringLen::N(191))"
    )]
    pub code_challenge_method: Option<String>,
    #[sea_orm(column_name = "expiresAt", column_type = "custom(\"datetime(3)\")")]
    pub expires_at: DateTime,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"datetime(3)\")",
        default_expr = "Expr::cust(\"CURRENT_TIMESTAMP(3)\")"
    )]
    pub created_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::oauth_app::Entity",
        from = "Column::ClientId",
        to = "super::oauth_app::Column::ClientId",
        on_delete = "Cascade",
        fk_name = "OAuthCode_clientId_fkey"
    )]
    OAuthApp,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "OAuthCode_userId_fkey"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_delete = "Cascade",
        fk_name = "OAuthCode_projectId_fkey"
    )]
    Project,
}

impl Related<super::oauth_app::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::OAuthApp.def()
    }
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
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
        Ok(self)
    }
}
