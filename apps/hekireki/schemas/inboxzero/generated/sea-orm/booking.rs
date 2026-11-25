use super::booking_status::BookingStatus;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Booking")]
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
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "guestName", column_type = "Text")]
    pub guest_name: String,
    #[sea_orm(column_name = "guestEmail", column_type = "Text")]
    pub guest_email: String,
    #[sea_orm(column_name = "guestNote", column_type = "Text")]
    pub guest_note: Option<String>,
    #[sea_orm(column_name = "startTime", column_type = "custom(\"timestamp(3)\")")]
    pub start_time: DateTime,
    #[sea_orm(column_name = "endTime", column_type = "custom(\"timestamp(3)\")")]
    pub end_time: DateTime,
    #[sea_orm(default_value = "PENDING_PROVIDER_EVENT")]
    pub status: BookingStatus,
    #[sea_orm(column_type = "Text")]
    pub provider: Option<String>,
    #[sea_orm(column_name = "providerConnectionId", column_type = "Text")]
    pub provider_connection_id: Option<String>,
    #[sea_orm(column_name = "providerCalendarId", column_type = "Text")]
    pub provider_calendar_id: Option<String>,
    #[sea_orm(column_name = "providerEventId", column_type = "Text")]
    pub provider_event_id: Option<String>,
    #[sea_orm(column_name = "videoConferenceLink", column_type = "Text")]
    pub video_conference_link: Option<String>,
    #[sea_orm(column_name = "cancelTokenHash", column_type = "Text")]
    pub cancel_token_hash: String,
    #[sea_orm(column_name = "cancellationReason", column_type = "Text")]
    pub cancellation_reason: Option<String>,
    #[sea_orm(column_name = "idempotencyToken", column_type = "Text")]
    pub idempotency_token: Option<String>,
    #[sea_orm(column_name = "bookingLinkId", column_type = "Text")]
    pub booking_link_id: String,
    #[sea_orm(column_name = "emailAccountId", column_type = "Text")]
    pub email_account_id: String,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::booking_link::Entity",
        from = "Column::BookingLinkId",
        to = "super::booking_link::Column::Id",
        on_delete = "Cascade",
        fk_name = "Booking_bookingLinkId_fkey"
    )]
    BookingLink,
    #[sea_orm(
        belongs_to = "super::email_account::Entity",
        from = "Column::EmailAccountId",
        to = "super::email_account::Column::Id",
        on_delete = "Cascade",
        fk_name = "Booking_emailAccountId_fkey"
    )]
    EmailAccount,
}

impl Related<super::booking_link::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::BookingLink.def()
    }
}

impl Related<super::email_account::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::EmailAccount.def()
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
