use super::date_format_preference::DateFormatPreference;
use super::tier_id::TierId;
use super::time_format_preference::TimeFormatPreference;
use super::week_start_preference::WeekStartPreference;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "User")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    pub id: String,
    #[sea_orm(unique)]
    pub email: String,
    #[sea_orm(unique)]
    pub username: String,
    #[sea_orm(column_name = "firstName")]
    pub first_name: Option<String>,
    #[sea_orm(column_name = "lastName")]
    pub last_name: Option<String>,
    #[sea_orm(column_name = "displayName")]
    pub display_name: Option<String>,
    #[sea_orm(column_name = "profilePicture")]
    pub profile_picture: Option<String>,
    #[sea_orm(column_name = "usedFreeTrial", default_value = false)]
    pub used_free_trial: bool,
    #[sea_orm(default_value = false)]
    pub onboarded: bool,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "customerId")]
    pub customer_id: Option<String>,
    #[sea_orm(default_value = false)]
    pub sso: bool,
    #[sea_orm(column_name = "createdWithInvite", default_value = false)]
    pub created_with_invite: bool,
    #[sea_orm(column_name = "skipSubscriptionCheck", default_value = false)]
    pub skip_subscription_check: bool,
    #[sea_orm(column_name = "hasUnpaidInvoice", default_value = false)]
    pub has_unpaid_invoice: bool,
    #[sea_orm(column_name = "warnForNoPaymentMethod", default_value = false)]
    pub warn_for_no_payment_method: bool,
    #[sea_orm(column_name = "lastSelectedOrganizationId")]
    pub last_selected_organization_id: Option<String>,
    #[sea_orm(column_name = "createdAt")]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "deletedAt")]
    pub deleted_at: Option<DateTime>,
    #[sea_orm(column_name = "lastMobileActiveAt")]
    pub last_mobile_active_at: Option<DateTime>,
    #[sea_orm(column_name = "dateFormat")]
    pub date_format: Option<DateFormatPreference>,
    #[sea_orm(column_name = "timeFormat")]
    pub time_format: Option<TimeFormatPreference>,
    #[sea_orm(column_name = "weekStart")]
    pub week_start: Option<WeekStartPreference>,
    #[sea_orm(column_name = "timeZone")]
    pub time_zone: Option<String>,
    #[sea_orm(column_name = "tierId", default_value = "free")]
    pub tier_id: TierId,
    #[sea_orm(column_name = "referralSource")]
    pub referral_source: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::organization::Entity",
        from = "Column::LastSelectedOrganizationId",
        to = "super::organization::Column::Id",
        on_delete = "SetNull"
    )]
    LastSelectedOrganization,
    #[sea_orm(
        belongs_to = "super::tier::Entity",
        from = "Column::TierId",
        to = "super::tier::Column::Id"
    )]
    Tier,
    #[sea_orm(has_many = "super::asset::Entity")]
    Assets,
    #[sea_orm(has_many = "super::category::Entity")]
    Categories,
    #[sea_orm(has_many = "super::note::Entity")]
    Notes,
    #[sea_orm(has_many = "super::booking_note::Entity")]
    BookingNotes,
    #[sea_orm(has_many = "super::audit_note::Entity")]
    AuditNotes,
    #[sea_orm(has_many = "super::location_note::Entity")]
    LocationNotes,
    #[sea_orm(has_many = "super::qr::Entity")]
    QrCodes,
    #[sea_orm(has_many = "super::scan::Entity")]
    Scans,
    #[sea_orm(has_many = "super::tag::Entity")]
    Tags,
    #[sea_orm(has_many = "super::location::Entity")]
    Locations,
    #[sea_orm(has_many = "super::image::Entity")]
    Images,
    #[sea_orm(
        has_many = "super::organization::Entity",
        from = "Column::Id",
        to = "super::organization::Column::UserId"
    )]
    Organizations,
    #[sea_orm(has_many = "super::custom_field::Entity")]
    CustomFields,
    #[sea_orm(
        has_many = "super::invite::Entity",
        from = "Column::Id",
        to = "super::invite::Column::InviterId"
    )]
    SentInvites,
    #[sea_orm(
        has_many = "super::invite::Entity",
        from = "Column::Id",
        to = "super::invite::Column::InviteeUserId"
    )]
    ReceivedInvites,
    #[sea_orm(has_many = "super::team_member::Entity")]
    TeamMembers,
    #[sea_orm(has_many = "super::user_organization::Entity")]
    UserOrganizations,
    #[sea_orm(
        has_many = "super::booking::Entity",
        from = "Column::Id",
        to = "super::booking::Column::CreatorId"
    )]
    Bookings,
    #[sea_orm(
        has_many = "super::booking::Entity",
        from = "Column::Id",
        to = "super::booking::Column::CustodianUserId"
    )]
    Custodies,
    #[sea_orm(
        has_many = "super::booking::Entity",
        from = "Column::Id",
        to = "super::booking::Column::ApprovedById"
    )]
    ApprovedBookings,
    #[sea_orm(has_many = "super::kit::Entity")]
    CreatedKits,
    #[sea_orm(has_many = "super::asset_reminder::Entity")]
    AssetReminders,
    #[sea_orm(has_many = "super::audit_session::Entity")]
    AuditsCreated,
    #[sea_orm(has_many = "super::asset_import::Entity")]
    AssetImports,
    #[sea_orm(has_many = "super::audit_assignment::Entity")]
    AuditAssignments,
    #[sea_orm(has_many = "super::audit_asset::Entity")]
    AuditAssetsScanned,
    #[sea_orm(has_many = "super::audit_scan::Entity")]
    AuditScans,
    #[sea_orm(has_many = "super::audit_image::Entity")]
    AuditImagesUploaded,
    #[sea_orm(has_many = "super::asset_index_settings::Entity")]
    AssetIndexSettings,
    #[sea_orm(has_many = "super::partial_booking_checkin::Entity")]
    PartialCheckins,
    #[sea_orm(has_many = "super::partial_booking_checkout::Entity")]
    PartialCheckouts,
    #[sea_orm(has_many = "super::mobile_auth_code::Entity")]
    MobileAuthCodes,
    #[sea_orm(has_many = "super::push_token::Entity")]
    PushTokens,
    #[sea_orm(has_many = "super::update::Entity")]
    CreatedUpdates,
    #[sea_orm(has_many = "super::user_update_read::Entity")]
    UpdateReads,
    #[sea_orm(has_many = "super::asset_filter_preset::Entity")]
    AssetFilterPresets,
    #[sea_orm(has_many = "super::asset_model::Entity")]
    AssetModels,
    #[sea_orm(has_many = "super::consumption_log::Entity")]
    ConsumptionLogs,
    #[sea_orm(has_many = "super::user_scim_external_id::Entity")]
    ScimExternalIds,
    #[sea_orm(
        has_many = "super::role_change_log::Entity",
        from = "Column::Id",
        to = "super::role_change_log::Column::UserId"
    )]
    RoleChangesReceived,
    #[sea_orm(
        has_many = "super::role_change_log::Entity",
        from = "Column::Id",
        to = "super::role_change_log::Column::ChangedById"
    )]
    RoleChangesPerformed,
    #[sea_orm(has_many = "super::team_member_note::Entity")]
    TeamMemberNotesAuthored,
    #[sea_orm(has_one = "super::user_contact::Entity")]
    Contact,
    #[sea_orm(has_one = "super::user_business_intel::Entity")]
    BusinessIntel,
    #[sea_orm(has_one = "super::notification_preference::Entity")]
    NotificationPreference,
    #[sea_orm(has_one = "super::custom_tier_limit::Entity")]
    CustomTierLimit,
}

impl Related<super::organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::LastSelectedOrganization.def()
    }
}

impl Related<super::tier::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Tier.def()
    }
}

impl Related<super::asset::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Assets.def()
    }
}

impl Related<super::category::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Categories.def()
    }
}

impl Related<super::note::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Notes.def()
    }
}

impl Related<super::booking_note::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::BookingNotes.def()
    }
}

impl Related<super::audit_note::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AuditNotes.def()
    }
}

impl Related<super::location_note::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::LocationNotes.def()
    }
}

impl Related<super::qr::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::QrCodes.def()
    }
}

impl Related<super::scan::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Scans.def()
    }
}

impl Related<super::tag::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Tags.def()
    }
}

impl Related<super::location::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Locations.def()
    }
}

impl Related<super::image::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Images.def()
    }
}

impl Related<super::custom_field::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CustomFields.def()
    }
}

impl Related<super::invite::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::SentInvites.def()
    }
}

impl Related<super::team_member::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::TeamMembers.def()
    }
}

impl Related<super::user_organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::UserOrganizations.def()
    }
}

impl Related<super::booking::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Bookings.def()
    }
}

impl Related<super::kit::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CreatedKits.def()
    }
}

impl Related<super::asset_reminder::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AssetReminders.def()
    }
}

impl Related<super::audit_session::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AuditsCreated.def()
    }
}

impl Related<super::asset_import::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AssetImports.def()
    }
}

impl Related<super::audit_assignment::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AuditAssignments.def()
    }
}

impl Related<super::audit_asset::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AuditAssetsScanned.def()
    }
}

impl Related<super::audit_scan::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AuditScans.def()
    }
}

impl Related<super::audit_image::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AuditImagesUploaded.def()
    }
}

impl Related<super::asset_index_settings::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AssetIndexSettings.def()
    }
}

impl Related<super::partial_booking_checkin::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::PartialCheckins.def()
    }
}

impl Related<super::partial_booking_checkout::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::PartialCheckouts.def()
    }
}

impl Related<super::mobile_auth_code::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::MobileAuthCodes.def()
    }
}

impl Related<super::push_token::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::PushTokens.def()
    }
}

impl Related<super::update::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CreatedUpdates.def()
    }
}

impl Related<super::user_update_read::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::UpdateReads.def()
    }
}

impl Related<super::asset_filter_preset::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AssetFilterPresets.def()
    }
}

impl Related<super::asset_model::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AssetModels.def()
    }
}

impl Related<super::consumption_log::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ConsumptionLogs.def()
    }
}

impl Related<super::user_scim_external_id::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ScimExternalIds.def()
    }
}

impl Related<super::role_change_log::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::RoleChangesReceived.def()
    }
}

impl Related<super::team_member_note::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::TeamMemberNotesAuthored.def()
    }
}

impl Related<super::user_contact::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Contact.def()
    }
}

impl Related<super::user_business_intel::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::BusinessIntel.def()
    }
}

impl Related<super::notification_preference::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::NotificationPreference.def()
    }
}

impl Related<super::custom_tier_limit::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CustomTierLimit.def()
    }
}

impl Related<super::role::Entity> for Entity {
    fn to() -> RelationDef {
        super::role_to_user::Relation::Role.def()
    }
    fn via() -> Option<RelationDef> {
        Some(super::role_to_user::Relation::User.def().rev())
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
