from sqlalchemy import ARRAY, BigInteger, Dialect, Enum, ForeignKey, ForeignKeyConstraint, Index, Integer, MetaData, Numeric, String, Text, TypeDecorator, text
from sqlalchemy.dialects.postgresql import JSONB, TIMESTAMP
from sqlalchemy.orm import DeclarativeBase, Mapped, mapped_column, relationship
from typing import Any, Optional
from decimal import Decimal as DecimalType
from datetime import datetime, timezone
import cuid  # type: ignore[import-untyped, unused-ignore]
import nanoid


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


class Account(Base):
    __tablename__ = "Account"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="Account_userId_fkey"))
    type: Mapped[str] = mapped_column(default="oauth", server_default=text("'oauth'"))
    provider: Mapped[str]
    issuer: Mapped[Optional[str]]
    """WHO asserted `providerAccountId`. Better Auth keys accounts by
    `(provider, providerAccountId)`; auth's account create hook writes this
    (modules/auth/specs/account-issuer.feature). Nullable only so the column
    could be added to a populated table; the migration backfilled every row.
    """
    provider_account_id: Mapped[str] = mapped_column("providerAccountId")
    refresh_token: Mapped[Optional[str]]
    access_token: Mapped[Optional[str]]
    expires_at: Mapped[Optional[datetime]]
    ext_expires_in: Mapped[Optional[int]]
    token_type: Mapped[Optional[str]]
    scope: Mapped[Optional[str]]
    id_token: Mapped[Optional[str]]
    session_state: Mapped[Optional[str]]
    password: Mapped[Optional[str]]
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("Account_userId_idx", "userId"),
        Index("Account_issuer_providerAccountId_idx", "issuer", "providerAccountId"),
        Index("Account_provider_providerAccountId_key", "provider", "providerAccountId", unique=True),
    )

    user: Mapped["User"] = relationship(back_populates="accounts")

class AccountCredential(Base):
    __tablename__ = "AccountCredential"

    id: Mapped[str] = mapped_column(primary_key=True)
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="AccountCredential_userId_fkey"))
    provider: Mapped[str]
    password: Mapped[Optional[str]]
    access_token: Mapped[Optional[str]] = mapped_column("accessToken")
    refresh_token: Mapped[Optional[str]] = mapped_column("refreshToken")
    id_token: Mapped[Optional[str]] = mapped_column("idToken")
    access_token_expires_at: Mapped[Optional[datetime]] = mapped_column("accessTokenExpiresAt")
    refresh_token_expires_at: Mapped[Optional[datetime]] = mapped_column("refreshTokenExpiresAt")
    scope: Mapped[Optional[str]]
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("AccountCredential_userId_idx", "userId"),
    )

    user: Mapped["User"] = relationship(back_populates="account_credentials")

class Session(Base):
    __tablename__ = "Session"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    session_token: Mapped[str] = mapped_column("sessionToken")
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="Session_userId_fkey"))
    expires: Mapped[datetime]
    ip_address: Mapped[Optional[str]] = mapped_column("ipAddress")
    user_agent: Mapped[Optional[str]] = mapped_column("userAgent")
    identifier_id: Mapped[Optional[str]] = mapped_column("identifierId")
    """Which of the user's sign-in methods minted this session (D06). NULL on
    every session that predates the column and on any mint we could not
    attribute - an ordinary value, not an error. Nothing reads it to decide
    whether the session is valid, so no session was ended to add it.
    """
    amr: Mapped[list[str]] = mapped_column(ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))
    """What the sign-in that minted this session PROVED, as RFC 8176 method
    references (D06). Empty on every session minted before the column, and
    on any ordinary password sign-in. An organization that requires a second
    factor reads this when a member reaches its data; it is never read at
    mint and never ends a session.
    """
    started_by_sign_up: Mapped[bool] = mapped_column("startedBySignUp", default=False, server_default=text("false"))
    """Minted while the account had never signed in: the session its sign-up
    started, which skips the secure-account offer. False on every later one
    and on every session that predates the column (specs/identity/passkeys.feature).
    """
    impersonating: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))
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
    actor_user_id: Mapped[Optional[str]] = mapped_column("actorUserId")
    """The impersonation claims (D06), aligned with the authz Principal
    {actor, subject}. The actor is the operator who really made the request
    and the subject is the person whose access they are borrowing; both are
    recorded on every authorization decision, so the audit trail can answer
    who really did it. These replace the legacy `impersonating` JSON blob,
    which is dropped in the release AFTER this one.

    NULL on every ordinary session, which is all of them: an impersonation
    is a handful of rows at a time and nothing else reads these columns.
    """
    subject_user_id: Mapped[Optional[str]] = mapped_column("subjectUserId")
    impersonation_reason: Mapped[Optional[str]] = mapped_column("impersonationReason")
    """Why the operator started it. Required to start one
    (specs/features/backoffice-user-impersonation-reason.feature) and kept
    beside both people rather than only in the audit log, so a request under
    impersonation carries its own justification.
    """
    impersonation_expires_at: Mapped[Optional[datetime]] = mapped_column("impersonationExpiresAt")
    """When the borrowed access lapses. A session outlives its impersonation
    window: past this the claims are ignored and the operator is themselves
    again, with nothing ended.
    """
    last_seen_at: Mapped[Optional[datetime]] = mapped_column("lastSeenAt")
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
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("Session_sessionToken_key", "sessionToken", unique=True),
        Index("Session_userId_idx", "userId"),
        Index("Session_identifierId_idx", "identifierId"),
    )

    user: Mapped["User"] = relationship(back_populates="sessions")

class UserErasure(Base):
    """One row per erased account, holding only its creation time: no id, name or address. It keeps
    the install's first-created account knowable after an erasure deletes that row.
    Spec: modules/user/specs/user.feature
    """

    __tablename__ = "UserErasure"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    account_created_at: Mapped[datetime] = mapped_column("accountCreatedAt")
    erased_at: Mapped[datetime] = mapped_column("erasedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("UserErasure_accountCreatedAt_idx", "accountCreatedAt"),
    )

class User(Base):
    __tablename__ = "User"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    name: Mapped[Optional[str]]
    email: Mapped[Optional[str]]
    email_verified: Mapped[bool] = mapped_column("emailVerified", default=False, server_default=text("false"))
    signup_confirmation_pending: Mapped[bool] = mapped_column("signupConfirmationPending", default=False, server_default=text("false"))
    passkey_signup_claim_hash: Mapped[Optional[str]] = mapped_column("passkeySignupClaimHash")
    image: Mapped[Optional[str]]
    pending_sso_setup: Mapped[bool] = mapped_column("pendingSsoSetup", default=False, server_default=text("false"))
    user_hash_key: Mapped[Optional[str]] = mapped_column("userHashKey")
    two_factor_enabled: Mapped[bool] = mapped_column("twoFactorEnabled", default=False, server_default=text("false"))
    """Owned by the two-factor plugin (D06), which is why it is a plain
    boolean rather than a state: the plugin flips it when a setup is
    confirmed and when it is turned off, and challenges every sign-in while
    it is true. The lifecycle - and the history - lives on MfaEnrollment.
    """
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    last_login_at: Mapped[Optional[datetime]] = mapped_column("lastLoginAt")
    deactivated_at: Mapped[Optional[datetime]] = mapped_column("deactivatedAt")
    last_home_path: Mapped[Optional[str]] = mapped_column("lastHomePath")
    traces_explorer_tour_dismissed_at: Mapped[Optional[datetime]] = mapped_column("tracesExplorerTourDismissedAt")
    langy_code_access_preference: Mapped[Optional[str]] = mapped_column("langyCodeAccessPreference")
    """How Langy should reach this person's code when a task needs a change
    (ADR-129). "github" means they chose the pull request path and asked to
    be remembered; NULL means ask. A shared local folder is never stored
    here, because a folder has to be shared again each time.
    Spec: specs/langy/langy-code-access.feature
    """
    passkey_nudge_dismissed_at: Mapped[Optional[datetime]] = mapped_column("passkeyNudgeDismissedAt")
    """When this person last said "not now" to being offered a passkey
    (ADR-120). On the account rather than in browser storage for two
    reasons: a per-browser record forgets on a new device, which is exactly
    where the offer should be MORE eager, and an interval nobody can enforce
    is not an interval. NULL means never asked.
    """
    join_offer_dismissed_domains: Mapped[list[str]] = mapped_column("joinOfferDismissedDomains", ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))
    """The domains this person has said "no thanks" to being offered their
    colleagues on (D12). Per domain rather than one flag: somebody who
    changes jobs and verifies a second address should be offered their new
    employer's organization, not silenced forever by a decision they made
    about the old one. Empty means never dismissed anything.
    """
    notification_preferences: Mapped[dict[str, Any]] = mapped_column("notificationPreferences", JSONB, default=lambda: {}, server_default=text("'{}'"))
    """This person's answer about browser notifications, one entry per topic:
    { "langy": "enabled" | "declined" }. A topic that is absent was never
    answered. The browser's own permission is not stored here.
    Spec: specs/langy/langy-notifications.feature
    """
    password_reset_required_at: Mapped[Optional[datetime]] = mapped_column("passwordResetRequiredAt")
    """When an operator marked this person's password as no longer trusted. A
    password sign-in opens no session while it is set; any write of a new
    password hash clears it. Spec: modules/auth/specs/forced-password-reset.feature
    """

    __table_args__ = (
        Index("User_email_key", "email", unique=True),
    )

    accounts: Mapped[list["Account"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")
    account_credentials: Mapped[list["AccountCredential"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")
    sessions: Mapped[list["Session"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")
    two_factors: Mapped[list["TwoFactor"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")
    passkeys: Mapped[list["Passkey"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")
    team_memberships: Mapped[list["TeamUser"]] = relationship(passive_deletes="all", back_populates="user")
    org_memberships: Mapped[list["OrganizationUser"]] = relationship(passive_deletes="all", back_populates="user")
    scim_user_resources: Mapped[list["ScimUserResource"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")
    annotation: Mapped[list["Annotation"]] = relationship(back_populates="user")
    share_links: Mapped[list["ShareLink"]] = relationship(back_populates="user")
    workflow: Mapped[list["Workflow"]] = relationship(back_populates="published_by")
    workflow_version: Mapped[list["WorkflowVersion"]] = relationship(passive_deletes="all", back_populates="author")
    annotation_queues: Mapped[list["AnnotationQueueMembers"]] = relationship(passive_deletes="all", back_populates="user")
    assigned_queue_items: Mapped[list["AnnotationQueueItem"]] = relationship(foreign_keys="AnnotationQueueItem.user_id", back_populates="user")
    created_queue_items: Mapped[list["AnnotationQueueItem"]] = relationship(foreign_keys="AnnotationQueueItem.created_by_user_id", back_populates="created_by_user")
    llm_prompt_config_versions: Mapped[list["LlmPromptConfigVersion"]] = relationship(back_populates="author")
    scenarios: Mapped[list["Scenario"]] = relationship(back_populates="last_updated_by")
    invite_requests: Mapped[list["OrganizationInvite"]] = relationship(back_populates="requested_by_user")
    secrets_created: Mapped[list["ProjectSecret"]] = relationship(foreign_keys="ProjectSecret.created_by_id", passive_deletes="all", back_populates="created_by")
    secrets_updated: Mapped[list["ProjectSecret"]] = relationship(foreign_keys="ProjectSecret.updated_by_id", passive_deletes="all", back_populates="updated_by")
    saved_views: Mapped[list["SavedView"]] = relationship(back_populates="user")
    version_tags_created: Mapped[list["PromptTagAssignment"]] = relationship(foreign_keys="PromptTagAssignment.created_by_id", back_populates="created_by")
    version_tags_updated: Mapped[list["PromptTagAssignment"]] = relationship(foreign_keys="PromptTagAssignment.updated_by_id", back_populates="updated_by")
    prompt_tags_created: Mapped[list["PromptTag"]] = relationship(foreign_keys="PromptTag.created_by_id", back_populates="created_by")
    prompt_tags_updated: Mapped[list["PromptTag"]] = relationship(foreign_keys="PromptTag.updated_by_id", back_populates="updated_by")
    group_memberships: Mapped[list["GroupMembership"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")
    role_bindings: Mapped[list["RoleBinding"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")
    scim_external_ids: Mapped[list["ScimExternalId"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="user")
    api_keys: Mapped[list["ApiKey"]] = relationship(passive_deletes="all", back_populates="user")
    gateway_budgets_created: Mapped[list["GatewayBudget"]] = relationship(passive_deletes="all", back_populates="created_by")
    principal_virtual_keys: Mapped[list["VirtualKey"]] = relationship(back_populates="principal_user")
    personal_teams: Mapped[list["Team"]] = relationship(back_populates="owner_user")
    personal_projects: Mapped[list["Project"]] = relationship(back_populates="owner_user")
    routing_policies_created: Mapped[list["RoutingPolicy"]] = relationship(foreign_keys="RoutingPolicy.created_by_id", back_populates="created_by")
    routing_policies_updated: Mapped[list["RoutingPolicy"]] = relationship(foreign_keys="RoutingPolicy.updated_by_id", back_populates="updated_by")
    gateway_guardrails_created: Mapped[list["GatewayGuardrail"]] = relationship(foreign_keys="GatewayGuardrail.created_by_id", back_populates="created_by")
    gateway_guardrails_updated: Mapped[list["GatewayGuardrail"]] = relationship(foreign_keys="GatewayGuardrail.updated_by_id", back_populates="updated_by")
    ingestion_sources_created: Mapped[list["IngestionSource"]] = relationship(back_populates="created_by")
    anomaly_rules_created: Mapped[list["AnomalyRule"]] = relationship(back_populates="created_by")
    pinned_traces: Mapped[list["PinnedTrace"]] = relationship(back_populates="user")
    trace_edit_overlays_created: Mapped[list["TraceEditOverlay"]] = relationship(foreign_keys="TraceEditOverlay.created_by_id", back_populates="created_by")
    trace_edit_overlays_updated: Mapped[list["TraceEditOverlay"]] = relationship(foreign_keys="TraceEditOverlay.updated_by_id", back_populates="updated_by")

class VerificationToken(Base):
    __tablename__ = "VerificationToken"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    identifier: Mapped[str]
    token: Mapped[str]
    """Not unique alone: better-auth stores plain values here (a user id, a counter). See WEB-19201."""
    expires: Mapped[datetime]
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("VerificationToken_identifier_token_key", "identifier", "token", unique=True),
    )

class Identifier(Base):
    __tablename__ = "Identifier"

    id: Mapped[str] = mapped_column(primary_key=True)
    user_id: Mapped[str] = mapped_column("userId")
    provider: Mapped[str]
    value: Mapped[Optional[str]]
    domain: Mapped[Optional[str]]
    identifier_hash: Mapped[Optional[str]] = mapped_column("identifierHash")
    account_id: Mapped[Optional[str]] = mapped_column("accountId")
    provider_id: Mapped[Optional[str]] = mapped_column("providerId")
    issuer: Mapped[Optional[str]]
    provider_account_id: Mapped[Optional[str]] = mapped_column("providerAccountId")
    state: Mapped[str]
    connection_id: Mapped[Optional[str]] = mapped_column("connectionId")
    verified_at: Mapped[Optional[datetime]] = mapped_column("verifiedAt")
    attached_at: Mapped[datetime] = mapped_column("attachedAt")
    detached_at: Mapped[Optional[datetime]] = mapped_column("detachedAt")
    last_used_at: Mapped[Optional[datetime]] = mapped_column("lastUsedAt")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("Identifier_userId_idx", "userId"),
        Index("Identifier_identifierHash_idx", "identifierHash"),
        Index("Identifier_value_idx", "value"),
        Index("Identifier_domain_idx", "domain"),
        Index("Identifier_provider_providerAccountId_idx", "provider", "providerAccountId"),
        Index("Identifier_issuer_providerAccountId_idx", "issuer", "providerAccountId"),
        Index("Identifier_accountId_idx", "accountId"),
    )

class IdentitySignInSwitch(Base):
    """Whether a user signs in on the identifier path: identity's own switch, set when the
    identifier backfill proves the user and pinned off by an operator's rollback. No row is the
    legacy path. Spec: modules/identity/specs/sign-in-switch.feature. Column refs only, no FK.
    """

    __tablename__ = "IdentitySignInSwitch"

    user_id: Mapped[str] = mapped_column("userId", primary_key=True)
    on_identifiers: Mapped[bool] = mapped_column("onIdentifiers")
    pinned: Mapped[bool] = mapped_column(default=False, server_default=text("false"))
    switched_at: Mapped[datetime] = mapped_column("switchedAt")
    last_applied_event_id: Mapped[Optional[str]] = mapped_column("lastAppliedEventId")
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

class IdentifierReservation(Base):
    __tablename__ = "IdentifierReservation"

    normalized_value: Mapped[str] = mapped_column("normalizedValue", primary_key=True)
    user_id: Mapped[str] = mapped_column("userId")
    identifier_id: Mapped[str] = mapped_column("identifierId")
    command_id: Mapped[str] = mapped_column("commandId")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("IdentifierReservation_userId_idx", "userId"),
        Index("IdentifierReservation_identifierId_idx", "identifierId"),
    )

class IdentityProjectionCursor(Base):
    __tablename__ = "IdentityProjectionCursor"

    user_id: Mapped[str] = mapped_column("userId", primary_key=True)
    last_event_id: Mapped[str] = mapped_column("lastEventId")
    accepted_at: Mapped[datetime] = mapped_column("acceptedAt")
    occurred_at: Mapped[datetime] = mapped_column("occurredAt")
    projection_version: Mapped[str] = mapped_column("projectionVersion")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

class TwoFactor(Base):
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

    __tablename__ = "TwoFactor"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="TwoFactor_userId_fkey"))
    secret: Mapped[str]
    backup_codes: Mapped[str] = mapped_column("backupCodes")
    verified: Mapped[bool] = mapped_column(default=True, server_default=text("true"))
    failed_verification_count: Mapped[int] = mapped_column("failedVerificationCount", default=0, server_default=text("0"))
    locked_until: Mapped[Optional[datetime]] = mapped_column("lockedUntil")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("TwoFactor_userId_idx", "userId"),
        Index("TwoFactor_secret_idx", "secret"),
    )

    user: Mapped["User"] = relationship(back_populates="two_factors")

class SignInAttemptLock(Base):
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

    __tablename__ = "SignInAttemptLock"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    identifier_hash: Mapped[str] = mapped_column("identifierHash")
    """Keyed hash of the normalised address. Never the address itself."""
    failed_count: Mapped[int] = mapped_column("failedCount", default=0, server_default=text("0"))
    """Consecutive failures since the last success. Reset by a success, not by
    time - "five consecutive", not five ever.
    """
    locked_until: Mapped[Optional[datetime]] = mapped_column("lockedUntil")
    """When the current lock lifts. Null when not locked."""
    consecutive_lockouts: Mapped[int] = mapped_column("consecutiveLockouts", default=0, server_default=text("0"))
    """Locks taken since the last success. At five this stops lifting by
    itself and becomes an incident - see `heldForReview`.
    """
    held_for_review: Mapped[bool] = mapped_column("heldForReview", default=False, server_default=text("false"))
    """Held past its lock because the count of locks reached the ceiling the
    control names. Cleared by an administrator, or by the owner proving
    mailbox control through a password reset - the second is what keeps a
    held account from being a denial of service anyone can inflict by
    typing a colleague's address wrong often enough.
    """
    user_id: Mapped[Optional[str]] = mapped_column("userId")
    """Whose account this resolved to at the time, when it resolved to one.
    Null for an address with no account - those rows are the reason this is
    nullable, and the reason this table is not hung off User.
    """
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("SignInAttemptLock_identifierHash_key", "identifierHash", unique=True),
        Index("SignInAttemptLock_lockedUntil_idx", "lockedUntil"),
        Index("SignInAttemptLock_userId_idx", "userId"),
    )

class Passkey(Base):
    """The passkey plugin's own table (D07). Protocol state, like TwoFactor: the
    public key and the signature counter are the plugin's to read and write.
    A passkey also becomes an `Identifier` row with provider `passkey`, which
    is what makes it a sign-in method like any other - this table is the
    credential, that one is the identity.

    `aaguid` names the authenticator MODEL, never a device or a person, and is
    only present at registration. Privacy-preserving platforms report all
    zeroes, so a label derived from it is best-effort by construction.
    """

    __tablename__ = "Passkey"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    name: Mapped[Optional[str]]
    public_key: Mapped[str] = mapped_column("publicKey")
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="Passkey_userId_fkey"))
    credential_id: Mapped[str] = mapped_column("credentialID")
    counter: Mapped[int] = mapped_column(default=0, server_default=text("0"))
    device_type: Mapped[str] = mapped_column("deviceType")
    backed_up: Mapped[bool] = mapped_column("backedUp", default=False, server_default=text("false"))
    transports: Mapped[Optional[str]]
    aaguid: Mapped[Optional[str]]
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("Passkey_userId_idx", "userId"),
        Index("Passkey_credentialID_idx", "credentialID"),
    )

    user: Mapped["User"] = relationship(back_populates="passkeys")

class MfaEnrollment(Base):
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

    __tablename__ = "MfaEnrollment"

    user_id: Mapped[str] = mapped_column("userId", primary_key=True)
    """The user. Also the aggregate id and the tenant - one identity history
    per person (ADR-029 §4).
    """
    enrollment_id: Mapped[Optional[str]] = mapped_column("enrollmentId")
    """The current enrollment, or NULL for somebody who never started one."""
    method: Mapped[Optional[str]]
    """totp. Never sms."""
    state: Mapped[str]
    """NONE | PENDING | ENABLED | EXPIRED | DISABLED"""
    enrolled_at: Mapped[Optional[datetime]] = mapped_column("enrolledAt")
    confirmed_at: Mapped[Optional[datetime]] = mapped_column("confirmedAt")
    expired_at: Mapped[Optional[datetime]] = mapped_column("expiredAt")
    disabled_at: Mapped[Optional[datetime]] = mapped_column("disabledAt")
    disabled_via: Mapped[Optional[str]] = mapped_column("disabledVia")
    """password+totp | org-admin - who ended it, and on whose authority."""
    backup_code_count: Mapped[int] = mapped_column("backupCodeCount", default=0, server_default=text("0"))
    consumed_backup_code_indexes: Mapped[list[int]] = mapped_column("consumedBackupCodeIndexes", ARRAY(Integer), nullable=True, default=lambda: [], server_default=text("ARRAY[]::integer[]"))
    """Positions already spent, ascending. Never codes."""
    failed_count: Mapped[int] = mapped_column("failedCount", default=0, server_default=text("0"))
    occurred_at: Mapped[datetime] = mapped_column("occurredAt")
    last_event_id: Mapped[str] = mapped_column("lastEventId")
    accepted_at: Mapped[datetime] = mapped_column("acceptedAt")
    projection_version: Mapped[str] = mapped_column("projectionVersion")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("MfaEnrollment_state_idx", "state"),
    )

class SsoConnection(Base):
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

    __tablename__ = "SsoConnection"

    id: Mapped[str] = mapped_column(primary_key=True)
    organization_id: Mapped[str] = mapped_column("organizationId")
    type: Mapped[str]
    """oidc | saml"""
    state: Mapped[str]
    """DRAFT | CLAIMED | APPROVED | REJECTED | DISCARDED |
    VERIFICATION_PENDING | VERIFIED | ACTIVE | SUSPENDED |
    TEARDOWN_PENDING | TORN_DOWN
    """
    claimed_domains: Mapped[list[str]] = mapped_column("claimedDomains", ARRAY(Text), nullable=True)
    domain_claims: Mapped[dict[str, Any]] = mapped_column("domainClaims", JSONB, default=lambda: [], server_default=text("'[]'"))
    """[{ domain, state, claimedAtMs, claimedByActorId, decidedAtMs,
    decidedByActorId, authority, waitedMs, note }] - every claim this
    connection has made and how long each waited to be decided. The tier-3
    approval queue is a read of this column, and the epic's Open Q2 wants
    that latency measured from the day the queue exists.
    """
    approved_domains: Mapped[list[str]] = mapped_column("approvedDomains", ARRAY(Text), nullable=True)
    verified_domains: Mapped[list[str]] = mapped_column("verifiedDomains", ARRAY(Text), nullable=True)
    """The only domains that ever route, and only while state is ACTIVE."""
    lapsed_domains: Mapped[list[str]] = mapped_column("lapsedDomains", ARRAY(Text), nullable=True)
    """The subset of `verifiedDomains` whose published record stayed missing
    through its grace window (ADR-123). They still ROUTE - everyone already
    here signs in unchanged - and they stop vouching for anybody NEW: no
    provisioning on first sign-in, no walking in by domain. A column rather
    than a read of the JSON below, because both questions are asked on a
    sign-in path and have to be one indexed predicate.
    """
    domain_verifications: Mapped[dict[str, Any]] = mapped_column("domainVerifications", JSONB, default=lambda: [], server_default=text("'[]'"))
    """[{ domain, method, actorId, verifiedAtMs, proofState, firstAbsentAtMs,
    graceEndsAtMs }] - what proved each verified domain, who, and whether
    that evidence is still there. Kept on the row rather than only in the
    log so that an operator-attested domain can never be read as one the
    customer proved, and a wavering one never as one nothing is doubting.
    """
    pending_verification: Mapped[Optional[dict[str, Any]]] = mapped_column("pendingVerification", JSONB(none_as_null=True))
    """{ domain, method, tokenHash } while a ceremony is in flight."""
    idp_metadata: Mapped[dict[str, Any]] = mapped_column("idpMetadata", JSONB)
    arrival_policy: Mapped[str] = mapped_column("arrivalPolicy", default="refuse", server_default=text("'refuse'"))
    """admit | request | refuse - what this connection does with somebody who
    signs in through it and is not a member yet (ADR-117 3). Stated at
    registration and changed by the setup journey, so it is never absent
    and no reader has to decide what absence means. This was two fields
    for a while - a boolean `allowsJit` beside it, kept in step by the
    fold - and they disagreed about the middle answer.
    """
    allows_jit: Mapped[bool] = mapped_column("allowsJit", default=False, server_default=text("false"))
    """DEAD, AND STAYING FOR ONE RELEASE. Nothing reads or writes this; the
    routing repository derives the same answer from `arrivalPolicy`. It is
    declared so the schema still describes the table, and left in place so
    the PREVIOUS release's pods — which select it by name on every
    connection read — keep working through the rollout. Dropping it in the
    release that stopped reading it would stop single sign-on routing and
    the connection projection for the length of the deploy. Dropped by
    `20260827120000_drop_sso_connection_allows_jit` in the next one.
    """
    arrival_policy_decided_at: Mapped[Optional[datetime]] = mapped_column("arrivalPolicyDecidedAt")
    """When somebody CHOSE the policy above, and null while the registration
    default stands. A different fact from the policy rather than a second
    copy of it: a connection always has a behaviour, and separately
    somebody has or has not decided on it. Going live waits for the
    deciding - "turn everybody away" is a decision too.
    """
    source: Mapped[str]
    """self-serve | legacy-grandfathered"""
    test_login_account_id: Mapped[Optional[str]] = mapped_column("testLoginAccountId")
    replaces_connection_id: Mapped[Optional[str]] = mapped_column("replacesConnectionId")
    """The grandfathered connection this direct connection replaces. Null for
    ordinary connections and for the grandfathered predecessor itself.

    Partial uniqueness for nonterminal replacements is installed by the
    flattened migration; Prisma cannot express that predicate here.
    """
    migration_phase: Mapped[Optional[str]] = mapped_column("migrationPhase", Enum("SETUP", "GRACE_LEGACY", "GRACE_DIRECT", "FINALIZING", "FINALIZED", name="SsoConnectionMigrationPhase"))
    """SETUP | GRACE_LEGACY | GRACE_DIRECT | FINALIZING | FINALIZED.
    Null means this connection is not a legacy-to-direct replacement.
    """
    grace_started_at: Mapped[Optional[datetime]] = mapped_column("graceStartedAt")
    route_changed_at: Mapped[Optional[datetime]] = mapped_column("routeChangedAt")
    """When an administrator last selected the normal route. This is audit
    evidence only: routing is derived from `migrationPhase`, never recency.
    """
    finalization_requested_at: Mapped[Optional[datetime]] = mapped_column("finalizationRequestedAt")
    finalized_at: Mapped[Optional[datetime]] = mapped_column("finalizedAt")
    rejection: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))
    """{ domain, note } - the last ops rejection, read back on a re-claim."""
    created_by: Mapped[Optional[str]] = mapped_column("createdBy")
    tear_down_after: Mapped[Optional[datetime]] = mapped_column("tearDownAfter")
    """While TEARDOWN_PENDING: when the process manager's wake completes it."""
    occurred_at: Mapped[datetime] = mapped_column("occurredAt")
    last_event_id: Mapped[str] = mapped_column("lastEventId")
    accepted_at: Mapped[datetime] = mapped_column("acceptedAt")
    projection_version: Mapped[str] = mapped_column("projectionVersion")
    created_at: Mapped[datetime] = mapped_column("createdAt")
    """Event-derived, not database-managed: `@default(now())` / `@updatedAt`
    would make a replayed row differ from the row it rebuilds, which is
    exactly the parity this projection has to keep.
    """
    updated_at: Mapped[datetime] = mapped_column("updatedAt")

    __table_args__ = (
        Index("SsoConnection_organizationId_idx", "organizationId"),
        Index("SsoConnection_organizationId_migrationPhase_idx", "organizationId", "migrationPhase"),
        Index("SsoConnection_state_idx", "state"),
        Index("SsoConnection_verifiedDomains_idx", "verifiedDomains"),
        Index("SsoConnection_lapsedDomains_idx", "lapsedDomains"),
    )

    reproof_cursor: Mapped[Optional["SsoConnectionReproofCursor"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="connection")

class SsoConnectionRegistrationSlot(Base):
    """The one-statement pre-event registration lock. `SsoConnection` remains
    event truth; this row only decides which command may append that truth.
    A slot whose projected connection is terminal may be atomically replaced.
    """

    __tablename__ = "SsoConnectionRegistrationSlot"

    organization_id: Mapped[str] = mapped_column("organizationId", primary_key=True)
    kind: Mapped[str] = mapped_column(primary_key=True)
    """legacy | direct"""
    connection_id: Mapped[str] = mapped_column("connectionId")
    replaces_connection_id: Mapped[Optional[str]] = mapped_column("replacesConnectionId")
    command_id: Mapped[str] = mapped_column("commandId")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("SsoConnectionRegistrationSlot_connectionId_key", "connectionId", unique=True),
    )

class SsoConnectionReproofCursor(Base):
    """Operational round-robin position for ADR-123's domain re-proof sweep.

    This is deliberately not a column on `SsoConnection`: that row is a
    whole-row event projection and its fold is its only writer, while a
    healthy re-read states no fact. Keeping the scheduling cursor beside the
    projection preserves both rules and still lets a fixed-size sweep revisit
    every connection, including one whose proof has started wavering.
    """

    __tablename__ = "SsoConnectionReproofCursor"

    connection_id: Mapped[str] = mapped_column("connectionId", ForeignKey("SsoConnection.id", ondelete="CASCADE", onupdate="CASCADE", name="SsoConnectionReproofCursor_connectionId_fkey"), primary_key=True)
    last_reproof_at: Mapped[datetime] = mapped_column("lastReproofAt")

    __table_args__ = (
        Index("SsoConnectionReproofCursor_lastReproofAt_idx", "lastReproofAt"),
    )

    connection: Mapped["SsoConnection"] = relationship(back_populates="reproof_cursor")

class SsoCredential(Base):
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

    __tablename__ = "SsoCredential"

    id: Mapped[str] = mapped_column(primary_key=True)
    organization_id: Mapped[str] = mapped_column("organizationId")
    connection_id: Mapped[str] = mapped_column("connectionId")
    """The connection whose registration minted this. Not a relation, matching
    `SsoConnection` itself: a projection rebuilt from events must not depend
    on relational state the log does not carry.
    """
    kind: Mapped[str]
    """oidc-client-id | oidc-client-secret | saml-idp-config"""
    ciphertext: Mapped[str]
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("SsoCredential_organizationId_connectionId_idx", "organizationId", "connectionId"),
    )

class SsoVerifiedDomain(Base):
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

    __tablename__ = "SsoVerifiedDomain"

    domain: Mapped[str] = mapped_column(primary_key=True)
    organization_id: Mapped[str] = mapped_column("organizationId")

    __table_args__ = (
        Index("SsoVerifiedDomain_organizationId_idx", "organizationId"),
        Index("SsoVerifiedDomain_domain_organizationId_key", "domain", "organizationId", unique=True),
    )

    holders: Mapped[list["SsoVerifiedDomainHolder"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="ownership")

class SsoVerifiedDomainHolder(Base):
    """A connection currently entitled to use an organization-owned verified
    domain. Usually there is one holder. During an Auth0 migration there are
    exactly two: the grandfathered predecessor and its direct replacement.

    The projection repository enforces that relationship before inserting a
    second holder. The flattened migration also installs partial unique live
    legacy/direct slots on `SsoConnection`, closing concurrent third-writer
    races that Prisma cannot describe declaratively.
    """

    __tablename__ = "SsoVerifiedDomainHolder"

    domain: Mapped[str] = mapped_column(primary_key=True)
    connection_id: Mapped[str] = mapped_column("connectionId", primary_key=True)
    organization_id: Mapped[str] = mapped_column("organizationId")

    __table_args__ = (
        Index("SsoVerifiedDomainHolder_connectionId_idx", "connectionId"),
        Index("SsoVerifiedDomainHolder_organizationId_idx", "organizationId"),
        ForeignKeyConstraint(["domain", "organizationId"], ["SsoVerifiedDomain.domain", "SsoVerifiedDomain.organizationId"], ondelete="CASCADE", onupdate="CASCADE", name="SsoVerifiedDomainHolder_domain_organizationId_fkey"),
    )

    ownership: Mapped["SsoVerifiedDomain"] = relationship(back_populates="holders")

class SsoAuthenticationActivity(Base):
    """A successful callback through one SSO connection. Identifier.lastUsedAt
    also advances during routing and resolution, so it cannot prove that a
    member actually authenticated through the replacement before cutover.
    """

    __tablename__ = "SsoAuthenticationActivity"

    id: Mapped[str] = mapped_column(primary_key=True)
    organization_id: Mapped[str] = mapped_column("organizationId")
    connection_id: Mapped[str] = mapped_column("connectionId")
    user_id: Mapped[str] = mapped_column("userId")
    authenticated_at: Mapped[datetime] = mapped_column("authenticatedAt")
    provider_account_id: Mapped[Optional[str]] = mapped_column("providerAccountId")
    """The provider-side account the assertion named; null before #7633's arrival door recorded it."""

    __table_args__ = (
        Index("SsoAuthenticationActivity_organizationId_connectionId_authe_idx", "organizationId", "connectionId", "authenticatedAt"),
        Index("SsoAuthenticationActivity_organizationId_userId_idx", "organizationId", "userId"),
    )

class SsoProvider(Base):
    __tablename__ = "SsoProvider"

    id: Mapped[str] = mapped_column(primary_key=True)
    issuer: Mapped[str]
    oidc_config: Mapped[Optional[str]] = mapped_column("oidcConfig")
    saml_config: Mapped[Optional[str]] = mapped_column("samlConfig")
    user_id: Mapped[Optional[str]] = mapped_column("userId")
    provider_id: Mapped[str] = mapped_column("providerId")
    organization_id: Mapped[Optional[str]] = mapped_column("organizationId")
    domain: Mapped[str]

    __table_args__ = (
        Index("SsoProvider_providerId_key", "providerId", unique=True),
        Index("SsoProvider_organizationId_idx", "organizationId"),
        Index("SsoProvider_domain_idx", "domain"),
    )

class SsoBreakGlassBinding(Base):
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

    __tablename__ = "SsoBreakGlassBinding"

    id: Mapped[str] = mapped_column(primary_key=True)
    organization_id: Mapped[str] = mapped_column("organizationId")
    user_id: Mapped[str] = mapped_column("userId")
    """Who holds the way in."""
    granted_by_user_id: Mapped[str] = mapped_column("grantedByUserId")
    """Who granted it. A way back in is never self-served."""
    granted_at: Mapped[datetime] = mapped_column("grantedAt")
    expires_at: Mapped[datetime] = mapped_column("expiresAt")
    superseded_at: Mapped[Optional[datetime]] = mapped_column("supersededAt")
    """Set when a renewal replaced this row; the row itself is never edited."""
    renewed_from_id: Mapped[Optional[str]] = mapped_column("renewedFromId")
    """The binding this one renewed, so the previous end date stays readable."""
    warned_days: Mapped[list[int]] = mapped_column("warnedDays", ARRAY(Integer), nullable=True)
    """Which of the fourteen, seven and one day warnings have been sent, so a
    sweep that runs twice in a day never says the same thing twice.
    """
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("SsoBreakGlassBinding_organizationId_expiresAt_idx", "organizationId", "expiresAt"),
        Index("SsoBreakGlassBinding_userId_idx", "userId"),
    )

class SsoActivationRecoveryReservation(Base):
    """A durable claim on one live recovery path while activation is between its
    guard and projected ACTIVE fact. It has no expiry: only projection of the
    matching activation, or an explicit cancellation, may consume it.
    """

    __tablename__ = "SsoActivationRecoveryReservation"

    command_id: Mapped[str] = mapped_column("commandId", primary_key=True)
    organization_id: Mapped[str] = mapped_column("organizationId")
    connection_id: Mapped[str] = mapped_column("connectionId")
    created_at: Mapped[datetime] = mapped_column("createdAt")

    __table_args__ = (
        Index("SsoActivationRecoveryReservation_organizationId_idx", "organizationId"),
        Index("SsoActivationRecoveryReservation_organizationId_connectionI_key", "organizationId", "connectionId", unique=True),
    )

class JoinRequest(Base):
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

    __tablename__ = "JoinRequest"

    id: Mapped[str] = mapped_column(primary_key=True)
    user_id: Mapped[str] = mapped_column("userId")
    organization_id: Mapped[str] = mapped_column("organizationId")
    domain: Mapped[str]
    """The normalized domain the match was made on."""
    state: Mapped[str]
    """PENDING | APPROVED | REJECTED | EXPIRED | WITHDRAWN"""
    matched_via: Mapped[str] = mapped_column("matchedVia")
    """How the domain was matched: `verified-identifier-domain` or `sso-connection-domain`."""
    origin: Mapped[str] = mapped_column(default="web", server_default=text("'web'"))
    """Where the request was made: `web` or `cli` (ADR-171 v6). A request from
    the terminal lands a Developer seat whatever the joiner seat says.
    """
    expires_at: Mapped[Optional[datetime]] = mapped_column("expiresAt")
    """While PENDING: when the day-14 wake lapses it. Null once resolved."""
    resolved_at: Mapped[Optional[datetime]] = mapped_column("resolvedAt")
    resolved_by_type: Mapped[Optional[str]] = mapped_column("resolvedByType")
    """user | policy | invite - who ended it. `policy` is domain auto-join's
    own principal, deliberately not the one SSO auto-join uses.
    """
    resolved_by_id: Mapped[Optional[str]] = mapped_column("resolvedById")
    """The admin's user id, `domain-auto`, or the invitation id."""
    withdrawal_cause: Mapped[Optional[str]] = mapped_column("withdrawalCause")
    """user | invite-accepted, on a withdrawal."""
    connection_id: Mapped[Optional[str]] = mapped_column("connectionId")
    """The SSO connection a single sign-on arrival was admitted through; null otherwise."""
    occurred_at: Mapped[datetime] = mapped_column("occurredAt")
    last_event_id: Mapped[str] = mapped_column("lastEventId")
    accepted_at: Mapped[datetime] = mapped_column("acceptedAt")
    projection_version: Mapped[str] = mapped_column("projectionVersion")
    created_at: Mapped[datetime] = mapped_column("createdAt")
    """Event-derived, not database-managed: `@default(now())` / `@updatedAt`
    would make a replayed row differ from the row it rebuilds, which is
    exactly the parity this projection has to keep.
    """
    updated_at: Mapped[datetime] = mapped_column("updatedAt")

    __table_args__ = (
        Index("JoinRequest_organizationId_state_idx", "organizationId", "state"),
        Index("JoinRequest_userId_state_idx", "userId", "state"),
        Index("JoinRequest_userId_organizationId_idx", "userId", "organizationId"),
        Index("JoinRequest_domain_idx", "domain"),
    )

class TeamUser(Base):
    __tablename__ = "TeamUser"

    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="RESTRICT", onupdate="CASCADE", name="TeamUser_userId_fkey"), primary_key=True)
    team_id: Mapped[str] = mapped_column("teamId", ForeignKey("Team.id", ondelete="RESTRICT", onupdate="CASCADE", name="TeamUser_teamId_fkey"), primary_key=True)
    role: Mapped[str] = mapped_column(Enum("ADMIN", "MEMBER", "VIEWER", "CUSTOM", name="TeamUserRole"))
    assigned_role_id: Mapped[Optional[str]] = mapped_column("assignedRoleId", ForeignKey("CustomRole.id", ondelete="SET NULL", onupdate="CASCADE", name="TeamUser_assignedRoleId_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("TeamUser_teamId_idx", "teamId"),
        Index("TeamUser_userId_idx", "userId"),
        Index("TeamUser_assignedRoleId_idx", "assignedRoleId"),
    )

    assigned_role: Mapped[Optional["CustomRole"]] = relationship(back_populates="assigned_users")
    user: Mapped["User"] = relationship(back_populates="team_memberships")
    team: Mapped["Team"] = relationship(back_populates="members")

class OrganizationUser(Base):
    __tablename__ = "OrganizationUser"

    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="RESTRICT", onupdate="CASCADE", name="OrganizationUser_userId_fkey"), primary_key=True)
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="RESTRICT", onupdate="CASCADE", name="OrganizationUser_organizationId_fkey"), primary_key=True)
    role: Mapped[str] = mapped_column(Enum("ADMIN", "MEMBER", "EXTERNAL", "DEVELOPER", name="OrganizationUserRole"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    department_id: Mapped[Optional[str]] = mapped_column("departmentId")
    """Department assignment for spend attribution (the people lens).
    Single-valued, org-scoped, accounting-only (never gates access).
    Plain column ref (relationMode="prisma"); the service validates the
    id belongs to this org. See departments.feature.
    """
    disabled_at: Mapped[Optional[datetime]] = mapped_column("disabledAt")
    """Set when an admin disables this membership to stay within the licensed
    seat count. The person loses access to this organization only, keeps
    their role, department and everything they did, and frees their seat.
    Reversible: re-enabling restores access if a seat is available.
    Seat counting excludes disabled memberships. See
    seat-reconciliation.feature.
    """
    pending_sso_grant_id: Mapped[Optional[str]] = mapped_column("pendingSsoGrantId")
    """Stable grant intent for an unfinished automatic SSO admission.
    Cleared after the authoritative grant is observed; never inferred from missing access.
    """
    membership_stamp: Mapped[str] = mapped_column("membershipStamp", server_default=text("gen_random_uuid()::text"))
    """A new generation for each membership lifetime. Stamped USER grant facts
    must match the current active membership before they can be projected;
    removing and rejoining therefore cannot revive a delayed old attach.
    """
    full_seat_since: Mapped[Optional[datetime]] = mapped_column("fullSeatSince")
    """When this membership started holding the full seat it holds now; null means createdAt.
    Anchors the 30-day seat hold (specs/licensing/seat-rotation.feature).
    """
    seat_pending_since: Mapped[Optional[datetime]] = mapped_column("seatPendingSince")
    """Set when admission wrote this membership disabled because no seat was free: the person
    is waiting for one. Cleared when the membership is enabled.
    """

    __table_args__ = (
        Index("OrganizationUser_organizationId_idx", "organizationId"),
        Index("OrganizationUser_userId_idx", "userId"),
        Index("OrganizationUser_departmentId_idx", "departmentId"),
    )

    user: Mapped["User"] = relationship(back_populates="org_memberships")
    organization: Mapped["Organization"] = relationship(back_populates="members")

class SeatHold(Base):
    """A full seat still counted after the person who took it stopped holding it (removed,
    disabled or demoted), until 30 days after they took it (specs/licensing/seat-rotation.feature).
    One per organization and person; a return within the hold takes the seat back.
    """

    __tablename__ = "SeatHold"

    id: Mapped[str] = mapped_column(primary_key=True)
    organization_id: Mapped[str] = mapped_column("organizationId")
    user_id: Mapped[str] = mapped_column("userId")
    email: Mapped[Optional[str]]
    """Display copy for the Users list, kept because a removed person has no membership row."""
    name: Mapped[Optional[str]]
    seat_taken_at: Mapped[datetime] = mapped_column("seatTakenAt")
    """When the person took the seat this hold keeps."""
    held_until: Mapped[datetime] = mapped_column("heldUntil")
    """seatTakenAt plus 30 days; expired holds are ignored by every count and deleted hourly."""
    reason: Mapped[str]
    """REMOVED | DISABLED | DEMOTED"""
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("SeatHold_heldUntil_idx", "heldUntil"),
        Index("SeatHold_organizationId_userId_key", "organizationId", "userId", unique=True),
    )

class SeatRequest(Base):
    """A person asked for a full seat the organization had no room for (modules/organization/specs/
    seat-requests.feature). One row per organization and address; reopening reuses the row.
    """

    __tablename__ = "SeatRequest"

    id: Mapped[str] = mapped_column(primary_key=True)
    organization_id: Mapped[str] = mapped_column("organizationId")
    email: Mapped[str]
    """Lowercased address of the person asking."""
    user_id: Mapped[Optional[str]] = mapped_column("userId")
    """Null when only the address is known (a sign-in the identity provider refused)."""
    name: Mapped[Optional[str]]
    status: Mapped[str]
    """OPEN | APPROVED | DECLINED"""
    source: Mapped[str]
    """pending | lite | idp_refused: how the person was left when the request opened."""
    idp_refused: Mapped[bool] = mapped_column("idpRefused", default=False, server_default=text("false"))
    """The identity provider also refused this person's sign-in."""
    idp_kind: Mapped[Optional[str]] = mapped_column("idpKind")
    """okta | entra | google | onelogin | jumpcloud | ping | auth0 | oidc | saml"""
    sales_notified_at: Mapped[Optional[datetime]] = mapped_column("salesNotifiedAt")
    """When sales were told about this request; reset when it reopens."""
    opened_at: Mapped[datetime] = mapped_column("openedAt")
    resolved_at: Mapped[Optional[datetime]] = mapped_column("resolvedAt")
    resolved_by_user_id: Mapped[Optional[str]] = mapped_column("resolvedByUserId")
    seat_pending_since: Mapped[Optional[datetime]] = mapped_column("seatPendingSince")
    """Set when an approval claims the request, cleared once its seat is given; an approval
    still pending long after is one whose seat never landed, and is reopened.
    """
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("SeatRequest_organizationId_status_idx", "organizationId", "status"),
        Index("SeatRequest_userId_idx", "userId"),
        Index("SeatRequest_organizationId_email_key", "organizationId", "email", unique=True),
    )

class OrganizationSeatLease(Base):
    """Held by the one seat-taking write running on an organization (seat-requests.feature). A
    row past expiresAt is free to take; a holder deletes its own row when its write is done.
    """

    __tablename__ = "OrganizationSeatLease"

    organization_id: Mapped[str] = mapped_column("organizationId", primary_key=True)
    holder: Mapped[str]
    """Random per hold, so a release deletes only its own lease."""
    expires_at: Mapped[datetime] = mapped_column("expiresAt")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

class OrganizationAccessNotice(Base):
    """The mail log for seat requests and refused sign-ins: at most one row per organization,
    kind, address and UTC day, each mailed at once.
    """

    __tablename__ = "OrganizationAccessNotice"

    id: Mapped[str] = mapped_column(primary_key=True)
    organization_id: Mapped[str] = mapped_column("organizationId")
    kind: Mapped[str]
    """seat_request | idp_refused"""
    email: Mapped[str]
    name: Mapped[Optional[str]]
    idp: Mapped[Optional[str]]
    day_key: Mapped[str] = mapped_column("dayKey")
    """UTC day the notice was recorded on, YYYY-MM-DD."""
    idp_refused: Mapped[bool] = mapped_column("idpRefused", default=False, server_default=text("false"))
    """Whether the seat request it announces was also refused by the identity provider."""
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    mailed_at: Mapped[Optional[datetime]] = mapped_column("mailedAt")
    """Set once the notice was handed to the mailer; null if a send never started."""

    __table_args__ = (
        Index("OrganizationAccessNotice_organizationId_kind_email_dayKey_key", "organizationId", "kind", "email", "dayKey", unique=True),
    )

class Team(Base):
    __tablename__ = "Team"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    name: Mapped[str]
    slug: Mapped[str]
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="RESTRICT", onupdate="CASCADE", name="Team_organizationId_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    archived_at: Mapped[Optional[datetime]] = mapped_column("archivedAt")
    is_personal: Mapped[bool] = mapped_column("isPersonal", default=False, server_default=text("false"))
    owner_user_id: Mapped[Optional[str]] = mapped_column("ownerUserId", String(255), ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="Team_ownerUserId_fkey"))
    department_id: Mapped[Optional[str]] = mapped_column("departmentId")
    """Department assignment inherited as the default for members and
    projects that have none of their own. Accounting-only.
    """

    __table_args__ = (
        Index("Team_slug_key", "slug", unique=True),
        Index("Team_organizationId_idx", "organizationId"),
        Index("Team_isPersonal_idx", "isPersonal"),
        Index("Team_departmentId_idx", "departmentId"),
    )

    organization: Mapped["Organization"] = relationship(back_populates="teams")
    owner_user: Mapped[Optional["User"]] = relationship(back_populates="personal_teams")
    members: Mapped[list["TeamUser"]] = relationship(passive_deletes="all", back_populates="team")
    projects: Mapped[list["Project"]] = relationship(passive_deletes="all", back_populates="team")
    ingestion_sources: Mapped[list["IngestionSource"]] = relationship(back_populates="team")
    ai_tool_entry_teams: Mapped[list["AiToolEntryTeam"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="team")

class Organization(Base):
    __tablename__ = "Organization"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    name: Mapped[str]
    kind: Mapped[str] = mapped_column(default="customer", server_default=text("'customer'"))
    """`customer` for every organization a person founds or an operator provisions; `system` for
    the install's one hidden organization the Ops checkup canaries run in, never listed to a person.
    Spec: modules/ops/specs/checkup-canary-project.feature.
    """
    phone_number: Mapped[Optional[str]] = mapped_column("phoneNumber")
    slug: Mapped[str]
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    usage_spending_max_limit: Mapped[Optional[int]] = mapped_column("usageSpendingMaxLimit")
    dataset_attachment_max_mb: Mapped[Optional[int]] = mapped_column("datasetAttachmentMaxMb")
    """The largest file a dataset image or file cell accepts for this
    organization, in mebibytes. Null means the default. Set by an operator
    from the back office only; every other dataset size limit scales with it.
    """
    private_click_house_label: Mapped[Optional[str]] = mapped_column("privateClickHouseLabel")
    """The label of this organization's `CLICKHOUSE_URL__<label>__<id>` route, recorded by the
    organization module the first time it is seen. Once set, a missing route refuses the
    organization instead of serving it from the shared ClickHouse. Null: never private.
    Spec: specs/private-dataplane/clickhouse-routing.feature.
    """
    private_s3_label: Mapped[Optional[str]] = mapped_column("privateS3Label")
    """The same for `DATAPLANE_S3__<label>__<id>`. Spec: specs/private-dataplane/s3-routing.feature."""
    private_dataplane_seen_at: Mapped[Optional[datetime]] = mapped_column("privateDataplaneSeenAt")
    """When either label was first recorded."""
    max_session_duration_days: Mapped[int] = mapped_column("maxSessionDurationDays", default=0, server_default=text("0"))
    """Phase 8 - admin-configurable max lifetime for CLI device-flow
    sessions, in days. 0 = unbounded (default, mirrors GitHub CLI /
    gh-style flows). Companies that need finite session lifetime
    (eg. SOC 2 / ISO 27001 customers with mandated session
    expiration) can set a positive value. The /api/auth/cli/refresh
    endpoint rejects refresh attempts where now - issued_at exceeds
    this value, regardless of the underlying refresh-token TTL.
    Spec: specs/ai-governance/sessions/session-ttl.feature.
    """
    mfa_required: Mapped[bool] = mapped_column("mfaRequired", default=False, server_default=text("false"))
    """Every member of this organization can prove a second factor (D06). A
    MEMBERSHIP CONDITION, not a session policy: it is asked when a member
    reaches this organization's data, it holds the ones who cannot prove
    one at an enrollment gate for THIS organization alone, and it ends no
    session, ever. Turning it on signs nobody out.
    Spec: specs/identity/mfa-and-session-shape.feature.
    """
    docs_playground_sessions_enabled: Mapped[bool] = mapped_column("docsPlaygroundSessionsEnabled", default=True, server_default=text("true"))
    """Whether members may connect the documentation playground to their own access with a
    short-lived docs session key (ADR-185, amendment of 2026-11-13). On by default; turned off,
    no key is minted and live ones stop authenticating. The pasted-key path is untouched.
    Spec: modules/api-key/specs/docs-session-keys.feature.
    """
    docs_playground_sessions_disabled_at: Mapped[Optional[datetime]] = mapped_column("docsPlaygroundSessionsDisabledAt")
    """When the docs playground was last turned off, written in the same update as the switch. A
    docs session key minted at or before it never authenticates again, nor any key minted under
    one, so turning it back on revives none of them (ADR-185, amendment of 2026-11-14).
    """
    lockout_after_failed_attempts: Mapped[int] = mapped_column("lockoutAfterFailedAttempts", default=0, server_default=text("0"))
    """How many consecutive failed sign-ins lock an account (GAC-09).
    0 = never lock, which is the default and what every organization that
    has not asked for this keeps. Counts EVERY credential together - a
    password, a one-time code, a backup code, a passkey - because five per
    factor hands an attacker one budget per factor.
    Spec: specs/identity/org-account-lockout.feature.
    """
    lockout_minutes: Mapped[int] = mapped_column("lockoutMinutes", default=30, server_default=text("30"))
    """How long a lock lasts, in minutes. Ignored while the threshold is 0."""
    session_idle_timeout_minutes: Mapped[int] = mapped_column("sessionIdleTimeoutMinutes", default=1440, server_default=text("1440"))
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
    session_max_lifetime_minutes: Mapped[int] = mapped_column("sessionMaxLifetimeMinutes", default=0, server_default=text("0"))
    """The ceiling on a browser session from the moment it was minted, in
    minutes, which no amount of activity extends. 0 = no ceiling, the
    default. Offered beside the idle timeout because a session kept alive
    forever by a polling page is not really bounded.
    """
    signup_data: Mapped[Optional[dict[str, Any]]] = mapped_column("signupData", JSONB(none_as_null=True))
    signed_dpa: Mapped[bool] = mapped_column("signedDPA", default=False, server_default=text("false"))
    elasticsearch_node_url: Mapped[Optional[str]] = mapped_column("elasticsearchNodeUrl")
    elasticsearch_api_key: Mapped[Optional[str]] = mapped_column("elasticsearchApiKey")
    use_custom_elasticsearch: Mapped[bool] = mapped_column("useCustomElasticsearch", default=False, server_default=text("false"))
    s3_endpoint: Mapped[Optional[str]] = mapped_column("s3Endpoint")
    s3_access_key_id: Mapped[Optional[str]] = mapped_column("s3AccessKeyId")
    s3_secret_access_key: Mapped[Optional[str]] = mapped_column("s3SecretAccessKey")
    s3_bucket: Mapped[Optional[str]] = mapped_column("s3Bucket")
    use_custom_s3: Mapped[bool] = mapped_column("useCustomS3", default=False, server_default=text("false"))
    sent_plan_limit_alert: Mapped[Optional[datetime]] = mapped_column("sentPlanLimitAlert")
    sso_domain: Mapped[Optional[str]] = mapped_column("ssoDomain")
    sso_provider: Mapped[Optional[str]] = mapped_column("ssoProvider")
    domain_join: Mapped[str] = mapped_column("domainJoin", default="request", server_default=text("'request'"))
    """How colleagues on a matching domain get in (D12, ADR-117):
    `off` (invisible to them entirely) | `request` (they ask, an admin
    approves) | `auto` (a verified colleague walks straight in).

    `request` is the default for a self-serve organization; `auto` is never
    a default and never inferred - an administrator turns it on, names the
    domain in `joinDomains` while doing it, and the licence gate holds that
    value while letting `request` through.
    """
    joiner_role: Mapped[str] = mapped_column("joinerRole", Enum("ADMIN", "MEMBER", "EXTERNAL", "DEVELOPER", name="OrganizationUserRole"), default="MEMBER", server_default=text("'MEMBER'"))
    """The organisation role handed to a person admitted without an invitation:
    a self-service domain join (`domainJoin`) or an SSO-admitted login.
    `MEMBER` (default, a Full seat) or `DEVELOPER` (ADR-171). Set on the
    organisation settings page only; the join flow never asks, and an
    invitation always names its own role.
    """
    join_domains: Mapped[list[str]] = mapped_column("joinDomains", ARRAY(Text), nullable=True)
    """The domains an administrator NAMED when turning automatic joining on.
    Empty means automatic joining admits nobody whatever `domainJoin` says -
    which is what keeps a setting left on from admitting a domain nobody
    ever agreed to.
    """
    presence_enabled: Mapped[bool] = mapped_column("presenceEnabled", default=True, server_default=text("true"))
    trace_sharing_enabled: Mapped[bool] = mapped_column("traceSharingEnabled", default=True, server_default=text("true"))
    """Org-level trace-sharing kill switch. Off = trace share links cannot be
    created in ANY project of this org, and existing links stop resolving —
    the org toggle is the global switch, the per-project
    `Project.traceSharingEnabled` scopes the kill to a single project (mirrors
    `presenceEnabled`). Effective sharing = org AND project. ADR-039.
    """
    support_contact: Mapped[Optional[str]] = mapped_column("supportContact")
    """Org-configurable contact surfaced to users in CLI "contact your
    admin" messages and the in-app budget-exceeded banner. Free text:
    accepts an email, a URL (internal ticketing system), or any other
    instruction. When null, callers fall back to the first ADMIN
    member's email (`resolveSupportContact` helper).
    """
    primary_intent: Mapped[Optional[str]] = mapped_column("primaryIntent", Enum("AGENT_GOVERNANCE", "LLM_OPS", name="OrganizationIntent"))
    """Signup intent, editable later as the org "Primary use" setting.
    AGENT_GOVERNANCE lands "/" on /me, LLM_OPS on the project home.
    NULL means intent unset (every pre-fork org and any org created
    outside onboarding) — the home resolver keeps legacy behavior.
    ADR-038.
    """
    promo_code: Mapped[Optional[str]] = mapped_column("promoCode")
    stripe_customer_id: Mapped[Optional[str]] = mapped_column("stripeCustomerId")
    currency: Mapped[str] = mapped_column(Enum("USD", "EUR", name="Currency"), default="EUR", server_default=text("'EUR'"))
    pricing_model: Mapped[str] = mapped_column("pricingModel", Enum("TIERED", "SEAT_EVENT", name="PricingModel"), default="TIERED", server_default=text("'TIERED'"))
    license: Mapped[Optional[str]]
    license_expires_at: Mapped[Optional[datetime]] = mapped_column("licenseExpiresAt")
    license_last_validated_at: Mapped[Optional[datetime]] = mapped_column("licenseLastValidatedAt")
    self_hosted_customer: Mapped[bool] = mapped_column("selfHostedCustomer", default=False, server_default=text("false"))
    """A customer that runs LangWatch self-hosted. The hosted usage, budgets and
    invoices of its connected installs attach to this organization (ADR-139).
    """
    connect_services_disabled: Mapped[list[str]] = mapped_column("connectServicesDisabled", ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))
    """LangWatch-hosted services this organization has switched OFF, named the
    way its license names them ("instant_evals"). A service the license is
    entitled to is on unless it is listed here, so a customer who bought
    hosted judging gets it without finding a settings page first, and an
    administrator who switched one off keeps it off across license changes
    (ADR-139, section 2). A license that names no service reaches nothing,
    whatever this column holds.
    """
    instant_evals_enabled_at: Mapped[Optional[datetime]] = mapped_column("instantEvalsEnabledAt")
    """When a member of this organization switched Instant Evals on from the
    search bar, which is the organization's consent to have judged text sent
    to the judge's provider. Null until then: a self-serve organization is
    never opted in on its behalf, and an enterprise one is switched on by an
    operator through the release flag after asking, never here.
    """
    instant_evals_enabled_by_user_id: Mapped[Optional[str]] = mapped_column("instantEvalsEnabledByUserId")
    """The member who switched them on, so the audit trail names a person and
    not just a moment.
    """
    connect_last_sync_at: Mapped[Optional[datetime]] = mapped_column("connectLastSyncAt")
    """When the last license sync succeeded."""
    connect_last_sync_error: Mapped[Optional[str]] = mapped_column("connectLastSyncError")
    """The code of the failure the last sync ended on, cleared on the next
    success. Settings, Connect renders the copy for it.
    """

    __table_args__ = (
        Index("Organization_slug_key", "slug", unique=True),
        Index("Organization_ssoDomain_key", "ssoDomain", unique=True),
        Index("Organization_stripeCustomerId_key", "stripeCustomerId", unique=True),
        Index("Organization_ssoProvider_idx", "ssoProvider"),
    )

    members: Mapped[list["OrganizationUser"]] = relationship(passive_deletes="all", back_populates="organization")
    teams: Mapped[list["Team"]] = relationship(passive_deletes="all", back_populates="organization")
    webhook_endpoints: Mapped[list["WebhookEndpoint"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    organization_invite: Mapped[list["OrganizationInvite"]] = relationship(passive_deletes="all", back_populates="organization")
    custom_roles: Mapped[list["CustomRole"]] = relationship(passive_deletes="all", back_populates="organization")
    notifications: Mapped[list["Notification"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    subscriptions: Mapped[list["Subscription"]] = relationship(passive_deletes="all", back_populates="organization")
    scim_tokens: Mapped[list["ScimToken"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    scim_request_log: Mapped[list["ScimRequestLog"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    prompt_tags: Mapped[list["PromptTag"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    groups: Mapped[list["Group"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    role_bindings: Mapped[list["RoleBinding"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    api_keys: Mapped[list["ApiKey"]] = relationship(passive_deletes="all", back_populates="organization")
    issued_licenses: Mapped[list["IssuedLicense"]] = relationship(passive_deletes="all", back_populates="organization")
    gateway_change_events: Mapped[list["GatewayChangeEvent"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    routing_policies: Mapped[list["RoutingPolicy"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    ingestion_sources: Mapped[list["IngestionSource"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    anomaly_rules: Mapped[list["AnomalyRule"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    anomaly_alerts: Mapped[list["AnomalyAlert"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    ai_tool_entries: Mapped[list["AiToolEntry"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    ingestion_templates: Mapped[list["IngestionTemplate"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    github_installations: Mapped[list["GithubInstallation"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    github_pull_requests: Mapped[list["GithubPullRequest"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    github_branch_pr_checks: Mapped[list["GithubBranchPullRequestCheck"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="organization")
    connected_billing: Mapped[Optional["ConnectedBillingAccount"]] = relationship(passive_deletes="all", back_populates="organization")

class IssuedLicense(Base):
    """The registry of licenses LangWatch has issued (ADR-139). It lives on
    LangWatch Cloud; a self-hosted install has the table and no rows.

    The row never holds the token an install presents. `tokenHash` is the
    SHA-256 of that token, so read access to this table is not credential
    access. Status is derived: revoked, superseded, past `expiresAt`, else active.
    """

    __tablename__ = "IssuedLicense"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    license_id: Mapped[str] = mapped_column("licenseId")
    """The `licenseId` inside the signed license (`lic-<uuid>`)."""
    token_hash: Mapped[str] = mapped_column("tokenHash")
    """SHA-256 of the `lwl_` token. Looked up on every credential resolution."""
    organization_id: Mapped[Optional[str]] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="RESTRICT", onupdate="CASCADE", name="IssuedLicense_organizationId_fkey"))
    """The customer organization on LangWatch Cloud. Null for a license that is
    recorded but not linked to a customer yet, such as a self-serve purchase:
    it resolves to nothing until an operator links it.
    """
    organization_name: Mapped[str] = mapped_column("organizationName")
    """As signed into the license."""
    email: Mapped[str]
    plan_type: Mapped[str] = mapped_column("planType")
    max_members: Mapped[int] = mapped_column("maxMembers")
    max_members_lite: Mapped[Optional[int]] = mapped_column("maxMembersLite")
    """The lite seats sold: null when not stated or when unlimited. Never a stand-in number."""
    lite_seats_unlimited: Mapped[bool] = mapped_column("liteSeatsUnlimited", default=False, server_default=text("false"))
    """True when the license sells unlimited lite seats; maxMembersLite is then null."""
    max_members_developer: Mapped[Optional[int]] = mapped_column("maxMembersDeveloper")
    """The Developer seats sold: null when not stated, which sells none."""
    issued_at: Mapped[datetime] = mapped_column("issuedAt")
    expires_at: Mapped[datetime] = mapped_column("expiresAt")
    source: Mapped[str] = mapped_column(Enum("BACKOFFICE", "PURCHASE", "SCRIPT", "LEGACY_IMPORT", name="IssuedLicenseSource"))
    issued_by_id: Mapped[Optional[str]] = mapped_column("issuedById")
    """The operator who issued or registered it. Null for the purchase flow and the script."""
    revoked_at: Mapped[Optional[datetime]] = mapped_column("revokedAt")
    revoked_by_id: Mapped[Optional[str]] = mapped_column("revokedById")
    revoked_reason: Mapped[Optional[str]] = mapped_column("revokedReason")
    superseded_at: Mapped[Optional[datetime]] = mapped_column("supersededAt")
    """Set when the install first presents the license that replaced this one."""
    replaces_id: Mapped[Optional[str]] = mapped_column("replacesId")
    """The id of the IssuedLicense this one replaces, when it was reissued. A
    plain column ref (relationMode="prisma" cannot declare a self-relation on
    Postgres); the registry service validates it. Unique, so a license is
    replaced by one license at most.
    """
    pending_delivery_license: Mapped[Optional[str]] = mapped_column("pendingDeliveryLicense")
    """A reissued license waiting to be delivered over sync, encrypted. Erased
    when the install first presents it. The only license text the registry holds.
    """
    services: Mapped[list[str]] = mapped_column(ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))
    """Hosted services this license may call, such as `instant_evals`."""
    seat_rate_cents: Mapped[Optional[int]] = mapped_column("seatRateCents")
    """The annual rate of one seat, in cents of `seatCurrency`. Needed to prorate added seats."""
    seat_currency: Mapped[Optional[str]] = mapped_column("seatCurrency", Enum("USD", "EUR", name="Currency"))
    commit_usd_cents: Mapped[int] = mapped_column("commitUsdCents", default=0, server_default=text("0"))
    """The prepaid usage commit for the term, in USD cents."""
    overage_enabled: Mapped[bool] = mapped_column("overageEnabled", default=False, server_default=text("false"))
    """Whether usage may continue past the commit, up to `overageMaxUsdCents`."""
    overage_max_usd_cents: Mapped[Optional[int]] = mapped_column("overageMaxUsdCents")
    instance_id: Mapped[Optional[str]] = mapped_column("instanceId")
    """The install this license is bound to. Set on first use, cleared by an operator."""
    instance_bound_at: Mapped[Optional[datetime]] = mapped_column("instanceBoundAt")
    last_sync_at: Mapped[Optional[datetime]] = mapped_column("lastSyncAt")
    """The last license sync this install reported: when it arrived, the
    LangWatch version that sent it, and the seats in use at that moment.
    Null until the install syncs for the first time.
    """
    last_sync_version: Mapped[Optional[str]] = mapped_column("lastSyncVersion")
    reported_members: Mapped[Optional[int]] = mapped_column("reportedMembers")
    reported_members_lite: Mapped[Optional[int]] = mapped_column("reportedMembersLite")
    virtual_key_id: Mapped[Optional[str]] = mapped_column("virtualKeyId")
    """The managed gateway key (purpose CONNECT) that hosted-service traffic of
    this license runs under. One per license, so spend rows name the install.
    Created when the license first resolves. A plain column, not a relation.
    """
    seats_raised_from: Mapped[Optional[int]] = mapped_column("seatsRaisedFrom")
    """The seats of the license this one replaced, set only when a seat change
    raised a linked license. Billing reads it through licensing to invoice
    the added seats, so the fact commits with the license or not at all.
    """
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("IssuedLicense_licenseId_key", "licenseId", unique=True),
        Index("IssuedLicense_tokenHash_key", "tokenHash", unique=True),
        Index("IssuedLicense_replacesId_key", "replacesId", unique=True),
        Index("IssuedLicense_virtualKeyId_key", "virtualKeyId", unique=True),
        Index("IssuedLicense_organizationId_idx", "organizationId"),
        Index("IssuedLicense_expiresAt_idx", "expiresAt"),
        Index("IssuedLicense_seatsRaisedFrom_idx", "seatsRaisedFrom"),
    )

    organization: Mapped[Optional["Organization"]] = relationship(back_populates="issued_licenses")

class OrganizationLicense(Base):
    """The licence an organization activated and its dates, owned by licensing
    (round 37 D6). A row whose key is null is a cleared licence. An organization
    with no row is read from its Organization columns until the copy step runs.
    """

    __tablename__ = "OrganizationLicense"

    organization_id: Mapped[str] = mapped_column("organizationId", primary_key=True)
    license_key: Mapped[Optional[str]] = mapped_column("licenseKey")
    expires_at: Mapped[Optional[datetime]] = mapped_column("expiresAt")
    validated_at: Mapped[Optional[datetime]] = mapped_column("validatedAt")
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

class ActivationCode(Base):
    """A short code a fresh install pastes instead of a license blob (ADR-139,
    section 5). Redeeming it mints the license the code describes and returns
    it signed, which is what identifies a new customer without anybody emailing
    a license key around.

    The row holds a hash of the code, never the code. Every lookup hashes what
    the caller presented first, so the stored value is a step away from anything
    that could be presented: read access to this table is not credential access.
    """

    __tablename__ = "ActivationCode"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    code_hash: Mapped[str] = mapped_column("codeHash")
    """SHA-256 of the normalised code. Looked up on every redemption."""
    code_hint: Mapped[str] = mapped_column("codeHint")
    """The last four characters, so an operator tells two codes apart in a list
    without the list holding either one.
    """
    organization_id: Mapped[str] = mapped_column("organizationId")
    """The customer this code belongs to, and the terms it mints."""
    organization_name: Mapped[str] = mapped_column("organizationName")
    email: Mapped[str]
    plan_type: Mapped[str] = mapped_column("planType")
    max_members: Mapped[int] = mapped_column("maxMembers")
    max_members_lite: Mapped[Optional[int]] = mapped_column("maxMembersLite", default=0, server_default=text("0"))
    """The lite seats the minted license carries: null when unlimited. Never a stand-in number."""
    lite_seats_unlimited: Mapped[bool] = mapped_column("liteSeatsUnlimited", default=False, server_default=text("false"))
    """True when the minted license carries unlimited lite seats; maxMembersLite is then null."""
    max_members_developer: Mapped[Optional[int]] = mapped_column("maxMembersDeveloper")
    """The Developer seats the minted license carries: null when not stated, which sells none."""
    license_term_days: Mapped[int] = mapped_column("licenseTermDays")
    """How long the license this code mints runs for."""
    services: Mapped[list[str]] = mapped_column(ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))
    """Hosted services the minted license may call."""
    expires_at: Mapped[datetime] = mapped_column("expiresAt")
    """When the code stops working, which is not the term of what it mints."""
    reusable: Mapped[bool] = mapped_column(default=False, server_default=text("false"))
    """A code a customer may redeem on more than one install, for a rollout."""
    redeemed_at: Mapped[Optional[datetime]] = mapped_column("redeemedAt")
    """Set by the redemption that won, which is what makes a single-use code
    single use: the write itself carries the condition.
    """
    redeemed_by_instance_id: Mapped[Optional[str]] = mapped_column("redeemedByInstanceId")
    issued_license_id: Mapped[Optional[str]] = mapped_column("issuedLicenseId")
    """The license the winning redemption minted."""
    redemption_count: Mapped[int] = mapped_column("redemptionCount", default=0, server_default=text("0"))
    """How many installs have redeemed it, which only moves on a reusable code."""
    revoked_at: Mapped[Optional[datetime]] = mapped_column("revokedAt")
    revoked_by_id: Mapped[Optional[str]] = mapped_column("revokedById")
    created_by_id: Mapped[Optional[str]] = mapped_column("createdById")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("ActivationCode_codeHash_key", "codeHash", unique=True),
        Index("ActivationCode_organizationId_idx", "organizationId"),
        Index("ActivationCode_expiresAt_idx", "expiresAt"),
    )

class ConnectedBillingAccount(Base):
    """How a connected customer is invoiced (ADR-139, section 7): one payment
    provider customer, one quarterly usage subscription in USD, and the
    commercial terms the annual invoice was raised from.

    One row per organization. Every write through it is addressed by a stored
    provider id, so running onboarding or a renewal twice changes nothing.
    """

    __tablename__ = "ConnectedBillingAccount"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="RESTRICT", onupdate="CASCADE", name="ConnectedBillingAccount_organizationId_fkey"))
    stripe_customer_id: Mapped[str] = mapped_column("stripeCustomerId")
    """The customer at the payment provider. Invoiced, never charged."""
    usage_subscription_id: Mapped[Optional[str]] = mapped_column("usageSubscriptionId")
    """The one metered subscription, and the item hosted usage is reported on.
    Null between the customer being created and the subscription existing,
    which is where a half-finished onboarding resumes from.
    """
    usage_subscription_item_id: Mapped[Optional[str]] = mapped_column("usageSubscriptionItemId")
    term_starts_at: Mapped[datetime] = mapped_column("termStartsAt")
    term_ends_at: Mapped[datetime] = mapped_column("termEndsAt")
    commit_usd_cents: Mapped[int] = mapped_column("commitUsdCents", default=0, server_default=text("0"))
    """Prepaid usage for the term, in USD cents. The credit grants add up to it."""
    seat_currency: Mapped[str] = mapped_column("seatCurrency", Enum("USD", "EUR", name="Currency"))
    """The seat contract, which may be in another currency than hosted usage."""
    seat_rate_cents: Mapped[int] = mapped_column("seatRateCents")
    seats: Mapped[int]
    bank_transfer_type: Mapped[Optional[str]] = mapped_column("bankTransferType")
    """`us_bank_transfer` or `eu_bank_transfer` when the customer pays into a
    virtual bank account at the provider. Null when it wires to LangWatch's
    own account and finance marks the invoice paid.
    """
    bank_transfer_country: Mapped[Optional[str]] = mapped_column("bankTransferCountry")
    billing_email: Mapped[str] = mapped_column("billingEmail")
    pending_renewal: Mapped[Optional[dict[str, Any]]] = mapped_column("pendingRenewal", JSONB(none_as_null=True))
    """A renewal whose credit is waiting for the last usage invoice of the old
    term to be finalized: the new term's end and commit, and who asked.
    """
    created_by_id: Mapped[Optional[str]] = mapped_column("createdById")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("ConnectedBillingAccount_organizationId_key", "organizationId", unique=True),
        Index("ConnectedBillingAccount_stripeCustomerId_key", "stripeCustomerId", unique=True),
    )

    organization: Mapped["Organization"] = relationship(back_populates="connected_billing")

class ConnectedCreditGrant(Base):
    """A paid credit at the payment provider, which only metered usage draws down.
    Kept here so a grant is created once per reason and the backoffice can show
    what the customer holds without asking the provider.
    """

    __tablename__ = "ConnectedCreditGrant"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    account_id: Mapped[str] = mapped_column("accountId")
    """The `ConnectedBillingAccount` row. A plain column ref, like
    `IssuedLicense.replacesId`: relationMode="prisma" carries no foreign keys.
    """
    stripe_credit_grant_id: Mapped[str] = mapped_column("stripeCreditGrantId")
    """The grant at the payment provider. Unique, so a retry records one row."""
    amount_usd_cents: Mapped[int] = mapped_column("amountUsdCents")
    kind: Mapped[str] = mapped_column(Enum("commit", "added", "renewal", name="ConnectedCreditGrantKind"))
    term_ends_at: Mapped[datetime] = mapped_column("termEndsAt")
    """The term this credit was granted for, and when the provider lets it lapse:
    after the last invoice of that term has been finalized, never on the day
    the term ends.
    """
    expires_at: Mapped[datetime] = mapped_column("expiresAt")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("ConnectedCreditGrant_stripeCreditGrantId_key", "stripeCreditGrantId", unique=True),
        Index("ConnectedCreditGrant_accountId_idx", "accountId"),
    )

class ConnectedInvoice(Base):
    """An invoice LangWatch raised for a connected customer, as the backoffice
    reads it.
    """

    __tablename__ = "ConnectedInvoice"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    account_id: Mapped[str] = mapped_column("accountId")
    """The `ConnectedBillingAccount` row, as a plain column ref."""
    stripe_invoice_id: Mapped[str] = mapped_column("stripeInvoiceId")
    kind: Mapped[str] = mapped_column(Enum("annual", "seat_change", "usage", name="ConnectedInvoiceKind"))
    currency: Mapped[str] = mapped_column(Enum("USD", "EUR", name="Currency"))
    amount_cents: Mapped[int] = mapped_column("amountCents")
    status: Mapped[str]
    """The provider's own status: `draft`, `open`, `paid`, `void` or `uncollectible`."""
    paid_out_of_band_at: Mapped[Optional[datetime]] = mapped_column("paidOutOfBandAt")
    term_starts_at: Mapped[Optional[datetime]] = mapped_column("termStartsAt")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("ConnectedInvoice_stripeInvoiceId_key", "stripeInvoiceId", unique=True),
        Index("ConnectedInvoice_accountId_idx", "accountId"),
    )

class ScimRequestLog(Base):
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

    __tablename__ = "ScimRequestLog"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="ScimRequestLog_organizationId_fkey"))
    connection_id: Mapped[Optional[str]] = mapped_column("connectionId")
    """Null when the refusal resolved an organization and stopped before the
    connection — a lapsed plan is answered off the token alone.
    """
    method: Mapped[str]
    resource: Mapped[str]
    """The SCIM resource asked for ("Users", "Users/:id", "Groups"), never the
    raw path: a path carries ids and query strings, and this is read by
    people rather than matched by machines.
    """
    status: Mapped[int]
    reason: Mapped[Optional[str]]
    """A stable slug, only on a refusal, so a reader branches on this and never
    on the sentence beside it.
    """
    detail: Mapped[Optional[str]]
    """Our own short sentence, customer-safe."""
    occurred_at: Mapped[datetime] = mapped_column("occurredAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("ScimRequestLog_organizationId_occurredAt_idx", "organizationId", "occurredAt"),
        Index("ScimRequestLog_connectionId_occurredAt_idx", "connectionId", "occurredAt"),
        Index("ScimRequestLog_occurredAt_idx", "occurredAt"),
    )

    organization: Mapped["Organization"] = relationship(back_populates="scim_request_log")

class ScimToken(Base):
    __tablename__ = "ScimToken"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="ScimToken_organizationId_fkey"))
    connection_id: Mapped[Optional[str]] = mapped_column("connectionId")
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
    hashed_token: Mapped[str] = mapped_column("hashedToken")
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
    hash_scheme: Mapped[str] = mapped_column("hashScheme", default="sha256", server_default=text("'sha256'"))
    """How `hashedToken` was derived, so the two schemes can coexist while
    tokens minted under the old one are still in use.

    `sha256` is a bare digest, which is correct for 32 random bytes and wrong
    for a value a person chose — a database dump of those is a wordlist away
    from live credentials. `hmac-sha256` is keyed on the deployment's own
    secret, so a dump on its own is inert. Every new row is the latter.
    """
    description: Mapped[Optional[str]]
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    last_used_at: Mapped[Optional[datetime]] = mapped_column("lastUsedAt")

    __table_args__ = (
        Index("ScimToken_hashedToken_key", "hashedToken", unique=True),
        Index("ScimToken_organizationId_idx", "organizationId"),
        Index("ScimToken_connectionId_idx", "connectionId"),
    )

    organization: Mapped["Organization"] = relationship(back_populates="scim_tokens")

class ScimExternalId(Base):
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

    __tablename__ = "ScimExternalId"

    organization_id: Mapped[str] = mapped_column("organizationId")
    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    connection_id: Mapped[str] = mapped_column("connectionId")
    external_id: Mapped[str] = mapped_column("externalId")
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="ScimExternalId_userId_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("ScimExternalId_organizationId_userId_idx", "organizationId", "userId"),
        Index("ScimExternalId_userId_idx", "userId"),
        Index("ScimExternalId_connectionId_idx", "connectionId"),
        Index("ScimExternalId_connectionId_externalId_key", "connectionId", "externalId", unique=True),
    )

    user: Mapped["User"] = relationship(back_populates="scim_external_ids")

class ScimUserResource(Base):
    """Directory profile and lifecycle are tenant facts, independent of the shared account."""

    __tablename__ = "ScimUserResource"

    organization_id: Mapped[str] = mapped_column("organizationId", primary_key=True)
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="ScimUserResource_userId_fkey"), primary_key=True)
    user_name: Mapped[str] = mapped_column("userName")
    name: Mapped[Optional[str]]
    active: Mapped[bool] = mapped_column(default=True, server_default=text("true"))
    deleted_at: Mapped[Optional[datetime]] = mapped_column("deletedAt")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("ScimUserResource_organizationId_userName_idx", "organizationId", "userName"),
        Index("ScimUserResource_userId_idx", "userId"),
    )

    user: Mapped["User"] = relationship(back_populates="scim_user_resources")

class ScimDirectoryUser(Base):
    """A connection's managed users, including resources without an externalId.
    Scalar keys only; readers exclude users that no longer exist.
    """

    __tablename__ = "ScimDirectoryUser"

    organization_id: Mapped[str] = mapped_column("organizationId")
    connection_id: Mapped[str] = mapped_column("connectionId", primary_key=True)
    user_id: Mapped[str] = mapped_column("userId", primary_key=True)

    __table_args__ = (
        Index("ScimDirectoryUser_organizationId_userId_idx", "organizationId", "userId"),
    )

class ScimSyncState(Base):
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

    __tablename__ = "ScimSyncState"

    id: Mapped[str] = mapped_column(primary_key=True)
    """The aggregate: the connection's id, which IS its sync's id."""
    connection_id: Mapped[str] = mapped_column("connectionId")
    organization_id: Mapped[str] = mapped_column("organizationId")
    state: Mapped[str]
    """TOKEN_ISSUED | SYNCING | ERROR | REVOKED"""
    last_pushed_at: Mapped[Optional[datetime]] = mapped_column("lastPushedAt")
    last_failure: Mapped[Optional[dict[str, Any]]] = mapped_column("lastFailure", JSONB(none_as_null=True))
    """The failure standing right now: { op, errorCode, attempts, retiredAtMs,
    userId, occurredAtMs }. A reason CODE and ids — never a token, never a
    secret, never an internal hostname.
    """
    dead_letters: Mapped[dict[str, Any]] = mapped_column("deadLetters", JSONB, default=lambda: [], server_default=text("'[]'"))
    """Every failure retired without ever being applied, newest last. The dead
    letters an administrator has to see; retiring never drops one.
    """
    revoked_cause: Mapped[Optional[str]] = mapped_column("revokedCause")
    """revoke | teardown, once the sync has ended."""
    occurred_at: Mapped[datetime] = mapped_column("occurredAt")
    last_event_id: Mapped[str] = mapped_column("lastEventId")
    accepted_at: Mapped[datetime] = mapped_column("acceptedAt")
    projection_version: Mapped[str] = mapped_column("projectionVersion")
    created_at: Mapped[datetime] = mapped_column("createdAt")
    """Event-derived, not database-managed: `@default(now())` / `@updatedAt`
    would make a replayed row differ from the row it rebuilds, which is
    exactly the parity this projection has to keep.
    """
    updated_at: Mapped[datetime] = mapped_column("updatedAt")

    __table_args__ = (
        Index("ScimSyncState_organizationId_idx", "organizationId"),
        Index("ScimSyncState_state_idx", "state"),
        Index("ScimSyncState_connectionId_idx", "connectionId"),
    )

class ScimSsoConnectionView(Base):
    """SCIM's own copy of an organization's SSO connections: a peer fold over
    identity's connection facts (ARCHITECTURE §9), folded with identity-contract's
    reducer, one row per connection. SCIM reads this instead of asking identity.
    Rebuilt from identity's event log by a projection replay, so no relation:
    a foreign key would make a replay depend on state the log does not carry.
    """

    __tablename__ = "ScimSsoConnectionView"

    id: Mapped[str] = mapped_column(primary_key=True)
    """The aggregate: identity's connection id."""
    organization_id: Mapped[str] = mapped_column("organizationId")
    folded: Mapped[dict[str, Any]] = mapped_column(JSONB)
    """The connection as the reducer folded it (SsoConnectionState + LastEventOccurredAt)."""
    applied_event_ids: Mapped[list[str]] = mapped_column("appliedEventIds", ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))
    """The event ids of the last delivery folded in, so a redelivery folds once."""
    projection_version: Mapped[str] = mapped_column("projectionVersion")
    """A row of another version reads as undecodable and is re-folded from the log."""
    created_at: Mapped[datetime] = mapped_column("createdAt")
    """Event-derived, so a replayed row equals the row it rebuilds."""
    updated_at: Mapped[datetime] = mapped_column("updatedAt")

    __table_args__ = (
        Index("ScimSsoConnectionView_organizationId_createdAt_idx", "organizationId", "createdAt"),
    )

class OrganizationSeatPeak(Base):
    """Organization's monthly seat peak, folded from its seat changes (modules/organization, ADR-002).
    A read model, one row per organization: rebuildable by replaying the event log, so no relation.
    Epochs in milliseconds.
    Organization's seat-peak counts (ADR-002): one row per organization once counting opened.
    """

    __tablename__ = "OrganizationSeatPeak"

    organization_id: Mapped[str] = mapped_column("organizationId", primary_key=True)
    """The organization, which is also the aggregate id."""
    opened_at: Mapped[float] = mapped_column("openedAt")
    """When the opening snapshot was taken, epoch milliseconds."""
    full_seats: Mapped[int] = mapped_column("fullSeats")
    lite_seats: Mapped[int] = mapped_column("liteSeats")
    developer_seats: Mapped[int] = mapped_column("developerSeats")
    latest_change_at: Mapped[float] = mapped_column("latestChangeAt")
    """The latest change time applied to any holder; invitations expire up to it."""
    mismatches: Mapped[int]
    """Changes whose `from` disagreed with the seat the fold held."""

class OrganizationSeatPeakHolder(Base):
    """One member's or pending invitation's seat as the seat peak last applied it (ADR-002)."""

    __tablename__ = "OrganizationSeatPeakHolder"

    organization_id: Mapped[str] = mapped_column("organizationId", primary_key=True)
    holder_kind: Mapped[str] = mapped_column("holderKind", primary_key=True)
    """`member`, `invitation` or `hold` (a seat held after somebody left)."""
    holder_id: Mapped[str] = mapped_column("holderId", primary_key=True)
    """The member's or the hold's userId, or the invitation's id."""
    seat: Mapped[str]
    """`full`, `lite`, `developer` or `none`."""
    last_changed_at: Mapped[float] = mapped_column("lastChangedAt")
    last_change_id: Mapped[Optional[str]] = mapped_column("lastChangeId")
    """Absent for a seat the opening snapshot recorded."""
    expires_at: Mapped[Optional[float]] = mapped_column("expiresAt")
    """An invitation's expiry; null for a member or an invitation that never expires."""
    expired_by_fold: Mapped[bool] = mapped_column("expiredByFold", default=False, server_default=text("false"))
    """The fold gave the invitation's seat back at `expiresAt`."""
    suppressed_by: Mapped[Optional[str]] = mapped_column("suppressedBy")
    """A hold whose seat a live full invitation to its address took: that invitation's id."""
    restored_by_fold: Mapped[bool] = mapped_column("restoredByFold", default=False, server_default=text("false"))
    """The fold gave a suppressed hold its seat back when that invitation expired."""

    __table_args__ = (
        Index("OrganizationSeatPeakHolder_organizationId_holderKind_seat_e_idx", "organizationId", "holderKind", "seat", "expiresAt"),
        Index("OrganizationSeatPeakHolder_organizationId_suppressedBy_idx", "organizationId", "suppressedBy"),
    )

class OrganizationSeatPeakMonth(Base):
    """One calendar month's (UTC) seat peaks for an organization (ADR-002)."""

    __tablename__ = "OrganizationSeatPeakMonth"

    organization_id: Mapped[str] = mapped_column("organizationId", primary_key=True)
    month: Mapped[str] = mapped_column(primary_key=True)
    """`YYYY-MM`."""
    full_peak: Mapped[int] = mapped_column("fullPeak")
    full_peak_at: Mapped[float] = mapped_column("fullPeakAt")
    lite_peak: Mapped[int] = mapped_column("litePeak")
    lite_peak_at: Mapped[float] = mapped_column("litePeakAt")
    full_at_end: Mapped[int] = mapped_column("fullAtEnd")
    lite_at_end: Mapped[int] = mapped_column("liteAtEnd")

class Project(Base):
    __tablename__ = "Project"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    name: Mapped[str]
    slug: Mapped[str]
    api_key: Mapped[str] = mapped_column("apiKey")
    legacy_api_key_last_used_at: Mapped[Optional[datetime]] = mapped_column("legacyApiKeyLastUsedAt")
    """When the legacy apiKey last authenticated a request, written at most once a
    minute per project and process. Spec: modules/api-key/specs/request-credential-type.feature
    """
    lwql_key: Mapped[str] = mapped_column("lwqlKey", server_default=text("(gen_random_uuid())::text"))
    """Random per-project secret whose SHA-256 digest is this project's LangWatchQL
    SQL tenant capability in ClickHouse. Deliberately NOT the API key: the
    credential a caller authenticates with and the capability the database
    resolves a tenant from stay decoupled, so rotating one never invalidates
    the other, and a leak of the ClickHouse key map's input discloses nothing
    a caller could authenticate with. Generated by the database on insert.
    """
    team_id: Mapped[str] = mapped_column("teamId", ForeignKey("Team.id", ondelete="RESTRICT", onupdate="CASCADE", name="Project_teamId_fkey"))
    language: Mapped[str]
    framework: Mapped[str]
    kind: Mapped[str] = mapped_column(default="application", server_default=text("'application'"))
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
    aggregate_rule: Mapped[Optional[dict[str, Any]]] = mapped_column("aggregateRule", JSONB(none_as_null=True))
    """The scope rule an aggregate project materialises into grants. Null on
    every other kind; decided in ADR-177. Shapes: { kind: "all-personal" } |
    { kind: "personal-by-department", departmentId } |
    { kind: "explicit", projectIds: string[] }.
    """
    first_message: Mapped[bool] = mapped_column("firstMessage", default=False, server_default=text("false"))
    integrated: Mapped[bool] = mapped_column(default=False, server_default=text("false"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    user_link_template: Mapped[Optional[str]] = mapped_column("userLinkTemplate")
    trace_sharing_enabled: Mapped[bool] = mapped_column("traceSharingEnabled", default=True, server_default=text("true"))
    presence_enabled: Mapped[bool] = mapped_column("presenceEnabled", default=True, server_default=text("true"))
    s3_endpoint: Mapped[Optional[str]] = mapped_column("s3Endpoint")
    s3_access_key_id: Mapped[Optional[str]] = mapped_column("s3AccessKeyId")
    s3_secret_access_key: Mapped[Optional[str]] = mapped_column("s3SecretAccessKey")
    s3_bucket: Mapped[Optional[str]] = mapped_column("s3Bucket")
    archived_at: Mapped[Optional[datetime]] = mapped_column("archivedAt")
    is_personal: Mapped[bool] = mapped_column("isPersonal", default=False, server_default=text("false"))
    owner_user_id: Mapped[Optional[str]] = mapped_column("ownerUserId", String(255), ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="Project_ownerUserId_fkey"))
    personal_features: Mapped[dict[str, Any]] = mapped_column("personalFeatures", JSONB, default=lambda: {}, server_default=text("'{}'"))
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
    department_id: Mapped[Optional[str]] = mapped_column("departmentId")
    """Department assignment for spend attribution (the project / agent
    lens, used when a trace has no principal user). Accounting-only.
    """
    langy_egress_allowlist: Mapped[Optional[dict[str, Any]]] = mapped_column("langyEgressAllowlist", JSONB(none_as_null=True))
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
    last_coding_agent_session_at: Mapped[Optional[datetime]] = mapped_column("lastCodingAgentSessionAt")
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
    last_coding_agent_pull_request_at: Mapped[Optional[datetime]] = mapped_column("lastCodingAgentPullRequestAt")

    __table_args__ = (
        Index("Project_slug_key", "slug", unique=True),
        Index("Project_apiKey_key", "apiKey", unique=True),
        Index("Project_lwqlKey_key", "lwqlKey", unique=True),
        Index("Project_teamId_idx", "teamId"),
        Index("Project_isPersonal_idx", "isPersonal"),
        Index("Project_ownerUserId_idx", "ownerUserId"),
        Index("Project_teamId_kind_idx", "teamId", "kind"),
        Index("Project_departmentId_idx", "departmentId"),
    )

    team: Mapped["Team"] = relationship(back_populates="projects")
    owner_user: Mapped[Optional["User"]] = relationship(back_populates="personal_projects")
    checks: Mapped[list["Monitor"]] = relationship(passive_deletes="all", back_populates="project")
    costs: Mapped[list["Cost"]] = relationship(passive_deletes="all", back_populates="project")
    topics: Mapped[list["Topic"]] = relationship(passive_deletes="all", back_populates="project")
    datasets: Mapped[list["Dataset"]] = relationship(passive_deletes="all", back_populates="project")
    dataset_records: Mapped[list["DatasetRecord"]] = relationship(passive_deletes="all", back_populates="project")
    custom_graphs: Mapped[list["CustomGraph"]] = relationship(passive_deletes="all", back_populates="project")
    dashboards: Mapped[list["Dashboard"]] = relationship(passive_deletes="all", back_populates="project")
    batch_evaluations: Mapped[list["BatchEvaluation"]] = relationship(passive_deletes="all", back_populates="project")
    triggers: Mapped[list["Trigger"]] = relationship(passive_deletes="all", back_populates="project")
    experiments: Mapped[list["Experiment"]] = relationship(passive_deletes="all", back_populates="project")
    annotations: Mapped[list["Annotation"]] = relationship(passive_deletes="all", back_populates="project")
    project_secrets: Mapped[list["ProjectSecret"]] = relationship(passive_deletes="all", back_populates="project")
    trigger_sent: Mapped[list["TriggerSent"]] = relationship(passive_deletes="all", back_populates="project")
    webhook_deliveries: Mapped[list["WebhookEndpointDelivery"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="project")
    annotation_scores: Mapped[list["AnnotationScore"]] = relationship(passive_deletes="all", back_populates="project")
    share_links: Mapped[list["ShareLink"]] = relationship(passive_deletes="all", back_populates="project")
    workflows: Mapped[list["Workflow"]] = relationship(passive_deletes="all", back_populates="project")
    workflow_version: Mapped[list["WorkflowVersion"]] = relationship(passive_deletes="all", back_populates="project")
    annotation_queue: Mapped[list["AnnotationQueue"]] = relationship(passive_deletes="all", back_populates="project")
    annotation_queue_item: Mapped[list["AnnotationQueueItem"]] = relationship(passive_deletes="all", back_populates="project")
    llm_prompt_configs: Mapped[list["LlmPromptConfig"]] = relationship(passive_deletes="all", back_populates="project")
    analytics: Mapped[list["Analytics"]] = relationship(passive_deletes="all", back_populates="project")
    notifications: Mapped[list["Notification"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="project")
    agents: Mapped[list["Agent"]] = relationship(passive_deletes="all", back_populates="project")
    evaluators: Mapped[list["Evaluator"]] = relationship(passive_deletes="all", back_populates="project")
    scenarios: Mapped[list["Scenario"]] = relationship(passive_deletes="all", back_populates="project")
    suite_configurations: Mapped[list["SimulationSuite"]] = relationship(passive_deletes="all", back_populates="project")
    saved_views: Mapped[list["SavedView"]] = relationship(passive_deletes="all", back_populates="project")
    gateway_guardrails: Mapped[list["GatewayGuardrail"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="project")
    gateway_change_events: Mapped[list["GatewayChangeEvent"]] = relationship(back_populates="project")
    pinned_traces: Mapped[list["PinnedTrace"]] = relationship(passive_deletes="all", back_populates="project")
    trace_edit_overlays: Mapped[list["TraceEditOverlay"]] = relationship(passive_deletes="all", back_populates="project")
    email_suppressions: Mapped[list["EmailSuppression"]] = relationship(passive_deletes="all", back_populates="project")

class Department(Base):
    """A department is a single-valued accounting dimension, separate from
    RBAC, for rolling up AI spend by org unit (Engineering, Marketing,
    ...). Assigned to OrganizationUser, Team, and Project via their
    `departmentId` column. Pure accounting: never grants or restricts
    access. See specs/ai-gateway/governance/departments.feature.
    """

    __tablename__ = "Department"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId")
    name: Mapped[str]
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    archived_at: Mapped[Optional[datetime]] = mapped_column("archivedAt")

    __table_args__ = (
        Index("Department_organizationId_idx", "organizationId"),
    )

class DepartmentMembershipHistory(Base):
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

    __tablename__ = "DepartmentMembershipHistory"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId")
    user_id: Mapped[str] = mapped_column("userId")
    department_id: Mapped[str] = mapped_column("departmentId")
    valid_from: Mapped[datetime] = mapped_column("validFrom")
    valid_to: Mapped[Optional[datetime]] = mapped_column("validTo")
    """Null means the link is open. Reassignment and clearing close it."""
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("DepartmentMembershipHistory_organizationId_userId_validFrom_idx", "organizationId", "userId", "validFrom"),
    )

class PlatformToolPolicy(Base):
    """Per-(organization, tool) override for which `langwatch <tool>` paths the
    CLI may use: `allowVk` gates the gateway / virtual-key path (Path A) and
    `allowOtelDirect` gates the direct-OTLP ingestion path (Path B). A missing
    row resolves to the hardcoded defaults (both true for claude/codex/gemini/
    opencode; cursor allows the gateway path only), so the policy is purely
    additive: disabling a path is an explicit admin choice. The CLI caches the
    resolved map at login and gates path selection on it.
    See specs/ai-gateway/governance/cli-tool-mode-policy.feature.
    """

    __tablename__ = "PlatformToolPolicy"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId")
    tool_slug: Mapped[str] = mapped_column("toolSlug")
    allow_vk: Mapped[bool] = mapped_column("allowVk", default=True, server_default=text("true"))
    allow_otel_direct: Mapped[bool] = mapped_column("allowOtelDirect", default=True, server_default=text("true"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("PlatformToolPolicy_organizationId_idx", "organizationId"),
        Index("PlatformToolPolicy_organizationId_toolSlug_key", "organizationId", "toolSlug", unique=True),
    )

class OrganizationInvite(Base):
    __tablename__ = "OrganizationInvite"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    email: Mapped[str]
    invite_code: Mapped[str] = mapped_column("inviteCode")
    expiration: Mapped[Optional[datetime]]
    status: Mapped[str] = mapped_column(Enum("PENDING", "ACCEPTED", "WAITING_APPROVAL", "PAYMENT_PENDING", "REVOKED", name="INVITE_STATUS"), default="PENDING", server_default=text("'PENDING'"))
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="RESTRICT", onupdate="CASCADE", name="OrganizationInvite_organizationId_fkey"))
    team_ids: Mapped[str] = mapped_column("teamIds")
    team_assignments: Mapped[Optional[dict[str, Any]]] = mapped_column("teamAssignments", JSONB(none_as_null=True))
    role: Mapped[str] = mapped_column(Enum("ADMIN", "MEMBER", "EXTERNAL", "DEVELOPER", name="OrganizationUserRole"))
    requested_by: Mapped[Optional[str]] = mapped_column("requestedBy", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="OrganizationInvite_requestedBy_fkey"))
    subscription_id: Mapped[Optional[str]] = mapped_column("subscriptionId")
    accepted_by_user_id: Mapped[Optional[str]] = mapped_column("acceptedByUserId")
    accepted_via_identifier_id: Mapped[Optional[str]] = mapped_column("acceptedViaIdentifierId")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("OrganizationInvite_inviteCode_key", "inviteCode", unique=True),
        Index("OrganizationInvite_organizationId_idx", "organizationId"),
        Index("OrganizationInvite_requestedBy_idx", "requestedBy"),
    )

    organization: Mapped["Organization"] = relationship(back_populates="organization_invite")
    requested_by_user: Mapped[Optional["User"]] = relationship(back_populates="invite_requests")

class Monitor(Base):
    __tablename__ = "Monitor"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    project_id: Mapped[str] = mapped_column("projectId", ForeignKey("Project.id", ondelete="RESTRICT", onupdate="CASCADE", name="Monitor_projectId_fkey"))
    experiment_id: Mapped[Optional[str]] = mapped_column("experimentId", ForeignKey("Experiment.id", ondelete="RESTRICT", onupdate="RESTRICT", name="Monitor_experimentId_fkey"))
    evaluator_id: Mapped[Optional[str]] = mapped_column("evaluatorId", ForeignKey("Evaluator.id", ondelete="SET NULL", onupdate="CASCADE", name="Monitor_evaluatorId_fkey"))
    check_type: Mapped[str] = mapped_column("checkType")
    name: Mapped[str]
    slug: Mapped[str]
    execution_mode: Mapped[str] = mapped_column("executionMode", Enum("ON_MESSAGE", "AS_GUARDRAIL", "MANUALLY", name="EvaluationExecutionMode"), default="ON_MESSAGE", server_default=text("'ON_MESSAGE'"))
    enabled: Mapped[bool] = mapped_column(default=True, server_default=text("true"))
    preconditions: Mapped[dict[str, Any]] = mapped_column(JSONB)
    parameters: Mapped[dict[str, Any]] = mapped_column(JSONB)
    mappings: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))
    sample: Mapped[float] = mapped_column(default=1, server_default=text("1"))
    level: Mapped[str] = mapped_column(default="trace", server_default=text("'trace'"))
    thread_idle_timeout: Mapped[Optional[int]] = mapped_column("threadIdleTimeout")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    auto_paused_at: Mapped[Optional[datetime]] = mapped_column("autoPausedAt")
    auto_paused_reason: Mapped[Optional[str]] = mapped_column("autoPausedReason")
    evaluator_label: Mapped[Optional[str]] = mapped_column("evaluatorLabel")
    created_by_id: Mapped[Optional[str]] = mapped_column("createdById")
    updated_by_id: Mapped[Optional[str]] = mapped_column("updatedById")

    __table_args__ = (
        Index("Monitor_experimentId_key", "experimentId", unique=True),
        Index("Monitor_projectId_idx", "projectId"),
        Index("Monitor_evaluatorId_idx", "evaluatorId"),
        Index("Monitor_projectId_slug_key", "projectId", "slug", unique=True),
    )

    project: Mapped["Project"] = relationship(back_populates="checks")
    experiment: Mapped[Optional["Experiment"]] = relationship(back_populates="monitor")
    evaluator: Mapped[Optional["Evaluator"]] = relationship(back_populates="monitors")

class Cost(Base):
    __tablename__ = "Cost"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    project_id: Mapped[str] = mapped_column("projectId", ForeignKey("Project.id", ondelete="RESTRICT", onupdate="CASCADE", name="Cost_projectId_fkey"))
    cost_type: Mapped[str] = mapped_column("costType", Enum("TRACE_CHECK", "GUARDRAIL", "CLUSTERING", "BATCH_EVALUATION", name="CostType"))
    cost_name: Mapped[Optional[str]] = mapped_column("costName")
    reference_type: Mapped[str] = mapped_column("referenceType", Enum("CHECK", "TRACE", "PROJECT", "BATCH", name="CostReferenceType"))
    reference_id: Mapped[str] = mapped_column("referenceId")
    amount: Mapped[float]
    currency: Mapped[str]
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    extra_info: Mapped[Optional[dict[str, Any]]] = mapped_column("extraInfo", JSONB(none_as_null=True))

    __table_args__ = (
        Index("Cost_referenceType_referenceId_idx", "referenceType", "referenceId"),
        Index("Cost_costType_idx", "costType"),
        Index("Cost_projectId_idx", "projectId"),
        Index("Cost_projectId_createdAt_idx", "projectId", "createdAt"),
    )

    project: Mapped["Project"] = relationship(back_populates="costs")

class Topic(Base):
    """PROJECTION of the topic-clustering event stream (ADR-051): topics are
    `topics_recorded` events; this table is written ONLY by
    TopicModelFoldProjection (cursor: TopicModelProjection) and is
    rebuildable by replay. Kept named `Topic` because it is the domain
    entity every surface reads — the `*Projection` suffix is reserved for
    cursor/read-model tables.
    """

    __tablename__ = "Topic"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    project_id: Mapped[str] = mapped_column("projectId", ForeignKey("Project.id", ondelete="RESTRICT", onupdate="CASCADE", name="Topic_projectId_fkey"))
    name: Mapped[str]
    parent_id: Mapped[Optional[str]] = mapped_column("parentId", ForeignKey("Topic.id", ondelete="RESTRICT", onupdate="RESTRICT", name="Topic_parentId_fkey"))
    embeddings_model: Mapped[str]
    centroid: Mapped[dict[str, Any]] = mapped_column(JSONB)
    p95_distance: Mapped[float] = mapped_column("p95Distance")
    automatically_generated: Mapped[bool] = mapped_column("automaticallyGenerated", default=True, server_default=text("true"))
    last_event_id: Mapped[Optional[str]] = mapped_column("lastEventId")
    """The `topics_recorded` event that recorded this topic (its id in the
    event log). Null means the row predates event ownership and has not
    been seeded onto the stream yet; the projection sets it on every row
    it writes.
    """
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("Topic_parentId_idx", "parentId"),
        Index("Topic_projectId_idx", "projectId"),
    )

    project: Mapped["Project"] = relationship(back_populates="topics")
    parent: Mapped[Optional["Topic"]] = relationship(remote_side=[id], back_populates="subtopics")
    subtopics: Mapped[list["Topic"]] = relationship(passive_deletes="all", back_populates="parent")

class Dataset(Base):
    __tablename__ = "Dataset"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    project_id: Mapped[str] = mapped_column("projectId", ForeignKey("Project.id", ondelete="RESTRICT", onupdate="CASCADE", name="Dataset_projectId_fkey"))
    name: Mapped[str]
    slug: Mapped[str]
    column_types: Mapped[dict[str, Any]] = mapped_column("columnTypes", JSONB)
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    archived_at: Mapped[Optional[datetime]] = mapped_column("archivedAt")
    mapping: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))
    use_s3: Mapped[bool] = mapped_column("useS3", default=False, server_default=text("false"))
    s3_record_count: Mapped[Optional[int]] = mapped_column("s3RecordCount")
    content_layout: Mapped[str] = mapped_column("contentLayout", default="postgres", server_default=text("'postgres'"))
    status: Mapped[str] = mapped_column(default="ready", server_default=text("'ready'"))
    status_error: Mapped[Optional[str]] = mapped_column("statusError")
    staging_key: Mapped[Optional[str]] = mapped_column("stagingKey")
    source_stored_object_id: Mapped[Optional[str]] = mapped_column("sourceStoredObjectId")
    upload_filename: Mapped[Optional[str]] = mapped_column("uploadFilename")
    row_count: Mapped[Optional[int]] = mapped_column("rowCount")
    size_bytes: Mapped[Optional[int]] = mapped_column("sizeBytes", BigInteger)
    chunk_count: Mapped[Optional[int]] = mapped_column("chunkCount")
    chunk_offsets: Mapped[Optional[dict[str, Any]]] = mapped_column("chunkOffsets", JSONB(none_as_null=True))

    __table_args__ = (
        Index("Dataset_projectId_slug_key", "projectId", "slug", unique=True),
    )

    project: Mapped["Project"] = relationship(back_populates="datasets")
    dataset_records: Mapped[list["DatasetRecord"]] = relationship(passive_deletes="all", back_populates="dataset")
    batch_evaluations: Mapped[list["BatchEvaluation"]] = relationship(passive_deletes="all", back_populates="dataset")

class DatasetRecord(Base):
    __tablename__ = "DatasetRecord"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    dataset_id: Mapped[str] = mapped_column("datasetId", ForeignKey("Dataset.id", ondelete="RESTRICT", onupdate="CASCADE", name="DatasetRecord_datasetId_fkey"))
    project_id: Mapped[str] = mapped_column("projectId", ForeignKey("Project.id", ondelete="RESTRICT", onupdate="CASCADE", name="DatasetRecord_projectId_fkey"))
    entry: Mapped[dict[str, Any]] = mapped_column(JSONB)
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("DatasetRecord_datasetId_idx", "datasetId"),
        Index("DatasetRecord_projectId_idx", "projectId"),
        Index("DatasetRecord_createdAt_idx", "createdAt"),
        Index("DatasetRecord_projectId_datasetId_idx", "projectId", "datasetId"),
    )

    dataset: Mapped["Dataset"] = relationship(back_populates="dataset_records")
    project: Mapped["Project"] = relationship(back_populates="dataset_records")

class DatasetVersion(Base):
    """One immutable snapshot of a dataset's columns and rows; the rows are JSONL chunks in object storage."""

    __tablename__ = "DatasetVersion"

    id: Mapped[str] = mapped_column(primary_key=True)
    project_id: Mapped[str] = mapped_column("projectId")
    dataset_id: Mapped[str] = mapped_column("datasetId")
    version: Mapped[int]
    name: Mapped[Optional[str]]
    """A name given by hand; a named version is never thinned by retention."""
    trigger: Mapped[str]
    """"autosave", "manual", "experiment_run", "restore" or "initial"."""
    content_hash: Mapped[str] = mapped_column("contentHash")
    """SHA-256 over the columns and the rows in order; equal content saves no new version."""
    column_types: Mapped[dict[str, Any]] = mapped_column("columnTypes", JSONB)
    row_count: Mapped[int] = mapped_column("rowCount")
    size_bytes: Mapped[int] = mapped_column("sizeBytes", BigInteger)
    chunk_count: Mapped[int] = mapped_column("chunkCount")
    restored_from_version: Mapped[Optional[int]] = mapped_column("restoredFromVersion")
    author_id: Mapped[Optional[str]] = mapped_column("authorId")
    author_kind: Mapped[Optional[str]] = mapped_column("authorKind")
    """"user", "api_key" or "system": who saved the version when authorId is null."""
    pinned_at: Mapped[Optional[datetime]] = mapped_column("pinnedAt")
    """Set when an experiment pins this version; a pinned version is never thinned."""
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("DatasetVersion_projectId_datasetId_idx", "projectId", "datasetId"),
        Index("DatasetVersion_datasetId_version_key", "datasetId", "version", unique=True),
    )

class DatasetVersionRun(Base):
    """An experiment run that read one dataset version; a version a run read is never thinned."""

    __tablename__ = "DatasetVersionRun"

    id: Mapped[str] = mapped_column(primary_key=True)
    project_id: Mapped[str] = mapped_column("projectId")
    dataset_id: Mapped[str] = mapped_column("datasetId")
    version_id: Mapped[str] = mapped_column("versionId")
    run_id: Mapped[str] = mapped_column("runId")
    experiment_id: Mapped[Optional[str]] = mapped_column("experimentId")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("DatasetVersionRun_versionId_idx", "versionId"),
        Index("DatasetVersionRun_projectId_runId_datasetId_key", "projectId", "runId", "datasetId", unique=True),
    )

class Dashboard(Base):
    __tablename__ = "Dashboard"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    project_id: Mapped[str] = mapped_column("projectId", ForeignKey("Project.id", ondelete="RESTRICT", onupdate="CASCADE", name="Dashboard_projectId_fkey"))
    """The project that owns the board; only there can it be edited."""
    name: Mapped[str]
    order: Mapped[int] = mapped_column(default=0, server_default=text("0"))
    """Legacy board sort position. The Dashboards v2 app orders boards by a
    member's favourites and never writes this; kept because the legacy
    Analytics reports section and the REST reorder endpoint still write it,
    and a rolling deploy's older image still reads it.
    """
    description: Mapped[Optional[str]]
    """Free text under the board's name; null until a member writes one.
    Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
    """
    created_by_id: Mapped[Optional[str]] = mapped_column("createdById")
    """The member who created it; null for boards that predate the column.
    Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
    """
    scope: Mapped[str] = mapped_column(Enum("PRIVATE", "PROJECT", "ORGANIZATION", name="DashboardScope"), default="PROJECT", server_default=text("'PROJECT'"))
    """Who sees it. Only the author changes it; a board with no author stays PROJECT.
    Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
    """
    organization_id: Mapped[Optional[str]] = mapped_column("organizationId")
    """The owning project's organization, stamped when the board is first set to
    ORGANIZATION; null before. No relation, per the database rule.
    Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
    """
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("Dashboard_projectId_idx", "projectId"),
        Index("Dashboard_projectId_order_idx", "projectId", "order"),
        Index("Dashboard_organizationId_scope_idx", "organizationId", "scope"),
    )

    project: Mapped["Project"] = relationship(back_populates="dashboards")
    graphs: Mapped[list["CustomGraph"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="dashboard")

class DashboardFavourite(Base):
    """A member's star, in their own order within a project: a board or a From
    LangWatch template. Personal: one member's stars never show to another. No
    foreign key (per the database rule): the dashboard service removes a board's
    favourites when it deletes the board, so a star never outlives its dashboard.
    """

    __tablename__ = "DashboardFavourite"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    user_id: Mapped[str] = mapped_column("userId")
    dashboard_id: Mapped[Optional[str]] = mapped_column("dashboardId")
    """A DashboardsV2 board's id (copied boards kept their Dashboard id); null when
    the star is a template.
    """
    template_id: Mapped[Optional[str]] = mapped_column("templateId")
    """The starred From LangWatch template id; null when the star is a board."""
    project_id: Mapped[str] = mapped_column("projectId")
    position: Mapped[int]
    """The member's drag order among their own stars in the project."""
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("DashboardFavourite_userId_projectId_position_idx", "userId", "projectId", "position"),
        Index("DashboardFavourite_dashboardId_idx", "dashboardId"),
        Index("DashboardFavourite_userId_dashboardId_key", "userId", "dashboardId", unique=True),
        Index("DashboardFavourite_userId_projectId_templateId_key", "userId", "projectId", "templateId", unique=True),
    )

class DashboardsV2(Base):
    """A Dashboards v2 board, apart from the Analytics Reports boards in Dashboard. Boards
    made before the split were copied here with their ids, so stars and insight
    pointers still name them. No relation, per the database rule.
    """

    __tablename__ = "DashboardsV2"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    project_id: Mapped[str] = mapped_column("projectId")
    """The project that owns the board; only there can it be edited."""
    name: Mapped[str]
    description: Mapped[Optional[str]]
    """Free text under the board's name; null until a member writes one."""
    created_by_id: Mapped[Optional[str]] = mapped_column("createdById")
    """The member who created it; null for boards that predate authorship."""
    scope: Mapped[str] = mapped_column(Enum("PRIVATE", "PROJECT", "ORGANIZATION", name="DashboardScope"), default="PROJECT", server_default=text("'PROJECT'"))
    """Who sees it. Only the author changes it; a board with no author stays PROJECT."""
    organization_id: Mapped[Optional[str]] = mapped_column("organizationId")
    """The owning project's organization, stamped when the board is first set to
    ORGANIZATION; null before.
    """
    excluded_origins: Mapped[Optional[dict[str, Any]]] = mapped_column("excludedOrigins", JSONB(none_as_null=True))
    """Deprecated: unused since each widget query names its own sources (dashboards-v2 AC200,
    2026-10-28); dropped in a later migration, once no serving release reads it.
    """
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("DashboardsV2_projectId_idx", "projectId"),
        Index("DashboardsV2_organizationId_scope_idx", "organizationId", "scope"),
    )

class DashboardWidget(Base):
    """A Dashboards v2 widget: a srcdoc definition placed on a DashboardsV2 board, or
    unplaced. Widgets made before the split were copied from CustomGraph with their
    ids. No relation, per the database rule.
    """

    __tablename__ = "DashboardWidget"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    project_id: Mapped[str] = mapped_column("projectId")
    dashboard_id: Mapped[Optional[str]] = mapped_column("dashboardId")
    """The DashboardsV2 board it sits on; null while unplaced."""
    name: Mapped[str]
    definition: Mapped[dict[str, Any]] = mapped_column(JSONB)
    """The widget's srcdoc definition."""
    grid_column: Mapped[int] = mapped_column("gridColumn", default=0, server_default=text("0"))
    grid_row: Mapped[int] = mapped_column("gridRow", default=0, server_default=text("0"))
    col_span: Mapped[int] = mapped_column("colSpan", default=1, server_default=text("1"))
    row_span: Mapped[int] = mapped_column("rowSpan", default=1, server_default=text("1"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("DashboardWidget_projectId_dashboardId_idx", "projectId", "dashboardId"),
    )

class InsightProjection(Base):
    """One insight in its owner's inbox in a project, folded from its `insight` aggregate
    (modules/insight). A read model: rebuildable by replaying the event log. Epochs in
    milliseconds.
    """

    __tablename__ = "InsightProjection"

    id: Mapped[str] = mapped_column(primary_key=True)
    """The insight id, which is also its aggregate id."""
    project_id: Mapped[str] = mapped_column("projectId")
    title: Mapped[str]
    body: Mapped[str]
    tone: Mapped[str]
    """bad | watch | good"""
    topic: Mapped[Optional[str]]
    valid_days: Mapped[int] = mapped_column("validDays")
    lwql: Mapped[Optional[str]]
    source_conversation_id: Mapped[Optional[str]] = mapped_column("sourceConversationId")
    source_message_id: Mapped[Optional[str]] = mapped_column("sourceMessageId")
    filed_by_user_id: Mapped[Optional[str]] = mapped_column("filedByUserId")
    filed_at: Mapped[float] = mapped_column("filedAt")
    renewed_at: Mapped[Optional[float]] = mapped_column("renewedAt")
    """Set by a later run that found the insight still true."""
    created_at: Mapped[float] = mapped_column("createdAt")
    updated_at: Mapped[float] = mapped_column("updatedAt")
    occurred_at: Mapped[float] = mapped_column("occurredAt")
    accepted_at: Mapped[float] = mapped_column("acceptedAt")
    last_event_id: Mapped[str] = mapped_column("lastEventId")
    projection_version: Mapped[str] = mapped_column("projectionVersion")
    board_id: Mapped[Optional[str]] = mapped_column("boardId")
    """The board and widget it came from. A pointer only, never a relation: the row reads the
    same once either is deleted, so the names are kept as filed.
    """
    board_name: Mapped[Optional[str]] = mapped_column("boardName")
    widget_id: Mapped[Optional[str]] = mapped_column("widgetId")
    widget_name: Mapped[Optional[str]] = mapped_column("widgetName")
    filed_via: Mapped[str] = mapped_column("filedVia", default="chat", server_default=text("'chat'"))
    """chat | run"""
    replay_start: Mapped[Optional[float]] = mapped_column("replayStart")
    """The fixed window `lwql` replays over (epoch milliseconds, half-open) and its step."""
    replay_end: Mapped[Optional[float]] = mapped_column("replayEnd")
    replay_granularity_seconds: Mapped[Optional[int]] = mapped_column("replayGranularitySeconds")
    replay_context: Mapped[Optional[dict[str, Any]]] = mapped_column("replayContext", JSONB(none_as_null=True))
    """The values in force when it was filed: `{ period, parameters }`."""
    owner_user_id: Mapped[Optional[str]] = mapped_column("ownerUserId")
    """Whose insight it is: the one person who reads it. NULL on a row folded before the
    column existed, which is owned by whoever filed it (`filedByUserId`).
    """
    finding_key: Mapped[Optional[str]] = mapped_column("findingKey")
    """The stable key a daily run matches a repeated finding by, so it renews the insight
    and files no twin. NULL for an insight saved from a chat and on a row folded before
    the column existed.
    """

    __table_args__ = (
        Index("InsightProjection_projectId_filedAt_idx", "projectId", "filedAt"),
        Index("InsightProjection_projectId_ownerUserId_filedAt_idx", "projectId", "ownerUserId", "filedAt"),
    )

class InsightReaderProjection(Base):
    """One reader's own state on one insight (seen, done, kept), folded from the same stream.
    An insight has one reader today, its owner; a later forward may add more.
    """

    __tablename__ = "InsightReaderProjection"

    id: Mapped[str] = mapped_column(primary_key=True)
    project_id: Mapped[str] = mapped_column("projectId")
    insight_id: Mapped[str] = mapped_column("insightId")
    user_id: Mapped[str] = mapped_column("userId")
    seen_at: Mapped[Optional[float]] = mapped_column("seenAt")
    archived_at: Mapped[Optional[float]] = mapped_column("archivedAt")
    kept_at: Mapped[Optional[float]] = mapped_column("keptAt")
    created_at: Mapped[float] = mapped_column("createdAt")
    updated_at: Mapped[float] = mapped_column("updatedAt")
    occurred_at: Mapped[float] = mapped_column("occurredAt")
    accepted_at: Mapped[float] = mapped_column("acceptedAt")
    last_event_id: Mapped[str] = mapped_column("lastEventId")
    projection_version: Mapped[str] = mapped_column("projectionVersion")

    __table_args__ = (
        Index("InsightReaderProjection_projectId_userId_idx", "projectId", "userId"),
        Index("InsightReaderProjection_projectId_insightId_userId_key", "projectId", "insightId", "userId", unique=True),
    )

class InsightDailyScheduleProjection(Base):
    """One person's daily insights run on one board in a project, folded from its
    `insight_daily_schedule` aggregate (modules/insight). A read model: rebuildable by
    replaying the event log. Epochs in milliseconds.
    """

    __tablename__ = "InsightDailyScheduleProjection"

    id: Mapped[str] = mapped_column(primary_key=True)
    """The schedule id, which is also its aggregate id: one per project, person and board."""
    project_id: Mapped[str] = mapped_column("projectId")
    user_id: Mapped[str] = mapped_column("userId")
    board_kind: Mapped[str] = mapped_column("boardKind")
    """dashboard | template. A pointer only, never a relation: a run checks the board is there."""
    board_id: Mapped[str] = mapped_column("boardId")
    board_name: Mapped[str] = mapped_column("boardName")
    state: Mapped[str] = mapped_column(default="undecided", server_default=text("'undecided'"))
    """on | off | undecided."""
    off_reason: Mapped[Optional[str]] = mapped_column("offReason")
    """board_deleted | skipped_days | member_removed | project_archived. Why the system turned
    it off. NULL while it is on, and when the person turned it off.
    """
    hour: Mapped[Optional[int]]
    timezone_: Mapped[Optional[str]] = mapped_column("timezone")
    max_insights: Mapped[Optional[int]] = mapped_column("maxInsights")
    last_run_id: Mapped[Optional[str]] = mapped_column("lastRunId")
    """The last run that recorded an outcome. NULL until one did."""
    last_run_at: Mapped[Optional[float]] = mapped_column("lastRunAt")
    last_run_outcome: Mapped[Optional[str]] = mapped_column("lastRunOutcome")
    """filed | nothing | failed | skipped"""
    last_run_reason: Mapped[Optional[str]] = mapped_column("lastRunReason")
    last_run_filed: Mapped[Optional[int]] = mapped_column("lastRunFiled")
    last_run_renewed: Mapped[Optional[int]] = mapped_column("lastRunRenewed")
    """How many insights the run found still true and renewed."""
    last_run_conversation_id: Mapped[Optional[str]] = mapped_column("lastRunConversationId")
    consecutive_skips: Mapped[int] = mapped_column("consecutiveSkips", default=0, server_default=text("0"))
    """The scheduled days in a row whose run ended skipped, and the slot of the newest counted."""
    skip_streak_slot: Mapped[Optional[float]] = mapped_column("skipStreakSlot")
    created_at: Mapped[float] = mapped_column("createdAt")
    updated_at: Mapped[float] = mapped_column("updatedAt")
    occurred_at: Mapped[float] = mapped_column("occurredAt")
    accepted_at: Mapped[float] = mapped_column("acceptedAt")
    last_event_id: Mapped[str] = mapped_column("lastEventId")
    projection_version: Mapped[str] = mapped_column("projectionVersion")

    __table_args__ = (
        Index("InsightDailySchedule_projectId_userId_board_key", "projectId", "userId", "boardKind", "boardId", unique=True),
    )

class SavedView(Base):
    __tablename__ = "SavedView"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    project_id: Mapped[str] = mapped_column("projectId", ForeignKey("Project.id", ondelete="RESTRICT", onupdate="CASCADE", name="SavedView_projectId_fkey"))
    user_id: Mapped[Optional[str]] = mapped_column("userId", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="SavedView_userId_fkey"))
    name: Mapped[str]
    filters: Mapped[dict[str, Any]] = mapped_column(JSONB, default=lambda: {}, server_default=text("'{}'"))
    query: Mapped[Optional[str]]
    period: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))
    order: Mapped[int] = mapped_column(default=0, server_default=text("0"))
    kind: Mapped[str] = mapped_column(default="v1-traces-filter", server_default=text("'v1-traces-filter'"))
    """Discriminator for the two storage shapes that share this table:
    - "v1-traces-filter" (default) - legacy /messages page filter map saved views
    - "v2-traces-lens"             - new traces v2 lens configs (columns, addons, grouping, sort, filterText) encoded inside `filters` JSON
    Each client filters by its own kind on read so neither sees the other's rows.
    """
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("SavedView_projectId_idx", "projectId"),
        Index("SavedView_projectId_order_idx", "projectId", "order"),
        Index("SavedView_userId_idx", "userId"),
        Index("SavedView_projectId_kind_order_idx", "projectId", "kind", "order"),
    )

    project: Mapped["Project"] = relationship(back_populates="saved_views")
    user: Mapped[Optional["User"]] = relationship(back_populates="saved_views")

class CustomGraph(Base):
    __tablename__ = "CustomGraph"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    project_id: Mapped[str] = mapped_column("projectId", ForeignKey("Project.id", ondelete="RESTRICT", onupdate="CASCADE", name="CustomGraph_projectId_fkey"))
    name: Mapped[str]
    graph: Mapped[dict[str, Any]] = mapped_column(JSONB)
    """The chart's definition, in the shape its `kind` names."""
    filters: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))
    kind: Mapped[str] = mapped_column(default="builder", server_default=text("'builder'"))
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
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    dashboard_id: Mapped[Optional[str]] = mapped_column("dashboardId", ForeignKey("Dashboard.id", ondelete="CASCADE", onupdate="CASCADE", name="CustomGraph_dashboardId_fkey"))
    grid_column: Mapped[int] = mapped_column("gridColumn", default=0, server_default=text("0"))
    grid_row: Mapped[int] = mapped_column("gridRow", default=0, server_default=text("0"))
    col_span: Mapped[int] = mapped_column("colSpan", default=1, server_default=text("1"))
    row_span: Mapped[int] = mapped_column("rowSpan", default=1, server_default=text("1"))

    __table_args__ = (
        Index("CustomGraph_projectId_idx", "projectId"),
        Index("CustomGraph_dashboardId_idx", "dashboardId"),
        Index("CustomGraph_projectId_kind_idx", "projectId", "kind"),
    )

    project: Mapped["Project"] = relationship(back_populates="custom_graphs")
    dashboard: Mapped[Optional["Dashboard"]] = relationship(back_populates="graphs")
    trigger_sent: Mapped[list["TriggerSent"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="custom_graph")
    trigger: Mapped[Optional["Trigger"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="custom_graph")

class BatchEvaluation(Base):
    __tablename__ = "BatchEvaluation"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    experiment_id: Mapped[str] = mapped_column("experimentId", ForeignKey("Experiment.id", ondelete="RESTRICT", onupdate="CASCADE", name="BatchEvaluation_experimentId_fkey"))
    project_id: Mapped[str] = mapped_column("projectId", ForeignKey("Project.id", ondelete="RESTRICT", onupdate="CASCADE", name="BatchEvaluation_projectId_fkey"))
    data: Mapped[dict[str, Any]] = mapped_column(JSONB, default=lambda: {}, server_default=text("'{}'"))
    status: Mapped[str]
    score: Mapped[float]
    label: Mapped[Optional[str]]
    passed: Mapped[bool]
    details: Mapped[str]
    cost: Mapped[float]
    dataset_slug: Mapped[str] = mapped_column("datasetSlug")
    dataset_id: Mapped[str] = mapped_column("datasetId", ForeignKey("Dataset.id", ondelete="RESTRICT", onupdate="CASCADE", name="BatchEvaluation_datasetId_fkey"))
    evaluation: Mapped[str]
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("BatchEvaluation_projectId_idx", "projectId"),
        Index("BatchEvaluation_datasetId_idx", "datasetId"),
        Index("BatchEvaluation_experimentId_idx", "experimentId"),
        Index("BatchEvaluation_id_key", "id", unique=True),
    )

    experiment: Mapped["Experiment"] = relationship(back_populates="batch_evaluations")
    project: Mapped["Project"] = relationship(back_populates="batch_evaluations")
    dataset: Mapped["Dataset"] = relationship(back_populates="batch_evaluations")

class Trigger(Base):
    __tablename__ = "Trigger"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    name: Mapped[str]
    project_id: Mapped[str] = mapped_column("projectId", ForeignKey("Project.id", ondelete="RESTRICT", onupdate="CASCADE", name="Trigger_projectId_fkey"))
    action: Mapped[str] = mapped_column(Enum("SEND_EMAIL", "ADD_TO_DATASET", "ADD_TO_ANNOTATION_QUEUE", "SEND_SLACK_MESSAGE", "SEND_WEBHOOK", name="TriggerAction"))
    trigger_kind: Mapped[str] = mapped_column("triggerKind", Enum("AUTOMATION", "ALERT", "REPORT", name="TriggerKind"), default="AUTOMATION", server_default=text("'AUTOMATION'"))
    action_params: Mapped[dict[str, Any]] = mapped_column("actionParams", JSONB)
    filters: Mapped[dict[str, Any]] = mapped_column(JSONB)
    filter_query: Mapped[Optional[str]] = mapped_column("filterQuery")
    last_run_at: Mapped[float] = mapped_column("lastRunAt", default=0, server_default=text("0"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    active: Mapped[bool] = mapped_column(default=True, server_default=text("true"))
    paused_reason: Mapped[Optional[str]] = mapped_column("pausedReason")
    paused_at: Mapped[Optional[datetime]] = mapped_column("pausedAt", UtcDateTimeTz(precision=3))
    message: Mapped[Optional[str]]
    deleted: Mapped[bool] = mapped_column(default=False, server_default=text("false"))
    alert_type: Mapped[Optional[str]] = mapped_column("alertType", Enum("CRITICAL", "WARNING", "INFO", name="AlertType"))
    slack_template_type: Mapped[Optional[str]] = mapped_column("slackTemplateType")
    slack_template: Mapped[Optional[str]] = mapped_column("slackTemplate")
    email_subject_template: Mapped[Optional[str]] = mapped_column("emailSubjectTemplate")
    email_body_template: Mapped[Optional[str]] = mapped_column("emailBodyTemplate")
    notification_cadence: Mapped[str] = mapped_column("notificationCadence", default="immediate", server_default=text("'immediate'"))
    trace_debounce_ms: Mapped[int] = mapped_column("traceDebounceMs", default=30000, server_default=text("30000"))
    custom_graph_id: Mapped[Optional[str]] = mapped_column("customGraphId", ForeignKey("CustomGraph.id", ondelete="CASCADE", onupdate="CASCADE", name="Trigger_customGraphId_fkey"))

    __table_args__ = (
        Index("Trigger_customGraphId_key", "customGraphId", unique=True),
        Index("Trigger_projectId_idx", "projectId"),
    )

    project: Mapped["Project"] = relationship(back_populates="triggers")
    custom_graph: Mapped[Optional["CustomGraph"]] = relationship(back_populates="trigger")
    trigger_sent: Mapped[list["TriggerSent"]] = relationship(passive_deletes="all", back_populates="trigger")
    webhook_deliveries: Mapped[list["WebhookEndpointDelivery"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="trigger")
    latest_evaluation: Mapped[Optional["TriggerLatestEvaluation"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="trigger")

class SlackIntegration(Base):
    """One named Slack connection (ADR-093 §5a); an organization holds as many as
    it likes, each scoped to the organization or to one project. The secret is
    AES-256-GCM ciphertext and never returned to a client. `secretFingerprint`
    is an HMAC of the plaintext, so one secret is one connection per scope
    without decrypting anything. No foreign key: `scopeId` is
    polymorphic (ADR-021), as on RetentionPolicy and CustomLLMModelCost.
    """

    __tablename__ = "SlackIntegration"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    name: Mapped[str]
    kind: Mapped[str] = mapped_column(Enum("BOT", "INCOMING_WEBHOOK", name="SlackIntegrationKind"))
    scope_type: Mapped[str] = mapped_column("scopeType", Enum("ORGANIZATION", "PROJECT", name="SlackIntegrationScopeType"))
    scope_id: Mapped[str] = mapped_column("scopeId")
    organization_id: Mapped[str] = mapped_column("organizationId")
    bot_token_encrypted: Mapped[Optional[str]] = mapped_column("botTokenEncrypted")
    webhook_url_encrypted: Mapped[Optional[str]] = mapped_column("webhookUrlEncrypted")
    secret_fingerprint: Mapped[str] = mapped_column("secretFingerprint")
    secret_hint: Mapped[str] = mapped_column("secretHint")
    slack_team_id: Mapped[Optional[str]] = mapped_column("slackTeamId")
    slack_team_name: Mapped[Optional[str]] = mapped_column("slackTeamName")
    created_by_id: Mapped[str] = mapped_column("createdById")
    updated_by_id: Mapped[str] = mapped_column("updatedById")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("SlackIntegration_scopeType_scopeId_idx", "scopeType", "scopeId"),
        Index("SlackIntegration_organizationId_scopeType_scopeId_secretFin_key", "organizationId", "scopeType", "scopeId", "secretFingerprint", unique=True),
    )

class SlackConnectionClaim(Base):
    """An automation holding a Slack connection in use (ARCHITECTURE.md §3, the
    claims ruling): a claimed connection is never deleted. `projectId` is the
    claimant's project; `organizationId` is the connection's tenancy anchor.
    No foreign key: slack's service refuses to delete a claimed connection.
    """

    __tablename__ = "slack_connection_claim"

    connection_id: Mapped[str] = mapped_column("connectionId", primary_key=True)
    claimant_id: Mapped[str] = mapped_column("claimantId", primary_key=True)
    claimant_label: Mapped[str] = mapped_column("claimantLabel")
    organization_id: Mapped[str] = mapped_column("organizationId")
    project_id: Mapped[str] = mapped_column("projectId")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("slack_connection_claim_organizationId_idx", "organizationId"),
    )

class WebhookEndpoint(Base):
    """Org-anchored outbound webhook endpoint (the webhook platform, enterprise
    only). The signing secret is AES-256-GCM encrypted at rest (shared
    encrypt/decrypt, CREDENTIALS_SECRET) and returned in plaintext exactly
    once, at create or roll time. `enabledEvents` holds registry event types
    or family wildcards ("gateway.*", "*"). Failure-streak columns drive the
    72h auto-disable. Spec: specs/webhooks/webhook-endpoints.feature.
    """

    __tablename__ = "WebhookEndpoint"

    id: Mapped[str] = mapped_column(primary_key=True)
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="WebhookEndpoint_organizationId_fkey"))
    destination_kind: Mapped[str] = mapped_column("destinationKind", Enum("http", "sqs", name="WebhookDestinationKind"), default="http", server_default=text("'http'"))
    """Where this endpoint delivers. Everything above the last hop is shared:
    same batching, same retry ladder, same delivery log, same signature over
    the same bytes.
    """
    url: Mapped[Optional[str]]
    """The receiver URL. Present on `http` endpoints and null on the others;
    the per-kind shape is pinned by a CHECK constraint Prisma cannot model.
    """
    sqs_queue_url: Mapped[Optional[str]] = mapped_column("sqsQueueUrl")
    """The Amazon SQS queue this endpoint delivers to, on `sqs` endpoints. The
    region and the owning account are read off it rather than configured
    separately, so they can never disagree with the queue.
    """
    sqs_role_arn: Mapped[Optional[str]] = mapped_column("sqsRoleArn")
    """The role we assume to write to the queue, and the ExternalId that proves
    the request came from us rather than from anyone who learned the role's
    name. The recommended credential mode.
    """
    sqs_external_id: Mapped[Optional[str]] = mapped_column("sqsExternalId")
    sqs_access_key_id: Mapped[Optional[str]] = mapped_column("sqsAccessKeyId")
    """A static key pair for the queue, encrypted at rest exactly like the
    signing secret. Both halves are present or both are absent: a half-filled
    pair is what the AWS SDK reads as "we have credentials" and stops looking
    with, so the CHECK constraint refuses it.
    """
    sqs_secret_access_key_encrypted: Mapped[Optional[str]] = mapped_column("sqsSecretAccessKeyEncrypted")
    secret_encrypted: Mapped[str] = mapped_column("secretEncrypted")
    signature_scheme: Mapped[Optional[str]] = mapped_column("signatureScheme")
    """How deliveries are formatted and signed. Null: the batch envelope signed
    `t=,v1=`. `legacy_sha256`: one raw message per POST signed `sha256=`, kept
    only for endpoints migrated from governance anomaly destinations.
    """
    idempotency_key: Mapped[Optional[str]] = mapped_column("idempotencyKey")
    """Set only by governance's anomaly destination migration: a create that
    repeats it in the same organization answers the endpoint it made. An
    archived endpoint gives its key up, so a retry after cleanup creates anew.
    """
    previous_secret_encrypted: Mapped[Optional[str]] = mapped_column("previousSecretEncrypted")
    """The secret this endpoint signed with before the last roll. Deliveries
    carry a second `v1=` computed from it until it expires, so a receiver
    swaps secrets on its own schedule instead of dropping deliveries in the
    gap. Null once the window closes or if the secret was never rolled.
    """
    previous_secret_expires_at: Mapped[Optional[datetime]] = mapped_column("previousSecretExpiresAt")
    """When the previous secret stops being signed with and accepted."""
    enabled_events: Mapped[list[str]] = mapped_column("enabledEvents", ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))
    status: Mapped[str] = mapped_column(Enum("ACTIVE", "DISABLED", name="WebhookEndpointStatus"), default="ACTIVE", server_default=text("'ACTIVE'"))
    disabled_reason: Mapped[Optional[str]] = mapped_column("disabledReason")
    """Why the endpoint is disabled: "auto_failures_72h" or "manual"."""
    disabled_at: Mapped[Optional[datetime]] = mapped_column("disabledAt")
    failing_since: Mapped[Optional[datetime]] = mapped_column("failingSince")
    """Start of the current consecutive-failure streak; cleared on any 2xx.
    72 hours of unbroken failures from this instant auto-disables.
    """
    max_batch_size: Mapped[int] = mapped_column("maxBatchSize", default=100, server_default=text("100"))
    """Delivery controls, each server-bounded (see webhookDeliveryControls
    constants): envelopes per POST, coalescing window before a partial
    batch ships, and concurrent POSTs to this receiver.
    """
    max_batch_delay_ms: Mapped[int] = mapped_column("maxBatchDelayMs", default=250, server_default=text("250"))
    max_in_flight: Mapped[int] = mapped_column("maxInFlight", default=4, server_default=text("4"))
    allow_self_signed_certificate: Mapped[bool] = mapped_column("allowSelfSignedCertificate", default=False, server_default=text("false"))
    """Self-hosted only: deliveries accept the receiver's self-signed TLS
    certificate. Off by default; the hosted product ignores it.
    """
    last_success_at: Mapped[Optional[datetime]] = mapped_column("lastSuccessAt")
    last_failure_at: Mapped[Optional[datetime]] = mapped_column("lastFailureAt")
    archived_at: Mapped[Optional[datetime]] = mapped_column("archivedAt")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("WebhookEndpoint_organizationId_status_idx", "organizationId", "status"),
        Index("WebhookEndpoint_organizationId_idempotencyKey_key", "organizationId", "idempotencyKey", unique=True),
    )

    organization: Mapped["Organization"] = relationship(back_populates="webhook_endpoints")
    deliveries: Mapped[list["WebhookEndpointDelivery"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="endpoint")

class WebhookEndpointDelivery(Base):
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

    __tablename__ = "WebhookEndpointDelivery"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    channel: Mapped[str] = mapped_column(Enum("platform", "automations", name="WebhookDeliveryChannel"), default="platform", server_default=text("'platform'"))
    organization_id: Mapped[Optional[str]] = mapped_column("organizationId")
    """Platform tenancy. Null on automations rows."""
    endpoint_id: Mapped[Optional[str]] = mapped_column("endpointId", ForeignKey("WebhookEndpoint.id", ondelete="CASCADE", onupdate="CASCADE", name="WebhookEndpointDelivery_endpointId_fkey"))
    project_id: Mapped[Optional[str]] = mapped_column("projectId", ForeignKey("Project.id", ondelete="CASCADE", onupdate="CASCADE", name="WebhookEndpointDelivery_projectId_fkey"))
    """Automations tenancy. Null on platform rows."""
    trigger_id: Mapped[Optional[str]] = mapped_column("triggerId", ForeignKey("Trigger.id", ondelete="CASCADE", onupdate="CASCADE", name="WebhookEndpointDelivery_triggerId_fkey"))
    dispatch_id: Mapped[str] = mapped_column("dispatchId")
    attempt: Mapped[Optional[int]]
    """1-based attempt and the batch's envelope count. Both null on automations
    rows, which record neither: one trigger fire is one event, and the
    channel's sender does not thread its attempt number into the log.
    """
    event_count: Mapped[Optional[int]] = mapped_column("eventCount")
    response_status: Mapped[Optional[int]] = mapped_column("responseStatus")
    latency_ms: Mapped[Optional[int]] = mapped_column("latencyMs")
    error: Mapped[Optional[str]]
    response: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))
    outcome: Mapped[str] = mapped_column(Enum("success", "retryable", "terminal", "pending", name="WebhookDeliveryOutcome"))
    fired_at: Mapped[datetime] = mapped_column("firedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("WebhookEndpointDelivery_endpointId_firedAt_idx", "endpointId", "firedAt"),
        Index("WebhookEndpointDelivery_organizationId_firedAt_idx", "organizationId", "firedAt"),
        Index("WebhookEndpointDelivery_firedAt_idx", "firedAt"),
        Index("WebhookEndpointDelivery_endpointId_dispatchId_idx", "endpointId", "dispatchId"),
        Index("WebhookEndpointDelivery_projectId_triggerId_firedAt_idx", "projectId", "triggerId", "firedAt"),
        Index("WebhookEndpointDelivery_projectId_triggerId_dispatchId_idx", "projectId", "triggerId", "dispatchId"),
    )

    endpoint: Mapped[Optional["WebhookEndpoint"]] = relationship(back_populates="deliveries")
    project: Mapped[Optional["Project"]] = relationship(back_populates="webhook_deliveries")
    trigger: Mapped[Optional["Trigger"]] = relationship(back_populates="webhook_deliveries")

class Experiment(Base):
    __tablename__ = "Experiment"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    name: Mapped[Optional[str]]
    type: Mapped[str] = mapped_column(Enum("DSPY", "BATCH_EVALUATION", "BATCH_EVALUATION_V2", "EVALUATIONS_V3", name="ExperimentType"))
    slug: Mapped[str]
    project_id: Mapped[str] = mapped_column("projectId", ForeignKey("Project.id", ondelete="RESTRICT", onupdate="CASCADE", name="Experiment_projectId_fkey"))
    workflow_id: Mapped[Optional[str]] = mapped_column("workflowId", ForeignKey("Workflow.id", ondelete="SET NULL", onupdate="CASCADE", name="Experiment_workflowId_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    archived_at: Mapped[Optional[datetime]] = mapped_column("archivedAt")
    workbench_state: Mapped[Optional[dict[str, Any]]] = mapped_column("workbenchState", JSONB(none_as_null=True))
    workbench_version: Mapped[int] = mapped_column("workbenchVersion", default=0, server_default=text("0"))

    __table_args__ = (
        Index("Experiment_projectId_idx", "projectId"),
        Index("Experiment_workflowId_idx", "workflowId"),
        Index("Experiment_projectId_archivedAt_idx", "projectId", "archivedAt"),
        Index("Experiment_projectId_slug_key", "projectId", "slug", unique=True),
    )

    project: Mapped["Project"] = relationship(back_populates="experiments")
    workflow: Mapped[Optional["Workflow"]] = relationship(back_populates="experiments")
    batch_evaluations: Mapped[list["BatchEvaluation"]] = relationship(passive_deletes="all", back_populates="experiment")
    monitor: Mapped[Optional["Monitor"]] = relationship(passive_deletes="all", back_populates="experiment")

class ExperimentVersion(Base):
    __tablename__ = "ExperimentVersion"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    experiment_id: Mapped[str] = mapped_column("experimentId")
    project_id: Mapped[str] = mapped_column("projectId")
    version: Mapped[int]
    """The number the version is addressed and displayed by. Numbered rows hold
    a gap-free sequence. The rolling autosave row rides on the counter
    instead, which is what keeps it clear of that sequence.
    """
    counter_version: Mapped[int] = mapped_column("counterVersion")
    """`Experiment.workbenchVersion` at the moment this row was last written.
    The rolling row is rewritten in place, so this is the only column that
    says how recent its content is, and the version list sorts and pages by
    it. It also joins a row to the live state: the row whose `counterVersion`
    equals the experiment's counter holds what the workbench shows now.
    """
    auto_saved: Mapped[bool] = mapped_column("autoSaved", default=False, server_default=text("false"))
    commit_message: Mapped[Optional[str]] = mapped_column("commitMessage")
    author_id: Mapped[Optional[str]] = mapped_column("authorId")
    author_label: Mapped[str] = mapped_column("authorLabel")
    """Who wrote this version: "user", "langy" or "api"."""
    run_id: Mapped[Optional[str]] = mapped_column("runId")
    """The run that wrote this version, when a run wrote it.

    A run writes its cells into the workbench state, which advances the
    counter. Without this, the page that started that run reads the bump as
    somebody else's write, refuses to save, and asks the reader to reload
    over their own unsaved edits. Named here, the page adopts the version and
    carries on.
    """
    state: Mapped[dict[str, Any]] = mapped_column(JSONB)
    schema_version: Mapped[str] = mapped_column("schemaVersion", default="1", server_default=text("'1'"))
    """Schema version of `state`, so a later shape change can be read forward."""
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("ExperimentVersion_projectId_experimentId_autoSaved_idx", "projectId", "experimentId", "autoSaved"),
        Index("ExperimentVersion_projectId_experimentId_version_idx", "projectId", "experimentId", "version"),
        Index("ExperimentVersion_projectId_experimentId_counterVersion_idx", "projectId", "experimentId", "counterVersion"),
        Index("ExperimentVersion_createdAt_idx", "createdAt"),
        Index("ExperimentVersion_experimentId_version_key", "experimentId", "version", unique=True),
    )

class Annotation(Base):
    __tablename__ = "Annotation"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    project_id: Mapped[str] = mapped_column("projectId", ForeignKey("Project.id", ondelete="RESTRICT", onupdate="CASCADE", name="Annotation_projectId_fkey"))
    trace_id: Mapped[str] = mapped_column("traceId")
    comment: Mapped[Optional[str]]
    is_thumbs_up: Mapped[Optional[bool]] = mapped_column("isThumbsUp")
    user_id: Mapped[Optional[str]] = mapped_column("userId", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="Annotation_userId_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    email: Mapped[Optional[str]]
    score_options: Mapped[Optional[dict[str, Any]]] = mapped_column("scoreOptions", JSONB(none_as_null=True))
    expected_output: Mapped[Optional[str]] = mapped_column("expectedOutput")
    anchor_kind: Mapped[Optional[str]] = mapped_column("anchorKind")
    anchor_id: Mapped[Optional[str]] = mapped_column("anchorId")
    anchor_path: Mapped[Optional[str]] = mapped_column("anchorPath")

    __table_args__ = (
        Index("Annotation_projectId_idx", "projectId"),
        Index("Annotation_traceId_idx", "traceId"),
        Index("Annotation_userId_idx", "userId"),
        Index("Annotation_projectId_traceId_idx", "projectId", "traceId"),
        Index("Annotation_projectId_traceId_anchorKind_idx", "projectId", "traceId", "anchorKind"),
    )

    project: Mapped["Project"] = relationship(back_populates="annotations")
    user: Mapped[Optional["User"]] = relationship(back_populates="annotation")

class ModelProvider(Base):
    __tablename__ = "ModelProvider"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    name: Mapped[str]
    provider: Mapped[str]
    routing_handle: Mapped[Optional[str]] = mapped_column("routingHandle")
    enabled: Mapped[bool]
    custom_keys: Mapped[Optional[dict[str, Any]]] = mapped_column("customKeys", JSONB(none_as_null=True))
    extra_headers: Mapped[Optional[dict[str, Any]]] = mapped_column("extraHeaders", JSONB(none_as_null=True))
    custom_models: Mapped[Optional[dict[str, Any]]] = mapped_column("customModels", JSONB(none_as_null=True))
    custom_embeddings_models: Mapped[Optional[dict[str, Any]]] = mapped_column("customEmbeddingsModels", JSONB(none_as_null=True))
    deployment_mapping: Mapped[Optional[dict[str, Any]]] = mapped_column("deploymentMapping", JSONB(none_as_null=True))
    rate_limit_rpm: Mapped[Optional[int]] = mapped_column("rateLimitRpm")
    rate_limit_tpm: Mapped[Optional[int]] = mapped_column("rateLimitTpm")
    rate_limit_rpd: Mapped[Optional[int]] = mapped_column("rateLimitRpd")
    rotation_policy: Mapped[str] = mapped_column("rotationPolicy", Enum("MANUAL", name="GatewayProviderRotationPolicy"), default="MANUAL", server_default=text("'MANUAL'"))
    provider_config: Mapped[Optional[dict[str, Any]]] = mapped_column("providerConfig", JSONB(none_as_null=True))
    fallback_priority_global: Mapped[Optional[int]] = mapped_column("fallbackPriorityGlobal")
    langy_skip_permissions_models: Mapped[Optional[dict[str, Any]]] = mapped_column("langySkipPermissionsModels", JSONB(none_as_null=True))
    health_status: Mapped[str] = mapped_column("healthStatus", Enum("UNKNOWN", "HEALTHY", "DEGRADED", "CIRCUIT_OPEN", name="GatewayProviderHealthStatus"), default="UNKNOWN", server_default=text("'UNKNOWN'"))
    circuit_opened_at: Mapped[Optional[datetime]] = mapped_column("circuitOpenedAt")
    last_health_check_at: Mapped[Optional[datetime]] = mapped_column("lastHealthCheckAt")
    disabled_at: Mapped[Optional[datetime]] = mapped_column("disabledAt")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    organization_id: Mapped[str] = mapped_column("organizationId")

    __table_args__ = (
        Index("ModelProvider_healthStatus_idx", "healthStatus"),
        Index("ModelProvider_fallbackPriorityGlobal_idx", "fallbackPriorityGlobal"),
        Index("ModelProvider_organizationId_idx", "organizationId"),
        Index("ModelProvider_organizationId_routingHandle_key", "organizationId", "routingHandle", unique=True),
    )

    scopes: Mapped[list["ModelProviderScope"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="model_provider")

class ModelProviderScope(Base):
    __tablename__ = "ModelProviderScope"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    model_provider_id: Mapped[str] = mapped_column("modelProviderId", ForeignKey("ModelProvider.id", ondelete="CASCADE", onupdate="CASCADE", name="ModelProviderScope_modelProviderId_fkey"))
    scope_type: Mapped[str] = mapped_column("scopeType", Enum("ORGANIZATION", "TEAM", "PROJECT", name="ModelProviderScopeType"))
    scope_id: Mapped[str] = mapped_column("scopeId")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("ModelProviderScope_scopeType_scopeId_idx", "scopeType", "scopeId"),
        Index("ModelProviderScope_modelProviderId_idx", "modelProviderId"),
        Index("ModelProviderScope_modelProviderId_scopeType_scopeId_key", "modelProviderId", "scopeType", "scopeId", unique=True),
    )

    model_provider: Mapped["ModelProvider"] = relationship(back_populates="scopes")

class ModelDefaultConfig(Base):
    __tablename__ = "ModelDefaultConfig"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    config: Mapped[dict[str, Any]] = mapped_column(JSONB)
    author_id: Mapped[Optional[str]] = mapped_column("authorId")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)
    organization_id: Mapped[str] = mapped_column("organizationId")

    __table_args__ = (
        Index("ModelDefaultConfig_organizationId_idx", "organizationId"),
    )

    scopes: Mapped[list["ModelDefaultConfigScope"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="config")

class ModelDefaultConfigScope(Base):
    __tablename__ = "ModelDefaultConfigScope"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    config_id: Mapped[str] = mapped_column("configId", ForeignKey("ModelDefaultConfig.id", ondelete="CASCADE", onupdate="CASCADE", name="ModelDefaultConfigScope_configId_fkey"))
    scope_type: Mapped[str] = mapped_column("scopeType", Enum("ORGANIZATION", "TEAM", "PROJECT", name="ModelDefaultScopeType"))
    scope_id: Mapped[str] = mapped_column("scopeId")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("ModelDefaultConfigScope_scopeType_scopeId_idx", "scopeType", "scopeId"),
        Index("ModelDefaultConfigScope_configId_idx", "configId"),
        Index("ModelDefaultConfigScope_configId_scopeType_scopeId_key", "configId", "scopeType", "scopeId", unique=True),
    )

    config: Mapped["ModelDefaultConfig"] = relationship(back_populates="scopes")

class GithubInstallation(Base):
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

    __tablename__ = "GithubInstallation"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    installation_id: Mapped[str] = mapped_column("installationId")
    """GitHub installation id (numeric; stored as string). Globally unique."""
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="GithubInstallation_organizationId_fkey"))
    account_login: Mapped[str] = mapped_column("accountLogin")
    """The GitHub account (organization or user) the App is installed on."""
    account_type: Mapped[str] = mapped_column("accountType")
    """"Organization" or "User"."""
    account_id: Mapped[str] = mapped_column("accountId")
    """GitHub numeric account id (stable across renames)."""
    repository_selection: Mapped[str] = mapped_column("repositorySelection")
    """"all" or "selected" — whether the App can see every repo on the account
    or only a chosen subset.
    """
    repositories: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))
    """Cached repository list [{ id, fullName }] when repositorySelection =
    "selected". Null when "all" (the full set is resolved live at mint time).
    """
    suspended_at: Mapped[Optional[datetime]] = mapped_column("suspendedAt")
    """Set when GitHub suspends the installation; cleared on unsuspend. A
    suspended installation cannot mint tokens.
    """
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("GithubInstallation_installationId_key", "installationId", unique=True),
        Index("GithubInstallation_organizationId_idx", "organizationId"),
    )

    organization: Mapped["Organization"] = relationship(back_populates="github_installations")

class GithubPullRequest(Base):
    """A pull request discovered through the organization's GitHub connection and
    mapped to the branch it was opened from, so a coding-agent session that
    knows only its branch can name the pull request it produced.
    Spec: specs/integrations/github-connection.feature.

    The stored title, state, draft and timestamp columns are SNAPSHOTS taken
    when the mapping was last refreshed. They are fallback labels only: the
    frontend reads live status from GitHub, so a stale row never presents
    itself as current truth.
    """

    __tablename__ = "GithubPullRequest"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="GithubPullRequest_organizationId_fkey"))
    repository_host: Mapped[str] = mapped_column("repositoryHost")
    """The GitHub host the repository lives on, e.g. "github.com"."""
    repository_full_name: Mapped[str] = mapped_column("repositoryFullName")
    """"owner/repo", stored LOWERCASED so a lookup matches whatever casing the
    session reported.
    """
    head_branch: Mapped[str] = mapped_column("headBranch")
    """The branch the pull request was opened from."""
    pr_number: Mapped[int] = mapped_column("prNumber")
    html_url: Mapped[str] = mapped_column("htmlUrl")
    title: Mapped[str]
    state: Mapped[str]
    """GitHub's own state string: "open" or "closed"."""
    is_draft: Mapped[bool] = mapped_column("isDraft", default=False, server_default=text("false"))
    author_login: Mapped[Optional[str]] = mapped_column("authorLogin")
    pr_created_at: Mapped[datetime] = mapped_column("prCreatedAt")
    pr_closed_at: Mapped[Optional[datetime]] = mapped_column("prClosedAt")
    pr_merged_at: Mapped[Optional[datetime]] = mapped_column("prMergedAt")
    pr_updated_at: Mapped[Optional[datetime]] = mapped_column("prUpdatedAt")
    """GitHub's own `updated_at` for the snapshot this row holds, and the
    freshness guard on every write to it: a snapshot older than the stored one
    is skipped. GitHub may deliver webhooks out of order, and without the
    guard a late `opened` lands after `closed` and reopens a merged pull
    request. NULL means the row predates the column, so any write is accepted.
    """
    mapped_at: Mapped[datetime] = mapped_column("mappedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    """When this row was first mapped to its branch."""
    last_checked_at: Mapped[datetime] = mapped_column("lastCheckedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    """When the snapshot columns were last refreshed from GitHub."""
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("GithubPullRequest_organizationId_repositoryHost_repositoryF_idx", "organizationId", "repositoryHost", "repositoryFullName", "headBranch"),
        Index("GithubPullRequest_organizationId_repositoryHost_repositoryF_key", "organizationId", "repositoryHost", "repositoryFullName", "prNumber", unique=True),
    )

    organization: Mapped["Organization"] = relationship(back_populates="github_pull_requests")

class GithubBranchPullRequestCheck(Base):
    """The per-branch bookkeeping behind pull-request mapping: when a branch was
    last looked up on GitHub, how many pull requests it resolved to, and when it
    may be asked about again. A branch with no pull request is the common case,
    so `notFoundAt` plus `recheckAfter` is what keeps an unmapped branch from
    costing a GitHub call on every read.
    Spec: specs/integrations/github-connection.feature.
    """

    __tablename__ = "GithubBranchPullRequestCheck"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="GithubBranchPullRequestCheck_organizationId_fkey"))
    repository_host: Mapped[str] = mapped_column("repositoryHost")
    repository_full_name: Mapped[str] = mapped_column("repositoryFullName")
    """"owner/repo", stored LOWERCASED, matching GithubPullRequest."""
    head_branch: Mapped[str] = mapped_column("headBranch")
    last_checked_at: Mapped[datetime] = mapped_column("lastCheckedAt")
    pr_count: Mapped[int] = mapped_column("prCount", default=0, server_default=text("0"))
    not_found_at: Mapped[Optional[datetime]] = mapped_column("notFoundAt")
    """Set while the branch resolves to no pull request; cleared once one is
    found.
    """
    recheck_after: Mapped[Optional[datetime]] = mapped_column("recheckAfter")
    """The earliest instant GitHub may be asked about this branch again."""
    attempts: Mapped[int] = mapped_column(default=0, server_default=text("0"))
    """Consecutive lookups that found nothing, backing off the recheck."""
    last_requested_at: Mapped[datetime] = mapped_column("lastRequestedAt")
    """When a reader last asked about this branch, whether or not GitHub was
    called.
    """
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("GithubBranchPullRequestCheck_recheckAfter_lastRequestedAt_n_idx", "recheckAfter", "lastRequestedAt", "notFoundAt"),
        Index("GithubBranchPullRequestCheck_organizationId_repositoryHost__key", "organizationId", "repositoryHost", "repositoryFullName", "headBranch", unique=True),
    )

    organization: Mapped["Organization"] = relationship(back_populates="github_branch_pr_checks")

class LangyConversationProjection(Base):
    """Rebuildable operational state folded from Langy's canonical event stream.
    The source of truth remains the ClickHouse event log.
    """

    __tablename__ = "LangyConversationProjection"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    project_id: Mapped[str] = mapped_column("projectId")
    conversation_id: Mapped[str] = mapped_column("conversationId")
    user_id: Mapped[str] = mapped_column("userId")
    """Owner; first-writer-wins from the first message's userId."""
    title: Mapped[Optional[str]]
    title_source: Mapped[str] = mapped_column("titleSource", Enum("derived", "auto", "user", name="LangyProjectionTitleSource"))
    """derived | auto | user — governs auto-title precedence (user is sticky)."""
    origin: Mapped[str] = mapped_column(default="interactive", server_default=text("'interactive'"))
    """Who started the conversation: `interactive` for the person in the panel, `run` for a
    module acting for them with nobody at the keyboard. A row folded before the column
    existed reads as interactive.
    """
    status: Mapped[str]
    is_shared: Mapped[bool] = mapped_column("isShared", default=False, server_default=text("false"))
    shared_at: Mapped[Optional[float]] = mapped_column("sharedAt")
    shared_by_id: Mapped[Optional[str]] = mapped_column("sharedById")
    message_count: Mapped[int] = mapped_column("messageCount", default=0, server_default=text("0"))
    last_activity_at: Mapped[Optional[float]] = mapped_column("lastActivityAt")
    current_turn_id: Mapped[Optional[str]] = mapped_column("currentTurnId")
    """The turn in flight, or null when idle (busy guard reads this)."""
    last_error: Mapped[Optional[str]] = mapped_column("lastError")
    last_model: Mapped[Optional[str]] = mapped_column("lastModel")
    """The model the latest accepted turn ran on; reopening seeds the picker from it."""
    pending_handoff_token: Mapped[Optional[str]] = mapped_column("pendingHandoffToken")
    """ADR-048 shutdown-handoff resume token (opaque, worker-authored)."""
    pending_handoff_turn_id: Mapped[Optional[str]] = mapped_column("pendingHandoffTurnId")
    run_token: Mapped[Optional[str]] = mapped_column("runToken")
    """Per-conversation HMAC key for worker stream frames; server-side only."""
    archived_at: Mapped[Optional[float]] = mapped_column("archivedAt")
    created_at: Mapped[float] = mapped_column("createdAt")
    updated_at: Mapped[float] = mapped_column("updatedAt")
    occurred_at: Mapped[float] = mapped_column("occurredAt")
    """Business time of the latest applied event."""
    accepted_at: Mapped[float] = mapped_column("acceptedAt")
    """Event-log acceptance time. Together with lastEventId this is the cursor."""
    last_event_id: Mapped[str] = mapped_column("lastEventId")
    projection_version: Mapped[str] = mapped_column("projectionVersion")

    __table_args__ = (
        Index("LangyConversationProjection_projectId_userId_conversationId_idx", "projectId", "userId", "conversationId"),
        Index("LangyConversationProjection_projectId_userId_lastActivityAt_idx", "projectId", "userId", "lastActivityAt"),
        Index("LangyConversationProjection_projectId_conversationId_key", "projectId", "conversationId", unique=True),
    )

class LangyTurnRequest(Base):
    """Durable idempotency receipt for one logical Langy send. A transport retry
    reuses the same requestId and receives the original conversation/turn ids
    instead of minting a second message or turn.
    """

    __tablename__ = "LangyTurnRequest"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    project_id: Mapped[str] = mapped_column("projectId")
    user_id: Mapped[str] = mapped_column("userId")
    request_id: Mapped[str] = mapped_column("requestId")
    conversation_id: Mapped[str] = mapped_column("conversationId")
    turn_id: Mapped[str] = mapped_column("turnId")
    status: Mapped[str]
    lease_owner: Mapped[str] = mapped_column("leaseOwner")
    lease_expires_at: Mapped[datetime] = mapped_column("leaseExpiresAt")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("LangyTurnRequest_projectId_conversationId_idx", "projectId", "conversationId"),
        Index("LangyTurnRequest_projectId_userId_requestId_key", "projectId", "userId", "requestId", unique=True),
    )

class LangyActiveTurn(Base):
    """Postgres-authoritative one-turn admission for a conversation. PREPARING
    claims are leased so a crash before the durable start event can recover;
    COMMITTED claims are released only by the matching terminal event.
    """

    __tablename__ = "LangyActiveTurn"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    project_id: Mapped[str] = mapped_column("projectId")
    conversation_id: Mapped[str] = mapped_column("conversationId")
    turn_id: Mapped[str] = mapped_column("turnId")
    request_id: Mapped[str] = mapped_column("requestId")
    user_id: Mapped[str] = mapped_column("userId")
    status: Mapped[str]
    lease_owner: Mapped[str] = mapped_column("leaseOwner")
    lease_expires_at: Mapped[datetime] = mapped_column("leaseExpiresAt")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("LangyActiveTurn_projectId_userId_idx", "projectId", "userId"),
        Index("LangyActiveTurn_projectId_conversationId_key", "projectId", "conversationId", unique=True),
    )

class LangyConversationTurnProjection(Base):
    """Rebuildable per-turn operational render document."""

    __tablename__ = "LangyConversationTurnProjection"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    project_id: Mapped[str] = mapped_column("projectId")
    conversation_id: Mapped[str] = mapped_column("conversationId")
    turn_id: Mapped[str] = mapped_column("turnId")
    status: Mapped[str]
    """LangyConversationTurnStatus. Deliberately TEXT, not a Postgres enum: this
    is a PROJECTION — one writer (the fold), rebuildable from the event log,
    and nothing queries or indexes on it. An enum bought no selectivity here,
    only typo protection, and charged a migration for every new status that
    Postgres then could not remove. The one definition lives in
    `LANGY_CONVERSATION_TURN_STATUS`; the repository validates against it.
    """
    question_parts: Mapped[dict[str, Any]] = mapped_column("questionParts", JSONB)
    """LangyMessagePart[] — the user question that opened the turn."""
    answer_parts: Mapped[dict[str, Any]] = mapped_column("answerParts", JSONB)
    """LangyMessagePart[] — the agent's whole final answer."""
    tool_calls: Mapped[dict[str, Any]] = mapped_column("toolCalls", JSONB)
    """LangyTurnToolCall[] — tool-call lifecycle audit in initiation order."""
    plan: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))
    """LangyPlanItemData[] | null — null means the turn never kept a plan."""
    error: Mapped[Optional[str]]
    started_at: Mapped[Optional[float]] = mapped_column("startedAt")
    ended_at: Mapped[Optional[float]] = mapped_column("endedAt")
    created_at: Mapped[float] = mapped_column("createdAt")
    updated_at: Mapped[float] = mapped_column("updatedAt")
    occurred_at: Mapped[float] = mapped_column("occurredAt")
    accepted_at: Mapped[float] = mapped_column("acceptedAt")
    last_event_id: Mapped[str] = mapped_column("lastEventId")
    projection_version: Mapped[str] = mapped_column("projectionVersion")

    __table_args__ = (
        Index("LangyConversationTurnProjection_projectId_conversationId_up_idx", "projectId", "conversationId", "updatedAt"),
        Index("LangyConversationTurnProjection_projectId_conversationId_tu_key", "projectId", "conversationId", "turnId", unique=True),
    )

class LangyMessageProjection(Base):
    """Rebuildable operational message fact. One source event produces one row."""

    __tablename__ = "LangyMessageProjection"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    project_id: Mapped[str] = mapped_column("projectId")
    conversation_id: Mapped[str] = mapped_column("conversationId")
    message_id: Mapped[str] = mapped_column("messageId")
    role: Mapped[str]
    parts: Mapped[dict[str, Any]] = mapped_column(JSONB)
    source_event_id: Mapped[str] = mapped_column("sourceEventId")
    occurred_at: Mapped[float] = mapped_column("occurredAt")
    accepted_at: Mapped[float] = mapped_column("acceptedAt")
    created_at: Mapped[float] = mapped_column("createdAt")
    updated_at: Mapped[float] = mapped_column("updatedAt")

    __table_args__ = (
        Index("LangyMessageProjection_projectId_conversationId_createdAt_idx", "projectId", "conversationId", "createdAt"),
        Index("LangyMessageProjection_projectId_conversationId_messageId_key", "projectId", "conversationId", "messageId", unique=True),
        Index("LangyMessageProjection_projectId_sourceEventId_key", "projectId", "sourceEventId", unique=True),
    )

class ProjectSecret(Base):
    __tablename__ = "ProjectSecret"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    project_id: Mapped[str] = mapped_column("projectId", ForeignKey("Project.id", ondelete="RESTRICT", onupdate="CASCADE", name="ProjectSecret_projectId_fkey"))
    name: Mapped[str]
    encrypted_value: Mapped[str] = mapped_column("encryptedValue")
    created_by_id: Mapped[str] = mapped_column("createdById", ForeignKey("User.id", ondelete="RESTRICT", onupdate="CASCADE", name="ProjectSecret_createdById_fkey"))
    updated_by_id: Mapped[str] = mapped_column("updatedById", ForeignKey("User.id", ondelete="RESTRICT", onupdate="CASCADE", name="ProjectSecret_updatedById_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("ProjectSecret_projectId_idx", "projectId"),
        Index("ProjectSecret_projectId_name_key", "projectId", "name", unique=True),
    )

    project: Mapped["Project"] = relationship(back_populates="project_secrets")
    created_by: Mapped["User"] = relationship(foreign_keys=[created_by_id], back_populates="secrets_created")
    updated_by: Mapped["User"] = relationship(foreign_keys=[updated_by_id], back_populates="secrets_updated")

class TriggerSent(Base):
    __tablename__ = "TriggerSent"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    trigger_id: Mapped[str] = mapped_column("triggerId", ForeignKey("Trigger.id", ondelete="RESTRICT", onupdate="CASCADE", name="TriggerSent_triggerId_fkey"))
    trace_id: Mapped[Optional[str]] = mapped_column("traceId")
    custom_graph_id: Mapped[Optional[str]] = mapped_column("customGraphId", ForeignKey("CustomGraph.id", ondelete="CASCADE", onupdate="CASCADE", name="TriggerSent_customGraphId_fkey"))
    project_id: Mapped[str] = mapped_column("projectId", ForeignKey("Project.id", ondelete="RESTRICT", onupdate="CASCADE", name="TriggerSent_projectId_fkey"))
    resolved_at: Mapped[Optional[datetime]] = mapped_column("resolvedAt")
    open_incident_key: Mapped[Optional[str]] = mapped_column("openIncidentKey")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("TriggerSent_openIncidentKey_key", "openIncidentKey", unique=True),
        Index("TriggerSent_triggerId_idx", "triggerId"),
        Index("TriggerSent_projectId_idx", "projectId"),
        Index("TriggerSent_customGraphId_idx", "customGraphId"),
        Index("TriggerSent_resolvedAt_idx", "resolvedAt"),
        Index("TriggerSent_projectId_triggerId_createdAt_id_idx", "projectId", "triggerId", "createdAt", "id"),
        Index("TriggerSent_triggerId_traceId_key", "triggerId", "traceId", unique=True),
    )

    custom_graph: Mapped[Optional["CustomGraph"]] = relationship(back_populates="trigger_sent")
    project: Mapped["Project"] = relationship(back_populates="trigger_sent")
    trigger: Mapped["Trigger"] = relationship(back_populates="trigger_sent")

class TriggerLatestEvaluation(Base):
    __tablename__ = "TriggerLatestEvaluation"

    trigger_id: Mapped[str] = mapped_column("triggerId", ForeignKey("Trigger.id", ondelete="CASCADE", onupdate="CASCADE", name="TriggerLatestEvaluation_triggerId_fkey"), primary_key=True)
    project_id: Mapped[str] = mapped_column("projectId")
    evaluated_at: Mapped[datetime] = mapped_column("evaluatedAt", UtcDateTimeTz(precision=3))
    verdict: Mapped[str]
    observed_value: Mapped[Optional[float]] = mapped_column("observedValue")
    threshold: Mapped[Optional[float]]
    operator: Mapped[Optional[str]]
    time_period_minutes: Mapped[Optional[int]] = mapped_column("timePeriodMinutes")
    skip_code: Mapped[Optional[str]] = mapped_column("skipCode")
    updated_at: Mapped[datetime] = mapped_column("updatedAt", UtcDateTimeTz(precision=3), default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("TriggerLatestEvaluation_projectId_idx", "projectId"),
    )

    trigger: Mapped["Trigger"] = relationship(back_populates="latest_evaluation")

class EmailSuppression(Base):
    __tablename__ = "EmailSuppression"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    project_id: Mapped[str] = mapped_column("projectId", ForeignKey("Project.id", ondelete="RESTRICT", onupdate="CASCADE", name="EmailSuppression_projectId_fkey"))
    email: Mapped[str]
    trigger_id: Mapped[Optional[str]] = mapped_column("triggerId")
    reason: Mapped[str] = mapped_column(default="unsubscribe", server_default=text("'unsubscribe'"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("EmailSuppression_projectId_idx", "projectId"),
        Index("EmailSuppression_projectId_email_triggerId_key", "projectId", "email", "triggerId", unique=True),
    )

    project: Mapped["Project"] = relationship(back_populates="email_suppressions")

class AnnotationScore(Base):
    __tablename__ = "AnnotationScore"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    name: Mapped[str]
    project_id: Mapped[str] = mapped_column("projectId", ForeignKey("Project.id", ondelete="RESTRICT", onupdate="CASCADE", name="AnnotationScore_projectId_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    deleted_at: Mapped[Optional[datetime]] = mapped_column("deletedAt")
    description: Mapped[Optional[str]]
    active: Mapped[bool] = mapped_column(default=True, server_default=text("true"))
    data_type: Mapped[Optional[str]] = mapped_column("dataType", Enum("CATEGORICAL", "BOOLEAN", "LIKERT", "OPTION", "CHECKBOX", name="AnnotationScoreDataType"))
    options: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))
    default_value: Mapped[Optional[dict[str, Any]]] = mapped_column("defaultValue", JSONB(none_as_null=True))
    global_: Mapped[bool] = mapped_column("global", default=False, server_default=text("false"))

    __table_args__ = (
        Index("AnnotationScore_projectId_idx", "projectId"),
    )

    project: Mapped["Project"] = relationship(back_populates="annotation_scores")
    annotation_queue_scores: Mapped[list["AnnotationQueueScores"]] = relationship(passive_deletes="all", back_populates="annotation_score")

class AnnotationQueue(Base):
    __tablename__ = "AnnotationQueue"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    name: Mapped[str]
    slug: Mapped[str]
    project_id: Mapped[str] = mapped_column("projectId", ForeignKey("Project.id", ondelete="RESTRICT", onupdate="CASCADE", name="AnnotationQueue_projectId_fkey"))
    description: Mapped[Optional[str]]
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("AnnotationQueue_projectId_idx", "projectId"),
        Index("AnnotationQueue_projectId_slug_key", "projectId", "slug", unique=True),
    )

    project: Mapped["Project"] = relationship(back_populates="annotation_queue")
    members: Mapped[list["AnnotationQueueMembers"]] = relationship(passive_deletes="all", back_populates="annotation_queue")
    annotation_queue_scores: Mapped[list["AnnotationQueueScores"]] = relationship(passive_deletes="all", back_populates="annotation_queue")
    annotation_queue_items: Mapped[list["AnnotationQueueItem"]] = relationship(back_populates="annotation_queue")

class AnnotationQueueMembers(Base):
    __tablename__ = "AnnotationQueueMembers"

    annotation_queue_id: Mapped[str] = mapped_column("annotationQueueId", ForeignKey("AnnotationQueue.id", ondelete="RESTRICT", onupdate="CASCADE", name="AnnotationQueueMembers_annotationQueueId_fkey"), primary_key=True)
    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="RESTRICT", onupdate="CASCADE", name="AnnotationQueueMembers_userId_fkey"), primary_key=True)

    __table_args__ = (
        Index("AnnotationQueueMembers_annotationQueueId_idx", "annotationQueueId"),
        Index("AnnotationQueueMembers_userId_idx", "userId"),
    )

    user: Mapped["User"] = relationship(back_populates="annotation_queues")
    annotation_queue: Mapped["AnnotationQueue"] = relationship(back_populates="members")

class AnnotationQueueScores(Base):
    __tablename__ = "AnnotationQueueScores"

    annotation_queue_id: Mapped[str] = mapped_column("annotationQueueId", ForeignKey("AnnotationQueue.id", ondelete="RESTRICT", onupdate="CASCADE", name="AnnotationQueueScores_annotationQueueId_fkey"), primary_key=True)
    annotation_score_id: Mapped[str] = mapped_column("annotationScoreId", ForeignKey("AnnotationScore.id", ondelete="RESTRICT", onupdate="CASCADE", name="AnnotationQueueScores_annotationScoreId_fkey"), primary_key=True)

    __table_args__ = (
        Index("AnnotationQueueScores_annotationQueueId_idx", "annotationQueueId"),
        Index("AnnotationQueueScores_annotationScoreId_idx", "annotationScoreId"),
    )

    annotation_score: Mapped["AnnotationScore"] = relationship(back_populates="annotation_queue_scores")
    annotation_queue: Mapped["AnnotationQueue"] = relationship(back_populates="annotation_queue_scores")

class AnnotationQueueItem(Base):
    __tablename__ = "AnnotationQueueItem"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    annotation_queue_id: Mapped[Optional[str]] = mapped_column("annotationQueueId", ForeignKey("AnnotationQueue.id", ondelete="SET NULL", onupdate="CASCADE", name="AnnotationQueueItem_annotationQueueId_fkey"))
    user_id: Mapped[Optional[str]] = mapped_column("userId", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="AnnotationQueueItem_userId_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    created_by_user_id: Mapped[Optional[str]] = mapped_column("createdByUserId", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="AnnotationQueueItem_createdByUserId_fkey"))
    trace_id: Mapped[str] = mapped_column("traceId")
    project_id: Mapped[str] = mapped_column("projectId", ForeignKey("Project.id", ondelete="RESTRICT", onupdate="CASCADE", name="AnnotationQueueItem_projectId_fkey"))
    done_at: Mapped[Optional[datetime]] = mapped_column("doneAt")
    marked_for_dataset_at: Mapped[Optional[datetime]] = mapped_column("markedForDatasetAt")
    """Set while the annotator wants this item's trace added to a dataset once
    the queue walk is finished. Persisted rather than kept in the browser so
    the selection survives a refresh and stays visible to teammates.
    """

    __table_args__ = (
        Index("AnnotationQueueItem_annotationQueueId_idx", "annotationQueueId"),
        Index("AnnotationQueueItem_userId_idx", "userId"),
        Index("AnnotationQueueItem_projectId_idx", "projectId"),
        Index("AnnotationQueueItem_projectId_doneAt_idx", "projectId", "doneAt"),
        Index("AnnotationQueueItem_projectId_userId_doneAt_idx", "projectId", "userId", "doneAt"),
        Index("AnnotationQueueItem_annotationQueueId_doneAt_idx", "annotationQueueId", "doneAt"),
        Index("AnnotationQueueItem_traceId_userId_projectId_key", "traceId", "userId", "projectId", unique=True),
        Index("AnnotationQueueItem_traceId_annotationQueueId_projectId_key", "traceId", "annotationQueueId", "projectId", unique=True),
    )

    annotation_queue: Mapped[Optional["AnnotationQueue"]] = relationship(back_populates="annotation_queue_items")
    user: Mapped[Optional["User"]] = relationship(foreign_keys=[user_id], back_populates="assigned_queue_items")
    created_by_user: Mapped[Optional["User"]] = relationship(foreign_keys=[created_by_user_id], back_populates="created_queue_items")
    project: Mapped["Project"] = relationship(back_populates="annotation_queue_item")

class ShareLink(Base):
    __tablename__ = "ShareLink"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    token: Mapped[str]
    resource_type: Mapped[str] = mapped_column("resourceType", Enum("TRACE", "THREAD", name="ShareResourceType"))
    resource_id: Mapped[str] = mapped_column("resourceId")
    thread_id: Mapped[Optional[str]] = mapped_column("threadId")
    project_id: Mapped[str] = mapped_column("projectId", ForeignKey("Project.id", ondelete="RESTRICT", onupdate="CASCADE", name="ShareLink_projectId_fkey"))
    user_id: Mapped[Optional[str]] = mapped_column("userId", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="ShareLink_userId_fkey"))
    visibility: Mapped[str] = mapped_column(Enum("PUBLIC", "ORGANIZATION", "PROJECT", name="ShareVisibility"), default="PUBLIC", server_default=text("'PUBLIC'"))
    expires_at: Mapped[Optional[datetime]] = mapped_column("expiresAt")
    max_views: Mapped[Optional[int]] = mapped_column("maxViews")
    view_count: Mapped[int] = mapped_column("viewCount", default=0, server_default=text("0"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("ShareLink_token_key", "token", unique=True),
        Index("ShareLink_projectId_resourceType_resourceId_idx", "projectId", "resourceType", "resourceId"),
        Index("ShareLink_userId_idx", "userId"),
    )

    project: Mapped["Project"] = relationship(back_populates="share_links")
    user: Mapped[Optional["User"]] = relationship(back_populates="share_links")

class PinnedTrace(Base):
    __tablename__ = "PinnedTrace"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    project_id: Mapped[str] = mapped_column("projectId", ForeignKey("Project.id", ondelete="RESTRICT", onupdate="CASCADE", name="PinnedTrace_projectId_fkey"))
    trace_id: Mapped[str] = mapped_column("traceId")
    user_id: Mapped[Optional[str]] = mapped_column("userId", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="PinnedTrace_userId_fkey"))
    source: Mapped[str] = mapped_column(Enum("manual", "share", name="PinSource"), default="manual", server_default=text("'manual'"))
    reason: Mapped[Optional[str]]
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("PinnedTrace_projectId_idx", "projectId"),
        Index("PinnedTrace_projectId_traceId_key", "projectId", "traceId", unique=True),
    )

    project: Mapped["Project"] = relationship(back_populates="pinned_traces")
    user: Mapped[Optional["User"]] = relationship(back_populates="pinned_traces")

class TraceEditOverlay(Base):
    __tablename__ = "TraceEditOverlay"

    id: Mapped[str] = mapped_column(primary_key=True)
    project_id: Mapped[str] = mapped_column("projectId", ForeignKey("Project.id", ondelete="RESTRICT", onupdate="CASCADE", name="TraceEditOverlay_projectId_fkey"))
    trace_id: Mapped[str] = mapped_column("traceId")
    patch: Mapped[dict[str, Any]] = mapped_column(JSONB)
    created_by_id: Mapped[Optional[str]] = mapped_column("createdById", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="TraceEditOverlay_createdById_fkey"))
    updated_by_id: Mapped[Optional[str]] = mapped_column("updatedById", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="TraceEditOverlay_updatedById_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("TraceEditOverlay_projectId_idx", "projectId"),
        Index("TraceEditOverlay_projectId_traceId_key", "projectId", "traceId", unique=True),
    )

    project: Mapped["Project"] = relationship(back_populates="trace_edit_overlays")
    created_by: Mapped[Optional["User"]] = relationship(foreign_keys=[created_by_id], back_populates="trace_edit_overlays_created")
    updated_by: Mapped[Optional["User"]] = relationship(foreign_keys=[updated_by_id], back_populates="trace_edit_overlays_updated")

class TraceIngestSourceBilling(Base):
    __tablename__ = "TraceIngestSourceBilling"

    organization_id: Mapped[str] = mapped_column("organizationId", primary_key=True)
    source_type: Mapped[str] = mapped_column("sourceType", primary_key=True)
    billed: Mapped[bool]
    recorded_at: Mapped[datetime] = mapped_column("recordedAt")
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

class RetentionPolicy(Base):
    __tablename__ = "RetentionPolicy"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId")
    scope_type: Mapped[str] = mapped_column("scopeType", Enum("ORGANIZATION", "TEAM", "PROJECT", name="RetentionPolicyScopeType"))
    scope_id: Mapped[str] = mapped_column("scopeId")
    category: Mapped[str]
    retention_days: Mapped[int] = mapped_column("retentionDays")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("RetentionPolicy_organizationId_idx", "organizationId"),
        Index("RetentionPolicy_scopeType_scopeId_idx", "scopeType", "scopeId"),
        Index("RetentionPolicy_scopeType_scopeId_category_key", "scopeType", "scopeId", "category", unique=True),
    )

class DataPrivacyPolicy(Base):
    __tablename__ = "DataPrivacyPolicy"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId")
    scope_type: Mapped[str] = mapped_column("scopeType", Enum("ORGANIZATION", "DEPARTMENT", "TEAM", "PROJECT", name="DataPrivacyScopeType"))
    scope_id: Mapped[str] = mapped_column("scopeId")
    personal_only: Mapped[bool] = mapped_column("personalOnly", default=False, server_default=text("false"))
    config: Mapped[dict[str, Any]] = mapped_column(JSONB)
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("DataPrivacyPolicy_organizationId_idx", "organizationId"),
        Index("DataPrivacyPolicy_scopeType_scopeId_idx", "scopeType", "scopeId"),
        Index("DataPrivacyPolicy_scopeType_scopeId_personalOnly_key", "scopeType", "scopeId", "personalOnly", unique=True),
    )

class InstantEvalJudgeProject(Base):
    __tablename__ = "InstantEvalJudgeProject"

    project_id: Mapped[str] = mapped_column("projectId", primary_key=True)
    organization_id: Mapped[str] = mapped_column("organizationId")
    created_at: Mapped[datetime] = mapped_column("createdAt")

class InstantEvalJudgeUsageBilling(Base):
    __tablename__ = "InstantEvalJudgeUsageBilling"

    organization_id: Mapped[str] = mapped_column("organizationId", primary_key=True)
    usage_billed: Mapped[bool] = mapped_column("usageBilled")
    occurred_at: Mapped[datetime] = mapped_column("occurredAt")
    from_catch_up: Mapped[bool] = mapped_column("fromCatchUp")

class InstantEvalJudgeSpend(Base):
    __tablename__ = "InstantEvalJudgeSpend"

    organization_id: Mapped[str] = mapped_column("organizationId", primary_key=True)
    request_id: Mapped[str] = mapped_column("requestId", primary_key=True)
    spend_nano_usd: Mapped[int] = mapped_column("spendNanoUsd", BigInteger)
    occurred_at: Mapped[datetime] = mapped_column("occurredAt")

class CustomLLMModelCost(Base):
    __tablename__ = "CustomLLMModelCost"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId")
    scope_type: Mapped[str] = mapped_column("scopeType", Enum("ORGANIZATION", "TEAM", "PROJECT", name="CustomLLMModelCostScopeType"), default="PROJECT", server_default=text("'PROJECT'"))
    scope_id: Mapped[str] = mapped_column("scopeId")
    project_id: Mapped[Optional[str]] = mapped_column("projectId")
    model: Mapped[str]
    regex: Mapped[str]
    input_cost_per_token: Mapped[Optional[float]] = mapped_column("inputCostPerToken")
    output_cost_per_token: Mapped[Optional[float]] = mapped_column("outputCostPerToken")
    cache_read_cost_per_token: Mapped[Optional[float]] = mapped_column("cacheReadCostPerToken")
    cache_creation_cost_per_token: Mapped[Optional[float]] = mapped_column("cacheCreationCostPerToken")
    cache_creation1h_cost_per_token: Mapped[Optional[float]] = mapped_column("cacheCreation1hCostPerToken")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("CustomLLMModelCost_organizationId_idx", "organizationId"),
        Index("CustomLLMModelCost_scopeType_scopeId_idx", "scopeType", "scopeId"),
        Index("CustomLLMModelCost_projectId_idx", "projectId"),
    )

class Workflow(Base):
    __tablename__ = "Workflow"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    project_id: Mapped[str] = mapped_column("projectId", ForeignKey("Project.id", ondelete="RESTRICT", onupdate="CASCADE", name="Workflow_projectId_fkey"))
    name: Mapped[str]
    icon: Mapped[str]
    description: Mapped[str]
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    latest_version_id: Mapped[Optional[str]] = mapped_column("latestVersionId", ForeignKey("WorkflowVersion.id", ondelete="RESTRICT", onupdate="RESTRICT", name="Workflow_latestVersionId_fkey"))
    current_version_id: Mapped[Optional[str]] = mapped_column("currentVersionId", ForeignKey("WorkflowVersion.id", ondelete="RESTRICT", onupdate="RESTRICT", name="Workflow_currentVersionId_fkey"))
    published_id: Mapped[Optional[str]] = mapped_column("publishedId")
    published_by_id: Mapped[Optional[str]] = mapped_column("publishedById", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="Workflow_publishedById_fkey"))
    archived_at: Mapped[Optional[datetime]] = mapped_column("archivedAt")
    is_evaluator: Mapped[bool] = mapped_column("isEvaluator", default=False, server_default=text("false"))
    is_component: Mapped[bool] = mapped_column("isComponent", default=False, server_default=text("false"))
    copied_from_workflow_id: Mapped[Optional[str]] = mapped_column("copiedFromWorkflowId", ForeignKey("Workflow.id", ondelete="RESTRICT", onupdate="RESTRICT", name="Workflow_copiedFromWorkflowId_fkey"))

    __table_args__ = (
        Index("Workflow_projectId_idx", "projectId"),
        Index("Workflow_publishedById_idx", "publishedById"),
        Index("Workflow_latestVersionId_idx", "latestVersionId"),
        Index("Workflow_currentVersionId_idx", "currentVersionId"),
        Index("Workflow_copiedFromWorkflowId_idx", "copiedFromWorkflowId"),
    )

    project: Mapped["Project"] = relationship(back_populates="workflows")
    latest_version: Mapped[Optional["WorkflowVersion"]] = relationship(foreign_keys=[latest_version_id], back_populates="workflow_as_latest_version")
    current_version: Mapped[Optional["WorkflowVersion"]] = relationship(foreign_keys=[current_version_id], back_populates="workflow_as_current_version")
    published_by: Mapped[Optional["User"]] = relationship(back_populates="workflow")
    copied_from: Mapped[Optional["Workflow"]] = relationship(remote_side=[id], back_populates="copied_workflows")
    versions: Mapped[list["WorkflowVersion"]] = relationship(foreign_keys="WorkflowVersion.workflow_id", passive_deletes="all", back_populates="workflow")
    experiments: Mapped[list["Experiment"]] = relationship(back_populates="workflow")
    copied_workflows: Mapped[list["Workflow"]] = relationship(passive_deletes="all", back_populates="copied_from")
    agents: Mapped[list["Agent"]] = relationship(back_populates="workflow")
    evaluators: Mapped[list["Evaluator"]] = relationship(back_populates="workflow")

class WorkflowVersion(Base):
    __tablename__ = "WorkflowVersion"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    version: Mapped[str]
    commit_message: Mapped[str] = mapped_column("commitMessage")
    author_id: Mapped[str] = mapped_column("authorId", ForeignKey("User.id", ondelete="RESTRICT", onupdate="CASCADE", name="WorkflowVersion_authorId_fkey"))
    project_id: Mapped[str] = mapped_column("projectId", ForeignKey("Project.id", ondelete="RESTRICT", onupdate="CASCADE", name="WorkflowVersion_projectId_fkey"))
    workflow_id: Mapped[str] = mapped_column("workflowId", ForeignKey("Workflow.id", ondelete="RESTRICT", onupdate="CASCADE", name="WorkflowVersion_workflowId_fkey"))
    parent_id: Mapped[Optional[str]] = mapped_column("parentId", ForeignKey("WorkflowVersion.id", ondelete="RESTRICT", onupdate="RESTRICT", name="WorkflowVersion_parentId_fkey"))
    auto_saved: Mapped[bool] = mapped_column("autoSaved", default=False, server_default=text("false"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    dsl: Mapped[dict[str, Any]] = mapped_column(JSONB)

    __table_args__ = (
        Index("WorkflowVersion_workflowId_idx", "workflowId"),
        Index("WorkflowVersion_projectId_idx", "projectId"),
        Index("WorkflowVersion_parentId_idx", "parentId"),
        Index("WorkflowVersion_authorId_idx", "authorId"),
        Index("WorkflowVersion_createdAt_idx", "createdAt"),
        Index("WorkflowVersion_workflowId_version_key", "workflowId", "version", unique=True),
    )

    author: Mapped["User"] = relationship(back_populates="workflow_version")
    project: Mapped["Project"] = relationship(back_populates="workflow_version")
    workflow: Mapped["Workflow"] = relationship(foreign_keys=[workflow_id], back_populates="versions")
    parent: Mapped[Optional["WorkflowVersion"]] = relationship(remote_side=[id], back_populates="children")
    children: Mapped[list["WorkflowVersion"]] = relationship(passive_deletes="all", back_populates="parent")
    workflow_as_latest_version: Mapped[list["Workflow"]] = relationship(foreign_keys="Workflow.latest_version_id", passive_deletes="all", back_populates="latest_version")
    workflow_as_current_version: Mapped[list["Workflow"]] = relationship(foreign_keys="Workflow.current_version_id", passive_deletes="all", back_populates="current_version")

class AuditLog(Base):
    __tablename__ = "AuditLog"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    user_id: Mapped[Optional[str]] = mapped_column("userId")
    project_id: Mapped[Optional[str]] = mapped_column("projectId")
    organization_id: Mapped[Optional[str]] = mapped_column("organizationId")
    action: Mapped[str]
    args: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))
    error: Mapped[Optional[str]]
    ip_address: Mapped[Optional[str]] = mapped_column("ipAddress")
    user_agent: Mapped[Optional[str]] = mapped_column("userAgent")
    metadata_: Mapped[Optional[dict[str, Any]]] = mapped_column("metadata", JSONB(none_as_null=True))
    actor_user_id: Mapped[Optional[str]] = mapped_column("actorUserId")
    """WHO REALLY DID IT, when that is not `userId`.

    Under an impersonation the session's `user.id` is the SUBJECT — the
    customer whose access an operator is borrowing — so a row naming only
    `userId` filed the operator's act against the customer. An operator
    granting a permanent way into somebody's organization appeared, in the
    only durable record of it, to be that somebody.

    Null on an ordinary request, where the two would be the same person and
    saying so twice adds nothing. ADR-118 / D06: every decision names both.
    """
    target_kind: Mapped[Optional[str]] = mapped_column("targetKind")
    target_id: Mapped[Optional[str]] = mapped_column("targetId")
    before: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))
    after: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))
    idempotency_key: Mapped[Optional[str]] = mapped_column("idempotencyKey")
    """A producer's outbox key: a redelivered audit intent finds its row here (Alex, audit R2)."""

    __table_args__ = (
        Index("AuditLog_idempotencyKey_key", "idempotencyKey", unique=True),
        Index("AuditLog_createdAt_idx", "createdAt"),
        Index("AuditLog_userId_idx", "userId"),
        Index("AuditLog_actorUserId_idx", "actorUserId"),
        Index("AuditLog_projectId_idx", "projectId"),
        Index("AuditLog_action_idx", "action"),
        Index("AuditLog_ipAddress_idx", "ipAddress"),
        Index("AuditLog_organizationId_createdAt_idx", "organizationId", "createdAt"),
        Index("AuditLog_targetKind_targetId_idx", "targetKind", "targetId"),
    )

class LlmPromptConfig(Base):
    __tablename__ = "LlmPromptConfig"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    handle: Mapped[Optional[str]]
    name: Mapped[str]
    project_id: Mapped[str] = mapped_column("projectId", ForeignKey("Project.id", ondelete="RESTRICT", onupdate="CASCADE", name="LlmPromptConfig_projectId_fkey"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    deleted_at: Mapped[Optional[datetime]] = mapped_column("deletedAt")
    organization_id: Mapped[str] = mapped_column("organizationId")
    scope: Mapped[str] = mapped_column(Enum("ORGANIZATION", "PROJECT", name="PromptScope"), default="PROJECT", server_default=text("'PROJECT'"))
    copied_from_prompt_id: Mapped[Optional[str]] = mapped_column("copiedFromPromptId", ForeignKey("LlmPromptConfig.id", ondelete="RESTRICT", onupdate="RESTRICT", name="LlmPromptConfig_copiedFromPromptId_fkey"))

    __table_args__ = (
        Index("LlmPromptConfig_projectId_idx", "projectId"),
        Index("LlmPromptConfig_organizationId_scope_idx", "organizationId", "scope"),
        Index("LlmPromptConfig_copiedFromPromptId_idx", "copiedFromPromptId"),
        Index("LlmPromptConfig_handle_key", "handle", unique=True),
    )

    project: Mapped["Project"] = relationship(back_populates="llm_prompt_configs")
    copied_from: Mapped[Optional["LlmPromptConfig"]] = relationship(remote_side=[id], back_populates="copied_prompts")
    copied_prompts: Mapped[list["LlmPromptConfig"]] = relationship(passive_deletes="all", back_populates="copied_from")
    versions: Mapped[list["LlmPromptConfigVersion"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="config")
    version_tags: Mapped[list["PromptTagAssignment"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="config")

class LlmPromptConfigVersion(Base):
    __tablename__ = "LlmPromptConfigVersion"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    version: Mapped[int]
    commit_message: Mapped[Optional[str]] = mapped_column("commitMessage")
    author_id: Mapped[Optional[str]] = mapped_column("authorId", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="LlmPromptConfigVersion_authorId_fkey"))
    config_id: Mapped[str] = mapped_column("configId", ForeignKey("LlmPromptConfig.id", ondelete="CASCADE", onupdate="CASCADE", name="LlmPromptConfigVersion_configId_fkey"))
    config_data: Mapped[dict[str, Any]] = mapped_column("config", JSONB)
    schema_version: Mapped[str] = mapped_column("schemaVersion")
    runtime_parameters: Mapped[dict[str, Any]] = mapped_column("runtimeParameters", JSONB, default=lambda: {}, server_default=text("'{}'"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    project_id: Mapped[str] = mapped_column("projectId")

    __table_args__ = (
        Index("LlmPromptConfigVersion_configId_idx", "configId"),
        Index("LlmPromptConfigVersion_authorId_idx", "authorId"),
        Index("LlmPromptConfigVersion_createdAt_idx", "createdAt"),
        Index("LlmPromptConfigVersion_configId_version_key", "configId", "version", unique=True),
    )

    author: Mapped[Optional["User"]] = relationship(back_populates="llm_prompt_config_versions")
    config: Mapped["LlmPromptConfig"] = relationship(back_populates="versions")
    version_tags: Mapped[list["PromptTagAssignment"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="version")

class PromptTagAssignment(Base):
    __tablename__ = "PromptTagAssignment"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    config_id: Mapped[str] = mapped_column("configId", ForeignKey("LlmPromptConfig.id", ondelete="CASCADE", onupdate="CASCADE", name="PromptTagAssignment_configId_fkey"))
    version_id: Mapped[str] = mapped_column("versionId", ForeignKey("LlmPromptConfigVersion.id", ondelete="CASCADE", onupdate="CASCADE", name="PromptTagAssignment_versionId_fkey"))
    tag_id: Mapped[str] = mapped_column("tagId", ForeignKey("PromptTag.id", ondelete="CASCADE", onupdate="CASCADE", name="PromptTagAssignment_tagId_fkey"))
    project_id: Mapped[str] = mapped_column("projectId")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    created_by_id: Mapped[Optional[str]] = mapped_column("createdById", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="PromptTagAssignment_createdById_fkey"))
    updated_by_id: Mapped[Optional[str]] = mapped_column("updatedById", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="PromptTagAssignment_updatedById_fkey"))

    __table_args__ = (
        Index("PromptTagAssignment_configId_idx", "configId"),
        Index("PromptTagAssignment_versionId_idx", "versionId"),
        Index("PromptTagAssignment_projectId_idx", "projectId"),
        Index("PromptTagAssignment_tagId_idx", "tagId"),
        Index("PromptTagAssignment_configId_tagId_key", "configId", "tagId", unique=True),
    )

    config: Mapped["LlmPromptConfig"] = relationship(back_populates="version_tags")
    version: Mapped["LlmPromptConfigVersion"] = relationship(back_populates="version_tags")
    prompt_tag: Mapped["PromptTag"] = relationship(back_populates="assignments")
    created_by: Mapped[Optional["User"]] = relationship(foreign_keys=[created_by_id], back_populates="version_tags_created")
    updated_by: Mapped[Optional["User"]] = relationship(foreign_keys=[updated_by_id], back_populates="version_tags_updated")

class PromptTag(Base):
    __tablename__ = "PromptTag"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="PromptTag_organizationId_fkey"))
    name: Mapped[str]
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    created_by_id: Mapped[Optional[str]] = mapped_column("createdById", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="PromptTag_createdById_fkey"))
    updated_by_id: Mapped[Optional[str]] = mapped_column("updatedById", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="PromptTag_updatedById_fkey"))

    __table_args__ = (
        Index("PromptTag_organizationId_idx", "organizationId"),
        Index("PromptTag_organizationId_name_key", "organizationId", "name", unique=True),
    )

    organization: Mapped["Organization"] = relationship(back_populates="prompt_tags")
    created_by: Mapped[Optional["User"]] = relationship(foreign_keys=[created_by_id], back_populates="prompt_tags_created")
    updated_by: Mapped[Optional["User"]] = relationship(foreign_keys=[updated_by_id], back_populates="prompt_tags_updated")
    assignments: Mapped[list["PromptTagAssignment"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="prompt_tag")

class Analytics(Base):
    __tablename__ = "Analytics"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    project_id: Mapped[str] = mapped_column("projectId", ForeignKey("Project.id", ondelete="RESTRICT", onupdate="CASCADE", name="Analytics_projectId_fkey"))
    key: Mapped[str]
    value: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))
    numeric_value: Mapped[Optional[float]] = mapped_column("numericValue")
    string_value: Mapped[Optional[str]] = mapped_column("stringValue")
    bool_value: Mapped[Optional[bool]] = mapped_column("boolValue")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("Analytics_createdAt_key_projectId_numericValue_idx", "createdAt", "key", "projectId", "numericValue"),
        Index("Analytics_projectId_idx", "projectId"),
    )

    project: Mapped["Project"] = relationship(back_populates="analytics")

class CustomRole(Base):
    __tablename__ = "CustomRole"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="RESTRICT", onupdate="CASCADE", name="CustomRole_organizationId_fkey"))
    name: Mapped[str]
    description: Mapped[Optional[str]]
    permissions: Mapped[dict[str, Any]] = mapped_column(JSONB)
    kind: Mapped[str] = mapped_column(default="custom", server_default=text("'custom'"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("CustomRole_organizationId_idx", "organizationId"),
        Index("CustomRole_organizationId_name_key", "organizationId", "name", unique=True),
    )

    organization: Mapped["Organization"] = relationship(back_populates="custom_roles")
    assigned_users: Mapped[list["TeamUser"]] = relationship(back_populates="assigned_role")
    role_bindings: Mapped[list["RoleBinding"]] = relationship(back_populates="custom_role")

class Notification(Base):
    __tablename__ = "Notification"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[Optional[str]] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="Notification_organizationId_fkey"))
    project_id: Mapped[Optional[str]] = mapped_column("projectId", ForeignKey("Project.id", ondelete="CASCADE", onupdate="CASCADE", name="Notification_projectId_fkey"))
    metadata_: Mapped[dict[str, Any]] = mapped_column("metadata", JSONB)
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    sent_at: Mapped[datetime] = mapped_column("sentAt")

    __table_args__ = (
        Index("Notification_organizationId_idx", "organizationId"),
        Index("Notification_projectId_idx", "projectId"),
        Index("Notification_createdAt_idx", "createdAt"),
        Index("Notification_sentAt_idx", "sentAt"),
    )

    organization: Mapped[Optional["Organization"]] = relationship(back_populates="notifications")
    project: Mapped[Optional["Project"]] = relationship(back_populates="notifications")

class Agent(Base):
    __tablename__ = "Agent"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    project_id: Mapped[str] = mapped_column("projectId", ForeignKey("Project.id", ondelete="RESTRICT", onupdate="CASCADE", name="Agent_projectId_fkey"))
    name: Mapped[str]
    type: Mapped[str]
    config: Mapped[dict[str, Any]] = mapped_column(JSONB)
    workflow_id: Mapped[Optional[str]] = mapped_column("workflowId", ForeignKey("Workflow.id", ondelete="SET NULL", onupdate="CASCADE", name="Agent_workflowId_fkey"))
    copied_from_agent_id: Mapped[Optional[str]] = mapped_column("copiedFromAgentId", ForeignKey("Agent.id", ondelete="RESTRICT", onupdate="RESTRICT", name="Agent_copiedFromAgentId_fkey"))
    environment: Mapped[Optional[str]]
    owner_user_id: Mapped[Optional[str]] = mapped_column("ownerUserId")
    host_label: Mapped[Optional[str]] = mapped_column("hostLabel")
    identity_key: Mapped[Optional[str]] = mapped_column("identityKey")
    last_seen_at: Mapped[Optional[datetime]] = mapped_column("lastSeenAt")
    archived_at: Mapped[Optional[datetime]] = mapped_column("archivedAt")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    created_by_id: Mapped[Optional[str]] = mapped_column("createdById")
    updated_by_id: Mapped[Optional[str]] = mapped_column("updatedById")

    __table_args__ = (
        Index("Agent_projectId_idx", "projectId"),
        Index("Agent_workflowId_idx", "workflowId"),
        Index("Agent_copiedFromAgentId_idx", "copiedFromAgentId"),
        Index("Agent_projectId_lastSeenAt_idx", "projectId", "lastSeenAt"),
        Index("Agent_projectId_identityKey_key", "projectId", "identityKey", unique=True),
    )

    project: Mapped["Project"] = relationship(back_populates="agents")
    workflow: Mapped[Optional["Workflow"]] = relationship(back_populates="agents")
    copied_from: Mapped[Optional["Agent"]] = relationship(remote_side=[id], back_populates="copied_agents")
    copied_agents: Mapped[list["Agent"]] = relationship(passive_deletes="all", back_populates="copied_from")

class AgentVersion(Base):
    """One immutable snapshot of an agent; the Agent row is always the latest."""

    __tablename__ = "AgentVersion"

    id: Mapped[str] = mapped_column(primary_key=True)
    project_id: Mapped[str] = mapped_column("projectId")
    agent_id: Mapped[str] = mapped_column("agentId")
    version: Mapped[int]
    name: Mapped[str]
    type: Mapped[str]
    config: Mapped[dict[str, Any]] = mapped_column(JSONB)
    workflow_id: Mapped[Optional[str]] = mapped_column("workflowId")
    commit_message: Mapped[Optional[str]] = mapped_column("commitMessage")
    author_id: Mapped[Optional[str]] = mapped_column("authorId")
    author_kind: Mapped[Optional[str]] = mapped_column("authorKind")
    """"user", "api_key" or "system": who wrote the version when authorId is null."""
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("AgentVersion_projectId_idx", "projectId"),
        Index("AgentVersion_agentId_version_key", "agentId", "version", unique=True),
    )

class AgentVersionLabel(Base):
    """A deploy label ("production", "staging") pointing at one version of an agent."""

    __tablename__ = "AgentVersionLabel"

    id: Mapped[str] = mapped_column(primary_key=True)
    project_id: Mapped[str] = mapped_column("projectId")
    agent_id: Mapped[str] = mapped_column("agentId")
    label: Mapped[str]
    version_id: Mapped[str] = mapped_column("versionId")
    created_by_id: Mapped[Optional[str]] = mapped_column("createdById")
    updated_by_id: Mapped[Optional[str]] = mapped_column("updatedById")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("AgentVersionLabel_projectId_idx", "projectId"),
        Index("AgentVersionLabel_versionId_idx", "versionId"),
        Index("AgentVersionLabel_agentId_label_key", "agentId", "label", unique=True),
    )

class Evaluator(Base):
    __tablename__ = "Evaluator"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    project_id: Mapped[str] = mapped_column("projectId", ForeignKey("Project.id", ondelete="RESTRICT", onupdate="CASCADE", name="Evaluator_projectId_fkey"))
    name: Mapped[str]
    slug: Mapped[Optional[str]]
    type: Mapped[str]
    config: Mapped[dict[str, Any]] = mapped_column(JSONB)
    workflow_id: Mapped[Optional[str]] = mapped_column("workflowId", ForeignKey("Workflow.id", ondelete="SET NULL", onupdate="CASCADE", name="Evaluator_workflowId_fkey"))
    copied_from_evaluator_id: Mapped[Optional[str]] = mapped_column("copiedFromEvaluatorId", ForeignKey("Evaluator.id", ondelete="RESTRICT", onupdate="RESTRICT", name="Evaluator_copiedFromEvaluatorId_fkey"))
    archived_at: Mapped[Optional[datetime]] = mapped_column("archivedAt")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    created_by_id: Mapped[Optional[str]] = mapped_column("createdById")
    updated_by_id: Mapped[Optional[str]] = mapped_column("updatedById")

    __table_args__ = (
        Index("Evaluator_projectId_idx", "projectId"),
        Index("Evaluator_workflowId_idx", "workflowId"),
        Index("Evaluator_copiedFromEvaluatorId_idx", "copiedFromEvaluatorId"),
        Index("Evaluator_projectId_slug_key", "projectId", "slug", unique=True),
    )

    project: Mapped["Project"] = relationship(back_populates="evaluators")
    workflow: Mapped[Optional["Workflow"]] = relationship(back_populates="evaluators")
    copied_from: Mapped[Optional["Evaluator"]] = relationship(remote_side=[id], back_populates="copied_evaluators")
    copied_evaluators: Mapped[list["Evaluator"]] = relationship(passive_deletes="all", back_populates="copied_from")
    monitors: Mapped[list["Monitor"]] = relationship(back_populates="evaluator")
    gateway_guardrails: Mapped[list["GatewayGuardrail"]] = relationship(passive_deletes="all", back_populates="evaluator")

class EvaluatorVersion(Base):
    """One immutable snapshot of an evaluator; the Evaluator row is always the latest."""

    __tablename__ = "EvaluatorVersion"

    id: Mapped[str] = mapped_column(primary_key=True)
    project_id: Mapped[str] = mapped_column("projectId")
    evaluator_id: Mapped[str] = mapped_column("evaluatorId")
    version: Mapped[int]
    name: Mapped[str]
    type: Mapped[str]
    config: Mapped[dict[str, Any]] = mapped_column(JSONB)
    workflow_id: Mapped[Optional[str]] = mapped_column("workflowId")
    commit_message: Mapped[Optional[str]] = mapped_column("commitMessage")
    author_id: Mapped[Optional[str]] = mapped_column("authorId")
    author_kind: Mapped[Optional[str]] = mapped_column("authorKind")
    """"user", "api_key" or "system": who wrote the version when authorId is null."""
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("EvaluatorVersion_projectId_idx", "projectId"),
        Index("EvaluatorVersion_evaluatorId_version_key", "evaluatorId", "version", unique=True),
    )

class EvaluatorVersionLabel(Base):
    """A deploy label ("production", "staging") pointing at one version of an evaluator."""

    __tablename__ = "EvaluatorVersionLabel"

    id: Mapped[str] = mapped_column(primary_key=True)
    project_id: Mapped[str] = mapped_column("projectId")
    evaluator_id: Mapped[str] = mapped_column("evaluatorId")
    label: Mapped[str]
    version_id: Mapped[str] = mapped_column("versionId")
    created_by_id: Mapped[Optional[str]] = mapped_column("createdById")
    updated_by_id: Mapped[Optional[str]] = mapped_column("updatedById")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("EvaluatorVersionLabel_projectId_idx", "projectId"),
        Index("EvaluatorVersionLabel_versionId_idx", "versionId"),
        Index("EvaluatorVersionLabel_evaluatorId_label_key", "evaluatorId", "label", unique=True),
    )

class Scenario(Base):
    __tablename__ = "Scenario"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    project_id: Mapped[str] = mapped_column("projectId", ForeignKey("Project.id", ondelete="RESTRICT", onupdate="CASCADE", name="Scenario_projectId_fkey"))
    name: Mapped[str]
    situation: Mapped[str]
    criteria: Mapped[list[str]] = mapped_column(ARRAY(Text), nullable=True)
    labels: Mapped[list[str]] = mapped_column(ARRAY(Text), nullable=True)
    parameters: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))
    simulator_model: Mapped[Optional[str]] = mapped_column("simulatorModel")
    judge_model: Mapped[Optional[str]] = mapped_column("judgeModel")
    max_turns: Mapped[Optional[int]] = mapped_column("maxTurns")
    min_turns: Mapped[Optional[int]] = mapped_column("minTurns")
    fields: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))
    caller_voice: Mapped[Optional[dict[str, Any]]] = mapped_column("callerVoice", JSONB(none_as_null=True))
    test_suite_id: Mapped[Optional[str]] = mapped_column("testSuiteId", ForeignKey("SimulationSuite.id", ondelete="SET NULL", onupdate="CASCADE", name="Scenario_testSuiteId_fkey"))
    version: Mapped[int] = mapped_column(default=1, server_default=text("1"))
    last_updated_by_id: Mapped[Optional[str]] = mapped_column("lastUpdatedById", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="Scenario_lastUpdatedById_fkey"))
    archived_at: Mapped[Optional[datetime]] = mapped_column("archivedAt")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("Scenario_projectId_idx", "projectId"),
        Index("Scenario_projectId_testSuiteId_idx", "projectId", "testSuiteId"),
    )

    project: Mapped["Project"] = relationship(back_populates="scenarios")
    test_suite: Mapped[Optional["SimulationSuite"]] = relationship(back_populates="test_suite_scenarios")
    last_updated_by: Mapped[Optional["User"]] = relationship(back_populates="scenarios")
    versions: Mapped[list["ScenarioVersion"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="scenario")

class ScenarioVersion(Base):
    __tablename__ = "ScenarioVersion"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    scenario_id: Mapped[str] = mapped_column("scenarioId", ForeignKey("Scenario.id", ondelete="CASCADE", onupdate="CASCADE", name="ScenarioVersion_scenarioId_fkey"))
    project_id: Mapped[str] = mapped_column("projectId")
    """The project of the parent case, carried here so every read filters by
    tenant without a join. The service is what keeps the pair honest: no
    composite foreign key can, under relationMode="prisma".
    """
    version: Mapped[int]
    author_id: Mapped[Optional[str]] = mapped_column("authorId")
    author_label: Mapped[str] = mapped_column("authorLabel")
    """Who wrote this version: "user", "api", "cli" or "langy"."""
    change_description: Mapped[Optional[str]] = mapped_column("changeDescription")
    snapshot: Mapped[dict[str, Any]] = mapped_column(JSONB)
    """Envelope { schemaVersion, fields, changedFields } holding the editable
    field set as saved and the names of the fields this save changed.
    """
    schema_version: Mapped[int] = mapped_column("schemaVersion", default=1, server_default=text("1"))
    """Schema version of `snapshot`, so a later shape change can be read forward."""
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("ScenarioVersion_projectId_scenarioId_version_idx", "projectId", "scenarioId", "version"),
        Index("ScenarioVersion_createdAt_idx", "createdAt"),
        Index("ScenarioVersion_scenarioId_version_key", "scenarioId", "version", unique=True),
    )

    scenario: Mapped["Scenario"] = relationship(back_populates="versions")
    """Deleting the case takes its history with it. Under relationMode="prisma"
    the cascade is the client's, so a raw SQL delete still has to clear these
    rows itself.
    """

class SimulationSuite(Base):
    __tablename__ = "SimulationSuite"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    project_id: Mapped[str] = mapped_column("projectId", ForeignKey("Project.id", ondelete="RESTRICT", onupdate="CASCADE", name="SimulationSuite_projectId_fkey"))
    name: Mapped[str]
    slug: Mapped[str]
    description: Mapped[Optional[str]]
    scenario_ids: Mapped[list[str]] = mapped_column("scenarioIds", ARRAY(Text), nullable=True)
    targets: Mapped[dict[str, Any]] = mapped_column(JSONB)
    repeat_count: Mapped[int] = mapped_column("repeatCount", default=1, server_default=text("1"))
    labels: Mapped[list[str]] = mapped_column(ARRAY(Text), nullable=True)
    simulator_model: Mapped[Optional[str]] = mapped_column("simulatorModel")
    judge_model: Mapped[Optional[str]] = mapped_column("judgeModel")
    kind: Mapped[str] = mapped_column(default="run_plan", server_default=text("'run_plan'"))
    scope: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))
    fields: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))
    evaluators: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))
    archived_at: Mapped[Optional[datetime]] = mapped_column("archivedAt")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("SimulationSuite_projectId_idx", "projectId"),
        Index("SimulationSuite_projectId_kind_idx", "projectId", "kind"),
        Index("SimulationSuite_projectId_slug_key", "projectId", "slug", unique=True),
    )

    project: Mapped["Project"] = relationship(back_populates="suite_configurations")
    test_suite_scenarios: Mapped[list["Scenario"]] = relationship(back_populates="test_suite")

class Subscription(Base):
    __tablename__ = "Subscription"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="RESTRICT", onupdate="CASCADE", name="Subscription_organizationId_fkey"))
    plan: Mapped[str] = mapped_column(Enum("FREE", "PRO", "GROWTH", "GROWTH_SEAT_EUR_MONTHLY", "GROWTH_SEAT_EUR_ANNUAL", "GROWTH_SEAT_USD_MONTHLY", "GROWTH_SEAT_USD_ANNUAL", "ENTERPRISE", "LAUNCH", "ACCELERATE", "LAUNCH_ANNUAL", "ACCELERATE_ANNUAL", name="PlanTypes"))
    stripe_subscription_id: Mapped[Optional[str]] = mapped_column("stripeSubscriptionId")
    status: Mapped[str] = mapped_column(Enum("PENDING", "FAILED", "ACTIVE", "CANCELLED", name="SubscriptionStatus"))
    start_date: Mapped[Optional[datetime]] = mapped_column("startDate")
    end_date: Mapped[Optional[datetime]] = mapped_column("endDate")
    last_payment_failed_date: Mapped[Optional[datetime]] = mapped_column("lastPaymentFailedDate")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    max_members: Mapped[Optional[int]] = mapped_column("maxMembers")
    max_projects: Mapped[Optional[int]] = mapped_column("maxProjects")
    max_messages_per_month: Mapped[Optional[int]] = mapped_column("maxMessagesPerMonth")
    max_workflows: Mapped[Optional[int]] = mapped_column("maxWorkflows")
    max_retention_days: Mapped[Optional[int]] = mapped_column("maxRetentionDays")
    max_members_lite: Mapped[Optional[int]] = mapped_column("maxMembersLite")
    max_members_developer: Mapped[Optional[int]] = mapped_column("maxMembersDeveloper")
    """Developer seats the organization pays for; null falls to the plan's own cap."""
    max_teams: Mapped[Optional[int]] = mapped_column("maxTeams")
    max_prompts: Mapped[Optional[int]] = mapped_column("maxPrompts")
    max_evaluators: Mapped[Optional[int]] = mapped_column("maxEvaluators")
    max_scenarios: Mapped[Optional[int]] = mapped_column("maxScenarios")
    max_agents: Mapped[Optional[int]] = mapped_column("maxAgents")
    max_experiments: Mapped[Optional[int]] = mapped_column("maxExperiments")
    max_online_evaluations: Mapped[Optional[int]] = mapped_column("maxOnlineEvaluations")
    max_datasets: Mapped[Optional[int]] = mapped_column("maxDatasets")
    max_dashboards: Mapped[Optional[int]] = mapped_column("maxDashboards")
    max_custom_graphs: Mapped[Optional[int]] = mapped_column("maxCustomGraphs")
    max_automations: Mapped[Optional[int]] = mapped_column("maxAutomations")
    seats_kept: Mapped[Optional[int]] = mapped_column("seatsKept")
    """Seats paid for this period that stay usable after a seat decrease, until seatsKeptUntil."""
    seats_kept_until: Mapped[Optional[datetime]] = mapped_column("seatsKeptUntil")
    seat_line_quantity: Mapped[Optional[int]] = mapped_column("seatLineQuantity")
    """The seat quantity the last finished seat change set the seat line to; provider events never write it."""
    seat_line_quantity_at: Mapped[Optional[datetime]] = mapped_column("seatLineQuantityAt")
    seat_line_write_quantities: Mapped[list[int]] = mapped_column("seatLineWriteQuantities", ARRAY(Integer), nullable=True, default=lambda: [], server_default=text("ARRAY[]::integer[]"))
    """A Stripe seat line write recorded before it is sent and cleared once it is known to have
    finished under the lease; one left behind is settled by the seat line repair.
    Every seat quantity written while this record was pending.
    """
    seat_line_write_since: Mapped[Optional[datetime]] = mapped_column("seatLineWriteSince")
    """The first write still pending; a later holder's write keeps it."""
    seat_line_write_at: Mapped[Optional[datetime]] = mapped_column("seatLineWriteAt")
    """The latest write sent; the repair waits a grace window after it."""
    seat_line_write_attempts: Mapped[Optional[int]] = mapped_column("seatLineWriteAttempts")
    seat_line_write_flagged_at: Mapped[Optional[datetime]] = mapped_column("seatLineWriteFlaggedAt")
    seat_repair_quantity: Mapped[Optional[int]] = mapped_column("seatRepairQuantity")
    """The seat quantity Stripe should bill if a seat change stops after lifting the seat line; null once settled."""
    seat_repair_lifted_quantity: Mapped[Optional[int]] = mapped_column("seatRepairLiftedQuantity")
    """The quantity the change lifted the line to; the repair moves the line back only while Stripe still bills it."""
    seat_repair_since: Mapped[Optional[datetime]] = mapped_column("seatRepairSince")
    seat_repair_attempts: Mapped[Optional[int]] = mapped_column("seatRepairAttempts")
    seat_repair_flagged_at: Mapped[Optional[datetime]] = mapped_column("seatRepairFlaggedAt")
    """Set once the seat line repair has failed often enough that Ops should look at it."""
    seat_line_lease_token: Mapped[Optional[str]] = mapped_column("seatLineLeaseToken")
    """The lease a seat change or seat line repair holds while it reads and moves the seat line."""
    seat_line_lease_until: Mapped[Optional[datetime]] = mapped_column("seatLineLeaseUntil")

    __table_args__ = (
        Index("Subscription_stripeSubscriptionId_key", "stripeSubscriptionId", unique=True),
        Index("Subscription_organizationId_idx", "organizationId"),
    )

    organization: Mapped["Organization"] = relationship(back_populates="subscriptions")
    invoices: Mapped[list["Invoice"]] = relationship(passive_deletes="all", back_populates="subscription")

class Invoice(Base):
    __tablename__ = "Invoice"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    subscription_id: Mapped[str] = mapped_column("subscriptionId", ForeignKey("Subscription.id", ondelete="RESTRICT", onupdate="CASCADE", name="Invoice_subscriptionId_fkey"))
    total_amount: Mapped[float] = mapped_column("totalAmount")
    currency: Mapped[str]
    due_date: Mapped[datetime] = mapped_column("dueDate")
    paid: Mapped[bool] = mapped_column(default=False, server_default=text("false"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("Invoice_subscriptionId_idx", "subscriptionId"),
    )

    subscription: Mapped["Subscription"] = relationship(back_populates="invoices")
    line_items: Mapped[list["InvoiceItem"]] = relationship(passive_deletes="all", back_populates="invoice")

class InvoiceItem(Base):
    __tablename__ = "InvoiceItem"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    invoice_id: Mapped[str] = mapped_column("invoiceId", ForeignKey("Invoice.id", ondelete="RESTRICT", onupdate="CASCADE", name="InvoiceItem_invoiceId_fkey"))
    description: Mapped[str]
    amount: Mapped[float]
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("InvoiceItem_invoiceId_idx", "invoiceId"),
    )

    invoice: Mapped["Invoice"] = relationship(back_populates="line_items")

class Group(Base):
    __tablename__ = "Group"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="Group_organizationId_fkey"))
    name: Mapped[str]
    slug: Mapped[str]
    external_id: Mapped[Optional[str]] = mapped_column("externalId")
    scim_source: Mapped[Optional[str]] = mapped_column("scimSource")
    scim_connection_id: Mapped[Optional[str]] = mapped_column("scimConnectionId")
    """D08: which SSO connection provisioned this group. Null for manually
    created groups and for the SCIM groups that predate connection scoping.
    """
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("Group_organizationId_idx", "organizationId"),
        Index("Group_scimConnectionId_idx", "scimConnectionId"),
        Index("Group_organizationId_slug_key", "organizationId", "slug", unique=True),
        Index("Group_scimConnectionId_externalId_key", "scimConnectionId", "externalId", unique=True),
    )

    organization: Mapped["Organization"] = relationship(back_populates="groups")
    members: Mapped[list["GroupMembership"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="group")
    role_bindings: Mapped[list["RoleBinding"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="group")

class GroupMembership(Base):
    __tablename__ = "GroupMembership"

    user_id: Mapped[str] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="GroupMembership_userId_fkey"), primary_key=True)
    group_id: Mapped[str] = mapped_column("groupId", ForeignKey("Group.id", ondelete="CASCADE", onupdate="CASCADE", name="GroupMembership_groupId_fkey"), primary_key=True)
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("GroupMembership_userId_idx", "userId"),
        Index("GroupMembership_groupId_idx", "groupId"),
    )

    user: Mapped["User"] = relationship(back_populates="group_memberships")
    group: Mapped["Group"] = relationship(back_populates="members")

class AuthzUserStanding(Base):
    """Who authz treats as gone, from user's and identity's facts; authz never reads User (ARCHITECTURE.md, "Platform operators are a grant")."""

    __tablename__ = "AuthzUserStanding"

    user_id: Mapped[str] = mapped_column("userId", primary_key=True)
    deactivated_at: Mapped[Optional[datetime]] = mapped_column("deactivatedAt")
    erased_at: Mapped[Optional[datetime]] = mapped_column("erasedAt")
    standing_changed_at: Mapped[datetime] = mapped_column("standingChangedAt")

class AuthzEngineSwitch(Base):
    """Whether an organization's permission checks run on the authorization engine: authz's own
    switch, set when the engine migration proves the organization and pinned off by an operator's
    rollback. No row is the legacy path. Spec: modules/authz/specs/engine-switch.feature.
    """

    __tablename__ = "AuthzEngineSwitch"

    organization_id: Mapped[str] = mapped_column("organizationId", primary_key=True)
    on_engine: Mapped[bool] = mapped_column("onEngine")
    pinned: Mapped[bool] = mapped_column(default=False, server_default=text("false"))
    switched_at: Mapped[datetime] = mapped_column("switchedAt")
    last_applied_event_id: Mapped[Optional[str]] = mapped_column("lastAppliedEventId")
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

class AuthzOrganizationPolicy(Base):
    """ADR-201: an organisation's braid setting. Owned by authz; read into the organisation's snapshot."""

    __tablename__ = "AuthzOrganizationPolicy"

    organization_id: Mapped[str] = mapped_column("organizationId", primary_key=True)
    braids_organization_admins_only: Mapped[bool] = mapped_column("braidsOrganizationAdminsOnly", default=False, server_default=text("false"))
    updated_by_user_id: Mapped[str] = mapped_column("updatedByUserId")
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

class RoleBinding(Base):
    __tablename__ = "RoleBinding"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="RoleBinding_organizationId_fkey"))
    user_id: Mapped[Optional[str]] = mapped_column("userId", ForeignKey("User.id", ondelete="CASCADE", onupdate="CASCADE", name="RoleBinding_userId_fkey"))
    group_id: Mapped[Optional[str]] = mapped_column("groupId", ForeignKey("Group.id", ondelete="CASCADE", onupdate="CASCADE", name="RoleBinding_groupId_fkey"))
    api_key_id: Mapped[Optional[str]] = mapped_column("apiKeyId", ForeignKey("ApiKey.id", ondelete="RESTRICT", onupdate="CASCADE", name="RoleBinding_apiKeyId_fkey"))
    role: Mapped[str] = mapped_column(Enum("ADMIN", "MEMBER", "VIEWER", "CUSTOM", name="TeamUserRole"))
    custom_role_id: Mapped[Optional[str]] = mapped_column("customRoleId", ForeignKey("CustomRole.id", ondelete="SET NULL", onupdate="CASCADE", name="RoleBinding_customRoleId_fkey"))
    scope_type: Mapped[str] = mapped_column("scopeType", Enum("ORGANIZATION", "TEAM", "PROJECT", name="RoleBindingScopeType"))
    scope_id: Mapped[str] = mapped_column("scopeId")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("RoleBinding_organizationId_idx", "organizationId"),
        Index("RoleBinding_userId_idx", "userId"),
        Index("RoleBinding_groupId_idx", "groupId"),
        Index("RoleBinding_apiKeyId_idx", "apiKeyId"),
        Index("RoleBinding_scopeType_scopeId_idx", "scopeType", "scopeId"),
    )

    organization: Mapped["Organization"] = relationship(back_populates="role_bindings")
    user: Mapped[Optional["User"]] = relationship(back_populates="role_bindings")
    group: Mapped[Optional["Group"]] = relationship(back_populates="role_bindings")
    api_key: Mapped[Optional["ApiKey"]] = relationship(back_populates="role_bindings")
    custom_role: Mapped[Optional["CustomRole"]] = relationship(back_populates="role_bindings")

class Grant(Base):
    __tablename__ = "Grant"

    id: Mapped[str] = mapped_column(primary_key=True)
    organization_id: Mapped[str] = mapped_column("organizationId")
    principal_type: Mapped[str] = mapped_column("principalType", Enum("USER", "API_KEY", "GROUP", "TEAM", "ORGANIZATION", "PROJECT", "ANYONE", name="GrantPrincipalType"))
    principal_id: Mapped[Optional[str]] = mapped_column("principalId")
    role_key: Mapped[Optional[str]] = mapped_column("roleKey")
    legacy_role: Mapped[Optional[str]] = mapped_column("legacyRole")
    source: Mapped[str]
    scope_type: Mapped[str] = mapped_column("scopeType", Enum("ORGANIZATION", "TEAM", "PROJECT", "RESOURCE", "PLATFORM", name="GrantScopeType"))
    scope_id: Mapped[str] = mapped_column("scopeId")
    token: Mapped[Optional[str]]
    permission: Mapped[Optional[str]]
    resource_kind: Mapped[Optional[str]] = mapped_column("resourceKind")
    project_id: Mapped[Optional[str]] = mapped_column("projectId")
    created_by_user_id: Mapped[Optional[str]] = mapped_column("createdByUserId")
    expires_at: Mapped[Optional[datetime]] = mapped_column("expiresAt")
    max_views: Mapped[Optional[int]] = mapped_column("maxViews")
    condition: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))
    replaces_grant_id: Mapped[Optional[str]] = mapped_column("replacesGrantId")
    occurred_at: Mapped[datetime] = mapped_column("occurredAt")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    revoked_at: Mapped[Optional[datetime]] = mapped_column("revokedAt")
    revoked_reason: Mapped[Optional[str]] = mapped_column("revokedReason")

    __table_args__ = (
        Index("Grant_token_key", "token", unique=True),
        Index("Grant_organizationId_scopeType_scopeId_idx", "organizationId", "scopeType", "scopeId"),
        Index("Grant_projectId_resourceKind_scopeId_idx", "projectId", "resourceKind", "scopeId"),
    )

class GrantUsage(Base):
    __tablename__ = "GrantUsage"

    grant_id: Mapped[str] = mapped_column("grantId", primary_key=True)
    organization_id: Mapped[str] = mapped_column("organizationId")
    project_id: Mapped[str] = mapped_column("projectId")
    view_count: Mapped[int] = mapped_column("viewCount", default=0, server_default=text("0"))
    last_viewed_at: Mapped[Optional[datetime]] = mapped_column("lastViewedAt")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("GrantUsage_organizationId_idx", "organizationId"),
        Index("GrantUsage_projectId_idx", "projectId"),
    )

class Role(Base):
    __tablename__ = "Role"

    id: Mapped[str] = mapped_column(primary_key=True)
    organization_id: Mapped[str] = mapped_column("organizationId")
    name: Mapped[str]
    description: Mapped[Optional[str]]
    permissions: Mapped[dict[str, Any]] = mapped_column(JSONB)
    kind: Mapped[str] = mapped_column(default="custom", server_default=text("'custom'"))
    occurred_at: Mapped[datetime] = mapped_column("occurredAt")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    deleted_at: Mapped[Optional[datetime]] = mapped_column("deletedAt")

    __table_args__ = (
        Index("Role_organizationId_idx", "organizationId"),
    )

class ApiKey(Base):
    __tablename__ = "ApiKey"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    name: Mapped[str]
    description: Mapped[Optional[str]]
    lookup_id: Mapped[str] = mapped_column("lookupId")
    hashed_secret: Mapped[str] = mapped_column("hashedSecret")
    permission_mode: Mapped[str] = mapped_column("permissionMode", default="all", server_default=text("'all'"))
    user_id: Mapped[Optional[str]] = mapped_column("userId", ForeignKey("User.id", ondelete="RESTRICT", onupdate="CASCADE", name="ApiKey_userId_fkey"))
    created_by_user_id: Mapped[Optional[str]] = mapped_column("createdByUserId")
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="RESTRICT", onupdate="CASCADE", name="ApiKey_organizationId_fkey"))
    expires_at: Mapped[Optional[datetime]] = mapped_column("expiresAt")
    last_used_at: Mapped[Optional[datetime]] = mapped_column("lastUsedAt")
    revoked_at: Mapped[Optional[datetime]] = mapped_column("revokedAt")
    revocation_cause: Mapped[Optional[str]] = mapped_column("revocationCause")
    """Why the key was revoked, set with `revokedAt`: "user" when a person
    revoked it (the API-keys page, the REST API, the devices tab), "rotation"
    when a mint replaced it, "session" when the CLI login key it was parented
    to was revoked, "expired" when that session ran out, "offboarded" when the
    person's membership of the organization ended. "cap" survives on rows the
    retired per-tool cap revoked. The CLI reads it to decide whether a dead
    ingest key may be re-minted: a key a person revoked stays dead.
    NULL for live keys and for keys revoked before the cause was recorded.
    """
    parent_api_key_id: Mapped[Optional[str]] = mapped_column("parentApiKeyId")
    """The CLI login key of the device session that minted this ingestion key.
    Set by the CLI's personal mint, NULL for every other key. Revoking the
    parent (logout, the devices tab, a re-login from the same device, session
    expiry) revokes the children with it. Column ref only, no FK and no
    index: lookups filter by the indexed `userId` first and match the parent
    in memory over a person's few live ingest keys.
    """
    ingest_source_type: Mapped[Optional[str]] = mapped_column("ingestSourceType")
    """Non-null marks this key as an "ingestion key": a project-scoped,
    ingest-only ApiKey the unified `langwatch <tool>` CLI (and template
    installs) mint for OTLP Path B. The slug (claude_code / codex / gemini /
    opencode / claude_cowork) is stamped as `langwatch.source` provenance on
    every span the key authorizes. NULL for normal API / service keys.
    """
    ingestion_template_id: Mapped[Optional[str]] = mapped_column("ingestionTemplateId")
    """Set only for ingestion keys derived from an admin IngestionTemplate
    (e.g. claude_cowork). Drives receiver OTTL application + the
    `langwatch.template.id` provenance stamp. NULL for unified-CLI tools
    (which have no template) and for normal keys. Column ref only (no FK,
    per the app-layer-integrity convention); resolved by the ingest-key service.
    """
    created_by_device_label: Mapped[Optional[str]] = mapped_column("createdByDeviceLabel")
    """Human label of the CLI device session that minted this ingestion key
    (e.g. "Rogerio's MacBook Pro"), snapshotted at mint time from the device
    flow's `client_info`. Denormalized so the org-wide API-keys settings page
    can show which device a key came from without resolving another user's
    Redis-stored sessions. NULL for non-ingestion keys and for keys minted by
    a CLI old enough to predate device metadata.
    """
    is_system_managed: Mapped[bool] = mapped_column("isSystemManaged", default=False, server_default=text("false"))
    """True for a key LangWatch minted for itself (a run, the gateway's trace
    export, a Langy session). Hides it from listings, refuses customer edits
    and revokes, and lets an ownerless run key act as the system. Never set
    by a customer create; rows before the column stay false.
    """
    created_by_browser_session_id: Mapped[Optional[str]] = mapped_column("createdByBrowserSessionId")
    """The browser session that minted a docs session key (ADR-185), so signing
    out of that session revokes it. NULL for every other key. Unindexed:
    sign-out filters by the indexed `userId` first.
    """
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("ApiKey_lookupId_key", "lookupId", unique=True),
        Index("ApiKey_userId_idx", "userId"),
        Index("ApiKey_organizationId_idx", "organizationId"),
        Index("ApiKey_lookupId_idx", "lookupId"),
        Index("ApiKey_ingestSourceType_idx", "ingestSourceType"),
        Index("ApiKey_ingestionTemplateId_idx", "ingestionTemplateId"),
    )

    user: Mapped[Optional["User"]] = relationship(back_populates="api_keys")
    organization: Mapped["Organization"] = relationship(back_populates="api_keys")
    role_bindings: Mapped[list["RoleBinding"]] = relationship(passive_deletes="all", back_populates="api_key")

class SigningKey(Base):
    """ADR-201: a project's key for signing the tags on its calls. Owned by the API key module."""

    __tablename__ = "SigningKey"

    id: Mapped[str] = mapped_column(primary_key=True)
    organization_id: Mapped[str] = mapped_column("organizationId")
    project_id: Mapped[str] = mapped_column("projectId")
    name: Mapped[str]
    public_id: Mapped[str] = mapped_column("publicId")
    secret_ciphertext: Mapped[str] = mapped_column("secretCiphertext")
    created_by_user_id: Mapped[str] = mapped_column("createdByUserId")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    rotated_at: Mapped[Optional[datetime]] = mapped_column("rotatedAt")
    verifies_until: Mapped[Optional[datetime]] = mapped_column("verifiesUntil")
    revoked_at: Mapped[Optional[datetime]] = mapped_column("revokedAt")
    last_used_at: Mapped[Optional[datetime]] = mapped_column("lastUsedAt")

    __table_args__ = (
        Index("SigningKey_publicId_key", "publicId", unique=True),
        Index("SigningKey_projectId_idx", "projectId"),
        Index("SigningKey_organizationId_idx", "organizationId"),
    )

class ProjectActiveDay(Base):
    __tablename__ = "ProjectActiveDay"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    project_id: Mapped[str] = mapped_column("projectId")
    day: Mapped[str]
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("ProjectActiveDay_projectId_day_key", "projectId", "day", unique=True),
    )

class NurturingProject(Base):
    """Nurturing's read model: which organization a project belongs to, from project's created event.
    The per-email usage places a project and lists an organization's projects from these rows only.
    """

    __tablename__ = "NurturingProject"

    project_id: Mapped[str] = mapped_column("projectId", primary_key=True)
    organization_id: Mapped[str] = mapped_column("organizationId")

    __table_args__ = (
        Index("NurturingProject_organizationId_idx", "organizationId"),
    )

class NurturingUncopiedProject(Base):
    """A project nurturing saw traces of with no placement copy, and when it first did: every worker
    retries its trace batches only within one window from then, across restarts. Deleted when the
    copy lands.
    """

    __tablename__ = "NurturingUncopiedProject"

    project_id: Mapped[str] = mapped_column("projectId", primary_key=True)
    first_seen_at: Mapped[datetime] = mapped_column("firstSeenAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

class NurturingOrganizationMember(Base):
    """Nurturing's own copy of each organization's members, from the owners' member facts (§9):
    organization's signed_up and invite_accepted and auth's sso_auto_added join, member_removed
    and member_disabled leave, member_enabled returns. Each column keeps its latest time.
    """

    __tablename__ = "NurturingOrganizationMember"

    organization_id: Mapped[str] = mapped_column("organizationId", primary_key=True)
    user_id: Mapped[str] = mapped_column("userId", primary_key=True)
    joined_at: Mapped[Optional[datetime]] = mapped_column("joinedAt")
    left_at: Mapped[Optional[datetime]] = mapped_column("leftAt")
    returned_at: Mapped[Optional[datetime]] = mapped_column("returnedAt")

    __table_args__ = (
        Index("NurturingOrganizationMember_userId_idx", "userId"),
    )

class NurturingOrganizationCreated(Base):
    """When an organization was created, from organization's created event: nurturing's own copy for
    the per-email usage's day offsets. Its own table, so an early creation never adds an unseeded
    NurturingOrganization row.
    """

    __tablename__ = "NurturingOrganizationCreated"

    organization_id: Mapped[str] = mapped_column("organizationId", primary_key=True)
    created_at: Mapped[datetime] = mapped_column("createdAt")

    __table_args__ = (
        Index("NurturingOrganizationCreated_createdAt_idx", "createdAt"),
    )

class NurturingOrganization(Base):
    """Nurturing's read model per organization: its admin and its milestone counts.
    `seeded` marks an organization learned from the project backfill: its first_* already happened.
    """

    __tablename__ = "NurturingOrganization"

    organization_id: Mapped[str] = mapped_column("organizationId", primary_key=True)
    admin_user_id: Mapped[Optional[str]] = mapped_column("adminUserId")
    seeded: Mapped[bool] = mapped_column(default=False, server_default=text("false"))
    evaluation_count: Mapped[int] = mapped_column("evaluationCount", default=0, server_default=text("0"))
    simulation_run_count: Mapped[int] = mapped_column("simulationRunCount", default=0, server_default=text("0"))
    coding_first_session_at: Mapped[Optional[datetime]] = mapped_column("codingFirstSessionAt")
    coding_agent_types: Mapped[list[str]] = mapped_column("codingAgentTypes", ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))
    coding_active_days: Mapped[int] = mapped_column("codingActiveDays", default=0, server_default=text("0"))
    coding_last_active_day: Mapped[Optional[str]] = mapped_column("codingLastActiveDay")
    coding_terminal_types: Mapped[list[str]] = mapped_column("codingTerminalTypes", ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))
    coding_repository_keys: Mapped[list[str]] = mapped_column("codingRepositoryKeys", ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))
    coding_first_days: Mapped[list[int]] = mapped_column("codingFirstDays", ARRAY(Integer), nullable=True, default=lambda: [], server_default=text("ARRAY[]::integer[]"))
    last_trace_at: Mapped[Optional[datetime]] = mapped_column("lastTraceAt")
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

class NurturingProjectUsageView(Base):
    """Nurturing's own per-project usage: a peer fold over the owners' prompt, evaluator, scenario
    and simulation-run events (ARCHITECTURE §9), one peer fold per owner pipeline, summed per
    organization for the per-email traits. Traces are counted in NurturingProjectTraceDay.
    """

    __tablename__ = "NurturingProjectUsageView"

    project_id: Mapped[str] = mapped_column("projectId", primary_key=True)
    owner: Mapped[str] = mapped_column(primary_key=True)
    """The owner pipeline whose peer fold wrote the row (prompt, evaluator, scenario, simulation),
    or `evaluation`: the project's evaluation days, written by a subscriber, at most once a day.
    """
    folded: Mapped[dict[str, Any]] = mapped_column(JSONB)
    """The project's usage as the fold folded it (NurturingProjectUsageState)."""
    applied_event_ids: Mapped[list[str]] = mapped_column("appliedEventIds", ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))
    """The event ids of the last delivery folded in, so a redelivery folds once."""
    projection_version: Mapped[str] = mapped_column("projectionVersion")
    """A row of another version reads as undecodable and is re-folded from the log."""
    updated_at: Mapped[datetime] = mapped_column("updatedAt")
    """Event-derived, so a replayed row equals the row it rebuilds."""

class NurturingProjectTraceDay(Base):
    """Nurturing's own trace counts from trace's span_received and trace_failed (ARCHITECTURE §9), a
    root span being one trace, one row per project and UTC day of append time, incremented in place
    by each batch. A redelivered batch can count a trace twice; accepted (issue 9378). Kept 30 days.
    """

    __tablename__ = "NurturingProjectTraceDay"

    project_id: Mapped[str] = mapped_column("projectId", primary_key=True)
    day: Mapped[str] = mapped_column(primary_key=True)
    """UTC `YYYY-MM-DD` of the root spans' append time."""
    trace_count: Mapped[int] = mapped_column("traceCount", default=0, server_default=text("0"))
    error_count: Mapped[int] = mapped_column("errorCount", default=0, server_default=text("0"))
    """Traces trace recorded failed (lw.trace.trace_failed)."""
    last_trace_at: Mapped[Optional[datetime]] = mapped_column("lastTraceAt")
    """Append time of the newest root span."""
    last_failed_trace_at: Mapped[Optional[datetime]] = mapped_column("lastFailedTraceAt")
    """Append time of the newest trace_failed."""

class NurturingPersonalProject(Base):
    """Nurturing's copy of each personal project's owner, from organization's and project's facts.
    The first owner recorded stays: a personal project's owner never changes.
    """

    __tablename__ = "NurturingPersonalProject"

    project_id: Mapped[str] = mapped_column("projectId", primary_key=True)
    organization_id: Mapped[str] = mapped_column("organizationId")
    owner_user_id: Mapped[str] = mapped_column("ownerUserId")

class NurturingTracedProject(Base):
    """Nurturing's copy of each project that received a real trace, from trace's and project's facts."""

    __tablename__ = "NurturingTracedProject"

    project_id: Mapped[str] = mapped_column("projectId", primary_key=True)
    organization_id: Mapped[str] = mapped_column("organizationId")

    __table_args__ = (
        Index("NurturingTracedProject_organizationId_idx", "organizationId"),
    )

class BillingMeterCheckpoint(Base):
    __tablename__ = "BillingMeterCheckpoint"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: cuid.cuid())
    organization_id: Mapped[str] = mapped_column("organizationId")
    billing_month: Mapped[str] = mapped_column("billingMonth")
    meter: Mapped[str] = mapped_column(default="langwatch_billable_events", server_default=text("'langwatch_billable_events'"))
    last_reported_total: Mapped[int] = mapped_column("lastReportedTotal", default=0, server_default=text("0"))
    pending_reported_total: Mapped[Optional[int]] = mapped_column("pendingReportedTotal")
    consecutive_failures: Mapped[int] = mapped_column("consecutiveFailures", default=0, server_default=text("0"))
    last_counted_event_id: Mapped[Optional[str]] = mapped_column("lastCountedEventId")
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("BillingMeterCheckpoint_organizationId_billingMonth_meter_key", "organizationId", "billingMonth", "meter", unique=True),
    )

class VirtualKey(Base):
    __tablename__ = "VirtualKey"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId")
    name: Mapped[str]
    description: Mapped[Optional[str]]
    status: Mapped[str] = mapped_column(Enum("ACTIVE", "DISABLED", "REVOKED", name="VirtualKeyStatus"), default="ACTIVE", server_default=text("'ACTIVE'"))
    purpose: Mapped[str] = mapped_column(Enum("USER", "LANGY", "CONNECT", name="VirtualKeyPurpose"), default="USER", server_default=text("'USER'"))
    external_id: Mapped[Optional[str]] = mapped_column("externalId")
    metadata_: Mapped[dict[str, Any]] = mapped_column("metadata", JSONB, default=lambda: {}, server_default=text("'{}'"))
    disabled_at: Mapped[Optional[datetime]] = mapped_column("disabledAt")
    disabled_reason: Mapped[Optional[str]] = mapped_column("disabledReason")
    expires_at: Mapped[Optional[datetime]] = mapped_column("expiresAt")
    hashed_secret: Mapped[str] = mapped_column("hashedSecret")
    display_prefix: Mapped[str] = mapped_column("displayPrefix")
    principal_user_id: Mapped[Optional[str]] = mapped_column("principalUserId", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="VirtualKey_principalUserId_fkey"))
    trace_project_id: Mapped[Optional[str]] = mapped_column("traceProjectId")
    config: Mapped[dict[str, Any]] = mapped_column(JSONB, default=lambda: {}, server_default=text("'{}'"))
    revision: Mapped[int] = mapped_column(BigInteger, default=0, server_default=text("0"))
    previous_hashed_secret: Mapped[Optional[str]] = mapped_column("previousHashedSecret")
    previous_secret_valid_until: Mapped[Optional[datetime]] = mapped_column("previousSecretValidUntil")
    revoked_at: Mapped[Optional[datetime]] = mapped_column("revokedAt")
    revoked_by_id: Mapped[Optional[str]] = mapped_column("revokedById")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    created_by_id: Mapped[str] = mapped_column("createdById")
    last_used_at: Mapped[Optional[datetime]] = mapped_column("lastUsedAt")
    routing_policy_id: Mapped[Optional[str]] = mapped_column("routingPolicyId", ForeignKey("RoutingPolicy.id", ondelete="SET NULL", onupdate="CASCADE", name="VirtualKey_routingPolicyId_fkey"))
    routing_mode: Mapped[str] = mapped_column("routingMode", Enum("NONE", "FALLBACK_ALL", "POLICY", name="VirtualKeyRoutingMode"), default="NONE", server_default=text("'NONE'"))
    connect_services: Mapped[list[str]] = mapped_column("connectServices", ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))
    license_token_hash: Mapped[Optional[str]] = mapped_column("licenseTokenHash")
    license_instance_id: Mapped[Optional[str]] = mapped_column("licenseInstanceId")
    license_expires_at: Mapped[Optional[datetime]] = mapped_column("licenseExpiresAt")
    license_id: Mapped[Optional[str]] = mapped_column("licenseId")

    __table_args__ = (
        Index("VirtualKey_licenseTokenHash_key", "licenseTokenHash", unique=True),
        Index("VirtualKey_organizationId_idx", "organizationId"),
        Index("VirtualKey_displayPrefix_idx", "displayPrefix"),
        Index("VirtualKey_status_idx", "status"),
        Index("VirtualKey_principalUserId_idx", "principalUserId"),
        Index("VirtualKey_revision_idx", "revision"),
        Index("VirtualKey_routingPolicyId_idx", "routingPolicyId"),
        Index("VirtualKey_organizationId_purpose_idx", "organizationId", "purpose"),
        Index("VirtualKey_hashedSecret_key", "hashedSecret", unique=True),
        Index("VirtualKey_organizationId_externalId_key", "organizationId", "externalId", unique=True),
    )

    principal_user: Mapped[Optional["User"]] = relationship(back_populates="principal_virtual_keys")
    routing_policy: Mapped[Optional["RoutingPolicy"]] = relationship(back_populates="virtual_keys")
    scopes: Mapped[list["VirtualKeyScope"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="virtual_key")
    ledger_entries: Mapped[list["GatewayBudgetLedger"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="virtual_key")

class VirtualKeyScope(Base):
    __tablename__ = "VirtualKeyScope"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    virtual_key_id: Mapped[str] = mapped_column("virtualKeyId", ForeignKey("VirtualKey.id", ondelete="CASCADE", onupdate="CASCADE", name="VirtualKeyScope_virtualKeyId_fkey"))
    scope_type: Mapped[str] = mapped_column("scopeType", Enum("ORGANIZATION", "TEAM", "PROJECT", name="VirtualKeyScopeType"))
    scope_id: Mapped[str] = mapped_column("scopeId")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("VirtualKeyScope_scopeType_scopeId_idx", "scopeType", "scopeId"),
        Index("VirtualKeyScope_virtualKeyId_idx", "virtualKeyId"),
        Index("VirtualKeyScope_virtualKeyId_scopeType_scopeId_key", "virtualKeyId", "scopeType", "scopeId", unique=True),
    )

    virtual_key: Mapped["VirtualKey"] = relationship(back_populates="scopes")

class RoutingPolicy(Base):
    __tablename__ = "RoutingPolicy"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="RoutingPolicy_organizationId_fkey"))
    name: Mapped[str]
    description: Mapped[Optional[str]]
    model_provider_ids: Mapped[dict[str, Any]] = mapped_column("modelProviderIds", JSONB, default=lambda: [], server_default=text("'[]'"))
    model_aliases: Mapped[dict[str, Any]] = mapped_column("modelAliases", JSONB, default=lambda: {}, server_default=text("'{}'"))
    default_model: Mapped[Optional[str]] = mapped_column("defaultModel")
    policy_rules: Mapped[dict[str, Any]] = mapped_column("policyRules", JSONB, default=lambda: {}, server_default=text("'{}'"))
    is_default: Mapped[bool] = mapped_column("isDefault", default=False, server_default=text("false"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    created_by_id: Mapped[Optional[str]] = mapped_column("createdById", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="RoutingPolicy_createdById_fkey"))
    updated_by_id: Mapped[Optional[str]] = mapped_column("updatedById", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="RoutingPolicy_updatedById_fkey"))

    __table_args__ = (
        Index("RoutingPolicy_organizationId_idx", "organizationId"),
    )

    organization: Mapped["Organization"] = relationship(back_populates="routing_policies")
    created_by: Mapped[Optional["User"]] = relationship(foreign_keys=[created_by_id], back_populates="routing_policies_created")
    updated_by: Mapped[Optional["User"]] = relationship(foreign_keys=[updated_by_id], back_populates="routing_policies_updated")
    virtual_keys: Mapped[list["VirtualKey"]] = relationship(back_populates="routing_policy")
    scopes: Mapped[list["RoutingPolicyScope"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="routing_policy")

class RoutingPolicyScope(Base):
    __tablename__ = "RoutingPolicyScope"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    routing_policy_id: Mapped[str] = mapped_column("routingPolicyId", ForeignKey("RoutingPolicy.id", ondelete="CASCADE", onupdate="CASCADE", name="RoutingPolicyScope_routingPolicyId_fkey"))
    scope_type: Mapped[str] = mapped_column("scopeType", Enum("ORGANIZATION", "TEAM", "PROJECT", name="RoutingPolicyScopeType"))
    scope_id: Mapped[str] = mapped_column("scopeId")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("RoutingPolicyScope_scopeType_scopeId_idx", "scopeType", "scopeId"),
        Index("RoutingPolicyScope_routingPolicyId_idx", "routingPolicyId"),
        Index("RoutingPolicyScope_routingPolicyId_scopeType_scopeId_key", "routingPolicyId", "scopeType", "scopeId", unique=True),
    )

    routing_policy: Mapped["RoutingPolicy"] = relationship(back_populates="scopes")

class GatewayGuardrail(Base):
    __tablename__ = "GatewayGuardrail"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    project_id: Mapped[str] = mapped_column("projectId", ForeignKey("Project.id", ondelete="CASCADE", onupdate="CASCADE", name="GatewayGuardrail_projectId_fkey"))
    name: Mapped[str]
    description: Mapped[Optional[str]]
    evaluator_id: Mapped[str] = mapped_column("evaluatorId", ForeignKey("Evaluator.id", ondelete="RESTRICT", onupdate="CASCADE", name="GatewayGuardrail_evaluatorId_fkey"))
    direction: Mapped[str] = mapped_column(Enum("PRE", "POST", "STREAM_CHUNK", name="GatewayGuardrailDirection"))
    failure_mode: Mapped[str] = mapped_column("failureMode", Enum("FAIL_OPEN", "FAIL_CLOSED", name="GatewayGuardrailFailureMode"), default="FAIL_CLOSED", server_default=text("'FAIL_CLOSED'"))
    archived_at: Mapped[Optional[datetime]] = mapped_column("archivedAt")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    created_by_id: Mapped[Optional[str]] = mapped_column("createdById", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="GatewayGuardrail_createdById_fkey"))
    updated_by_id: Mapped[Optional[str]] = mapped_column("updatedById", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="GatewayGuardrail_updatedById_fkey"))

    __table_args__ = (
        Index("GatewayGuardrail_projectId_idx", "projectId"),
        Index("GatewayGuardrail_evaluatorId_idx", "evaluatorId"),
        Index("GatewayGuardrail_projectId_archivedAt_idx", "projectId", "archivedAt"),
    )

    project: Mapped["Project"] = relationship(back_populates="gateway_guardrails")
    evaluator: Mapped["Evaluator"] = relationship(back_populates="gateway_guardrails")
    created_by: Mapped[Optional["User"]] = relationship(foreign_keys=[created_by_id], back_populates="gateway_guardrails_created")
    updated_by: Mapped[Optional["User"]] = relationship(foreign_keys=[updated_by_id], back_populates="gateway_guardrails_updated")

class IngestionSource(Base):
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

    __tablename__ = "IngestionSource"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="IngestionSource_organizationId_fkey"))
    team_id: Mapped[Optional[str]] = mapped_column("teamId", ForeignKey("Team.id", ondelete="SET NULL", onupdate="CASCADE", name="IngestionSource_teamId_fkey"))
    """Optional team-scope. null = org-wide."""
    source_type: Mapped[str] = mapped_column("sourceType")
    """'otel_generic' | 'claude_cowork' | 'workato' | 'copilot_studio' |
    'openai_compliance' | 'claude_compliance' | 's3_custom' (extensible).
    String not enum so adapters can be added without migrations - the
    supported set lives in code (activity-monitor adapter registry).
    """
    name: Mapped[str]
    """Human-friendly label for the admin UI. Unique per-org so two
    "Cowork" sources can coexist with different names.
    """
    description: Mapped[Optional[str]]
    ingest_secret_hash: Mapped[str] = mapped_column("ingestSecretHash")
    """HMAC bearer token the upstream operator pastes into the platform
    (or LangWatch sends in pull mode). Stored encrypted at rest via
    the same pepper-derived KEK as VirtualKey secrets. Auto-generated
    on create, rotatable via the admin UI with a 24h grace window.
    """
    parser_config: Mapped[dict[str, Any]] = mapped_column("parserConfig", JSONB, default=lambda: {}, server_default=text("'{}'"))
    """Per-source-type configuration. Shape is platform-specific:
    - otel_generic: { allowedSourceTypeLabel?: string }
    - claude_cowork: { workspaceId: string }
    - workato: { sharedSecretLastFour: string }
    - copilot_studio: { tenantId, clientId, clientSecretHash, pollEverySec }
    - openai_compliance: { bucket, prefix, roleArn, pollEverySec }
    - claude_compliance: { workspaceApiKeyHash, pollEverySec }
    - s3_custom: { bucket, prefix, roleArn, parserDsl, pollEverySec }
    """
    poller_cursor: Mapped[Optional[dict[str, Any]]] = mapped_column("pollerCursor", JSONB(none_as_null=True))
    """For pull-based sources, persisted between polls so the next poll
    resumes where the last left off (timestamp / cursor / file marker).
    """
    error_count: Mapped[int] = mapped_column("errorCount", default=0, server_default=text("0"))
    """Phase 10 - number of consecutive errors observed by the puller
    worker. Reset to 0 on a successful run. Drives admin-UI banners
    ("Microsoft authentication failed - re-authorize") + back-off
    scheduling when the source has been broken for a while.
    Spec: specs/ai-governance/puller-framework/puller-adapter-contract.feature
    """
    pull_schedule: Mapped[Optional[str]] = mapped_column("pullSchedule")
    """Phase 10 - cron schedule for pull-mode sources. Validated by the
    adapter's pullConfig schema; mirrored here for fast worker
    dispatch. Null for push-mode sources.
    """
    status: Mapped[str] = mapped_column(default="awaiting_first_event", server_default=text("'awaiting_first_event'"))
    """Status flag. 'active' (default), 'disabled' (admin paused), or
    'awaiting_first_event' (created but never received).
    """
    trace_project_id: Mapped[Optional[str]] = mapped_column("traceProjectId")
    """ADR-088 v7: optional trace DESTINATION for conversation-bearing pulls
    (Databricks Genie). Null = don't route. Same shape as
    VirtualKey.traceProjectId: a destination is NOT a scope row and grants
    no access to the source; it only decides where routed conversation
    traces land. Must be a project of this org (validated at write time).
    Freely editable; already-routed history stays where it landed.
    NOT the deferred cost-attribution projectId (ADR-088 Decision 4).
    """
    provider_account_id: Mapped[Optional[str]] = mapped_column("providerAccountId")
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
    last_event_at: Mapped[Optional[datetime]] = mapped_column("lastEventAt")
    """Last event timestamp observed (any direction). Surfaces in the
    admin UI as "last event 2 minutes ago".
    """
    last_success_at: Mapped[Optional[datetime]] = mapped_column("lastSuccessAt")
    """ADR-128 - when a pull run last reached the provider without erroring,
    mirrored from the ingestionPullRunStatus projection. Deliberately NOT
    lastEventAt: an empty period is a working puller, and a run that
    delivered data and then failed is not. Drives the health badge and
    the "no data since" cut-off, below which a day's spend is unknown
    rather than zero.
    Spec: specs/governance/ingestion-source-health.feature
    """
    last_read_through_at: Mapped[Optional[datetime]] = mapped_column("lastReadThroughAt")
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
    last_run_completeness: Mapped[Optional[str]] = mapped_column("lastRunCompleteness")
    unpriced_usage_since: Mapped[Optional[datetime]] = mapped_column("unpricedUsageSince")
    """ADR-088 - the window this source read while pulled cost recording was
    off for the organization, so its spend was dropped rather than priced.
    Both null = nothing dropped. The pull cursor advances whether or not the
    money path is live (audit-only is a supported mode), so without this the
    loss is unrecoverable AND invisible: those days read as no spend rather
    than as unknown. Widened by any run that drops a priced item, cleared
    only when a later run re-reads back across the whole window.
    Spec: specs/governance/pulled-usage-cost-gap.feature
    """
    unpriced_usage_through: Mapped[Optional[datetime]] = mapped_column("unpricedUsageThrough")
    archived_at: Mapped[Optional[datetime]] = mapped_column("archivedAt")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    created_by_id: Mapped[Optional[str]] = mapped_column("createdById", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="IngestionSource_createdById_fkey"))

    __table_args__ = (
        Index("IngestionSource_organizationId_idx", "organizationId"),
        Index("IngestionSource_teamId_idx", "teamId"),
        Index("IngestionSource_sourceType_idx", "sourceType"),
        Index("IngestionSource_organizationId_name_key", "organizationId", "name", unique=True),
    )

    organization: Mapped["Organization"] = relationship(back_populates="ingestion_sources")
    team: Mapped[Optional["Team"]] = relationship(back_populates="ingestion_sources")
    created_by: Mapped[Optional["User"]] = relationship(back_populates="ingestion_sources_created")

class AnomalyRule(Base):
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

    __tablename__ = "AnomalyRule"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="AnomalyRule_organizationId_fkey"))
    scope: Mapped[str]
    """'organization' | 'team' | 'project' | 'source_type' | 'source'.
    String not enum so new scopes ship without migrations.
    """
    scope_id: Mapped[str] = mapped_column("scopeId")
    """Org id / team id / project id / source-type string / source id
    - narrows what the rule applies to.
    """
    name: Mapped[str]
    description: Mapped[Optional[str]]
    severity: Mapped[str]
    """'critical' | 'warning' | 'info'. Drives default destination
    routing (e.g. critical → PagerDuty, warning → Slack only).
    """
    rule_type: Mapped[str] = mapped_column("ruleType")
    """Extensible rule taxonomy: 'spend_spike' | 'rate_limit' |
    'tool_mismatch' | 'after_hours' | 'unusual_model' | 'pii_leak'
    | 'custom'. New types ship without migrations - evaluator
    picks a per-type implementation from a code registry.
    """
    threshold_config: Mapped[dict[str, Any]] = mapped_column("thresholdConfig", JSONB, default=lambda: {}, server_default=text("'{}'"))
    """Per-rule-type threshold shape. Examples:
    spend_spike: { windowSec, ratioVsBaseline, minBaselineUsd }
    rate_limit: { windowSec, requestsThreshold }
    after_hours: { startHour, endHour, timezone, requestsThreshold }
    """
    destination_config: Mapped[dict[str, Any]] = mapped_column("destinationConfig", JSONB, default=lambda: {}, server_default=text("'{}'"))
    """Alert routing per destination. Shape:
    { slack?: { channelId }, siem?: { endpointUrl, authHeader },
    webhook?: { url, authHeader }, pagerduty?: { routingKey },
    email?: { to: string[] } }
    Empty map = use org defaults (when defaults exist; today no-op).
    """
    status: Mapped[str] = mapped_column(default="active", server_default=text("'active'"))
    """'active' (default - evaluator picks it up) | 'disabled' (admin
    paused; rule rows kept for audit). 'archived' uses archivedAt
    timestamp instead so we can filter `archivedAt IS NULL` cheaply.
    """
    archived_at: Mapped[Optional[datetime]] = mapped_column("archivedAt")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    created_by_id: Mapped[Optional[str]] = mapped_column("createdById", ForeignKey("User.id", ondelete="SET NULL", onupdate="CASCADE", name="AnomalyRule_createdById_fkey"))

    __table_args__ = (
        Index("AnomalyRule_organizationId_idx", "organizationId"),
        Index("AnomalyRule_scope_scopeId_idx", "scope", "scopeId"),
        Index("AnomalyRule_ruleType_idx", "ruleType"),
        Index("AnomalyRule_status_idx", "status"),
        Index("AnomalyRule_organizationId_name_key", "organizationId", "name", unique=True),
    )

    organization: Mapped["Organization"] = relationship(back_populates="anomaly_rules")
    created_by: Mapped[Optional["User"]] = relationship(back_populates="anomaly_rules_created")
    alerts: Mapped[list["AnomalyAlert"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="rule")

class AnomalyAlert(Base):
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

    __tablename__ = "AnomalyAlert"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="AnomalyAlert_organizationId_fkey"))
    rule_id: Mapped[str] = mapped_column("ruleId", ForeignKey("AnomalyRule.id", ondelete="CASCADE", onupdate="CASCADE", name="AnomalyAlert_ruleId_fkey"))
    severity: Mapped[str]
    """Denormalised from rule for fast WHERE filtering on the
    dashboard query (no join needed for the common path).
    """
    rule_name: Mapped[str] = mapped_column("ruleName")
    rule_type: Mapped[str] = mapped_column("ruleType")
    trigger_window_start: Mapped[datetime] = mapped_column("triggerWindowStart")
    """Window the trigger refers to (UTC)."""
    trigger_window_end: Mapped[datetime] = mapped_column("triggerWindowEnd")
    trigger_spend_usd: Mapped[Optional[DecimalType]] = mapped_column("triggerSpendUsd", Numeric(precision=18, scale=6))
    """Numeric trigger context. Per-rule-type detail in `detail` JSONB."""
    trigger_event_count: Mapped[Optional[int]] = mapped_column("triggerEventCount")
    detail: Mapped[dict[str, Any]] = mapped_column(JSONB, default=lambda: {}, server_default=text("'{}'"))
    """Per-rule-type extra context: {topActors: [...], baseline: ...,
    observed: ..., timezone: ..., etc}.
    """
    state: Mapped[str] = mapped_column(default="open", server_default=text("'open'"))
    """'open' | 'acknowledged' | 'resolved'."""
    destination_status: Mapped[dict[str, Any]] = mapped_column("destinationStatus", JSONB, default=lambda: {}, server_default=text("'{}'"))
    """Per-destination dispatch outcome:
    { webhook: { lastAttemptIso, lastError?, attemptCount, statusCode? },
    slack: ..., pagerduty: ..., siem: ..., email: ... }
    """
    detected_at: Mapped[datetime] = mapped_column("detectedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    acknowledged_at: Mapped[Optional[datetime]] = mapped_column("acknowledgedAt")
    resolved_at: Mapped[Optional[datetime]] = mapped_column("resolvedAt")

    __table_args__ = (
        Index("AnomalyAlert_organizationId_detectedAt_idx", "organizationId", "detectedAt"),
        Index("AnomalyAlert_ruleId_idx", "ruleId"),
        Index("AnomalyAlert_state_idx", "state"),
        Index("AnomalyAlert_ruleId_triggerWindowStart_key", "ruleId", "triggerWindowStart", unique=True),
    )

    organization: Mapped["Organization"] = relationship(back_populates="anomaly_alerts")
    rule: Mapped["AnomalyRule"] = relationship(back_populates="alerts")

class AiToolEntry(Base):
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

    __tablename__ = "AiToolEntry"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="AiToolEntry_organizationId_fkey"))
    scope: Mapped[str]
    """"organization" | "team". String not enum so new scopes ship
    without migrations.
    """
    scope_id: Mapped[str] = mapped_column("scopeId")
    """organizationId or teamId, narrowing what scope this entry applies
    to. Always equals organizationId when scope=organization.
    """
    type: Mapped[str]
    """"coding_assistant" | "model_provider" | "external_tool".
    Discriminator drives the per-tile UX shape and validates
    the `config` JSON column at the app layer.
    """
    display_name: Mapped[str] = mapped_column("displayName")
    slug: Mapped[str]
    """Stable slug for team-overrides-org resolution (a team-scoped
    entry with the same slug as an org-scoped one hides the org
    default for users in that team). Server-owned: generated from
    displayName + nanoid(6) on create, immutable on update.
    """
    icon_key: Mapped[Optional[str]] = mapped_column("iconKey")
    """@deprecated Use `iconAsset` instead. Kept for back-compat with
    pre-refactor rows; new writes should populate iconAsset.
    """
    icon_asset: Mapped[Optional[str]] = mapped_column("iconAsset")
    """Icon source - single column, prefix-discriminated:
    "preset:claude_code" / "preset:codex" / ...   (built-in)
    "data:image/svg+xml;base64,..."               (admin-uploaded)
    null                                          → fall back to
    type-default in UI
    """
    order: Mapped[int] = mapped_column(default=0, server_default=text("0"))
    """Admin-curated sort within the user-facing portal grid."""
    enabled: Mapped[bool] = mapped_column(default=True, server_default=text("true"))
    """When false the entry is hidden from the user-facing list but
    remains in the admin list (so admins can re-enable without
    losing per-entry config). Disabled is "soft-hide", archived is
    "soft-delete" (preserves audit trail + scope binding history).
    """
    config: Mapped[dict[str, Any]] = mapped_column(JSONB, default=lambda: {}, server_default=text("'{}'"))
    """Per-type config JSON. Validated via zod discriminated union
    at the service layer (see AiToolEntryService).
    """
    archived_at: Mapped[Optional[datetime]] = mapped_column("archivedAt")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    created_by_id: Mapped[Optional[str]] = mapped_column("createdById")
    updated_by_id: Mapped[Optional[str]] = mapped_column("updatedById")

    __table_args__ = (
        Index("AiToolEntry_organizationId_scope_scopeId_idx", "organizationId", "scope", "scopeId"),
        Index("AiToolEntry_organizationId_enabled_archivedAt_idx", "organizationId", "enabled", "archivedAt"),
        Index("AiToolEntry_organizationId_type_idx", "organizationId", "type"),
    )

    organization: Mapped["Organization"] = relationship(back_populates="ai_tool_entries")
    teams: Mapped[list["AiToolEntryTeam"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="entry")
    """@deprecated Multi-team scope. Superseded by `departments` - the
    "Visible to" model is now WHOLE ORGANIZATION or a set of
    DEPARTMENTS. Rows are no longer written here (the table is kept
    for now to avoid a destructive drop); readers prefer `departments`.
    """
    departments: Mapped[list["AiToolEntryDepartment"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="entry")
    """Department scope. Empty = org-wide (visible to every member).
    Non-empty = visible only to members whose `departmentId` is in
    this set. Mirrored onto the legacy scope/scopeId pair for
    back-compat reads (scope='organization' when empty,
    scope='department' otherwise).
    """

class AiToolEntryTeam(Base):
    """@deprecated Per-team scope binding for AiToolEntry. Superseded by
    AiToolEntryDepartment; kept so the rename migration is non-destructive.
    No new rows are written here.
    """

    __tablename__ = "AiToolEntryTeam"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    entry_id: Mapped[str] = mapped_column("entryId", ForeignKey("AiToolEntry.id", ondelete="CASCADE", onupdate="CASCADE", name="AiToolEntryTeam_entryId_fkey"))
    team_id: Mapped[str] = mapped_column("teamId", ForeignKey("Team.id", ondelete="CASCADE", onupdate="CASCADE", name="AiToolEntryTeam_teamId_fkey"))

    __table_args__ = (
        Index("AiToolEntryTeam_teamId_idx", "teamId"),
        Index("AiToolEntryTeam_entryId_teamId_key", "entryId", "teamId", unique=True),
    )

    entry: Mapped["AiToolEntry"] = relationship(back_populates="teams")
    team: Mapped["Team"] = relationship(back_populates="ai_tool_entry_teams")

class AiToolEntryDepartment(Base):
    """Per-department scope binding for AiToolEntry. Empty set = org-wide.
    Non-empty = entry visible only to members of those departments (an
    OrganizationUser whose `departmentId` is in the set). Mirrors
    AiToolEntryTeam's shape. `departmentId` is a plain column ref
    (relationMode="prisma"), matching how Department is referenced
    elsewhere; the service validates the id belongs to the org.
    """

    __tablename__ = "AiToolEntryDepartment"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    entry_id: Mapped[str] = mapped_column("entryId", ForeignKey("AiToolEntry.id", ondelete="CASCADE", onupdate="CASCADE", name="AiToolEntryDepartment_entryId_fkey"))
    department_id: Mapped[str] = mapped_column("departmentId")

    __table_args__ = (
        Index("AiToolEntryDepartment_departmentId_idx", "departmentId"),
        Index("AiToolEntryDepartment_entryId_departmentId_key", "entryId", "departmentId", unique=True),
    )

    entry: Mapped["AiToolEntry"] = relationship(back_populates="departments")

class IngestionTemplate(Base):
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

    __tablename__ = "IngestionTemplate"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[Optional[str]] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="IngestionTemplate_organizationId_fkey"))
    slug: Mapped[str]
    """Stable slug - unique within (organizationId, slug) scope. Generated
    from displayName + nanoid(6) on create for org-authored; canonical
    values for platform-published (`claude_code` / `cursor` /
    `claude_cowork` / `raw_otlp_advanced`).
    """
    source_type: Mapped[str] = mapped_column("sourceType")
    """Mirrors the slug for v1 platform-published rows. Kept as a separate
    column so a future v2 may publish two templates that share an
    upstream sourceType (e.g. claude_code-personal vs claude_code-team)
    without a migration. String not enum - extensibility.
    """
    display_name: Mapped[str] = mapped_column("displayName")
    description: Mapped[Optional[str]]
    icon_asset: Mapped[Optional[str]] = mapped_column("iconAsset")
    """Icon source - single column, prefix-discriminated, mirrors
    AiToolEntry.iconAsset:
    `preset:<kind>`                 (built-in slot)
    `data:image/svg+xml;base64,…`   (admin-uploaded for org-authored)
    null                            → fall back to slug-default in UI
    """
    credential_schema: Mapped[Optional[str]] = mapped_column("credentialSchema")
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
    ottl_rules: Mapped[str] = mapped_column("ottlRules", default="", server_default=text("''"))
    """Admin/platform-authored OTTL statements. Applied at receive time
    AFTER the receiver re-stamps the binding's authoritative principal
    + provenance keys (see `protectedTemplateAttributeKeys` 19-key
    closed list). Templates write canonical OTTL outputs (gen_ai.*,
    cost.usd, response.model) - they CANNOT write attribution or
    provenance keys; the receiver guard restores those post-OTTL.
    """
    platform_published: Mapped[bool] = mapped_column("platformPublished", default=False, server_default=text("false"))
    """True for the v1 platform-published default templates
    (organizationId IS NULL). Future org-authored rows write false.
    Drives the admin UI "READ-ONLY view + Request a custom template"
    surface - admins consume + opt-out only in v1.
    """
    enabled: Mapped[bool] = mapped_column(default=True, server_default=text("true"))
    """Admin/platform enable. False = hidden from /me catalog. Existing
    ingestion keys continue to land traces; new installs are blocked.
    """
    archived_at: Mapped[Optional[datetime]] = mapped_column("archivedAt")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    created_by_id: Mapped[Optional[str]] = mapped_column("createdById")
    updated_by_id: Mapped[Optional[str]] = mapped_column("updatedById")

    __table_args__ = (
        Index("IngestionTemplate_organizationId_archivedAt_idx", "organizationId", "archivedAt"),
        Index("IngestionTemplate_sourceType_idx", "sourceType"),
        Index("IngestionTemplate_organizationId_slug_key", "organizationId", "slug", unique=True),
    )

    organization: Mapped[Optional["Organization"]] = relationship(back_populates="ingestion_templates")

class GatewayConnectUpstream(Base):
    __tablename__ = "GatewayConnectUpstream"

    organization_id: Mapped[str] = mapped_column("organizationId", primary_key=True)
    base_url: Mapped[str] = mapped_column("baseUrl")
    encrypted_token: Mapped[str] = mapped_column("encryptedToken")
    instance_id: Mapped[str] = mapped_column("instanceId")
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

class GatewayTraceExportKey(Base):
    __tablename__ = "GatewayTraceExportKey"

    project_id: Mapped[str] = mapped_column("projectId", primary_key=True)
    api_key_id: Mapped[str] = mapped_column("apiKeyId")
    encrypted_token: Mapped[str] = mapped_column("encryptedToken")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

class GatewayBudget(Base):
    __tablename__ = "GatewayBudget"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId")
    scope_type: Mapped[str] = mapped_column("scopeType", Enum("ORGANIZATION", "TEAM", "PROJECT", "VIRTUAL_KEY", "PRINCIPAL", "GROUP", "ATTRIBUTED_USER", name="GatewayBudgetScopeType"))
    scope_id: Mapped[str] = mapped_column("scopeId")
    provider_key: Mapped[Optional[str]] = mapped_column("providerKey")
    name: Mapped[str]
    description: Mapped[Optional[str]]
    window: Mapped[str] = mapped_column(Enum("MINUTE", "HOUR", "DAY", "WEEK", "MONTH", "TOTAL", "MANUAL", name="GatewayBudgetWindow"))
    limit_usd: Mapped[DecimalType] = mapped_column("limitUsd", Numeric(precision=18, scale=6))
    on_breach: Mapped[str] = mapped_column("onBreach", Enum("BLOCK", "WARN", name="GatewayBudgetBreachAction"), default="BLOCK", server_default=text("'BLOCK'"))
    timezone_: Mapped[Optional[str]] = mapped_column("timezone")
    external_id: Mapped[Optional[str]] = mapped_column("externalId")
    metadata_: Mapped[dict[str, Any]] = mapped_column("metadata", JSONB, default=lambda: {}, server_default=text("'{}'"))
    spent_usd: Mapped[DecimalType] = mapped_column("spentUsd", Numeric(precision=18, scale=6), default=DecimalType("0"), server_default=text("0"))
    current_period_started_at: Mapped[datetime] = mapped_column("currentPeriodStartedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    resets_at: Mapped[datetime] = mapped_column("resetsAt")
    last_reset_at: Mapped[Optional[datetime]] = mapped_column("lastResetAt")
    cycle_anchor_at: Mapped[Optional[datetime]] = mapped_column("cycleAnchorAt")
    archived_at: Mapped[Optional[datetime]] = mapped_column("archivedAt")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    created_by_id: Mapped[str] = mapped_column("createdById", ForeignKey("User.id", ondelete="RESTRICT", onupdate="CASCADE", name="GatewayBudget_createdById_fkey"))
    managed_by_virtual_key_id: Mapped[Optional[str]] = mapped_column("managedByVirtualKeyId")

    __table_args__ = (
        Index("GatewayBudget_organizationId_idx", "organizationId"),
        Index("GatewayBudget_scopeType_scopeId_providerKey_idx", "scopeType", "scopeId", "providerKey"),
        Index("GatewayBudget_managedByVirtualKeyId_idx", "managedByVirtualKeyId"),
        Index("GatewayBudget_resetsAt_idx", "resetsAt"),
        Index("GatewayBudget_organizationId_externalId_key", "organizationId", "externalId", unique=True),
    )

    created_by: Mapped["User"] = relationship(back_populates="gateway_budgets_created")
    ledger_entries: Mapped[list["GatewayBudgetLedger"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="budget")
    bucket_boundaries: Mapped[list["GatewayBudgetBucketBoundary"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="budget")

class GatewayBudgetBucketBoundary(Base):
    __tablename__ = "GatewayBudgetBucketBoundary"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId")
    budget_id: Mapped[str] = mapped_column("budgetId", ForeignKey("GatewayBudget.id", ondelete="CASCADE", onupdate="CASCADE", name="GatewayBudgetBucketBoundary_budgetId_fkey"))
    bucket_scope_id: Mapped[str] = mapped_column("bucketScopeId")
    period_started_at: Mapped[datetime] = mapped_column("periodStartedAt")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("GatewayBudgetBucketBoundary_organizationId_idx", "organizationId"),
        Index("GatewayBudgetBucketBoundary_budgetId_bucketScopeId_key", "budgetId", "bucketScopeId", unique=True),
    )

    budget: Mapped["GatewayBudget"] = relationship(back_populates="bucket_boundaries")

class GatewayBudgetLedger(Base):
    __tablename__ = "GatewayBudgetLedger"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    budget_id: Mapped[str] = mapped_column("budgetId", ForeignKey("GatewayBudget.id", ondelete="CASCADE", onupdate="CASCADE", name="GatewayBudgetLedger_budgetId_fkey"))
    virtual_key_id: Mapped[str] = mapped_column("virtualKeyId", ForeignKey("VirtualKey.id", ondelete="CASCADE", onupdate="CASCADE", name="GatewayBudgetLedger_virtualKeyId_fkey"))
    model_provider_id: Mapped[Optional[str]] = mapped_column("modelProviderId")
    provider_key: Mapped[Optional[str]] = mapped_column("providerKey")
    gateway_request_id: Mapped[str] = mapped_column("gatewayRequestId")
    amount_usd: Mapped[DecimalType] = mapped_column("amountUsd", Numeric(precision=18, scale=6))
    tokens_input: Mapped[int] = mapped_column("tokensInput", default=0, server_default=text("0"))
    tokens_output: Mapped[int] = mapped_column("tokensOutput", default=0, server_default=text("0"))
    tokens_cache_read: Mapped[int] = mapped_column("tokensCacheRead", default=0, server_default=text("0"))
    tokens_cache_write: Mapped[int] = mapped_column("tokensCacheWrite", default=0, server_default=text("0"))
    model: Mapped[str]
    duration_ms: Mapped[Optional[int]] = mapped_column("durationMs")
    status: Mapped[str] = mapped_column(Enum("SUCCESS", "PROVIDER_ERROR", "BLOCKED_BY_GUARDRAIL", "CANCELLED", name="GatewayBudgetLedgerStatus"), default="SUCCESS", server_default=text("'SUCCESS'"))
    occurred_at: Mapped[datetime] = mapped_column("occurredAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("GatewayBudgetLedger_virtualKeyId_idx", "virtualKeyId"),
        Index("GatewayBudgetLedger_gatewayRequestId_idx", "gatewayRequestId"),
        Index("GatewayBudgetLedger_occurredAt_idx", "occurredAt"),
        Index("GatewayBudgetLedger_budgetId_gatewayRequestId_key", "budgetId", "gatewayRequestId", unique=True),
    )

    budget: Mapped["GatewayBudget"] = relationship(back_populates="ledger_entries")
    virtual_key: Mapped["VirtualKey"] = relationship(back_populates="ledger_entries")

class GatewayChangeEvent(Base):
    __tablename__ = "GatewayChangeEvent"

    revision: Mapped[int] = mapped_column(BigInteger, primary_key=True, autoincrement=True)
    organization_id: Mapped[str] = mapped_column("organizationId", ForeignKey("Organization.id", ondelete="CASCADE", onupdate="CASCADE", name="GatewayChangeEvent_organizationId_fkey"))
    kind: Mapped[str] = mapped_column(Enum("VK_CREATED", "VK_CONFIG_UPDATED", "VK_REVOKED", "VK_ROTATED", "VK_DISABLED", "VK_ENABLED", "BUDGET_CREATED", "BUDGET_UPDATED", "BUDGET_DELETED", "MODEL_PROVIDER_UPDATED", "CACHE_RULE_CREATED", "CACHE_RULE_UPDATED", "CACHE_RULE_DELETED", "ROUTING_POLICY_UPDATED", "ROUTING_POLICY_DELETED", name="GatewayChangeEventKind"))
    virtual_key_id: Mapped[Optional[str]] = mapped_column("virtualKeyId")
    budget_id: Mapped[Optional[str]] = mapped_column("budgetId")
    model_provider_id: Mapped[Optional[str]] = mapped_column("modelProviderId")
    project_id: Mapped[Optional[str]] = mapped_column("projectId", ForeignKey("Project.id", ondelete="SET NULL", onupdate="CASCADE", name="GatewayChangeEvent_projectId_fkey"))
    payload: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("GatewayChangeEvent_organizationId_revision_idx", "organizationId", "revision"),
        Index("GatewayChangeEvent_projectId_revision_idx", "projectId", "revision"),
        Index("GatewayChangeEvent_virtualKeyId_idx", "virtualKeyId"),
        Index("GatewayChangeEvent_budgetId_idx", "budgetId"),
    )

    organization: Mapped["Organization"] = relationship(back_populates="gateway_change_events")
    project: Mapped[Optional["Project"]] = relationship(back_populates="gateway_change_events")

class GatewayCacheRule(Base):
    __tablename__ = "GatewayCacheRule"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId")
    name: Mapped[str]
    description: Mapped[Optional[str]]
    priority: Mapped[int] = mapped_column(default=100, server_default=text("100"))
    enabled: Mapped[bool] = mapped_column(default=True, server_default=text("true"))
    matchers: Mapped[dict[str, Any]] = mapped_column(JSONB)
    action: Mapped[dict[str, Any]] = mapped_column(JSONB)
    mode_enum: Mapped[str] = mapped_column("modeEnum", Enum("RESPECT", "FORCE", "DISABLE", name="GatewayCacheRuleMode"))
    archived_at: Mapped[Optional[datetime]] = mapped_column("archivedAt")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
    created_by_id: Mapped[str] = mapped_column("createdById")

    __table_args__ = (
        Index("GatewayCacheRule_organizationId_archivedAt_priority_idx", "organizationId", "archivedAt", "priority"),
        Index("GatewayCacheRule_organizationId_enabled_priority_idx", "organizationId", "enabled", "priority"),
    )

class FeatureFlag(Base):
    __tablename__ = "FeatureFlag"

    key: Mapped[str] = mapped_column(primary_key=True)
    enabled: Mapped[bool]
    rules: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))
    last_edited_by: Mapped[Optional[str]] = mapped_column("lastEditedBy")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

class FeatureFlagExperimentSetting(Base):
    """One experiment setting per (flag, subject). A USER row is that person's
    own opt-in; an ORGANIZATION or PROJECT row is an owner's explicit policy
    for that exact scope. No row means inherit.
    """

    __tablename__ = "FeatureFlagExperimentSetting"

    flag_key: Mapped[str] = mapped_column("flagKey", primary_key=True)
    subject_type: Mapped[str] = mapped_column("subjectType", Enum("USER", "ORGANIZATION", "PROJECT", name="FeatureFlagExperimentSubjectType"), primary_key=True)
    subject_id: Mapped[str] = mapped_column("subjectId", primary_key=True)
    enabled: Mapped[bool]
    changed_by_user_id: Mapped[Optional[str]] = mapped_column("changedByUserId")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("FeatureFlagExperimentSetting_subjectType_subjectId_idx", "subjectType", "subjectId"),
    )

class ProcessManagerInstance(Base):
    """Durable state for a long-running event-sourced process."""

    __tablename__ = "ProcessManagerInstance"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    process_name: Mapped[str] = mapped_column("processName")
    project_id: Mapped[str] = mapped_column("projectId")
    process_key: Mapped[str] = mapped_column("processKey")
    tenant_id: Mapped[str] = mapped_column("tenantId")
    user_id: Mapped[Optional[str]] = mapped_column("userId")
    state: Mapped[dict[str, Any]] = mapped_column(JSONB)
    revision: Mapped[int]
    next_wake_at: Mapped[Optional[datetime]] = mapped_column("nextWakeAt", UtcDateTimeTz(precision=3))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", UtcDateTimeTz(precision=3))

    __table_args__ = (
        Index("ProcessManagerInstance_nextWakeAt_idx", "nextWakeAt"),
        Index("ProcessManagerInstance_processName_projectId_processKey_key", "processName", "projectId", "processKey", unique=True),
    )

class ProcessManagerInbox(Base):
    """Idempotency marker for a source event consumed by a process definition."""

    __tablename__ = "ProcessManagerInbox"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    process_name: Mapped[str] = mapped_column("processName")
    project_id: Mapped[str] = mapped_column("projectId")
    process_key: Mapped[str] = mapped_column("processKey")
    tenant_id: Mapped[str] = mapped_column("tenantId")
    source_event_id: Mapped[str] = mapped_column("sourceEventId")
    """The raw source event id, kept UNINDEXED for diagnostics. A caller composes
    this and its length is that domain's business, so it must never carry a
    btree constraint — see `sourceEventKey`.
    """
    source_event_key: Mapped[Optional[str]] = mapped_column("sourceEventKey")
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
    consumed_at: Mapped[datetime] = mapped_column("consumedAt", UtcDateTimeTz(precision=3))

    __table_args__ = (
        Index("ProcessManagerInbox_processName_projectId_processKey_consum_idx", "processName", "projectId", "processKey", "consumedAt"),
        Index("ProcessManagerInbox_consumedAt_idx", "consumedAt"),
        Index("ProcessManagerInbox_processName_projectId_sourceEventKey_key", "processName", "projectId", "sourceEventKey", unique=True),
    )

class ProcessManagerOutbox(Base):
    """Transactional intents emitted by a durable process transition."""

    __tablename__ = "ProcessManagerOutbox"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    process_name: Mapped[str] = mapped_column("processName")
    project_id: Mapped[str] = mapped_column("projectId")
    process_key: Mapped[str] = mapped_column("processKey")
    tenant_id: Mapped[str] = mapped_column("tenantId")
    user_id: Mapped[Optional[str]] = mapped_column("userId")
    message_key: Mapped[str] = mapped_column("messageKey")
    intent_type: Mapped[str] = mapped_column("intentType")
    payload: Mapped[dict[str, Any]] = mapped_column(JSONB)
    trace_carrier: Mapped[dict[str, Any]] = mapped_column("traceCarrier", JSONB)
    source_event_id: Mapped[Optional[str]] = mapped_column("sourceEventId")
    status: Mapped[str] = mapped_column(Enum("pending", "dispatched", "dead", "discarded", name="ProcessManagerOutboxStatus"), default="pending", server_default=text("'pending'"))
    attempts: Mapped[int] = mapped_column(default=0, server_default=text("0"))
    next_attempt_at: Mapped[datetime] = mapped_column("nextAttemptAt", UtcDateTimeTz(precision=3))
    leased_until: Mapped[Optional[datetime]] = mapped_column("leasedUntil", UtcDateTimeTz(precision=3))
    lease_token: Mapped[Optional[str]] = mapped_column("leaseToken")
    dispatched_at: Mapped[Optional[datetime]] = mapped_column("dispatchedAt", UtcDateTimeTz(precision=3))
    created_at: Mapped[datetime] = mapped_column("createdAt", UtcDateTimeTz(precision=3))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", UtcDateTimeTz(precision=3))

    __table_args__ = (
        Index("ProcessManagerOutbox_leaseToken_key", "leaseToken", unique=True),
        Index("ProcessManagerOutbox_processName_projectId_processKey_creat_idx", "processName", "projectId", "processKey", "createdAt"),
        Index("ProcessManagerOutbox_status_nextAttemptAt_leasedUntil_idx", "status", "nextAttemptAt", "leasedUntil"),
        Index("ProcessManagerOutbox_status_processName_nextAttemptAt_idx", "status", "processName", "nextAttemptAt"),
        Index("ProcessManagerOutbox_processName_projectId_messageKey_key", "processName", "projectId", "messageKey", unique=True),
    )

    attempt_log: Mapped[list["ProcessManagerOutboxAttempt"]] = relationship(cascade="all, delete", passive_deletes=True, back_populates="outbox")

class ProcessManagerOutboxAttempt(Base):
    """Why an outbox delivery failed, one row per FAILED attempt
    (specs/ops/dead-letter-recovery.feature). Successes write nothing, so the
    table grows with trouble, not with traffic. Rows die with their message:
    the cascade delete is the retention story, riding the existing sweep.
    """

    __tablename__ = "ProcessManagerOutboxAttempt"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    outbox_id: Mapped[str] = mapped_column("outboxId", ForeignKey("ProcessManagerOutbox.id", ondelete="CASCADE", onupdate="CASCADE", name="ProcessManagerOutboxAttempt_outboxId_fkey"))
    project_id: Mapped[str] = mapped_column("projectId")
    attempt: Mapped[int]
    """1-based attempt number, as the dispatcher counts it."""
    occurred_at: Mapped[datetime] = mapped_column("occurredAt", UtcDateTimeTz(precision=3))
    outcome: Mapped[str] = mapped_column(Enum("retry_scheduled", "dead", name="ProcessManagerOutboxAttemptOutcome"))
    error_type: Mapped[str] = mapped_column("errorType")
    error_message: Mapped[str] = mapped_column("errorMessage")
    """The safe failure diagnostic the dispatcher already computes for its
    span — never a raw provider body.
    """
    retry_after_ms: Mapped[Optional[int]] = mapped_column("retryAfterMs")

    __table_args__ = (
        Index("ProcessManagerOutboxAttempt_outboxId_attempt_idx", "outboxId", "attempt"),
    )

    outbox: Mapped["ProcessManagerOutbox"] = relationship(back_populates="attempt_log")

class TopicClusteringRunProjection(Base):
    """Per-project topic clustering run status (ADR-051): the rebuildable
    operational read model behind the settings page. One row per project;
    folded from run_completed/run_failed events, never written directly.
    """

    __tablename__ = "TopicClusteringRunProjection"

    id: Mapped[str] = mapped_column(primary_key=True)
    """KSUID (`topicrun_…`), generated by the repository on first upsert."""
    project_id: Mapped[str] = mapped_column("projectId")
    last_requested_at: Mapped[Optional[float]] = mapped_column("lastRequestedAt")
    last_request_trigger: Mapped[Optional[str]] = mapped_column("lastRequestTrigger")
    last_run_at: Mapped[Optional[float]] = mapped_column("lastRunAt")
    last_run_outcome: Mapped[Optional[str]] = mapped_column("lastRunOutcome")
    """completed | skipped | failed"""
    last_run_mode: Mapped[Optional[str]] = mapped_column("lastRunMode")
    last_run_skipped_reason: Mapped[Optional[str]] = mapped_column("lastRunSkippedReason")
    last_run_error: Mapped[Optional[str]] = mapped_column("lastRunError")
    last_run_error_code: Mapped[Optional[str]] = mapped_column("lastRunErrorCode")
    last_run_error_user_actionable: Mapped[bool] = mapped_column("lastRunErrorUserActionable", default=False, server_default=text("false"))
    """True when the customer can resolve the failure (credentials, quota)."""
    last_run_traces_processed: Mapped[int] = mapped_column("lastRunTracesProcessed", default=0, server_default=text("0"))
    last_run_topics_count: Mapped[int] = mapped_column("lastRunTopicsCount", default=0, server_default=text("0"))
    last_run_subtopics_count: Mapped[int] = mapped_column("lastRunSubtopicsCount", default=0, server_default=text("0"))
    last_run_pages: Mapped[int] = mapped_column("lastRunPages", default=0, server_default=text("0"))
    in_progress_run_id: Mapped[Optional[str]] = mapped_column("inProgressRunId")
    in_progress_traces: Mapped[int] = mapped_column("inProgressTraces", default=0, server_default=text("0"))
    in_progress_pages: Mapped[int] = mapped_column("inProgressPages", default=0, server_default=text("0"))
    in_progress_started_at: Mapped[Optional[float]] = mapped_column("inProgressStartedAt")
    """Business time the in-progress run opened; bounds how long the read
    model may report it as running (mirrors the process's stale-run guard).
    """
    created_at: Mapped[float] = mapped_column("createdAt")
    updated_at: Mapped[float] = mapped_column("updatedAt")
    occurred_at: Mapped[float] = mapped_column("occurredAt")
    """Business time of the latest applied event."""
    accepted_at: Mapped[float] = mapped_column("acceptedAt")
    """Event-log acceptance time. Together with lastEventId this is the cursor."""
    last_event_id: Mapped[str] = mapped_column("lastEventId")
    projection_version: Mapped[str] = mapped_column("projectionVersion")

    __table_args__ = (
        Index("TopicClusteringRunProjection_projectId_key", "projectId", unique=True),
    )

class TopicModelProjection(Base):
    """Cursor row for the topic-model projection (ADR-051). The model itself is
    the `Topic` table, written through by the projection; this row only
    records how far the fold has applied the event log.
    """

    __tablename__ = "TopicModelProjection"

    id: Mapped[str] = mapped_column(primary_key=True)
    """KSUID (`topicmodel_…`), generated by the repository on first upsert."""
    project_id: Mapped[str] = mapped_column("projectId")
    created_at: Mapped[float] = mapped_column("createdAt")
    updated_at: Mapped[float] = mapped_column("updatedAt")
    occurred_at: Mapped[float] = mapped_column("occurredAt")
    """Business time of the latest applied event."""
    accepted_at: Mapped[float] = mapped_column("acceptedAt")
    """Event-log acceptance time. Together with lastEventId this is the cursor."""
    last_event_id: Mapped[str] = mapped_column("lastEventId")
    projection_version: Mapped[str] = mapped_column("projectionVersion")

    __table_args__ = (
        Index("TopicModelProjection_projectId_key", "projectId", unique=True),
    )

class TopicClusteringRunHistoryProjection(Base):
    """Per-project topic clustering run history (audit read model, ADR-051).
    One row per project; `runs` is the bounded newest-first entry list.
    Rebuildable by replaying the event log.
    """

    __tablename__ = "TopicClusteringRunHistoryProjection"

    id: Mapped[str] = mapped_column(primary_key=True)
    """KSUID (`topicrunhist_…`), generated by the repository on first upsert."""
    project_id: Mapped[str] = mapped_column("projectId")
    runs: Mapped[dict[str, Any]] = mapped_column(JSONB)
    created_at: Mapped[float] = mapped_column("createdAt")
    updated_at: Mapped[float] = mapped_column("updatedAt")
    occurred_at: Mapped[float] = mapped_column("occurredAt")
    """Business time of the latest applied event."""
    accepted_at: Mapped[float] = mapped_column("acceptedAt")
    """Event-log acceptance time. Together with lastEventId this is the cursor."""
    last_event_id: Mapped[str] = mapped_column("lastEventId")
    projection_version: Mapped[str] = mapped_column("projectionVersion")

    __table_args__ = (
        Index("TopicClusteringRunHistoryProjection_projectId_key", "projectId", unique=True),
    )

class IngestionPullRunProjection(Base):
    """Rebuildable per-source pull checkpoint and operator-facing run status."""

    __tablename__ = "IngestionPullRunProjection"

    id: Mapped[str] = mapped_column(primary_key=True)
    source_id: Mapped[str] = mapped_column("sourceId")
    project_id: Mapped[str] = mapped_column("projectId")
    enabled: Mapped[bool] = mapped_column(default=False, server_default=text("false"))
    cron: Mapped[Optional[str]]
    cursor: Mapped[Optional[str]]
    last_run_at: Mapped[Optional[float]] = mapped_column("lastRunAt")
    last_run_outcome: Mapped[Optional[str]] = mapped_column("lastRunOutcome")
    last_run_event_count: Mapped[int] = mapped_column("lastRunEventCount", default=0, server_default=text("0"))
    last_run_error: Mapped[Optional[str]] = mapped_column("lastRunError")
    last_run_error_code: Mapped[Optional[str]] = mapped_column("lastRunErrorCode")
    consecutive_errors: Mapped[int] = mapped_column("consecutiveErrors", default=0, server_default=text("0"))
    last_success_at: Mapped[Optional[float]] = mapped_column("lastSuccessAt")
    """When a run last completed without erroring. Backfilled by replay, which
    is why the projection version moved with this column.
    """
    last_run_scheduled_for: Mapped[Optional[float]] = mapped_column("lastRunScheduledFor")
    """Identity of the run whose outcome this row reflects. Outcomes from a
    superseded run arrive with a smaller value and are ignored, so a late
    completion cannot regress Cursor or the status fields.
    """
    last_read_through_at: Mapped[Optional[float]] = mapped_column("lastReadThroughAt")
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
    last_run_completeness: Mapped[Optional[str]] = mapped_column("lastRunCompleteness")
    last_agents_listing_at: Mapped[Optional[float]] = mapped_column("lastAgentsListingAt")
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
    last_agents_listing_outcome: Mapped[Optional[str]] = mapped_column("lastAgentsListingOutcome")
    last_agents_listing_count: Mapped[Optional[int]] = mapped_column("lastAgentsListingCount")
    last_agents_listing_reason: Mapped[Optional[str]] = mapped_column("lastAgentsListingReason")
    last_agents_listing_status: Mapped[Optional[int]] = mapped_column("lastAgentsListingStatus")
    last_people_listing_at: Mapped[Optional[float]] = mapped_column("lastPeopleListingAt")
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
    last_people_listing_outcome: Mapped[Optional[str]] = mapped_column("lastPeopleListingOutcome")
    last_people_directory_count: Mapped[Optional[int]] = mapped_column("lastPeopleDirectoryCount")
    last_people_withheld_count: Mapped[Optional[int]] = mapped_column("lastPeopleWithheldCount")
    last_people_listing_reason: Mapped[Optional[str]] = mapped_column("lastPeopleListingReason")
    last_people_listing_status: Mapped[Optional[int]] = mapped_column("lastPeopleListingStatus")
    created_at: Mapped[float] = mapped_column("createdAt")
    updated_at: Mapped[float] = mapped_column("updatedAt")
    occurred_at: Mapped[float] = mapped_column("occurredAt")
    accepted_at: Mapped[float] = mapped_column("acceptedAt")
    last_event_id: Mapped[str] = mapped_column("lastEventId")
    projection_version: Mapped[str] = mapped_column("projectionVersion")

    __table_args__ = (
        Index("IngestionPullRunProjection_sourceId_key", "sourceId", unique=True),
        Index("IngestionPullRunProjection_projectId_idx", "projectId"),
    )

class ScheduledJob(Base):
    __tablename__ = "ScheduledJob"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    project_id: Mapped[str] = mapped_column("projectId")
    target_type: Mapped[str] = mapped_column("targetType")
    target_id: Mapped[str] = mapped_column("targetId")
    cron: Mapped[str]
    timezone_: Mapped[str] = mapped_column("timezone")
    next_run_at: Mapped[datetime] = mapped_column("nextRunAt")
    last_slot: Mapped[Optional[datetime]] = mapped_column("lastSlot")
    current_slot: Mapped[Optional[datetime]] = mapped_column("currentSlot")
    attempts: Mapped[int] = mapped_column(default=0, server_default=text("0"))
    last_error: Mapped[Optional[str]] = mapped_column("lastError")
    active: Mapped[bool] = mapped_column(default=True, server_default=text("true"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("ScheduledJob_active_nextRunAt_idx", "active", "nextRunAt"),
        Index("ScheduledJob_projectId_idx", "projectId"),
        Index("ScheduledJob_targetType_targetId_key", "targetType", "targetId", unique=True),
    )

class BugReport(Base):
    __tablename__ = "BugReport"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    created_at: Mapped[datetime] = mapped_column("createdAt", UtcDateTimeTz(precision=3), default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    source: Mapped[str]
    kind: Mapped[str]
    title: Mapped[str]
    summary: Mapped[Optional[str]]
    session_data: Mapped[Optional[str]] = mapped_column("sessionData")
    session_truncated: Mapped[bool] = mapped_column("sessionTruncated", default=False, server_default=text("false"))
    agent: Mapped[Optional[str]]
    contact_email: Mapped[Optional[str]] = mapped_column("contactEmail")
    cli_version: Mapped[Optional[str]] = mapped_column("cliVersion")
    linked_project_id: Mapped[Optional[str]] = mapped_column("linkedProjectId")
    metadata_: Mapped[Optional[dict[str, Any]]] = mapped_column("metadata", JSONB(none_as_null=True))

    __table_args__ = (
        Index("BugReport_createdAt_idx", "createdAt"),
    )

class IdempotencyReceipt(Base):
    __tablename__ = "IdempotencyReceipt"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    scope_id: Mapped[str] = mapped_column("scopeId")
    key: Mapped[str]
    claim_id: Mapped[str] = mapped_column("claimId")
    request_fingerprint: Mapped[str] = mapped_column("requestFingerprint")
    response_status: Mapped[Optional[int]] = mapped_column("responseStatus")
    response_body: Mapped[Optional[str]] = mapped_column("responseBody")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    heartbeat_at: Mapped[datetime] = mapped_column("heartbeatAt")
    expires_at: Mapped[datetime] = mapped_column("expiresAt")

    __table_args__ = (
        Index("IdempotencyReceipt_expiresAt_idx", "expiresAt"),
        Index("IdempotencyReceipt_scopeId_key_key", "scopeId", "key", unique=True),
    )

class GatewayRealtimeSession(Base):
    __tablename__ = "GatewayRealtimeSession"

    id: Mapped[str] = mapped_column(primary_key=True)
    project_id: Mapped[str] = mapped_column("projectId")
    organization_id: Mapped[str] = mapped_column("organizationId")
    virtual_key_id: Mapped[str] = mapped_column("virtualKeyId")
    model_provider_id: Mapped[str] = mapped_column("modelProviderId")
    vendor: Mapped[str]
    agent_id: Mapped[Optional[str]] = mapped_column("agentId")
    model: Mapped[str]
    requested_model: Mapped[Optional[str]] = mapped_column("requestedModel")
    vendor_conversation_id: Mapped[Optional[str]] = mapped_column("vendorConversationId")
    status: Mapped[str] = mapped_column(Enum("OPEN", "CLOSED", "FAILED", "EXPIRED", name="GatewayRealtimeSessionStatus"), default="OPEN", server_default=text("'OPEN'"))
    minted_at: Mapped[datetime] = mapped_column("mintedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    closed_at: Mapped[Optional[datetime]] = mapped_column("closedAt")
    close_reason: Mapped[Optional[str]] = mapped_column("closeReason")
    trace_id: Mapped[Optional[str]] = mapped_column("traceId")
    vendor_cost_raw: Mapped[Optional[dict[str, Any]]] = mapped_column("vendorCostRaw", JSONB(none_as_null=True))
    kind: Mapped[Optional[str]]
    metering: Mapped[Optional[str]]
    credential_expires_at: Mapped[Optional[datetime]] = mapped_column("credentialExpiresAt")
    transcription_model: Mapped[Optional[str]] = mapped_column("transcriptionModel")
    end_user_id: Mapped[Optional[str]] = mapped_column("endUserId")
    last_report_at: Mapped[Optional[datetime]] = mapped_column("lastReportAt")
    reported_cost_nano_usd: Mapped[int] = mapped_column("reportedCostNanoUsd", BigInteger, default=0, server_default=text("0"))
    report_count: Mapped[int] = mapped_column("reportCount", default=0, server_default=text("0"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("GatewayRealtimeSession_virtualKeyId_status_mintedAt_idx", "virtualKeyId", "status", "mintedAt"),
        Index("GatewayRealtimeSession_projectId_status_mintedAt_idx", "projectId", "status", "mintedAt"),
        Index("GatewayRealtimeSession_status_metering_mintedAt_idx", "status", "metering", "mintedAt"),
        Index("GatewayRealtimeSession_vendor_vendorConversationId_key", "vendor", "vendorConversationId", unique=True),
    )

class GatewayRealtimeSessionReport(Base):
    __tablename__ = "GatewayRealtimeSessionReport"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    session_id: Mapped[str] = mapped_column("sessionId")
    report_key: Mapped[str] = mapped_column("reportKey")
    project_id: Mapped[str] = mapped_column("projectId")
    model: Mapped[str]
    usage: Mapped[dict[str, Any]] = mapped_column(JSONB)
    cost_nano_usd: Mapped[int] = mapped_column("costNanoUsd", BigInteger)
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("GatewayRealtimeSessionReport_projectId_idx", "projectId"),
        Index("GatewayRealtimeSessionReport_sessionId_reportKey_key", "sessionId", "reportKey", unique=True),
    )

class SystemMigrationTenantState(Base):
    __tablename__ = "SystemMigrationTenantState"

    migration_name: Mapped[str] = mapped_column("migrationName", primary_key=True)
    tenant_id: Mapped[str] = mapped_column("tenantId", primary_key=True)
    status: Mapped[str]
    report: Mapped[Optional[dict[str, Any]]] = mapped_column(JSONB(none_as_null=True))
    occurred_at: Mapped[datetime] = mapped_column("occurredAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("SystemMigrationTenantState_migrationName_status_idx", "migrationName", "status"),
    )

class SystemMigrationEnrollment(Base):
    __tablename__ = "SystemMigrationEnrollment"

    organization_id: Mapped[str] = mapped_column("organizationId", primary_key=True)
    migration_name: Mapped[str] = mapped_column("migrationName", primary_key=True)
    enrolled_by_user_id: Mapped[str] = mapped_column("enrolledByUserId")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("SystemMigrationEnrollment_migrationName_idx", "migrationName"),
    )

class StoredObject(Base):
    __tablename__ = "StoredObject"

    tenant_id: Mapped[str] = mapped_column("tenantId", primary_key=True)
    id: Mapped[str] = mapped_column(primary_key=True)
    status: Mapped[str]
    purpose: Mapped[str]
    owner_kind: Mapped[str] = mapped_column("ownerKind")
    owner_id: Mapped[str] = mapped_column("ownerId")
    filename: Mapped[str]
    sha256: Mapped[str]
    size_bytes: Mapped[int] = mapped_column("sizeBytes", BigInteger)
    media_type: Mapped[str] = mapped_column("mediaType")
    media_type_verified: Mapped[bool] = mapped_column("mediaTypeVerified", default=False, server_default=text("false"))
    storage_provider: Mapped[Optional[str]] = mapped_column("storageProvider")
    storage_destination_id: Mapped[Optional[str]] = mapped_column("storageDestinationId")
    storage_provider_relative_id: Mapped[Optional[str]] = mapped_column("storageProviderRelativeId")
    generation: Mapped[int] = mapped_column(default=0, server_default=text("0"))
    audiences: Mapped[dict[str, Any]] = mapped_column(JSONB)
    expires_at: Mapped[Optional[datetime]] = mapped_column("expiresAt")
    available_at: Mapped[Optional[datetime]] = mapped_column("availableAt")
    deleted_at: Mapped[Optional[datetime]] = mapped_column("deletedAt")
    source: Mapped[str]
    legacy_fingerprint: Mapped[Optional[str]] = mapped_column("legacyFingerprint")
    created_at: Mapped[datetime] = mapped_column("createdAt")
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("StoredObject_tenantId_status_id_idx", "tenantId", "status", "id"),
        Index("StoredObject_ownerKind_ownerId_status_idx", "ownerKind", "ownerId", "status"),
        Index("StoredObject_status_expiresAt_idx", "status", "expiresAt"),
    )

class DiscoveredPerson(Base):
    """A name-text seen on a provider's cost or audit rows: "the provider told us
    someone called m.silva exists". One row per provider-scoped identity.

    This is the unit of GDPR erasure (ADR-128 §9), and deliberately so: most of
    these people have no LangWatch login at all - contractors, seat holders who
    never signed in, anyone whose email a provider put on a cost row - so there
    is no platform user to key on and a user-deletion event can never fire for
    them.
    """

    __tablename__ = "DiscoveredPerson"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId")
    provider: Mapped[str]
    """"anthropic" | "openai" | "databricks" | ... - the provider that named them."""
    raw_actor_id: Mapped[str] = mapped_column("rawActorId")
    """What the provider calls them, verbatim. The join key onto the money rows'
    RawActorId. Replaced by a deterministic pseudonym on erasure, which keeps
    the row addressable for spend attribution without keeping the identifier.
    """
    display_text: Mapped[str] = mapped_column("displayText")
    """Name or email text as seen. Pseudonymized alongside rawActorId on erasure."""
    kind: Mapped[str]
    """person | service_account - determined deterministically per ADR-128 §10
    (a bare UUID in Databricks `executed_by` is a service account, never an
    agent), never guessed from the shape of a name.
    """
    department: Mapped[Optional[str]]
    """The department the provider's directory filed this person under, verbatim.
    An OBSERVATION, not the `Department` entity: it is free text a tenant owns
    on their side, it exists for people who hold no LangWatch account, and
    nothing rolls spend up by it. Spend attribution stays on `Department`, via
    the membership the match engine proves (directoryDepartmentSync.service).
    Null means no directory has named one - never "no department", because a
    blank field in the tenant's directory and a person no directory lists are
    the same null here.
    """
    first_seen_at: Mapped[datetime] = mapped_column("firstSeenAt")
    last_seen_at: Mapped[datetime] = mapped_column("lastSeenAt")
    erased_at: Mapped[Optional[datetime]] = mapped_column("erasedAt")
    """Set when this row's identity has been erased. The row survives so its
    spend stays attributable to *someone*; the identifier does not.
    """
    money_rows_pending_at: Mapped[Optional[datetime]] = mapped_column("moneyRowsPendingAt")
    """Set alongside `erasedAt` and cleared once the daily cost rows have been
    removed and their rebuild requested. Non-null means the ClickHouse half
    of the erasure did not finish — a crash between the delete and the
    replay leaves days short by the erased amount, and without this marker a
    re-run reads `erasedAt`, short-circuits, and reports a clean erasure that
    did no work. The re-run resumes at that step instead.
    """
    money_rebuild_since: Mapped[Optional[str]] = mapped_column("moneyRebuildSince")
    """The earliest day to rebuild from, recorded BEFORE the rows are deleted.
    Once they are gone there is nothing left to ask which days they were on,
    so a resume with no record of this could only replay everything or
    nothing. Null means nothing was rebuildable — either no day carried the
    identifier, or every one of them is older than the event log keeps.
    """
    suspended_at: Mapped[Optional[datetime]] = mapped_column("suspendedAt")
    """ADR-128 §12 conflict rule: automatic linking has stopped for this
    identity because evidence pointed at two candidates, or contradicted an
    existing link. Stored HERE rather than on the suggestion rows, because a
    halt on auto-linking that a recompute or a restart clears is not a halt.
    """
    suspended_reason: Mapped[Optional[str]] = mapped_column("suspendedReason")
    """What tripped the suspension, for the human who reviews it."""
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("DiscoveredPerson_organizationId_lastSeenAt_idx", "organizationId", "lastSeenAt"),
        Index("DiscoveredPerson_organizationId_provider_rawActorId_key", "organizationId", "provider", "rawActorId", unique=True),
    )

class DiscoveredAgent(Base):
    """The same shape as DiscoveredPerson for non-human actors - Genie spaces,
    Copilot bots, OpenAI projects - plus whatever provider-native metadata the
    source published. Discovering an agent grants it nothing either.
    """

    __tablename__ = "DiscoveredAgent"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId")
    provider: Mapped[str]
    raw_agent_id: Mapped[str] = mapped_column("rawAgentId")
    """Genie space / Copilot bot / OpenAI project id, verbatim."""
    display_text: Mapped[str] = mapped_column("displayText")
    metadata_: Mapped[dict[str, Any]] = mapped_column("metadata", JSONB)
    """Workspace, plan, and kin - provider-native descriptive fields only. Never
    credentials: this row is read by admin screens, not by any pull path.
    """
    first_seen_at: Mapped[datetime] = mapped_column("firstSeenAt")
    last_seen_at: Mapped[datetime] = mapped_column("lastSeenAt")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("DiscoveredAgent_organizationId_lastSeenAt_idx", "organizationId", "lastSeenAt"),
        Index("DiscoveredAgent_organizationId_provider_rawAgentId_key", "organizationId", "provider", "rawAgentId", unique=True),
    )

class AgentRegistryEntry(Base):
    """One agent in the organisation's governance registry, from any source governance reads
    (enterprise/modules/governance/specs/agent-registry.feature). `entry` is the latest
    machine-readable snapshot; `version` grows by one on every material change.
    """

    __tablename__ = "AgentRegistryEntry"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId")
    source_key: Mapped[str] = mapped_column("sourceKey")
    """Where the agent came from and its id there, e.g. `langwatch:<agentId>`."""
    source: Mapped[str]
    status: Mapped[str]
    version: Mapped[int]
    fingerprint: Mapped[str]
    entry: Mapped[dict[str, Any]] = mapped_column(JSONB)
    first_seen_at: Mapped[datetime] = mapped_column("firstSeenAt")
    last_seen_at: Mapped[Optional[datetime]] = mapped_column("lastSeenAt")
    removed_at: Mapped[Optional[datetime]] = mapped_column("removedAt")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)

    __table_args__ = (
        Index("AgentRegistryEntry_organizationId_updatedAt_idx", "organizationId", "updatedAt"),
        Index("AgentRegistryEntry_organizationId_sourceKey_key", "organizationId", "sourceKey", unique=True),
    )

class AgentRegistryEvent(Base):
    """The registry's append-only change feed: one row per registered, versioned, removed or
    tool-deprecated fact. `sequence` is the pull cursor; `deliveryRecordedAt` marks the webhook
    fan-out as recorded in governance's outbox.
    """

    __tablename__ = "AgentRegistryEvent"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId")
    sequence: Mapped[int] = mapped_column(BigInteger)
    type: Mapped[str]
    entry_id: Mapped[str] = mapped_column("entryId")
    from_version: Mapped[Optional[int]] = mapped_column("fromVersion")
    to_version: Mapped[Optional[int]] = mapped_column("toVersion")
    payload: Mapped[dict[str, Any]] = mapped_column(JSONB)
    occurred_at: Mapped[datetime] = mapped_column("occurredAt")
    delivery_recorded_at: Mapped[Optional[datetime]] = mapped_column("deliveryRecordedAt")

    __table_args__ = (
        Index("AgentRegistryEvent_organizationId_deliveryRecordedAt_idx", "organizationId", "deliveryRecordedAt"),
        Index("AgentRegistryEvent_organizationId_entryId_idx", "organizationId", "entryId"),
        Index("AgentRegistryEvent_organizationId_sequence_key", "organizationId", "sequence", unique=True),
    )

class AgentRegistryBaseline(Base):
    """An organisation's registry baseline: its first reconcile, written as already pushed so the
    agents it already had are not announced. Every reconcile after it pushes its changes.
    """

    __tablename__ = "AgentRegistryBaseline"

    organization_id: Mapped[str] = mapped_column("organizationId", primary_key=True)
    baselined_at: Mapped[datetime] = mapped_column("baselinedAt")

class IdentityMatch(Base):
    """The dated match table: "`user_abc123` on Anthropic is platform user #42,
    valid from March 3", closed when it stops being true.

    Dating is what makes a re-issued email address survivable: the leaver's link
    closes at offboarding, the new hire gets a new link, and January's spend
    stays with January's person. A correction closes the old row and opens a new
    one; nothing is ever silently rewritten.
    """

    __tablename__ = "IdentityMatch"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId")
    discovered_person_id: Mapped[str] = mapped_column("discoveredPersonId")
    """The provider-side identity this link is about."""
    user_id: Mapped[Optional[str]] = mapped_column("userId")
    """The platform user. Nullable, and blanked by erasure (ADR-128 §9 step 2) -
    the row and its dates remain, so the history of the link is not rewritten
    by the erasure either.
    """
    evidence_kind: Mapped[str] = mapped_column("evidenceKind")
    """directory_id | verified_email | human_confirmed. Exact verified-email
    equality is the primary evidence; a directory id corroborates but never
    stands alone (ADR-128 §12).
    """
    valid_from: Mapped[datetime] = mapped_column("validFrom")
    valid_to: Mapped[Optional[datetime]] = mapped_column("validTo")
    """Null means the link is open. Offboarding and corrections close it."""
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

    __table_args__ = (
        Index("IdentityMatch_organizationId_discoveredPersonId_idx", "organizationId", "discoveredPersonId"),
        Index("IdentityMatch_organizationId_userId_idx", "organizationId", "userId"),
    )

class GovernanceTenantHistory(Base):
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

    __tablename__ = "GovernanceTenantHistory"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId")
    tenant_id: Mapped[str] = mapped_column("tenantId")
    """A hidden governance project id this organization has written rows under."""
    first_used_at: Mapped[datetime] = mapped_column("firstUsedAt")
    last_used_at: Mapped[datetime] = mapped_column("lastUsedAt")

    __table_args__ = (
        Index("GovernanceTenantHistory_organizationId_idx", "organizationId"),
        Index("GovernanceTenantHistory_organizationId_tenantId_key", "organizationId", "tenantId", unique=True),
    )

class ErasedIdentifierSuppression(Base):
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

    __tablename__ = "ErasedIdentifierSuppression"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId")
    provider: Mapped[str]
    identifier_hash: Mapped[str] = mapped_column("identifierHash")
    erased_at: Mapped[datetime] = mapped_column("erasedAt")

    __table_args__ = (
        Index("ErasedIdentifierSuppression_organizationId_idx", "organizationId"),
        Index("ErasedIdentifierSuppression_organizationId_provider_identif_key", "organizationId", "provider", "identifierHash", unique=True),
    )

class IdentityMatchSuggestion(Base):
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

    __tablename__ = "IdentityMatchSuggestion"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    organization_id: Mapped[str] = mapped_column("organizationId")
    discovered_person_id: Mapped[str] = mapped_column("discoveredPersonId")
    user_id: Mapped[str] = mapped_column("userId")
    """The platform user this might be. Not nullable: a suggestion with no
    candidate is not a suggestion.
    """
    score: Mapped[float]
    """Similarity of the two name texts, after the prefilter, in [0, 1]. A
    display hint and an ordering key - never a threshold anything auto-links
    on, which is the whole distinction this table exists to hold.
    """
    computed_at: Mapped[datetime] = mapped_column("computedAt")

    __table_args__ = (
        Index("IdentityMatchSuggestion_organizationId_score_idx", "organizationId", "score"),
        Index("IdentityMatchSuggestion_organizationId_discoveredPersonId_u_key", "organizationId", "discoveredPersonId", "userId", unique=True),
    )

class ConnectedSeatChange(Base):
    """What a mid-term seat change of one connected license owes (ADR-141,
    section 7): the seats added on the reissued license, invoiced prorated to
    the end of the term.

    The row is written before the payment provider is called and completed
    after it, so a call that failed halfway is retried by the next seat
    invoicing pass without inventing a second invoice. One row per reissued license.
    """

    __tablename__ = "ConnectedSeatChange"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    license_id: Mapped[str] = mapped_column("licenseId")
    """The reissued `IssuedLicense` row that carries the new seat count. A
    plain column ref, like `replacesId`: relationMode="prisma" carries no
    foreign keys.
    """
    account_id: Mapped[Optional[str]] = mapped_column("accountId")
    """The `ConnectedBillingAccount` row the invoice is raised against. Null
    when the customer had no billing account (`not_onboarded`).
    """
    organization_id: Mapped[Optional[str]] = mapped_column("organizationId")
    """The customer organization the license belongs to."""
    changed_at: Mapped[datetime] = mapped_column("changedAt")
    """When the seats changed, which is the day the proration counts from."""
    added_seats: Mapped[int] = mapped_column("addedSeats")
    """Seats on the new license above the seats on the one it replaced."""
    unit_amount_cents: Mapped[int] = mapped_column("unitAmountCents")
    """What one added seat costs for the rest of the term, in cents of `currency`."""
    amount_cents: Mapped[int] = mapped_column("amountCents")
    """`unitAmountCents` times `addedSeats`. Zero when nothing is owed."""
    currency: Mapped[Optional[str]] = mapped_column(Enum("USD", "EUR", name="Currency"))
    """Null when no billing account named one (`not_onboarded`)."""
    stripe_invoice_id: Mapped[Optional[str]] = mapped_column("stripeInvoiceId")
    """The one-off invoice the added seats were billed on."""
    state: Mapped[str] = mapped_column(Enum("intent", "invoiced", "nothing_to_invoice", "not_onboarded", name="ConnectedSeatChangeState"))
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("ConnectedSeatChange_licenseId_key", "licenseId", unique=True),
        Index("ConnectedSeatChange_accountId_idx", "accountId"),
        Index("ConnectedSeatChange_organizationId_idx", "organizationId"),
        Index("ConnectedSeatChange_state_idx", "state"),
    )

class ConnectedStatement(Base):
    """The monthly usage statement one connected customer was sent (ADR-139,
    section 7). One row per account per month keeps the statement to one send,
    whatever runs the tick.
    """

    __tablename__ = "ConnectedStatement"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    account_id: Mapped[str] = mapped_column("accountId")
    """The `ConnectedBillingAccount` row. A plain column ref."""
    month: Mapped[datetime]
    """First day of the month the statement covers, in UTC."""
    sent_at: Mapped[datetime] = mapped_column("sentAt")

    __table_args__ = (
        Index("ConnectedStatement_accountId_idx", "accountId"),
        Index("ConnectedStatement_accountId_month_key", "accountId", "month", unique=True),
    )

class LicenseRenewalClaim(Base):
    """One renewed licence per paid Stripe invoice of a self-serve licence
    subscription (specs/licensing/self-serving-license-purchase.feature). A
    redelivered invoice event finds the row and mints nothing. A claim never
    delivered is taken over once `claimedAt` is older than the claim lease.
    """

    __tablename__ = "LicenseRenewalClaim"

    stripe_invoice_id: Mapped[str] = mapped_column("stripeInvoiceId", primary_key=True)
    stripe_subscription_id: Mapped[str] = mapped_column("stripeSubscriptionId")
    claim_token: Mapped[str] = mapped_column("claimToken")
    """Fences completion and release to the delivery that holds the claim."""
    claimed_at: Mapped[datetime] = mapped_column("claimedAt")
    issued_license: Mapped[Optional[dict[str, Any]]] = mapped_column("issuedLicense", JSONB(none_as_null=True))
    """The signed licence between minting and mailing, so a retry mails it again instead of minting another; cleared once delivered."""
    delivered_at: Mapped[Optional[datetime]] = mapped_column("deliveredAt")
    license_id: Mapped[Optional[str]] = mapped_column("licenseId")
    """The `licenseId` inside the delivered licence."""

class InstanceIdentity(Base):
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

    __tablename__ = "InstanceIdentity"

    id: Mapped[str] = mapped_column(primary_key=True, default="self", server_default=text("'self'"))
    """Fixed, so this table holds one row. A second insert loses the race on the
    primary key and re-reads the winner's id.
    """
    instance_id: Mapped[str] = mapped_column("instanceId")
    """A random UUID. Carries nothing about the customer."""
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    last_report_at: Mapped[Optional[datetime]] = mapped_column("lastReportAt")
    """When the usage report last reached LangWatch, and why the last attempt
    did not. Both are what the checkup page reads to say whether reporting
    works, which nothing recorded before: the sender logged success as soon
    as `fetch` resolved and never looked at the response.
    """
    last_report_error: Mapped[Optional[str]] = mapped_column("lastReportError")
    optional_metrics_opt_out: Mapped[bool] = mapped_column("optionalMetricsOptOut", default=False, server_default=text("false"))
    """The optional category of the usage report, switched off from the checkup
    page. With it off the report still says which release runs and how many
    projects the install carries, and says nothing about what anyone does
    with it (ADR-139, section 10).
    """
    hostname_opt_out: Mapped[bool] = mapped_column("hostnameOptOut", default=False, server_default=text("false"))
    """Hostname, which names the customer's own network and so has a switch of
    its own rather than travelling with its category.
    """
    upgrade_diagnostics_opt_out: Mapped[bool] = mapped_column("upgradeDiagnosticsOptOut", default=False, server_default=text("false"))
    """Sending a failed upgrade step's diagnostics to LangWatch on its own,
    switched off by the operator (upgrade-diagnostics-via-connect, decision 4).
    """

    __table_args__ = (
        Index("InstanceIdentity_instanceId_key", "instanceId", unique=True),
    )

class SelfHostedInstance(Base):
    """One self-hosted install, as its daily usage report describes it (ADR-139,
    section 10). The row is created by the first report that names an instance
    id and is updated by every report after that, so the backoffice reads one
    row per install rather than folding a history on every page load.

    The report carries no credential, so it identifies an install and never a
    customer. `organizationId` and `issuedLicenseId` are resolved from the
    license bound to this instance, which the install proved by presenting its
    token to license sync. A field in the posted body never reaches them.
    """

    __tablename__ = "SelfHostedInstance"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    instance_id: Mapped[str] = mapped_column("instanceId")
    """The UUID the install minted for itself. The join key for everything else."""
    first_seen_at: Mapped[datetime] = mapped_column("firstSeenAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    last_seen_at: Mapped[datetime] = mapped_column("lastSeenAt")
    version: Mapped[Optional[str]]
    """What the install says about itself."""
    install_method: Mapped[Optional[str]] = mapped_column("installMethod")
    chart_version: Mapped[Optional[str]] = mapped_column("chartVersion")
    hostname: Mapped[Optional[str]]
    """The host it answers on, when the operator left hostname reporting on."""
    environment: Mapped[Optional[str]]
    installed_at: Mapped[Optional[datetime]] = mapped_column("installedAt")
    """The day the install itself says it first came up, which predates the
    first report on an install that switched reporting on later.
    """
    report_schema_version: Mapped[Optional[int]] = mapped_column("reportSchemaVersion")
    organization_id: Mapped[Optional[str]] = mapped_column("organizationId")
    """The customer, and the license that named them. Plain column refs rather
    than relations, like `IssuedLicense.replacesId`: relationMode="prisma"
    carries no foreign keys, and the receiver resolves both from the license
    registry rather than from the report.
    """
    issued_license_id: Mapped[Optional[str]] = mapped_column("issuedLicenseId")
    user_email_domains: Mapped[Optional[dict[str, Any]]] = mapped_column("userEmailDomains", JSONB(none_as_null=True))
    """Domains with counts, as `{"acme.test": 14}`. Never an address. Null on an
    install that switched the optional category off.
    """
    user_domains: Mapped[list[str]] = mapped_column("userDomains", ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))
    """The same domains as the keys of `userEmailDomains`, flattened so an
    operator can search for one. A JSON object's keys are the domains
    themselves, and no query language addresses a key it does not already
    know the name of.
    """
    latest_report: Mapped[Optional[dict[str, Any]]] = mapped_column("latestReport", JSONB(none_as_null=True))
    """What the last report carried, so the list and the drawer read one row:
    the onboarding ladder, the usage counts and the operational block, under
    the same keys the dictionary declares.
    """
    optional_metrics_reported: Mapped[bool] = mapped_column("optionalMetricsReported", default=True, server_default=text("true"))
    """What the last report said the operator had switched off, read from which
    categories arrived rather than from a field the install could assert.
    """
    hostname_reported: Mapped[bool] = mapped_column("hostnameReported", default=True, server_default=text("true"))
    report_count: Mapped[int] = mapped_column("reportCount", default=0, server_default=text("0"))
    """How many reports have landed, and how many fields the last one carried
    that this release has no name for. A count above zero says the senders
    are ahead of the receiver.
    """
    last_unknown_fields: Mapped[int] = mapped_column("lastUnknownFields", default=0, server_default=text("0"))
    raised_signals: Mapped[list[str]] = mapped_column("raisedSignals", ARRAY(Text), nullable=True, default=lambda: [], server_default=text("ARRAY[]::text[]"))
    """The lead signals already raised for this install. A signal on this list
    is never raised again, which is what keeps a large install from
    reporting that it is large every morning.
    """
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

    __table_args__ = (
        Index("SelfHostedInstance_instanceId_key", "instanceId", unique=True),
        Index("SelfHostedInstance_lastSeenAt_idx", "lastSeenAt"),
        Index("SelfHostedInstance_organizationId_idx", "organizationId"),
        Index("SelfHostedInstance_issuedLicenseId_idx", "issuedLicenseId"),
    )

class SelfHostedInstanceDiagnostic(Base):
    """One upgrade diagnostics bundle a self-hosted install sent over Connect, on
    LangWatch Cloud. The bytes live in object storage under
    `upgrade-diagnostics/` and expire at 30 days by a lifecycle rule; this row
    points at them and feeds Cloud admin and the Slack signal. Plain column
    refs, no relation, as the usage report rows.
    """

    __tablename__ = "SelfHostedInstanceDiagnostic"

    id: Mapped[str] = mapped_column(primary_key=True)
    instance_id: Mapped[str] = mapped_column("instanceId")
    """The `SelfHostedInstance.instanceId` the licence caller presented."""
    step_id: Mapped[str] = mapped_column("stepId")
    step_description: Mapped[Optional[str]] = mapped_column("stepDescription")
    release: Mapped[Optional[str]]
    attempt: Mapped[int]
    trigger: Mapped[str]
    """`operator` or `automatic`."""
    byte_length: Mapped[int] = mapped_column("byteLength")
    object_key: Mapped[str] = mapped_column("objectKey")
    received_at: Mapped[datetime] = mapped_column("receivedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    expires_at: Mapped[datetime] = mapped_column("expiresAt")
    """When the lifecycle rule removes the bundle; the row is listed until then."""
    slack_announced_at: Mapped[Optional[datetime]] = mapped_column("slackAnnouncedAt")
    """Set on the report that posted this step and release's one Slack message."""

    __table_args__ = (
        Index("SelfHostedInstanceDiagnostic_instanceId_receivedAt_idx", "instanceId", "receivedAt"),
    )

class SelfHostedNotice(Base):
    """A notice LangWatch staff send every self-hosted install in Cloud's reply to its daily usage
    report (enterprise/modules/licensing/specs/self-hosted-notices.feature). Cloud only, and no
    tenant: every install receives every published notice that has not expired.
    """

    __tablename__ = "SelfHostedNotice"

    id: Mapped[str] = mapped_column(primary_key=True)
    severity: Mapped[str]
    """`info`, `warning` or `security`."""
    title: Mapped[str]
    body: Mapped[str]
    url: Mapped[Optional[str]]
    """An https link to read more; null where the notice names none."""
    fixed_in: Mapped[Optional[str]] = mapped_column("fixedIn")
    """The release that fixes what the notice is about; an install on it or newer hides it."""
    status: Mapped[str]
    """`draft`, `published` or `withdrawn`. Only a published notice that has not expired is sent."""
    expires_at: Mapped[Optional[datetime]] = mapped_column("expiresAt")
    published_at: Mapped[Optional[datetime]] = mapped_column("publishedAt")
    published_by: Mapped[Optional[str]] = mapped_column("publishedBy")
    """The staff member who published it, by user id."""
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, onupdate=utc_now)

class SelfHostedInstanceReport(Base):
    """One usage report, exactly as it arrived. Kept so a number on an install's
    row can be traced back to the report that set it, and so a question asked
    later can be answered over the history rather than over the latest snapshot.
    """

    __tablename__ = "SelfHostedInstanceReport"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    instance_id: Mapped[str] = mapped_column("instanceId")
    """The `SelfHostedInstance.instanceId` this report came from. A plain column
    ref for the same reason the instance row's refs are.
    """
    received_at: Mapped[datetime] = mapped_column("receivedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    version: Mapped[Optional[str]]
    report_schema_version: Mapped[Optional[int]] = mapped_column("reportSchemaVersion")
    unknown_fields: Mapped[int] = mapped_column("unknownFields", default=0, server_default=text("0"))
    """Fields this release had no name for. Counted, never stored by name."""
    payload: Mapped[dict[str, Any]] = mapped_column(JSONB)
    """The validated report, with every unknown field already dropped."""

    __table_args__ = (
        Index("SelfHostedInstanceReport_instanceId_receivedAt_idx", "instanceId", "receivedAt"),
        Index("SelfHostedInstanceReport_receivedAt_idx", "receivedAt"),
    )

class WebPushSubscription(Base):
    """One browser a person subscribed to Web Push, owned by notification. A
    person holds one row per device; the endpoint names the browser, so a
    browser that subscribes again replaces its own row.
    """

    __tablename__ = "WebPushSubscription"

    id: Mapped[str] = mapped_column(primary_key=True, default=lambda: nanoid.generate())
    user_id: Mapped[str] = mapped_column("userId")
    """The person the pushes are for. A plain column: the row leaves with the
    person through notification's own subscriber on the user lifecycle.
    """
    endpoint: Mapped[str]
    """The push service URL the browser handed out. Unique per browser."""
    p256dh: Mapped[str]
    """The browser's P-256 public key and auth secret, base64url, as the
    browser reports them. Payloads are encrypted to these.
    """
    auth: Mapped[str]
    user_agent: Mapped[Optional[str]] = mapped_column("userAgent")
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))
    last_success_at: Mapped[Optional[datetime]] = mapped_column("lastSuccessAt")
    """When the push service last accepted a push for this browser."""

    __table_args__ = (
        Index("WebPushSubscription_endpoint_key", "endpoint", unique=True),
        Index("WebPushSubscription_userId_idx", "userId"),
    )

class WebPushVapidKey(Base):
    """This installation's VAPID key pair, generated on first use so Web Push works
    without configuration. Fixed id, so the table holds one row: a second
    insert loses the race on the primary key and reads the winner's pair.
    """

    __tablename__ = "WebPushVapidKey"

    id: Mapped[str] = mapped_column(primary_key=True, default="self", server_default=text("'self'"))
    public_key: Mapped[str] = mapped_column("publicKey")
    """The uncompressed P-256 public key, base64url. The browser subscribes with it."""
    private_key_encrypted: Mapped[str] = mapped_column("privateKeyEncrypted")
    """The private key, base64url, encrypted with the process encryption key."""
    created_at: Mapped[datetime] = mapped_column("createdAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"))

class SignInProviderMount(Base):
    __tablename__ = "SignInProviderMount"

    id: Mapped[str] = mapped_column(primary_key=True, default="self", server_default=text("'self'"))
    """Fixed, so this table holds one row: licensing's copy of sso's sign-in
    provider fact (ADR-179).
    """
    provider: Mapped[Optional[str]]
    """The configured provider's name; null when the deployment signs in by email."""
    mounted: Mapped[bool]
    last_applied_event_id: Mapped[str] = mapped_column("lastAppliedEventId")
    """The last sso event applied: the row's idempotency guard and its version."""
    updated_at: Mapped[datetime] = mapped_column("updatedAt", default=utc_now, server_default=text("CURRENT_TIMESTAMP"), onupdate=utc_now)
