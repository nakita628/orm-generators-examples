from datetime import datetime
from datetime import timezone as dt_timezone
from decimal import Decimal
from typing import Any, TypeVar

import cuid  # type: ignore[import-untyped, unused-ignore]
import nanoid
from django.conf import settings
from django.contrib.postgres.fields import ArrayField
from django.contrib.postgres.indexes import GinIndex
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


class AutoNowDateTimeField(models.DateTimeField):  # type: ignore[type-arg]
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


def cuid_str() -> str:
    return str(cuid.cuid())


def nanoid_str(size: int = 21) -> str:
    return nanoid.generate(size=size)


def now_milliseconds() -> datetime:
    now = timezone.now()
    return now.replace(microsecond=now.microsecond // 1000 * 1000)


class SsoConnectionMigrationPhase(models.TextChoices):
    SETUP = "SETUP"
    GRACE_LEGACY = "GRACE_LEGACY"
    GRACE_DIRECT = "GRACE_DIRECT"
    FINALIZING = "FINALIZING"
    FINALIZED = "FINALIZED"


class TeamUserRole(models.TextChoices):
    ADMIN = "ADMIN"
    MEMBER = "MEMBER"
    VIEWER = "VIEWER"
    CUSTOM = "CUSTOM"


class RoleBindingScopeType(models.TextChoices):
    ORGANIZATION = "ORGANIZATION"
    TEAM = "TEAM"
    PROJECT = "PROJECT"


class GrantScopeType(models.TextChoices):
    ORGANIZATION = "ORGANIZATION"
    TEAM = "TEAM"
    PROJECT = "PROJECT"
    RESOURCE = "RESOURCE"
    PLATFORM = "PLATFORM"


class GrantPrincipalType(models.TextChoices):
    USER = "USER"
    API_KEY = "API_KEY"
    GROUP = "GROUP"
    TEAM = "TEAM"
    ORGANIZATION = "ORGANIZATION"
    PROJECT = "PROJECT"
    ANYONE = "ANYONE"


class OrganizationUserRole(models.TextChoices):
    ADMIN = "ADMIN"
    MEMBER = "MEMBER"
    EXTERNAL = "EXTERNAL"
    DEVELOPER = "DEVELOPER"


class IssuedLicenseSource(models.TextChoices):
    """Where a license in the registry came from."""

    BACKOFFICE = "BACKOFFICE"
    PURCHASE = "PURCHASE"
    SCRIPT = "SCRIPT"
    LEGACY_IMPORT = "LEGACY_IMPORT"


class ConnectedCreditGrantKind(models.TextChoices):
    """Why a credit was granted. The amount is always prepaid, never promotional."""

    commit = "commit"
    added = "added"
    renewal = "renewal"


class ConnectedInvoiceKind(models.TextChoices):
    """What an invoice of a connected customer is for."""

    annual = "annual"
    seat_change = "seat_change"
    usage = "usage"


class PlanTypes(models.TextChoices):
    FREE = "FREE"
    PRO = "PRO"
    GROWTH = "GROWTH"
    GROWTH_SEAT_EUR_MONTHLY = "GROWTH_SEAT_EUR_MONTHLY"
    GROWTH_SEAT_EUR_ANNUAL = "GROWTH_SEAT_EUR_ANNUAL"
    GROWTH_SEAT_USD_MONTHLY = "GROWTH_SEAT_USD_MONTHLY"
    GROWTH_SEAT_USD_ANNUAL = "GROWTH_SEAT_USD_ANNUAL"
    ENTERPRISE = "ENTERPRISE"
    LAUNCH = "LAUNCH"
    ACCELERATE = "ACCELERATE"
    LAUNCH_ANNUAL = "LAUNCH_ANNUAL"
    ACCELERATE_ANNUAL = "ACCELERATE_ANNUAL"


class PricingModel(models.TextChoices):
    TIERED = "TIERED"
    SEAT_EVENT = "SEAT_EVENT"


class OrganizationIntent(models.TextChoices):
    """Organization signup intent (ADR-038): which product the org came for."""

    AGENT_GOVERNANCE = "AGENT_GOVERNANCE"
    LLM_OPS = "LLM_OPS"


class SubscriptionStatus(models.TextChoices):
    PENDING = "PENDING"
    FAILED = "FAILED"
    ACTIVE = "ACTIVE"
    CANCELLED = "CANCELLED"


class Currency(models.TextChoices):
    USD = "USD"
    EUR = "EUR"


class INVITESTATUS(models.TextChoices):
    PENDING = "PENDING"
    ACCEPTED = "ACCEPTED"
    WAITING_APPROVAL = "WAITING_APPROVAL"
    PAYMENT_PENDING = "PAYMENT_PENDING"
    REVOKED = "REVOKED"


class EvaluationExecutionMode(models.TextChoices):
    ON_MESSAGE = "ON_MESSAGE"
    AS_GUARDRAIL = "AS_GUARDRAIL"
    MANUALLY = "MANUALLY"


class CostType(models.TextChoices):
    TRACE_CHECK = "TRACE_CHECK"
    GUARDRAIL = "GUARDRAIL"
    CLUSTERING = "CLUSTERING"
    BATCH_EVALUATION = "BATCH_EVALUATION"


class CostReferenceType(models.TextChoices):
    CHECK = "CHECK"
    TRACE = "TRACE"
    PROJECT = "PROJECT"
    BATCH = "BATCH"


class DashboardScope(models.TextChoices):
    """Who sees a dashboard, named as the prompt scope is, with one more level."""

    PRIVATE = "PRIVATE"
    PROJECT = "PROJECT"
    ORGANIZATION = "ORGANIZATION"


class TriggerAction(models.TextChoices):
    SEND_EMAIL = "SEND_EMAIL"
    ADD_TO_DATASET = "ADD_TO_DATASET"
    ADD_TO_ANNOTATION_QUEUE = "ADD_TO_ANNOTATION_QUEUE"
    SEND_SLACK_MESSAGE = "SEND_SLACK_MESSAGE"
    SEND_WEBHOOK = "SEND_WEBHOOK"


class TriggerKind(models.TextChoices):
    AUTOMATION = "AUTOMATION"
    ALERT = "ALERT"
    REPORT = "REPORT"


class AlertType(models.TextChoices):
    CRITICAL = "CRITICAL"
    WARNING = "WARNING"
    INFO = "INFO"


class SlackIntegrationScopeType(models.TextChoices):
    """Who can use a Slack connection (ADR-093 §5a): every project in the
    organization, or one project. Chosen per connection.
    """

    ORGANIZATION = "ORGANIZATION"
    PROJECT = "PROJECT"


class SlackIntegrationKind(models.TextChoices):
    """How a Slack connection posts: a bot token over the Web API (any channel it
    is invited to, every block renders) or an incoming webhook (one channel).
    """

    BOT = "BOT"
    INCOMING_WEBHOOK = "INCOMING_WEBHOOK"


class WebhookDeliveryOutcome(models.TextChoices):
    success = "success"
    retryable = "retryable"
    terminal = "terminal"
    pending = "pending"


class WebhookDestinationKind(models.TextChoices):
    """The last hop a webhook endpoint delivers over. Lowercase because these
    values are the wire spelling: the same strings appear in the REST DTO,
    both SDKs and the CLI.
    """

    http = "http"
    sqs = "sqs"


class WebhookEndpointStatus(models.TextChoices):
    ACTIVE = "ACTIVE"
    DISABLED = "DISABLED"


class WebhookDeliveryChannel(models.TextChoices):
    """Which webhook pipeline wrote a delivery row, and therefore which columns
    carry its tenancy: `platform` rows anchor on organizationId + endpointId,
    `automations` rows on projectId + triggerId.
    """

    platform = "platform"
    automations = "automations"


class ExperimentType(models.TextChoices):
    DSPY = "DSPY"
    BATCH_EVALUATION = "BATCH_EVALUATION"
    BATCH_EVALUATION_V2 = "BATCH_EVALUATION_V2"
    EVALUATIONS_V3 = "EVALUATIONS_V3"


class ModelProviderScopeType(models.TextChoices):
    ORGANIZATION = "ORGANIZATION"
    TEAM = "TEAM"
    PROJECT = "PROJECT"


class ModelDefaultScopeType(models.TextChoices):
    ORGANIZATION = "ORGANIZATION"
    TEAM = "TEAM"
    PROJECT = "PROJECT"


class LangyProjectionTitleSource(models.TextChoices):
    derived = "derived"
    auto = "auto"
    user = "user"


class AnnotationScoreDataType(models.TextChoices):
    CATEGORICAL = "CATEGORICAL"
    BOOLEAN = "BOOLEAN"
    LIKERT = "LIKERT"
    OPTION = "OPTION"
    CHECKBOX = "CHECKBOX"


class ShareResourceType(models.TextChoices):
    TRACE = "TRACE"
    THREAD = "THREAD"


class ShareVisibility(models.TextChoices):
    PUBLIC = "PUBLIC"
    ORGANIZATION = "ORGANIZATION"
    PROJECT = "PROJECT"


class PinSource(models.TextChoices):
    manual = "manual"
    share = "share"


class RetentionPolicyScopeType(models.TextChoices):
    ORGANIZATION = "ORGANIZATION"
    TEAM = "TEAM"
    PROJECT = "PROJECT"


class DataPrivacyScopeType(models.TextChoices):
    ORGANIZATION = "ORGANIZATION"
    DEPARTMENT = "DEPARTMENT"
    TEAM = "TEAM"
    PROJECT = "PROJECT"


class CustomLLMModelCostScopeType(models.TextChoices):
    ORGANIZATION = "ORGANIZATION"
    TEAM = "TEAM"
    PROJECT = "PROJECT"


class PromptScope(models.TextChoices):
    ORGANIZATION = "ORGANIZATION"
    PROJECT = "PROJECT"


class VirtualKeyStatus(models.TextChoices):
    ACTIVE = "ACTIVE"
    DISABLED = "DISABLED"
    REVOKED = "REVOKED"


class VirtualKeyPurpose(models.TextChoices):
    USER = "USER"
    LANGY = "LANGY"
    CONNECT = "CONNECT"


class GatewayProviderRotationPolicy(models.TextChoices):
    MANUAL = "MANUAL"


class GatewayProviderHealthStatus(models.TextChoices):
    UNKNOWN = "UNKNOWN"
    HEALTHY = "HEALTHY"
    DEGRADED = "DEGRADED"
    CIRCUIT_OPEN = "CIRCUIT_OPEN"


class GatewayBudgetScopeType(models.TextChoices):
    ORGANIZATION = "ORGANIZATION"
    TEAM = "TEAM"
    PROJECT = "PROJECT"
    VIRTUAL_KEY = "VIRTUAL_KEY"
    PRINCIPAL = "PRINCIPAL"
    GROUP = "GROUP"
    ATTRIBUTED_USER = "ATTRIBUTED_USER"


class GatewayBudgetWindow(models.TextChoices):
    MINUTE = "MINUTE"
    HOUR = "HOUR"
    DAY = "DAY"
    WEEK = "WEEK"
    MONTH = "MONTH"
    TOTAL = "TOTAL"
    MANUAL = "MANUAL"


class GatewayBudgetBreachAction(models.TextChoices):
    BLOCK = "BLOCK"
    WARN = "WARN"


class GatewayBudgetLedgerStatus(models.TextChoices):
    SUCCESS = "SUCCESS"
    PROVIDER_ERROR = "PROVIDER_ERROR"
    BLOCKED_BY_GUARDRAIL = "BLOCKED_BY_GUARDRAIL"
    CANCELLED = "CANCELLED"


class GatewayChangeEventKind(models.TextChoices):
    VK_CREATED = "VK_CREATED"
    VK_CONFIG_UPDATED = "VK_CONFIG_UPDATED"
    VK_REVOKED = "VK_REVOKED"
    VK_ROTATED = "VK_ROTATED"
    VK_DISABLED = "VK_DISABLED"
    VK_ENABLED = "VK_ENABLED"
    BUDGET_CREATED = "BUDGET_CREATED"
    BUDGET_UPDATED = "BUDGET_UPDATED"
    BUDGET_DELETED = "BUDGET_DELETED"
    MODEL_PROVIDER_UPDATED = "MODEL_PROVIDER_UPDATED"
    CACHE_RULE_CREATED = "CACHE_RULE_CREATED"
    CACHE_RULE_UPDATED = "CACHE_RULE_UPDATED"
    CACHE_RULE_DELETED = "CACHE_RULE_DELETED"
    ROUTING_POLICY_UPDATED = "ROUTING_POLICY_UPDATED"
    ROUTING_POLICY_DELETED = "ROUTING_POLICY_DELETED"


class GatewayCacheRuleMode(models.TextChoices):
    RESPECT = "RESPECT"
    FORCE = "FORCE"
    DISABLE = "DISABLE"


class VirtualKeyRoutingMode(models.TextChoices):
    NONE = "NONE"
    FALLBACK_ALL = "FALLBACK_ALL"
    POLICY = "POLICY"


class VirtualKeyScopeType(models.TextChoices):
    ORGANIZATION = "ORGANIZATION"
    TEAM = "TEAM"
    PROJECT = "PROJECT"


class RoutingPolicyScopeType(models.TextChoices):
    ORGANIZATION = "ORGANIZATION"
    TEAM = "TEAM"
    PROJECT = "PROJECT"


class GatewayGuardrailDirection(models.TextChoices):
    PRE = "PRE"
    POST = "POST"
    STREAM_CHUNK = "STREAM_CHUNK"


class GatewayGuardrailFailureMode(models.TextChoices):
    FAIL_OPEN = "FAIL_OPEN"
    FAIL_CLOSED = "FAIL_CLOSED"


class FeatureFlagExperimentSubjectType(models.TextChoices):
    USER = "USER"
    ORGANIZATION = "ORGANIZATION"
    PROJECT = "PROJECT"


class ProcessManagerOutboxStatus(models.TextChoices):
    pending = "pending"
    dispatched = "dispatched"
    dead = "dead"
    discarded = "discarded"


class ProcessManagerOutboxAttemptOutcome(models.TextChoices):
    retry_scheduled = "retry_scheduled"
    dead = "dead"


class GatewayRealtimeSessionStatus(models.TextChoices):
    OPEN = "OPEN"
    CLOSED = "CLOSED"
    FAILED = "FAILED"
    EXPIRED = "EXPIRED"


class ConnectedSeatChangeState(models.TextChoices):
    """How the invoicing of one seat change ended."""

    intent = "intent"
    invoiced = "invoiced"
    nothing_to_invoice = "nothing_to_invoice"
    not_onboarded = "not_onboarded"


class Account(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="accounts", db_column="userId", db_index=False)
    type = models.TextField(default="oauth")
    provider = models.TextField()
    issuer = models.TextField(null=True)
    """WHO asserted `providerAccountId`. Better Auth keys accounts by
    `(provider, providerAccountId)`; auth's account create hook writes this
    (modules/auth/specs/account-issuer.feature). Nullable only so the column
    could be added to a populated table; the migration backfilled every row.
    """
    provider_account_id = models.TextField(db_column="providerAccountId")
    refresh_token = models.TextField(null=True)
    access_token = models.TextField(null=True)
    expires_at = UtcDateTimeField(null=True)
    ext_expires_in = models.IntegerField(null=True)
    token_type = models.TextField(null=True)
    scope = models.TextField(null=True)
    id_token = models.TextField(null=True)
    session_state = models.TextField(null=True)
    password = models.TextField(null=True)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Account"
        constraints = [
            models.UniqueConstraint(fields=["provider", "provider_account_id"], name="Account_provider_providerAccountId_key"),
        ]
        indexes = [
            models.Index(fields=["user"]),
            models.Index(fields=["issuer", "provider_account_id"]),
        ]


class AccountCredential(models.Model):
    id = models.TextField(primary_key=True, default=None)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="account_credentials", db_column="userId", db_index=False)
    provider = models.TextField()
    password = models.TextField(null=True)
    access_token = models.TextField(null=True, db_column="accessToken")
    refresh_token = models.TextField(null=True, db_column="refreshToken")
    id_token = models.TextField(null=True, db_column="idToken")
    access_token_expires_at = UtcDateTimeField(null=True, db_column="accessTokenExpiresAt")
    refresh_token_expires_at = UtcDateTimeField(null=True, db_column="refreshTokenExpiresAt")
    scope = models.TextField(null=True)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "AccountCredential"
        indexes = [
            models.Index(fields=["user"]),
        ]


class Session(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    session_token = models.TextField(unique=True, db_column="sessionToken")
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="sessions", db_column="userId", db_index=False)
    expires = UtcDateTimeField()
    ip_address = models.TextField(null=True, db_column="ipAddress")
    user_agent = models.TextField(null=True, db_column="userAgent")
    identifier_id = models.TextField(null=True, db_column="identifierId")
    """Which of the user's sign-in methods minted this session (D06). NULL on
    every session that predates the column and on any mint we could not
    attribute - an ordinary value, not an error. Nothing reads it to decide
    whether the session is valid, so no session was ended to add it.
    """
    amr = ArrayField(models.TextField(), default=list)
    """What the sign-in that minted this session PROVED, as RFC 8176 method
    references (D06). Empty on every session minted before the column, and
    on any ordinary password sign-in. An organization that requires a second
    factor reads this when a member reaches its data; it is never read at
    mint and never ends a session.
    """
    started_by_sign_up = models.BooleanField(db_column="startedBySignUp", default=False)
    """Minted while the account had never signed in: the session its sign-up
    started, which skips the secure-account offer. False on every later one
    and on every session that predates the column (specs/identity/passkeys.feature).
    """
    impersonating = models.JSONField(null=True)
    """The legacy impersonation payload. NOTHING READS OR WRITES IT any more -
    the {actor, subject} columns below replaced it, and the migration beside
    them ended every session that still carried a value.

    It stays for ONE release, and that is the whole reason it is still here.
    `prisma migrate deploy` runs at container start, so the first pod of a
    new release drops a column while the previous release's pods are still
    serving - and those pods select this column on EVERY authenticated
    request and fail closed when the row does not come back. Dropping it in
    the same release that stops reading it signs everybody out for the length
    of the rollout. Expand now, contract next release.
    """
    actor_user_id = models.TextField(null=True, db_column="actorUserId")
    """The impersonation claims (D06), aligned with the authz Principal
    {actor, subject}. The actor is the operator who really made the request
    and the subject is the person whose access they are borrowing; both are
    recorded on every authorization decision, so the audit trail can answer
    who really did it. These replace the legacy `impersonating` JSON blob,
    which is dropped in the release AFTER this one.

    NULL on every ordinary session, which is all of them: an impersonation
    is a handful of rows at a time and nothing else reads these columns.
    """
    subject_user_id = models.TextField(null=True, db_column="subjectUserId")
    impersonation_reason = models.TextField(null=True, db_column="impersonationReason")
    """Why the operator started it. Required to start one
    (specs/features/backoffice-user-impersonation-reason.feature) and kept
    beside both people rather than only in the audit log, so a request under
    impersonation carries its own justification.
    """
    impersonation_expires_at = UtcDateTimeField(null=True, db_column="impersonationExpiresAt")
    """When the borrowed access lapses. A session outlives its impersonation
    window: past this the claims are ignored and the operator is themselves
    again, with nothing ended.
    """
    last_seen_at = UtcDateTimeField(null=True, db_column="lastSeenAt")
    """When this session was last used, for the idle timeout an organization
    may set (GAC-10, specs/identity/org-session-lifetime.feature).

    A COLUMN OF OUR OWN rather than `updatedAt`, because better-auth rolls a
    live session's expiry once per `updateAge` - a day - so `updatedAt` is
    activity to the nearest day and cannot tell an idle hour from an idle
    minute. Written only for sessions whose owner belongs to an organization
    that has set a window, and then at most once per quarter-window, so a
    deployment where nobody has turned this on writes it never.

    NULL on every session minted before the column and on every session
    under no window at all. A null READS AS `updatedAt`, which better-auth
    keeps rolled to within a day of the last use. That fallback is
    conservative in the direction that matters - `updatedAt` is never
    earlier than the real last use, so it can only ever judge a session MORE
    active than it was, never less - and it is what lets a window turned on
    today end the sessions that have been idle since February instead of
    waiting a day for a stamp they have never had.
    """
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Session"
        indexes = [
            models.Index(fields=["user"]),
            models.Index(fields=["identifier_id"]),
        ]


class UserErasure(models.Model):
    """One row per erased account, holding only its creation time: no id, name or address. It keeps
    the install's first-created account knowable after an erasure deletes that row.
    Spec: modules/user/specs/user.feature
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    account_created_at = UtcDateTimeField(db_column="accountCreatedAt")
    erased_at = UtcDateTimeField(db_column="erasedAt", default=now_milliseconds)

    class Meta:
        db_table = "UserErasure"
        indexes = [
            models.Index(fields=["account_created_at"]),
        ]


class User(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    name = models.TextField(null=True)
    email = models.TextField(null=True, unique=True)
    email_verified = models.BooleanField(db_column="emailVerified", default=False)
    signup_confirmation_pending = models.BooleanField(db_column="signupConfirmationPending", default=False)
    passkey_signup_claim_hash = models.TextField(null=True, db_column="passkeySignupClaimHash")
    image = models.TextField(null=True)
    pending_sso_setup = models.BooleanField(db_column="pendingSsoSetup", default=False)
    user_hash_key = models.TextField(null=True, db_column="userHashKey")
    two_factor_enabled = models.BooleanField(db_column="twoFactorEnabled", default=False)
    """Owned by the two-factor plugin (D06), which is why it is a plain
    boolean rather than a state: the plugin flips it when a setup is
    confirmed and when it is turned off, and challenges every sign-in while
    it is true. The lifecycle - and the history - lives on MfaEnrollment.
    """
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    last_login_at = UtcDateTimeField(null=True, db_column="lastLoginAt")
    deactivated_at = UtcDateTimeField(null=True, db_column="deactivatedAt")
    last_home_path = models.TextField(null=True, db_column="lastHomePath")
    traces_explorer_tour_dismissed_at = UtcDateTimeField(null=True, db_column="tracesExplorerTourDismissedAt")
    langy_code_access_preference = models.TextField(null=True, db_column="langyCodeAccessPreference")
    """How Langy should reach this person's code when a task needs a change
    (ADR-129). "github" means they chose the pull request path and asked to
    be remembered; NULL means ask. A shared local folder is never stored
    here, because a folder has to be shared again each time.
    Spec: specs/langy/langy-code-access.feature
    """
    passkey_nudge_dismissed_at = UtcDateTimeField(null=True, db_column="passkeyNudgeDismissedAt")
    """When this person last said "not now" to being offered a passkey
    (ADR-120). On the account rather than in browser storage for two
    reasons: a per-browser record forgets on a new device, which is exactly
    where the offer should be MORE eager, and an interval nobody can enforce
    is not an interval. NULL means never asked.
    """
    join_offer_dismissed_domains = ArrayField(models.TextField(), db_column="joinOfferDismissedDomains", default=list)
    """The domains this person has said "no thanks" to being offered their
    colleagues on (D12). Per domain rather than one flag: somebody who
    changes jobs and verifies a second address should be offered their new
    employer's organization, not silenced forever by a decision they made
    about the old one. Empty means never dismissed anything.
    """
    notification_preferences = models.JSONField(db_column="notificationPreferences", default=dict)
    """This person's answer about browser notifications, one entry per topic:
    { "langy": "enabled" | "declined" }. A topic that is absent was never
    answered. The browser's own permission is not stored here.
    Spec: specs/langy/langy-notifications.feature
    """
    password_reset_required_at = UtcDateTimeField(null=True, db_column="passwordResetRequiredAt")
    """When an operator marked this person's password as no longer trusted. A
    password sign-in opens no session while it is set; any write of a new
    password hash clears it. Spec: modules/auth/specs/forced-password-reset.feature
    """
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "User"


class VerificationToken(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    identifier = models.TextField()
    token = models.TextField()
    """Not unique alone: better-auth stores plain values here (a user id, a counter). See WEB-19201."""
    expires = UtcDateTimeField()
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "VerificationToken"
        constraints = [
            models.UniqueConstraint(fields=["identifier", "token"], name="VerificationToken_identifier_token_key"),
        ]


class Identifier(models.Model):
    id = models.TextField(primary_key=True, default=None)
    user_id = models.TextField(db_column="userId")
    provider = models.TextField()
    value = models.TextField(null=True)
    domain = models.TextField(null=True)
    identifier_hash = models.TextField(null=True, db_column="identifierHash")
    account_id = models.TextField(null=True, db_column="accountId")
    provider_id = models.TextField(null=True, db_column="providerId")
    issuer = models.TextField(null=True)
    provider_account_id = models.TextField(null=True, db_column="providerAccountId")
    state = models.TextField()
    connection_id = models.TextField(null=True, db_column="connectionId")
    verified_at = UtcDateTimeField(null=True, db_column="verifiedAt")
    attached_at = UtcDateTimeField(db_column="attachedAt")
    detached_at = UtcDateTimeField(null=True, db_column="detachedAt")
    last_used_at = UtcDateTimeField(null=True, db_column="lastUsedAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Identifier"
        indexes = [
            models.Index(fields=["user_id"]),
            models.Index(fields=["identifier_hash"]),
            models.Index(fields=["value"]),
            models.Index(fields=["domain"]),
            models.Index(fields=["provider", "provider_account_id"]),
            models.Index(fields=["issuer", "provider_account_id"]),
            models.Index(fields=["account_id"]),
        ]


class IdentitySignInSwitch(models.Model):
    """Whether a user signs in on the identifier path: identity's own switch, set when the
    identifier backfill proves the user and pinned off by an operator's rollback. No row is the
    legacy path. Spec: modules/identity/specs/sign-in-switch.feature. Column refs only, no FK.
    """

    user_id = models.TextField(primary_key=True, db_column="userId", default=None)
    on_identifiers = models.BooleanField(db_column="onIdentifiers")
    pinned = models.BooleanField(default=False)
    switched_at = UtcDateTimeField(db_column="switchedAt")
    last_applied_event_id = models.TextField(null=True, db_column="lastAppliedEventId")
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "IdentitySignInSwitch"


class IdentifierReservation(models.Model):
    normalized_value = models.TextField(primary_key=True, db_column="normalizedValue", default=None)
    user_id = models.TextField(db_column="userId")
    identifier_id = models.TextField(db_column="identifierId")
    command_id = models.TextField(db_column="commandId")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "IdentifierReservation"
        indexes = [
            models.Index(fields=["user_id"]),
            models.Index(fields=["identifier_id"]),
        ]


class IdentityProjectionCursor(models.Model):
    user_id = models.TextField(primary_key=True, db_column="userId", default=None)
    last_event_id = models.TextField(db_column="lastEventId")
    accepted_at = UtcDateTimeField(db_column="acceptedAt")
    occurred_at = UtcDateTimeField(db_column="occurredAt")
    projection_version = models.TextField(db_column="projectionVersion")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "IdentityProjectionCursor"


class TwoFactor(models.Model):
    """The two-factor plugin's own table (D06). Protocol state, and row-truth for
    good (ADR-101 R12): the shared secret and the backup codes live here and
    in no event, ever. Nothing in this repo reads or writes these columns -
    the plugin owns them end to end, including the lockout counters, which is
    why `failedVerificationCount` and `lockedUntil` are here rather than
    rebuilt on the projection.

    `backupCodes` holds an ENCRYPTED blob, not hashes. A backup code has to be
    compared against what somebody types, and the plugin decrypts and compares;
    hashing them would leave it unable to verify its own codes.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="two_factors", db_column="userId", db_index=False)
    secret = models.TextField()
    backup_codes = models.TextField(db_column="backupCodes")
    verified = models.BooleanField(default=True)
    failed_verification_count = models.IntegerField(db_column="failedVerificationCount", default=0)
    locked_until = UtcDateTimeField(null=True, db_column="lockedUntil")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "TwoFactor"
        indexes = [
            models.Index(fields=["user"]),
            models.Index(fields=["secret"]),
        ]


class SignInAttemptLock(models.Model):
    """Consecutive failed sign-ins, and the lock they earn (GAC-09).
    Spec: specs/identity/org-account-lockout.feature.

    KEYED ON THE ADDRESS THAT WAS TYPED, NOT ON A USER, and that is the whole
    design. We already answer "no such account" and "wrong password" with one
    refusal so that sign-in cannot be used to discover who has an account. A
    lock that existed only for real accounts would put that back: five wrong
    guesses saying "locked" for one address and "incorrect" for another tells
    an attacker which addresses are worth attacking. So an address with
    nothing behind it gets a row, a count and a lock exactly like one that
    resolves, and both answer identically.

    `identifierHash` is a KEYED hash of the normalised address, never the
    address. The key is the deployment's own secret, so this table cannot be
    lifted and reversed into the list of every address anybody has ever tried
    to sign in as - which is what it would otherwise become, including every
    typo and every address that has no account here at all.

    Row-truth, deliberately (ADR-101 R12): this is protocol state read and
    written on the sign-in path, like TwoFactor's own counters next door. The
    EVIDENCE an auditor reads is the identity event emitted when a lock is
    taken, which is durable; these rows are reaped once they stop meaning
    anything.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    identifier_hash = models.TextField(unique=True, db_column="identifierHash")
    """Keyed hash of the normalised address. Never the address itself."""
    failed_count = models.IntegerField(db_column="failedCount", default=0)
    """Consecutive failures since the last success. Reset by a success, not by
    time - "five consecutive", not five ever.
    """
    locked_until = UtcDateTimeField(null=True, db_column="lockedUntil")
    """When the current lock lifts. Null when not locked."""
    consecutive_lockouts = models.IntegerField(db_column="consecutiveLockouts", default=0)
    """Locks taken since the last success. At five this stops lifting by
    itself and becomes an incident - see `heldForReview`.
    """
    held_for_review = models.BooleanField(db_column="heldForReview", default=False)
    """Held past its lock because the count of locks reached the ceiling the
    control names. Cleared by an administrator, or by the owner proving
    mailbox control through a password reset - the second is what keeps a
    held account from being a denial of service anyone can inflict by
    typing a colleague's address wrong often enough.
    """
    user_id = models.TextField(null=True, db_column="userId")
    """Whose account this resolved to at the time, when it resolved to one.
    Null for an address with no account - those rows are the reason this is
    nullable, and the reason this table is not hung off User.
    """
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "SignInAttemptLock"
        indexes = [
            models.Index(fields=["locked_until"]),
            models.Index(fields=["user_id"]),
        ]


class Passkey(models.Model):
    """The passkey plugin's own table (D07). Protocol state, like TwoFactor: the
    public key and the signature counter are the plugin's to read and write.
    A passkey also becomes an `Identifier` row with provider `passkey`, which
    is what makes it a sign-in method like any other - this table is the
    credential, that one is the identity.

    `aaguid` names the authenticator MODEL, never a device or a person, and is
    only present at registration. Privacy-preserving platforms report all
    zeroes, so a label derived from it is best-effort by construction.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    name = models.TextField(null=True)
    public_key = models.TextField(db_column="publicKey")
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="passkeys", db_column="userId", db_index=False)
    credential_id = models.TextField(db_column="credentialID")
    counter = models.IntegerField(default=0)
    device_type = models.TextField(db_column="deviceType")
    backed_up = models.BooleanField(db_column="backedUp", default=False)
    transports = models.TextField(null=True)
    aaguid = models.TextField(null=True)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Passkey"
        indexes = [
            models.Index(fields=["user"]),
            models.Index(fields=["credential_id"]),
        ]


class MfaEnrollment(models.Model):
    """The two-step verification pipeline's projection (D06 - see
    specs/identity/mfa-and-session-shape.feature). One row per PERSON, because
    one person has one setup however many organizations they belong to.

    A pure event-truth head like `SsoConnection`: every column is fold-written,
    rows are never deleted (DISABLED is a tombstone, so history survives being
    turned off), and a replay rebuilds it whole-row from the event log. The
    cursor columns are inline because the fold produces exactly one row per
    aggregate - the `Identifier` fan-out needed its own cursor table, this
    does not.

    Every column here is a lifecycle fact. There is deliberately nowhere in
    this model a secret or a backup code could be put: `consumedBackupCodeIndexes`
    holds POSITIONS, which is what lets "how many are left" be answered
    without the projection ever knowing a code.
    """

    user_id = models.TextField(primary_key=True, db_column="userId", default=None)
    """The user. Also the aggregate id and the tenant - one identity history
    per person (ADR-029 §4).
    """
    enrollment_id = models.TextField(null=True, db_column="enrollmentId")
    """The current enrollment, or NULL for somebody who never started one."""
    method = models.TextField(null=True)
    """totp. Never sms."""
    state = models.TextField()
    """NONE | PENDING | ENABLED | EXPIRED | DISABLED"""
    enrolled_at = UtcDateTimeField(null=True, db_column="enrolledAt")
    confirmed_at = UtcDateTimeField(null=True, db_column="confirmedAt")
    expired_at = UtcDateTimeField(null=True, db_column="expiredAt")
    disabled_at = UtcDateTimeField(null=True, db_column="disabledAt")
    disabled_via = models.TextField(null=True, db_column="disabledVia")
    """password+totp | org-admin - who ended it, and on whose authority."""
    backup_code_count = models.IntegerField(db_column="backupCodeCount", default=0)
    consumed_backup_code_indexes = ArrayField(models.IntegerField(), db_column="consumedBackupCodeIndexes", default=list)
    """Positions already spent, ascending. Never codes."""
    failed_count = models.IntegerField(db_column="failedCount", default=0)
    occurred_at = UtcDateTimeField(db_column="occurredAt")
    last_event_id = models.TextField(db_column="lastEventId")
    accepted_at = UtcDateTimeField(db_column="acceptedAt")
    projection_version = models.TextField(db_column="projectionVersion")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "MfaEnrollment"
        indexes = [
            models.Index(fields=["state"]),
        ]


class SsoConnection(models.Model):
    """The SSO connection pipeline's projection (ADR-117 §5, D04 - see
    specs/identity/sso-connection-lifecycle.feature). A pure event-truth head
    like `Identifier`: every column is fold-written, rows are never deleted
    (TORN_DOWN is a tombstone), and a replay rebuilds it whole-row from the
    event log.

    The fold cursor lives on the row rather than in a sibling table, because
    one row IS one aggregate here - a connection - so there is nothing for a
    separate cursor table to key on that this row does not already key on.

    No relation to Organization on purpose, matching `Identifier`: a
    projection is rebuilt from events, and a foreign key would make a replay
    depend on relational state the log does not carry. `idpMetadata` holds
    endpoints and credential REFERENCES; no secret is ever stored here or in
    any event.
    """

    id = models.TextField(primary_key=True, default=None)
    organization_id = models.TextField(db_column="organizationId")
    type = models.TextField()
    """oidc | saml"""
    state = models.TextField()
    """DRAFT | CLAIMED | APPROVED | REJECTED | DISCARDED |
    VERIFICATION_PENDING | VERIFIED | ACTIVE | SUSPENDED |
    TEARDOWN_PENDING | TORN_DOWN
    """
    claimed_domains = ArrayField(models.TextField(), db_column="claimedDomains")
    domain_claims = models.JSONField(db_column="domainClaims", default=list)
    """[{ domain, state, claimedAtMs, claimedByActorId, decidedAtMs,
    decidedByActorId, authority, waitedMs, note }] - every claim this
    connection has made and how long each waited to be decided. The tier-3
    approval queue is a read of this column, and the epic's Open Q2 wants
    that latency measured from the day the queue exists.
    """
    approved_domains = ArrayField(models.TextField(), db_column="approvedDomains")
    verified_domains = ArrayField(models.TextField(), db_column="verifiedDomains")
    """The only domains that ever route, and only while state is ACTIVE."""
    lapsed_domains = ArrayField(models.TextField(), db_column="lapsedDomains")
    """The subset of `verifiedDomains` whose published record stayed missing
    through its grace window (ADR-123). They still ROUTE - everyone already
    here signs in unchanged - and they stop vouching for anybody NEW: no
    provisioning on first sign-in, no walking in by domain. A column rather
    than a read of the JSON below, because both questions are asked on a
    sign-in path and have to be one indexed predicate.
    """
    domain_verifications = models.JSONField(db_column="domainVerifications", default=list)
    """[{ domain, method, actorId, verifiedAtMs, proofState, firstAbsentAtMs,
    graceEndsAtMs }] - what proved each verified domain, who, and whether
    that evidence is still there. Kept on the row rather than only in the
    log so that an operator-attested domain can never be read as one the
    customer proved, and a wavering one never as one nothing is doubting.
    """
    pending_verification = models.JSONField(null=True, db_column="pendingVerification")
    """{ domain, method, tokenHash } while a ceremony is in flight."""
    idp_metadata = models.JSONField(db_column="idpMetadata")
    arrival_policy = models.TextField(db_column="arrivalPolicy", default="refuse")
    """admit | request | refuse - what this connection does with somebody who
    signs in through it and is not a member yet (ADR-117 3). Stated at
    registration and changed by the setup journey, so it is never absent
    and no reader has to decide what absence means. This was two fields
    for a while - a boolean `allowsJit` beside it, kept in step by the
    fold - and they disagreed about the middle answer.
    """
    allows_jit = models.BooleanField(db_column="allowsJit", default=False)
    """DEAD, AND STAYING FOR ONE RELEASE. Nothing reads or writes this; the
    routing repository derives the same answer from `arrivalPolicy`. It is
    declared so the schema still describes the table, and left in place so
    the PREVIOUS release's pods — which select it by name on every
    connection read — keep working through the rollout. Dropping it in the
    release that stopped reading it would stop single sign-on routing and
    the connection projection for the length of the deploy. Dropped by
    `20260827120000_drop_sso_connection_allows_jit` in the next one.
    """
    arrival_policy_decided_at = UtcDateTimeField(null=True, db_column="arrivalPolicyDecidedAt")
    """When somebody CHOSE the policy above, and null while the registration
    default stands. A different fact from the policy rather than a second
    copy of it: a connection always has a behaviour, and separately
    somebody has or has not decided on it. Going live waits for the
    deciding - "turn everybody away" is a decision too.
    """
    source = models.TextField()
    """self-serve | legacy-grandfathered"""
    test_login_account_id = models.TextField(null=True, db_column="testLoginAccountId")
    replaces_connection_id = models.TextField(null=True, db_column="replacesConnectionId")
    """The grandfathered connection this direct connection replaces. Null for
    ordinary connections and for the grandfathered predecessor itself.

    Partial uniqueness for nonterminal replacements is installed by the
    flattened migration; Prisma cannot express that predicate here.
    """
    migration_phase = models.TextField(choices=SsoConnectionMigrationPhase.choices, null=True, db_column="migrationPhase")
    """SETUP | GRACE_LEGACY | GRACE_DIRECT | FINALIZING | FINALIZED.
    Null means this connection is not a legacy-to-direct replacement.
    """
    grace_started_at = UtcDateTimeField(null=True, db_column="graceStartedAt")
    route_changed_at = UtcDateTimeField(null=True, db_column="routeChangedAt")
    """When an administrator last selected the normal route. This is audit
    evidence only: routing is derived from `migrationPhase`, never recency.
    """
    finalization_requested_at = UtcDateTimeField(null=True, db_column="finalizationRequestedAt")
    finalized_at = UtcDateTimeField(null=True, db_column="finalizedAt")
    rejection = models.JSONField(null=True)
    """{ domain, note } - the last ops rejection, read back on a re-claim."""
    created_by = models.TextField(null=True, db_column="createdBy")
    tear_down_after = UtcDateTimeField(null=True, db_column="tearDownAfter")
    """While TEARDOWN_PENDING: when the process manager's wake completes it."""
    occurred_at = UtcDateTimeField(db_column="occurredAt")
    last_event_id = models.TextField(db_column="lastEventId")
    accepted_at = UtcDateTimeField(db_column="acceptedAt")
    projection_version = models.TextField(db_column="projectionVersion")
    created_at = UtcDateTimeField(db_column="createdAt")
    """Event-derived, not database-managed: `@default(now())` / `@updatedAt`
    would make a replayed row differ from the row it rebuilds, which is
    exactly the parity this projection has to keep.
    """
    updated_at = UtcDateTimeField(db_column="updatedAt")

    class Meta:
        db_table = "SsoConnection"
        indexes = [
            models.Index(fields=["organization_id"]),
            models.Index(fields=["organization_id", "migration_phase"]),
            models.Index(fields=["state"]),
            GinIndex(fields=["verified_domains"]),
            GinIndex(fields=["lapsed_domains"]),
        ]


class SsoConnectionRegistrationSlot(models.Model):
    """The one-statement pre-event registration lock. `SsoConnection` remains
    event truth; this row only decides which command may append that truth.
    A slot whose projected connection is terminal may be atomically replaced.
    """

    pk = models.CompositePrimaryKey("organization_id", "kind")
    organization_id = models.TextField(db_column="organizationId")
    kind = models.TextField()
    """legacy | direct"""
    connection_id = models.TextField(unique=True, db_column="connectionId")
    replaces_connection_id = models.TextField(null=True, db_column="replacesConnectionId")
    command_id = models.TextField(db_column="commandId")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "SsoConnectionRegistrationSlot"


class SsoConnectionReproofCursor(models.Model):
    """Operational round-robin position for ADR-123's domain re-proof sweep.

    This is deliberately not a column on `SsoConnection`: that row is a
    whole-row event projection and its fold is its only writer, while a
    healthy re-read states no fact. Keeping the scheduling cursor beside the
    projection preserves both rules and still lets a fixed-size sweep revisit
    every connection, including one whose proof has started wavering.
    """

    connection = models.OneToOneField("SsoConnection", on_delete=models.CASCADE, related_name="reproof_cursor", primary_key=True, db_column="connectionId")
    last_reproof_at = UtcDateTimeField(db_column="lastReproofAt")

    class Meta:
        db_table = "SsoConnectionReproofCursor"
        indexes = [
            models.Index(fields=["last_reproof_at"]),
        ]


class SsoCredential(models.Model):
    """The values an `SsoConnection`'s credential references point AT (D09 - see
    specs/identity/sso-idp-termination.feature).

    The aggregate has always carried `clientIdRef` / `secretRef` / `certRefs`
    rather than values, because a fact may not carry a secret (ADR-101 §4).
    Until D09 there was nothing behind the references and every writer passed
    null. This is the thing behind them.

    `ciphertext` is AES-256-GCM under the deployment's `CREDENTIALS_SECRET`,
    the same key every other credential in the product is kept under. Nothing
    reads a row except the projection that builds the engine's dialing
    configuration; no surface renders one back.

    Written by the command path rather than the fold, and that split is
    deliberate: the VALUE is not event truth and cannot be, so the log records
    WHEN a credential changed (a new reference in a new fact) and never what
    it was. A replay rebuilds every reference from the log and finds the value
    still here.
    """

    id = models.TextField(primary_key=True, default=None)
    organization_id = models.TextField(db_column="organizationId")
    connection_id = models.TextField(db_column="connectionId")
    """The connection whose registration minted this. Not a relation, matching
    `SsoConnection` itself: a projection rebuilt from events must not depend
    on relational state the log does not carry.
    """
    kind = models.TextField()
    """oidc-client-id | oidc-client-secret | saml-idp-config"""
    ciphertext = models.TextField()
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "SsoCredential"
        indexes = [
            models.Index(fields=["organization_id", "connection_id"]),
        ]


class SsoVerifiedDomain(models.Model):
    """better-auth's single sign-on plugin's provider table (D09), and a
    PROJECTION of the connection log rather than a second source of truth.

    The plugin owns the columns and the reads: it is the only thing that
    speaks OpenID Connect discovery or reads a SAML assertion, and it dials
    what it finds here. What it never does is decide that a row should exist —
    that is folded from `SsoConnection`'s events by the same projection that
    writes the connection head, in the same transaction, so a replay rebuilds
    both together and a row can never outlive the connection that justified
    it. Nothing else writes here.

    `providerId` is the CONNECTION ID rather than the name the customer typed.
    The plugin's column is globally unique, and two organizations both calling
    their provider `okta` must both be able to.

    `oidcConfig` and `samlConfig` are the plugin's own JSON documents, rebuilt
    on every fold from `SsoCredential` — the vault is where the material
    lives, and this row is a derived copy the engine can read synchronously.

    SEALED AT REST, under the same key the vault uses. The document carries
    the client secret the derivation reads OUT of the vault, so a plaintext
    copy here would hand every customer's live identity-provider credential
    to anybody holding a database dump — the exact threat `SsoCredential`
    exists to defeat, reopened one table over. A stored value announces its
    form (`enc:v1:`), so a row written before the seal existed still opens;
    the one reader is the identity storage adapter, on the path that dials.
    The reader tolerates both forms, so this needs no backfill — but a pod
    without it cannot read a sealed row, so it must not be rolled back past.
    WHO OWNS A PROVED DOMAIN, as a row the database can refuse twice.

    "First verifier owns" was enforced by reading `findDomainOwner` and then
    writing the fact — two different aggregates, so the per-connection queue
    never ordered them against each other, and `SsoConnection.verifiedDomains`
    is a `String[]` with only a GIN index behind it, which enforces nothing.
    Two organizations proving `acme.com` at the same moment both read "nobody
    owns it" and both committed; the sign-in router then picked between two
    live connections by planner order, so a member of one company could be
    handed to another company's identity provider — and with `arrivalPolicy`
    on `admit`, provisioned into it.

    The legacy single-domain model had this: `Organization.ssoDomain` is
    `@unique`. The multi-domain rewrite dropped the guarantee, and self-serve
    proof is what made the race reachable by customers rather than by two
    operators working at once.

    `domain` is the primary key, so a second organization loses on the
    constraint rather than on a check it raced. Connection holders are
    separate because an exact predecessor/replacement pair must coexist
    during migration. Both are projected with the connection head.
    """

    domain = models.TextField(primary_key=True, default=None)
    organization_id = models.TextField(db_column="organizationId")

    class Meta:
        db_table = "SsoVerifiedDomain"
        constraints = [
            models.UniqueConstraint(fields=["domain", "organization_id"], name="SsoVerifiedDomain_domain_organizationId_key"),
        ]
        indexes = [
            models.Index(fields=["organization_id"]),
        ]


class SsoVerifiedDomainHolder(models.Model):
    """A connection currently entitled to use an organization-owned verified
    domain. Usually there is one holder. During an Auth0 migration there are
    exactly two: the grandfathered predecessor and its direct replacement.

    The projection repository enforces that relationship before inserting a
    second holder. The flattened migration also installs partial unique live
    legacy/direct slots on `SsoConnection`, closing concurrent third-writer
    races that Prisma cannot describe declaratively.
    """

    pk = models.CompositePrimaryKey("domain", "connection_id")
    domain = models.TextField()
    connection_id = models.TextField(db_column="connectionId")
    organization_id = models.TextField(db_column="organizationId")

    class Meta:
        db_table = "SsoVerifiedDomainHolder"
        indexes = [
            models.Index(fields=["connection_id"]),
            models.Index(fields=["organization_id"]),
        ]


class SsoAuthenticationActivity(models.Model):
    """A successful callback through one SSO connection. Identifier.lastUsedAt
    also advances during routing and resolution, so it cannot prove that a
    member actually authenticated through the replacement before cutover.
    """

    id = models.TextField(primary_key=True, default=None)
    organization_id = models.TextField(db_column="organizationId")
    connection_id = models.TextField(db_column="connectionId")
    user_id = models.TextField(db_column="userId")
    authenticated_at = UtcDateTimeField(db_column="authenticatedAt")
    provider_account_id = models.TextField(null=True, db_column="providerAccountId")
    """The provider-side account the assertion named; null before #7633's arrival door recorded it."""

    class Meta:
        db_table = "SsoAuthenticationActivity"
        indexes = [
            models.Index(fields=["organization_id", "connection_id", "authenticated_at"]),
            models.Index(fields=["organization_id", "user_id"]),
        ]


class SsoProvider(models.Model):
    id = models.TextField(primary_key=True, default=None)
    issuer = models.TextField()
    oidc_config = models.TextField(null=True, db_column="oidcConfig")
    saml_config = models.TextField(null=True, db_column="samlConfig")
    user_id = models.TextField(null=True, db_column="userId")
    provider_id = models.TextField(unique=True, db_column="providerId")
    organization_id = models.TextField(null=True, db_column="organizationId")
    domain = models.TextField()

    class Meta:
        db_table = "SsoProvider"
        indexes = [
            models.Index(fields=["organization_id"]),
            models.Index(fields=["domain"]),
        ]


class SsoBreakGlassBinding(models.Model):
    """A way back in (D05 - see specs/identity/sso-onboarding-tiers.feature).

    One named person who may still sign in without the identity provider
    while their organization's sign-in belongs to one, and the date on which
    that stops. Activation asks whether such a person exists, because the
    failure it prevents is the only one that cannot be recovered from inside
    the product: a misconfigured connection goes live, every member is handed
    to an identity provider that will not authenticate them, and nobody is
    left who can turn it off.

    Rows are immutable except for `supersededAt` and `warnedDays`. A renewal
    writes a NEW row naming the one it replaced, so the date a way in
    previously ended is still readable in the history - which is what makes a
    renewal auditable rather than a column somebody moved.

    No relation to Organization or User, matching the identity heads above:
    the binding is the record of who could get in, and a cascade would erase
    exactly the history a lockout post-mortem reads.
    """

    id = models.TextField(primary_key=True, default=None)
    organization_id = models.TextField(db_column="organizationId")
    user_id = models.TextField(db_column="userId")
    """Who holds the way in."""
    granted_by_user_id = models.TextField(db_column="grantedByUserId")
    """Who granted it. A way back in is never self-served."""
    granted_at = UtcDateTimeField(db_column="grantedAt")
    expires_at = UtcDateTimeField(db_column="expiresAt")
    superseded_at = UtcDateTimeField(null=True, db_column="supersededAt")
    """Set when a renewal replaced this row; the row itself is never edited."""
    renewed_from_id = models.TextField(null=True, db_column="renewedFromId")
    """The binding this one renewed, so the previous end date stays readable."""
    warned_days = ArrayField(models.IntegerField(), db_column="warnedDays")
    """Which of the fourteen, seven and one day warnings have been sent, so a
    sweep that runs twice in a day never says the same thing twice.
    """
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "SsoBreakGlassBinding"
        indexes = [
            models.Index(fields=["organization_id", "expires_at"]),
            models.Index(fields=["user_id"]),
        ]


class SsoActivationRecoveryReservation(models.Model):
    """A durable claim on one live recovery path while activation is between its
    guard and projected ACTIVE fact. It has no expiry: only projection of the
    matching activation, or an explicit cancellation, may consume it.
    """

    command_id = models.TextField(primary_key=True, db_column="commandId", default=None)
    organization_id = models.TextField(db_column="organizationId")
    connection_id = models.TextField(db_column="connectionId")
    created_at = UtcDateTimeField(db_column="createdAt")

    class Meta:
        db_table = "SsoActivationRecoveryReservation"
        constraints = [
            models.UniqueConstraint(fields=["organization_id", "connection_id"], name="SsoActivationRecoveryReservation_organizationId_connectionId_key"),
        ]
        indexes = [
            models.Index(fields=["organization_id"]),
        ]


class JoinRequest(models.Model):
    """The join-request pipeline's projection (ADR-117, D12 - see
    specs/identity/join-requests.feature). A pure event-truth head like
    `Identifier` and `SsoConnection`: every column is fold-written, rows are
    never deleted (a terminal state is a tombstone the admin panel stops
    listing, not an absence), and a replay rebuilds it whole-row from the
    event log.

    The fold cursor lives on the row rather than in a sibling table, because
    one row IS one aggregate here - a request.

    No relation to Organization or User on purpose, matching the two heads
    above: a projection is rebuilt from events, and a foreign key would make a
    replay depend on relational state the log does not carry.

    Membership is NOT here. An approval dispatches an attach on the grants
    ledger, exactly as accepting an invitation does, and that ledger fact
    carries `source: "join-request"` as its provenance - which is what puts a
    join approval on the customer's audit page beside every other membership
    change rather than in a history of its own.

    The requester's address is not here either: the DOMAIN is the fact, and it
    is the only part of an address a join decision may depend on.
    """

    id = models.TextField(primary_key=True, default=None)
    user_id = models.TextField(db_column="userId")
    organization_id = models.TextField(db_column="organizationId")
    domain = models.TextField()
    """The normalized domain the match was made on."""
    state = models.TextField()
    """PENDING | APPROVED | REJECTED | EXPIRED | WITHDRAWN"""
    matched_via = models.TextField(db_column="matchedVia")
    """How the domain was matched: `verified-identifier-domain` or `sso-connection-domain`."""
    origin = models.TextField(default="web")
    """Where the request was made: `web` or `cli` (ADR-171 v6). A request from
    the terminal lands a Developer seat whatever the joiner seat says.
    """
    expires_at = UtcDateTimeField(null=True, db_column="expiresAt")
    """While PENDING: when the day-14 wake lapses it. Null once resolved."""
    resolved_at = UtcDateTimeField(null=True, db_column="resolvedAt")
    resolved_by_type = models.TextField(null=True, db_column="resolvedByType")
    """user | policy | invite - who ended it. `policy` is domain auto-join's
    own principal, deliberately not the one SSO auto-join uses.
    """
    resolved_by_id = models.TextField(null=True, db_column="resolvedById")
    """The admin's user id, `domain-auto`, or the invitation id."""
    withdrawal_cause = models.TextField(null=True, db_column="withdrawalCause")
    """user | invite-accepted, on a withdrawal."""
    connection_id = models.TextField(null=True, db_column="connectionId")
    """The SSO connection a single sign-on arrival was admitted through; null otherwise."""
    occurred_at = UtcDateTimeField(db_column="occurredAt")
    last_event_id = models.TextField(db_column="lastEventId")
    accepted_at = UtcDateTimeField(db_column="acceptedAt")
    projection_version = models.TextField(db_column="projectionVersion")
    created_at = UtcDateTimeField(db_column="createdAt")
    """Event-derived, not database-managed: `@default(now())` / `@updatedAt`
    would make a replayed row differ from the row it rebuilds, which is
    exactly the parity this projection has to keep.
    """
    updated_at = UtcDateTimeField(db_column="updatedAt")

    class Meta:
        db_table = "JoinRequest"
        indexes = [
            models.Index(fields=["organization_id", "state"]),
            models.Index(fields=["user_id", "state"]),
            models.Index(fields=["user_id", "organization_id"]),
            models.Index(fields=["domain"]),
        ]


class TeamUser(models.Model):
    pk = models.CompositePrimaryKey("user_id", "team_id")
    user = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="team_memberships", db_column="userId", db_index=False)
    team = models.ForeignKey("Team", on_delete=models.RESTRICT, related_name="members", db_column="teamId", db_index=False)
    role = models.TextField(choices=TeamUserRole.choices)
    assigned_role = models.ForeignKey("CustomRole", on_delete=models.SET_NULL, related_name="assigned_users", db_column="assignedRoleId", null=True, db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "TeamUser"
        indexes = [
            models.Index(fields=["team"]),
            models.Index(fields=["user"]),
            models.Index(fields=["assigned_role"]),
        ]


class OrganizationUser(models.Model):
    pk = models.CompositePrimaryKey("user_id", "organization_id")
    user = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="org_memberships", db_column="userId", db_index=False)
    organization = models.ForeignKey("Organization", on_delete=models.RESTRICT, related_name="members", db_column="organizationId", db_index=False)
    role = models.TextField(choices=OrganizationUserRole.choices)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    department_id = models.TextField(null=True, db_column="departmentId")
    """Department assignment for spend attribution (the people lens).
    Single-valued, org-scoped, accounting-only (never gates access).
    Plain column ref (relationMode="prisma"); the service validates the
    id belongs to this org. See departments.feature.
    """
    disabled_at = UtcDateTimeField(null=True, db_column="disabledAt")
    """Set when an admin disables this membership to stay within the licensed
    seat count. The person loses access to this organization only, keeps
    their role, department and everything they did, and frees their seat.
    Reversible: re-enabling restores access if a seat is available.
    Seat counting excludes disabled memberships. See
    seat-reconciliation.feature.
    """
    pending_sso_grant_id = models.TextField(null=True, db_column="pendingSsoGrantId")
    """Stable grant intent for an unfinished automatic SSO admission.
    Cleared after the authoritative grant is observed; never inferred from missing access.
    """
    membership_stamp = models.TextField(db_column="membershipStamp", db_default=RawSQL("gen_random_uuid()::text", []))
    """A new generation for each membership lifetime. Stamped USER grant facts
    must match the current active membership before they can be projected;
    removing and rejoining therefore cannot revive a delayed old attach.
    """
    full_seat_since = UtcDateTimeField(null=True, db_column="fullSeatSince")
    """When this membership started holding the full seat it holds now; null means createdAt.
    Anchors the 30-day seat hold (specs/licensing/seat-rotation.feature).
    """
    seat_pending_since = UtcDateTimeField(null=True, db_column="seatPendingSince")
    """Set when admission wrote this membership disabled because no seat was free: the person
    is waiting for one. Cleared when the membership is enabled.
    """
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "OrganizationUser"
        indexes = [
            models.Index(fields=["organization"]),
            models.Index(fields=["user"]),
            models.Index(fields=["department_id"]),
        ]


class SeatHold(models.Model):
    """A full seat still counted after the person who took it stopped holding it (removed,
    disabled or demoted), until 30 days after they took it (specs/licensing/seat-rotation.feature).
    One per organization and person; a return within the hold takes the seat back.
    """

    id = models.TextField(primary_key=True, default=None)
    organization_id = models.TextField(db_column="organizationId")
    user_id = models.TextField(db_column="userId")
    email = models.TextField(null=True)
    """Display copy for the Users list, kept because a removed person has no membership row."""
    name = models.TextField(null=True)
    seat_taken_at = UtcDateTimeField(db_column="seatTakenAt")
    """When the person took the seat this hold keeps."""
    held_until = UtcDateTimeField(db_column="heldUntil")
    """seatTakenAt plus 30 days; expired holds are ignored by every count and deleted hourly."""
    reason = models.TextField()
    """REMOVED | DISABLED | DEMOTED"""
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "SeatHold"
        constraints = [
            models.UniqueConstraint(fields=["organization_id", "user_id"], name="SeatHold_organizationId_userId_key"),
        ]
        indexes = [
            models.Index(fields=["held_until"]),
        ]


class SeatRequest(models.Model):
    """A person asked for a full seat the organization had no room for (modules/organization/specs/
    seat-requests.feature). One row per organization and address; reopening reuses the row.
    """

    id = models.TextField(primary_key=True, default=None)
    organization_id = models.TextField(db_column="organizationId")
    email = models.TextField()
    """Lowercased address of the person asking."""
    user_id = models.TextField(null=True, db_column="userId")
    """Null when only the address is known (a sign-in the identity provider refused)."""
    name = models.TextField(null=True)
    status = models.TextField()
    """OPEN | APPROVED | DECLINED"""
    source = models.TextField()
    """pending | lite | idp_refused: how the person was left when the request opened."""
    idp_refused = models.BooleanField(db_column="idpRefused", default=False)
    """The identity provider also refused this person's sign-in."""
    idp_kind = models.TextField(null=True, db_column="idpKind")
    """okta | entra | google | onelogin | jumpcloud | ping | auth0 | oidc | saml"""
    sales_notified_at = UtcDateTimeField(null=True, db_column="salesNotifiedAt")
    """When sales were told about this request; reset when it reopens."""
    opened_at = UtcDateTimeField(db_column="openedAt")
    resolved_at = UtcDateTimeField(null=True, db_column="resolvedAt")
    resolved_by_user_id = models.TextField(null=True, db_column="resolvedByUserId")
    seat_pending_since = UtcDateTimeField(null=True, db_column="seatPendingSince")
    """Set when an approval claims the request, cleared once its seat is given; an approval
    still pending long after is one whose seat never landed, and is reopened.
    """
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "SeatRequest"
        constraints = [
            models.UniqueConstraint(fields=["organization_id", "email"], name="SeatRequest_organizationId_email_key"),
        ]
        indexes = [
            models.Index(fields=["organization_id", "status"]),
            models.Index(fields=["user_id"]),
        ]


class OrganizationSeatLease(models.Model):
    """Held by the one seat-taking write running on an organization (seat-requests.feature). A
    row past expiresAt is free to take; a holder deletes its own row when its write is done.
    """

    organization_id = models.TextField(primary_key=True, db_column="organizationId", default=None)
    holder = models.TextField()
    """Random per hold, so a release deletes only its own lease."""
    expires_at = UtcDateTimeField(db_column="expiresAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "OrganizationSeatLease"


class OrganizationAccessNotice(models.Model):
    """The mail log for seat requests and refused sign-ins: at most one row per organization,
    kind, address and UTC day, each mailed at once.
    """

    id = models.TextField(primary_key=True, default=None)
    organization_id = models.TextField(db_column="organizationId")
    kind = models.TextField()
    """seat_request | idp_refused"""
    email = models.TextField()
    name = models.TextField(null=True)
    idp = models.TextField(null=True)
    day_key = models.TextField(db_column="dayKey")
    """UTC day the notice was recorded on, YYYY-MM-DD."""
    idp_refused = models.BooleanField(db_column="idpRefused", default=False)
    """Whether the seat request it announces was also refused by the identity provider."""
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    mailed_at = UtcDateTimeField(null=True, db_column="mailedAt")
    """Set once the notice was handed to the mailer; null if a send never started."""

    class Meta:
        db_table = "OrganizationAccessNotice"
        constraints = [
            models.UniqueConstraint(fields=["organization_id", "kind", "email", "day_key"], name="OrganizationAccessNotice_organizationId_kind_email_dayKey_key"),
        ]


class Team(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    name = models.TextField()
    slug = models.TextField(unique=True)
    organization = models.ForeignKey("Organization", on_delete=models.RESTRICT, related_name="teams", db_column="organizationId", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    archived_at = UtcDateTimeField(null=True, db_column="archivedAt")
    is_personal = models.BooleanField(db_column="isPersonal", default=False)
    owner_user = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="personal_teams", db_column="ownerUserId", null=True, db_index=False)
    department_id = models.TextField(null=True, db_column="departmentId")
    """Department assignment inherited as the default for members and
    projects that have none of their own. Accounting-only.
    """
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Team"
        indexes = [
            models.Index(fields=["organization"]),
            models.Index(fields=["is_personal"]),
            models.Index(fields=["department_id"]),
        ]


class Organization(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    name = models.TextField()
    kind = models.TextField(default="customer")
    """`customer` for every organization a person founds or an operator provisions; `system` for
    the install's one hidden organization the Ops checkup canaries run in, never listed to a person.
    Spec: modules/ops/specs/checkup-canary-project.feature.
    """
    phone_number = models.TextField(null=True, db_column="phoneNumber")
    slug = models.TextField(unique=True)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    usage_spending_max_limit = models.IntegerField(null=True, db_column="usageSpendingMaxLimit")
    dataset_attachment_max_mb = models.IntegerField(null=True, db_column="datasetAttachmentMaxMb")
    """The largest file a dataset image or file cell accepts for this
    organization, in mebibytes. Null means the default. Set by an operator
    from the back office only; every other dataset size limit scales with it.
    """
    private_click_house_label = models.TextField(null=True, db_column="privateClickHouseLabel")
    """The label of this organization's `CLICKHOUSE_URL__<label>__<id>` route, recorded by the
    organization module the first time it is seen. Once set, a missing route refuses the
    organization instead of serving it from the shared ClickHouse. Null: never private.
    Spec: specs/private-dataplane/clickhouse-routing.feature.
    """
    private_s3_label = models.TextField(null=True, db_column="privateS3Label")
    """The same for `DATAPLANE_S3__<label>__<id>`. Spec: specs/private-dataplane/s3-routing.feature."""
    private_dataplane_seen_at = UtcDateTimeField(null=True, db_column="privateDataplaneSeenAt")
    """When either label was first recorded."""
    max_session_duration_days = models.IntegerField(db_column="maxSessionDurationDays", default=0)
    """Phase 8 - admin-configurable max lifetime for CLI device-flow
    sessions, in days. 0 = unbounded (default, mirrors GitHub CLI /
    gh-style flows). Companies that need finite session lifetime
    (eg. SOC 2 / ISO 27001 customers with mandated session
    expiration) can set a positive value. The /api/auth/cli/refresh
    endpoint rejects refresh attempts where now - issued_at exceeds
    this value, regardless of the underlying refresh-token TTL.
    Spec: specs/ai-governance/sessions/session-ttl.feature.
    """
    mfa_required = models.BooleanField(db_column="mfaRequired", default=False)
    """Every member of this organization can prove a second factor (D06). A
    MEMBERSHIP CONDITION, not a session policy: it is asked when a member
    reaches this organization's data, it holds the ones who cannot prove
    one at an enrollment gate for THIS organization alone, and it ends no
    session, ever. Turning it on signs nobody out.
    Spec: specs/identity/mfa-and-session-shape.feature.
    """
    docs_playground_sessions_enabled = models.BooleanField(db_column="docsPlaygroundSessionsEnabled", default=True)
    """Whether members may connect the documentation playground to their own access with a
    short-lived docs session key (ADR-185, amendment of 2026-11-13). On by default; turned off,
    no key is minted and live ones stop authenticating. The pasted-key path is untouched.
    Spec: modules/api-key/specs/docs-session-keys.feature.
    """
    docs_playground_sessions_disabled_at = UtcDateTimeField(null=True, db_column="docsPlaygroundSessionsDisabledAt")
    """When the docs playground was last turned off, written in the same update as the switch. A
    docs session key minted at or before it never authenticates again, nor any key minted under
    one, so turning it back on revives none of them (ADR-185, amendment of 2026-11-14).
    """
    lockout_after_failed_attempts = models.IntegerField(db_column="lockoutAfterFailedAttempts", default=0)
    """How many consecutive failed sign-ins lock an account (GAC-09).
    0 = never lock, which is the default and what every organization that
    has not asked for this keeps. Counts EVERY credential together - a
    password, a one-time code, a backup code, a passkey - because five per
    factor hands an attacker one budget per factor.
    Spec: specs/identity/org-account-lockout.feature.
    """
    lockout_minutes = models.IntegerField(db_column="lockoutMinutes", default=30)
    """How long a lock lasts, in minutes. Ignored while the threshold is 0."""
    session_idle_timeout_minutes = models.IntegerField(db_column="sessionIdleTimeoutMinutes", default=1440)
    """How long a BROWSER session may sit idle before it ends, in minutes
    (GAC-10). 0 = no idle timeout. Activity rolls the window, which is the
    "refreshed every hour" half of the control.

    A day for organizations created from now on; organizations that predate
    the column keep 0 and are unchanged, because the migration adds the
    column at zero and moves the default afterwards.

    NOT `maxSessionDurationDays`, which bounds CLI and device sessions and
    is enforced when the CLI refreshes. Two surfaces, two lifetimes, two
    decisions an administrator makes separately.
    Spec: specs/identity/org-session-lifetime.feature.
    """
    session_max_lifetime_minutes = models.IntegerField(db_column="sessionMaxLifetimeMinutes", default=0)
    """The ceiling on a browser session from the moment it was minted, in
    minutes, which no amount of activity extends. 0 = no ceiling, the
    default. Offered beside the idle timeout because a session kept alive
    forever by a polling page is not really bounded.
    """
    signup_data = models.JSONField(null=True, db_column="signupData")
    signed_dpa = models.BooleanField(db_column="signedDPA", default=False)
    elasticsearch_node_url = models.TextField(null=True, db_column="elasticsearchNodeUrl")
    elasticsearch_api_key = models.TextField(null=True, db_column="elasticsearchApiKey")
    use_custom_elasticsearch = models.BooleanField(db_column="useCustomElasticsearch", default=False)
    s3_endpoint = models.TextField(null=True, db_column="s3Endpoint")
    s3_access_key_id = models.TextField(null=True, db_column="s3AccessKeyId")
    s3_secret_access_key = models.TextField(null=True, db_column="s3SecretAccessKey")
    s3_bucket = models.TextField(null=True, db_column="s3Bucket")
    use_custom_s3 = models.BooleanField(db_column="useCustomS3", default=False)
    sent_plan_limit_alert = UtcDateTimeField(null=True, db_column="sentPlanLimitAlert")
    sso_domain = models.TextField(null=True, unique=True, db_column="ssoDomain")
    sso_provider = models.TextField(null=True, db_column="ssoProvider")
    domain_join = models.TextField(db_column="domainJoin", default="request")
    """How colleagues on a matching domain get in (D12, ADR-117):
    `off` (invisible to them entirely) | `request` (they ask, an admin
    approves) | `auto` (a verified colleague walks straight in).

    `request` is the default for a self-serve organization; `auto` is never
    a default and never inferred - an administrator turns it on, names the
    domain in `joinDomains` while doing it, and the licence gate holds that
    value while letting `request` through.
    """
    joiner_role = models.TextField(choices=OrganizationUserRole.choices, db_column="joinerRole", default=OrganizationUserRole.MEMBER)
    """The organisation role handed to a person admitted without an invitation:
    a self-service domain join (`domainJoin`) or an SSO-admitted login.
    `MEMBER` (default, a Full seat) or `DEVELOPER` (ADR-171). Set on the
    organisation settings page only; the join flow never asks, and an
    invitation always names its own role.
    """
    join_domains = ArrayField(models.TextField(), db_column="joinDomains")
    """The domains an administrator NAMED when turning automatic joining on.
    Empty means automatic joining admits nobody whatever `domainJoin` says -
    which is what keeps a setting left on from admitting a domain nobody
    ever agreed to.
    """
    presence_enabled = models.BooleanField(db_column="presenceEnabled", default=True)
    trace_sharing_enabled = models.BooleanField(db_column="traceSharingEnabled", default=True)
    """Org-level trace-sharing kill switch. Off = trace share links cannot be
    created in ANY project of this org, and existing links stop resolving —
    the org toggle is the global switch, the per-project
    `Project.traceSharingEnabled` scopes the kill to a single project (mirrors
    `presenceEnabled`). Effective sharing = org AND project. ADR-039.
    """
    support_contact = models.TextField(null=True, db_column="supportContact")
    """Org-configurable contact surfaced to users in CLI "contact your
    admin" messages and the in-app budget-exceeded banner. Free text:
    accepts an email, a URL (internal ticketing system), or any other
    instruction. When null, callers fall back to the first ADMIN
    member's email (`resolveSupportContact` helper).
    """
    primary_intent = models.TextField(choices=OrganizationIntent.choices, null=True, db_column="primaryIntent")
    """Signup intent, editable later as the org "Primary use" setting.
    AGENT_GOVERNANCE lands "/" on /me, LLM_OPS on the project home.
    NULL means intent unset (every pre-fork org and any org created
    outside onboarding) — the home resolver keeps legacy behavior.
    ADR-038.
    """
    promo_code = models.TextField(null=True, db_column="promoCode")
    stripe_customer_id = models.TextField(null=True, unique=True, db_column="stripeCustomerId")
    currency = models.TextField(choices=Currency.choices, default=Currency.EUR)
    pricing_model = models.TextField(choices=PricingModel.choices, db_column="pricingModel", default=PricingModel.TIERED)
    license = models.TextField(null=True)
    license_expires_at = UtcDateTimeField(null=True, db_column="licenseExpiresAt")
    license_last_validated_at = UtcDateTimeField(null=True, db_column="licenseLastValidatedAt")
    self_hosted_customer = models.BooleanField(db_column="selfHostedCustomer", default=False)
    """A customer that runs LangWatch self-hosted. The hosted usage, budgets and
    invoices of its connected installs attach to this organization (ADR-139).
    """
    connect_services_disabled = ArrayField(models.TextField(), db_column="connectServicesDisabled", default=list)
    """LangWatch-hosted services this organization has switched OFF, named the
    way its license names them ("instant_evals"). A service the license is
    entitled to is on unless it is listed here, so a customer who bought
    hosted judging gets it without finding a settings page first, and an
    administrator who switched one off keeps it off across license changes
    (ADR-139, section 2). A license that names no service reaches nothing,
    whatever this column holds.
    """
    instant_evals_enabled_at = UtcDateTimeField(null=True, db_column="instantEvalsEnabledAt")
    """When a member of this organization switched Instant Evals on from the
    search bar, which is the organization's consent to have judged text sent
    to the judge's provider. Null until then: a self-serve organization is
    never opted in on its behalf, and an enterprise one is switched on by an
    operator through the release flag after asking, never here.
    """
    instant_evals_enabled_by_user_id = models.TextField(null=True, db_column="instantEvalsEnabledByUserId")
    """The member who switched them on, so the audit trail names a person and
    not just a moment.
    """
    connect_last_sync_at = UtcDateTimeField(null=True, db_column="connectLastSyncAt")
    """When the last license sync succeeded."""
    connect_last_sync_error = models.TextField(null=True, db_column="connectLastSyncError")
    """The code of the failure the last sync ended on, cleared on the next
    success. Settings, Connect renders the copy for it.
    """
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Organization"
        indexes = [
            models.Index(fields=["sso_provider"]),
        ]


class IssuedLicense(models.Model):
    """The registry of licenses LangWatch has issued (ADR-139). It lives on
    LangWatch Cloud; a self-hosted install has the table and no rows.

    The row never holds the token an install presents. `tokenHash` is the
    SHA-256 of that token, so read access to this table is not credential
    access. Status is derived: revoked, superseded, past `expiresAt`, else active.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    license_id = models.TextField(unique=True, db_column="licenseId")
    """The `licenseId` inside the signed license (`lic-<uuid>`)."""
    token_hash = models.TextField(unique=True, db_column="tokenHash")
    """SHA-256 of the `lwl_` token. Looked up on every credential resolution."""
    organization = models.ForeignKey("Organization", on_delete=models.RESTRICT, related_name="issued_licenses", db_column="organizationId", null=True, db_index=False)
    """The customer organization on LangWatch Cloud. Null for a license that is
    recorded but not linked to a customer yet, such as a self-serve purchase:
    it resolves to nothing until an operator links it.
    """
    organization_name = models.TextField(db_column="organizationName")
    """As signed into the license."""
    email = models.TextField()
    plan_type = models.TextField(db_column="planType")
    max_members = models.IntegerField(db_column="maxMembers")
    max_members_lite = models.IntegerField(null=True, db_column="maxMembersLite")
    """The lite seats sold: null when not stated or when unlimited. Never a stand-in number."""
    lite_seats_unlimited = models.BooleanField(db_column="liteSeatsUnlimited", default=False)
    """True when the license sells unlimited lite seats; maxMembersLite is then null."""
    max_members_developer = models.IntegerField(null=True, db_column="maxMembersDeveloper")
    """The Developer seats sold: null when not stated, which sells none."""
    issued_at = UtcDateTimeField(db_column="issuedAt")
    expires_at = UtcDateTimeField(db_column="expiresAt")
    source = models.TextField(choices=IssuedLicenseSource.choices)
    issued_by_id = models.TextField(null=True, db_column="issuedById")
    """The operator who issued or registered it. Null for the purchase flow and the script."""
    revoked_at = UtcDateTimeField(null=True, db_column="revokedAt")
    revoked_by_id = models.TextField(null=True, db_column="revokedById")
    revoked_reason = models.TextField(null=True, db_column="revokedReason")
    superseded_at = UtcDateTimeField(null=True, db_column="supersededAt")
    """Set when the install first presents the license that replaced this one."""
    replaces_id = models.TextField(null=True, unique=True, db_column="replacesId")
    """The id of the IssuedLicense this one replaces, when it was reissued. A
    plain column ref (relationMode="prisma" cannot declare a self-relation on
    Postgres); the registry service validates it. Unique, so a license is
    replaced by one license at most.
    """
    pending_delivery_license = models.TextField(null=True, db_column="pendingDeliveryLicense")
    """A reissued license waiting to be delivered over sync, encrypted. Erased
    when the install first presents it. The only license text the registry holds.
    """
    services = ArrayField(models.TextField(), default=list)
    """Hosted services this license may call, such as `instant_evals`."""
    seat_rate_cents = models.IntegerField(null=True, db_column="seatRateCents")
    """The annual rate of one seat, in cents of `seatCurrency`. Needed to prorate added seats."""
    seat_currency = models.TextField(choices=Currency.choices, null=True, db_column="seatCurrency")
    commit_usd_cents = models.IntegerField(db_column="commitUsdCents", default=0)
    """The prepaid usage commit for the term, in USD cents."""
    overage_enabled = models.BooleanField(db_column="overageEnabled", default=False)
    """Whether usage may continue past the commit, up to `overageMaxUsdCents`."""
    overage_max_usd_cents = models.IntegerField(null=True, db_column="overageMaxUsdCents")
    instance_id = models.TextField(null=True, db_column="instanceId")
    """The install this license is bound to. Set on first use, cleared by an operator."""
    instance_bound_at = UtcDateTimeField(null=True, db_column="instanceBoundAt")
    last_sync_at = UtcDateTimeField(null=True, db_column="lastSyncAt")
    """The last license sync this install reported: when it arrived, the
    LangWatch version that sent it, and the seats in use at that moment.
    Null until the install syncs for the first time.
    """
    last_sync_version = models.TextField(null=True, db_column="lastSyncVersion")
    reported_members = models.IntegerField(null=True, db_column="reportedMembers")
    reported_members_lite = models.IntegerField(null=True, db_column="reportedMembersLite")
    virtual_key_id = models.TextField(null=True, unique=True, db_column="virtualKeyId")
    """The managed gateway key (purpose CONNECT) that hosted-service traffic of
    this license runs under. One per license, so spend rows name the install.
    Created when the license first resolves. A plain column, not a relation.
    """
    seats_raised_from = models.IntegerField(null=True, db_column="seatsRaisedFrom")
    """The seats of the license this one replaced, set only when a seat change
    raised a linked license. Billing reads it through licensing to invoice
    the added seats, so the fact commits with the license or not at all.
    """
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "IssuedLicense"
        indexes = [
            models.Index(fields=["organization"]),
            models.Index(fields=["expires_at"]),
            models.Index(fields=["seats_raised_from"]),
        ]


class OrganizationLicense(models.Model):
    """The licence an organization activated and its dates, owned by licensing
    (round 37 D6). A row whose key is null is a cleared licence. An organization
    with no row is read from its Organization columns until the copy step runs.
    """

    organization_id = models.TextField(primary_key=True, db_column="organizationId", default=None)
    license_key = models.TextField(null=True, db_column="licenseKey")
    expires_at = UtcDateTimeField(null=True, db_column="expiresAt")
    validated_at = UtcDateTimeField(null=True, db_column="validatedAt")
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "OrganizationLicense"


class ActivationCode(models.Model):
    """A short code a fresh install pastes instead of a license blob (ADR-139,
    section 5). Redeeming it mints the license the code describes and returns
    it signed, which is what identifies a new customer without anybody emailing
    a license key around.

    The row holds a hash of the code, never the code. Every lookup hashes what
    the caller presented first, so the stored value is a step away from anything
    that could be presented: read access to this table is not credential access.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    code_hash = models.TextField(unique=True, db_column="codeHash")
    """SHA-256 of the normalised code. Looked up on every redemption."""
    code_hint = models.TextField(db_column="codeHint")
    """The last four characters, so an operator tells two codes apart in a list
    without the list holding either one.
    """
    organization_id = models.TextField(db_column="organizationId")
    """The customer this code belongs to, and the terms it mints."""
    organization_name = models.TextField(db_column="organizationName")
    email = models.TextField()
    plan_type = models.TextField(db_column="planType")
    max_members = models.IntegerField(db_column="maxMembers")
    max_members_lite = models.IntegerField(null=True, db_column="maxMembersLite", default=0)
    """The lite seats the minted license carries: null when unlimited. Never a stand-in number."""
    lite_seats_unlimited = models.BooleanField(db_column="liteSeatsUnlimited", default=False)
    """True when the minted license carries unlimited lite seats; maxMembersLite is then null."""
    max_members_developer = models.IntegerField(null=True, db_column="maxMembersDeveloper")
    """The Developer seats the minted license carries: null when not stated, which sells none."""
    license_term_days = models.IntegerField(db_column="licenseTermDays")
    """How long the license this code mints runs for."""
    services = ArrayField(models.TextField(), default=list)
    """Hosted services the minted license may call."""
    expires_at = UtcDateTimeField(db_column="expiresAt")
    """When the code stops working, which is not the term of what it mints."""
    reusable = models.BooleanField(default=False)
    """A code a customer may redeem on more than one install, for a rollout."""
    redeemed_at = UtcDateTimeField(null=True, db_column="redeemedAt")
    """Set by the redemption that won, which is what makes a single-use code
    single use: the write itself carries the condition.
    """
    redeemed_by_instance_id = models.TextField(null=True, db_column="redeemedByInstanceId")
    issued_license_id = models.TextField(null=True, db_column="issuedLicenseId")
    """The license the winning redemption minted."""
    redemption_count = models.IntegerField(db_column="redemptionCount", default=0)
    """How many installs have redeemed it, which only moves on a reusable code."""
    revoked_at = UtcDateTimeField(null=True, db_column="revokedAt")
    revoked_by_id = models.TextField(null=True, db_column="revokedById")
    created_by_id = models.TextField(null=True, db_column="createdById")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "ActivationCode"
        indexes = [
            models.Index(fields=["organization_id"]),
            models.Index(fields=["expires_at"]),
        ]


class ConnectedBillingAccount(models.Model):
    """How a connected customer is invoiced (ADR-139, section 7): one payment
    provider customer, one quarterly usage subscription in USD, and the
    commercial terms the annual invoice was raised from.

    One row per organization. Every write through it is addressed by a stored
    provider id, so running onboarding or a renewal twice changes nothing.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    organization = models.OneToOneField("Organization", on_delete=models.RESTRICT, related_name="connected_billing", db_column="organizationId")
    stripe_customer_id = models.TextField(unique=True, db_column="stripeCustomerId")
    """The customer at the payment provider. Invoiced, never charged."""
    usage_subscription_id = models.TextField(null=True, db_column="usageSubscriptionId")
    """The one metered subscription, and the item hosted usage is reported on.
    Null between the customer being created and the subscription existing,
    which is where a half-finished onboarding resumes from.
    """
    usage_subscription_item_id = models.TextField(null=True, db_column="usageSubscriptionItemId")
    term_starts_at = UtcDateTimeField(db_column="termStartsAt")
    term_ends_at = UtcDateTimeField(db_column="termEndsAt")
    commit_usd_cents = models.IntegerField(db_column="commitUsdCents", default=0)
    """Prepaid usage for the term, in USD cents. The credit grants add up to it."""
    seat_currency = models.TextField(choices=Currency.choices, db_column="seatCurrency")
    """The seat contract, which may be in another currency than hosted usage."""
    seat_rate_cents = models.IntegerField(db_column="seatRateCents")
    seats = models.IntegerField()
    bank_transfer_type = models.TextField(null=True, db_column="bankTransferType")
    """`us_bank_transfer` or `eu_bank_transfer` when the customer pays into a
    virtual bank account at the provider. Null when it wires to LangWatch's
    own account and finance marks the invoice paid.
    """
    bank_transfer_country = models.TextField(null=True, db_column="bankTransferCountry")
    billing_email = models.TextField(db_column="billingEmail")
    pending_renewal = models.JSONField(null=True, db_column="pendingRenewal")
    """A renewal whose credit is waiting for the last usage invoice of the old
    term to be finalized: the new term's end and commit, and who asked.
    """
    created_by_id = models.TextField(null=True, db_column="createdById")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "ConnectedBillingAccount"


class ConnectedCreditGrant(models.Model):
    """A paid credit at the payment provider, which only metered usage draws down.
    Kept here so a grant is created once per reason and the backoffice can show
    what the customer holds without asking the provider.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    account_id = models.TextField(db_column="accountId")
    """The `ConnectedBillingAccount` row. A plain column ref, like
    `IssuedLicense.replacesId`: relationMode="prisma" carries no foreign keys.
    """
    stripe_credit_grant_id = models.TextField(unique=True, db_column="stripeCreditGrantId")
    """The grant at the payment provider. Unique, so a retry records one row."""
    amount_usd_cents = models.IntegerField(db_column="amountUsdCents")
    kind = models.TextField(choices=ConnectedCreditGrantKind.choices)
    term_ends_at = UtcDateTimeField(db_column="termEndsAt")
    """The term this credit was granted for, and when the provider lets it lapse:
    after the last invoice of that term has been finalized, never on the day
    the term ends.
    """
    expires_at = UtcDateTimeField(db_column="expiresAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "ConnectedCreditGrant"
        indexes = [
            models.Index(fields=["account_id"]),
        ]


class ConnectedInvoice(models.Model):
    """An invoice LangWatch raised for a connected customer, as the backoffice
    reads it.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    account_id = models.TextField(db_column="accountId")
    """The `ConnectedBillingAccount` row, as a plain column ref."""
    stripe_invoice_id = models.TextField(unique=True, db_column="stripeInvoiceId")
    kind = models.TextField(choices=ConnectedInvoiceKind.choices)
    currency = models.TextField(choices=Currency.choices)
    amount_cents = models.IntegerField(db_column="amountCents")
    status = models.TextField()
    """The provider's own status: `draft`, `open`, `paid`, `void` or `uncollectible`."""
    paid_out_of_band_at = UtcDateTimeField(null=True, db_column="paidOutOfBandAt")
    term_starts_at = UtcDateTimeField(null=True, db_column="termStartsAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "ConnectedInvoice"
        indexes = [
            models.Index(fields=["account_id"]),
        ]


class ScimRequestLog(models.Model):
    """Every SCIM request we could attribute, and what we answered (ADR-126).

    EVIDENCE, NOT TRUTH. A request authors nothing — an HTTP round trip that
    was refused decided nothing — so this is a table and not an event: nothing
    downstream derives from it, rebuilding the world from the log must not
    depend on how often a provider retried a GET, and rows age out of a
    retention window an event log would never have.

    NO ROW FOR A REQUEST WE CANNOT ATTRIBUTE. A token that does not verify is
    answered 401 and recorded nowhere: we do not know whose organization to
    file it under, and a table unauthenticated traffic can write is a table
    anybody on the internet can fill. What answers that case instead is
    `ScimToken.lastUsedAt` being null. The 403 IS recorded — a lapsed plan is
    a credential we recognize.

    THE PAYLOAD RULE HOLDS. A method, a resource name, a status, a reason slug
    and our own sentence. Never a token, never a header, never a provider's
    raw message — the same reason `scim_apply_failed` carries a code and not a
    message.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="scim_request_log", db_column="organizationId", db_index=False)
    connection_id = models.TextField(null=True, db_column="connectionId")
    """Null when the refusal resolved an organization and stopped before the
    connection — a lapsed plan is answered off the token alone.
    """
    method = models.TextField()
    resource = models.TextField()
    """The SCIM resource asked for ("Users", "Users/:id", "Groups"), never the
    raw path: a path carries ids and query strings, and this is read by
    people rather than matched by machines.
    """
    status = models.IntegerField()
    reason = models.TextField(null=True)
    """A stable slug, only on a refusal, so a reader branches on this and never
    on the sentence beside it.
    """
    detail = models.TextField(null=True)
    """Our own short sentence, customer-safe."""
    occurred_at = UtcDateTimeField(db_column="occurredAt", default=now_milliseconds)

    class Meta:
        db_table = "ScimRequestLog"
        indexes = [
            models.Index(fields=["organization", "occurred_at"]),
            models.Index(fields=["connection_id", "occurred_at"]),
            models.Index(fields=["occurred_at"]),
        ]


class ScimToken(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="scim_tokens", db_column="organizationId", db_index=False)
    connection_id = models.TextField(null=True, db_column="connectionId")
    """D08: the SSO connection this token was issued for, and its ENTIRE write
    authority — a token minted for one connection can never touch another
    connection's people. No relation, matching `SsoConnection` itself: that
    projection carries no foreign keys, so neither may a row pointing at it.

    Nullable only for the tokens that predate the column. They were minted
    when "which organization" was the whole answer, so they keep the
    organization-wide authority they were sold with; the backfill points
    each at its organization's first connection where there is one. Minting
    a NEW token without a connection is refused (`scim_connection_required`).
    """
    hashed_token = models.TextField(unique=True, db_column="hashedToken")
    """UNIQUE, and that is a tenancy boundary rather than tidiness.

    The lookup that turns a bearer token into an organization is keyed on
    this column alone — it has to be, since the token is the only thing the
    caller presents. While every token was 32 random bytes a collision was
    impossible and the missing constraint cost nothing. An administrator may
    now supply a value they already hold, so two organizations CAN choose the
    same string, and without this the lookup would answer with whichever row
    the planner reached first: one customer's directory provisioning and
    deleting another customer's people.
    """
    hash_scheme = models.TextField(db_column="hashScheme", default="sha256")
    """How `hashedToken` was derived, so the two schemes can coexist while
    tokens minted under the old one are still in use.

    `sha256` is a bare digest, which is correct for 32 random bytes and wrong
    for a value a person chose — a database dump of those is a wordlist away
    from live credentials. `hmac-sha256` is keyed on the deployment's own
    secret, so a dump on its own is inert. Every new row is the latter.
    """
    description = models.TextField(null=True)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    last_used_at = UtcDateTimeField(null=True, db_column="lastUsedAt")

    class Meta:
        db_table = "ScimToken"
        indexes = [
            models.Index(fields=["organization"]),
            models.Index(fields=["connection_id"]),
        ]


class ScimExternalId(models.Model):
    """D08 — who the directory means, per connection.

    The identity provider's own identifier for a person, which is what
    survives their email changing: an address is not identity, people change
    theirs, and the same person carries different identifiers on two different
    connections. So the PAIR is the key and `externalId` alone never is.

    No relation to `SsoConnection` (that projection carries no foreign keys);
    the `User` relation is real, so a deleted user takes their directory
    identities with them rather than stranding a mapping that resolves to
    nobody.
    """

    organization_id = models.TextField(db_column="organizationId")
    id = models.TextField(primary_key=True, default=nanoid_str)
    connection_id = models.TextField(db_column="connectionId")
    external_id = models.TextField(db_column="externalId")
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="scim_external_ids", db_column="userId", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "ScimExternalId"
        constraints = [
            models.UniqueConstraint(fields=["connection_id", "external_id"], name="ScimExternalId_connectionId_externalId_key"),
        ]
        indexes = [
            models.Index(fields=["organization_id", "user"]),
            models.Index(fields=["user"]),
            models.Index(fields=["connection_id"]),
        ]


class ScimUserResource(models.Model):
    """Directory profile and lifecycle are tenant facts, independent of the shared account."""

    pk = models.CompositePrimaryKey("organization_id", "user_id")
    organization_id = models.TextField(db_column="organizationId")
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="scim_user_resources", db_column="userId", db_index=False)
    user_name = models.TextField(db_column="userName")
    name = models.TextField(null=True)
    active = models.BooleanField(default=True)
    deleted_at = UtcDateTimeField(null=True, db_column="deletedAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "ScimUserResource"
        indexes = [
            models.Index(fields=["organization_id", "user_name"]),
            models.Index(fields=["user"]),
        ]


class ScimDirectoryUser(models.Model):
    """A connection's managed users, including resources without an externalId.
    Scalar keys only; readers exclude users that no longer exist.
    """

    pk = models.CompositePrimaryKey("connection_id", "user_id")
    organization_id = models.TextField(db_column="organizationId")
    connection_id = models.TextField(db_column="connectionId")
    user_id = models.TextField(db_column="userId")

    class Meta:
        db_table = "ScimDirectoryUser"
        indexes = [
            models.Index(fields=["organization_id", "user_id"]),
        ]


class ScimSyncState(models.Model):
    """D08 — one connection's directory sync, as the `scim-sync` pipeline's fold
    writes it.

    Pure event truth: every column is fold-written, rows are never deleted,
    and REVOKED is a tombstone rather than an absence — a dead letter that
    vanished when a connection was torn down would be a removal nobody could
    check afterwards.

    No relation to Organization or SsoConnection, matching `SsoConnection`
    itself: a projection is rebuilt from events, and a foreign key would make
    a replay depend on relational state the log does not carry.
    """

    id = models.TextField(primary_key=True, default=None)
    """The aggregate: the connection's id, which IS its sync's id."""
    connection_id = models.TextField(db_column="connectionId")
    organization_id = models.TextField(db_column="organizationId")
    state = models.TextField()
    """TOKEN_ISSUED | SYNCING | ERROR | REVOKED"""
    last_pushed_at = UtcDateTimeField(null=True, db_column="lastPushedAt")
    last_failure = models.JSONField(null=True, db_column="lastFailure")
    """The failure standing right now: { op, errorCode, attempts, retiredAtMs,
    userId, occurredAtMs }. A reason CODE and ids — never a token, never a
    secret, never an internal hostname.
    """
    dead_letters = models.JSONField(db_column="deadLetters", default=list)
    """Every failure retired without ever being applied, newest last. The dead
    letters an administrator has to see; retiring never drops one.
    """
    revoked_cause = models.TextField(null=True, db_column="revokedCause")
    """revoke | teardown, once the sync has ended."""
    occurred_at = UtcDateTimeField(db_column="occurredAt")
    last_event_id = models.TextField(db_column="lastEventId")
    accepted_at = UtcDateTimeField(db_column="acceptedAt")
    projection_version = models.TextField(db_column="projectionVersion")
    created_at = UtcDateTimeField(db_column="createdAt")
    """Event-derived, not database-managed: `@default(now())` / `@updatedAt`
    would make a replayed row differ from the row it rebuilds, which is
    exactly the parity this projection has to keep.
    """
    updated_at = UtcDateTimeField(db_column="updatedAt")

    class Meta:
        db_table = "ScimSyncState"
        indexes = [
            models.Index(fields=["organization_id"]),
            models.Index(fields=["state"]),
            models.Index(fields=["connection_id"]),
        ]


class ScimSsoConnectionView(models.Model):
    """SCIM's own copy of an organization's SSO connections: a peer fold over
    identity's connection facts (ARCHITECTURE §9), folded with identity-contract's
    reducer, one row per connection. SCIM reads this instead of asking identity.
    Rebuilt from identity's event log by a projection replay, so no relation:
    a foreign key would make a replay depend on state the log does not carry.
    """

    id = models.TextField(primary_key=True, default=None)
    """The aggregate: identity's connection id."""
    organization_id = models.TextField(db_column="organizationId")
    folded = models.JSONField()
    """The connection as the reducer folded it (SsoConnectionState + LastEventOccurredAt)."""
    applied_event_ids = ArrayField(models.TextField(), db_column="appliedEventIds", default=list)
    """The event ids of the last delivery folded in, so a redelivery folds once."""
    projection_version = models.TextField(db_column="projectionVersion")
    """A row of another version reads as undecodable and is re-folded from the log."""
    created_at = UtcDateTimeField(db_column="createdAt")
    """Event-derived, so a replayed row equals the row it rebuilds."""
    updated_at = UtcDateTimeField(db_column="updatedAt")

    class Meta:
        db_table = "ScimSsoConnectionView"
        indexes = [
            models.Index(fields=["organization_id", "created_at"]),
        ]


class OrganizationSeatPeak(models.Model):
    """Organization's monthly seat peak, folded from its seat changes (modules/organization, ADR-002).
    A read model, one row per organization: rebuildable by replaying the event log, so no relation.
    Epochs in milliseconds.
    Organization's seat-peak counts (ADR-002): one row per organization once counting opened.
    """

    organization_id = models.TextField(primary_key=True, db_column="organizationId", default=None)
    """The organization, which is also the aggregate id."""
    opened_at = models.FloatField(db_column="openedAt")
    """When the opening snapshot was taken, epoch milliseconds."""
    full_seats = models.IntegerField(db_column="fullSeats")
    lite_seats = models.IntegerField(db_column="liteSeats")
    developer_seats = models.IntegerField(db_column="developerSeats")
    latest_change_at = models.FloatField(db_column="latestChangeAt")
    """The latest change time applied to any holder; invitations expire up to it."""
    mismatches = models.IntegerField()
    """Changes whose `from` disagreed with the seat the fold held."""

    class Meta:
        db_table = "OrganizationSeatPeak"


class OrganizationSeatPeakHolder(models.Model):
    """One member's or pending invitation's seat as the seat peak last applied it (ADR-002)."""

    pk = models.CompositePrimaryKey("organization_id", "holder_kind", "holder_id")
    organization_id = models.TextField(db_column="organizationId")
    holder_kind = models.TextField(db_column="holderKind")
    """`member`, `invitation` or `hold` (a seat held after somebody left)."""
    holder_id = models.TextField(db_column="holderId")
    """The member's or the hold's userId, or the invitation's id."""
    seat = models.TextField()
    """`full`, `lite`, `developer` or `none`."""
    last_changed_at = models.FloatField(db_column="lastChangedAt")
    last_change_id = models.TextField(null=True, db_column="lastChangeId")
    """Absent for a seat the opening snapshot recorded."""
    expires_at = models.FloatField(null=True, db_column="expiresAt")
    """An invitation's expiry; null for a member or an invitation that never expires."""
    expired_by_fold = models.BooleanField(db_column="expiredByFold", default=False)
    """The fold gave the invitation's seat back at `expiresAt`."""
    suppressed_by = models.TextField(null=True, db_column="suppressedBy")
    """A hold whose seat a live full invitation to its address took: that invitation's id."""
    restored_by_fold = models.BooleanField(db_column="restoredByFold", default=False)
    """The fold gave a suppressed hold its seat back when that invitation expired."""

    class Meta:
        db_table = "OrganizationSeatPeakHolder"
        indexes = [
            models.Index(fields=["organization_id", "holder_kind", "seat", "expires_at"]),
            models.Index(fields=["organization_id", "suppressed_by"]),
        ]


class OrganizationSeatPeakMonth(models.Model):
    """One calendar month's (UTC) seat peaks for an organization (ADR-002)."""

    pk = models.CompositePrimaryKey("organization_id", "month")
    organization_id = models.TextField(db_column="organizationId")
    month = models.TextField()
    """`YYYY-MM`."""
    full_peak = models.IntegerField(db_column="fullPeak")
    full_peak_at = models.FloatField(db_column="fullPeakAt")
    lite_peak = models.IntegerField(db_column="litePeak")
    lite_peak_at = models.FloatField(db_column="litePeakAt")
    full_at_end = models.IntegerField(db_column="fullAtEnd")
    lite_at_end = models.IntegerField(db_column="liteAtEnd")

    class Meta:
        db_table = "OrganizationSeatPeakMonth"


class Project(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    name = models.TextField()
    slug = models.TextField(unique=True)
    api_key = models.TextField(unique=True, db_column="apiKey")
    legacy_api_key_last_used_at = UtcDateTimeField(null=True, db_column="legacyApiKeyLastUsedAt")
    """When the legacy apiKey last authenticated a request, written at most once a
    minute per project and process. Spec: modules/api-key/specs/request-credential-type.feature
    """
    lwql_key = models.TextField(unique=True, db_column="lwqlKey", db_default=RawSQL("(gen_random_uuid())::text", []))
    """Random per-project secret whose SHA-256 digest is this project's LangWatchQL
    SQL tenant capability in ClickHouse. Deliberately NOT the API key: the
    credential a caller authenticates with and the capability the database
    resolves a tenant from stay decoupled, so rotating one never invalidates
    the other, and a leak of the ClickHouse key map's input discloses nothing
    a caller could authenticate with. Generated by the database on insert.
    """
    team = models.ForeignKey("Team", on_delete=models.RESTRICT, related_name="projects", db_column="teamId", db_index=False)
    language = models.TextField()
    framework = models.TextField()
    kind = models.TextField(default="application")
    """Project kind:
    - "application": user-facing workspace where developers send their
    own app traces. The default for everything created via the
    onboarding/quickstart flow.
    - "internal_governance": hidden per-org routing/tenancy artifact
    for IngestionSource ingest data. Auto-created lazily on the
    first IngestionSource mint via ensureHiddenGovernanceProject().
    MUST be filtered out of every user-visible Project surface
    (project picker, /api/v1/projects, billing exports, RBAC
    pickers, etc.) - the regression-test invariant in
    specs/ai-gateway/governance/ui-contract.feature codifies this.
    - "aggregate": a project that owns no traces and reads its member
    projects through shared project-reader grants (ADR-177). Hidden
    from every "send traces here" picker and from ingestion; visible in
    the project switcher to organisation admins only.
    Stored as a free-form string (matching IngestionSource.sourceType
    pattern) so future routing kinds can ship without schema migrations.
    """
    aggregate_rule = models.JSONField(null=True, db_column="aggregateRule")
    """The scope rule an aggregate project materialises into grants. Null on
    every other kind; decided in ADR-177. Shapes: { kind: "all-personal" } |
    { kind: "personal-by-department", departmentId } |
    { kind: "explicit", projectIds: string[] }.
    """
    first_message = models.BooleanField(db_column="firstMessage", default=False)
    integrated = models.BooleanField(default=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    user_link_template = models.TextField(null=True, db_column="userLinkTemplate")
    trace_sharing_enabled = models.BooleanField(db_column="traceSharingEnabled", default=True)
    presence_enabled = models.BooleanField(db_column="presenceEnabled", default=True)
    s3_endpoint = models.TextField(null=True, db_column="s3Endpoint")
    s3_access_key_id = models.TextField(null=True, db_column="s3AccessKeyId")
    s3_secret_access_key = models.TextField(null=True, db_column="s3SecretAccessKey")
    s3_bucket = models.TextField(null=True, db_column="s3Bucket")
    archived_at = UtcDateTimeField(null=True, db_column="archivedAt")
    is_personal = models.BooleanField(db_column="isPersonal", default=False)
    owner_user = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="personal_projects", db_column="ownerUserId", null=True, db_index=False)
    personal_features = models.JSONField(db_column="personalFeatures", default=dict)
    """Per-feature unlock flags for personal projects only. Shape:
    `{ evaluations: bool, datasets: bool, annotations: bool, automations: bool }`.
    Empty `{}` (default) reads as all-false at the service layer -
    existing rows pre-migration require no backfill. The bundle is
    a UI/nav predicate, NOT an auth gate (the underlying tRPC
    routers stay open even when the bundle is off; admin tooling +
    offline migration scripts can still operate on the data).
    Team / org-scope projects are not feature-gated - this column
    is meaningful only when `isPersonal=true`.
    """
    department_id = models.TextField(null=True, db_column="departmentId")
    """Department assignment for spend attribution (the project / agent
    lens, used when a trace has no principal user). Accounting-only.
    """
    langy_egress_allowlist = models.JSONField(null=True, db_column="langyEgressAllowlist")
    """Per-project Langy egress allow-list (ADR-076). A nullable JSON array of
    host patterns (`["registry.npmjs.org", "*.internal.acme.com"]`) that the
    project's Langy workers may reach outbound. The *presence* of the list is
    the enforcement mode: null/absent ⇒ monitor-only (watch, never block);
    non-empty ⇒ the worker's egress adapter restricts outbound to
    floor ∪ this list. Enforced in the agent pod's egress adapter — NOT the
    gateway — so it lives on the project, not the VirtualKey. Resolved by
    LangyCredentialService.getEgressAllowlist and threaded through the
    per-request credentials envelope. See specs/langy/langy-egress-enforcement.feature.
    """
    last_coding_agent_session_at = UtcDateTimeField(null=True, db_column="lastCodingAgentSessionAt")
    """When a coding-agent session was last folded for this project, and when a
    pull request was last mapped for a branch one of those sessions ran on.
    Both are instrumentation signals, not counters: the project sidebar shows
    its Sessions and Pull requests destinations only while the matching column
    is inside the recency window, so a project that never sends coding-agent
    telemetry never grows the two links, and one that stops sending loses
    them again. Written through the throttled touch methods on
    ProjectService, which skip the write while the stored value is still
    recent, so a busy project costs at most one update an hour per column.
    No index: every read is by project id, which is the primary key.
    See specs/coding-agent/project-menu-links.feature.
    """
    last_coding_agent_pull_request_at = UtcDateTimeField(null=True, db_column="lastCodingAgentPullRequestAt")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Project"
        indexes = [
            models.Index(fields=["team"]),
            models.Index(fields=["is_personal"]),
            models.Index(fields=["owner_user"]),
            models.Index(fields=["team", "kind"]),
            models.Index(fields=["department_id"]),
        ]


class Department(models.Model):
    """A department is a single-valued accounting dimension, separate from
    RBAC, for rolling up AI spend by org unit (Engineering, Marketing,
    ...). Assigned to OrganizationUser, Team, and Project via their
    `departmentId` column. Pure accounting: never grants or restricts
    access. See specs/ai-gateway/governance/departments.feature.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    organization_id = models.TextField(db_column="organizationId")
    name = models.TextField()
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    archived_at = UtcDateTimeField(null=True, db_column="archivedAt")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Department"
        indexes = [
            models.Index(fields=["organization_id"]),
        ]


class DepartmentMembershipHistory(models.Model):
    """One dated person → department link (ADR-128 §13, #7882).

    `OrganizationUser.departmentId` stays the CURRENT assignment every screen
    reads; these rows are its history, appended by the same `assignUser` write
    so the two cannot drift. They exist for the cost reads: January's spend
    stays with January's department, so a read resolves against the link that
    was open on the cost's day, not against today's pointer. Members with no
    link on a day resolve to "unassigned", never to an error.

    Same shape as IdentityMatch's dated links, same one-open-link guard: the
    partial unique index in the migration rejects a second open row per
    (organizationId, userId) with SQLSTATE 23505.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    organization_id = models.TextField(db_column="organizationId")
    user_id = models.TextField(db_column="userId")
    department_id = models.TextField(db_column="departmentId")
    valid_from = UtcDateTimeField(db_column="validFrom")
    valid_to = UtcDateTimeField(null=True, db_column="validTo")
    """Null means the link is open. Reassignment and clearing close it."""
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "DepartmentMembershipHistory"
        indexes = [
            models.Index(fields=["organization_id", "user_id", "valid_from"]),
        ]


class PlatformToolPolicy(models.Model):
    """Per-(organization, tool) override for which `langwatch <tool>` paths the
    CLI may use: `allowVk` gates the gateway / virtual-key path (Path A) and
    `allowOtelDirect` gates the direct-OTLP ingestion path (Path B). A missing
    row resolves to the hardcoded defaults (both true for claude/codex/gemini/
    opencode; cursor allows the gateway path only), so the policy is purely
    additive: disabling a path is an explicit admin choice. The CLI caches the
    resolved map at login and gates path selection on it.
    See specs/ai-gateway/governance/cli-tool-mode-policy.feature.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    organization_id = models.TextField(db_column="organizationId")
    tool_slug = models.TextField(db_column="toolSlug")
    allow_vk = models.BooleanField(db_column="allowVk", default=True)
    allow_otel_direct = models.BooleanField(db_column="allowOtelDirect", default=True)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "PlatformToolPolicy"
        constraints = [
            models.UniqueConstraint(fields=["organization_id", "tool_slug"], name="PlatformToolPolicy_organizationId_toolSlug_key"),
        ]
        indexes = [
            models.Index(fields=["organization_id"]),
        ]


class OrganizationInvite(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    email = models.TextField()
    invite_code = models.TextField(unique=True, db_column="inviteCode")
    expiration = UtcDateTimeField(null=True)
    status = models.TextField(choices=INVITESTATUS.choices, default=INVITESTATUS.PENDING)
    organization = models.ForeignKey("Organization", on_delete=models.RESTRICT, related_name="organization_invite", db_column="organizationId", db_index=False)
    team_ids = models.TextField(db_column="teamIds")
    team_assignments = models.JSONField(null=True, db_column="teamAssignments")
    role = models.TextField(choices=OrganizationUserRole.choices)
    requested_by_user = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="invite_requests", db_column="requestedBy", null=True, db_index=False)
    subscription_id = models.TextField(null=True, db_column="subscriptionId")
    accepted_by_user_id = models.TextField(null=True, db_column="acceptedByUserId")
    accepted_via_identifier_id = models.TextField(null=True, db_column="acceptedViaIdentifierId")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "OrganizationInvite"
        indexes = [
            models.Index(fields=["organization"]),
            models.Index(fields=["requested_by_user"]),
        ]


class Monitor(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    project = models.ForeignKey("Project", on_delete=models.RESTRICT, related_name="checks", db_column="projectId", db_index=False)
    experiment = models.OneToOneField("Experiment", on_delete=models.RESTRICT, related_name="monitor", db_column="experimentId", null=True)
    evaluator = models.ForeignKey("Evaluator", on_delete=models.SET_NULL, related_name="monitors", db_column="evaluatorId", null=True, db_index=False)
    check_type = models.TextField(db_column="checkType")
    name = models.TextField()
    slug = models.TextField()
    execution_mode = models.TextField(choices=EvaluationExecutionMode.choices, db_column="executionMode", default=EvaluationExecutionMode.ON_MESSAGE)
    enabled = models.BooleanField(default=True)
    preconditions = models.JSONField()
    parameters = models.JSONField()
    mappings = models.JSONField(null=True)
    sample = models.FloatField(default=1)
    level = models.TextField(default="trace")
    thread_idle_timeout = models.IntegerField(null=True, db_column="threadIdleTimeout")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    auto_paused_at = UtcDateTimeField(null=True, db_column="autoPausedAt")
    auto_paused_reason = models.TextField(null=True, db_column="autoPausedReason")
    evaluator_label = models.TextField(null=True, db_column="evaluatorLabel")
    created_by_id = models.TextField(null=True, db_column="createdById")
    updated_by_id = models.TextField(null=True, db_column="updatedById")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Monitor"
        constraints = [
            models.UniqueConstraint(fields=["project", "slug"], name="Monitor_projectId_slug_key"),
        ]
        indexes = [
            models.Index(fields=["project"]),
            models.Index(fields=["evaluator"]),
        ]


class Cost(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    project = models.ForeignKey("Project", on_delete=models.RESTRICT, related_name="costs", db_column="projectId", db_index=False)
    cost_type = models.TextField(choices=CostType.choices, db_column="costType")
    cost_name = models.TextField(null=True, db_column="costName")
    reference_type = models.TextField(choices=CostReferenceType.choices, db_column="referenceType")
    reference_id = models.TextField(db_column="referenceId")
    amount = models.FloatField()
    currency = models.TextField()
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    extra_info = models.JSONField(null=True, db_column="extraInfo")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Cost"
        indexes = [
            models.Index(fields=["reference_type", "reference_id"]),
            models.Index(fields=["cost_type"]),
            models.Index(fields=["project"]),
            models.Index(fields=["project", "created_at"]),
        ]


class Topic(models.Model):
    """PROJECTION of the topic-clustering event stream (ADR-051): topics are
    `topics_recorded` events; this table is written ONLY by
    TopicModelFoldProjection (cursor: TopicModelProjection) and is
    rebuildable by replay. Kept named `Topic` because it is the domain
    entity every surface reads — the `*Projection` suffix is reserved for
    cursor/read-model tables.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    project = models.ForeignKey("Project", on_delete=models.RESTRICT, related_name="topics", db_column="projectId", db_index=False)
    name = models.TextField()
    parent = models.ForeignKey("self", on_delete=models.RESTRICT, related_name="subtopics", db_column="parentId", null=True, db_index=False)
    embeddings_model = models.TextField()
    centroid = models.JSONField()
    p95_distance = models.FloatField(db_column="p95Distance")
    automatically_generated = models.BooleanField(db_column="automaticallyGenerated", default=True)
    last_event_id = models.TextField(null=True, db_column="lastEventId")
    """The `topics_recorded` event that recorded this topic (its id in the
    event log). Null means the row predates event ownership and has not
    been seeded onto the stream yet; the projection sets it on every row
    it writes.
    """
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Topic"
        indexes = [
            models.Index(fields=["parent"]),
            models.Index(fields=["project"]),
        ]


class Dataset(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    project = models.ForeignKey("Project", on_delete=models.RESTRICT, related_name="datasets", db_column="projectId", db_index=False)
    name = models.TextField()
    slug = models.TextField()
    column_types = models.JSONField(db_column="columnTypes")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    archived_at = UtcDateTimeField(null=True, db_column="archivedAt")
    mapping = models.JSONField(null=True)
    use_s3 = models.BooleanField(db_column="useS3", default=False)
    s3_record_count = models.IntegerField(null=True, db_column="s3RecordCount")
    content_layout = models.TextField(db_column="contentLayout", default="postgres")
    status = models.TextField(default="ready")
    status_error = models.TextField(null=True, db_column="statusError")
    staging_key = models.TextField(null=True, db_column="stagingKey")
    source_stored_object_id = models.TextField(null=True, db_column="sourceStoredObjectId")
    upload_filename = models.TextField(null=True, db_column="uploadFilename")
    row_count = models.IntegerField(null=True, db_column="rowCount")
    size_bytes = models.BigIntegerField(null=True, db_column="sizeBytes")
    chunk_count = models.IntegerField(null=True, db_column="chunkCount")
    chunk_offsets = models.JSONField(null=True, db_column="chunkOffsets")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Dataset"
        constraints = [
            models.UniqueConstraint(fields=["project", "slug"], name="Dataset_projectId_slug_key"),
        ]


class DatasetRecord(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    dataset = models.ForeignKey("Dataset", on_delete=models.RESTRICT, related_name="dataset_records", db_column="datasetId", db_index=False)
    project = models.ForeignKey("Project", on_delete=models.RESTRICT, related_name="dataset_records", db_column="projectId", db_index=False)
    entry = models.JSONField()
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "DatasetRecord"
        indexes = [
            models.Index(fields=["dataset"]),
            models.Index(fields=["project"]),
            models.Index(fields=["created_at"]),
            models.Index(fields=["project", "dataset"]),
        ]


class DatasetVersion(models.Model):
    """One immutable snapshot of a dataset's columns and rows; the rows are JSONL chunks in object storage."""

    id = models.TextField(primary_key=True, default=None)
    project_id = models.TextField(db_column="projectId")
    dataset_id = models.TextField(db_column="datasetId")
    version = models.IntegerField()
    name = models.TextField(null=True)
    """A name given by hand; a named version is never thinned by retention."""
    trigger = models.TextField()
    """"autosave", "manual", "experiment_run", "restore" or "initial"."""
    content_hash = models.TextField(db_column="contentHash")
    """SHA-256 over the columns and the rows in order; equal content saves no new version."""
    column_types = models.JSONField(db_column="columnTypes")
    row_count = models.IntegerField(db_column="rowCount")
    size_bytes = models.BigIntegerField(db_column="sizeBytes")
    chunk_count = models.IntegerField(db_column="chunkCount")
    restored_from_version = models.IntegerField(null=True, db_column="restoredFromVersion")
    author_id = models.TextField(null=True, db_column="authorId")
    author_kind = models.TextField(null=True, db_column="authorKind")
    """"user", "api_key" or "system": who saved the version when authorId is null."""
    pinned_at = UtcDateTimeField(null=True, db_column="pinnedAt")
    """Set when an experiment pins this version; a pinned version is never thinned."""
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "DatasetVersion"
        constraints = [
            models.UniqueConstraint(fields=["dataset_id", "version"], name="DatasetVersion_datasetId_version_key"),
        ]
        indexes = [
            models.Index(fields=["project_id", "dataset_id"]),
        ]


class DatasetVersionRun(models.Model):
    """An experiment run that read one dataset version; a version a run read is never thinned."""

    id = models.TextField(primary_key=True, default=None)
    project_id = models.TextField(db_column="projectId")
    dataset_id = models.TextField(db_column="datasetId")
    version_id = models.TextField(db_column="versionId")
    run_id = models.TextField(db_column="runId")
    experiment_id = models.TextField(null=True, db_column="experimentId")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "DatasetVersionRun"
        constraints = [
            models.UniqueConstraint(fields=["project_id", "run_id", "dataset_id"], name="DatasetVersionRun_projectId_runId_datasetId_key"),
        ]
        indexes = [
            models.Index(fields=["version_id"]),
        ]


class Dashboard(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    project = models.ForeignKey("Project", on_delete=models.RESTRICT, related_name="dashboards", db_column="projectId", db_index=False)
    """The project that owns the board; only there can it be edited."""
    name = models.TextField()
    order = models.IntegerField(default=0)
    """Legacy board sort position. The Dashboards v2 app orders boards by a
    member's favourites and never writes this; kept because the legacy
    Analytics reports section and the REST reorder endpoint still write it,
    and a rolling deploy's older image still reads it.
    """
    description = models.TextField(null=True)
    """Free text under the board's name; null until a member writes one.
    Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
    """
    created_by_id = models.TextField(null=True, db_column="createdById")
    """The member who created it; null for boards that predate the column.
    Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
    """
    scope = models.TextField(choices=DashboardScope.choices, default=DashboardScope.PROJECT)
    """Who sees it. Only the author changes it; a board with no author stays PROJECT.
    Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
    """
    organization_id = models.TextField(null=True, db_column="organizationId")
    """The owning project's organization, stamped when the board is first set to
    ORGANIZATION; null before. No relation, per the database rule.
    Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
    """
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Dashboard"
        indexes = [
            models.Index(fields=["project"]),
            models.Index(fields=["project", "order"]),
            models.Index(fields=["organization_id", "scope"]),
        ]


class DashboardFavourite(models.Model):
    """A member's star, in their own order within a project: a board or a From
    LangWatch template. Personal: one member's stars never show to another. No
    foreign key (per the database rule): the dashboard service removes a board's
    favourites when it deletes the board, so a star never outlives its dashboard.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    user_id = models.TextField(db_column="userId")
    dashboard_id = models.TextField(null=True, db_column="dashboardId")
    """A DashboardsV2 board's id (copied boards kept their Dashboard id); null when
    the star is a template.
    """
    template_id = models.TextField(null=True, db_column="templateId")
    """The starred From LangWatch template id; null when the star is a board."""
    project_id = models.TextField(db_column="projectId")
    position = models.IntegerField()
    """The member's drag order among their own stars in the project."""
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "DashboardFavourite"
        constraints = [
            models.UniqueConstraint(fields=["user_id", "dashboard_id"], name="DashboardFavourite_userId_dashboardId_key"),
            models.UniqueConstraint(fields=["user_id", "project_id", "template_id"], name="DashboardFavourite_userId_projectId_templateId_key"),
        ]
        indexes = [
            models.Index(fields=["user_id", "project_id", "position"]),
            models.Index(fields=["dashboard_id"]),
        ]


class DashboardsV2(models.Model):
    """A Dashboards v2 board, apart from the Analytics Reports boards in Dashboard. Boards
    made before the split were copied here with their ids, so stars and insight
    pointers still name them. No relation, per the database rule.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    project_id = models.TextField(db_column="projectId")
    """The project that owns the board; only there can it be edited."""
    name = models.TextField()
    description = models.TextField(null=True)
    """Free text under the board's name; null until a member writes one."""
    created_by_id = models.TextField(null=True, db_column="createdById")
    """The member who created it; null for boards that predate authorship."""
    scope = models.TextField(choices=DashboardScope.choices, default=DashboardScope.PROJECT)
    """Who sees it. Only the author changes it; a board with no author stays PROJECT."""
    organization_id = models.TextField(null=True, db_column="organizationId")
    """The owning project's organization, stamped when the board is first set to
    ORGANIZATION; null before.
    """
    excluded_origins = models.JSONField(null=True, db_column="excludedOrigins")
    """Deprecated: unused since each widget query names its own sources (dashboards-v2 AC200,
    2026-10-28); dropped in a later migration, once no serving release reads it.
    """
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "DashboardsV2"
        indexes = [
            models.Index(fields=["project_id"]),
            models.Index(fields=["organization_id", "scope"]),
        ]


class DashboardWidget(models.Model):
    """A Dashboards v2 widget: a srcdoc definition placed on a DashboardsV2 board, or
    unplaced. Widgets made before the split were copied from CustomGraph with their
    ids. No relation, per the database rule.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    project_id = models.TextField(db_column="projectId")
    dashboard_id = models.TextField(null=True, db_column="dashboardId")
    """The DashboardsV2 board it sits on; null while unplaced."""
    name = models.TextField()
    definition = models.JSONField()
    """The widget's srcdoc definition."""
    grid_column = models.IntegerField(db_column="gridColumn", default=0)
    grid_row = models.IntegerField(db_column="gridRow", default=0)
    col_span = models.IntegerField(db_column="colSpan", default=1)
    row_span = models.IntegerField(db_column="rowSpan", default=1)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "DashboardWidget"
        indexes = [
            models.Index(fields=["project_id", "dashboard_id"]),
        ]


class InsightProjection(models.Model):
    """One insight in its owner's inbox in a project, folded from its `insight` aggregate
    (modules/insight). A read model: rebuildable by replaying the event log. Epochs in
    milliseconds.
    """

    id = models.TextField(primary_key=True, default=None)
    """The insight id, which is also its aggregate id."""
    project_id = models.TextField(db_column="projectId")
    title = models.TextField()
    body = models.TextField()
    tone = models.TextField()
    """bad | watch | good"""
    topic = models.TextField(null=True)
    valid_days = models.IntegerField(db_column="validDays")
    lwql = models.TextField(null=True)
    source_conversation_id = models.TextField(null=True, db_column="sourceConversationId")
    source_message_id = models.TextField(null=True, db_column="sourceMessageId")
    filed_by_user_id = models.TextField(null=True, db_column="filedByUserId")
    filed_at = models.FloatField(db_column="filedAt")
    renewed_at = models.FloatField(null=True, db_column="renewedAt")
    """Set by a later run that found the insight still true."""
    created_at = models.FloatField(db_column="createdAt")
    updated_at = models.FloatField(db_column="updatedAt")
    occurred_at = models.FloatField(db_column="occurredAt")
    accepted_at = models.FloatField(db_column="acceptedAt")
    last_event_id = models.TextField(db_column="lastEventId")
    projection_version = models.TextField(db_column="projectionVersion")
    board_id = models.TextField(null=True, db_column="boardId")
    """The board and widget it came from. A pointer only, never a relation: the row reads the
    same once either is deleted, so the names are kept as filed.
    """
    board_name = models.TextField(null=True, db_column="boardName")
    widget_id = models.TextField(null=True, db_column="widgetId")
    widget_name = models.TextField(null=True, db_column="widgetName")
    filed_via = models.TextField(db_column="filedVia", default="chat")
    """chat | run"""
    replay_start = models.FloatField(null=True, db_column="replayStart")
    """The fixed window `lwql` replays over (epoch milliseconds, half-open) and its step."""
    replay_end = models.FloatField(null=True, db_column="replayEnd")
    replay_granularity_seconds = models.IntegerField(null=True, db_column="replayGranularitySeconds")
    replay_context = models.JSONField(null=True, db_column="replayContext")
    """The values in force when it was filed: `{ period, parameters }`."""
    owner_user_id = models.TextField(null=True, db_column="ownerUserId")
    """Whose insight it is: the one person who reads it. NULL on a row folded before the
    column existed, which is owned by whoever filed it (`filedByUserId`).
    """
    finding_key = models.TextField(null=True, db_column="findingKey")
    """The stable key a daily run matches a repeated finding by, so it renews the insight
    and files no twin. NULL for an insight saved from a chat and on a row folded before
    the column existed.
    """

    class Meta:
        db_table = "InsightProjection"
        indexes = [
            models.Index(fields=["project_id", "filed_at"]),
            models.Index(fields=["project_id", "owner_user_id", "filed_at"]),
        ]


class InsightReaderProjection(models.Model):
    """One reader's own state on one insight (seen, done, kept), folded from the same stream.
    An insight has one reader today, its owner; a later forward may add more.
    """

    id = models.TextField(primary_key=True, default=None)
    project_id = models.TextField(db_column="projectId")
    insight_id = models.TextField(db_column="insightId")
    user_id = models.TextField(db_column="userId")
    seen_at = models.FloatField(null=True, db_column="seenAt")
    archived_at = models.FloatField(null=True, db_column="archivedAt")
    kept_at = models.FloatField(null=True, db_column="keptAt")
    created_at = models.FloatField(db_column="createdAt")
    updated_at = models.FloatField(db_column="updatedAt")
    occurred_at = models.FloatField(db_column="occurredAt")
    accepted_at = models.FloatField(db_column="acceptedAt")
    last_event_id = models.TextField(db_column="lastEventId")
    projection_version = models.TextField(db_column="projectionVersion")

    class Meta:
        db_table = "InsightReaderProjection"
        constraints = [
            models.UniqueConstraint(fields=["project_id", "insight_id", "user_id"], name="InsightReaderProjection_projectId_insightId_userId_key"),
        ]
        indexes = [
            models.Index(fields=["project_id", "user_id"]),
        ]


class InsightDailyScheduleProjection(models.Model):
    """One person's daily insights run on one board in a project, folded from its
    `insight_daily_schedule` aggregate (modules/insight). A read model: rebuildable by
    replaying the event log. Epochs in milliseconds.
    """

    id = models.TextField(primary_key=True, default=None)
    """The schedule id, which is also its aggregate id: one per project, person and board."""
    project_id = models.TextField(db_column="projectId")
    user_id = models.TextField(db_column="userId")
    board_kind = models.TextField(db_column="boardKind")
    """dashboard | template. A pointer only, never a relation: a run checks the board is there."""
    board_id = models.TextField(db_column="boardId")
    board_name = models.TextField(db_column="boardName")
    state = models.TextField(default="undecided")
    """on | off | undecided."""
    off_reason = models.TextField(null=True, db_column="offReason")
    """board_deleted | skipped_days | member_removed | project_archived. Why the system turned
    it off. NULL while it is on, and when the person turned it off.
    """
    hour = models.IntegerField(null=True)
    timezone = models.TextField(null=True)
    max_insights = models.IntegerField(null=True, db_column="maxInsights")
    last_run_id = models.TextField(null=True, db_column="lastRunId")
    """The last run that recorded an outcome. NULL until one did."""
    last_run_at = models.FloatField(null=True, db_column="lastRunAt")
    last_run_outcome = models.TextField(null=True, db_column="lastRunOutcome")
    """filed | nothing | failed | skipped"""
    last_run_reason = models.TextField(null=True, db_column="lastRunReason")
    last_run_filed = models.IntegerField(null=True, db_column="lastRunFiled")
    last_run_renewed = models.IntegerField(null=True, db_column="lastRunRenewed")
    """How many insights the run found still true and renewed."""
    last_run_conversation_id = models.TextField(null=True, db_column="lastRunConversationId")
    consecutive_skips = models.IntegerField(db_column="consecutiveSkips", default=0)
    """The scheduled days in a row whose run ended skipped, and the slot of the newest counted."""
    skip_streak_slot = models.FloatField(null=True, db_column="skipStreakSlot")
    created_at = models.FloatField(db_column="createdAt")
    updated_at = models.FloatField(db_column="updatedAt")
    occurred_at = models.FloatField(db_column="occurredAt")
    accepted_at = models.FloatField(db_column="acceptedAt")
    last_event_id = models.TextField(db_column="lastEventId")
    projection_version = models.TextField(db_column="projectionVersion")

    class Meta:
        db_table = "InsightDailyScheduleProjection"
        constraints = [
            models.UniqueConstraint(fields=["project_id", "user_id", "board_kind", "board_id"], name="InsightDailySchedule_projectId_userId_board_key"),
        ]


class SavedView(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    project = models.ForeignKey("Project", on_delete=models.RESTRICT, related_name="saved_views", db_column="projectId", db_index=False)
    user = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="saved_views", db_column="userId", null=True, db_index=False)
    name = models.TextField()
    filters = models.JSONField(default=dict)
    query = models.TextField(null=True)
    period = models.JSONField(null=True)
    order = models.IntegerField(default=0)
    kind = models.TextField(default="v1-traces-filter")
    """Discriminator for the two storage shapes that share this table:
    - "v1-traces-filter" (default) - legacy /messages page filter map saved views
    - "v2-traces-lens"             - new traces v2 lens configs (columns, addons, grouping, sort, filterText) encoded inside `filters` JSON
    Each client filters by its own kind on read so neither sees the other's rows.
    """
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "SavedView"
        indexes = [
            models.Index(fields=["project"]),
            models.Index(fields=["project", "order"]),
            models.Index(fields=["user"]),
            models.Index(fields=["project", "kind", "order"]),
        ]


class CustomGraph(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    project = models.ForeignKey("Project", on_delete=models.RESTRICT, related_name="custom_graphs", db_column="projectId", db_index=False)
    name = models.TextField()
    graph = models.JSONField()
    """The chart's definition, in the shape its `kind` names."""
    filters = models.JSONField(null=True)
    kind = models.TextField(default="builder")
    """Discriminator for the chart shapes that share this table:
    - "builder" (default) - the chart builder's own `graph` payload
    - "workbench_sql"     - a saved LangWatchQL workbench chart, whose
    `graph` holds the versioned { sql, parameters, vegaLiteSpec }
    definition (see server/analytics/saved-workbench-charts)
    - "dashboard_srcdoc"  - frozen legacy copies of Dashboards v2 widgets, which
    now live in DashboardWidget; nothing new writes this kind
    Each reader filters by its own kind, so neither sees the other's rows and
    a builder chart can never be interpreted as a workbench definition.
    """
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    dashboard = models.ForeignKey("Dashboard", on_delete=models.CASCADE, related_name="graphs", db_column="dashboardId", null=True, db_index=False)
    grid_column = models.IntegerField(db_column="gridColumn", default=0)
    grid_row = models.IntegerField(db_column="gridRow", default=0)
    col_span = models.IntegerField(db_column="colSpan", default=1)
    row_span = models.IntegerField(db_column="rowSpan", default=1)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "CustomGraph"
        indexes = [
            models.Index(fields=["project"]),
            models.Index(fields=["dashboard"]),
            models.Index(fields=["project", "kind"]),
        ]


class BatchEvaluation(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    experiment = models.ForeignKey("Experiment", on_delete=models.RESTRICT, related_name="batch_evaluations", db_column="experimentId", db_index=False)
    project = models.ForeignKey("Project", on_delete=models.RESTRICT, related_name="batch_evaluations", db_column="projectId", db_index=False)
    data = models.JSONField(default=dict)
    status = models.TextField()
    score = models.FloatField()
    label = models.TextField(null=True)
    passed = models.BooleanField()
    details = models.TextField()
    cost = models.FloatField()
    dataset_slug = models.TextField(db_column="datasetSlug")
    dataset = models.ForeignKey("Dataset", on_delete=models.RESTRICT, related_name="batch_evaluations", db_column="datasetId", db_index=False)
    evaluation = models.TextField()
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "BatchEvaluation"
        constraints = [
            models.UniqueConstraint(fields=["id"], name="BatchEvaluation_id_key"),
        ]
        indexes = [
            models.Index(fields=["project"]),
            models.Index(fields=["dataset"]),
            models.Index(fields=["experiment"]),
        ]


class Trigger(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    name = models.TextField()
    project = models.ForeignKey("Project", on_delete=models.RESTRICT, related_name="triggers", db_column="projectId", db_index=False)
    action = models.TextField(choices=TriggerAction.choices)
    trigger_kind = models.TextField(choices=TriggerKind.choices, db_column="triggerKind", default=TriggerKind.AUTOMATION)
    action_params = models.JSONField(db_column="actionParams")
    filters = models.JSONField()
    filter_query = models.TextField(null=True, db_column="filterQuery")
    last_run_at = models.FloatField(db_column="lastRunAt", default=0)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    active = models.BooleanField(default=True)
    paused_reason = models.TextField(null=True, db_column="pausedReason")
    paused_at = models.DateTimeField(null=True, db_column="pausedAt")
    message = models.TextField(null=True)
    deleted = models.BooleanField(default=False)
    alert_type = models.TextField(choices=AlertType.choices, null=True, db_column="alertType")
    slack_template_type = models.TextField(null=True, db_column="slackTemplateType")
    slack_template = models.TextField(null=True, db_column="slackTemplate")
    email_subject_template = models.TextField(null=True, db_column="emailSubjectTemplate")
    email_body_template = models.TextField(null=True, db_column="emailBodyTemplate")
    notification_cadence = models.TextField(db_column="notificationCadence", default="immediate")
    trace_debounce_ms = models.IntegerField(db_column="traceDebounceMs", default=30000)
    custom_graph = models.OneToOneField("CustomGraph", on_delete=models.CASCADE, related_name="trigger", db_column="customGraphId", null=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Trigger"
        indexes = [
            models.Index(fields=["project"]),
        ]


class SlackIntegration(models.Model):
    """One named Slack connection (ADR-093 §5a); an organization holds as many as
    it likes, each scoped to the organization or to one project. The secret is
    AES-256-GCM ciphertext and never returned to a client. `secretFingerprint`
    is an HMAC of the plaintext, so one secret is one connection per scope
    without decrypting anything. No foreign key: `scopeId` is
    polymorphic (ADR-021), as on RetentionPolicy and CustomLLMModelCost.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    name = models.TextField()
    kind = models.TextField(choices=SlackIntegrationKind.choices)
    scope_type = models.TextField(choices=SlackIntegrationScopeType.choices, db_column="scopeType")
    scope_id = models.TextField(db_column="scopeId")
    organization_id = models.TextField(db_column="organizationId")
    bot_token_encrypted = models.TextField(null=True, db_column="botTokenEncrypted")
    webhook_url_encrypted = models.TextField(null=True, db_column="webhookUrlEncrypted")
    secret_fingerprint = models.TextField(db_column="secretFingerprint")
    secret_hint = models.TextField(db_column="secretHint")
    slack_team_id = models.TextField(null=True, db_column="slackTeamId")
    slack_team_name = models.TextField(null=True, db_column="slackTeamName")
    created_by_id = models.TextField(db_column="createdById")
    updated_by_id = models.TextField(db_column="updatedById")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "SlackIntegration"
        constraints = [
            models.UniqueConstraint(fields=["organization_id", "scope_type", "scope_id", "secret_fingerprint"], name="SlackIntegration_organizationId_scopeType_scopeId_secretFingerprint_key"),
        ]
        indexes = [
            models.Index(fields=["scope_type", "scope_id"]),
        ]


class SlackConnectionClaim(models.Model):
    """An automation holding a Slack connection in use (ARCHITECTURE.md §3, the
    claims ruling): a claimed connection is never deleted. `projectId` is the
    claimant's project; `organizationId` is the connection's tenancy anchor.
    No foreign key: slack's service refuses to delete a claimed connection.
    """

    pk = models.CompositePrimaryKey("connection_id", "claimant_id")
    connection_id = models.TextField(db_column="connectionId")
    claimant_id = models.TextField(db_column="claimantId")
    claimant_label = models.TextField(db_column="claimantLabel")
    organization_id = models.TextField(db_column="organizationId")
    project_id = models.TextField(db_column="projectId")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "slack_connection_claim"
        indexes = [
            models.Index(fields=["organization_id"]),
        ]


class WebhookEndpoint(models.Model):
    """Org-anchored outbound webhook endpoint (the webhook platform, enterprise
    only). The signing secret is AES-256-GCM encrypted at rest (shared
    encrypt/decrypt, CREDENTIALS_SECRET) and returned in plaintext exactly
    once, at create or roll time. `enabledEvents` holds registry event types
    or family wildcards ("gateway.*", "*"). Failure-streak columns drive the
    72h auto-disable. Spec: specs/webhooks/webhook-endpoints.feature.
    """

    id = models.TextField(primary_key=True, default=None)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="webhook_endpoints", db_column="organizationId", db_index=False)
    destination_kind = models.TextField(choices=WebhookDestinationKind.choices, db_column="destinationKind", default=WebhookDestinationKind.http)
    """Where this endpoint delivers. Everything above the last hop is shared:
    same batching, same retry ladder, same delivery log, same signature over
    the same bytes.
    """
    url = models.TextField(null=True)
    """The receiver URL. Present on `http` endpoints and null on the others;
    the per-kind shape is pinned by a CHECK constraint Prisma cannot model.
    """
    sqs_queue_url = models.TextField(null=True, db_column="sqsQueueUrl")
    """The Amazon SQS queue this endpoint delivers to, on `sqs` endpoints. The
    region and the owning account are read off it rather than configured
    separately, so they can never disagree with the queue.
    """
    sqs_role_arn = models.TextField(null=True, db_column="sqsRoleArn")
    """The role we assume to write to the queue, and the ExternalId that proves
    the request came from us rather than from anyone who learned the role's
    name. The recommended credential mode.
    """
    sqs_external_id = models.TextField(null=True, db_column="sqsExternalId")
    sqs_access_key_id = models.TextField(null=True, db_column="sqsAccessKeyId")
    """A static key pair for the queue, encrypted at rest exactly like the
    signing secret. Both halves are present or both are absent: a half-filled
    pair is what the AWS SDK reads as "we have credentials" and stops looking
    with, so the CHECK constraint refuses it.
    """
    sqs_secret_access_key_encrypted = models.TextField(null=True, db_column="sqsSecretAccessKeyEncrypted")
    secret_encrypted = models.TextField(db_column="secretEncrypted")
    signature_scheme = models.TextField(null=True, db_column="signatureScheme")
    """How deliveries are formatted and signed. Null: the batch envelope signed
    `t=,v1=`. `legacy_sha256`: one raw message per POST signed `sha256=`, kept
    only for endpoints migrated from governance anomaly destinations.
    """
    idempotency_key = models.TextField(null=True, db_column="idempotencyKey")
    """Set only by governance's anomaly destination migration: a create that
    repeats it in the same organization answers the endpoint it made. An
    archived endpoint gives its key up, so a retry after cleanup creates anew.
    """
    previous_secret_encrypted = models.TextField(null=True, db_column="previousSecretEncrypted")
    """The secret this endpoint signed with before the last roll. Deliveries
    carry a second `v1=` computed from it until it expires, so a receiver
    swaps secrets on its own schedule instead of dropping deliveries in the
    gap. Null once the window closes or if the secret was never rolled.
    """
    previous_secret_expires_at = UtcDateTimeField(null=True, db_column="previousSecretExpiresAt")
    """When the previous secret stops being signed with and accepted."""
    enabled_events = ArrayField(models.TextField(), db_column="enabledEvents", default=list)
    status = models.TextField(choices=WebhookEndpointStatus.choices, default=WebhookEndpointStatus.ACTIVE)
    disabled_reason = models.TextField(null=True, db_column="disabledReason")
    """Why the endpoint is disabled: "auto_failures_72h" or "manual"."""
    disabled_at = UtcDateTimeField(null=True, db_column="disabledAt")
    failing_since = UtcDateTimeField(null=True, db_column="failingSince")
    """Start of the current consecutive-failure streak; cleared on any 2xx.
    72 hours of unbroken failures from this instant auto-disables.
    """
    max_batch_size = models.IntegerField(db_column="maxBatchSize", default=100)
    """Delivery controls, each server-bounded (see webhookDeliveryControls
    constants): envelopes per POST, coalescing window before a partial
    batch ships, and concurrent POSTs to this receiver.
    """
    max_batch_delay_ms = models.IntegerField(db_column="maxBatchDelayMs", default=250)
    max_in_flight = models.IntegerField(db_column="maxInFlight", default=4)
    allow_self_signed_certificate = models.BooleanField(db_column="allowSelfSignedCertificate", default=False)
    """Self-hosted only: deliveries accept the receiver's self-signed TLS
    certificate. Off by default; the hosted product ignores it.
    """
    last_success_at = UtcDateTimeField(null=True, db_column="lastSuccessAt")
    last_failure_at = UtcDateTimeField(null=True, db_column="lastFailureAt")
    archived_at = UtcDateTimeField(null=True, db_column="archivedAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "WebhookEndpoint"
        constraints = [
            models.UniqueConstraint(fields=["organization", "idempotency_key"], name="WebhookEndpoint_organizationId_idempotencyKey_key"),
        ]
        indexes = [
            models.Index(fields=["organization", "status"]),
        ]


class WebhookEndpointDelivery(models.Model):
    """The per-ATTEMPT delivery log for BOTH webhook channels: the endpoints
    platform's batches and the automations channel's trigger fires. One row per
    HTTP attempt; `dispatchId` groups every attempt of one logical send (the
    value of the X-LangWatch-Delivery-Id or X-LangWatch-Event-Id header). A slim
    facts table: outcome, status, latency, a capped error message, and a
    truncated copy of a failure response. Our request content (URL, headers,
    body) is never stored. Pruned after 30 days.

    The two channels used to keep separate tables with the same columns and two
    prune implementations. They are one table with a discriminator now, so the
    tenancy columns of the channel that did not write a row are null.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    channel = models.TextField(choices=WebhookDeliveryChannel.choices, default=WebhookDeliveryChannel.platform)
    organization_id = models.TextField(null=True, db_column="organizationId")
    """Platform tenancy. Null on automations rows."""
    endpoint = models.ForeignKey("WebhookEndpoint", on_delete=models.CASCADE, related_name="deliveries", db_column="endpointId", null=True, db_index=False)
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="webhook_deliveries", db_column="projectId", null=True, db_index=False)
    """Automations tenancy. Null on platform rows."""
    trigger = models.ForeignKey("Trigger", on_delete=models.CASCADE, related_name="webhook_deliveries", db_column="triggerId", null=True, db_index=False)
    dispatch_id = models.TextField(db_column="dispatchId")
    attempt = models.IntegerField(null=True)
    """1-based attempt and the batch's envelope count. Both null on automations
    rows, which record neither: one trigger fire is one event, and the
    channel's sender does not thread its attempt number into the log.
    """
    event_count = models.IntegerField(null=True, db_column="eventCount")
    response_status = models.IntegerField(null=True, db_column="responseStatus")
    latency_ms = models.IntegerField(null=True, db_column="latencyMs")
    error = models.TextField(null=True)
    response = models.JSONField(null=True)
    outcome = models.TextField(choices=WebhookDeliveryOutcome.choices)
    fired_at = UtcDateTimeField(db_column="firedAt", default=now_milliseconds)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "WebhookEndpointDelivery"
        indexes = [
            models.Index(fields=["endpoint", "fired_at"]),
            models.Index(fields=["organization_id", "fired_at"]),
            models.Index(fields=["fired_at"]),
            models.Index(fields=["endpoint", "dispatch_id"]),
            models.Index(fields=["project", "trigger", "fired_at"]),
            models.Index(fields=["project", "trigger", "dispatch_id"]),
        ]


class Experiment(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    name = models.TextField(null=True)
    type = models.TextField(choices=ExperimentType.choices)
    slug = models.TextField()
    project = models.ForeignKey("Project", on_delete=models.RESTRICT, related_name="experiments", db_column="projectId", db_index=False)
    workflow = models.ForeignKey("Workflow", on_delete=models.SET_NULL, related_name="experiments", db_column="workflowId", null=True, db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    archived_at = UtcDateTimeField(null=True, db_column="archivedAt")
    workbench_state = models.JSONField(null=True, db_column="workbenchState")
    workbench_version = models.IntegerField(db_column="workbenchVersion", default=0)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Experiment"
        constraints = [
            models.UniqueConstraint(fields=["project", "slug"], name="Experiment_projectId_slug_key"),
        ]
        indexes = [
            models.Index(fields=["project"]),
            models.Index(fields=["workflow"]),
            models.Index(fields=["project", "archived_at"]),
        ]


class ExperimentVersion(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    experiment_id = models.TextField(db_column="experimentId")
    project_id = models.TextField(db_column="projectId")
    version = models.IntegerField()
    """The number the version is addressed and displayed by. Numbered rows hold
    a gap-free sequence. The rolling autosave row rides on the counter
    instead, which is what keeps it clear of that sequence.
    """
    counter_version = models.IntegerField(db_column="counterVersion")
    """`Experiment.workbenchVersion` at the moment this row was last written.
    The rolling row is rewritten in place, so this is the only column that
    says how recent its content is, and the version list sorts and pages by
    it. It also joins a row to the live state: the row whose `counterVersion`
    equals the experiment's counter holds what the workbench shows now.
    """
    auto_saved = models.BooleanField(db_column="autoSaved", default=False)
    commit_message = models.TextField(null=True, db_column="commitMessage")
    author_id = models.TextField(null=True, db_column="authorId")
    author_label = models.TextField(db_column="authorLabel")
    """Who wrote this version: "user", "langy" or "api"."""
    run_id = models.TextField(null=True, db_column="runId")
    """The run that wrote this version, when a run wrote it.

    A run writes its cells into the workbench state, which advances the
    counter. Without this, the page that started that run reads the bump as
    somebody else's write, refuses to save, and asks the reader to reload
    over their own unsaved edits. Named here, the page adopts the version and
    carries on.
    """
    state = models.JSONField()
    schema_version = models.TextField(db_column="schemaVersion", default="1")
    """Schema version of `state`, so a later shape change can be read forward."""
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "ExperimentVersion"
        constraints = [
            models.UniqueConstraint(fields=["experiment_id", "version"], name="ExperimentVersion_experimentId_version_key"),
        ]
        indexes = [
            models.Index(fields=["project_id", "experiment_id", "auto_saved"]),
            models.Index(fields=["project_id", "experiment_id", "version"]),
            models.Index(fields=["project_id", "experiment_id", "counter_version"]),
            models.Index(fields=["created_at"]),
        ]


class Annotation(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    project = models.ForeignKey("Project", on_delete=models.RESTRICT, related_name="annotations", db_column="projectId", db_index=False)
    trace_id = models.TextField(db_column="traceId")
    comment = models.TextField(null=True)
    is_thumbs_up = models.BooleanField(null=True, db_column="isThumbsUp")
    user = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="annotation", db_column="userId", null=True, db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    email = models.TextField(null=True)
    score_options = models.JSONField(null=True, db_column="scoreOptions")
    expected_output = models.TextField(null=True, db_column="expectedOutput")
    anchor_kind = models.TextField(null=True, db_column="anchorKind")
    anchor_id = models.TextField(null=True, db_column="anchorId")
    anchor_path = models.TextField(null=True, db_column="anchorPath")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Annotation"
        indexes = [
            models.Index(fields=["project"]),
            models.Index(fields=["trace_id"]),
            models.Index(fields=["user"]),
            models.Index(fields=["project", "trace_id"]),
            models.Index(fields=["project", "trace_id", "anchor_kind"]),
        ]


class ModelProvider(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    name = models.TextField()
    provider = models.TextField()
    routing_handle = models.TextField(null=True, db_column="routingHandle")
    enabled = models.BooleanField()
    custom_keys = models.JSONField(null=True, db_column="customKeys")
    extra_headers = models.JSONField(null=True, db_column="extraHeaders")
    custom_models = models.JSONField(null=True, db_column="customModels")
    custom_embeddings_models = models.JSONField(null=True, db_column="customEmbeddingsModels")
    deployment_mapping = models.JSONField(null=True, db_column="deploymentMapping")
    rate_limit_rpm = models.IntegerField(null=True, db_column="rateLimitRpm")
    rate_limit_tpm = models.IntegerField(null=True, db_column="rateLimitTpm")
    rate_limit_rpd = models.IntegerField(null=True, db_column="rateLimitRpd")
    rotation_policy = models.TextField(choices=GatewayProviderRotationPolicy.choices, db_column="rotationPolicy", default=GatewayProviderRotationPolicy.MANUAL)
    provider_config = models.JSONField(null=True, db_column="providerConfig")
    fallback_priority_global = models.IntegerField(null=True, db_column="fallbackPriorityGlobal")
    langy_skip_permissions_models = models.JSONField(null=True, db_column="langySkipPermissionsModels")
    health_status = models.TextField(choices=GatewayProviderHealthStatus.choices, db_column="healthStatus", default=GatewayProviderHealthStatus.UNKNOWN)
    circuit_opened_at = UtcDateTimeField(null=True, db_column="circuitOpenedAt")
    last_health_check_at = UtcDateTimeField(null=True, db_column="lastHealthCheckAt")
    disabled_at = UtcDateTimeField(null=True, db_column="disabledAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    organization_id = models.TextField(db_column="organizationId")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "ModelProvider"
        constraints = [
            models.UniqueConstraint(fields=["organization_id", "routing_handle"], name="ModelProvider_organizationId_routingHandle_key"),
        ]
        indexes = [
            models.Index(fields=["health_status"]),
            models.Index(fields=["fallback_priority_global"]),
            models.Index(fields=["organization_id"]),
        ]


class ModelProviderScope(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    model_provider = models.ForeignKey("ModelProvider", on_delete=models.CASCADE, related_name="scopes", db_column="modelProviderId", db_index=False)
    scope_type = models.TextField(choices=ModelProviderScopeType.choices, db_column="scopeType")
    scope_id = models.TextField(db_column="scopeId")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "ModelProviderScope"
        constraints = [
            models.UniqueConstraint(fields=["model_provider", "scope_type", "scope_id"], name="ModelProviderScope_modelProviderId_scopeType_scopeId_key"),
        ]
        indexes = [
            models.Index(fields=["scope_type", "scope_id"]),
            models.Index(fields=["model_provider"]),
        ]


class ModelDefaultConfig(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    config = models.JSONField()
    author_id = models.TextField(null=True, db_column="authorId")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    organization_id = models.TextField(db_column="organizationId")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "ModelDefaultConfig"
        indexes = [
            models.Index(fields=["organization_id"]),
        ]


class ModelDefaultConfigScope(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    config = models.ForeignKey("ModelDefaultConfig", on_delete=models.CASCADE, related_name="scopes", db_column="configId", db_index=False)
    scope_type = models.TextField(choices=ModelDefaultScopeType.choices, db_column="scopeType")
    scope_id = models.TextField(db_column="scopeId")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "ModelDefaultConfigScope"
        constraints = [
            models.UniqueConstraint(fields=["config", "scope_type", "scope_id"], name="ModelDefaultConfigScope_configId_scopeType_scopeId_key"),
        ]
        indexes = [
            models.Index(fields=["scope_type", "scope_id"]),
            models.Index(fields=["config"]),
        ]


class GithubInstallation(models.Model):
    """The organization's GitHub connection: a GitHub App installation mapped to a
    LangWatch organization. Short-lived (1h) installation access tokens are
    minted against these rows on demand, by Langy for writes (bot-authored pull
    requests, scoped to the installation's repositories) and by pull-request
    linkage for reads. No user token and no secret are stored here, the crown
    jewel is the App private key, held only in the control-plane env.
    Spec: specs/integrations/github-connection.feature.

    installationId is globally unique (one GitHub installation, one row); a
    single LangWatch organization may own MANY installations (different GitHub
    accounts), so the org to installation relationship is one-to-many.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    installation_id = models.TextField(unique=True, db_column="installationId")
    """GitHub installation id (numeric; stored as string). Globally unique."""
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="github_installations", db_column="organizationId", db_index=False)
    account_login = models.TextField(db_column="accountLogin")
    """The GitHub account (organization or user) the App is installed on."""
    account_type = models.TextField(db_column="accountType")
    """"Organization" or "User"."""
    account_id = models.TextField(db_column="accountId")
    """GitHub numeric account id (stable across renames)."""
    repository_selection = models.TextField(db_column="repositorySelection")
    """"all" or "selected" — whether the App can see every repo on the account
    or only a chosen subset.
    """
    repositories = models.JSONField(null=True)
    """Cached repository list [{ id, fullName }] when repositorySelection =
    "selected". Null when "all" (the full set is resolved live at mint time).
    """
    suspended_at = UtcDateTimeField(null=True, db_column="suspendedAt")
    """Set when GitHub suspends the installation; cleared on unsuspend. A
    suspended installation cannot mint tokens.
    """
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "GithubInstallation"
        indexes = [
            models.Index(fields=["organization"]),
        ]


class GithubPullRequest(models.Model):
    """A pull request discovered through the organization's GitHub connection and
    mapped to the branch it was opened from, so a coding-agent session that
    knows only its branch can name the pull request it produced.
    Spec: specs/integrations/github-connection.feature.

    The stored title, state, draft and timestamp columns are SNAPSHOTS taken
    when the mapping was last refreshed. They are fallback labels only: the
    frontend reads live status from GitHub, so a stale row never presents
    itself as current truth.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="github_pull_requests", db_column="organizationId", db_index=False)
    repository_host = models.TextField(db_column="repositoryHost")
    """The GitHub host the repository lives on, e.g. "github.com"."""
    repository_full_name = models.TextField(db_column="repositoryFullName")
    """"owner/repo", stored LOWERCASED so a lookup matches whatever casing the
    session reported.
    """
    head_branch = models.TextField(db_column="headBranch")
    """The branch the pull request was opened from."""
    pr_number = models.IntegerField(db_column="prNumber")
    html_url = models.TextField(db_column="htmlUrl")
    title = models.TextField()
    state = models.TextField()
    """GitHub's own state string: "open" or "closed"."""
    is_draft = models.BooleanField(db_column="isDraft", default=False)
    author_login = models.TextField(null=True, db_column="authorLogin")
    pr_created_at = UtcDateTimeField(db_column="prCreatedAt")
    pr_closed_at = UtcDateTimeField(null=True, db_column="prClosedAt")
    pr_merged_at = UtcDateTimeField(null=True, db_column="prMergedAt")
    pr_updated_at = UtcDateTimeField(null=True, db_column="prUpdatedAt")
    """GitHub's own `updated_at` for the snapshot this row holds, and the
    freshness guard on every write to it: a snapshot older than the stored one
    is skipped. GitHub may deliver webhooks out of order, and without the
    guard a late `opened` lands after `closed` and reopens a merged pull
    request. NULL means the row predates the column, so any write is accepted.
    """
    mapped_at = UtcDateTimeField(db_column="mappedAt", default=now_milliseconds)
    """When this row was first mapped to its branch."""
    last_checked_at = UtcDateTimeField(db_column="lastCheckedAt", default=now_milliseconds)
    """When the snapshot columns were last refreshed from GitHub."""
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "GithubPullRequest"
        constraints = [
            models.UniqueConstraint(fields=["organization", "repository_host", "repository_full_name", "pr_number"], name="GithubPullRequest_organizationId_repositoryHost_repositoryFullName_prNumber_key"),
        ]
        indexes = [
            models.Index(fields=["organization", "repository_host", "repository_full_name", "head_branch"]),
        ]


class GithubBranchPullRequestCheck(models.Model):
    """The per-branch bookkeeping behind pull-request mapping: when a branch was
    last looked up on GitHub, how many pull requests it resolved to, and when it
    may be asked about again. A branch with no pull request is the common case,
    so `notFoundAt` plus `recheckAfter` is what keeps an unmapped branch from
    costing a GitHub call on every read.
    Spec: specs/integrations/github-connection.feature.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="github_branch_pr_checks", db_column="organizationId", db_index=False)
    repository_host = models.TextField(db_column="repositoryHost")
    repository_full_name = models.TextField(db_column="repositoryFullName")
    """"owner/repo", stored LOWERCASED, matching GithubPullRequest."""
    head_branch = models.TextField(db_column="headBranch")
    last_checked_at = UtcDateTimeField(db_column="lastCheckedAt")
    pr_count = models.IntegerField(db_column="prCount", default=0)
    not_found_at = UtcDateTimeField(null=True, db_column="notFoundAt")
    """Set while the branch resolves to no pull request; cleared once one is
    found.
    """
    recheck_after = UtcDateTimeField(null=True, db_column="recheckAfter")
    """The earliest instant GitHub may be asked about this branch again."""
    attempts = models.IntegerField(default=0)
    """Consecutive lookups that found nothing, backing off the recheck."""
    last_requested_at = UtcDateTimeField(db_column="lastRequestedAt")
    """When a reader last asked about this branch, whether or not GitHub was
    called.
    """
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "GithubBranchPullRequestCheck"
        constraints = [
            models.UniqueConstraint(fields=["organization", "repository_host", "repository_full_name", "head_branch"], name="GithubBranchPullRequestCheck_organizationId_repositoryHost_repositoryFullName_headBranch_key"),
        ]
        indexes = [
            models.Index(fields=["recheck_after", "last_requested_at", "not_found_at"]),
        ]


class LangyConversationProjection(models.Model):
    """Rebuildable operational state folded from Langy's canonical event stream.
    The source of truth remains the ClickHouse event log.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    project_id = models.TextField(db_column="projectId")
    conversation_id = models.TextField(db_column="conversationId")
    user_id = models.TextField(db_column="userId")
    """Owner; first-writer-wins from the first message's userId."""
    title = models.TextField(null=True)
    title_source = models.TextField(choices=LangyProjectionTitleSource.choices, db_column="titleSource")
    """derived | auto | user — governs auto-title precedence (user is sticky)."""
    origin = models.TextField(default="interactive")
    """Who started the conversation: `interactive` for the person in the panel, `run` for a
    module acting for them with nobody at the keyboard. A row folded before the column
    existed reads as interactive.
    """
    status = models.TextField()
    is_shared = models.BooleanField(db_column="isShared", default=False)
    shared_at = models.FloatField(null=True, db_column="sharedAt")
    shared_by_id = models.TextField(null=True, db_column="sharedById")
    message_count = models.IntegerField(db_column="messageCount", default=0)
    last_activity_at = models.FloatField(null=True, db_column="lastActivityAt")
    current_turn_id = models.TextField(null=True, db_column="currentTurnId")
    """The turn in flight, or null when idle (busy guard reads this)."""
    last_error = models.TextField(null=True, db_column="lastError")
    last_model = models.TextField(null=True, db_column="lastModel")
    """The model the latest accepted turn ran on; reopening seeds the picker from it."""
    pending_handoff_token = models.TextField(null=True, db_column="pendingHandoffToken")
    """ADR-048 shutdown-handoff resume token (opaque, worker-authored)."""
    pending_handoff_turn_id = models.TextField(null=True, db_column="pendingHandoffTurnId")
    run_token = models.TextField(null=True, db_column="runToken")
    """Per-conversation HMAC key for worker stream frames; server-side only."""
    archived_at = models.FloatField(null=True, db_column="archivedAt")
    created_at = models.FloatField(db_column="createdAt")
    updated_at = models.FloatField(db_column="updatedAt")
    occurred_at = models.FloatField(db_column="occurredAt")
    """Business time of the latest applied event."""
    accepted_at = models.FloatField(db_column="acceptedAt")
    """Event-log acceptance time. Together with lastEventId this is the cursor."""
    last_event_id = models.TextField(db_column="lastEventId")
    projection_version = models.TextField(db_column="projectionVersion")

    class Meta:
        db_table = "LangyConversationProjection"
        constraints = [
            models.UniqueConstraint(fields=["project_id", "conversation_id"], name="LangyConversationProjection_projectId_conversationId_key"),
        ]
        indexes = [
            models.Index(fields=["project_id", "user_id", "conversation_id"]),
            models.Index(fields=["project_id", "user_id", "last_activity_at"]),
        ]


class LangyTurnRequest(models.Model):
    """Durable idempotency receipt for one logical Langy send. A transport retry
    reuses the same requestId and receives the original conversation/turn ids
    instead of minting a second message or turn.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    project_id = models.TextField(db_column="projectId")
    user_id = models.TextField(db_column="userId")
    request_id = models.TextField(db_column="requestId")
    conversation_id = models.TextField(db_column="conversationId")
    turn_id = models.TextField(db_column="turnId")
    status = models.TextField()
    lease_owner = models.TextField(db_column="leaseOwner")
    lease_expires_at = UtcDateTimeField(db_column="leaseExpiresAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "LangyTurnRequest"
        constraints = [
            models.UniqueConstraint(fields=["project_id", "user_id", "request_id"], name="LangyTurnRequest_projectId_userId_requestId_key"),
        ]
        indexes = [
            models.Index(fields=["project_id", "conversation_id"]),
        ]


class LangyActiveTurn(models.Model):
    """Postgres-authoritative one-turn admission for a conversation. PREPARING
    claims are leased so a crash before the durable start event can recover;
    COMMITTED claims are released only by the matching terminal event.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    project_id = models.TextField(db_column="projectId")
    conversation_id = models.TextField(db_column="conversationId")
    turn_id = models.TextField(db_column="turnId")
    request_id = models.TextField(db_column="requestId")
    user_id = models.TextField(db_column="userId")
    status = models.TextField()
    lease_owner = models.TextField(db_column="leaseOwner")
    lease_expires_at = UtcDateTimeField(db_column="leaseExpiresAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "LangyActiveTurn"
        constraints = [
            models.UniqueConstraint(fields=["project_id", "conversation_id"], name="LangyActiveTurn_projectId_conversationId_key"),
        ]
        indexes = [
            models.Index(fields=["project_id", "user_id"]),
        ]


class LangyConversationTurnProjection(models.Model):
    """Rebuildable per-turn operational render document."""

    id = models.TextField(primary_key=True, default=nanoid_str)
    project_id = models.TextField(db_column="projectId")
    conversation_id = models.TextField(db_column="conversationId")
    turn_id = models.TextField(db_column="turnId")
    status = models.TextField()
    """LangyConversationTurnStatus. Deliberately TEXT, not a Postgres enum: this
    is a PROJECTION — one writer (the fold), rebuildable from the event log,
    and nothing queries or indexes on it. An enum bought no selectivity here,
    only typo protection, and charged a migration for every new status that
    Postgres then could not remove. The one definition lives in
    `LANGY_CONVERSATION_TURN_STATUS`; the repository validates against it.
    """
    question_parts = models.JSONField(db_column="questionParts")
    """LangyMessagePart[] — the user question that opened the turn."""
    answer_parts = models.JSONField(db_column="answerParts")
    """LangyMessagePart[] — the agent's whole final answer."""
    tool_calls = models.JSONField(db_column="toolCalls")
    """LangyTurnToolCall[] — tool-call lifecycle audit in initiation order."""
    plan = models.JSONField(null=True)
    """LangyPlanItemData[] | null — null means the turn never kept a plan."""
    error = models.TextField(null=True)
    started_at = models.FloatField(null=True, db_column="startedAt")
    ended_at = models.FloatField(null=True, db_column="endedAt")
    created_at = models.FloatField(db_column="createdAt")
    updated_at = models.FloatField(db_column="updatedAt")
    occurred_at = models.FloatField(db_column="occurredAt")
    accepted_at = models.FloatField(db_column="acceptedAt")
    last_event_id = models.TextField(db_column="lastEventId")
    projection_version = models.TextField(db_column="projectionVersion")

    class Meta:
        db_table = "LangyConversationTurnProjection"
        constraints = [
            models.UniqueConstraint(fields=["project_id", "conversation_id", "turn_id"], name="LangyConversationTurnProjection_projectId_conversationId_turnId_key"),
        ]
        indexes = [
            models.Index(fields=["project_id", "conversation_id", "updated_at"]),
        ]


class LangyMessageProjection(models.Model):
    """Rebuildable operational message fact. One source event produces one row."""

    id = models.TextField(primary_key=True, default=nanoid_str)
    project_id = models.TextField(db_column="projectId")
    conversation_id = models.TextField(db_column="conversationId")
    message_id = models.TextField(db_column="messageId")
    role = models.TextField()
    parts = models.JSONField()
    source_event_id = models.TextField(db_column="sourceEventId")
    occurred_at = models.FloatField(db_column="occurredAt")
    accepted_at = models.FloatField(db_column="acceptedAt")
    created_at = models.FloatField(db_column="createdAt")
    updated_at = models.FloatField(db_column="updatedAt")

    class Meta:
        db_table = "LangyMessageProjection"
        constraints = [
            models.UniqueConstraint(fields=["project_id", "conversation_id", "message_id"], name="LangyMessageProjection_projectId_conversationId_messageId_key"),
            models.UniqueConstraint(fields=["project_id", "source_event_id"], name="LangyMessageProjection_projectId_sourceEventId_key"),
        ]
        indexes = [
            models.Index(fields=["project_id", "conversation_id", "created_at"]),
        ]


class ProjectSecret(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    project = models.ForeignKey("Project", on_delete=models.RESTRICT, related_name="project_secrets", db_column="projectId", db_index=False)
    name = models.TextField()
    encrypted_value = models.TextField(db_column="encryptedValue")
    created_by = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="secrets_created", db_column="createdById", db_index=False)
    updated_by = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="secrets_updated", db_column="updatedById", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "ProjectSecret"
        constraints = [
            models.UniqueConstraint(fields=["project", "name"], name="ProjectSecret_projectId_name_key"),
        ]
        indexes = [
            models.Index(fields=["project"]),
        ]


class TriggerSent(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    trigger = models.ForeignKey("Trigger", on_delete=models.RESTRICT, related_name="trigger_sent", db_column="triggerId", db_index=False)
    trace_id = models.TextField(null=True, db_column="traceId")
    custom_graph = models.ForeignKey("CustomGraph", on_delete=models.CASCADE, related_name="trigger_sent", db_column="customGraphId", null=True, db_index=False)
    project = models.ForeignKey("Project", on_delete=models.RESTRICT, related_name="trigger_sent", db_column="projectId", db_index=False)
    resolved_at = UtcDateTimeField(null=True, db_column="resolvedAt")
    open_incident_key = models.TextField(null=True, unique=True, db_column="openIncidentKey")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "TriggerSent"
        constraints = [
            models.UniqueConstraint(fields=["trigger", "trace_id"], name="TriggerSent_triggerId_traceId_key"),
        ]
        indexes = [
            models.Index(fields=["trigger"]),
            models.Index(fields=["project"]),
            models.Index(fields=["custom_graph"]),
            models.Index(fields=["resolved_at"]),
            models.Index(fields=["project", "trigger", "created_at", "id"]),
        ]


class TriggerLatestEvaluation(models.Model):
    trigger = models.OneToOneField("Trigger", on_delete=models.CASCADE, related_name="latest_evaluation", primary_key=True, db_column="triggerId")
    project_id = models.TextField(db_column="projectId")
    evaluated_at = models.DateTimeField(db_column="evaluatedAt")
    verdict = models.TextField()
    observed_value = models.FloatField(null=True, db_column="observedValue")
    threshold = models.FloatField(null=True)
    operator = models.TextField(null=True)
    time_period_minutes = models.IntegerField(null=True, db_column="timePeriodMinutes")
    skip_code = models.TextField(null=True, db_column="skipCode")
    updated_at = AutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "TriggerLatestEvaluation"
        indexes = [
            models.Index(fields=["project_id"]),
        ]


class EmailSuppression(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    project = models.ForeignKey("Project", on_delete=models.RESTRICT, related_name="email_suppressions", db_column="projectId", db_index=False)
    email = models.TextField()
    trigger_id = models.TextField(null=True, db_column="triggerId")
    reason = models.TextField(default="unsubscribe")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "EmailSuppression"
        constraints = [
            models.UniqueConstraint(fields=["project", "email", "trigger_id"], name="EmailSuppression_projectId_email_triggerId_key"),
        ]
        indexes = [
            models.Index(fields=["project"]),
        ]


class AnnotationScore(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    name = models.TextField()
    project = models.ForeignKey("Project", on_delete=models.RESTRICT, related_name="annotation_scores", db_column="projectId", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    deleted_at = UtcDateTimeField(null=True, db_column="deletedAt")
    description = models.TextField(null=True)
    active = models.BooleanField(default=True)
    data_type = models.TextField(choices=AnnotationScoreDataType.choices, null=True, db_column="dataType")
    options = models.JSONField(null=True)
    default_value = models.JSONField(null=True, db_column="defaultValue")
    global_field = models.BooleanField(db_column="global", default=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "AnnotationScore"
        indexes = [
            models.Index(fields=["project"]),
        ]


class AnnotationQueue(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    name = models.TextField()
    slug = models.TextField()
    project = models.ForeignKey("Project", on_delete=models.RESTRICT, related_name="annotation_queue", db_column="projectId", db_index=False)
    description = models.TextField(null=True)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "AnnotationQueue"
        constraints = [
            models.UniqueConstraint(fields=["project", "slug"], name="AnnotationQueue_projectId_slug_key"),
        ]
        indexes = [
            models.Index(fields=["project"]),
        ]


class AnnotationQueueMembers(models.Model):
    pk = models.CompositePrimaryKey("annotation_queue_id", "user_id")
    annotation_queue = models.ForeignKey("AnnotationQueue", on_delete=models.RESTRICT, related_name="members", db_column="annotationQueueId", db_index=False)
    user = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="annotation_queues", db_column="userId", db_index=False)

    class Meta:
        db_table = "AnnotationQueueMembers"
        indexes = [
            models.Index(fields=["annotation_queue"]),
            models.Index(fields=["user"]),
        ]


class AnnotationQueueScores(models.Model):
    pk = models.CompositePrimaryKey("annotation_queue_id", "annotation_score_id")
    annotation_queue = models.ForeignKey("AnnotationQueue", on_delete=models.RESTRICT, related_name="annotation_queue_scores", db_column="annotationQueueId", db_index=False)
    annotation_score = models.ForeignKey("AnnotationScore", on_delete=models.RESTRICT, related_name="annotation_queue_scores", db_column="annotationScoreId", db_index=False)

    class Meta:
        db_table = "AnnotationQueueScores"
        indexes = [
            models.Index(fields=["annotation_queue"]),
            models.Index(fields=["annotation_score"]),
        ]


class AnnotationQueueItem(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    annotation_queue = models.ForeignKey("AnnotationQueue", on_delete=models.SET_NULL, related_name="annotation_queue_items", db_column="annotationQueueId", null=True, db_index=False)
    user = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="assigned_queue_items", db_column="userId", null=True, db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    created_by_user = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="created_queue_items", db_column="createdByUserId", null=True, db_index=False)
    trace_id = models.TextField(db_column="traceId")
    project = models.ForeignKey("Project", on_delete=models.RESTRICT, related_name="annotation_queue_item", db_column="projectId", db_index=False)
    done_at = UtcDateTimeField(null=True, db_column="doneAt")
    marked_for_dataset_at = UtcDateTimeField(null=True, db_column="markedForDatasetAt")
    """Set while the annotator wants this item's trace added to a dataset once
    the queue walk is finished. Persisted rather than kept in the browser so
    the selection survives a refresh and stays visible to teammates.
    """
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "AnnotationQueueItem"
        constraints = [
            models.UniqueConstraint(fields=["trace_id", "user", "project"], name="AnnotationQueueItem_traceId_userId_projectId_key"),
            models.UniqueConstraint(fields=["trace_id", "annotation_queue", "project"], name="AnnotationQueueItem_traceId_annotationQueueId_projectId_key"),
        ]
        indexes = [
            models.Index(fields=["annotation_queue"]),
            models.Index(fields=["user"]),
            models.Index(fields=["project"]),
            models.Index(fields=["project", "done_at"]),
            models.Index(fields=["project", "user", "done_at"]),
            models.Index(fields=["annotation_queue", "done_at"]),
        ]


class ShareLink(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    token = models.TextField(unique=True)
    resource_type = models.TextField(choices=ShareResourceType.choices, db_column="resourceType")
    resource_id = models.TextField(db_column="resourceId")
    thread_id = models.TextField(null=True, db_column="threadId")
    project = models.ForeignKey("Project", on_delete=models.RESTRICT, related_name="share_links", db_column="projectId", db_index=False)
    user = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="share_links", db_column="userId", null=True, db_index=False)
    visibility = models.TextField(choices=ShareVisibility.choices, default=ShareVisibility.PUBLIC)
    expires_at = UtcDateTimeField(null=True, db_column="expiresAt")
    max_views = models.IntegerField(null=True, db_column="maxViews")
    view_count = models.IntegerField(db_column="viewCount", default=0)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "ShareLink"
        indexes = [
            models.Index(fields=["project", "resource_type", "resource_id"]),
            models.Index(fields=["user"]),
        ]


class PinnedTrace(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    project = models.ForeignKey("Project", on_delete=models.RESTRICT, related_name="pinned_traces", db_column="projectId", db_index=False)
    trace_id = models.TextField(db_column="traceId")
    user = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="pinned_traces", db_column="userId", null=True, db_index=False)
    source = models.TextField(choices=PinSource.choices, default=PinSource.manual)
    reason = models.TextField(null=True)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "PinnedTrace"
        constraints = [
            models.UniqueConstraint(fields=["project", "trace_id"], name="PinnedTrace_projectId_traceId_key"),
        ]
        indexes = [
            models.Index(fields=["project"]),
        ]


class TraceEditOverlay(models.Model):
    id = models.TextField(primary_key=True, default=None)
    project = models.ForeignKey("Project", on_delete=models.RESTRICT, related_name="trace_edit_overlays", db_column="projectId", db_index=False)
    trace_id = models.TextField(db_column="traceId")
    patch = models.JSONField()
    created_by = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="trace_edit_overlays_created", db_column="createdById", null=True, db_index=False)
    updated_by = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="trace_edit_overlays_updated", db_column="updatedById", null=True, db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "TraceEditOverlay"
        constraints = [
            models.UniqueConstraint(fields=["project", "trace_id"], name="TraceEditOverlay_projectId_traceId_key"),
        ]
        indexes = [
            models.Index(fields=["project"]),
        ]


class TraceIngestSourceBilling(models.Model):
    pk = models.CompositePrimaryKey("organization_id", "source_type")
    organization_id = models.TextField(db_column="organizationId")
    source_type = models.TextField(db_column="sourceType")
    billed = models.BooleanField()
    recorded_at = UtcDateTimeField(db_column="recordedAt")
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "TraceIngestSourceBilling"


class RetentionPolicy(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    organization_id = models.TextField(db_column="organizationId")
    scope_type = models.TextField(choices=RetentionPolicyScopeType.choices, db_column="scopeType")
    scope_id = models.TextField(db_column="scopeId")
    category = models.TextField()
    retention_days = models.IntegerField(db_column="retentionDays")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "RetentionPolicy"
        constraints = [
            models.UniqueConstraint(fields=["scope_type", "scope_id", "category"], name="RetentionPolicy_scopeType_scopeId_category_key"),
        ]
        indexes = [
            models.Index(fields=["organization_id"]),
            models.Index(fields=["scope_type", "scope_id"]),
        ]


class DataPrivacyPolicy(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    organization_id = models.TextField(db_column="organizationId")
    scope_type = models.TextField(choices=DataPrivacyScopeType.choices, db_column="scopeType")
    scope_id = models.TextField(db_column="scopeId")
    personal_only = models.BooleanField(db_column="personalOnly", default=False)
    config = models.JSONField()
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "DataPrivacyPolicy"
        constraints = [
            models.UniqueConstraint(fields=["scope_type", "scope_id", "personal_only"], name="DataPrivacyPolicy_scopeType_scopeId_personalOnly_key"),
        ]
        indexes = [
            models.Index(fields=["organization_id"]),
            models.Index(fields=["scope_type", "scope_id"]),
        ]


class InstantEvalJudgeProject(models.Model):
    project_id = models.TextField(primary_key=True, db_column="projectId", default=None)
    organization_id = models.TextField(db_column="organizationId")
    created_at = UtcDateTimeField(db_column="createdAt")

    class Meta:
        db_table = "InstantEvalJudgeProject"


class InstantEvalJudgeUsageBilling(models.Model):
    organization_id = models.TextField(primary_key=True, db_column="organizationId", default=None)
    usage_billed = models.BooleanField(db_column="usageBilled")
    occurred_at = UtcDateTimeField(db_column="occurredAt")
    from_catch_up = models.BooleanField(db_column="fromCatchUp")

    class Meta:
        db_table = "InstantEvalJudgeUsageBilling"


class InstantEvalJudgeSpend(models.Model):
    pk = models.CompositePrimaryKey("organization_id", "request_id")
    organization_id = models.TextField(db_column="organizationId")
    request_id = models.TextField(db_column="requestId")
    spend_nano_usd = models.BigIntegerField(db_column="spendNanoUsd")
    occurred_at = UtcDateTimeField(db_column="occurredAt")

    class Meta:
        db_table = "InstantEvalJudgeSpend"


class CustomLLMModelCost(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    organization_id = models.TextField(db_column="organizationId")
    scope_type = models.TextField(choices=CustomLLMModelCostScopeType.choices, db_column="scopeType", default=CustomLLMModelCostScopeType.PROJECT)
    scope_id = models.TextField(db_column="scopeId")
    project_id = models.TextField(null=True, db_column="projectId")
    model = models.TextField()
    regex = models.TextField()
    input_cost_per_token = models.FloatField(null=True, db_column="inputCostPerToken")
    output_cost_per_token = models.FloatField(null=True, db_column="outputCostPerToken")
    cache_read_cost_per_token = models.FloatField(null=True, db_column="cacheReadCostPerToken")
    cache_creation_cost_per_token = models.FloatField(null=True, db_column="cacheCreationCostPerToken")
    cache_creation1h_cost_per_token = models.FloatField(null=True, db_column="cacheCreation1hCostPerToken")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "CustomLLMModelCost"
        indexes = [
            models.Index(fields=["organization_id"]),
            models.Index(fields=["scope_type", "scope_id"]),
            models.Index(fields=["project_id"]),
        ]


class Workflow(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    project = models.ForeignKey("Project", on_delete=models.RESTRICT, related_name="workflows", db_column="projectId", db_index=False)
    name = models.TextField()
    icon = models.TextField()
    description = models.TextField()
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    latest_version = models.ForeignKey("WorkflowVersion", on_delete=models.RESTRICT, related_name="workflow_as_latest_version", db_column="latestVersionId", null=True, db_index=False)
    current_version = models.ForeignKey("WorkflowVersion", on_delete=models.RESTRICT, related_name="workflow_as_current_version", db_column="currentVersionId", null=True, db_index=False)
    published_id = models.TextField(null=True, db_column="publishedId")
    published_by = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="workflow", db_column="publishedById", null=True, db_index=False)
    archived_at = UtcDateTimeField(null=True, db_column="archivedAt")
    is_evaluator = models.BooleanField(db_column="isEvaluator", default=False)
    is_component = models.BooleanField(db_column="isComponent", default=False)
    copied_from = models.ForeignKey("self", on_delete=models.RESTRICT, related_name="copied_workflows", db_column="copiedFromWorkflowId", null=True, db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Workflow"
        indexes = [
            models.Index(fields=["project"]),
            models.Index(fields=["published_by"]),
            models.Index(fields=["latest_version"]),
            models.Index(fields=["current_version"]),
            models.Index(fields=["copied_from"]),
        ]


class WorkflowVersion(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    version = models.TextField()
    commit_message = models.TextField(db_column="commitMessage")
    author = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="workflow_version", db_column="authorId", db_index=False)
    project = models.ForeignKey("Project", on_delete=models.RESTRICT, related_name="workflow_version", db_column="projectId", db_index=False)
    workflow = models.ForeignKey("Workflow", on_delete=models.RESTRICT, related_name="versions", db_column="workflowId", db_index=False)
    parent = models.ForeignKey("self", on_delete=models.RESTRICT, related_name="children", db_column="parentId", null=True, db_index=False)
    auto_saved = models.BooleanField(db_column="autoSaved", default=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    dsl = models.JSONField()
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "WorkflowVersion"
        constraints = [
            models.UniqueConstraint(fields=["workflow", "version"], name="WorkflowVersion_workflowId_version_key"),
        ]
        indexes = [
            models.Index(fields=["workflow"]),
            models.Index(fields=["project"]),
            models.Index(fields=["parent"]),
            models.Index(fields=["author"]),
            models.Index(fields=["created_at"]),
        ]


class AuditLog(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    user_id = models.TextField(null=True, db_column="userId")
    project_id = models.TextField(null=True, db_column="projectId")
    organization_id = models.TextField(null=True, db_column="organizationId")
    action = models.TextField()
    args = models.JSONField(null=True)
    error = models.TextField(null=True)
    ip_address = models.TextField(null=True, db_column="ipAddress")
    user_agent = models.TextField(null=True, db_column="userAgent")
    metadata = models.JSONField(null=True)
    actor_user_id = models.TextField(null=True, db_column="actorUserId")
    """WHO REALLY DID IT, when that is not `userId`.

    Under an impersonation the session's `user.id` is the SUBJECT — the
    customer whose access an operator is borrowing — so a row naming only
    `userId` filed the operator's act against the customer. An operator
    granting a permanent way into somebody's organization appeared, in the
    only durable record of it, to be that somebody.

    Null on an ordinary request, where the two would be the same person and
    saying so twice adds nothing. ADR-118 / D06: every decision names both.
    """
    target_kind = models.TextField(null=True, db_column="targetKind")
    target_id = models.TextField(null=True, db_column="targetId")
    before = models.JSONField(null=True)
    after = models.JSONField(null=True)
    idempotency_key = models.TextField(null=True, unique=True, db_column="idempotencyKey")
    """A producer's outbox key: a redelivered audit intent finds its row here (Alex, audit R2)."""

    class Meta:
        db_table = "AuditLog"
        indexes = [
            models.Index(fields=["created_at"]),
            models.Index(fields=["user_id"]),
            models.Index(fields=["actor_user_id"]),
            models.Index(fields=["project_id"]),
            models.Index(fields=["action"]),
            models.Index(fields=["ip_address"]),
            models.Index(fields=["organization_id", "created_at"]),
            models.Index(fields=["target_kind", "target_id"]),
        ]


class LlmPromptConfig(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    handle = models.TextField(null=True, unique=True)
    name = models.TextField()
    project = models.ForeignKey("Project", on_delete=models.RESTRICT, related_name="llm_prompt_configs", db_column="projectId", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    deleted_at = UtcDateTimeField(null=True, db_column="deletedAt")
    organization_id = models.TextField(db_column="organizationId")
    scope = models.TextField(choices=PromptScope.choices, default=PromptScope.PROJECT)
    copied_from = models.ForeignKey("self", on_delete=models.RESTRICT, related_name="copied_prompts", db_column="copiedFromPromptId", null=True, db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "LlmPromptConfig"
        constraints = [
            models.UniqueConstraint(fields=["handle"], name="LlmPromptConfig_handle_key"),
        ]
        indexes = [
            models.Index(fields=["project"]),
            models.Index(fields=["organization_id", "scope"]),
            models.Index(fields=["copied_from"]),
        ]


class LlmPromptConfigVersion(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    version = models.IntegerField()
    commit_message = models.TextField(null=True, db_column="commitMessage")
    author = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="llm_prompt_config_versions", db_column="authorId", null=True, db_index=False)
    config = models.ForeignKey("LlmPromptConfig", on_delete=models.CASCADE, related_name="versions", db_column="configId", db_index=False)
    config_data = models.JSONField(db_column="config")
    schema_version = models.TextField(db_column="schemaVersion")
    runtime_parameters = models.JSONField(db_column="runtimeParameters", default=dict)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    project_id = models.TextField(db_column="projectId")

    class Meta:
        db_table = "LlmPromptConfigVersion"
        constraints = [
            models.UniqueConstraint(fields=["config", "version"], name="LlmPromptConfigVersion_configId_version_key"),
        ]
        indexes = [
            models.Index(fields=["config"]),
            models.Index(fields=["author"]),
            models.Index(fields=["created_at"]),
        ]


class PromptTagAssignment(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    config = models.ForeignKey("LlmPromptConfig", on_delete=models.CASCADE, related_name="version_tags", db_column="configId", db_index=False)
    version = models.ForeignKey("LlmPromptConfigVersion", on_delete=models.CASCADE, related_name="version_tags", db_column="versionId", db_index=False)
    prompt_tag = models.ForeignKey("PromptTag", on_delete=models.CASCADE, related_name="assignments", db_column="tagId", db_index=False)
    project_id = models.TextField(db_column="projectId")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    created_by = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="version_tags_created", db_column="createdById", null=True, db_index=False)
    updated_by = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="version_tags_updated", db_column="updatedById", null=True, db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "PromptTagAssignment"
        constraints = [
            models.UniqueConstraint(fields=["config", "prompt_tag"], name="PromptTagAssignment_configId_tagId_key"),
        ]
        indexes = [
            models.Index(fields=["config"]),
            models.Index(fields=["version"]),
            models.Index(fields=["project_id"]),
            models.Index(fields=["prompt_tag"]),
        ]


class PromptTag(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="prompt_tags", db_column="organizationId", db_index=False)
    name = models.TextField()
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    created_by = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="prompt_tags_created", db_column="createdById", null=True, db_index=False)
    updated_by = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="prompt_tags_updated", db_column="updatedById", null=True, db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "PromptTag"
        constraints = [
            models.UniqueConstraint(fields=["organization", "name"], name="PromptTag_organizationId_name_key"),
        ]
        indexes = [
            models.Index(fields=["organization"]),
        ]


class Analytics(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    project = models.ForeignKey("Project", on_delete=models.RESTRICT, related_name="analytics", db_column="projectId", db_index=False)
    key = models.TextField()
    value = models.JSONField(null=True)
    numeric_value = models.FloatField(null=True, db_column="numericValue")
    string_value = models.TextField(null=True, db_column="stringValue")
    bool_value = models.BooleanField(null=True, db_column="boolValue")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "Analytics"
        indexes = [
            models.Index(fields=["created_at", "key", "project", "numeric_value"]),
            models.Index(fields=["project"]),
        ]


class CustomRole(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    organization = models.ForeignKey("Organization", on_delete=models.RESTRICT, related_name="custom_roles", db_column="organizationId", db_index=False)
    name = models.TextField()
    description = models.TextField(null=True)
    permissions = models.JSONField()
    kind = models.TextField(default="custom")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "CustomRole"
        constraints = [
            models.UniqueConstraint(fields=["organization", "name"], name="CustomRole_organizationId_name_key"),
        ]
        indexes = [
            models.Index(fields=["organization"]),
        ]


class Notification(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="notifications", db_column="organizationId", null=True, db_index=False)
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="notifications", db_column="projectId", null=True, db_index=False)
    metadata = models.JSONField()
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    sent_at = UtcDateTimeField(db_column="sentAt")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Notification"
        indexes = [
            models.Index(fields=["organization"]),
            models.Index(fields=["project"]),
            models.Index(fields=["created_at"]),
            models.Index(fields=["sent_at"]),
        ]


class Agent(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    project = models.ForeignKey("Project", on_delete=models.RESTRICT, related_name="agents", db_column="projectId", db_index=False)
    name = models.TextField()
    type = models.TextField()
    config = models.JSONField()
    workflow = models.ForeignKey("Workflow", on_delete=models.SET_NULL, related_name="agents", db_column="workflowId", null=True, db_index=False)
    copied_from = models.ForeignKey("self", on_delete=models.RESTRICT, related_name="copied_agents", db_column="copiedFromAgentId", null=True, db_index=False)
    environment = models.TextField(null=True)
    owner_user_id = models.TextField(null=True, db_column="ownerUserId")
    host_label = models.TextField(null=True, db_column="hostLabel")
    identity_key = models.TextField(null=True, db_column="identityKey")
    last_seen_at = UtcDateTimeField(null=True, db_column="lastSeenAt")
    archived_at = UtcDateTimeField(null=True, db_column="archivedAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    created_by_id = models.TextField(null=True, db_column="createdById")
    updated_by_id = models.TextField(null=True, db_column="updatedById")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Agent"
        constraints = [
            models.UniqueConstraint(fields=["project", "identity_key"], name="Agent_projectId_identityKey_key"),
        ]
        indexes = [
            models.Index(fields=["project"]),
            models.Index(fields=["workflow"]),
            models.Index(fields=["copied_from"]),
            models.Index(fields=["project", "last_seen_at"]),
        ]


class AgentVersion(models.Model):
    """One immutable snapshot of an agent; the Agent row is always the latest."""

    id = models.TextField(primary_key=True, default=None)
    project_id = models.TextField(db_column="projectId")
    agent_id = models.TextField(db_column="agentId")
    version = models.IntegerField()
    name = models.TextField()
    type = models.TextField()
    config = models.JSONField()
    workflow_id = models.TextField(null=True, db_column="workflowId")
    commit_message = models.TextField(null=True, db_column="commitMessage")
    author_id = models.TextField(null=True, db_column="authorId")
    author_kind = models.TextField(null=True, db_column="authorKind")
    """"user", "api_key" or "system": who wrote the version when authorId is null."""
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "AgentVersion"
        constraints = [
            models.UniqueConstraint(fields=["agent_id", "version"], name="AgentVersion_agentId_version_key"),
        ]
        indexes = [
            models.Index(fields=["project_id"]),
        ]


class AgentVersionLabel(models.Model):
    """A deploy label ("production", "staging") pointing at one version of an agent."""

    id = models.TextField(primary_key=True, default=None)
    project_id = models.TextField(db_column="projectId")
    agent_id = models.TextField(db_column="agentId")
    label = models.TextField()
    version_id = models.TextField(db_column="versionId")
    created_by_id = models.TextField(null=True, db_column="createdById")
    updated_by_id = models.TextField(null=True, db_column="updatedById")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "AgentVersionLabel"
        constraints = [
            models.UniqueConstraint(fields=["agent_id", "label"], name="AgentVersionLabel_agentId_label_key"),
        ]
        indexes = [
            models.Index(fields=["project_id"]),
            models.Index(fields=["version_id"]),
        ]


class Evaluator(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    project = models.ForeignKey("Project", on_delete=models.RESTRICT, related_name="evaluators", db_column="projectId", db_index=False)
    name = models.TextField()
    slug = models.TextField(null=True)
    type = models.TextField()
    config = models.JSONField()
    workflow = models.ForeignKey("Workflow", on_delete=models.SET_NULL, related_name="evaluators", db_column="workflowId", null=True, db_index=False)
    copied_from = models.ForeignKey("self", on_delete=models.RESTRICT, related_name="copied_evaluators", db_column="copiedFromEvaluatorId", null=True, db_index=False)
    archived_at = UtcDateTimeField(null=True, db_column="archivedAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    created_by_id = models.TextField(null=True, db_column="createdById")
    updated_by_id = models.TextField(null=True, db_column="updatedById")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Evaluator"
        constraints = [
            models.UniqueConstraint(fields=["project", "slug"], name="Evaluator_projectId_slug_key"),
        ]
        indexes = [
            models.Index(fields=["project"]),
            models.Index(fields=["workflow"]),
            models.Index(fields=["copied_from"]),
        ]


class EvaluatorVersion(models.Model):
    """One immutable snapshot of an evaluator; the Evaluator row is always the latest."""

    id = models.TextField(primary_key=True, default=None)
    project_id = models.TextField(db_column="projectId")
    evaluator_id = models.TextField(db_column="evaluatorId")
    version = models.IntegerField()
    name = models.TextField()
    type = models.TextField()
    config = models.JSONField()
    workflow_id = models.TextField(null=True, db_column="workflowId")
    commit_message = models.TextField(null=True, db_column="commitMessage")
    author_id = models.TextField(null=True, db_column="authorId")
    author_kind = models.TextField(null=True, db_column="authorKind")
    """"user", "api_key" or "system": who wrote the version when authorId is null."""
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "EvaluatorVersion"
        constraints = [
            models.UniqueConstraint(fields=["evaluator_id", "version"], name="EvaluatorVersion_evaluatorId_version_key"),
        ]
        indexes = [
            models.Index(fields=["project_id"]),
        ]


class EvaluatorVersionLabel(models.Model):
    """A deploy label ("production", "staging") pointing at one version of an evaluator."""

    id = models.TextField(primary_key=True, default=None)
    project_id = models.TextField(db_column="projectId")
    evaluator_id = models.TextField(db_column="evaluatorId")
    label = models.TextField()
    version_id = models.TextField(db_column="versionId")
    created_by_id = models.TextField(null=True, db_column="createdById")
    updated_by_id = models.TextField(null=True, db_column="updatedById")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "EvaluatorVersionLabel"
        constraints = [
            models.UniqueConstraint(fields=["evaluator_id", "label"], name="EvaluatorVersionLabel_evaluatorId_label_key"),
        ]
        indexes = [
            models.Index(fields=["project_id"]),
            models.Index(fields=["version_id"]),
        ]


class Scenario(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    project = models.ForeignKey("Project", on_delete=models.RESTRICT, related_name="scenarios", db_column="projectId", db_index=False)
    name = models.TextField()
    situation = models.TextField()
    criteria = ArrayField(models.TextField())
    labels = ArrayField(models.TextField())
    parameters = models.JSONField(null=True)
    simulator_model = models.TextField(null=True, db_column="simulatorModel")
    judge_model = models.TextField(null=True, db_column="judgeModel")
    max_turns = models.IntegerField(null=True, db_column="maxTurns")
    min_turns = models.IntegerField(null=True, db_column="minTurns")
    fields = models.JSONField(null=True)
    caller_voice = models.JSONField(null=True, db_column="callerVoice")
    test_suite = models.ForeignKey("SimulationSuite", on_delete=models.SET_NULL, related_name="test_suite_scenarios", db_column="testSuiteId", null=True, db_index=False)
    version = models.IntegerField(default=1)
    last_updated_by = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="scenarios", db_column="lastUpdatedById", null=True, db_index=False)
    archived_at = UtcDateTimeField(null=True, db_column="archivedAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Scenario"
        indexes = [
            models.Index(fields=["project"]),
            models.Index(fields=["project", "test_suite"]),
        ]


class ScenarioVersion(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    scenario = models.ForeignKey("Scenario", on_delete=models.CASCADE, related_name="versions", db_column="scenarioId", db_index=False)
    """Deleting the case takes its history with it. Under relationMode="prisma"
    the cascade is the client's, so a raw SQL delete still has to clear these
    rows itself.
    """
    project_id = models.TextField(db_column="projectId")
    """The project of the parent case, carried here so every read filters by
    tenant without a join. The service is what keeps the pair honest: no
    composite foreign key can, under relationMode="prisma".
    """
    version = models.IntegerField()
    author_id = models.TextField(null=True, db_column="authorId")
    author_label = models.TextField(db_column="authorLabel")
    """Who wrote this version: "user", "api", "cli" or "langy"."""
    change_description = models.TextField(null=True, db_column="changeDescription")
    snapshot = models.JSONField()
    """Envelope { schemaVersion, fields, changedFields } holding the editable
    field set as saved and the names of the fields this save changed.
    """
    schema_version = models.IntegerField(db_column="schemaVersion", default=1)
    """Schema version of `snapshot`, so a later shape change can be read forward."""
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "ScenarioVersion"
        constraints = [
            models.UniqueConstraint(fields=["scenario", "version"], name="ScenarioVersion_scenarioId_version_key"),
        ]
        indexes = [
            models.Index(fields=["project_id", "scenario", "version"]),
            models.Index(fields=["created_at"]),
        ]


class SimulationSuite(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    project = models.ForeignKey("Project", on_delete=models.RESTRICT, related_name="suite_configurations", db_column="projectId", db_index=False)
    name = models.TextField()
    slug = models.TextField()
    description = models.TextField(null=True)
    scenario_ids = ArrayField(models.TextField(), db_column="scenarioIds")
    targets = models.JSONField()
    repeat_count = models.IntegerField(db_column="repeatCount", default=1)
    labels = ArrayField(models.TextField())
    simulator_model = models.TextField(null=True, db_column="simulatorModel")
    judge_model = models.TextField(null=True, db_column="judgeModel")
    kind = models.TextField(default="run_plan")
    scope = models.JSONField(null=True)
    fields = models.JSONField(null=True)
    evaluators = models.JSONField(null=True)
    archived_at = UtcDateTimeField(null=True, db_column="archivedAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "SimulationSuite"
        constraints = [
            models.UniqueConstraint(fields=["project", "slug"], name="SimulationSuite_projectId_slug_key"),
        ]
        indexes = [
            models.Index(fields=["project"]),
            models.Index(fields=["project", "kind"]),
        ]


class Subscription(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    organization = models.ForeignKey("Organization", on_delete=models.RESTRICT, related_name="subscriptions", db_column="organizationId", db_index=False)
    plan = models.TextField(choices=PlanTypes.choices)
    stripe_subscription_id = models.TextField(null=True, unique=True, db_column="stripeSubscriptionId")
    status = models.TextField(choices=SubscriptionStatus.choices)
    start_date = UtcDateTimeField(null=True, db_column="startDate")
    end_date = UtcDateTimeField(null=True, db_column="endDate")
    last_payment_failed_date = UtcDateTimeField(null=True, db_column="lastPaymentFailedDate")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    max_members = models.IntegerField(null=True, db_column="maxMembers")
    max_projects = models.IntegerField(null=True, db_column="maxProjects")
    max_messages_per_month = models.IntegerField(null=True, db_column="maxMessagesPerMonth")
    max_workflows = models.IntegerField(null=True, db_column="maxWorkflows")
    max_retention_days = models.IntegerField(null=True, db_column="maxRetentionDays")
    max_members_lite = models.IntegerField(null=True, db_column="maxMembersLite")
    max_members_developer = models.IntegerField(null=True, db_column="maxMembersDeveloper")
    """Developer seats the organization pays for; null falls to the plan's own cap."""
    max_teams = models.IntegerField(null=True, db_column="maxTeams")
    max_prompts = models.IntegerField(null=True, db_column="maxPrompts")
    max_evaluators = models.IntegerField(null=True, db_column="maxEvaluators")
    max_scenarios = models.IntegerField(null=True, db_column="maxScenarios")
    max_agents = models.IntegerField(null=True, db_column="maxAgents")
    max_experiments = models.IntegerField(null=True, db_column="maxExperiments")
    max_online_evaluations = models.IntegerField(null=True, db_column="maxOnlineEvaluations")
    max_datasets = models.IntegerField(null=True, db_column="maxDatasets")
    max_dashboards = models.IntegerField(null=True, db_column="maxDashboards")
    max_custom_graphs = models.IntegerField(null=True, db_column="maxCustomGraphs")
    max_automations = models.IntegerField(null=True, db_column="maxAutomations")
    seats_kept = models.IntegerField(null=True, db_column="seatsKept")
    """Seats paid for this period that stay usable after a seat decrease, until seatsKeptUntil."""
    seats_kept_until = UtcDateTimeField(null=True, db_column="seatsKeptUntil")
    seat_line_quantity = models.IntegerField(null=True, db_column="seatLineQuantity")
    """The seat quantity the last finished seat change set the seat line to; provider events never write it."""
    seat_line_quantity_at = UtcDateTimeField(null=True, db_column="seatLineQuantityAt")
    seat_line_write_quantities = ArrayField(models.IntegerField(), db_column="seatLineWriteQuantities", default=list)
    """A Stripe seat line write recorded before it is sent and cleared once it is known to have
    finished under the lease; one left behind is settled by the seat line repair.
    Every seat quantity written while this record was pending.
    """
    seat_line_write_since = UtcDateTimeField(null=True, db_column="seatLineWriteSince")
    """The first write still pending; a later holder's write keeps it."""
    seat_line_write_at = UtcDateTimeField(null=True, db_column="seatLineWriteAt")
    """The latest write sent; the repair waits a grace window after it."""
    seat_line_write_attempts = models.IntegerField(null=True, db_column="seatLineWriteAttempts")
    seat_line_write_flagged_at = UtcDateTimeField(null=True, db_column="seatLineWriteFlaggedAt")
    seat_repair_quantity = models.IntegerField(null=True, db_column="seatRepairQuantity")
    """The seat quantity Stripe should bill if a seat change stops after lifting the seat line; null once settled."""
    seat_repair_lifted_quantity = models.IntegerField(null=True, db_column="seatRepairLiftedQuantity")
    """The quantity the change lifted the line to; the repair moves the line back only while Stripe still bills it."""
    seat_repair_since = UtcDateTimeField(null=True, db_column="seatRepairSince")
    seat_repair_attempts = models.IntegerField(null=True, db_column="seatRepairAttempts")
    seat_repair_flagged_at = UtcDateTimeField(null=True, db_column="seatRepairFlaggedAt")
    """Set once the seat line repair has failed often enough that Ops should look at it."""
    seat_line_lease_token = models.TextField(null=True, db_column="seatLineLeaseToken")
    """The lease a seat change or seat line repair holds while it reads and moves the seat line."""
    seat_line_lease_until = UtcDateTimeField(null=True, db_column="seatLineLeaseUntil")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Subscription"
        indexes = [
            models.Index(fields=["organization"]),
        ]


class Invoice(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    subscription = models.ForeignKey("Subscription", on_delete=models.RESTRICT, related_name="invoices", db_column="subscriptionId", db_index=False)
    total_amount = models.FloatField(db_column="totalAmount")
    currency = models.TextField()
    due_date = UtcDateTimeField(db_column="dueDate")
    paid = models.BooleanField(default=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Invoice"
        indexes = [
            models.Index(fields=["subscription"]),
        ]


class InvoiceItem(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    invoice = models.ForeignKey("Invoice", on_delete=models.RESTRICT, related_name="line_items", db_column="invoiceId", db_index=False)
    description = models.TextField()
    amount = models.FloatField()
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "InvoiceItem"
        indexes = [
            models.Index(fields=["invoice"]),
        ]


class Group(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="groups", db_column="organizationId", db_index=False)
    name = models.TextField()
    slug = models.TextField()
    external_id = models.TextField(null=True, db_column="externalId")
    scim_source = models.TextField(null=True, db_column="scimSource")
    scim_connection_id = models.TextField(null=True, db_column="scimConnectionId")
    """D08: which SSO connection provisioned this group. Null for manually
    created groups and for the SCIM groups that predate connection scoping.
    """
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Group"
        constraints = [
            models.UniqueConstraint(fields=["organization", "slug"], name="Group_organizationId_slug_key"),
            models.UniqueConstraint(fields=["scim_connection_id", "external_id"], name="Group_scimConnectionId_externalId_key"),
        ]
        indexes = [
            models.Index(fields=["organization"]),
            models.Index(fields=["scim_connection_id"]),
        ]


class GroupMembership(models.Model):
    pk = models.CompositePrimaryKey("user_id", "group_id")
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="group_memberships", db_column="userId", db_index=False)
    group = models.ForeignKey("Group", on_delete=models.CASCADE, related_name="members", db_column="groupId", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "GroupMembership"
        indexes = [
            models.Index(fields=["user"]),
            models.Index(fields=["group"]),
        ]


class AuthzUserStanding(models.Model):
    """Who authz treats as gone, from user's and identity's facts; authz never reads User (ARCHITECTURE.md, "Platform operators are a grant")."""

    user_id = models.TextField(primary_key=True, db_column="userId", default=None)
    deactivated_at = UtcDateTimeField(null=True, db_column="deactivatedAt")
    erased_at = UtcDateTimeField(null=True, db_column="erasedAt")
    standing_changed_at = UtcDateTimeField(db_column="standingChangedAt")

    class Meta:
        db_table = "AuthzUserStanding"


class AuthzEngineSwitch(models.Model):
    """Whether an organization's permission checks run on the authorization engine: authz's own
    switch, set when the engine migration proves the organization and pinned off by an operator's
    rollback. No row is the legacy path. Spec: modules/authz/specs/engine-switch.feature.
    """

    organization_id = models.TextField(primary_key=True, db_column="organizationId", default=None)
    on_engine = models.BooleanField(db_column="onEngine")
    pinned = models.BooleanField(default=False)
    switched_at = UtcDateTimeField(db_column="switchedAt")
    last_applied_event_id = models.TextField(null=True, db_column="lastAppliedEventId")
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "AuthzEngineSwitch"


class AuthzOrganizationPolicy(models.Model):
    """ADR-201: an organisation's braid setting. Owned by authz; read into the organisation's snapshot."""

    organization_id = models.TextField(primary_key=True, db_column="organizationId", default=None)
    braids_organization_admins_only = models.BooleanField(db_column="braidsOrganizationAdminsOnly", default=False)
    updated_by_user_id = models.TextField(db_column="updatedByUserId")
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "AuthzOrganizationPolicy"


class RoleBinding(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="role_bindings", db_column="organizationId", db_index=False)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="role_bindings", db_column="userId", null=True, db_index=False)
    group = models.ForeignKey("Group", on_delete=models.CASCADE, related_name="role_bindings", db_column="groupId", null=True, db_index=False)
    api_key = models.ForeignKey("ApiKey", on_delete=models.RESTRICT, related_name="role_bindings", db_column="apiKeyId", null=True, db_index=False)
    role = models.TextField(choices=TeamUserRole.choices)
    custom_role = models.ForeignKey("CustomRole", on_delete=models.SET_NULL, related_name="role_bindings", db_column="customRoleId", null=True, db_index=False)
    scope_type = models.TextField(choices=RoleBindingScopeType.choices, db_column="scopeType")
    scope_id = models.TextField(db_column="scopeId")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "RoleBinding"
        indexes = [
            models.Index(fields=["organization"]),
            models.Index(fields=["user"]),
            models.Index(fields=["group"]),
            models.Index(fields=["api_key"]),
            models.Index(fields=["scope_type", "scope_id"]),
        ]


class Grant(models.Model):
    id = models.TextField(primary_key=True, default=None)
    organization_id = models.TextField(db_column="organizationId")
    principal_type = models.TextField(choices=GrantPrincipalType.choices, db_column="principalType")
    principal_id = models.TextField(null=True, db_column="principalId")
    role_key = models.TextField(null=True, db_column="roleKey")
    legacy_role = models.TextField(null=True, db_column="legacyRole")
    source = models.TextField()
    scope_type = models.TextField(choices=GrantScopeType.choices, db_column="scopeType")
    scope_id = models.TextField(db_column="scopeId")
    token = models.TextField(null=True, unique=True)
    permission = models.TextField(null=True)
    resource_kind = models.TextField(null=True, db_column="resourceKind")
    project_id = models.TextField(null=True, db_column="projectId")
    created_by_user_id = models.TextField(null=True, db_column="createdByUserId")
    expires_at = UtcDateTimeField(null=True, db_column="expiresAt")
    max_views = models.IntegerField(null=True, db_column="maxViews")
    condition = models.JSONField(null=True)
    replaces_grant_id = models.TextField(null=True, db_column="replacesGrantId")
    occurred_at = UtcDateTimeField(db_column="occurredAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    revoked_at = UtcDateTimeField(null=True, db_column="revokedAt")
    revoked_reason = models.TextField(null=True, db_column="revokedReason")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Grant"
        indexes = [
            models.Index(fields=["organization_id", "scope_type", "scope_id"]),
            models.Index(fields=["project_id", "resource_kind", "scope_id"]),
        ]


class GrantUsage(models.Model):
    grant_id = models.TextField(primary_key=True, db_column="grantId", default=None)
    organization_id = models.TextField(db_column="organizationId")
    project_id = models.TextField(db_column="projectId")
    view_count = models.IntegerField(db_column="viewCount", default=0)
    last_viewed_at = UtcDateTimeField(null=True, db_column="lastViewedAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "GrantUsage"
        indexes = [
            models.Index(fields=["organization_id"]),
            models.Index(fields=["project_id"]),
        ]


class Role(models.Model):
    id = models.TextField(primary_key=True, default=None)
    organization_id = models.TextField(db_column="organizationId")
    name = models.TextField()
    description = models.TextField(null=True)
    permissions = models.JSONField()
    kind = models.TextField(default="custom")
    occurred_at = UtcDateTimeField(db_column="occurredAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    deleted_at = UtcDateTimeField(null=True, db_column="deletedAt")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Role"
        indexes = [
            models.Index(fields=["organization_id"]),
        ]


class ApiKey(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    name = models.TextField()
    description = models.TextField(null=True)
    lookup_id = models.TextField(unique=True, db_column="lookupId")
    hashed_secret = models.TextField(db_column="hashedSecret")
    permission_mode = models.TextField(db_column="permissionMode", default="all")
    user = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="api_keys", db_column="userId", null=True, db_index=False)
    created_by_user_id = models.TextField(null=True, db_column="createdByUserId")
    organization = models.ForeignKey("Organization", on_delete=models.RESTRICT, related_name="api_keys", db_column="organizationId", db_index=False)
    expires_at = UtcDateTimeField(null=True, db_column="expiresAt")
    last_used_at = UtcDateTimeField(null=True, db_column="lastUsedAt")
    revoked_at = UtcDateTimeField(null=True, db_column="revokedAt")
    revocation_cause = models.TextField(null=True, db_column="revocationCause")
    """Why the key was revoked, set with `revokedAt`: "user" when a person
    revoked it (the API-keys page, the REST API, the devices tab), "rotation"
    when a mint replaced it, "session" when the CLI login key it was parented
    to was revoked, "expired" when that session ran out, "offboarded" when the
    person's membership of the organization ended. "cap" survives on rows the
    retired per-tool cap revoked. The CLI reads it to decide whether a dead
    ingest key may be re-minted: a key a person revoked stays dead.
    NULL for live keys and for keys revoked before the cause was recorded.
    """
    parent_api_key_id = models.TextField(null=True, db_column="parentApiKeyId")
    """The CLI login key of the device session that minted this ingestion key.
    Set by the CLI's personal mint, NULL for every other key. Revoking the
    parent (logout, the devices tab, a re-login from the same device, session
    expiry) revokes the children with it. Column ref only, no FK and no
    index: lookups filter by the indexed `userId` first and match the parent
    in memory over a person's few live ingest keys.
    """
    ingest_source_type = models.TextField(null=True, db_column="ingestSourceType")
    """Non-null marks this key as an "ingestion key": a project-scoped,
    ingest-only ApiKey the unified `langwatch <tool>` CLI (and template
    installs) mint for OTLP Path B. The slug (claude_code / codex / gemini /
    opencode / claude_cowork) is stamped as `langwatch.source` provenance on
    every span the key authorizes. NULL for normal API / service keys.
    """
    ingestion_template_id = models.TextField(null=True, db_column="ingestionTemplateId")
    """Set only for ingestion keys derived from an admin IngestionTemplate
    (e.g. claude_cowork). Drives receiver OTTL application + the
    `langwatch.template.id` provenance stamp. NULL for unified-CLI tools
    (which have no template) and for normal keys. Column ref only (no FK,
    per the app-layer-integrity convention); resolved by the ingest-key service.
    """
    created_by_device_label = models.TextField(null=True, db_column="createdByDeviceLabel")
    """Human label of the CLI device session that minted this ingestion key
    (e.g. "Rogerio's MacBook Pro"), snapshotted at mint time from the device
    flow's `client_info`. Denormalized so the org-wide API-keys settings page
    can show which device a key came from without resolving another user's
    Redis-stored sessions. NULL for non-ingestion keys and for keys minted by
    a CLI old enough to predate device metadata.
    """
    is_system_managed = models.BooleanField(db_column="isSystemManaged", default=False)
    """True for a key LangWatch minted for itself (a run, the gateway's trace
    export, a Langy session). Hides it from listings, refuses customer edits
    and revokes, and lets an ownerless run key act as the system. Never set
    by a customer create; rows before the column stay false.
    """
    created_by_browser_session_id = models.TextField(null=True, db_column="createdByBrowserSessionId")
    """The browser session that minted a docs session key (ADR-185), so signing
    out of that session revokes it. NULL for every other key. Unindexed:
    sign-out filters by the indexed `userId` first.
    """
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "ApiKey"
        indexes = [
            models.Index(fields=["user"]),
            models.Index(fields=["organization"]),
            models.Index(fields=["lookup_id"]),
            models.Index(fields=["ingest_source_type"]),
            models.Index(fields=["ingestion_template_id"]),
        ]


class SigningKey(models.Model):
    """ADR-201: a project's key for signing the tags on its calls. Owned by the API key module."""

    id = models.TextField(primary_key=True, default=None)
    organization_id = models.TextField(db_column="organizationId")
    project_id = models.TextField(db_column="projectId")
    name = models.TextField()
    public_id = models.TextField(unique=True, db_column="publicId")
    secret_ciphertext = models.TextField(db_column="secretCiphertext")
    created_by_user_id = models.TextField(db_column="createdByUserId")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    rotated_at = UtcDateTimeField(null=True, db_column="rotatedAt")
    verifies_until = UtcDateTimeField(null=True, db_column="verifiesUntil")
    revoked_at = UtcDateTimeField(null=True, db_column="revokedAt")
    last_used_at = UtcDateTimeField(null=True, db_column="lastUsedAt")

    class Meta:
        db_table = "SigningKey"
        indexes = [
            models.Index(fields=["project_id"]),
            models.Index(fields=["organization_id"]),
        ]


class ProjectActiveDay(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    project_id = models.TextField(db_column="projectId")
    day = models.TextField()
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "ProjectActiveDay"
        constraints = [
            models.UniqueConstraint(fields=["project_id", "day"], name="ProjectActiveDay_projectId_day_key"),
        ]


class NurturingProject(models.Model):
    """Nurturing's read model: which organization a project belongs to, from project's created event.
    The per-email usage places a project and lists an organization's projects from these rows only.
    """

    project_id = models.TextField(primary_key=True, db_column="projectId", default=None)
    organization_id = models.TextField(db_column="organizationId")

    class Meta:
        db_table = "NurturingProject"
        indexes = [
            models.Index(fields=["organization_id"]),
        ]


class NurturingUncopiedProject(models.Model):
    """A project nurturing saw traces of with no placement copy, and when it first did: every worker
    retries its trace batches only within one window from then, across restarts. Deleted when the
    copy lands.
    """

    project_id = models.TextField(primary_key=True, db_column="projectId", default=None)
    first_seen_at = UtcDateTimeField(db_column="firstSeenAt", default=now_milliseconds)

    class Meta:
        db_table = "NurturingUncopiedProject"


class NurturingOrganizationMember(models.Model):
    """Nurturing's own copy of each organization's members, from the owners' member facts (§9):
    organization's signed_up and invite_accepted and auth's sso_auto_added join, member_removed
    and member_disabled leave, member_enabled returns. Each column keeps its latest time.
    """

    pk = models.CompositePrimaryKey("organization_id", "user_id")
    organization_id = models.TextField(db_column="organizationId")
    user_id = models.TextField(db_column="userId")
    joined_at = UtcDateTimeField(null=True, db_column="joinedAt")
    left_at = UtcDateTimeField(null=True, db_column="leftAt")
    returned_at = UtcDateTimeField(null=True, db_column="returnedAt")

    class Meta:
        db_table = "NurturingOrganizationMember"
        indexes = [
            models.Index(fields=["user_id"]),
        ]


class NurturingOrganizationCreated(models.Model):
    """When an organization was created, from organization's created event: nurturing's own copy for
    the per-email usage's day offsets. Its own table, so an early creation never adds an unseeded
    NurturingOrganization row.
    """

    organization_id = models.TextField(primary_key=True, db_column="organizationId", default=None)
    created_at = UtcDateTimeField(db_column="createdAt")

    class Meta:
        db_table = "NurturingOrganizationCreated"
        indexes = [
            models.Index(fields=["created_at"]),
        ]


class NurturingOrganization(models.Model):
    """Nurturing's read model per organization: its admin and its milestone counts.
    `seeded` marks an organization learned from the project backfill: its first_* already happened.
    """

    organization_id = models.TextField(primary_key=True, db_column="organizationId", default=None)
    admin_user_id = models.TextField(null=True, db_column="adminUserId")
    seeded = models.BooleanField(default=False)
    evaluation_count = models.IntegerField(db_column="evaluationCount", default=0)
    simulation_run_count = models.IntegerField(db_column="simulationRunCount", default=0)
    coding_first_session_at = UtcDateTimeField(null=True, db_column="codingFirstSessionAt")
    coding_agent_types = ArrayField(models.TextField(), db_column="codingAgentTypes", default=list)
    coding_active_days = models.IntegerField(db_column="codingActiveDays", default=0)
    coding_last_active_day = models.TextField(null=True, db_column="codingLastActiveDay")
    coding_terminal_types = ArrayField(models.TextField(), db_column="codingTerminalTypes", default=list)
    coding_repository_keys = ArrayField(models.TextField(), db_column="codingRepositoryKeys", default=list)
    coding_first_days = ArrayField(models.IntegerField(), db_column="codingFirstDays", default=list)
    last_trace_at = UtcDateTimeField(null=True, db_column="lastTraceAt")
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "NurturingOrganization"


class NurturingProjectUsageView(models.Model):
    """Nurturing's own per-project usage: a peer fold over the owners' prompt, evaluator, scenario
    and simulation-run events (ARCHITECTURE §9), one peer fold per owner pipeline, summed per
    organization for the per-email traits. Traces are counted in NurturingProjectTraceDay.
    """

    pk = models.CompositePrimaryKey("project_id", "owner")
    project_id = models.TextField(db_column="projectId")
    owner = models.TextField()
    """The owner pipeline whose peer fold wrote the row (prompt, evaluator, scenario, simulation),
    or `evaluation`: the project's evaluation days, written by a subscriber, at most once a day.
    """
    folded = models.JSONField()
    """The project's usage as the fold folded it (NurturingProjectUsageState)."""
    applied_event_ids = ArrayField(models.TextField(), db_column="appliedEventIds", default=list)
    """The event ids of the last delivery folded in, so a redelivery folds once."""
    projection_version = models.TextField(db_column="projectionVersion")
    """A row of another version reads as undecodable and is re-folded from the log."""
    updated_at = UtcDateTimeField(db_column="updatedAt")
    """Event-derived, so a replayed row equals the row it rebuilds."""

    class Meta:
        db_table = "NurturingProjectUsageView"


class NurturingProjectTraceDay(models.Model):
    """Nurturing's own trace counts from trace's span_received and trace_failed (ARCHITECTURE §9), a
    root span being one trace, one row per project and UTC day of append time, incremented in place
    by each batch. A redelivered batch can count a trace twice; accepted (issue 9378). Kept 30 days.
    """

    pk = models.CompositePrimaryKey("project_id", "day")
    project_id = models.TextField(db_column="projectId")
    day = models.TextField()
    """UTC `YYYY-MM-DD` of the root spans' append time."""
    trace_count = models.IntegerField(db_column="traceCount", default=0)
    error_count = models.IntegerField(db_column="errorCount", default=0)
    """Traces trace recorded failed (lw.trace.trace_failed)."""
    last_trace_at = UtcDateTimeField(null=True, db_column="lastTraceAt")
    """Append time of the newest root span."""
    last_failed_trace_at = UtcDateTimeField(null=True, db_column="lastFailedTraceAt")
    """Append time of the newest trace_failed."""

    class Meta:
        db_table = "NurturingProjectTraceDay"


class NurturingPersonalProject(models.Model):
    """Nurturing's copy of each personal project's owner, from organization's and project's facts.
    The first owner recorded stays: a personal project's owner never changes.
    """

    project_id = models.TextField(primary_key=True, db_column="projectId", default=None)
    organization_id = models.TextField(db_column="organizationId")
    owner_user_id = models.TextField(db_column="ownerUserId")

    class Meta:
        db_table = "NurturingPersonalProject"


class NurturingTracedProject(models.Model):
    """Nurturing's copy of each project that received a real trace, from trace's and project's facts."""

    project_id = models.TextField(primary_key=True, db_column="projectId", default=None)
    organization_id = models.TextField(db_column="organizationId")

    class Meta:
        db_table = "NurturingTracedProject"
        indexes = [
            models.Index(fields=["organization_id"]),
        ]


class BillingMeterCheckpoint(models.Model):
    id = models.TextField(primary_key=True, default=cuid_str)
    organization_id = models.TextField(db_column="organizationId")
    billing_month = models.TextField(db_column="billingMonth")
    meter = models.TextField(default="langwatch_billable_events")
    last_reported_total = models.IntegerField(db_column="lastReportedTotal", default=0)
    pending_reported_total = models.IntegerField(null=True, db_column="pendingReportedTotal")
    consecutive_failures = models.IntegerField(db_column="consecutiveFailures", default=0)
    last_counted_event_id = models.TextField(null=True, db_column="lastCountedEventId")
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "BillingMeterCheckpoint"
        constraints = [
            models.UniqueConstraint(fields=["organization_id", "billing_month", "meter"], name="BillingMeterCheckpoint_organizationId_billingMonth_meter_key"),
        ]


class VirtualKey(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    organization_id = models.TextField(db_column="organizationId")
    name = models.TextField()
    description = models.TextField(null=True)
    status = models.TextField(choices=VirtualKeyStatus.choices, default=VirtualKeyStatus.ACTIVE)
    purpose = models.TextField(choices=VirtualKeyPurpose.choices, default=VirtualKeyPurpose.USER)
    external_id = models.TextField(null=True, db_column="externalId")
    metadata = models.JSONField(default=dict)
    disabled_at = UtcDateTimeField(null=True, db_column="disabledAt")
    disabled_reason = models.TextField(null=True, db_column="disabledReason")
    expires_at = UtcDateTimeField(null=True, db_column="expiresAt")
    hashed_secret = models.TextField(unique=True, db_column="hashedSecret")
    display_prefix = models.TextField(db_column="displayPrefix")
    principal_user = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="principal_virtual_keys", db_column="principalUserId", null=True, db_index=False)
    trace_project_id = models.TextField(null=True, db_column="traceProjectId")
    config = models.JSONField(default=dict)
    revision = models.BigIntegerField(default=0)
    previous_hashed_secret = models.TextField(null=True, db_column="previousHashedSecret")
    previous_secret_valid_until = UtcDateTimeField(null=True, db_column="previousSecretValidUntil")
    revoked_at = UtcDateTimeField(null=True, db_column="revokedAt")
    revoked_by_id = models.TextField(null=True, db_column="revokedById")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    created_by_id = models.TextField(db_column="createdById")
    last_used_at = UtcDateTimeField(null=True, db_column="lastUsedAt")
    routing_policy = models.ForeignKey("RoutingPolicy", on_delete=models.SET_NULL, related_name="virtual_keys", db_column="routingPolicyId", null=True, db_index=False)
    routing_mode = models.TextField(choices=VirtualKeyRoutingMode.choices, db_column="routingMode", default=VirtualKeyRoutingMode.NONE)
    connect_services = ArrayField(models.TextField(), db_column="connectServices", default=list)
    license_token_hash = models.TextField(null=True, unique=True, db_column="licenseTokenHash")
    license_instance_id = models.TextField(null=True, db_column="licenseInstanceId")
    license_expires_at = UtcDateTimeField(null=True, db_column="licenseExpiresAt")
    license_id = models.TextField(null=True, db_column="licenseId")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "VirtualKey"
        constraints = [
            models.UniqueConstraint(fields=["hashed_secret"], name="VirtualKey_hashedSecret_key"),
            models.UniqueConstraint(fields=["organization_id", "external_id"], name="VirtualKey_organizationId_externalId_key"),
        ]
        indexes = [
            models.Index(fields=["organization_id"]),
            models.Index(fields=["display_prefix"]),
            models.Index(fields=["status"]),
            models.Index(fields=["principal_user"]),
            models.Index(fields=["revision"]),
            models.Index(fields=["routing_policy"]),
            models.Index(fields=["organization_id", "purpose"]),
        ]


class VirtualKeyScope(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    virtual_key = models.ForeignKey("VirtualKey", on_delete=models.CASCADE, related_name="scopes", db_column="virtualKeyId", db_index=False)
    scope_type = models.TextField(choices=VirtualKeyScopeType.choices, db_column="scopeType")
    scope_id = models.TextField(db_column="scopeId")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "VirtualKeyScope"
        constraints = [
            models.UniqueConstraint(fields=["virtual_key", "scope_type", "scope_id"], name="VirtualKeyScope_virtualKeyId_scopeType_scopeId_key"),
        ]
        indexes = [
            models.Index(fields=["scope_type", "scope_id"]),
            models.Index(fields=["virtual_key"]),
        ]


class RoutingPolicy(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="routing_policies", db_column="organizationId", db_index=False)
    name = models.TextField()
    description = models.TextField(null=True)
    model_provider_ids = models.JSONField(db_column="modelProviderIds", default=list)
    model_aliases = models.JSONField(db_column="modelAliases", default=dict)
    default_model = models.TextField(null=True, db_column="defaultModel")
    policy_rules = models.JSONField(db_column="policyRules", default=dict)
    is_default = models.BooleanField(db_column="isDefault", default=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    created_by = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="routing_policies_created", db_column="createdById", null=True, db_index=False)
    updated_by = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="routing_policies_updated", db_column="updatedById", null=True, db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "RoutingPolicy"
        indexes = [
            models.Index(fields=["organization"]),
        ]


class RoutingPolicyScope(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    routing_policy = models.ForeignKey("RoutingPolicy", on_delete=models.CASCADE, related_name="scopes", db_column="routingPolicyId", db_index=False)
    scope_type = models.TextField(choices=RoutingPolicyScopeType.choices, db_column="scopeType")
    scope_id = models.TextField(db_column="scopeId")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "RoutingPolicyScope"
        constraints = [
            models.UniqueConstraint(fields=["routing_policy", "scope_type", "scope_id"], name="RoutingPolicyScope_routingPolicyId_scopeType_scopeId_key"),
        ]
        indexes = [
            models.Index(fields=["scope_type", "scope_id"]),
            models.Index(fields=["routing_policy"]),
        ]


class GatewayGuardrail(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="gateway_guardrails", db_column="projectId", db_index=False)
    name = models.TextField()
    description = models.TextField(null=True)
    evaluator = models.ForeignKey("Evaluator", on_delete=models.RESTRICT, related_name="gateway_guardrails", db_column="evaluatorId", db_index=False)
    direction = models.TextField(choices=GatewayGuardrailDirection.choices)
    failure_mode = models.TextField(choices=GatewayGuardrailFailureMode.choices, db_column="failureMode", default=GatewayGuardrailFailureMode.FAIL_CLOSED)
    archived_at = UtcDateTimeField(null=True, db_column="archivedAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    created_by = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="gateway_guardrails_created", db_column="createdById", null=True, db_index=False)
    updated_by = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="gateway_guardrails_updated", db_column="updatedById", null=True, db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "GatewayGuardrail"
        indexes = [
            models.Index(fields=["project"]),
            models.Index(fields=["evaluator"]),
            models.Index(fields=["project", "archived_at"]),
        ]


class IngestionSource(models.Model):
    """Per-platform fleet configuration that connects a closed SaaS
    product's audit / OTel / S3 stream to LangWatch's Activity
    Monitor (cf. specs/ai-gateway/governance/activity-monitor.feature
    + ingestion-sources.feature, docs/ai-gateway/governance/architecture.md).

    One IngestionSource = one upstream fleet (e.g. "Miro Cowork",
    "Acme Workato production"). The runtime ingests events on whatever
    cadence the source supports - push for OTel/webhook/S3 (Tier C/D),
    pull for poll-based admin APIs (Tier C). Per-source-type adapters
    normalise platform-specific shapes into OCSF + AOS before they
    land in trace_summaries with SourceType + SourceId tags.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="ingestion_sources", db_column="organizationId", db_index=False)
    team = models.ForeignKey("Team", on_delete=models.SET_NULL, related_name="ingestion_sources", db_column="teamId", null=True, db_index=False)
    """Optional team-scope. null = org-wide."""
    source_type = models.TextField(db_column="sourceType")
    """'otel_generic' | 'claude_cowork' | 'workato' | 'copilot_studio' |
    'openai_compliance' | 'claude_compliance' | 's3_custom' (extensible).
    String not enum so adapters can be added without migrations - the
    supported set lives in code (activity-monitor adapter registry).
    """
    name = models.TextField()
    """Human-friendly label for the admin UI. Unique per-org so two
    "Cowork" sources can coexist with different names.
    """
    description = models.TextField(null=True)
    ingest_secret_hash = models.TextField(db_column="ingestSecretHash")
    """HMAC bearer token the upstream operator pastes into the platform
    (or LangWatch sends in pull mode). Stored encrypted at rest via
    the same pepper-derived KEK as VirtualKey secrets. Auto-generated
    on create, rotatable via the admin UI with a 24h grace window.
    """
    parser_config = models.JSONField(db_column="parserConfig", default=dict)
    """Per-source-type configuration. Shape is platform-specific:
    - otel_generic: { allowedSourceTypeLabel?: string }
    - claude_cowork: { workspaceId: string }
    - workato: { sharedSecretLastFour: string }
    - copilot_studio: { tenantId, clientId, clientSecretHash, pollEverySec }
    - openai_compliance: { bucket, prefix, roleArn, pollEverySec }
    - claude_compliance: { workspaceApiKeyHash, pollEverySec }
    - s3_custom: { bucket, prefix, roleArn, parserDsl, pollEverySec }
    """
    poller_cursor = models.JSONField(null=True, db_column="pollerCursor")
    """For pull-based sources, persisted between polls so the next poll
    resumes where the last left off (timestamp / cursor / file marker).
    """
    error_count = models.IntegerField(db_column="errorCount", default=0)
    """Phase 10 - number of consecutive errors observed by the puller
    worker. Reset to 0 on a successful run. Drives admin-UI banners
    ("Microsoft authentication failed - re-authorize") + back-off
    scheduling when the source has been broken for a while.
    Spec: specs/ai-governance/puller-framework/puller-adapter-contract.feature
    """
    pull_schedule = models.TextField(null=True, db_column="pullSchedule")
    """Phase 10 - cron schedule for pull-mode sources. Validated by the
    adapter's pullConfig schema; mirrored here for fast worker
    dispatch. Null for push-mode sources.
    """
    status = models.TextField(default="awaiting_first_event")
    """Status flag. 'active' (default), 'disabled' (admin paused), or
    'awaiting_first_event' (created but never received).
    """
    trace_project_id = models.TextField(null=True, db_column="traceProjectId")
    """ADR-088 v7: optional trace DESTINATION for conversation-bearing pulls
    (Databricks Genie). Null = don't route. Same shape as
    VirtualKey.traceProjectId: a destination is NOT a scope row and grants
    no access to the source; it only decides where routed conversation
    traces land. Must be a project of this org (validated at write time).
    Freely editable; already-routed history stays where it landed.
    NOT the deferred cost-attribution projectId (ADR-088 Decision 4).
    """
    provider_account_id = models.TextField(null=True, db_column="providerAccountId")
    """The account the provider itself named when this connection was saved -
    an OpenAI or Anthropic organization id, read once per save and compared
    so a second connection cannot read one account's spend twice. NOT
    derived from the administrator key: nothing worked out from a customer
    credential is stored here or anywhere else, because the account id is
    not a secret and cannot be turned back into a key.
    Nullable because every connection that predates the guard has none, and
    because Azure and Power Platform name what they read in parserConfig
    instead. "Unknown account" has to be a storable state.
    Spec: specs/ai-gateway/governance/ingestion-sources.feature
    """
    last_event_at = UtcDateTimeField(null=True, db_column="lastEventAt")
    """Last event timestamp observed (any direction). Surfaces in the
    admin UI as "last event 2 minutes ago".
    """
    last_success_at = UtcDateTimeField(null=True, db_column="lastSuccessAt")
    """ADR-128 - when a pull run last reached the provider without erroring,
    mirrored from the ingestionPullRunStatus projection. Deliberately NOT
    lastEventAt: an empty period is a working puller, and a run that
    delivered data and then failed is not. Drives the health badge and
    the "no data since" cut-off, below which a day's spend is unknown
    rather than zero.
    Spec: specs/governance/ingestion-source-health.feature
    """
    last_read_through_at = UtcDateTimeField(null=True, db_column="lastReadThroughAt")
    """ADR-128 - how far the last run actually read, and whether it reached the
    end. Mirrored from the ingestionPullRunStatus projection, in the same way
    and for the same reason as lastSuccessAt above: every reader of this
    surface selects from this row, and two of them render a LIST of sources,
    so reaching a per-run projection would be one query per row on a screen
    built to show many at once.

    These two are the only evidence separating a source that truncated once
    and then finished from one that is stuck. A stuck source's run clock
    advances on every attempt, so the clock cannot be that evidence and is
    deliberately not stored here - only the point the run REACHED.

    Completeness is "complete" or "truncated". Null on both is unknown, which
    is what every row written before this existed reads as, and unknown must
    never render as either answer.
    Spec: specs/governance/ingestion-source-health.feature
    """
    last_run_completeness = models.TextField(null=True, db_column="lastRunCompleteness")
    unpriced_usage_since = UtcDateTimeField(null=True, db_column="unpricedUsageSince")
    """ADR-088 - the window this source read while pulled cost recording was
    off for the organization, so its spend was dropped rather than priced.
    Both null = nothing dropped. The pull cursor advances whether or not the
    money path is live (audit-only is a supported mode), so without this the
    loss is unrecoverable AND invisible: those days read as no spend rather
    than as unknown. Widened by any run that drops a priced item, cleared
    only when a later run re-reads back across the whole window.
    Spec: specs/governance/pulled-usage-cost-gap.feature
    """
    unpriced_usage_through = UtcDateTimeField(null=True, db_column="unpricedUsageThrough")
    archived_at = UtcDateTimeField(null=True, db_column="archivedAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    created_by = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="ingestion_sources_created", db_column="createdById", null=True, db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "IngestionSource"
        constraints = [
            models.UniqueConstraint(fields=["organization", "name"], name="IngestionSource_organizationId_name_key"),
        ]
        indexes = [
            models.Index(fields=["organization"]),
            models.Index(fields=["team"]),
            models.Index(fields=["source_type"]),
        ]


class AnomalyRule(models.Model):
    """Admin-defined anomaly detection rule. Evaluated server-side at
    ingest time (Option C - not in this slice yet) against
    gateway_activity_events. When triggered, an alert is dispatched
    to the destinations configured on the rule (Slack / SIEM / webhook
    / PagerDuty / email).

    This is a configuration entity - the actual rule evaluation +
    alert state lives in a future ClickHouse table
    (gateway_anomaly_alerts) and a server-side worker. Spec:
    specs/ai-gateway/governance/anomaly-rules.feature
    (Alexis lane B C-UI; mocked-v0 today, real persistence is THIS slice).
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="anomaly_rules", db_column="organizationId", db_index=False)
    scope = models.TextField()
    """'organization' | 'team' | 'project' | 'source_type' | 'source'.
    String not enum so new scopes ship without migrations.
    """
    scope_id = models.TextField(db_column="scopeId")
    """Org id / team id / project id / source-type string / source id
    - narrows what the rule applies to.
    """
    name = models.TextField()
    description = models.TextField(null=True)
    severity = models.TextField()
    """'critical' | 'warning' | 'info'. Drives default destination
    routing (e.g. critical → PagerDuty, warning → Slack only).
    """
    rule_type = models.TextField(db_column="ruleType")
    """Extensible rule taxonomy: 'spend_spike' | 'rate_limit' |
    'tool_mismatch' | 'after_hours' | 'unusual_model' | 'pii_leak'
    | 'custom'. New types ship without migrations - evaluator
    picks a per-type implementation from a code registry.
    """
    threshold_config = models.JSONField(db_column="thresholdConfig", default=dict)
    """Per-rule-type threshold shape. Examples:
    spend_spike: { windowSec, ratioVsBaseline, minBaselineUsd }
    rate_limit: { windowSec, requestsThreshold }
    after_hours: { startHour, endHour, timezone, requestsThreshold }
    """
    destination_config = models.JSONField(db_column="destinationConfig", default=dict)
    """Alert routing per destination. Shape:
    { slack?: { channelId }, siem?: { endpointUrl, authHeader },
    webhook?: { url, authHeader }, pagerduty?: { routingKey },
    email?: { to: string[] } }
    Empty map = use org defaults (when defaults exist; today no-op).
    """
    status = models.TextField(default="active")
    """'active' (default - evaluator picks it up) | 'disabled' (admin
    paused; rule rows kept for audit). 'archived' uses archivedAt
    timestamp instead so we can filter `archivedAt IS NULL` cheaply.
    """
    archived_at = UtcDateTimeField(null=True, db_column="archivedAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    created_by = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="anomaly_rules_created", db_column="createdById", null=True, db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "AnomalyRule"
        constraints = [
            models.UniqueConstraint(fields=["organization", "name"], name="AnomalyRule_organizationId_name_key"),
        ]
        indexes = [
            models.Index(fields=["organization"]),
            models.Index(fields=["scope", "scope_id"]),
            models.Index(fields=["rule_type"]),
            models.Index(fields=["status"]),
        ]


class AnomalyAlert(models.Model):
    """Persisted anomaly detection. Inserted by the anomaly-detection
    reactor (event-sourcing pipeline; modeled on PR #3351's
    alertTrigger pattern) when a rule's threshold trips while
    processing a freshly-arrived ActivityEventReceived event.
    Read by api.activityMonitor.recentAnomalies for the admin dashboard.

    Dedup: (ruleId, triggerWindowStart) - re-evaluating the same
    window updates the existing row rather than creating a duplicate.
    Acknowledgement / resolution are admin actions on the row.

    Spec: specs/ai-gateway/governance/anomaly-detection.feature
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="anomaly_alerts", db_column="organizationId", db_index=False)
    rule = models.ForeignKey("AnomalyRule", on_delete=models.CASCADE, related_name="alerts", db_column="ruleId", db_index=False)
    severity = models.TextField()
    """Denormalised from rule for fast WHERE filtering on the
    dashboard query (no join needed for the common path).
    """
    rule_name = models.TextField(db_column="ruleName")
    rule_type = models.TextField(db_column="ruleType")
    trigger_window_start = UtcDateTimeField(db_column="triggerWindowStart")
    """Window the trigger refers to (UTC)."""
    trigger_window_end = UtcDateTimeField(db_column="triggerWindowEnd")
    trigger_spend_usd = models.DecimalField(max_digits=18, decimal_places=6, null=True, db_column="triggerSpendUsd")
    """Numeric trigger context. Per-rule-type detail in `detail` JSONB."""
    trigger_event_count = models.IntegerField(null=True, db_column="triggerEventCount")
    detail = models.JSONField(default=dict)
    """Per-rule-type extra context: {topActors: [...], baseline: ...,
    observed: ..., timezone: ..., etc}.
    """
    state = models.TextField(default="open")
    """'open' | 'acknowledged' | 'resolved'."""
    destination_status = models.JSONField(db_column="destinationStatus", default=dict)
    """Per-destination dispatch outcome:
    { webhook: { lastAttemptIso, lastError?, attemptCount, statusCode? },
    slack: ..., pagerduty: ..., siem: ..., email: ... }
    """
    detected_at = UtcDateTimeField(db_column="detectedAt", default=now_milliseconds)
    acknowledged_at = UtcDateTimeField(null=True, db_column="acknowledgedAt")
    resolved_at = UtcDateTimeField(null=True, db_column="resolvedAt")

    class Meta:
        db_table = "AnomalyAlert"
        constraints = [
            models.UniqueConstraint(fields=["rule", "trigger_window_start"], name="AnomalyAlert_ruleId_triggerWindowStart_key"),
        ]
        indexes = [
            models.Index(fields=["organization", "detected_at"]),
            models.Index(fields=["rule"]),
            models.Index(fields=["state"]),
        ]


class AiToolEntry(models.Model):
    """Org-scoped catalog of AI tools surfaced on the personal home portal
    (/me) and managed by org admins at /governance/tool-catalog.

    Three first-class tile types via `type` discriminator:
    - "coding_assistant" - click expands to setup helper (e.g. `langwatch claude`)
    - "model_provider"   - click expands to inline personal-VK creation
    - "external_tool"    - click expands to admin-attached markdown + external link

    Scoping: organization-default entries can be OVERRIDDEN at the team
    tier (same `slug` + `scope=team` + matching `scopeId`) so admins can
    hide or re-skin a globally-listed tool for a specific team. The
    user-facing `list` endpoint resolves this hierarchy at read time.

    `config` is a discriminated JSON union per `type`:
    coding_assistant: { setupCommand, setupDocsUrl, helperText? }
    model_provider:   { providerKey, suggestedRoutingPolicyId?, defaultLabel?, projectSuggestionText? }
    external_tool:    { descriptionMarkdown, linkUrl, ctaLabel? }

    Spec: specs/ai-governance/personal-portal/tool-catalog-*.feature
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="ai_tool_entries", db_column="organizationId", db_index=False)
    scope = models.TextField()
    """"organization" | "team". String not enum so new scopes ship
    without migrations.
    """
    scope_id = models.TextField(db_column="scopeId")
    """organizationId or teamId, narrowing what scope this entry applies
    to. Always equals organizationId when scope=organization.
    """
    type = models.TextField()
    """"coding_assistant" | "model_provider" | "external_tool".
    Discriminator drives the per-tile UX shape and validates
    the `config` JSON column at the app layer.
    """
    display_name = models.TextField(db_column="displayName")
    slug = models.TextField()
    """Stable slug for team-overrides-org resolution (a team-scoped
    entry with the same slug as an org-scoped one hides the org
    default for users in that team). Server-owned: generated from
    displayName + nanoid(6) on create, immutable on update.
    """
    icon_key = models.TextField(null=True, db_column="iconKey")
    """@deprecated Use `iconAsset` instead. Kept for back-compat with
    pre-refactor rows; new writes should populate iconAsset.
    """
    icon_asset = models.TextField(null=True, db_column="iconAsset")
    """Icon source - single column, prefix-discriminated:
    "preset:claude_code" / "preset:codex" / ...   (built-in)
    "data:image/svg+xml;base64,..."               (admin-uploaded)
    null                                          → fall back to
    type-default in UI
    """
    order = models.IntegerField(default=0)
    """Admin-curated sort within the user-facing portal grid."""
    enabled = models.BooleanField(default=True)
    """When false the entry is hidden from the user-facing list but
    remains in the admin list (so admins can re-enable without
    losing per-entry config). Disabled is "soft-hide", archived is
    "soft-delete" (preserves audit trail + scope binding history).
    """
    config = models.JSONField(default=dict)
    """Per-type config JSON. Validated via zod discriminated union
    at the service layer (see AiToolEntryService).
    """
    archived_at = UtcDateTimeField(null=True, db_column="archivedAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    created_by_id = models.TextField(null=True, db_column="createdById")
    updated_by_id = models.TextField(null=True, db_column="updatedById")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "AiToolEntry"
        indexes = [
            models.Index(fields=["organization", "scope", "scope_id"]),
            models.Index(fields=["organization", "enabled", "archived_at"]),
            models.Index(fields=["organization", "type"]),
        ]


class AiToolEntryTeam(models.Model):
    """@deprecated Per-team scope binding for AiToolEntry. Superseded by
    AiToolEntryDepartment; kept so the rename migration is non-destructive.
    No new rows are written here.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    entry = models.ForeignKey("AiToolEntry", on_delete=models.CASCADE, related_name="teams", db_column="entryId", db_index=False)
    team = models.ForeignKey("Team", on_delete=models.CASCADE, related_name="ai_tool_entry_teams", db_column="teamId", db_index=False)

    class Meta:
        db_table = "AiToolEntryTeam"
        constraints = [
            models.UniqueConstraint(fields=["entry", "team"], name="AiToolEntryTeam_entryId_teamId_key"),
        ]
        indexes = [
            models.Index(fields=["team"]),
        ]


class AiToolEntryDepartment(models.Model):
    """Per-department scope binding for AiToolEntry. Empty set = org-wide.
    Non-empty = entry visible only to members of those departments (an
    OrganizationUser whose `departmentId` is in the set). Mirrors
    AiToolEntryTeam's shape. `departmentId` is a plain column ref
    (relationMode="prisma"), matching how Department is referenced
    elsewhere; the service validates the id belongs to the org.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    entry = models.ForeignKey("AiToolEntry", on_delete=models.CASCADE, related_name="departments", db_column="entryId", db_index=False)
    department_id = models.TextField(db_column="departmentId")

    class Meta:
        db_table = "AiToolEntryDepartment"
        constraints = [
            models.UniqueConstraint(fields=["entry", "department_id"], name="AiToolEntryDepartment_entryId_departmentId_key"),
        ]
        indexes = [
            models.Index(fields=["department_id"]),
        ]


class IngestionTemplate(models.Model):
    """Admin/platform-curated catalog row for upstream OTLP normalization.
    Owns the OTTL transform that reshapes a specific upstream tool's
    span shape into canonical gen_ai conventions. Users connect via an
    ingestion key (an ingest-only ApiKey carrying ingestionTemplateId) to
    land traces in their personal project.

    `organizationId` NULL = platform-published default (always-visible
    across every org). NOT NULL = org-authored (visible only within
    that org). v1 ships only platform-published; org-authored authoring
    UI lands v2 (the column ships now to avoid a future migration).

    Spec: specs/ai-gateway/governance/ingestion-templates-catalog.feature
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="ingestion_templates", db_column="organizationId", null=True, db_index=False)
    slug = models.TextField()
    """Stable slug - unique within (organizationId, slug) scope. Generated
    from displayName + nanoid(6) on create for org-authored; canonical
    values for platform-published (`claude_code` / `cursor` /
    `claude_cowork` / `raw_otlp_advanced`).
    """
    source_type = models.TextField(db_column="sourceType")
    """Mirrors the slug for v1 platform-published rows. Kept as a separate
    column so a future v2 may publish two templates that share an
    upstream sourceType (e.g. claude_code-personal vs claude_code-team)
    without a migration. String not enum - extensibility.
    """
    display_name = models.TextField(db_column="displayName")
    description = models.TextField(null=True)
    icon_asset = models.TextField(null=True, db_column="iconAsset")
    """Icon source - single column, prefix-discriminated, mirrors
    AiToolEntry.iconAsset:
    `preset:<kind>`                 (built-in slot)
    `data:image/svg+xml;base64,…`   (admin-uploaded for org-authored)
    null                            → fall back to slug-default in UI
    """
    credential_schema = models.TextField(null=True, db_column="credentialSchema")
    """Discriminator for what (if any) extra credential metadata the
    install drawer captures from the user:
    null              → otlp_token-only (drawer auto-issues the
    ingestion key, no input form)
    `static_api_key`  → drawer captures the user's upstream tool API key
    `agent_id`        → drawer captures an agent identifier (e.g.
    Cursor agentId)
    String not enum so v2 shapes (oauth_redirect / s3_bucket) ship
    without a migration.
    """
    ottl_rules = models.TextField(db_column="ottlRules", default="")
    """Admin/platform-authored OTTL statements. Applied at receive time
    AFTER the receiver re-stamps the binding's authoritative principal
    + provenance keys (see `protectedTemplateAttributeKeys` 19-key
    closed list). Templates write canonical OTTL outputs (gen_ai.*,
    cost.usd, response.model) - they CANNOT write attribution or
    provenance keys; the receiver guard restores those post-OTTL.
    """
    platform_published = models.BooleanField(db_column="platformPublished", default=False)
    """True for the v1 platform-published default templates
    (organizationId IS NULL). Future org-authored rows write false.
    Drives the admin UI "READ-ONLY view + Request a custom template"
    surface - admins consume + opt-out only in v1.
    """
    enabled = models.BooleanField(default=True)
    """Admin/platform enable. False = hidden from /me catalog. Existing
    ingestion keys continue to land traces; new installs are blocked.
    """
    archived_at = UtcDateTimeField(null=True, db_column="archivedAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    created_by_id = models.TextField(null=True, db_column="createdById")
    updated_by_id = models.TextField(null=True, db_column="updatedById")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "IngestionTemplate"
        constraints = [
            models.UniqueConstraint(fields=["organization", "slug"], name="IngestionTemplate_organizationId_slug_key"),
        ]
        indexes = [
            models.Index(fields=["organization", "archived_at"]),
            models.Index(fields=["source_type"]),
        ]


class GatewayConnectUpstream(models.Model):
    organization_id = models.TextField(primary_key=True, db_column="organizationId", default=None)
    base_url = models.TextField(db_column="baseUrl")
    encrypted_token = models.TextField(db_column="encryptedToken")
    instance_id = models.TextField(db_column="instanceId")
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "GatewayConnectUpstream"


class GatewayTraceExportKey(models.Model):
    project_id = models.TextField(primary_key=True, db_column="projectId", default=None)
    api_key_id = models.TextField(db_column="apiKeyId")
    encrypted_token = models.TextField(db_column="encryptedToken")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "GatewayTraceExportKey"


class GatewayBudget(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    organization_id = models.TextField(db_column="organizationId")
    scope_type = models.TextField(choices=GatewayBudgetScopeType.choices, db_column="scopeType")
    scope_id = models.TextField(db_column="scopeId")
    provider_key = models.TextField(null=True, db_column="providerKey")
    name = models.TextField()
    description = models.TextField(null=True)
    window = models.TextField(choices=GatewayBudgetWindow.choices)
    limit_usd = models.DecimalField(max_digits=18, decimal_places=6, db_column="limitUsd")
    on_breach = models.TextField(choices=GatewayBudgetBreachAction.choices, db_column="onBreach", default=GatewayBudgetBreachAction.BLOCK)
    timezone = models.TextField(null=True)
    external_id = models.TextField(null=True, db_column="externalId")
    metadata = models.JSONField(default=dict)
    spent_usd = models.DecimalField(max_digits=18, decimal_places=6, db_column="spentUsd", default=Decimal("0"))
    current_period_started_at = UtcDateTimeField(db_column="currentPeriodStartedAt", default=now_milliseconds)
    resets_at = UtcDateTimeField(db_column="resetsAt")
    last_reset_at = UtcDateTimeField(null=True, db_column="lastResetAt")
    cycle_anchor_at = UtcDateTimeField(null=True, db_column="cycleAnchorAt")
    archived_at = UtcDateTimeField(null=True, db_column="archivedAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    created_by = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="gateway_budgets_created", db_column="createdById", db_index=False)
    managed_by_virtual_key_id = models.TextField(null=True, db_column="managedByVirtualKeyId")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "GatewayBudget"
        constraints = [
            models.UniqueConstraint(fields=["organization_id", "external_id"], name="GatewayBudget_organizationId_externalId_key"),
        ]
        indexes = [
            models.Index(fields=["organization_id"]),
            models.Index(fields=["scope_type", "scope_id", "provider_key"]),
            models.Index(fields=["managed_by_virtual_key_id"]),
            models.Index(fields=["resets_at"]),
        ]


class GatewayBudgetBucketBoundary(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    organization_id = models.TextField(db_column="organizationId")
    budget = models.ForeignKey("GatewayBudget", on_delete=models.CASCADE, related_name="bucket_boundaries", db_column="budgetId", db_index=False)
    bucket_scope_id = models.TextField(db_column="bucketScopeId")
    period_started_at = UtcDateTimeField(db_column="periodStartedAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "GatewayBudgetBucketBoundary"
        constraints = [
            models.UniqueConstraint(fields=["budget", "bucket_scope_id"], name="GatewayBudgetBucketBoundary_budgetId_bucketScopeId_key"),
        ]
        indexes = [
            models.Index(fields=["organization_id"]),
        ]


class GatewayBudgetLedger(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    budget = models.ForeignKey("GatewayBudget", on_delete=models.CASCADE, related_name="ledger_entries", db_column="budgetId", db_index=False)
    virtual_key = models.ForeignKey("VirtualKey", on_delete=models.CASCADE, related_name="ledger_entries", db_column="virtualKeyId", db_index=False)
    model_provider_id = models.TextField(null=True, db_column="modelProviderId")
    provider_key = models.TextField(null=True, db_column="providerKey")
    gateway_request_id = models.TextField(db_column="gatewayRequestId")
    amount_usd = models.DecimalField(max_digits=18, decimal_places=6, db_column="amountUsd")
    tokens_input = models.IntegerField(db_column="tokensInput", default=0)
    tokens_output = models.IntegerField(db_column="tokensOutput", default=0)
    tokens_cache_read = models.IntegerField(db_column="tokensCacheRead", default=0)
    tokens_cache_write = models.IntegerField(db_column="tokensCacheWrite", default=0)
    model = models.TextField()
    duration_ms = models.IntegerField(null=True, db_column="durationMs")
    status = models.TextField(choices=GatewayBudgetLedgerStatus.choices, default=GatewayBudgetLedgerStatus.SUCCESS)
    occurred_at = UtcDateTimeField(db_column="occurredAt", default=now_milliseconds)

    class Meta:
        db_table = "GatewayBudgetLedger"
        constraints = [
            models.UniqueConstraint(fields=["budget", "gateway_request_id"], name="GatewayBudgetLedger_budgetId_gatewayRequestId_key"),
        ]
        indexes = [
            models.Index(fields=["virtual_key"]),
            models.Index(fields=["gateway_request_id"]),
            models.Index(fields=["occurred_at"]),
        ]


class GatewayChangeEvent(models.Model):
    revision = models.BigAutoField(primary_key=True)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="gateway_change_events", db_column="organizationId", db_index=False)
    kind = models.TextField(choices=GatewayChangeEventKind.choices)
    virtual_key_id = models.TextField(null=True, db_column="virtualKeyId")
    budget_id = models.TextField(null=True, db_column="budgetId")
    model_provider_id = models.TextField(null=True, db_column="modelProviderId")
    project = models.ForeignKey("Project", on_delete=models.SET_NULL, related_name="gateway_change_events", db_column="projectId", null=True, db_index=False)
    payload = models.JSONField(null=True)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "GatewayChangeEvent"
        indexes = [
            models.Index(fields=["organization", "revision"]),
            models.Index(fields=["project", "revision"]),
            models.Index(fields=["virtual_key_id"]),
            models.Index(fields=["budget_id"]),
        ]


class GatewayCacheRule(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    organization_id = models.TextField(db_column="organizationId")
    name = models.TextField()
    description = models.TextField(null=True)
    priority = models.IntegerField(default=100)
    enabled = models.BooleanField(default=True)
    matchers = models.JSONField()
    action = models.JSONField()
    mode_enum = models.TextField(choices=GatewayCacheRuleMode.choices, db_column="modeEnum")
    archived_at = UtcDateTimeField(null=True, db_column="archivedAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    created_by_id = models.TextField(db_column="createdById")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "GatewayCacheRule"
        indexes = [
            models.Index(fields=["organization_id", "archived_at", "priority"]),
            models.Index(fields=["organization_id", "enabled", "priority"]),
        ]


class FeatureFlag(models.Model):
    key = models.TextField(primary_key=True, default=None)
    enabled = models.BooleanField()
    rules = models.JSONField(null=True)
    last_edited_by = models.TextField(null=True, db_column="lastEditedBy")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "FeatureFlag"


class FeatureFlagExperimentSetting(models.Model):
    """One experiment setting per (flag, subject). A USER row is that person's
    own opt-in; an ORGANIZATION or PROJECT row is an owner's explicit policy
    for that exact scope. No row means inherit.
    """

    pk = models.CompositePrimaryKey("flag_key", "subject_type", "subject_id")
    flag_key = models.TextField(db_column="flagKey")
    subject_type = models.TextField(choices=FeatureFlagExperimentSubjectType.choices, db_column="subjectType")
    subject_id = models.TextField(db_column="subjectId")
    enabled = models.BooleanField()
    changed_by_user_id = models.TextField(null=True, db_column="changedByUserId")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "FeatureFlagExperimentSetting"
        indexes = [
            models.Index(fields=["subject_type", "subject_id"]),
        ]


class ProcessManagerInstance(models.Model):
    """Durable state for a long-running event-sourced process."""

    id = models.TextField(primary_key=True, default=nanoid_str)
    process_name = models.TextField(db_column="processName")
    project_id = models.TextField(db_column="projectId")
    process_key = models.TextField(db_column="processKey")
    tenant_id = models.TextField(db_column="tenantId")
    user_id = models.TextField(null=True, db_column="userId")
    state = models.JSONField()
    revision = models.IntegerField()
    next_wake_at = models.DateTimeField(null=True, db_column="nextWakeAt")
    updated_at = models.DateTimeField(db_column="updatedAt")

    class Meta:
        db_table = "ProcessManagerInstance"
        constraints = [
            models.UniqueConstraint(fields=["process_name", "project_id", "process_key"], name="ProcessManagerInstance_processName_projectId_processKey_key"),
        ]
        indexes = [
            models.Index(fields=["next_wake_at"]),
        ]


class ProcessManagerInbox(models.Model):
    """Idempotency marker for a source event consumed by a process definition."""

    id = models.TextField(primary_key=True, default=nanoid_str)
    process_name = models.TextField(db_column="processName")
    project_id = models.TextField(db_column="projectId")
    process_key = models.TextField(db_column="processKey")
    tenant_id = models.TextField(db_column="tenantId")
    source_event_id = models.TextField(db_column="sourceEventId")
    """The raw source event id, kept UNINDEXED for diagnostics. A caller composes
    this and its length is that domain's business, so it must never carry a
    btree constraint — see `sourceEventKey`.
    """
    source_event_key = models.TextField(null=True, db_column="sourceEventKey")
    """The fixed-width digest of `sourceEventId` the uniqueness is enforced on,
    derived by the store (`stores/inboxKey.ts`). Postgres cannot index a btree
    row past ~2704 bytes; keying on a digest makes the constraint the same
    width no matter what any pipeline concatenates into an idempotency key.

    Nullable ONLY for the rolling deploy that introduces it — pods running the
    previous build write rows without it, and a NOT NULL would fail their
    commits outright. The store always supplies it.

    The trade, stated plainly: a NULL is not constrained by the unique index
    (btree treats NULLs as distinct), so while an old writer is serving, the
    one case that loses its guard is the same source event reaching two
    different process keys at once — the store's own read-check plus its
    per-processKey advisory lock cover everything else. A follow-up migration
    adds NOT NULL once the fleet has cycled and closes it.
    """
    consumed_at = models.DateTimeField(db_column="consumedAt")

    class Meta:
        db_table = "ProcessManagerInbox"
        constraints = [
            models.UniqueConstraint(fields=["process_name", "project_id", "source_event_key"], name="ProcessManagerInbox_processName_projectId_sourceEventKey_key"),
        ]
        indexes = [
            models.Index(fields=["process_name", "project_id", "process_key", "consumed_at"]),
            models.Index(fields=["consumed_at"]),
        ]


class ProcessManagerOutbox(models.Model):
    """Transactional intents emitted by a durable process transition."""

    id = models.TextField(primary_key=True, default=nanoid_str)
    process_name = models.TextField(db_column="processName")
    project_id = models.TextField(db_column="projectId")
    process_key = models.TextField(db_column="processKey")
    tenant_id = models.TextField(db_column="tenantId")
    user_id = models.TextField(null=True, db_column="userId")
    message_key = models.TextField(db_column="messageKey")
    intent_type = models.TextField(db_column="intentType")
    payload = models.JSONField()
    trace_carrier = models.JSONField(db_column="traceCarrier")
    source_event_id = models.TextField(null=True, db_column="sourceEventId")
    status = models.TextField(choices=ProcessManagerOutboxStatus.choices, default=ProcessManagerOutboxStatus.pending)
    attempts = models.IntegerField(default=0)
    next_attempt_at = models.DateTimeField(db_column="nextAttemptAt")
    leased_until = models.DateTimeField(null=True, db_column="leasedUntil")
    lease_token = models.TextField(null=True, unique=True, db_column="leaseToken")
    dispatched_at = models.DateTimeField(null=True, db_column="dispatchedAt")
    created_at = models.DateTimeField(db_column="createdAt")
    updated_at = models.DateTimeField(db_column="updatedAt")

    class Meta:
        db_table = "ProcessManagerOutbox"
        constraints = [
            models.UniqueConstraint(fields=["process_name", "project_id", "message_key"], name="ProcessManagerOutbox_processName_projectId_messageKey_key"),
        ]
        indexes = [
            models.Index(fields=["process_name", "project_id", "process_key", "created_at"]),
            models.Index(fields=["status", "next_attempt_at", "leased_until"]),
            models.Index(fields=["status", "process_name", "next_attempt_at"]),
        ]


class ProcessManagerOutboxAttempt(models.Model):
    """Why an outbox delivery failed, one row per FAILED attempt
    (specs/ops/dead-letter-recovery.feature). Successes write nothing, so the
    table grows with trouble, not with traffic. Rows die with their message:
    the cascade delete is the retention story, riding the existing sweep.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    outbox = models.ForeignKey("ProcessManagerOutbox", on_delete=models.CASCADE, related_name="attempt_log", db_column="outboxId", db_index=False)
    project_id = models.TextField(db_column="projectId")
    attempt = models.IntegerField()
    """1-based attempt number, as the dispatcher counts it."""
    occurred_at = models.DateTimeField(db_column="occurredAt")
    outcome = models.TextField(choices=ProcessManagerOutboxAttemptOutcome.choices)
    error_type = models.TextField(db_column="errorType")
    error_message = models.TextField(db_column="errorMessage")
    """The safe failure diagnostic the dispatcher already computes for its
    span — never a raw provider body.
    """
    retry_after_ms = models.IntegerField(null=True, db_column="retryAfterMs")

    class Meta:
        db_table = "ProcessManagerOutboxAttempt"
        indexes = [
            models.Index(fields=["outbox", "attempt"]),
        ]


class TopicClusteringRunProjection(models.Model):
    """Per-project topic clustering run status (ADR-051): the rebuildable
    operational read model behind the settings page. One row per project;
    folded from run_completed/run_failed events, never written directly.
    """

    id = models.TextField(primary_key=True, default=None)
    """KSUID (`topicrun_…`), generated by the repository on first upsert."""
    project_id = models.TextField(unique=True, db_column="projectId")
    last_requested_at = models.FloatField(null=True, db_column="lastRequestedAt")
    last_request_trigger = models.TextField(null=True, db_column="lastRequestTrigger")
    last_run_at = models.FloatField(null=True, db_column="lastRunAt")
    last_run_outcome = models.TextField(null=True, db_column="lastRunOutcome")
    """completed | skipped | failed"""
    last_run_mode = models.TextField(null=True, db_column="lastRunMode")
    last_run_skipped_reason = models.TextField(null=True, db_column="lastRunSkippedReason")
    last_run_error = models.TextField(null=True, db_column="lastRunError")
    last_run_error_code = models.TextField(null=True, db_column="lastRunErrorCode")
    last_run_error_user_actionable = models.BooleanField(db_column="lastRunErrorUserActionable", default=False)
    """True when the customer can resolve the failure (credentials, quota)."""
    last_run_traces_processed = models.IntegerField(db_column="lastRunTracesProcessed", default=0)
    last_run_topics_count = models.IntegerField(db_column="lastRunTopicsCount", default=0)
    last_run_subtopics_count = models.IntegerField(db_column="lastRunSubtopicsCount", default=0)
    last_run_pages = models.IntegerField(db_column="lastRunPages", default=0)
    in_progress_run_id = models.TextField(null=True, db_column="inProgressRunId")
    in_progress_traces = models.IntegerField(db_column="inProgressTraces", default=0)
    in_progress_pages = models.IntegerField(db_column="inProgressPages", default=0)
    in_progress_started_at = models.FloatField(null=True, db_column="inProgressStartedAt")
    """Business time the in-progress run opened; bounds how long the read
    model may report it as running (mirrors the process's stale-run guard).
    """
    created_at = models.FloatField(db_column="createdAt")
    updated_at = models.FloatField(db_column="updatedAt")
    occurred_at = models.FloatField(db_column="occurredAt")
    """Business time of the latest applied event."""
    accepted_at = models.FloatField(db_column="acceptedAt")
    """Event-log acceptance time. Together with lastEventId this is the cursor."""
    last_event_id = models.TextField(db_column="lastEventId")
    projection_version = models.TextField(db_column="projectionVersion")

    class Meta:
        db_table = "TopicClusteringRunProjection"
        constraints = [
            models.UniqueConstraint(fields=["project_id"], name="TopicClusteringRunProjection_projectId_key"),
        ]


class TopicModelProjection(models.Model):
    """Cursor row for the topic-model projection (ADR-051). The model itself is
    the `Topic` table, written through by the projection; this row only
    records how far the fold has applied the event log.
    """

    id = models.TextField(primary_key=True, default=None)
    """KSUID (`topicmodel_…`), generated by the repository on first upsert."""
    project_id = models.TextField(unique=True, db_column="projectId")
    created_at = models.FloatField(db_column="createdAt")
    updated_at = models.FloatField(db_column="updatedAt")
    occurred_at = models.FloatField(db_column="occurredAt")
    """Business time of the latest applied event."""
    accepted_at = models.FloatField(db_column="acceptedAt")
    """Event-log acceptance time. Together with lastEventId this is the cursor."""
    last_event_id = models.TextField(db_column="lastEventId")
    projection_version = models.TextField(db_column="projectionVersion")

    class Meta:
        db_table = "TopicModelProjection"
        constraints = [
            models.UniqueConstraint(fields=["project_id"], name="TopicModelProjection_projectId_key"),
        ]


class TopicClusteringRunHistoryProjection(models.Model):
    """Per-project topic clustering run history (audit read model, ADR-051).
    One row per project; `runs` is the bounded newest-first entry list.
    Rebuildable by replaying the event log.
    """

    id = models.TextField(primary_key=True, default=None)
    """KSUID (`topicrunhist_…`), generated by the repository on first upsert."""
    project_id = models.TextField(unique=True, db_column="projectId")
    runs = models.JSONField()
    created_at = models.FloatField(db_column="createdAt")
    updated_at = models.FloatField(db_column="updatedAt")
    occurred_at = models.FloatField(db_column="occurredAt")
    """Business time of the latest applied event."""
    accepted_at = models.FloatField(db_column="acceptedAt")
    """Event-log acceptance time. Together with lastEventId this is the cursor."""
    last_event_id = models.TextField(db_column="lastEventId")
    projection_version = models.TextField(db_column="projectionVersion")

    class Meta:
        db_table = "TopicClusteringRunHistoryProjection"
        constraints = [
            models.UniqueConstraint(fields=["project_id"], name="TopicClusteringRunHistoryProjection_projectId_key"),
        ]


class IngestionPullRunProjection(models.Model):
    """Rebuildable per-source pull checkpoint and operator-facing run status."""

    id = models.TextField(primary_key=True, default=None)
    source_id = models.TextField(unique=True, db_column="sourceId")
    project_id = models.TextField(db_column="projectId")
    enabled = models.BooleanField(default=False)
    cron = models.TextField(null=True)
    cursor = models.TextField(null=True)
    last_run_at = models.FloatField(null=True, db_column="lastRunAt")
    last_run_outcome = models.TextField(null=True, db_column="lastRunOutcome")
    last_run_event_count = models.IntegerField(db_column="lastRunEventCount", default=0)
    last_run_error = models.TextField(null=True, db_column="lastRunError")
    last_run_error_code = models.TextField(null=True, db_column="lastRunErrorCode")
    consecutive_errors = models.IntegerField(db_column="consecutiveErrors", default=0)
    last_success_at = models.FloatField(null=True, db_column="lastSuccessAt")
    """When a run last completed without erroring. Backfilled by replay, which
    is why the projection version moved with this column.
    """
    last_run_scheduled_for = models.FloatField(null=True, db_column="lastRunScheduledFor")
    """Identity of the run whose outcome this row reflects. Outcomes from a
    superseded run arrive with a smaller value and are ignored, so a late
    completion cannot regress Cursor or the status fields.
    """
    last_read_through_at = models.FloatField(null=True, db_column="lastReadThroughAt")
    """ADR-128 - how far the last run read, in epoch milliseconds, and whether
    it reached the end. The fold writes both; the IngestionSource row mirrors
    them for the screens. Null read-through means no run has reported one.
    Null completeness reads as unknown, never as either answer, which is what
    every row written before this existed holds.

    Nullable on purpose, so replay is not required to make this migration
    safe: a row that has never seen a run carrying these stays null and every
    reader already handles unknown.
    Spec: specs/governance/ingestion-source-health.feature
    """
    last_run_completeness = models.TextField(null=True, db_column="lastRunCompleteness")
    last_agents_listing_at = models.FloatField(null=True, db_column="lastAgentsListingAt")
    """The last time somebody asked this source to enumerate its agents, and how
    that ask ended. A listing is a different scope on the same connection as
    the pull above: a credential that cannot enumerate a directory still pulls
    cost perfectly, so no health, coverage or completeness reader may treat a
    refusal here as the source failing to pull.

    Count is set ONLY when the outcome is listed, and stays null on a refusal.
    Zero is a real answer meaning the provider returned an empty list, so a
    zero default would manufacture that answer for a source that refused.
    Reason is set only on a refusal and is a plain string, never an enum and
    never a foreign key, because the log outlives the vocabulary. Status is
    nullable even when refused, because a refusal we generated ourselves has
    no HTTP status, and it is never rendered to a customer.

    The two `...ListingAt` columns hold the instant the ASK was accepted, not
    the instant its answer was recorded, and the fold orders outcomes by them
    so a superseded listing cannot overwrite the one that replaced it. The
    name predates that job and is kept because renaming a column is a
    migration; the fold's own field docs are the statement of record.
    """
    last_agents_listing_outcome = models.TextField(null=True, db_column="lastAgentsListingOutcome")
    last_agents_listing_count = models.IntegerField(null=True, db_column="lastAgentsListingCount")
    last_agents_listing_reason = models.TextField(null=True, db_column="lastAgentsListingReason")
    last_agents_listing_status = models.IntegerField(null=True, db_column="lastAgentsListingStatus")
    last_people_listing_at = models.FloatField(null=True, db_column="lastPeopleListingAt")
    """The same for people, in its own set of columns rather than sharing the
    agents set with a kind discriminator. A source can be refused for people
    and fine for agents, and one shared set holds only the most recent listing
    of either kind, so syncing people would erase an agents refusal and the
    agents page would go back to not knowing after briefly knowing.

    LastPeopleDirectoryCount is everyone the provider's directory named.
    LastPeopleWithheldCount is how many of those this deployment does not hold
    because erasure suppression removed them: a SUBSET of the directory count,
    never an addition to it. Adding the two counts the same people twice.
    Subtracting is the valid arithmetic, and the surviving total is that
    subtraction rather than a third column that could disagree.

    The withheld count is safe to show as a CURRENT figure and never as a
    series: it moving from zero to one at a known moment says an erasure
    happened then, which on a small tenant identifies the person as surely as
    a name would. So no trend line, no history drawer, no per-run export, and
    never beside a per-person list.
    """
    last_people_listing_outcome = models.TextField(null=True, db_column="lastPeopleListingOutcome")
    last_people_directory_count = models.IntegerField(null=True, db_column="lastPeopleDirectoryCount")
    last_people_withheld_count = models.IntegerField(null=True, db_column="lastPeopleWithheldCount")
    last_people_listing_reason = models.TextField(null=True, db_column="lastPeopleListingReason")
    last_people_listing_status = models.IntegerField(null=True, db_column="lastPeopleListingStatus")
    created_at = models.FloatField(db_column="createdAt")
    updated_at = models.FloatField(db_column="updatedAt")
    occurred_at = models.FloatField(db_column="occurredAt")
    accepted_at = models.FloatField(db_column="acceptedAt")
    last_event_id = models.TextField(db_column="lastEventId")
    projection_version = models.TextField(db_column="projectionVersion")

    class Meta:
        db_table = "IngestionPullRunProjection"
        indexes = [
            models.Index(fields=["project_id"]),
        ]


class ScheduledJob(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    project_id = models.TextField(db_column="projectId")
    target_type = models.TextField(db_column="targetType")
    target_id = models.TextField(db_column="targetId")
    cron = models.TextField()
    timezone = models.TextField()
    next_run_at = UtcDateTimeField(db_column="nextRunAt")
    last_slot = UtcDateTimeField(null=True, db_column="lastSlot")
    current_slot = UtcDateTimeField(null=True, db_column="currentSlot")
    attempts = models.IntegerField(default=0)
    last_error = models.TextField(null=True, db_column="lastError")
    active = models.BooleanField(default=True)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "ScheduledJob"
        constraints = [
            models.UniqueConstraint(fields=["target_type", "target_id"], name="ScheduledJob_targetType_targetId_key"),
        ]
        indexes = [
            models.Index(fields=["active", "next_run_at"]),
            models.Index(fields=["project_id"]),
        ]


class BugReport(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    created_at = models.DateTimeField(db_column="createdAt", default=now_milliseconds)
    source = models.TextField()
    kind = models.TextField()
    title = models.TextField()
    summary = models.TextField(null=True)
    session_data = models.TextField(null=True, db_column="sessionData")
    session_truncated = models.BooleanField(db_column="sessionTruncated", default=False)
    agent = models.TextField(null=True)
    contact_email = models.TextField(null=True, db_column="contactEmail")
    cli_version = models.TextField(null=True, db_column="cliVersion")
    linked_project_id = models.TextField(null=True, db_column="linkedProjectId")
    metadata = models.JSONField(null=True)

    class Meta:
        db_table = "BugReport"
        indexes = [
            models.Index(fields=["created_at"]),
        ]


class IdempotencyReceipt(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    scope_id = models.TextField(db_column="scopeId")
    key = models.TextField()
    claim_id = models.TextField(db_column="claimId")
    request_fingerprint = models.TextField(db_column="requestFingerprint")
    response_status = models.IntegerField(null=True, db_column="responseStatus")
    response_body = models.TextField(null=True, db_column="responseBody")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    heartbeat_at = UtcDateTimeField(db_column="heartbeatAt")
    expires_at = UtcDateTimeField(db_column="expiresAt")

    class Meta:
        db_table = "IdempotencyReceipt"
        constraints = [
            models.UniqueConstraint(fields=["scope_id", "key"], name="IdempotencyReceipt_scopeId_key_key"),
        ]
        indexes = [
            models.Index(fields=["expires_at"]),
        ]


class GatewayRealtimeSession(models.Model):
    id = models.TextField(primary_key=True, default=None)
    project_id = models.TextField(db_column="projectId")
    organization_id = models.TextField(db_column="organizationId")
    virtual_key_id = models.TextField(db_column="virtualKeyId")
    model_provider_id = models.TextField(db_column="modelProviderId")
    vendor = models.TextField()
    agent_id = models.TextField(null=True, db_column="agentId")
    model = models.TextField()
    requested_model = models.TextField(null=True, db_column="requestedModel")
    vendor_conversation_id = models.TextField(null=True, db_column="vendorConversationId")
    status = models.TextField(choices=GatewayRealtimeSessionStatus.choices, default=GatewayRealtimeSessionStatus.OPEN)
    minted_at = UtcDateTimeField(db_column="mintedAt", default=now_milliseconds)
    closed_at = UtcDateTimeField(null=True, db_column="closedAt")
    close_reason = models.TextField(null=True, db_column="closeReason")
    trace_id = models.TextField(null=True, db_column="traceId")
    vendor_cost_raw = models.JSONField(null=True, db_column="vendorCostRaw")
    kind = models.TextField(null=True)
    metering = models.TextField(null=True)
    credential_expires_at = UtcDateTimeField(null=True, db_column="credentialExpiresAt")
    transcription_model = models.TextField(null=True, db_column="transcriptionModel")
    end_user_id = models.TextField(null=True, db_column="endUserId")
    last_report_at = UtcDateTimeField(null=True, db_column="lastReportAt")
    reported_cost_nano_usd = models.BigIntegerField(db_column="reportedCostNanoUsd", default=0)
    report_count = models.IntegerField(db_column="reportCount", default=0)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "GatewayRealtimeSession"
        constraints = [
            models.UniqueConstraint(fields=["vendor", "vendor_conversation_id"], name="GatewayRealtimeSession_vendor_vendorConversationId_key"),
        ]
        indexes = [
            models.Index(fields=["virtual_key_id", "status", "minted_at"]),
            models.Index(fields=["project_id", "status", "minted_at"]),
            models.Index(fields=["status", "metering", "minted_at"]),
        ]


class GatewayRealtimeSessionReport(models.Model):
    id = models.TextField(primary_key=True, default=nanoid_str)
    session_id = models.TextField(db_column="sessionId")
    report_key = models.TextField(db_column="reportKey")
    project_id = models.TextField(db_column="projectId")
    model = models.TextField()
    usage = models.JSONField()
    cost_nano_usd = models.BigIntegerField(db_column="costNanoUsd")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "GatewayRealtimeSessionReport"
        constraints = [
            models.UniqueConstraint(fields=["session_id", "report_key"], name="GatewayRealtimeSessionReport_sessionId_reportKey_key"),
        ]
        indexes = [
            models.Index(fields=["project_id"]),
        ]


class SystemMigrationTenantState(models.Model):
    pk = models.CompositePrimaryKey("migration_name", "tenant_id")
    migration_name = models.TextField(db_column="migrationName")
    tenant_id = models.TextField(db_column="tenantId")
    status = models.TextField()
    report = models.JSONField(null=True)
    occurred_at = UtcDateTimeField(db_column="occurredAt", default=now_milliseconds)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "SystemMigrationTenantState"
        indexes = [
            models.Index(fields=["migration_name", "status"]),
        ]


class SystemMigrationEnrollment(models.Model):
    pk = models.CompositePrimaryKey("organization_id", "migration_name")
    organization_id = models.TextField(db_column="organizationId")
    migration_name = models.TextField(db_column="migrationName")
    enrolled_by_user_id = models.TextField(db_column="enrolledByUserId")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "SystemMigrationEnrollment"
        indexes = [
            models.Index(fields=["migration_name"]),
        ]


class StoredObject(models.Model):
    pk = models.CompositePrimaryKey("tenant_id", "id")
    tenant_id = models.TextField(db_column="tenantId")
    id = models.TextField()
    status = models.TextField()
    purpose = models.TextField()
    owner_kind = models.TextField(db_column="ownerKind")
    owner_id = models.TextField(db_column="ownerId")
    filename = models.TextField()
    sha256 = models.TextField()
    size_bytes = models.BigIntegerField(db_column="sizeBytes")
    media_type = models.TextField(db_column="mediaType")
    media_type_verified = models.BooleanField(db_column="mediaTypeVerified", default=False)
    storage_provider = models.TextField(null=True, db_column="storageProvider")
    storage_destination_id = models.TextField(null=True, db_column="storageDestinationId")
    storage_provider_relative_id = models.TextField(null=True, db_column="storageProviderRelativeId")
    generation = models.IntegerField(default=0)
    audiences = models.JSONField()
    expires_at = UtcDateTimeField(null=True, db_column="expiresAt")
    available_at = UtcDateTimeField(null=True, db_column="availableAt")
    deleted_at = UtcDateTimeField(null=True, db_column="deletedAt")
    source = models.TextField()
    legacy_fingerprint = models.TextField(null=True, db_column="legacyFingerprint")
    created_at = UtcDateTimeField(db_column="createdAt")
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "StoredObject"
        indexes = [
            models.Index(fields=["tenant_id", "status", "id"]),
            models.Index(fields=["owner_kind", "owner_id", "status"]),
            models.Index(fields=["status", "expires_at"]),
        ]


class DiscoveredPerson(models.Model):
    """A name-text seen on a provider's cost or audit rows: "the provider told us
    someone called m.silva exists". One row per provider-scoped identity.

    This is the unit of GDPR erasure (ADR-128 §9), and deliberately so: most of
    these people have no LangWatch login at all - contractors, seat holders who
    never signed in, anyone whose email a provider put on a cost row - so there
    is no platform user to key on and a user-deletion event can never fire for
    them.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    organization_id = models.TextField(db_column="organizationId")
    provider = models.TextField()
    """"anthropic" | "openai" | "databricks" | ... - the provider that named them."""
    raw_actor_id = models.TextField(db_column="rawActorId")
    """What the provider calls them, verbatim. The join key onto the money rows'
    RawActorId. Replaced by a deterministic pseudonym on erasure, which keeps
    the row addressable for spend attribution without keeping the identifier.
    """
    display_text = models.TextField(db_column="displayText")
    """Name or email text as seen. Pseudonymized alongside rawActorId on erasure."""
    kind = models.TextField()
    """person | service_account - determined deterministically per ADR-128 §10
    (a bare UUID in Databricks `executed_by` is a service account, never an
    agent), never guessed from the shape of a name.
    """
    department = models.TextField(null=True)
    """The department the provider's directory filed this person under, verbatim.
    An OBSERVATION, not the `Department` entity: it is free text a tenant owns
    on their side, it exists for people who hold no LangWatch account, and
    nothing rolls spend up by it. Spend attribution stays on `Department`, via
    the membership the match engine proves (directoryDepartmentSync.service).
    Null means no directory has named one - never "no department", because a
    blank field in the tenant's directory and a person no directory lists are
    the same null here.
    """
    first_seen_at = UtcDateTimeField(db_column="firstSeenAt")
    last_seen_at = UtcDateTimeField(db_column="lastSeenAt")
    erased_at = UtcDateTimeField(null=True, db_column="erasedAt")
    """Set when this row's identity has been erased. The row survives so its
    spend stays attributable to *someone*; the identifier does not.
    """
    money_rows_pending_at = UtcDateTimeField(null=True, db_column="moneyRowsPendingAt")
    """Set alongside `erasedAt` and cleared once the daily cost rows have been
    removed and their rebuild requested. Non-null means the ClickHouse half
    of the erasure did not finish — a crash between the delete and the
    replay leaves days short by the erased amount, and without this marker a
    re-run reads `erasedAt`, short-circuits, and reports a clean erasure that
    did no work. The re-run resumes at that step instead.
    """
    money_rebuild_since = models.TextField(null=True, db_column="moneyRebuildSince")
    """The earliest day to rebuild from, recorded BEFORE the rows are deleted.
    Once they are gone there is nothing left to ask which days they were on,
    so a resume with no record of this could only replay everything or
    nothing. Null means nothing was rebuildable — either no day carried the
    identifier, or every one of them is older than the event log keeps.
    """
    suspended_at = UtcDateTimeField(null=True, db_column="suspendedAt")
    """ADR-128 §12 conflict rule: automatic linking has stopped for this
    identity because evidence pointed at two candidates, or contradicted an
    existing link. Stored HERE rather than on the suggestion rows, because a
    halt on auto-linking that a recompute or a restart clears is not a halt.
    """
    suspended_reason = models.TextField(null=True, db_column="suspendedReason")
    """What tripped the suspension, for the human who reviews it."""
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "DiscoveredPerson"
        constraints = [
            models.UniqueConstraint(fields=["organization_id", "provider", "raw_actor_id"], name="DiscoveredPerson_organizationId_provider_rawActorId_key"),
        ]
        indexes = [
            models.Index(fields=["organization_id", "last_seen_at"]),
        ]


class DiscoveredAgent(models.Model):
    """The same shape as DiscoveredPerson for non-human actors - Genie spaces,
    Copilot bots, OpenAI projects - plus whatever provider-native metadata the
    source published. Discovering an agent grants it nothing either.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    organization_id = models.TextField(db_column="organizationId")
    provider = models.TextField()
    raw_agent_id = models.TextField(db_column="rawAgentId")
    """Genie space / Copilot bot / OpenAI project id, verbatim."""
    display_text = models.TextField(db_column="displayText")
    metadata = models.JSONField()
    """Workspace, plan, and kin - provider-native descriptive fields only. Never
    credentials: this row is read by admin screens, not by any pull path.
    """
    first_seen_at = UtcDateTimeField(db_column="firstSeenAt")
    last_seen_at = UtcDateTimeField(db_column="lastSeenAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "DiscoveredAgent"
        constraints = [
            models.UniqueConstraint(fields=["organization_id", "provider", "raw_agent_id"], name="DiscoveredAgent_organizationId_provider_rawAgentId_key"),
        ]
        indexes = [
            models.Index(fields=["organization_id", "last_seen_at"]),
        ]


class AgentRegistryEntry(models.Model):
    """One agent in the organisation's governance registry, from any source governance reads
    (enterprise/modules/governance/specs/agent-registry.feature). `entry` is the latest
    machine-readable snapshot; `version` grows by one on every material change.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    organization_id = models.TextField(db_column="organizationId")
    source_key = models.TextField(db_column="sourceKey")
    """Where the agent came from and its id there, e.g. `langwatch:<agentId>`."""
    source = models.TextField()
    status = models.TextField()
    version = models.IntegerField()
    fingerprint = models.TextField()
    entry = models.JSONField()
    first_seen_at = UtcDateTimeField(db_column="firstSeenAt")
    last_seen_at = UtcDateTimeField(null=True, db_column="lastSeenAt")
    removed_at = UtcDateTimeField(null=True, db_column="removedAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "AgentRegistryEntry"
        constraints = [
            models.UniqueConstraint(fields=["organization_id", "source_key"], name="AgentRegistryEntry_organizationId_sourceKey_key"),
        ]
        indexes = [
            models.Index(fields=["organization_id", "updated_at"]),
        ]


class AgentRegistryEvent(models.Model):
    """The registry's append-only change feed: one row per registered, versioned, removed or
    tool-deprecated fact. `sequence` is the pull cursor; `deliveryRecordedAt` marks the webhook
    fan-out as recorded in governance's outbox.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    organization_id = models.TextField(db_column="organizationId")
    sequence = models.BigIntegerField()
    type = models.TextField()
    entry_id = models.TextField(db_column="entryId")
    from_version = models.IntegerField(null=True, db_column="fromVersion")
    to_version = models.IntegerField(null=True, db_column="toVersion")
    payload = models.JSONField()
    occurred_at = UtcDateTimeField(db_column="occurredAt")
    delivery_recorded_at = UtcDateTimeField(null=True, db_column="deliveryRecordedAt")

    class Meta:
        db_table = "AgentRegistryEvent"
        constraints = [
            models.UniqueConstraint(fields=["organization_id", "sequence"], name="AgentRegistryEvent_organizationId_sequence_key"),
        ]
        indexes = [
            models.Index(fields=["organization_id", "delivery_recorded_at"]),
            models.Index(fields=["organization_id", "entry_id"]),
        ]


class AgentRegistryBaseline(models.Model):
    """An organisation's registry baseline: its first reconcile, written as already pushed so the
    agents it already had are not announced. Every reconcile after it pushes its changes.
    """

    organization_id = models.TextField(primary_key=True, db_column="organizationId", default=None)
    baselined_at = UtcDateTimeField(db_column="baselinedAt")

    class Meta:
        db_table = "AgentRegistryBaseline"


class IdentityMatch(models.Model):
    """The dated match table: "`user_abc123` on Anthropic is platform user #42,
    valid from March 3", closed when it stops being true.

    Dating is what makes a re-issued email address survivable: the leaver's link
    closes at offboarding, the new hire gets a new link, and January's spend
    stays with January's person. A correction closes the old row and opens a new
    one; nothing is ever silently rewritten.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    organization_id = models.TextField(db_column="organizationId")
    discovered_person_id = models.TextField(db_column="discoveredPersonId")
    """The provider-side identity this link is about."""
    user_id = models.TextField(null=True, db_column="userId")
    """The platform user. Nullable, and blanked by erasure (ADR-128 §9 step 2) -
    the row and its dates remain, so the history of the link is not rewritten
    by the erasure either.
    """
    evidence_kind = models.TextField(db_column="evidenceKind")
    """directory_id | verified_email | human_confirmed. Exact verified-email
    equality is the primary evidence; a directory id corroborates but never
    stands alone (ADR-128 §12).
    """
    valid_from = UtcDateTimeField(db_column="validFrom")
    valid_to = UtcDateTimeField(null=True, db_column="validTo")
    """Null means the link is open. Offboarding and corrections close it."""
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "IdentityMatch"
        indexes = [
            models.Index(fields=["organization_id", "discovered_person_id"]),
            models.Index(fields=["organization_id", "user_id"]),
        ]


class GovernanceTenantHistory(models.Model):
    """Every ClickHouse `TenantId` this organization has ever written governance
    rows under, appended on first use and never pruned.

    The live `resolveGovProjectId` lookup cannot be this translation: it filters
    `archivedAt: null` and so returns null forever once someone archives the
    governance project, while the write path re-reads the same project by slug
    with no such filter and keeps landing rows under the old tenant. The
    organization would map to zero tenants on read and one on write - a
    permanent split-brain in which the cost screen reports "no governance data"
    and an erasure job erases nothing and reports success.

    Reads resolve against the whole history; erasure walks all of it, because
    personal data does not stop existing in a tenant that stopped being current.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    organization_id = models.TextField(db_column="organizationId")
    tenant_id = models.TextField(db_column="tenantId")
    """A hidden governance project id this organization has written rows under."""
    first_used_at = UtcDateTimeField(db_column="firstUsedAt")
    last_used_at = UtcDateTimeField(db_column="lastUsedAt")

    class Meta:
        db_table = "GovernanceTenantHistory"
        constraints = [
            models.UniqueConstraint(fields=["organization_id", "tenant_id"], name="GovernanceTenantHistory_organizationId_tenantId_key"),
        ]
        indexes = [
            models.Index(fields=["organization_id"]),
        ]


class ErasedIdentifierSuppression(models.Model):
    """Hashes of erased identifiers - the do-not-reimport list.

    Without it the erasure is undone by the pipeline that produced it: the
    pullers look 30 days back, so the next day's pull re-ingests the same email
    and re-writes what we erased. Every write path carrying an identifier
    consults this list and drops a suppressed one (ADR-128 §9 step 1).

    It stores hashes and never identifiers, so it is not itself a copy of the
    data it exists to keep out. The stored `identifierHash` and the pseudonym
    written in place of the erased value are the same function of the same
    input - HMAC-SHA256(secret, identifier) - so a write path computes that digest
    once and uses it twice: as the membership test, and, on a hit, as the
    replacement value. That is why no original-to-pseudonym mapping table
    exists; one would hold the erased identifier in plaintext forever.

    Scoped to the provider as well as the organization: the same string can be a
    different person under a different provider.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    organization_id = models.TextField(db_column="organizationId")
    provider = models.TextField()
    identifier_hash = models.TextField(db_column="identifierHash")
    erased_at = UtcDateTimeField(db_column="erasedAt")

    class Meta:
        db_table = "ErasedIdentifierSuppression"
        constraints = [
            models.UniqueConstraint(fields=["organization_id", "provider", "identifier_hash"], name="ErasedIdentifierSuppression_organizationId_provider_identifierHash_key"),
        ]
        indexes = [
            models.Index(fields=["organization_id"]),
        ]


class IdentityMatchSuggestion(models.Model):
    """Match candidates a background job computed and a human has yet to confirm
    (ADR-128 §12).

    Stored rather than computed when the review screen asks, and the reason is
    measured rather than stylistic: fuzzy matching is quadratic, there is no
    database route to it here, and at this document's own example size - 2,000
    discovered people against 500 accounts - plain edit distance over the
    1,000,000 pairs blocked the event loop for 2.9 seconds. Per page load,
    uncached, stalling every other request on the instance. A stored row also
    makes a pending count answerable without paying for the whole sweep.

    Rows are written only by the job, which prefilters before it scores and
    recomputes when its inputs change. Confirming one opens an `IdentityMatch`
    and deletes the suggestion.

    Suspension deliberately does NOT live here - it is on `DiscoveredPerson`,
    because a halt on automatic linking that the next recompute clears is not a
    halt.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    organization_id = models.TextField(db_column="organizationId")
    discovered_person_id = models.TextField(db_column="discoveredPersonId")
    user_id = models.TextField(db_column="userId")
    """The platform user this might be. Not nullable: a suggestion with no
    candidate is not a suggestion.
    """
    score = models.FloatField()
    """Similarity of the two name texts, after the prefilter, in [0, 1]. A
    display hint and an ordering key - never a threshold anything auto-links
    on, which is the whole distinction this table exists to hold.
    """
    computed_at = UtcDateTimeField(db_column="computedAt")

    class Meta:
        db_table = "IdentityMatchSuggestion"
        constraints = [
            models.UniqueConstraint(fields=["organization_id", "discovered_person_id", "user_id"], name="IdentityMatchSuggestion_organizationId_discoveredPersonId_userId_key"),
        ]
        indexes = [
            models.Index(fields=["organization_id", "score"]),
        ]


class ConnectedSeatChange(models.Model):
    """What a mid-term seat change of one connected license owes (ADR-141,
    section 7): the seats added on the reissued license, invoiced prorated to
    the end of the term.

    The row is written before the payment provider is called and completed
    after it, so a call that failed halfway is retried by the next seat
    invoicing pass without inventing a second invoice. One row per reissued license.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    license_id = models.TextField(unique=True, db_column="licenseId")
    """The reissued `IssuedLicense` row that carries the new seat count. A
    plain column ref, like `replacesId`: relationMode="prisma" carries no
    foreign keys.
    """
    account_id = models.TextField(null=True, db_column="accountId")
    """The `ConnectedBillingAccount` row the invoice is raised against. Null
    when the customer had no billing account (`not_onboarded`).
    """
    organization_id = models.TextField(null=True, db_column="organizationId")
    """The customer organization the license belongs to."""
    changed_at = UtcDateTimeField(db_column="changedAt")
    """When the seats changed, which is the day the proration counts from."""
    added_seats = models.IntegerField(db_column="addedSeats")
    """Seats on the new license above the seats on the one it replaced."""
    unit_amount_cents = models.IntegerField(db_column="unitAmountCents")
    """What one added seat costs for the rest of the term, in cents of `currency`."""
    amount_cents = models.IntegerField(db_column="amountCents")
    """`unitAmountCents` times `addedSeats`. Zero when nothing is owed."""
    currency = models.TextField(choices=Currency.choices, null=True)
    """Null when no billing account named one (`not_onboarded`)."""
    stripe_invoice_id = models.TextField(null=True, db_column="stripeInvoiceId")
    """The one-off invoice the added seats were billed on."""
    state = models.TextField(choices=ConnectedSeatChangeState.choices)
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "ConnectedSeatChange"
        indexes = [
            models.Index(fields=["account_id"]),
            models.Index(fields=["organization_id"]),
            models.Index(fields=["state"]),
        ]


class ConnectedStatement(models.Model):
    """The monthly usage statement one connected customer was sent (ADR-139,
    section 7). One row per account per month keeps the statement to one send,
    whatever runs the tick.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    account_id = models.TextField(db_column="accountId")
    """The `ConnectedBillingAccount` row. A plain column ref."""
    month = UtcDateTimeField()
    """First day of the month the statement covers, in UTC."""
    sent_at = UtcDateTimeField(db_column="sentAt")

    class Meta:
        db_table = "ConnectedStatement"
        constraints = [
            models.UniqueConstraint(fields=["account_id", "month"], name="ConnectedStatement_accountId_month_key"),
        ]
        indexes = [
            models.Index(fields=["account_id"]),
        ]


class LicenseRenewalClaim(models.Model):
    """One renewed licence per paid Stripe invoice of a self-serve licence
    subscription (specs/licensing/self-serving-license-purchase.feature). A
    redelivered invoice event finds the row and mints nothing. A claim never
    delivered is taken over once `claimedAt` is older than the claim lease.
    """

    stripe_invoice_id = models.TextField(primary_key=True, db_column="stripeInvoiceId", default=None)
    stripe_subscription_id = models.TextField(db_column="stripeSubscriptionId")
    claim_token = models.TextField(db_column="claimToken")
    """Fences completion and release to the delivery that holds the claim."""
    claimed_at = UtcDateTimeField(db_column="claimedAt")
    issued_license = models.JSONField(null=True, db_column="issuedLicense")
    """The signed licence between minting and mailing, so a retry mails it again instead of minting another; cleared once delivered."""
    delivered_at = UtcDateTimeField(null=True, db_column="deliveredAt")
    license_id = models.TextField(null=True, db_column="licenseId")
    """The `licenseId` inside the delivered licence."""

    class Meta:
        db_table = "LicenseRenewalClaim"


class InstanceIdentity(models.Model):
    """The identity this install presents to LangWatch (ADR-139, section 2).

    Exactly one row, minted on first use and never rewritten. It survives
    restarts, backups and hostname changes, and an install restored into
    another database mints another one, which is what the registry refuses as
    the wrong instance.

    It replaced an identity built from the organization name and id, which
    leaked the customer's name in cleartext, named one organization rather than
    one install, and could not be joined to a license row without splitting a
    string.
    """

    id = models.TextField(primary_key=True, default="self")
    """Fixed, so this table holds one row. A second insert loses the race on the
    primary key and re-reads the winner's id.
    """
    instance_id = models.TextField(unique=True, db_column="instanceId")
    """A random UUID. Carries nothing about the customer."""
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    last_report_at = UtcDateTimeField(null=True, db_column="lastReportAt")
    """When the usage report last reached LangWatch, and why the last attempt
    did not. Both are what the checkup page reads to say whether reporting
    works, which nothing recorded before: the sender logged success as soon
    as `fetch` resolved and never looked at the response.
    """
    last_report_error = models.TextField(null=True, db_column="lastReportError")
    optional_metrics_opt_out = models.BooleanField(db_column="optionalMetricsOptOut", default=False)
    """The optional category of the usage report, switched off from the checkup
    page. With it off the report still says which release runs and how many
    projects the install carries, and says nothing about what anyone does
    with it (ADR-139, section 10).
    """
    hostname_opt_out = models.BooleanField(db_column="hostnameOptOut", default=False)
    """Hostname, which names the customer's own network and so has a switch of
    its own rather than travelling with its category.
    """
    upgrade_diagnostics_opt_out = models.BooleanField(db_column="upgradeDiagnosticsOptOut", default=False)
    """Sending a failed upgrade step's diagnostics to LangWatch on its own,
    switched off by the operator (upgrade-diagnostics-via-connect, decision 4).
    """

    class Meta:
        db_table = "InstanceIdentity"


class SelfHostedInstance(models.Model):
    """One self-hosted install, as its daily usage report describes it (ADR-139,
    section 10). The row is created by the first report that names an instance
    id and is updated by every report after that, so the backoffice reads one
    row per install rather than folding a history on every page load.

    The report carries no credential, so it identifies an install and never a
    customer. `organizationId` and `issuedLicenseId` are resolved from the
    license bound to this instance, which the install proved by presenting its
    token to license sync. A field in the posted body never reaches them.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    instance_id = models.TextField(unique=True, db_column="instanceId")
    """The UUID the install minted for itself. The join key for everything else."""
    first_seen_at = UtcDateTimeField(db_column="firstSeenAt", default=now_milliseconds)
    last_seen_at = UtcDateTimeField(db_column="lastSeenAt")
    version = models.TextField(null=True)
    """What the install says about itself."""
    install_method = models.TextField(null=True, db_column="installMethod")
    chart_version = models.TextField(null=True, db_column="chartVersion")
    hostname = models.TextField(null=True)
    """The host it answers on, when the operator left hostname reporting on."""
    environment = models.TextField(null=True)
    installed_at = UtcDateTimeField(null=True, db_column="installedAt")
    """The day the install itself says it first came up, which predates the
    first report on an install that switched reporting on later.
    """
    report_schema_version = models.IntegerField(null=True, db_column="reportSchemaVersion")
    organization_id = models.TextField(null=True, db_column="organizationId")
    """The customer, and the license that named them. Plain column refs rather
    than relations, like `IssuedLicense.replacesId`: relationMode="prisma"
    carries no foreign keys, and the receiver resolves both from the license
    registry rather than from the report.
    """
    issued_license_id = models.TextField(null=True, db_column="issuedLicenseId")
    user_email_domains = models.JSONField(null=True, db_column="userEmailDomains")
    """Domains with counts, as `{"acme.test": 14}`. Never an address. Null on an
    install that switched the optional category off.
    """
    user_domains = ArrayField(models.TextField(), db_column="userDomains", default=list)
    """The same domains as the keys of `userEmailDomains`, flattened so an
    operator can search for one. A JSON object's keys are the domains
    themselves, and no query language addresses a key it does not already
    know the name of.
    """
    latest_report = models.JSONField(null=True, db_column="latestReport")
    """What the last report carried, so the list and the drawer read one row:
    the onboarding ladder, the usage counts and the operational block, under
    the same keys the dictionary declares.
    """
    optional_metrics_reported = models.BooleanField(db_column="optionalMetricsReported", default=True)
    """What the last report said the operator had switched off, read from which
    categories arrived rather than from a field the install could assert.
    """
    hostname_reported = models.BooleanField(db_column="hostnameReported", default=True)
    report_count = models.IntegerField(db_column="reportCount", default=0)
    """How many reports have landed, and how many fields the last one carried
    that this release has no name for. A count above zero says the senders
    are ahead of the receiver.
    """
    last_unknown_fields = models.IntegerField(db_column="lastUnknownFields", default=0)
    raised_signals = ArrayField(models.TextField(), db_column="raisedSignals", default=list)
    """The lead signals already raised for this install. A signal on this list
    is never raised again, which is what keeps a large install from
    reporting that it is large every morning.
    """
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "SelfHostedInstance"
        indexes = [
            models.Index(fields=["last_seen_at"]),
            models.Index(fields=["organization_id"]),
            models.Index(fields=["issued_license_id"]),
        ]


class SelfHostedInstanceDiagnostic(models.Model):
    """One upgrade diagnostics bundle a self-hosted install sent over Connect, on
    LangWatch Cloud. The bytes live in object storage under
    `upgrade-diagnostics/` and expire at 30 days by a lifecycle rule; this row
    points at them and feeds Cloud admin and the Slack signal. Plain column
    refs, no relation, as the usage report rows.
    """

    id = models.TextField(primary_key=True, default=None)
    instance_id = models.TextField(db_column="instanceId")
    """The `SelfHostedInstance.instanceId` the licence caller presented."""
    step_id = models.TextField(db_column="stepId")
    step_description = models.TextField(null=True, db_column="stepDescription")
    release = models.TextField(null=True)
    attempt = models.IntegerField()
    trigger = models.TextField()
    """`operator` or `automatic`."""
    byte_length = models.IntegerField(db_column="byteLength")
    object_key = models.TextField(db_column="objectKey")
    received_at = UtcDateTimeField(db_column="receivedAt", default=now_milliseconds)
    expires_at = UtcDateTimeField(db_column="expiresAt")
    """When the lifecycle rule removes the bundle; the row is listed until then."""
    slack_announced_at = UtcDateTimeField(null=True, db_column="slackAnnouncedAt")
    """Set on the report that posted this step and release's one Slack message."""

    class Meta:
        db_table = "SelfHostedInstanceDiagnostic"
        indexes = [
            models.Index(fields=["instance_id", "received_at"]),
        ]


class SelfHostedNotice(models.Model):
    """A notice LangWatch staff send every self-hosted install in Cloud's reply to its daily usage
    report (enterprise/modules/licensing/specs/self-hosted-notices.feature). Cloud only, and no
    tenant: every install receives every published notice that has not expired.
    """

    id = models.TextField(primary_key=True, default=None)
    severity = models.TextField()
    """`info`, `warning` or `security`."""
    title = models.TextField()
    body = models.TextField()
    url = models.TextField(null=True)
    """An https link to read more; null where the notice names none."""
    fixed_in = models.TextField(null=True, db_column="fixedIn")
    """The release that fixes what the notice is about; an install on it or newer hides it."""
    status = models.TextField()
    """`draft`, `published` or `withdrawn`. Only a published notice that has not expired is sent."""
    expires_at = UtcDateTimeField(null=True, db_column="expiresAt")
    published_at = UtcDateTimeField(null=True, db_column="publishedAt")
    published_by = models.TextField(null=True, db_column="publishedBy")
    """The staff member who published it, by user id."""
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "SelfHostedNotice"


class SelfHostedInstanceReport(models.Model):
    """One usage report, exactly as it arrived. Kept so a number on an install's
    row can be traced back to the report that set it, and so a question asked
    later can be answered over the history rather than over the latest snapshot.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    instance_id = models.TextField(db_column="instanceId")
    """The `SelfHostedInstance.instanceId` this report came from. A plain column
    ref for the same reason the instance row's refs are.
    """
    received_at = UtcDateTimeField(db_column="receivedAt", default=now_milliseconds)
    version = models.TextField(null=True)
    report_schema_version = models.IntegerField(null=True, db_column="reportSchemaVersion")
    unknown_fields = models.IntegerField(db_column="unknownFields", default=0)
    """Fields this release had no name for. Counted, never stored by name."""
    payload = models.JSONField()
    """The validated report, with every unknown field already dropped."""

    class Meta:
        db_table = "SelfHostedInstanceReport"
        indexes = [
            models.Index(fields=["instance_id", "received_at"]),
            models.Index(fields=["received_at"]),
        ]


class WebPushSubscription(models.Model):
    """One browser a person subscribed to Web Push, owned by notification. A
    person holds one row per device; the endpoint names the browser, so a
    browser that subscribes again replaces its own row.
    """

    id = models.TextField(primary_key=True, default=nanoid_str)
    user_id = models.TextField(db_column="userId")
    """The person the pushes are for. A plain column: the row leaves with the
    person through notification's own subscriber on the user lifecycle.
    """
    endpoint = models.TextField(unique=True)
    """The push service URL the browser handed out. Unique per browser."""
    p256dh = models.TextField()
    """The browser's P-256 public key and auth secret, base64url, as the
    browser reports them. Payloads are encrypted to these.
    """
    auth = models.TextField()
    user_agent = models.TextField(null=True, db_column="userAgent")
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)
    last_success_at = UtcDateTimeField(null=True, db_column="lastSuccessAt")
    """When the push service last accepted a push for this browser."""

    class Meta:
        db_table = "WebPushSubscription"
        indexes = [
            models.Index(fields=["user_id"]),
        ]


class WebPushVapidKey(models.Model):
    """This installation's VAPID key pair, generated on first use so Web Push works
    without configuration. Fixed id, so the table holds one row: a second
    insert loses the race on the primary key and reads the winner's pair.
    """

    id = models.TextField(primary_key=True, default="self")
    public_key = models.TextField(db_column="publicKey")
    """The uncompressed P-256 public key, base64url. The browser subscribes with it."""
    private_key_encrypted = models.TextField(db_column="privateKeyEncrypted")
    """The private key, base64url, encrypted with the process encryption key."""
    created_at = UtcDateTimeField(db_column="createdAt", default=now_milliseconds)

    class Meta:
        db_table = "WebPushVapidKey"


class SignInProviderMount(models.Model):
    id = models.TextField(primary_key=True, default="self")
    """Fixed, so this table holds one row: licensing's copy of sso's sign-in
    provider fact (ADR-179).
    """
    provider = models.TextField(null=True)
    """The configured provider's name; null when the deployment signs in by email."""
    mounted = models.BooleanField()
    last_applied_event_id = models.TextField(db_column="lastAppliedEventId")
    """The last sso event applied: the row's idempotency guard and its version."""
    updated_at = UtcAutoNowDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "SignInProviderMount"
