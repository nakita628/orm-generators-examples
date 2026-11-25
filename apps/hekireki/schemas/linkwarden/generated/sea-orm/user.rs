use super::ai_tagging_method::AiTaggingMethod;
use super::links_route_to::LinksRouteTo;
use super::theme::Theme;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "User")]
pub struct Model {
    #[sea_orm(primary_key)]
    pub id: i32,
    #[sea_orm(unique)]
    pub uuid: Uuid,
    #[sea_orm(column_type = "Text")]
    pub name: Option<String>,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub username: Option<String>,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub email: Option<String>,
    #[sea_orm(
        column_name = "emailVerified",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub email_verified: Option<DateTime>,
    #[sea_orm(column_name = "unverifiedNewEmail", column_type = "Text")]
    pub unverified_new_email: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub image: Option<String>,
    #[sea_orm(column_type = "Text")]
    pub password: Option<String>,
    #[sea_orm(column_type = "Text", default_value = "en")]
    pub locale: String,
    #[sea_orm(column_name = "parentSubscriptionId")]
    pub parent_subscription_id: Option<i32>,
    #[sea_orm(column_name = "collectionOrder")]
    pub collection_order: Vec<i32>,
    #[sea_orm(column_name = "linksRouteTo", default_value = "ORIGINAL")]
    pub links_route_to: LinksRouteTo,
    #[sea_orm(column_name = "aiTaggingMethod", default_value = "DISABLED")]
    pub ai_tagging_method: AiTaggingMethod,
    #[sea_orm(
        column_name = "aiPredefinedTags",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub ai_predefined_tags: Vec<String>,
    #[sea_orm(column_name = "aiTagExistingLinks", default_value = false)]
    pub ai_tag_existing_links: bool,
    #[sea_orm(default_value = "dark")]
    pub theme: Theme,
    #[sea_orm(column_name = "dismissedAnnouncementId", column_type = "Text")]
    pub dismissed_announcement_id: Option<String>,
    #[sea_orm(
        column_name = "readableFontFamily",
        column_type = "Text",
        default_value = "sans-serif"
    )]
    pub readable_font_family: Option<String>,
    #[sea_orm(
        column_name = "readableFontSize",
        column_type = "Text",
        default_value = "20px"
    )]
    pub readable_font_size: Option<String>,
    #[sea_orm(
        column_name = "readableLineHeight",
        column_type = "Text",
        default_value = "1.8"
    )]
    pub readable_line_height: Option<String>,
    #[sea_orm(
        column_name = "readableLineWidth",
        column_type = "Text",
        default_value = "normal"
    )]
    pub readable_line_width: Option<String>,
    #[sea_orm(column_name = "preventDuplicateLinks", default_value = false)]
    pub prevent_duplicate_links: bool,
    #[sea_orm(column_name = "archiveAsScreenshot", default_value = true)]
    pub archive_as_screenshot: bool,
    #[sea_orm(column_name = "archiveAsMonolith", default_value = true)]
    pub archive_as_monolith: bool,
    #[sea_orm(column_name = "archiveAsPDF", default_value = true)]
    pub archive_as_pdf: bool,
    #[sea_orm(column_name = "archiveAsReadable", default_value = true)]
    pub archive_as_readable: bool,
    #[sea_orm(column_name = "archiveAsWaybackMachine", default_value = false)]
    pub archive_as_wayback_machine: bool,
    #[sea_orm(column_name = "referredBy", column_type = "Text")]
    pub referred_by: Option<String>,
    #[sea_orm(column_name = "lastPickedAt", column_type = "custom(\"timestamp(3)\")")]
    pub last_picked_at: Option<DateTime>,
    #[sea_orm(column_name = "acceptPromotionalEmails", default_value = false)]
    pub accept_promotional_emails: bool,
    #[sea_orm(column_name = "shareUsageAnalytics", default_value = true)]
    pub share_usage_analytics: bool,
    #[sea_orm(column_name = "trialEndEmailSent", default_value = false)]
    pub trial_end_email_sent: bool,
    #[sea_orm(column_name = "sessionVersion", default_value = 0)]
    pub session_version: i32,
    #[sea_orm(column_name = "passwordRetries", default_value = 0)]
    pub password_retries: i32,
    #[sea_orm(
        column_name = "passwordRetryAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub password_retry_at: Option<DateTime>,
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
        belongs_to = "super::subscription::Entity",
        from = "Column::ParentSubscriptionId",
        to = "super::subscription::Column::Id",
        fk_name = "User_parentSubscriptionId_fkey"
    )]
    ParentSubscription,
    #[sea_orm(has_many = "super::account::Entity")]
    Accounts,
    #[sea_orm(
        has_many = "super::collection::Entity",
        from = "Column::Id",
        to = "super::collection::Column::OwnerId"
    )]
    Collections,
    #[sea_orm(has_many = "super::tag::Entity")]
    Tags,
    #[sea_orm(has_many = "super::link::Entity")]
    CreatedLinks,
    #[sea_orm(
        has_many = "super::collection::Entity",
        from = "Column::Id",
        to = "super::collection::Column::CreatedById"
    )]
    CreatedCollections,
    #[sea_orm(has_many = "super::highlight::Entity")]
    Highlights,
    #[sea_orm(has_many = "super::users_and_collections::Entity")]
    CollectionsJoined,
    #[sea_orm(has_many = "super::access_token::Entity")]
    AccessTokens,
    #[sea_orm(has_many = "super::dashboard_section::Entity")]
    DashboardSections,
    #[sea_orm(
        has_one = "super::subscription::Entity",
        from = "Column::Id",
        to = "super::subscription::Column::UserId"
    )]
    Subscriptions,
}

impl Related<super::subscription::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ParentSubscription.def()
    }
}

impl Related<super::account::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Accounts.def()
    }
}

impl Related<super::collection::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Collections.def()
    }
}

impl Related<super::tag::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Tags.def()
    }
}

impl Related<super::link::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CreatedLinks.def()
    }
}

impl Related<super::highlight::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Highlights.def()
    }
}

impl Related<super::users_and_collections::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CollectionsJoined.def()
    }
}

impl Related<super::access_token::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AccessTokens.def()
    }
}

impl Related<super::dashboard_section::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::DashboardSections.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            uuid: Set(uuid::Uuid::new_v4()),
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
