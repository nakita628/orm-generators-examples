from datetime import datetime
from datetime import timezone as dt_timezone
from typing import Any, TypeVar

import cuid  # type: ignore[import-untyped, unused-ignore]
from django.conf import settings
from django.contrib.postgres.fields import ArrayField
from django.contrib.postgres.indexes import GinIndex
from django.db import models
from django.db.backends.base.base import BaseDatabaseWrapper
from django.utils import timezone


class UtcDateTimeField(models.DateTimeField):  # type: ignore[type-arg]
    def get_db_prep_value(self, value: Any, connection: BaseDatabaseWrapper, prepared: bool = False) -> Any:
        if not prepared:
            value = self.get_prep_value(value)
        if not isinstance(value, datetime):
            return super().get_db_prep_value(value, connection, prepared=True)
        if timezone.is_naive(value):
            value = timezone.make_aware(value, timezone.get_default_timezone())
        value = value.astimezone(dt_timezone.utc)
        if connection.vendor == "sqlite":
            return value.isoformat(timespec="milliseconds")
        return connection.ops.adapt_datetimefield_value(value.replace(tzinfo=None))

    def from_db_value(self, value: datetime | None, expression: Any, connection: BaseDatabaseWrapper) -> datetime | None:
        if value is None:
            return None
        if timezone.is_naive(value):
            value = value.replace(tzinfo=dt_timezone.utc)
        return value if settings.USE_TZ else timezone.make_naive(value, timezone.get_default_timezone())


class UtcAutoNowDateTimeField(UtcDateTimeField):
    def stamp(self) -> datetime:
        return now_milliseconds()

    def pre_save(self, model_instance: models.Model, add: bool) -> Any:
        if not self.auto_now:
            return super().pre_save(model_instance, add)
        value = self.stamp()
        setattr(model_instance, self.attname, value)
        return value


class AutoNowDateTimeField(models.DateTimeField):  # type: ignore[type-arg]
    def stamp(self) -> datetime:
        return now_milliseconds()

    def pre_save(self, model_instance: models.Model, add: bool) -> Any:
        if not self.auto_now:
            return super().pre_save(model_instance, add)
        value = self.stamp()
        setattr(model_instance, self.attname, value)
        return value


class EnumArrayField(ArrayField):  # type: ignore[type-arg]
    def __init__(self, base_field: Any, db_enum: str, **kwargs: Any) -> None:
        self.db_enum = db_enum
        super().__init__(base_field, **kwargs)

    def deconstruct(self) -> Any:
        name, path, args, kwargs = super().deconstruct()
        kwargs["db_enum"] = self.db_enum
        return name, path, args, kwargs

    def db_type(self, connection: BaseDatabaseWrapper) -> str:
        return f"{self.db_enum}[]"

    def cast_db_type(self, connection: BaseDatabaseWrapper) -> str:
        return self.db_type(connection)

    def select_format(self, compiler: Any, sql: str, params: Any) -> Any:
        return f"{sql}::text[]", params


_M = TypeVar("_M", bound=models.Model)


class AutoNowQuerySet(models.QuerySet[_M]):
    def update(self, **kwargs: Any) -> int:
        for field in self.model._meta.concrete_fields:
            if getattr(field, "auto_now", False):
                kwargs.setdefault(field.name, getattr(field, "stamp", timezone.now)())
        return super().update(**kwargs)


def cuid_str() -> str:
    return str(cuid.cuid())


def now_milliseconds() -> datetime:
    now = timezone.now()
    return now.replace(microsecond=now.microsecond // 1000 * 1000)


def asset_index_settings_columns_default() -> list[Any]:
    return [{"name": "id", "visible": True, "position": 0}, {"name": "status", "visible": True, "position": 1}, {"name": "description", "visible": True, "position": 2}, {"name": "valuation", "visible": True, "position": 3}, {"name": "createdAt", "visible": True, "position": 4}, {"name": "category", "visible": True, "position": 5}, {"name": "tags", "visible": True, "position": 6}, {"name": "location", "visible": True, "position": 7}, {"name": "kit", "visible": True, "position": 8}, {"name": "custody", "visible": True, "position": 9}]


def tag_use_for_default() -> list[str]:
    return [TagUseFor.ASSET]


def working_hours_weekly_schedule_default() -> dict[str, Any]:
    return {"0": {"isOpen": False}, "1": {"isOpen": True, "openTime": "09:00", "closeTime": "17:00"}, "2": {"isOpen": True, "openTime": "09:00", "closeTime": "17:00"}, "3": {"isOpen": True, "openTime": "09:00", "closeTime": "17:00"}, "4": {"isOpen": True, "openTime": "09:00", "closeTime": "17:00"}, "5": {"isOpen": True, "openTime": "09:00", "closeTime": "17:00"}, "6": {"isOpen": False}}


class AssetStatus(models.TextChoices):
    AVAILABLE = "AVAILABLE"
    IN_CUSTODY = "IN_CUSTODY"
    CHECKED_OUT = "CHECKED_OUT"


class AssetType(models.TextChoices):
    """* Determines how an asset is tracked throughout the system. Immutable after creation."""

    INDIVIDUAL = "INDIVIDUAL"
    QUANTITY_TRACKED = "QUANTITY_TRACKED"


class ConsumptionType(models.TextChoices):
    """* Controls what happens after checkout for quantity-tracked assets."""

    ONE_WAY = "ONE_WAY"
    TWO_WAY = "TWO_WAY"


class ConsumptionCategory(models.TextChoices):
    """* Classifies the type of quantity change in a ConsumptionLog entry."""

    CHECKOUT = "CHECKOUT"
    RETURN = "RETURN"
    RESTOCK = "RESTOCK"
    ADJUSTMENT = "ADJUSTMENT"
    LOSS = "LOSS"
    CONSUME = "CONSUME"
    DAMAGE = "DAMAGE"
    MOVE = "MOVE"
    INITIAL = "INITIAL"
    OPENING_BALANCE = "OPENING_BALANCE"


class PdfTemplateKind(models.TextChoices):
    """The kind of PDF a PdfTemplate lays out. Only the booking PDF has saved
    layouts today; the check-in receipt and the audit receipt can join as new
    kinds.
    """

    BOOKING = "BOOKING"


class AssetIndexMode(models.TextChoices):
    SIMPLE = "SIMPLE"
    ADVANCED = "ADVANCED"


class TagUseFor(models.TextChoices):
    ASSET = "ASSET"
    BOOKING = "BOOKING"


class NoteType(models.TextChoices):
    COMMENT = "COMMENT"
    UPDATE = "UPDATE"


class BarcodeType(models.TextChoices):
    Code128 = "Code128"
    Code39 = "Code39"
    DataMatrix = "DataMatrix"
    ExternalQR = "ExternalQR"
    EAN13 = "EAN13"


class ErrorCorrection(models.TextChoices):
    L = "L"
    M = "M"
    Q = "Q"
    H = "H"


class ScanCodeType(models.TextChoices):
    """The kind of code a `Scan` recorded."""

    QR = "QR"
    BARCODE = "BARCODE"
    SAM_ID = "SAM_ID"


class ScanSource(models.TextChoices):
    """The surface that recorded a `Scan`."""

    QR_LINK = "QR_LINK"
    WEB_SCANNER = "WEB_SCANNER"
    WEB_DRAWER = "WEB_DRAWER"
    AUDIT = "AUDIT"
    COMPANION = "COMPANION"
    GPS_UPDATE = "GPS_UPDATE"


class Roles(models.TextChoices):
    USER = "USER"
    ADMIN = "ADMIN"


class OrganizationType(models.TextChoices):
    PERSONAL = "PERSONAL"
    TEAM = "TEAM"


class QrIdDisplayPreference(models.TextChoices):
    QR_ID = "QR_ID"
    SAM_ID = "SAM_ID"
    Code128 = "Code128"
    Code39 = "Code39"
    DataMatrix = "DataMatrix"
    ExternalQR = "ExternalQR"
    EAN13 = "EAN13"


class DateFormatPreference(models.TextChoices):
    DD_MM_YYYY = "DD_MM_YYYY"
    MM_DD_YYYY = "MM_DD_YYYY"
    YYYY_MM_DD = "YYYY_MM_DD"
    MMM_DD_YYYY = "MMM_DD_YYYY"
    DD_MMM_YYYY = "DD_MMM_YYYY"


class TimeFormatPreference(models.TextChoices):
    H12 = "H12"
    H24 = "H24"


class WeekStartPreference(models.TextChoices):
    MONDAY = "MONDAY"
    SUNDAY = "SUNDAY"
    SATURDAY = "SATURDAY"


class OrganizationRoles(models.TextChoices):
    ADMIN = "ADMIN"
    BASE = "BASE"
    OWNER = "OWNER"
    SELF_SERVICE = "SELF_SERVICE"
    MANAGER = "MANAGER"


class TierId(models.TextChoices):
    free = "free"
    tier_1 = "tier_1"
    tier_2 = "tier_2"
    custom = "custom"


class CustomFieldType(models.TextChoices):
    TEXT = "TEXT"
    OPTION = "OPTION"
    BOOLEAN = "BOOLEAN"
    DATE = "DATE"
    MULTILINE_TEXT = "MULTILINE_TEXT"
    AMOUNT = "AMOUNT"
    NUMBER = "NUMBER"


class CustomFieldFormat(models.TextChoices):
    """The format of a TEXT custom field. Unlike the type, it can change after the
    field exists.
    """

    PLAIN = "PLAIN"
    LINK = "LINK"


class Currency(models.TextChoices):
    AED = "AED"
    AFN = "AFN"
    ALL = "ALL"
    AMD = "AMD"
    ANG = "ANG"
    AOA = "AOA"
    ARS = "ARS"
    AUD = "AUD"
    AWG = "AWG"
    AZN = "AZN"
    BAM = "BAM"
    BBD = "BBD"
    BDT = "BDT"
    BGN = "BGN"
    BHD = "BHD"
    BIF = "BIF"
    BMD = "BMD"
    BND = "BND"
    BOB = "BOB"
    BRL = "BRL"
    BSD = "BSD"
    BTN = "BTN"
    BWP = "BWP"
    BYN = "BYN"
    BZD = "BZD"
    CAD = "CAD"
    CDF = "CDF"
    CHF = "CHF"
    CLP = "CLP"
    CNY = "CNY"
    COP = "COP"
    CRC = "CRC"
    CUP = "CUP"
    CVE = "CVE"
    CZK = "CZK"
    DJF = "DJF"
    DKK = "DKK"
    DOP = "DOP"
    DZD = "DZD"
    EGP = "EGP"
    ERN = "ERN"
    ETB = "ETB"
    EUR = "EUR"
    FJD = "FJD"
    FKP = "FKP"
    GBP = "GBP"
    GEL = "GEL"
    GHS = "GHS"
    GIP = "GIP"
    GMD = "GMD"
    GNF = "GNF"
    GTQ = "GTQ"
    GYD = "GYD"
    HKD = "HKD"
    HNL = "HNL"
    HTG = "HTG"
    HUF = "HUF"
    IDR = "IDR"
    ILS = "ILS"
    INR = "INR"
    IQD = "IQD"
    IRR = "IRR"
    ISK = "ISK"
    JMD = "JMD"
    JOD = "JOD"
    JPY = "JPY"
    KES = "KES"
    KGS = "KGS"
    KHR = "KHR"
    KMF = "KMF"
    KPW = "KPW"
    KRW = "KRW"
    KWD = "KWD"
    KYD = "KYD"
    KZT = "KZT"
    LAK = "LAK"
    LBP = "LBP"
    LKR = "LKR"
    LRD = "LRD"
    LSL = "LSL"
    LYD = "LYD"
    MAD = "MAD"
    MDL = "MDL"
    MGA = "MGA"
    MKD = "MKD"
    MMK = "MMK"
    MNT = "MNT"
    MOP = "MOP"
    MRU = "MRU"
    MUR = "MUR"
    MVR = "MVR"
    MWK = "MWK"
    MXN = "MXN"
    MYR = "MYR"
    MZN = "MZN"
    NAD = "NAD"
    NGN = "NGN"
    NIO = "NIO"
    NOK = "NOK"
    NPR = "NPR"
    NZD = "NZD"
    OMR = "OMR"
    PAB = "PAB"
    PEN = "PEN"
    PGK = "PGK"
    PHP = "PHP"
    PKR = "PKR"
    PLN = "PLN"
    PYG = "PYG"
    QAR = "QAR"
    RON = "RON"
    RSD = "RSD"
    RUB = "RUB"
    RWF = "RWF"
    SAR = "SAR"
    SBD = "SBD"
    SCR = "SCR"
    SDG = "SDG"
    SEK = "SEK"
    SGD = "SGD"
    SHP = "SHP"
    SLE = "SLE"
    SOS = "SOS"
    SRD = "SRD"
    SSP = "SSP"
    STN = "STN"
    SVC = "SVC"
    SYP = "SYP"
    SZL = "SZL"
    THB = "THB"
    TJS = "TJS"
    TMT = "TMT"
    TND = "TND"
    TOP = "TOP"
    TRY = "TRY"
    TTD = "TTD"
    TWD = "TWD"
    TZS = "TZS"
    UAH = "UAH"
    UGX = "UGX"
    USD = "USD"
    UYU = "UYU"
    UZS = "UZS"
    VES = "VES"
    VND = "VND"
    VUV = "VUV"
    WST = "WST"
    XAF = "XAF"
    XCD = "XCD"
    XOF = "XOF"
    XPF = "XPF"
    YER = "YER"
    ZAR = "ZAR"
    ZMW = "ZMW"
    ZWL = "ZWL"


class InviteStatuses(models.TextChoices):
    PENDING = "PENDING"
    ACCEPTED = "ACCEPTED"
    REJECTED = "REJECTED"
    INVALIDATED = "INVALIDATED"


class BookingStatus(models.TextChoices):
    DRAFT = "DRAFT"
    RESERVED = "RESERVED"
    ONGOING = "ONGOING"
    OVERDUE = "OVERDUE"
    COMPLETE = "COMPLETE"
    ARCHIVED = "ARCHIVED"
    CANCELLED = "CANCELLED"


class DevicePlatform(models.TextChoices):
    """The operating system a companion-app push token was issued for."""

    IOS = "IOS"
    ANDROID = "ANDROID"


class KitStatus(models.TextChoices):
    AVAILABLE = "AVAILABLE"
    IN_CUSTODY = "IN_CUSTODY"
    CHECKED_OUT = "CHECKED_OUT"


class ReminderRecurrenceUnit(models.TextChoices):
    DAY = "DAY"
    WEEK = "WEEK"
    MONTH = "MONTH"
    YEAR = "YEAR"


class UpdateStatus(models.TextChoices):
    DRAFT = "DRAFT"
    PUBLISHED = "PUBLISHED"


class AuditStatus(models.TextChoices):
    PENDING = "PENDING"
    ACTIVE = "ACTIVE"
    COMPLETED = "COMPLETED"
    CANCELLED = "CANCELLED"
    ARCHIVED = "ARCHIVED"


class AuditAssetStatus(models.TextChoices):
    PENDING = "PENDING"
    FOUND = "FOUND"
    MISSING = "MISSING"
    UNEXPECTED = "UNEXPECTED"


class AuditAssignmentRole(models.TextChoices):
    LEAD = "LEAD"
    PARTICIPANT = "PARTICIPANT"


class AssetImportStatus(models.TextChoices):
    """Lifecycle of a background asset import."""

    DRAFT = "DRAFT"
    QUEUED = "QUEUED"
    RUNNING = "RUNNING"
    COMPLETED = "COMPLETED"
    FAILED = "FAILED"


class AssetImportPhase(models.TextChoices):
    """Which step of the background import runs next (or failed)."""

    PREPARE = "PREPARE"
    ROWS = "ROWS"
    FINALIZE = "FINALIZE"


class LinkPreviewOutcome(models.TextChoices):
    """How the last fetch of a link's page went."""

    TITLE = "TITLE"
    NO_TITLE = "NO_TITLE"
    SIGN_IN = "SIGN_IN"
    BLOCKED = "BLOCKED"
    UNREACHABLE = "UNREACHABLE"


class AuditImageCaptureSource(models.TextChoices):
    """How an audit photo reached Shelf, as reported by the client. The phone app
    reports CAMERA only for its own camera, which cannot hand back an old
    picture, so it is the strongest evidence short of a modified client; a
    library photo's date is only what its file claims; the web app cannot tell
    a camera shot from a file; PHONE is a phone app version that does not say
    which.
    """

    CAMERA = "CAMERA"
    LIBRARY = "LIBRARY"
    WEB_UPLOAD = "WEB_UPLOAD"
    PHONE = "PHONE"


class RoleChangeSource(models.TextChoices):
    """How a membership's role came to change."""

    MANUAL = "MANUAL"
    SSO = "SSO"


class ActivityEntity(models.TextChoices):
    ASSET = "ASSET"
    BOOKING = "BOOKING"
    AUDIT = "AUDIT"
    KIT = "KIT"
    LOCATION = "LOCATION"
    TEAM_MEMBER = "TEAM_MEMBER"
    CUSTODY = "CUSTODY"
    USER = "USER"
    ORGANIZATION = "ORGANIZATION"


class ActivityAction(models.TextChoices):
    ASSET_CREATED = "ASSET_CREATED"
    ASSET_NAME_CHANGED = "ASSET_NAME_CHANGED"
    ASSET_DESCRIPTION_CHANGED = "ASSET_DESCRIPTION_CHANGED"
    ASSET_CATEGORY_CHANGED = "ASSET_CATEGORY_CHANGED"
    ASSET_MODEL_CHANGED = "ASSET_MODEL_CHANGED"
    ASSET_KIT_CHANGED = "ASSET_KIT_CHANGED"
    ASSET_LOCATION_CHANGED = "ASSET_LOCATION_CHANGED"
    ASSET_TAGS_CHANGED = "ASSET_TAGS_CHANGED"
    ASSET_STATUS_CHANGED = "ASSET_STATUS_CHANGED"
    ASSET_VALUATION_CHANGED = "ASSET_VALUATION_CHANGED"
    ASSET_QUANTITY_CHANGED = "ASSET_QUANTITY_CHANGED"
    ASSET_MIN_QUANTITY_CHANGED = "ASSET_MIN_QUANTITY_CHANGED"
    ASSET_CUSTOM_FIELD_CHANGED = "ASSET_CUSTOM_FIELD_CHANGED"
    ASSET_PREFERRED_BARCODE_CHANGED = "ASSET_PREFERRED_BARCODE_CHANGED"
    ASSET_AVAILABLE_TO_BOOK_CHANGED = "ASSET_AVAILABLE_TO_BOOK_CHANGED"
    ASSET_ARCHIVED = "ASSET_ARCHIVED"
    ASSET_UNARCHIVED = "ASSET_UNARCHIVED"
    ASSET_DELETED = "ASSET_DELETED"
    CUSTODY_ASSIGNED = "CUSTODY_ASSIGNED"
    CUSTODY_RELEASED = "CUSTODY_RELEASED"
    BOOKING_CREATED = "BOOKING_CREATED"
    BOOKING_STATUS_CHANGED = "BOOKING_STATUS_CHANGED"
    BOOKING_DATES_CHANGED = "BOOKING_DATES_CHANGED"
    BOOKING_ASSETS_ADDED = "BOOKING_ASSETS_ADDED"
    BOOKING_ASSETS_REMOVED = "BOOKING_ASSETS_REMOVED"
    BOOKING_CHECKED_OUT = "BOOKING_CHECKED_OUT"
    BOOKING_CHECKED_IN = "BOOKING_CHECKED_IN"
    BOOKING_PARTIAL_CHECKIN = "BOOKING_PARTIAL_CHECKIN"
    BOOKING_PARTIAL_CHECKOUT = "BOOKING_PARTIAL_CHECKOUT"
    BOOKING_CANCELLED = "BOOKING_CANCELLED"
    BOOKING_ARCHIVED = "BOOKING_ARCHIVED"
    BOOKING_MODEL_REQUESTED = "BOOKING_MODEL_REQUESTED"
    BOOKING_MODEL_REQUEST_CHANGED = "BOOKING_MODEL_REQUEST_CHANGED"
    BOOKING_MODEL_REQUEST_REMOVED = "BOOKING_MODEL_REQUEST_REMOVED"
    BOOKING_MODEL_REQUEST_FULFILLED = "BOOKING_MODEL_REQUEST_FULFILLED"
    AUDIT_CREATED = "AUDIT_CREATED"
    AUDIT_STARTED = "AUDIT_STARTED"
    AUDIT_ASSETS_ADDED = "AUDIT_ASSETS_ADDED"
    AUDIT_ASSETS_REMOVED = "AUDIT_ASSETS_REMOVED"
    AUDIT_ASSET_SCANNED = "AUDIT_ASSET_SCANNED"
    AUDIT_ASSET_SCAN_REMOVED = "AUDIT_ASSET_SCAN_REMOVED"
    AUDIT_DUE_DATE_CHANGED = "AUDIT_DUE_DATE_CHANGED"
    AUDIT_ASSIGNEE_ADDED = "AUDIT_ASSIGNEE_ADDED"
    AUDIT_ASSIGNEE_REMOVED = "AUDIT_ASSIGNEE_REMOVED"
    AUDIT_UPDATED = "AUDIT_UPDATED"
    AUDIT_COMPLETED = "AUDIT_COMPLETED"
    AUDIT_CANCELLED = "AUDIT_CANCELLED"
    AUDIT_ARCHIVED = "AUDIT_ARCHIVED"
    LOCATION_CREATED = "LOCATION_CREATED"
    LOCATION_UPDATED = "LOCATION_UPDATED"
    KIT_CREATED = "KIT_CREATED"
    KIT_UPDATED = "KIT_UPDATED"
    ORGANIZATION_QR_ID_DISPLAY_PREFERENCE_CHANGED = "ORGANIZATION_QR_ID_DISPLAY_PREFERENCE_CHANGED"


class Image(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    content_type = models.TextField(db_column="contentType")
    alt_text = models.TextField(null=True, db_column="altText")
    blob = models.BinaryField()
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    owner_org = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="images", db_column="ownerOrgId", db_index=False)
    user = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="images", db_column="userId", db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Image"
        indexes = [
            models.Index(fields=["owner_org"]),
            models.Index(fields=["user"]),
        ]


class User(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    email = models.TextField(unique=True)
    username = models.TextField(unique=True, default=cuid_str)
    first_name = models.TextField(null=True, db_column="firstName")
    last_name = models.TextField(null=True, db_column="lastName")
    display_name = models.TextField(null=True, db_column="displayName")
    profile_picture = models.TextField(null=True, db_column="profilePicture")
    used_free_trial = models.BooleanField(db_column="usedFreeTrial", default=False)
    onboarded = models.BooleanField(default=False)
    customer_id = models.TextField(null=True, unique=True, db_column="customerId")
    sso = models.BooleanField(default=False)
    created_with_invite = models.BooleanField(db_column="createdWithInvite", default=False)
    skip_subscription_check = models.BooleanField(db_column="skipSubscriptionCheck", default=False)
    has_unpaid_invoice = models.BooleanField(db_column="hasUnpaidInvoice", default=False)
    warn_for_no_payment_method = models.BooleanField(db_column="warnForNoPaymentMethod", default=False)
    last_selected_organization = models.ForeignKey("Organization", on_delete=models.SET_NULL, related_name="users_with_last_selected", db_column="lastSelectedOrganizationId", null=True, db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    deleted_at = UtcDateTimeField(null=True, db_column="deletedAt")
    last_mobile_active_at = UtcDateTimeField(null=True, db_column="lastMobileActiveAt")
    date_format = models.TextField(choices=DateFormatPreference.choices, null=True, db_column="dateFormat")
    time_format = models.TextField(choices=TimeFormatPreference.choices, null=True, db_column="timeFormat")
    week_start = models.TextField(choices=WeekStartPreference.choices, null=True, db_column="weekStart")
    time_zone = models.TextField(null=True, db_column="timeZone")
    tier = models.ForeignKey("Tier", on_delete=models.RESTRICT, related_name="subscribers", db_column="tierId", db_index=False, default=TierId.free)
    referral_source = models.TextField(null=True, db_column="referralSource")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "User"
        constraints = [
            models.UniqueConstraint(fields=["email", "username"], name="User_email_username_key"),
        ]
        indexes = [
            models.Index(fields=["first_name", "last_name"]),
            models.Index(fields=["tier"]),
            models.Index(fields=["last_selected_organization"]),
            models.Index(fields=["last_mobile_active_at"]),
        ]


class UserContact(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    phone = models.TextField(null=True)
    street = models.TextField(null=True)
    city = models.TextField(null=True)
    state_province = models.TextField(null=True, db_column="stateProvince")
    zip_postal_code = models.TextField(null=True, db_column="zipPostalCode")
    country_region = models.TextField(null=True, db_column="countryRegion")
    user = models.OneToOneField("User", on_delete=models.CASCADE, related_name="contact", db_column="userId")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "UserContact"
        indexes = [
            models.Index(fields=["user"]),
            models.Index(fields=["phone"]),
            models.Index(fields=["city", "state_province"]),
            models.Index(fields=["country_region"]),
            models.Index(fields=["zip_postal_code"]),
            models.Index(fields=["city", "country_region"]),
        ]


class UserBusinessIntel(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    how_did_you_hear_about_us = models.TextField(null=True, db_column="howDidYouHearAboutUs")
    job_title = models.TextField(null=True, db_column="jobTitle")
    team_size = models.TextField(null=True, db_column="teamSize")
    company_name = models.TextField(null=True, db_column="companyName")
    primary_use_case = models.TextField(null=True, db_column="primaryUseCase")
    current_solution = models.TextField(null=True, db_column="currentSolution")
    timeline = models.TextField(null=True)
    signup_plan = models.TextField(null=True, db_column="signupPlan")
    signup_trial = models.BooleanField(null=True, db_column="signupTrial")
    utm_source = models.TextField(null=True, db_column="utmSource")
    utm_medium = models.TextField(null=True, db_column="utmMedium")
    utm_campaign = models.TextField(null=True, db_column="utmCampaign")
    utm_content = models.TextField(null=True, db_column="utmContent")
    user = models.OneToOneField("User", on_delete=models.CASCADE, related_name="business_intel", db_column="userId")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "UserBusinessIntel"
        indexes = [
            models.Index(fields=["user"]),
            models.Index(fields=["company_name"]),
            models.Index(fields=["job_title"]),
            models.Index(fields=["team_size"]),
            models.Index(fields=["utm_source"]),
            models.Index(fields=["utm_campaign"]),
        ]


class Asset(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    title = models.TextField()
    description = models.TextField(null=True)
    main_image = models.TextField(null=True, db_column="mainImage")
    thumbnail_image = models.TextField(null=True, db_column="thumbnailImage")
    main_image_expiration = UtcDateTimeField(null=True, db_column="mainImageExpiration")
    status = models.TextField(choices=AssetStatus.choices, default=AssetStatus.AVAILABLE)
    valuation = models.FloatField(null=True, db_column="value")
    available_to_book = models.BooleanField(db_column="availableToBook", default=True)
    sequential_id = models.TextField(null=True, db_column="sequentialId")
    archived_at = models.DateTimeField(null=True, db_column="archivedAt")
    type = models.TextField(choices=AssetType.choices, default=AssetType.INDIVIDUAL)
    asset_model = models.ForeignKey("AssetModel", on_delete=models.SET_NULL, related_name="assets", db_column="assetModelId", null=True, db_index=False)
    quantity = models.IntegerField(null=True)
    min_quantity = models.IntegerField(null=True, db_column="minQuantity")
    consumption_type = models.TextField(choices=ConsumptionType.choices, null=True, db_column="consumptionType")
    unit_of_measure = models.TextField(null=True, db_column="unitOfMeasure")
    low_stock_notified_at = UtcDateTimeField(null=True, db_column="lowStockNotifiedAt")
    stock_ledger_started_at = UtcDateTimeField(null=True, db_column="stockLedgerStartedAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="assets", db_column="userId", db_index=False)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="assets", db_column="organizationId", db_index=False)
    category = models.ForeignKey("Category", on_delete=models.SET_NULL, related_name="assets", db_column="categoryId", null=True, db_index=False)
    tags: "models.ManyToManyField[Tag, AssetToTag]" = models.ManyToManyField("Tag", through="AssetToTag", related_name="assets")
    preferred_barcode = models.OneToOneField("Barcode", on_delete=models.SET_NULL, related_name="preferred_by_asset", db_column="preferredBarcodeId", null=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Asset"
        constraints = [
            models.UniqueConstraint(fields=["organization", "sequential_id"], name="Asset_organizationId_sequentialId_key"),
        ]
        indexes = [
            GinIndex(fields=["title", "description"]),
            GinIndex(fields=["sequential_id"], name="Asset_sequentialId_trgm_idx"),
            models.Index(fields=["organization", "title", "status", "available_to_book"]),
            models.Index(fields=["status", "organization"]),
            models.Index(fields=["organization", "archived_at"]),
            models.Index(fields=["created_at", "organization"]),
            models.Index(fields=["organization", "-created_at", "id"]),
            models.Index(fields=["valuation", "organization"]),
            models.Index(fields=["category", "organization"]),
            models.Index(fields=["asset_model", "organization"]),
            models.Index(fields=["sequential_id"], name="Asset_sequentialId_idx"),
            models.Index(fields=["user"]),
        ]


class AssetFilterPreset(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="asset_filter_presets", db_column="organizationId", db_index=False)
    owner = models.ForeignKey("User", on_delete=models.CASCADE, related_name="asset_filter_presets", db_column="ownerId", db_index=False)
    name = models.TextField()
    query = models.TextField()
    starred = models.BooleanField(default=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "AssetFilterPreset"
        constraints = [
            models.UniqueConstraint(fields=["organization", "owner", "name"], name="asset_filter_presets_owner_name_unique"),
        ]
        indexes = [
            models.Index(fields=["organization", "owner"]),
        ]


class AssetIndexSettings(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="asset_index_settings", db_column="userId", db_index=False)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="asset_index_settings", db_column="organizationId", db_index=False)
    mode = models.TextField(choices=AssetIndexMode.choices, default=AssetIndexMode.SIMPLE)
    columns = models.JSONField(default=asset_index_settings_columns_default)
    freeze_column = models.BooleanField(db_column="freezeColumn", default=True)
    show_asset_image = models.BooleanField(db_column="showAssetImage", default=True)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "AssetIndexSettings"
        constraints = [
            models.UniqueConstraint(fields=["user", "organization"], name="AssetIndexSettings_userId_organizationId_key"),
        ]
        indexes = [
            models.Index(fields=["organization"]),
        ]


class PdfTemplate(models.Model):
    """A saved layout for a generated PDF, kept per team workspace under a name:
    which columns, rows and header lines print, and in what order the rows
    come. A row is a layout, not a generated file; the person generating a PDF
    picks a template by name.
    """

    id = models.TextField(primary_key=True, default=cuid_str)
    name = models.TextField()
    kind = models.TextField(choices=PdfTemplateKind.choices, default=PdfTemplateKind.BOOKING)
    config = models.JSONField(default=dict)
    """The layout's switches. Read through `parseBookingTemplateConfig`,
    which fills any key a stored row lacks from the Checklist template's
    values.
    """
    is_default = models.BooleanField(db_column="isDefault", default=False)
    """The template the PDF dialog offers first. One per workspace and kind,
    enforced by a partial unique index in raw SQL (migration
    20261103120000_add_pdf_templates). Names are also unique per workspace
    and kind ignoring case, through an expression index in the same
    migration.
    """
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="pdf_templates", db_column="organizationId", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "PdfTemplate"
        constraints = [
            models.UniqueConstraint(fields=["organization", "kind", "name"], name="PdfTemplate_organizationId_kind_name_key"),
        ]


class Category(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    name = models.TextField()
    description = models.TextField(null=True)
    color = models.TextField()
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    user = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="categories", db_column="userId", db_index=False)
    custom_fields: "models.ManyToManyField[CustomField, CategoryToCustomField]" = models.ManyToManyField("CustomField", through="CategoryToCustomField", related_name="categories")
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="categories", db_column="organizationId", db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Category"
        indexes = [
            models.Index(fields=["organization"]),
            models.Index(fields=["user"]),
            GinIndex(fields=["name"], name="Category_name_trgm_idx"),
        ]


class AssetModel(models.Model):
    """*
    * A template/grouping entity for assets.
    * Groups individual assets of the same type (e.g., "Dell Latitude 5550")
    * and provides default field values when creating new assets from the model.
    """

    id = models.TextField(primary_key=True, default=cuid_str)
    name = models.TextField()
    description = models.TextField(null=True)
    image = models.TextField(null=True)
    """Public URL (bucket `files`) of the model's cover image. Rendered by any
    asset of this model that has no image of its own. Public — and therefore
    never-expiring and cacheable — like Location.imageUrl; see
    resolveAssetImage() for the cascade.
    """
    thumbnail_image = models.TextField(null=True, db_column="thumbnailImage")
    """Public URL of the 108px thumbnail generated alongside `image`. Used at
    list sizes so inheriting assets never render the full-size original.
    """
    default_category = models.ForeignKey("Category", on_delete=models.SET_NULL, related_name="asset_model_defaults", db_column="defaultCategoryId", null=True, db_index=False)
    default_valuation = models.FloatField(null=True, db_column="defaultValuation")
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="asset_models", db_column="organizationId", db_index=False)
    created_by = models.ForeignKey("User", on_delete=models.CASCADE, related_name="asset_models", db_column="userId", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "AssetModel"
        indexes = [
            models.Index(fields=["organization", "name"]),
            models.Index(fields=["created_by"]),
            GinIndex(fields=["name"], name="AssetModel_name_trgm_idx"),
        ]


class BookingModelRequest(models.Model):
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

    id = models.TextField(primary_key=True, default=cuid_str)
    booking = models.ForeignKey("Booking", on_delete=models.CASCADE, related_name="model_requests", db_column="bookingId", db_index=False)
    asset_model = models.ForeignKey("AssetModel", on_delete=models.CASCADE, related_name="model_requests", db_column="assetModelId", db_index=False)
    quantity = models.IntegerField()
    """Total units originally reserved. Mutated only by explicit
    operator edits (upsert) — never decremented on materialisation.
    This preserves the audit trail of "how many were planned".
    """
    fulfilled_quantity = models.IntegerField(db_column="fulfilledQuantity", default=0)
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
    fulfilled_at = UtcDateTimeField(null=True, db_column="fulfilledAt")
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
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "BookingModelRequest"
        constraints = [
            models.UniqueConstraint(fields=["booking", "asset_model"], name="BookingModelRequest_bookingId_assetModelId_key"),
        ]
        indexes = [
            models.Index(fields=["asset_model"]),
        ]


class Tag(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    name = models.TextField()
    description = models.TextField(null=True)
    color = models.TextField(null=True)
    use_for = EnumArrayField(models.TextField(choices=TagUseFor.choices), db_enum="\"TagUseFor\"", db_column="useFor", default=tag_use_for_default)
    user = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="tags", db_column="userId", db_index=False)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="tags", db_column="organizationId", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Tag"
        indexes = [
            models.Index(fields=["organization"]),
            models.Index(fields=["user"]),
            GinIndex(fields=["name"], name="Tag_name_trgm_idx"),
        ]


class Note(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    content = models.TextField()
    type = models.TextField(choices=NoteType.choices, default=NoteType.COMMENT)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    user = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="notes", db_column="userId", null=True, db_index=False)
    asset = models.ForeignKey("Asset", on_delete=models.CASCADE, related_name="notes", db_column="assetId", db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Note"
        indexes = [
            models.Index(fields=["asset"]),
            models.Index(fields=["user"]),
        ]


class BookingNote(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    content = models.TextField()
    type = models.TextField(choices=NoteType.choices, default=NoteType.COMMENT)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    user = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="booking_notes", db_column="userId", null=True, db_index=False)
    booking = models.ForeignKey("Booking", on_delete=models.CASCADE, related_name="notes", db_column="bookingId", db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "BookingNote"
        indexes = [
            models.Index(fields=["booking"]),
            models.Index(fields=["user"]),
        ]


class LocationNote(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    content = models.TextField()
    type = models.TextField(choices=NoteType.choices, default=NoteType.COMMENT)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    user = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="location_notes", db_column="userId", null=True, db_index=False)
    location = models.ForeignKey("Location", on_delete=models.CASCADE, related_name="notes", db_column="locationId", db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "LocationNote"
        indexes = [
            models.Index(fields=["location"]),
            models.Index(fields=["user"]),
        ]


class Qr(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    version = models.IntegerField(default=0)
    error_correction = models.TextField(choices=ErrorCorrection.choices, db_column="errorCorrection", default=ErrorCorrection.L)
    asset = models.ForeignKey("Asset", on_delete=models.SET_NULL, related_name="qr_codes", db_column="assetId", null=True, db_index=False)
    kit = models.ForeignKey("Kit", on_delete=models.SET_NULL, related_name="qr_codes", db_column="kitId", null=True, db_index=False)
    user = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="qr_codes", db_column="userId", null=True, db_index=False)
    organization = models.ForeignKey("Organization", on_delete=models.SET_NULL, related_name="qr_codes", db_column="organizationId", null=True, db_index=False)
    batch = models.ForeignKey("PrintBatch", on_delete=models.SET_NULL, related_name="qr_codes", db_column="batchId", null=True, db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Qr"
        indexes = [
            models.Index(fields=["asset"], name="Qr_assetId_idx"),
            models.Index(fields=["kit"]),
            models.Index(fields=["user"]),
            models.Index(fields=["organization"]),
            models.Index(fields=["batch"]),
            GinIndex(fields=["id"], name="Qr_id_trgm_idx"),
        ]


class Barcode(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    value = models.TextField()
    type = models.TextField(choices=BarcodeType.choices, default=BarcodeType.Code128)
    asset = models.ForeignKey("Asset", on_delete=models.SET_NULL, related_name="barcodes", db_column="assetId", null=True, db_index=False)
    kit = models.ForeignKey("Kit", on_delete=models.SET_NULL, related_name="barcodes", db_column="kitId", null=True, db_index=False)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="barcodes", db_column="organizationId", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Barcode"
        constraints = [
            models.UniqueConstraint(fields=["organization", "value"], name="Barcode_organizationId_value_key"),
        ]
        indexes = [
            models.Index(fields=["organization", "value"]),
            models.Index(fields=["asset"]),
            models.Index(fields=["kit"]),
            models.Index(fields=["organization"]),
            GinIndex(fields=["value"], name="Barcode_value_trgm_idx"),
        ]


class PrintBatch(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    name = models.TextField(unique=True)
    printed = models.BooleanField(default=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "PrintBatch"


class ReportFound(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    email = models.TextField()
    content = models.TextField()
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    asset = models.ForeignKey("Asset", on_delete=models.CASCADE, related_name="reports", db_column="assetId", null=True, db_index=False)
    kit = models.ForeignKey("Kit", on_delete=models.CASCADE, related_name="reports", db_column="kitId", null=True, db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "ReportFound"
        indexes = [
            models.Index(fields=["asset"]),
            models.Index(fields=["kit"]),
        ]


class Scan(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    latitude = models.TextField(null=True)
    longitude = models.TextField(null=True)
    user_agent = models.TextField(null=True, db_column="userAgent")
    user = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="scans", db_column="userId", null=True, db_index=False)
    qr = models.ForeignKey("Qr", on_delete=models.SET_NULL, related_name="scans", db_column="qrId", null=True, db_index=False)
    raw_qr_id = models.TextField(null=True, db_column="rawQrId")
    code_type = models.TextField(choices=ScanCodeType.choices, db_column="codeType", default=ScanCodeType.QR)
    code = models.TextField(null=True)
    barcode = models.ForeignKey("Barcode", on_delete=models.SET_NULL, related_name="scans", db_column="barcodeId", null=True, db_index=False)
    asset = models.ForeignKey("Asset", on_delete=models.SET_NULL, related_name="scans", db_column="assetId", null=True, db_index=False)
    kit = models.ForeignKey("Kit", on_delete=models.SET_NULL, related_name="scans", db_column="kitId", null=True, db_index=False)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="scans", db_column="organizationId", db_index=False)
    source = models.TextField(choices=ScanSource.choices, null=True)
    manually_generated = models.BooleanField(db_column="manuallyGenerated", default=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Scan"
        indexes = [
            models.Index(fields=["qr"]),
            models.Index(fields=["user"]),
            models.Index(fields=["raw_qr_id", "created_at"]),
            models.Index(fields=["asset", "created_at"]),
            models.Index(fields=["kit", "created_at"]),
            models.Index(fields=["barcode"]),
            models.Index(fields=["organization"]),
        ]


class Location(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    name = models.TextField()
    description = models.TextField(null=True)
    address = models.TextField(null=True)
    latitude = models.FloatField(null=True)
    longitude = models.FloatField(null=True)
    image_url = models.TextField(null=True, db_column="imageUrl")
    thumbnail_url = models.TextField(null=True, db_column="thumbnailUrl")
    image = models.OneToOneField("Image", on_delete=models.SET_NULL, related_name="location", db_column="imageId", null=True)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="locations", db_column="userId", db_index=False)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="locations", db_column="organizationId", db_index=False)
    parent = models.ForeignKey("self", on_delete=models.SET_NULL, related_name="children", db_column="parentId", null=True, db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Location"
        indexes = [
            models.Index(fields=["organization"]),
            models.Index(fields=["user"]),
            models.Index(fields=["organization", "parent"]),
            GinIndex(fields=["name"], name="Location_name_trgm_idx"),
            GinIndex(fields=["description"], name="Location_description_trgm_idx"),
            GinIndex(fields=["address"], name="Location_address_trgm_idx"),
        ]


class Role(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    name = models.TextField(choices=Roles.choices, unique=True, default=Roles.USER)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    users: "models.ManyToManyField[User, RoleToUser]" = models.ManyToManyField("User", through="RoleToUser", related_name="roles")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Role"


class TeamMember(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    name = models.TextField()
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="members", db_column="organizationId", db_index=False)
    user = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="team_members", db_column="userId", null=True, db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    deleted_at = UtcDateTimeField(null=True, db_column="deletedAt")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "TeamMember"
        indexes = [
            GinIndex(fields=["name"]),
            models.Index(fields=["organization"]),
            models.Index(fields=["user"]),
        ]


class Custody(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    custodian = models.ForeignKey("TeamMember", on_delete=models.RESTRICT, related_name="custodies", db_column="teamMemberId", db_index=False)
    asset = models.ForeignKey("Asset", on_delete=models.CASCADE, related_name="custody", db_column="assetId", db_index=False)
    kit_custody = models.ForeignKey("KitCustody", on_delete=models.CASCADE, related_name="inherited_custody", db_column="kitCustodyId", null=True, db_index=False)
    quantity = models.IntegerField(default=1)
    location = models.ForeignKey("Location", on_delete=models.SET_NULL, related_name="custodies", db_column="locationId", null=True, db_index=False)
    source_unknown = models.BooleanField(db_column="sourceUnknown", default=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Custody"
        indexes = [
            models.Index(fields=["asset", "custodian"]),
            models.Index(fields=["custodian"]),
            models.Index(fields=["kit_custody"]),
            models.Index(fields=["location"]),
        ]


class Organization(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    name = models.TextField(default="Personal")
    type = models.TextField(choices=OrganizationType.choices, default=OrganizationType.PERSONAL)
    owner = models.ForeignKey("User", on_delete=models.CASCADE, related_name="organizations", db_column="userId", db_index=False)
    currency = models.TextField(choices=Currency.choices, default=Currency.USD)
    image = models.OneToOneField("Image", on_delete=models.SET_NULL, related_name="organization", db_column="imageId", null=True)
    enabled_sso = models.BooleanField(db_column="enabledSso", default=False)
    sso_details = models.ForeignKey("SsoDetails", on_delete=models.SET_NULL, related_name="organizations", db_column="ssoDetailsId", null=True, db_index=False)
    self_service_can_see_custody = models.BooleanField(db_column="selfServiceCanSeeCustody", default=False)
    self_service_can_see_bookings = models.BooleanField(db_column="selfServiceCanSeeBookings", default=False)
    base_user_can_see_custody = models.BooleanField(db_column="baseUserCanSeeCustody", default=False)
    base_user_can_see_bookings = models.BooleanField(db_column="baseUserCanSeeBookings", default=False)
    managers_can_edit_assets = models.BooleanField(db_column="managersCanEditAssets", default=False)
    barcodes_enabled = models.BooleanField(db_column="barcodesEnabled", default=False)
    barcodes_enabled_at = UtcDateTimeField(null=True, db_column="barcodesEnabledAt")
    used_barcode_trial = models.BooleanField(db_column="usedBarcodeTrial", default=False)
    audits_enabled = models.BooleanField(db_column="auditsEnabled", default=False)
    audits_enabled_at = UtcDateTimeField(null=True, db_column="auditsEnabledAt")
    used_audit_trial = models.BooleanField(db_column="usedAuditTrial", default=False)
    workspace_disabled = models.BooleanField(db_column="workspaceDisabled", default=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    has_sequential_ids_migrated = models.BooleanField(db_column="hasSequentialIdsMigrated", default=False)
    qr_id_display_preference = models.TextField(choices=QrIdDisplayPreference.choices, db_column="qrIdDisplayPreference", default=QrIdDisplayPreference.QR_ID)
    show_shelf_branding = models.BooleanField(db_column="showShelfBranding", default=True)
    show_qr_codes_on_pdfs = models.BooleanField(db_column="showQrCodesOnPdfs", default=True)
    custom_email_footer = models.CharField(max_length=500, null=True, db_column="customEmailFooter")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Organization"
        indexes = [
            models.Index(fields=["owner"]),
            models.Index(fields=["sso_details"]),
        ]


class UserOrganization(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="user_organizations", db_column="userId", db_index=False)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="user_organizations", db_column="organizationId", db_index=False)
    roles = EnumArrayField(models.TextField(choices=OrganizationRoles.choices), db_enum="\"OrganizationRoles\"")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    calendar_token_id = models.TextField(null=True, unique=True, db_column="calendarTokenId")
    """Secret, unguessable token for this member's iCal calendar-feed subscription
    in this workspace. Null until the user generates a feed; regenerating it
    rotates the value (invalidating any previously-shared subscription URL).
    `@unique` doubles as the lookup index for the public feed route.
    """
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "UserOrganization"
        constraints = [
            models.UniqueConstraint(fields=["user", "organization"], name="UserOrganization_userId_organizationId_key"),
        ]
        indexes = [
            models.Index(fields=["organization"]),
        ]


class SsoDetails(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    domain = models.TextField()
    base_user_group_id = models.TextField(null=True, db_column="baseUserGroupId")
    self_service_group_id = models.TextField(null=True, db_column="selfServiceGroupId")
    admin_group_id = models.TextField(null=True, db_column="adminGroupId")
    manager_group_id = models.TextField(null=True, db_column="managerGroupId")
    require_sso_login = models.BooleanField(db_column="requireSsoLogin", default=True)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "SsoDetails"


class ScimToken(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    token_hash = models.TextField(unique=True, db_column="tokenHash")
    label = models.TextField()
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="scim_tokens", db_column="organizationId", db_index=False)
    created_by_id = models.TextField(db_column="createdById")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    last_used_at = UtcDateTimeField(null=True, db_column="lastUsedAt")

    class Meta:
        db_table = "ScimToken"
        indexes = [
            models.Index(fields=["organization"]),
        ]


class UserScimExternalId(models.Model):
    """*
    * Org-scoped SCIM external ID linking a Shelf user to an identity-provider
    * object within a specific organization.  A single Shelf user can appear in
    * multiple organizations, each managed by a different identity provider, so the
    * external ID must be scoped to the (user, organization) pair.
    """

    id = models.TextField(primary_key=True, default=cuid_str)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="scim_external_ids", db_column="userId", db_index=False)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="scim_external_ids", db_column="organizationId", db_index=False)
    scim_external_id = models.TextField(db_column="scimExternalId")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "UserScimExternalId"
        constraints = [
            models.UniqueConstraint(fields=["user", "organization"], name="UserScimExternalId_userId_organizationId_key"),
            models.UniqueConstraint(fields=["organization", "scim_external_id"], name="UserScimExternalId_organizationId_scimExternalId_key"),
        ]
        indexes = [
            models.Index(fields=["organization"]),
        ]


class Tier(models.Model):
    id = models.TextField(choices=TierId.choices, primary_key=True, default=None)
    name = models.TextField()
    tier_limit = models.OneToOneField("TierLimit", on_delete=models.SET_NULL, related_name="tier", db_column="tierLimitId", null=True)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Tier"


class TierLimit(models.Model):
    id = models.TextField(choices=TierId.choices, primary_key=True, default=None)
    can_import_assets = models.BooleanField(db_column="canImportAssets", default=False)
    can_export_assets = models.BooleanField(db_column="canExportAssets", default=False)
    can_import_nrm = models.BooleanField(db_column="canImportNRM", default=False)
    can_hide_shelf_branding = models.BooleanField(db_column="canHideShelfBranding", default=False)
    can_use_reports = models.BooleanField(db_column="canUseReports", default=False)
    can_use_recurring_reminders = models.BooleanField(db_column="canUseRecurringReminders", default=False)
    can_use_push_notifications = models.BooleanField(db_column="canUsePushNotifications", default=False)
    max_custom_fields = models.IntegerField(db_column="maxCustomFields", default=0)
    max_organizations = models.IntegerField(db_column="maxOrganizations", default=1)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "TierLimit"


class CustomTierLimit(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    user = models.OneToOneField("User", on_delete=models.SET_NULL, related_name="custom_tier_limit", db_column="userId", null=True)
    can_import_assets = models.BooleanField(db_column="canImportAssets", default=True)
    can_export_assets = models.BooleanField(db_column="canExportAssets", default=True)
    can_import_nrm = models.BooleanField(db_column="canImportNRM", default=True)
    can_hide_shelf_branding = models.BooleanField(db_column="canHideShelfBranding", default=True)
    can_use_reports = models.BooleanField(db_column="canUseReports", default=True)
    can_use_recurring_reminders = models.BooleanField(db_column="canUseRecurringReminders", default=True)
    can_use_push_notifications = models.BooleanField(db_column="canUsePushNotifications", default=True)
    max_custom_fields = models.IntegerField(db_column="maxCustomFields", default=1000)
    max_organizations = models.IntegerField(db_column="maxOrganizations", default=1)
    is_enterprise = models.BooleanField(db_column="isEnterprise", default=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "CustomTierLimit"


class CustomField(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    name = models.TextField()
    help_text = models.TextField(null=True, db_column="helpText")
    required = models.BooleanField(default=False)
    active = models.BooleanField(default=True)
    type = models.TextField(choices=CustomFieldType.choices, default=CustomFieldType.TEXT)
    options = ArrayField(models.TextField())
    format = models.TextField(choices=CustomFieldFormat.choices, default=CustomFieldFormat.PLAIN)
    """How a TEXT field's values are read. LINK requires a web link in every
    value; any other type stays PLAIN (enforced by a CHECK constraint).
    """
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="custom_fields", db_column="organizationId", db_index=False)
    created_by = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="custom_fields", db_column="userId", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    deleted_at = UtcDateTimeField(null=True, db_column="deletedAt")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "CustomField"
        indexes = [
            models.Index(fields=["organization"]),
            models.Index(fields=["created_by"]),
            models.Index(fields=["organization", "deleted_at"]),
        ]


class AssetCustomFieldValue(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    value = models.JSONField()
    asset = models.ForeignKey("Asset", on_delete=models.CASCADE, related_name="custom_fields", db_column="assetId", db_index=False)
    custom_field = models.ForeignKey("CustomField", on_delete=models.CASCADE, related_name="asset_custom_fields_values", db_column="customFieldId", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "AssetCustomFieldValue"
        indexes = [
            models.Index(fields=["asset", "custom_field"]),
            models.Index(fields=["custom_field"]),
        ]


class AssetModelCustomFieldDefault(models.Model):
    """A custom field value an AssetModel gives the assets that join it.

    `value` has the exact shape of `AssetCustomFieldValue.value`, so filling an
    asset is a plain copy. Rows cascade away with their model or field.
    """

    id = models.TextField(primary_key=True, default=cuid_str)
    value = models.JSONField()
    asset_model = models.ForeignKey("AssetModel", on_delete=models.CASCADE, related_name="custom_field_defaults", db_column="assetModelId", db_index=False)
    custom_field_id = models.TextField(db_column="customFieldId")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "AssetModelCustomFieldDefault"
        constraints = [
            models.UniqueConstraint(fields=["asset_model", "custom_field_id"], name="AssetModelCustomFieldDefault_assetModelId_customFieldId_key"),
        ]
        indexes = [
            models.Index(fields=["custom_field_id"]),
        ]


class Invite(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    inviter = models.ForeignKey("User", on_delete=models.CASCADE, related_name="sent_invites", db_column="inviterId", db_index=False)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="invites", db_column="organizationId", db_index=False)
    invitee_user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="received_invites", db_column="inviteeUserId", null=True, db_index=False)
    invitee_team_member = models.ForeignKey("TeamMember", on_delete=models.RESTRICT, related_name="received_invites", db_column="teamMemberId", db_index=False)
    invitee_email = models.TextField(db_column="inviteeEmail")
    status = models.TextField(choices=InviteStatuses.choices, default=InviteStatuses.PENDING)
    invite_code = models.TextField(db_column="inviteCode")
    roles = EnumArrayField(models.TextField(choices=OrganizationRoles.choices), db_enum="\"OrganizationRoles\"")
    invite_message = models.CharField(max_length=1000, null=True, db_column="inviteMessage")
    expires_at = UtcDateTimeField(db_column="expiresAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Invite"
        indexes = [
            models.Index(fields=["invitee_user"]),
            models.Index(fields=["inviter"]),
            models.Index(fields=["organization"]),
            models.Index(fields=["invitee_team_member"]),
        ]


class Announcement(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    name = models.TextField()
    content = models.TextField()
    link = models.TextField(null=True)
    link_text = models.TextField(null=True, db_column="linkText")
    published = models.BooleanField(default=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Announcement"


class ConsumptionLog(models.Model):
    """*
    * Audit trail for all quantity changes on quantity-tracked assets.
    * Every event records who performed it, who received/returned items,
    * which booking it relates to, and the category of change.
    * Quantity is always positive — direction is determined by category.
    """

    id = models.TextField(primary_key=True, default=cuid_str)
    asset = models.ForeignKey("Asset", on_delete=models.CASCADE, related_name="consumption_logs", db_column="assetId", db_index=False)
    category = models.TextField(choices=ConsumptionCategory.choices)
    quantity = models.IntegerField()
    note = models.TextField(null=True)
    performed_by = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="consumption_logs", db_column="userId", db_index=False)
    booking = models.ForeignKey("Booking", on_delete=models.SET_NULL, related_name="consumption_logs", db_column="bookingId", null=True, db_index=False)
    custodian = models.ForeignKey("TeamMember", on_delete=models.SET_NULL, related_name="consumption_logs", db_column="custodianId", null=True, db_index=False)
    booking_asset_id = models.TextField(null=True, db_column="bookingAssetId")
    location = models.ForeignKey("Location", on_delete=models.SET_NULL, related_name="consumption_logs", db_column="locationId", null=True, db_index=False)
    stock_change = models.IntegerField(null=True, db_column="stockChange")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "ConsumptionLog"
        indexes = [
            models.Index(fields=["asset", "created_at"]),
            models.Index(fields=["performed_by"]),
            models.Index(fields=["booking"]),
            models.Index(fields=["custodian"]),
            models.Index(fields=["booking_asset_id"]),
            models.Index(fields=["location"]),
        ]


class BookingAsset(models.Model):
    """*
    * Explicit pivot table between Booking and Asset.
    * Replaces the implicit M2M (_AssetToBooking) via a rename migration in Phase 3.
    """

    id = models.TextField(primary_key=True, default=cuid_str)
    booking = models.ForeignKey("Booking", on_delete=models.CASCADE, related_name="booking_assets", db_column="bookingId", db_index=False)
    asset = models.ForeignKey("Asset", on_delete=models.CASCADE, related_name="booking_assets", db_column="assetId", db_index=False)
    quantity = models.IntegerField(default=1)
    booking_model_request_id = models.TextField(null=True, db_column="bookingModelRequestId")
    asset_kit_id = models.TextField(null=True, db_column="assetKitId")
    source_kit_id = models.TextField(null=True, db_column="sourceKitId")
    checked_out_at = models.DateTimeField(null=True, db_column="checkedOutAt")
    checked_out_by_id = models.TextField(null=True, db_column="checkedOutById")
    checked_in_at = models.DateTimeField(null=True, db_column="checkedInAt")
    checked_in_by_id = models.TextField(null=True, db_column="checkedInById")
    checked_out_quantity = models.IntegerField(db_column="checkedOutQuantity", default=0)
    source_location_id = models.TextField(null=True, db_column="sourceLocationId")

    class Meta:
        db_table = "BookingAsset"
        indexes = [
            models.Index(fields=["asset"]),
            models.Index(fields=["asset_kit_id"]),
            models.Index(fields=["source_kit_id"]),
            models.Index(fields=["source_location_id"]),
        ]


class Booking(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    name = models.TextField()
    status = models.TextField(choices=BookingStatus.choices, default=BookingStatus.DRAFT)
    description = models.TextField(null=True, default="")
    active_scheduler_reference = models.TextField(null=True, db_column="activeSchedulerReference")
    creator = models.ForeignKey("User", on_delete=models.CASCADE, related_name="bookings", db_column="creatorId", db_index=False)
    custodian_user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="custodies", db_column="custodianUserId", null=True, db_index=False)
    custodian_team_member = models.ForeignKey("TeamMember", on_delete=models.CASCADE, related_name="bookings", db_column="custodianTeamMemberId", null=True, db_index=False)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="bookings", db_column="organizationId", db_index=False)
    tags: "models.ManyToManyField[Tag, BookingToTag]" = models.ManyToManyField("Tag", through="BookingToTag", related_name="bookings")
    created_at = models.DateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = AutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    from_field = models.DateTimeField(db_column="from")
    to = models.DateTimeField()
    original_from = models.DateTimeField(null=True, db_column="originalFrom")
    original_to = models.DateTimeField(null=True, db_column="originalTo")
    auto_archived_at = models.DateTimeField(null=True, db_column="autoArchivedAt")
    archived_without_checkin = models.BooleanField(db_column="archivedWithoutCheckin", default=False)
    cancellation_reason = models.TextField(null=True, db_column="cancellationReason")
    notification_recipients: "models.ManyToManyField[TeamMember, BookingNotificationRecipients]" = models.ManyToManyField("TeamMember", through="BookingNotificationRecipients", related_name="booking_notifications")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Booking"
        indexes = [
            models.Index(fields=["creator"]),
            models.Index(fields=["custodian_team_member"]),
            models.Index(fields=["custodian_user"]),
            models.Index(fields=["organization"]),
            models.Index(fields=["status"]),
            models.Index(fields=["organization", "from_field", "id"]),
        ]


class BookingSettings(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    buffer_start_time = models.IntegerField(db_column="bufferStartTime", default=0)
    tags_required = models.BooleanField(db_column="tagsRequired", default=False)
    max_booking_length = models.IntegerField(null=True, db_column="maxBookingLength")
    max_booking_length_skip_closed_days = models.BooleanField(db_column="maxBookingLengthSkipClosedDays", default=False)
    auto_archive_bookings = models.BooleanField(db_column="autoArchiveBookings", default=False)
    auto_archive_days = models.IntegerField(db_column="autoArchiveDays", default=2)
    auto_archive_expired_reservations = models.BooleanField(db_column="autoArchiveExpiredReservations", default=False)
    require_explicit_checkin_for_admin = models.BooleanField(db_column="requireExplicitCheckinForAdmin", default=False)
    require_explicit_checkin_for_self_service = models.BooleanField(db_column="requireExplicitCheckinForSelfService", default=False)
    require_explicit_checkout_for_admin = models.BooleanField(db_column="requireExplicitCheckoutForAdmin", default=False)
    require_explicit_checkout_for_self_service = models.BooleanField(db_column="requireExplicitCheckoutForSelfService", default=False)
    count_kits_as_single_unit = models.BooleanField(db_column="countKitsAsSingleUnit", default=False)
    notify_booking_creator = models.BooleanField(db_column="notifyBookingCreator", default=True)
    notify_admins_on_new_booking = models.BooleanField(db_column="notifyAdminsOnNewBooking", default=True)
    always_notify_team_members: "models.ManyToManyField[TeamMember, BookingSettingsAlwaysNotify]" = models.ManyToManyField("TeamMember", through="BookingSettingsAlwaysNotify", related_name="always_notify_booking_settings")
    organization = models.OneToOneField("Organization", on_delete=models.CASCADE, related_name="booking_settings", db_column="organizationId")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "BookingSettings"
        indexes = [
            models.Index(fields=["organization"]),
        ]


class PartialBookingCheckin(models.Model):
    """A progressive check-in SESSION, mirroring PartialBookingCheckout.

    `assetIds` lists only FULLY reconciled assets: INDIVIDUAL always, and
    QUANTITY_TRACKED only once its remaining units hit zero (partial returns
    live in `ConsumptionLog`). Completing a booking outright writes no row here,
    so the same warning applies — absence is not evidence of anything.
    `BookingAsset.checkedInAt` is the per-slice fact.
    """

    id = models.TextField(primary_key=True, default=cuid_str)
    asset_ids = ArrayField(models.TextField(), db_column="assetIds")
    checkin_count = models.IntegerField(db_column="checkinCount")
    checkin_timestamp = models.DateTimeField(db_column="checkinTimestamp", default=now_milliseconds)
    booking = models.ForeignKey("Booking", on_delete=models.CASCADE, related_name="partial_checkins", db_column="bookingId", db_index=False)
    checked_in_by = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="partial_checkins", db_column="checkedInById", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "PartialBookingCheckin"
        indexes = [
            models.Index(fields=["booking"]),
            models.Index(fields=["checked_in_by"]),
            models.Index(fields=["checkin_timestamp"]),
            models.Index(fields=["booking", "checkin_timestamp"]),
        ]


class PartialBookingCheckout(models.Model):
    """A progressive checkout SESSION: one scan batch, covering N assets via the
    positional `assetIds` / `quantities` / `bookingAssetIds` arrays.

    Records sessions and the units they claimed — NOT whether an asset is out.
    The all-at-once checkout (`checkoutBooking`, the Check out button) flips
    statuses and writes NO row here, so **absence proves nothing** and must never
    be read as "was not checked out". Reading it that way is what let a single
    later scan mark every button-checked-out asset on a booking as
    "never checked out". `BookingAsset.checkedOutAt` answers that question.
    """

    id = models.TextField(primary_key=True, default=cuid_str)
    asset_ids = ArrayField(models.TextField(), db_column="assetIds")
    quantities = ArrayField(models.IntegerField(), default=list)
    booking_asset_ids = ArrayField(models.TextField(), db_column="bookingAssetIds", default=list)
    checkout_count = models.IntegerField(db_column="checkoutCount")
    checkout_timestamp = models.DateTimeField(db_column="checkoutTimestamp", default=now_milliseconds)
    booking = models.ForeignKey("Booking", on_delete=models.CASCADE, related_name="partial_checkouts", db_column="bookingId", db_index=False)
    checked_out_by = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="partial_checkouts", db_column="checkedOutById", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "PartialBookingCheckout"
        indexes = [
            models.Index(fields=["booking"]),
            models.Index(fields=["checked_out_by"]),
            models.Index(fields=["checkout_timestamp"]),
            models.Index(fields=["booking", "checkout_timestamp"]),
        ]


class MobileAuthCode(models.Model):
    """Native-app SSO web-delegated login — single-use session-handoff codes.
    @see apps/webapp/app/routes/_auth+/oauth.callback.mobile.tsx
    """

    id = models.TextField(primary_key=True, default=cuid_str)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="mobile_auth_codes", db_column="userId", db_index=False)
    code_hash = models.TextField(unique=True, db_column="codeHash")
    code_challenge = models.TextField(null=True, db_column="codeChallenge")
    session_ciphertext = models.TextField(null=True, db_column="sessionCiphertext")
    consumed_at = UtcDateTimeField(null=True, db_column="consumedAt")
    expires_at = UtcDateTimeField(db_column="expiresAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "MobileAuthCode"
        indexes = [
            models.Index(fields=["user"]),
            models.Index(fields=["expires_at"]),
        ]


class PushToken(models.Model):
    """One Expo push token per installed companion app. A token identifies a
    device, not a person: when someone else signs in on the same phone, the
    row moves to them. Pushes go only to rows where `disabledAt` is null.
    @see apps/webapp/app/modules/push/service.server.ts
    """

    id = models.TextField(primary_key=True, default=cuid_str)
    token = models.TextField(unique=True)
    platform = models.TextField(choices=DevicePlatform.choices)
    device_name = models.TextField(null=True, db_column="deviceName")
    app_version = models.TextField(null=True, db_column="appVersion")
    last_seen_at = UtcDateTimeField(db_column="lastSeenAt")
    disabled_at = UtcDateTimeField(null=True, db_column="disabledAt")
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="push_tokens", db_column="userId", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "PushToken"
        indexes = [
            models.Index(fields=["user"]),
        ]


class NotificationPreference(models.Model):
    """A person's push notification switches. A missing row means every switch
    is on, so no backfill is needed. The three category columns are named
    after `PushCategory` in apps/webapp/app/modules/push/types.ts.
    """

    user = models.OneToOneField("User", on_delete=models.CASCADE, related_name="notification_preference", primary_key=True, db_column="userId")
    push_enabled = models.BooleanField(db_column="pushEnabled", default=True)
    bookings = models.BooleanField(default=True)
    audits = models.BooleanField(default=True)
    stock_and_reminders = models.BooleanField(db_column="stockAndReminders", default=True)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "NotificationPreference"


class Kit(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    name = models.TextField()
    description = models.TextField(null=True)
    status = models.TextField(choices=KitStatus.choices, default=KitStatus.AVAILABLE)
    image = models.TextField(null=True)
    image_expiration = UtcDateTimeField(null=True, db_column="imageExpiration")
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="kits", db_column="organizationId", db_index=False)
    created_by = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="created_kits", db_column="createdById", db_index=False)
    category = models.ForeignKey("Category", on_delete=models.SET_NULL, related_name="kits", db_column="categoryId", null=True, db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    location = models.ForeignKey("Location", on_delete=models.SET_NULL, related_name="kits", db_column="locationId", null=True, db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Kit"
        indexes = [
            models.Index(fields=["created_by"]),
            models.Index(fields=["organization"]),
            models.Index(fields=["category", "organization"]),
            models.Index(fields=["category", "organization", "created_at"]),
            models.Index(fields=["category", "organization", "name"]),
            models.Index(fields=["category", "organization", "status"]),
        ]


class KitCustody(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    custodian = models.ForeignKey("TeamMember", on_delete=models.RESTRICT, related_name="kit_custodies", db_column="custodianId", db_index=False)
    kit = models.OneToOneField("Kit", on_delete=models.CASCADE, related_name="custody", db_column="kitId")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "KitCustody"
        indexes = [
            models.Index(fields=["custodian"]),
        ]


class AssetKit(models.Model):
    """Pivot linking an `Asset` to the `Kit` it belongs to.

    `@@unique([assetId])` enforces "at most one kit per asset" today.
    The structure already allows multi-row-per-asset; if multi-kit
    allocation lands later, drop the `assetId` unique constraint and
    add a type-aware single-row trigger + sum-within-total CONSTRAINT
    TRIGGER (parallel to the Custody / BookingAsset pivots).
    """

    id = models.TextField(primary_key=True, default=cuid_str)
    asset = models.ForeignKey("Asset", on_delete=models.CASCADE, related_name="asset_kits", db_column="assetId", db_index=False)
    kit = models.ForeignKey("Kit", on_delete=models.CASCADE, related_name="asset_kits", db_column="kitId", db_index=False)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="asset_kits", db_column="organizationId", db_index=False)
    quantity = models.IntegerField(default=1)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "AssetKit"
        constraints = [
            models.UniqueConstraint(fields=["asset", "kit"], name="AssetKit_assetId_kitId_key"),
        ]
        indexes = [
            models.Index(fields=["kit"]),
            models.Index(fields=["organization"]),
        ]


class AssetLocation(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    asset = models.ForeignKey("Asset", on_delete=models.CASCADE, related_name="asset_locations", db_column="assetId", db_index=False)
    location = models.ForeignKey("Location", on_delete=models.CASCADE, related_name="asset_locations", db_column="locationId", db_index=False)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="asset_locations", db_column="organizationId", db_index=False)
    quantity = models.IntegerField(default=1)
    asset_kit = models.ForeignKey("AssetKit", on_delete=models.CASCADE, related_name="asset_locations", db_column="assetKitId", null=True, db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "AssetLocation"
        indexes = [
            models.Index(fields=["location"]),
            models.Index(fields=["organization"]),
            models.Index(fields=["asset"]),
            models.Index(fields=["asset_kit"]),
        ]


class AssetReminder(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    name = models.TextField()
    message = models.TextField()
    alert_date_time = UtcDateTimeField(db_column="alertDateTime")
    active_scheduler_reference = models.TextField(null=True, db_column="activeSchedulerReference")
    recurrence_unit = models.TextField(choices=ReminderRecurrenceUnit.choices, null=True, db_column="recurrenceUnit")
    recurrence_interval = models.IntegerField(null=True, db_column="recurrenceInterval")
    recurrence_timezone = models.TextField(null=True, db_column="recurrenceTimezone")
    recurrence_ends_at = UtcDateTimeField(null=True, db_column="recurrenceEndsAt")
    organization = models.ForeignKey("Organization", on_delete=models.RESTRICT, related_name="asset_reminders", db_column="organizationId", db_index=False)
    asset = models.ForeignKey("Asset", on_delete=models.CASCADE, related_name="reminders", db_column="assetId", db_index=False)
    team_members: "models.ManyToManyField[TeamMember, AssetReminderToTeamMember]" = models.ManyToManyField("TeamMember", through="AssetReminderToTeamMember", related_name="asset_reminders")
    created_by = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="asset_reminders", db_column="createdById", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "AssetReminder"
        indexes = [
            models.Index(fields=["asset", "alert_date_time"]),
            GinIndex(fields=["name", "message"]),
            models.Index(fields=["organization", "alert_date_time", "asset"]),
            models.Index(fields=["alert_date_time", "active_scheduler_reference"]),
            models.Index(fields=["created_by"]),
        ]


class WorkingHours(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    enabled = models.BooleanField(default=False)
    weekly_schedule = models.JSONField(db_column="weeklySchedule", default=working_hours_weekly_schedule_default)
    organization = models.OneToOneField("Organization", on_delete=models.CASCADE, related_name="working_hours", db_column="organizationId")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "WorkingHours"
        indexes = [
            models.Index(fields=["organization"]),
        ]


class WorkingHoursOverride(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    date = models.DateField()
    is_open = models.BooleanField(db_column="isOpen", default=False)
    open_time = models.TextField(null=True, db_column="openTime")
    close_time = models.TextField(null=True, db_column="closeTime")
    reason = models.TextField(null=True)
    working_hours = models.ForeignKey("WorkingHours", on_delete=models.CASCADE, related_name="overrides", db_column="workingHoursId", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "WorkingHoursOverride"
        constraints = [
            models.UniqueConstraint(fields=["working_hours", "date"], name="WorkingHoursOverride_workingHoursId_date_key"),
        ]
        indexes = [
            models.Index(fields=["working_hours", "date"]),
            models.Index(fields=["date", "is_open"]),
        ]


class Update(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    title = models.TextField()
    content = models.TextField()
    url = models.TextField(null=True)
    image_url = models.TextField(null=True, db_column="imageUrl")
    publish_date = UtcDateTimeField(db_column="publishDate")
    status = models.TextField(choices=UpdateStatus.choices, default=UpdateStatus.DRAFT)
    target_roles = EnumArrayField(models.TextField(choices=OrganizationRoles.choices), db_enum="\"OrganizationRoles\"", db_column="targetRoles")
    click_count = models.IntegerField(db_column="clickCount", default=0)
    view_count = models.IntegerField(db_column="viewCount", default=0)
    created_by = models.ForeignKey("User", on_delete=models.CASCADE, related_name="created_updates", db_column="createdById", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Update"
        indexes = [
            models.Index(fields=["status", "publish_date"]),
            models.Index(fields=["publish_date"]),
            models.Index(fields=["created_by"]),
        ]


class UserUpdateRead(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="update_reads", db_column="userId", db_index=False)
    update = models.ForeignKey("Update", on_delete=models.CASCADE, related_name="user_reads", db_column="updateId", db_index=False)
    read_at = UtcDateTimeField(db_column="readAt", default=now_milliseconds)

    class Meta:
        db_table = "UserUpdateRead"
        constraints = [
            models.UniqueConstraint(fields=["user", "update"], name="UserUpdateRead_userId_updateId_key"),
        ]
        indexes = [
            models.Index(fields=["user"]),
            models.Index(fields=["update"]),
            models.Index(fields=["read_at"]),
        ]


class AuditNote(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    content = models.TextField()
    type = models.TextField(choices=NoteType.choices, default=NoteType.COMMENT)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    user = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="audit_notes", db_column="userId", null=True, db_index=False)
    audit_session = models.ForeignKey("AuditSession", on_delete=models.CASCADE, related_name="notes", db_column="auditSessionId", db_index=False)
    audit_asset = models.ForeignKey("AuditAsset", on_delete=models.SET_NULL, related_name="notes", db_column="auditAssetId", null=True, db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "AuditNote"
        indexes = [
            models.Index(fields=["audit_session"]),
            models.Index(fields=["user"]),
            models.Index(fields=["audit_asset"]),
        ]


class TeamMemberNote(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    content = models.TextField()
    type = models.TextField(choices=NoteType.choices, default=NoteType.COMMENT)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    user = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="team_member_notes_authored", db_column="userId", null=True, db_index=False)
    team_member = models.ForeignKey("TeamMember", on_delete=models.CASCADE, related_name="team_member_notes", db_column="teamMemberId", db_index=False)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="team_member_notes", db_column="organizationId", db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "TeamMemberNote"
        indexes = [
            models.Index(fields=["team_member"]),
            models.Index(fields=["user"]),
            models.Index(fields=["organization"]),
            models.Index(fields=["team_member", "organization"]),
        ]


class AuditSession(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    name = models.TextField()
    description = models.TextField(null=True)
    target_id = models.TextField(null=True, db_column="targetId")
    status = models.TextField(choices=AuditStatus.choices, default=AuditStatus.PENDING)
    scope_meta = models.JSONField(null=True, db_column="scopeMeta")
    expected_asset_count = models.IntegerField(db_column="expectedAssetCount", default=0)
    found_asset_count = models.IntegerField(db_column="foundAssetCount", default=0)
    missing_asset_count = models.IntegerField(db_column="missingAssetCount", default=0)
    unexpected_asset_count = models.IntegerField(db_column="unexpectedAssetCount", default=0)
    started_at = UtcDateTimeField(null=True, db_column="startedAt")
    due_date = UtcDateTimeField(null=True, db_column="dueDate")
    completed_at = UtcDateTimeField(null=True, db_column="completedAt")
    cancelled_at = UtcDateTimeField(null=True, db_column="cancelledAt")
    active_scheduler_reference = models.TextField(null=True, db_column="activeSchedulerReference")
    created_by = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="audits_created", db_column="createdById", db_index=False)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="audit_sessions", db_column="organizationId", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "AuditSession"
        indexes = [
            models.Index(fields=["organization", "status"]),
            models.Index(fields=["created_by"]),
            models.Index(fields=["status", "created_at"]),
        ]


class AssetImport(models.Model):
    """One CSV content import. Holds the parsed rows from preview until the import
    completes, so the job imports exactly the file the user previewed.
    A partial unique index on (organizationId) WHERE status IN (QUEUED, RUNNING) allows one
    active import per workspace. Prisma cannot express it; it lives in migration
    20261015120100_asset_import_one_active_per_workspace.
    """

    id = models.TextField(primary_key=True, default=cuid_str)
    status = models.TextField(choices=AssetImportStatus.choices, default=AssetImportStatus.DRAFT)
    phase = models.TextField(choices=AssetImportPhase.choices, default=AssetImportPhase.PREPARE)
    file_name = models.TextField(null=True, db_column="fileName")
    total_rows = models.IntegerField(db_column="totalRows")
    processed_rows = models.IntegerField(db_column="processedRows", default=0)
    """Rows committed so far. Advanced in the same transaction as each chunk."""
    rows = models.JSONField(null=True)
    """Parsed rows (`StoredImportRow[]`). Cleared when the import completes."""
    resolved = models.JSONField(null=True)
    """Name to id maps written by the PREPARE step (`ResolvedImportEntities`)."""
    can_use_barcodes = models.BooleanField(db_column="canUseBarcodes", default=False)
    error_message = models.TextField(null=True, db_column="errorMessage")
    """User-facing reason the import failed. Null unless status is FAILED."""
    warnings = models.JSONField(null=True)
    """Rows imported without a conflicting piece (`ImportWarning[]`), capped at 500."""
    warning_count = models.IntegerField(db_column="warningCount", default=0)
    """Exact number of warnings, including any beyond the stored cap."""
    locked_until = UtcDateTimeField(null=True, db_column="lockedUntil")
    """A step holds the import until this time. Null when no step is running."""
    confirmed_at = UtcDateTimeField(null=True, db_column="confirmedAt")
    started_at = UtcDateTimeField(null=True, db_column="startedAt")
    completed_at = UtcDateTimeField(null=True, db_column="completedAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    created_by = models.ForeignKey("User", on_delete=models.CASCADE, related_name="asset_imports", db_column="createdById", db_index=False)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="asset_imports", db_column="organizationId", db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "AssetImport"
        indexes = [
            models.Index(fields=["organization", "status"]),
            models.Index(fields=["created_by", "organization", "status"]),
        ]


class LinkPreview(models.Model):
    """What Shelf knows about one web link used in a workspace's custom fields,
    keyed by its `LinkKey` from @shelf/links. Two writers on disjoint columns:
    the link preview job (fetch columns) and people (customTitle).
    """

    id = models.TextField(primary_key=True, default=cuid_str)
    url = models.TextField()
    """`LinkKey` from @shelf/links, never a raw field value. At most 2048
    characters (CHECK constraint in the migration).
    """
    page_title = models.TextField(null=True, db_column="pageTitle")
    outcome = models.TextField(choices=LinkPreviewOutcome.choices, null=True)
    """Null until the first fetch finishes."""
    attempts = models.IntegerField(default=0)
    fetched_at = UtcDateTimeField(null=True, db_column="fetchedAt")
    next_fetch_at = UtcDateTimeField(null=True, db_column="nextFetchAt", default=now_milliseconds)
    """When the link may be fetched next; null means never again. A claim moves
    it a few minutes ahead, which is also the lease on a fetch in flight.
    """
    custom_title = models.TextField(null=True, db_column="customTitle")
    """The name a person gave the link in this workspace; wins over pageTitle."""
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="link_previews", db_column="organizationId", db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "LinkPreview"
        constraints = [
            models.UniqueConstraint(fields=["organization", "url"], name="LinkPreview_organizationId_url_key"),
        ]


class AuditAssignment(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    audit_session = models.ForeignKey("AuditSession", on_delete=models.CASCADE, related_name="assignments", db_column="auditSessionId", db_index=False)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="audit_assignments", db_column="userId", db_index=False)
    role = models.TextField(choices=AuditAssignmentRole.choices, null=True)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "AuditAssignment"
        constraints = [
            models.UniqueConstraint(fields=["audit_session", "user"], name="AuditAssignment_auditSessionId_userId_key"),
        ]
        indexes = [
            models.Index(fields=["user"]),
        ]


class AuditAsset(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    audit_session = models.ForeignKey("AuditSession", on_delete=models.CASCADE, related_name="assets", db_column="auditSessionId", db_index=False)
    asset = models.ForeignKey("Asset", on_delete=models.CASCADE, related_name="audit_entries", db_column="assetId", db_index=False)
    expected = models.BooleanField(default=True)
    status = models.TextField(choices=AuditAssetStatus.choices, default=AuditAssetStatus.PENDING)
    scanned_by = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="audit_assets_scanned", db_column="scannedById", null=True, db_index=False)
    scanned_at = UtcDateTimeField(null=True, db_column="scannedAt")
    metadata = models.JSONField(null=True)
    expected_quantity = models.IntegerField(null=True, db_column="expectedQuantity")
    """Units the audit expects to find, snapshotted when the row is created: the
    units in the audit's scope (placed at its locations, held in its kits, or
    held by its person; at most `Asset.quantity`), or `Asset.quantity` when
    none are in scope or the audit has no scope. Null for INDIVIDUAL assets,
    which count as one.

    A snapshot, not a live read: stock that moves during the audit must not
    move the target the auditor is counting against.
    """
    found_quantity = models.IntegerField(null=True, db_column="foundQuantity")
    """Units the auditor counted. Null until a quantity-tracked asset is
    counted, and always null for INDIVIDUAL assets. A FOUND row with a count
    below `expectedQuantity` is a shortfall: the status stays FOUND (the pool
    was located), the gap is this number.

    Both columns are non-negative, enforced by CHECK constraints in raw SQL
    (migration `20261104120000_add_audit_asset_quantities`).
    """
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "AuditAsset"
        constraints = [
            models.UniqueConstraint(fields=["audit_session", "asset"], name="AuditAsset_auditSessionId_assetId_key"),
        ]
        indexes = [
            models.Index(fields=["status"]),
            models.Index(fields=["scanned_by"]),
            models.Index(fields=["asset"]),
        ]


class AuditScan(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    audit_session = models.ForeignKey("AuditSession", on_delete=models.CASCADE, related_name="scans", db_column="auditSessionId", db_index=False)
    audit_asset = models.ForeignKey("AuditAsset", on_delete=models.SET_NULL, related_name="scans", db_column="auditAssetId", null=True, db_index=False)
    asset = models.ForeignKey("Asset", on_delete=models.SET_NULL, related_name="audit_scans", db_column="assetId", null=True, db_index=False)
    scanned_by = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="audit_scans", db_column="scannedById", null=True, db_index=False)
    code = models.TextField(null=True)
    asset_title = models.TextField(null=True, db_column="assetTitle")
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
    was_expected = models.BooleanField(null=True, db_column="wasExpected")
    metadata = models.JSONField(null=True)
    scanned_at = UtcDateTimeField(db_column="scannedAt", default=now_milliseconds)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "AuditScan"
        indexes = [
            models.Index(fields=["audit_session", "scanned_at"]),
            models.Index(fields=["audit_asset"]),
            models.Index(fields=["asset"]),
        ]


class AuditImage(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    image_url = models.TextField(db_column="imageUrl")
    thumbnail_url = models.TextField(null=True, db_column="thumbnailUrl")
    description = models.TextField(null=True)
    audit_session = models.ForeignKey("AuditSession", on_delete=models.CASCADE, related_name="images", db_column="auditSessionId", db_index=False)
    audit_asset = models.ForeignKey("AuditAsset", on_delete=models.SET_NULL, related_name="images", db_column="auditAssetId", null=True, db_index=False)
    uploaded_by = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="audit_images_uploaded", db_column="uploadedById", null=True, db_index=False)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="audit_images", db_column="organizationId", db_index=False)
    captured_at = UtcDateTimeField(null=True, db_column="capturedAt")
    """When the photo was taken, from its EXIF data (or the phone app's copy of
    it). Null when the photo carried no usable date. EXIF can be edited, so
    this is the best signal available, not proof.
    """
    capture_source = models.TextField(choices=AuditImageCaptureSource.choices, null=True, db_column="captureSource")
    """How the photo reached Shelf. Null only for photos uploaded before this
    field existed; a phone upload with no or an unknown source is PHONE.
    """
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "AuditImage"
        indexes = [
            models.Index(fields=["audit_session"]),
            models.Index(fields=["audit_asset"]),
            models.Index(fields=["organization"]),
            models.Index(fields=["uploaded_by"]),
        ]


class RoleChangeLog(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    previous_role = models.TextField(choices=OrganizationRoles.choices, db_column="previousRole")
    new_role = models.TextField(choices=OrganizationRoles.choices, db_column="newRole")
    source = models.TextField(choices=RoleChangeSource.choices, default=RoleChangeSource.MANUAL)
    """How the change was made. SSO rows name the member themself as `changedBy`."""
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    user = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="role_changes_received", db_column="userId", db_index=False)
    changed_by = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="role_changes_performed", db_column="changedById", db_index=False)
    organization = models.ForeignKey("Organization", on_delete=models.RESTRICT, related_name="role_change_logs", db_column="organizationId", db_index=False)

    class Meta:
        db_table = "RoleChangeLog"
        indexes = [
            models.Index(fields=["user", "organization"]),
            models.Index(fields=["organization", "created_at"]),
        ]


class ActivityEvent(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="activity_events", db_column="organizationId", db_index=False)
    occurred_at = models.DateTimeField(db_column="occurredAt", default=now_milliseconds)
    actor_user_id = models.TextField(null=True, db_column="actorUserId")
    actor_snapshot = models.JSONField(null=True, db_column="actorSnapshot")
    action = models.TextField(choices=ActivityAction.choices)
    entity_type = models.TextField(choices=ActivityEntity.choices, db_column="entityType")
    entity_id = models.TextField(db_column="entityId")
    asset_id = models.TextField(null=True, db_column="assetId")
    booking_id = models.TextField(null=True, db_column="bookingId")
    audit_session_id = models.TextField(null=True, db_column="auditSessionId")
    audit_asset_id = models.TextField(null=True, db_column="auditAssetId")
    kit_id = models.TextField(null=True, db_column="kitId")
    location_id = models.TextField(null=True, db_column="locationId")
    team_member_id = models.TextField(null=True, db_column="teamMemberId")
    target_user_id = models.TextField(null=True, db_column="targetUserId")
    field = models.TextField(null=True)
    from_value = models.JSONField(null=True, db_column="fromValue")
    to_value = models.JSONField(null=True, db_column="toValue")
    meta = models.JSONField(null=True)

    class Meta:
        db_table = "ActivityEvent"
        indexes = [
            models.Index(fields=["organization", "occurred_at"]),
            models.Index(fields=["organization", "action", "occurred_at"]),
            models.Index(fields=["organization", "entity_type", "entity_id", "occurred_at"]),
            models.Index(fields=["actor_user_id", "occurred_at"]),
            models.Index(fields=["asset_id", "occurred_at"]),
            models.Index(fields=["booking_id", "occurred_at"]),
            models.Index(fields=["audit_session_id", "occurred_at"]),
        ]


class RoleToUser(models.Model):
    pk = models.CompositePrimaryKey("a_id", "b_id")
    a = models.ForeignKey("Role", on_delete=models.CASCADE, related_name="+", db_column="A", db_index=False)
    b = models.ForeignKey("User", on_delete=models.CASCADE, related_name="+", db_column="B", db_index=False)

    class Meta:
        db_table = "_RoleToUser"
        indexes = [
            models.Index(fields=["b"]),
        ]


class AssetToTag(models.Model):
    pk = models.CompositePrimaryKey("a_id", "b_id")
    a = models.ForeignKey("Asset", on_delete=models.CASCADE, related_name="+", db_column="A", db_index=False)
    b = models.ForeignKey("Tag", on_delete=models.CASCADE, related_name="+", db_column="B", db_index=False)

    class Meta:
        db_table = "_AssetToTag"
        indexes = [
            models.Index(fields=["b"]),
        ]


class CategoryToCustomField(models.Model):
    pk = models.CompositePrimaryKey("a_id", "b_id")
    a = models.ForeignKey("Category", on_delete=models.CASCADE, related_name="+", db_column="A", db_index=False)
    b = models.ForeignKey("CustomField", on_delete=models.CASCADE, related_name="+", db_column="B", db_index=False)

    class Meta:
        db_table = "_CategoryToCustomField"
        indexes = [
            models.Index(fields=["b"]),
        ]


class BookingToTag(models.Model):
    pk = models.CompositePrimaryKey("a_id", "b_id")
    a = models.ForeignKey("Booking", on_delete=models.CASCADE, related_name="+", db_column="A", db_index=False)
    b = models.ForeignKey("Tag", on_delete=models.CASCADE, related_name="+", db_column="B", db_index=False)

    class Meta:
        db_table = "_BookingToTag"
        indexes = [
            models.Index(fields=["b"]),
        ]


class AssetReminderToTeamMember(models.Model):
    pk = models.CompositePrimaryKey("a_id", "b_id")
    a = models.ForeignKey("AssetReminder", on_delete=models.CASCADE, related_name="+", db_column="A", db_index=False)
    b = models.ForeignKey("TeamMember", on_delete=models.CASCADE, related_name="+", db_column="B", db_index=False)

    class Meta:
        db_table = "_AssetReminderToTeamMember"
        indexes = [
            models.Index(fields=["b"]),
        ]


class BookingSettingsAlwaysNotify(models.Model):
    pk = models.CompositePrimaryKey("a_id", "b_id")
    a = models.ForeignKey("BookingSettings", on_delete=models.CASCADE, related_name="+", db_column="A", db_index=False)
    b = models.ForeignKey("TeamMember", on_delete=models.CASCADE, related_name="+", db_column="B", db_index=False)

    class Meta:
        db_table = "_BookingSettingsAlwaysNotify"
        indexes = [
            models.Index(fields=["b"]),
        ]


class BookingNotificationRecipients(models.Model):
    pk = models.CompositePrimaryKey("a_id", "b_id")
    a = models.ForeignKey("Booking", on_delete=models.CASCADE, related_name="+", db_column="A", db_index=False)
    b = models.ForeignKey("TeamMember", on_delete=models.CASCADE, related_name="+", db_column="B", db_index=False)

    class Meta:
        db_table = "_BookingNotificationRecipients"
        indexes = [
            models.Index(fields=["b"]),
        ]
