import uuid
from datetime import datetime
from datetime import timezone as dt_timezone
from typing import Any, TypeVar

from django.conf import settings
from django.contrib.postgres.fields import ArrayField
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


_M = TypeVar("_M", bound=models.Model)


class AutoNowQuerySet(models.QuerySet[_M]):
    def update(self, **kwargs: Any) -> int:
        for field in self.model._meta.concrete_fields:
            if getattr(field, "auto_now", False):
                kwargs.setdefault(field.name, getattr(field, "stamp", timezone.now)())
        return super().update(**kwargs)


def uuid4_str() -> str:
    return str(uuid.uuid4())


def survey_hidden_fields_default() -> dict[str, Any]:
    return {"enabled": False}


def survey_welcome_card_default() -> dict[str, Any]:
    return {"enabled": False}


def survey_single_use_default() -> dict[str, Any]:
    return {"enabled": False, "isEncrypted": True}


def survey_recaptcha_default() -> dict[str, Any]:
    return {"enabled": False, "threshold": 0.1}


def workspace_styling_default() -> dict[str, Any]:
    return {"allowStyleOverwrite": True}


def dashboard_widget_layout_default() -> dict[str, Any]:
    return {"x": 0, "y": 0, "w": 4, "h": 3}


class CreativeType(models.TextChoices):
    """Creatives (ENG-3993): page elements the web SDK renders next to surveys. Every type lives in
    `Creative`; `content` and `styling` are JSON validated per type by the API v3 schemas, so a new type is
    a new payload variant, not new columns.
    """

    feedbackButton = "feedbackButton"
    banner = "banner"


class CreativeStatus(models.TextChoices):
    """The author's intent only. Scheduled, Active, "Paused with its survey" and Closed after the
    end time are derived (`apps/web/modules/ee/creatives/lib/display-status.ts` and its SQL twin); nothing
    writes this column when a schedule ends or a survey pauses.
    """

    draft = "draft"
    inProgress = "inProgress"
    paused = "paused"
    completed = "completed"


class CreativeShowOn(models.TextChoices):
    """What makes a creative appear. A creative left with no action never shows; it never falls back to
    every page view.
    """

    pageView = "pageView"
    action = "action"


class CreativeDeviceType(models.TextChoices):
    """The device type counts are split by: the values of the Show when `deviceType` operand."""

    phone = "phone"
    tablet = "tablet"
    desktop = "desktop"


class PipelineTriggers(models.TextChoices):
    responseCreated = "responseCreated"
    responseUpdated = "responseUpdated"
    responseFinished = "responseFinished"


class WebhookSource(models.TextChoices):
    user = "user"
    zapier = "zapier"
    make = "make"
    n8n = "n8n"
    activepieces = "activepieces"


class ContactAttributeType(models.TextChoices):
    default = "default"
    custom = "custom"


class ContactAttributeDataType(models.TextChoices):
    string = "string"
    number = "number"
    date = "date"


class ResponseOutcome(models.TextChoices):
    """What a reviewer decided about a response. One value today; null on `Response.outcome` means none."""

    lowQuality = "lowQuality"


class SurveyStatus(models.TextChoices):
    draft = "draft"
    inProgress = "inProgress"
    paused = "paused"
    completed = "completed"


class SurveyAttributeFilterCondition(models.TextChoices):
    equals = "equals"
    notEquals = "notEquals"


class SurveyQuotaAction(models.TextChoices):
    endSurvey = "endSurvey"
    continueSurvey = "continueSurvey"


class ResponseQuotaLinkStatus(models.TextChoices):
    screenedIn = "screenedIn"
    screenedOut = "screenedOut"


class SurveyType(models.TextChoices):
    link = "link"
    app = "app"


class DisplayOptions(models.TextChoices):
    displayOnce = "displayOnce"
    displayMultiple = "displayMultiple"
    displaySome = "displaySome"
    respondMultiple = "respondMultiple"


class SurveyVisibility(models.TextChoices):
    """Who may read a survey. `workspace` is everyone with workspace access (the behaviour before ENG-3282);
    `restricted` is the owner and the organization's owners and managers only.
    """

    restricted = "restricted"
    workspace = "workspace"


class SurveyScriptMode(models.TextChoices):
    add = "add"
    replace_ = "replace"


class ActionType(models.TextChoices):
    code = "code"
    noCode = "noCode"


class IntegrationType(models.TextChoices):
    googleSheets = "googleSheets"
    notion = "notion"
    airtable = "airtable"
    slack = "slack"


class DataMigrationStatus(models.TextChoices):
    pending = "pending"
    applied = "applied"
    failed = "failed"


class WidgetPlacement(models.TextChoices):
    bottomLeft = "bottomLeft"
    bottomRight = "bottomRight"
    topLeft = "topLeft"
    topRight = "topRight"
    center_ = "center"


class SurveyOverlay(models.TextChoices):
    none = "none"
    light = "light"
    dark = "dark"


class BootstrapAdminMarker(models.TextChoices):
    """The bootstrap marker's only value (ENG-2247). A one-variant enum rather than a boolean on
    purpose: a unique index over a nullable boolean has a slot for `true` AND one for `false`, so the
    first `false` ever written would take the second slot and break every later sign-up. There is no
    `false` to write here, which is the same guarantee without a CHECK constraint Prisma cannot model.
    """

    bootstrapAdmin = "bootstrapAdmin"


class OrganizationRole(models.TextChoices):
    owner = "owner"
    manager = "manager"
    member = "member"
    billing = "billing"
    dashboard_viewer = "dashboard_viewer"


class ApiKeyPermission(models.TextChoices):
    """Defines permission levels for API keys.
    Controls what operations an API key can perform.
    """

    read = "read"
    write = "write"
    manage = "manage"


class IdentityProvider(models.TextChoices):
    email = "email"
    github = "github"
    google = "google"
    azuread = "azuread"
    openid = "openid"
    saml = "saml"


class EmbeddedDataSource(models.TextChoices):
    """Where an Embedded Data field's value comes from.
    `computed` is set by survey logic, `ingested` arrives from a URL param or the SDK, and
    `reserved` is auto-captured system metadata. Reserved fields are never stored as rows —
    they are a static catalog in code — but the value exists so every layer shares one type.
    """

    computed = "computed"
    ingested = "ingested"
    reserved = "reserved"


class EmbeddedDataType(models.TextChoices):
    """The declared type of an Embedded Data field's value."""

    string = "string"
    number = "number"
    boolean = "boolean"
    date = "date"


class TeamUserRole(models.TextChoices):
    admin = "admin"
    contributor = "contributor"


class WorkspaceTeamPermission(models.TextChoices):
    read = "read"
    readWrite = "readWrite"
    manage = "manage"


class ChartType(models.TextChoices):
    area = "area"
    bar = "bar"
    pie = "pie"
    big_number = "big_number"
    matrix = "matrix"


class FeedbackSourceType(models.TextChoices):
    formbricks_survey = "formbricks_survey"
    csv = "csv"


class FeedbackSourceStatus(models.TextChoices):
    active = "active"
    paused = "paused"
    error = "error"


class FeedbackSourceImportMode(models.TextChoices):
    """Which survey responses a Formbricks feedback source imports into Hub.
    completedOnly matches the live ingestion path, which only runs on responseFinished.
    all additionally imports partial responses — answers a respondent typed but never submitted.
    """

    completedOnly = "completedOnly"
    all = "all"


class FeedbackSourceElementScope(models.TextChoices):
    """Whether a formbricks_survey source tracks every supported element of the surveys it maps, or an
    explicit subset the operator curated. Ignored by csv sources, which map fields instead.

    Orthogonal to FeedbackSourceImportMode: that decides which responses are imported, this decides
    which questions a response is imported for.

    This exists because a question added to a survey after the source was connected is
    indistinguishable, from stored state alone, from a question the operator deliberately excluded —
    survey elements carry no timestamps. Recording the intent at write time (where the full candidate
    set IS known) is the same approach webhooks take with `surveyIds: []` meaning "all surveys", and
    segment survey-interaction filters with `surveyScope: "any" | "specific"`.

    Defaults to `specific` so sources that predate the column keep their current behaviour: they gain
    no mappings on their own until the operator next saves the source, at which point the scope is
    derived from the selection they submit.
    """

    all = "all"
    specific = "specific"


class HubFieldType(models.TextChoices):
    text = "text"
    categorical = "categorical"
    nps = "nps"
    csat = "csat"
    ces = "ces"
    rating = "rating"
    number = "number"
    boolean = "boolean"
    date = "date"


class DeletionCleanupKind(models.TextChoices):
    hubSurvey = "hubSurvey"
    hubResponses = "hubResponses"
    storageFiles = "storageFiles"
    storageSurveyFolder = "storageSurveyFolder"
    hubResponseOutcome = "hubResponseOutcome"


class AttachmentExportStatus(models.TextChoices):
    queued = "queued"
    running = "running"
    completed = "completed"
    failed = "failed"
    expired = "expired"


class RetentionEntity(models.TextChoices):
    """Data retention (ENG-3697). The kind of data a policy governs."""

    responses = "responses"
    surveys = "surveys"
    members = "members"


class RetentionSurveyCondition(models.TextChoices):
    """The surveys policy's conditions. Every ticked one must hold, and they share the policy's period."""

    noResponse = "noResponse"
    noChange = "noChange"
    createdBefore = "createdBefore"


class RetentionRunKind(models.TextChoices):
    """What wrote a `RetentionRun`: the nightly sweep, or the archive purge recording the deletion of surveys
    the surveys policy archived. Only sweep runs count as the sweep having run (its lease, its warning
    restart after a gap, where its scan resumes, the health check).
    """

    sweep = "sweep"
    purge = "purge"


class RetentionExemptionSource(models.TextChoices):
    """Who made an exemption: a person (with a reason), or a restore from the archive under a surveys policy
    whose only condition is "created before" (`restartSurveyRetention`), which carries no reason.
    """

    manual = "manual"
    restore = "restore"


class RetentionTargetType(models.TextChoices):
    """What a `RetentionRunItem` row points at. Response deletions are grouped into one `survey` row per
    survey, with the number of responses in `count`.
    """

    survey = "survey"
    user = "user"


class RetentionRunItemAction(models.TextChoices):
    notified = "notified"
    archived = "archived"
    deactivated = "deactivated"
    deleted = "deleted"
    skipped = "skipped"


class RetentionSkipReason(models.TextChoices):
    exempt = "exempt"
    lastOwner = "lastOwner"
    otherOrganization = "otherOrganization"
    noRecipient = "noRecipient"


class WorkflowStatus(models.TextChoices):
    draft = "draft"
    enabled = "enabled"
    disabled = "disabled"
    archived = "archived"


class WorkflowRunStatus(models.TextChoices):
    queued = "queued"
    running = "running"
    completed = "completed"
    failed = "failed"
    canceled = "canceled"


class WorkflowRunLogStatus(models.TextChoices):
    pending = "pending"
    running = "running"
    succeeded = "succeeded"
    failed = "failed"
    skipped = "skipped"


class Creative(models.Model):
    """A feedback button or a banner. Opens at most one Destination: a survey (`surveyId`) or a link
    (`linkUrl`), never both.

    @property type - Fixed at creation.
    @property content - Per-type JSON (`FeedbackButtonContent` / `BannerContent` in the v3 spec), texts keyed by locale code.
    @property styling - Per-type JSON (`FeedbackButtonStyling` / `BannerStyling` in the v3 spec).
    @property survey - The Destination survey, in the same workspace by construction of the composite key.
    @property showConditions - Show when (`ZShowConditions`, shared with surveys); null always shows.
    @property autoSelectLanguage - The browser's language picks the texts' language when the site sets none; null (never chosen, as on Survey) is off.
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    name = models.TextField()
    type = models.TextField(choices=CreativeType.choices)
    status = models.TextField(choices=CreativeStatus.choices, default=CreativeStatus.draft)
    published_at = UtcDateTimeField(null=True, db_column="publishedAt")
    archived_at = UtcDateTimeField(null=True, db_column="archivedAt")
    workspace = models.ForeignKey("Workspace", on_delete=models.CASCADE, related_name="creatives", db_column="workspaceId", db_index=False)
    creator = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="creatives", db_column="createdBy", null=True, db_index=False)
    content = models.JSONField()
    styling = models.JSONField()
    survey_id = models.TextField(null=True, db_column="surveyId")
    link_url = models.TextField(null=True, db_column="linkUrl")
    show_on = models.TextField(choices=CreativeShowOn.choices, db_column="showOn", default=CreativeShowOn.pageView)
    show_conditions = models.JSONField(null=True, db_column="showConditions")
    display_percentage = models.FloatField(null=True, db_column="displayPercentage")
    start_at = UtcDateTimeField(null=True, db_column="startAt")
    end_at = UtcDateTimeField(null=True, db_column="endAt")
    auto_select_language = models.BooleanField(null=True, db_column="autoSelectLanguage")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Creative"
        indexes = [
            models.Index(fields=["workspace", "updated_at"]),
            models.Index(fields=["workspace", "created_at"]),
            models.Index(fields=["workspace", "name"]),
            models.Index(fields=["survey_id"]),
        ]


class CreativeDailyStat(models.Model):
    """A creative's Impressions and Clicks on one UTC day on one device type (ENG-4000). Written
    only by the counter flush job as `GREATEST(stored, counted)`, from Redis counters that hold the day's
    running totals, so a repeated flush changes nothing. At most 3 rows per creative per day. Deleting
    the creative deletes its counts.
    """

    pk = models.CompositePrimaryKey("creative_id", "date", "device_type")
    creative = models.ForeignKey("Creative", on_delete=models.CASCADE, related_name="daily_stats", db_column="creativeId", db_index=False)
    date = models.DateField()
    """UTC day."""
    device_type = models.TextField(choices=CreativeDeviceType.choices, db_column="deviceType")
    impressions = models.IntegerField(default=0)
    clicks = models.IntegerField(default=0)

    class Meta:
        db_table = "CreativeDailyStat"


class CreativeDailyPageStat(models.Model):
    """A creative's Impressions and Clicks on one UTC day on one page (host + path), written like
    `CreativeDailyStat`. At most 1000 pages per creative per day; the rest are counted under `(other)`.
    """

    pk = models.CompositePrimaryKey("creative_id", "date", "page")
    creative = models.ForeignKey("Creative", on_delete=models.CASCADE, related_name="daily_page_stats", db_column="creativeId", db_index=False)
    date = models.DateField()
    """UTC day."""
    page = models.TextField()
    """Host + path, without query string or hash, or `(other)`."""
    impressions = models.IntegerField(default=0)
    clicks = models.IntegerField(default=0)

    class Meta:
        db_table = "CreativeDailyPageStat"


class CreativeTrigger(models.Model):
    """Links a creative to the actions that show it (`showOn: action`). Mirrors `SurveyTrigger`."""

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    creative = models.ForeignKey("Creative", on_delete=models.CASCADE, related_name="triggers", db_column="creativeId", db_index=False)
    action_class = models.ForeignKey("ActionClass", on_delete=models.CASCADE, related_name="creative_triggers", db_column="actionClassId", db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "CreativeTrigger"
        constraints = [
            models.UniqueConstraint(fields=["creative", "action_class"], name="CreativeTrigger_creativeId_actionClassId_key"),
        ]
        indexes = [
            models.Index(fields=["action_class"]),
        ]


class CreativeLanguage(models.Model):
    """The languages a creative's texts are written in. Mirrors `SurveyLanguage`; exactly one is the default."""

    pk = models.CompositePrimaryKey("language_id", "creative_id")
    language = models.ForeignKey("Language", on_delete=models.CASCADE, related_name="creative_languages", db_column="languageId", db_index=False)
    creative = models.ForeignKey("Creative", on_delete=models.CASCADE, related_name="languages", db_column="creativeId", db_index=False)
    default = models.BooleanField(default=False)
    enabled = models.BooleanField(default=True)

    class Meta:
        db_table = "CreativeLanguage"
        indexes = [
            models.Index(fields=["creative"]),
        ]


class OutboundEmailAudit(models.Model):
    """Metadata for an outbound email attempt. Content and recipient addresses are never stored here.
    Identifiers are deliberately not foreign keys so erasure can scrub a row without retaining source data.
    """

    id = models.TextField(primary_key=True, default=uuid4_str)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    expires_at = UtcDateTimeField(db_column="expiresAt")
    content_expires_at = UtcDateTimeField(db_column="contentExpiresAt")
    environment = models.TextField()
    email_type = models.TextField(db_column="emailType")
    trigger = models.TextField()
    actor_kind = models.TextField(db_column="actorKind")
    actor_user_id = models.TextField(null=True, db_column="actorUserId")
    configured_by_user_id = models.TextField(null=True, db_column="configuredByUserId")
    actor_user_hash = models.TextField(null=True, db_column="actorUserHash")
    configured_by_user_hash = models.TextField(null=True, db_column="configuredByUserHash")
    organization_id = models.TextField(null=True, db_column="organizationId")
    survey_id = models.TextField(null=True, db_column="surveyId")
    response_id = models.TextField(null=True, db_column="responseId")
    follow_up_id = models.TextField(null=True, db_column="followUpId")
    workflow_id = models.TextField(null=True, db_column="workflowId")
    workflow_run_id = models.TextField(null=True, db_column="workflowRunId")
    invite_id = models.TextField(null=True, db_column="inviteId")
    recipient_mac = models.TextField(null=True, db_column="recipientMac")
    ses_message_hash = models.TextField(null=True, db_column="sesMessageHash")
    snapshot_key = models.TextField(null=True, db_column="snapshotKey")
    snapshot_sha256 = models.TextField(null=True, db_column="snapshotSha256")
    status = models.TextField(default="prepared")
    status_recorded_at = UtcDateTimeField(null=True, db_column="statusRecordedAt")
    audit_gap = models.BooleanField(db_column="auditGap", default=False)
    erasure_requested_at = UtcDateTimeField(null=True, db_column="erasureRequestedAt")
    erased_at = UtcDateTimeField(null=True, db_column="erasedAt")
    content_delete_retry_at = UtcDateTimeField(null=True, db_column="contentDeleteRetryAt")
    content_delete_failures = models.IntegerField(db_column="contentDeleteFailures", default=0)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "OutboundEmailAudit"
        indexes = [
            models.Index(fields=["environment", "created_at", "id"]),
            models.Index(fields=["organization_id", "created_at", "id"]),
            models.Index(fields=["survey_id", "created_at", "id"]),
            models.Index(fields=["response_id", "created_at", "id"]),
            models.Index(fields=["actor_user_id", "created_at", "id"]),
            models.Index(fields=["configured_by_user_id", "created_at", "id"]),
            models.Index(fields=["actor_user_hash", "created_at", "id"]),
            models.Index(fields=["configured_by_user_hash", "created_at", "id"]),
            models.Index(fields=["recipient_mac", "created_at", "id"]),
            models.Index(fields=["ses_message_hash"]),
            models.Index(fields=["status", "created_at", "id"]),
            models.Index(fields=["content_expires_at", "id"]),
            models.Index(fields=["expires_at", "id"]),
            models.Index(fields=["erasure_requested_at", "id"]),
            models.Index(fields=["content_delete_retry_at", "id"]),
        ]


class OutboundEmailAuditErasureIntent(models.Model):
    """Holds an account erasure request across an in-flight email send. The address is never stored.
    The user id is removed once deletion is confirmed; one-way recovery keys expire after 90 days.
    """

    id = models.TextField(primary_key=True, default=uuid4_str)
    user_id = models.TextField(null=True, db_column="userId")
    user_hash = models.TextField(db_column="userHash")
    recipient_mac = models.TextField(db_column="recipientMac")
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    activated_at = UtcDateTimeField(null=True, db_column="activatedAt")
    checked_at = UtcDateTimeField(null=True, db_column="checkedAt")
    expires_at = UtcDateTimeField(db_column="expiresAt")

    class Meta:
        db_table = "OutboundEmailAuditErasureIntent"
        indexes = [
            models.Index(fields=["user_id"]),
            models.Index(fields=["recipient_mac", "activated_at"]),
            models.Index(fields=["user_hash", "activated_at"]),
            models.Index(fields=["checked_at", "id"]),
            models.Index(fields=["expires_at"]),
        ]


class Webhook(models.Model):
    """Represents a webhook endpoint for receiving survey-related events.
    Webhooks can be configured to receive notifications about response creation, updates, and completion.

    @property id - Unique identifier for the webhook
    @property name - Optional display name for the webhook
    @property url - The endpoint URL where events will be sent
    @property source - Origin of the webhook (user, zapier, make, etc.)
    @property workspace - Associated workspace
    @property triggers - Types of events that trigger this webhook
    @property surveyIds - List of surveys this webhook is monitoring
    """

    id = models.TextField(primary_key=True, default=None)
    name = models.TextField(null=True)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    url = models.TextField()
    source = models.TextField(choices=WebhookSource.choices, default=WebhookSource.user)
    workspace = models.ForeignKey("Workspace", on_delete=models.CASCADE, related_name="webhooks", db_column="workspaceId", db_index=False)
    triggers = ArrayField(models.TextField(choices=PipelineTriggers.choices))
    survey_ids = ArrayField(models.TextField(), db_column="surveyIds")
    secret = models.TextField(null=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Webhook"
        indexes = [
            models.Index(fields=["workspace"]),
        ]


class ContactAttribute(models.Model):
    """Represents an attribute value associated with a contact.
    Used to store custom properties and characteristics of contacts.

    @property id - Unique identifier for the attribute
    @property attributeKey - Reference to the attribute definition
    @property contact - The contact this attribute belongs to
    @property value - The string value of the attribute (used for string type + backwards compatibility)
    @property valueNumber - Native numeric storage for number type attributes
    @property valueDate - Native date storage for date type attributes
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    attribute_key = models.ForeignKey("ContactAttributeKey", on_delete=models.CASCADE, related_name="attributes", db_column="attributeKeyId", db_index=False)
    contact = models.ForeignKey("Contact", on_delete=models.CASCADE, related_name="attributes", db_column="contactId", db_index=False)
    value = models.TextField()
    value_number = models.FloatField(null=True, db_column="valueNumber")
    value_date = UtcDateTimeField(null=True, db_column="valueDate")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "ContactAttribute"
        constraints = [
            models.UniqueConstraint(fields=["contact", "attribute_key"], name="ContactAttribute_contactId_attributeKeyId_key"),
        ]
        indexes = [
            models.Index(fields=["attribute_key", "value"]),
            models.Index(fields=["attribute_key", "value_number"]),
            models.Index(fields=["attribute_key", "value_date"]),
        ]


class ContactAttributeKey(models.Model):
    """Defines the possible attributes that can be assigned to contacts.
    Acts as a schema for contact attributes within a workspace.

    @property id - Unique identifier for the attribute key
    @property isUnique - Whether the attribute must have unique values across contacts
    @property key - The attribute identifier used in the system
    @property name - Display name for the attribute
    @property type - Whether this is a default or custom attribute
    @property dataType - The data type of the attribute (string, number, date)
    @property workspace - The workspace this attribute belongs to
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    is_unique = models.BooleanField(db_column="isUnique", default=False)
    key = models.TextField()
    name = models.TextField(null=True)
    description = models.TextField(null=True)
    type = models.TextField(choices=ContactAttributeType.choices, default=ContactAttributeType.custom)
    data_type = models.TextField(choices=ContactAttributeDataType.choices, db_column="dataType", default=ContactAttributeDataType.string)
    workspace = models.ForeignKey("Workspace", on_delete=models.CASCADE, related_name="contact_attribute_keys", db_column="workspaceId", db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "ContactAttributeKey"
        constraints = [
            models.UniqueConstraint(fields=["key", "workspace"], name="ContactAttributeKey_key_workspaceId_key"),
        ]
        indexes = [
            models.Index(fields=["workspace", "created_at"]),
        ]


class Contact(models.Model):
    """Represents a person or user who can receive and respond to surveys.
    Contacts are workspace-specific and can have multiple attributes and responses.

    @property id - Unique identifier for the contact
    @property workspace - The workspace this contact belongs to
    @property responses - Survey responses from this contact
    @property attributes - Custom attributes associated with this contact
    @property displays - Record of surveys shown to this contact
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    workspace = models.ForeignKey("Workspace", on_delete=models.CASCADE, related_name="contacts", db_column="workspaceId", db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Contact"
        indexes = [
            models.Index(fields=["workspace"]),
        ]


class Response(models.Model):
    """Stores a user's response to a survey, including their answers and metadata.
    Each response is linked to a specific survey and optionally to a contact.

    @property id - Unique identifier for the response
    @property finished - Whether the survey was completed
    @property survey - The associated survey
    @property contact - The optional contact who provided the response
    @property data - JSON object containing the actual response data
    @property variables - Custom variables used in the response
    @property ttc - Time to completion metrics
    @property meta - Additional metadata about the response
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    finished = models.BooleanField(default=False)
    survey = models.ForeignKey("Survey", on_delete=models.CASCADE, related_name="responses", db_column="surveyId", db_index=False)
    contact = models.ForeignKey("Contact", on_delete=models.CASCADE, related_name="responses", db_column="contactId", null=True, db_index=False)
    ending_id = models.TextField(null=True, db_column="endingId")
    data = models.JSONField(default=dict)
    """[ResponseData]"""
    variables = models.JSONField(default=dict)
    """[ResponseVariables]"""
    ttc = models.JSONField(default=dict)
    """[ResponseTtc]"""
    meta = models.JSONField(default=dict)
    """[ResponseMeta]"""
    ingest_flags = models.JSONField(null=True)
    """Why a value in `data` will not read back as its declared Embedded Data type — a failed coercion
    or a truncated oversize value (ENG-1845). Server-computed on every ingest and never accepted
    from the client. Null means no ingest boundary has written it, which is not the same claim as
    an empty list: without this column a wrong-typed value is indistinguishable from an absent one,
    and surfacing that in the UI (v2) would have nothing to read.
    [ResponseIngestFlags]
    """
    outcome = models.TextField(choices=ResponseOutcome.choices, null=True)
    """A reviewer's verdict on the response; null means none. `lowQuality` hides it from the Analysis
    export by default and from Feedback Records, and nothing else: it is never deleted and every
    count, list, webhook and integration still includes it (ENG-6780).
    """
    not_shown = models.JSONField(null=True)
    """What the respondent was not shown (ENG-6936, ADR 0029): the blocks and elements a show-if kept
    from them, and per element its options, matrix rows and columns, by stable ids. Written by the
    survey while they answer; null means unknown (older responses and bundles), read as before.
    Omitted from reads without a `select` (PRISMA_GLOBAL_OMIT), so it never reaches a v1/v2 payload,
    webhook or integration.
    [ResponseNotShown]
    """
    contact_attributes = models.JSONField(null=True, db_column="contactAttributes")
    """[ResponseContactAttributes]"""
    single_use_id = models.TextField(null=True, db_column="singleUseId")
    language = models.TextField(null=True)
    display = models.OneToOneField("Display", on_delete=models.SET_NULL, related_name="response", db_column="displayId", null=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Response"
        constraints = [
            models.UniqueConstraint(fields=["survey", "single_use_id"], name="Response_surveyId_singleUseId_key"),
        ]
        indexes = [
            models.Index(fields=["created_at"]),
            models.Index(fields=["created_at", "id"]),
            models.Index(fields=["survey", "created_at"]),
            models.Index(fields=["survey", "created_at", "id"]),
            models.Index(fields=["contact", "created_at"]),
            models.Index(fields=["contact", "created_at", "id"]),
            models.Index(fields=["survey", "finished"]),
        ]


class Tag(models.Model):
    """Represents a label that can be applied to survey responses.
    Used for categorizing and organizing responses within a workspace.

    @property id - Unique identifier for the tag
    @property name - Display name of the tag
    @property responses - Survey responses tagged with this label
    @property workspace - The workspace this tag belongs to
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    name = models.TextField()
    workspace = models.ForeignKey("Workspace", on_delete=models.CASCADE, related_name="tags", db_column="workspaceId", db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Tag"
        constraints = [
            models.UniqueConstraint(fields=["workspace", "name"], name="Tag_workspaceId_name_key"),
        ]
        indexes = [
            models.Index(fields=["workspace"]),
        ]


class TagsOnResponses(models.Model):
    """Junction table linking tags to responses.
    Enables many-to-many relationship between tags and responses.

    @property response - The tagged response
    @property tag - The tag applied to the response
    """

    pk = models.CompositePrimaryKey("response_id", "tag_id")
    response = models.ForeignKey("Response", on_delete=models.CASCADE, related_name="tags", db_column="responseId", db_index=False)
    tag = models.ForeignKey("Tag", on_delete=models.CASCADE, related_name="responses", db_column="tagId", db_index=False)

    class Meta:
        db_table = "TagsOnResponses"


class Display(models.Model):
    """Records when a survey is shown to a user.
    Tracks survey display history and response status.

    @property id - Unique identifier for the display event
    @property survey - The survey that was displayed
    @property contact - The contact who saw the survey
    @property response - The associated response if one exists
    @property creativeId - The creative that opened the survey. No foreign key, so the record survives the creative's deletion.
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    survey = models.ForeignKey("Survey", on_delete=models.CASCADE, related_name="displays", db_column="surveyId", db_index=False)
    contact = models.ForeignKey("Contact", on_delete=models.CASCADE, related_name="displays", db_column="contactId", null=True, db_index=False)
    creative_id = models.TextField(null=True, db_column="creativeId")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Display"
        indexes = [
            models.Index(fields=["survey"]),
            models.Index(fields=["contact", "created_at"]),
        ]


class SurveyTrigger(models.Model):
    """Links surveys to specific trigger actions.
    Defines which user actions should cause a survey to be displayed.
    This is the connection table between Surveys and ActionClasses that determines
    when and under what conditions a survey should be triggered.

    @property id - Unique identifier for the trigger
    @property survey - The survey to be triggered
    @property actionClass - The action that triggers the survey
    @property createdAt - When the trigger was created
    @property updatedAt - When the trigger was last modified
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    survey = models.ForeignKey("Survey", on_delete=models.CASCADE, related_name="triggers", db_column="surveyId", db_index=False)
    action_class = models.ForeignKey("ActionClass", on_delete=models.CASCADE, related_name="survey_triggers", db_column="actionClassId", db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "SurveyTrigger"
        constraints = [
            models.UniqueConstraint(fields=["survey", "action_class"], name="SurveyTrigger_surveyId_actionClassId_key"),
        ]


class SurveyAttributeFilter(models.Model):
    """Defines targeting rules for surveys based on contact attributes.
    Used to show surveys only to contacts matching specific criteria.

    @property id - Unique identifier for the filter
    @property attributeKey - The contact attribute to filter on
    @property survey - The survey being filtered
    @property condition - The comparison operator to use
    @property value - The value to compare against
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    attribute_key = models.ForeignKey("ContactAttributeKey", on_delete=models.CASCADE, related_name="attribute_filters", db_column="attributeKeyId", db_index=False)
    survey = models.ForeignKey("Survey", on_delete=models.CASCADE, related_name="attribute_filters", db_column="surveyId", db_index=False)
    condition = models.TextField(choices=SurveyAttributeFilterCondition.choices)
    value = models.TextField()
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "SurveyAttributeFilter"
        constraints = [
            models.UniqueConstraint(fields=["survey", "attribute_key"], name="SurveyAttributeFilter_surveyId_attributeKeyId_key"),
        ]
        indexes = [
            models.Index(fields=["attribute_key"]),
        ]


class Survey(models.Model):
    """Represents a complete survey configuration including questions, styling, and display rules.
    Core model for the survey functionality in Formbricks.

    @property id - Unique identifier for the survey
    @property name - Display name of the survey
    @property type - Survey delivery method (link, web, website, app)
    @property status - Current state of the survey (draft, active, completed, etc)
    @property workspace - The workspace this survey belongs to
    @property questions - JSON array containing survey questions configuration
    @property displayOption - Rules for how often the survey can be shown
    @property triggers - Actions that can trigger this survey
    @property attributeFilters - Rules for targeting specific contacts
    @property customHeadScripts - Survey-specific custom HTML scripts (self-hosted only)
    @property customHeadScriptsMode - "add" (merge with workspace) or "replace" (override workspace)
    """

    hidden_fields = models.JSONField(db_column="hiddenFields", default=survey_hidden_fields_default)
    """[SurveyHiddenFields]"""
    variables = models.JSONField(default=list)
    """[SurveyVariables]"""
    ownership_initialized = models.BooleanField(db_column="ownershipInitialized", default=False)
    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    name = models.TextField()
    redirect_url = models.TextField(null=True, db_column="redirectUrl")
    type = models.TextField(choices=SurveyType.choices, default=SurveyType.app)
    workspace = models.ForeignKey("Workspace", on_delete=models.CASCADE, related_name="surveys", db_column="workspaceId", db_index=False)
    creator = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="surveys", db_column="createdBy", null=True, db_index=False)
    visibility = models.TextField(choices=SurveyVisibility.choices, default=SurveyVisibility.workspace)
    """ENG-3282: authorization facts, projected into SpiceDB by the survey outbox trigger. `ownerId` is
    the durable ownership source; `createdBy` stays pure attribution and is never rewritten.
    """
    owner = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="owned_surveys", db_column="ownerId", null=True, db_index=False)
    visibility_version = models.IntegerField(db_column="visibilityVersion", default=0)
    """Incremented by every accepted visibility change; fences projection work. An INSERT starts at 0
    with `visibilityProjectedVersion` -1 (the `survey_visibility_pending` trigger): the initial
    projection, decided from PostgreSQL facts until the projector acknowledges it.
    """
    visibility_projected_version = models.IntegerField(db_column="visibilityProjectedVersion", default=0)
    """The `visibilityVersion` SpiceDB last acknowledged; -1 while none ever was."""
    visibility_pending = models.BooleanField(db_column="visibilityPending", default=False)
    """"visibilityVersion" <> "visibilityProjectedVersion", kept by the `survey_visibility_pending`
    trigger (a BEFORE trigger overrides whatever an INSERT or UPDATE supplies). Read-only: never put it
    in create/update data.
    """
    visibility_changed_at = UtcDateTimeField(null=True, db_column="visibilityChangedAt")
    visibility_changed_by_id = models.TextField(null=True, db_column="visibilityChangedById")
    status = models.TextField(choices=SurveyStatus.choices, default=SurveyStatus.draft)
    welcome_card = models.JSONField(db_column="welcomeCard", default=survey_welcome_card_default)
    """[SurveyWelcomeCard]"""
    questions = models.JSONField(default=list)
    """[SurveyQuestions]"""
    blocks = ArrayField(models.JSONField(), default=list)
    """[SurveyBlocks]"""
    endings = ArrayField(models.JSONField(), default=list)
    """[SurveyEnding]"""
    display_option = models.TextField(choices=DisplayOptions.choices, db_column="displayOption", default=DisplayOptions.displayOnce)
    recontact_days = models.IntegerField(null=True, db_column="recontactDays")
    display_limit = models.IntegerField(null=True, db_column="displayLimit")
    inline_triggers = models.JSONField(null=True, db_column="inlineTriggers")
    """[SurveyInlineTriggers]"""
    auto_close = models.IntegerField(null=True, db_column="autoClose")
    keep_ending_card_open = models.BooleanField(db_column="keepEndingCardOpen", default=False)
    auto_complete = models.IntegerField(null=True, db_column="autoComplete")
    delay = models.IntegerField(default=0)
    publish_on = UtcDateTimeField(null=True, db_column="publishOn")
    close_on = UtcDateTimeField(null=True, db_column="closeOn")
    archived_at = UtcDateTimeField(null=True, db_column="archivedAt")
    survey_closed_message = models.JSONField(null=True, db_column="surveyClosedMessage")
    """[SurveyClosedMessage]"""
    segment = models.ForeignKey("Segment", on_delete=models.SET_NULL, related_name="surveys", db_column="segmentId", null=True, db_index=False)
    workspace_overwrites = models.JSONField(null=True, db_column="workspaceOverwrites")
    """[SurveyWorkspaceOverwrites]"""
    styling = models.JSONField(null=True)
    """[SurveyStyling]"""
    single_use = models.JSONField(null=True, db_column="singleUse", default=survey_single_use_default)
    """[SurveySingleUse]"""
    is_verify_email_enabled = models.BooleanField(db_column="isVerifyEmailEnabled", default=False)
    is_back_button_hidden = models.BooleanField(db_column="isBackButtonHidden", default=False)
    is_auto_progressing_enabled = models.BooleanField(db_column="isAutoProgressingEnabled", default=False)
    is_capture_ip_enabled = models.BooleanField(db_column="isCaptureIpEnabled", default=False)
    is_anonymize_responses_enabled = models.BooleanField(default=False)
    pin = models.TextField(null=True)
    display_percentage = models.DecimalField(max_digits=65, decimal_places=30, null=True, db_column="displayPercentage")
    show_language_switch = models.BooleanField(null=True, db_column="showLanguageSwitch")
    auto_select_language = models.BooleanField(null=True, db_column="autoSelectLanguage")
    """Open the survey in the respondent's browser language when it matches an enabled survey
    language. Null on surveys that predate the setting, which reads as off.
    """
    recaptcha = models.JSONField(null=True, default=survey_recaptcha_default)
    """[SurveyRecaptcha]"""
    metadata = models.JSONField(default=dict)
    """[SurveyLinkMetadata]"""
    slug = models.TextField(null=True, unique=True)
    custom_head_scripts = models.TextField(null=True, db_column="customHeadScripts")
    custom_head_scripts_mode = models.TextField(choices=SurveyScriptMode.choices, null=True, db_column="customHeadScriptsMode", default=SurveyScriptMode.add)
    custom_css = models.JSONField(null=True, db_column="customCss")
    """[CustomCssStored] Survey custom CSS (ENG-2949): source plus trusted compiled output, written only
    by the shared custom CSS save service. null = no survey CSS. Kept outside `styling` (M2.07).
    """
    show_conditions = models.JSONField(null=True, db_column="showConditions")
    """[SurveyShowConditions] App surveys: conditions checked in the browser when a trigger fires
    (ENG-3873). null = no conditions, which is what every existing survey keeps.
    """
    start_logic = models.JSONField(null=True, db_column="startLogic")
    """[SurveyStartLogic] Rules before the first question (ENG-6932): block-rule format, run by the
    survey on load. null = none, which is what every existing survey keeps.
    """
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Survey"
        constraints = [
            models.UniqueConstraint(fields=["id", "workspace"], name="Survey_id_workspaceId_key"),
        ]
        indexes = [
            models.Index(fields=["workspace", "updated_at"]),
            models.Index(fields=["workspace", "owner"]),
            models.Index(fields=["segment"]),
            models.Index(fields=["status", "publish_on"]),
            models.Index(fields=["status", "close_on"]),
            models.Index(fields=["archived_at"]),
        ]


class SurveyQuota(models.Model):
    """Represents a quota configuration for a survey.
    Defines response limits and conditions for quota management.

    @property id - Unique identifier for the quota
    @property survey - The survey this quota belongs to
    @property name - Display name for the quota
    @property limit - Maximum number of responses allowed
    @property conditions - JSON object containing quota conditions
    @property action - Action to take when quota is reached
    @property endingCardId - Optional ending card to show when quota is reached
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    survey = models.ForeignKey("Survey", on_delete=models.CASCADE, related_name="quotas", db_column="surveyId", db_index=False)
    name = models.TextField()
    limit = models.IntegerField()
    logic = models.JSONField(default=dict)
    """[SurveyQuotaLogic]"""
    action = models.TextField(choices=SurveyQuotaAction.choices)
    ending_card_id = models.TextField(null=True, db_column="endingCardId")
    count_partial_submissions = models.BooleanField(db_column="countPartialSubmissions", default=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "SurveyQuota"
        constraints = [
            models.UniqueConstraint(fields=["survey", "name"], name="SurveyQuota_surveyId_name_key"),
        ]


class ResponseQuotaLink(models.Model):
    """Junction table linking responses to quotas.
    Tracks which responses counted towards specific quotas and their status.

    @property response - The response that counted towards the quota
    @property quota - The quota this response counted towards
    @property status - Whether the response was screened in or out
    """

    pk = models.CompositePrimaryKey("response_id", "quota_id")
    response = models.ForeignKey("Response", on_delete=models.CASCADE, related_name="quota_links", db_column="responseId", db_index=False)
    quota = models.ForeignKey("SurveyQuota", on_delete=models.CASCADE, related_name="quota_links", db_column="quotaId", db_index=False)
    status = models.TextField(choices=ResponseQuotaLinkStatus.choices)

    class Meta:
        db_table = "ResponseQuotaLink"
        indexes = [
            models.Index(fields=["quota", "status"]),
        ]


class SurveyFollowUp(models.Model):
    """Defines follow-up actions for survey responses.
    Enables automated actions based on specific survey response conditions.

    @property id - Unique identifier for the follow-up
    @property survey - The associated survey
    @property name - Display name for the follow-up
    @property trigger - Conditions that activate the follow-up
    @property action - Actions to take when triggered
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    survey = models.ForeignKey("Survey", on_delete=models.CASCADE, related_name="follow_ups", db_column="surveyId", db_index=False)
    name = models.TextField()
    trigger = models.JSONField()
    """[SurveyFollowUpTrigger]"""
    action = models.JSONField()
    """[SurveyFollowUpAction]"""
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "SurveyFollowUp"


class ActionClass(models.Model):
    """Represents a user action that can trigger surveys.
    Used to define points in the user journey where surveys can be shown.

    @property id - Unique identifier for the action
    @property name - Display name of the action
    @property type - Whether this is a code or no-code action
    @property key - Unique identifier used in code implementation
    @property noCodeConfig - Configuration for no-code setup
    @property workspace - The workspace this action belongs to
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    name = models.TextField()
    description = models.TextField(null=True)
    type = models.TextField(choices=ActionType.choices)
    key = models.TextField(null=True)
    no_code_config = models.JSONField(null=True, db_column="noCodeConfig")
    """[ActionClassNoCodeConfig]"""
    workspace = models.ForeignKey("Workspace", on_delete=models.CASCADE, related_name="action_classes", db_column="workspaceId", db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "ActionClass"
        constraints = [
            models.UniqueConstraint(fields=["key", "workspace"], name="ActionClass_key_workspaceId_key"),
            models.UniqueConstraint(fields=["name", "workspace"], name="ActionClass_name_workspaceId_key"),
        ]
        indexes = [
            models.Index(fields=["workspace", "created_at"]),
        ]


class Integration(models.Model):
    """Represents third-party service integrations.
    Enables data flow between Formbricks and external services.

    @property id - Unique identifier for the integration
    @property type - The service being integrated (Google Sheets, Notion, etc.)
    @property workspace - The workspace this integration belongs to
    @property config - Service-specific configuration details
    """

    id = models.TextField(primary_key=True, default=None)
    type = models.TextField(choices=IntegrationType.choices)
    config = models.JSONField()
    """[IntegrationConfig]"""
    workspace = models.ForeignKey("Workspace", on_delete=models.CASCADE, related_name="integrations", db_column="workspaceId", db_index=False)

    class Meta:
        db_table = "Integration"
        constraints = [
            models.UniqueConstraint(fields=["type", "workspace"], name="Integration_type_workspaceId_key"),
        ]
        indexes = [
            models.Index(fields=["workspace"]),
        ]


class DataMigration(models.Model):
    """Tracks database schema migrations.
    Used to manage and track the state of data structure changes.

    @property id - Unique identifier for the migration
    @property name - Name of the migration
    @property status - Current state of the migration (pending, applied, failed)
    @property startedAt - When the migration began
    @property finishedAt - When the migration completed
    """

    id = models.TextField(primary_key=True, default=None)
    started_at = UtcDateTimeField(default=timezone.now)
    finished_at = UtcDateTimeField(null=True)
    name = models.TextField(unique=True)
    status = models.TextField(choices=DataMigrationStatus.choices)

    class Meta:
        db_table = "DataMigration"


class Workspace(models.Model):
    """Main grouping mechanism for resources in Formbricks.
    Each organization can have multiple workspaces to separate different applications or products.

    @property id - Unique identifier for the workspace
    @property name - Display name of the workspace
    @property organization - Reference to parent organization
    @property surveys - Collection of surveys in this workspace
    @property styling - Workspace-wide styling configuration
    @property config - Workspace-specific configuration
    @property recontactDays - Default recontact delay for surveys
    @property placement - Default widget placement for in-app surveys
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    name = models.TextField()
    legacy_environment_id = models.TextField(null=True, unique=True, db_column="legacyEnvironmentId")
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="workspaces", db_column="organizationId", db_index=False)
    styling = models.JSONField(default=workspace_styling_default)
    """[Styling]"""
    config = models.JSONField(default=dict)
    """[WorkspaceConfig]"""
    recontact_days = models.IntegerField(db_column="recontactDays", default=7)
    link_survey_branding = models.BooleanField(db_column="linkSurveyBranding", default=True)
    in_app_survey_branding = models.BooleanField(db_column="inAppSurveyBranding", default=True)
    placement = models.TextField(choices=WidgetPlacement.choices, default=WidgetPlacement.bottomRight)
    click_outside_close = models.BooleanField(db_column="clickOutsideClose", default=True)
    overlay = models.TextField(choices=SurveyOverlay.choices, default=SurveyOverlay.none)
    overlay_color = models.TextField(null=True, db_column="overlayColor")
    """Custom overlay colour (hex, no alpha) and opacity (10-100%). null = the preset for `overlay`."""
    overlay_opacity = models.IntegerField(null=True, db_column="overlayOpacity")
    logo = models.JSONField(null=True)
    """[Logo]"""
    app_setup_completed = models.BooleanField(db_column="appSetupCompleted", default=False)
    custom_head_scripts = models.TextField(null=True, db_column="customHeadScripts")
    custom_css = models.JSONField(null=True, db_column="customCss")
    """[CustomCssStored] Workspace custom CSS (ENG-2949), written only by the shared save service."""
    custom_css_previous = models.JSONField(null=True, db_column="customCssPrevious")
    """[CustomCssStored] The one recoverable previous workspace CSS revision, source included. Restored by
    submitting its source again through the current processor and authorization checks.
    """
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Workspace"
        constraints = [
            models.UniqueConstraint(fields=["organization", "name"], name="Workspace_organizationId_name_key"),
        ]


class Organization(models.Model):
    """Represents the top-level organizational hierarchy in Formbricks.
    Self-hosting instances typically have a single organization, while Formbricks Cloud
    supports multiple organizations with multi-tenancy.

    @property id - Unique identifier for the organization
    @property name - Display name of the organization
    @property memberships - User memberships within the organization
    @property workspaces - Collection of workspaces owned by the organization
    @property billing - JSON field containing billing information
    @property whitelabel - Whitelabel configuration for the organization
    @property isAISmartToolsEnabled - Controls access to AI smart tools (e.g. translations) that never touch collected data
    @property displayTimeZone - Optional IANA time zone used for human-facing response timestamps (exports, integrations) and for the calendar days behind the analytics date filters (survey summary, charts, dashboards); null means UTC
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    name = models.TextField()
    whitelabel = models.JSONField(default=dict)
    """[OrganizationWhitelabel]"""
    is_aismart_tools_enabled = models.BooleanField(db_column="isAISmartToolsEnabled", default=False)
    display_time_zone = models.TextField(null=True, db_column="displayTimeZone")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Organization"


class OrganizationBilling(models.Model):
    """Stores billing and Stripe synchronization data for an organization.

    @property organizationId - Reference to the organization
    @property stripeCustomerId - Stripe customer ID for cloud billing
    @property limits - Plan limits used by feature and quota checks
    @property usageCycleAnchor - Stable anchor used to derive the current monthly usage cycle
    @property stripe - Stripe synchronization snapshot data
    """

    organization = models.OneToOneField("Organization", on_delete=models.CASCADE, related_name="billing", primary_key=True)
    stripe_customer_id = models.TextField(null=True, unique=True)
    limits = models.JSONField()
    """[OrganizationBillingPlanLimits]"""
    usage_cycle_anchor = UtcDateTimeField(null=True)
    stripe = models.JSONField(null=True)
    """[OrganizationStripeBilling]"""
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "OrganizationBilling"


class Membership(models.Model):
    """Links users to organizations with specific roles.
    Manages organization membership and permissions.
    Core model for managing user access within organizations.

    @property organization - The organization the user belongs to
    @property user - The member user
    @property accepted - Whether the user has accepted the membership
    @property role - User's role within the organization (owner, manager, member, billing)
    """

    pk = models.CompositePrimaryKey("user_id", "organization_id")
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="memberships", db_column="organizationId", db_index=False)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="memberships", db_column="userId", db_index=False)
    accepted = models.BooleanField(default=False)
    role = models.TextField(choices=OrganizationRole.choices, default=OrganizationRole.member)

    class Meta:
        db_table = "Membership"
        indexes = [
            models.Index(fields=["organization"]),
        ]


class Invite(models.Model):
    """Represents pending invitations to join an organization.
    Used to manage the process of adding new users to an organization.
    Once accepted, invites are converted into memberships.

    @property id - Unique identifier for the invite
    @property email - Email address of the invited user
    @property name - Optional display name for the invited user
    @property organization - The organization sending the invite
    @property creator - The user who created the invite
    @property acceptor - The user who accepted the invite (if accepted)
    @property expiresAt - When the invite becomes invalid
    @property role - Intended role for the invited user
    @property teamIds - Teams the user will be added to upon acceptance
    """

    id = models.TextField(primary_key=True, default=uuid4_str)
    email = models.TextField()
    name = models.TextField(null=True)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="invites", db_column="organizationId", db_index=False)
    creator = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="invites_created", db_column="creatorId", db_index=False)
    acceptor = models.ForeignKey("User", on_delete=models.CASCADE, related_name="invites_accepted", db_column="acceptorId", null=True, db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    expires_at = UtcDateTimeField(db_column="expiresAt")
    role = models.TextField(choices=OrganizationRole.choices, default=OrganizationRole.member)
    team_ids = ArrayField(models.TextField(), db_column="teamIds", default=list)

    class Meta:
        db_table = "Invite"
        indexes = [
            models.Index(fields=["email", "organization"]),
            models.Index(fields=["organization"]),
        ]


class ApiKey(models.Model):
    """Represents enhanced API authentication keys with organization-level ownership.
    Used for authenticating API requests to Formbricks with more granular permissions.

    @property id - Unique identifier for the API key
    @property label - Optional descriptive name for the key
    @property hashedKey - Securely stored API key
    @property organization - The organization this key belongs to
    @property createdBy - User ID who created this key
    @property lastUsedAt - Timestamp of last usage
    @property apiKeyWorkspaces - Workspaces this key has access to
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    created_by = models.TextField(null=True, db_column="createdBy")
    last_used_at = UtcDateTimeField(null=True, db_column="lastUsedAt")
    label = models.TextField()
    hashed_key = models.TextField(db_column="hashedKey")
    lookup_hash = models.TextField(null=True, unique=True, db_column="lookupHash")
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="api_keys", db_column="organizationId", db_index=False)
    organization_access = models.JSONField(db_column="organizationAccess", default=dict)
    """[OrganizationAccess]"""

    class Meta:
        db_table = "ApiKey"
        indexes = [
            models.Index(fields=["organization"]),
        ]


class ApiKeyWorkspace(models.Model):
    """Links API keys to workspaces with specific permissions.
    Enables granular access control for API keys across workspaces.

    @property id - Unique identifier for the workspace access entry
    @property apiKey - The associated API key
    @property workspace - The workspace being accessed
    @property permission - Level of access granted
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    api_key = models.ForeignKey("ApiKey", on_delete=models.CASCADE, related_name="api_key_workspaces", db_column="apiKeyId", db_index=False)
    workspace = models.ForeignKey("Workspace", on_delete=models.CASCADE, related_name="api_key_workspaces", db_column="workspaceId", db_index=False)
    permission = models.TextField(choices=ApiKeyPermission.choices)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "ApiKeyWorkspace"
        constraints = [
            models.UniqueConstraint(fields=["api_key", "workspace"], name="ApiKeyWorkspace_apiKeyId_workspaceId_key"),
        ]
        indexes = [
            models.Index(fields=["workspace"]),
        ]


class Account(models.Model):
    """Stores third-party authentication account information.
    Enables OAuth and other external authentication methods.

    @property id - Unique identifier for the account
    @property user - The Formbricks user who owns this account
    @property provider - The authentication provider (GitHub, Google, etc.)
    @property providerAccountId - User ID from the provider
    @property access_token - OAuth access token
    @property refresh_token - OAuth refresh token
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="accounts", db_column="userId", db_index=False)
    type = models.TextField(null=True)
    provider = models.TextField()
    provider_account_id = models.TextField(db_column="providerAccountId")
    access_token = models.TextField(null=True)
    refresh_token = models.TextField(null=True)
    expires_at = models.IntegerField(null=True)
    ext_expires_in = models.IntegerField(null=True)
    token_type = models.TextField(null=True)
    scope = models.TextField(null=True)
    id_token = models.TextField(null=True)
    session_state = models.TextField(null=True)
    password = models.TextField(null=True)
    access_token_expires_at = UtcDateTimeField(null=True, db_column="accessTokenExpiresAt")
    refresh_token_expires_at = UtcDateTimeField(null=True, db_column="refreshTokenExpiresAt")
    issuer = models.TextField(null=True)
    """Better Auth 1.7.0 (ENG-2343): the account's identity namespace. 1.7.0 keyed accounts on
    (issuer, accountId) and filtered every lookup on it — findAccountByKey, findAccountOwnerByKey,
    findCredentialAccount and updatePassword — so on 1.7.0 a null here locks the user out.
    1.7.3 dropped the field and 1.7.7 keys accounts on (provider, providerAccountId) again
    (ENG-3880). We keep writing it (see `account.create.before` in better-auth-hooks.ts) so a
    rollback to 1.7.0 still finds every row; drop it once that rollback is off the table.

    NULLABLE on purpose. A NOT NULL column would fail every INSERT from a 1.6 pod still serving
    during a rolling deploy. The migration backfills every existing row; see the migration SQL for
    the value derivation.
    """
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Account"
        constraints = [
            models.UniqueConstraint(fields=["provider", "provider_account_id"], name="Account_provider_providerAccountId_key"),
            models.UniqueConstraint(fields=["issuer", "provider_account_id"], name="Account_issuer_providerAccountId_key"),
        ]
        indexes = [
            models.Index(fields=["user"]),
        ]


class Session(models.Model):
    """Stores active authentication sessions for revocable server-side login state.

    @property sessionToken - Opaque token stored in the browser cookie
    @property user - The Formbricks user who owns this session
    @property expires - Hard expiry for the session
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    session_token = models.TextField(unique=True, db_column="sessionToken")
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="sessions", db_column="userId", db_index=False)
    expires = UtcDateTimeField()
    ip_address = models.TextField(null=True, db_column="ipAddress")
    user_agent = models.TextField(null=True, db_column="userAgent")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Session"
        indexes = [
            models.Index(fields=["user"]),
        ]


class VerificationToken(models.Model):
    """Stores one-time verification tokens used by Auth.js adapter flows."""

    pk = models.CompositePrimaryKey("identifier", "token")
    identifier = models.TextField()
    token = models.TextField()
    expires = UtcDateTimeField()

    class Meta:
        db_table = "VerificationToken"
        constraints = [
            models.UniqueConstraint(fields=["identifier", "token"], name="VerificationToken_identifier_token_key"),
        ]


class TwoFactor(models.Model):
    """Better Auth (ENG-1054) two-factor secrets. Supersedes the legacy User.twoFactorSecret /
    User.backupCodes columns (kept until cutover). `secret` and `backupCodes` are encrypted
    at rest by Better Auth using BETTER_AUTH_SECRET.
    """

    id = models.TextField(primary_key=True, default=None)
    secret = models.TextField()
    backup_codes = models.TextField(db_column="backupCodes")
    user = models.OneToOneField("User", on_delete=models.CASCADE, related_name="two_factor", db_column="userId")
    verified = models.BooleanField(default=True)
    failed_verification_count = models.IntegerField(db_column="failedVerificationCount", default=0)
    """Better Auth 1.6.23+ two-factor brute-force lockout: consecutive failed verification attempts."""
    locked_until = UtcDateTimeField(null=True, db_column="lockedUntil")
    """Better Auth 1.6.23+ two-factor brute-force lockout: set once failedVerificationCount trips the limit."""

    class Meta:
        db_table = "TwoFactor"
        constraints = [
            models.UniqueConstraint(fields=["user"], name="TwoFactor_userId_key"),
        ]


class Jwks(models.Model):
    """Better Auth JWT signing keys for OAuth/OIDC access and ID tokens."""

    id = models.TextField(primary_key=True, default=None)
    public_key = models.TextField(db_column="publicKey")
    private_key = models.TextField(db_column="privateKey")
    created_at = UtcDateTimeField(db_column="createdAt")
    expires_at = UtcDateTimeField(null=True, db_column="expiresAt")
    alg = models.TextField(null=True)
    """Better Auth 1.7 (ENG-2343): the JWS algorithm this key signs with, so a keyring can hold more
    than one alg and `getLatestKeyByAlg` can pick the right key per resource.

    NOT optional upstream in practice, despite being declared `required: false`: `createJwk` always
    puts `alg` in the row it writes (defaulting to EdDSA), so without this column every key mint is
    rejected by Prisma as an unknown argument — which breaks JWT signing, and with it the whole MCP
    OAuth flow, on any deployment that has yet to mint a key. Nullable because rows written before
    1.7 have no value; upstream reads treat `alg: null` as the configured default alg, so existing
    keys keep working untouched.
    """
    crv = models.TextField(null=True)
    """Better Auth 1.7 (ENG-2343): the curve for EC/OKP keys. Written whenever it can be derived from
    the generated key, which it can for the default EdDSA/Ed25519 config we run, so it is on the
    same always-written path as `alg` above.
    """

    class Meta:
        db_table = "jwks"


class OauthClient(models.Model):
    """Better Auth OAuth 2.1 client registrations for MCP and future OAuth clients."""

    id = models.TextField(primary_key=True, default=None)
    client_id = models.TextField(unique=True, db_column="clientId")
    client_secret = models.TextField(null=True, db_column="clientSecret")
    disabled = models.BooleanField(null=True, default=False)
    skip_consent = models.BooleanField(null=True, db_column="skipConsent")
    enable_end_session = models.BooleanField(null=True, db_column="enableEndSession")
    subject_type = models.TextField(null=True, db_column="subjectType")
    scopes = ArrayField(models.TextField(), default=list)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="oauth_clients", db_column="userId", null=True, db_index=False)
    created_at = UtcDateTimeField(null=True, db_column="createdAt")
    updated_at = UtcDateTimeField(null=True, db_column="updatedAt")
    name = models.TextField(null=True)
    uri = models.TextField(null=True)
    icon = models.TextField(null=True)
    contacts = ArrayField(models.TextField(), default=list)
    tos = models.TextField(null=True)
    policy = models.TextField(null=True)
    software_id = models.TextField(null=True, db_column="softwareId")
    software_version = models.TextField(null=True, db_column="softwareVersion")
    software_statement = models.TextField(null=True, db_column="softwareStatement")
    redirect_uris = ArrayField(models.TextField(), db_column="redirectUris")
    post_logout_redirect_uris = ArrayField(models.TextField(), db_column="postLogoutRedirectUris", default=list)
    token_endpoint_auth_method = models.TextField(null=True, db_column="tokenEndpointAuthMethod")
    grant_types = ArrayField(models.TextField(), db_column="grantTypes", default=list)
    response_types = ArrayField(models.TextField(), db_column="responseTypes", default=list)
    public = models.BooleanField(null=True)
    """Unwritten by Better Auth 1.7 (it dropped both from its schema), kept deliberately: they are
    nullable, so keeping them makes the 1.7 migration purely additive and a rollback to 1.6 needs no
    down-migration. Drop them in a separate cleanup after a bake period.
    """
    type = models.TextField(null=True)
    require_pkce = models.BooleanField(null=True, db_column="requirePKCE")
    reference_id = models.TextField(null=True, db_column="referenceId")
    metadata = models.JSONField(null=True)
    client_discovery_id = models.TextField(null=True, db_column="clientDiscoveryId")
    """Better Auth 1.7 (ENG-2343)."""
    client_credentials_scopes = ArrayField(models.TextField(), db_column="clientCredentialsScopes", default=list)
    backchannel_logout_uri = models.TextField(null=True, db_column="backchannelLogoutUri")
    backchannel_logout_session_required = models.BooleanField(null=True, db_column="backchannelLogoutSessionRequired")
    application_type = models.TextField(null=True, db_column="applicationType")
    jwks = models.TextField(null=True)
    jwks_uri = models.TextField(null=True, db_column="jwksUri")
    dpop_bound_access_tokens = models.BooleanField(null=True, db_column="dpopBoundAccessTokens", default=False)

    class Meta:
        db_table = "oauthClient"
        indexes = [
            models.Index(fields=["user"]),
        ]


class OauthAccessToken(models.Model):
    """Better Auth opaque access token records. The token column stores Better Auth's hashed lookup value.
    JWT access tokens are verified through JWKS and bounded by the configured short lifetime.
    """

    id = models.TextField(primary_key=True, default=None)
    token = models.TextField(unique=True)
    client = models.ForeignKey("OauthClient", on_delete=models.CASCADE, related_name="access_tokens", to_field="client_id", db_column="clientId", db_index=False)
    session = models.ForeignKey("Session", on_delete=models.SET_NULL, related_name="oauth_access_tokens", db_column="sessionId", null=True, db_index=False)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="oauth_access_tokens", db_column="userId", null=True, db_index=False)
    reference_id = models.TextField(null=True, db_column="referenceId")
    refresh = models.ForeignKey("OauthRefreshToken", on_delete=models.CASCADE, related_name="access_tokens", db_column="refreshId", null=True, db_index=False)
    expires_at = UtcDateTimeField(db_column="expiresAt")
    created_at = UtcDateTimeField(db_column="createdAt")
    scopes = ArrayField(models.TextField())
    authorization_code_id = models.TextField(null=True, db_column="authorizationCodeId")
    """Better Auth 1.7 (ENG-2343). `resources` is the RFC 8707 resource set the grant approved — the
    binding this whole upgrade exists to get (GHSA-p2fr-6hmx-4528).
    """
    resources = ArrayField(models.TextField(), default=list)
    requested_user_info_claims = ArrayField(models.TextField(), db_column="requestedUserInfoClaims", default=list)
    revoked = UtcDateTimeField(null=True)
    confirmation = models.JSONField(null=True)

    class Meta:
        db_table = "oauthAccessToken"
        indexes = [
            models.Index(fields=["client"]),
            models.Index(fields=["session"]),
            models.Index(fields=["user"]),
            models.Index(fields=["refresh"]),
            models.Index(fields=["authorization_code_id"]),
        ]


class OauthRefreshToken(models.Model):
    """Better Auth OAuth refresh tokens used by delegated MCP clients. The token column stores a hashed lookup value."""

    id = models.TextField(primary_key=True, default=None)
    token = models.TextField(unique=True)
    client = models.ForeignKey("OauthClient", on_delete=models.CASCADE, related_name="refresh_tokens", to_field="client_id", db_column="clientId", db_index=False)
    session = models.ForeignKey("Session", on_delete=models.SET_NULL, related_name="oauth_refresh_tokens", db_column="sessionId", null=True, db_index=False)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="oauth_refresh_tokens", db_column="userId", db_index=False)
    reference_id = models.TextField(null=True, db_column="referenceId")
    expires_at = UtcDateTimeField(db_column="expiresAt")
    created_at = UtcDateTimeField(db_column="createdAt")
    revoked = UtcDateTimeField(null=True)
    auth_time = UtcDateTimeField(null=True, db_column="authTime")
    scopes = ArrayField(models.TextField())
    authorization_code_id = models.TextField(null=True, db_column="authorizationCodeId")
    """Better Auth 1.7 (ENG-2343). `resources` carries the full original grant so a refresh can narrow
    the resource set but never widen it.
    """
    resources = ArrayField(models.TextField(), default=list)
    requested_user_info_claims = ArrayField(models.TextField(), db_column="requestedUserInfoClaims", default=list)
    rotated_at = UtcDateTimeField(null=True, db_column="rotatedAt")
    rotation_replay_response = models.TextField(null=True, db_column="rotationReplayResponse")
    rotation_replay_expires_at = UtcDateTimeField(null=True, db_column="rotationReplayExpiresAt")
    confirmation = models.JSONField(null=True)

    class Meta:
        db_table = "oauthRefreshToken"
        indexes = [
            models.Index(fields=["client"]),
            models.Index(fields=["session"]),
            models.Index(fields=["user"]),
            models.Index(fields=["authorization_code_id"]),
        ]


class OauthConsent(models.Model):
    """Better Auth OAuth user consent grants for delegated clients."""

    id = models.TextField(primary_key=True, default=None)
    client = models.ForeignKey("OauthClient", on_delete=models.CASCADE, related_name="consents", to_field="client_id", db_column="clientId", db_index=False)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="oauth_consents", db_column="userId", null=True, db_index=False)
    reference_id = models.TextField(null=True, db_column="referenceId")
    scopes = ArrayField(models.TextField())
    created_at = UtcDateTimeField(db_column="createdAt")
    updated_at = UtcDateTimeField(db_column="updatedAt")
    resources = ArrayField(models.TextField(), default=list)
    """Better Auth 1.7 (ENG-2343). /authorize compares the requested resources against these; without
    them every pre-existing connection is forced back through the consent screen.
    """
    requested_user_info_claims = ArrayField(models.TextField(), db_column="requestedUserInfoClaims", default=list)

    class Meta:
        db_table = "oauthConsent"
        indexes = [
            models.Index(fields=["client"]),
            models.Index(fields=["user"]),
        ]


class OauthResource(models.Model):
    """Better Auth 1.7 (ENG-2343): a protected resource the authorization server issues tokens for, with
    its own token lifetimes, scope allow-list and claims. Replaces the flat `validAudiences` option —
    a token is now bound to the resource it was approved for instead of to a static allow-list.
    `identifier` is the RFC 8707 `resource` value and is the natural key other rows point at.
    """

    id = models.TextField(primary_key=True, default=None)
    identifier = models.TextField(unique=True)
    name = models.TextField()
    access_token_ttl = models.IntegerField(null=True, db_column="accessTokenTtl")
    refresh_token_ttl = models.IntegerField(null=True, db_column="refreshTokenTtl")
    signing_algorithm = models.TextField(null=True, db_column="signingAlgorithm")
    signing_key_id = models.TextField(null=True, db_column="signingKeyId")
    allowed_scopes = ArrayField(models.TextField(), db_column="allowedScopes", default=list)
    custom_claims = models.JSONField(null=True, db_column="customClaims")
    dpop_bound_access_tokens_required = models.BooleanField(null=True, db_column="dpopBoundAccessTokensRequired", default=False)
    disabled = models.BooleanField(null=True, default=False)
    created_at = UtcDateTimeField(null=True, db_column="createdAt")
    updated_at = UtcDateTimeField(null=True, db_column="updatedAt")
    policy_version = models.IntegerField(null=True, db_column="policyVersion", default=1)
    metadata = models.JSONField(null=True)

    class Meta:
        db_table = "oauthResource"


class OauthClientResource(models.Model):
    """Better Auth 1.7 (ENG-2343): links a client to the resources it may request tokens for. With
    `enforcePerClientResources` (the default) a missing row means `invalid_target` at the token
    endpoint, which is why the migration backfills one row per existing client.
    Both sides reference natural keys, matching the other oauth child tables.
    """

    id = models.TextField(primary_key=True, default=None)
    client = models.ForeignKey("OauthClient", on_delete=models.CASCADE, related_name="resources", to_field="client_id", db_column="clientId", db_index=False)
    resource = models.ForeignKey("OauthResource", on_delete=models.CASCADE, related_name="clients", to_field="identifier", db_column="resourceId", db_index=False)
    metadata = models.JSONField(null=True)
    created_at = UtcDateTimeField(null=True, db_column="createdAt")

    class Meta:
        db_table = "oauthClientResource"
        constraints = [
            models.UniqueConstraint(fields=["client", "resource"], name="oauthClientResource_clientId_resourceId_key"),
        ]
        indexes = [
            models.Index(fields=["resource"]),
        ]


class OauthClientAssertion(models.Model):
    """Better Auth 1.7 (ENG-2343): replay store for `private_key_jwt` client assertions (the row id is the
    assertion jti). We use no private_key_jwt client authentication today, but the plugin declares the
    table, so omitting it risks an adapter error on a code path we do not control.
    """

    id = models.TextField(primary_key=True, default=None)
    expires_at = UtcDateTimeField(db_column="expiresAt")

    class Meta:
        db_table = "oauthClientAssertion"
        indexes = [
            models.Index(fields=["expires_at"]),
        ]


class PasswordResetToken(models.Model):
    """Stores the active password reset token for a user.
    Tokens are opaque, hashed at rest, revocable, and single-use.
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    token_hash = models.TextField(unique=True)
    expires_at = UtcDateTimeField()
    user = models.OneToOneField("User", on_delete=models.CASCADE, related_name="password_reset_token", db_column="userId")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "PasswordResetToken"
        indexes = [
            models.Index(fields=["expires_at"]),
        ]


class User(models.Model):
    """Represents a user in the Formbricks system.
    Central model for user authentication and profile management.

    @property id - Unique identifier for the user
    @property name - Display name of the user
    @property email - User's email address
    @property twoFactorEnabled - Whether 2FA is active
    @property memberships - Organizations the user belongs to
    @property notificationSettings - User's notification preferences
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    name = models.TextField()
    email = models.TextField(unique=True)
    email_verified = models.BooleanField(default=False)
    two_factor_secret = models.TextField(null=True, db_column="twoFactorSecret")
    two_factor_enabled = models.BooleanField(db_column="twoFactorEnabled", default=False)
    backup_codes = models.TextField(null=True, db_column="backupCodes")
    password = models.TextField(null=True)
    identity_provider = models.TextField(choices=IdentityProvider.choices, db_column="identityProvider", default=IdentityProvider.email)
    identity_provider_account_id = models.TextField(null=True, db_column="identityProviderAccountId")
    group_id = models.TextField(null=True, db_column="groupId")
    notification_settings = models.JSONField(db_column="notificationSettings", default=dict)
    """[UserNotificationSettings]"""
    locale = models.TextField(default="en-US")
    """[Locale]"""
    last_login_at = UtcDateTimeField(null=True, db_column="lastLoginAt")
    is_active = models.BooleanField(db_column="isActive", default=True)
    reactivated_at = UtcDateTimeField(null=True, db_column="reactivatedAt")
    """When an admin last reactivated this user. Restarts the members retention clock (ENG-3713)
    without touching `lastLoginAt`, which the member list, usage counts and APIs report as a real
    sign-in.
    """
    last_active_at = UtcDateTimeField(null=True, db_column="lastActiveAt")
    """When the user last completed a sign-in (`lastActivePlugin`, after any second factor) or renewed a
    session (Better Auth `databaseHooks.session.update`), on the database clock and never moved
    backwards. The members retention clock reads it: sessions are
    deleted on sign-out and expiry, so their own timestamps can't serve as a record of activity.
    """
    is_bootstrap_admin = models.TextField(choices=BootstrapAdminMarker.choices, null=True, unique=True, db_column="isBootstrapAdmin")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "User"


class Segment(models.Model):
    """Defines a segment of contacts based on attributes.
    Used for targeting surveys to specific user groups.

    @property id - Unique identifier for the segment
    @property title - Display name of the segment
    @property filters - Rules defining the segment
    @property isPrivate - Whether the segment is private
    @property workspace - The workspace this segment belongs to
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    title = models.TextField()
    description = models.TextField(null=True)
    is_private = models.BooleanField(db_column="isPrivate", default=True)
    filters = models.JSONField(default=list)
    """[SegmentFilter]"""
    workspace = models.ForeignKey("Workspace", on_delete=models.CASCADE, related_name="segments", db_column="workspaceId", db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Segment"
        constraints = [
            models.UniqueConstraint(fields=["workspace", "title"], name="Segment_workspaceId_title_key"),
        ]
        indexes = [
            models.Index(fields=["workspace"]),
        ]


class Language(models.Model):
    """Represents a supported language in the system.
    Used for multilingual survey support.

    @property id - Unique identifier for the language
    @property code - Language code (e.g., 'en-US')
    @property alias - Optional friendly name
    @property workspace - The workspace this language is enabled for
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    code = models.TextField()
    alias = models.TextField(null=True)
    workspace = models.ForeignKey("Workspace", on_delete=models.CASCADE, related_name="languages", db_column="workspaceId", db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Language"
        constraints = [
            models.UniqueConstraint(fields=["workspace", "code"], name="Language_workspaceId_code_key"),
        ]


class SurveyLanguage(models.Model):
    """Links surveys to their supported languages.
    Manages which languages are available for each survey.

    @property language - The supported language
    @property survey - The associated survey
    @property default - Whether this is the default language
    @property enabled - Whether this language is currently active
    """

    pk = models.CompositePrimaryKey("language_id", "survey_id")
    language = models.ForeignKey("Language", on_delete=models.CASCADE, related_name="survey_languages", db_column="languageId", db_index=False)
    survey = models.ForeignKey("Survey", on_delete=models.CASCADE, related_name="languages", db_column="surveyId", db_index=False)
    default = models.BooleanField(default=False)
    enabled = models.BooleanField(default=True)

    class Meta:
        db_table = "SurveyLanguage"
        indexes = [
            models.Index(fields=["survey"]),
        ]


class EmbeddedData(models.Model):
    """Defines one Embedded Data field: the extra context stored on a response that isn't an
    answer to a question (brand, plan, page type, a computed score, ...).

    A field is either **local** — used by a single survey, which is how existing variables and
    hidden fields are migrated — or **shared**, meaning it is part of the workspace library and can
    be linked by many surveys. There is no stored flag for that: `surveyId` is the owner and `key`
    is the library name, so a local field has `surveyId` set and `key` null, and a shared field the
    other way round. Deriving it rather than storing a boolean removes a column that could drift
    out of step with the two fields it stands for.

    @property key - Library name, unique per workspace and immutable. Null for local fields
    @property name - Human-readable display label
    @property source - Where the value comes from (computed | ingested). `reserved` is part of the
    shared enum but never stored as a row: those fields are a static catalog in code
    @property dataType - Declared value type, used for coercion at ingest
    @property defaultValue - Initial value (computed) or fallback when nothing arrives (ingested)
    @property locked - Ingested only: ignore external writes and use the default instead
    @property survey - The owning survey for a local field; cascades so nothing is orphaned
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    key = models.TextField(null=True)
    name = models.TextField()
    description = models.TextField(null=True)
    source = models.TextField(choices=EmbeddedDataSource.choices)
    data_type = models.TextField(choices=EmbeddedDataType.choices, db_column="dataType", default=EmbeddedDataType.string)
    default_value = models.JSONField(null=True, db_column="defaultValue")
    """[EmbeddedDataDefaultValue]"""
    locked = models.BooleanField(default=False)
    survey_id = models.TextField(null=True, db_column="surveyId")
    workspace = models.ForeignKey("Workspace", on_delete=models.CASCADE, related_name="embedded_data", db_column="workspaceId", db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "EmbeddedData"
        constraints = [
            models.UniqueConstraint(fields=["workspace", "key"], name="EmbeddedData_workspaceId_key_key"),
            models.UniqueConstraint(fields=["id", "workspace"], name="EmbeddedData_id_workspaceId_key"),
        ]
        indexes = [
            models.Index(fields=["survey_id"]),
        ]


class SurveyEmbeddedData(models.Model):
    """Links a survey to an Embedded Data field and records the key that field's value is stored
    under inside that survey.

    `workspaceId` is part of both foreign keys on purpose: `surveyId` and `embeddedDataId` alone are
    independent, so nothing would stop a survey in one workspace linking a field defined in another.
    Referencing `(id, workspaceId)` on both sides makes a cross-workspace pair unrepresentable, and
    puts the workspace scope every query needs on the row instead of behind a join through Survey.

    @property storageKey - The key the value lives under: a cuid for computed fields, the field
    name for ingested ones. Migrated fields keep their original key, which is what lets existing
    recall tokens, logic conditions and stored responses keep resolving untouched.

    @property order - Position within this survey's field list, dense from 0. Lives on the link
    rather than on `EmbeddedData` because a shared definition sits at a different position in each
    survey that links it. Deliberately has no default: every writer knows the position it is writing,
    so a missing one should be a type error rather than a row that silently sorts to the front.
    """

    id = models.TextField(primary_key=True, default=None)
    workspace_id = models.TextField(db_column="workspaceId")
    survey_id = models.TextField(db_column="surveyId")
    embedded_data_id = models.TextField(db_column="embeddedDataId")
    storage_key = models.TextField(db_column="storageKey")
    order = models.IntegerField()

    class Meta:
        db_table = "SurveyEmbeddedData"
        constraints = [
            models.UniqueConstraint(fields=["survey_id", "embedded_data_id"], name="SurveyEmbeddedData_surveyId_embeddedDataId_key"),
            models.UniqueConstraint(fields=["survey_id", "storage_key"], name="SurveyEmbeddedData_surveyId_storageKey_key"),
        ]
        indexes = [
            models.Index(fields=["embedded_data_id"]),
            models.Index(fields=["survey_id", "order"]),
        ]


class Team(models.Model):
    """Represents a team within an organization.
    Enables group-based access control and collaboration.

    @property id - Unique identifier for the team
    @property name - Display name of the team
    @property organization - The parent organization
    @property teamUsers - Users who are part of this team
    @property workspaceTeams - Workspaces this team has access to
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    name = models.TextField()
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="teams", db_column="organizationId", db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Team"
        constraints = [
            models.UniqueConstraint(fields=["organization", "name"], name="Team_organizationId_name_key"),
        ]


class TeamUser(models.Model):
    """Links users to teams with specific roles.
    Manages team membership and permissions.

    @property team - The associated team
    @property user - The team member
    @property role - User's role within the team
    """

    pk = models.CompositePrimaryKey("team_id", "user_id")
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    team = models.ForeignKey("Team", on_delete=models.CASCADE, related_name="team_users", db_column="teamId", db_index=False)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="team_users", db_column="userId", db_index=False)
    role = models.TextField(choices=TeamUserRole.choices)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "TeamUser"
        indexes = [
            models.Index(fields=["user"]),
        ]


class WorkspaceTeam(models.Model):
    """Defines team access to specific workspaces.
    Manages workspace-level permissions for teams.

    @property workspace - The accessed workspace
    @property team - The team receiving access
    @property permission - Level of access granted
    """

    pk = models.CompositePrimaryKey("workspace_id", "team_id")
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    workspace = models.ForeignKey("Workspace", on_delete=models.CASCADE, related_name="workspace_teams", db_column="workspaceId", db_index=False)
    team = models.ForeignKey("Team", on_delete=models.CASCADE, related_name="workspace_teams", db_column="teamId", db_index=False)
    permission = models.TextField(choices=WorkspaceTeamPermission.choices, default=WorkspaceTeamPermission.read)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "WorkspaceTeam"
        indexes = [
            models.Index(fields=["team"]),
        ]


class Chart(models.Model):
    """Represents a chart/visualization that can be used in multiple dashboards.
    Charts are reusable components that query analytics data.

    @property id - Unique identifier for the chart
    @property name - Display name of the chart
    @property type - Type of visualization (bar, line, pie, etc.)
    @property workspace - The workspace this chart belongs to
    @property query - Cube.js query configuration (JSON)
    @property config - Chart-specific configuration (colors, labels, etc.)
    @property createdBy - User who created the chart
    @property dashboards - Dashboards that use this chart
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    name = models.TextField()
    type = models.TextField(choices=ChartType.choices)
    workspace = models.ForeignKey("Workspace", on_delete=models.CASCADE, related_name="charts", db_column="workspaceId", db_index=False)
    query = models.JSONField(default=dict)
    """[ChartQuery] - Cube.js query configuration"""
    config = models.JSONField(default=dict)
    """[ChartConfig] - Visualization configuration (colors, labels, formatting)"""
    creator = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="charts", db_column="createdBy", null=True, db_index=False)
    feedback_directory = models.ForeignKey("FeedbackDirectory", on_delete=models.RESTRICT, related_name="charts", db_column="feedbackDirectoryId", db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Chart"
        constraints = [
            models.UniqueConstraint(fields=["workspace", "name"], name="Chart_workspaceId_name_key"),
        ]
        indexes = [
            models.Index(fields=["workspace", "created_at"]),
            models.Index(fields=["feedback_directory"]),
        ]


class Dashboard(models.Model):
    """Represents a dashboard containing multiple charts.
    Dashboards aggregate analytics insights at the workspace level.

    @property id - Unique identifier for the dashboard
    @property name - Display name of the dashboard
    @property workspace - The workspace this dashboard belongs to
    @property widgets - Charts on this dashboard
    @property filters - The dashboard filters pinned to its top bar, in order: `{ id, field, operator, label }`,
    at most 10 (ADR-0026). Each viewer's picks live in their own browser, not here.
    @property createdBy - User who created the dashboard
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    name = models.TextField()
    workspace = models.ForeignKey("Workspace", on_delete=models.CASCADE, related_name="dashboards", db_column="workspaceId", db_index=False)
    creator = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="dashboards", db_column="createdBy", null=True, db_index=False)
    filters = models.JSONField(default=list)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Dashboard"
        constraints = [
            models.UniqueConstraint(fields=["workspace", "name"], name="Dashboard_workspaceId_name_key"),
        ]
        indexes = [
            models.Index(fields=["workspace", "created_at"]),
        ]


class DashboardWidget(models.Model):
    """Represents a chart widget on a dashboard.
    Widgets are positioned using a grid layout system.

    @property id - Unique identifier for the widget
    @property dashboard - The dashboard this widget belongs to
    @property chart - The chart displayed in this widget (chart name is used for display)
    @property layout - Grid layout configuration (x, y, width, height)
    @property order - Display order within the dashboard
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    dashboard = models.ForeignKey("Dashboard", on_delete=models.CASCADE, related_name="widgets", db_column="dashboardId", db_index=False)
    chart = models.ForeignKey("Chart", on_delete=models.CASCADE, related_name="widgets", db_column="chartId", db_index=False)
    layout = models.JSONField(default=dashboard_widget_layout_default)
    """[WidgetLayout] - Grid layout: { x, y, w, h }"""
    order = models.IntegerField(default=0)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "DashboardWidget"
        constraints = [
            models.UniqueConstraint(fields=["dashboard", "chart"], name="DashboardWidget_dashboardId_chartId_key"),
        ]
        indexes = [
            models.Index(fields=["dashboard", "order"]),
        ]


class DeletionCleanup(models.Model):
    """Work left over after a delete commits (ENG-3612, ENG-3614: deleted means deleted): Hub records and
    storage files that can't be removed inside the database transaction. A row is written in the same
    transaction as the delete, so a delete that commits always gets its cleanup and one that rolls back
    queues nothing. A worker drains it with retries, never giving up, and deletes each row once done.
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    kind = models.TextField(choices=DeletionCleanupKind.choices)
    organization_id = models.TextField(db_column="organizationId")
    """Scoping and logs. Not foreign keys: the rows they name are usually gone already."""
    workspace_id = models.TextField(db_column="workspaceId")
    survey_id = models.TextField(db_column="surveyId")
    tenant_ids = ArrayField(models.TextField(), db_column="tenantIds", default=list)
    """Hub kinds: the organisation's feedback directories, read inside the delete's transaction."""
    response_ids = ArrayField(models.TextField(), db_column="responseIds", default=list)
    """`hubResponses`: the deleted responses whose records go. `hubResponseOutcome`: exactly one live response."""
    file_keys = ArrayField(models.TextField(), db_column="fileKeys", default=list)
    """`storageFiles`: the object keys or URLs to delete."""
    attempts = models.IntegerField(default=0)
    next_attempt_at = UtcDateTimeField(db_column="nextAttemptAt", default=timezone.now)
    """Not drained before this, so a late pipeline write can't slip in behind a survey's cleanup."""
    last_error = models.TextField(null=True, db_column="lastError")
    """The last failure: a short code and message, never record or file content."""
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "DeletionCleanup"
        indexes = [
            models.Index(fields=["next_attempt_at"]),
            models.Index(fields=["organization_id", "created_at"]),
        ]


class AttachmentExport(models.Model):
    """A survey's response attachments built into one ZIP in object storage by a background job and
    downloaded through a short-lived presigned URL, so an interrupted download can resume (ENG-6862).
    Only the requester reads it. The reconcile job deletes the archive at `expiresAt`; deleting the
    survey deletes it with the survey's storage prefix.

    A partial unique index in the migration allows one queued or running export per requester and
    survey, so a double click returns the export already in flight.
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    survey_id = models.TextField(db_column="surveyId")
    workspace = models.ForeignKey("Workspace", on_delete=models.CASCADE, related_name="attachment_exports", db_column="workspaceId", db_index=False)
    requested_by = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="attachment_exports", db_column="requestedById", null=True, db_index=False)
    status = models.TextField(choices=AttachmentExportStatus.choices, default=AttachmentExportStatus.queued)
    filter_criteria = models.JSONField(null=True, db_column="filterCriteria")
    """The response filter at request time, a `TResponseFilterCriteria`; null exports every response."""
    time_zone = models.TextField(db_column="timeZone")
    """IANA zone the archive's folder names are stamped in: the org's display zone at request time."""
    dispatch_count = models.IntegerField(db_column="dispatchCount", default=0)
    """Bumped on every enqueue and part of the job id, so a re-dispatch never dedupes against a job an
    older pod dropped as unknown during a rolling upgrade.
    """
    file_count = models.IntegerField(db_column="fileCount")
    """The pre-flight's counts, shown while the archive builds."""
    response_count = models.IntegerField(db_column="responseCount")
    included_file_count = models.IntegerField(null=True, db_column="includedFileCount")
    """What the archive actually holds, once built."""
    total_bytes = models.BigIntegerField(null=True, db_column="totalBytes")
    truncated = models.BooleanField(default=False)
    storage_key = models.TextField(null=True, db_column="storageKey")
    error_code = models.TextField(null=True, db_column="errorCode")
    """A short code (`stale`, `storage_upload_failed`, …), never a message or file name."""
    started_at = UtcDateTimeField(null=True, db_column="startedAt")
    finished_at = UtcDateTimeField(null=True, db_column="finishedAt")
    expires_at = UtcDateTimeField(null=True, db_column="expiresAt")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "AttachmentExport"
        indexes = [
            models.Index(fields=["survey_id", "requested_by", "created_at"]),
            models.Index(fields=["status", "created_at"]),
            models.Index(fields=["expires_at"]),
            models.Index(fields=["workspace"]),
        ]


class FeedbackSource(models.Model):
    """A feedback source — connects external data into Unify Feedback for record creation.
    Naming note: "FeedbackSource" is the Unify Feedback domain term; the word "Connector"
    is reserved for the Workflows domain (see docs/development/standards/practices/naming-conventions.mdx).

    @property id - Unique identifier for the feedback source
    @property name - Display name for the feedback source
    @property type - Type of feedback source (formbricks_survey, csv, …)
    @property status - Current state of the feedback source (active, paused, error)
    @property workspace - The workspace this feedback source belongs to
    @property feedbackDirectory - Hub directory the feedback source writes into
    @property formbricksMappings - Element mappings for Formbricks feedback sources
    @property fieldMappings - Field mappings for other feedback source types
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    name = models.TextField()
    type = models.TextField(choices=FeedbackSourceType.choices)
    status = models.TextField(choices=FeedbackSourceStatus.choices, default=FeedbackSourceStatus.active)
    import_mode = models.TextField(choices=FeedbackSourceImportMode.choices, default=FeedbackSourceImportMode.completedOnly)
    """Defaults to completedOnly so a new source agrees with the live ingestion path, which only
    fires on responseFinished. Existing rows take the default; nothing is re-imported.
    """
    element_scope = models.TextField(choices=FeedbackSourceElementScope.choices, db_column="elementScope", default=FeedbackSourceElementScope.specific)
    workspace = models.ForeignKey("Workspace", on_delete=models.CASCADE, related_name="feedback_sources", db_column="workspaceId", db_index=False)
    feedback_directory = models.ForeignKey("FeedbackDirectory", on_delete=models.CASCADE, related_name="feedback_sources", db_column="feedbackDirectoryId", db_index=False)
    last_sync_at = UtcDateTimeField(null=True)
    creator = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="feedback_sources", db_column="created_by", null=True, db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "FeedbackSource"
        constraints = [
            models.UniqueConstraint(fields=["id", "workspace"], name="FeedbackSource_id_workspaceId_key"),
            models.UniqueConstraint(fields=["workspace", "name"], name="FeedbackSource_workspaceId_name_key"),
        ]
        indexes = [
            models.Index(fields=["type"]),
            models.Index(fields=["feedback_directory", "workspace"]),
        ]


class FeedbackSourceFormbricksMapping(models.Model):
    """Maps survey elements to Hub FeedbackRecords for Formbricks feedback sources.
    Each row represents one element that will create FeedbackRecords when answered.

    @property id - Unique identifier for the mapping
    @property feedbackSource - The parent feedback source
    @property survey - The survey containing the element
    @property elementId - The element ID within the survey (from blocks[].elements[].id)
    @property hubFieldType - The field_type to use in Hub (text, nps, rating, etc.)
    @property customFieldLabel - Optional override for the element headline as field_label in Hub
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    feedback_source_id = models.TextField()
    workspace_id = models.TextField(db_column="workspaceId")
    survey_id = models.TextField(db_column="surveyId")
    element_id = models.TextField(db_column="elementId")
    hub_field_type = models.TextField(choices=HubFieldType.choices, db_column="hubFieldType")
    custom_field_label = models.TextField(null=True)

    class Meta:
        db_table = "FeedbackSourceFormbricksMapping"
        constraints = [
            models.UniqueConstraint(fields=["workspace_id", "feedback_source_id", "survey_id", "element_id"], name="FeedbackSourceFormbricksMapping_workspaceId_feedbackSourceId_su"),
        ]
        indexes = [
            models.Index(fields=["workspace_id", "survey_id"]),
            models.Index(fields=["survey_id"]),
        ]


class FeedbackSourceFieldMapping(models.Model):
    """Generic field mapping for non-Formbricks feedback sources (CSV today; webhook/email/slack later).
    Maps source fields to Hub FeedbackRecord fields.

    @property id - Unique identifier for the mapping
    @property feedbackSource - The parent feedback source
    @property sourceFieldId - Field path for webhook (e.g., "user.id"), column name for CSV
    @property targetFieldId - Hub field (collected_at, field_id, value_text, etc.)
    @property staticValue - If set, use this value instead of reading from sourceFieldId
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    feedback_source_id = models.TextField()
    workspace_id = models.TextField(db_column="workspaceId")
    source_field_id = models.TextField()
    target_field_id = models.TextField()
    static_value = models.TextField(null=True)

    class Meta:
        db_table = "FeedbackSourceFieldMapping"
        constraints = [
            models.UniqueConstraint(fields=["workspace_id", "feedback_source_id", "source_field_id", "target_field_id"], name="FeedbackSourceFieldMapping_workspaceId_feedbackSourceId_sourceF"),
        ]


class FeedbackDirectory(models.Model):
    """Represents a feedback directory (Hub tenant) owned by an organization.
    Directories group feedback data and are assigned to workspaces for access control.

    @property id - Unique identifier for the directory
    @property name - Display name of the directory
    @property isArchived - Soft delete flag
    @property organization - The parent organization
    @property workspaces - Workspaces assigned to this directory
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    name = models.TextField()
    is_archived = models.BooleanField(db_column="isArchived", default=False)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="feedback_directories", db_column="organizationId", db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "FeedbackDirectory"
        constraints = [
            models.UniqueConstraint(fields=["organization", "name"], name="FeedbackDirectory_organizationId_name_key"),
        ]


class FeedbackDirectoryWorkspace(models.Model):
    """Links feedback directories to workspaces (workspaces).
    Manages which workspaces can access a given directory.

    @property feedbackDirectory - The directory being accessed
    @property workspace - The workspace receiving access
    """

    pk = models.CompositePrimaryKey("feedback_directory_id", "workspace_id")
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    feedback_directory = models.ForeignKey("FeedbackDirectory", on_delete=models.CASCADE, related_name="workspaces", db_column="feedbackDirectoryId", db_index=False)
    workspace = models.ForeignKey("Workspace", on_delete=models.CASCADE, related_name="feedback_directory_workspaces", db_column="workspaceId", db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "FeedbackDirectoryWorkspace"
        indexes = [
            models.Index(fields=["workspace"]),
        ]


class AuthzedProjectionScopeState(models.Model):
    """ENG-3282: per-scope readiness of a SpiceDB projection. A row with `readyAt` set means the scope was
    backfilled and audited clean, so the evaluator may enforce it. Written only by the backfill CLI.
    """

    scope = models.TextField(primary_key=True, default=None)
    ready_at = UtcDateTimeField(null=True, db_column="readyAt")
    ready_by = models.TextField(null=True, db_column="readyBy")
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "AuthzedProjectionScopeState"


class AuthzedProjectionOutbox(models.Model):
    """Durable PostgreSQL queue for rebuilding SpiceDB relationships after an authorization source changes.
    Source-table triggers insert these rows in the same transaction as the mutation. Identifiers are
    operationally sensitive and must never be included in logs, metrics, or CLI output.

    This model declares NO indexes on purpose. All three live in
    `migration/20260818120000_add_authzed_projection_outbox/migration.sql`, because each one is
    partial (`WHERE "processedAt" IS NULL AND ...`) and Prisma's `@@index` has no `where`. Declaring
    approximate non-partial copies here would be worse than declaring none: `prisma db push` would
    see a name match with a different definition, drop and recreate them without the predicate, and
    the migration's `CREATE INDEX IF NOT EXISTS` would then skip them — leaving a silently wrong
    shape that carries a week of delivered rows on every hot-path lookup. Deployments are unaffected
    either way (`prisma migrate deploy` never reads this file); a dev who runs `db push` restores
    them by rerunning the migration, which is the same contract its triggers already live under.
    """

    id = models.TextField(primary_key=True, default=uuid4_str)
    target_type = models.TextField(db_column="targetType")
    primary_id = models.TextField(db_column="primaryId")
    secondary_id = models.TextField(null=True, db_column="secondaryId")
    is_revocation = models.BooleanField(db_column="isRevocation", default=False)
    attempts = models.IntegerField(default=0)
    """Total delivery attempts. Drives the retry backoff and reports how hard delivery was tried."""
    permanent_failures = models.IntegerField(db_column="permanentFailures", default=0)
    """Non-retryable failures attributable to this event alone. The only input to dead-lettering, so
    that a SpiceDB outage of any duration cannot dead-letter events that were never the problem.
    """
    available_at = UtcDateTimeField(db_column="availableAt", default=timezone.now)
    leased_at = UtcDateTimeField(null=True, db_column="leasedAt")
    lease_expires_at = UtcDateTimeField(null=True, db_column="leaseExpiresAt")
    lease_owner = models.TextField(null=True, db_column="leaseOwner")
    processed_at = UtcDateTimeField(null=True, db_column="processedAt")
    dead_lettered_at = UtcDateTimeField(null=True, db_column="deadLetteredAt")
    last_attempt_at = UtcDateTimeField(null=True, db_column="lastAttemptAt")
    last_error_code = models.TextField(null=True, db_column="lastErrorCode")
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "AuthzedProjectionOutbox"


class OnlineUpgradeState(models.Model):
    release = models.TextField(primary_key=True, default=None)
    target_image = models.TextField(db_column="targetImage")
    phase = models.TextField()
    completed_phases = ArrayField(models.TextField(), db_column="completedPhases", default=list)
    survey_cursor = models.TextField(null=True, db_column="surveyCursor")
    survey_count = models.BigIntegerField(db_column="surveyCount", default=0)
    baseline_completed_at = UtcDateTimeField(null=True, db_column="baselineCompletedAt")
    forward_only_at = UtcDateTimeField(null=True, db_column="forwardOnlyAt")
    queue_paused_at = UtcDateTimeField(null=True, db_column="queuePausedAt")
    queue_pause_deadline_at = UtcDateTimeField(null=True, db_column="queuePauseDeadlineAt")
    queue_resumed_at = UtcDateTimeField(null=True, db_column="queueResumedAt")
    activated_at = UtcDateTimeField(null=True, db_column="activatedAt")
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "OnlineUpgradeState"


class RetentionPolicy(models.Model):
    """One policy per organisation and kind of data: a notice `warnDays` before the action, and the action
    `periodDays` after the target's clock. The action follows from the kind of data: responses are
    deleted, surveys archived (the archive purge deletes them 30 days later, as it does every archived
    survey), members deactivated. The dates themselves are never stored; they are computed by the shared
    helper in `apps/web/modules/ee/data-retention/lib/schedule.ts`, so the UI, the emails and the sweep
    always agree.

    The allowed ranges are validated by the API. The migration repeats the ones that guard against
    acting early or re-notifying (a notice of at least 14 days and shorter than the period, an enabled
    policy always has `enabledAt`) and the structural ones (conditions on the surveys policy only) as
    CHECK constraints, which Prisma cannot express.
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="retention_policies", db_column="organizationId", db_index=False)
    entity = models.TextField(choices=RetentionEntity.choices)
    enabled = models.BooleanField(default=False)
    enabled_at = UtcDateTimeField(null=True, db_column="enabledAt")
    """When the current configuration took effect: the policy was switched on or unpaused, or changed
    in a way that could bring an action earlier (a shorter period or warning, different
    conditions). Lengthening does not move it. The warning always runs in full from here, so no
    change can delete data on the night it is made (ENG-3614). Also the clock for members with no
    recorded sign-in.
    """
    warn_days = models.IntegerField(db_column="warnDays", default=60)
    """How long before the action the notice goes out. The action never runs before the notice has
    been out for this long, so the warning always runs in full.
    """
    period_days = models.IntegerField(db_column="periodDays")
    """The period, counted from the target's clock, after which the policy acts: responses are
    deleted, surveys archived (the archive purge deletes them 30 days later), members deactivated.
    """
    conditions = ArrayField(models.TextField(choices=RetentionSurveyCondition.choices), default=list)
    """Surveys policy only."""
    deferred_at = UtcDateTimeField(null=True, db_column="deferredAt")
    """The last night the sweep's time ran out before it reached this policy's organisation, licensed.
    The organisation was due and the sweep was running, so it counts as reached for the gap that
    restarts the warning (`RETENTION_SWEEP_GAP_MS`): a backlog of organisations never restarts warnings.
    """
    updated_by = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="retention_policies_updated", db_column="updatedById", null=True, db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "RetentionPolicy"
        constraints = [
            models.UniqueConstraint(fields=["organization", "entity"], name="RetentionPolicy_organizationId_entity_key"),
        ]
        indexes = [
            models.Index(fields=["updated_by"]),
        ]


class RetentionExemption(models.Model):
    """Holds one survey back from one policy until `until`. Keyed by organisation + entity rather than a
    policy id, so removing and re-adding a policy doesn't orphan its exemptions.

    Active means `revokedAt IS NULL AND until > now()`, so expiry needs no job. "One active exemption
    per survey and policy" is the partial unique index `RetentionExemption_surveyId_entity_active_key`
    (`WHERE "revokedAt" IS NULL`) in
    `migration/20261007120000_add_data_retention_tables/migration.sql`; Prisma's `@@unique` has no
    `where`, so it is deliberately not declared here (see "Indexes Prisma cannot express" in the
    README). An expired row still holds that slot until it is revoked, so creating a new exemption
    must close an expired one in the same transaction.

    `surveyId` is nullable and a CHECK requires exactly one target, so a later `workflowId` is one more
    column in that CHECK. The survey's organisation can't be tied to `organizationId` in the database
    (a survey carries its workspace, not its organisation): writers derive it from the survey.
    """

    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="retention_exemptions", db_column="organizationId", db_index=False)
    entity = models.TextField(choices=RetentionEntity.choices)
    survey = models.ForeignKey("Survey", on_delete=models.CASCADE, related_name="retention_exemptions", db_column="surveyId", null=True, db_index=False)
    until = UtcDateTimeField()
    source = models.TextField(choices=RetentionExemptionSource.choices, default=RetentionExemptionSource.manual)
    reason = models.TextField(null=True)
    """Required for `manual`, absent for `restore` (`RetentionExemption_reason_check`)."""
    created_by = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="exemptions_created", db_column="createdById", null=True, db_index=False)
    revoked_at = UtcDateTimeField(null=True, db_column="revokedAt")
    revoked_by = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="exemptions_revoked", db_column="revokedById", null=True, db_index=False)

    class Meta:
        db_table = "RetentionExemption"
        indexes = [
            models.Index(fields=["organization", "created_at", "id"]),
            models.Index(fields=["survey"]),
            models.Index(fields=["created_by"]),
            models.Index(fields=["revoked_by"]),
        ]


class RetentionRun(models.Model):
    """One sweep of one policy for one organisation: the History tab's rows. `finishedAt` stays null
    while the run is in progress, or if it never finished.
    """

    id = models.TextField(primary_key=True, default=None)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="retention_runs", db_column="organizationId", db_index=False)
    entity = models.TextField(choices=RetentionEntity.choices)
    kind = models.TextField(choices=RetentionRunKind.choices, default=RetentionRunKind.sweep)
    started_at = UtcDateTimeField(db_column="startedAt", default=timezone.now)
    finished_at = UtcDateTimeField(null=True, db_column="finishedAt")
    notified_count = models.IntegerField(db_column="notifiedCount", default=0)
    archived_count = models.IntegerField(db_column="archivedCount", default=0)
    """Archived surveys, or deactivated members."""
    deleted_count = models.IntegerField(db_column="deletedCount", default=0)
    """Deleted responses, or, on a purge run, deleted surveys."""
    skipped_count = models.IntegerField(db_column="skippedCount", default=0)
    has_changes = models.BooleanField(db_column="hasChanges", default=False)
    """False when the run changed nothing. History hides those by default."""
    scan_cursor = models.TextField(null=True, db_column="scanCursor")
    """Where the run's scan of candidates stopped for want of time, so the next run resumes there instead
    of rescanning the same first targets every night. Null when the scan reached the end.
    """

    class Meta:
        db_table = "RetentionRun"
        indexes = [
            models.Index(fields=["organization", "started_at", "id"]),
        ]


class RetentionRunItem(models.Model):
    """What a run did, one row per target: the History CSV. Ids, counts, names and the recipient's email,
    never response content. `targetId` is not a foreign key on purpose: the target is often gone, and
    `targetName` is a snapshot taken when the row is written so the CSV still says what it was.
    """

    id = models.TextField(primary_key=True, default=None)
    run = models.ForeignKey("RetentionRun", on_delete=models.CASCADE, related_name="items", db_column="runId", db_index=False)
    target_type = models.TextField(choices=RetentionTargetType.choices, db_column="targetType")
    target_id = models.TextField(db_column="targetId")
    target_name = models.TextField(null=True, db_column="targetName")
    action = models.TextField(choices=RetentionRunItemAction.choices)
    count = models.IntegerField(default=1)
    """1, except for grouped response deletions, where it is the number of responses deleted."""
    recipient = models.TextField(null=True)
    """Who was emailed, for `notified` rows."""
    skip_reason = models.TextField(choices=RetentionSkipReason.choices, null=True, db_column="skipReason")

    class Meta:
        db_table = "RetentionRunItem"
        indexes = [
            models.Index(fields=["run", "id"]),
            models.Index(fields=["target_id"]),
        ]


class RetentionNotice(models.Model):
    """The notice for one target under one policy, so each notice goes out once: a survey or a member, or
    the one-time responses reminder for a survey. It counts only once delivered, and is void if it was
    claimed before the policy's `enabledAt`, before the target's latest exemption ended or (surveys and
    members) before the target's clock: then a later cycle claims it again and sends a new one.

    A CHECK in the migration ties each entity to its target: `surveys` and `responses` → `surveyId`,
    `members` → `userId`; another makes a sent email imply a delivery time. The survey notices that still
    hold a clock are the partial index `RetentionNotice_surveys_clock_pending_idx` (`WHERE "entity" =
    'surveys' AND "clockAt" IS NOT NULL`) in
    `migration/20261012120000_retention_notice_void_on_response_delete/migration.sql`, with the triggers
    that clear `clockAt`; Prisma's `@@index` has no `where`, so it is deliberately not declared here (see
    "Indexes Prisma cannot express" in the README).
    """

    id = models.TextField(primary_key=True, default=None)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="retention_notices", db_column="organizationId", db_index=False)
    entity = models.TextField(choices=RetentionEntity.choices)
    survey = models.ForeignKey("Survey", on_delete=models.CASCADE, related_name="retention_notices", db_column="surveyId", null=True, db_index=False)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="retention_notices", db_column="userId", null=True, db_index=False)
    sent_at = UtcDateTimeField(db_column="sentAt", default=timezone.now)
    """When the notice was claimed for sending. A claim that never got delivered can be taken over."""
    delivered_at = UtcDateTimeField(null=True, db_column="deliveredAt")
    """When the notice was handed to the mail transport, or recorded without one (no SMTP). The
    warning only counts from a delivered notice, so a crash between claim and send can't let the
    action run without it.
    """
    email_sent = models.BooleanField(db_column="emailSent", default=False)
    """Whether an email actually went out. False when SMTP isn't configured."""
    claim_token = models.TextField(null=True, db_column="claimToken")
    """Set by each claim; delivery is recorded only for the claim that sent the email, so a slow sender
    can't mark a newer claim delivered.
    """
    clock_at = UtcDateTimeField(null=True, db_column="clockAt")
    """The target's clock the notice was computed for (surveys and members). A notice counts only while
    the target's clock is still this one, so activity between the sweep's read and its claim voids
    it. Null for the responses reminder, which is sent once per survey and follows no one clock, and
    for a survey notice voided for good: a person deleted or moved one of its responses, or the
    responses policy's deletion left it computed for another clock (see `NOTICE_FOLLOWS_CLOCK`).
    """

    class Meta:
        db_table = "RetentionNotice"
        constraints = [
            models.UniqueConstraint(fields=["survey", "entity"], name="RetentionNotice_surveyId_entity_key"),
            models.UniqueConstraint(fields=["user", "organization", "entity"], name="RetentionNotice_userId_organizationId_entity_key"),
        ]
        indexes = [
            models.Index(fields=["organization"]),
        ]


class RetentionSurveyRestore(models.Model):
    """When a survey was last restored from the archive. A restore starts the surveys policy's "no response"
    and "no change" periods over (`getSurveyRetentionClock`); only ever moves forward. A table of its own
    so retention adds nothing to the survey row.
    """

    survey = models.OneToOneField("Survey", on_delete=models.CASCADE, related_name="retention_restore", primary_key=True, db_column="surveyId")
    restored_at = UtcDateTimeField(db_column="restoredAt")

    class Meta:
        db_table = "RetentionSurveyRestore"


class Workflow(models.Model):
    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    name = models.TextField()
    description = models.TextField(null=True)
    status = models.TextField(choices=WorkflowStatus.choices, default=WorkflowStatus.draft)
    workspace = models.ForeignKey("Workspace", on_delete=models.CASCADE, related_name="workflows", db_column="workspaceId", db_index=False)
    creator = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="workflows", db_column="createdBy", null=True, db_index=False)
    definition = models.JSONField()
    """[WorkflowDefinition]"""
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Workflow"
        constraints = [
            models.UniqueConstraint(fields=["id", "workspace"], name="Workflow_id_workspaceId_key"),
        ]
        indexes = [
            models.Index(fields=["workspace", "status"]),
            models.Index(fields=["workspace", "updated_at"]),
        ]


class WorkflowVersion(models.Model):
    id = models.TextField(primary_key=True, default=None)
    workflow_id = models.TextField(db_column="workflowId")
    workspace_id = models.TextField(db_column="workspaceId")
    version = models.IntegerField()
    definition = models.JSONField()
    """[WorkflowExecutableDefinition]"""
    published_at = UtcDateTimeField(db_column="publishedAt", default=timezone.now)
    publisher = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="workflow_versions_published", db_column="publishedBy", null=True, db_index=False)

    class Meta:
        db_table = "WorkflowVersion"
        constraints = [
            models.UniqueConstraint(fields=["workflow_id", "version"], name="WorkflowVersion_workflowId_version_key"),
        ]
        indexes = [
            models.Index(fields=["workflow_id", "workspace_id"]),
            models.Index(fields=["publisher"]),
        ]


class WorkflowRun(models.Model):
    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    workflow_id = models.TextField(db_column="workflowId")
    workspace = models.ForeignKey("Workspace", on_delete=models.CASCADE, related_name="workflow_runs", db_column="workspaceId", db_index=False)
    workflow_version = models.ForeignKey("WorkflowVersion", on_delete=models.SET_NULL, related_name="runs", db_column="workflowVersionId", null=True, db_index=False)
    response = models.ForeignKey("Response", on_delete=models.CASCADE, related_name="workflow_runs", db_column="responseId", null=True, db_index=False)
    status = models.TextField(choices=WorkflowRunStatus.choices, default=WorkflowRunStatus.queued)
    trigger_type = models.TextField(db_column="triggerType")
    survey_id = models.TextField(null=True, db_column="surveyId")
    is_dry_run = models.BooleanField(db_column="isDryRun", default=False)
    idempotency_key = models.TextField(null=True, db_column="idempotencyKey")
    attempt = models.IntegerField(default=0)
    next_attempt_at = UtcDateTimeField(null=True, db_column="nextAttemptAt")
    last_error_at = UtcDateTimeField(null=True, db_column="lastErrorAt")
    trigger_payload = models.JSONField(db_column="triggerPayload")
    """[WorkflowTriggerRunPayload]"""
    data = models.JSONField(default=dict)
    """[WorkflowRunData]"""
    error = models.TextField(null=True)
    dispatched_at = UtcDateTimeField(null=True, db_column="dispatchedAt")
    started_at = UtcDateTimeField(null=True, db_column="startedAt")
    finished_at = UtcDateTimeField(null=True, db_column="finishedAt")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "WorkflowRun"
        constraints = [
            models.UniqueConstraint(fields=["workflow_id", "idempotency_key"], name="WorkflowRun_workflowId_idempotencyKey_key"),
        ]
        indexes = [
            models.Index(fields=["workspace", "created_at"]),
            models.Index(fields=["workflow_id", "workspace", "created_at"]),
            models.Index(fields=["status", "created_at"]),
            models.Index(fields=["response"]),
            models.Index(fields=["is_dry_run"]),
            models.Index(fields=["next_attempt_at"]),
        ]


class WorkflowRunLog(models.Model):
    id = models.TextField(primary_key=True, default=None)
    run = models.ForeignKey("WorkflowRun", on_delete=models.CASCADE, related_name="logs", db_column="runId", db_index=False)
    sequence = models.IntegerField()
    step_id = models.TextField(db_column="stepId")
    step_type = models.TextField(db_column="stepType")
    status = models.TextField(choices=WorkflowRunLogStatus.choices, default=WorkflowRunLogStatus.pending)
    input = models.JSONField(default=dict)
    """[WorkflowRunLogInput]"""
    output = models.JSONField(default=dict)
    """[WorkflowRunLogOutput]"""
    error = models.TextField(null=True)
    started_at = UtcDateTimeField(null=True, db_column="startedAt")
    finished_at = UtcDateTimeField(null=True, db_column="finishedAt")

    class Meta:
        db_table = "WorkflowRunLog"
        constraints = [
            models.UniqueConstraint(fields=["run", "step_id"], name="WorkflowRunLog_runId_stepId_key"),
        ]
        indexes = [
            models.Index(fields=["run", "sequence"]),
        ]
