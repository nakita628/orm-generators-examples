use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Integration")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub id: String,
    #[sea_orm(column_name = "userId", column_type = "String(StringLen::N(191))")]
    pub user_id: Option<String>,
    #[sea_orm(column_name = "projectId", column_type = "String(StringLen::N(191))")]
    pub project_id: String,
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub name: String,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub slug: String,
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub description: Option<String>,
    #[sea_orm(column_type = "custom(\"longtext\")")]
    pub readme: Option<String>,
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub developer: String,
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub website: String,
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub logo: Option<String>,
    pub screenshots: Option<Json>,
    #[sea_orm(default_value = false)]
    pub verified: bool,
    #[sea_orm(column_name = "installUrl", column_type = "Text")]
    pub install_url: Option<String>,
    #[sea_orm(column_name = "guideUrl", column_type = "Text")]
    pub guide_url: Option<String>,
    #[sea_orm(column_type = "String(StringLen::N(191))")]
    pub category: Option<String>,
    #[sea_orm(column_name = "comingSoon", default_value = false)]
    pub coming_soon: bool,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"datetime(3)\")",
        default_expr = "Expr::cust(\"CURRENT_TIMESTAMP(3)\")"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"datetime(3)\")")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "SetNull",
        fk_name = "Integration_userId_fkey"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::project::Entity",
        from = "Column::ProjectId",
        to = "super::project::Column::Id",
        on_delete = "Cascade",
        fk_name = "Integration_projectId_fkey"
    )]
    Project,
    #[sea_orm(has_many = "super::installed_integration::Entity")]
    Installations,
    #[sea_orm(has_one = "super::oauth_app::Entity")]
    OAuthApp,
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
        Relation::Installations.def()
    }
}

impl Related<super::oauth_app::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::OAuthApp.def()
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
