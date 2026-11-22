from sqlalchemy import ARRAY, Column, Date, Dialect, Enum, ForeignKey, Index, Integer, MetaData, PrimaryKeyConstraint, String, Table, Text, TypeDecorator, text
from sqlalchemy.dialects.postgresql import JSONB, TIMESTAMP
from sqlalchemy.orm import DeclarativeBase, Mapped, mapped_column, relationship
from typing import Any, Optional
from datetime import datetime, date, timezone
import cuid  # type: ignore[import-untyped, unused-ignore]


class UtcDateTime(TypeDecorator[datetime]):
    impl = TIMESTAMP
    cache_ok = True

    def process_bind_param(self, value: Optional[datetime], dialect: Dialect) -> Optional[datetime]:
        if value is None or value.tzinfo is None:
            return value
        return value.astimezone(timezone.utc).replace(tzinfo=None)

    def process_result_value(self, value: Optional[datetime], dialect: Dialect) -> Optional[datetime]:
        return None if value is None else value.replace(tzinfo=timezone.utc)


class UtcDateTimeTz(TypeDecorator[datetime]):
    impl = TIMESTAMP
    cache_ok = True

    def __init__(self, precision: Optional[int] = None) -> None:
        super().__init__(timezone=True, precision=precision)

    def process_bind_param(self, value: Optional[datetime], dialect: Dialect) -> Optional[datetime]:
        return value if value is None or value.tzinfo else value.replace(tzinfo=timezone.utc)

    def process_result_value(self, value: Optional[datetime], dialect: Dialect) -> Optional[datetime]:
        return None if value is None else value.astimezone(timezone.utc)


def utc_now() -> datetime:
    now = datetime.now(timezone.utc)
    return now.replace(microsecond=now.microsecond // 1000 * 1000)


class Base(DeclarativeBase):
    metadata = MetaData(naming_convention={"pk": "%(table_name)s_pkey"})
    type_annotation_map = {datetime: UtcDateTime(precision=3), str: Text}

role_to_user = Table(
    "_RoleToUser",
    Base.metadata,
    Column("A", Text, ForeignKey("Role.id", ondelete="CASCADE", onupdate="CASCADE", name="_RoleToUser_A_fkey"), nullable=False),
    Column("B", Text, ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="_RoleToUser_B_fkey"), nullable=False),
    PrimaryKeyConstraint("A", "B", name="_RoleToUser_AB_pkey"),
    Index("_RoleToUser_B_index", "B"),
)

asset_to_tag = Table(
    "_AssetToTag",
    Base.metadata,
    Column("A", Text, ForeignKey("Asset.id", ondelete="CASCADE", onupdate="CASCADE", name="_AssetToTag_A_fkey"), nullable=False),
    Column("B", Text, ForeignKey("Tag.id", ondelete="CASCADE", onupdate="CASCADE", name="_AssetToTag_B_fkey"), nullable=False),
    PrimaryKeyConstraint("A", "B", name="_AssetToTag_AB_pkey"),
    Index("_AssetToTag_B_index", "B"),
)

category_to_custom_field = Table(
    "_CategoryToCustomField",
    Base.metadata,
    Column("A", Text, ForeignKey("Category.id", ondelete="CASCADE", onupdate="CASCADE", name="_CategoryToCustomField_A_fkey"), nullable=False),
    Column("B", Text, ForeignKey("CustomField.id", ondelete="CASCADE", onupdate="CASCADE", name="_CategoryToCustomField_B_fkey"), nullable=False),
    PrimaryKeyConstraint("A", "B", name="_CategoryToCustomField_AB_pkey"),
    Index("_CategoryToCustomField_B_index", "B"),
)

booking_to_tag = Table(
    "_BookingToTag",
    Base.metadata,
    Column("A", Text, ForeignKey("Booking.id", ondelete="CASCADE", onupdate="CASCADE", name="_BookingToTag_A_fkey"), nullable=False),
    Column("B", Text, ForeignKey("Tag.id", ondelete="CASCADE", onupdate="CASCADE", name="_BookingToTag_B_fkey"), nullable=False),
    PrimaryKeyConstraint("A", "B", name="_BookingToTag_AB_pkey"),
    Index("_BookingToTag_B_index", "B"),
)

asset_reminder_to_team_member = Table(
    "_AssetReminderToTeamMember",
    Base.metadata,
    Column("A", Text, ForeignKey("AssetReminder.id", ondelete="CASCADE", onupdate="CASCADE", name="_AssetReminderToTeamMember_A_fkey"), nullable=False),
    Column("B", Text, ForeignKey("TeamMember.id", ondelete="CASCADE", onupdate="CASCADE", name="_AssetReminderToTeamMember_B_fkey"), nullable=False),
    PrimaryKeyConstraint("A", "B", name="_AssetReminderToTeamMember_AB_pkey"),
    Index("_AssetReminderToTeamMember_B_index", "B"),
)

booking_settings_always_notify = Table(
    "_BookingSettingsAlwaysNotify",
    Base.metadata,
    Column("A", Text, ForeignKey("BookingSettings.id", ondelete="CASCADE", onupdate="CASCADE", name="_BookingSettingsAlwaysNotify_A_fkey"), nullable=False),
    Column("B", Text, ForeignKey("TeamMember.id", ondelete="CASCADE", onupdate="CASCADE", name="_BookingSettingsAlwaysNotify_B_fkey"), nullable=False),
    PrimaryKeyConstraint("A", "B", name="_BookingSettingsAlwaysNotify_AB_pkey"),
    Index("_BookingSettingsAlwaysNotify_B_index", "B"),
)

booking_notification_recipients = Table(
    "_BookingNotificationRecipients",
    Base.metadata,
    Column("A", Text, ForeignKey("Booking.id", ondelete="CASCADE", onupdate="CASCADE", name="_BookingNotificationRecipients_A_fkey"), nullable=False),
    Column("B", Text, ForeignKey("TeamMember.id", ondelete="CASCADE", onupdate="CASCADE", name="_BookingNotificationRecipients_B_fkey"), nullable=False),
    PrimaryKeyConstraint("A", "B", name="_BookingNotificationRecipients_AB_pkey"),
    Index("_BookingNotificationRecipients_B_index", "B"),
)


class Image(Base):
    __tablename__ = "Image"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    content_type: Mapped[str] = mapped_column("contentType")
    alt_text: Mapped[Optional[str]] = mapped_column("altText")
    blob: Mapped[bytes]
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)
    owner_org_id: Mapped[str] = mapped_column("ownerOrgId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="Image_ownerOrgId_fkey"))
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="RESTRICT", onupdate="CASCADE", name="Image_userId_fkey"))

    __table_args__ = (
        Index("Image_ownerOrgId_idx", "ownerOrgId"),
        Index("Image_userId_idx", "userId"),
    )

    owner_org: Mapped["Organization"] = relationship(foreign_keys=[owner_org_id], back_populates="images")
    user: Mapped["User"] = relationship(back_populates="images")
    location: Mapped[Optional["Location"]] = relationship(back_populates="image")
    organization: Mapped[Optional["Organization"]] = relationship(foreign_keys="Organization.image_id", back_populates="image")

class User(Base):
    __tablename__ = "User"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    email: Mapped[str]
    username: Mapped[str] = mapped_column(default=lambda: cuid.cuid())
    first_name: Mapped[Optional[str]] = mapped_column("firstName")
    last_name: Mapped[Optional[str]] = mapped_column("lastName")
    display_name: Mapped[Optional[str]] = mapped_column("displayName")
    profile_picture: Mapped[Optional[str]] = mapped_column("profilePicture")
    used_free_trial: Mapped[bool] = mapped_column("usedFreeTrial", default=False, server_default=text("false"))
    onboarded: Mapped[bool] = mapped_column(default=False, server_default=text("false"))
    customer_id: Mapped[Optional[str]] = mapped_column("customerId")
    sso: Mapped[bool] = mapped_column(default=False, server_default=text("false"))
    created_with_invite: Mapped[bool] = mapped_column("createdWithInvite", default=False, server_default=text("false"))
    skip_subscription_check: Mapped[bool] = mapped_column("skipSubscriptionCheck", default=False, server_default=text("false"))
    has_unpaid_invoice: Mapped[bool] = mapped_column("hasUnpaidInvoice", default=False, server_default=text("false"))
    warn_for_no_payment_method: Mapped[bool] = mapped_column("warnForNoPaymentMethod", default=False, server_default=text("false"))
    last_selected_organization_id: Mapped[Optional[str]] = mapped_column("lastSelectedOrganizationId", ForeignKey("Organization.id", ondelete="SET NULL", onupdate="CASCADE", name="User_lastSelectedOrganizationId_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)
    deleted_at: Mapped[Optional[datetime]] = mapped_column("deletedAt")
    last_mobile_active_at: Mapped[Optional[datetime]] = mapped_column("lastMobileActiveAt")
    date_format: Mapped[Optional[str]] = mapped_column("dateFormat", Enum("DD_MM_YYYY", "MM_DD_YYYY", "YYYY_MM_DD", "MMM_DD_YYYY", "DD_MMM_YYYY", name="DateFormatPreference"))
    time_format: Mapped[Optional[str]] = mapped_column("timeFormat", Enum("H12", "H24", name="TimeFormatPreference"))
    week_start: Mapped[Optional[str]] = mapped_column("weekStart", Enum("MONDAY", "SUNDAY", "SATURDAY", name="WeekStartPreference"))
    time_zone: Mapped[Optional[str]] = mapped_column("timeZone")
    tier_id: Mapped[str] = mapped_column("tierId", Enum("free", "tier_1", "tier_2", "custom", name="TierId"), ForeignKey("Tier.id", ondelete="RESTRICT", onupdate="CASCADE", name="User_tierId_fkey"), default="free", server_default=text("'free'"))
    referral_source: Mapped[Optional[str]] = mapped_column("referralSource")

    __table_args__ = (
        Index("User_email_key", "email", unique=True),
        Index("User_username_key", "username", unique=True),
        Index("User_customerId_key", "customerId", unique=True),
        Index("User_firstName_lastName_idx", "firstName", "lastName"),
        Index("User_tierId_idx", "tierId"),
        Index("User_lastSelectedOrganizationId_idx", "lastSelectedOrganizationId"),
        Index("User_lastMobileActiveAt_idx", "lastMobileActiveAt"),
        Index("User_email_username_key", "email", "username", unique=True),
    )

    last_selected_organization: Mapped[Optional["Organization"]] = relationship(foreign_keys=[last_selected_organization_id], back_populates="users_with_last_selected")
    tier: Mapped["Tier"] = relationship(back_populates="subscribers")
    assets: Mapped[list["Asset"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")
    categories: Mapped[list["Category"]] = relationship(passive_deletes="all", back_populates="user")
    notes: Mapped[list["Note"]] = relationship(back_populates="user")
    booking_notes: Mapped[list["BookingNote"]] = relationship(back_populates="user")
    audit_notes: Mapped[list["AuditNote"]] = relationship(back_populates="user")
    location_notes: Mapped[list["LocationNote"]] = relationship(back_populates="user")
    qr_codes: Mapped[list["Qr"]] = relationship(back_populates="user")
    scans: Mapped[list["Scan"]] = relationship(back_populates="user")
    tags: Mapped[list["Tag"]] = relationship(passive_deletes="all", back_populates="user")
    locations: Mapped[list["Location"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")
    images: Mapped[list["Image"]] = relationship(passive_deletes="all", back_populates="user")
    organizations: Mapped[list["Organization"]] = relationship(foreign_keys="Organization.user_id", cascade="all, delete", passive_deletes=True, back_populates="owner")
    custom_fields: Mapped[list["CustomField"]] = relationship(passive_deletes="all", back_populates="created_by")
    sent_invites: Mapped[list["Invite"]] = relationship(foreign_keys="Invite.inviter_id", cascade="all, delete", passive_deletes=True, back_populates="inviter")
    received_invites: Mapped[list["Invite"]] = relationship(foreign_keys="Invite.invitee_user_id", cascade="all, delete", passive_deletes=True, back_populates="invitee_user")
    team_members: Mapped[list["TeamMember"]] = relationship(back_populates="user")
    user_organizations: Mapped[list["UserOrganization"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")
    bookings: Mapped[list["Booking"]] = relationship(foreign_keys="Booking.creator_id", cascade="all, delete", passive_deletes=True, back_populates="creator")
    custodies: Mapped[list["Booking"]] = relationship(foreign_keys="Booking.custodian_user_id", cascade="all, delete", passive_deletes=True, back_populates="custodian_user")
    created_kits: Mapped[list["Kit"]] = relationship(passive_deletes="all", back_populates="created_by")
    asset_reminders: Mapped[list["AssetReminder"]] = relationship(passive_deletes="all", back_populates="created_by")
    audits_created: Mapped[list["AuditSession"]] = relationship(passive_deletes="all", back_populates="created_by")
    asset_imports: Mapped[list["AssetImport"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="created_by")
    audit_assignments: Mapped[list["AuditAssignment"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")
    audit_assets_scanned: Mapped[list["AuditAsset"]] = relationship(back_populates="scanned_by")
    audit_scans: Mapped[list["AuditScan"]] = relationship(back_populates="scanned_by")
    audit_images_uploaded: Mapped[list["AuditImage"]] = relationship(back_populates="uploaded_by")
    asset_index_settings: Mapped[list["AssetIndexSettings"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")
    partial_checkins: Mapped[list["PartialBookingCheckin"]] = relationship(passive_deletes="all", back_populates="checked_in_by")
    partial_checkouts: Mapped[list["PartialBookingCheckout"]] = relationship(passive_deletes="all", back_populates="checked_out_by")
    mobile_auth_codes: Mapped[list["MobileAuthCode"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")
    push_tokens: Mapped[list["PushToken"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")
    created_updates: Mapped[list["Update"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="created_by")
    update_reads: Mapped[list["UserUpdateRead"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")
    asset_filter_presets: Mapped[list["AssetFilterPreset"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="owner")
    asset_models: Mapped[list["AssetModel"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="created_by")
    consumption_logs: Mapped[list["ConsumptionLog"]] = relationship(passive_deletes="all", back_populates="performed_by")
    scim_external_ids: Mapped[list["UserScimExternalId"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")
    role_changes_received: Mapped[list["RoleChangeLog"]] = relationship(foreign_keys="RoleChangeLog.user_id", passive_deletes="all", back_populates="user")
    role_changes_performed: Mapped[list["RoleChangeLog"]] = relationship(foreign_keys="RoleChangeLog.changed_by_id", passive_deletes="all", back_populates="changed_by")
    team_member_notes_authored: Mapped[list["TeamMemberNote"]] = relationship(back_populates="user")
    contact: Mapped[Optional["UserContact"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")
    business_intel: Mapped[Optional["UserBusinessIntel"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")
    notification_preference: Mapped[Optional["NotificationPreference"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")
    custom_tier_limit: Mapped[Optional["CustomTierLimit"]] = relationship(back_populates="user")
    roles: Mapped[list["Role"]] = relationship(secondary=role_to_user, back_populates="users")

class UserContact(Base):
    __tablename__ = "UserContact"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    phone: Mapped[Optional[str]]
    street: Mapped[Optional[str]]
    city: Mapped[Optional[str]]
    state_province: Mapped[Optional[str]] = mapped_column("stateProvince")
    zip_postal_code: Mapped[Optional[str]] = mapped_column("zipPostalCode")
    country_region: Mapped[Optional[str]] = mapped_column("countryRegion")
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="UserContact_userId_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("UserContact_userId_key", "userId", unique=True),
        Index("UserContact_userId_idx", "userId"),
        Index("UserContact_phone_idx", "phone"),
        Index("UserContact_city_stateProvince_idx", "city", "stateProvince"),
        Index("UserContact_countryRegion_idx", "countryRegion"),
        Index("UserContact_zipPostalCode_idx", "zipPostalCode"),
        Index("UserContact_city_countryRegion_idx", "city", "countryRegion"),
    )

    user: Mapped["User"] = relationship(back_populates="contact")

class UserBusinessIntel(Base):
    __tablename__ = "UserBusinessIntel"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    how_did_you_hear_about_us: Mapped[Optional[str]] = mapped_column("howDidYouHearAboutUs")
    job_title: Mapped[Optional[str]] = mapped_column("jobTitle")
    team_size: Mapped[Optional[str]] = mapped_column("teamSize")
    company_name: Mapped[Optional[str]] = mapped_column("companyName")
    primary_use_case: Mapped[Optional[str]] = mapped_column("primaryUseCase")
    current_solution: Mapped[Optional[str]] = mapped_column("currentSolution")
    timeline: Mapped[Optional[str]]
    signup_plan: Mapped[Optional[str]] = mapped_column("signupPlan")
    signup_trial: Mapped[Optional[bool]] = mapped_column("signupTrial")
    utm_source: Mapped[Optional[str]] = mapped_column("utmSource")
    utm_medium: Mapped[Optional[str]] = mapped_column("utmMedium")
    utm_campaign: Mapped[Optional[str]] = mapped_column("utmCampaign")
    utm_content: Mapped[Optional[str]] = mapped_column("utmContent")
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="UserBusinessIntel_userId_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("UserBusinessIntel_userId_key", "userId", unique=True),
        Index("UserBusinessIntel_userId_idx", "userId"),
        Index("UserBusinessIntel_companyName_idx", "companyName"),
        Index("UserBusinessIntel_jobTitle_idx", "jobTitle"),
        Index("UserBusinessIntel_teamSize_idx", "teamSize"),
        Index("UserBusinessIntel_utmSource_idx", "utmSource"),
        Index("UserBusinessIntel_utmCampaign_idx", "utmCampaign"),
    )

    user: Mapped["User"] = relationship(back_populates="business_intel")

class Asset(Base):
    __tablename__ = "Asset"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    title: Mapped[str]
    description: Mapped[Optional[str]]
    main_image: Mapped[Optional[str]] = mapped_column("mainImage")
    thumbnail_image: Mapped[Optional[str]] = mapped_column("thumbnailImage")
    main_image_expiration: Mapped[Optional[datetime]] = mapped_column("mainImageExpiration")
    status: Mapped[str] = mapped_column(Enum("AVAILABLE", "IN_CUSTODY", "CHECKED_OUT", name="AssetStatus"), default="AVAILABLE", server_default=text("'AVAILABLE'"))
    valuation: Mapped[Optional[float]] = mapped_column("value")
    available_to_book: Mapped[bool] = mapped_column("availableToBook", default=True, server_default=text("true"))
    sequential_id: Mapped[Optional[str]] = mapped_column("sequentialId")
    archived_at: Mapped[Optional[datetime]] = mapped_column("archivedAt", UtcDateTimeTz(precision=3))
    type: Mapped[str] = mapped_column(Enum("INDIVIDUAL", "QUANTITY_TRACKED", name="AssetType"), default="INDIVIDUAL", server_default=text("'INDIVIDUAL'"))
    asset_model_id: Mapped[Optional[str]] = mapped_column("assetModelId", ForeignKey("AssetModel.id", ondelete="SET NULL", onupdate="CASCADE", name="Asset_assetModelId_fkey"))
    quantity: Mapped[Optional[int]]
    min_quantity: Mapped[Optional[int]] = mapped_column("minQuantity")
    consumption_type: Mapped[Optional[str]] = mapped_column("consumptionType", Enum("ONE_WAY", "TWO_WAY", name="ConsumptionType"))
    unit_of_measure: Mapped[Optional[str]] = mapped_column("unitOfMeasure")
    low_stock_notified_at: Mapped[Optional[datetime]] = mapped_column("lowStockNotifiedAt")
    stock_ledger_started_at: Mapped[Optional[datetime]] = mapped_column("stockLedgerStartedAt")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="Asset_userId_fkey"))
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="Asset_organizationId_fkey"))
    category_id: Mapped[Optional[str]] = mapped_column("categoryId", ForeignKey("Category.id", ondelete="SET NULL", onupdate="CASCADE", name="Asset_categoryId_fkey"))
    preferred_barcode_id: Mapped[Optional[str]] = mapped_column("preferredBarcodeId", ForeignKey("Barcode.id", ondelete="SET NULL", onupdate="CASCADE", name="Asset_preferredBarcodeId_fkey"))

    __table_args__ = (
        Index("Asset_preferredBarcodeId_key", "preferredBarcodeId", unique=True),
        Index("Asset_title_description_idx", "title", "description"),
        Index("Asset_sequentialId_trgm_idx", "sequentialId"),
        Index("Asset_organizationId_compound_idx", "organizationId", "title", "status", "availableToBook"),
        Index("Asset_status_organizationId_idx", "status", "organizationId"),
        Index("Asset_organizationId_archivedAt_idx", "organizationId", "archivedAt"),
        Index("Asset_createdAt_organizationId_idx", "createdAt", "organizationId"),
        Index("Asset_organizationId_createdAt_id_idx", "organizationId", "createdAt", "id"),
        Index("Asset_valuation_organizationId_idx", "value", "organizationId"),
        Index("Asset_categoryId_organizationId_idx", "categoryId", "organizationId"),
        Index("Asset_assetModelId_organizationId_idx", "assetModelId", "organizationId"),
        Index("Asset_sequentialId_idx", "sequentialId"),
        Index("Asset_userId_idx", "userId"),
        Index("Asset_organizationId_sequentialId_key", "organizationId", "sequentialId", unique=True),
    )

    asset_model: Mapped[Optional["AssetModel"]] = relationship(back_populates="assets")
    user: Mapped["User"] = relationship(back_populates="assets")
    organization: Mapped["Organization"] = relationship(back_populates="assets")
    category: Mapped[Optional["Category"]] = relationship(back_populates="assets")
    preferred_barcode: Mapped[Optional["Barcode"]] = relationship(foreign_keys=[preferred_barcode_id], back_populates="preferred_by_asset")
    custody: Mapped[list["Custody"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="asset")
    notes: Mapped[list["Note"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="asset")
    qr_codes: Mapped[list["Qr"]] = relationship(back_populates="asset")
    barcodes: Mapped[list["Barcode"]] = relationship(foreign_keys="Barcode.asset_id", back_populates="asset")
    reports: Mapped[list["ReportFound"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="asset")
    custom_fields: Mapped[list["AssetCustomFieldValue"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="asset")
    booking_assets: Mapped[list["BookingAsset"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="asset")
    consumption_logs: Mapped[list["ConsumptionLog"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="asset")
    reminders: Mapped[list["AssetReminder"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="asset")
    audit_entries: Mapped[list["AuditAsset"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="asset")
    audit_scans: Mapped[list["AuditScan"]] = relationship(back_populates="asset")
    scans: Mapped[list["Scan"]] = relationship(back_populates="asset")
    asset_kits: Mapped[list["AssetKit"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="asset")
    asset_locations: Mapped[list["AssetLocation"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="asset")
    tags: Mapped[list["Tag"]] = relationship(secondary=asset_to_tag, back_populates="assets")

class AssetFilterPreset(Base):
    __tablename__ = "AssetFilterPreset"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="AssetFilterPreset_organizationId_fkey"))
    owner_id: Mapped[str] = mapped_column("ownerId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="AssetFilterPreset_ownerId_fkey"))
    name: Mapped[str]
    query: Mapped[str]
    starred: Mapped[bool] = mapped_column(default=False, server_default=text("false"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("asset_filter_presets_owner_lookup_idx", "organizationId", "ownerId"),
        Index("asset_filter_presets_owner_name_unique", "organizationId", "ownerId", "name", unique=True),
    )

    organization: Mapped["Organization"] = relationship(back_populates="asset_filter_presets")
    owner: Mapped["User"] = relationship(back_populates="asset_filter_presets")

class AssetIndexSettings(Base):
    __tablename__ = "AssetIndexSettings"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="AssetIndexSettings_userId_fkey"))
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="AssetIndexSettings_organizationId_fkey"))
    mode: Mapped[str] = mapped_column(Enum("SIMPLE", "ADVANCED", name="AssetIndexMode"), default="SIMPLE", server_default=text("'SIMPLE'"))
    columns: Mapped[dict[str, Any]] = mapped_column(JSONB, default=lambda: [{"name": "id", "visible": True, "position": 0}, {"name": "status", "visible": True, "position": 1}, {"name": "description", "visible": True, "position": 2}, {"name": "valuation", "visible": True, "position": 3}, {"name": "createdAt", "visible": True, "position": 4}, {"name": "category", "visible": True, "position": 5}, {"name": "tags", "visible": True, "position": 6}, {"name": "location", "visible": True, "position": 7}, {"name": "kit", "visible": True, "position": 8}, {"name": "custody", "visible": True, "position": 9}], server_default=text("'[ {\"name\": \"id\", \"visible\": true, \"position\": 0}, {\"name\": \"status\", \"visible\": true, \"position\": 1}, {\"name\": \"description\", \"visible\": true, \"position\": 2}, {\"name\": \"valuation\", \"visible\": true, \"position\": 3}, {\"name\": \"createdAt\", \"visible\": true, \"position\": 4}, {\"name\": \"category\", \"visible\": true, \"position\": 5}, {\"name\": \"tags\", \"visible\": true, \"position\": 6}, {\"name\": \"location\", \"visible\": true, \"position\": 7}, {\"name\": \"kit\", \"visible\": true, \"position\": 8}, {\"name\": \"custody\", \"visible\": true, \"position\": 9} ]'"))
    freeze_column: Mapped[bool] = mapped_column("freezeColumn", default=True, server_default=text("true"))
    show_asset_image: Mapped[bool] = mapped_column("showAssetImage", default=True, server_default=text("true"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("AssetIndexSettings_organizationId_idx", "organizationId"),
        Index("AssetIndexSettings_userId_organizationId_key", "userId", "organizationId", unique=True),
    )

    user: Mapped["User"] = relationship(back_populates="asset_index_settings")
    organization: Mapped["Organization"] = relationship(back_populates="asset_index_settings")

class PdfTemplate(Base):
    """A saved layout for a generated PDF, kept per team workspace under a name:
    which columns, rows and header lines print, and in what order the rows
    come. A row is a layout, not a generated file; the person generating a PDF
    picks a template by name.
    """

    __tablename__ = "PdfTemplate"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    name: Mapped[str]
    kind: Mapped[str] = mapped_column(Enum("BOOKING", name="PdfTemplateKind"), default="BOOKING", server_default=text("'BOOKING'"))
    config: Mapped[dict[str, Any]] = mapped_column(JSONB, default=lambda: {}, server_default=text("'{}'"))
    """The layout's switches. Read through `parseBookingTemplateConfig`,
    which fills any key a stored row lacks from the Checklist template's
    values.
    """
    is_default: Mapped[bool] = mapped_column("isDefault", default=False, server_default=text("false"))
    """The template the PDF dialog offers first. One per workspace and kind,
    enforced by a partial unique index in raw SQL (migration
    20261103120000_add_pdf_templates). Names are also unique per workspace
    and kind ignoring case, through an expression index in the same
    migration.
    """
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="PdfTemplate_organizationId_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("PdfTemplate_organizationId_kind_name_key", "organizationId", "kind", "name", unique=True),
    )

    organization: Mapped["Organization"] = relationship(back_populates="pdf_templates")

class Category(Base):
    __tablename__ = "Category"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    name: Mapped[str]
    description: Mapped[Optional[str]]
    color: Mapped[str]
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="RESTRICT", onupdate="CASCADE", name="Category_userId_fkey"))
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="Category_organizationId_fkey"))

    __table_args__ = (
        Index("Category_organizationId_idx", "organizationId"),
        Index("Category_userId_idx", "userId"),
        Index("Category_name_trgm_idx", "name"),
    )

    user: Mapped["User"] = relationship(back_populates="categories")
    organization: Mapped["Organization"] = relationship(back_populates="categories")
    assets: Mapped[list["Asset"]] = relationship(back_populates="category")
    kits: Mapped[list["Kit"]] = relationship(back_populates="category")
    asset_model_defaults: Mapped[list["AssetModel"]] = relationship(back_populates="default_category")
    custom_fields: Mapped[list["CustomField"]] = relationship(secondary=category_to_custom_field, back_populates="categories")

class AssetModel(Base):
    """*
    * A template/grouping entity for assets.
    * Groups individual assets of the same type (e.g., "Dell Latitude 5550")
    * and provides default field values when creating new assets from the model.
    """

    __tablename__ = "AssetModel"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    name: Mapped[str]
    description: Mapped[Optional[str]]
    image: Mapped[Optional[str]]
    """Public URL (bucket `files`) of the model's cover image. Rendered by any
    asset of this model that has no image of its own. Public — and therefore
    never-expiring and cacheable — like Location.imageUrl; see
    resolveAssetImage() for the cascade.
    """
    thumbnail_image: Mapped[Optional[str]] = mapped_column("thumbnailImage")
    """Public URL of the 108px thumbnail generated alongside `image`. Used at
    list sizes so inheriting assets never render the full-size original.
    """
    default_category_id: Mapped[Optional[str]] = mapped_column("defaultCategoryId", ForeignKey("Category.id", ondelete="SET NULL", onupdate="CASCADE", name="AssetModel_defaultCategoryId_fkey"))
    default_valuation: Mapped[Optional[float]] = mapped_column("defaultValuation")
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="AssetModel_organizationId_fkey"))
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="AssetModel_userId_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("AssetModel_organizationId_name_idx", "organizationId", "name"),
        Index("AssetModel_userId_idx", "userId"),
        Index("AssetModel_name_trgm_idx", "name"),
    )

    default_category: Mapped[Optional["Category"]] = relationship(back_populates="asset_model_defaults")
    organization: Mapped["Organization"] = relationship(back_populates="asset_models")
    created_by: Mapped["User"] = relationship(back_populates="asset_models")
    custom_field_defaults: Mapped[list["AssetModelCustomFieldDefault"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="asset_model")
    """Default custom field values. Like the category and valuation defaults,
    each one fills the matching field only when it is empty, at the moment
    an asset joins this model. Nothing stays linked afterwards.
    """
    assets: Mapped[list["Asset"]] = relationship(back_populates="asset_model")
    model_requests: Mapped[list["BookingModelRequest"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="asset_model")

class BookingModelRequest(Base):
    """Book-by-Model: intent row attached to a booking for a quantity of
    assets of a given AssetModel, without yet picking specific assets.
    Concrete `BookingAsset` rows are only created at
    scan-to-assign time; when the last unit is assigned, the request
    row is deleted. `quantity` is the number of units still needing
    assignment (starts at the full requested total, decrements on scan).

    Keeping model requests separate from `BookingAsset` means the ~40
    code paths that assume `BookingAsset.assetId` points to a real
    asset stay untouched — model-only bookings become concrete the
    moment they're checked out.
    """

    __tablename__ = "BookingModelRequest"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    booking_id: Mapped[str] = mapped_column("bookingId", ForeignKey("Booking.id", ondelete="CASCADE", onupdate="CASCADE", name="BookingModelRequest_bookingId_fkey"))
    asset_model_id: Mapped[str] = mapped_column("assetModelId", ForeignKey("AssetModel.id", ondelete="CASCADE", onupdate="CASCADE", name="BookingModelRequest_assetModelId_fkey"))
    quantity: Mapped[int]
    """Total units originally reserved. Mutated only by explicit
    operator edits (upsert) — never decremented on materialisation.
    This preserves the audit trail of "how many were planned".
    """
    fulfilled_quantity: Mapped[int] = mapped_column("fulfilledQuantity", default=0, server_default=text("0"))
    """Units already assigned to concrete `BookingAsset` rows via
    `materializeModelRequestForAsset`. Scans increment it.

    It DOES decrease: `removeAssets` (booking service) decrements it by the
    number of removed `BookingAsset` rows whose `bookingModelRequestId`
    points at this request, floored at 0, and clears `fulfilledAt` when the
    result drops below `quantity`.

    Counted from that column, never from a shared `assetModelId`: an asset
    added directly alongside a reservation shares the model without having
    discharged anything, so grouping by model re-opened reservations that
    were still correctly satisfied.
    """
    fulfilled_at: Mapped[Optional[datetime]] = mapped_column("fulfilledAt")
    """Timestamp of the moment the reservation became fully fulfilled — set
    when `fulfilledQuantity` catches up to `quantity`, and preserved (not
    re-stamped) if an unchanged quantity is saved again. Cleared back to
    null when outstanding units reappear: the operator edits `quantity`
    upwards, or `removeAssets` drops `fulfilledQuantity` below it.

    The invariant that makes this the primary "is this outstanding?" filter:
    non-null IFF `fulfilledQuantity === quantity`, so `fulfilledAt IS NULL`
    means units remain to materialise. `fulfilledQuantity` can never exceed
    `quantity` (the claim in `materializeModelRequestForAsset` is a single
    statement guarded by `WHERE "fulfilledQuantity" < "quantity"`; upsert
    rejects a shrink below it), so any decrement re-opens the row.
    """
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("BookingModelRequest_assetModelId_idx", "assetModelId"),
        Index("BookingModelRequest_bookingId_assetModelId_key", "bookingId", "assetModelId", unique=True),
    )

    booking: Mapped["Booking"] = relationship(back_populates="model_requests")
    asset_model: Mapped["AssetModel"] = relationship(back_populates="model_requests")

class Tag(Base):
    __tablename__ = "Tag"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    name: Mapped[str]
    description: Mapped[Optional[str]]
    color: Mapped[Optional[str]]
    use_for: Mapped[list[str]] = mapped_column("useFor", ARRAY(Enum("ASSET", "BOOKING", name="TagUseFor")), nullable=True, default=lambda: ["ASSET"], server_default=text("ARRAY['ASSET']::\"TagUseFor\"[]"))
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="RESTRICT", onupdate="CASCADE", name="Tag_userId_fkey"))
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="Tag_organizationId_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("Tag_organizationId_idx", "organizationId"),
        Index("Tag_userId_idx", "userId"),
        Index("Tag_name_trgm_idx", "name"),
    )

    user: Mapped["User"] = relationship(back_populates="tags")
    organization: Mapped["Organization"] = relationship(back_populates="tags")
    assets: Mapped[list["Asset"]] = relationship(secondary=asset_to_tag, back_populates="tags")
    bookings: Mapped[list["Booking"]] = relationship(secondary=booking_to_tag, back_populates="tags")

class Note(Base):
    __tablename__ = "Note"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    content: Mapped[str]
    type: Mapped[str] = mapped_column(Enum("COMMENT", "UPDATE", name="NoteType"), default="COMMENT", server_default=text("'COMMENT'"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)
    user_id: Mapped[Optional[str]] = mapped_column("userId", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="Note_userId_fkey"))
    asset_id: Mapped[str] = mapped_column("assetId", ForeignKey("Asset.id", ondelete="CASCADE", onupdate="CASCADE", name="Note_assetId_fkey"))

    __table_args__ = (
        Index("Note_assetId_idx", "assetId"),
        Index("Note_userId_idx", "userId"),
    )

    user: Mapped[Optional["User"]] = relationship(back_populates="notes")
    asset: Mapped["Asset"] = relationship(back_populates="notes")

class BookingNote(Base):
    __tablename__ = "BookingNote"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    content: Mapped[str]
    type: Mapped[str] = mapped_column(Enum("COMMENT", "UPDATE", name="NoteType"), default="COMMENT", server_default=text("'COMMENT'"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)
    user_id: Mapped[Optional[str]] = mapped_column("userId", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="BookingNote_userId_fkey"))
    booking_id: Mapped[str] = mapped_column("bookingId", ForeignKey("Booking.id", ondelete="CASCADE", onupdate="CASCADE", name="BookingNote_bookingId_fkey"))

    __table_args__ = (
        Index("BookingNote_bookingId_idx", "bookingId"),
        Index("BookingNote_userId_idx", "userId"),
    )

    user: Mapped[Optional["User"]] = relationship(back_populates="booking_notes")
    booking: Mapped["Booking"] = relationship(back_populates="notes")

class LocationNote(Base):
    __tablename__ = "LocationNote"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    content: Mapped[str]
    type: Mapped[str] = mapped_column(Enum("COMMENT", "UPDATE", name="NoteType"), default="COMMENT", server_default=text("'COMMENT'"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)
    user_id: Mapped[Optional[str]] = mapped_column("userId", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="LocationNote_userId_fkey"))
    location_id: Mapped[str] = mapped_column("locationId", ForeignKey("Location.id", ondelete="CASCADE", onupdate="CASCADE", name="LocationNote_locationId_fkey"))

    __table_args__ = (
        Index("LocationNote_locationId_idx", "locationId"),
        Index("LocationNote_userId_idx", "userId"),
    )

    user: Mapped[Optional["User"]] = relationship(back_populates="location_notes")
    location: Mapped["Location"] = relationship(back_populates="notes")

class Qr(Base):
    __tablename__ = "Qr"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    version: Mapped[int] = mapped_column(default=0, server_default=text("0"))
    error_correction: Mapped[str] = mapped_column("errorCorrection", Enum("L", "M", "Q", "H", name="ErrorCorrection"), default="L", server_default=text("'L'"))
    asset_id: Mapped[Optional[str]] = mapped_column("assetId", ForeignKey("Asset.id", ondelete="SET NULL", onupdate="CASCADE", name="Qr_assetId_fkey"))
    kit_id: Mapped[Optional[str]] = mapped_column("kitId", ForeignKey("Kit.id", ondelete="SET NULL", onupdate="CASCADE", name="Qr_kitId_fkey"))
    user_id: Mapped[Optional[str]] = mapped_column("userId", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="Qr_userId_fkey"))
    organization_id: Mapped[Optional[str]] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="SET NULL", onupdate="CASCADE", name="Qr_organizationId_fkey"))
    batch_id: Mapped[Optional[str]] = mapped_column("batchId", ForeignKey("PrintBatch.id", ondelete="SET NULL", onupdate="CASCADE", name="Qr_batchId_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("Qr_assetId_idx", "assetId"),
        Index("Qr_kitId_idx", "kitId"),
        Index("Qr_userId_idx", "userId"),
        Index("Qr_organizationId_idx", "organizationId"),
        Index("Qr_batchId_idx", "batchId"),
        Index("Qr_id_trgm_idx", "id"),
    )

    asset: Mapped[Optional["Asset"]] = relationship(back_populates="qr_codes")
    kit: Mapped[Optional["Kit"]] = relationship(back_populates="qr_codes")
    user: Mapped[Optional["User"]] = relationship(back_populates="qr_codes")
    organization: Mapped[Optional["Organization"]] = relationship(back_populates="qr_codes")
    batch: Mapped[Optional["PrintBatch"]] = relationship(back_populates="qr_codes")
    scans: Mapped[list["Scan"]] = relationship(back_populates="qr")

class Barcode(Base):
    __tablename__ = "Barcode"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    value: Mapped[str]
    type: Mapped[str] = mapped_column(Enum("Code128", "Code39", "DataMatrix", "ExternalQR", "EAN13", name="BarcodeType"), default="Code128", server_default=text("'Code128'"))
    asset_id: Mapped[Optional[str]] = mapped_column("assetId", ForeignKey("Asset.id", ondelete="SET NULL", onupdate="CASCADE", name="Barcode_assetId_fkey"))
    kit_id: Mapped[Optional[str]] = mapped_column("kitId", ForeignKey("Kit.id", ondelete="SET NULL", onupdate="CASCADE", name="Barcode_kitId_fkey"))
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="Barcode_organizationId_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("Barcode_organizationId_value_idx", "organizationId", "value"),
        Index("Barcode_assetId_idx", "assetId"),
        Index("Barcode_kitId_idx", "kitId"),
        Index("Barcode_organizationId_idx", "organizationId"),
        Index("Barcode_value_trgm_idx", "value"),
        Index("Barcode_organizationId_value_key", "organizationId", "value", unique=True),
    )

    asset: Mapped[Optional["Asset"]] = relationship(foreign_keys=[asset_id], back_populates="barcodes")
    kit: Mapped[Optional["Kit"]] = relationship(back_populates="barcodes")
    organization: Mapped["Organization"] = relationship(back_populates="barcodes")
    scans: Mapped[list["Scan"]] = relationship(back_populates="barcode")
    preferred_by_asset: Mapped[Optional["Asset"]] = relationship(foreign_keys="Asset.preferred_barcode_id", back_populates="preferred_barcode")

class PrintBatch(Base):
    __tablename__ = "PrintBatch"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    name: Mapped[str]
    printed: Mapped[bool] = mapped_column(default=False, server_default=text("false"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("PrintBatch_name_key", "name", unique=True),
    )

    qr_codes: Mapped[list["Qr"]] = relationship(back_populates="batch")

class ReportFound(Base):
    __tablename__ = "ReportFound"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    email: Mapped[str]
    content: Mapped[str]
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)
    asset_id: Mapped[Optional[str]] = mapped_column("assetId", ForeignKey("Asset.id", ondelete="CASCADE", onupdate="CASCADE", name="ReportFound_assetId_fkey"))
    kit_id: Mapped[Optional[str]] = mapped_column("kitId", ForeignKey("Kit.id", ondelete="CASCADE", onupdate="CASCADE", name="ReportFound_kitId_fkey"))

    __table_args__ = (
        Index("ReportFound_assetId_idx", "assetId"),
        Index("ReportFound_kitId_idx", "kitId"),
    )

    asset: Mapped[Optional["Asset"]] = relationship(back_populates="reports")
    kit: Mapped[Optional["Kit"]] = relationship(back_populates="reports")

class Scan(Base):
    __tablename__ = "Scan"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    latitude: Mapped[Optional[str]]
    longitude: Mapped[Optional[str]]
    user_agent: Mapped[Optional[str]] = mapped_column("userAgent")
    user_id: Mapped[Optional[str]] = mapped_column("userId", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="Scan_userId_fkey"))
    qr_id: Mapped[Optional[str]] = mapped_column("qrId", ForeignKey("Qr.id", ondelete="SET NULL", onupdate="CASCADE", name="Scan_qrId_fkey"))
    raw_qr_id: Mapped[Optional[str]] = mapped_column("rawQrId")
    code_type: Mapped[str] = mapped_column("codeType", Enum("QR", "BARCODE", "SAM_ID", name="ScanCodeType"), default="QR", server_default=text("'QR'"))
    code: Mapped[Optional[str]]
    barcode_id: Mapped[Optional[str]] = mapped_column("barcodeId", ForeignKey("Barcode.id", ondelete="SET NULL", onupdate="CASCADE", name="Scan_barcodeId_fkey"))
    asset_id: Mapped[Optional[str]] = mapped_column("assetId", ForeignKey("Asset.id", ondelete="SET NULL", onupdate="CASCADE", name="Scan_assetId_fkey"))
    kit_id: Mapped[Optional[str]] = mapped_column("kitId", ForeignKey("Kit.id", ondelete="SET NULL", onupdate="CASCADE", name="Scan_kitId_fkey"))
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="Scan_organizationId_fkey"))
    source: Mapped[Optional[str]] = mapped_column(Enum("QR_LINK", "WEB_SCANNER", "WEB_DRAWER", "AUDIT", "COMPANION", "GPS_UPDATE", name="ScanSource"))
    manually_generated: Mapped[bool] = mapped_column("manuallyGenerated", default=False, server_default=text("false"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("Scan_qrId_idx", "qrId"),
        Index("Scan_userId_idx", "userId"),
        Index("Scan_rawQrId_createdAt_idx", "rawQrId", "createdAt"),
        Index("Scan_assetId_createdAt_idx", "assetId", "createdAt"),
        Index("Scan_kitId_createdAt_idx", "kitId", "createdAt"),
        Index("Scan_barcodeId_idx", "barcodeId"),
        Index("Scan_organizationId_idx", "organizationId"),
    )

    user: Mapped[Optional["User"]] = relationship(back_populates="scans")
    qr: Mapped[Optional["Qr"]] = relationship(back_populates="scans")
    barcode: Mapped[Optional["Barcode"]] = relationship(back_populates="scans")
    asset: Mapped[Optional["Asset"]] = relationship(back_populates="scans")
    kit: Mapped[Optional["Kit"]] = relationship(back_populates="scans")
    organization: Mapped["Organization"] = relationship(back_populates="scans")

class Location(Base):
    __tablename__ = "Location"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    name: Mapped[str]
    description: Mapped[Optional[str]]
    address: Mapped[Optional[str]]
    latitude: Mapped[Optional[float]]
    longitude: Mapped[Optional[float]]
    image_url: Mapped[Optional[str]] = mapped_column("imageUrl")
    thumbnail_url: Mapped[Optional[str]] = mapped_column("thumbnailUrl")
    image_id: Mapped[Optional[str]] = mapped_column("imageId", ForeignKey("Image.id", ondelete="SET NULL", onupdate="CASCADE", name="Location_imageId_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="Location_userId_fkey"))
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="Location_organizationId_fkey"))
    parent_id: Mapped[Optional[str]] = mapped_column("parentId", ForeignKey("Location.id", ondelete="SET NULL", onupdate="CASCADE", name="Location_parentId_fkey"))

    __table_args__ = (
        Index("Location_imageId_key", "imageId", unique=True),
        Index("Location_organizationId_idx", "organizationId"),
        Index("Location_userId_idx", "userId"),
        Index("Location_organizationId_parentId_idx", "organizationId", "parentId"),
        Index("Location_name_trgm_idx", "name"),
        Index("Location_description_trgm_idx", "description"),
        Index("Location_address_trgm_idx", "address"),
    )

    image: Mapped[Optional["Image"]] = relationship(back_populates="location")
    user: Mapped["User"] = relationship(back_populates="locations")
    organization: Mapped["Organization"] = relationship(back_populates="locations")
    parent: Mapped[Optional["Location"]] = relationship(remote_side=[id], back_populates="children")
    asset_locations: Mapped[list["AssetLocation"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="location")
    kits: Mapped[list["Kit"]] = relationship(back_populates="location")
    notes: Mapped[list["LocationNote"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="location")
    custodies: Mapped[list["Custody"]] = relationship(back_populates="location")
    consumption_logs: Mapped[list["ConsumptionLog"]] = relationship(back_populates="location")
    children: Mapped[list["Location"]] = relationship(back_populates="parent")

class Role(Base):
    __tablename__ = "Role"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    name: Mapped[str] = mapped_column(Enum("USER", "ADMIN", name="Roles"), default="USER", server_default=text("'USER'"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("Role_name_key", "name", unique=True),
    )

    users: Mapped[list["User"]] = relationship(secondary=role_to_user, back_populates="roles")

class TeamMember(Base):
    __tablename__ = "TeamMember"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    name: Mapped[str]
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="TeamMember_organizationId_fkey"))
    user_id: Mapped[Optional[str]] = mapped_column("userId", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="TeamMember_userId_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)
    deleted_at: Mapped[Optional[datetime]] = mapped_column("deletedAt")

    __table_args__ = (
        Index("TeamMember_name_idx", "name"),
        Index("TeamMember_organizationId_idx", "organizationId"),
        Index("TeamMember_userId_idx", "userId"),
    )

    organization: Mapped["Organization"] = relationship(back_populates="members")
    user: Mapped[Optional["User"]] = relationship(back_populates="team_members")
    custodies: Mapped[list["Custody"]] = relationship(passive_deletes="all", back_populates="custodian")
    received_invites: Mapped[list["Invite"]] = relationship(passive_deletes="all", back_populates="invitee_team_member")
    kit_custodies: Mapped[list["KitCustody"]] = relationship(passive_deletes="all", back_populates="custodian")
    bookings: Mapped[list["Booking"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="custodian_team_member")
    consumption_logs: Mapped[list["ConsumptionLog"]] = relationship(back_populates="custodian")
    team_member_notes: Mapped[list["TeamMemberNote"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="team_member")
    asset_reminders: Mapped[list["AssetReminder"]] = relationship(secondary=asset_reminder_to_team_member, back_populates="team_members")
    always_notify_booking_settings: Mapped[list["BookingSettings"]] = relationship(secondary=booking_settings_always_notify, back_populates="always_notify_team_members")
    booking_notifications: Mapped[list["Booking"]] = relationship(secondary=booking_notification_recipients, back_populates="notification_recipients")

class Custody(Base):
    __tablename__ = "Custody"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    team_member_id: Mapped[str] = mapped_column("teamMemberId", ForeignKey("TeamMember.id", ondelete="RESTRICT", onupdate="CASCADE", name="Custody_teamMemberId_fkey"))
    asset_id: Mapped[str] = mapped_column("assetId", ForeignKey("Asset.id", ondelete="CASCADE", onupdate="CASCADE", name="Custody_assetId_fkey"))
    kit_custody_id: Mapped[Optional[str]] = mapped_column("kitCustodyId", ForeignKey("KitCustody.id", ondelete="CASCADE", onupdate="CASCADE", name="Custody_kitCustodyId_fkey"))
    quantity: Mapped[int] = mapped_column(default=1, server_default=text("1"))
    location_id: Mapped[Optional[str]] = mapped_column("locationId", ForeignKey("Location.id", ondelete="SET NULL", onupdate="CASCADE", name="Custody_locationId_fkey"))
    source_unknown: Mapped[bool] = mapped_column("sourceUnknown", default=False, server_default=text("false"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("Custody_assetId_teamMemberId_idx", "assetId", "teamMemberId"),
        Index("Custody_teamMemberId_idx", "teamMemberId"),
        Index("Custody_kitCustodyId_idx", "kitCustodyId"),
        Index("Custody_locationId_idx", "locationId"),
    )

    custodian: Mapped["TeamMember"] = relationship(back_populates="custodies")
    asset: Mapped["Asset"] = relationship(back_populates="custody")
    kit_custody: Mapped[Optional["KitCustody"]] = relationship(back_populates="inherited_custody")
    location: Mapped[Optional["Location"]] = relationship(back_populates="custodies")

class Organization(Base):
    __tablename__ = "Organization"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    name: Mapped[str] = mapped_column(default="Personal", server_default=text("'Personal'"))
    type: Mapped[str] = mapped_column(Enum("PERSONAL", "TEAM", name="OrganizationType"), default="PERSONAL", server_default=text("'PERSONAL'"))
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="Organization_userId_fkey"))
    currency: Mapped[str] = mapped_column(Enum("AED", "AFN", "ALL", "AMD", "ANG", "AOA", "ARS", "AUD", "AWG", "AZN", "BAM", "BBD", "BDT", "BGN", "BHD", "BIF", "BMD", "BND", "BOB", "BRL", "BSD", "BTN", "BWP", "BYN", "BZD", "CAD", "CDF", "CHF", "CLP", "CNY", "COP", "CRC", "CUP", "CVE", "CZK", "DJF", "DKK", "DOP", "DZD", "EGP", "ERN", "ETB", "EUR", "FJD", "FKP", "GBP", "GEL", "GHS", "GIP", "GMD", "GNF", "GTQ", "GYD", "HKD", "HNL", "HTG", "HUF", "IDR", "ILS", "INR", "IQD", "IRR", "ISK", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LYD", "MAD", "MDL", "MGA", "MKD", "MMK", "MNT", "MOP", "MRU", "MUR", "MVR", "MWK", "MXN", "MYR", "MZN", "NAD", "NGN", "NIO", "NOK", "NPR", "NZD", "OMR", "PAB", "PEN", "PGK", "PHP", "PKR", "PLN", "PYG", "QAR", "RON", "RSD", "RUB", "RWF", "SAR", "SBD", "SCR", "SDG", "SEK", "SGD", "SHP", "SLE", "SOS", "SRD", "SSP", "STN", "SVC", "SYP", "SZL", "THB", "TJS", "TMT", "TND", "TOP", "TRY", "TTD", "TWD", "TZS", "UAH", "UGX", "USD", "UYU", "UZS", "VES", "VND", "VUV", "WST", "XAF", "XCD", "XOF", "XPF", "YER", "ZAR", "ZMW", "ZWL", name="Currency"), default="USD", server_default=text("'USD'"))
    image_id: Mapped[Optional[str]] = mapped_column("imageId", ForeignKey("Image.id", ondelete="SET NULL", onupdate="CASCADE", name="Organization_imageId_fkey"))
    enabled_sso: Mapped[bool] = mapped_column("enabledSso", default=False, server_default=text("false"))
    sso_details_id: Mapped[Optional[str]] = mapped_column("ssoDetailsId", ForeignKey("SsoDetails.id", ondelete="SET NULL", onupdate="CASCADE", name="Organization_ssoDetailsId_fkey"))
    self_service_can_see_custody: Mapped[bool] = mapped_column("selfServiceCanSeeCustody", default=False, server_default=text("false"))
    self_service_can_see_bookings: Mapped[bool] = mapped_column("selfServiceCanSeeBookings", default=False, server_default=text("false"))
    base_user_can_see_custody: Mapped[bool] = mapped_column("baseUserCanSeeCustody", default=False, server_default=text("false"))
    base_user_can_see_bookings: Mapped[bool] = mapped_column("baseUserCanSeeBookings", default=False, server_default=text("false"))
    managers_can_edit_assets: Mapped[bool] = mapped_column("managersCanEditAssets", default=False, server_default=text("false"))
    barcodes_enabled: Mapped[bool] = mapped_column("barcodesEnabled", default=False, server_default=text("false"))
    barcodes_enabled_at: Mapped[Optional[datetime]] = mapped_column("barcodesEnabledAt")
    used_barcode_trial: Mapped[bool] = mapped_column("usedBarcodeTrial", default=False, server_default=text("false"))
    audits_enabled: Mapped[bool] = mapped_column("auditsEnabled", default=False, server_default=text("false"))
    audits_enabled_at: Mapped[Optional[datetime]] = mapped_column("auditsEnabledAt")
    used_audit_trial: Mapped[bool] = mapped_column("usedAuditTrial", default=False, server_default=text("false"))
    workspace_disabled: Mapped[bool] = mapped_column("workspaceDisabled", default=False, server_default=text("false"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)
    has_sequential_ids_migrated: Mapped[bool] = mapped_column("hasSequentialIdsMigrated", default=False, server_default=text("false"))
    qr_id_display_preference: Mapped[str] = mapped_column("qrIdDisplayPreference", Enum("QR_ID", "SAM_ID", "Code128", "Code39", "DataMatrix", "ExternalQR", "EAN13", name="QrIdDisplayPreference"), default="QR_ID", server_default=text("'QR_ID'"))
    show_shelf_branding: Mapped[bool] = mapped_column("showShelfBranding", default=True, server_default=text("true"))
    show_qr_codes_on_pdfs: Mapped[bool] = mapped_column("showQrCodesOnPdfs", default=True, server_default=text("true"))
    custom_email_footer: Mapped[Optional[str]] = mapped_column("customEmailFooter", String(500))

    __table_args__ = (
        Index("Organization_imageId_key", "imageId", unique=True),
        Index("Organization_userId_idx", "userId"),
        Index("Organization_ssoDetailsId_idx", "ssoDetailsId"),
    )

    owner: Mapped["User"] = relationship(foreign_keys=[user_id], back_populates="organizations")
    image: Mapped[Optional["Image"]] = relationship(foreign_keys=[image_id], back_populates="organization")
    sso_details: Mapped[Optional["SsoDetails"]] = relationship(back_populates="organizations")
    members: Mapped[list["TeamMember"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    assets: Mapped[list["Asset"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    asset_kits: Mapped[list["AssetKit"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    asset_locations: Mapped[list["AssetLocation"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    scans: Mapped[list["Scan"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    locations: Mapped[list["Location"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    categories: Mapped[list["Category"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    tags: Mapped[list["Tag"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    qr_codes: Mapped[list["Qr"]] = relationship(back_populates="organization")
    barcodes: Mapped[list["Barcode"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    invites: Mapped[list["Invite"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    user_organizations: Mapped[list["UserOrganization"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    custom_fields: Mapped[list["CustomField"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    images: Mapped[list["Image"]] = relationship(foreign_keys="Image.owner_org_id", cascade="all, delete", passive_deletes=True, back_populates="owner_org")
    bookings: Mapped[list["Booking"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    kits: Mapped[list["Kit"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    asset_models: Mapped[list["AssetModel"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    asset_index_settings: Mapped[list["AssetIndexSettings"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    asset_reminders: Mapped[list["AssetReminder"]] = relationship(passive_deletes="all", back_populates="organization")
    asset_filter_presets: Mapped[list["AssetFilterPreset"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    audit_sessions: Mapped[list["AuditSession"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    audit_images: Mapped[list["AuditImage"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    asset_imports: Mapped[list["AssetImport"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    link_previews: Mapped[list["LinkPreview"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    pdf_templates: Mapped[list["PdfTemplate"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    users_with_last_selected: Mapped[list["User"]] = relationship(foreign_keys="User.last_selected_organization_id", back_populates="last_selected_organization")
    scim_tokens: Mapped[list["ScimToken"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    scim_external_ids: Mapped[list["UserScimExternalId"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    role_change_logs: Mapped[list["RoleChangeLog"]] = relationship(passive_deletes="all", back_populates="organization")
    team_member_notes: Mapped[list["TeamMemberNote"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    activity_events: Mapped[list["ActivityEvent"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    working_hours: Mapped[Optional["WorkingHours"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    booking_settings: Mapped[Optional["BookingSettings"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")

class UserOrganization(Base):
    __tablename__ = "UserOrganization"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="UserOrganization_userId_fkey"))
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="UserOrganization_organizationId_fkey"))
    roles: Mapped[list[str]] = mapped_column(ARRAY(Enum("ADMIN", "BASE", "OWNER", "SELF_SERVICE", "MANAGER", name="OrganizationRoles")), nullable=True)
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)
    calendar_token_id: Mapped[Optional[str]] = mapped_column("calendarTokenId")
    """Secret, unguessable token for this member's iCal calendar-feed subscription
    in this workspace. Null until the user generates a feed; regenerating it
    rotates the value (invalidating any previously-shared subscription URL).
    `@unique` doubles as the lookup index for the public feed route.
    """

    __table_args__ = (
        Index("UserOrganization_calendarTokenId_key", "calendarTokenId", unique=True),
        Index("UserOrganization_organizationId_idx", "organizationId"),
        Index("UserOrganization_userId_organizationId_key", "userId", "organizationId", unique=True),
    )

    user: Mapped["User"] = relationship(back_populates="user_organizations")
    organization: Mapped["Organization"] = relationship(back_populates="user_organizations")

class SsoDetails(Base):
    __tablename__ = "SsoDetails"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    domain: Mapped[str]
    base_user_group_id: Mapped[Optional[str]] = mapped_column("baseUserGroupId")
    self_service_group_id: Mapped[Optional[str]] = mapped_column("selfServiceGroupId")
    admin_group_id: Mapped[Optional[str]] = mapped_column("adminGroupId")
    manager_group_id: Mapped[Optional[str]] = mapped_column("managerGroupId")
    require_sso_login: Mapped[bool] = mapped_column("requireSsoLogin", default=True, server_default=text("true"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    organizations: Mapped[list["Organization"]] = relationship(back_populates="sso_details")

class ScimToken(Base):
    __tablename__ = "ScimToken"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    token_hash: Mapped[str] = mapped_column("tokenHash")
    label: Mapped[str]
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="ScimToken_organizationId_fkey"))
    created_by_id: Mapped[str] = mapped_column("createdById")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    last_used_at: Mapped[Optional[datetime]] = mapped_column("lastUsedAt")

    __table_args__ = (
        Index("ScimToken_tokenHash_key", "tokenHash", unique=True),
        Index("ScimToken_organizationId_idx", "organizationId"),
    )

    organization: Mapped["Organization"] = relationship(back_populates="scim_tokens")

class UserScimExternalId(Base):
    """*
    * Org-scoped SCIM external ID linking a Shelf user to an identity-provider
    * object within a specific organization.  A single Shelf user can appear in
    * multiple organizations, each managed by a different identity provider, so the
    * external ID must be scoped to the (user, organization) pair.
    """

    __tablename__ = "UserScimExternalId"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="UserScimExternalId_userId_fkey"))
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="UserScimExternalId_organizationId_fkey"))
    scim_external_id: Mapped[str] = mapped_column("scimExternalId")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("UserScimExternalId_organizationId_idx", "organizationId"),
        Index("UserScimExternalId_userId_organizationId_key", "userId", "organizationId", unique=True),
        Index("UserScimExternalId_organizationId_scimExternalId_key", "organizationId", "scimExternalId", unique=True),
    )

    user: Mapped["User"] = relationship(back_populates="scim_external_ids")
    organization: Mapped["Organization"] = relationship(back_populates="scim_external_ids")

class Tier(Base):
    __tablename__ = "Tier"

    id: Mapped[str] = mapped_column(Enum("free", "tier_1", "tier_2", "custom", name="TierId"), primary_key=True)
    name: Mapped[str]
    tier_limit_id: Mapped[Optional[str]] = mapped_column("tierLimitId", Enum("free", "tier_1", "tier_2", "custom", name="TierId"), ForeignKey("TierLimit.id", ondelete="SET NULL", onupdate="CASCADE", name="Tier_tierLimitId_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("Tier_tierLimitId_key", "tierLimitId", unique=True),
    )

    tier_limit: Mapped[Optional["TierLimit"]] = relationship(back_populates="tier")
    subscribers: Mapped[list["User"]] = relationship(passive_deletes="all", back_populates="tier")

class TierLimit(Base):
    __tablename__ = "TierLimit"

    id: Mapped[str] = mapped_column(Enum("free", "tier_1", "tier_2", "custom", name="TierId"), primary_key=True)
    can_import_assets: Mapped[bool] = mapped_column("canImportAssets", default=False, server_default=text("false"))
    can_export_assets: Mapped[bool] = mapped_column("canExportAssets", default=False, server_default=text("false"))
    can_import_nrm: Mapped[bool] = mapped_column("canImportNRM", default=False, server_default=text("false"))
    can_hide_shelf_branding: Mapped[bool] = mapped_column("canHideShelfBranding", default=False, server_default=text("false"))
    can_use_reports: Mapped[bool] = mapped_column("canUseReports", default=False, server_default=text("false"))
    can_use_recurring_reminders: Mapped[bool] = mapped_column("canUseRecurringReminders", default=False, server_default=text("false"))
    can_use_push_notifications: Mapped[bool] = mapped_column("canUsePushNotifications", default=False, server_default=text("false"))
    max_custom_fields: Mapped[int] = mapped_column("maxCustomFields", default=0, server_default=text("0"))
    max_organizations: Mapped[int] = mapped_column("maxOrganizations", default=1, server_default=text("1"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    tier: Mapped[Optional["Tier"]] = relationship(back_populates="tier_limit")

class CustomTierLimit(Base):
    __tablename__ = "CustomTierLimit"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    user_id: Mapped[Optional[str]] = mapped_column("userId", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="CustomTierLimit_userId_fkey"))
    can_import_assets: Mapped[bool] = mapped_column("canImportAssets", default=True, server_default=text("true"))
    can_export_assets: Mapped[bool] = mapped_column("canExportAssets", default=True, server_default=text("true"))
    can_import_nrm: Mapped[bool] = mapped_column("canImportNRM", default=True, server_default=text("true"))
    can_hide_shelf_branding: Mapped[bool] = mapped_column("canHideShelfBranding", default=True, server_default=text("true"))
    can_use_reports: Mapped[bool] = mapped_column("canUseReports", default=True, server_default=text("true"))
    can_use_recurring_reminders: Mapped[bool] = mapped_column("canUseRecurringReminders", default=True, server_default=text("true"))
    can_use_push_notifications: Mapped[bool] = mapped_column("canUsePushNotifications", default=True, server_default=text("true"))
    max_custom_fields: Mapped[int] = mapped_column("maxCustomFields", default=1000, server_default=text("1000"))
    max_organizations: Mapped[int] = mapped_column("maxOrganizations", default=1, server_default=text("1"))
    is_enterprise: Mapped[bool] = mapped_column("isEnterprise", default=False, server_default=text("false"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("CustomTierLimit_userId_key", "userId", unique=True),
    )

    user: Mapped[Optional["User"]] = relationship(back_populates="custom_tier_limit")

class CustomField(Base):
    __tablename__ = "CustomField"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    name: Mapped[str]
    help_text: Mapped[Optional[str]] = mapped_column("helpText")
    required: Mapped[bool] = mapped_column(default=False, server_default=text("false"))
    active: Mapped[bool] = mapped_column(default=True, server_default=text("true"))
    type: Mapped[str] = mapped_column(Enum("TEXT", "OPTION", "BOOLEAN", "DATE", "MULTILINE_TEXT", "AMOUNT", "NUMBER", name="CustomFieldType"), default="TEXT", server_default=text("'TEXT'"))
    options: Mapped[list[str]] = mapped_column(ARRAY(Text), nullable=True)
    format: Mapped[str] = mapped_column(Enum("PLAIN", "LINK", name="CustomFieldFormat"), default="PLAIN", server_default=text("'PLAIN'"))
    """How a TEXT field's values are read. LINK requires a web link in every
    value; any other type stays PLAIN (enforced by a CHECK constraint).
    """
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="CustomField_organizationId_fkey"))
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="RESTRICT", onupdate="CASCADE", name="CustomField_userId_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)
    deleted_at: Mapped[Optional[datetime]] = mapped_column("deletedAt")

    __table_args__ = (
        Index("CustomField_organizationId_idx", "organizationId"),
        Index("CustomField_userId_idx", "userId"),
        Index("CustomField_organizationId_deletedAt_idx", "organizationId", "deletedAt"),
    )

    organization: Mapped["Organization"] = relationship(back_populates="custom_fields")
    created_by: Mapped["User"] = relationship(back_populates="custom_fields")
    asset_custom_fields_values: Mapped[list["AssetCustomFieldValue"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="custom_field")
    categories: Mapped[list["Category"]] = relationship(secondary=category_to_custom_field, back_populates="custom_fields")

class AssetCustomFieldValue(Base):
    __tablename__ = "AssetCustomFieldValue"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    value: Mapped[dict[str, Any]] = mapped_column(JSONB)
    asset_id: Mapped[str] = mapped_column("assetId", ForeignKey("Asset.id", ondelete="CASCADE", onupdate="CASCADE", name="AssetCustomFieldValue_assetId_fkey"))
    custom_field_id: Mapped[str] = mapped_column("customFieldId", ForeignKey("CustomField.id", ondelete="CASCADE", onupdate="CASCADE", name="AssetCustomFieldValue_customFieldId_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("AssetCustomFieldValue_lookup_idx", "assetId", "customFieldId"),
        Index("AssetCustomFieldValue_customFieldId_idx", "customFieldId"),
    )

    asset: Mapped["Asset"] = relationship(back_populates="custom_fields")
    custom_field: Mapped["CustomField"] = relationship(back_populates="asset_custom_fields_values")

class AssetModelCustomFieldDefault(Base):
    """A custom field value an AssetModel gives the assets that join it.

    `value` has the exact shape of `AssetCustomFieldValue.value`, so filling an
    asset is a plain copy. Rows cascade away with their model or field.
    """

    __tablename__ = "AssetModelCustomFieldDefault"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    value: Mapped[dict[str, Any]] = mapped_column(JSONB)
    asset_model_id: Mapped[str] = mapped_column("assetModelId", ForeignKey("AssetModel.id", ondelete="CASCADE", onupdate="CASCADE", name="AssetModelCustomFieldDefault_assetModelId_fkey"))
    custom_field_id: Mapped[str] = mapped_column("customFieldId")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("AssetModelCustomFieldDefault_customFieldId_idx", "customFieldId"),
        Index("AssetModelCustomFieldDefault_assetModelId_customFieldId_key", "assetModelId", "customFieldId", unique=True),
    )

    asset_model: Mapped["AssetModel"] = relationship(back_populates="custom_field_defaults")

class Invite(Base):
    __tablename__ = "Invite"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    inviter_id: Mapped[str] = mapped_column("inviterId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="Invite_inviterId_fkey"))
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="Invite_organizationId_fkey"))
    invitee_user_id: Mapped[Optional[str]] = mapped_column("inviteeUserId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="Invite_inviteeUserId_fkey"))
    team_member_id: Mapped[str] = mapped_column("teamMemberId", ForeignKey("TeamMember.id", ondelete="RESTRICT", onupdate="CASCADE", name="Invite_teamMemberId_fkey"))
    invitee_email: Mapped[str] = mapped_column("inviteeEmail")
    status: Mapped[str] = mapped_column(Enum("PENDING", "ACCEPTED", "REJECTED", "INVALIDATED", name="InviteStatuses"), default="PENDING", server_default=text("'PENDING'"))
    invite_code: Mapped[str] = mapped_column("inviteCode")
    roles: Mapped[list[str]] = mapped_column(ARRAY(Enum("ADMIN", "BASE", "OWNER", "SELF_SERVICE", "MANAGER", name="OrganizationRoles")), nullable=True)
    invite_message: Mapped[Optional[str]] = mapped_column("inviteMessage", String(1000))
    expires_at: Mapped[datetime] = mapped_column("expiresAt")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("Invite_inviteeUserId_idx", "inviteeUserId"),
        Index("Invite_inviterId_idx", "inviterId"),
        Index("Invite_organizationId_idx", "organizationId"),
        Index("Invite_teamMemberId_idx", "teamMemberId"),
    )

    inviter: Mapped["User"] = relationship(foreign_keys=[inviter_id], back_populates="sent_invites")
    organization: Mapped["Organization"] = relationship(back_populates="invites")
    invitee_user: Mapped[Optional["User"]] = relationship(foreign_keys=[invitee_user_id], back_populates="received_invites")
    invitee_team_member: Mapped["TeamMember"] = relationship(back_populates="received_invites")

class Announcement(Base):
    __tablename__ = "Announcement"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    name: Mapped[str]
    content: Mapped[str]
    link: Mapped[Optional[str]]
    link_text: Mapped[Optional[str]] = mapped_column("linkText")
    published: Mapped[bool] = mapped_column(default=False, server_default=text("false"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

class ConsumptionLog(Base):
    """*
    * Audit trail for all quantity changes on quantity-tracked assets.
    * Every event records who performed it, who received/returned items,
    * which booking it relates to, and the category of change.
    * Quantity is always positive — direction is determined by category.
    """

    __tablename__ = "ConsumptionLog"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    asset_id: Mapped[str] = mapped_column("assetId", ForeignKey("Asset.id", ondelete="CASCADE", onupdate="CASCADE", name="ConsumptionLog_assetId_fkey"))
    category: Mapped[str] = mapped_column(Enum("CHECKOUT", "RETURN", "RESTOCK", "ADJUSTMENT", "LOSS", "CONSUME", "DAMAGE", "MOVE", "INITIAL", "OPENING_BALANCE", name="ConsumptionCategory"))
    quantity: Mapped[int]
    note: Mapped[Optional[str]]
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="RESTRICT", onupdate="CASCADE", name="ConsumptionLog_userId_fkey"))
    booking_id: Mapped[Optional[str]] = mapped_column("bookingId", ForeignKey("Booking.id", ondelete="SET NULL", onupdate="CASCADE", name="ConsumptionLog_bookingId_fkey"))
    custodian_id: Mapped[Optional[str]] = mapped_column("custodianId", ForeignKey("TeamMember.id", ondelete="SET NULL", onupdate="CASCADE", name="ConsumptionLog_custodianId_fkey"))
    booking_asset_id: Mapped[Optional[str]] = mapped_column("bookingAssetId")
    location_id: Mapped[Optional[str]] = mapped_column("locationId", ForeignKey("Location.id", ondelete="SET NULL", onupdate="CASCADE", name="ConsumptionLog_locationId_fkey"))
    stock_change: Mapped[Optional[int]] = mapped_column("stockChange")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("ConsumptionLog_assetId_createdAt_idx", "assetId", "createdAt"),
        Index("ConsumptionLog_userId_idx", "userId"),
        Index("ConsumptionLog_bookingId_idx", "bookingId"),
        Index("ConsumptionLog_custodianId_idx", "custodianId"),
        Index("ConsumptionLog_bookingAssetId_idx", "bookingAssetId"),
        Index("ConsumptionLog_locationId_idx", "locationId"),
    )

    asset: Mapped["Asset"] = relationship(back_populates="consumption_logs")
    performed_by: Mapped["User"] = relationship(back_populates="consumption_logs")
    booking: Mapped[Optional["Booking"]] = relationship(back_populates="consumption_logs")
    custodian: Mapped[Optional["TeamMember"]] = relationship(back_populates="consumption_logs")
    location: Mapped[Optional["Location"]] = relationship(back_populates="consumption_logs")

class BookingAsset(Base):
    """*
    * Explicit pivot table between Booking and Asset.
    * Replaces the implicit M2M (_AssetToBooking) via a rename migration in Phase 3.
    """

    __tablename__ = "BookingAsset"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    booking_id: Mapped[str] = mapped_column("bookingId", ForeignKey("Booking.id", ondelete="CASCADE", onupdate="CASCADE", name="BookingAsset_bookingId_fkey"))
    asset_id: Mapped[str] = mapped_column("assetId", ForeignKey("Asset.id", ondelete="CASCADE", onupdate="CASCADE", name="BookingAsset_assetId_fkey"))
    quantity: Mapped[int] = mapped_column(default=1, server_default=text("1"))
    booking_model_request_id: Mapped[Optional[str]] = mapped_column("bookingModelRequestId")
    asset_kit_id: Mapped[Optional[str]] = mapped_column("assetKitId")
    source_kit_id: Mapped[Optional[str]] = mapped_column("sourceKitId")
    checked_out_at: Mapped[Optional[datetime]] = mapped_column("checkedOutAt", UtcDateTimeTz(precision=3))
    checked_out_by_id: Mapped[Optional[str]] = mapped_column("checkedOutById")
    checked_in_at: Mapped[Optional[datetime]] = mapped_column("checkedInAt", UtcDateTimeTz(precision=3))
    checked_in_by_id: Mapped[Optional[str]] = mapped_column("checkedInById")
    checked_out_quantity: Mapped[int] = mapped_column("checkedOutQuantity", default=0, server_default=text("0"))
    source_location_id: Mapped[Optional[str]] = mapped_column("sourceLocationId")

    __table_args__ = (
        Index("BookingAsset_assetId_idx", "assetId"),
        Index("BookingAsset_assetKitId_idx", "assetKitId"),
        Index("BookingAsset_sourceKitId_idx", "sourceKitId"),
        Index("BookingAsset_sourceLocationId_idx", "sourceLocationId"),
    )

    booking: Mapped["Booking"] = relationship(back_populates="booking_assets")
    asset: Mapped["Asset"] = relationship(back_populates="booking_assets")

class Booking(Base):
    __tablename__ = "Booking"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    name: Mapped[str]
    status: Mapped[str] = mapped_column(Enum("DRAFT", "RESERVED", "ONGOING", "OVERDUE", "COMPLETE", "ARCHIVED", "CANCELLED", name="BookingStatus"), default="DRAFT", server_default=text("'DRAFT'"))
    description: Mapped[Optional[str]] = mapped_column(default="", server_default=text("''"))
    active_scheduler_reference: Mapped[Optional[str]] = mapped_column("activeSchedulerReference")
    creator_id: Mapped[str] = mapped_column("creatorId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="Booking_creatorId_fkey"))
    custodian_user_id: Mapped[Optional[str]] = mapped_column("custodianUserId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="Booking_custodianUserId_fkey"))
    custodian_team_member_id: Mapped[Optional[str]] = mapped_column("custodianTeamMemberId", ForeignKey("TeamMember.id", ondelete="CASCADE", onupdate="CASCADE", name="Booking_custodianTeamMemberId_fkey"))
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="Booking_organizationId_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", UtcDateTimeTz(precision=3), default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", UtcDateTimeTz(precision=3), default=utc_now, onupdate=utc_now)
    from_: Mapped[datetime] = mapped_column("from", UtcDateTimeTz(precision=3))
    to: Mapped[datetime] = mapped_column(UtcDateTimeTz(precision=3))
    original_from: Mapped[Optional[datetime]] = mapped_column("originalFrom", UtcDateTimeTz(precision=3))
    original_to: Mapped[Optional[datetime]] = mapped_column("originalTo", UtcDateTimeTz(precision=3))
    auto_archived_at: Mapped[Optional[datetime]] = mapped_column("autoArchivedAt", UtcDateTimeTz(precision=3))
    archived_without_checkin: Mapped[bool] = mapped_column("archivedWithoutCheckin", default=False, server_default=text("false"))
    cancellation_reason: Mapped[Optional[str]] = mapped_column("cancellationReason")

    __table_args__ = (
        Index("Booking_creatorId_idx", "creatorId"),
        Index("Booking_custodianTeamMemberId_idx", "custodianTeamMemberId"),
        Index("Booking_custodianUserId_idx", "custodianUserId"),
        Index("Booking_organizationId_idx", "organizationId"),
        Index("Booking_status_idx", "status"),
        Index("Booking_organizationId_from_id_idx", "organizationId", "from", "id"),
    )

    creator: Mapped["User"] = relationship(foreign_keys=[creator_id], back_populates="bookings")
    custodian_user: Mapped[Optional["User"]] = relationship(foreign_keys=[custodian_user_id], back_populates="custodies")
    custodian_team_member: Mapped[Optional["TeamMember"]] = relationship(back_populates="bookings")
    organization: Mapped["Organization"] = relationship(back_populates="bookings")
    booking_assets: Mapped[list["BookingAsset"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="booking")
    consumption_logs: Mapped[list["ConsumptionLog"]] = relationship(back_populates="booking")
    notes: Mapped[list["BookingNote"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="booking")
    model_requests: Mapped[list["BookingModelRequest"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="booking")
    partial_checkins: Mapped[list["PartialBookingCheckin"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="booking")
    partial_checkouts: Mapped[list["PartialBookingCheckout"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="booking")
    tags: Mapped[list["Tag"]] = relationship(secondary=booking_to_tag, back_populates="bookings")
    notification_recipients: Mapped[list["TeamMember"]] = relationship(secondary=booking_notification_recipients, back_populates="booking_notifications")

class BookingSettings(Base):
    __tablename__ = "BookingSettings"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    buffer_start_time: Mapped[int] = mapped_column("bufferStartTime", default=0, server_default=text("0"))
    tags_required: Mapped[bool] = mapped_column("tagsRequired", default=False, server_default=text("false"))
    max_booking_length: Mapped[Optional[int]] = mapped_column("maxBookingLength")
    max_booking_length_skip_closed_days: Mapped[bool] = mapped_column("maxBookingLengthSkipClosedDays", default=False, server_default=text("false"))
    auto_archive_bookings: Mapped[bool] = mapped_column("autoArchiveBookings", default=False, server_default=text("false"))
    auto_archive_days: Mapped[int] = mapped_column("autoArchiveDays", default=2, server_default=text("2"))
    auto_archive_expired_reservations: Mapped[bool] = mapped_column("autoArchiveExpiredReservations", default=False, server_default=text("false"))
    require_explicit_checkin_for_admin: Mapped[bool] = mapped_column("requireExplicitCheckinForAdmin", default=False, server_default=text("false"))
    require_explicit_checkin_for_self_service: Mapped[bool] = mapped_column("requireExplicitCheckinForSelfService", default=False, server_default=text("false"))
    require_explicit_checkout_for_admin: Mapped[bool] = mapped_column("requireExplicitCheckoutForAdmin", default=False, server_default=text("false"))
    require_explicit_checkout_for_self_service: Mapped[bool] = mapped_column("requireExplicitCheckoutForSelfService", default=False, server_default=text("false"))
    count_kits_as_single_unit: Mapped[bool] = mapped_column("countKitsAsSingleUnit", default=False, server_default=text("false"))
    notify_booking_creator: Mapped[bool] = mapped_column("notifyBookingCreator", default=True, server_default=text("true"))
    notify_admins_on_new_booking: Mapped[bool] = mapped_column("notifyAdminsOnNewBooking", default=True, server_default=text("true"))
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="BookingSettings_organizationId_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("BookingSettings_organizationId_key", "organizationId", unique=True),
        Index("BookingSettings_organizationId_idx", "organizationId"),
    )

    organization: Mapped["Organization"] = relationship(back_populates="booking_settings")
    always_notify_team_members: Mapped[list["TeamMember"]] = relationship(secondary=booking_settings_always_notify, back_populates="always_notify_booking_settings")

class PartialBookingCheckin(Base):
    """A progressive check-in SESSION, mirroring PartialBookingCheckout.

    `assetIds` lists only FULLY reconciled assets: INDIVIDUAL always, and
    QUANTITY_TRACKED only once its remaining units hit zero (partial returns
    live in `ConsumptionLog`). Completing a booking outright writes no row here,
    so the same warning applies — absence is not evidence of anything.
    `BookingAsset.checkedInAt` is the per-slice fact.
    """

    __tablename__ = "PartialBookingCheckin"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    asset_ids: Mapped[list[str]] = mapped_column("assetIds", ARRAY(Text), nullable=True)
    checkin_count: Mapped[int] = mapped_column("checkinCount")
    checkin_timestamp: Mapped[datetime] = mapped_column("checkinTimestamp", UtcDateTimeTz(precision=3), default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    booking_id: Mapped[str] = mapped_column("bookingId", ForeignKey("Booking.id", ondelete="CASCADE", onupdate="CASCADE", name="PartialBookingCheckin_bookingId_fkey"))
    checked_in_by_id: Mapped[str] = mapped_column("checkedInById", ForeignKey("User.id", ondelete="RESTRICT", onupdate="CASCADE", name="PartialBookingCheckin_checkedInById_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("PartialBookingCheckin_bookingId_idx", "bookingId"),
        Index("PartialBookingCheckin_checkedInById_idx", "checkedInById"),
        Index("PartialBookingCheckin_checkinTimestamp_idx", "checkinTimestamp"),
        Index("PartialBookingCheckin_bookingId_checkinTimestamp_idx", "bookingId", "checkinTimestamp"),
    )

    booking: Mapped["Booking"] = relationship(back_populates="partial_checkins")
    checked_in_by: Mapped["User"] = relationship(back_populates="partial_checkins")

class PartialBookingCheckout(Base):
    """A progressive checkout SESSION: one scan batch, covering N assets via the
    positional `assetIds` / `quantities` / `bookingAssetIds` arrays.

    Records sessions and the units they claimed — NOT whether an asset is out.
    The all-at-once checkout (`checkoutBooking`, the Check out button) flips
    statuses and writes NO row here, so **absence proves nothing** and must never
    be read as "was not checked out". Reading it that way is what let a single
    later scan mark every button-checked-out asset on a booking as
    "never checked out". `BookingAsset.checkedOutAt` answers that question.
    """

    __tablename__ = "PartialBookingCheckout"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    asset_ids: Mapped[list[str]] = mapped_column("assetIds", ARRAY(Text), nullable=True)
    quantities: Mapped[list[int]] = mapped_column(ARRAY(Integer), nullable=True, default=lambda: [], server_default=text("ARRAY[]::integer[]"))
    booking_asset_ids: Mapped[list[str]] = mapped_column("bookingAssetIds", ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))
    checkout_count: Mapped[int] = mapped_column("checkoutCount")
    checkout_timestamp: Mapped[datetime] = mapped_column("checkoutTimestamp", UtcDateTimeTz(precision=3), default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    booking_id: Mapped[str] = mapped_column("bookingId", ForeignKey("Booking.id", ondelete="CASCADE", onupdate="CASCADE", name="PartialBookingCheckout_bookingId_fkey"))
    checked_out_by_id: Mapped[str] = mapped_column("checkedOutById", ForeignKey("User.id", ondelete="RESTRICT", onupdate="CASCADE", name="PartialBookingCheckout_checkedOutById_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("PartialBookingCheckout_bookingId_idx", "bookingId"),
        Index("PartialBookingCheckout_checkedOutById_idx", "checkedOutById"),
        Index("PartialBookingCheckout_checkoutTimestamp_idx", "checkoutTimestamp"),
        Index("PartialBookingCheckout_bookingId_checkoutTimestamp_idx", "bookingId", "checkoutTimestamp"),
    )

    booking: Mapped["Booking"] = relationship(back_populates="partial_checkouts")
    checked_out_by: Mapped["User"] = relationship(back_populates="partial_checkouts")

class MobileAuthCode(Base):
    """Native-app SSO web-delegated login — single-use session-handoff codes.
    @see apps/webapp/app/routes/_auth+/oauth.callback.mobile.tsx
    """

    __tablename__ = "MobileAuthCode"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="MobileAuthCode_userId_fkey"))
    code_hash: Mapped[str] = mapped_column("codeHash")
    code_challenge: Mapped[Optional[str]] = mapped_column("codeChallenge")
    session_ciphertext: Mapped[Optional[str]] = mapped_column("sessionCiphertext")
    consumed_at: Mapped[Optional[datetime]] = mapped_column("consumedAt")
    expires_at: Mapped[datetime] = mapped_column("expiresAt")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("MobileAuthCode_codeHash_key", "codeHash", unique=True),
        Index("MobileAuthCode_userId_idx", "userId"),
        Index("MobileAuthCode_expiresAt_idx", "expiresAt"),
    )

    user: Mapped["User"] = relationship(back_populates="mobile_auth_codes")

class PushToken(Base):
    """One Expo push token per installed companion app. A token identifies a
    device, not a person: when someone else signs in on the same phone, the
    row moves to them. Pushes go only to rows where `disabledAt` is null.
    @see apps/webapp/app/modules/push/service.server.ts
    """

    __tablename__ = "PushToken"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    token: Mapped[str]
    platform: Mapped[str] = mapped_column(Enum("IOS", "ANDROID", name="DevicePlatform"))
    device_name: Mapped[Optional[str]] = mapped_column("deviceName")
    app_version: Mapped[Optional[str]] = mapped_column("appVersion")
    last_seen_at: Mapped[datetime] = mapped_column("lastSeenAt")
    disabled_at: Mapped[Optional[datetime]] = mapped_column("disabledAt")
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="PushToken_userId_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("PushToken_token_key", "token", unique=True),
        Index("PushToken_userId_idx", "userId"),
    )

    user: Mapped["User"] = relationship(back_populates="push_tokens")

class NotificationPreference(Base):
    """A person's push notification switches. A missing row means every switch
    is on, so no backfill is needed. The three category columns are named
    after `PushCategory` in apps/webapp/app/modules/push/types.ts.
    """

    __tablename__ = "NotificationPreference"

    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="NotificationPreference_userId_fkey"), primary_key=True)
    push_enabled: Mapped[bool] = mapped_column("pushEnabled", default=True, server_default=text("true"))
    bookings: Mapped[bool] = mapped_column(default=True, server_default=text("true"))
    audits: Mapped[bool] = mapped_column(default=True, server_default=text("true"))
    stock_and_reminders: Mapped[bool] = mapped_column("stockAndReminders", default=True, server_default=text("true"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    user: Mapped["User"] = relationship(back_populates="notification_preference")

class Kit(Base):
    __tablename__ = "Kit"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    name: Mapped[str]
    description: Mapped[Optional[str]]
    status: Mapped[str] = mapped_column(Enum("AVAILABLE", "IN_CUSTODY", "CHECKED_OUT", name="KitStatus"), default="AVAILABLE", server_default=text("'AVAILABLE'"))
    image: Mapped[Optional[str]]
    image_expiration: Mapped[Optional[datetime]] = mapped_column("imageExpiration")
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="Kit_organizationId_fkey"))
    created_by_id: Mapped[str] = mapped_column("createdById", ForeignKey("User.id", ondelete="RESTRICT", onupdate="CASCADE", name="Kit_createdById_fkey"))
    category_id: Mapped[Optional[str]] = mapped_column("categoryId", ForeignKey("Category.id", ondelete="SET NULL", onupdate="CASCADE", name="Kit_categoryId_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)
    location_id: Mapped[Optional[str]] = mapped_column("locationId", ForeignKey("Location.id", ondelete="SET NULL", onupdate="CASCADE", name="Kit_locationId_fkey"))

    __table_args__ = (
        Index("Kit_createdById_idx", "createdById"),
        Index("Kit_organizationId_idx", "organizationId"),
        Index("Kit_categoryId_organizationId_idx", "categoryId", "organizationId"),
        Index("Kit_categoryId_organizationId_createdAt_idx", "categoryId", "organizationId", "createdAt"),
        Index("Kit_categoryId_organizationId_name_idx", "categoryId", "organizationId", "name"),
        Index("Kit_categoryId_organizationId_status_idx", "categoryId", "organizationId", "status"),
    )

    organization: Mapped["Organization"] = relationship(back_populates="kits")
    created_by: Mapped["User"] = relationship(back_populates="created_kits")
    category: Mapped[Optional["Category"]] = relationship(back_populates="kits")
    location: Mapped[Optional["Location"]] = relationship(back_populates="kits")
    asset_kits: Mapped[list["AssetKit"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="kit")
    qr_codes: Mapped[list["Qr"]] = relationship(back_populates="kit")
    barcodes: Mapped[list["Barcode"]] = relationship(back_populates="kit")
    reports: Mapped[list["ReportFound"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="kit")
    scans: Mapped[list["Scan"]] = relationship(back_populates="kit")
    custody: Mapped[Optional["KitCustody"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="kit")

class KitCustody(Base):
    __tablename__ = "KitCustody"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    custodian_id: Mapped[str] = mapped_column("custodianId", ForeignKey("TeamMember.id", ondelete="RESTRICT", onupdate="CASCADE", name="KitCustody_custodianId_fkey"))
    kit_id: Mapped[str] = mapped_column("kitId", ForeignKey("Kit.id", ondelete="CASCADE", onupdate="CASCADE", name="KitCustody_kitId_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("KitCustody_kitId_key", "kitId", unique=True),
        Index("KitCustody_custodianId_idx", "custodianId"),
    )

    custodian: Mapped["TeamMember"] = relationship(back_populates="kit_custodies")
    kit: Mapped["Kit"] = relationship(back_populates="custody")
    inherited_custody: Mapped[list["Custody"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="kit_custody")

class AssetKit(Base):
    """Pivot linking an `Asset` to the `Kit` it belongs to.

    `@@unique([assetId])` enforces "at most one kit per asset" today.
    The structure already allows multi-row-per-asset; if multi-kit
    allocation lands later, drop the `assetId` unique constraint and
    add a type-aware single-row trigger + sum-within-total CONSTRAINT
    TRIGGER (parallel to the Custody / BookingAsset pivots).
    """

    __tablename__ = "AssetKit"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    asset_id: Mapped[str] = mapped_column("assetId", ForeignKey("Asset.id", ondelete="CASCADE", onupdate="CASCADE", name="AssetKit_assetId_fkey"))
    kit_id: Mapped[str] = mapped_column("kitId", ForeignKey("Kit.id", ondelete="CASCADE", onupdate="CASCADE", name="AssetKit_kitId_fkey"))
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="AssetKit_organizationId_fkey"))
    quantity: Mapped[int] = mapped_column(default=1, server_default=text("1"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("AssetKit_kitId_idx", "kitId"),
        Index("AssetKit_organizationId_idx", "organizationId"),
        Index("AssetKit_assetId_kitId_key", "assetId", "kitId", unique=True),
    )

    asset: Mapped["Asset"] = relationship(back_populates="asset_kits")
    kit: Mapped["Kit"] = relationship(back_populates="asset_kits")
    organization: Mapped["Organization"] = relationship(back_populates="asset_kits")
    asset_locations: Mapped[list["AssetLocation"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="asset_kit")

class AssetLocation(Base):
    __tablename__ = "AssetLocation"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    asset_id: Mapped[str] = mapped_column("assetId", ForeignKey("Asset.id", ondelete="CASCADE", onupdate="CASCADE", name="AssetLocation_assetId_fkey"))
    location_id: Mapped[str] = mapped_column("locationId", ForeignKey("Location.id", ondelete="CASCADE", onupdate="CASCADE", name="AssetLocation_locationId_fkey"))
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="AssetLocation_organizationId_fkey"))
    quantity: Mapped[int] = mapped_column(default=1, server_default=text("1"))
    asset_kit_id: Mapped[Optional[str]] = mapped_column("assetKitId", ForeignKey("AssetKit.id", ondelete="CASCADE", onupdate="CASCADE", name="AssetLocation_assetKitId_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("AssetLocation_locationId_idx", "locationId"),
        Index("AssetLocation_organizationId_idx", "organizationId"),
        Index("AssetLocation_assetId_idx", "assetId"),
        Index("AssetLocation_assetKitId_idx", "assetKitId"),
    )

    asset: Mapped["Asset"] = relationship(back_populates="asset_locations")
    location: Mapped["Location"] = relationship(back_populates="asset_locations")
    organization: Mapped["Organization"] = relationship(back_populates="asset_locations")
    asset_kit: Mapped[Optional["AssetKit"]] = relationship(back_populates="asset_locations")

class AssetReminder(Base):
    __tablename__ = "AssetReminder"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    name: Mapped[str]
    message: Mapped[str]
    alert_date_time: Mapped[datetime] = mapped_column("alertDateTime")
    active_scheduler_reference: Mapped[Optional[str]] = mapped_column("activeSchedulerReference")
    recurrence_unit: Mapped[Optional[str]] = mapped_column("recurrenceUnit", Enum("DAY", "WEEK", "MONTH", "YEAR", name="ReminderRecurrenceUnit"))
    recurrence_interval: Mapped[Optional[int]] = mapped_column("recurrenceInterval")
    recurrence_timezone: Mapped[Optional[str]] = mapped_column("recurrenceTimezone")
    recurrence_ends_at: Mapped[Optional[datetime]] = mapped_column("recurrenceEndsAt")
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="RESTRICT", onupdate="CASCADE", name="AssetReminder_organizationId_fkey"))
    asset_id: Mapped[str] = mapped_column("assetId", ForeignKey("Asset.id", ondelete="CASCADE", onupdate="CASCADE", name="AssetReminder_assetId_fkey"))
    created_by_id: Mapped[str] = mapped_column("createdById", ForeignKey("User.id", ondelete="RESTRICT", onupdate="CASCADE", name="AssetReminder_createdById_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("AssetReminder_assetId_alertDateTime_idx", "assetId", "alertDateTime"),
        Index("AssetReminder_name_message_idx", "name", "message"),
        Index("AssetReminder_organizationId_alertDateTime_assetId_idx", "organizationId", "alertDateTime", "assetId"),
        Index("AssetReminder_alertDateTime_activeSchedulerReference_idx", "alertDateTime", "activeSchedulerReference"),
        Index("AssetReminder_createdById_idx", "createdById"),
    )

    organization: Mapped["Organization"] = relationship(back_populates="asset_reminders")
    asset: Mapped["Asset"] = relationship(back_populates="reminders")
    created_by: Mapped["User"] = relationship(back_populates="asset_reminders")
    team_members: Mapped[list["TeamMember"]] = relationship(secondary=asset_reminder_to_team_member, back_populates="asset_reminders")

class WorkingHours(Base):
    __tablename__ = "WorkingHours"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    enabled: Mapped[bool] = mapped_column(default=False, server_default=text("false"))
    weekly_schedule: Mapped[dict[str, Any]] = mapped_column("weeklySchedule", JSONB, default=lambda: {"0": {"isOpen": False}, "1": {"isOpen": True, "openTime": "09:00", "closeTime": "17:00"}, "2": {"isOpen": True, "openTime": "09:00", "closeTime": "17:00"}, "3": {"isOpen": True, "openTime": "09:00", "closeTime": "17:00"}, "4": {"isOpen": True, "openTime": "09:00", "closeTime": "17:00"}, "5": {"isOpen": True, "openTime": "09:00", "closeTime": "17:00"}, "6": {"isOpen": False}}, server_default=text("'{\"0\":{\"isOpen\":false},\"1\":{\"isOpen\":true,\"openTime\":\"09:00\",\"closeTime\":\"17:00\"},\"2\":{\"isOpen\":true,\"openTime\":\"09:00\",\"closeTime\":\"17:00\"},\"3\":{\"isOpen\":true,\"openTime\":\"09:00\",\"closeTime\":\"17:00\"},\"4\":{\"isOpen\":true,\"openTime\":\"09:00\",\"closeTime\":\"17:00\"},\"5\":{\"isOpen\":true,\"openTime\":\"09:00\",\"closeTime\":\"17:00\"},\"6\":{\"isOpen\":false}}'"))
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="WorkingHours_organizationId_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("WorkingHours_organizationId_key", "organizationId", unique=True),
        Index("WorkingHours_organizationId_idx", "organizationId"),
    )

    organization: Mapped["Organization"] = relationship(back_populates="working_hours")
    overrides: Mapped[list["WorkingHoursOverride"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="working_hours")

class WorkingHoursOverride(Base):
    __tablename__ = "WorkingHoursOverride"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    date_: Mapped[date] = mapped_column("date", Date)
    is_open: Mapped[bool] = mapped_column("isOpen", default=False, server_default=text("false"))
    open_time: Mapped[Optional[str]] = mapped_column("openTime")
    close_time: Mapped[Optional[str]] = mapped_column("closeTime")
    reason: Mapped[Optional[str]]
    working_hours_id: Mapped[str] = mapped_column("workingHoursId", ForeignKey("WorkingHours.id", ondelete="CASCADE", onupdate="CASCADE", name="WorkingHoursOverride_workingHoursId_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("WorkingHoursOverride_workingHoursId_date_idx", "workingHoursId", "date"),
        Index("WorkingHoursOverride_date_isOpen_idx", "date", "isOpen"),
        Index("WorkingHoursOverride_workingHoursId_date_key", "workingHoursId", "date", unique=True),
    )

    working_hours: Mapped["WorkingHours"] = relationship(back_populates="overrides")

class Update(Base):
    __tablename__ = "Update"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    title: Mapped[str]
    content: Mapped[str]
    url: Mapped[Optional[str]]
    image_url: Mapped[Optional[str]] = mapped_column("imageUrl")
    publish_date: Mapped[datetime] = mapped_column("publishDate")
    status: Mapped[str] = mapped_column(Enum("DRAFT", "PUBLISHED", name="UpdateStatus"), default="DRAFT", server_default=text("'DRAFT'"))
    target_roles: Mapped[list[str]] = mapped_column("targetRoles", ARRAY(Enum("ADMIN", "BASE", "OWNER", "SELF_SERVICE", "MANAGER", name="OrganizationRoles")), nullable=True)
    click_count: Mapped[int] = mapped_column("clickCount", default=0, server_default=text("0"))
    view_count: Mapped[int] = mapped_column("viewCount", default=0, server_default=text("0"))
    created_by_id: Mapped[str] = mapped_column("createdById", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="Update_createdById_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("Update_status_publishDate_idx", "status", "publishDate"),
        Index("Update_publishDate_idx", "publishDate"),
        Index("Update_createdById_idx", "createdById"),
    )

    created_by: Mapped["User"] = relationship(back_populates="created_updates")
    user_reads: Mapped[list["UserUpdateRead"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="update")

class UserUpdateRead(Base):
    __tablename__ = "UserUpdateRead"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="UserUpdateRead_userId_fkey"))
    update_id: Mapped[str] = mapped_column("updateId", ForeignKey("Update.id", ondelete="CASCADE", onupdate="CASCADE", name="UserUpdateRead_updateId_fkey"))
    read_at: Mapped[datetime] = mapped_column("readAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("UserUpdateRead_userId_idx", "userId"),
        Index("UserUpdateRead_updateId_idx", "updateId"),
        Index("UserUpdateRead_readAt_idx", "readAt"),
        Index("UserUpdateRead_userId_updateId_key", "userId", "updateId", unique=True),
    )

    user: Mapped["User"] = relationship(back_populates="update_reads")
    update: Mapped["Update"] = relationship(back_populates="user_reads")

class AuditNote(Base):
    __tablename__ = "AuditNote"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    content: Mapped[str]
    type: Mapped[str] = mapped_column(Enum("COMMENT", "UPDATE", name="NoteType"), default="COMMENT", server_default=text("'COMMENT'"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)
    user_id: Mapped[Optional[str]] = mapped_column("userId", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="AuditNote_userId_fkey"))
    audit_session_id: Mapped[str] = mapped_column("auditSessionId", ForeignKey("AuditSession.id", ondelete="CASCADE", onupdate="CASCADE", name="AuditNote_auditSessionId_fkey"))
    audit_asset_id: Mapped[Optional[str]] = mapped_column("auditAssetId", ForeignKey("AuditAsset.id", ondelete="SET NULL", onupdate="CASCADE", name="AuditNote_auditAssetId_fkey"))

    __table_args__ = (
        Index("AuditNote_auditSessionId_idx", "auditSessionId"),
        Index("AuditNote_userId_idx", "userId"),
        Index("AuditNote_auditAssetId_idx", "auditAssetId"),
    )

    user: Mapped[Optional["User"]] = relationship(back_populates="audit_notes")
    audit_session: Mapped["AuditSession"] = relationship(back_populates="notes")
    audit_asset: Mapped[Optional["AuditAsset"]] = relationship(back_populates="notes")

class TeamMemberNote(Base):
    __tablename__ = "TeamMemberNote"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    content: Mapped[str]
    type: Mapped[str] = mapped_column(Enum("COMMENT", "UPDATE", name="NoteType"), default="COMMENT", server_default=text("'COMMENT'"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)
    user_id: Mapped[Optional[str]] = mapped_column("userId", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="TeamMemberNote_userId_fkey"))
    team_member_id: Mapped[str] = mapped_column("teamMemberId", ForeignKey("TeamMember.id", ondelete="CASCADE", onupdate="CASCADE", name="TeamMemberNote_teamMemberId_fkey"))
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="TeamMemberNote_organizationId_fkey"))

    __table_args__ = (
        Index("TeamMemberNote_teamMemberId_idx", "teamMemberId"),
        Index("TeamMemberNote_userId_idx", "userId"),
        Index("TeamMemberNote_organizationId_idx", "organizationId"),
        Index("TeamMemberNote_teamMemberId_organizationId_idx", "teamMemberId", "organizationId"),
    )

    user: Mapped[Optional["User"]] = relationship(back_populates="team_member_notes_authored")
    team_member: Mapped["TeamMember"] = relationship(back_populates="team_member_notes")
    organization: Mapped["Organization"] = relationship(back_populates="team_member_notes")

class AuditSession(Base):
    __tablename__ = "AuditSession"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    name: Mapped[str]
    description: Mapped[Optional[str]]
    target_id: Mapped[Optional[str]] = mapped_column("targetId")
    status: Mapped[str] = mapped_column(Enum("PENDING", "ACTIVE", "COMPLETED", "CANCELLED", "ARCHIVED", name="AuditStatus"), default="PENDING", server_default=text("'PENDING'"))
    scope_meta: Mapped[Optional[dict[str, Any]]] = mapped_column("scopeMeta", JSONB(none_as_null=True))
    expected_asset_count: Mapped[int] = mapped_column("expectedAssetCount", default=0, server_default=text("0"))
    found_asset_count: Mapped[int] = mapped_column("foundAssetCount", default=0, server_default=text("0"))
    missing_asset_count: Mapped[int] = mapped_column("missingAssetCount", default=0, server_default=text("0"))
    unexpected_asset_count: Mapped[int] = mapped_column("unexpectedAssetCount", default=0, server_default=text("0"))
    started_at: Mapped[Optional[datetime]] = mapped_column("startedAt")
    due_date: Mapped[Optional[datetime]] = mapped_column("dueDate")
    completed_at: Mapped[Optional[datetime]] = mapped_column("completedAt")
    cancelled_at: Mapped[Optional[datetime]] = mapped_column("cancelledAt")
    active_scheduler_reference: Mapped[Optional[str]] = mapped_column("activeSchedulerReference")
    created_by_id: Mapped[str] = mapped_column("createdById", ForeignKey("User.id", ondelete="RESTRICT", onupdate="CASCADE", name="AuditSession_createdById_fkey"))
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="AuditSession_organizationId_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("AuditSession_organizationId_status_idx", "organizationId", "status"),
        Index("AuditSession_createdById_idx", "createdById"),
        Index("AuditSession_status_createdAt_idx", "status", "createdAt"),
    )

    created_by: Mapped["User"] = relationship(back_populates="audits_created")
    organization: Mapped["Organization"] = relationship(back_populates="audit_sessions")
    assignments: Mapped[list["AuditAssignment"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="audit_session")
    assets: Mapped[list["AuditAsset"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="audit_session")
    scans: Mapped[list["AuditScan"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="audit_session")
    notes: Mapped[list["AuditNote"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="audit_session")
    images: Mapped[list["AuditImage"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="audit_session")

class AssetImport(Base):
    """One CSV content import. Holds the parsed rows from preview until the import
    completes, so the job imports exactly the file the user previewed.
    A partial unique index on (organizationId) WHERE status IN (QUEUED, RUNNING) allows one
    active import per workspace. Prisma cannot express it; it lives in migration
    20261015120100_asset_import_one_active_per_workspace.
    """

    __tablename__ = "AssetImport"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    status: Mapped[str] = mapped_column(Enum("DRAFT", "QUEUED", "RUNNING", "COMPLETED", "FAILED", name="AssetImportStatus"), default="DRAFT", server_default=text("'DRAFT'"))
    phase: Mapped[str] = mapped_column(Enum("PREPARE", "ROWS", "FINALIZE", name="AssetImportPhase"), default="PREPARE", server_default=text("'PREPARE'"))
    file_name: Mapped[Optional[str]] = mapped_column("fileName")
    total_rows: Mapped[int] = mapped_column("totalRows")
    processed_rows: Mapped[int] = mapped_column("processedRows", default=0, server_default=text("0"))
    """Rows committed so far. Advanced in the same transaction as each chunk."""
    rows: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))
    """Parsed rows (`StoredImportRow[]`). Cleared when the import completes."""
    resolved: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))
    """Name to id maps written by the PREPARE step (`ResolvedImportEntities`)."""
    can_use_barcodes: Mapped[bool] = mapped_column("canUseBarcodes", default=False, server_default=text("false"))
    error_message: Mapped[Optional[str]] = mapped_column("errorMessage")
    """User-facing reason the import failed. Null unless status is FAILED."""
    warnings: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))
    """Rows imported without a conflicting piece (`ImportWarning[]`), capped at 500."""
    warning_count: Mapped[int] = mapped_column("warningCount", default=0, server_default=text("0"))
    """Exact number of warnings, including any beyond the stored cap."""
    locked_until: Mapped[Optional[datetime]] = mapped_column("lockedUntil")
    """A step holds the import until this time. Null when no step is running."""
    confirmed_at: Mapped[Optional[datetime]] = mapped_column("confirmedAt")
    started_at: Mapped[Optional[datetime]] = mapped_column("startedAt")
    completed_at: Mapped[Optional[datetime]] = mapped_column("completedAt")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)
    created_by_id: Mapped[str] = mapped_column("createdById", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="AssetImport_createdById_fkey"))
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="AssetImport_organizationId_fkey"))

    __table_args__ = (
        Index("AssetImport_organizationId_status_idx", "organizationId", "status"),
        Index("AssetImport_createdById_organizationId_status_idx", "createdById", "organizationId", "status"),
    )

    created_by: Mapped["User"] = relationship(back_populates="asset_imports")
    organization: Mapped["Organization"] = relationship(back_populates="asset_imports")

class LinkPreview(Base):
    """What Shelf knows about one web link used in a workspace's custom fields,
    keyed by its `LinkKey` from @shelf/links. Two writers on disjoint columns:
    the link preview job (fetch columns) and people (customTitle).
    """

    __tablename__ = "LinkPreview"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    url: Mapped[str]
    """`LinkKey` from @shelf/links, never a raw field value. At most 2048
    characters (CHECK constraint in the migration).
    """
    page_title: Mapped[Optional[str]] = mapped_column("pageTitle")
    outcome: Mapped[Optional[str]] = mapped_column(Enum("TITLE", "NO_TITLE", "SIGN_IN", "BLOCKED", "UNREACHABLE", name="LinkPreviewOutcome"))
    """Null until the first fetch finishes."""
    attempts: Mapped[int] = mapped_column(default=0, server_default=text("0"))
    fetched_at: Mapped[Optional[datetime]] = mapped_column("fetchedAt")
    next_fetch_at: Mapped[Optional[datetime]] = mapped_column("nextFetchAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    """When the link may be fetched next; null means never again. A claim moves
    it a few minutes ahead, which is also the lease on a fetch in flight.
    """
    custom_title: Mapped[Optional[str]] = mapped_column("customTitle")
    """The name a person gave the link in this workspace; wins over pageTitle."""
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="LinkPreview_organizationId_fkey"))

    __table_args__ = (
        Index("LinkPreview_organizationId_url_key", "organizationId", "url", unique=True),
    )

    organization: Mapped["Organization"] = relationship(back_populates="link_previews")

class AuditAssignment(Base):
    __tablename__ = "AuditAssignment"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    audit_session_id: Mapped[str] = mapped_column("auditSessionId", ForeignKey("AuditSession.id", ondelete="CASCADE", onupdate="CASCADE", name="AuditAssignment_auditSessionId_fkey"))
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="AuditAssignment_userId_fkey"))
    role: Mapped[Optional[str]] = mapped_column(Enum("LEAD", "PARTICIPANT", name="AuditAssignmentRole"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("AuditAssignment_userId_idx", "userId"),
        Index("AuditAssignment_auditSessionId_userId_key", "auditSessionId", "userId", unique=True),
    )

    audit_session: Mapped["AuditSession"] = relationship(back_populates="assignments")
    user: Mapped["User"] = relationship(back_populates="audit_assignments")

class AuditAsset(Base):
    __tablename__ = "AuditAsset"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    audit_session_id: Mapped[str] = mapped_column("auditSessionId", ForeignKey("AuditSession.id", ondelete="CASCADE", onupdate="CASCADE", name="AuditAsset_auditSessionId_fkey"))
    asset_id: Mapped[str] = mapped_column("assetId", ForeignKey("Asset.id", ondelete="CASCADE", onupdate="CASCADE", name="AuditAsset_assetId_fkey"))
    expected: Mapped[bool] = mapped_column(default=True, server_default=text("true"))
    status: Mapped[str] = mapped_column(Enum("PENDING", "FOUND", "MISSING", "UNEXPECTED", name="AuditAssetStatus"), default="PENDING", server_default=text("'PENDING'"))
    scanned_by_id: Mapped[Optional[str]] = mapped_column("scannedById", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="AuditAsset_scannedById_fkey"))
    scanned_at: Mapped[Optional[datetime]] = mapped_column("scannedAt")
    metadata_: Mapped[Optional[dict[str, Any]]] = mapped_column("metadata", JSONB(none_as_null=True))
    expected_quantity: Mapped[Optional[int]] = mapped_column("expectedQuantity")
    """Units the audit expects to find, snapshotted when the row is created: the
    units in the audit's scope (placed at its locations, held in its kits, or
    held by its person; at most `Asset.quantity`), or `Asset.quantity` when
    none are in scope or the audit has no scope. Null for INDIVIDUAL assets,
    which count as one.

    A snapshot, not a live read: stock that moves during the audit must not
    move the target the auditor is counting against.
    """
    found_quantity: Mapped[Optional[int]] = mapped_column("foundQuantity")
    """Units the auditor counted. Null until a quantity-tracked asset is
    counted, and always null for INDIVIDUAL assets. A FOUND row with a count
    below `expectedQuantity` is a shortfall: the status stays FOUND (the pool
    was located), the gap is this number.

    Both columns are non-negative, enforced by CHECK constraints in raw SQL
    (migration `20261104120000_add_audit_asset_quantities`).
    """
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("AuditAsset_status_idx", "status"),
        Index("AuditAsset_scannedById_idx", "scannedById"),
        Index("AuditAsset_assetId_idx", "assetId"),
        Index("AuditAsset_auditSessionId_assetId_key", "auditSessionId", "assetId", unique=True),
    )

    audit_session: Mapped["AuditSession"] = relationship(back_populates="assets")
    asset: Mapped["Asset"] = relationship(back_populates="audit_entries")
    scanned_by: Mapped[Optional["User"]] = relationship(back_populates="audit_assets_scanned")
    scans: Mapped[list["AuditScan"]] = relationship(back_populates="audit_asset")
    images: Mapped[list["AuditImage"]] = relationship(back_populates="audit_asset")
    notes: Mapped[list["AuditNote"]] = relationship(back_populates="audit_asset")

class AuditScan(Base):
    __tablename__ = "AuditScan"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    audit_session_id: Mapped[str] = mapped_column("auditSessionId", ForeignKey("AuditSession.id", ondelete="CASCADE", onupdate="CASCADE", name="AuditScan_auditSessionId_fkey"))
    audit_asset_id: Mapped[Optional[str]] = mapped_column("auditAssetId", ForeignKey("AuditAsset.id", ondelete="SET NULL", onupdate="CASCADE", name="AuditScan_auditAssetId_fkey"))
    asset_id: Mapped[Optional[str]] = mapped_column("assetId", ForeignKey("Asset.id", ondelete="SET NULL", onupdate="CASCADE", name="AuditScan_assetId_fkey"))
    scanned_by_id: Mapped[Optional[str]] = mapped_column("scannedById", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="AuditScan_scannedById_fkey"))
    code: Mapped[Optional[str]]
    asset_title: Mapped[Optional[str]] = mapped_column("assetTitle")
    """What was scanned, recorded BY VALUE at scan time.

    `asset` and `auditAsset` are both SetNull/Cascade, so deleting an asset
    leaves this row pointing at nothing: no name, and no way to tell whether
    it had been expected. An audit is a historical record, and a record whose
    entries can silently lose their meaning is not one — so the two
    descriptive facts are snapshotted here rather than only referenced.

    Readers prefer the LIVE asset when it still exists (a rename should show
    the current name) and fall back to these.

    Null means the snapshot could not be taken: scans whose asset had already
    been deleted before these columns were backfilled. Nothing can recover
    those, and they degrade to the scanned `code`, which also survives. Every
    other row carries a snapshot, so null is never merely "old".
    """
    was_expected: Mapped[Optional[bool]] = mapped_column("wasExpected")
    metadata_: Mapped[Optional[dict[str, Any]]] = mapped_column("metadata", JSONB(none_as_null=True))
    scanned_at: Mapped[datetime] = mapped_column("scannedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("AuditScan_auditSessionId_scannedAt_idx", "auditSessionId", "scannedAt"),
        Index("AuditScan_auditAssetId_idx", "auditAssetId"),
        Index("AuditScan_assetId_idx", "assetId"),
    )

    audit_session: Mapped["AuditSession"] = relationship(back_populates="scans")
    audit_asset: Mapped[Optional["AuditAsset"]] = relationship(back_populates="scans")
    asset: Mapped[Optional["Asset"]] = relationship(back_populates="audit_scans")
    scanned_by: Mapped[Optional["User"]] = relationship(back_populates="audit_scans")

class AuditImage(Base):
    __tablename__ = "AuditImage"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    image_url: Mapped[str] = mapped_column("imageUrl")
    thumbnail_url: Mapped[Optional[str]] = mapped_column("thumbnailUrl")
    description: Mapped[Optional[str]]
    audit_session_id: Mapped[str] = mapped_column("auditSessionId", ForeignKey("AuditSession.id", ondelete="CASCADE", onupdate="CASCADE", name="AuditImage_auditSessionId_fkey"))
    audit_asset_id: Mapped[Optional[str]] = mapped_column("auditAssetId", ForeignKey("AuditAsset.id", ondelete="SET NULL", onupdate="CASCADE", name="AuditImage_auditAssetId_fkey"))
    uploaded_by_id: Mapped[Optional[str]] = mapped_column("uploadedById", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="AuditImage_uploadedById_fkey"))
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="AuditImage_organizationId_fkey"))
    captured_at: Mapped[Optional[datetime]] = mapped_column("capturedAt")
    """When the photo was taken, from its EXIF data (or the phone app's copy of
    it). Null when the photo carried no usable date. EXIF can be edited, so
    this is the best signal available, not proof.
    """
    capture_source: Mapped[Optional[str]] = mapped_column("captureSource", Enum("CAMERA", "LIBRARY", "WEB_UPLOAD", "PHONE", name="AuditImageCaptureSource"))
    """How the photo reached Shelf. Null only for photos uploaded before this
    field existed; a phone upload with no or an unknown source is PHONE.
    """
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("AuditImage_auditSessionId_idx", "auditSessionId"),
        Index("AuditImage_auditAssetId_idx", "auditAssetId"),
        Index("AuditImage_organizationId_idx", "organizationId"),
        Index("AuditImage_uploadedById_idx", "uploadedById"),
    )

    audit_session: Mapped["AuditSession"] = relationship(back_populates="images")
    audit_asset: Mapped[Optional["AuditAsset"]] = relationship(back_populates="images")
    uploaded_by: Mapped[Optional["User"]] = relationship(back_populates="audit_images_uploaded")
    organization: Mapped["Organization"] = relationship(back_populates="audit_images")

class RoleChangeLog(Base):
    __tablename__ = "RoleChangeLog"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    previous_role: Mapped[str] = mapped_column("previousRole", Enum("ADMIN", "BASE", "OWNER", "SELF_SERVICE", "MANAGER", name="OrganizationRoles"))
    new_role: Mapped[str] = mapped_column("newRole", Enum("ADMIN", "BASE", "OWNER", "SELF_SERVICE", "MANAGER", name="OrganizationRoles"))
    source: Mapped[str] = mapped_column(Enum("MANUAL", "SSO", name="RoleChangeSource"), default="MANUAL", server_default=text("'MANUAL'"))
    """How the change was made. SSO rows name the member themself as `changedBy`."""
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="RESTRICT", onupdate="CASCADE", name="RoleChangeLog_userId_fkey"))
    changed_by_id: Mapped[str] = mapped_column("changedById", ForeignKey("User.id", ondelete="RESTRICT", onupdate="CASCADE", name="RoleChangeLog_changedById_fkey"))
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="RESTRICT", onupdate="CASCADE", name="RoleChangeLog_organizationId_fkey"))

    __table_args__ = (
        Index("RoleChangeLog_userId_organizationId_idx", "userId", "organizationId"),
        Index("RoleChangeLog_organizationId_createdAt_idx", "organizationId", "createdAt"),
    )

    user: Mapped["User"] = relationship(foreign_keys=[user_id], back_populates="role_changes_received")
    changed_by: Mapped["User"] = relationship(foreign_keys=[changed_by_id], back_populates="role_changes_performed")
    organization: Mapped["Organization"] = relationship(back_populates="role_change_logs")

class ActivityEvent(Base):
    __tablename__ = "ActivityEvent"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="ActivityEvent_organizationId_fkey"))
    occurred_at: Mapped[datetime] = mapped_column("occurredAt", UtcDateTimeTz(precision=3), default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    actor_user_id: Mapped[Optional[str]] = mapped_column("actorUserId")
    actor_snapshot: Mapped[Optional[dict[str, Any]]] = mapped_column("actorSnapshot", JSONB(none_as_null=True))
    action: Mapped[str] = mapped_column(Enum("ASSET_CREATED", "ASSET_NAME_CHANGED", "ASSET_DESCRIPTION_CHANGED", "ASSET_CATEGORY_CHANGED", "ASSET_MODEL_CHANGED", "ASSET_KIT_CHANGED", "ASSET_LOCATION_CHANGED", "ASSET_TAGS_CHANGED", "ASSET_STATUS_CHANGED", "ASSET_VALUATION_CHANGED", "ASSET_QUANTITY_CHANGED", "ASSET_MIN_QUANTITY_CHANGED", "ASSET_CUSTOM_FIELD_CHANGED", "ASSET_PREFERRED_BARCODE_CHANGED", "ASSET_AVAILABLE_TO_BOOK_CHANGED", "ASSET_ARCHIVED", "ASSET_UNARCHIVED", "ASSET_DELETED", "CUSTODY_ASSIGNED", "CUSTODY_RELEASED", "BOOKING_CREATED", "BOOKING_STATUS_CHANGED", "BOOKING_DATES_CHANGED", "BOOKING_ASSETS_ADDED", "BOOKING_ASSETS_REMOVED", "BOOKING_CHECKED_OUT", "BOOKING_CHECKED_IN", "BOOKING_PARTIAL_CHECKIN", "BOOKING_PARTIAL_CHECKOUT", "BOOKING_CANCELLED", "BOOKING_ARCHIVED", "BOOKING_MODEL_REQUESTED", "BOOKING_MODEL_REQUEST_CHANGED", "BOOKING_MODEL_REQUEST_REMOVED", "BOOKING_MODEL_REQUEST_FULFILLED", "AUDIT_CREATED", "AUDIT_STARTED", "AUDIT_ASSETS_ADDED", "AUDIT_ASSETS_REMOVED", "AUDIT_ASSET_SCANNED", "AUDIT_ASSET_SCAN_REMOVED", "AUDIT_DUE_DATE_CHANGED", "AUDIT_ASSIGNEE_ADDED", "AUDIT_ASSIGNEE_REMOVED", "AUDIT_UPDATED", "AUDIT_COMPLETED", "AUDIT_CANCELLED", "AUDIT_ARCHIVED", "LOCATION_CREATED", "LOCATION_UPDATED", "KIT_CREATED", "KIT_UPDATED", "ORGANIZATION_QR_ID_DISPLAY_PREFERENCE_CHANGED", name="ActivityAction"))
    entity_type: Mapped[str] = mapped_column("entityType", Enum("ASSET", "BOOKING", "AUDIT", "KIT", "LOCATION", "TEAM_MEMBER", "CUSTODY", "USER", "ORGANIZATION", name="ActivityEntity"))
    entity_id: Mapped[str] = mapped_column("entityId")
    asset_id: Mapped[Optional[str]] = mapped_column("assetId")
    booking_id: Mapped[Optional[str]] = mapped_column("bookingId")
    audit_session_id: Mapped[Optional[str]] = mapped_column("auditSessionId")
    audit_asset_id: Mapped[Optional[str]] = mapped_column("auditAssetId")
    kit_id: Mapped[Optional[str]] = mapped_column("kitId")
    location_id: Mapped[Optional[str]] = mapped_column("locationId")
    team_member_id: Mapped[Optional[str]] = mapped_column("teamMemberId")
    target_user_id: Mapped[Optional[str]] = mapped_column("targetUserId")
    field: Mapped[Optional[str]]
    from_value: Mapped[Optional[dict[str, Any]]] = mapped_column("fromValue", JSONB(none_as_null=True))
    to_value: Mapped[Optional[dict[str, Any]]] = mapped_column("toValue", JSONB(none_as_null=True))
    meta: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))

    __table_args__ = (
        Index("ActivityEvent_organizationId_occurredAt_idx", "organizationId", "occurredAt"),
        Index("ActivityEvent_organizationId_action_occurredAt_idx", "organizationId", "action", "occurredAt"),
        Index("ActivityEvent_organizationId_entityType_entityId_occurredAt_idx", "organizationId", "entityType", "entityId", "occurredAt"),
        Index("ActivityEvent_actorUserId_occurredAt_idx", "actorUserId", "occurredAt"),
        Index("ActivityEvent_assetId_occurredAt_idx", "assetId", "occurredAt"),
        Index("ActivityEvent_bookingId_occurredAt_idx", "bookingId", "occurredAt"),
        Index("ActivityEvent_auditSessionId_occurredAt_idx", "auditSessionId", "occurredAt"),
    )

    organization: Mapped["Organization"] = relationship(back_populates="activity_events")
