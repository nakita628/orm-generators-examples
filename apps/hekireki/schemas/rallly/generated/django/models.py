import uuid
from datetime import datetime
from datetime import timezone as dt_timezone
from typing import Any, TypeVar

import cuid  # type: ignore[import-untyped, unused-ignore]
from django.conf import settings
from django.contrib.postgres.fields import ArrayField
from django.contrib.postgres.indexes import HashIndex
from django.db import models
from django.db.backends.base.base import BaseDatabaseWrapper
from django.db.models.expressions import RawSQL
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


_M = TypeVar("_M", bound=models.Model)


class AutoNowQuerySet(models.QuerySet[_M]):
    def update(self, **kwargs: Any) -> int:
        for field in self.model._meta.concrete_fields:
            if getattr(field, "auto_now", False):
                kwargs.setdefault(field.name, getattr(field, "stamp", timezone.now)())
        return super().update(**kwargs)


def uuid4_str() -> str:
    return str(uuid.uuid4())


def cuid_str() -> str:
    return str(cuid.cuid())


def now_milliseconds() -> datetime:
    now = timezone.now()
    return now.replace(microsecond=now.microsecond // 1000 * 1000)


class ApiUsageResource(models.TextChoices):
    poll = "poll"


class SubscriptionStatus(models.TextChoices):
    incomplete = "incomplete"
    incomplete_expired = "incomplete_expired"
    active = "active"
    paused = "paused"
    trialing = "trialing"
    past_due = "past_due"
    canceled = "canceled"
    unpaid = "unpaid"


class SubscriptionInterval(models.TextChoices):
    month = "month"
    year = "year"


class NonprofitApplicationStatus(models.TextChoices):
    approved = "approved"
    rejected = "rejected"
    failed = "failed"


class QueuedEmailStatus(models.TextChoices):
    pending = "pending"
    sent = "sent"
    skipped = "skipped"
    failed = "failed"


class ScheduledEventStatus(models.TextChoices):
    confirmed = "confirmed"
    canceled = "canceled"
    unconfirmed = "unconfirmed"


class ScheduledEventInviteStatus(models.TextChoices):
    pending = "pending"
    accepted = "accepted"
    declined = "declined"
    tentative = "tentative"


class CredentialType(models.TextChoices):
    OAUTH = "OAUTH"


class LicenseType(models.TextChoices):
    PLUS = "PLUS"
    ORGANIZATION = "ORGANIZATION"
    ENTERPRISE = "ENTERPRISE"


class LicenseStatus(models.TextChoices):
    ACTIVE = "ACTIVE"
    REVOKED = "REVOKED"


class ParticipantVisibility(models.TextChoices):
    full = "full"
    scoresOnly = "scoresOnly"
    limited = "limited"


class PollStatus(models.TextChoices):
    open = "open"
    closed = "closed"
    scheduled = "scheduled"
    canceled = "canceled"


class PollKind(models.TextChoices):
    date = "date"
    time = "time"


class PollClosedReason(models.TextChoices):
    auto = "auto"
    manual = "manual"


class VoteType(models.TextChoices):
    yes = "yes"
    no = "no"
    ifNeedBe = "ifNeedBe"


class SpaceUsageResource(models.TextChoices):
    api_poll = "api_poll"
    email_invite = "email_invite"


class SpaceMemberRole(models.TextChoices):
    ADMIN = "ADMIN"
    MEMBER = "MEMBER"


class SpaceType(models.TextChoices):
    personal = "personal"
    work = "work"


class SpaceTier(models.TextChoices):
    hobby = "hobby"
    pro = "pro"


class WebhookDeliveryStatus(models.TextChoices):
    pending = "pending"
    in_flight = "in_flight"
    succeeded = "succeeded"
    failed = "failed"
    exhausted = "exhausted"


class TimeFormat(models.TextChoices):
    hours12 = "hours12"
    hours24 = "hours24"


class UserRole(models.TextChoices):
    admin = "admin"
    user = "user"


class ApiUsage(models.Model):
    """Superseded by SpaceUsage (resource api_poll). Unused; kept so old
    deployments keep working during the switch, then dropped.
    """

    pk = models.CompositePrimaryKey("space_id", "resource", "period")
    space = models.ForeignKey("Space", on_delete=models.CASCADE, related_name="api_usage", db_index=False)
    resource = models.TextField(choices=ApiUsageResource.choices)
    period = models.DateField()
    count = models.IntegerField(default=0)
    near_limit_warned_at = UtcDateTimeField(null=True)
    at_limit_warned_at = UtcDateTimeField(null=True)
    updated_at = UtcAutoNowDateTimeField(auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "api_usage"


class Subscription(models.Model):
    id = models.TextField(primary_key=True, default=None)
    price_id = models.TextField()
    quantity = models.IntegerField(default=1)
    subscription_item_id = models.TextField()
    amount = models.IntegerField()
    discount_percent_off = models.FloatField(null=True)
    discount_amount_off = models.IntegerField(null=True)
    status = models.TextField(choices=SubscriptionStatus.choices)
    active = models.BooleanField()
    currency = models.TextField()
    interval = models.TextField(choices=SubscriptionInterval.choices)
    created_at = UtcDateTimeField(default=now_milliseconds)
    period_start = UtcDateTimeField()
    period_end = UtcDateTimeField()
    cancel_at_period_end = models.BooleanField(default=False)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="subscriptions", db_index=False)
    space = models.ForeignKey("Space", on_delete=models.CASCADE, related_name="subscriptions", db_index=False)
    api_addon_item_id = models.TextField(null=True)
    api_addon_price_id = models.TextField(null=True)
    api_addon_amount = models.IntegerField(null=True)
    api_addon_tier = models.TextField(null=True)
    api_poll_retained_tier = models.TextField(null=True)
    api_poll_retained_until = UtcDateTimeField(null=True)

    class Meta:
        db_table = "subscriptions"
        indexes = [
            models.Index(fields=["user"]),
            models.Index(fields=["space"]),
        ]


class PaymentMethod(models.Model):
    id = models.TextField(primary_key=True, default=None)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="payment_methods", db_index=False)
    type = models.TextField()
    data = models.JSONField()
    created_at = UtcDateTimeField(default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "payment_methods"


class NonprofitApplication(models.Model):
    id = models.TextField(primary_key=True, default=uuid4_str)
    space = models.ForeignKey("Space", on_delete=models.CASCADE, related_name="nonprofit_applications", db_index=False)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="nonprofit_applications", db_index=False)
    organization_name = models.TextField()
    website = models.TextField()
    email_domain = models.TextField()
    status = models.TextField(choices=NonprofitApplicationStatus.choices)
    reason = models.TextField(null=True)
    model = models.TextField(null=True)
    created_at = UtcDateTimeField(default=now_milliseconds)

    class Meta:
        db_table = "nonprofit_applications"
        indexes = [
            models.Index(fields=["space", "created_at"]),
        ]


class QueuedEmail(models.Model):
    """An email waiting to be sent by the queue, one row per recipient. The
    email is built when it is sent, from the row its subject points at.
    """

    id = models.TextField(primary_key=True, default=cuid_str)
    kind = models.TextField()
    """Which handler builds the email. A string rather than an enum so a new
    kind needs no migration; the claim only takes kinds the build knows.
    """
    subject_id = models.TextField()
    """The row the email is about, interpreted by the kind. Not a foreign key:
    kinds point at different tables, and a subject deleted before the send
    makes its handler skip the email.
    """
    batch_id = models.TextField()
    """Groups the emails one action queued, so the action can send its own
    first batch at once and its progress can be read back.
    """
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="queued_emails", db_index=False)
    """The account whose action queued the email. Banning it skips whatever
    is still pending.
    """
    status = models.TextField(choices=QueuedEmailStatus.choices, default=QueuedEmailStatus.pending)
    send_after = UtcDateTimeField(null=True)
    """Not claimed before this time. Unset sends at the next run."""
    attempts = models.IntegerField(default=0)
    claimed_at = UtcDateTimeField(null=True)
    sent_at = UtcDateTimeField(null=True)
    last_error = models.TextField(null=True)
    created_at = UtcDateTimeField(default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "queued_emails"
        indexes = [
            models.Index(fields=["status", "created_at"]),
            models.Index(fields=["status", "updated_at"]),
            models.Index(fields=["batch_id"]),
            models.Index(fields=["user", "status"]),
        ]


class EventType(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    space = models.ForeignKey("Space", on_delete=models.CASCADE, related_name="event_types", db_index=False)
    host = models.ForeignKey("User", on_delete=models.CASCADE, related_name="hosted_event_types", db_index=False)
    name = models.TextField()
    duration = models.IntegerField(db_column="duration_minutes")
    capacity = models.IntegerField(null=True)
    description = models.TextField(null=True)
    location = models.JSONField(null=True)
    created_at = UtcDateTimeField(default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(auto_now=True)
    deleted = models.BooleanField(default=False)
    deleted_at = UtcDateTimeField(null=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "event_types"
        indexes = [
            models.Index(fields=["space"]),
            models.Index(fields=["host"]),
        ]


class ScheduledEvent(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="scheduled_events", db_index=False)
    space = models.ForeignKey("Space", on_delete=models.CASCADE, related_name="scheduled_events", db_index=False)
    event_type = models.ForeignKey("EventType", on_delete=models.SET_NULL, related_name="scheduled_events", null=True, db_index=False)
    sheet_slot = models.OneToOneField("SheetSlot", on_delete=models.SET_NULL, related_name="scheduled_event", null=True)
    title = models.TextField()
    description = models.TextField(null=True)
    location = models.JSONField(null=True)
    conferencing = models.JSONField(null=True)
    capacity = models.IntegerField(null=True)
    created_at = UtcDateTimeField(default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(auto_now=True)
    status = models.TextField(choices=ScheduledEventStatus.choices, default=ScheduledEventStatus.confirmed)
    time_zone = models.TextField(null=True)
    """Venue / anchor timezone for in-person or hybrid events. Null means no fixed venue (e.g. remote-only)."""
    start = UtcDateTimeField()
    end = UtcDateTimeField()
    all_day = models.BooleanField(default=False)
    deleted_at = UtcDateTimeField(null=True)
    sequence = models.IntegerField(default=0)
    uid = models.TextField(unique=True)
    hide_attendees = models.BooleanField(default=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "scheduled_events"
        indexes = [
            models.Index(fields=["space"]),
            models.Index(fields=["user"]),
            models.Index(fields=["event_type"]),
            models.Index(fields=["space", "status", "start"]),
        ]


class RescheduledEventDate(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    scheduled_event = models.ForeignKey("ScheduledEvent", on_delete=models.CASCADE, related_name="rescheduled_dates", db_index=False)
    start = UtcDateTimeField()
    end = UtcDateTimeField()
    all_day = models.BooleanField(default=False)
    created_at = UtcDateTimeField(default=now_milliseconds)

    class Meta:
        db_table = "rescheduled_event_dates"
        indexes = [
            models.Index(fields=["scheduled_event"]),
        ]


class ScheduledEventInvite(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    uid = models.TextField(unique=True)
    scheduled_event = models.ForeignKey("ScheduledEvent", on_delete=models.CASCADE, related_name="invites", db_index=False)
    invitee_name = models.TextField()
    invitee_email = models.TextField()
    user = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="scheduled_event_invites", db_column="invitee_id", null=True, db_index=False)
    invitee_time_zone = models.TextField(null=True)
    invitee_locale = models.TextField(null=True)
    status = models.TextField(choices=ScheduledEventInviteStatus.choices, default=ScheduledEventInviteStatus.pending)
    created_at = UtcDateTimeField(default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "scheduled_event_invites"
        constraints = [
            models.UniqueConstraint(fields=["scheduled_event", "invitee_email"], name="scheduled_event_invites_scheduled_event_id_invitee_email_key"),
        ]
        indexes = [
            models.Index(fields=["user"]),
            models.Index(fields=["invitee_email"]),
        ]


class InstanceSettings(models.Model):
    id = models.IntegerField(primary_key=True, default=1)
    disable_user_registration = models.BooleanField(default=False)
    instance_id = models.UUIDField(unique=True, db_default=RawSQL("gen_random_uuid()", []))
    app_name = models.TextField(null=True)
    primary_color = models.TextField(null=True)
    primary_color_dark = models.TextField(null=True)
    logo = models.TextField(null=True)
    logo_dark = models.TextField(null=True)
    logo_icon = models.TextField(null=True)
    hide_attribution = models.BooleanField(null=True)
    footer_links = models.JSONField(default=list)
    created_at = UtcDateTimeField(default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "instance_settings"


class Credential(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="credentials", db_index=False)
    provider = models.TextField()
    provider_account_id = models.TextField()
    type = models.TextField(choices=CredentialType.choices)
    secret = models.TextField()
    scopes = ArrayField(models.TextField())
    expires_at = UtcDateTimeField(null=True)
    created_at = UtcDateTimeField(default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "credentials"
        constraints = [
            models.UniqueConstraint(fields=["user", "provider", "provider_account_id"], name="credentials_user_id_provider_provider_account_id_key"),
        ]
        indexes = [
            models.Index(fields=["expires_at"], name="credential_expiry_idx"),
        ]


class CalendarConnection(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="calendar_connections", db_index=False)
    provider = models.TextField()
    integration_id = models.TextField()
    provider_account_id = models.TextField()
    email = models.TextField()
    display_name = models.TextField(null=True)
    created_at = UtcDateTimeField(default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(auto_now=True)
    credential = models.ForeignKey("Credential", on_delete=models.CASCADE, related_name="calendar_connections", db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "calendar_connections"
        constraints = [
            models.UniqueConstraint(fields=["user", "provider", "provider_account_id"], name="calendar_connections_user_id_provider_provider_account_id_key"),
        ]


class ProviderCalendar(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    calendar_connection = models.ForeignKey("CalendarConnection", on_delete=models.CASCADE, related_name="provider_calendars", db_index=False)
    provider_calendar_id = models.TextField()
    name = models.TextField()
    is_primary = models.BooleanField(db_column="primary", default=False)
    time_zone = models.TextField(null=True)
    is_selected = models.BooleanField(db_column="selected", default=False)
    is_deleted = models.BooleanField(db_column="deleted", default=False)
    is_writable = models.BooleanField(db_column="writable", default=False)
    provider_data = models.JSONField(null=True)
    last_synced_at = UtcDateTimeField(default=now_milliseconds)
    created_at = UtcDateTimeField(default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "provider_calendars"
        constraints = [
            models.UniqueConstraint(fields=["calendar_connection", "provider_calendar_id"], name="provider_calendars_calendar_connection_id_provider_calendar_id_key"),
        ]
        indexes = [
            models.Index(fields=["calendar_connection", "is_selected"], name="connection_selected_idx"),
            models.Index(fields=["is_primary"], name="primary_calendar_idx"),
            models.Index(fields=["last_synced_at"], name="sync_time_idx"),
        ]


class ConferencingConnection(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="conferencing_connections", db_index=False)
    provider = models.TextField()
    integration_id = models.TextField()
    provider_account_id = models.TextField()
    email = models.TextField()
    display_name = models.TextField(null=True)
    created_at = UtcDateTimeField(default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(auto_now=True)
    credential = models.ForeignKey("Credential", on_delete=models.CASCADE, related_name="conferencing_connections", db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "conferencing_connections"
        constraints = [
            models.UniqueConstraint(fields=["user", "provider", "provider_account_id"], name="conferencing_connections_user_id_provider_provider_account_id_key"),
        ]


class License(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    license_key = models.TextField(unique=True)
    idempotency_key = models.TextField(null=True, unique=True)
    version = models.IntegerField(null=True)
    type = models.TextField(choices=LicenseType.choices)
    seats = models.IntegerField(null=True)
    issued_at = UtcDateTimeField(default=now_milliseconds)
    expires_at = UtcDateTimeField(null=True)
    licensee_email = models.TextField(null=True)
    licensee_name = models.TextField(null=True)
    white_label_addon = models.BooleanField(default=False)
    status = models.TextField(choices=LicenseStatus.choices, default=LicenseStatus.ACTIVE)

    class Meta:
        db_table = "licenses"


class LicenseValidation(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    license = models.ForeignKey("License", on_delete=models.CASCADE, related_name="validations", db_index=False)
    ip_address = models.TextField(null=True)
    fingerprint = models.TextField(null=True)
    validated_at = UtcDateTimeField(default=now_milliseconds)
    user_agent = models.TextField(null=True)

    class Meta:
        db_table = "license_validations"


class InstanceLicense(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    license_key = models.TextField(unique=True)
    version = models.IntegerField(null=True)
    type = models.TextField(choices=LicenseType.choices)
    seats = models.IntegerField(null=True)
    issued_at = UtcDateTimeField(default=now_milliseconds)
    expires_at = UtcDateTimeField(null=True)
    licensee_email = models.TextField(null=True)
    licensee_name = models.TextField(null=True)
    white_label_addon = models.BooleanField(default=False)
    status = models.TextField(choices=LicenseStatus.choices, default=LicenseStatus.ACTIVE)

    class Meta:
        db_table = "instance_licenses"


class UserNotificationPreferences(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    user = models.OneToOneField("User", on_delete=models.CASCADE, related_name="notification_preferences")
    prefs = models.JSONField(default=dict)
    created_at = UtcDateTimeField(default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "user_notification_preferences"


class Poll(models.Model):
    id = models.TextField(primary_key=True, default=None)
    created_at = UtcDateTimeField(default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(auto_now=True)
    deadline = UtcDateTimeField(null=True)
    title = models.TextField()
    description = models.TextField(null=True)
    location = models.TextField(null=True)
    conferencing = models.JSONField(null=True)
    """What the organizer asked for: a provider to mint a link with at
    finalize ("zoom" | "meet" | "teams" | "webex"), or a link they pasted
    themselves.
    """
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="polls", null=True, db_index=False)
    time_zone = models.TextField(null=True)
    status = models.TextField(choices=PollStatus.choices, default=PollStatus.open)
    closed_reason = models.TextField(choices=PollClosedReason.choices, null=True)
    kind = models.TextField(choices=PollKind.choices, default=PollKind.date)
    deleted = models.BooleanField(default=False)
    deleted_at = UtcDateTimeField(null=True)
    event_id = models.TextField(null=True, unique=True)
    scheduled_event = models.ForeignKey("ScheduledEvent", on_delete=models.SET_NULL, related_name="polls", null=True, db_index=False)
    hide_participants = models.BooleanField(default=False)
    hide_scores = models.BooleanField(default=False)
    disable_comments = models.BooleanField(default=True)
    require_participant_email = models.BooleanField(default=False)
    allow_tentative_votes = models.BooleanField(default=True)
    """Whether participants may cast the tentative "if need be" vote. Positive
    flag so a future vote-type expansion adds columns rather than inverting a
    negative. Turning it off only blocks new tentative votes; existing ones
    are kept and stay visible.
    """
    muted = models.BooleanField(default=False)
    space = models.ForeignKey("Space", on_delete=models.SET_NULL, related_name="polls", null=True, db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "polls"
        indexes = [
            models.Index(fields=["space"]),
            models.Index(fields=["user"]),
        ]


class Participant(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    name = models.TextField()
    email = models.TextField(null=True)
    note = models.TextField(null=True)
    token = models.TextField(unique=True)
    """Per response edit link: grants edit access to this response without a session. App minted at creation, CSPRNG alphanumeric ≥32 chars (nanoid customAlphabet — url safe with no linkifier edge cases, unlike base64url); never a DB default so a code path that forgets to mint fails loudly. A response created through an emailed invite takes the invite's token. Rows that predate minting were backfilled with gen_random_uuid() hex"""
    user = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="participants", null=True, db_index=False)
    poll = models.ForeignKey("Poll", on_delete=models.CASCADE, related_name="participants", db_index=False)
    locale = models.TextField(null=True)
    time_zone = models.TextField(null=True)
    created_at = UtcDateTimeField(default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(null=True, auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "participants"
        indexes = [
            HashIndex(fields=["poll"]),
            models.Index(fields=["user"]),
        ]


class PollInvite(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    poll = models.ForeignKey("Poll", on_delete=models.CASCADE, related_name="invites", db_index=False)
    email = models.TextField()
    token = models.TextField(unique=True)
    """Per invitee link. App minted, CSPRNG alphanumeric ≥32 chars (nanoid customAlphabet) — no schema default, so forgetting to mint fails loudly. Rotated on reactivation so a revoked (possibly leaked) token never regains access"""
    opened_at = UtcDateTimeField(null=True)
    reminded_at = UtcDateTimeField(null=True)
    revoked_at = UtcDateTimeField(null=True)
    """Revocation is always explicit: set when the host removes a pending invite, never as a side effect of close, schedule or reopen. Deleting a response instead hard deletes the participant, so the SetNull FK clears participantId and reverts the invite to pending with its token intact — the token the deleted response held is freed with it, so the next response through the same link can take it again. Cleared only by re-inviting (reactivation), which also rotates the token"""
    participant = models.OneToOneField("Participant", on_delete=models.SET_NULL, related_name="invite", null=True)
    """Conversion join, merges invite history into the response timeline. Set only when the invitee responds via their token: a typed email is unverified, so a generic link response never claims an invite"""
    created_at = UtcDateTimeField(default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "poll_invites"
        constraints = [
            models.UniqueConstraint(fields=["poll", "email"], name="poll_invites_poll_id_email_key"),
        ]


class PollActivity(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    poll = models.ForeignKey("Poll", on_delete=models.CASCADE, related_name="activities", db_index=False)
    type = models.TextField()
    """Validated by a zod union in the feature, not a DB enum: the activity vocabulary grows without migrations"""
    user_id = models.TextField(null=True)
    """Subject ids are soft references (no FKs): activity must outlive its subjects — including account deletion — so refs persist as identity and the payload carries the display snapshot; whether a ref still resolves is the renderer's question"""
    participant_id = models.TextField(null=True)
    invite_id = models.TextField(null=True)
    option_id = models.TextField(null=True)
    payload = models.JSONField(default=dict)
    created_at = UtcDateTimeField(default=now_milliseconds)

    class Meta:
        db_table = "poll_activities"
        indexes = [
            models.Index(fields=["poll", "-created_at"]),
            models.Index(fields=["poll", "participant_id"]),
            models.Index(fields=["poll", "option_id"]),
            models.Index(fields=["poll", "invite_id"]),
            models.Index(fields=["user_id"]),
        ]


class Option(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    start_time = UtcDateTimeField()
    duration = models.IntegerField(db_column="duration_minutes", default=0)
    poll = models.ForeignKey("Poll", on_delete=models.CASCADE, related_name="options", db_index=False)
    created_at = UtcDateTimeField(default=now_milliseconds)

    class Meta:
        db_table = "options"
        indexes = [
            HashIndex(fields=["poll"]),
        ]


class Vote(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    participant = models.ForeignKey("Participant", on_delete=models.CASCADE, related_name="votes", db_index=False)
    option = models.ForeignKey("Option", on_delete=models.CASCADE, related_name="votes", db_index=False)
    poll = models.ForeignKey("Poll", on_delete=models.CASCADE, related_name="votes", db_index=False)
    type = models.TextField(choices=VoteType.choices, default=VoteType.yes)
    created_at = UtcDateTimeField(default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(null=True, auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "votes"
        indexes = [
            HashIndex(fields=["poll"]),
            HashIndex(fields=["participant"]),
            HashIndex(fields=["option"]),
        ]


class Comment(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    content = models.TextField()
    poll = models.ForeignKey("Poll", on_delete=models.CASCADE, related_name="comments", db_index=False)
    author_name = models.TextField()
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="comments", null=True, db_index=False)
    created_at = UtcDateTimeField(default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(null=True, auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "comments"
        indexes = [
            models.Index(fields=["poll"]),
            models.Index(fields=["user"]),
        ]


class RegisteredInstance(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    instance_id = models.TextField(unique=True)
    version = models.TextField()
    first_seen_at = UtcDateTimeField(default=now_milliseconds)
    last_seen_at = UtcAutoNowDateTimeField(auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "registered_instances"


class Sheet(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    space = models.ForeignKey("Space", on_delete=models.CASCADE, related_name="sheets", db_index=False)
    host = models.ForeignKey("User", on_delete=models.CASCADE, related_name="hosted_sheets", db_index=False)
    title = models.TextField()
    description = models.TextField(null=True)
    url_id = models.TextField(unique=True)
    created_at = UtcDateTimeField(default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(auto_now=True)
    deleted = models.BooleanField(default=False)
    deleted_at = UtcDateTimeField(null=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "sheets"
        indexes = [
            models.Index(fields=["space"]),
            models.Index(fields=["host"]),
        ]


class SheetSlot(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    sheet = models.ForeignKey("Sheet", on_delete=models.CASCADE, related_name="slots", db_index=False)
    event_type = models.ForeignKey("EventType", on_delete=models.RESTRICT, related_name="slots", db_index=False)
    start_time = UtcDateTimeField()
    created_at = UtcDateTimeField(default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "sheet_slots"
        indexes = [
            models.Index(fields=["sheet", "start_time"]),
            models.Index(fields=["event_type"]),
        ]


class SpaceUsage(models.Model):
    """Monthly counters for metered actions charged to a space, independent of
    the records they produce, so deleting a poll does not give back what it
    spent.
    """

    pk = models.CompositePrimaryKey("space_id", "resource", "period")
    space = models.ForeignKey("Space", on_delete=models.CASCADE, related_name="usage", db_index=False)
    resource = models.TextField(choices=SpaceUsageResource.choices)
    period = models.DateField()
    count = models.IntegerField(default=0)
    near_limit_warned_at = UtcDateTimeField(null=True)
    at_limit_warned_at = UtcDateTimeField(null=True)
    updated_at = UtcAutoNowDateTimeField(auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "space_usage"


class Space(models.Model):
    id = models.TextField(primary_key=True, default=uuid4_str)
    name = models.TextField()
    image = models.TextField(null=True)
    owner = models.ForeignKey("User", on_delete=models.CASCADE, related_name="spaces", db_index=False)
    tier = models.TextField(choices=SpaceTier.choices, default=SpaceTier.hobby)
    space_type = models.TextField(choices=SpaceType.choices, null=True)
    industry = models.TextField(null=True)
    primary_color = models.TextField(null=True)
    show_branding = models.BooleanField(default=False)
    hide_attribution = models.BooleanField(default=False)
    shared = models.BooleanField(default=False)
    nonprofit_discount_granted_at = UtcDateTimeField(null=True)
    api_poll_quota_override = models.IntegerField(null=True)
    created_at = UtcDateTimeField(default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "spaces"
        indexes = [
            models.Index(fields=["owner"]),
        ]


class SpaceMember(models.Model):
    id = models.TextField(primary_key=True, default=uuid4_str)
    space = models.ForeignKey("Space", on_delete=models.CASCADE, related_name="members", db_index=False)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="member_of", db_index=False)
    created_at = UtcDateTimeField(default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(auto_now=True)
    role = models.TextField(choices=SpaceMemberRole.choices, default=SpaceMemberRole.MEMBER)
    last_selected_at = UtcDateTimeField(default=now_milliseconds)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "space_members"
        constraints = [
            models.UniqueConstraint(fields=["space", "user"], name="space_members_space_id_user_id_key"),
        ]
        indexes = [
            models.Index(fields=["user"]),
        ]


class SpaceMemberInvite(models.Model):
    id = models.TextField(primary_key=True, default=uuid4_str)
    space = models.ForeignKey("Space", on_delete=models.CASCADE, related_name="member_invites", db_index=False)
    email = models.TextField()
    created_at = UtcDateTimeField(default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(auto_now=True)
    role = models.TextField(choices=SpaceMemberRole.choices, default=SpaceMemberRole.MEMBER)
    invited_by = models.ForeignKey("User", on_delete=models.CASCADE, related_name="space_member_invites", db_column="inviter_id", db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "space_member_invites"
        constraints = [
            models.UniqueConstraint(fields=["space", "email"], name="space_member_invites_space_id_email_key"),
        ]


class SpaceApiKey(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    space = models.ForeignKey("Space", on_delete=models.CASCADE, related_name="api_keys", db_index=False)
    name = models.TextField()
    prefix = models.TextField(unique=True)
    hashed_key = models.TextField()
    last_used_at = UtcDateTimeField(null=True)
    expires_at = UtcDateTimeField(null=True)
    revoked_at = UtcDateTimeField(null=True)
    created_at = UtcDateTimeField(default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "space_api_keys"
        indexes = [
            models.Index(fields=["space"], name="space_api_key_space_idx"),
            models.Index(fields=["expires_at"], name="space_api_key_expires_idx"),
            models.Index(fields=["revoked_at"], name="space_api_key_revoked_idx"),
        ]


class SpaceWebhook(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    space = models.ForeignKey("Space", on_delete=models.CASCADE, related_name="webhooks", db_index=False)
    url = models.TextField()
    secret = models.TextField()
    """AES-256-GCM under SECRET_PASSWORD (see @rallly/utils/encryption); the plaintext is shown once at creation"""
    events = ArrayField(models.TextField())
    """Subscribed event names (`poll.closed`, ...); validated by the feature schema, not a DB enum, so the vocabulary grows without migrations"""
    enabled = models.BooleanField(default=True)
    version = models.TextField()
    """The payload contract this endpoint was built against (WEBHOOK_VERSION at creation). No default: a writer that forgets it fails instead of recording a stale contract. A breaking payload change branches on this so existing endpoints keep their shape"""
    cursor = UtcDateTimeField(default=now_milliseconds)
    """Fan-out position in the poll activity log: activities created after this instant are still to be delivered. Starts at creation so a new endpoint never replays history"""
    last_delivered_at = UtcDateTimeField(null=True)
    last_failed_at = UtcDateTimeField(null=True)
    consecutive_failures = models.IntegerField(default=0)
    """Exhausted deliveries in a row; reset by any success. The dispatcher disables the endpoint once it reaches its cap"""
    created_at = UtcDateTimeField(default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "space_webhooks"
        indexes = [
            models.Index(fields=["space"]),
        ]


class WebhookDelivery(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    webhook = models.ForeignKey("SpaceWebhook", on_delete=models.CASCADE, related_name="deliveries", db_index=False)
    activity_id = models.TextField()
    """Soft reference to the PollActivity row this delivery was fanned out from; with webhookId it makes fan-out idempotent"""
    event_type = models.TextField()
    payload = models.JSONField()
    """Frozen at fan-out: retries resend exactly what the first attempt sent"""
    status = models.TextField(choices=WebhookDeliveryStatus.choices, default=WebhookDeliveryStatus.pending)
    attempts = models.IntegerField(default=0)
    next_attempt_at = UtcDateTimeField(default=now_milliseconds)
    last_response_status = models.IntegerField(null=True)
    last_error = models.TextField(null=True)
    created_at = UtcDateTimeField(default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "webhook_deliveries"
        constraints = [
            models.UniqueConstraint(fields=["webhook", "activity_id"], name="webhook_deliveries_webhook_id_activity_id_key"),
        ]
        indexes = [
            models.Index(fields=["status", "next_attempt_at"]),
        ]


class Account(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="accounts", db_index=False)
    provider = models.TextField()
    provider_account_id = models.TextField()
    refresh_token = models.TextField(null=True)
    access_token = models.TextField(null=True)
    expires_at = models.IntegerField(null=True)
    scope = models.TextField(null=True)
    id_token = models.TextField(null=True)
    access_token_expires_at = UtcDateTimeField(null=True)
    refresh_token_expires_at = UtcDateTimeField(null=True)
    password = models.TextField(null=True)
    created_at = UtcDateTimeField(default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "accounts"
        constraints = [
            models.UniqueConstraint(fields=["provider", "provider_account_id"], name="accounts_provider_provider_account_id_key"),
        ]
        indexes = [
            models.Index(fields=["user"]),
        ]


class User(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    name = models.TextField()
    email = models.TextField(unique=True)
    email_verified = models.BooleanField(null=True)
    image = models.TextField(null=True)
    time_zone = models.TextField(null=True)
    week_start = models.IntegerField(null=True)
    time_format = models.TextField(choices=TimeFormat.choices, null=True)
    locale = models.TextField(null=True)
    created_at = UtcDateTimeField(default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(null=True, auto_now=True)
    customer_id = models.TextField(null=True)
    banned = models.BooleanField(default=False)
    banned_at = UtcDateTimeField(null=True)
    ban_reason = models.TextField(null=True)
    ban_expires = UtcDateTimeField(null=True)
    role = models.TextField(choices=UserRole.choices, default=UserRole.user)
    last_login_method = models.TextField(null=True)
    is_anonymous = models.BooleanField(db_column="anonymous", default=False)
    last_seen_at = UtcDateTimeField(default=now_milliseconds)
    deleted_at = UtcDateTimeField(null=True)
    job_title = models.TextField(null=True)
    review_requested_at = UtcDateTimeField(null=True)
    default_destination_calendar = models.ForeignKey("ProviderCalendar", on_delete=models.SET_NULL, related_name="user_default_destination", null=True, db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "users"
        indexes = [
            models.Index(fields=["deleted_at"]),
            models.Index(fields=["is_anonymous", "last_seen_at"]),
        ]


class Session(models.Model):
    id = models.TextField(primary_key=True, default=None)
    expires_at = UtcDateTimeField()
    token = models.TextField(unique=True)
    created_at = UtcDateTimeField(default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(auto_now=True)
    ip_address = models.TextField(null=True)
    user_agent = models.TextField(null=True)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="sessions", db_index=False)
    impersonated_by_user = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="impersonated_sessions", db_column="impersonated_by", null=True, db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "sessions"
        constraints = [
            models.UniqueConstraint(fields=["token"], name="sessions_token_key"),
        ]


class Verification(models.Model):
    id = models.TextField(primary_key=True, default=None)
    identifier = models.TextField(unique=True)
    value = models.TextField()
    expires_at = UtcDateTimeField()
    created_at = UtcDateTimeField(default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "verifications"
        indexes = [
            models.Index(fields=["identifier"]),
        ]
