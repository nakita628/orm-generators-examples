use super::graph_navigation::GraphNavigation;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "User")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
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
    #[sea_orm(
        column_name = "lastActivityAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub last_activity_at: DateTime,
    #[sea_orm(column_type = "Text")]
    pub name: Option<String>,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub email: Option<String>,
    #[sea_orm(
        column_name = "emailVerified",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub email_verified: Option<DateTime>,
    #[sea_orm(column_type = "Text")]
    pub image: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub company: Option<String>,
    #[sea_orm(column_name = "onboardingCategories", column_type = "JsonBinary")]
    pub onboarding_categories: Json,
    #[sea_orm(column_type = "Text")]
    pub referral: Option<String>,
    #[sea_orm(column_name = "graphNavigation")]
    pub graph_navigation: Option<GraphNavigation>,
    #[sea_orm(column_name = "preferredAppAppearance", column_type = "Text")]
    pub preferred_app_appearance: Option<String>,
    #[sea_orm(
        column_name = "displayedInAppNotifications",
        column_type = "JsonBinary"
    )]
    pub displayed_in_app_notifications: Option<Json>,
    #[sea_orm(column_name = "groupTitlesAutoGeneration", column_type = "JsonBinary")]
    pub group_titles_auto_generation: Option<Json>,
    #[sea_orm(column_name = "preferredLanguage", column_type = "Text")]
    pub preferred_language: Option<String>,
    #[sea_orm(
        column_name = "termsAcceptedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub terms_accepted_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(has_many = "super::account::Entity")]
    Accounts,
    #[sea_orm(has_many = "super::api_token::Entity")]
    ApiTokens,
    #[sea_orm(has_many = "super::collaborators_on_typebots::Entity")]
    CollaboratorsOnTypebots,
    #[sea_orm(has_many = "super::member_in_workspace::Entity")]
    Workspaces,
    #[sea_orm(has_many = "super::session::Entity")]
    Sessions,
    #[sea_orm(has_many = "super::banned_ip::Entity")]
    BannedIps,
    #[sea_orm(has_many = "super::user_credentials::Entity")]
    Credentials,
}

impl Related<super::account::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Accounts.def()
    }
}

impl Related<super::api_token::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ApiTokens.def()
    }
}

impl Related<super::collaborators_on_typebots::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CollaboratorsOnTypebots.def()
    }
}

impl Related<super::member_in_workspace::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Workspaces.def()
    }
}

impl Related<super::session::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Sessions.def()
    }
}

impl Related<super::banned_ip::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::BannedIps.def()
    }
}

impl Related<super::user_credentials::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Credentials.def()
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
        if insert && self.last_activity_at.is_not_set() {
            self.last_activity_at = Set(now.naive_utc());
        }
        Ok(self)
    }
}
