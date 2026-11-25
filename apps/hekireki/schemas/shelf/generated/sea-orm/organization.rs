use super::currency::Currency;
use super::organization_type::OrganizationType;
use super::qr_id_display_preference::QrIdDisplayPreference;
use chrono::SubsecRound;
use sea_orm::entity::prelude::*;
use sea_orm::Set;
use serde::{Deserialize, Serialize};

#[derive(Clone, Debug, PartialEq, Eq, DeriveEntityModel, Serialize, Deserialize)]
#[sea_orm(table_name = "Organization")]
pub struct Model {
    #[sea_orm(primary_key, auto_increment = false)]
    #[sea_orm(column_type = "Text")]
    pub id: String,
    #[sea_orm(column_type = "Text", default_value = "Personal")]
    pub name: String,
    #[sea_orm(default_value = "PERSONAL")]
    pub r#type: OrganizationType,
    #[sea_orm(column_name = "userId", column_type = "Text")]
    pub user_id: String,
    #[sea_orm(default_value = "USD")]
    pub currency: Currency,
    #[sea_orm(unique)]
    #[sea_orm(column_name = "imageId", column_type = "Text")]
    pub image_id: Option<String>,
    #[sea_orm(column_name = "enabledSso", default_value = false)]
    pub enabled_sso: bool,
    #[sea_orm(column_name = "ssoDetailsId", column_type = "Text")]
    pub sso_details_id: Option<String>,
    #[sea_orm(column_name = "selfServiceCanSeeCustody", default_value = false)]
    pub self_service_can_see_custody: bool,
    #[sea_orm(column_name = "selfServiceCanSeeBookings", default_value = false)]
    pub self_service_can_see_bookings: bool,
    #[sea_orm(column_name = "baseUserCanSeeCustody", default_value = false)]
    pub base_user_can_see_custody: bool,
    #[sea_orm(column_name = "baseUserCanSeeBookings", default_value = false)]
    pub base_user_can_see_bookings: bool,
    #[sea_orm(column_name = "managersCanEditAssets", default_value = false)]
    pub managers_can_edit_assets: bool,
    #[sea_orm(column_name = "barcodesEnabled", default_value = false)]
    pub barcodes_enabled: bool,
    #[sea_orm(
        column_name = "barcodesEnabledAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub barcodes_enabled_at: Option<DateTime>,
    #[sea_orm(column_name = "usedBarcodeTrial", default_value = false)]
    pub used_barcode_trial: bool,
    #[sea_orm(column_name = "auditsEnabled", default_value = false)]
    pub audits_enabled: bool,
    #[sea_orm(
        column_name = "auditsEnabledAt",
        column_type = "custom(\"timestamp(3)\")"
    )]
    pub audits_enabled_at: Option<DateTime>,
    #[sea_orm(column_name = "usedAuditTrial", default_value = false)]
    pub used_audit_trial: bool,
    #[sea_orm(column_name = "workspaceDisabled", default_value = false)]
    pub workspace_disabled: bool,
    #[sea_orm(
        column_name = "createdAt",
        column_type = "custom(\"timestamp(3)\")",
        default_expr = "Expr::current_timestamp()"
    )]
    pub created_at: DateTime,
    #[sea_orm(column_name = "updatedAt", column_type = "custom(\"timestamp(3)\")")]
    pub updated_at: DateTime,
    #[sea_orm(column_name = "hasSequentialIdsMigrated", default_value = false)]
    pub has_sequential_ids_migrated: bool,
    #[sea_orm(column_name = "qrIdDisplayPreference", default_value = "QR_ID")]
    pub qr_id_display_preference: QrIdDisplayPreference,
    #[sea_orm(column_name = "showShelfBranding", default_value = true)]
    pub show_shelf_branding: bool,
    #[sea_orm(column_name = "showQrCodesOnPdfs", default_value = true)]
    pub show_qr_codes_on_pdfs: bool,
    #[sea_orm(
        column_name = "customEmailFooter",
        column_type = "String(StringLen::N(500))"
    )]
    pub custom_email_footer: Option<String>,
}

#[derive(Copy, Clone, Debug, EnumIter, DeriveRelation)]
pub enum Relation {
    #[sea_orm(
        belongs_to = "super::user::Entity",
        from = "Column::UserId",
        to = "super::user::Column::Id",
        on_update = "Cascade",
        on_delete = "Cascade",
        fk_name = "Organization_userId_fkey"
    )]
    Owner,
    #[sea_orm(
        belongs_to = "super::image::Entity",
        from = "Column::ImageId",
        to = "super::image::Column::Id",
        fk_name = "Organization_imageId_fkey"
    )]
    Image,
    #[sea_orm(
        belongs_to = "super::sso_details::Entity",
        from = "Column::SsoDetailsId",
        to = "super::sso_details::Column::Id",
        fk_name = "Organization_ssoDetailsId_fkey"
    )]
    SsoDetails,
    #[sea_orm(has_many = "super::team_member::Entity")]
    Members,
    #[sea_orm(has_many = "super::asset::Entity")]
    Assets,
    #[sea_orm(has_many = "super::asset_kit::Entity")]
    AssetKits,
    #[sea_orm(has_many = "super::asset_location::Entity")]
    AssetLocations,
    #[sea_orm(has_many = "super::scan::Entity")]
    Scans,
    #[sea_orm(has_many = "super::location::Entity")]
    Locations,
    #[sea_orm(has_many = "super::category::Entity")]
    Categories,
    #[sea_orm(has_many = "super::tag::Entity")]
    Tags,
    #[sea_orm(has_many = "super::qr::Entity")]
    QrCodes,
    #[sea_orm(has_many = "super::barcode::Entity")]
    Barcodes,
    #[sea_orm(has_many = "super::invite::Entity")]
    Invites,
    #[sea_orm(has_many = "super::user_organization::Entity")]
    UserOrganizations,
    #[sea_orm(has_many = "super::custom_field::Entity")]
    CustomFields,
    #[sea_orm(
        has_many = "super::image::Entity",
        from = "Column::Id",
        to = "super::image::Column::OwnerOrgId"
    )]
    Images,
    #[sea_orm(has_many = "super::booking::Entity")]
    Bookings,
    #[sea_orm(has_many = "super::kit::Entity")]
    Kits,
    #[sea_orm(has_many = "super::asset_model::Entity")]
    AssetModels,
    #[sea_orm(has_many = "super::asset_index_settings::Entity")]
    AssetIndexSettings,
    #[sea_orm(has_many = "super::asset_reminder::Entity")]
    AssetReminders,
    #[sea_orm(has_many = "super::asset_filter_preset::Entity")]
    AssetFilterPresets,
    #[sea_orm(has_many = "super::audit_session::Entity")]
    AuditSessions,
    #[sea_orm(has_many = "super::audit_image::Entity")]
    AuditImages,
    #[sea_orm(has_many = "super::asset_import::Entity")]
    AssetImports,
    #[sea_orm(has_many = "super::link_preview::Entity")]
    LinkPreviews,
    #[sea_orm(has_many = "super::pdf_template::Entity")]
    PdfTemplates,
    #[sea_orm(
        has_many = "super::user::Entity",
        from = "Column::Id",
        to = "super::user::Column::LastSelectedOrganizationId"
    )]
    UsersWithLastSelected,
    #[sea_orm(has_many = "super::scim_token::Entity")]
    ScimTokens,
    #[sea_orm(has_many = "super::user_scim_external_id::Entity")]
    ScimExternalIds,
    #[sea_orm(has_many = "super::role_change_log::Entity")]
    RoleChangeLogs,
    #[sea_orm(has_many = "super::team_member_note::Entity")]
    TeamMemberNotes,
    #[sea_orm(has_many = "super::activity_event::Entity")]
    ActivityEvents,
    #[sea_orm(has_one = "super::working_hours::Entity")]
    WorkingHours,
    #[sea_orm(has_one = "super::booking_settings::Entity")]
    BookingSettings,
}

impl Related<super::user::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Owner.def()
    }
}

impl Related<super::image::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Image.def()
    }
}

impl Related<super::sso_details::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::SsoDetails.def()
    }
}

impl Related<super::team_member::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Members.def()
    }
}

impl Related<super::asset::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Assets.def()
    }
}

impl Related<super::asset_kit::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AssetKits.def()
    }
}

impl Related<super::asset_location::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AssetLocations.def()
    }
}

impl Related<super::scan::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Scans.def()
    }
}

impl Related<super::location::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Locations.def()
    }
}

impl Related<super::category::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Categories.def()
    }
}

impl Related<super::tag::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Tags.def()
    }
}

impl Related<super::qr::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::QrCodes.def()
    }
}

impl Related<super::barcode::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Barcodes.def()
    }
}

impl Related<super::invite::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Invites.def()
    }
}

impl Related<super::user_organization::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::UserOrganizations.def()
    }
}

impl Related<super::custom_field::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::CustomFields.def()
    }
}

impl Related<super::booking::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Bookings.def()
    }
}

impl Related<super::kit::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::Kits.def()
    }
}

impl Related<super::asset_model::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AssetModels.def()
    }
}

impl Related<super::asset_index_settings::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AssetIndexSettings.def()
    }
}

impl Related<super::asset_reminder::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AssetReminders.def()
    }
}

impl Related<super::asset_filter_preset::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AssetFilterPresets.def()
    }
}

impl Related<super::audit_session::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AuditSessions.def()
    }
}

impl Related<super::audit_image::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AuditImages.def()
    }
}

impl Related<super::asset_import::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::AssetImports.def()
    }
}

impl Related<super::link_preview::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::LinkPreviews.def()
    }
}

impl Related<super::pdf_template::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::PdfTemplates.def()
    }
}

impl Related<super::scim_token::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ScimTokens.def()
    }
}

impl Related<super::user_scim_external_id::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ScimExternalIds.def()
    }
}

impl Related<super::role_change_log::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::RoleChangeLogs.def()
    }
}

impl Related<super::team_member_note::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::TeamMemberNotes.def()
    }
}

impl Related<super::activity_event::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::ActivityEvents.def()
    }
}

impl Related<super::working_hours::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::WorkingHours.def()
    }
}

impl Related<super::booking_settings::Entity> for Entity {
    fn to() -> RelationDef {
        Relation::BookingSettings.def()
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
