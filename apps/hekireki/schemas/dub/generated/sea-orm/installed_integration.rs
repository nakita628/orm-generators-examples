use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "InstalledIntegration")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub id: String,
    #[sea_orm(column_name = "userId", column_type = "String(StringLen::N(191))")]
    pub user_id: String,
    #[sea_orm(
        column_name = "integrationId",
        column_type = "String(StringLen::N(191))"
    )]
    pub integration_id: String,
    #[sea_orm(column_name = "projectId", column_type = "String(StringLen::N(191))")]
    pub project_id: String,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"datetime(3)\")",
        default_expr = "Expr::cust(\"CURRENT_TIMESTAMP(3)\")"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"datetime(3)\")")]
    pub updated_at: DateTime,
    pub credentials: Option<Json>,
    pub settings: Option<Json>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "InstalledIntegration_userId_fkey"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::integration::Entity",
        from = "Column::IntegrationId",
        to = "super::integration::Column::Id",
        on_delete = "Cascade",
        fk_name = "InstalledIntegration_integrationId_fkey"
    )]
    Integration,
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_delete = "Cascade",
        fk_name = "InstalledIntegration_projectId_fkey"
    )]
    Project,
    #[sea_orm(has_many = "super::oauth_refresh_token::Entity")]
    RefreshTokens,
    #[sea_orm(has_many = "super::restricted_token::Entity")]
    AccessTokens,
    #[sea_orm(has_many = "super::webhook::Entity")]
    Webhooks,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::integration::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Integration.def()
    }
}

impl Related<super::project::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Project.def()
    }
}

impl Related<super::oauth_refresh_token::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::RefreshTokens.def()
    }
}

impl Related<super::restricted_token::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AccessTokens.def()
    }
}

impl Related<super::webhook::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Webhooks.def()
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
