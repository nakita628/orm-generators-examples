from datetime import datetime
from datetime import timezone as dt_timezone
from decimal import Decimal
from typing import Any, TypeVar

from django.conf import settings
from django.contrib.postgres.fields import ArrayField
from django.contrib.postgres.indexes import BTreeIndex, BrinIndex, GinIndex
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


def project_alert_channel_environment_types_default() -> list[str]:
    return [RuntimeEnvironmentType.STAGING, RuntimeEnvironmentType.PRODUCTION]


class AuthenticationMethod(models.TextChoices):
    GITHUB = "GITHUB"
    MAGIC_LINK = "MAGIC_LINK"
    GOOGLE = "GOOGLE"
    SSO = "SSO"


class OrganizationAccessTokenType(models.TextChoices):
    USER = "USER"
    SYSTEM = "SYSTEM"


class OrgMemberRole(models.TextChoices):
    ADMIN = "ADMIN"
    MEMBER = "MEMBER"


class ApiKeyIssuer(models.TextChoices):
    USER = "USER"
    SYSTEM = "SYSTEM"


class RuntimeEnvironmentType(models.TextChoices):
    PRODUCTION = "PRODUCTION"
    STAGING = "STAGING"
    DEVELOPMENT = "DEVELOPMENT"
    PREVIEW = "PREVIEW"


class EnvironmentPauseSource(models.TextChoices):
    BILLING_LIMIT = "BILLING_LIMIT"
    ADMIN_BAN = "ADMIN_BAN"


class ProjectVersion(models.TextChoices):
    V2 = "V2"
    V3 = "V3"


class EnvironmentKeyWrappingProvider(models.TextChoices):
    PLATFORM_KEYRING = "PLATFORM_KEYRING"
    AWS_KMS = "AWS_KMS"


class SecretStoreProvider(models.TextChoices):
    DATABASE = "DATABASE"
    AWS_PARAM_STORE = "AWS_PARAM_STORE"


class OidcSigningKeyStatus(models.TextChoices):
    STAGED = "STAGED"
    ACTIVE = "ACTIVE"
    RETIRED = "RETIRED"


class OidcSigningKeyPurpose(models.TextChoices):
    """Each purpose is a separate key set with its own issuer, rotation and publication."""

    EXTERNAL_OIDC = "EXTERNAL_OIDC"
    INTERNAL_RUN = "INTERNAL_RUN"


class TaskTriggerSource(models.TextChoices):
    STANDARD = "STANDARD"
    SCHEDULED = "SCHEDULED"
    AGENT = "AGENT"
    WEBHOOK = "WEBHOOK"


class WebhookEndpointStatus(models.TextChoices):
    ACTIVE = "ACTIVE"
    INACTIVE = "INACTIVE"
    DELETING = "DELETING"


class WebhookDeliveryStatus(models.TextChoices):
    PENDING = "PENDING"
    PROCESSING = "PROCESSING"
    SUCCEEDED = "SUCCEEDED"
    FAILED = "FAILED"
    FILTERED = "FILTERED"
    UNMATCHED = "UNMATCHED"


class TaskRunStatus(models.TextChoices):
    DELAYED = "DELAYED"
    PENDING = "PENDING"
    PENDING_VERSION = "PENDING_VERSION"
    WAITING_FOR_DEPLOY = "WAITING_FOR_DEPLOY"
    DEQUEUED = "DEQUEUED"
    EXECUTING = "EXECUTING"
    WAITING_TO_RESUME = "WAITING_TO_RESUME"
    RETRYING_AFTER_FAILURE = "RETRYING_AFTER_FAILURE"
    PAUSED = "PAUSED"
    CANCELED = "CANCELED"
    INTERRUPTED = "INTERRUPTED"
    COMPLETED_SUCCESSFULLY = "COMPLETED_SUCCESSFULLY"
    COMPLETED_WITH_ERRORS = "COMPLETED_WITH_ERRORS"
    SYSTEM_FAILURE = "SYSTEM_FAILURE"
    CRASHED = "CRASHED"
    EXPIRED = "EXPIRED"
    TIMED_OUT = "TIMED_OUT"


class RunEngineVersion(models.TextChoices):
    V1 = "V1"
    V2 = "V2"


class TaskRunExecutionStatus(models.TextChoices):
    RUN_CREATED = "RUN_CREATED"
    DELAYED = "DELAYED"
    QUEUED = "QUEUED"
    QUEUED_EXECUTING = "QUEUED_EXECUTING"
    PENDING_EXECUTING = "PENDING_EXECUTING"
    EXECUTING = "EXECUTING"
    EXECUTING_WITH_WAITPOINTS = "EXECUTING_WITH_WAITPOINTS"
    SUSPENDED = "SUSPENDED"
    PENDING_CANCEL = "PENDING_CANCEL"
    FINISHED = "FINISHED"


class TaskRunCheckpointType(models.TextChoices):
    DOCKER = "DOCKER"
    KUBERNETES = "KUBERNETES"
    COMPUTE = "COMPUTE"


class WaitpointType(models.TextChoices):
    RUN = "RUN"
    DATETIME = "DATETIME"
    MANUAL = "MANUAL"
    BATCH = "BATCH"


class WaitpointStatus(models.TextChoices):
    PENDING = "PENDING"
    COMPLETED = "COMPLETED"


class WorkerInstanceGroupType(models.TextChoices):
    MANAGED = "MANAGED"
    UNMANAGED = "UNMANAGED"


class WorkloadType(models.TextChoices):
    CONTAINER = "CONTAINER"
    MICROVM = "MICROVM"


class TaskRunAttemptStatus(models.TextChoices):
    PENDING = "PENDING"
    EXECUTING = "EXECUTING"
    PAUSED = "PAUSED"
    FAILED = "FAILED"
    CANCELED = "CANCELED"
    COMPLETED = "COMPLETED"


class TaskEventLevel(models.TextChoices):
    TRACE = "TRACE"
    DEBUG = "DEBUG"
    LOG = "LOG"
    INFO = "INFO"
    WARN = "WARN"
    ERROR = "ERROR"


class TaskEventKind(models.TextChoices):
    UNSPECIFIED = "UNSPECIFIED"
    INTERNAL = "INTERNAL"
    SERVER = "SERVER"
    CLIENT = "CLIENT"
    PRODUCER = "PRODUCER"
    CONSUMER = "CONSUMER"
    UNRECOGNIZED = "UNRECOGNIZED"
    LOG = "LOG"


class TaskEventStatus(models.TextChoices):
    UNSET = "UNSET"
    OK = "OK"
    ERROR = "ERROR"
    UNRECOGNIZED = "UNRECOGNIZED"


class TaskQueueType(models.TextChoices):
    VIRTUAL = "VIRTUAL"
    NAMED = "NAMED"


class TaskQueueVersion(models.TextChoices):
    V1 = "V1"
    V2 = "V2"


class TaskQueueConcurrencyVersion(models.TextChoices):
    V1 = "V1"
    V2 = "V2"


class TaskQueueRole(models.TextChoices):
    QUEUE = "QUEUE"
    LIMIT = "LIMIT"


class BatchTaskRunStatus(models.TextChoices):
    PENDING = "PENDING"
    PROCESSING = "PROCESSING"
    COMPLETED = "COMPLETED"
    PARTIAL_FAILED = "PARTIAL_FAILED"
    ABORTED = "ABORTED"


class BatchTaskRunItemStatus(models.TextChoices):
    PENDING = "PENDING"
    FAILED = "FAILED"
    CANCELED = "CANCELED"
    COMPLETED = "COMPLETED"


class CheckpointType(models.TextChoices):
    DOCKER = "DOCKER"
    KUBERNETES = "KUBERNETES"


class CheckpointRestoreEventType(models.TextChoices):
    CHECKPOINT = "CHECKPOINT"
    RESTORE = "RESTORE"


class WorkerDeploymentType(models.TextChoices):
    MANAGED = "MANAGED"
    UNMANAGED = "UNMANAGED"
    V1 = "V1"


class WorkerDeploymentStatus(models.TextChoices):
    PENDING = "PENDING"
    INSTALLING = "INSTALLING"
    BUILDING = "BUILDING"
    DEPLOYING = "DEPLOYING"
    DEPLOYED = "DEPLOYED"
    FAILED = "FAILED"
    CANCELED = "CANCELED"
    TIMED_OUT = "TIMED_OUT"


class ScheduleType(models.TextChoices):
    DECLARATIVE = "DECLARATIVE"
    IMPERATIVE = "IMPERATIVE"


class ScheduleGeneratorType(models.TextChoices):
    CRON = "CRON"


class ProjectAlertChannelType(models.TextChoices):
    EMAIL = "EMAIL"
    SLACK = "SLACK"
    WEBHOOK = "WEBHOOK"


class ProjectAlertType(models.TextChoices):
    TASK_RUN = "TASK_RUN"
    TASK_RUN_ATTEMPT = "TASK_RUN_ATTEMPT"
    DEPLOYMENT_FAILURE = "DEPLOYMENT_FAILURE"
    DEPLOYMENT_SUCCESS = "DEPLOYMENT_SUCCESS"
    ERROR_GROUP = "ERROR_GROUP"
    DASHBOARD_AGENT_WATCH = "DASHBOARD_AGENT_WATCH"


class ProjectAlertStatus(models.TextChoices):
    PENDING = "PENDING"
    SENT = "SENT"
    FAILED = "FAILED"


class IntegrationService(models.TextChoices):
    SLACK = "SLACK"
    VERCEL = "VERCEL"


class BulkActionType(models.TextChoices):
    CANCEL = "CANCEL"
    REPLAY = "REPLAY"


class BulkActionStatus(models.TextChoices):
    PENDING = "PENDING"
    COMPLETED = "COMPLETED"
    ABORTED = "ABORTED"


class BulkActionNotificationType(models.TextChoices):
    NONE = "NONE"
    EMAIL = "EMAIL"


class BulkActionItemStatus(models.TextChoices):
    PENDING = "PENDING"
    COMPLETED = "COMPLETED"
    FAILED = "FAILED"


class GithubRepositorySelection(models.TextChoices):
    ALL = "ALL"
    SELECTED = "SELECTED"


class TrustedDeployScope(models.TextChoices):
    ANY = "ANY"
    BRANCH = "BRANCH"
    TAG = "TAG"
    ENVIRONMENT = "ENVIRONMENT"


class ImpersonationAuditLogAction(models.TextChoices):
    START = "START"
    STOP = "STOP"


class SupportAccessMode(models.TextChoices):
    ALLOW = "ALLOW"
    REQUIRES_REQUEST = "REQUIRES_REQUEST"


class SupportAccessRequestStatus(models.TextChoices):
    PENDING = "PENDING"
    APPROVED = "APPROVED"
    CANCELLED = "CANCELLED"


class CustomerQuerySource(models.TextChoices):
    DASHBOARD = "DASHBOARD"
    API = "API"


class CustomerQueryScope(models.TextChoices):
    ORGANIZATION = "ORGANIZATION"
    PROJECT = "PROJECT"
    ENVIRONMENT = "ENVIRONMENT"


class PlatformNotificationSurface(models.TextChoices):
    WEBAPP = "WEBAPP"
    CLI = "CLI"


class PlatformNotificationScope(models.TextChoices):
    USER = "USER"
    PROJECT = "PROJECT"
    ORGANIZATION = "ORGANIZATION"
    GLOBAL = "GLOBAL"


class ErrorGroupStatus(models.TextChoices):
    UNRESOLVED = "UNRESOLVED"
    RESOLVED = "RESOLVED"
    IGNORED = "IGNORED"


class DataStoreKind(models.TextChoices):
    CLICKHOUSE = "CLICKHOUSE"


class User(models.Model):
    id = models.TextField(primary_key=True, default=None)
    email = models.TextField(unique=True)
    pending_email = models.TextField(null=True, db_column="pendingEmail")
    """Address awaiting confirmation; becomes `email` only after a link sent to it is opened"""
    authentication_method = models.TextField(choices=AuthenticationMethod.choices, db_column="authenticationMethod")
    authentication_profile = models.JSONField(null=True, db_column="authenticationProfile")
    authentication_extra_params = models.JSONField(null=True, db_column="authenticationExtraParams")
    auth_identifier = models.TextField(null=True, unique=True, db_column="authIdentifier")
    display_name = models.TextField(null=True, db_column="displayName")
    name = models.TextField(null=True)
    avatar_url = models.TextField(null=True, db_column="avatarUrl")
    admin = models.BooleanField(default=False)
    dashboard_preferences = models.JSONField(null=True, db_column="dashboardPreferences")
    """Preferences for the dashboard"""
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    is_on_cloud_waitlist = models.BooleanField(db_column="isOnCloudWaitlist", default=False)
    """@deprecated"""
    feature_cloud = models.BooleanField(db_column="featureCloud", default=False)
    """@deprecated"""
    is_on_hosted_repo_waitlist = models.BooleanField(db_column="isOnHostedRepoWaitlist", default=False)
    """@deprecated"""
    marketing_emails = models.BooleanField(db_column="marketingEmails", default=True)
    confirmed_basic_details = models.BooleanField(db_column="confirmedBasicDetails", default=False)
    referral_source = models.TextField(null=True, db_column="referralSource")
    onboarding_data = models.JSONField(null=True, db_column="onboardingData")
    mfa_enabled_at = UtcDateTimeField(null=True, db_column="mfaEnabledAt")
    mfa_secret_reference = models.ForeignKey("SecretReference", on_delete=models.SET_NULL, related_name="user", db_column="mfaSecretReferenceId", null=True, db_index=False)
    mfa_last_used_code = models.TextField(null=True, db_column="mfaLastUsedCode")
    """Hash of the last used code to prevent replay attacks"""
    session_duration = models.IntegerField(db_column="sessionDuration", default=31556952)
    """Maximum session lifetime in seconds for this user. Default = 1 year (Gregorian).
    May be further restricted by Organization.maxSessionDuration on any of the user's orgs.
    """
    next_session_end = UtcDateTimeField(null=True, db_column="nextSessionEnd")
    """Absolute timestamp at which the user's current session must end. Written
    at login (and any time the effective duration is recomputed) and shortened
    when an admin lowers an org cap. Read on requireUser / getUser; null is
    treated as "no enforced deadline" (legacy sessions from before this
    shipped). Read-on-demand has zero per-request DB cost — the value
    piggybacks on the User fetch that requireUser already does.
    """
    invitation_code = models.ForeignKey("InvitationCode", on_delete=models.SET_NULL, related_name="users", db_column="invitationCodeId", null=True, db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "User"


class MfaBackupCode(models.Model):
    id = models.TextField(primary_key=True, default=None)
    code = models.TextField()
    """Hash of the actual code"""
    user = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="backup_codes", db_column="userId", db_index=False)
    used_at = UtcDateTimeField(null=True, db_column="usedAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "MfaBackupCode"
        constraints = [
            models.UniqueConstraint(fields=["user", "code"], name="MfaBackupCode_userId_code_key"),
        ]


class InvitationCode(models.Model):
    id = models.TextField(primary_key=True, default=None)
    code = models.TextField(unique=True)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)

    class Meta:
        db_table = "InvitationCode"


class AuthorizationCode(models.Model):
    """Used to generate PersonalAccessTokens, they're one-time use"""

    id = models.TextField(primary_key=True, default=None)
    code = models.TextField(unique=True)
    personal_access_token = models.ForeignKey("PersonalAccessToken", on_delete=models.CASCADE, related_name="authorization_codes", db_column="personalAccessTokenId", null=True, db_index=False)
    requester_ip = models.TextField(null=True, db_column="requesterIp")
    """Client IP of the request that minted the code, compared with the approving browser's IP. Deleted with the code once it expires."""
    code_challenge = models.TextField(null=True, db_column="codeChallenge")
    """S256 PKCE challenge from the client that minted the code. Null for codes minted by CLIs that predate it."""
    callback_port = models.IntegerField(null=True, db_column="callbackPort")
    """Loopback port the approving browser is redirected to with the approval code. Null means the user pastes it."""
    approval_code_hash = models.TextField(null=True, db_column="approvalCodeHash")
    """sha256 of the approval code, which only the approving browser ever sees."""
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "AuthorizationCode"


class PersonalAccessToken(models.Model):
    id = models.TextField(primary_key=True, default=None)
    name = models.TextField()
    """If generated by the CLI this will be "cli", otherwise user-provided"""
    encrypted_token = models.JSONField(db_column="encryptedToken")
    """This is the token encrypted using the ENCRYPTION_KEY"""
    obfuscated_token = models.TextField(db_column="obfuscatedToken")
    """This is shown in the UI, with ********"""
    hashed_token = models.TextField(unique=True, db_column="hashedToken")
    """This is used to find the token in the database"""
    user = models.ForeignKey("User", on_delete=models.RESTRICT, related_name="personal_access_tokens", db_column="userId", db_index=False)
    revoked_at = UtcDateTimeField(null=True, db_column="revokedAt")
    last_accessed_at = UtcDateTimeField(null=True, db_column="lastAccessedAt")
    expires_at = UtcDateTimeField(null=True, db_column="expiresAt")
    """After this the token no longer authenticates. Null means it never expires."""
    idle_timeout_seconds = models.IntegerField(null=True, db_column="idleTimeoutSeconds")
    """Set on CLI login tokens: they stop authenticating this many seconds after their last use
    (or creation, if never used). Null means no idle limit.
    """
    expiry_reminder_sent_at = UtcDateTimeField(null=True, db_column="expiryReminderSentAt")
    """When the "about to expire" email was sent. Set before sending, so it's sent at most once."""
    expired_notice_sent_at = UtcDateTimeField(null=True, db_column="expiredNoticeSentAt")
    """When the "has expired" email was sent. Set before sending, so it's sent at most once."""
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "PersonalAccessToken"
        indexes = [
            models.Index(fields=["user"]),
        ]


class OrganizationAccessToken(models.Model):
    id = models.TextField(primary_key=True, default=None)
    name = models.TextField()
    """User-provided name for the token"""
    type = models.TextField(choices=OrganizationAccessTokenType.choices, default=OrganizationAccessTokenType.USER)
    """Used to differentiate between user-generated and system-generated tokens"""
    hashed_token = models.TextField(unique=True, db_column="hashedToken")
    """This is used to find the token in the database"""
    organization = models.ForeignKey("Organization", on_delete=models.RESTRICT, related_name="organization_access_tokens", db_column="organizationId", db_index=False)
    expires_at = UtcDateTimeField(null=True, db_column="expiresAt")
    """Optional expiration date for the token"""
    revoked_at = UtcDateTimeField(null=True, db_column="revokedAt")
    last_accessed_at = UtcDateTimeField(null=True, db_column="lastAccessedAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "OrganizationAccessToken"
        indexes = [
            models.Index(fields=["organization", "type", "-created_at"]),
        ]


class Organization(models.Model):
    id = models.TextField(primary_key=True, default=None)
    slug = models.TextField(unique=True)
    title = models.TextField()
    maximum_execution_time_per_run_in_ms = models.IntegerField(db_column="maximumExecutionTimePerRunInMs", default=900000)
    maximum_concurrency_limit = models.IntegerField(db_column="maximumConcurrencyLimit", default=10)
    maximum_schedules_limit = models.IntegerField(db_column="maximumSchedulesLimit", default=5)
    """This is deprecated and will be removed in the future"""
    maximum_dev_queue_size = models.IntegerField(null=True, db_column="maximumDevQueueSize")
    maximum_deployed_queue_size = models.IntegerField(null=True, db_column="maximumDeployedQueueSize")
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    deleted_at = UtcDateTimeField(null=True, db_column="deletedAt")
    company_size = models.TextField(null=True, db_column="companySize")
    onboarding_data = models.JSONField(null=True, db_column="onboardingData")
    avatar = models.JSONField(null=True)
    runs_enabled = models.BooleanField(db_column="runsEnabled", default=True)
    is_activated = models.BooleanField(db_column="v3Enabled", default=False)
    v2_enabled = models.BooleanField(db_column="v2Enabled", default=False)
    """@deprecated"""
    v2_marqs_enabled = models.BooleanField(db_column="v2MarqsEnabled", default=False)
    """@deprecated"""
    has_requested_v3 = models.BooleanField(db_column="hasRequestedV3", default=False)
    """@deprecated"""
    api_rate_limiter_config = models.JSONField(null=True, db_column="apiRateLimiterConfig")
    realtime_rate_limiter_config = models.JSONField(null=True, db_column="realtimeRateLimiterConfig")
    batch_rate_limit_config = models.JSONField(null=True, db_column="batchRateLimitConfig")
    batch_queue_concurrency_config = models.JSONField(null=True, db_column="batchQueueConcurrencyConfig")
    webhook_limits_config = models.JSONField(null=True, db_column="webhookLimitsConfig")
    """Webhook limits from the org's plan or the back office: { maxWaitersPerEnvironment, maxWaitersPerEndpoint, concurrency }. Null uses the env defaults."""
    feature_flags = models.JSONField(null=True, db_column="featureFlags")
    maximum_project_count = models.IntegerField(db_column="maximumProjectCount", default=25)
    max_session_duration = models.IntegerField(null=True, db_column="maxSessionDuration")
    """Maximum session lifetime in seconds for members of this organization.
    When set, caps each member's User.sessionDuration. The most restrictive cap across
    all of a user's orgs (excluding nulls) wins.
    """
    stream_basin_name = models.TextField(null=True, db_column="streamBasinName")
    """S2 basin that holds this org's realtime streams. Null until the
    per-org basin has been provisioned (OSS / s2-lite installs leave
    it null forever; reads fall back to the global basin env var).
    Set once at provisioning time; retention is reconfigured in-place
    when the org's plan changes.
    """
    support_access_mode = models.TextField(choices=SupportAccessMode.choices, db_column="supportAccessMode", default=SupportAccessMode.ALLOW)
    support_access_mode_updated_at = UtcDateTimeField(null=True, db_column="supportAccessModeUpdatedAt")
    support_access_mode_updated_by_id = models.TextField(null=True, db_column="supportAccessModeUpdatedById")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Organization"


class OrgMember(models.Model):
    id = models.TextField(primary_key=True, default=None)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="members", db_column="organizationId", db_index=False)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="org_memberships", db_column="userId", db_index=False)
    role = models.TextField(choices=OrgMemberRole.choices, default=OrgMemberRole.MEMBER)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "OrgMember"
        constraints = [
            models.UniqueConstraint(fields=["organization", "user"], name="OrgMember_organizationId_userId_key"),
        ]
        indexes = [
            models.Index(fields=["user"]),
        ]


class OrgMemberInvite(models.Model):
    id = models.TextField(primary_key=True, default=None)
    token = models.TextField(unique=True)
    email = models.TextField()
    role = models.TextField(choices=OrgMemberRole.choices, default=OrgMemberRole.MEMBER)
    rbac_role_id = models.TextField(null=True, db_column="rbacRoleId")
    """Optional RBAC role to assign on invite acceptance. When set, the
    accept-invite flow calls the loaded RBAC plugin's setUserRole with
    this id after creating the OrgMember. Null = legacy behaviour, the
    runtime fallback derives the role from `role` above.

    Plain text (not an FK) — the RBAC plugin's RbacRole table lives on
    a separate schema (Drizzle, not Prisma) so we can't model the FK
    here. Validation happens at write time (action) and read time
    (acceptInvite).
    """
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="invites", db_column="organizationId", db_index=False)
    inviter = models.ForeignKey("User", on_delete=models.CASCADE, related_name="sent_invites", db_column="inviterId", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "OrgMemberInvite"
        constraints = [
            models.UniqueConstraint(fields=["organization", "email"], name="OrgMemberInvite_organizationId_email_key"),
        ]


class RuntimeEnvironment(models.Model):
    id = models.TextField(primary_key=True, default=None)
    slug = models.TextField()
    api_key = models.TextField(unique=True, db_column="apiKey")
    root_api_key_hidden_at = UtcDateTimeField(null=True, db_column="rootApiKeyHiddenAt")
    """Non-null prevents the root API key from being returned to users."""
    pk_api_key = models.TextField(unique=True, db_column="pkApiKey")
    """@deprecated was for v2"""
    type = models.TextField(choices=RuntimeEnvironmentType.choices, default=RuntimeEnvironmentType.DEVELOPMENT)
    is_branchable_environment = models.BooleanField(db_column="isBranchableEnvironment", default=False)
    """If true, this environment has branches and is treated differently in the dashboard/API
    NB: this flag is NOT used for Development branches, instead (type, parentEnvironmentId) = (DEVELOPMENT, NULL) is used
    """
    branch_name = models.TextField(null=True, db_column="branchName")
    parent_environment = models.ForeignKey("self", on_delete=models.CASCADE, related_name="child_environments", db_column="parentEnvironmentId", null=True, db_index=False)
    preview_auto_archive_after_days = models.IntegerField(null=True, db_column="previewAutoArchiveAfterDays")
    """Only configured on the PREVIEW root. Null disables automatic archiving."""
    preview_auto_archive_excluded_branches = ArrayField(models.TextField(), db_column="previewAutoArchiveExcludedBranches", default=list)
    preview_auto_archive_next_check_at = UtcDateTimeField(null=True, db_column="previewAutoArchiveNextCheckAt")
    """Durable per-project cleanup progress. Null next-check disables scheduling."""
    preview_auto_archive_cursor_created_at = UtcDateTimeField(null=True, db_column="previewAutoArchiveCursorCreatedAt")
    preview_auto_archive_cursor_id = models.TextField(null=True, db_column="previewAutoArchiveCursorId")
    git = models.JSONField(null=True)
    archived_at = UtcDateTimeField(null=True, db_column="archivedAt")
    """When set API calls will fail"""
    shortcode = models.TextField()
    """A memorable code for the environment"""
    maximum_concurrency_limit = models.IntegerField(db_column="maximumConcurrencyLimit", default=5)
    concurrency_limit_burst_factor = models.DecimalField(max_digits=4, decimal_places=2, db_column="concurrencyLimitBurstFactor", default=Decimal("2"))
    paused = models.BooleanField(default=False)
    pause_source = models.TextField(choices=EnvironmentPauseSource.choices, null=True, db_column="pauseSource")
    auto_enable_internal_sources = models.BooleanField(db_column="autoEnableInternalSources", default=True)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="environments", db_column="organizationId", db_index=False)
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="environments", db_column="projectId", db_index=False)
    org_member = models.ForeignKey("OrgMember", on_delete=models.SET_NULL, related_name="environments", db_column="orgMemberId", null=True, db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    built_in_environment_variable_overrides = models.JSONField(null=True, db_column="builtInEnvironmentVariableOverrides")
    """Allows us to customize the built-in environment variables for a specific environment, like TRIGGER_OTEL_SPAN_ATTRIBUTE_VALUE_LENGTH_LIMIT"""
    tunnel_id = models.TextField(null=True, db_column="tunnelId")
    current_session = models.ForeignKey("RuntimeEnvironmentSession", on_delete=models.SET_NULL, related_name="current_environments", db_column="currentSessionId", null=True, db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "RuntimeEnvironment"
        constraints = [
            models.UniqueConstraint(fields=["project", "slug", "org_member"], name="RuntimeEnvironment_projectId_slug_orgMemberId_key"),
            models.UniqueConstraint(fields=["project", "shortcode"], name="RuntimeEnvironment_projectId_shortcode_key"),
        ]
        indexes = [
            models.Index(fields=["parent_environment"]),
            models.Index(fields=["project"]),
            models.Index(fields=["organization"]),
        ]


class ApiKey(models.Model):
    id = models.TextField(primary_key=True, default=None)
    name = models.TextField()
    key_hash = models.TextField(unique=True)
    last_four = models.TextField()
    description = models.TextField(null=True)
    """Optional one-line note the creator adds to say what the key is for."""
    runtime_environment = models.ForeignKey("RuntimeEnvironment", on_delete=models.CASCADE, related_name="api_keys", db_index=False)
    created_by = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="created_api_keys", db_column="created_by_user_id", null=True, db_index=False)
    issued_by = models.TextField(choices=ApiKeyIssuer.choices, default=ApiKeyIssuer.USER)
    preset_id = models.TextField(null=True)
    scopes = ArrayField(models.TextField())
    allowed_ip_ranges = ArrayField(models.TextField(), default=list)
    """CIDR ranges this key may be used from. Empty means any IP."""
    last_used_at = UtcDateTimeField(null=True)
    last_used_ip = models.TextField(null=True)
    revoked_at = UtcDateTimeField(null=True)
    expires_at = UtcDateTimeField(null=True)
    updated_at = UtcDateTimeField(auto_now=True)
    created_at = UtcDateTimeField(default=timezone.now)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "api_keys"
        indexes = [
            models.Index(fields=["runtime_environment", "revoked_at", "-created_at"]),
            models.Index(fields=["runtime_environment", "issued_by", "revoked_at", "-created_at"]),
        ]


class RevokedApiKey(models.Model):
    """Records of previously-valid API keys that are still accepted for authentication
    during a grace window after rotation. Extend or end the grace period by updating `expiresAt`.
    """

    id = models.TextField(primary_key=True, default=None)
    api_key = models.TextField(db_column="apiKey")
    runtime_environment = models.ForeignKey("RuntimeEnvironment", on_delete=models.CASCADE, related_name="revoked_api_keys", db_column="runtimeEnvironmentId", db_index=False)
    expires_at = UtcDateTimeField(db_column="expiresAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)

    class Meta:
        db_table = "RevokedApiKey"
        indexes = [
            models.Index(fields=["api_key"]),
            models.Index(fields=["runtime_environment"]),
        ]


class Project(models.Model):
    id = models.TextField(primary_key=True, default=None)
    slug = models.TextField(unique=True)
    name = models.TextField()
    external_ref = models.TextField(unique=True, db_column="externalRef")
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="projects", db_column="organizationId", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    deleted_at = UtcDateTimeField(null=True, db_column="deletedAt")
    initialized_at = UtcDateTimeField(null=True, db_column="initializedAt")
    """Set the first time the CLI `init` command completes against this project. Drives the dev onboarding progress."""
    default_runtime = models.TextField(null=True, db_column="defaultRuntime")
    """Runtime used when a deployment config does not specify one. Null preserves the legacy Node 20 fallback."""
    version = models.TextField(choices=ProjectVersion.choices, default=ProjectVersion.V2)
    engine = models.TextField(choices=RunEngineVersion.choices, default=RunEngineVersion.V1)
    builder_project_id = models.TextField(null=True, db_column="builderProjectId")
    default_worker_group = models.ForeignKey("WorkerInstanceGroup", on_delete=models.SET_NULL, related_name="default_for_projects", db_column="defaultWorkerGroupId", null=True, db_index=False)
    allowed_worker_queues = ArrayField(models.TextField(), db_column="allowedMasterQueues", default=list)
    """The master queues they are allowed to use (impacts what they can set as default and trigger runs with)"""
    build_settings = models.JSONField(null=True, db_column="buildSettings")
    onboarding_data = models.JSONField(null=True, db_column="onboardingData")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Project"


class EnvironmentEncryptionPolicy(models.Model):
    environment = models.OneToOneField("RuntimeEnvironment", on_delete=models.CASCADE, related_name="encryption_policy", primary_key=True, db_column="environmentId")
    revision = models.IntegerField(default=0)
    wrapping_provider = models.TextField(choices=EnvironmentKeyWrappingProvider.choices, null=True, db_column="wrappingProvider")
    wrapping_key_reference = models.TextField(null=True, db_column="wrappingKeyReference")
    active_key_id = models.TextField(null=True, db_column="activeKeyId")

    class Meta:
        db_table = "EnvironmentEncryptionPolicy"


class EnvironmentEncryptionKey(models.Model):
    id = models.TextField(primary_key=True, default=None)
    environment = models.ForeignKey("RuntimeEnvironment", on_delete=models.CASCADE, related_name="encryption_keys", db_column="environmentId", db_index=False)
    wrapped_key = models.JSONField(db_column="wrappedKey")
    wrapping_provider = models.TextField(choices=EnvironmentKeyWrappingProvider.choices, db_column="wrappingProvider")
    wrapping_key_reference = models.TextField(db_column="wrappingKeyReference")
    revision = models.IntegerField(default=0)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)

    class Meta:
        db_table = "EnvironmentEncryptionKey"
        constraints = [
            models.UniqueConstraint(fields=["id", "environment"], name="EnvironmentEncryptionKey_id_environmentId_key"),
        ]
        indexes = [
            models.Index(fields=["environment"]),
        ]


class SecretReference(models.Model):
    id = models.TextField(primary_key=True, default=None)
    key = models.TextField(unique=True)
    provider = models.TextField(choices=SecretStoreProvider.choices, default=SecretStoreProvider.DATABASE)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "SecretReference"


class SecretStore(models.Model):
    key = models.TextField(primary_key=True, default=None)
    value = models.JSONField()
    version = models.TextField(default="1")
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "SecretStore"
        indexes = [
            BTreeIndex(fields=["key"]),
        ]


class OidcSigningKey(models.Model):
    id = models.TextField(primary_key=True, default=None)
    kid = models.TextField(unique=True)
    purpose = models.TextField(choices=OidcSigningKeyPurpose.choices, default=OidcSigningKeyPurpose.EXTERNAL_OIDC)
    status = models.TextField(choices=OidcSigningKeyStatus.choices)
    """At most one STAGED and one ACTIVE key per purpose: enforced by a partial unique index in the migration."""
    public_jwk = models.JSONField(db_column="publicJwk")
    ready_at = UtcDateTimeField(db_column="readyAt")
    public_until = UtcDateTimeField(null=True, db_column="publicUntil")
    generated_by_user_id = models.TextField(db_column="generatedByUserId")
    activated_at = UtcDateTimeField(null=True, db_column="activatedAt")
    activated_by_user_id = models.TextField(null=True, db_column="activatedByUserId")
    retired_at = UtcDateTimeField(null=True, db_column="retiredAt")
    retired_by_user_id = models.TextField(null=True, db_column="retiredByUserId")
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "OidcSigningKey"


class DataMigration(models.Model):
    id = models.TextField(primary_key=True, default=None)
    name = models.TextField(unique=True)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    completed_at = UtcDateTimeField(null=True, db_column="completedAt")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "DataMigration"


class BackgroundWorker(models.Model):
    id = models.TextField(primary_key=True, default=None)
    friendly_id = models.TextField(unique=True, db_column="friendlyId")
    engine = models.TextField(choices=RunEngineVersion.choices, default=RunEngineVersion.V1)
    content_hash = models.TextField(db_column="contentHash")
    sdk_version = models.TextField(db_column="sdkVersion", default="unknown")
    cli_version = models.TextField(db_column="cliVersion", default="unknown")
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="background_workers", db_column="projectId", db_index=False)
    runtime_environment = models.ForeignKey("RuntimeEnvironment", on_delete=models.CASCADE, related_name="background_workers", db_column="runtimeEnvironmentId", db_index=False)
    version = models.TextField()
    metadata = models.JSONField()
    runtime = models.TextField(null=True)
    runtime_version = models.TextField(null=True, db_column="runtimeVersion")
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    files: "models.ManyToManyField[BackgroundWorkerFile, BackgroundWorkerToBackgroundWorkerFile]" = models.ManyToManyField("BackgroundWorkerFile", through="BackgroundWorkerToBackgroundWorkerFile", related_name="background_workers")
    queues: "models.ManyToManyField[TaskQueue, BackgroundWorkerToTaskQueue]" = models.ManyToManyField("TaskQueue", through="BackgroundWorkerToTaskQueue", related_name="workers")
    worker_group = models.ForeignKey("WorkerInstanceGroup", on_delete=models.SET_NULL, related_name="background_workers", db_column="workerGroupId", null=True, db_index=False)
    supports_lazy_attempts = models.BooleanField(db_column="supportsLazyAttempts", default=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "BackgroundWorker"
        constraints = [
            models.UniqueConstraint(fields=["project", "runtime_environment", "version"], name="BackgroundWorker_projectId_runtimeEnvironmentId_version_key"),
        ]
        indexes = [
            models.Index(fields=["runtime_environment"]),
            models.Index(fields=["runtime_environment", "-created_at"]),
        ]


class BackgroundWorkerFile(models.Model):
    id = models.TextField(primary_key=True, default=None)
    friendly_id = models.TextField(unique=True, db_column="friendlyId")
    file_path = models.TextField(db_column="filePath")
    content_hash = models.TextField(db_column="contentHash")
    contents = models.BinaryField()
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="background_worker_file", db_column="projectId", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)

    class Meta:
        db_table = "BackgroundWorkerFile"
        constraints = [
            models.UniqueConstraint(fields=["project", "content_hash"], name="BackgroundWorkerFile_projectId_contentHash_key"),
        ]


class Prompt(models.Model):
    id = models.TextField(primary_key=True, default=None)
    friendly_id = models.TextField(unique=True)
    slug = models.TextField()
    description = models.TextField(null=True)
    type = models.TextField(default="text")
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="prompts", db_column="organizationId", db_index=False)
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="prompts", db_column="projectId", db_index=False)
    runtime_environment = models.ForeignKey("RuntimeEnvironment", on_delete=models.CASCADE, related_name="prompts", db_column="runtimeEnvironmentId", db_index=False)
    file_path = models.TextField(null=True, db_column="filePath")
    export_name = models.TextField(null=True, db_column="exportName")
    variable_schema = models.JSONField(null=True, db_column="variableSchema")
    default_model = models.TextField(null=True, db_column="defaultModel")
    default_config = models.JSONField(null=True, db_column="defaultConfig")
    tags = ArrayField(models.TextField(), default=list)
    archived_at = UtcDateTimeField(null=True, db_column="archivedAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "prompts"
        constraints = [
            models.UniqueConstraint(fields=["project", "runtime_environment", "slug"], name="prompts_projectId_runtimeEnvironmentId_slug_key"),
        ]
        indexes = [
            models.Index(fields=["project"]),
            models.Index(fields=["runtime_environment"]),
        ]


class PromptVersion(models.Model):
    id = models.TextField(primary_key=True, default=None)
    version = models.IntegerField()
    prompt = models.ForeignKey("Prompt", on_delete=models.CASCADE, related_name="versions", db_column="promptId", db_index=False)
    text_content = models.TextField(null=True, db_column="textContent")
    chat_content = models.JSONField(null=True, db_column="chatContent")
    model = models.TextField(null=True)
    config = models.JSONField(null=True)
    source = models.TextField()
    commit_message = models.TextField(null=True, db_column="commitMessage")
    content_hash = models.TextField(db_column="contentHash")
    labels = ArrayField(models.TextField(), default=list)
    worker = models.ForeignKey("BackgroundWorker", on_delete=models.SET_NULL, related_name="prompt_versions", db_column="workerId", null=True, db_index=False)
    created_by = models.TextField(null=True, db_column="createdBy")
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)

    class Meta:
        db_table = "prompt_versions"
        constraints = [
            models.UniqueConstraint(fields=["prompt", "version"], name="prompt_versions_promptId_version_key"),
        ]
        indexes = [
            models.Index(fields=["prompt"]),
            models.Index(fields=["content_hash"]),
        ]


class BackgroundWorkerTask(models.Model):
    id = models.TextField(primary_key=True, default=None)
    slug = models.TextField()
    description = models.TextField(null=True)
    friendly_id = models.TextField(unique=True, db_column="friendlyId")
    file_path = models.TextField(db_column="filePath")
    export_name = models.TextField(null=True, db_column="exportName")
    worker = models.ForeignKey("BackgroundWorker", on_delete=models.CASCADE, related_name="tasks", db_column="workerId", db_index=False)
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="background_worker_tasks", db_column="projectId", db_index=False)
    file = models.ForeignKey("BackgroundWorkerFile", on_delete=models.CASCADE, related_name="tasks", db_column="fileId", null=True, db_index=False)
    runtime_environment = models.ForeignKey("RuntimeEnvironment", on_delete=models.CASCADE, related_name="background_worker_tasks", db_column="runtimeEnvironmentId", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    queue_config = models.JSONField(null=True, db_column="queueConfig")
    retry_config = models.JSONField(null=True, db_column="retryConfig")
    machine_config = models.JSONField(null=True, db_column="machineConfig")
    gates = models.JSONField(null=True)
    """Gates declared on the task: other queues its runs must also hold a concurrency
    slot in while executing. Shape: [{ queue: string, concurrencyKey?: string }]
    """
    queue = models.ForeignKey("TaskQueue", on_delete=models.SET_NULL, related_name="tasks", db_column="queueId", null=True, db_index=False)
    max_duration_in_seconds = models.IntegerField(null=True, db_column="maxDurationInSeconds")
    ttl = models.TextField(null=True)
    trigger_source = models.TextField(choices=TaskTriggerSource.choices, db_column="triggerSource", default=TaskTriggerSource.STANDARD)
    config = models.JSONField(null=True)
    """Extra task configuration JSON. Shape depends on triggerSource.
    AGENT: { type: "ai-sdk-chat" }
    """
    payload_schema = models.JSONField(null=True, db_column="payloadSchema")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "BackgroundWorkerTask"
        constraints = [
            models.UniqueConstraint(fields=["worker", "slug"], name="BackgroundWorkerTask_workerId_slug_key"),
        ]
        indexes = [
            models.Index(fields=["project", "slug"]),
            models.Index(fields=["project", "slug", "created_at"]),
            models.Index(fields=["runtime_environment", "project"]),
            models.Index(fields=["runtime_environment", "slug", "trigger_source"]),
        ]


class WebhookEndpoint(models.Model):
    id = models.TextField(primary_key=True, default=None)
    friendly_id = models.TextField(unique=True, db_column="friendlyId")
    opaque_id = models.TextField(unique=True, db_column="opaqueId")
    """URL key. CSPRNG, NOT a friendlyId: crypto.randomBytes(16) base64url.
    Generated by the create path (deploy-sync / P2 API), NOT by @default.
    """
    organization_id = models.TextField(db_column="organizationId")
    project_id = models.TextField(db_column="projectId")
    runtime_environment_id = models.TextField(db_column="runtimeEnvironmentId")
    environment_type = models.TextField(choices=RuntimeEnvironmentType.choices, db_column="environmentType")
    endpoint_tenant_id = models.TextField(db_column="endpointTenantId", default="")
    endpoint_external_ref = models.TextField(db_column="endpointExternalRef", default="")
    source = models.TextField()
    declared_id = models.TextField(db_column="declaredId")
    """The declared webhooks.endpoint.define id (string ref, no relation)."""
    routing_targets = models.JSONField(db_column="routingTargets", default=list)
    """Subscribers of this endpoint: StoredWebhookRoutingTarget[] ({ id, type: "task" | "session", ..., filterAst })."""
    verifier_artifact = models.JSONField(db_column="verifierArtifact")
    metadata = models.JSONField(default=dict)
    secret_provisioning = models.TextField(db_column="secretProvisioning", default="either")
    setup_prompt = models.TextField(null=True, db_column="setupPrompt")
    """The source's instructions for an AI agent registering this endpoint with the provider."""
    signing_secret_key = models.TextField(null=True, db_column="signingSecretKey")
    """SecretReference.key string. Plain String, NO @relation -> no FK to SecretReference."""
    status = models.TextField(choices=WebhookEndpointStatus.choices, default=WebhookEndpointStatus.ACTIVE)
    manually_deactivated_at = UtcDateTimeField(null=True, db_column="manuallyDeactivatedAt")
    """When an operator disabled the endpoint via the dashboard/API. Null means the declarative sync
    owns the status: a redeploy that re-declares a previously-removed (auto-deactivated) webhook
    reactivates it. Non-null means the operator disabled it, so the sync leaves the status alone.
    """
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "WebhookEndpoint"
        constraints = [
            models.UniqueConstraint(fields=["runtime_environment_id", "declared_id", "endpoint_tenant_id", "endpoint_external_ref"], name="WebhookEndpoint_runtimeEnvironmentId_declaredId_endpointTenantId_endpointExternalRef_key"),
        ]
        indexes = [
            models.Index(fields=["runtime_environment_id", "source"]),
            models.Index(fields=["project_id", "-created_at"]),
        ]


class WebhookDelivery(models.Model):
    pk = models.CompositePrimaryKey("id", "created_at", "retention_days")
    id = models.TextField()
    friendly_id = models.TextField(db_column="friendlyId")
    webhook_endpoint_id = models.TextField(db_column="webhookEndpointId")
    organization_id = models.TextField(db_column="organizationId")
    project_id = models.TextField(db_column="projectId")
    runtime_environment_id = models.TextField(db_column="runtimeEnvironmentId")
    environment_type = models.TextField(choices=RuntimeEnvironmentType.choices, db_column="environmentType")
    external_delivery_id = models.TextField(db_column="externalDeliveryId")
    idempotency_key = models.TextField(db_column="idempotencyKey")
    run_id = models.TextField(null=True, db_column="runId")
    status = models.TextField(choices=WebhookDeliveryStatus.choices, default=WebhookDeliveryStatus.PENDING)
    is_test = models.BooleanField(db_column="isTest", default=False)
    """Set from the x-trigger-test ingress header; marks console/test-send deliveries so the list can filter them."""
    parsed_event = models.JSONField(null=True, db_column="parsedEvent")
    headers = models.JSONField(null=True)
    raw_body_hash = models.TextField(null=True, db_column="rawBodyHash")
    error_message = models.TextField(null=True, db_column="errorMessage")
    filter_reason = models.TextField(null=True, db_column="filterReason")
    targets = models.JSONField(default=list)
    """Per-target results: WebhookDeliveryTargetResult[] ({ id, type, status, reason?, runId?, error? })."""
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    processed_at = UtcDateTimeField(null=True, db_column="processedAt")
    retention_days = models.IntegerField(db_column="retentionDays", default=30)
    """Retention class in days: the LIST partition key, stamped at ingest from the org's plan and encoded in the id."""
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "WebhookDelivery"
        indexes = [
            models.Index(fields=["webhook_endpoint_id", "-created_at"]),
        ]


class Alert(models.Model):
    id = models.TextField(primary_key=True, default=None)
    organization_id = models.TextField(db_column="organizationId")
    project_id = models.TextField(db_column="projectId")
    environment_id = models.TextField(db_column="environmentId")
    capability_key = models.TextField(db_column="capabilityKey")
    params = models.JSONField()
    condition = models.JSONField()
    query = models.TextField(null=True)
    expires_at = models.DateTimeField(null=True, db_column="expiresAt")
    lifecycle = models.TextField(default="active")
    attention = models.TextField(default="clear")
    observed_condition = models.TextField(db_column="observedCondition", default="unknown")
    observed_condition_changed_at = models.DateTimeField(null=True, db_column="observedConditionChangedAt")
    pending_met_since = models.DateTimeField(null=True, db_column="pendingMetSince")
    evaluator_state = models.JSONField(null=True, db_column="evaluatorState")
    evaluation_lease_id = models.TextField(null=True, db_column="evaluationLeaseId")
    evaluation_lease_until = models.DateTimeField(null=True, db_column="evaluationLeaseUntil")
    last_evaluated_at = models.DateTimeField(null=True, db_column="lastEvaluatedAt")
    next_evaluation_at = models.DateTimeField(null=True, db_column="nextEvaluationAt")
    notifiers = models.JSONField(default=list)
    name = models.TextField(null=True)
    created_by = models.TextField(db_column="createdBy")
    created_at = models.DateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = models.DateTimeField(db_column="updatedAt", default=timezone.now)
    resolved_at = models.DateTimeField(null=True, db_column="resolvedAt")
    resolved_by = models.TextField(null=True, db_column="resolvedBy")
    stopped_at = models.DateTimeField(null=True, db_column="stoppedAt")
    stopped_by = models.TextField(null=True, db_column="stoppedBy")
    last_occurrence_at = models.DateTimeField(null=True, db_column="lastOccurrenceAt")
    occurrences_since_raise = models.IntegerField(db_column="occurrencesSinceRaise", default=0)
    occurrences_total = models.BigIntegerField(db_column="occurrencesTotal", default=0)
    group_id = models.TextField(null=True, db_column="groupId")

    class Meta:
        db_table = "Alert"
        indexes = [
            models.Index(fields=["environment_id", "lifecycle"]),
            models.Index(fields=["lifecycle", "evaluation_lease_until"]),
            models.Index(fields=["project_id"]),
            models.Index(fields=["organization_id"]),
            models.Index(fields=["group_id"]),
        ]


class AlertEvent(models.Model):
    """RANGE-partitioned by month on `at` (see migration)."""

    pk = models.CompositePrimaryKey("id", "at")
    id = models.TextField()
    alert_id = models.TextField(db_column="alertId")
    type = models.TextField()
    at = models.DateTimeField(default=timezone.now)
    facts = models.JSONField(default=dict)
    actor_id = models.TextField(null=True, db_column="actorId")
    source_event_id = models.TextField(null=True, db_column="sourceEventId")

    class Meta:
        db_table = "AlertEvent"
        indexes = [
            models.Index(fields=["alert_id", "-at"]),
        ]


class AlertOccurrenceSeen(models.Model):
    pk = models.CompositePrimaryKey("alert_id", "source_event_id")
    alert_id = models.TextField(db_column="alertId")
    source_event_id = models.TextField(db_column="sourceEventId")
    row_time = models.DateTimeField(db_column="rowTime")
    seen_at = models.DateTimeField(db_column="seenAt", default=timezone.now)
    expires_at = models.DateTimeField(db_column="expiresAt")

    class Meta:
        db_table = "AlertOccurrenceSeen"
        indexes = [
            models.Index(fields=["expires_at"]),
        ]


class AlertDelivery(models.Model):
    """RANGE-partitioned by month on `eventAt` (see migration)."""

    pk = models.CompositePrimaryKey("id", "event_at")
    id = models.TextField()
    alert_event_id = models.TextField(db_column="alertEventId")
    event_at = models.DateTimeField(db_column="eventAt")
    notifier_id = models.TextField(db_column="notifierId")
    notifier = models.JSONField()
    status = models.TextField(default="pending")
    claim_id = models.TextField(null=True, db_column="claimId")
    claimed_at = models.DateTimeField(null=True, db_column="claimedAt")
    delivered_at = models.DateTimeField(null=True, db_column="deliveredAt")
    failure_reason = models.TextField(null=True, db_column="failureReason")
    provider_metadata = models.JSONField(null=True, db_column="providerMetadata")
    attempts = models.IntegerField(default=0)
    first_attempt_at = models.DateTimeField(null=True, db_column="firstAttemptAt")
    created_at = models.DateTimeField(db_column="createdAt", default=timezone.now)

    class Meta:
        db_table = "AlertDelivery"
        constraints = [
            models.UniqueConstraint(fields=["alert_event_id", "notifier_id", "event_at"], name="AlertDelivery_alertEventId_notifierId_eventAt_key"),
        ]
        indexes = [
            models.Index(fields=["status", "created_at"]),
            models.Index(fields=["status", "claimed_at"]),
        ]


class PlaygroundConversation(models.Model):
    id = models.TextField(primary_key=True, default=None)
    chat_id = models.TextField(db_column="chatId")
    """The chat session ID used by the transport"""
    title = models.TextField(default="New conversation")
    """User-editable conversation title (auto-generated from first message)"""
    agent_slug = models.TextField(db_column="agentSlug")
    """Which agent this conversation is with"""
    run = models.ForeignKey("TaskRun", on_delete=models.SET_NULL, related_name="playground_conversations", db_column="runId", null=True, db_index=False)
    """The current active run backing this conversation (null if no run yet)"""
    client_data = models.JSONField(null=True, db_column="clientData")
    """The client data JSON used for this conversation"""
    messages = models.JSONField(null=True)
    """Accumulated UIMessages from completed turns (for resume without stream replay)"""
    last_event_id = models.TextField(null=True, db_column="lastEventId")
    """Last SSE event ID — resume from this position to avoid replaying old turns"""
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="playground_conversations", db_column="projectId", db_index=False)
    runtime_environment = models.ForeignKey("RuntimeEnvironment", on_delete=models.CASCADE, related_name="playground_conversations", db_column="runtimeEnvironmentId", db_index=False)
    user_id = models.TextField(db_column="userId")
    """The user who started this conversation"""
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "PlaygroundConversation"
        constraints = [
            models.UniqueConstraint(fields=["chat_id", "runtime_environment"], name="PlaygroundConversation_chatId_runtimeEnvironmentId_key"),
        ]
        indexes = [
            models.Index(fields=["runtime_environment", "agent_slug", "-updated_at"]),
            models.Index(fields=["user_id", "runtime_environment"]),
        ]


class Session(models.Model):
    """Durable, typed, bidirectional I/O primitive. Owns two S2 streams (.out / .in).
    The row is essentially static — no status, no counters, no pointers. No
    foreign keys: project/runtimeEnvironment/organization ids are plain
    scalar columns (matches TaskRun pattern). List-style queries are served
    from ClickHouse, not Postgres, so only point-lookup indexes live here.
    """

    id = models.TextField(primary_key=True, default=None)
    friendly_id = models.TextField(unique=True, db_column="friendlyId")
    external_id = models.TextField(null=True, db_column="externalId")
    """User-supplied identifier scoped to the environment. Used for
    idempotent upsert and for resolving sessions via the public API.
    """
    type = models.TextField()
    """Plain string — intentionally not an enum."""
    project_id = models.TextField(db_column="projectId")
    """Denormalized scoping columns — no FK relations."""
    runtime_environment_id = models.TextField(db_column="runtimeEnvironmentId")
    environment_type = models.TextField(choices=RuntimeEnvironmentType.choices, db_column="environmentType")
    organization_id = models.TextField(db_column="organizationId")
    task_identifier = models.TextField(db_column="taskIdentifier")
    """Task this session triggers runs against. Required — Sessions are
    task-bound: creating a session also triggers its first run, and
    every subsequent re-trigger uses this same identifier.
    """
    trigger_config = models.JSONField(db_column="triggerConfig")
    """Trigger config used for every run this session schedules. Shape
    (validated at the route layer, opaque to the DB):
    { basePayload: object, machine?: string, queue?: string,
    tags?: string[], maxAttempts?: number,
    idleTimeoutInSeconds?: number }
    `basePayload` carries the customer's client-data; runtime fields
    (chatId, messages, trigger) are merged at trigger time.
    """
    is_test = models.BooleanField(db_column="isTest", default=False)
    """Whether this session was created from the Test/playground UI rather
    than by a real chat.agent() trigger. Mirrors TaskRun.isTest.
    """
    tags = ArrayField(models.TextField(), default=list)
    metadata = models.JSONField(null=True)
    current_run_id = models.TextField(null=True, db_column="currentRunId")
    """Live run pointer — non-FK so run deletion never cascades. Can lag
    reality; the `.in/append` handler re-checks the snapshot status
    before reusing it.
    """
    current_run_version = models.IntegerField(db_column="currentRunVersion", default=0)
    """Monotonic counter used for optimistic locking on `currentRunId`
    swaps. Bumped atomically alongside any update that changes
    `currentRunId`.
    """
    closed_at = UtcDateTimeField(null=True, db_column="closedAt")
    """Terminal markers — written once, never flipped back."""
    closed_reason = models.TextField(null=True, db_column="closedReason")
    expires_at = UtcDateTimeField(null=True, db_column="expiresAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    stream_basin_name = models.TextField(null=True, db_column="streamBasinName")
    """S2 basin where this session's stream pair lives. Stamped at create
    time from `Organization.streamBasinName` so reads can resolve the
    basin without joining org. Null when the org has no per-org basin
    (OSS, or pre-backfill); reads fall back to the global basin.
    """
    chat_snapshot_storage_path = models.TextField(null=True, db_column="chatSnapshotStoragePath")
    """Storage URI (with protocol prefix) for this session's chat.agent
    snapshot blob. Set on first snapshot write. Null = pre-column session,
    fall back to computed default path.
    """
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Session"
        constraints = [
            models.UniqueConstraint(fields=["runtime_environment_id", "external_id"], name="Session_runtimeEnvironmentId_externalId_key"),
        ]
        indexes = [
            models.Index(fields=["expires_at"]),
            models.Index(fields=["current_run_id"]),
        ]


class SessionRun(models.Model):
    """Historical record of every run a Session has owned. Append-only —
    rows are inserted on each `ensureRunForSession` claim, never updated.
    Lets us reconstruct the run timeline of a chat for debugging /
    dashboard surfaces. The relation cascades on Session delete (tied to
    the session lifecycle) but `runId` is a plain string column with no
    FK to TaskRun so run pruning is independent.
    """

    id = models.TextField(primary_key=True, default=None)
    session = models.ForeignKey("Session", on_delete=models.CASCADE, related_name="runs", db_column="sessionId", db_index=False)
    run_id = models.TextField(unique=True, db_column="runId")
    """TaskRun.id (no FK — runs may be archived independently of session history)"""
    reason = models.TextField()
    """One of: "initial" | "continuation" | "upgrade" | "manual".
    Plain string for forward-compat with future trigger reasons.
    """
    triggered_at = UtcDateTimeField(db_column="triggeredAt", default=timezone.now)

    class Meta:
        db_table = "SessionRun"
        indexes = [
            models.Index(fields=["session"]),
            models.Index(fields=["session", "triggered_at", "id"]),
        ]


class TaskRun(models.Model):
    id = models.TextField(primary_key=True, default=None)
    number = models.IntegerField(default=0)
    friendly_id = models.TextField(unique=True, db_column="friendlyId")
    engine = models.TextField(choices=RunEngineVersion.choices, default=RunEngineVersion.V1)
    status = models.TextField(choices=TaskRunStatus.choices, default=TaskRunStatus.PENDING)
    status_reason = models.TextField(null=True, db_column="statusReason")
    idempotency_key = models.TextField(null=True, db_column="idempotencyKey")
    idempotency_key_expires_at = UtcDateTimeField(null=True, db_column="idempotencyKeyExpiresAt")
    idempotency_key_options = models.JSONField(null=True, db_column="idempotencyKeyOptions")
    """Stores the user-provided key and scope: { key: string, scope: "run" | "attempt" | "global" }"""
    debounce = models.JSONField(null=True)
    """Debounce options: { key: string, delay: string, createdAt: Date }"""
    task_identifier = models.TextField(db_column="taskIdentifier")
    is_test = models.BooleanField(db_column="isTest", default=False)
    payload = models.TextField()
    payload_type = models.TextField(db_column="payloadType", default="application/json")
    context = models.JSONField(null=True)
    trace_context = models.JSONField(null=True, db_column="traceContext")
    trace_id = models.TextField(db_column="traceId")
    span_id = models.TextField(db_column="spanId")
    runtime_environment = models.ForeignKey("RuntimeEnvironment", on_delete=models.CASCADE, related_name="task_runs", db_column="runtimeEnvironmentId", db_index=False)
    environment_type = models.TextField(choices=RuntimeEnvironmentType.choices, null=True, db_column="environmentType")
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="task_runs", db_column="projectId", db_index=False)
    organization_id = models.TextField(null=True, db_column="organizationId")
    queue = models.TextField()
    locked_queue_id = models.TextField(null=True, db_column="lockedQueueId")
    worker_queue = models.TextField(db_column="masterQueue", default="main")
    """The main queue that this run is part of"""
    region = models.TextField(null=True)
    """User-facing geo region, stamped at trigger; workerQueue is where it actually ran."""
    secondary_master_queue = models.TextField(null=True, db_column="secondaryMasterQueue")
    """@deprecated"""
    attempt_number = models.IntegerField(null=True, db_column="attemptNumber")
    """From engine v2+ this will be defined after a run has been dequeued (starting at 1)"""
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    run_tags = ArrayField(models.TextField(), db_column="runTags")
    """Denormized column that holds the raw tags"""
    task_version = models.TextField(null=True, db_column="taskVersion")
    """Denormalized version of the background worker task"""
    sdk_version = models.TextField(null=True, db_column="sdkVersion")
    cli_version = models.TextField(null=True, db_column="cliVersion")
    started_at = UtcDateTimeField(null=True, db_column="startedAt")
    """startedAt marks the point at which a run is dequeued from MarQS"""
    executed_at = UtcDateTimeField(null=True, db_column="executedAt")
    """executedAt is set when the first attempt is about to execute"""
    completed_at = UtcDateTimeField(null=True, db_column="completedAt")
    machine_preset = models.TextField(null=True, db_column="machinePreset")
    usage_duration_ms = models.IntegerField(db_column="usageDurationMs", default=0)
    cost_in_cents = models.FloatField(db_column="costInCents", default=0)
    base_cost_in_cents = models.FloatField(db_column="baseCostInCents", default=0)
    locked_at = UtcDateTimeField(null=True, db_column="lockedAt")
    locked_by = models.ForeignKey("BackgroundWorkerTask", on_delete=models.SET_NULL, related_name="runs", db_column="lockedById", null=True, db_index=False)
    locked_to_version = models.ForeignKey("BackgroundWorker", on_delete=models.SET_NULL, related_name="locked_runs", db_column="lockedToVersionId", null=True, db_index=False)
    priority_ms = models.IntegerField(db_column="priorityMs", default=0)
    """The "priority" of the run. This is just a negative offset in ms for the queue timestamp
    E.g. a value of 60_000 would put the run into the queue 60s ago.
    """
    concurrency_key = models.TextField(null=True, db_column="concurrencyKey")
    gates = models.JSONField(null=True)
    """Gates for this run: other queues it must also hold a concurrency slot in while
    executing. Shape: [{ queue: string, concurrencyKey?: string }]
    """
    delay_until = UtcDateTimeField(null=True, db_column="delayUntil")
    queued_at = UtcDateTimeField(null=True, db_column="queuedAt")
    ttl = models.TextField(null=True)
    expired_at = UtcDateTimeField(null=True, db_column="expiredAt")
    max_attempts = models.IntegerField(null=True, db_column="maxAttempts")
    locked_retry_config = models.JSONField(null=True, db_column="lockedRetryConfig")
    one_time_use_token = models.TextField(null=True, unique=True, db_column="oneTimeUseToken")
    """optional token that can be used to authenticate the task run"""
    connected_waitpoints: "models.ManyToManyField[Waitpoint, WaitpointRunConnections]" = models.ManyToManyField("Waitpoint", through="WaitpointRunConnections", related_name="connected_runs")
    """All waitpoints that blocked this run at some point, used for display purposes"""
    task_event_store = models.TextField(db_column="taskEventStore", default="taskEvent")
    """Where the logs are stored"""
    queue_timestamp = UtcDateTimeField(null=True, db_column="queueTimestamp")
    schedule_instance_id = models.TextField(null=True, db_column="scheduleInstanceId")
    schedule_id = models.TextField(null=True, db_column="scheduleId")
    bulk_action_group_ids = ArrayField(models.TextField(), db_column="bulkActionGroupIds", default=list)
    logs_deleted_at = UtcDateTimeField(null=True, db_column="logsDeletedAt")
    replayed_from_task_run_friendly_id = models.TextField(null=True, db_column="replayedFromTaskRunFriendlyId")
    root_task_run = models.ForeignKey("self", on_delete=models.SET_NULL, related_name="descendant_runs", db_column="rootTaskRunId", null=True, db_index=False)
    """This represents the original task that that was triggered outside of a Trigger.dev task"""
    parent_task_run = models.ForeignKey("self", on_delete=models.SET_NULL, related_name="child_runs", db_column="parentTaskRunId", null=True, db_index=False)
    """The immediate parent run of this task run"""
    parent_task_run_attempt = models.ForeignKey("TaskRunAttempt", on_delete=models.SET_NULL, related_name="child_runs", db_column="parentTaskRunAttemptId", null=True, db_index=False)
    """The immediate parent attempt of this task run"""
    batch = models.ForeignKey("BatchTaskRun", on_delete=models.SET_NULL, related_name="runs", db_column="batchId", null=True, db_index=False)
    """The batch run that this task run is a part of"""
    resume_parent_on_completion = models.BooleanField(db_column="resumeParentOnCompletion", default=False)
    """whether or not the task run was created because of a triggerAndWait for batchTriggerAndWait"""
    depth = models.IntegerField(default=0)
    """The depth of this task run in the task run hierarchy"""
    parent_span_id = models.TextField(null=True, db_column="parentSpanId")
    """The span ID of the "trigger" span in the parent task run"""
    run_chain_state = models.JSONField(null=True, db_column="runChainState")
    """Holds the state of the run chain for deadlock detection"""
    seed_metadata = models.TextField(null=True, db_column="seedMetadata")
    """seed run metadata"""
    seed_metadata_type = models.TextField(db_column="seedMetadataType", default="application/json")
    metadata = models.TextField(null=True)
    """Run metadata"""
    metadata_type = models.TextField(db_column="metadataType", default="application/json")
    metadata_version = models.IntegerField(db_column="metadataVersion", default=1)
    annotations = models.JSONField(null=True)
    """Structured annotations: triggerSource, triggerAction, rootTriggerSource, rootScheduleId"""
    is_warm_start = models.BooleanField(null=True, db_column="isWarmStart")
    """Whether the latest attempt was a warm start. Null until first attempt starts."""
    output = models.TextField(null=True)
    """Run output"""
    output_type = models.TextField(db_column="outputType", default="application/json")
    error = models.JSONField(null=True)
    """Run error"""
    plan_type = models.TextField(null=True, db_column="planType")
    """Organization's billing plan type (cached for fallback when billing API fails)"""
    max_duration_in_seconds = models.IntegerField(null=True, db_column="maxDurationInSeconds")
    realtime_streams_version = models.TextField(db_column="realtimeStreamsVersion", default="v1")
    """The version of the realtime streams implementation used by the run"""
    realtime_streams = ArrayField(models.TextField(), db_column="realtimeStreams", default=list)
    """Store the stream keys that are being used by the run"""
    stream_basin_name = models.TextField(null=True, db_column="streamBasinName")
    """S2 basin where this run's realtime streams live. Stamped at create
    time from `Organization.streamBasinName` so reads can resolve the
    basin without joining org. Null when the org has no per-org basin
    (OSS, or pre-backfill); reads fall back to the global basin.
    """
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "TaskRun"
        constraints = [
            models.UniqueConstraint(fields=["one_time_use_token"], name="TaskRun_oneTimeUseToken_key"),
            models.UniqueConstraint(fields=["runtime_environment", "task_identifier", "idempotency_key"], name="TaskRun_runtimeEnvironmentId_taskIdentifier_idempotencyKey_key"),
        ]
        indexes = [
            models.Index(fields=["parent_task_run"]),
            models.Index(fields=["span_id"]),
            models.Index(fields=["parent_span_id"]),
            GinIndex(fields=["run_tags"]),
            models.Index(fields=["runtime_environment", "batch"]),
            models.Index(fields=["runtime_environment", "-created_at"]),
            BrinIndex(fields=["created_at"]),
        ]


class TaskRunTemplate(models.Model):
    id = models.TextField(primary_key=True, default=None)
    task_slug = models.TextField(db_column="taskSlug")
    trigger_source = models.TextField(choices=TaskTriggerSource.choices, db_column="triggerSource")
    label = models.TextField()
    payload = models.TextField(null=True)
    payload_type = models.TextField(db_column="payloadType", default="application/json")
    metadata = models.TextField(null=True)
    metadata_type = models.TextField(db_column="metadataType", default="application/json")
    queue = models.TextField(null=True)
    concurrency_key = models.TextField(null=True, db_column="concurrencyKey")
    ttl_seconds = models.IntegerField(null=True, db_column="ttlSeconds")
    max_attempts = models.IntegerField(null=True, db_column="maxAttempts")
    max_duration_seconds = models.IntegerField(null=True, db_column="maxDurationSeconds")
    tags = ArrayField(models.TextField())
    machine_preset = models.TextField(null=True, db_column="machinePreset")
    project_id = models.TextField(db_column="projectId")
    organization_id = models.TextField(db_column="organizationId")
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "TaskRunTemplate"
        indexes = [
            models.Index(fields=["project_id", "task_slug", "trigger_source", "-created_at"]),
        ]


class TaskRunExecutionSnapshot(models.Model):
    """Used by the RunEngine during TaskRun execution
    It has the required information to transactionally progress a run through states,
    and prevent side effects like heartbeats failing a run that has progressed.
    It is optimised for performance and is designed to be cleared at some point,
    so there are no cascading relationships to other models.
    """

    id = models.TextField(primary_key=True, default=None)
    engine = models.TextField(choices=RunEngineVersion.choices, default=RunEngineVersion.V2)
    """This should always be 2+ (V1 didn't use the run engine or snapshots)"""
    execution_status = models.TextField(choices=TaskRunExecutionStatus.choices, db_column="executionStatus")
    """The execution status"""
    description = models.TextField()
    """For debugging"""
    is_valid = models.BooleanField(db_column="isValid", default=True)
    """We store invalid snapshots as a record of the run state when we tried to move"""
    error = models.TextField(null=True)
    previous_snapshot_id = models.TextField(null=True, db_column="previousSnapshotId")
    """The previous snapshot ID"""
    run = models.ForeignKey("TaskRun", on_delete=models.RESTRICT, related_name="execution_snapshots", db_column="runId", db_index=False)
    """Run"""
    run_status = models.TextField(choices=TaskRunStatus.choices, db_column="runStatus")
    batch_id = models.TextField(null=True, db_column="batchId")
    attempt_number = models.IntegerField(null=True, db_column="attemptNumber")
    """This is the current run attempt number. Users can define how many attempts they want for a run."""
    environment_id = models.TextField(db_column="environmentId")
    """Environment"""
    environment_type = models.TextField(choices=RuntimeEnvironmentType.choices, db_column="environmentType")
    project_id = models.TextField(db_column="projectId")
    organization_id = models.TextField(db_column="organizationId")
    completed_waitpoints: "models.ManyToManyField[Waitpoint, CompletedWaitpoints]" = models.ManyToManyField("Waitpoint", through="CompletedWaitpoints", related_name="completed_execution_snapshots")
    """Waitpoints that have been completed for this execution"""
    completed_waitpoint_order = ArrayField(models.TextField(), db_column="completedWaitpointOrder")
    """An array of waitpoint IDs in the correct order, used for batches"""
    checkpoint = models.ForeignKey("TaskRunCheckpoint", on_delete=models.SET_NULL, related_name="execution_snapshot", db_column="checkpointId", null=True, db_index=False)
    """Checkpoint"""
    worker_id = models.TextField(null=True, db_column="workerId")
    """Worker"""
    runner_id = models.TextField(null=True, db_column="runnerId")
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    last_heartbeat_at = UtcDateTimeField(null=True, db_column="lastHeartbeatAt")
    metadata = models.JSONField(null=True)
    """Metadata used by various systems in the run engine"""
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "TaskRunExecutionSnapshot"
        indexes = [
            models.Index(fields=["run", "is_valid", "-created_at"]),
        ]


class TaskRunCheckpoint(models.Model):
    id = models.TextField(primary_key=True, default=None)
    friendly_id = models.TextField(unique=True, db_column="friendlyId")
    type = models.TextField(choices=TaskRunCheckpointType.choices)
    location = models.TextField()
    image_ref = models.TextField(null=True, db_column="imageRef")
    reason = models.TextField(null=True)
    metadata = models.TextField(null=True)
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="task_run_checkpoints", db_column="projectId", db_index=False)
    runtime_environment = models.ForeignKey("RuntimeEnvironment", on_delete=models.CASCADE, related_name="task_run_checkpoints", db_column="runtimeEnvironmentId", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "TaskRunCheckpoint"


class Waitpoint(models.Model):
    """A Waitpoint blocks a run from continuing until it's completed
    If there's a waitpoint blocking a run, it shouldn't be in the queue
    """

    id = models.TextField(primary_key=True, default=None)
    friendly_id = models.TextField(unique=True, db_column="friendlyId")
    type = models.TextField(choices=WaitpointType.choices)
    status = models.TextField(choices=WaitpointStatus.choices, default=WaitpointStatus.PENDING)
    completed_at = UtcDateTimeField(null=True, db_column="completedAt")
    idempotency_key = models.TextField(db_column="idempotencyKey")
    """If it's an Event type waitpoint, this is the event. It can also be provided for the DATETIME type"""
    user_provided_idempotency_key = models.BooleanField(db_column="userProvidedIdempotencyKey")
    """If this is true then we can show it in the dashboard/return it from the SDK"""
    idempotency_key_expires_at = UtcDateTimeField(null=True, db_column="idempotencyKeyExpiresAt")
    """If there's a user provided idempotency key, this is the time it expires at"""
    inactive_idempotency_key = models.TextField(null=True, db_column="inactiveIdempotencyKey")
    """If an idempotencyKey is no longer active, we store it here and generate a new one for the idempotencyKey field.
    Clearing an idempotencyKey is useful for debounce or cancelling child runs.
    This is a workaround because Prisma doesn't support partial indexes.
    """
    completed_by_task_run = models.OneToOneField("TaskRun", on_delete=models.SET_NULL, related_name="associated_waitpoint", db_column="completedByTaskRunId", null=True)
    """If it's a RUN type waitpoint, this is the associated run"""
    completed_after = UtcDateTimeField(null=True, db_column="completedAfter")
    """If it's a DATETIME type waitpoint, this is the date.
    If it's a MANUAL waitpoint, this can be set as the `timeout`.
    """
    completed_by_batch = models.ForeignKey("BatchTaskRun", on_delete=models.SET_NULL, related_name="waitpoints", db_column="completedByBatchId", null=True, db_index=False)
    """If it's a BATCH type waitpoint, this is the associated batch"""
    output = models.TextField(null=True)
    """When completed, an output can be stored here"""
    output_type = models.TextField(db_column="outputType", default="application/json")
    output_is_error = models.BooleanField(db_column="outputIsError", default=False)
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="waitpoints", db_column="projectId", db_index=False)
    environment = models.ForeignKey("RuntimeEnvironment", on_delete=models.CASCADE, related_name="waitpoints", db_column="environmentId", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    tags = ArrayField(models.TextField())
    """Denormized column that holds the raw tags
    Denormalized column that holds the raw tags
    """
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Waitpoint"
        constraints = [
            models.UniqueConstraint(fields=["environment", "idempotency_key"], name="Waitpoint_environmentId_idempotencyKey_key"),
        ]
        indexes = [
            models.Index(fields=["completed_by_batch"]),
            models.Index(fields=["environment", "type", "-created_at"]),
            models.Index(fields=["environment", "type", "status"]),
            models.Index(fields=["environment", "type", "-id"]),
        ]


class TaskRunWaitpoint(models.Model):
    id = models.TextField(primary_key=True, default=None)
    task_run = models.ForeignKey("TaskRun", on_delete=models.RESTRICT, related_name="blocked_by_waitpoints", db_column="taskRunId", db_index=False)
    waitpoint = models.ForeignKey("Waitpoint", on_delete=models.RESTRICT, related_name="blocking_task_runs", db_column="waitpointId", db_index=False)
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="task_run_waitpoints", db_column="projectId", db_index=False)
    span_id_to_complete = models.TextField(null=True, db_column="spanIdToComplete")
    """This span id is completed when the waitpoint is completed. This is used with cached runs (idempotent)"""
    batch = models.ForeignKey("BatchTaskRun", on_delete=models.SET_NULL, related_name="runs_blocked", db_column="batchId", null=True, db_index=False)
    batch_index = models.IntegerField(null=True, db_column="batchIndex")
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "TaskRunWaitpoint"
        constraints = [
            models.UniqueConstraint(fields=["task_run", "waitpoint", "batch_index"], name="TaskRunWaitpoint_taskRunId_waitpointId_batchIndex_key"),
        ]
        indexes = [
            models.Index(fields=["task_run"]),
            models.Index(fields=["waitpoint"]),
        ]


class WaitpointTag(models.Model):
    id = models.TextField(primary_key=True, default=None)
    name = models.TextField()
    environment = models.ForeignKey("RuntimeEnvironment", on_delete=models.CASCADE, related_name="waitpoint_tags", db_column="environmentId", db_index=False)
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="waitpoint_tags", db_column="projectId", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)

    class Meta:
        db_table = "WaitpointTag"
        constraints = [
            models.UniqueConstraint(fields=["environment", "name"], name="WaitpointTag_environmentId_name_key"),
        ]


class FeatureFlag(models.Model):
    id = models.TextField(primary_key=True, default=None)
    key = models.TextField(unique=True)
    value = models.JSONField(null=True)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "FeatureFlag"


class WorkerInstance(models.Model):
    id = models.TextField(primary_key=True, default=None)
    name = models.TextField()
    """For example "worker-1\""""
    resource_identifier = models.TextField(db_column="resourceIdentifier")
    """If managed, it will default to the name, e.g. "worker-1"
    If unmanged, it will be prefixed with the deployment ID e.g. "deploy-123-worker-1\"
    """
    metadata = models.JSONField(null=True)
    worker_group = models.ForeignKey("WorkerInstanceGroup", on_delete=models.RESTRICT, related_name="workers", db_column="workerGroupId", db_index=False)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="worker_instances", db_column="organizationId", null=True, db_index=False)
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="workers", db_column="projectId", null=True, db_index=False)
    environment = models.ForeignKey("RuntimeEnvironment", on_delete=models.CASCADE, related_name="worker_instances", db_column="environmentId", null=True, db_index=False)
    deployment = models.ForeignKey("WorkerDeployment", on_delete=models.SET_NULL, related_name="worker_instance", db_column="deploymentId", null=True, db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    last_dequeue_at = UtcDateTimeField(null=True, db_column="lastDequeueAt")
    last_heartbeat_at = UtcDateTimeField(null=True, db_column="lastHeartbeatAt")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "WorkerInstance"
        constraints = [
            models.UniqueConstraint(fields=["worker_group", "resource_identifier"], name="WorkerInstance_workerGroupId_resourceIdentifier_key"),
        ]


class WorkerInstanceGroup(models.Model):
    id = models.TextField(primary_key=True, default=None)
    type = models.TextField(choices=WorkerInstanceGroupType.choices)
    name = models.TextField()
    """For example "us-east-1\""""
    master_queue = models.TextField(unique=True, db_column="masterQueue")
    """If managed, it will default to the name, e.g. "us-east-1"
    If unmanged, it will be prefixed with the project ID e.g. "project_1-us-east-1\"
    """
    description = models.TextField(null=True)
    """"N. Virginia, USA", "Frankfurt, Germany", etc. Used for display purposes"""
    hidden = models.BooleanField(default=False)
    token = models.OneToOneField("WorkerGroupToken", on_delete=models.CASCADE, related_name="worker_group", db_column="tokenId")
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="worker_groups", db_column="organizationId", null=True, db_index=False)
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="worker_groups", db_column="projectId", null=True, db_index=False)
    cloud_provider = models.TextField(null=True, db_column="cloudProvider")
    location = models.TextField(null=True)
    """"usa", "europe", etc. Used like a pseudo enum for things like flags"""
    static_ips = models.TextField(null=True, db_column="staticIPs")
    workload_type = models.TextField(choices=WorkloadType.choices, db_column="workloadType", default=WorkloadType.CONTAINER)
    region = models.TextField(null=True)
    """Geo region; container + MICROVM groups for one geo share it. Set-once once it has runs."""
    enable_fast_path = models.BooleanField(db_column="enableFastPath", default=False)
    """When true, runs enqueued to this worker queue may skip the intermediate queue
    and be pushed directly to the worker queue when concurrency is available.
    """
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "WorkerInstanceGroup"


class WorkerGroupToken(models.Model):
    id = models.TextField(primary_key=True, default=None)
    token_hash = models.TextField(unique=True, db_column="tokenHash")
    previous_token_hash = models.TextField(null=True, unique=True, db_column="previousTokenHash")
    previous_token_valid_until = UtcDateTimeField(null=True, db_column="previousTokenValidUntil")
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "WorkerGroupToken"


class TaskRunTag(models.Model):
    id = models.TextField(primary_key=True, default=None)
    name = models.TextField()
    friendly_id = models.TextField(unique=True, db_column="friendlyId")
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="run_tags", db_column="projectId", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)

    class Meta:
        db_table = "TaskRunTag"
        constraints = [
            models.UniqueConstraint(fields=["project", "name"], name="TaskRunTag_projectId_name_key"),
        ]
        indexes = [
            models.Index(fields=["name", "id"]),
        ]


class TaskRunDependency(models.Model):
    """This is used for triggerAndWait and batchTriggerAndWait. The taskRun is the child task, it points at a parent attempt or a batch"""

    id = models.TextField(primary_key=True, default=None)
    task_run = models.OneToOneField("TaskRun", on_delete=models.CASCADE, related_name="dependency", db_column="taskRunId")
    """The child run"""
    checkpoint_event = models.OneToOneField("CheckpointRestoreEvent", on_delete=models.CASCADE, related_name="task_run_dependency", db_column="checkpointEventId", null=True)
    dependent_attempt = models.ForeignKey("TaskRunAttempt", on_delete=models.SET_NULL, related_name="dependencies", db_column="dependentAttemptId", null=True, db_index=False)
    """An attempt that is dependent on this task run."""
    dependent_batch_run = models.ForeignKey("BatchTaskRun", on_delete=models.SET_NULL, related_name="run_dependencies", db_column="dependentBatchRunId", null=True, db_index=False)
    """A batch run that is dependent on this task run"""
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    resumed_at = UtcDateTimeField(null=True, db_column="resumedAt")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "TaskRunDependency"
        indexes = [
            models.Index(fields=["dependent_attempt"]),
            models.Index(fields=["dependent_batch_run"]),
        ]


class TaskRunCounter(models.Model):
    """deprecated, we hadn't included the project id in the unique constraint"""

    task_identifier = models.TextField(primary_key=True, db_column="taskIdentifier", default=None)
    last_number = models.IntegerField(db_column="lastNumber", default=0)

    class Meta:
        db_table = "TaskRunCounter"


class TaskRunNumberCounter(models.Model):
    """Used for the TaskRun number (e.g. #1,421)"""

    id = models.TextField(primary_key=True, default=None)
    task_identifier = models.TextField(db_column="taskIdentifier")
    environment = models.ForeignKey("RuntimeEnvironment", on_delete=models.CASCADE, related_name="task_run_number_counter", db_column="environmentId", db_index=False)
    last_number = models.IntegerField(db_column="lastNumber", default=0)

    class Meta:
        db_table = "TaskRunNumberCounter"
        constraints = [
            models.UniqueConstraint(fields=["task_identifier", "environment"], name="TaskRunNumberCounter_taskIdentifier_environmentId_key"),
        ]


class TaskRunAttempt(models.Model):
    """This is not used from engine v2+, attempts use the TaskRunExecutionSnapshot and TaskRun"""

    id = models.TextField(primary_key=True, default=None)
    number = models.IntegerField(default=0)
    friendly_id = models.TextField(unique=True, db_column="friendlyId")
    task_run = models.ForeignKey("TaskRun", on_delete=models.CASCADE, related_name="attempts", db_column="taskRunId", db_index=False)
    background_worker = models.ForeignKey("BackgroundWorker", on_delete=models.CASCADE, related_name="attempts", db_column="backgroundWorkerId", db_index=False)
    background_worker_task = models.ForeignKey("BackgroundWorkerTask", on_delete=models.CASCADE, related_name="attempts", db_column="backgroundWorkerTaskId", db_index=False)
    runtime_environment = models.ForeignKey("RuntimeEnvironment", on_delete=models.CASCADE, related_name="task_run_attempts", db_column="runtimeEnvironmentId", db_index=False)
    queue = models.ForeignKey("TaskQueue", on_delete=models.CASCADE, related_name="attempts", db_column="queueId", db_index=False)
    status = models.TextField(choices=TaskRunAttemptStatus.choices, default=TaskRunAttemptStatus.PENDING)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    started_at = UtcDateTimeField(null=True, db_column="startedAt")
    completed_at = UtcDateTimeField(null=True, db_column="completedAt")
    usage_duration_ms = models.IntegerField(db_column="usageDurationMs", default=0)
    error = models.JSONField(null=True)
    output = models.TextField(null=True)
    output_type = models.TextField(db_column="outputType", default="application/json")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "TaskRunAttempt"
        constraints = [
            models.UniqueConstraint(fields=["task_run", "number"], name="TaskRunAttempt_taskRunId_number_key"),
        ]
        indexes = [
            models.Index(fields=["task_run"]),
        ]


class TaskEvent(models.Model):
    """This is the unified otel span/log model that will eventually be replaced by clickhouse"""

    id = models.TextField(primary_key=True, default=None)
    message = models.TextField()
    """This matches the span name for a trace event, or the log body for a log event"""
    trace_id = models.TextField(db_column="traceId")
    span_id = models.TextField(db_column="spanId")
    parent_id = models.TextField(null=True, db_column="parentId")
    tracestate = models.TextField(null=True)
    is_error = models.BooleanField(db_column="isError", default=False)
    is_partial = models.BooleanField(db_column="isPartial", default=False)
    is_cancelled = models.BooleanField(db_column="isCancelled", default=False)
    is_debug = models.BooleanField(db_column="isDebug", default=False)
    """deprecated: don't use this, moving this to properties, this now uses TaskEventKind.LOG"""
    service_name = models.TextField(db_column="serviceName")
    service_namespace = models.TextField(db_column="serviceNamespace")
    level = models.TextField(choices=TaskEventLevel.choices, default=TaskEventLevel.TRACE)
    kind = models.TextField(choices=TaskEventKind.choices, default=TaskEventKind.INTERNAL)
    status = models.TextField(choices=TaskEventStatus.choices, default=TaskEventStatus.UNSET)
    links = models.JSONField(null=True)
    events = models.JSONField(null=True)
    start_time = models.BigIntegerField(db_column="startTime")
    """This is the time the event started in nanoseconds since the epoch"""
    duration = models.BigIntegerField(default=0)
    """This is the duration of the event in nanoseconds"""
    attempt_id = models.TextField(null=True, db_column="attemptId")
    attempt_number = models.IntegerField(null=True, db_column="attemptNumber")
    environment_id = models.TextField(db_column="environmentId")
    environment_type = models.TextField(choices=RuntimeEnvironmentType.choices, db_column="environmentType")
    organization_id = models.TextField(db_column="organizationId")
    project_id = models.TextField(db_column="projectId")
    project_ref = models.TextField(db_column="projectRef")
    run_id = models.TextField(db_column="runId")
    run_is_test = models.BooleanField(db_column="runIsTest", default=False)
    idempotency_key = models.TextField(null=True, db_column="idempotencyKey")
    task_slug = models.TextField(db_column="taskSlug")
    task_path = models.TextField(null=True, db_column="taskPath")
    task_export_name = models.TextField(null=True, db_column="taskExportName")
    worker_id = models.TextField(null=True, db_column="workerId")
    worker_version = models.TextField(null=True, db_column="workerVersion")
    queue_id = models.TextField(null=True, db_column="queueId")
    queue_name = models.TextField(null=True, db_column="queueName")
    batch_id = models.TextField(null=True, db_column="batchId")
    properties = models.JSONField()
    """This represents all the span attributes available, like http.status_code, and special attributes like $style.icon, $output, $metadata.payload.userId, as it's used for searching and filtering"""
    metadata = models.JSONField(null=True)
    """This represents all span attributes in the $metadata namespace, like $metadata.payload"""
    style = models.JSONField(null=True)
    """This represents all span attributes in the $style namespace, like $style"""
    output = models.JSONField(null=True)
    """This represents all span attributes in the $output namespace, like $output"""
    output_type = models.TextField(null=True, db_column="outputType")
    """This represents the mimetype of the output, such as application/json or application/super+json"""
    payload = models.JSONField(null=True)
    payload_type = models.TextField(null=True, db_column="payloadType")
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    usage_duration_ms = models.IntegerField(db_column="usageDurationMs", default=0)
    usage_cost_in_cents = models.FloatField(db_column="usageCostInCents", default=0)
    machine_preset = models.TextField(null=True, db_column="machinePreset")
    machine_preset_cpu = models.FloatField(null=True, db_column="machinePresetCpu")
    machine_preset_memory = models.FloatField(null=True, db_column="machinePresetMemory")
    machine_preset_cents_per_ms = models.FloatField(null=True, db_column="machinePresetCentsPerMs")

    class Meta:
        db_table = "TaskEvent"
        indexes = [
            models.Index(fields=["trace_id"]),
            models.Index(fields=["span_id"]),
            models.Index(fields=["run_id"]),
        ]


class TaskQueue(models.Model):
    id = models.TextField(primary_key=True, default=None)
    friendly_id = models.TextField(unique=True, db_column="friendlyId")
    name = models.TextField()
    type = models.TextField(choices=TaskQueueType.choices, default=TaskQueueType.VIRTUAL)
    version = models.TextField(choices=TaskQueueVersion.choices, default=TaskQueueVersion.V1)
    orderable_name = models.TextField(null=True, db_column="orderableName")
    concurrency_version = models.TextField(choices=TaskQueueConcurrencyVersion.choices, db_column="concurrencyVersion", default=TaskQueueConcurrencyVersion.V1)
    """How the queue's concurrency was declared: V1 = legacy concurrencyLimit semantics
    (per key when keyed, whole queue otherwise), V2 = the explicit perKey/total shape.
    """
    role = models.TextField(choices=TaskQueueRole.choices, default=TaskQueueRole.QUEUE)
    """QUEUE rows are real queues; LIMIT rows back named concurrency limits and are
    excluded from every queue-facing surface.
    """
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="task_queues", db_column="projectId", db_index=False)
    runtime_environment = models.ForeignKey("RuntimeEnvironment", on_delete=models.CASCADE, related_name="task_queues", db_column="runtimeEnvironmentId", db_index=False)
    concurrency_limit = models.IntegerField(null=True, db_column="concurrencyLimit")
    """Represents the current concurrency limit for the queue"""
    concurrency_limit_overridden_at = UtcDateTimeField(null=True, db_column="concurrencyLimitOverriddenAt")
    """When the concurrency limit was overridden"""
    concurrency_limit_overridden_by = models.TextField(null=True, db_column="concurrencyLimitOverriddenBy")
    """Who overrode the concurrency limit (will be null if overridden via the API)"""
    concurrency_limit_base = models.IntegerField(null=True, db_column="concurrencyLimitBase")
    """If concurrencyLimit is overridden, this is the overridden value"""
    concurrency_limit_override_percent = models.DecimalField(max_digits=5, decimal_places=2, null=True, db_column="concurrencyLimitOverridePercent")
    """If the override was expressed as a percentage of the environment limit, this stores that
    percentage (the source of truth). The absolute concurrencyLimit is materialized from it.
    Decimal(5,2) allows fractional percentages like 12.50% (0.01–100.00).
    """
    total_concurrency_limit = models.IntegerField(null=True, db_column="totalConcurrencyLimit")
    """Caps total concurrent runs across ALL concurrencyKey values of this queue
    (concurrencyLimit applies per key value). Null = no total cap.
    """
    total_concurrency_limit_overridden_at = UtcDateTimeField(null=True, db_column="totalConcurrencyLimitOverriddenAt")
    """When the total concurrency limit was overridden"""
    total_concurrency_limit_overridden_by = models.TextField(null=True, db_column="totalConcurrencyLimitOverriddenBy")
    """Who overrode the total concurrency limit (null when overridden via the API)"""
    total_concurrency_limit_base = models.IntegerField(null=True, db_column="totalConcurrencyLimitBase")
    """If totalConcurrencyLimit is overridden, the declared value it reverts to on reset"""
    rate_limit = models.JSONField(null=True, db_column="rateLimit")
    paused = models.BooleanField(default=False)
    archived_at = UtcDateTimeField(null=True, db_column="archivedAt")
    """Hidden from the dashboard queue list. Display-only: never read on trigger/dequeue paths."""
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "TaskQueue"
        constraints = [
            models.UniqueConstraint(fields=["runtime_environment", "name"], name="TaskQueue_runtimeEnvironmentId_name_key"),
        ]


class BatchTaskRun(models.Model):
    id = models.TextField(primary_key=True, default=None)
    friendly_id = models.TextField(unique=True, db_column="friendlyId")
    idempotency_key = models.TextField(null=True, db_column="idempotencyKey")
    idempotency_key_expires_at = UtcDateTimeField(null=True, db_column="idempotencyKeyExpiresAt")
    status = models.TextField(choices=BatchTaskRunStatus.choices, default=BatchTaskRunStatus.PENDING)
    runtime_environment_id = models.TextField(db_column="runtimeEnvironmentId")
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    run_ids = ArrayField(models.TextField(), db_column="runIds", default=list)
    """Friendly IDs"""
    run_count = models.IntegerField(db_column="runCount", default=0)
    payload = models.TextField(null=True)
    payload_type = models.TextField(db_column="payloadType", default="application/json")
    options = models.JSONField(null=True)
    batch_version = models.TextField(db_column="batchVersion", default="v1")
    sealed = models.BooleanField(default=False)
    """sealed is set to true once no more items can be added to the batch"""
    sealed_at = UtcDateTimeField(null=True, db_column="sealedAt")
    expected_count = models.IntegerField(db_column="expectedCount", default=0)
    """this is the expected number of items in the batch"""
    completed_count = models.IntegerField(db_column="completedCount", default=0)
    """this is the completed number of items in the batch. once this reaches expectedCount, and the batch is sealed, the batch is considered completed"""
    completed_at = UtcDateTimeField(null=True, db_column="completedAt")
    resumed_at = UtcDateTimeField(null=True, db_column="resumedAt")
    processing_jobs_count = models.IntegerField(db_column="processingJobsCount", default=0)
    """this is used to be able to "seal" this BatchTaskRun when all of the runs have been triggered asynchronously, and using the "parallel" processing strategy"""
    processing_jobs_expected_count = models.IntegerField(db_column="processingJobsExpectedCount", default=0)
    one_time_use_token = models.TextField(null=True, unique=True, db_column="oneTimeUseToken")
    """optional token that can be used to authenticate the task run"""
    processing_started_at = UtcDateTimeField(null=True, db_column="processingStartedAt")
    """When processing started (status changed to PROCESSING)"""
    processing_completed_at = UtcDateTimeField(null=True, db_column="processingCompletedAt")
    """When processing completed (all items processed)"""
    successful_run_count = models.IntegerField(null=True, db_column="successfulRunCount")
    """Count of successfully created runs"""
    failed_run_count = models.IntegerField(null=True, db_column="failedRunCount")
    """Count of failed run creations"""
    task_identifier = models.TextField(null=True, db_column="taskIdentifier")
    checkpoint_event = models.OneToOneField("CheckpointRestoreEvent", on_delete=models.CASCADE, related_name="batch_task_run_dependency", db_column="checkpointEventId", null=True)
    dependent_task_attempt = models.ForeignKey("TaskRunAttempt", on_delete=models.CASCADE, related_name="batch_dependencies", db_column="dependentTaskAttemptId", null=True, db_index=False)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "BatchTaskRun"
        constraints = [
            models.UniqueConstraint(fields=["one_time_use_token"], name="BatchTaskRun_oneTimeUseToken_key"),
            models.UniqueConstraint(fields=["runtime_environment_id", "idempotency_key"], name="BatchTaskRun_runtimeEnvironmentId_idempotencyKey_key"),
        ]
        indexes = [
            models.Index(fields=["dependent_task_attempt"]),
            models.Index(fields=["runtime_environment_id", "-id"]),
            models.Index(fields=["runtime_environment_id", "-created_at", "-id"]),
        ]


class BatchTaskRunItem(models.Model):
    """Used in engine V1 only"""

    id = models.TextField(primary_key=True, default=None)
    status = models.TextField(choices=BatchTaskRunItemStatus.choices, default=BatchTaskRunItemStatus.PENDING)
    batch_task_run = models.ForeignKey("BatchTaskRun", on_delete=models.CASCADE, related_name="items", db_column="batchTaskRunId", db_index=False)
    task_run = models.ForeignKey("TaskRun", on_delete=models.CASCADE, related_name="batch_items", db_column="taskRunId", db_index=False)
    task_run_attempt = models.ForeignKey("TaskRunAttempt", on_delete=models.SET_NULL, related_name="batch_task_run_items", db_column="taskRunAttemptId", null=True, db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    completed_at = UtcDateTimeField(null=True, db_column="completedAt")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "BatchTaskRunItem"
        constraints = [
            models.UniqueConstraint(fields=["batch_task_run", "task_run"], name="BatchTaskRunItem_batchTaskRunId_taskRunId_key"),
        ]
        indexes = [
            models.Index(fields=["task_run_attempt"]),
            models.Index(fields=["task_run"], name="idx_batchtaskrunitem_taskrun"),
        ]


class BatchTaskRunError(models.Model):
    """Track individual run creation failures in batch processing (Run Engine v2)"""

    id = models.TextField(primary_key=True, default=None)
    batch_task_run = models.ForeignKey("BatchTaskRun", on_delete=models.CASCADE, related_name="errors", db_column="batchTaskRunId", db_index=False)
    index = models.IntegerField()
    """Which item in the batch (0-based index)"""
    task_identifier = models.TextField(db_column="taskIdentifier")
    """The task identifier that was being triggered"""
    payload = models.TextField(null=True)
    """The payload that failed (JSON, may be truncated)"""
    options = models.JSONField(null=True)
    """The options that were used"""
    error = models.TextField()
    """Error message"""
    error_code = models.TextField(null=True, db_column="errorCode")
    """Error code if available"""
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)

    class Meta:
        db_table = "BatchTaskRunError"
        constraints = [
            models.UniqueConstraint(fields=["batch_task_run", "index"], name="BatchTaskRunError_batchTaskRunId_index_key"),
        ]
        indexes = [
            models.Index(fields=["batch_task_run"]),
        ]


class EnvironmentVariable(models.Model):
    id = models.TextField(primary_key=True, default=None)
    friendly_id = models.TextField(unique=True, db_column="friendlyId")
    key = models.TextField()
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="environment_variables", db_column="projectId", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "EnvironmentVariable"
        constraints = [
            models.UniqueConstraint(fields=["project", "key"], name="EnvironmentVariable_projectId_key_key"),
        ]


class EnvironmentVariableValue(models.Model):
    id = models.TextField(primary_key=True, default=None)
    value_reference = models.ForeignKey("SecretReference", on_delete=models.SET_NULL, related_name="environment_variable_values", db_column="valueReferenceId", null=True, db_index=False)
    variable = models.ForeignKey("EnvironmentVariable", on_delete=models.CASCADE, related_name="values", db_column="variableId", db_index=False)
    environment = models.ForeignKey("RuntimeEnvironment", on_delete=models.CASCADE, related_name="environment_variable_values", db_column="environmentId", db_index=False)
    is_secret = models.BooleanField(db_column="isSecret", default=False)
    """If true, the value is secret and cannot be revealed"""
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    version = models.IntegerField(default=1)
    last_updated_by = models.JSONField(null=True, db_column="lastUpdatedBy")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "EnvironmentVariableValue"
        constraints = [
            models.UniqueConstraint(fields=["variable", "environment"], name="EnvironmentVariableValue_variableId_environmentId_key"),
        ]
        indexes = [
            models.Index(fields=["environment"]),
            models.Index(fields=["value_reference"]),
            models.Index(fields=["environment", "-updated_at", "variable"]),
        ]


class Checkpoint(models.Model):
    id = models.TextField(primary_key=True, default=None)
    friendly_id = models.TextField(unique=True, db_column="friendlyId")
    type = models.TextField(choices=CheckpointType.choices)
    location = models.TextField()
    image_ref = models.TextField(db_column="imageRef")
    reason = models.TextField(null=True)
    metadata = models.TextField(null=True)
    run = models.ForeignKey("TaskRun", on_delete=models.CASCADE, related_name="checkpoints", db_column="runId", db_index=False)
    attempt = models.ForeignKey("TaskRunAttempt", on_delete=models.CASCADE, related_name="checkpoints", db_column="attemptId", db_index=False)
    attempt_number = models.IntegerField(null=True, db_column="attemptNumber")
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="checkpoints", db_column="projectId", db_index=False)
    runtime_environment = models.ForeignKey("RuntimeEnvironment", on_delete=models.CASCADE, related_name="checkpoints", db_column="runtimeEnvironmentId", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "Checkpoint"
        indexes = [
            models.Index(fields=["attempt"]),
            models.Index(fields=["run"]),
        ]


class CheckpointRestoreEvent(models.Model):
    id = models.TextField(primary_key=True, default=None)
    type = models.TextField(choices=CheckpointRestoreEventType.choices)
    reason = models.TextField(null=True)
    metadata = models.TextField(null=True)
    checkpoint = models.ForeignKey("Checkpoint", on_delete=models.CASCADE, related_name="events", db_column="checkpointId", db_index=False)
    run = models.ForeignKey("TaskRun", on_delete=models.CASCADE, related_name="checkpoint_restore_event", db_column="runId", db_index=False)
    attempt = models.ForeignKey("TaskRunAttempt", on_delete=models.CASCADE, related_name="checkpoint_restore_event", db_column="attemptId", db_index=False)
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="checkpoint_restore_event", db_column="projectId", db_index=False)
    runtime_environment = models.ForeignKey("RuntimeEnvironment", on_delete=models.CASCADE, related_name="checkpoint_restore_event", db_column="runtimeEnvironmentId", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "CheckpointRestoreEvent"
        indexes = [
            models.Index(fields=["checkpoint"]),
            models.Index(fields=["run"]),
        ]


class WorkerDeployment(models.Model):
    id = models.TextField(primary_key=True, default=None)
    content_hash = models.TextField(db_column="contentHash")
    friendly_id = models.TextField(unique=True, db_column="friendlyId")
    short_code = models.TextField(db_column="shortCode")
    version = models.TextField()
    runtime = models.TextField(null=True)
    runtime_version = models.TextField(null=True, db_column="runtimeVersion")
    cli_version = models.TextField(null=True, db_column="cliVersion")
    """CLI version that initiated the deploy, stamped at initialization"""
    image_reference = models.TextField(null=True, db_column="imageReference")
    image_platform = models.TextField(db_column="imagePlatform", default="linux/amd64")
    template_key = models.TextField(null=True, db_column="templateKey")
    """Key of the microVM template prebuilt from this deployment's image, set when the builder succeeds"""
    index_object_key = models.TextField(null=True, db_column="indexObjectKey")
    """Object key of the full index snapshot in the deployment index store. Set only after the upload succeeds."""
    index_sha256 = models.TextField(null=True, db_column="indexSha256")
    """sha256 hex of the snapshot's uncompressed JSON"""
    index_format_version = models.IntegerField(null=True, db_column="indexFormatVersion")
    external_build_data = models.JSONField(null=True, db_column="externalBuildData")
    build_server_metadata = models.JSONField(null=True, db_column="buildServerMetadata")
    build_env_vars = models.JSONField(null=True, db_column="buildEnvVars")
    """Encrypted build-time env vars for pre-bundled (fromBundle) deploys, as an
    EncryptedSecretValue envelope. Kept for the redeploy window, then cleared by a sweep.
    """
    status = models.TextField(choices=WorkerDeploymentStatus.choices, default=WorkerDeploymentStatus.PENDING)
    type = models.TextField(choices=WorkerDeploymentType.choices, default=WorkerDeploymentType.V1)
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="worker_deployment", db_column="projectId", db_index=False)
    environment = models.ForeignKey("RuntimeEnvironment", on_delete=models.CASCADE, related_name="worker_deployments", db_column="environmentId", db_index=False)
    worker = models.OneToOneField("BackgroundWorker", on_delete=models.CASCADE, related_name="deployment", db_column="workerId", null=True)
    triggered_by = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="deployments", db_column="triggeredById", null=True, db_index=False)
    triggered_via = models.TextField(null=True, db_column="triggeredVia")
    commit_sha = models.TextField(null=True, db_column="commitSHA")
    external_id = models.TextField(null=True, db_column="externalId")
    started_at = UtcDateTimeField(null=True, db_column="startedAt")
    installed_at = UtcDateTimeField(null=True, db_column="installedAt")
    built_at = UtcDateTimeField(null=True, db_column="builtAt")
    deployed_at = UtcDateTimeField(null=True, db_column="deployedAt")
    failed_at = UtcDateTimeField(null=True, db_column="failedAt")
    error_data = models.JSONField(null=True, db_column="errorData")
    canceled_at = UtcDateTimeField(null=True, db_column="canceledAt")
    canceled_reason = models.TextField(null=True, db_column="canceledReason")
    git = models.JSONField(null=True)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "WorkerDeployment"
        constraints = [
            models.UniqueConstraint(fields=["project", "short_code"], name="WorkerDeployment_projectId_shortCode_key"),
            models.UniqueConstraint(fields=["environment", "version"], name="WorkerDeployment_environmentId_version_key"),
        ]
        indexes = [
            models.Index(fields=["commit_sha"]),
            models.Index(fields=["environment", "created_at"]),
            models.Index(fields=["environment", "status", "id"]),
            models.Index(fields=["environment", "external_id"]),
        ]


class WorkerDeploymentPromotion(models.Model):
    id = models.TextField(primary_key=True, default=None)
    label = models.TextField()
    """This is the promotion label, e.g. "current\""""
    deployment = models.ForeignKey("WorkerDeployment", on_delete=models.CASCADE, related_name="promotions", db_column="deploymentId", db_index=False)
    environment = models.ForeignKey("RuntimeEnvironment", on_delete=models.CASCADE, related_name="worker_deployment_promotions", db_column="environmentId", db_index=False)

    class Meta:
        db_table = "WorkerDeploymentPromotion"
        constraints = [
            models.UniqueConstraint(fields=["environment", "label"], name="WorkerDeploymentPromotion_environmentId_label_key"),
        ]


class TaskSchedule(models.Model):
    """Schedules can be attached to tasks to trigger them on a schedule"""

    id = models.TextField(primary_key=True, default=None)
    type = models.TextField(choices=ScheduleType.choices, default=ScheduleType.IMPERATIVE)
    friendly_id = models.TextField(unique=True, db_column="friendlyId")
    """users see this as `id`. They start with schedule_"""
    task_identifier = models.TextField(db_column="taskIdentifier")
    """a reference to a task (not a foreign key because it's across versions)"""
    deduplication_key = models.TextField(db_column="deduplicationKey")
    """can be provided and we won't create another with the same key"""
    user_provided_deduplication_key = models.BooleanField(db_column="userProvidedDeduplicationKey", default=False)
    generator_expression = models.TextField(db_column="generatorExpression")
    """the CRON pattern"""
    generator_description = models.TextField(db_column="generatorDescription", default="")
    generator_type = models.TextField(choices=ScheduleGeneratorType.choices, db_column="generatorType", default=ScheduleGeneratorType.CRON)
    timezone_field = models.TextField(db_column="timezone", default="UTC")
    """These are IANA format string, or the default "UTC". E.g. "America/New_York\""""
    window_duration_seconds = models.IntegerField(null=True, db_column="windowDurationSeconds")
    window_percentage = models.IntegerField(null=True, db_column="windowPercentage")
    default_window_duration_seconds = models.IntegerField(null=True, db_column="defaultWindowDurationSeconds")
    minimum_window_duration_seconds = models.IntegerField(null=True, db_column="minimumWindowDurationSeconds")
    external_id = models.TextField(null=True, db_column="externalId")
    """Can be provided by the user then accessed inside a run"""
    last_run_triggered_at = UtcDateTimeField(null=True, db_column="lastRunTriggeredAt")
    """@deprecated stop writing 2026-04-30; reads moved out of code (UI now derives from cron's previous slot). Drop in follow-up."""
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="task_schedules", db_column="projectId", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    active = models.BooleanField(default=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "TaskSchedule"
        constraints = [
            models.UniqueConstraint(fields=["project", "deduplication_key"], name="TaskSchedule_projectId_deduplicationKey_key"),
        ]
        indexes = [
            models.Index(fields=["project"]),
            models.Index(fields=["project", "-created_at"]),
        ]


class TaskScheduleInstance(models.Model):
    """An instance links a schedule with an environment"""

    id = models.TextField(primary_key=True, default=None)
    task_schedule = models.ForeignKey("TaskSchedule", on_delete=models.CASCADE, related_name="instances", db_column="taskScheduleId", db_index=False)
    environment = models.ForeignKey("RuntimeEnvironment", on_delete=models.CASCADE, related_name="task_schedule_instances", db_column="environmentId", db_index=False)
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="task_schedule_instances", db_column="projectId", db_index=False)
    schedule_phase = models.IntegerField(null=True, db_column="schedulePhase")
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    active = models.BooleanField(default=True)
    last_scheduled_timestamp = UtcDateTimeField(null=True, db_column="lastScheduledTimestamp")
    """@deprecated stop writing 2026-04-30; engine derives from cron + exactScheduleTime. Drop in follow-up."""
    next_scheduled_timestamp = UtcDateTimeField(null=True, db_column="nextScheduledTimestamp")
    """@deprecated stop writing 2026-04-30; engine derives from cron + now(). Drop in follow-up."""
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "TaskScheduleInstance"
        constraints = [
            models.UniqueConstraint(fields=["task_schedule", "environment"], name="TaskScheduleInstance_taskScheduleId_environmentId_key"),
        ]
        indexes = [
            models.Index(fields=["environment"]),
            models.Index(fields=["project", "active"]),
        ]


class RuntimeEnvironmentSession(models.Model):
    id = models.TextField(primary_key=True, default=None)
    ip_address = models.TextField(db_column="ipAddress")
    environment = models.ForeignKey("RuntimeEnvironment", on_delete=models.CASCADE, related_name="sessions", db_column="environmentId", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    disconnected_at = UtcDateTimeField(null=True, db_column="disconnectedAt")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "RuntimeEnvironmentSession"


class ProjectAlertChannel(models.Model):
    id = models.TextField(primary_key=True, default=None)
    friendly_id = models.TextField(unique=True, db_column="friendlyId")
    deduplication_key = models.TextField(db_column="deduplicationKey")
    """can be provided and we won't create another with the same key"""
    user_provided_deduplication_key = models.BooleanField(db_column="userProvidedDeduplicationKey", default=False)
    integration = models.ForeignKey("OrganizationIntegration", on_delete=models.SET_NULL, related_name="alert_channels", db_column="integrationId", null=True, db_index=False)
    enabled = models.BooleanField(default=True)
    type = models.TextField(choices=ProjectAlertChannelType.choices)
    name = models.TextField()
    properties = models.JSONField()
    alert_types = ArrayField(models.TextField(choices=ProjectAlertType.choices), db_column="alertTypes")
    environment_types = ArrayField(models.TextField(choices=RuntimeEnvironmentType.choices), db_column="environmentTypes", default=project_alert_channel_environment_types_default)
    error_alert_config = models.JSONField(null=True, db_column="errorAlertConfig")
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="alert_channels", db_column="projectId", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "ProjectAlertChannel"
        constraints = [
            models.UniqueConstraint(fields=["project", "deduplication_key"], name="ProjectAlertChannel_projectId_deduplicationKey_key"),
        ]


class ProjectAlert(models.Model):
    id = models.TextField(primary_key=True, default=None)
    friendly_id = models.TextField(unique=True, db_column="friendlyId")
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="alerts", db_column="projectId", db_index=False)
    environment = models.ForeignKey("RuntimeEnvironment", on_delete=models.CASCADE, related_name="alerts", db_column="environmentId", db_index=False)
    channel = models.ForeignKey("ProjectAlertChannel", on_delete=models.CASCADE, related_name="alerts", db_column="channelId", db_index=False)
    status = models.TextField(choices=ProjectAlertStatus.choices, default=ProjectAlertStatus.PENDING)
    type = models.TextField(choices=ProjectAlertType.choices)
    task_run_attempt_id = models.TextField(null=True, db_column="taskRunAttemptId")
    task_run_id = models.TextField(null=True, db_column="taskRunId")
    worker_deployment = models.ForeignKey("WorkerDeployment", on_delete=models.CASCADE, related_name="alerts", db_column="workerDeploymentId", null=True, db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "ProjectAlert"
        indexes = [
            models.Index(fields=["channel"]),
        ]


class ProjectAlertStorage(models.Model):
    id = models.TextField(primary_key=True, default=None)
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="alert_storages", db_column="projectId", db_index=False)
    alert_channel = models.ForeignKey("ProjectAlertChannel", on_delete=models.CASCADE, related_name="alert_storages", db_column="alertChannelId", db_index=False)
    alert_type = models.TextField(choices=ProjectAlertType.choices, db_column="alertType")
    storage_id = models.TextField(db_column="storageId")
    storage_data = models.JSONField(db_column="storageData")
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "ProjectAlertStorage"
        indexes = [
            models.Index(fields=["alert_channel", "alert_type", "storage_id"]),
        ]


class OrganizationIntegration(models.Model):
    id = models.TextField(primary_key=True, default=None)
    friendly_id = models.TextField(unique=True, db_column="friendlyId")
    service = models.TextField(choices=IntegrationService.choices)
    external_organization_id = models.TextField(null=True, db_column="externalOrganizationId")
    """Identifier for external, integration's organization (e.g. Vercel's team)"""
    integration_data = models.JSONField(db_column="integrationData")
    token_reference = models.ForeignKey("SecretReference", on_delete=models.CASCADE, related_name="organization_integration", db_column="tokenReferenceId", db_index=False)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="organization_integrations", db_column="organizationId", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    deleted_at = UtcDateTimeField(null=True, db_column="deletedAt")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "OrganizationIntegration"
        indexes = [
            models.Index(fields=["external_organization_id"]),
        ]


class OrganizationProjectIntegration(models.Model):
    id = models.TextField(primary_key=True, default=None)
    organization_integration = models.ForeignKey("OrganizationIntegration", on_delete=models.CASCADE, related_name="organization_project_integration", db_column="organizationIntegrationId", db_index=False)
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="organization_project_integration", db_column="projectId", db_index=False)
    external_entity_id = models.TextField(db_column="externalEntityId")
    """Identifier for webhooks, for example Vercel's projectId"""
    integration_data = models.JSONField(db_column="integrationData")
    """Save useful data like config or external entity name"""
    installed_by = models.TextField(null=True, db_column="installedBy")
    """UserId who installed the integration"""
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    deleted_at = UtcDateTimeField(null=True, db_column="deletedAt")
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "OrganizationProjectIntegration"
        indexes = [
            models.Index(fields=["project"]),
            models.Index(fields=["project", "organization_integration"]),
            models.Index(fields=["external_entity_id"]),
        ]


class BulkActionGroup(models.Model):
    """Bulk actions, like canceling and replaying runs"""

    id = models.TextField(primary_key=True, default=None)
    friendly_id = models.TextField(unique=True, db_column="friendlyId")
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="bulk_action_groups", db_column="projectId", db_index=False)
    environment = models.ForeignKey("RuntimeEnvironment", on_delete=models.CASCADE, related_name="bulk_action_group", db_column="environmentId", null=True, db_index=False)
    """If this is set then it's a V2 Bulk Action that supports queries"""
    type = models.TextField(choices=BulkActionType.choices)
    status = models.TextField(choices=BulkActionStatus.choices, default=BulkActionStatus.PENDING)
    """When the group is created it's pending. After we've processed all the items it's completed. This does not mean the associated runs are completed."""
    query_name = models.TextField(null=True, db_column="queryName")
    """The query that will be executed to get the runs to process (if null, we are passing run ids directly)"""
    params = models.JSONField(null=True)
    """The params that will be passed to the query"""
    dedupe_key = models.TextField(null=True, db_column="dedupeKey")
    """Billing-limit dedupe key (grace hitAt or resolve job key). Null for dashboard bulk actions."""
    cursor = models.JSONField(null=True)
    """The cursor that will be passed to the query (null for the first page)"""
    success_count = models.IntegerField(db_column="successCount", default=0)
    """The number of runs that have been processed successfully"""
    failure_count = models.IntegerField(db_column="failureCount", default=0)
    """The number of runs that have been failed"""
    total_count = models.IntegerField(db_column="totalCount", default=0)
    """The total number of runs that will be processed"""
    completion_notification = models.TextField(choices=BulkActionNotificationType.choices, db_column="completionNotification", default=BulkActionNotificationType.NONE)
    user = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="bulk_actions", db_column="userId", null=True, db_index=False)
    """The userId who did the bulk action"""
    name = models.TextField(null=True)
    """The name of the bulk action"""
    completed_at = UtcDateTimeField(null=True, db_column="completedAt")
    """The time the bulk action was completed"""
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "BulkActionGroup"
        indexes = [
            models.Index(fields=["environment", "-created_at"]),
            models.Index(fields=["environment", "type", "dedupe_key"]),
            models.Index(fields=["environment", "status", "type"]),
        ]


class BulkActionItem(models.Model):
    id = models.TextField(primary_key=True, default=None)
    friendly_id = models.TextField(null=True, db_column="friendlyId")
    """@deprecated not used in new BulkActions"""
    group_id = models.TextField(db_column="groupId")
    type = models.TextField(choices=BulkActionType.choices)
    status = models.TextField(choices=BulkActionItemStatus.choices, default=BulkActionItemStatus.PENDING)
    """When the item is created it's pending. After we've processed the item it's completed. This does not mean the associated runs are completed."""
    source_run = models.ForeignKey("TaskRun", on_delete=models.CASCADE, related_name="source_bulk_action_items", db_column="sourceRunId", db_index=False)
    """The run that is the source of the action, e.g. when replaying this is the original run"""
    destination_run = models.ForeignKey("TaskRun", on_delete=models.CASCADE, related_name="destination_bulk_action_items", db_column="destinationRunId", null=True, db_index=False)
    """The run that's a result of the action, this will be set when the run has been created"""
    error = models.TextField(null=True)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "BulkActionItem"


class RealtimeStreamChunk(models.Model):
    id = models.TextField(primary_key=True, default=None)
    key = models.TextField()
    value = models.TextField()
    sequence = models.IntegerField()
    run_id = models.TextField(db_column="runId")
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)

    class Meta:
        db_table = "RealtimeStreamChunk"
        indexes = [
            models.Index(fields=["run_id"]),
            models.Index(fields=["created_at"]),
        ]


class TaskEventPartitioned(models.Model):
    """This is the unified otel span/log model that will eventually be replaced by clickhouse"""

    pk = models.CompositePrimaryKey("id", "created_at")
    id = models.TextField()
    message = models.TextField()
    """This matches the span name for a trace event, or the log body for a log event"""
    trace_id = models.TextField(db_column="traceId")
    span_id = models.TextField(db_column="spanId")
    parent_id = models.TextField(null=True, db_column="parentId")
    tracestate = models.TextField(null=True)
    is_error = models.BooleanField(db_column="isError", default=False)
    is_partial = models.BooleanField(db_column="isPartial", default=False)
    is_cancelled = models.BooleanField(db_column="isCancelled", default=False)
    service_name = models.TextField(db_column="serviceName")
    service_namespace = models.TextField(db_column="serviceNamespace")
    level = models.TextField(choices=TaskEventLevel.choices, default=TaskEventLevel.TRACE)
    kind = models.TextField(choices=TaskEventKind.choices, default=TaskEventKind.INTERNAL)
    status = models.TextField(choices=TaskEventStatus.choices, default=TaskEventStatus.UNSET)
    links = models.JSONField(null=True)
    events = models.JSONField(null=True)
    start_time = models.BigIntegerField(db_column="startTime")
    """This is the time the event started in nanoseconds since the epoch"""
    duration = models.BigIntegerField(default=0)
    """This is the duration of the event in nanoseconds"""
    attempt_id = models.TextField(null=True, db_column="attemptId")
    attempt_number = models.IntegerField(null=True, db_column="attemptNumber")
    environment_id = models.TextField(db_column="environmentId")
    environment_type = models.TextField(choices=RuntimeEnvironmentType.choices, db_column="environmentType")
    organization_id = models.TextField(db_column="organizationId")
    project_id = models.TextField(db_column="projectId")
    project_ref = models.TextField(db_column="projectRef")
    run_id = models.TextField(db_column="runId")
    run_is_test = models.BooleanField(db_column="runIsTest", default=False)
    idempotency_key = models.TextField(null=True, db_column="idempotencyKey")
    task_slug = models.TextField(db_column="taskSlug")
    task_path = models.TextField(null=True, db_column="taskPath")
    task_export_name = models.TextField(null=True, db_column="taskExportName")
    worker_id = models.TextField(null=True, db_column="workerId")
    worker_version = models.TextField(null=True, db_column="workerVersion")
    queue_id = models.TextField(null=True, db_column="queueId")
    queue_name = models.TextField(null=True, db_column="queueName")
    batch_id = models.TextField(null=True, db_column="batchId")
    properties = models.JSONField()
    """This represents all the span attributes available, like http.status_code, and special attributes like $style.icon, $output, $metadata.payload.userId, as it's used for searching and filtering"""
    metadata = models.JSONField(null=True)
    """This represents all span attributes in the $metadata namespace, like $metadata.payload"""
    style = models.JSONField(null=True)
    """This represents all span attributes in the $style namespace, like $style"""
    output = models.JSONField(null=True)
    """This represents all span attributes in the $output namespace, like $output"""
    output_type = models.TextField(null=True, db_column="outputType")
    """This represents the mimetype of the output, such as application/json or application/super+json"""
    payload = models.JSONField(null=True)
    payload_type = models.TextField(null=True, db_column="payloadType")
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    usage_duration_ms = models.IntegerField(db_column="usageDurationMs", default=0)
    usage_cost_in_cents = models.FloatField(db_column="usageCostInCents", default=0)
    machine_preset = models.TextField(null=True, db_column="machinePreset")
    machine_preset_cpu = models.FloatField(null=True, db_column="machinePresetCpu")
    machine_preset_memory = models.FloatField(null=True, db_column="machinePresetMemory")
    machine_preset_cents_per_ms = models.FloatField(null=True, db_column="machinePresetCentsPerMs")

    class Meta:
        db_table = "TaskEventPartitioned"
        indexes = [
            models.Index(fields=["trace_id"]),
            models.Index(fields=["span_id"]),
            models.Index(fields=["run_id"]),
        ]


class GithubAppInstallation(models.Model):
    id = models.TextField(primary_key=True, default=None)
    app_installation_id = models.BigIntegerField(unique=True, db_column="appInstallationId")
    target_id = models.BigIntegerField(db_column="targetId")
    target_type = models.TextField(db_column="targetType")
    account_handle = models.TextField(db_column="accountHandle")
    permissions = models.JSONField(null=True)
    repository_selection = models.TextField(choices=GithubRepositorySelection.choices, db_column="repositorySelection")
    installed_by = models.TextField(null=True, db_column="installedBy")
    installed_by_user_id = models.TextField(null=True, db_column="installedByUserId")
    organization = models.ForeignKey("Organization", on_delete=models.RESTRICT, related_name="github_app_installations", db_column="organizationId", db_index=False)
    deleted_at = UtcDateTimeField(null=True, db_column="deletedAt")
    suspended_at = UtcDateTimeField(null=True, db_column="suspendedAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "GithubAppInstallation"
        indexes = [
            models.Index(fields=["organization"]),
        ]


class GithubRepository(models.Model):
    id = models.TextField(primary_key=True, default=None)
    github_id = models.BigIntegerField(db_column="githubId")
    name = models.TextField()
    full_name = models.TextField(db_column="fullName")
    html_url = models.TextField(db_column="htmlUrl")
    private = models.BooleanField()
    default_branch = models.TextField(db_column="defaultBranch")
    installation = models.ForeignKey("GithubAppInstallation", on_delete=models.CASCADE, related_name="repositories", db_column="installationId", db_index=False)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "GithubRepository"
        constraints = [
            models.UniqueConstraint(fields=["installation", "github_id"], name="GithubRepository_installationId_githubId_key"),
        ]
        indexes = [
            models.Index(fields=["installation"]),
        ]


class TrustedDeployRepository(models.Model):
    id = models.TextField(primary_key=True, default=None)
    project = models.OneToOneField("Project", on_delete=models.CASCADE, related_name="trusted_deploy_repository", db_column="projectId")
    issuer = models.TextField(default="https://token.actions.githubusercontent.com")
    repository_owner = models.TextField(db_column="repositoryOwner")
    repository_name = models.TextField(db_column="repositoryName")
    repository_id = models.TextField(null=True, db_column="repositoryId")
    repository_owner_id = models.TextField(null=True, db_column="repositoryOwnerId")
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "TrustedDeployRepository"


class TrustedDeployEnvironment(models.Model):
    id = models.TextField(primary_key=True, default=None)
    trusted_deploy_repository = models.ForeignKey("TrustedDeployRepository", on_delete=models.CASCADE, related_name="environments", db_column="trustedDeployRepositoryId", db_index=False)
    runtime_environment = models.OneToOneField("RuntimeEnvironment", on_delete=models.CASCADE, related_name="trusted_deploy_grant", db_column="runtimeEnvironmentId")
    enabled = models.BooleanField(default=False)
    scope = models.TextField(choices=TrustedDeployScope.choices)
    scope_value = models.TextField(null=True, db_column="scopeValue")
    additional_claims = models.JSONField(db_column="additionalClaims", default=dict)
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "TrustedDeployEnvironment"
        indexes = [
            models.Index(fields=["trusted_deploy_repository"]),
        ]


class ConnectedGithubRepository(models.Model):
    id = models.TextField(primary_key=True, default=None)
    project = models.OneToOneField("Project", on_delete=models.CASCADE, related_name="connected_github_repository", db_column="projectId")
    repository = models.ForeignKey("GithubRepository", on_delete=models.CASCADE, related_name="connected_github_repository", db_column="repositoryId", db_index=False)
    branch_tracking = models.JSONField(db_column="branchTracking")
    preview_deployments_enabled = models.BooleanField(db_column="previewDeploymentsEnabled", default=True)
    build_filters = models.JSONField(null=True, db_column="buildFilters")
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "ConnectedGithubRepository"
        constraints = [
            models.UniqueConstraint(fields=["project"], name="ConnectedGithubRepository_projectId_key"),
        ]
        indexes = [
            models.Index(fields=["repository"]),
        ]


class ImpersonationAuditLog(models.Model):
    id = models.TextField(primary_key=True, default=None)
    action = models.TextField(choices=ImpersonationAuditLogAction.choices)
    admin = models.ForeignKey("User", on_delete=models.CASCADE, related_name="impersonations_performed", db_column="adminId", db_index=False)
    """The admin user who initiated/ended the impersonation"""
    target = models.ForeignKey("User", on_delete=models.CASCADE, related_name="impersonations_received", db_column="targetId", db_index=False)
    """The user being impersonated"""
    ip_address = models.TextField(null=True, db_column="ipAddress")
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)

    class Meta:
        db_table = "ImpersonationAuditLog"
        indexes = [
            models.Index(fields=["admin"]),
            models.Index(fields=["target"]),
            models.Index(fields=["created_at"]),
        ]


class SupportAccessRequest(models.Model):
    id = models.TextField(primary_key=True, default=None)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="support_access_requests", db_column="organizationId", db_index=False)
    requested_by = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="support_access_requests_made", db_column="requestedById", null=True, db_index=False)
    reason = models.TextField()
    status = models.TextField(choices=SupportAccessRequestStatus.choices, default=SupportAccessRequestStatus.PENDING)
    approved_by = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="support_access_requests_approved", db_column="approvedById", null=True, db_index=False)
    approved_at = UtcDateTimeField(null=True, db_column="approvedAt")
    expires_at = UtcDateTimeField(null=True, db_column="expiresAt")
    cancelled_by_id = models.TextField(null=True, db_column="cancelledById")
    cancelled_at = UtcDateTimeField(null=True, db_column="cancelledAt")
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "SupportAccessRequest"
        indexes = [
            models.Index(fields=["organization", "status", "expires_at"]),
            models.Index(fields=["requested_by"]),
            models.Index(fields=["approved_by"]),
        ]


class CustomerQuery(models.Model):
    id = models.TextField(primary_key=True, default=None)
    query = models.TextField()
    """The TSQL query text that was executed"""
    scope = models.TextField(choices=CustomerQueryScope.choices)
    """The scope of the query (determines which tenant IDs were used)"""
    stats = models.JSONField()
    """Query execution statistics from ClickHouse"""
    title = models.TextField(null=True)
    """AI-generated title summarizing the query"""
    source = models.TextField(choices=CustomerQuerySource.choices, default=CustomerQuerySource.DASHBOARD)
    """Where the query originated from"""
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="customer_queries", db_column="organizationId", db_index=False)
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="customer_queries", db_column="projectId", null=True, db_index=False)
    environment = models.ForeignKey("RuntimeEnvironment", on_delete=models.CASCADE, related_name="customer_queries", db_column="environmentId", null=True, db_index=False)
    user = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="customer_queries", db_column="userId", null=True, db_index=False)
    """Optional user who executed the query (null for API calls)"""
    filter_period = models.TextField(null=True, db_column="filterPeriod")
    """Time filter settings used when the query was executed
    Period like "7d", "24h", etc.
    """
    filter_from = UtcDateTimeField(null=True, db_column="filterFrom")
    """Custom start date (ISO 8601)"""
    filter_to = UtcDateTimeField(null=True, db_column="filterTo")
    """Custom end date (ISO 8601)"""
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)

    class Meta:
        db_table = "CustomerQuery"
        indexes = [
            models.Index(fields=["organization", "-created_at"]),
            models.Index(fields=["created_at"]),
        ]


class IntegrationDeployment(models.Model):
    id = models.TextField(primary_key=True, default=None)
    integration_name = models.TextField(db_column="integrationName")
    """For example Vercel"""
    integration_deployment_id = models.TextField(db_column="integrationDeploymentId")
    """External ID"""
    commit_sha = models.TextField(db_column="commitSHA")
    worker_deployment = models.ForeignKey("WorkerDeployment", on_delete=models.SET_NULL, related_name="integration_deployments", db_column="deploymentId", null=True, db_index=False)
    status = models.TextField(null=True)
    """External deployment status"""
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "IntegrationDeployment"
        indexes = [
            models.Index(fields=["commit_sha"]),
            models.Index(fields=["worker_deployment"]),
        ]


class MetricsDashboard(models.Model):
    """A user-defined metrics dashboard"""

    id = models.TextField(primary_key=True, default=None)
    friendly_id = models.TextField(unique=True, db_column="friendlyId")
    title = models.TextField()
    description = models.TextField(default="")
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="metrics_dashboards", db_column="organizationId", db_index=False)
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="metrics_dashboards", db_column="projectId", null=True, db_index=False)
    owner = models.ForeignKey("User", on_delete=models.SET_NULL, related_name="metrics_dashboards", db_column="ownerId", null=True, db_index=False)
    """Who created the dashboard"""
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    layout = models.TextField()
    """JSON that defines the config, queries, layout and config of all widgets.
    There will be a version field for the format.
    """
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "MetricsDashboard"
        indexes = [
            models.Index(fields=["project", "-created_at"]),
        ]


class LlmModel(models.Model):
    """A known LLM model or model pattern for cost tracking"""

    id = models.TextField(primary_key=True, default=None)
    friendly_id = models.TextField(unique=True)
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="llm_models", null=True, db_index=False)
    model_name = models.TextField()
    match_pattern = models.TextField()
    start_date = UtcDateTimeField(null=True)
    source = models.TextField(default="default")
    created_at = UtcDateTimeField(default=timezone.now)
    updated_at = UtcDateTimeField(auto_now=True)
    provider = models.TextField(null=True)
    description = models.TextField(null=True)
    context_window = models.IntegerField(null=True)
    max_output_tokens = models.IntegerField(null=True)
    capabilities = ArrayField(models.TextField(), default=list)
    is_hidden = models.BooleanField(default=False)
    base_model_name = models.TextField(null=True)
    pricing_unit = models.TextField(null=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "llm_models"
        constraints = [
            models.UniqueConstraint(fields=["project", "model_name", "start_date"], name="llm_models_project_id_model_name_start_date_key"),
        ]
        indexes = [
            models.Index(fields=["project"]),
        ]


class LlmPricingTier(models.Model):
    """A pricing tier for a model (supports volume-based or conditional pricing)"""

    id = models.TextField(primary_key=True, default=None)
    model = models.ForeignKey("LlmModel", on_delete=models.CASCADE, related_name="pricing_tiers", db_index=False)
    name = models.TextField()
    is_default = models.BooleanField(default=True)
    priority = models.IntegerField(default=0)
    conditions = models.JSONField(default=list)

    class Meta:
        db_table = "llm_pricing_tiers"
        constraints = [
            models.UniqueConstraint(fields=["model", "priority"], name="llm_pricing_tiers_model_id_priority_key"),
            models.UniqueConstraint(fields=["model", "name"], name="llm_pricing_tiers_model_id_name_key"),
        ]


class LlmPrice(models.Model):
    """A price point for a usage type within a pricing tier"""

    id = models.TextField(primary_key=True, default=None)
    model = models.ForeignKey("LlmModel", on_delete=models.CASCADE, related_name="prices", db_index=False)
    pricing_tier = models.ForeignKey("LlmPricingTier", on_delete=models.CASCADE, related_name="prices", db_index=False)
    usage_type = models.TextField()
    price = models.DecimalField(max_digits=20, decimal_places=12)

    class Meta:
        db_table = "llm_prices"
        constraints = [
            models.UniqueConstraint(fields=["model", "usage_type", "pricing_tier"], name="llm_prices_model_id_usage_type_pricing_tier_id_key"),
        ]


class PlatformNotification(models.Model):
    """Admin-created notification definitions"""

    id = models.TextField(primary_key=True, default=None)
    friendly_id = models.TextField(unique=True, db_column="friendlyId")
    title = models.TextField()
    """Admin-facing title for identification in admin tools"""
    payload = models.JSONField()
    """Versioned JSON rendering content (see payload schema in spec)"""
    surface = models.TextField(choices=PlatformNotificationSurface.choices)
    scope = models.TextField(choices=PlatformNotificationScope.choices)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="platform_notifications", db_column="userId", null=True, db_index=False)
    """Set when scope = USER"""
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="platform_notifications", db_column="organizationId", null=True, db_index=False)
    """Set when scope = ORGANIZATION"""
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="platform_notifications", db_column="projectId", null=True, db_index=False)
    """Set when scope = PROJECT"""
    starts_at = UtcDateTimeField(db_column="startsAt")
    """When notification becomes active"""
    ends_at = UtcDateTimeField(db_column="endsAt")
    """When notification expires"""
    cli_max_days_after_first_seen = models.IntegerField(null=True, db_column="cliMaxDaysAfterFirstSeen")
    """CLI: auto-expire N days after user first saw it"""
    cli_max_show_count = models.IntegerField(null=True, db_column="cliMaxShowCount")
    """CLI: auto-expire after shown N times to user"""
    cli_show_every = models.IntegerField(null=True, db_column="cliShowEvery")
    """CLI: only show every N-th request (e.g. cliShowEvery=3 means show on 3rd, 6th, 9th…)"""
    priority = models.IntegerField(default=0)
    """Ordering within same scope level (higher = more important)"""
    is_draft = models.BooleanField(db_column="isDraft", default=False)
    """Staged-but-unpublished. While true the notification is hidden from all
    user-facing reads regardless of startsAt/endsAt; publishing sets real
    dates and flips this to false.
    """
    archived_at = UtcDateTimeField(null=True, db_column="archivedAt")
    """Soft delete"""
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "PlatformNotification"
        indexes = [
            models.Index(fields=["surface", "scope", "starts_at"]),
        ]


class PlatformNotificationInteraction(models.Model):
    """Per-user tracking of notification views and dismissals"""

    id = models.TextField(primary_key=True, default=None)
    notification = models.ForeignKey("PlatformNotification", on_delete=models.CASCADE, related_name="interactions", db_column="notificationId", db_index=False)
    user = models.ForeignKey("User", on_delete=models.CASCADE, related_name="platform_notification_interactions", db_column="userId", db_index=False)
    first_seen_at = UtcDateTimeField(db_column="firstSeenAt", default=timezone.now)
    """Set by beacon/CLI GET on first impression"""
    show_count = models.IntegerField(db_column="showCount", default=1)
    """Times shown (incremented per beacon/CLI view)"""
    webapp_dismissed_at = UtcDateTimeField(null=True, db_column="webappDismissedAt")
    """User dismissed in webapp"""
    webapp_clicked_at = UtcDateTimeField(null=True, db_column="webappClickedAt")
    """User clicked a link in webapp"""
    cli_dismissed_at = UtcDateTimeField(null=True, db_column="cliDismissedAt")
    """Auto-dismissed or expired in CLI"""
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "PlatformNotificationInteraction"
        constraints = [
            models.UniqueConstraint(fields=["notification", "user"], name="PlatformNotificationInteraction_notificationId_userId_key"),
        ]


class ErrorGroupState(models.Model):
    """*
    * Error group state is used to track when a user has interacted with an error (ignored/resolved)
    * The actual error data is in ClickHouse.
    """

    id = models.TextField(primary_key=True, default=None)
    organization = models.ForeignKey("Organization", on_delete=models.CASCADE, related_name="error_group_states", db_column="organizationId", db_index=False)
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="error_group_states", db_column="projectId", db_index=False)
    environment = models.ForeignKey("RuntimeEnvironment", on_delete=models.CASCADE, related_name="error_group_states", db_column="environmentId", null=True, db_index=False)
    """*
    * You can ignore/resolve an error across all environments, or specific ones
    """
    task_identifier = models.TextField(db_column="taskIdentifier")
    error_fingerprint = models.TextField(db_column="errorFingerprint")
    status = models.TextField(choices=ErrorGroupStatus.choices, default=ErrorGroupStatus.UNRESOLVED)
    ignored_until = UtcDateTimeField(null=True, db_column="ignoredUntil")
    """*
    * Error is ignored until this date
    """
    ignored_until_occurrence_rate = models.IntegerField(null=True, db_column="ignoredUntilOccurrenceRate")
    """*
    * Error is ignored until this occurrence rate
    """
    ignored_until_total_occurrences = models.IntegerField(null=True, db_column="ignoredUntilTotalOccurrences")
    """*
    * Error is ignored until this total occurrences
    """
    ignored_at_occurrence_count = models.BigIntegerField(null=True, db_column="ignoredAtOccurrenceCount")
    """Total occurrence count at the time the error was ignored (from ClickHouse).
    Used with ignoredUntilTotalOccurrences to compute occurrences since ignoring.
    """
    ignored_at = UtcDateTimeField(null=True, db_column="ignoredAt")
    """*
    * Error was ignored at this date
    """
    ignored_reason = models.TextField(null=True, db_column="ignoredReason")
    """*
    * Reason for ignoring the error
    """
    ignored_by_user_id = models.TextField(null=True, db_column="ignoredByUserId")
    """*
    * User who ignored the error
    """
    resolved_at = UtcDateTimeField(null=True, db_column="resolvedAt")
    """*
    * Error was resolved at this date
    """
    resolved_in_version = models.TextField(null=True, db_column="resolvedInVersion")
    """*
    * Error was resolved in this version
    """
    resolved_by = models.TextField(null=True, db_column="resolvedBy")
    """*
    * User who resolved the error
    """
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "ErrorGroupState"
        constraints = [
            models.UniqueConstraint(fields=["environment", "task_identifier", "error_fingerprint"], name="ErrorGroupState_environmentId_taskIdentifier_errorFingerprint_key"),
        ]
        indexes = [
            models.Index(fields=["environment", "status"]),
        ]


class TaskIdentifier(models.Model):
    id = models.TextField(primary_key=True, default=None)
    runtime_environment = models.ForeignKey("RuntimeEnvironment", on_delete=models.CASCADE, related_name="task_identifiers", db_column="runtimeEnvironmentId", db_index=False)
    project = models.ForeignKey("Project", on_delete=models.CASCADE, related_name="task_identifiers", db_column="projectId", db_index=False)
    slug = models.TextField()
    current_trigger_source = models.TextField(choices=TaskTriggerSource.choices, db_column="currentTriggerSource", default=TaskTriggerSource.STANDARD)
    current_worker = models.ForeignKey("BackgroundWorker", on_delete=models.SET_NULL, related_name="task_identifiers", db_column="currentWorkerId", null=True, db_index=False)
    first_seen_at = UtcDateTimeField(db_column="firstSeenAt", default=timezone.now)
    last_seen_at = UtcDateTimeField(db_column="lastSeenAt", default=timezone.now)
    is_in_latest_deployment = models.BooleanField(db_column="isInLatestDeployment", default=True)

    class Meta:
        db_table = "TaskIdentifier"
        constraints = [
            models.UniqueConstraint(fields=["runtime_environment", "slug"], name="TaskIdentifier_runtimeEnvironmentId_slug_key"),
        ]
        indexes = [
            models.Index(fields=["runtime_environment", "is_in_latest_deployment"]),
        ]


class OrganizationDataStore(models.Model):
    """Defines org-scoped data store overrides (e.g. dedicated ClickHouse for HIPAA orgs).
    Multiple organizations can share a single data store row via organizationIds.
    """

    id = models.TextField(primary_key=True, default=None)
    key = models.TextField(unique=True)
    """Human-readable unique key (e.g. "hipaa-clickhouse-us-east")"""
    organization_ids = ArrayField(models.TextField(), db_column="organizationIds")
    """Organization IDs that use this data store"""
    kind = models.TextField(choices=DataStoreKind.choices)
    config = models.JSONField()
    """Versioned config JSON. Structure is discriminated by the top-level `version` field."""
    created_at = UtcDateTimeField(db_column="createdAt", default=timezone.now)
    updated_at = UtcDateTimeField(db_column="updatedAt", auto_now=True)
    objects = AutoNowQuerySet.as_manager()

    class Meta:
        db_table = "OrganizationDataStore"
        indexes = [
            models.Index(fields=["kind"]),
        ]


class BackgroundWorkerToBackgroundWorkerFile(models.Model):
    pk = models.CompositePrimaryKey("a_id", "b_id")
    a = models.ForeignKey("BackgroundWorker", on_delete=models.CASCADE, related_name="+", db_column="A", db_index=False)
    b = models.ForeignKey("BackgroundWorkerFile", on_delete=models.CASCADE, related_name="+", db_column="B", db_index=False)

    class Meta:
        db_table = "_BackgroundWorkerToBackgroundWorkerFile"
        indexes = [
            models.Index(fields=["b"]),
        ]


class BackgroundWorkerToTaskQueue(models.Model):
    pk = models.CompositePrimaryKey("a_id", "b_id")
    a = models.ForeignKey("BackgroundWorker", on_delete=models.CASCADE, related_name="+", db_column="A", db_index=False)
    b = models.ForeignKey("TaskQueue", on_delete=models.CASCADE, related_name="+", db_column="B", db_index=False)

    class Meta:
        db_table = "_BackgroundWorkerToTaskQueue"
        indexes = [
            models.Index(fields=["b"]),
        ]


class WaitpointRunConnections(models.Model):
    pk = models.CompositePrimaryKey("a_id", "b_id")
    a = models.ForeignKey("TaskRun", on_delete=models.CASCADE, related_name="+", db_column="A", db_index=False)
    b = models.ForeignKey("Waitpoint", on_delete=models.CASCADE, related_name="+", db_column="B", db_index=False)

    class Meta:
        db_table = "_WaitpointRunConnections"
        indexes = [
            models.Index(fields=["b"]),
        ]


class CompletedWaitpoints(models.Model):
    pk = models.CompositePrimaryKey("a_id", "b_id")
    a = models.ForeignKey("TaskRunExecutionSnapshot", on_delete=models.CASCADE, related_name="+", db_column="A", db_index=False)
    b = models.ForeignKey("Waitpoint", on_delete=models.CASCADE, related_name="+", db_column="B", db_index=False)

    class Meta:
        db_table = "_completedWaitpoints"
        indexes = [
            models.Index(fields=["b"]),
        ]
