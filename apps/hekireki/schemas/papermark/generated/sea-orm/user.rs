use super::referral_track::ReferralTrack;
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
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "contactId", column_type = "Text")]
    pub contact_id: Option<String>,
    #[sea_orm(column_type = "Text", default_value = "free")]
    pub plan: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "stripeId", column_type = "Text")]
    pub stripe_id: Option<String>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "subscriptionId", column_type = "Text")]
    pub subscription_id: Option<String>,
    #[sea_orm(column_name = "startsAt", column_type = "custom(\"timestamp(3)\")")]
    pub starts_at: Option<DateTime>,
    #[sea_orm(column_name = "endsAt", column_type = "custom(\"timestamp(3)\")")]
    pub ends_at: Option<DateTime>,
    #[sea_orm(column_name = "referralTrack")]
    pub referral_track: Option<ReferralTrack>,
    #[sea_orm(
        column_name = "referralTrackSelectedAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub referral_track_selected_at: Option<DateTime>,
    #[sea_orm(column_name = "referralProfile", column_type = "JsonBinary")]
    pub referral_profile: Option<Json>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(has_many = "super::account::Entity")]
    Accounts,
    #[sea_orm(has_many = "super::session::Entity")]
    Sessions,
    #[sea_orm(has_many = "super::document::Entity")]
    Documents,
    #[sea_orm(has_many = "super::user_team::Entity")]
    Teams,
    #[sea_orm(has_many = "super::user_dataroom::Entity")]
    UserDatarooms,
    #[sea_orm(has_many = "super::domain::Entity")]
    Domains,
    #[sea_orm(has_many = "super::chat::Entity")]
    Chats,
    #[sea_orm(has_many = "super::restricted_token::Entity")]
    RestrictedTokens,
    #[sea_orm(has_many = "super::conversation_participant::Entity")]
    ParticipatedConversations,
    #[sea_orm(has_many = "super::message::Entity")]
    Messages,
    #[sea_orm(has_many = "super::dataroom_question::Entity")]
    CreatedDataroomQuestions,
    #[sea_orm(has_many = "super::dataroom_faq_item::Entity")]
    PublishedFaqItems,
    #[sea_orm(has_many = "super::document_annotation::Entity")]
    CreatedAnnotations,
    #[sea_orm(has_many = "super::document_redaction_job::Entity")]
    CreatedRedactionJobs,
    #[sea_orm(has_many = "super::installed_integration::Entity")]
    InstalledIntegrations,
    #[sea_orm(has_many = "super::link::Entity")]
    Links,
    #[sea_orm(has_many = "super::dataroom::Entity")]
    FrozenDatarooms,
    #[sea_orm(has_many = "super::task::Entity")]
    CreatedTasks,
    #[sea_orm(has_many = "super::task_activity::Entity")]
    TaskActivities,
    #[sea_orm(has_many = "super::notification_preference::Entity")]
    NotificationPreferences,
    #[sea_orm(has_many = "super::referral_invite::Entity")]
    ReferralInvites,
}

impl Related<super::account::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Accounts.def()
    }
}

impl Related<super::session::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Sessions.def()
    }
}

impl Related<super::document::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Documents.def()
    }
}

impl Related<super::user_team::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Teams.def()
    }
}

impl Related<super::user_dataroom::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::UserDatarooms.def()
    }
}

impl Related<super::domain::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Domains.def()
    }
}

impl Related<super::chat::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Chats.def()
    }
}

impl Related<super::restricted_token::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::RestrictedTokens.def()
    }
}

impl Related<super::conversation_participant::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ParticipatedConversations.def()
    }
}

impl Related<super::message::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Messages.def()
    }
}

impl Related<super::dataroom_question::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CreatedDataroomQuestions.def()
    }
}

impl Related<super::dataroom_faq_item::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::PublishedFaqItems.def()
    }
}

impl Related<super::document_annotation::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CreatedAnnotations.def()
    }
}

impl Related<super::document_redaction_job::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CreatedRedactionJobs.def()
    }
}

impl Related<super::installed_integration::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::InstalledIntegrations.def()
    }
}

impl Related<super::link::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Links.def()
    }
}

impl Related<super::dataroom::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::FrozenDatarooms.def()
    }
}

impl Related<super::task::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CreatedTasks.def()
    }
}

impl Related<super::task_activity::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::TaskActivities.def()
    }
}

impl Related<super::notification_preference::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::NotificationPreferences.def()
    }
}

impl Related<super::referral_invite::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ReferralInvites.def()
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
