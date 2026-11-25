use super::cancellation_reason_requirement::CancellationReasonRequirement;
use super::period_type::PeriodType;
use super::scheduling_type::SchedulingType;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "EventType")]
pub struct Model {
    #[sea_orm(primary_key)]
    pub id: i32,
    /// @zod.string.min(1)
    #[sea_orm(column_type = "Text")]
    pub title: String,
    /// @zod.import(["import { eventTypeSlug } from '../../zod-utils'"]).custom.use(eventTypeSlug)
    #[sea_orm(column_type = "Text")]
    pub slug: String,
    #[sea_orm(column_type = "Text")]
    pub description: Option<String>,
    #[sea_orm(column_name = "interfaceLanguage", column_type = "Text")]
    pub interface_language: Option<String>,
    #[sea_orm(default_value = 0)]
    pub position: i32,
    /// @zod.import(["import { eventTypeLocations } from '../../zod-utils'"]).custom.use(eventTypeLocations)
    #[sea_orm(column_type = "JsonBinary")]
    pub locations: Option<Json>,
    /// @zod.number.min(1)
    pub length: i32,
    #[sea_orm(column_name = "offsetStart", default_value = 0)]
    pub offset_start: i32,
    #[sea_orm(default_value = false)]
    pub hidden: bool,
    #[sea_orm(column_name = "userId")]
    pub user_id: Option<i32>,
    #[sea_orm(column_name = "profileId")]
    pub profile_id: Option<i32>,
    #[sea_orm(column_name = "teamId")]
    pub team_id: Option<i32>,
    #[sea_orm(column_name = "useEventLevelSelectedCalendars", default_value = false)]
    pub use_event_level_selected_calendars: bool,
    #[sea_orm(column_name = "eventName", column_type = "Text")]
    pub event_name: Option<String>,
    #[sea_orm(column_name = "parentId")]
    pub parent_id: Option<i32>,
    /// @zod.import(["import { eventTypeBookingFields } from '../../zod-utils'"]).custom.use(eventTypeBookingFields)
    #[sea_orm(column_name = "bookingFields", column_type = "JsonBinary")]
    pub booking_fields: Option<Json>,
    #[sea_orm(column_name = "timeZone", column_type = "Text")]
    pub time_zone: Option<String>,
    #[sea_orm(column_name = "periodType", default_value = "unlimited")]
    pub period_type: PeriodType,
    /// @zod.import(["import { coerceToDate } from '../../zod-utils'"]).custom.use(coerceToDate)
    #[sea_orm(
        column_name = "periodStartDate",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub period_start_date: Option<DateTime>,
    /// @zod.import(["import { coerceToDate } from '../../zod-utils'"]).custom.use(coerceToDate)
    #[sea_orm(
        column_name = "periodEndDate",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub period_end_date: Option<DateTime>,
    #[sea_orm(column_name = "periodDays")]
    pub period_days: Option<i32>,
    #[sea_orm(column_name = "periodCountCalendarDays")]
    pub period_count_calendar_days: Option<bool>,
    #[sea_orm(column_name = "lockTimeZoneToggleOnBookingPage", default_value = false)]
    pub lock_time_zone_toggle_on_booking_page: bool,
    #[sea_orm(column_name = "lockedTimeZone", column_type = "Text")]
    pub locked_time_zone: Option<String>,
    #[sea_orm(column_name = "requiresConfirmation", default_value = false)]
    pub requires_confirmation: bool,
    #[sea_orm(
        column_name = "requiresConfirmationWillBlockSlot",
        default_value = false
    )]
    pub requires_confirmation_will_block_slot: bool,
    #[sea_orm(
        column_name = "requiresConfirmationForFreeEmail",
        default_value = false
    )]
    pub requires_confirmation_for_free_email: bool,
    #[sea_orm(column_name = "requiresBookerEmailVerification", default_value = false)]
    pub requires_booker_email_verification: bool,
    #[sea_orm(
        column_name = "canSendCalVideoTranscriptionEmails",
        default_value = true
    )]
    pub can_send_cal_video_transcription_emails: bool,
    #[sea_orm(column_name = "autoTranslateDescriptionEnabled", default_value = false)]
    pub auto_translate_description_enabled: bool,
    #[sea_orm(
        column_name = "autoTranslateInstantMeetingTitleEnabled",
        default_value = true
    )]
    pub auto_translate_instant_meeting_title_enabled: bool,
    /// @zod.import(["import { recurringEventType } from '../../zod-utils'"]).custom.use(recurringEventType)
    #[sea_orm(column_name = "recurringEvent", column_type = "JsonBinary")]
    pub recurring_event: Option<Json>,
    #[sea_orm(column_name = "disableGuests", default_value = false)]
    pub disable_guests: bool,
    #[sea_orm(column_name = "hideCalendarNotes", default_value = false)]
    pub hide_calendar_notes: bool,
    #[sea_orm(column_name = "hideCalendarEventDetails", default_value = false)]
    pub hide_calendar_event_details: bool,
    /// @zod.number.min(0)
    #[sea_orm(column_name = "minimumBookingNotice", default_value = 120)]
    pub minimum_booking_notice: i32,
    #[sea_orm(column_name = "beforeEventBuffer", default_value = 0)]
    pub before_event_buffer: i32,
    #[sea_orm(column_name = "afterEventBuffer", default_value = 0)]
    pub after_event_buffer: i32,
    #[sea_orm(column_name = "seatsPerTimeSlot")]
    pub seats_per_time_slot: Option<i32>,
    #[sea_orm(column_name = "onlyShowFirstAvailableSlot", default_value = false)]
    pub only_show_first_available_slot: bool,
    #[sea_orm(column_name = "showOptimizedSlots", default_value = false)]
    pub show_optimized_slots: Option<bool>,
    #[sea_orm(column_name = "disableCancelling", default_value = false)]
    pub disable_cancelling: Option<bool>,
    #[sea_orm(column_name = "disableRescheduling", default_value = false)]
    pub disable_rescheduling: Option<bool>,
    /// @zod.number.min(0)
    #[sea_orm(column_name = "minimumRescheduleNotice")]
    pub minimum_reschedule_notice: Option<i32>,
    #[sea_orm(column_name = "seatsShowAttendees", default_value = false)]
    pub seats_show_attendees: Option<bool>,
    #[sea_orm(column_name = "seatsShowAvailabilityCount", default_value = true)]
    pub seats_show_availability_count: Option<bool>,
    #[sea_orm(column_name = "schedulingType")]
    pub scheduling_type: Option<SchedulingType>,
    #[sea_orm(column_name = "scheduleId")]
    pub schedule_id: Option<i32>,
    #[sea_orm(
        column_name = "allowReschedulingCancelledBookings",
        default_value = false
    )]
    pub allow_rescheduling_cancelled_bookings: Option<bool>,
    #[sea_orm(default_value = 0)]
    pub price: i32,
    #[sea_orm(column_type = "Text", default_value = "usd")]
    pub currency: String,
    #[sea_orm(column_name = "slotInterval")]
    pub slot_interval: Option<i32>,
    /// @zod.import(["import { EventTypeMetaDataSchema } from '../../zod-utils'"]).custom.use(EventTypeMetaDataSchema)
    #[sea_orm(column_type = "JsonBinary")]
    pub metadata: Option<Json>,
    /// @zod.import(["import { successRedirectUrl } from '../../zod-utils'"]).custom.use(successRedirectUrl)
    #[sea_orm(column_name = "successRedirectUrl", column_type = "Text")]
    pub success_redirect_url: Option<String>,
    #[sea_orm(column_name = "forwardParamsSuccessRedirect", default_value = true)]
    pub forward_params_success_redirect: Option<bool>,
    /// @zod.import(["import { intervalLimitsType } from '../../zod-utils'"]).custom.use(intervalLimitsType)
    #[sea_orm(column_name = "bookingLimits", column_type = "JsonBinary")]
    pub booking_limits: Option<Json>,
    /// @zod.import(["import { intervalLimitsType } from '../../zod-utils'"]).custom.use(intervalLimitsType)
    #[sea_orm(column_name = "durationLimits", column_type = "JsonBinary")]
    pub duration_limits: Option<Json>,
    #[sea_orm(column_name = "isInstantEvent", default_value = false)]
    pub is_instant_event: bool,
    #[sea_orm(
        column_name = "instantMeetingExpiryTimeOffsetInSeconds",
        default_value = 90
    )]
    pub instant_meeting_expiry_time_offset_in_seconds: i32,
    #[sea_orm(column_name = "instantMeetingScheduleId")]
    pub instant_meeting_schedule_id: Option<i32>,
    #[sea_orm(
        column_name = "instantMeetingParameters",
        column_type = "Array(RcOrArc::new(ColumnType::Text))"
    )]
    pub instant_meeting_parameters: Vec<String>,
    #[sea_orm(column_name = "assignAllTeamMembers", default_value = false)]
    pub assign_all_team_members: bool,
    #[sea_orm(column_name = "assignRRMembersUsingSegment", default_value = false)]
    pub assign_rrmembers_using_segment: bool,
    /// @zod.import(["import { rrSegmentQueryValueSchema } from '../../zod-utils'"]).custom.use(rrSegmentQueryValueSchema)
    #[sea_orm(column_name = "rrSegmentQueryValue", column_type = "JsonBinary")]
    pub rr_segment_query_value: Option<Json>,
    #[sea_orm(
        column_name = "useEventTypeDestinationCalendarEmail",
        default_value = false
    )]
    pub use_event_type_destination_calendar_email: bool,
    #[sea_orm(column_name = "isRRWeightsEnabled", default_value = false)]
    pub is_rrweights_enabled: bool,
    #[sea_orm(column_name = "maxLeadThreshold")]
    pub max_lead_threshold: Option<i32>,
    #[sea_orm(column_name = "includeNoShowInRRCalculation", default_value = false)]
    pub include_no_show_in_rrcalculation: bool,
    #[sea_orm(column_name = "allowReschedulingPastBookings", default_value = false)]
    pub allow_rescheduling_past_bookings: bool,
    #[sea_orm(column_name = "hideOrganizerEmail", default_value = false)]
    pub hide_organizer_email: bool,
    #[sea_orm(column_name = "maxActiveBookingsPerBooker")]
    pub max_active_bookings_per_booker: Option<i32>,
    #[sea_orm(
        column_name = "maxActiveBookingPerBookerOfferReschedule",
        default_value = false
    )]
    pub max_active_booking_per_booker_offer_reschedule: bool,
    /// @zod.import(["import { emailSchema } from '../../zod-utils'"]).custom.use(emailSchema)
    #[sea_orm(column_name = "customReplyToEmail", column_type = "Text")]
    pub custom_reply_to_email: Option<String>,
    /// @zod.import(["import { eventTypeColor } from '../../zod-utils'"]).custom.use(eventTypeColor)
    #[sea_orm(column_name = "eventTypeColor", column_type = "JsonBinary")]
    pub event_type_color: Option<Json>,
    #[sea_orm(
        column_name = "rescheduleWithSameRoundRobinHost",
        default_value = false
    )]
    pub reschedule_with_same_round_robin_host: bool,
    #[sea_orm(column_name = "secondaryEmailId")]
    pub secondary_email_id: Option<i32>,
    #[sea_orm(column_name = "useBookerTimezone", default_value = false)]
    pub use_booker_timezone: bool,
    #[sea_orm(column_name = "restrictionScheduleId")]
    pub restriction_schedule_id: Option<i32>,
    #[sea_orm(column_name = "bookingRequiresAuthentication", default_value = false)]
    pub booking_requires_authentication: bool,
    #[sea_orm(column_name = "rrHostSubsetEnabled", default_value = false)]
    pub rr_host_subset_enabled: bool,
    #[sea_orm(
        column_name = "requiresCancellationReason",
        default_value = "MANDATORY_HOST_ONLY"
    )]
    pub requires_cancellation_reason: Option<CancellationReasonRequirement>,
    #[sea_orm(column_name = "enablePerHostLocations", default_value = false)]
    pub enable_per_host_locations: bool,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: Option<DateTime>,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: Option<DateTime>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_delete = "Cascade",
        fk_name = "EventType_userId_fkey"
    )]
    Owner,
    #[sea_orm(
        belongs_to = "super::profile::Entity",
        from = "Column::ProfileId",
        to = "super::profile::Column::Id",
        fk_name = "EventType_profileId_fkey"
    )]
    Profile,
    #[sea_orm(
        belongs_to = "super::team::Entity",
        from = "Column::TeamId",
        to = "super::team::Column::Id",
        on_delete = "Cascade",
        fk_name = "EventType_teamId_fkey"
    )]
    Team,
    #[sea_orm(
        belongs_to = "super::event_type::Entity",
        from = "Column::ParentId",
        to = "super::event_type::Column::Id",
        on_delete = "Cascade",
        fk_name = "EventType_parentId_fkey"
    )]
    Parent,
    #[sea_orm(
        belongs_to = "super::schedule::Entity",
        from = "Column::ScheduleId",
        to = "super::schedule::Column::Id",
        fk_name = "EventType_scheduleId_fkey"
    )]
    Schedule,
    #[sea_orm(
        belongs_to = "super::schedule::Entity",
        from = "Column::InstantMeetingScheduleId",
        to = "super::schedule::Column::Id",
        fk_name = "EventType_instantMeetingScheduleId_fkey"
    )]
    InstantMeetingSchedule,
    #[sea_orm(
        belongs_to = "super::secondary_email::Entity",
        from = "Column::SecondaryEmailId",
        to = "super::secondary_email::Column::Id",
        on_delete = "Cascade",
        fk_name = "EventType_secondaryEmailId_fkey"
    )]
    SecondaryEmail,
    #[sea_orm(
        belongs_to = "super::schedule::Entity",
        from = "Column::RestrictionScheduleId",
        to = "super::schedule::Column::Id",
        fk_name = "EventType_restrictionScheduleId_fkey"
    )]
    RestrictionSchedule,
    #[sea_orm(has_many = "super::host::Entity")]
    Hosts,
    #[sea_orm(has_many = "super::hashed_link::Entity")]
    HashedLink,
    #[sea_orm(has_many = "super::booking::Entity")]
    Bookings,
    #[sea_orm(has_many = "super::availability::Entity")]
    Availability,
    #[sea_orm(has_many = "super::webhook::Entity")]
    Webhooks,
    #[sea_orm(has_many = "super::event_type_custom_input::Entity")]
    CustomInputs,
    #[sea_orm(
        has_many = "super::event_type::Entity",
        from = "Column::Id",
        to = "super::event_type::Column::ParentId"
    )]
    Children,
    #[sea_orm(has_many = "super::event_type_translation::Entity")]
    FieldTranslations,
    #[sea_orm(has_many = "super::selected_calendar::Entity")]
    SelectedCalendars,
    #[sea_orm(has_many = "super::host_group::Entity")]
    HostGroups,
    #[sea_orm(has_one = "super::destination_calendar::Entity")]
    DestinationCalendar,
    #[sea_orm(has_one = "super::cal_video_settings::Entity")]
    CalVideoSettings,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Owner.def()
    }
}

impl Related<super::profile::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Profile.def()
    }
}

impl Related<super::team::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Team.def()
    }
}

impl Related<super::event_type::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Parent.def()
    }
}

impl Related<super::schedule::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Schedule.def()
    }
}

impl Related<super::secondary_email::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::SecondaryEmail.def()
    }
}

impl Related<super::host::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Hosts.def()
    }
}

impl Related<super::hashed_link::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::HashedLink.def()
    }
}

impl Related<super::booking::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Bookings.def()
    }
}

impl Related<super::availability::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Availability.def()
    }
}

impl Related<super::webhook::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Webhooks.def()
    }
}

impl Related<super::event_type_custom_input::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CustomInputs.def()
    }
}

impl Related<super::event_type_translation::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::FieldTranslations.def()
    }
}

impl Related<super::selected_calendar::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::SelectedCalendars.def()
    }
}

impl Related<super::host_group::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::HostGroups.def()
    }
}

impl Related<super::destination_calendar::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::DestinationCalendar.def()
    }
}

impl Related<super::cal_video_settings::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CalVideoSettings.def()
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
            self.created_at = Set(Some(now.naive_utc()));
        }
        if !self.updated_at.is_set() {
            self.updated_at = Set(Some(now.naive_utc()));
        }
        Ok(self)
    }
}
