use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "RestrictedToken")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub id: String,
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub name: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "hashedKey", column_type = "String(StringLen::N(191))")]
    pub hashed_key: String,
    #[sea_orm(column_name = "partialKey", column_type = "String(StringLen::N(191))")]
    pub partial_key: String,
    #[sea_orm(column_type = "Text")]
    pub scopes: Option<String>,
    #[sea_orm(column_type = "custom(\"datetime(3)\")")]
    pub expires: Option<DateTime>,
    #[sea_orm(column_name = "lastUsed", column_type = "custom(\"datetime(3)\")")]
    pub last_used: Option<DateTime>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"datetime(3)\")",
        default_expr = "Expr::cust(\"CURRENT_TIMESTAMP(3)\")"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"datetime(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "userId", column_type = "String(StringLen::N(191))")]
    pub user_id: String,
    #[sea_orm(column_name = "projectId", column_type = "String(StringLen::N(191))")]
    pub project_id: String,
    #[sea_orm(
        column_name = "installationId",
        column_type = "String(StringLen::N(191))"
    )]
    pub installation_id: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "RestrictedToken_userId_fkey"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_delete = "Cascade",
        fk_name = "RestrictedToken_projectId_fkey"
    )]
    Project,
    #[sea_orm(
        belongs_to = "super::installed_integration::Entity",
        from = "Column::InstallationId",
        to = "super::installed_integration::Column::Id",
        on_delete = "Cascade",
        fk_name = "RestrictedToken_installationId_fkey"
    )]
    InstalledIntegration,
    #[sea_orm(has_many = "super::oauth_refresh_token::Entity")]
    RefreshTokens,
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

impl Related<super::installed_integration::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::InstalledIntegration.def()
    }
}

impl Related<super::oauth_refresh_token::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::RefreshTokens.def()
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
