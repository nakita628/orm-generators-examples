use super::booking_status::BookingStatus;
use super::creation_source::CreationSource;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Booking")]
pub struct Model {
    #[sea_orm(primary_key)]
    pub id: i32,
    #[sea_orm(unique)]
    #[sea_orm(column_type = "Text")]
    pub uid: String,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "idempotencyKey", column_type = "Text")]
    pub idempotency_key: Option<String>,
    #[sea_orm(column_name = "userId")]
    pub user_id: Option<i32>,
    /// @zod.import(["import { emailSchema } from '../../zod-utils'"]).custom.use(emailSchema)
    #[sea_orm(column_name = "userPrimaryEmail", column_type = "Text")]
    pub user_primary_email: Option<String>,
    #[sea_orm(column_name = "eventTypeId")]
    pub event_type_id: Option<i32>,
    #[sea_orm(column_type = "Text")]
    pub title: String,
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    #[sea_orm(column_name = "customInputs", column_type = "JsonBinary")]
    pub custom_inputs: Option<Json>,
    /// @zod.import(["import { bookingResponses } from '../../zod-utils'"]).custom.use(bookingResponses)
    #[sea_orm(column_type = "JsonBinary")]
    pub responses: Option<Json>,
    #[sea_orm(column_name = "startTime", column_type = "custom(\"timestamp(3)\")")]
    pub start_time: DateTime,
    #[sea_orm(column_name = "endTime", column_type = "custom(\"timestamp(3)\")")]
    pub end_time: DateTime,
    #[sea_orm(column_type = "Text")]
    pub location: Option<String>,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: Option<DateTime>,
    #[sea_orm(default_value = "accepted")]
    pub status: BookingStatus,
    #[sea_orm(default_value = false)]
    pub paid: bool,
    #[sea_orm(column_name = "destinationCalendarId")]
    pub destination_calendar_id: Option<i32>,
    #[sea_orm(column_name = "cancellationReason", column_type = "Text")]
    pub cancellation_reason: Option<String>,
    #[sea_orm(column_name = "rejectionReason", column_type = "Text")]
    pub rejection_reason: Option<String>,
    #[sea_orm(column_name = "reassignReason", column_type = "Text")]
    pub reassign_reason: Option<String>,
    #[sea_orm(column_name = "reassignById")]
    pub reassign_by_id: Option<i32>,
    #[sea_orm(column_name = "dynamicEventSlugRef", column_type = "Text")]
    pub dynamic_event_slug_ref: Option<String>,
    #[sea_orm(column_name = "dynamicGroupSlugRef", column_type = "Text")]
    pub dynamic_group_slug_ref: Option<String>,
    pub rescheduled: Option<bool>,
    #[sea_orm(column_name = "fromReschedule", column_type = "Text")]
    pub from_reschedule: Option<String>,
    #[sea_orm(column_name = "recurringEventId", column_type = "Text")]
    pub recurring_event_id: Option<String>,
    #[sea_orm(column_name = "smsReminderNumber", column_type = "Text")]
    pub sms_reminder_number: Option<String>,
    #[sea_orm(
        column_name = "scheduledJobs",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub scheduled_jobs: Vec<String>,
    /// @zod.import(["import { bookingMetadataSchema } from '../../zod-utils'"]).custom.use(bookingMetadataSchema)
    #[sea_orm(column_type = "JsonBinary")]
    pub metadata: Option<Json>,
    #[sea_orm(column_name = "isRecorded", default_value = false)]
    pub is_recorded: bool,
    #[sea_orm(column_name = "iCalUID", column_type = "Text", default_value = "")]
    pub i_cal_uid: Option<String>,
    #[sea_orm(column_name = "iCalSequence", default_value = 0)]
    pub i_cal_sequence: i32,
    pub rating: Option<i32>,
    #[sea_orm(column_name = "ratingFeedback", column_type = "Text")]
    pub rating_feedback: Option<String>,
    #[sea_orm(column_name = "noShowHost", default_value = false)]
    pub no_show_host: Option<bool>,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "oneTimePassword", column_type = "Text")]
    pub one_time_password: Option<String>,
    /// @zod.import(["import { emailSchema } from '../../zod-utils'"]).custom.use(emailSchema)
    #[sea_orm(column_name = "cancelledBy", column_type = "Text")]
    pub cancelled_by: Option<String>,
    /// @zod.import(["import { emailSchema } from '../../zod-utils'"]).custom.use(emailSchema)
    #[sea_orm(column_name = "rescheduledBy", column_type = "Text")]
    pub rescheduled_by: Option<String>,
    #[sea_orm(column_name = "creationSource")]
    pub creation_source: Option<CreationSource>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "Booking_userId_fkey"
    )]
    User,
    #[sea_orm(
        belongs_to = "super::event_type::Entity",
        from = "Column::EventTypeId",
        to = "super::event_type::Column::Id",
        fk_name = "Booking_eventTypeId_fkey"
    )]
    EventType,
    #[sea_orm(
        belongs_to = "super::destination_calendar::Entity",
        from = "Column::DestinationCalendarId",
        to = "super::destination_calendar::Column::Id",
        fk_name = "Booking_destinationCalendarId_fkey"
    )]
    DestinationCalendar,
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::ReassignById",
        to = "super::user::Column::Id",
        fk_name = "Booking_reassignById_fkey"
    )]
    ReassignBy,
    #[sea_orm(has_many = "super::booking_reference::Entity")]
    References,
    #[sea_orm(has_many = "super::attendee::Entity")]
    Attendees,
    #[sea_orm(has_many = "super::payment::Entity")]
    Payment,
    #[sea_orm(has_many = "super::booking_seat::Entity")]
    SeatsReferences,
    #[sea_orm(has_many = "super::webhook_scheduled_triggers::Entity")]
    ScheduledTriggers,
    #[sea_orm(has_many = "super::assignment_reason::Entity")]
    AssignmentReason,
    #[sea_orm(has_many = "super::booking_internal_note::Entity")]
    InternalNote,
    #[sea_orm(has_many = "super::credit_expense_log::Entity")]
    ExpenseLogs,
    #[sea_orm(has_many = "super::wrong_assignment_report::Entity")]
    WrongAssignmentReports,
    #[sea_orm(has_one = "super::instant_meeting_token::Entity")]
    InstantMeetingToken,
    #[sea_orm(has_one = "super::tracking::Entity")]
    Tracking,
    #[sea_orm(has_one = "super::booking_report::Entity")]
    Report,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::User.def()
    }
}

impl Related<super::event_type::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::EventType.def()
    }
}

impl Related<super::destination_calendar::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::DestinationCalendar.def()
    }
}

impl Related<super::booking_reference::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::References.def()
    }
}

impl Related<super::attendee::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Attendees.def()
    }
}

impl Related<super::payment::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Payment.def()
    }
}

impl Related<super::booking_seat::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::SeatsReferences.def()
    }
}

impl Related<super::webhook_scheduled_triggers::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ScheduledTriggers.def()
    }
}

impl Related<super::assignment_reason::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AssignmentReason.def()
    }
}

impl Related<super::booking_internal_note::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::InternalNote.def()
    }
}

impl Related<super::credit_expense_log::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ExpenseLogs.def()
    }
}

impl Related<super::wrong_assignment_report::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::WrongAssignmentReports.def()
    }
}

impl Related<super::instant_meeting_token::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::InstantMeetingToken.def()
    }
}

impl Related<super::tracking::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Tracking.def()
    }
}

impl Related<super::booking_report::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Report.def()
    }
}

#[async_trait::async_trait]
impl ActiveModelBehavior for ActiveModel {
    fn new() -> Self {
        Self {
            one_time_password: Set(Some(uuid::Uuid::new_v4().to_string())),
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
            self.updated_at = Set(Some(now.naive_utc()));
        }
        Ok(self)
    }
}
