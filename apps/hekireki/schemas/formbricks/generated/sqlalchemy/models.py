from sqlalchemy import ARRAY, BigInteger, Date, Dialect, Enum, ForeignKey, ForeignKeyConstraint, Index, MetaData, Numeric, Text, TypeDecorator, text
from sqlalchemy.dialects.postgresql import JSONB, TIMESTAMP
from sqlalchemy.orm import DeclarativeBase, Mapped, mapped_column, relationship
from typing import Any, Optional
from decimal import Decimal as DecimalType
from datetime import datetime, date, timezone
import uuid as uuid_mod
import cuid  # type: ignore[import-untyped, unused-ignore]
import cuid2


class UtcDateTime(TypeDecorator[datetime]):
    impl = TIMESTAMP
    cache_ok = True

    def process_bind_param(self, value: Optional[datetime], dialect: Dialect) -> Optional[datetime]:
        if value is None or value.tzinfo is None:
            return value
        return value.astimezone(timezone.utc).replace(tzinfo=None)

    def process_result_value(self, value: Optional[datetime], dialect: Dialect) -> Optional[datetime]:
        return None if value is None else value.replace(tzinfo=timezone.utc)


def utc_now() -> datetime:
    now = datetime.now(timezone.utc)
    return now.replace(microsecond=now.microsecond // 1000 * 1000)


class Base(DeclarativeBase):
    metadata = MetaData(naming_convention={"pk": "%(table_name)s_pkey"})
    type_annotation_map = {datetime: UtcDateTime(precision=3), str: Text, DecimalType: Numeric(precision=65, scale=30)}


class Creative(Base):
    """A feedback button or a banner. Opens at most one Destination: a survey (`surveyId`) or a link
    (`linkUrl`), never both.

    @property type - Fixed at creation.
    @property content - Per-type JSON (`FeedbackButtonContent` / `BannerContent` in the v3 spec), texts keyed by locale code.
    @property styling - Per-type JSON (`FeedbackButtonStyling` / `BannerStyling` in the v3 spec).
    @property survey - The Destination survey, in the same workspace by construction of the composite key.
    @property showConditions - Show when (`ZShowConditions`, shared with surveys); null always shows.
    @property autoSelectLanguage - The browser's language picks the texts' language when the site sets none; null (never chosen, as on Survey) is off.
    """

    __tablename__ = "Creative"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    name: Mapped[str]
    type: Mapped[str] = mapped_column(Enum("feedbackButton", "banner", name="CreativeType"))
    status: Mapped[str] = mapped_column(Enum("draft", "inProgress", "paused", "completed", name="CreativeStatus"), default="draft", server_default=text("'draft'"))
    published_at: Mapped[Optional[datetime]] = mapped_column("publishedAt")
    archived_at: Mapped[Optional[datetime]] = mapped_column("archivedAt")
    workspace_id: Mapped[str] = mapped_column("workspaceId", ForeignKey("Workspace.id", ondelete="CASCADE", onupdate="CASCADE", name="Creative_workspaceId_fkey"))
    created_by: Mapped[Optional[str]] = mapped_column("createdBy", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="Creative_createdBy_fkey"))
    content: Mapped[dict[str, Any]] = mapped_column(JSONB)
    styling: Mapped[dict[str, Any]] = mapped_column(JSONB)
    survey_id: Mapped[Optional[str]] = mapped_column("surveyId")
    link_url: Mapped[Optional[str]] = mapped_column("linkUrl")
    show_on: Mapped[str] = mapped_column("showOn", Enum("pageView", "action", name="CreativeShowOn"), default="pageView", server_default=text("'pageView'"))
    show_conditions: Mapped[Optional[dict[str, Any]]] = mapped_column("showConditions", JSONB(none_as_null=True))
    display_percentage: Mapped[Optional[float]] = mapped_column("displayPercentage")
    start_at: Mapped[Optional[datetime]] = mapped_column("startAt")
    end_at: Mapped[Optional[datetime]] = mapped_column("endAt")
    auto_select_language: Mapped[Optional[bool]] = mapped_column("autoSelectLanguage")

    __table_args__ = (
        Index("Creative_workspaceId_updated_at_idx", "workspaceId", "updated_at"),
        Index("Creative_workspaceId_created_at_idx", "workspaceId", "created_at"),
        Index("Creative_workspaceId_name_idx", "workspaceId", "name"),
        Index("Creative_surveyId_idx", "surveyId"),
        ForeignKeyConstraint(["surveyId", "workspaceId"], ["Survey.id", "Survey.workspaceId"], ondelete="NO ACTION", onupdate="NO ACTION", name="Creative_surveyId_workspaceId_fkey"),
    )

    workspace: Mapped["Workspace"] = relationship(overlaps="creatives,survey", back_populates="creatives")
    creator: Mapped[Optional["User"]] = relationship(back_populates="creatives")
    survey: Mapped[Optional["Survey"]] = relationship(overlaps="creatives,workspace", back_populates="creatives")
    triggers: Mapped[list["CreativeTrigger"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="creative")
    languages: Mapped[list["CreativeLanguage"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="creative")
    daily_stats: Mapped[list["CreativeDailyStat"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="creative")
    daily_page_stats: Mapped[list["CreativeDailyPageStat"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="creative")

class CreativeDailyStat(Base):
    """A creative's Impressions and Clicks on one UTC day on one device type (ENG-4000). Written
    only by the counter flush job as `GREATEST(stored, counted)`, from Redis counters that hold the day's
    running totals, so a repeated flush changes nothing. At most 3 rows per creative per day. Deleting
    the creative deletes its counts.
    """

    __tablename__ = "CreativeDailyStat"

    creative_id: Mapped[str] = mapped_column("creativeId", ForeignKey("Creative.id", ondelete="CASCADE", onupdate="CASCADE", name="CreativeDailyStat_creativeId_fkey"), primary_key=True)
    date_: Mapped[date] = mapped_column("date", Date, primary_key=True)
    """UTC day."""
    device_type: Mapped[str] = mapped_column("deviceType", Enum("phone", "tablet", "desktop", name="CreativeDeviceType"), primary_key=True)
    impressions: Mapped[int] = mapped_column(default=0, server_default=text("0"))
    clicks: Mapped[int] = mapped_column(default=0, server_default=text("0"))

    creative: Mapped["Creative"] = relationship(back_populates="daily_stats")

class CreativeDailyPageStat(Base):
    """A creative's Impressions and Clicks on one UTC day on one page (host + path), written like
    `CreativeDailyStat`. At most 1000 pages per creative per day; the rest are counted under `(other)`.
    """

    __tablename__ = "CreativeDailyPageStat"

    creative_id: Mapped[str] = mapped_column("creativeId", ForeignKey("Creative.id", ondelete="CASCADE", onupdate="CASCADE", name="CreativeDailyPageStat_creativeId_fkey"), primary_key=True)
    date_: Mapped[date] = mapped_column("date", Date, primary_key=True)
    """UTC day."""
    page: Mapped[str] = mapped_column(primary_key=True)
    """Host + path, without query string or hash, or `(other)`."""
    impressions: Mapped[int] = mapped_column(default=0, server_default=text("0"))
    clicks: Mapped[int] = mapped_column(default=0, server_default=text("0"))

    creative: Mapped["Creative"] = relationship(back_populates="daily_page_stats")

class CreativeTrigger(Base):
    """Links a creative to the actions that show it (`showOn: action`). Mirrors `SurveyTrigger`."""

    __tablename__ = "CreativeTrigger"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    creative_id: Mapped[str] = mapped_column("creativeId", ForeignKey("Creative.id", ondelete="CASCADE", onupdate="CASCADE", name="CreativeTrigger_creativeId_fkey"))
    action_class_id: Mapped[str] = mapped_column("actionClassId", ForeignKey("ActionClass.id", ondelete="CASCADE", onupdate="CASCADE", name="CreativeTrigger_actionClassId_fkey"))

    __table_args__ = (
        Index("CreativeTrigger_actionClassId_idx", "actionClassId"),
        Index("CreativeTrigger_creativeId_actionClassId_key", "creativeId", "actionClassId", unique=True),
    )

    creative: Mapped["Creative"] = relationship(back_populates="triggers")
    action_class: Mapped["ActionClass"] = relationship(back_populates="creative_triggers")

class CreativeLanguage(Base):
    """The languages a creative's texts are written in. Mirrors `SurveyLanguage`; exactly one is the default."""

    __tablename__ = "CreativeLanguage"

    language_id: Mapped[str] = mapped_column("languageId", ForeignKey("Language.id", ondelete="CASCADE", onupdate="CASCADE", name="CreativeLanguage_languageId_fkey"), primary_key=True)
    creative_id: Mapped[str] = mapped_column("creativeId", ForeignKey("Creative.id", ondelete="CASCADE", onupdate="CASCADE", name="CreativeLanguage_creativeId_fkey"), primary_key=True)
    default: Mapped[bool] = mapped_column(default=False, server_default=text("false"))
    enabled: Mapped[bool] = mapped_column(default=True, server_default=text("true"))

    __table_args__ = (
        Index("CreativeLanguage_creativeId_idx", "creativeId"),
    )

    language: Mapped["Language"] = relationship(back_populates="creative_languages")
    creative: Mapped["Creative"] = relationship(back_populates="languages")

class OutboundEmailAudit(Base):
    """Metadata for an outbound email attempt. Content and recipient addresses are never stored here.
    Identifiers are deliberately not foreign keys so erasure can scrub a row without retaining source data.
    """

    __tablename__ = "OutboundEmailAudit"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: str(uuid_mod.uuid4()))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)
    expires_at: Mapped[datetime] = mapped_column("expiresAt")
    content_expires_at: Mapped[datetime] = mapped_column("contentExpiresAt")
    environment: Mapped[str]
    email_type: Mapped[str] = mapped_column("emailType")
    trigger: Mapped[str]
    actor_kind: Mapped[str] = mapped_column("actorKind")
    actor_user_id: Mapped[Optional[str]] = mapped_column("actorUserId")
    configured_by_user_id: Mapped[Optional[str]] = mapped_column("configuredByUserId")
    actor_user_hash: Mapped[Optional[str]] = mapped_column("actorUserHash")
    configured_by_user_hash: Mapped[Optional[str]] = mapped_column("configuredByUserHash")
    organization_id: Mapped[Optional[str]] = mapped_column("organizationId")
    survey_id: Mapped[Optional[str]] = mapped_column("surveyId")
    response_id: Mapped[Optional[str]] = mapped_column("responseId")
    follow_up_id: Mapped[Optional[str]] = mapped_column("followUpId")
    workflow_id: Mapped[Optional[str]] = mapped_column("workflowId")
    workflow_run_id: Mapped[Optional[str]] = mapped_column("workflowRunId")
    invite_id: Mapped[Optional[str]] = mapped_column("inviteId")
    recipient_mac: Mapped[Optional[str]] = mapped_column("recipientMac")
    ses_message_hash: Mapped[Optional[str]] = mapped_column("sesMessageHash")
    snapshot_key: Mapped[Optional[str]] = mapped_column("snapshotKey")
    snapshot_sha256: Mapped[Optional[str]] = mapped_column("snapshotSha256")
    status: Mapped[str] = mapped_column(default="prepared", server_default=text("'prepared'"))
    status_recorded_at: Mapped[Optional[datetime]] = mapped_column("statusRecordedAt")
    audit_gap: Mapped[bool] = mapped_column("auditGap", default=False, server_default=text("false"))
    erasure_requested_at: Mapped[Optional[datetime]] = mapped_column("erasureRequestedAt")
    erased_at: Mapped[Optional[datetime]] = mapped_column("erasedAt")
    content_delete_retry_at: Mapped[Optional[datetime]] = mapped_column("contentDeleteRetryAt")
    content_delete_failures: Mapped[int] = mapped_column("contentDeleteFailures", default=0, server_default=text("0"))

    __table_args__ = (
        Index("OutboundEmailAudit_environment_createdAt_id_idx", "environment", "createdAt", "id"),
        Index("OutboundEmailAudit_organizationId_createdAt_id_idx", "organizationId", "createdAt", "id"),
        Index("OutboundEmailAudit_surveyId_createdAt_id_idx", "surveyId", "createdAt", "id"),
        Index("OutboundEmailAudit_responseId_createdAt_id_idx", "responseId", "createdAt", "id"),
        Index("OutboundEmailAudit_actorUserId_createdAt_id_idx", "actorUserId", "createdAt", "id"),
        Index("OutboundEmailAudit_configuredByUserId_createdAt_id_idx", "configuredByUserId", "createdAt", "id"),
        Index("OutboundEmailAudit_actorUserHash_createdAt_id_idx", "actorUserHash", "createdAt", "id"),
        Index("OutboundEmailAudit_configuredByUserHash_createdAt_id_idx", "configuredByUserHash", "createdAt", "id"),
        Index("OutboundEmailAudit_recipientMac_createdAt_id_idx", "recipientMac", "createdAt", "id"),
        Index("OutboundEmailAudit_sesMessageHash_idx", "sesMessageHash"),
        Index("OutboundEmailAudit_status_createdAt_id_idx", "status", "createdAt", "id"),
        Index("OutboundEmailAudit_contentExpiresAt_id_idx", "contentExpiresAt", "id"),
        Index("OutboundEmailAudit_expiresAt_id_idx", "expiresAt", "id"),
        Index("OutboundEmailAudit_erasureRequestedAt_id_idx", "erasureRequestedAt", "id"),
        Index("OutboundEmailAudit_contentDeleteRetryAt_id_idx", "contentDeleteRetryAt", "id"),
    )

class OutboundEmailAuditErasureIntent(Base):
    """Holds an account erasure request across an in-flight email send. The address is never stored.
    The user id is removed once deletion is confirmed; one-way recovery keys expire after 90 days.
    """

    __tablename__ = "OutboundEmailAuditErasureIntent"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: str(uuid_mod.uuid4()))
    user_id: Mapped[Optional[str]] = mapped_column("userId")
    user_hash: Mapped[str] = mapped_column("userHash")
    recipient_mac: Mapped[str] = mapped_column("recipientMac")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    activated_at: Mapped[Optional[datetime]] = mapped_column("activatedAt")
    checked_at: Mapped[Optional[datetime]] = mapped_column("checkedAt")
    expires_at: Mapped[datetime] = mapped_column("expiresAt")

    __table_args__ = (
        Index("OutboundEmailAuditErasureIntent_userId_idx", "userId"),
        Index("OutboundEmailAuditErasureIntent_recipientMac_activatedAt_idx", "recipientMac", "activatedAt"),
        Index("OutboundEmailAuditErasureIntent_userHash_activatedAt_idx", "userHash", "activatedAt"),
        Index("OutboundEmailAuditErasureIntent_checkedAt_id_idx", "checkedAt", "id"),
        Index("OutboundEmailAuditErasureIntent_expiresAt_idx", "expiresAt"),
    )

class Webhook(Base):
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

    __tablename__ = "Webhook"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    name: Mapped[Optional[str]]
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    url: Mapped[str]
    source: Mapped[str] = mapped_column(Enum("user", "zapier", "make", "n8n", "activepieces", name="WebhookSource"), default="user", server_default=text("'user'"))
    workspace_id: Mapped[str] = mapped_column("workspaceId", ForeignKey("Workspace.id", ondelete="CASCADE", onupdate="CASCADE", name="Webhook_workspaceId_fkey"))
    triggers: Mapped[list[str]] = mapped_column(ARRAY(Enum("responseCreated", "responseUpdated", "responseFinished", name="PipelineTriggers")), nullable=True)
    survey_ids: Mapped[list[str]] = mapped_column("surveyIds", ARRAY(Text), nullable=True)
    secret: Mapped[Optional[str]]

    __table_args__ = (
        Index("Webhook_workspaceId_idx", "workspaceId"),
    )

    workspace: Mapped["Workspace"] = relationship(back_populates="webhooks")

class ContactAttribute(Base):
    """Represents an attribute value associated with a contact.
    Used to store custom properties and characteristics of contacts.

    @property id - Unique identifier for the attribute
    @property attributeKey - Reference to the attribute definition
    @property contact - The contact this attribute belongs to
    @property value - The string value of the attribute (used for string type + backwards compatibility)
    @property valueNumber - Native numeric storage for number type attributes
    @property valueDate - Native date storage for date type attributes
    """

    __tablename__ = "ContactAttribute"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    attribute_key_id: Mapped[str] = mapped_column("attributeKeyId", ForeignKey("ContactAttributeKey.id", ondelete="CASCADE", onupdate="CASCADE", name="ContactAttribute_attributeKeyId_fkey"))
    contact_id: Mapped[str] = mapped_column("contactId", ForeignKey("Contact.id", ondelete="CASCADE", onupdate="CASCADE", name="ContactAttribute_contactId_fkey"))
    value: Mapped[str]
    value_number: Mapped[Optional[float]] = mapped_column("valueNumber")
    value_date: Mapped[Optional[datetime]] = mapped_column("valueDate")

    __table_args__ = (
        Index("ContactAttribute_attributeKeyId_value_idx", "attributeKeyId", "value"),
        Index("ContactAttribute_attributeKeyId_valueNumber_idx", "attributeKeyId", "valueNumber"),
        Index("ContactAttribute_attributeKeyId_valueDate_idx", "attributeKeyId", "valueDate"),
        Index("ContactAttribute_contactId_attributeKeyId_key", "contactId", "attributeKeyId", unique=True),
    )

    attribute_key: Mapped["ContactAttributeKey"] = relationship(back_populates="attributes")
    contact: Mapped["Contact"] = relationship(back_populates="attributes")

class ContactAttributeKey(Base):
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

    __tablename__ = "ContactAttributeKey"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    is_unique: Mapped[bool] = mapped_column("isUnique", default=False, server_default=text("false"))
    key: Mapped[str]
    name: Mapped[Optional[str]]
    description: Mapped[Optional[str]]
    type: Mapped[str] = mapped_column(Enum("default", "custom", name="ContactAttributeType"), default="custom", server_default=text("'custom'"))
    data_type: Mapped[str] = mapped_column("dataType", Enum("string", "number", "date", name="ContactAttributeDataType"), default="string", server_default=text("'string'"))
    workspace_id: Mapped[str] = mapped_column("workspaceId", ForeignKey("Workspace.id", ondelete="CASCADE", onupdate="CASCADE", name="ContactAttributeKey_workspaceId_fkey"))

    __table_args__ = (
        Index("ContactAttributeKey_workspaceId_created_at_idx", "workspaceId", "created_at"),
        Index("ContactAttributeKey_key_workspaceId_key", "key", "workspaceId", unique=True),
    )

    workspace: Mapped["Workspace"] = relationship(back_populates="contact_attribute_keys")
    attributes: Mapped[list["ContactAttribute"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="attribute_key")
    attribute_filters: Mapped[list["SurveyAttributeFilter"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="attribute_key")

class Contact(Base):
    """Represents a person or user who can receive and respond to surveys.
    Contacts are workspace-specific and can have multiple attributes and responses.

    @property id - Unique identifier for the contact
    @property workspace - The workspace this contact belongs to
    @property responses - Survey responses from this contact
    @property attributes - Custom attributes associated with this contact
    @property displays - Record of surveys shown to this contact
    """

    __tablename__ = "Contact"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    workspace_id: Mapped[str] = mapped_column("workspaceId", ForeignKey("Workspace.id", ondelete="CASCADE", onupdate="CASCADE", name="Contact_workspaceId_fkey"))

    __table_args__ = (
        Index("Contact_workspaceId_idx", "workspaceId"),
    )

    workspace: Mapped["Workspace"] = relationship(back_populates="contacts")
    responses: Mapped[list["Response"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="contact")
    attributes: Mapped[list["ContactAttribute"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="contact")
    displays: Mapped[list["Display"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="contact")

class Response(Base):
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

    __tablename__ = "Response"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    finished: Mapped[bool] = mapped_column(default=False, server_default=text("false"))
    survey_id: Mapped[str] = mapped_column("surveyId", ForeignKey("Survey.id", ondelete="CASCADE", onupdate="CASCADE", name="Response_surveyId_fkey"))
    contact_id: Mapped[Optional[str]] = mapped_column("contactId", ForeignKey("Contact.id", ondelete="CASCADE", onupdate="CASCADE", name="Response_contactId_fkey"))
    ending_id: Mapped[Optional[str]] = mapped_column("endingId")
    data: Mapped[dict[str, Any]] = mapped_column(JSONB, default=lambda: {}, server_default=text("'{}'"))
    """[ResponseData]"""
    variables: Mapped[dict[str, Any]] = mapped_column(JSONB, default=lambda: {}, server_default=text("'{}'"))
    """[ResponseVariables]"""
    ttc: Mapped[dict[str, Any]] = mapped_column(JSONB, default=lambda: {}, server_default=text("'{}'"))
    """[ResponseTtc]"""
    meta: Mapped[dict[str, Any]] = mapped_column(JSONB, default=lambda: {}, server_default=text("'{}'"))
    """[ResponseMeta]"""
    ingest_flags: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))
    """Why a value in `data` will not read back as its declared Embedded Data type — a failed coercion
    or a truncated oversize value (ENG-1845). Server-computed on every ingest and never accepted
    from the client. Null means no ingest boundary has written it, which is not the same claim as
    an empty list: without this column a wrong-typed value is indistinguishable from an absent one,
    and surfacing that in the UI (v2) would have nothing to read.
    [ResponseIngestFlags]
    """
    outcome: Mapped[Optional[str]] = mapped_column(Enum("lowQuality", name="ResponseOutcome"))
    """A reviewer's verdict on the response; null means none. `lowQuality` hides it from the Analysis
    export by default and from Feedback Records, and nothing else: it is never deleted and every
    count, list, webhook and integration still includes it (ENG-6780).
    """
    not_shown: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))
    """What the respondent was not shown (ENG-6936, ADR 0029): the blocks and elements a show-if kept
    from them, and per element its options, matrix rows and columns, by stable ids. Written by the
    survey while they answer; null means unknown (older responses and bundles), read as before.
    Omitted from reads without a `select` (PRISMA_GLOBAL_OMIT), so it never reaches a v1/v2 payload,
    webhook or integration.
    [ResponseNotShown]
    """
    contact_attributes: Mapped[Optional[dict[str, Any]]] = mapped_column("contactAttributes", JSONB(none_as_null=True))
    """[ResponseContactAttributes]"""
    single_use_id: Mapped[Optional[str]] = mapped_column("singleUseId")
    language: Mapped[Optional[str]]
    display_id: Mapped[Optional[str]] = mapped_column("displayId", ForeignKey("Display.id", ondelete="SET NULL", onupdate="CASCADE", name="Response_displayId_fkey"))

    __table_args__ = (
        Index("Response_displayId_key", "displayId", unique=True),
        Index("Response_created_at_idx", "created_at"),
        Index("Response_created_at_id_idx", "created_at", "id"),
        Index("Response_surveyId_created_at_idx", "surveyId", "created_at"),
        Index("Response_surveyId_created_at_id_idx", "surveyId", "created_at", "id"),
        Index("Response_contactId_created_at_idx", "contactId", "created_at"),
        Index("Response_contactId_created_at_id_idx", "contactId", "created_at", "id"),
        Index("Response_surveyId_finished_idx", "surveyId", "finished"),
        Index("Response_surveyId_singleUseId_key", "surveyId", "singleUseId", unique=True),
    )

    survey: Mapped["Survey"] = relationship(back_populates="responses")
    contact: Mapped[Optional["Contact"]] = relationship(back_populates="responses")
    display: Mapped[Optional["Display"]] = relationship(back_populates="response")
    tags: Mapped[list["TagsOnResponses"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="response")
    quota_links: Mapped[list["ResponseQuotaLink"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="response")
    workflow_runs: Mapped[list["WorkflowRun"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="response")

class Tag(Base):
    """Represents a label that can be applied to survey responses.
    Used for categorizing and organizing responses within a workspace.

    @property id - Unique identifier for the tag
    @property name - Display name of the tag
    @property responses - Survey responses tagged with this label
    @property workspace - The workspace this tag belongs to
    """

    __tablename__ = "Tag"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    name: Mapped[str]
    workspace_id: Mapped[str] = mapped_column("workspaceId", ForeignKey("Workspace.id", ondelete="CASCADE", onupdate="CASCADE", name="Tag_workspaceId_fkey"))

    __table_args__ = (
        Index("Tag_workspaceId_idx", "workspaceId"),
        Index("Tag_workspaceId_name_key", "workspaceId", "name", unique=True),
    )

    workspace: Mapped["Workspace"] = relationship(back_populates="tags")
    responses: Mapped[list["TagsOnResponses"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="tag")

class TagsOnResponses(Base):
    """Junction table linking tags to responses.
    Enables many-to-many relationship between tags and responses.

    @property response - The tagged response
    @property tag - The tag applied to the response
    """

    __tablename__ = "TagsOnResponses"

    response_id: Mapped[str] = mapped_column("responseId", ForeignKey("Response.id", ondelete="CASCADE", onupdate="CASCADE", name="TagsOnResponses_responseId_fkey"), primary_key=True)
    tag_id: Mapped[str] = mapped_column("tagId", ForeignKey("Tag.id", ondelete="CASCADE", onupdate="CASCADE", name="TagsOnResponses_tagId_fkey"), primary_key=True)

    response: Mapped["Response"] = relationship(back_populates="tags")
    tag: Mapped["Tag"] = relationship(back_populates="responses")

class Display(Base):
    """Records when a survey is shown to a user.
    Tracks survey display history and response status.

    @property id - Unique identifier for the display event
    @property survey - The survey that was displayed
    @property contact - The contact who saw the survey
    @property response - The associated response if one exists
    @property creativeId - The creative that opened the survey. No foreign key, so the record survives the creative's deletion.
    """

    __tablename__ = "Display"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    survey_id: Mapped[str] = mapped_column("surveyId", ForeignKey("Survey.id", ondelete="CASCADE", onupdate="CASCADE", name="Display_surveyId_fkey"))
    contact_id: Mapped[Optional[str]] = mapped_column("contactId", ForeignKey("Contact.id", ondelete="CASCADE", onupdate="CASCADE", name="Display_contactId_fkey"))
    creative_id: Mapped[Optional[str]] = mapped_column("creativeId")

    __table_args__ = (
        Index("Display_surveyId_idx", "surveyId"),
        Index("Display_contactId_created_at_idx", "contactId", "created_at"),
    )

    survey: Mapped["Survey"] = relationship(back_populates="displays")
    contact: Mapped[Optional["Contact"]] = relationship(back_populates="displays")
    response: Mapped[Optional["Response"]] = relationship(back_populates="display")

class SurveyTrigger(Base):
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

    __tablename__ = "SurveyTrigger"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    survey_id: Mapped[str] = mapped_column("surveyId", ForeignKey("Survey.id", ondelete="CASCADE", onupdate="CASCADE", name="SurveyTrigger_surveyId_fkey"))
    action_class_id: Mapped[str] = mapped_column("actionClassId", ForeignKey("ActionClass.id", ondelete="CASCADE", onupdate="CASCADE", name="SurveyTrigger_actionClassId_fkey"))

    __table_args__ = (
        Index("SurveyTrigger_surveyId_actionClassId_key", "surveyId", "actionClassId", unique=True),
    )

    survey: Mapped["Survey"] = relationship(back_populates="triggers")
    action_class: Mapped["ActionClass"] = relationship(back_populates="survey_triggers")

class SurveyAttributeFilter(Base):
    """Defines targeting rules for surveys based on contact attributes.
    Used to show surveys only to contacts matching specific criteria.

    @property id - Unique identifier for the filter
    @property attributeKey - The contact attribute to filter on
    @property survey - The survey being filtered
    @property condition - The comparison operator to use
    @property value - The value to compare against
    """

    __tablename__ = "SurveyAttributeFilter"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    attribute_key_id: Mapped[str] = mapped_column("attributeKeyId", ForeignKey("ContactAttributeKey.id", ondelete="CASCADE", onupdate="CASCADE", name="SurveyAttributeFilter_attributeKeyId_fkey"))
    survey_id: Mapped[str] = mapped_column("surveyId", ForeignKey("Survey.id", ondelete="CASCADE", onupdate="CASCADE", name="SurveyAttributeFilter_surveyId_fkey"))
    condition: Mapped[str] = mapped_column(Enum("equals", "notEquals", name="SurveyAttributeFilterCondition"))
    value: Mapped[str]

    __table_args__ = (
        Index("SurveyAttributeFilter_attributeKeyId_idx", "attributeKeyId"),
        Index("SurveyAttributeFilter_surveyId_attributeKeyId_key", "surveyId", "attributeKeyId", unique=True),
    )

    attribute_key: Mapped["ContactAttributeKey"] = relationship(back_populates="attribute_filters")
    survey: Mapped["Survey"] = relationship(back_populates="attribute_filters")

class Survey(Base):
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

    __tablename__ = "Survey"

    hidden_fields: Mapped[dict[str, Any]] = mapped_column("hiddenFields", JSONB, default=lambda: {"enabled": False}, server_default=text("'{\"enabled\": false}'"))
    """[SurveyHiddenFields]"""
    variables: Mapped[dict[str, Any]] = mapped_column(JSONB, default=lambda: [], server_default=text("'[]'"))
    """[SurveyVariables]"""
    ownership_initialized: Mapped[bool] = mapped_column("ownershipInitialized", default=False, server_default=text("false"))
    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    name: Mapped[str]
    redirect_url: Mapped[Optional[str]] = mapped_column("redirectUrl")
    type: Mapped[str] = mapped_column(Enum("link", "app", name="SurveyType"), default="app", server_default=text("'app'"))
    workspace_id: Mapped[str] = mapped_column("workspaceId", ForeignKey("Workspace.id", ondelete="CASCADE", onupdate="CASCADE", name="Survey_workspaceId_fkey"))
    created_by: Mapped[Optional[str]] = mapped_column("createdBy", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="Survey_createdBy_fkey"))
    visibility: Mapped[str] = mapped_column(Enum("restricted", "workspace", name="SurveyVisibility"), default="workspace", server_default=text("'workspace'"))
    """ENG-3282: authorization facts, projected into SpiceDB by the survey outbox trigger. `ownerId` is
    the durable ownership source; `createdBy` stays pure attribution and is never rewritten.
    """
    owner_id: Mapped[Optional[str]] = mapped_column("ownerId", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="Survey_ownerId_fkey"))
    visibility_version: Mapped[int] = mapped_column("visibilityVersion", default=0, server_default=text("0"))
    """Incremented by every accepted visibility change; fences projection work. An INSERT starts at 0
    with `visibilityProjectedVersion` -1 (the `survey_visibility_pending` trigger): the initial
    projection, decided from PostgreSQL facts until the projector acknowledges it.
    """
    visibility_projected_version: Mapped[int] = mapped_column("visibilityProjectedVersion", default=0, server_default=text("0"))
    """The `visibilityVersion` SpiceDB last acknowledged; -1 while none ever was."""
    visibility_pending: Mapped[bool] = mapped_column("visibilityPending", default=False, server_default=text("false"))
    """"visibilityVersion" <> "visibilityProjectedVersion", kept by the `survey_visibility_pending`
    trigger (a BEFORE trigger overrides whatever an INSERT or UPDATE supplies). Read-only: never put it
    in create/update data.
    """
    visibility_changed_at: Mapped[Optional[datetime]] = mapped_column("visibilityChangedAt")
    visibility_changed_by_id: Mapped[Optional[str]] = mapped_column("visibilityChangedById")
    status: Mapped[str] = mapped_column(Enum("draft", "inProgress", "paused", "completed", name="SurveyStatus"), default="draft", server_default=text("'draft'"))
    welcome_card: Mapped[dict[str, Any]] = mapped_column("welcomeCard", JSONB, default=lambda: {"enabled": False}, server_default=text("'{\"enabled\": false}'"))
    """[SurveyWelcomeCard]"""
    questions: Mapped[dict[str, Any]] = mapped_column(JSONB, default=lambda: [], server_default=text("'[]'"))
    """[SurveyQuestions]"""
    blocks: Mapped[list[dict[str, Any]]] = mapped_column(ARRAY(JSONB), nullable=True, default=lambda: [], server_default=text("ARRAY[]::jsonb[]"))
    """[SurveyBlocks]"""
    endings: Mapped[list[dict[str, Any]]] = mapped_column(ARRAY(JSONB), nullable=True, default=lambda: [], server_default=text("ARRAY[]::jsonb[]"))
    """[SurveyEnding]"""
    display_option: Mapped[str] = mapped_column("displayOption", Enum("displayOnce", "displayMultiple", "displaySome", "respondMultiple", name="displayOptions"), default="displayOnce", server_default=text("'displayOnce'"))
    recontact_days: Mapped[Optional[int]] = mapped_column("recontactDays")
    display_limit: Mapped[Optional[int]] = mapped_column("displayLimit")
    inline_triggers: Mapped[Optional[dict[str, Any]]] = mapped_column("inlineTriggers", JSONB(none_as_null=True))
    """[SurveyInlineTriggers]"""
    auto_close: Mapped[Optional[int]] = mapped_column("autoClose")
    keep_ending_card_open: Mapped[bool] = mapped_column("keepEndingCardOpen", default=False, server_default=text("false"))
    auto_complete: Mapped[Optional[int]] = mapped_column("autoComplete")
    delay: Mapped[int] = mapped_column(default=0, server_default=text("0"))
    publish_on: Mapped[Optional[datetime]] = mapped_column("publishOn")
    close_on: Mapped[Optional[datetime]] = mapped_column("closeOn")
    archived_at: Mapped[Optional[datetime]] = mapped_column("archivedAt")
    survey_closed_message: Mapped[Optional[dict[str, Any]]] = mapped_column("surveyClosedMessage", JSONB(none_as_null=True))
    """[SurveyClosedMessage]"""
    segment_id: Mapped[Optional[str]] = mapped_column("segmentId", ForeignKey("Segment.id", ondelete="SET NULL", onupdate="CASCADE", name="Survey_segmentId_fkey"))
    workspace_overwrites: Mapped[Optional[dict[str, Any]]] = mapped_column("workspaceOverwrites", JSONB(none_as_null=True))
    """[SurveyWorkspaceOverwrites]"""
    styling: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))
    """[SurveyStyling]"""
    single_use: Mapped[Optional[dict[str, Any]]] = mapped_column("singleUse", JSONB(none_as_null=True), default=lambda: {"enabled": False, "isEncrypted": True}, server_default=text("'{\"enabled\": false, \"isEncrypted\": true}'"))
    """[SurveySingleUse]"""
    is_verify_email_enabled: Mapped[bool] = mapped_column("isVerifyEmailEnabled", default=False, server_default=text("false"))
    is_back_button_hidden: Mapped[bool] = mapped_column("isBackButtonHidden", default=False, server_default=text("false"))
    is_auto_progressing_enabled: Mapped[bool] = mapped_column("isAutoProgressingEnabled", default=False, server_default=text("false"))
    is_capture_ip_enabled: Mapped[bool] = mapped_column("isCaptureIpEnabled", default=False, server_default=text("false"))
    is_anonymize_responses_enabled: Mapped[bool] = mapped_column(default=False, server_default=text("false"))
    pin: Mapped[Optional[str]]
    display_percentage: Mapped[Optional[DecimalType]] = mapped_column("displayPercentage")
    show_language_switch: Mapped[Optional[bool]] = mapped_column("showLanguageSwitch")
    auto_select_language: Mapped[Optional[bool]] = mapped_column("autoSelectLanguage")
    """Open the survey in the respondent's browser language when it matches an enabled survey
    language. Null on surveys that predate the setting, which reads as off.
    """
    recaptcha: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True), default=lambda: {"enabled": False, "threshold": 0.1}, server_default=text("'{\"enabled\": false, \"threshold\":0.1}'"))
    """[SurveyRecaptcha]"""
    metadata_: Mapped[dict[str, Any]] = mapped_column("metadata", JSONB, default=lambda: {}, server_default=text("'{}'"))
    """[SurveyLinkMetadata]"""
    slug: Mapped[Optional[str]]
    custom_head_scripts: Mapped[Optional[str]] = mapped_column("customHeadScripts")
    custom_head_scripts_mode: Mapped[Optional[str]] = mapped_column("customHeadScriptsMode", Enum("add", "replace", name="SurveyScriptMode"), default="add", server_default=text("'add'"))
    custom_css: Mapped[Optional[dict[str, Any]]] = mapped_column("customCss", JSONB(none_as_null=True))
    """[CustomCssStored] Survey custom CSS (ENG-2949): source plus trusted compiled output, written only
    by the shared custom CSS save service. null = no survey CSS. Kept outside `styling` (M2.07).
    """
    show_conditions: Mapped[Optional[dict[str, Any]]] = mapped_column("showConditions", JSONB(none_as_null=True))
    """[SurveyShowConditions] App surveys: conditions checked in the browser when a trigger fires
    (ENG-3873). null = no conditions, which is what every existing survey keeps.
    """
    start_logic: Mapped[Optional[dict[str, Any]]] = mapped_column("startLogic", JSONB(none_as_null=True))
    """[SurveyStartLogic] Rules before the first question (ENG-6932): block-rule format, run by the
    survey on load. null = none, which is what every existing survey keeps.
    """

    __table_args__ = (
        Index("Survey_slug_key", "slug", unique=True),
        Index("Survey_workspaceId_updated_at_idx", "workspaceId", "updated_at"),
        Index("Survey_workspaceId_ownerId_idx", "workspaceId", "ownerId"),
        Index("Survey_segmentId_idx", "segmentId"),
        Index("Survey_status_publishOn_idx", "status", "publishOn"),
        Index("Survey_status_closeOn_idx", "status", "closeOn"),
        Index("Survey_archivedAt_idx", "archivedAt"),
        Index("Survey_id_workspaceId_key", "id", "workspaceId", unique=True),
    )

    workspace: Mapped["Workspace"] = relationship(back_populates="surveys")
    creator: Mapped[Optional["User"]] = relationship(foreign_keys=[created_by], back_populates="surveys")
    owner: Mapped[Optional["User"]] = relationship(foreign_keys=[owner_id], back_populates="owned_surveys")
    segment: Mapped[Optional["Segment"]] = relationship(back_populates="surveys")
    responses: Mapped[list["Response"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="survey")
    quotas: Mapped[list["SurveyQuota"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="survey")
    triggers: Mapped[list["SurveyTrigger"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="survey")
    attribute_filters: Mapped[list["SurveyAttributeFilter"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="survey")
    displays: Mapped[list["Display"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="survey")
    languages: Mapped[list["SurveyLanguage"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="survey")
    embedded_data: Mapped[list["EmbeddedData"]] = relationship(cascade="all, delete", passive_deletes=True, overlaps="embedded_data,workspace", back_populates="survey")
    embedded_data_links: Mapped[list["SurveyEmbeddedData"]] = relationship(cascade="all, delete", passive_deletes=True, overlaps="embedded_data,survey_links", back_populates="survey")
    follow_ups: Mapped[list["SurveyFollowUp"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="survey")
    feedback_source_mappings: Mapped[list["FeedbackSourceFormbricksMapping"]] = relationship(cascade="all, delete", passive_deletes=True, overlaps="feedback_source,formbricks_mappings", back_populates="survey")
    retention_exemptions: Mapped[list["RetentionExemption"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="survey")
    retention_notices: Mapped[list["RetentionNotice"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="survey")
    creatives: Mapped[list["Creative"]] = relationship(passive_deletes="all", overlaps="creatives,workspace", back_populates="survey")
    """Creatives whose Destination is this survey (ENG-3993)."""
    attachment_exports: Mapped[list["AttachmentExport"]] = relationship(cascade="all, delete", passive_deletes=True, overlaps="attachment_exports,workspace", back_populates="survey")
    """Background attachment ZIP exports of this survey (ENG-6862)."""
    retention_restore: Mapped[Optional["RetentionSurveyRestore"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="survey")

class SurveyQuota(Base):
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

    __tablename__ = "SurveyQuota"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    survey_id: Mapped[str] = mapped_column("surveyId", ForeignKey("Survey.id", ondelete="CASCADE", onupdate="CASCADE", name="SurveyQuota_surveyId_fkey"))
    name: Mapped[str]
    limit: Mapped[int]
    logic: Mapped[dict[str, Any]] = mapped_column(JSONB, default=lambda: {}, server_default=text("'{}'"))
    """[SurveyQuotaLogic]"""
    action: Mapped[str] = mapped_column(Enum("endSurvey", "continueSurvey", name="SurveyQuotaAction"))
    ending_card_id: Mapped[Optional[str]] = mapped_column("endingCardId")
    count_partial_submissions: Mapped[bool] = mapped_column("countPartialSubmissions", default=False, server_default=text("false"))

    __table_args__ = (
        Index("SurveyQuota_surveyId_name_key", "surveyId", "name", unique=True),
    )

    survey: Mapped["Survey"] = relationship(back_populates="quotas")
    quota_links: Mapped[list["ResponseQuotaLink"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="quota")

class ResponseQuotaLink(Base):
    """Junction table linking responses to quotas.
    Tracks which responses counted towards specific quotas and their status.

    @property response - The response that counted towards the quota
    @property quota - The quota this response counted towards
    @property status - Whether the response was screened in or out
    """

    __tablename__ = "ResponseQuotaLink"

    response_id: Mapped[str] = mapped_column("responseId", ForeignKey("Response.id", ondelete="CASCADE", onupdate="CASCADE", name="ResponseQuotaLink_responseId_fkey"), primary_key=True)
    quota_id: Mapped[str] = mapped_column("quotaId", ForeignKey("SurveyQuota.id", ondelete="CASCADE", onupdate="CASCADE", name="ResponseQuotaLink_quotaId_fkey"), primary_key=True)
    status: Mapped[str] = mapped_column(Enum("screenedIn", "screenedOut", name="ResponseQuotaLinkStatus"))

    __table_args__ = (
        Index("ResponseQuotaLink_quotaId_status_idx", "quotaId", "status"),
    )

    response: Mapped["Response"] = relationship(back_populates="quota_links")
    quota: Mapped["SurveyQuota"] = relationship(back_populates="quota_links")

class SurveyFollowUp(Base):
    """Defines follow-up actions for survey responses.
    Enables automated actions based on specific survey response conditions.

    @property id - Unique identifier for the follow-up
    @property survey - The associated survey
    @property name - Display name for the follow-up
    @property trigger - Conditions that activate the follow-up
    @property action - Actions to take when triggered
    """

    __tablename__ = "SurveyFollowUp"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    survey_id: Mapped[str] = mapped_column("surveyId", ForeignKey("Survey.id", ondelete="CASCADE", onupdate="CASCADE", name="SurveyFollowUp_surveyId_fkey"))
    name: Mapped[str]
    trigger: Mapped[dict[str, Any]] = mapped_column(JSONB)
    """[SurveyFollowUpTrigger]"""
    action: Mapped[dict[str, Any]] = mapped_column(JSONB)
    """[SurveyFollowUpAction]"""

    survey: Mapped["Survey"] = relationship(back_populates="follow_ups")

class ActionClass(Base):
    """Represents a user action that can trigger surveys.
    Used to define points in the user journey where surveys can be shown.

    @property id - Unique identifier for the action
    @property name - Display name of the action
    @property type - Whether this is a code or no-code action
    @property key - Unique identifier used in code implementation
    @property noCodeConfig - Configuration for no-code setup
    @property workspace - The workspace this action belongs to
    """

    __tablename__ = "ActionClass"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    name: Mapped[str]
    description: Mapped[Optional[str]]
    type: Mapped[str] = mapped_column(Enum("code", "noCode", name="ActionType"))
    key: Mapped[Optional[str]]
    no_code_config: Mapped[Optional[dict[str, Any]]] = mapped_column("noCodeConfig", JSONB(none_as_null=True))
    """[ActionClassNoCodeConfig]"""
    workspace_id: Mapped[str] = mapped_column("workspaceId", ForeignKey("Workspace.id", ondelete="CASCADE", onupdate="CASCADE", name="ActionClass_workspaceId_fkey"))

    __table_args__ = (
        Index("ActionClass_workspaceId_created_at_idx", "workspaceId", "created_at"),
        Index("ActionClass_key_workspaceId_key", "key", "workspaceId", unique=True),
        Index("ActionClass_name_workspaceId_key", "name", "workspaceId", unique=True),
    )

    workspace: Mapped["Workspace"] = relationship(back_populates="action_classes")
    survey_triggers: Mapped[list["SurveyTrigger"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="action_class")
    creative_triggers: Mapped[list["CreativeTrigger"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="action_class")

class Integration(Base):
    """Represents third-party service integrations.
    Enables data flow between Formbricks and external services.

    @property id - Unique identifier for the integration
    @property type - The service being integrated (Google Sheets, Notion, etc.)
    @property workspace - The workspace this integration belongs to
    @property config - Service-specific configuration details
    """

    __tablename__ = "Integration"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    type: Mapped[str] = mapped_column(Enum("googleSheets", "notion", "airtable", "slack", name="IntegrationType"))
    config: Mapped[dict[str, Any]] = mapped_column(JSONB)
    """[IntegrationConfig]"""
    workspace_id: Mapped[str] = mapped_column("workspaceId", ForeignKey("Workspace.id", ondelete="CASCADE", onupdate="CASCADE", name="Integration_workspaceId_fkey"))

    __table_args__ = (
        Index("Integration_workspaceId_idx", "workspaceId"),
        Index("Integration_type_workspaceId_key", "type", "workspaceId", unique=True),
    )

    workspace: Mapped["Workspace"] = relationship(back_populates="integrations")

class DataMigration(Base):
    """Tracks database schema migrations.
    Used to manage and track the state of data structure changes.

    @property id - Unique identifier for the migration
    @property name - Name of the migration
    @property status - Current state of the migration (pending, applied, failed)
    @property startedAt - When the migration began
    @property finishedAt - When the migration completed
    """

    __tablename__ = "DataMigration"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    started_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    finished_at: Mapped[Optional[datetime]]
    name: Mapped[str]
    status: Mapped[str] = mapped_column(Enum("pending", "applied", "failed", name="DataMigrationStatus"))

    __table_args__ = (
        Index("DataMigration_name_key", "name", unique=True),
    )

class Workspace(Base):
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

    __tablename__ = "Workspace"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    name: Mapped[str]
    legacy_environment_id: Mapped[Optional[str]] = mapped_column("legacyEnvironmentId")
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="Workspace_organizationId_fkey"))
    styling: Mapped[dict[str, Any]] = mapped_column(JSONB, default=lambda: {"allowStyleOverwrite": True}, server_default=text("'{\"allowStyleOverwrite\":true}'"))
    """[Styling]"""
    config: Mapped[dict[str, Any]] = mapped_column(JSONB, default=lambda: {}, server_default=text("'{}'"))
    """[WorkspaceConfig]"""
    recontact_days: Mapped[int] = mapped_column("recontactDays", default=7, server_default=text("7"))
    link_survey_branding: Mapped[bool] = mapped_column("linkSurveyBranding", default=True, server_default=text("true"))
    in_app_survey_branding: Mapped[bool] = mapped_column("inAppSurveyBranding", default=True, server_default=text("true"))
    placement: Mapped[str] = mapped_column(Enum("bottomLeft", "bottomRight", "topLeft", "topRight", "center", name="WidgetPlacement"), default="bottomRight", server_default=text("'bottomRight'"))
    click_outside_close: Mapped[bool] = mapped_column("clickOutsideClose", default=True, server_default=text("true"))
    overlay: Mapped[str] = mapped_column(Enum("none", "light", "dark", name="SurveyOverlay"), default="none", server_default=text("'none'"))
    overlay_color: Mapped[Optional[str]] = mapped_column("overlayColor")
    """Custom overlay colour (hex, no alpha) and opacity (10-100%). null = the preset for `overlay`."""
    overlay_opacity: Mapped[Optional[int]] = mapped_column("overlayOpacity")
    logo: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))
    """[Logo]"""
    app_setup_completed: Mapped[bool] = mapped_column("appSetupCompleted", default=False, server_default=text("false"))
    custom_head_scripts: Mapped[Optional[str]] = mapped_column("customHeadScripts")
    custom_css: Mapped[Optional[dict[str, Any]]] = mapped_column("customCss", JSONB(none_as_null=True))
    """[CustomCssStored] Workspace custom CSS (ENG-2949), written only by the shared save service."""
    custom_css_previous: Mapped[Optional[dict[str, Any]]] = mapped_column("customCssPrevious", JSONB(none_as_null=True))
    """[CustomCssStored] The one recoverable previous workspace CSS revision, source included. Restored by
    submitting its source again through the current processor and authorization checks.
    """

    __table_args__ = (
        Index("Workspace_legacyEnvironmentId_key", "legacyEnvironmentId", unique=True),
        Index("Workspace_organizationId_name_key", "organizationId", "name", unique=True),
    )

    organization: Mapped["Organization"] = relationship(back_populates="workspaces")
    languages: Mapped[list["Language"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="workspace")
    embedded_data: Mapped[list["EmbeddedData"]] = relationship(cascade="all, delete", passive_deletes=True, overlaps="embedded_data,survey", back_populates="workspace")
    workspace_teams: Mapped[list["WorkspaceTeam"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="workspace")
    feedback_directory_workspaces: Mapped[list["FeedbackDirectoryWorkspace"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="workspace")
    charts: Mapped[list["Chart"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="workspace")
    dashboards: Mapped[list["Dashboard"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="workspace")
    creatives: Mapped[list["Creative"]] = relationship(cascade="all, delete", passive_deletes=True, overlaps="creatives,survey", back_populates="workspace")
    attachment_exports: Mapped[list["AttachmentExport"]] = relationship(cascade="all, delete", passive_deletes=True, overlaps="attachment_exports,survey", back_populates="workspace")
    surveys: Mapped[list["Survey"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="workspace")
    contacts: Mapped[list["Contact"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="workspace")
    action_classes: Mapped[list["ActionClass"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="workspace")
    contact_attribute_keys: Mapped[list["ContactAttributeKey"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="workspace")
    webhooks: Mapped[list["Webhook"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="workspace")
    tags: Mapped[list["Tag"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="workspace")
    segments: Mapped[list["Segment"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="workspace")
    integrations: Mapped[list["Integration"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="workspace")
    api_key_workspaces: Mapped[list["ApiKeyWorkspace"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="workspace")
    feedback_sources: Mapped[list["FeedbackSource"]] = relationship(cascade="all, delete", passive_deletes=True, overlaps="feedback_directory_workspace,feedback_sources", back_populates="workspace")
    workflows: Mapped[list["Workflow"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="workspace")
    workflow_runs: Mapped[list["WorkflowRun"]] = relationship(cascade="all, delete", passive_deletes=True, overlaps="runs,workflow", back_populates="workspace")

class Organization(Base):
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

    __tablename__ = "Organization"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    name: Mapped[str]
    whitelabel: Mapped[dict[str, Any]] = mapped_column(JSONB, default=lambda: {}, server_default=text("'{}'"))
    """[OrganizationWhitelabel]"""
    is_aismart_tools_enabled: Mapped[bool] = mapped_column("isAISmartToolsEnabled", default=False, server_default=text("false"))
    display_time_zone: Mapped[Optional[str]] = mapped_column("displayTimeZone")

    memberships: Mapped[list["Membership"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    workspaces: Mapped[list["Workspace"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    invites: Mapped[list["Invite"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    teams: Mapped[list["Team"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    api_keys: Mapped[list["ApiKey"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    feedback_directories: Mapped[list["FeedbackDirectory"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    retention_policies: Mapped[list["RetentionPolicy"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    retention_exemptions: Mapped[list["RetentionExemption"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    retention_runs: Mapped[list["RetentionRun"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    retention_notices: Mapped[list["RetentionNotice"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    billing: Mapped[Optional["OrganizationBilling"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")

class OrganizationBilling(Base):
    """Stores billing and Stripe synchronization data for an organization.

    @property organizationId - Reference to the organization
    @property stripeCustomerId - Stripe customer ID for cloud billing
    @property limits - Plan limits used by feature and quota checks
    @property usageCycleAnchor - Stable anchor used to derive the current monthly usage cycle
    @property stripe - Stripe synchronization snapshot data
    """

    __tablename__ = "OrganizationBilling"

    organization_id: Mapped[str] = mapped_column(ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="OrganizationBilling_organization_id_fkey"), primary_key=True)
    stripe_customer_id: Mapped[Optional[str]]
    limits: Mapped[dict[str, Any]] = mapped_column(JSONB)
    """[OrganizationBillingPlanLimits]"""
    usage_cycle_anchor: Mapped[Optional[datetime]]
    stripe: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))
    """[OrganizationStripeBilling]"""
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("OrganizationBilling_stripe_customer_id_key", "stripe_customer_id", unique=True),
    )

    organization: Mapped["Organization"] = relationship(back_populates="billing")

class Membership(Base):
    """Links users to organizations with specific roles.
    Manages organization membership and permissions.
    Core model for managing user access within organizations.

    @property organization - The organization the user belongs to
    @property user - The member user
    @property accepted - Whether the user has accepted the membership
    @property role - User's role within the organization (owner, manager, member, billing)
    """

    __tablename__ = "Membership"

    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="Membership_organizationId_fkey"), primary_key=True)
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="Membership_userId_fkey"), primary_key=True)
    accepted: Mapped[bool] = mapped_column(default=False, server_default=text("false"))
    role: Mapped[str] = mapped_column(Enum("owner", "manager", "member", "billing", "dashboard_viewer", name="OrganizationRole"), default="member", server_default=text("'member'"))

    __table_args__ = (
        Index("Membership_organizationId_idx", "organizationId"),
    )

    organization: Mapped["Organization"] = relationship(back_populates="memberships")
    user: Mapped["User"] = relationship(back_populates="memberships")

class Invite(Base):
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

    __tablename__ = "Invite"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: str(uuid_mod.uuid4()))
    email: Mapped[str]
    name: Mapped[Optional[str]]
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="Invite_organizationId_fkey"))
    creator_id: Mapped[str] = mapped_column("creatorId", ForeignKey("User.id", ondelete="RESTRICT", onupdate="CASCADE", name="Invite_creatorId_fkey"))
    acceptor_id: Mapped[Optional[str]] = mapped_column("acceptorId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="Invite_acceptorId_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    expires_at: Mapped[datetime] = mapped_column("expiresAt")
    role: Mapped[str] = mapped_column(Enum("owner", "manager", "member", "billing", "dashboard_viewer", name="OrganizationRole"), default="member", server_default=text("'member'"))
    team_ids: Mapped[list[str]] = mapped_column("teamIds", ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))

    __table_args__ = (
        Index("Invite_email_organizationId_idx", "email", "organizationId"),
        Index("Invite_organizationId_idx", "organizationId"),
    )

    organization: Mapped["Organization"] = relationship(back_populates="invites")
    creator: Mapped["User"] = relationship(foreign_keys=[creator_id], back_populates="invites_created")
    acceptor: Mapped[Optional["User"]] = relationship(foreign_keys=[acceptor_id], back_populates="invites_accepted")

class ApiKey(Base):
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

    __tablename__ = "ApiKey"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    created_by: Mapped[Optional[str]] = mapped_column("createdBy")
    last_used_at: Mapped[Optional[datetime]] = mapped_column("lastUsedAt")
    label: Mapped[str]
    hashed_key: Mapped[str] = mapped_column("hashedKey")
    lookup_hash: Mapped[Optional[str]] = mapped_column("lookupHash")
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="ApiKey_organizationId_fkey"))
    organization_access: Mapped[dict[str, Any]] = mapped_column("organizationAccess", JSONB, default=lambda: {}, server_default=text("'{}'"))
    """[OrganizationAccess]"""

    __table_args__ = (
        Index("ApiKey_lookupHash_key", "lookupHash", unique=True),
        Index("ApiKey_organizationId_idx", "organizationId"),
    )

    organization: Mapped["Organization"] = relationship(back_populates="api_keys")
    api_key_workspaces: Mapped[list["ApiKeyWorkspace"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="api_key")

class ApiKeyWorkspace(Base):
    """Links API keys to workspaces with specific permissions.
    Enables granular access control for API keys across workspaces.

    @property id - Unique identifier for the workspace access entry
    @property apiKey - The associated API key
    @property workspace - The workspace being accessed
    @property permission - Level of access granted
    """

    __tablename__ = "ApiKeyWorkspace"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)
    api_key_id: Mapped[str] = mapped_column("apiKeyId", ForeignKey("ApiKey.id", ondelete="CASCADE", onupdate="CASCADE", name="ApiKeyWorkspace_apiKeyId_fkey"))
    workspace_id: Mapped[str] = mapped_column("workspaceId", ForeignKey("Workspace.id", ondelete="CASCADE", onupdate="CASCADE", name="ApiKeyWorkspace_workspaceId_fkey"))
    permission: Mapped[str] = mapped_column(Enum("read", "write", "manage", name="ApiKeyPermission"))

    __table_args__ = (
        Index("ApiKeyWorkspace_workspaceId_idx", "workspaceId"),
        Index("ApiKeyWorkspace_apiKeyId_workspaceId_key", "apiKeyId", "workspaceId", unique=True),
    )

    api_key: Mapped["ApiKey"] = relationship(back_populates="api_key_workspaces")
    workspace: Mapped["Workspace"] = relationship(back_populates="api_key_workspaces")

class Account(Base):
    """Stores third-party authentication account information.
    Enables OAuth and other external authentication methods.

    @property id - Unique identifier for the account
    @property user - The Formbricks user who owns this account
    @property provider - The authentication provider (GitHub, Google, etc.)
    @property providerAccountId - User ID from the provider
    @property access_token - OAuth access token
    @property refresh_token - OAuth refresh token
    """

    __tablename__ = "Account"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="Account_userId_fkey"))
    type: Mapped[Optional[str]]
    provider: Mapped[str]
    provider_account_id: Mapped[str] = mapped_column("providerAccountId")
    access_token: Mapped[Optional[str]]
    refresh_token: Mapped[Optional[str]]
    expires_at: Mapped[Optional[int]]
    ext_expires_in: Mapped[Optional[int]]
    token_type: Mapped[Optional[str]]
    scope: Mapped[Optional[str]]
    id_token: Mapped[Optional[str]]
    session_state: Mapped[Optional[str]]
    password: Mapped[Optional[str]]
    access_token_expires_at: Mapped[Optional[datetime]] = mapped_column("accessTokenExpiresAt")
    refresh_token_expires_at: Mapped[Optional[datetime]] = mapped_column("refreshTokenExpiresAt")
    issuer: Mapped[Optional[str]]
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

    __table_args__ = (
        Index("Account_userId_idx", "userId"),
        Index("Account_provider_providerAccountId_key", "provider", "providerAccountId", unique=True),
        Index("Account_issuer_providerAccountId_key", "issuer", "providerAccountId", unique=True),
    )

    user: Mapped[Optional["User"]] = relationship(back_populates="accounts")

class Session(Base):
    """Stores active authentication sessions for revocable server-side login state.

    @property sessionToken - Opaque token stored in the browser cookie
    @property user - The Formbricks user who owns this session
    @property expires - Hard expiry for the session
    """

    __tablename__ = "Session"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    session_token: Mapped[str] = mapped_column("sessionToken")
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="Session_userId_fkey"))
    expires: Mapped[datetime]
    ip_address: Mapped[Optional[str]] = mapped_column("ipAddress")
    user_agent: Mapped[Optional[str]] = mapped_column("userAgent")

    __table_args__ = (
        Index("Session_sessionToken_key", "sessionToken", unique=True),
        Index("Session_userId_idx", "userId"),
    )

    user: Mapped["User"] = relationship(back_populates="sessions")
    oauth_access_tokens: Mapped[list["OauthAccessToken"]] = relationship(back_populates="session")
    oauth_refresh_tokens: Mapped[list["OauthRefreshToken"]] = relationship(back_populates="session")

class VerificationToken(Base):
    """Stores one-time verification tokens used by Auth.js adapter flows."""

    __tablename__ = "VerificationToken"

    identifier: Mapped[str] = mapped_column()
    token: Mapped[str] = mapped_column()
    expires: Mapped[datetime]

    __table_args__ = (
        Index("VerificationToken_identifier_token_key", "identifier", "token", unique=True),
    )

    __mapper_args__ = {"primary_key": [identifier, token]}

class TwoFactor(Base):
    """Better Auth (ENG-1054) two-factor secrets. Supersedes the legacy User.twoFactorSecret /
    User.backupCodes columns (kept until cutover). `secret` and `backupCodes` are encrypted
    at rest by Better Auth using BETTER_AUTH_SECRET.
    """

    __tablename__ = "TwoFactor"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    secret: Mapped[str]
    backup_codes: Mapped[str] = mapped_column("backupCodes")
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="TwoFactor_userId_fkey"))
    verified: Mapped[bool] = mapped_column(default=True, server_default=text("true"))
    failed_verification_count: Mapped[int] = mapped_column("failedVerificationCount", default=0, server_default=text("0"))
    """Better Auth 1.6.23+ two-factor brute-force lockout: consecutive failed verification attempts."""
    locked_until: Mapped[Optional[datetime]] = mapped_column("lockedUntil")
    """Better Auth 1.6.23+ two-factor brute-force lockout: set once failedVerificationCount trips the limit."""

    __table_args__ = (
        Index("TwoFactor_userId_key", "userId", unique=True),
    )

    user: Mapped["User"] = relationship(back_populates="two_factor")

class Jwks(Base):
    """Better Auth JWT signing keys for OAuth/OIDC access and ID tokens."""

    __tablename__ = "jwks"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    public_key: Mapped[str] = mapped_column("publicKey")
    private_key: Mapped[str] = mapped_column("privateKey")
    created_at: Mapped[datetime] = mapped_column("createdAt")
    expires_at: Mapped[Optional[datetime]] = mapped_column("expiresAt")
    alg: Mapped[Optional[str]]
    """Better Auth 1.7 (ENG-2343): the JWS algorithm this key signs with, so a keyring can hold more
    than one alg and `getLatestKeyByAlg` can pick the right key per resource.

    NOT optional upstream in practice, despite being declared `required: false`: `createJwk` always
    puts `alg` in the row it writes (defaulting to EdDSA), so without this column every key mint is
    rejected by Prisma as an unknown argument — which breaks JWT signing, and with it the whole MCP
    OAuth flow, on any deployment that has yet to mint a key. Nullable because rows written before
    1.7 have no value; upstream reads treat `alg: null` as the configured default alg, so existing
    keys keep working untouched.
    """
    crv: Mapped[Optional[str]]
    """Better Auth 1.7 (ENG-2343): the curve for EC/OKP keys. Written whenever it can be derived from
    the generated key, which it can for the default EdDSA/Ed25519 config we run, so it is on the
    same always-written path as `alg` above.
    """

class OauthClient(Base):
    """Better Auth OAuth 2.1 client registrations for MCP and future OAuth clients."""

    __tablename__ = "oauthClient"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    client_id: Mapped[str] = mapped_column("clientId")
    client_secret: Mapped[Optional[str]] = mapped_column("clientSecret")
    disabled: Mapped[Optional[bool]] = mapped_column(default=False, server_default=text("false"))
    skip_consent: Mapped[Optional[bool]] = mapped_column("skipConsent")
    enable_end_session: Mapped[Optional[bool]] = mapped_column("enableEndSession")
    subject_type: Mapped[Optional[str]] = mapped_column("subjectType")
    scopes: Mapped[list[str]] = mapped_column(ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))
    user_id: Mapped[Optional[str]] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="oauthClient_userId_fkey"))
    created_at: Mapped[Optional[datetime]] = mapped_column("createdAt")
    updated_at: Mapped[Optional[datetime]] = mapped_column("updatedAt")
    name: Mapped[Optional[str]]
    uri: Mapped[Optional[str]]
    icon: Mapped[Optional[str]]
    contacts: Mapped[list[str]] = mapped_column(ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))
    tos: Mapped[Optional[str]]
    policy: Mapped[Optional[str]]
    software_id: Mapped[Optional[str]] = mapped_column("softwareId")
    software_version: Mapped[Optional[str]] = mapped_column("softwareVersion")
    software_statement: Mapped[Optional[str]] = mapped_column("softwareStatement")
    redirect_uris: Mapped[list[str]] = mapped_column("redirectUris", ARRAY(Text), nullable=True)
    post_logout_redirect_uris: Mapped[list[str]] = mapped_column("postLogoutRedirectUris", ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))
    token_endpoint_auth_method: Mapped[Optional[str]] = mapped_column("tokenEndpointAuthMethod")
    grant_types: Mapped[list[str]] = mapped_column("grantTypes", ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))
    response_types: Mapped[list[str]] = mapped_column("responseTypes", ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))
    public: Mapped[Optional[bool]]
    """Unwritten by Better Auth 1.7 (it dropped both from its schema), kept deliberately: they are
    nullable, so keeping them makes the 1.7 migration purely additive and a rollback to 1.6 needs no
    down-migration. Drop them in a separate cleanup after a bake period.
    """
    type: Mapped[Optional[str]]
    require_pkce: Mapped[Optional[bool]] = mapped_column("requirePKCE")
    reference_id: Mapped[Optional[str]] = mapped_column("referenceId")
    metadata_: Mapped[Optional[dict[str, Any]]] = mapped_column("metadata", JSONB(none_as_null=True))
    client_discovery_id: Mapped[Optional[str]] = mapped_column("clientDiscoveryId")
    """Better Auth 1.7 (ENG-2343)."""
    client_credentials_scopes: Mapped[list[str]] = mapped_column("clientCredentialsScopes", ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))
    backchannel_logout_uri: Mapped[Optional[str]] = mapped_column("backchannelLogoutUri")
    backchannel_logout_session_required: Mapped[Optional[bool]] = mapped_column("backchannelLogoutSessionRequired")
    application_type: Mapped[Optional[str]] = mapped_column("applicationType")
    jwks: Mapped[Optional[str]]
    jwks_uri: Mapped[Optional[str]] = mapped_column("jwksUri")
    dpop_bound_access_tokens: Mapped[Optional[bool]] = mapped_column("dpopBoundAccessTokens", default=False, server_default=text("false"))

    __table_args__ = (
        Index("oauthClient_clientId_key", "clientId", unique=True),
        Index("oauthClient_userId_idx", "userId"),
    )

    user: Mapped[Optional["User"]] = relationship(back_populates="oauth_clients")
    access_tokens: Mapped[list["OauthAccessToken"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="client")
    refresh_tokens: Mapped[list["OauthRefreshToken"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="client")
    consents: Mapped[list["OauthConsent"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="client")
    resources: Mapped[list["OauthClientResource"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="client")

class OauthAccessToken(Base):
    """Better Auth opaque access token records. The token column stores Better Auth's hashed lookup value.
    JWT access tokens are verified through JWKS and bounded by the configured short lifetime.
    """

    __tablename__ = "oauthAccessToken"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    token: Mapped[str]
    client_id: Mapped[str] = mapped_column("clientId", ForeignKey("oauthClient.clientId", ondelete="CASCADE", onupdate="CASCADE", name="oauthAccessToken_clientId_fkey"))
    session_id: Mapped[Optional[str]] = mapped_column("sessionId", ForeignKey("Session.id", ondelete="SET NULL", onupdate="CASCADE", name="oauthAccessToken_sessionId_fkey"))
    user_id: Mapped[Optional[str]] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="oauthAccessToken_userId_fkey"))
    reference_id: Mapped[Optional[str]] = mapped_column("referenceId")
    refresh_id: Mapped[Optional[str]] = mapped_column("refreshId", ForeignKey("oauthRefreshToken.id", ondelete="CASCADE", onupdate="CASCADE", name="oauthAccessToken_refreshId_fkey"))
    expires_at: Mapped[datetime] = mapped_column("expiresAt")
    created_at: Mapped[datetime] = mapped_column("createdAt")
    scopes: Mapped[list[str]] = mapped_column(ARRAY(Text), nullable=True)
    authorization_code_id: Mapped[Optional[str]] = mapped_column("authorizationCodeId")
    """Better Auth 1.7 (ENG-2343). `resources` is the RFC 8707 resource set the grant approved — the
    binding this whole upgrade exists to get (GHSA-p2fr-6hmx-4528).
    """
    resources: Mapped[list[str]] = mapped_column(ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))
    requested_user_info_claims: Mapped[list[str]] = mapped_column("requestedUserInfoClaims", ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))
    revoked: Mapped[Optional[datetime]]
    confirmation: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))

    __table_args__ = (
        Index("oauthAccessToken_token_key", "token", unique=True),
        Index("oauthAccessToken_clientId_idx", "clientId"),
        Index("oauthAccessToken_sessionId_idx", "sessionId"),
        Index("oauthAccessToken_userId_idx", "userId"),
        Index("oauthAccessToken_refreshId_idx", "refreshId"),
        Index("oauthAccessToken_authorizationCodeId_idx", "authorizationCodeId"),
    )

    client: Mapped["OauthClient"] = relationship(back_populates="access_tokens")
    session: Mapped[Optional["Session"]] = relationship(back_populates="oauth_access_tokens")
    user: Mapped[Optional["User"]] = relationship(back_populates="oauth_access_tokens")
    refresh: Mapped[Optional["OauthRefreshToken"]] = relationship(back_populates="access_tokens")

class OauthRefreshToken(Base):
    """Better Auth OAuth refresh tokens used by delegated MCP clients. The token column stores a hashed lookup value."""

    __tablename__ = "oauthRefreshToken"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    token: Mapped[str]
    client_id: Mapped[str] = mapped_column("clientId", ForeignKey("oauthClient.clientId", ondelete="CASCADE", onupdate="CASCADE", name="oauthRefreshToken_clientId_fkey"))
    session_id: Mapped[Optional[str]] = mapped_column("sessionId", ForeignKey("Session.id", ondelete="SET NULL", onupdate="CASCADE", name="oauthRefreshToken_sessionId_fkey"))
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="oauthRefreshToken_userId_fkey"))
    reference_id: Mapped[Optional[str]] = mapped_column("referenceId")
    expires_at: Mapped[datetime] = mapped_column("expiresAt")
    created_at: Mapped[datetime] = mapped_column("createdAt")
    revoked: Mapped[Optional[datetime]]
    auth_time: Mapped[Optional[datetime]] = mapped_column("authTime")
    scopes: Mapped[list[str]] = mapped_column(ARRAY(Text), nullable=True)
    authorization_code_id: Mapped[Optional[str]] = mapped_column("authorizationCodeId")
    """Better Auth 1.7 (ENG-2343). `resources` carries the full original grant so a refresh can narrow
    the resource set but never widen it.
    """
    resources: Mapped[list[str]] = mapped_column(ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))
    requested_user_info_claims: Mapped[list[str]] = mapped_column("requestedUserInfoClaims", ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))
    rotated_at: Mapped[Optional[datetime]] = mapped_column("rotatedAt")
    rotation_replay_response: Mapped[Optional[str]] = mapped_column("rotationReplayResponse")
    rotation_replay_expires_at: Mapped[Optional[datetime]] = mapped_column("rotationReplayExpiresAt")
    confirmation: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))

    __table_args__ = (
        Index("oauthRefreshToken_token_key", "token", unique=True),
        Index("oauthRefreshToken_clientId_idx", "clientId"),
        Index("oauthRefreshToken_sessionId_idx", "sessionId"),
        Index("oauthRefreshToken_userId_idx", "userId"),
        Index("oauthRefreshToken_authorizationCodeId_idx", "authorizationCodeId"),
    )

    client: Mapped["OauthClient"] = relationship(back_populates="refresh_tokens")
    session: Mapped[Optional["Session"]] = relationship(back_populates="oauth_refresh_tokens")
    user: Mapped["User"] = relationship(back_populates="oauth_refresh_tokens")
    access_tokens: Mapped[list["OauthAccessToken"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="refresh")

class OauthConsent(Base):
    """Better Auth OAuth user consent grants for delegated clients."""

    __tablename__ = "oauthConsent"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    client_id: Mapped[str] = mapped_column("clientId", ForeignKey("oauthClient.clientId", ondelete="CASCADE", onupdate="CASCADE", name="oauthConsent_clientId_fkey"))
    user_id: Mapped[Optional[str]] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="oauthConsent_userId_fkey"))
    reference_id: Mapped[Optional[str]] = mapped_column("referenceId")
    scopes: Mapped[list[str]] = mapped_column(ARRAY(Text), nullable=True)
    created_at: Mapped[datetime] = mapped_column("createdAt")
    updated_at: Mapped[datetime] = mapped_column("updatedAt")
    resources: Mapped[list[str]] = mapped_column(ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))
    """Better Auth 1.7 (ENG-2343). /authorize compares the requested resources against these; without
    them every pre-existing connection is forced back through the consent screen.
    """
    requested_user_info_claims: Mapped[list[str]] = mapped_column("requestedUserInfoClaims", ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))

    __table_args__ = (
        Index("oauthConsent_clientId_idx", "clientId"),
        Index("oauthConsent_userId_idx", "userId"),
    )

    client: Mapped["OauthClient"] = relationship(back_populates="consents")
    user: Mapped[Optional["User"]] = relationship(back_populates="oauth_consents")

class OauthResource(Base):
    """Better Auth 1.7 (ENG-2343): a protected resource the authorization server issues tokens for, with
    its own token lifetimes, scope allow-list and claims. Replaces the flat `validAudiences` option —
    a token is now bound to the resource it was approved for instead of to a static allow-list.
    `identifier` is the RFC 8707 `resource` value and is the natural key other rows point at.
    """

    __tablename__ = "oauthResource"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    identifier: Mapped[str]
    name: Mapped[str]
    access_token_ttl: Mapped[Optional[int]] = mapped_column("accessTokenTtl")
    refresh_token_ttl: Mapped[Optional[int]] = mapped_column("refreshTokenTtl")
    signing_algorithm: Mapped[Optional[str]] = mapped_column("signingAlgorithm")
    signing_key_id: Mapped[Optional[str]] = mapped_column("signingKeyId")
    allowed_scopes: Mapped[list[str]] = mapped_column("allowedScopes", ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))
    custom_claims: Mapped[Optional[dict[str, Any]]] = mapped_column("customClaims", JSONB(none_as_null=True))
    dpop_bound_access_tokens_required: Mapped[Optional[bool]] = mapped_column("dpopBoundAccessTokensRequired", default=False, server_default=text("false"))
    disabled: Mapped[Optional[bool]] = mapped_column(default=False, server_default=text("false"))
    created_at: Mapped[Optional[datetime]] = mapped_column("createdAt")
    updated_at: Mapped[Optional[datetime]] = mapped_column("updatedAt")
    policy_version: Mapped[Optional[int]] = mapped_column("policyVersion", default=1, server_default=text("1"))
    metadata_: Mapped[Optional[dict[str, Any]]] = mapped_column("metadata", JSONB(none_as_null=True))

    __table_args__ = (
        Index("oauthResource_identifier_key", "identifier", unique=True),
    )

    clients: Mapped[list["OauthClientResource"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="resource")

class OauthClientResource(Base):
    """Better Auth 1.7 (ENG-2343): links a client to the resources it may request tokens for. With
    `enforcePerClientResources` (the default) a missing row means `invalid_target` at the token
    endpoint, which is why the migration backfills one row per existing client.
    Both sides reference natural keys, matching the other oauth child tables.
    """

    __tablename__ = "oauthClientResource"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    client_id: Mapped[str] = mapped_column("clientId", ForeignKey("oauthClient.clientId", ondelete="CASCADE", onupdate="CASCADE", name="oauthClientResource_clientId_fkey"))
    resource_id: Mapped[str] = mapped_column("resourceId", ForeignKey("oauthResource.identifier", ondelete="CASCADE", onupdate="CASCADE", name="oauthClientResource_resourceId_fkey"))
    metadata_: Mapped[Optional[dict[str, Any]]] = mapped_column("metadata", JSONB(none_as_null=True))
    created_at: Mapped[Optional[datetime]] = mapped_column("createdAt")

    __table_args__ = (
        Index("oauthClientResource_resourceId_idx", "resourceId"),
        Index("oauthClientResource_clientId_resourceId_key", "clientId", "resourceId", unique=True),
    )

    client: Mapped["OauthClient"] = relationship(back_populates="resources")
    resource: Mapped["OauthResource"] = relationship(back_populates="clients")

class OauthClientAssertion(Base):
    """Better Auth 1.7 (ENG-2343): replay store for `private_key_jwt` client assertions (the row id is the
    assertion jti). We use no private_key_jwt client authentication today, but the plugin declares the
    table, so omitting it risks an adapter error on a code path we do not control.
    """

    __tablename__ = "oauthClientAssertion"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    expires_at: Mapped[datetime] = mapped_column("expiresAt")

    __table_args__ = (
        Index("oauthClientAssertion_expiresAt_idx", "expiresAt"),
    )

class PasswordResetToken(Base):
    """Stores the active password reset token for a user.
    Tokens are opaque, hashed at rest, revocable, and single-use.
    """

    __tablename__ = "PasswordResetToken"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    token_hash: Mapped[str]
    expires_at: Mapped[datetime]
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="PasswordResetToken_userId_fkey"))

    __table_args__ = (
        Index("PasswordResetToken_token_hash_key", "token_hash", unique=True),
        Index("PasswordResetToken_userId_key", "userId", unique=True),
        Index("PasswordResetToken_expires_at_idx", "expires_at"),
    )

    user: Mapped["User"] = relationship(back_populates="password_reset_token")

class User(Base):
    """Represents a user in the Formbricks system.
    Central model for user authentication and profile management.

    @property id - Unique identifier for the user
    @property name - Display name of the user
    @property email - User's email address
    @property twoFactorEnabled - Whether 2FA is active
    @property memberships - Organizations the user belongs to
    @property notificationSettings - User's notification preferences
    """

    __tablename__ = "User"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    name: Mapped[str]
    email: Mapped[str]
    email_verified: Mapped[bool] = mapped_column(default=False, server_default=text("false"))
    two_factor_secret: Mapped[Optional[str]] = mapped_column("twoFactorSecret")
    two_factor_enabled: Mapped[bool] = mapped_column("twoFactorEnabled", default=False, server_default=text("false"))
    backup_codes: Mapped[Optional[str]] = mapped_column("backupCodes")
    password: Mapped[Optional[str]]
    identity_provider: Mapped[str] = mapped_column("identityProvider", Enum("email", "github", "google", "azuread", "openid", "saml", name="IdentityProvider"), default="email", server_default=text("'email'"))
    identity_provider_account_id: Mapped[Optional[str]] = mapped_column("identityProviderAccountId")
    group_id: Mapped[Optional[str]] = mapped_column("groupId")
    notification_settings: Mapped[dict[str, Any]] = mapped_column("notificationSettings", JSONB, default=lambda: {}, server_default=text("'{}'"))
    """[UserNotificationSettings]"""
    locale: Mapped[str] = mapped_column(default="en-US", server_default=text("'en-US'"))
    """[Locale]"""
    last_login_at: Mapped[Optional[datetime]] = mapped_column("lastLoginAt")
    is_active: Mapped[bool] = mapped_column("isActive", default=True, server_default=text("true"))
    reactivated_at: Mapped[Optional[datetime]] = mapped_column("reactivatedAt")
    """When an admin last reactivated this user. Restarts the members retention clock (ENG-3713)
    without touching `lastLoginAt`, which the member list, usage counts and APIs report as a real
    sign-in.
    """
    last_active_at: Mapped[Optional[datetime]] = mapped_column("lastActiveAt")
    """When the user last completed a sign-in (`lastActivePlugin`, after any second factor) or renewed a
    session (Better Auth `databaseHooks.session.update`), on the database clock and never moved
    backwards. The members retention clock reads it: sessions are
    deleted on sign-out and expiry, so their own timestamps can't serve as a record of activity.
    """
    is_bootstrap_admin: Mapped[Optional[str]] = mapped_column("isBootstrapAdmin", Enum("bootstrapAdmin", name="BootstrapAdminMarker"))

    __table_args__ = (
        Index("User_email_key", "email", unique=True),
        Index("User_isBootstrapAdmin_key", "isBootstrapAdmin", unique=True),
    )

    memberships: Mapped[list["Membership"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")
    accounts: Mapped[list["Account"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")
    sessions: Mapped[list["Session"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")
    oauth_clients: Mapped[list["OauthClient"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")
    oauth_access_tokens: Mapped[list["OauthAccessToken"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")
    oauth_refresh_tokens: Mapped[list["OauthRefreshToken"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")
    oauth_consents: Mapped[list["OauthConsent"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")
    invites_created: Mapped[list["Invite"]] = relationship(foreign_keys="Invite.creator_id", passive_deletes="all", back_populates="creator")
    invites_accepted: Mapped[list["Invite"]] = relationship(foreign_keys="Invite.acceptor_id", cascade="all, delete", passive_deletes=True, back_populates="acceptor")
    surveys: Mapped[list["Survey"]] = relationship(foreign_keys="Survey.created_by", back_populates="creator")
    owned_surveys: Mapped[list["Survey"]] = relationship(foreign_keys="Survey.owner_id", back_populates="owner")
    charts: Mapped[list["Chart"]] = relationship(back_populates="creator")
    dashboards: Mapped[list["Dashboard"]] = relationship(back_populates="creator")
    creatives: Mapped[list["Creative"]] = relationship(back_populates="creator")
    feedback_sources: Mapped[list["FeedbackSource"]] = relationship(back_populates="creator")
    team_users: Mapped[list["TeamUser"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")
    workflows: Mapped[list["Workflow"]] = relationship(back_populates="creator")
    workflow_versions_published: Mapped[list["WorkflowVersion"]] = relationship(back_populates="publisher")
    retention_policies_updated: Mapped[list["RetentionPolicy"]] = relationship(back_populates="updated_by")
    exemptions_created: Mapped[list["RetentionExemption"]] = relationship(foreign_keys="RetentionExemption.created_by_id", back_populates="created_by")
    exemptions_revoked: Mapped[list["RetentionExemption"]] = relationship(foreign_keys="RetentionExemption.revoked_by_id", back_populates="revoked_by")
    retention_notices: Mapped[list["RetentionNotice"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")
    attachment_exports: Mapped[list["AttachmentExport"]] = relationship(back_populates="requested_by")
    two_factor: Mapped[Optional["TwoFactor"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")
    password_reset_token: Mapped[Optional["PasswordResetToken"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")

class Segment(Base):
    """Defines a segment of contacts based on attributes.
    Used for targeting surveys to specific user groups.

    @property id - Unique identifier for the segment
    @property title - Display name of the segment
    @property filters - Rules defining the segment
    @property isPrivate - Whether the segment is private
    @property workspace - The workspace this segment belongs to
    """

    __tablename__ = "Segment"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    title: Mapped[str]
    description: Mapped[Optional[str]]
    is_private: Mapped[bool] = mapped_column("isPrivate", default=True, server_default=text("true"))
    filters: Mapped[dict[str, Any]] = mapped_column(JSONB, default=lambda: [], server_default=text("'[]'"))
    """[SegmentFilter]"""
    workspace_id: Mapped[str] = mapped_column("workspaceId", ForeignKey("Workspace.id", ondelete="CASCADE", onupdate="CASCADE", name="Segment_workspaceId_fkey"))

    __table_args__ = (
        Index("Segment_workspaceId_idx", "workspaceId"),
        Index("Segment_workspaceId_title_key", "workspaceId", "title", unique=True),
    )

    workspace: Mapped["Workspace"] = relationship(back_populates="segments")
    surveys: Mapped[list["Survey"]] = relationship(back_populates="segment")

class Language(Base):
    """Represents a supported language in the system.
    Used for multilingual survey support.

    @property id - Unique identifier for the language
    @property code - Language code (e.g., 'en-US')
    @property alias - Optional friendly name
    @property workspace - The workspace this language is enabled for
    """

    __tablename__ = "Language"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    code: Mapped[str]
    alias: Mapped[Optional[str]]
    workspace_id: Mapped[str] = mapped_column("workspaceId", ForeignKey("Workspace.id", ondelete="CASCADE", onupdate="CASCADE", name="Language_workspaceId_fkey"))

    __table_args__ = (
        Index("Language_workspaceId_code_key", "workspaceId", "code", unique=True),
    )

    workspace: Mapped["Workspace"] = relationship(back_populates="languages")
    survey_languages: Mapped[list["SurveyLanguage"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="language")
    creative_languages: Mapped[list["CreativeLanguage"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="language")

class SurveyLanguage(Base):
    """Links surveys to their supported languages.
    Manages which languages are available for each survey.

    @property language - The supported language
    @property survey - The associated survey
    @property default - Whether this is the default language
    @property enabled - Whether this language is currently active
    """

    __tablename__ = "SurveyLanguage"

    language_id: Mapped[str] = mapped_column("languageId", ForeignKey("Language.id", ondelete="CASCADE", onupdate="CASCADE", name="SurveyLanguage_languageId_fkey"), primary_key=True)
    survey_id: Mapped[str] = mapped_column("surveyId", ForeignKey("Survey.id", ondelete="CASCADE", onupdate="CASCADE", name="SurveyLanguage_surveyId_fkey"), primary_key=True)
    default: Mapped[bool] = mapped_column(default=False, server_default=text("false"))
    enabled: Mapped[bool] = mapped_column(default=True, server_default=text("true"))

    __table_args__ = (
        Index("SurveyLanguage_surveyId_idx", "surveyId"),
    )

    language: Mapped["Language"] = relationship(back_populates="survey_languages")
    survey: Mapped["Survey"] = relationship(back_populates="languages")

class EmbeddedData(Base):
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

    __tablename__ = "EmbeddedData"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    key: Mapped[Optional[str]]
    name: Mapped[str]
    description: Mapped[Optional[str]]
    source: Mapped[str] = mapped_column(Enum("computed", "ingested", "reserved", name="EmbeddedDataSource"))
    data_type: Mapped[str] = mapped_column("dataType", Enum("string", "number", "boolean", "date", name="EmbeddedDataType"), default="string", server_default=text("'string'"))
    default_value: Mapped[Optional[dict[str, Any]]] = mapped_column("defaultValue", JSONB(none_as_null=True))
    """[EmbeddedDataDefaultValue]"""
    locked: Mapped[bool] = mapped_column(default=False, server_default=text("false"))
    survey_id: Mapped[Optional[str]] = mapped_column("surveyId")
    workspace_id: Mapped[str] = mapped_column("workspaceId", ForeignKey("Workspace.id", ondelete="CASCADE", onupdate="CASCADE", name="EmbeddedData_workspaceId_fkey"))

    __table_args__ = (
        Index("EmbeddedData_surveyId_idx", "surveyId"),
        Index("EmbeddedData_workspaceId_key_key", "workspaceId", "key", unique=True),
        Index("EmbeddedData_id_workspaceId_key", "id", "workspaceId", unique=True),
        ForeignKeyConstraint(["surveyId", "workspaceId"], ["Survey.id", "Survey.workspaceId"], ondelete="CASCADE", onupdate="CASCADE", name="EmbeddedData_surveyId_workspaceId_fkey"),
    )

    survey: Mapped[Optional["Survey"]] = relationship(overlaps="embedded_data,workspace", back_populates="embedded_data")
    workspace: Mapped["Workspace"] = relationship(overlaps="embedded_data,survey", back_populates="embedded_data")
    survey_links: Mapped[list["SurveyEmbeddedData"]] = relationship(cascade="all, delete", passive_deletes=True, overlaps="embedded_data_links,survey", back_populates="embedded_data")

class SurveyEmbeddedData(Base):
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

    __tablename__ = "SurveyEmbeddedData"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    workspace_id: Mapped[str] = mapped_column("workspaceId")
    survey_id: Mapped[str] = mapped_column("surveyId")
    embedded_data_id: Mapped[str] = mapped_column("embeddedDataId")
    storage_key: Mapped[str] = mapped_column("storageKey")
    order: Mapped[int]

    __table_args__ = (
        Index("SurveyEmbeddedData_embeddedDataId_idx", "embeddedDataId"),
        Index("SurveyEmbeddedData_surveyId_order_idx", "surveyId", "order"),
        Index("SurveyEmbeddedData_surveyId_embeddedDataId_key", "surveyId", "embeddedDataId", unique=True),
        Index("SurveyEmbeddedData_surveyId_storageKey_key", "surveyId", "storageKey", unique=True),
        ForeignKeyConstraint(["surveyId", "workspaceId"], ["Survey.id", "Survey.workspaceId"], ondelete="CASCADE", onupdate="CASCADE", name="SurveyEmbeddedData_surveyId_workspaceId_fkey"),
        ForeignKeyConstraint(["embeddedDataId", "workspaceId"], ["EmbeddedData.id", "EmbeddedData.workspaceId"], ondelete="CASCADE", onupdate="CASCADE", name="SurveyEmbeddedData_embeddedDataId_workspaceId_fkey"),
    )

    survey: Mapped["Survey"] = relationship(overlaps="embedded_data,survey_links", back_populates="embedded_data_links")
    embedded_data: Mapped["EmbeddedData"] = relationship(overlaps="embedded_data_links,survey", back_populates="survey_links")

class Team(Base):
    """Represents a team within an organization.
    Enables group-based access control and collaboration.

    @property id - Unique identifier for the team
    @property name - Display name of the team
    @property organization - The parent organization
    @property teamUsers - Users who are part of this team
    @property workspaceTeams - Workspaces this team has access to
    """

    __tablename__ = "Team"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    name: Mapped[str]
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="Team_organizationId_fkey"))

    __table_args__ = (
        Index("Team_organizationId_name_key", "organizationId", "name", unique=True),
    )

    organization: Mapped["Organization"] = relationship(back_populates="teams")
    team_users: Mapped[list["TeamUser"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="team")
    workspace_teams: Mapped[list["WorkspaceTeam"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="team")

class TeamUser(Base):
    """Links users to teams with specific roles.
    Manages team membership and permissions.

    @property team - The associated team
    @property user - The team member
    @property role - User's role within the team
    """

    __tablename__ = "TeamUser"

    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    team_id: Mapped[str] = mapped_column("teamId", ForeignKey("Team.id", ondelete="CASCADE", onupdate="CASCADE", name="TeamUser_teamId_fkey"), primary_key=True)
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="TeamUser_userId_fkey"), primary_key=True)
    role: Mapped[str] = mapped_column(Enum("admin", "contributor", name="TeamUserRole"))

    __table_args__ = (
        Index("TeamUser_userId_idx", "userId"),
    )

    team: Mapped["Team"] = relationship(back_populates="team_users")
    user: Mapped["User"] = relationship(back_populates="team_users")

class WorkspaceTeam(Base):
    """Defines team access to specific workspaces.
    Manages workspace-level permissions for teams.

    @property workspace - The accessed workspace
    @property team - The team receiving access
    @property permission - Level of access granted
    """

    __tablename__ = "WorkspaceTeam"

    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    workspace_id: Mapped[str] = mapped_column("workspaceId", ForeignKey("Workspace.id", ondelete="CASCADE", onupdate="CASCADE", name="WorkspaceTeam_workspaceId_fkey"), primary_key=True)
    team_id: Mapped[str] = mapped_column("teamId", ForeignKey("Team.id", ondelete="CASCADE", onupdate="CASCADE", name="WorkspaceTeam_teamId_fkey"), primary_key=True)
    permission: Mapped[str] = mapped_column(Enum("read", "readWrite", "manage", name="WorkspaceTeamPermission"), default="read", server_default=text("'read'"))

    __table_args__ = (
        Index("WorkspaceTeam_teamId_idx", "teamId"),
    )

    workspace: Mapped["Workspace"] = relationship(back_populates="workspace_teams")
    team: Mapped["Team"] = relationship(back_populates="workspace_teams")

class Chart(Base):
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

    __tablename__ = "Chart"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    name: Mapped[str]
    type: Mapped[str] = mapped_column(Enum("area", "bar", "pie", "big_number", "matrix", name="ChartType"))
    workspace_id: Mapped[str] = mapped_column("workspaceId", ForeignKey("Workspace.id", ondelete="CASCADE", onupdate="CASCADE", name="Chart_workspaceId_fkey"))
    query: Mapped[dict[str, Any]] = mapped_column(JSONB, default=lambda: {}, server_default=text("'{}'"))
    """[ChartQuery] - Cube.js query configuration"""
    config: Mapped[dict[str, Any]] = mapped_column(JSONB, default=lambda: {}, server_default=text("'{}'"))
    """[ChartConfig] - Visualization configuration (colors, labels, formatting)"""
    created_by: Mapped[Optional[str]] = mapped_column("createdBy", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="Chart_createdBy_fkey"))
    feedback_directory_id: Mapped[str] = mapped_column("feedbackDirectoryId", ForeignKey("FeedbackDirectory.id", ondelete="RESTRICT", onupdate="CASCADE", name="Chart_feedbackDirectoryId_fkey"))

    __table_args__ = (
        Index("Chart_workspaceId_created_at_idx", "workspaceId", "created_at"),
        Index("Chart_feedbackDirectoryId_idx", "feedbackDirectoryId"),
        Index("Chart_workspaceId_name_key", "workspaceId", "name", unique=True),
    )

    workspace: Mapped["Workspace"] = relationship(back_populates="charts")
    creator: Mapped[Optional["User"]] = relationship(back_populates="charts")
    feedback_directory: Mapped["FeedbackDirectory"] = relationship(back_populates="charts")
    widgets: Mapped[list["DashboardWidget"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="chart")

class Dashboard(Base):
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

    __tablename__ = "Dashboard"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    name: Mapped[str]
    workspace_id: Mapped[str] = mapped_column("workspaceId", ForeignKey("Workspace.id", ondelete="CASCADE", onupdate="CASCADE", name="Dashboard_workspaceId_fkey"))
    created_by: Mapped[Optional[str]] = mapped_column("createdBy", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="Dashboard_createdBy_fkey"))
    filters: Mapped[dict[str, Any]] = mapped_column(JSONB, default=lambda: [], server_default=text("'[]'"))

    __table_args__ = (
        Index("Dashboard_workspaceId_created_at_idx", "workspaceId", "created_at"),
        Index("Dashboard_workspaceId_name_key", "workspaceId", "name", unique=True),
    )

    workspace: Mapped["Workspace"] = relationship(back_populates="dashboards")
    creator: Mapped[Optional["User"]] = relationship(back_populates="dashboards")
    widgets: Mapped[list["DashboardWidget"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="dashboard")

class DashboardWidget(Base):
    """Represents a chart widget on a dashboard.
    Widgets are positioned using a grid layout system.

    @property id - Unique identifier for the widget
    @property dashboard - The dashboard this widget belongs to
    @property chart - The chart displayed in this widget (chart name is used for display)
    @property layout - Grid layout configuration (x, y, width, height)
    @property order - Display order within the dashboard
    """

    __tablename__ = "DashboardWidget"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    dashboard_id: Mapped[str] = mapped_column("dashboardId", ForeignKey("Dashboard.id", ondelete="CASCADE", onupdate="CASCADE", name="DashboardWidget_dashboardId_fkey"))
    chart_id: Mapped[str] = mapped_column("chartId", ForeignKey("Chart.id", ondelete="CASCADE", onupdate="CASCADE", name="DashboardWidget_chartId_fkey"))
    layout: Mapped[dict[str, Any]] = mapped_column(JSONB, default=lambda: {"x": 0, "y": 0, "w": 4, "h": 3}, server_default=text("'{\"x\":0,\"y\":0,\"w\":4,\"h\":3}'"))
    """[WidgetLayout] - Grid layout: { x, y, w, h }"""
    order: Mapped[int] = mapped_column(default=0, server_default=text("0"))

    __table_args__ = (
        Index("DashboardWidget_dashboardId_order_idx", "dashboardId", "order"),
        Index("DashboardWidget_dashboardId_chartId_key", "dashboardId", "chartId", unique=True),
    )

    dashboard: Mapped["Dashboard"] = relationship(back_populates="widgets")
    chart: Mapped["Chart"] = relationship(back_populates="widgets")

class DeletionCleanup(Base):
    """Work left over after a delete commits (ENG-3612, ENG-3614: deleted means deleted): Hub records and
    storage files that can't be removed inside the database transaction. A row is written in the same
    transaction as the delete, so a delete that commits always gets its cleanup and one that rolls back
    queues nothing. A worker drains it with retries, never giving up, and deletes each row once done.
    """

    __tablename__ = "DeletionCleanup"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid2.create_id())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    kind: Mapped[str] = mapped_column(Enum("hubSurvey", "hubResponses", "storageFiles", "storageSurveyFolder", "hubResponseOutcome", name="DeletionCleanupKind"))
    organization_id: Mapped[str] = mapped_column("organizationId")
    """Scoping and logs. Not foreign keys: the rows they name are usually gone already."""
    workspace_id: Mapped[str] = mapped_column("workspaceId")
    survey_id: Mapped[str] = mapped_column("surveyId")
    tenant_ids: Mapped[list[str]] = mapped_column("tenantIds", ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))
    """Hub kinds: the organisation's feedback directories, read inside the delete's transaction."""
    response_ids: Mapped[list[str]] = mapped_column("responseIds", ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))
    """`hubResponses`: the deleted responses whose records go. `hubResponseOutcome`: exactly one live response."""
    file_keys: Mapped[list[str]] = mapped_column("fileKeys", ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))
    """`storageFiles`: the object keys or URLs to delete."""
    attempts: Mapped[int] = mapped_column(default=0, server_default=text("0"))
    next_attempt_at: Mapped[datetime] = mapped_column("nextAttemptAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    """Not drained before this, so a late pipeline write can't slip in behind a survey's cleanup."""
    last_error: Mapped[Optional[str]] = mapped_column("lastError")
    """The last failure: a short code and message, never record or file content."""

    __table_args__ = (
        Index("DeletionCleanup_nextAttemptAt_idx", "nextAttemptAt"),
        Index("DeletionCleanup_organizationId_created_at_idx", "organizationId", "created_at"),
    )

class AttachmentExport(Base):
    """A survey's response attachments built into one ZIP in object storage by a background job and
    downloaded through a short-lived presigned URL, so an interrupted download can resume (ENG-6862).
    Only the requester reads it. The reconcile job deletes the archive at `expiresAt`; deleting the
    survey deletes it with the survey's storage prefix.

    A partial unique index in the migration allows one queued or running export per requester and
    survey, so a double click returns the export already in flight.
    """

    __tablename__ = "AttachmentExport"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid2.create_id())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    survey_id: Mapped[str] = mapped_column("surveyId")
    workspace_id: Mapped[str] = mapped_column("workspaceId", ForeignKey("Workspace.id", ondelete="CASCADE", onupdate="CASCADE", name="AttachmentExport_workspaceId_fkey"))
    requested_by_id: Mapped[Optional[str]] = mapped_column("requestedById", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="AttachmentExport_requestedById_fkey"))
    status: Mapped[str] = mapped_column(Enum("queued", "running", "completed", "failed", "expired", name="AttachmentExportStatus"), default="queued", server_default=text("'queued'"))
    filter_criteria: Mapped[Optional[dict[str, Any]]] = mapped_column("filterCriteria", JSONB(none_as_null=True))
    """The response filter at request time, a `TResponseFilterCriteria`; null exports every response."""
    time_zone: Mapped[str] = mapped_column("timeZone")
    """IANA zone the archive's folder names are stamped in: the org's display zone at request time."""
    dispatch_count: Mapped[int] = mapped_column("dispatchCount", default=0, server_default=text("0"))
    """Bumped on every enqueue and part of the job id, so a re-dispatch never dedupes against a job an
    older pod dropped as unknown during a rolling upgrade.
    """
    file_count: Mapped[int] = mapped_column("fileCount")
    """The pre-flight's counts, shown while the archive builds."""
    response_count: Mapped[int] = mapped_column("responseCount")
    included_file_count: Mapped[Optional[int]] = mapped_column("includedFileCount")
    """What the archive actually holds, once built."""
    total_bytes: Mapped[Optional[int]] = mapped_column("totalBytes", BigInteger)
    truncated: Mapped[bool] = mapped_column(default=False, server_default=text("false"))
    storage_key: Mapped[Optional[str]] = mapped_column("storageKey")
    error_code: Mapped[Optional[str]] = mapped_column("errorCode")
    """A short code (`stale`, `storage_upload_failed`, …), never a message or file name."""
    started_at: Mapped[Optional[datetime]] = mapped_column("startedAt")
    finished_at: Mapped[Optional[datetime]] = mapped_column("finishedAt")
    expires_at: Mapped[Optional[datetime]] = mapped_column("expiresAt")

    __table_args__ = (
        Index("AttachmentExport_surveyId_requestedById_created_at_idx", "surveyId", "requestedById", "created_at"),
        Index("AttachmentExport_status_created_at_idx", "status", "created_at"),
        Index("AttachmentExport_expiresAt_idx", "expiresAt"),
        Index("AttachmentExport_workspaceId_idx", "workspaceId"),
        ForeignKeyConstraint(["surveyId", "workspaceId"], ["Survey.id", "Survey.workspaceId"], ondelete="CASCADE", onupdate="CASCADE", name="AttachmentExport_surveyId_workspaceId_fkey"),
    )

    survey: Mapped["Survey"] = relationship(overlaps="attachment_exports,workspace", back_populates="attachment_exports")
    workspace: Mapped["Workspace"] = relationship(overlaps="attachment_exports,survey", back_populates="attachment_exports")
    requested_by: Mapped[Optional["User"]] = relationship(back_populates="attachment_exports")

class FeedbackSource(Base):
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

    __tablename__ = "FeedbackSource"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    name: Mapped[str]
    type: Mapped[str] = mapped_column(Enum("formbricks_survey", "csv", name="FeedbackSourceType"))
    status: Mapped[str] = mapped_column(Enum("active", "paused", "error", name="FeedbackSourceStatus"), default="active", server_default=text("'active'"))
    import_mode: Mapped[str] = mapped_column(Enum("completedOnly", "all", name="FeedbackSourceImportMode"), default="completedOnly", server_default=text("'completedOnly'"))
    """Defaults to completedOnly so a new source agrees with the live ingestion path, which only
    fires on responseFinished. Existing rows take the default; nothing is re-imported.
    """
    element_scope: Mapped[str] = mapped_column("elementScope", Enum("all", "specific", name="FeedbackSourceElementScope"), default="specific", server_default=text("'specific'"))
    workspace_id: Mapped[str] = mapped_column("workspaceId", ForeignKey("Workspace.id", ondelete="CASCADE", onupdate="CASCADE", name="FeedbackSource_workspaceId_fkey"))
    feedback_directory_id: Mapped[str] = mapped_column("feedbackDirectoryId", ForeignKey("FeedbackDirectory.id", ondelete="CASCADE", onupdate="CASCADE", name="FeedbackSource_feedbackDirectoryId_fkey"))
    last_sync_at: Mapped[Optional[datetime]]
    created_by: Mapped[Optional[str]] = mapped_column(ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="FeedbackSource_created_by_fkey"))

    __table_args__ = (
        Index("FeedbackSource_type_idx", "type"),
        Index("FeedbackSource_feedbackDirectoryId_workspaceId_idx", "feedbackDirectoryId", "workspaceId"),
        Index("FeedbackSource_id_workspaceId_key", "id", "workspaceId", unique=True),
        Index("FeedbackSource_workspaceId_name_key", "workspaceId", "name", unique=True),
        ForeignKeyConstraint(["feedbackDirectoryId", "workspaceId"], ["FeedbackDirectoryWorkspace.feedbackDirectoryId", "FeedbackDirectoryWorkspace.workspaceId"], ondelete="NO ACTION", onupdate="CASCADE", name="FeedbackSource_feedbackDirectoryId_workspaceId_fkey"),
    )

    workspace: Mapped["Workspace"] = relationship(overlaps="feedback_directory_workspace,feedback_sources", back_populates="feedback_sources")
    feedback_directory: Mapped["FeedbackDirectory"] = relationship(overlaps="feedback_directory_workspace,feedback_sources", back_populates="feedback_sources")
    creator: Mapped[Optional["User"]] = relationship(back_populates="feedback_sources")
    feedback_directory_workspace: Mapped["FeedbackDirectoryWorkspace"] = relationship(overlaps="feedback_directory,feedback_sources,workspace", back_populates="feedback_sources")
    formbricks_mappings: Mapped[list["FeedbackSourceFormbricksMapping"]] = relationship(cascade="all, delete", passive_deletes=True, overlaps="feedback_source_mappings,survey", back_populates="feedback_source")
    field_mappings: Mapped[list["FeedbackSourceFieldMapping"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="feedback_source")

class FeedbackSourceFormbricksMapping(Base):
    """Maps survey elements to Hub FeedbackRecords for Formbricks feedback sources.
    Each row represents one element that will create FeedbackRecords when answered.

    @property id - Unique identifier for the mapping
    @property feedbackSource - The parent feedback source
    @property survey - The survey containing the element
    @property elementId - The element ID within the survey (from blocks[].elements[].id)
    @property hubFieldType - The field_type to use in Hub (text, nps, rating, etc.)
    @property customFieldLabel - Optional override for the element headline as field_label in Hub
    """

    __tablename__ = "FeedbackSourceFormbricksMapping"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    feedback_source_id: Mapped[str]
    workspace_id: Mapped[str] = mapped_column("workspaceId")
    survey_id: Mapped[str] = mapped_column("surveyId")
    element_id: Mapped[str] = mapped_column("elementId")
    hub_field_type: Mapped[str] = mapped_column("hubFieldType", Enum("text", "categorical", "nps", "csat", "ces", "rating", "number", "boolean", "date", name="HubFieldType"))
    custom_field_label: Mapped[Optional[str]]

    __table_args__ = (
        Index("FeedbackSourceFormbricksMapping_workspaceId_surveyId_idx", "workspaceId", "surveyId"),
        Index("FeedbackSourceFormbricksMapping_surveyId_idx", "surveyId"),
        Index("FeedbackSourceFormbricksMapping_workspaceId_feedbackSourceId_su", "workspaceId", "feedback_source_id", "surveyId", "elementId", unique=True),
        ForeignKeyConstraint(["feedback_source_id", "workspaceId"], ["FeedbackSource.id", "FeedbackSource.workspaceId"], ondelete="CASCADE", onupdate="CASCADE", name="FeedbackSourceFormbricksMapping_feedback_source_id_workspa_fkey"),
        ForeignKeyConstraint(["surveyId", "workspaceId"], ["Survey.id", "Survey.workspaceId"], ondelete="CASCADE", onupdate="CASCADE", name="FeedbackSourceFormbricksMapping_surveyId_workspaceId_fkey"),
    )

    feedback_source: Mapped["FeedbackSource"] = relationship(overlaps="feedback_source_mappings,survey", back_populates="formbricks_mappings")
    survey: Mapped["Survey"] = relationship(overlaps="feedback_source,formbricks_mappings", back_populates="feedback_source_mappings")

class FeedbackSourceFieldMapping(Base):
    """Generic field mapping for non-Formbricks feedback sources (CSV today; webhook/email/slack later).
    Maps source fields to Hub FeedbackRecord fields.

    @property id - Unique identifier for the mapping
    @property feedbackSource - The parent feedback source
    @property sourceFieldId - Field path for webhook (e.g., "user.id"), column name for CSV
    @property targetFieldId - Hub field (collected_at, field_id, value_text, etc.)
    @property staticValue - If set, use this value instead of reading from sourceFieldId
    """

    __tablename__ = "FeedbackSourceFieldMapping"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    feedback_source_id: Mapped[str]
    workspace_id: Mapped[str] = mapped_column("workspaceId")
    source_field_id: Mapped[str]
    target_field_id: Mapped[str]
    static_value: Mapped[Optional[str]]

    __table_args__ = (
        Index("FeedbackSourceFieldMapping_workspaceId_feedbackSourceId_sourceF", "workspaceId", "feedback_source_id", "source_field_id", "target_field_id", unique=True),
        ForeignKeyConstraint(["feedback_source_id", "workspaceId"], ["FeedbackSource.id", "FeedbackSource.workspaceId"], ondelete="CASCADE", onupdate="CASCADE", name="FeedbackSourceFieldMapping_feedback_source_id_workspaceId_fkey"),
    )

    feedback_source: Mapped["FeedbackSource"] = relationship(back_populates="field_mappings")

class FeedbackDirectory(Base):
    """Represents a feedback directory (Hub tenant) owned by an organization.
    Directories group feedback data and are assigned to workspaces for access control.

    @property id - Unique identifier for the directory
    @property name - Display name of the directory
    @property isArchived - Soft delete flag
    @property organization - The parent organization
    @property workspaces - Workspaces assigned to this directory
    """

    __tablename__ = "FeedbackDirectory"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    name: Mapped[str]
    is_archived: Mapped[bool] = mapped_column("isArchived", default=False, server_default=text("false"))
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="FeedbackDirectory_organizationId_fkey"))

    __table_args__ = (
        Index("FeedbackDirectory_organizationId_name_key", "organizationId", "name", unique=True),
    )

    organization: Mapped["Organization"] = relationship(back_populates="feedback_directories")
    workspaces: Mapped[list["FeedbackDirectoryWorkspace"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="feedback_directory")
    feedback_sources: Mapped[list["FeedbackSource"]] = relationship(cascade="all, delete", passive_deletes=True, overlaps="feedback_directory_workspace,feedback_sources", back_populates="feedback_directory")
    charts: Mapped[list["Chart"]] = relationship(passive_deletes="all", back_populates="feedback_directory")

class FeedbackDirectoryWorkspace(Base):
    """Links feedback directories to workspaces (workspaces).
    Manages which workspaces can access a given directory.

    @property feedbackDirectory - The directory being accessed
    @property workspace - The workspace receiving access
    """

    __tablename__ = "FeedbackDirectoryWorkspace"

    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    feedback_directory_id: Mapped[str] = mapped_column("feedbackDirectoryId", ForeignKey("FeedbackDirectory.id", ondelete="CASCADE", onupdate="CASCADE", name="FeedbackDirectoryWorkspace_feedbackDirectoryId_fkey"), primary_key=True)
    workspace_id: Mapped[str] = mapped_column("workspaceId", ForeignKey("Workspace.id", ondelete="CASCADE", onupdate="CASCADE", name="FeedbackDirectoryWorkspace_workspaceId_fkey"), primary_key=True)

    __table_args__ = (
        Index("FeedbackDirectoryWorkspace_workspaceId_idx", "workspaceId"),
    )

    feedback_directory: Mapped["FeedbackDirectory"] = relationship(back_populates="workspaces")
    workspace: Mapped["Workspace"] = relationship(back_populates="feedback_directory_workspaces")
    feedback_sources: Mapped[list["FeedbackSource"]] = relationship(passive_deletes="all", overlaps="feedback_directory,feedback_sources,workspace", back_populates="feedback_directory_workspace")

class AuthzedProjectionScopeState(Base):
    """ENG-3282: per-scope readiness of a SpiceDB projection. A row with `readyAt` set means the scope was
    backfilled and audited clean, so the evaluator may enforce it. Written only by the backfill CLI.
    """

    __tablename__ = "AuthzedProjectionScopeState"

    scope: Mapped[str] = mapped_column(primary_key=True)
    ready_at: Mapped[Optional[datetime]] = mapped_column("readyAt")
    ready_by: Mapped[Optional[str]] = mapped_column("readyBy")
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

class AuthzedProjectionOutbox(Base):
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

    __tablename__ = "AuthzedProjectionOutbox"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: str(uuid_mod.uuid4()))
    target_type: Mapped[str] = mapped_column("targetType")
    primary_id: Mapped[str] = mapped_column("primaryId")
    secondary_id: Mapped[Optional[str]] = mapped_column("secondaryId")
    is_revocation: Mapped[bool] = mapped_column("isRevocation", default=False, server_default=text("false"))
    attempts: Mapped[int] = mapped_column(default=0, server_default=text("0"))
    """Total delivery attempts. Drives the retry backoff and reports how hard delivery was tried."""
    permanent_failures: Mapped[int] = mapped_column("permanentFailures", default=0, server_default=text("0"))
    """Non-retryable failures attributable to this event alone. The only input to dead-lettering, so
    that a SpiceDB outage of any duration cannot dead-letter events that were never the problem.
    """
    available_at: Mapped[datetime] = mapped_column("availableAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    leased_at: Mapped[Optional[datetime]] = mapped_column("leasedAt")
    lease_expires_at: Mapped[Optional[datetime]] = mapped_column("leaseExpiresAt")
    lease_owner: Mapped[Optional[str]] = mapped_column("leaseOwner")
    processed_at: Mapped[Optional[datetime]] = mapped_column("processedAt")
    dead_lettered_at: Mapped[Optional[datetime]] = mapped_column("deadLetteredAt")
    last_attempt_at: Mapped[Optional[datetime]] = mapped_column("lastAttemptAt")
    last_error_code: Mapped[Optional[str]] = mapped_column("lastErrorCode")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

class OnlineUpgradeState(Base):
    __tablename__ = "OnlineUpgradeState"

    release: Mapped[str] = mapped_column(primary_key=True)
    target_image: Mapped[str] = mapped_column("targetImage")
    phase: Mapped[str]
    completed_phases: Mapped[list[str]] = mapped_column("completedPhases", ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))
    survey_cursor: Mapped[Optional[str]] = mapped_column("surveyCursor")
    survey_count: Mapped[int] = mapped_column("surveyCount", BigInteger, default=0, server_default=text("0"))
    baseline_completed_at: Mapped[Optional[datetime]] = mapped_column("baselineCompletedAt")
    forward_only_at: Mapped[Optional[datetime]] = mapped_column("forwardOnlyAt")
    queue_paused_at: Mapped[Optional[datetime]] = mapped_column("queuePausedAt")
    queue_pause_deadline_at: Mapped[Optional[datetime]] = mapped_column("queuePauseDeadlineAt")
    queue_resumed_at: Mapped[Optional[datetime]] = mapped_column("queueResumedAt")
    activated_at: Mapped[Optional[datetime]] = mapped_column("activatedAt")
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

class RetentionPolicy(Base):
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

    __tablename__ = "RetentionPolicy"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid2.create_id())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="RetentionPolicy_organizationId_fkey"))
    entity: Mapped[str] = mapped_column(Enum("responses", "surveys", "members", name="RetentionEntity"))
    enabled: Mapped[bool] = mapped_column(default=False, server_default=text("false"))
    enabled_at: Mapped[Optional[datetime]] = mapped_column("enabledAt")
    """When the current configuration took effect: the policy was switched on or unpaused, or changed
    in a way that could bring an action earlier (a shorter period or warning, different
    conditions). Lengthening does not move it. The warning always runs in full from here, so no
    change can delete data on the night it is made (ENG-3614). Also the clock for members with no
    recorded sign-in.
    """
    warn_days: Mapped[int] = mapped_column("warnDays", default=60, server_default=text("60"))
    """How long before the action the notice goes out. The action never runs before the notice has
    been out for this long, so the warning always runs in full.
    """
    period_days: Mapped[int] = mapped_column("periodDays")
    """The period, counted from the target's clock, after which the policy acts: responses are
    deleted, surveys archived (the archive purge deletes them 30 days later), members deactivated.
    """
    conditions: Mapped[list[str]] = mapped_column(ARRAY(Enum("noResponse", "noChange", "createdBefore", name="RetentionSurveyCondition")), nullable=True, default=lambda: [], server_default=text("ARRAY[]::\"RetentionSurveyCondition\"[]"))
    """Surveys policy only."""
    deferred_at: Mapped[Optional[datetime]] = mapped_column("deferredAt")
    """The last night the sweep's time ran out before it reached this policy's organisation, licensed.
    The organisation was due and the sweep was running, so it counts as reached for the gap that
    restarts the warning (`RETENTION_SWEEP_GAP_MS`): a backlog of organisations never restarts warnings.
    """
    updated_by_id: Mapped[Optional[str]] = mapped_column("updatedById", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="RetentionPolicy_updatedById_fkey"))

    __table_args__ = (
        Index("RetentionPolicy_updatedById_idx", "updatedById"),
        Index("RetentionPolicy_organizationId_entity_key", "organizationId", "entity", unique=True),
    )

    organization: Mapped["Organization"] = relationship(back_populates="retention_policies")
    updated_by: Mapped[Optional["User"]] = relationship(back_populates="retention_policies_updated")

class RetentionExemption(Base):
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

    __tablename__ = "RetentionExemption"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid2.create_id())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="RetentionExemption_organizationId_fkey"))
    entity: Mapped[str] = mapped_column(Enum("responses", "surveys", "members", name="RetentionEntity"))
    survey_id: Mapped[Optional[str]] = mapped_column("surveyId", ForeignKey("Survey.id", ondelete="CASCADE", onupdate="CASCADE", name="RetentionExemption_surveyId_fkey"))
    until: Mapped[datetime]
    source: Mapped[str] = mapped_column(Enum("manual", "restore", name="RetentionExemptionSource"), default="manual", server_default=text("'manual'"))
    reason: Mapped[Optional[str]]
    """Required for `manual`, absent for `restore` (`RetentionExemption_reason_check`)."""
    created_by_id: Mapped[Optional[str]] = mapped_column("createdById", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="RetentionExemption_createdById_fkey"))
    revoked_at: Mapped[Optional[datetime]] = mapped_column("revokedAt")
    revoked_by_id: Mapped[Optional[str]] = mapped_column("revokedById", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="RetentionExemption_revokedById_fkey"))

    __table_args__ = (
        Index("RetentionExemption_organizationId_created_at_id_idx", "organizationId", "created_at", "id"),
        Index("RetentionExemption_surveyId_idx", "surveyId"),
        Index("RetentionExemption_createdById_idx", "createdById"),
        Index("RetentionExemption_revokedById_idx", "revokedById"),
    )

    organization: Mapped["Organization"] = relationship(back_populates="retention_exemptions")
    survey: Mapped[Optional["Survey"]] = relationship(back_populates="retention_exemptions")
    created_by: Mapped[Optional["User"]] = relationship(foreign_keys=[created_by_id], back_populates="exemptions_created")
    revoked_by: Mapped[Optional["User"]] = relationship(foreign_keys=[revoked_by_id], back_populates="exemptions_revoked")

class RetentionRun(Base):
    """One sweep of one policy for one organisation: the History tab's rows. `finishedAt` stays null
    while the run is in progress, or if it never finished.
    """

    __tablename__ = "RetentionRun"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid2.create_id())
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="RetentionRun_organizationId_fkey"))
    entity: Mapped[str] = mapped_column(Enum("responses", "surveys", "members", name="RetentionEntity"))
    kind: Mapped[str] = mapped_column(Enum("sweep", "purge", name="RetentionRunKind"), default="sweep", server_default=text("'sweep'"))
    started_at: Mapped[datetime] = mapped_column("startedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    finished_at: Mapped[Optional[datetime]] = mapped_column("finishedAt")
    notified_count: Mapped[int] = mapped_column("notifiedCount", default=0, server_default=text("0"))
    archived_count: Mapped[int] = mapped_column("archivedCount", default=0, server_default=text("0"))
    """Archived surveys, or deactivated members."""
    deleted_count: Mapped[int] = mapped_column("deletedCount", default=0, server_default=text("0"))
    """Deleted responses, or, on a purge run, deleted surveys."""
    skipped_count: Mapped[int] = mapped_column("skippedCount", default=0, server_default=text("0"))
    has_changes: Mapped[bool] = mapped_column("hasChanges", default=False, server_default=text("false"))
    """False when the run changed nothing. History hides those by default."""
    scan_cursor: Mapped[Optional[str]] = mapped_column("scanCursor")
    """Where the run's scan of candidates stopped for want of time, so the next run resumes there instead
    of rescanning the same first targets every night. Null when the scan reached the end.
    """

    __table_args__ = (
        Index("RetentionRun_organizationId_startedAt_id_idx", "organizationId", "startedAt", "id"),
    )

    organization: Mapped["Organization"] = relationship(back_populates="retention_runs")
    items: Mapped[list["RetentionRunItem"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="run")

class RetentionRunItem(Base):
    """What a run did, one row per target: the History CSV. Ids, counts, names and the recipient's email,
    never response content. `targetId` is not a foreign key on purpose: the target is often gone, and
    `targetName` is a snapshot taken when the row is written so the CSV still says what it was.
    """

    __tablename__ = "RetentionRunItem"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid2.create_id())
    run_id: Mapped[str] = mapped_column("runId", ForeignKey("RetentionRun.id", ondelete="CASCADE", onupdate="CASCADE", name="RetentionRunItem_runId_fkey"))
    target_type: Mapped[str] = mapped_column("targetType", Enum("survey", "user", name="RetentionTargetType"))
    target_id: Mapped[str] = mapped_column("targetId")
    target_name: Mapped[Optional[str]] = mapped_column("targetName")
    action: Mapped[str] = mapped_column(Enum("notified", "archived", "deactivated", "deleted", "skipped", name="RetentionRunItemAction"))
    count: Mapped[int] = mapped_column(default=1, server_default=text("1"))
    """1, except for grouped response deletions, where it is the number of responses deleted."""
    recipient: Mapped[Optional[str]]
    """Who was emailed, for `notified` rows."""
    skip_reason: Mapped[Optional[str]] = mapped_column("skipReason", Enum("exempt", "lastOwner", "otherOrganization", "noRecipient", name="RetentionSkipReason"))

    __table_args__ = (
        Index("RetentionRunItem_runId_id_idx", "runId", "id"),
        Index("RetentionRunItem_targetId_idx", "targetId"),
    )

    run: Mapped["RetentionRun"] = relationship(back_populates="items")

class RetentionNotice(Base):
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

    __tablename__ = "RetentionNotice"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid2.create_id())
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="RetentionNotice_organizationId_fkey"))
    entity: Mapped[str] = mapped_column(Enum("responses", "surveys", "members", name="RetentionEntity"))
    survey_id: Mapped[Optional[str]] = mapped_column("surveyId", ForeignKey("Survey.id", ondelete="CASCADE", onupdate="CASCADE", name="RetentionNotice_surveyId_fkey"))
    user_id: Mapped[Optional[str]] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="RetentionNotice_userId_fkey"))
    sent_at: Mapped[datetime] = mapped_column("sentAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    """When the notice was claimed for sending. A claim that never got delivered can be taken over."""
    delivered_at: Mapped[Optional[datetime]] = mapped_column("deliveredAt")
    """When the notice was handed to the mail transport, or recorded without one (no SMTP). The
    warning only counts from a delivered notice, so a crash between claim and send can't let the
    action run without it.
    """
    email_sent: Mapped[bool] = mapped_column("emailSent", default=False, server_default=text("false"))
    """Whether an email actually went out. False when SMTP isn't configured."""
    claim_token: Mapped[Optional[str]] = mapped_column("claimToken")
    """Set by each claim; delivery is recorded only for the claim that sent the email, so a slow sender
    can't mark a newer claim delivered.
    """
    clock_at: Mapped[Optional[datetime]] = mapped_column("clockAt")
    """The target's clock the notice was computed for (surveys and members). A notice counts only while
    the target's clock is still this one, so activity between the sweep's read and its claim voids
    it. Null for the responses reminder, which is sent once per survey and follows no one clock, and
    for a survey notice voided for good: a person deleted or moved one of its responses, or the
    responses policy's deletion left it computed for another clock (see `NOTICE_FOLLOWS_CLOCK`).
    """

    __table_args__ = (
        Index("RetentionNotice_organizationId_idx", "organizationId"),
        Index("RetentionNotice_surveyId_entity_key", "surveyId", "entity", unique=True),
        Index("RetentionNotice_userId_organizationId_entity_key", "userId", "organizationId", "entity", unique=True),
    )

    organization: Mapped["Organization"] = relationship(back_populates="retention_notices")
    survey: Mapped[Optional["Survey"]] = relationship(back_populates="retention_notices")
    user: Mapped[Optional["User"]] = relationship(back_populates="retention_notices")

class RetentionSurveyRestore(Base):
    """When a survey was last restored from the archive. A restore starts the surveys policy's "no response"
    and "no change" periods over (`getSurveyRetentionClock`); only ever moves forward. A table of its own
    so retention adds nothing to the survey row.
    """

    __tablename__ = "RetentionSurveyRestore"

    survey_id: Mapped[str] = mapped_column("surveyId", ForeignKey("Survey.id", ondelete="CASCADE", onupdate="CASCADE", name="RetentionSurveyRestore_surveyId_fkey"), primary_key=True)
    restored_at: Mapped[datetime] = mapped_column("restoredAt")

    survey: Mapped["Survey"] = relationship(back_populates="retention_restore")

class Workflow(Base):
    __tablename__ = "Workflow"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid2.create_id())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    name: Mapped[str]
    description: Mapped[Optional[str]]
    status: Mapped[str] = mapped_column(Enum("draft", "enabled", "disabled", "archived", name="WorkflowStatus"), default="draft", server_default=text("'draft'"))
    workspace_id: Mapped[str] = mapped_column("workspaceId", ForeignKey("Workspace.id", ondelete="CASCADE", onupdate="CASCADE", name="Workflow_workspaceId_fkey"))
    created_by: Mapped[Optional[str]] = mapped_column("createdBy", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="Workflow_createdBy_fkey"))
    definition: Mapped[dict[str, Any]] = mapped_column(JSONB)
    """[WorkflowDefinition]"""

    __table_args__ = (
        Index("Workflow_workspaceId_status_idx", "workspaceId", "status"),
        Index("Workflow_workspaceId_updated_at_idx", "workspaceId", "updated_at"),
        Index("Workflow_id_workspaceId_key", "id", "workspaceId", unique=True),
    )

    workspace: Mapped["Workspace"] = relationship(back_populates="workflows")
    creator: Mapped[Optional["User"]] = relationship(back_populates="workflows")
    versions: Mapped[list["WorkflowVersion"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="workflow")
    runs: Mapped[list["WorkflowRun"]] = relationship(cascade="all, delete", passive_deletes=True, overlaps="workflow_runs,workspace", back_populates="workflow")

class WorkflowVersion(Base):
    __tablename__ = "WorkflowVersion"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid2.create_id())
    workflow_id: Mapped[str] = mapped_column("workflowId")
    workspace_id: Mapped[str] = mapped_column("workspaceId")
    version: Mapped[int]
    definition: Mapped[dict[str, Any]] = mapped_column(JSONB)
    """[WorkflowExecutableDefinition]"""
    published_at: Mapped[datetime] = mapped_column("publishedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    published_by: Mapped[Optional[str]] = mapped_column("publishedBy", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="WorkflowVersion_publishedBy_fkey"))

    __table_args__ = (
        Index("WorkflowVersion_workflowId_workspaceId_idx", "workflowId", "workspaceId"),
        Index("WorkflowVersion_publishedBy_idx", "publishedBy"),
        Index("WorkflowVersion_workflowId_version_key", "workflowId", "version", unique=True),
        ForeignKeyConstraint(["workflowId", "workspaceId"], ["Workflow.id", "Workflow.workspaceId"], ondelete="CASCADE", onupdate="CASCADE", name="WorkflowVersion_workflowId_workspaceId_fkey"),
    )

    workflow: Mapped["Workflow"] = relationship(back_populates="versions")
    publisher: Mapped[Optional["User"]] = relationship(back_populates="workflow_versions_published")
    runs: Mapped[list["WorkflowRun"]] = relationship(back_populates="workflow_version")

class WorkflowRun(Base):
    __tablename__ = "WorkflowRun"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid2.create_id())
    created_at: Mapped[datetime] = mapped_column(default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column(default=utc_now, onupdate=utc_now)
    workflow_id: Mapped[str] = mapped_column("workflowId")
    workspace_id: Mapped[str] = mapped_column("workspaceId", ForeignKey("Workspace.id", ondelete="CASCADE", onupdate="CASCADE", name="WorkflowRun_workspaceId_fkey"))
    workflow_version_id: Mapped[Optional[str]] = mapped_column("workflowVersionId", ForeignKey("WorkflowVersion.id", ondelete="SET NULL", onupdate="CASCADE", name="WorkflowRun_workflowVersionId_fkey"))
    response_id: Mapped[Optional[str]] = mapped_column("responseId", ForeignKey("Response.id", ondelete="CASCADE", onupdate="CASCADE", name="WorkflowRun_responseId_fkey"))
    status: Mapped[str] = mapped_column(Enum("queued", "running", "completed", "failed", "canceled", name="WorkflowRunStatus"), default="queued", server_default=text("'queued'"))
    trigger_type: Mapped[str] = mapped_column("triggerType")
    survey_id: Mapped[Optional[str]] = mapped_column("surveyId")
    is_dry_run: Mapped[bool] = mapped_column("isDryRun", default=False, server_default=text("false"))
    idempotency_key: Mapped[Optional[str]] = mapped_column("idempotencyKey")
    attempt: Mapped[int] = mapped_column(default=0, server_default=text("0"))
    next_attempt_at: Mapped[Optional[datetime]] = mapped_column("nextAttemptAt")
    last_error_at: Mapped[Optional[datetime]] = mapped_column("lastErrorAt")
    trigger_payload: Mapped[dict[str, Any]] = mapped_column("triggerPayload", JSONB)
    """[WorkflowTriggerRunPayload]"""
    data: Mapped[dict[str, Any]] = mapped_column(JSONB, default=lambda: {}, server_default=text("'{}'"))
    """[WorkflowRunData]"""
    error: Mapped[Optional[str]]
    dispatched_at: Mapped[Optional[datetime]] = mapped_column("dispatchedAt")
    started_at: Mapped[Optional[datetime]] = mapped_column("startedAt")
    finished_at: Mapped[Optional[datetime]] = mapped_column("finishedAt")

    __table_args__ = (
        Index("WorkflowRun_workspaceId_created_at_idx", "workspaceId", "created_at"),
        Index("WorkflowRun_workflowId_workspaceId_created_at_idx", "workflowId", "workspaceId", "created_at"),
        Index("WorkflowRun_status_created_at_idx", "status", "created_at"),
        Index("WorkflowRun_responseId_idx", "responseId"),
        Index("WorkflowRun_isDryRun_idx", "isDryRun"),
        Index("WorkflowRun_nextAttemptAt_idx", "nextAttemptAt"),
        Index("WorkflowRun_workflowId_idempotencyKey_key", "workflowId", "idempotencyKey", unique=True),
        ForeignKeyConstraint(["workflowId", "workspaceId"], ["Workflow.id", "Workflow.workspaceId"], ondelete="CASCADE", onupdate="CASCADE", name="WorkflowRun_workflowId_workspaceId_fkey"),
    )

    workflow: Mapped["Workflow"] = relationship(overlaps="workflow_runs,workspace", back_populates="runs")
    workspace: Mapped["Workspace"] = relationship(overlaps="runs,workflow", back_populates="workflow_runs")
    workflow_version: Mapped[Optional["WorkflowVersion"]] = relationship(back_populates="runs")
    response: Mapped[Optional["Response"]] = relationship(back_populates="workflow_runs")
    logs: Mapped[list["WorkflowRunLog"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="run")

class WorkflowRunLog(Base):
    __tablename__ = "WorkflowRunLog"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid2.create_id())
    run_id: Mapped[str] = mapped_column("runId", ForeignKey("WorkflowRun.id", ondelete="CASCADE", onupdate="CASCADE", name="WorkflowRunLog_runId_fkey"))
    sequence: Mapped[int]
    step_id: Mapped[str] = mapped_column("stepId")
    step_type: Mapped[str] = mapped_column("stepType")
    status: Mapped[str] = mapped_column(Enum("pending", "running", "succeeded", "failed", "skipped", name="WorkflowRunLogStatus"), default="pending", server_default=text("'pending'"))
    input: Mapped[dict[str, Any]] = mapped_column(JSONB, default=lambda: {}, server_default=text("'{}'"))
    """[WorkflowRunLogInput]"""
    output: Mapped[dict[str, Any]] = mapped_column(JSONB, default=lambda: {}, server_default=text("'{}'"))
    """[WorkflowRunLogOutput]"""
    error: Mapped[Optional[str]]
    started_at: Mapped[Optional[datetime]] = mapped_column("startedAt")
    finished_at: Mapped[Optional[datetime]] = mapped_column("finishedAt")

    __table_args__ = (
        Index("WorkflowRunLog_runId_sequence_idx", "runId", "sequence"),
        Index("WorkflowRunLog_runId_stepId_key", "runId", "stepId", unique=True),
    )

    run: Mapped["WorkflowRun"] = relationship(back_populates="logs")
