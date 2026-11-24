use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "BookingSettings")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(column_name = "bufferStartTime", default_value = 0)]
    pub buffer_start_time: i32,
    #[sea_orm(column_name = "tagsRequired", default_value = false)]
    pub tags_required: bool,
    #[sea_orm(column_name = "maxBookingLength")]
    pub max_booking_length: Option<i32>,
    #[sea_orm(column_name = "maxBookingLengthSkipClosedDays", default_value = false)]
    pub max_booking_length_skip_closed_days: bool,
    #[sea_orm(column_name = "autoArchiveBookings", default_value = false)]
    pub auto_archive_bookings: bool,
    #[sea_orm(column_name = "autoArchiveDays", default_value = 2)]
    pub auto_archive_days: i32,
    #[sea_orm(column_name = "autoArchiveExpiredReservations", default_value = false)]
    pub auto_archive_expired_reservations: bool,
    #[sea_orm(column_name = "requireApprovalForBase", default_value = false)]
    pub require_approval_for_base: bool,
    #[sea_orm(column_name = "requireApprovalForSelfService", default_value = false)]
    pub require_approval_for_self_service: bool,
    #[sea_orm(column_name = "requireExplicitCheckinForAdmin", default_value = false)]
    pub require_explicit_checkin_for_admin: bool,
    #[sea_orm(
        column_name = "requireExplicitCheckinForSelfService",
        default_value = false
    )]
    pub require_explicit_checkin_for_self_service: bool,
    #[sea_orm(column_name = "requireExplicitCheckoutForAdmin", default_value = false)]
    pub require_explicit_checkout_for_admin: bool,
    #[sea_orm(
        column_name = "requireExplicitCheckoutForSelfService",
        default_value = false
    )]
    pub require_explicit_checkout_for_self_service: bool,
    #[sea_orm(column_name = "countKitsAsSingleUnit", default_value = false)]
    pub count_kits_as_single_unit: bool,
    #[sea_orm(column_name = "notifyBookingCreator", default_value = true)]
    pub notify_booking_creator: bool,
    #[sea_orm(column_name = "notifyAdminsOnNewBooking", default_value = true)]
    pub notify_admins_on_new_booking: bool,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "organizationId")]
    pub organization_id: String,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::OrganizationId",
        to = "super::organization::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade"
    )]
    Organization,
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Organization.def()
    }
}

impl Related<super::team_member::Entity> for Entity {
    fn to() -> RelationDef {
        super::booking_settings_always_notify::Relation::TeamMember.def()
    }
    fn via() -> Option<RelationDef> {
        Some(
            super::booking_settings_always_notify::Relation::BookingSettings
                .def()
                .rev(),
        )
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
