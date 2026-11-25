import type { ColumnType } from 'kysely'

export type Generated<T> =
  T extends ColumnType<infer S, infer I, infer U>
    ? ColumnType<S, I | undefined, U>
    : ColumnType<T, T | undefined, T>

export type Timestamp = ColumnType<Date, Date | string, Date | string>

export type SsoConnectionMigrationPhase =
  | 'SETUP'
  | 'GRACE_LEGACY'
  | 'GRACE_DIRECT'
  | 'FINALIZING'
  | 'FINALIZED'

export type TeamUserRole = 'ADMIN' | 'MEMBER' | 'VIEWER' | 'CUSTOM'

export type RoleBindingScopeType = 'ORGANIZATION' | 'TEAM' | 'PROJECT'

export type GrantScopeType = 'ORGANIZATION' | 'TEAM' | 'PROJECT' | 'RESOURCE' | 'PLATFORM'

export type GrantPrincipalType =
  | 'USER'
  | 'API_KEY'
  | 'GROUP'
  | 'TEAM'
  | 'ORGANIZATION'
  | 'PROJECT'
  | 'ANYONE'

export type OrganizationUserRole = 'ADMIN' | 'MEMBER' | 'EXTERNAL' | 'DEVELOPER'

/** Where a license in the registry came from. */
export type IssuedLicenseSource = 'BACKOFFICE' | 'PURCHASE' | 'SCRIPT' | 'LEGACY_IMPORT'

/** Why a credit was granted. The amount is always prepaid, never promotional. */
export type ConnectedCreditGrantKind = 'commit' | 'added' | 'renewal'

/** What an invoice of a connected customer is for. */
export type ConnectedInvoiceKind = 'annual' | 'seat_change' | 'usage'

export type PlanTypes =
  | 'FREE'
  | 'PRO'
  | 'GROWTH'
  | 'GROWTH_SEAT_EUR_MONTHLY'
  | 'GROWTH_SEAT_EUR_ANNUAL'
  | 'GROWTH_SEAT_USD_MONTHLY'
  | 'GROWTH_SEAT_USD_ANNUAL'
  | 'ENTERPRISE'
  | 'LAUNCH'
  | 'ACCELERATE'
  | 'LAUNCH_ANNUAL'
  | 'ACCELERATE_ANNUAL'

export type PricingModel = 'TIERED' | 'SEAT_EVENT'

/** Organization signup intent (ADR-038): which product the org came for. */
export type OrganizationIntent = 'AGENT_GOVERNANCE' | 'LLM_OPS'

export type SubscriptionStatus = 'PENDING' | 'FAILED' | 'ACTIVE' | 'CANCELLED'

export type Currency = 'USD' | 'EUR'

export type INVITE_STATUS =
  | 'PENDING'
  | 'ACCEPTED'
  | 'WAITING_APPROVAL'
  | 'PAYMENT_PENDING'
  | 'REVOKED'

export type EvaluationExecutionMode = 'ON_MESSAGE' | 'AS_GUARDRAIL' | 'MANUALLY'

export type CostType = 'TRACE_CHECK' | 'GUARDRAIL' | 'CLUSTERING' | 'BATCH_EVALUATION'

export type CostReferenceType = 'CHECK' | 'TRACE' | 'PROJECT' | 'BATCH'

/** Who sees a dashboard, named as the prompt scope is, with one more level. */
export type DashboardScope = 'PRIVATE' | 'PROJECT' | 'ORGANIZATION'

export type TriggerAction =
  | 'SEND_EMAIL'
  | 'ADD_TO_DATASET'
  | 'ADD_TO_ANNOTATION_QUEUE'
  | 'SEND_SLACK_MESSAGE'
  | 'SEND_WEBHOOK'

export type TriggerKind = 'AUTOMATION' | 'ALERT' | 'REPORT'

export type AlertType = 'CRITICAL' | 'WARNING' | 'INFO'

/**
 * Who can use a Slack connection (ADR-093 §5a): every project in the
 * organization, or one project. Chosen per connection.
 */
export type SlackIntegrationScopeType = 'ORGANIZATION' | 'PROJECT'

/**
 * How a Slack connection posts: a bot token over the Web API (any channel it
 * is invited to, every block renders) or an incoming webhook (one channel).
 */
export type SlackIntegrationKind = 'BOT' | 'INCOMING_WEBHOOK'

export type WebhookDeliveryOutcome = 'success' | 'retryable' | 'terminal' | 'pending'

/**
 * The last hop a webhook endpoint delivers over. Lowercase because these
 * values are the wire spelling: the same strings appear in the REST DTO,
 * both SDKs and the CLI.
 */
export type WebhookDestinationKind = 'http' | 'sqs'

export type WebhookEndpointStatus = 'ACTIVE' | 'DISABLED'

/**
 * Which webhook pipeline wrote a delivery row, and therefore which columns
 * carry its tenancy: `platform` rows anchor on organizationId + endpointId,
 * `automations` rows on projectId + triggerId.
 */
export type WebhookDeliveryChannel = 'platform' | 'automations'

export type ExperimentType = 'DSPY' | 'BATCH_EVALUATION' | 'BATCH_EVALUATION_V2' | 'EVALUATIONS_V3'

export type ModelProviderScopeType = 'ORGANIZATION' | 'TEAM' | 'PROJECT'

export type ModelDefaultScopeType = 'ORGANIZATION' | 'TEAM' | 'PROJECT'

export type LangyProjectionTitleSource = 'derived' | 'auto' | 'user'

export type AnnotationScoreDataType = 'CATEGORICAL' | 'BOOLEAN' | 'LIKERT' | 'OPTION' | 'CHECKBOX'

export type ShareResourceType = 'TRACE' | 'THREAD'

export type ShareVisibility = 'PUBLIC' | 'ORGANIZATION' | 'PROJECT'

export type PinSource = 'manual' | 'share'

export type RetentionPolicyScopeType = 'ORGANIZATION' | 'TEAM' | 'PROJECT'

export type DataPrivacyScopeType = 'ORGANIZATION' | 'DEPARTMENT' | 'TEAM' | 'PROJECT'

export type CustomLLMModelCostScopeType = 'ORGANIZATION' | 'TEAM' | 'PROJECT'

export type PromptScope = 'ORGANIZATION' | 'PROJECT'

export type VirtualKeyStatus = 'ACTIVE' | 'DISABLED' | 'REVOKED'

export type VirtualKeyPurpose = 'USER' | 'LANGY' | 'CONNECT'

export type GatewayProviderRotationPolicy = 'MANUAL'

export type GatewayProviderHealthStatus = 'UNKNOWN' | 'HEALTHY' | 'DEGRADED' | 'CIRCUIT_OPEN'

export type GatewayBudgetScopeType =
  | 'ORGANIZATION'
  | 'TEAM'
  | 'PROJECT'
  | 'VIRTUAL_KEY'
  | 'PRINCIPAL'
  | 'GROUP'
  | 'ATTRIBUTED_USER'

export type GatewayBudgetWindow = 'MINUTE' | 'HOUR' | 'DAY' | 'WEEK' | 'MONTH' | 'TOTAL' | 'MANUAL'

export type GatewayBudgetBreachAction = 'BLOCK' | 'WARN'

export type GatewayBudgetLedgerStatus =
  | 'SUCCESS'
  | 'PROVIDER_ERROR'
  | 'BLOCKED_BY_GUARDRAIL'
  | 'CANCELLED'

export type GatewayChangeEventKind =
  | 'VK_CREATED'
  | 'VK_CONFIG_UPDATED'
  | 'VK_REVOKED'
  | 'VK_ROTATED'
  | 'VK_DISABLED'
  | 'VK_ENABLED'
  | 'BUDGET_CREATED'
  | 'BUDGET_UPDATED'
  | 'BUDGET_DELETED'
  | 'MODEL_PROVIDER_UPDATED'
  | 'CACHE_RULE_CREATED'
  | 'CACHE_RULE_UPDATED'
  | 'CACHE_RULE_DELETED'
  | 'ROUTING_POLICY_UPDATED'
  | 'ROUTING_POLICY_DELETED'
  | 'A2A_TARGET_CREATED'
  | 'A2A_TARGET_UPDATED'
  | 'A2A_TARGET_DELETED'

export type GatewayCacheRuleMode = 'RESPECT' | 'FORCE' | 'DISABLE'

export type VirtualKeyRoutingMode = 'NONE' | 'FALLBACK_ALL' | 'POLICY'

export type VirtualKeyScopeType = 'ORGANIZATION' | 'TEAM' | 'PROJECT'

export type RoutingPolicyScopeType = 'ORGANIZATION' | 'TEAM' | 'PROJECT'

export type GatewayGuardrailDirection = 'PRE' | 'POST' | 'STREAM_CHUNK'

export type GatewayGuardrailFailureMode = 'FAIL_OPEN' | 'FAIL_CLOSED'

export type FeatureFlagExperimentSubjectType = 'USER' | 'ORGANIZATION' | 'PROJECT'

export type ProcessManagerOutboxStatus = 'pending' | 'dispatched' | 'dead' | 'discarded'

export type ProcessManagerOutboxAttemptOutcome = 'retry_scheduled' | 'dead'

export type GatewayRealtimeSessionStatus = 'OPEN' | 'CLOSED' | 'FAILED' | 'EXPIRED'

/** How the invoicing of one seat change ended. */
export type ConnectedSeatChangeState =
  | 'intent'
  | 'invoiced'
  | 'nothing_to_invoice'
  | 'not_onboarded'

export interface Account {
  id: string
  userId: string
  type: Generated<string>
  provider: string
  /**
   * WHO asserted `providerAccountId`. Better Auth keys accounts by
   * `(provider, providerAccountId)`; auth's account create hook writes this
   * (modules/auth/specs/account-issuer.feature). Nullable only so the column
   * could be added to a populated table; the migration backfilled every row.
   */
  issuer: string | null
  providerAccountId: string
  refresh_token: string | null
  access_token: string | null
  expires_at: Timestamp | null
  ext_expires_in: number | null
  token_type: string | null
  scope: string | null
  id_token: string | null
  session_state: string | null
  password: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface AccountCredential {
  id: string
  userId: string
  provider: string
  password: string | null
  accessToken: string | null
  refreshToken: string | null
  idToken: string | null
  accessTokenExpiresAt: Timestamp | null
  refreshTokenExpiresAt: Timestamp | null
  scope: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface Session {
  id: string
  sessionToken: string
  userId: string
  expires: Timestamp
  ipAddress: string | null
  userAgent: string | null
  /**
   * Which of the user's sign-in methods minted this session (D06). NULL on
   * every session that predates the column and on any mint we could not
   * attribute - an ordinary value, not an error. Nothing reads it to decide
   * whether the session is valid, so no session was ended to add it.
   */
  identifierId: string | null
  /**
   * What the sign-in that minted this session PROVED, as RFC 8176 method
   * references (D06). Empty on every session minted before the column, and
   * on any ordinary password sign-in. An organization that requires a second
   * factor reads this when a member reaches its data; it is never read at
   * mint and never ends a session.
   */
  amr: Generated<string[]>
  /**
   * Minted while the account had never signed in: the session its sign-up
   * started, which skips the secure-account offer. False on every later one
   * and on every session that predates the column (specs/identity/passkeys.feature).
   */
  startedBySignUp: Generated<boolean>
  /**
   * The legacy impersonation payload. NOTHING READS OR WRITES IT any more -
   * the {actor, subject} columns below replaced it, and the migration beside
   * them ended every session that still carried a value.
   *
   * It stays for ONE release, and that is the whole reason it is still here.
   * `prisma migrate deploy` runs at container start, so the first pod of a
   * new release drops a column while the previous release's pods are still
   * serving - and those pods select this column on EVERY authenticated
   * request and fail closed when the row does not come back. Dropping it in
   * the same release that stops reading it signs everybody out for the length
   * of the rollout. Expand now, contract next release.
   */
  impersonating: unknown
  /**
   * The impersonation claims (D06), aligned with the authz Principal
   * {actor, subject}. The actor is the operator who really made the request
   * and the subject is the person whose access they are borrowing; both are
   * recorded on every authorization decision, so the audit trail can answer
   * who really did it. These replace the legacy `impersonating` JSON blob,
   * which is dropped in the release AFTER this one.
   *
   * NULL on every ordinary session, which is all of them: an impersonation
   * is a handful of rows at a time and nothing else reads these columns.
   */
  actorUserId: string | null
  subjectUserId: string | null
  /**
   * Why the operator started it. Required to start one
   * (specs/features/backoffice-user-impersonation-reason.feature) and kept
   * beside both people rather than only in the audit log, so a request under
   * impersonation carries its own justification.
   */
  impersonationReason: string | null
  /**
   * When the borrowed access lapses. A session outlives its impersonation
   * window: past this the claims are ignored and the operator is themselves
   * again, with nothing ended.
   */
  impersonationExpiresAt: Timestamp | null
  /**
   * When this session was last used, for the idle timeout an organization
   * may set (GAC-10, specs/identity/org-session-lifetime.feature).
   *
   * A COLUMN OF OUR OWN rather than `updatedAt`, because better-auth rolls a
   * live session's expiry once per `updateAge` - a day - so `updatedAt` is
   * activity to the nearest day and cannot tell an idle hour from an idle
   * minute. Written only for sessions whose owner belongs to an organization
   * that has set a window, and then at most once per quarter-window, so a
   * deployment where nobody has turned this on writes it never.
   *
   * NULL on every session minted before the column and on every session
   * under no window at all. A null READS AS `updatedAt`, which better-auth
   * keeps rolled to within a day of the last use. That fallback is
   * conservative in the direction that matters - `updatedAt` is never
   * earlier than the real last use, so it can only ever judge a session MORE
   * active than it was, never less - and it is what lets a window turned on
   * today end the sessions that have been idle since February instead of
   * waiting a day for a stamp they have never had.
   */
  lastSeenAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/**
 * One row per erased account, holding only its creation time: no id, name or address. It keeps
 * the install's first-created account knowable after an erasure deletes that row.
 * Spec: modules/user/specs/user.feature
 */
export interface UserErasure {
  id: string
  accountCreatedAt: Timestamp
  erasedAt: Generated<Timestamp>
}

export interface User {
  id: string
  name: string | null
  email: string | null
  emailVerified: Generated<boolean>
  signupConfirmationPending: Generated<boolean>
  passkeySignupClaimHash: string | null
  image: string | null
  pendingSsoSetup: Generated<boolean>
  userHashKey: string | null
  /**
   * Owned by the two-factor plugin (D06), which is why it is a plain
   * boolean rather than a state: the plugin flips it when a setup is
   * confirmed and when it is turned off, and challenges every sign-in while
   * it is true. The lifecycle - and the history - lives on MfaEnrollment.
   */
  twoFactorEnabled: Generated<boolean>
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  lastLoginAt: Timestamp | null
  deactivatedAt: Timestamp | null
  lastHomePath: string | null
  tracesExplorerTourDismissedAt: Timestamp | null
  /**
   * How Langy should reach this person's code when a task needs a change
   * (ADR-129). "github" means they chose the pull request path and asked to
   * be remembered; NULL means ask. A shared local folder is never stored
   * here, because a folder has to be shared again each time.
   * Spec: specs/langy/langy-code-access.feature
   */
  langyCodeAccessPreference: string | null
  /**
   * When this person last said "not now" to being offered a passkey
   * (ADR-120). On the account rather than in browser storage for two
   * reasons: a per-browser record forgets on a new device, which is exactly
   * where the offer should be MORE eager, and an interval nobody can enforce
   * is not an interval. NULL means never asked.
   */
  passkeyNudgeDismissedAt: Timestamp | null
  /**
   * The domains this person has said "no thanks" to being offered their
   * colleagues on (D12). Per domain rather than one flag: somebody who
   * changes jobs and verifies a second address should be offered their new
   * employer's organization, not silenced forever by a decision they made
   * about the old one. Empty means never dismissed anything.
   */
  joinOfferDismissedDomains: Generated<string[]>
  /**
   * This person's answer about browser notifications, one entry per topic:
   * { "langy": "enabled" | "declined" }. A topic that is absent was never
   * answered. The browser's own permission is not stored here.
   * Spec: specs/langy/langy-notifications.feature
   */
  notificationPreferences: Generated<unknown>
  /**
   * When an operator marked this person's password as no longer trusted. A
   * password sign-in opens no session while it is set; any write of a new
   * password hash clears it. Spec: modules/auth/specs/forced-password-reset.feature
   */
  passwordResetRequiredAt: Timestamp | null
}

export interface VerificationToken {
  id: string
  identifier: string
  /** Not unique alone: better-auth stores plain values here (a user id, a counter). See WEB-19201. */
  token: string
  expires: Timestamp
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface Identifier {
  id: string
  userId: string
  provider: string
  value: string | null
  domain: string | null
  identifierHash: string | null
  accountId: string | null
  providerId: string | null
  issuer: string | null
  providerAccountId: string | null
  state: string
  connectionId: string | null
  verifiedAt: Timestamp | null
  attachedAt: Timestamp
  detachedAt: Timestamp | null
  lastUsedAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/**
 * Whether a user signs in on the identifier path: identity's own switch, set when the
 * identifier backfill proves the user and pinned off by an operator's rollback. No row is the
 * legacy path. Spec: modules/identity/specs/sign-in-switch.feature. Column refs only, no FK.
 */
export interface IdentitySignInSwitch {
  userId: string
  onIdentifiers: boolean
  pinned: Generated<boolean>
  switchedAt: Timestamp
  lastAppliedEventId: string | null
  updatedAt: Generated<Timestamp>
}

export interface IdentifierReservation {
  normalizedValue: string
  userId: string
  identifierId: string
  commandId: string
  createdAt: Generated<Timestamp>
}

export interface IdentityProjectionCursor {
  userId: string
  lastEventId: string
  acceptedAt: Timestamp
  occurredAt: Timestamp
  projectionVersion: string
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/**
 * The two-factor plugin's own table (D06). Protocol state, and row-truth for
 * good (ADR-101 R12): the shared secret and the backup codes live here and
 * in no event, ever. Nothing in this repo reads or writes these columns -
 * the plugin owns them end to end, including the lockout counters, which is
 * why `failedVerificationCount` and `lockedUntil` are here rather than
 * rebuilt on the projection.
 *
 * `backupCodes` holds an ENCRYPTED blob, not hashes. A backup code has to be
 * compared against what somebody types, and the plugin decrypts and compares;
 * hashing them would leave it unable to verify its own codes.
 */
export interface TwoFactor {
  id: string
  userId: string
  secret: string
  backupCodes: string
  verified: Generated<boolean>
  failedVerificationCount: Generated<number>
  lockedUntil: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/**
 * Consecutive failed sign-ins, and the lock they earn (GAC-09).
 * Spec: specs/identity/org-account-lockout.feature.
 *
 * KEYED ON THE ADDRESS THAT WAS TYPED, NOT ON A USER, and that is the whole
 * design. We already answer "no such account" and "wrong password" with one
 * refusal so that sign-in cannot be used to discover who has an account. A
 * lock that existed only for real accounts would put that back: five wrong
 * guesses saying "locked" for one address and "incorrect" for another tells
 * an attacker which addresses are worth attacking. So an address with
 * nothing behind it gets a row, a count and a lock exactly like one that
 * resolves, and both answer identically.
 *
 * `identifierHash` is a KEYED hash of the normalised address, never the
 * address. The key is the deployment's own secret, so this table cannot be
 * lifted and reversed into the list of every address anybody has ever tried
 * to sign in as - which is what it would otherwise become, including every
 * typo and every address that has no account here at all.
 *
 * Row-truth, deliberately (ADR-101 R12): this is protocol state read and
 * written on the sign-in path, like TwoFactor's own counters next door. The
 * EVIDENCE an auditor reads is the identity event emitted when a lock is
 * taken, which is durable; these rows are reaped once they stop meaning
 * anything.
 */
export interface SignInAttemptLock {
  id: string
  /** Keyed hash of the normalised address. Never the address itself. */
  identifierHash: string
  /**
   * Consecutive failures since the last success. Reset by a success, not by
   * time - "five consecutive", not five ever.
   */
  failedCount: Generated<number>
  /** When the current lock lifts. Null when not locked. */
  lockedUntil: Timestamp | null
  /**
   * Locks taken since the last success. At five this stops lifting by
   * itself and becomes an incident - see `heldForReview`.
   */
  consecutiveLockouts: Generated<number>
  /**
   * Held past its lock because the count of locks reached the ceiling the
   * control names. Cleared by an administrator, or by the owner proving
   * mailbox control through a password reset - the second is what keeps a
   * held account from being a denial of service anyone can inflict by
   * typing a colleague's address wrong often enough.
   */
  heldForReview: Generated<boolean>
  /**
   * Whose account this resolved to at the time, when it resolved to one.
   * Null for an address with no account - those rows are the reason this is
   * nullable, and the reason this table is not hung off User.
   */
  userId: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/**
 * The passkey plugin's own table (D07). Protocol state, like TwoFactor: the
 * public key and the signature counter are the plugin's to read and write.
 * A passkey also becomes an `Identifier` row with provider `passkey`, which
 * is what makes it a sign-in method like any other - this table is the
 * credential, that one is the identity.
 *
 * `aaguid` names the authenticator MODEL, never a device or a person, and is
 * only present at registration. Privacy-preserving platforms report all
 * zeroes, so a label derived from it is best-effort by construction.
 */
export interface Passkey {
  id: string
  name: string | null
  publicKey: string
  userId: string
  credentialID: string
  counter: Generated<number>
  deviceType: string
  backedUp: Generated<boolean>
  transports: string | null
  aaguid: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/**
 * The two-step verification pipeline's projection (D06 - see
 * specs/identity/mfa-and-session-shape.feature). One row per PERSON, because
 * one person has one setup however many organizations they belong to.
 *
 * A pure event-truth head like `SsoConnection`: every column is fold-written,
 * rows are never deleted (DISABLED is a tombstone, so history survives being
 * turned off), and a replay rebuilds it whole-row from the event log. The
 * cursor columns are inline because the fold produces exactly one row per
 * aggregate - the `Identifier` fan-out needed its own cursor table, this
 * does not.
 *
 * Every column here is a lifecycle fact. There is deliberately nowhere in
 * this model a secret or a backup code could be put: `consumedBackupCodeIndexes`
 * holds POSITIONS, which is what lets "how many are left" be answered
 * without the projection ever knowing a code.
 */
export interface MfaEnrollment {
  /**
   * The user. Also the aggregate id and the tenant - one identity history
   * per person (ADR-029 §4).
   */
  userId: string
  /** The current enrollment, or NULL for somebody who never started one. */
  enrollmentId: string | null
  /** totp. Never sms. */
  method: string | null
  /** NONE | PENDING | ENABLED | EXPIRED | DISABLED */
  state: string
  enrolledAt: Timestamp | null
  confirmedAt: Timestamp | null
  expiredAt: Timestamp | null
  disabledAt: Timestamp | null
  /** password+totp | org-admin - who ended it, and on whose authority. */
  disabledVia: string | null
  backupCodeCount: Generated<number>
  /** Positions already spent, ascending. Never codes. */
  consumedBackupCodeIndexes: Generated<number[]>
  failedCount: Generated<number>
  occurredAt: Timestamp
  lastEventId: string
  acceptedAt: Timestamp
  projectionVersion: string
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/**
 * The SSO connection pipeline's projection (ADR-117 §5, D04 - see
 * specs/identity/sso-connection-lifecycle.feature). A pure event-truth head
 * like `Identifier`: every column is fold-written, rows are never deleted
 * (TORN_DOWN is a tombstone), and a replay rebuilds it whole-row from the
 * event log.
 *
 * The fold cursor lives on the row rather than in a sibling table, because
 * one row IS one aggregate here - a connection - so there is nothing for a
 * separate cursor table to key on that this row does not already key on.
 *
 * No relation to Organization on purpose, matching `Identifier`: a
 * projection is rebuilt from events, and a foreign key would make a replay
 * depend on relational state the log does not carry. `idpMetadata` holds
 * endpoints and credential REFERENCES; no secret is ever stored here or in
 * any event.
 */
export interface SsoConnection {
  id: string
  organizationId: string
  /** oidc | saml */
  type: string
  /**
   * DRAFT | CLAIMED | APPROVED | REJECTED | DISCARDED |
   * VERIFICATION_PENDING | VERIFIED | ACTIVE | SUSPENDED |
   * TEARDOWN_PENDING | TORN_DOWN
   */
  state: string
  claimedDomains: string[]
  /**
   * [{ domain, state, claimedAtMs, claimedByActorId, decidedAtMs,
   * decidedByActorId, authority, waitedMs, note }] - every claim this
   * connection has made and how long each waited to be decided. The tier-3
   * approval queue is a read of this column, and the epic's Open Q2 wants
   * that latency measured from the day the queue exists.
   */
  domainClaims: Generated<unknown>
  approvedDomains: string[]
  /** The only domains that ever route, and only while state is ACTIVE. */
  verifiedDomains: string[]
  /**
   * The subset of `verifiedDomains` whose published record stayed missing
   * through its grace window (ADR-123). They still ROUTE - everyone already
   * here signs in unchanged - and they stop vouching for anybody NEW: no
   * provisioning on first sign-in, no walking in by domain. A column rather
   * than a read of the JSON below, because both questions are asked on a
   * sign-in path and have to be one indexed predicate.
   */
  lapsedDomains: string[]
  /**
   * [{ domain, method, actorId, verifiedAtMs, proofState, firstAbsentAtMs,
   * graceEndsAtMs }] - what proved each verified domain, who, and whether
   * that evidence is still there. Kept on the row rather than only in the
   * log so that an operator-attested domain can never be read as one the
   * customer proved, and a wavering one never as one nothing is doubting.
   */
  domainVerifications: Generated<unknown>
  /** { domain, method, tokenHash } while a ceremony is in flight. */
  pendingVerification: unknown
  idpMetadata: unknown
  /**
   * admit | request | refuse - what this connection does with somebody who
   * signs in through it and is not a member yet (ADR-117 3). Stated at
   * registration and changed by the setup journey, so it is never absent
   * and no reader has to decide what absence means. This was two fields
   * for a while - a boolean `allowsJit` beside it, kept in step by the
   * fold - and they disagreed about the middle answer.
   */
  arrivalPolicy: Generated<string>
  /**
   * DEAD, AND STAYING FOR ONE RELEASE. Nothing reads or writes this; the
   * routing repository derives the same answer from `arrivalPolicy`. It is
   * declared so the schema still describes the table, and left in place so
   * the PREVIOUS release's pods — which select it by name on every
   * connection read — keep working through the rollout. Dropping it in the
   * release that stopped reading it would stop single sign-on routing and
   * the connection projection for the length of the deploy. Dropped by
   * `20260827120000_drop_sso_connection_allows_jit` in the next one.
   */
  allowsJit: Generated<boolean>
  /**
   * When somebody CHOSE the policy above, and null while the registration
   * default stands. A different fact from the policy rather than a second
   * copy of it: a connection always has a behaviour, and separately
   * somebody has or has not decided on it. Going live waits for the
   * deciding - "turn everybody away" is a decision too.
   */
  arrivalPolicyDecidedAt: Timestamp | null
  /** self-serve | legacy-grandfathered */
  source: string
  testLoginAccountId: string | null
  /**
   * The grandfathered connection this direct connection replaces. Null for
   * ordinary connections and for the grandfathered predecessor itself.
   *
   * Partial uniqueness for nonterminal replacements is installed by the
   * flattened migration; Prisma cannot express that predicate here.
   */
  replacesConnectionId: string | null
  /**
   * SETUP | GRACE_LEGACY | GRACE_DIRECT | FINALIZING | FINALIZED.
   * Null means this connection is not a legacy-to-direct replacement.
   */
  migrationPhase: SsoConnectionMigrationPhase | null
  graceStartedAt: Timestamp | null
  /**
   * When an administrator last selected the normal route. This is audit
   * evidence only: routing is derived from `migrationPhase`, never recency.
   */
  routeChangedAt: Timestamp | null
  finalizationRequestedAt: Timestamp | null
  finalizedAt: Timestamp | null
  /** { domain, note } - the last ops rejection, read back on a re-claim. */
  rejection: unknown
  createdBy: string | null
  /** While TEARDOWN_PENDING: when the process manager's wake completes it. */
  tearDownAfter: Timestamp | null
  occurredAt: Timestamp
  lastEventId: string
  acceptedAt: Timestamp
  projectionVersion: string
  /**
   * Event-derived, not database-managed: `@default(now())` / `@updatedAt`
   * would make a replayed row differ from the row it rebuilds, which is
   * exactly the parity this projection has to keep.
   */
  createdAt: Timestamp
  updatedAt: Timestamp
}

/**
 * The one-statement pre-event registration lock. `SsoConnection` remains
 * event truth; this row only decides which command may append that truth.
 * A slot whose projected connection is terminal may be atomically replaced.
 */
export interface SsoConnectionRegistrationSlot {
  organizationId: string
  /** legacy | direct */
  kind: string
  connectionId: string
  replacesConnectionId: string | null
  commandId: string
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/**
 * Operational round-robin position for ADR-123's domain re-proof sweep.
 *
 * This is deliberately not a column on `SsoConnection`: that row is a
 * whole-row event projection and its fold is its only writer, while a
 * healthy re-read states no fact. Keeping the scheduling cursor beside the
 * projection preserves both rules and still lets a fixed-size sweep revisit
 * every connection, including one whose proof has started wavering.
 */
export interface SsoConnectionReproofCursor {
  connectionId: string
  lastReproofAt: Timestamp
}

/**
 * The values an `SsoConnection`'s credential references point AT (D09 - see
 * specs/identity/sso-idp-termination.feature).
 *
 * The aggregate has always carried `clientIdRef` / `secretRef` / `certRefs`
 * rather than values, because a fact may not carry a secret (ADR-101 §4).
 * Until D09 there was nothing behind the references and every writer passed
 * null. This is the thing behind them.
 *
 * `ciphertext` is AES-256-GCM under the deployment's `CREDENTIALS_SECRET`,
 * the same key every other credential in the product is kept under. Nothing
 * reads a row except the projection that builds the engine's dialing
 * configuration; no surface renders one back.
 *
 * Written by the command path rather than the fold, and that split is
 * deliberate: the VALUE is not event truth and cannot be, so the log records
 * WHEN a credential changed (a new reference in a new fact) and never what
 * it was. A replay rebuilds every reference from the log and finds the value
 * still here.
 */
export interface SsoCredential {
  id: string
  organizationId: string
  /**
   * The connection whose registration minted this. Not a relation, matching
   * `SsoConnection` itself: a projection rebuilt from events must not depend
   * on relational state the log does not carry.
   */
  connectionId: string
  /** oidc-client-id | oidc-client-secret | saml-idp-config */
  kind: string
  ciphertext: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

/**
 * better-auth's single sign-on plugin's provider table (D09), and a
 * PROJECTION of the connection log rather than a second source of truth.
 *
 * The plugin owns the columns and the reads: it is the only thing that
 * speaks OpenID Connect discovery or reads a SAML assertion, and it dials
 * what it finds here. What it never does is decide that a row should exist —
 * that is folded from `SsoConnection`'s events by the same projection that
 * writes the connection head, in the same transaction, so a replay rebuilds
 * both together and a row can never outlive the connection that justified
 * it. Nothing else writes here.
 *
 * `providerId` is the CONNECTION ID rather than the name the customer typed.
 * The plugin's column is globally unique, and two organizations both calling
 * their provider `okta` must both be able to.
 *
 * `oidcConfig` and `samlConfig` are the plugin's own JSON documents, rebuilt
 * on every fold from `SsoCredential` — the vault is where the material
 * lives, and this row is a derived copy the engine can read synchronously.
 *
 * SEALED AT REST, under the same key the vault uses. The document carries
 * the client secret the derivation reads OUT of the vault, so a plaintext
 * copy here would hand every customer's live identity-provider credential
 * to anybody holding a database dump — the exact threat `SsoCredential`
 * exists to defeat, reopened one table over. A stored value announces its
 * form (`enc:v1:`), so a row written before the seal existed still opens;
 * the one reader is the identity storage adapter, on the path that dials.
 * The reader tolerates both forms, so this needs no backfill — but a pod
 * without it cannot read a sealed row, so it must not be rolled back past.
 * WHO OWNS A PROVED DOMAIN, as a row the database can refuse twice.
 *
 * "First verifier owns" was enforced by reading `findDomainOwner` and then
 * writing the fact — two different aggregates, so the per-connection queue
 * never ordered them against each other, and `SsoConnection.verifiedDomains`
 * is a `String[]` with only a GIN index behind it, which enforces nothing.
 * Two organizations proving `acme.com` at the same moment both read "nobody
 * owns it" and both committed; the sign-in router then picked between two
 * live connections by planner order, so a member of one company could be
 * handed to another company's identity provider — and with `arrivalPolicy`
 * on `admit`, provisioned into it.
 *
 * The legacy single-domain model had this: `Organization.ssoDomain` is
 * `@unique`. The multi-domain rewrite dropped the guarantee, and self-serve
 * proof is what made the race reachable by customers rather than by two
 * operators working at once.
 *
 * `domain` is the primary key, so a second organization loses on the
 * constraint rather than on a check it raced. Connection holders are
 * separate because an exact predecessor/replacement pair must coexist
 * during migration. Both are projected with the connection head.
 */
export interface SsoVerifiedDomain {
  domain: string
  organizationId: string
}

/**
 * A connection currently entitled to use an organization-owned verified
 * domain. Usually there is one holder. During an Auth0 migration there are
 * exactly two: the grandfathered predecessor and its direct replacement.
 *
 * The projection repository enforces that relationship before inserting a
 * second holder. The flattened migration also installs partial unique live
 * legacy/direct slots on `SsoConnection`, closing concurrent third-writer
 * races that Prisma cannot describe declaratively.
 */
export interface SsoVerifiedDomainHolder {
  domain: string
  connectionId: string
  organizationId: string
}

/**
 * A successful callback through one SSO connection. Identifier.lastUsedAt
 * also advances during routing and resolution, so it cannot prove that a
 * member actually authenticated through the replacement before cutover.
 */
export interface SsoAuthenticationActivity {
  id: string
  organizationId: string
  connectionId: string
  userId: string
  authenticatedAt: Timestamp
  /** The provider-side account the assertion named; null before #7633's arrival door recorded it. */
  providerAccountId: string | null
}

export interface SsoProvider {
  id: string
  issuer: string
  oidcConfig: string | null
  samlConfig: string | null
  userId: string | null
  providerId: string
  organizationId: string | null
  domain: string
}

/**
 * A way back in (D05 - see specs/identity/sso-onboarding-tiers.feature).
 *
 * One named person who may still sign in without the identity provider
 * while their organization's sign-in belongs to one, and the date on which
 * that stops. Activation asks whether such a person exists, because the
 * failure it prevents is the only one that cannot be recovered from inside
 * the product: a misconfigured connection goes live, every member is handed
 * to an identity provider that will not authenticate them, and nobody is
 * left who can turn it off.
 *
 * Rows are immutable except for `supersededAt` and `warnedDays`. A renewal
 * writes a NEW row naming the one it replaced, so the date a way in
 * previously ended is still readable in the history - which is what makes a
 * renewal auditable rather than a column somebody moved.
 *
 * No relation to Organization or User, matching the identity heads above:
 * the binding is the record of who could get in, and a cascade would erase
 * exactly the history a lockout post-mortem reads.
 */
export interface SsoBreakGlassBinding {
  id: string
  organizationId: string
  /** Who holds the way in. */
  userId: string
  /** Who granted it. A way back in is never self-served. */
  grantedByUserId: string
  grantedAt: Timestamp
  expiresAt: Timestamp
  /** Set when a renewal replaced this row; the row itself is never edited. */
  supersededAt: Timestamp | null
  /** The binding this one renewed, so the previous end date stays readable. */
  renewedFromId: string | null
  /**
   * Which of the fourteen, seven and one day warnings have been sent, so a
   * sweep that runs twice in a day never says the same thing twice.
   */
  warnedDays: number[]
  createdAt: Generated<Timestamp>
}

/**
 * A durable claim on one live recovery path while activation is between its
 * guard and projected ACTIVE fact. It has no expiry: only projection of the
 * matching activation, or an explicit cancellation, may consume it.
 */
export interface SsoActivationRecoveryReservation {
  commandId: string
  organizationId: string
  connectionId: string
  createdAt: Timestamp
}

/**
 * The join-request pipeline's projection (ADR-117, D12 - see
 * specs/identity/join-requests.feature). A pure event-truth head like
 * `Identifier` and `SsoConnection`: every column is fold-written, rows are
 * never deleted (a terminal state is a tombstone the admin panel stops
 * listing, not an absence), and a replay rebuilds it whole-row from the
 * event log.
 *
 * The fold cursor lives on the row rather than in a sibling table, because
 * one row IS one aggregate here - a request.
 *
 * No relation to Organization or User on purpose, matching the two heads
 * above: a projection is rebuilt from events, and a foreign key would make a
 * replay depend on relational state the log does not carry.
 *
 * Membership is NOT here. An approval dispatches an attach on the grants
 * ledger, exactly as accepting an invitation does, and that ledger fact
 * carries `source: "join-request"` as its provenance - which is what puts a
 * join approval on the customer's audit page beside every other membership
 * change rather than in a history of its own.
 *
 * The requester's address is not here either: the DOMAIN is the fact, and it
 * is the only part of an address a join decision may depend on.
 */
export interface JoinRequest {
  id: string
  userId: string
  organizationId: string
  /** The normalized domain the match was made on. */
  domain: string
  /** PENDING | APPROVED | REJECTED | EXPIRED | WITHDRAWN */
  state: string
  /** How the domain was matched: `verified-identifier-domain` or `sso-connection-domain`. */
  matchedVia: string
  /**
   * Where the request was made: `web` or `cli` (ADR-171 v6). A request from
   * the terminal lands a Developer seat whatever the joiner seat says.
   */
  origin: Generated<string>
  /** While PENDING: when the day-14 wake lapses it. Null once resolved. */
  expiresAt: Timestamp | null
  resolvedAt: Timestamp | null
  /**
   * user | policy | invite - who ended it. `policy` is domain auto-join's
   * own principal, deliberately not the one SSO auto-join uses.
   */
  resolvedByType: string | null
  /** The admin's user id, `domain-auto`, or the invitation id. */
  resolvedById: string | null
  /** user | invite-accepted, on a withdrawal. */
  withdrawalCause: string | null
  /** The SSO connection a single sign-on arrival was admitted through; null otherwise. */
  connectionId: string | null
  occurredAt: Timestamp
  lastEventId: string
  acceptedAt: Timestamp
  projectionVersion: string
  /**
   * Event-derived, not database-managed: `@default(now())` / `@updatedAt`
   * would make a replayed row differ from the row it rebuilds, which is
   * exactly the parity this projection has to keep.
   */
  createdAt: Timestamp
  updatedAt: Timestamp
}

export interface TeamUser {
  userId: string
  teamId: string
  role: TeamUserRole
  assignedRoleId: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface OrganizationUser {
  userId: string
  organizationId: string
  role: OrganizationUserRole
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  /**
   * Department assignment for spend attribution (the people lens).
   * Single-valued, org-scoped, accounting-only (never gates access).
   * Plain column ref (relationMode="prisma"); the service validates the
   * id belongs to this org. See departments.feature.
   */
  departmentId: string | null
  /**
   * A Billing-only member (ADR-205): role EXTERNAL, the Billing grant, no team and
   * no seat. Any role change clears it. See billing-only-invite.feature.
   */
  billingOnly: Generated<boolean>
  /**
   * Set when an admin disables this membership to stay within the licensed
   * seat count. The person loses access to this organization only, keeps
   * their role, department and everything they did, and frees their seat.
   * Reversible: re-enabling restores access if a seat is available.
   * Seat counting excludes disabled memberships. See
   * seat-reconciliation.feature.
   */
  disabledAt: Timestamp | null
  /**
   * Stable grant intent for an unfinished automatic SSO admission.
   * Cleared after the authoritative grant is observed; never inferred from missing access.
   */
  pendingSsoGrantId: string | null
  /**
   * A new generation for each membership lifetime. Stamped USER grant facts
   * must match the current active membership before they can be projected;
   * removing and rejoining therefore cannot revive a delayed old attach.
   */
  membershipStamp: Generated<string>
  /**
   * When this membership started holding the full seat it holds now; null means createdAt.
   * Anchors the 30-day seat hold (specs/licensing/seat-rotation.feature).
   */
  fullSeatSince: Timestamp | null
  /**
   * Set when admission wrote this membership disabled because no seat was free: the person
   * is waiting for one. Cleared when the membership is enabled.
   */
  seatPendingSince: Timestamp | null
}

/**
 * A full seat still counted after the person who took it stopped holding it (removed,
 * disabled or demoted), until 30 days after they took it (specs/licensing/seat-rotation.feature).
 * One per organization and person; a return within the hold takes the seat back.
 */
export interface SeatHold {
  id: string
  organizationId: string
  userId: string
  /** Display copy for the Users list, kept because a removed person has no membership row. */
  email: string | null
  name: string | null
  /** When the person took the seat this hold keeps. */
  seatTakenAt: Timestamp
  /** seatTakenAt plus 30 days; expired holds are ignored by every count and deleted hourly. */
  heldUntil: Timestamp
  /** REMOVED | DISABLED | DEMOTED */
  reason: string
  createdAt: Generated<Timestamp>
}

/**
 * A person asked for a full seat the organization had no room for (modules/organization/specs/
 * seat-requests.feature). One row per organization and address; reopening reuses the row.
 */
export interface SeatRequest {
  id: string
  organizationId: string
  /** Lowercased address of the person asking. */
  email: string
  /** Null when only the address is known (a sign-in the identity provider refused). */
  userId: string | null
  name: string | null
  /** OPEN | APPROVED | DECLINED */
  status: string
  /** pending | lite | idp_refused: how the person was left when the request opened. */
  source: string
  /** The identity provider also refused this person's sign-in. */
  idpRefused: Generated<boolean>
  /** okta | entra | google | onelogin | jumpcloud | ping | auth0 | oidc | saml */
  idpKind: string | null
  /** When sales were told about this request; reset when it reopens. */
  salesNotifiedAt: Timestamp | null
  openedAt: Timestamp
  resolvedAt: Timestamp | null
  resolvedByUserId: string | null
  /**
   * Set when an approval claims the request, cleared once its seat is given; an approval
   * still pending long after is one whose seat never landed, and is reopened.
   */
  seatPendingSince: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/**
 * Held by the one seat-taking write running on an organization (seat-requests.feature). A
 * row past expiresAt is free to take; a holder deletes its own row when its write is done.
 */
export interface OrganizationSeatLease {
  organizationId: string
  /** Random per hold, so a release deletes only its own lease. */
  holder: string
  expiresAt: Timestamp
  createdAt: Generated<Timestamp>
}

/**
 * The mail log for seat requests and refused sign-ins: at most one row per organization,
 * kind, address and UTC day, each mailed at once.
 */
export interface OrganizationAccessNotice {
  id: string
  organizationId: string
  /** seat_request | idp_refused */
  kind: string
  email: string
  name: string | null
  idp: string | null
  /** UTC day the notice was recorded on, YYYY-MM-DD. */
  dayKey: string
  /** Whether the seat request it announces was also refused by the identity provider. */
  idpRefused: Generated<boolean>
  createdAt: Generated<Timestamp>
  /** Set once the notice was handed to the mailer; null if a send never started. */
  mailedAt: Timestamp | null
}

export interface Team {
  id: string
  name: string
  slug: string
  organizationId: string
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  archivedAt: Timestamp | null
  isPersonal: Generated<boolean>
  ownerUserId: string | null
  /**
   * Department assignment inherited as the default for members and
   * projects that have none of their own. Accounting-only.
   */
  departmentId: string | null
}

export interface Organization {
  id: string
  name: string
  /**
   * `customer` for every organization a person founds or an operator provisions; `system` for
   * the install's one hidden organization the Ops checkup canaries run in, never listed to a person.
   * Spec: modules/ops/specs/checkup-canary-project.feature.
   */
  kind: Generated<string>
  phoneNumber: string | null
  slug: string
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  usageSpendingMaxLimit: number | null
  /**
   * The largest file a dataset image or file cell accepts for this
   * organization, in mebibytes. Null means the default. Set by an operator
   * from the back office only; every other dataset size limit scales with it.
   */
  datasetAttachmentMaxMb: number | null
  /**
   * The label of this organization's `CLICKHOUSE_URL__<label>__<id>` route, recorded by the
   * organization module the first time it is seen. Once set, a missing route refuses the
   * organization instead of serving it from the shared ClickHouse. Null: never private.
   * Spec: specs/private-dataplane/clickhouse-routing.feature.
   */
  privateClickHouseLabel: string | null
  /** The same for `DATAPLANE_S3__<label>__<id>`. Spec: specs/private-dataplane/s3-routing.feature. */
  privateS3Label: string | null
  /** When either label was first recorded. */
  privateDataplaneSeenAt: Timestamp | null
  /**
   * Phase 8 - admin-configurable max lifetime for CLI device-flow
   * sessions, in days. 0 = unbounded (default, mirrors GitHub CLI /
   * gh-style flows). Companies that need finite session lifetime
   * (eg. SOC 2 / ISO 27001 customers with mandated session
   * expiration) can set a positive value. The /api/auth/cli/refresh
   * endpoint rejects refresh attempts where now - issued_at exceeds
   * this value, regardless of the underlying refresh-token TTL.
   * Spec: specs/ai-governance/sessions/session-ttl.feature.
   */
  maxSessionDurationDays: Generated<number>
  /**
   * Every member of this organization can prove a second factor (D06). A
   * MEMBERSHIP CONDITION, not a session policy: it is asked when a member
   * reaches this organization's data, it holds the ones who cannot prove
   * one at an enrollment gate for THIS organization alone, and it ends no
   * session, ever. Turning it on signs nobody out.
   * Spec: specs/identity/mfa-and-session-shape.feature.
   */
  mfaRequired: Generated<boolean>
  /**
   * Whether members may connect the documentation playground to their own access with a
   * short-lived docs session key (ADR-185, amendment of 2026-11-13). On by default; turned off,
   * no key is minted and live ones stop authenticating. The pasted-key path is untouched.
   * Spec: modules/api-key/specs/docs-session-keys.feature.
   */
  docsPlaygroundSessionsEnabled: Generated<boolean>
  /**
   * When the docs playground was last turned off, written in the same update as the switch. A
   * docs session key minted at or before it never authenticates again, nor any key minted under
   * one, so turning it back on revives none of them (ADR-185, amendment of 2026-11-14).
   */
  docsPlaygroundSessionsDisabledAt: Timestamp | null
  /**
   * How many consecutive failed sign-ins lock an account (GAC-09).
   * 0 = never lock, which is the default and what every organization that
   * has not asked for this keeps. Counts EVERY credential together - a
   * password, a one-time code, a backup code, a passkey - because five per
   * factor hands an attacker one budget per factor.
   * Spec: specs/identity/org-account-lockout.feature.
   */
  lockoutAfterFailedAttempts: Generated<number>
  /** How long a lock lasts, in minutes. Ignored while the threshold is 0. */
  lockoutMinutes: Generated<number>
  /**
   * How long a BROWSER session may sit idle before it ends, in minutes
   * (GAC-10). 0 = no idle timeout. Activity rolls the window, which is the
   * "refreshed every hour" half of the control.
   *
   * A day for organizations created from now on; organizations that predate
   * the column keep 0 and are unchanged, because the migration adds the
   * column at zero and moves the default afterwards.
   *
   * NOT `maxSessionDurationDays`, which bounds CLI and device sessions and
   * is enforced when the CLI refreshes. Two surfaces, two lifetimes, two
   * decisions an administrator makes separately.
   * Spec: specs/identity/org-session-lifetime.feature.
   */
  sessionIdleTimeoutMinutes: Generated<number>
  /**
   * The ceiling on a browser session from the moment it was minted, in
   * minutes, which no amount of activity extends. 0 = no ceiling, the
   * default. Offered beside the idle timeout because a session kept alive
   * forever by a polling page is not really bounded.
   */
  sessionMaxLifetimeMinutes: Generated<number>
  signupData: unknown
  signedDPA: Generated<boolean>
  elasticsearchNodeUrl: string | null
  elasticsearchApiKey: string | null
  useCustomElasticsearch: Generated<boolean>
  s3Endpoint: string | null
  s3AccessKeyId: string | null
  s3SecretAccessKey: string | null
  s3Bucket: string | null
  useCustomS3: Generated<boolean>
  sentPlanLimitAlert: Timestamp | null
  ssoDomain: string | null
  ssoProvider: string | null
  /**
   * How colleagues on a matching domain get in (D12, ADR-117):
   * `off` (invisible to them entirely) | `request` (they ask, an admin
   * approves) | `auto` (a verified colleague walks straight in).
   *
   * `request` is the default for a self-serve organization; `auto` is never
   * a default and never inferred - an administrator turns it on, names the
   * domain in `joinDomains` while doing it, and the licence gate holds that
   * value while letting `request` through.
   */
  domainJoin: Generated<string>
  /**
   * The organisation role handed to a person admitted without an invitation:
   * a self-service domain join (`domainJoin`) or an SSO-admitted login.
   * `MEMBER` (default, a Full seat) or `DEVELOPER` (ADR-171). Set on the
   * organisation settings page only; the join flow never asks, and an
   * invitation always names its own role.
   */
  joinerRole: Generated<OrganizationUserRole>
  /**
   * The domains an administrator NAMED when turning automatic joining on.
   * Empty means automatic joining admits nobody whatever `domainJoin` says -
   * which is what keeps a setting left on from admitting a domain nobody
   * ever agreed to.
   */
  joinDomains: string[]
  presenceEnabled: Generated<boolean>
  /**
   * Org-level trace-sharing kill switch. Off = trace share links cannot be
   * created in ANY project of this org, and existing links stop resolving —
   * the org toggle is the global switch, the per-project
   * `Project.traceSharingEnabled` scopes the kill to a single project (mirrors
   * `presenceEnabled`). Effective sharing = org AND project. ADR-039.
   */
  traceSharingEnabled: Generated<boolean>
  /**
   * Org-configurable contact surfaced to users in CLI "contact your
   * admin" messages and the in-app budget-exceeded banner. Free text:
   * accepts an email, a URL (internal ticketing system), or any other
   * instruction. When null, callers fall back to the first ADMIN
   * member's email (`resolveSupportContact` helper).
   */
  supportContact: string | null
  /**
   * Signup intent, editable later as the org "Primary use" setting.
   * AGENT_GOVERNANCE lands "/" on /me, LLM_OPS on the project home.
   * NULL means intent unset (every pre-fork org and any org created
   * outside onboarding) — the home resolver keeps legacy behavior.
   * ADR-038.
   */
  primaryIntent: OrganizationIntent | null
  promoCode: string | null
  stripeCustomerId: string | null
  currency: Generated<Currency>
  pricingModel: Generated<PricingModel>
  license: string | null
  licenseExpiresAt: Timestamp | null
  licenseLastValidatedAt: Timestamp | null
  /**
   * A customer that runs LangWatch self-hosted. The hosted usage, budgets and
   * invoices of its connected installs attach to this organization (ADR-139).
   */
  selfHostedCustomer: Generated<boolean>
  /**
   * LangWatch-hosted services this organization has switched OFF, named the
   * way its license names them ("instant_evals"). A service the license is
   * entitled to is on unless it is listed here, so a customer who bought
   * hosted judging gets it without finding a settings page first, and an
   * administrator who switched one off keeps it off across license changes
   * (ADR-139, section 2). A license that names no service reaches nothing,
   * whatever this column holds.
   */
  connectServicesDisabled: Generated<string[]>
  /**
   * When a member of this organization switched Instant Evals on from the
   * search bar, which is the organization's consent to have judged text sent
   * to the judge's provider. Null until then: a self-serve organization is
   * never opted in on its behalf, and an enterprise one is switched on by an
   * operator through the release flag after asking, never here.
   */
  instantEvalsEnabledAt: Timestamp | null
  /**
   * The member who switched them on, so the audit trail names a person and
   * not just a moment.
   */
  instantEvalsEnabledByUserId: string | null
  /** When the last license sync succeeded. */
  connectLastSyncAt: Timestamp | null
  /**
   * The code of the failure the last sync ended on, cleared on the next
   * success. Settings, Connect renders the copy for it.
   */
  connectLastSyncError: string | null
}

/**
 * The registry of licenses LangWatch has issued (ADR-139). It lives on
 * LangWatch Cloud; a self-hosted install has the table and no rows.
 *
 * The row never holds the token an install presents. `tokenHash` is the
 * SHA-256 of that token, so read access to this table is not credential
 * access. Status is derived: revoked, superseded, past `expiresAt`, else active.
 */
export interface IssuedLicense {
  id: string
  /** The `licenseId` inside the signed license (`lic-<uuid>`). */
  licenseId: string
  /** SHA-256 of the `lwl_` token. Looked up on every credential resolution. */
  tokenHash: string
  /**
   * The customer organization on LangWatch Cloud. Null for a license that is
   * recorded but not linked to a customer yet, such as a self-serve purchase:
   * it resolves to nothing until an operator links it.
   */
  organizationId: string | null
  /** As signed into the license. */
  organizationName: string
  email: string
  planType: string
  maxMembers: number
  /** The lite seats sold: null when not stated or when unlimited. Never a stand-in number. */
  maxMembersLite: number | null
  /** True when the license sells unlimited lite seats; maxMembersLite is then null. */
  liteSeatsUnlimited: Generated<boolean>
  /** The Developer seats sold: null when not stated, which sells none. */
  maxMembersDeveloper: number | null
  issuedAt: Timestamp
  expiresAt: Timestamp
  source: IssuedLicenseSource
  /** The operator who issued or registered it. Null for the purchase flow and the script. */
  issuedById: string | null
  revokedAt: Timestamp | null
  revokedById: string | null
  revokedReason: string | null
  /** Set when the install first presents the license that replaced this one. */
  supersededAt: Timestamp | null
  /**
   * The id of the IssuedLicense this one replaces, when it was reissued. A
   * plain column ref (relationMode="prisma" cannot declare a self-relation on
   * Postgres); the registry service validates it. Unique, so a license is
   * replaced by one license at most.
   */
  replacesId: string | null
  /**
   * A reissued license waiting to be delivered over sync, encrypted. Erased
   * when the install first presents it. The only license text the registry holds.
   */
  pendingDeliveryLicense: string | null
  /** Hosted services this license may call, such as `instant_evals`. */
  services: Generated<string[]>
  /** The annual rate of one seat, in cents of `seatCurrency`. Needed to prorate added seats. */
  seatRateCents: number | null
  seatCurrency: Currency | null
  /** The prepaid usage commit for the term, in USD cents. */
  commitUsdCents: Generated<number>
  /** Whether usage may continue past the commit, up to `overageMaxUsdCents`. */
  overageEnabled: Generated<boolean>
  overageMaxUsdCents: number | null
  /** The install this license is bound to. Set on first use, cleared by an operator. */
  instanceId: string | null
  instanceBoundAt: Timestamp | null
  /**
   * The last license sync this install reported: when it arrived, the
   * LangWatch version that sent it, and the seats in use at that moment.
   * Null until the install syncs for the first time.
   */
  lastSyncAt: Timestamp | null
  lastSyncVersion: string | null
  reportedMembers: number | null
  reportedMembersLite: number | null
  /**
   * The managed gateway key (purpose CONNECT) that hosted-service traffic of
   * this license runs under. One per license, so spend rows name the install.
   * Created when the license first resolves. A plain column, not a relation.
   */
  virtualKeyId: string | null
  /**
   * The seats of the license this one replaced, set only when a seat change
   * raised a linked license. Billing reads it through licensing to invoice
   * the added seats, so the fact commits with the license or not at all.
   */
  seatsRaisedFrom: number | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

/**
 * The licence an organization activated and its dates, owned by licensing
 * (round 37 D6). A row whose key is null is a cleared licence. An organization
 * with no row is read from its Organization columns until the copy step runs.
 */
export interface OrganizationLicense {
  organizationId: string
  licenseKey: string | null
  expiresAt: Timestamp | null
  validatedAt: Timestamp | null
  updatedAt: Generated<Timestamp>
}

/**
 * A short code a fresh install pastes instead of a license blob (ADR-139,
 * section 5). Redeeming it mints the license the code describes and returns
 * it signed, which is what identifies a new customer without anybody emailing
 * a license key around.
 *
 * The row holds a hash of the code, never the code. Every lookup hashes what
 * the caller presented first, so the stored value is a step away from anything
 * that could be presented: read access to this table is not credential access.
 */
export interface ActivationCode {
  id: string
  /** SHA-256 of the normalised code. Looked up on every redemption. */
  codeHash: string
  /**
   * The last four characters, so an operator tells two codes apart in a list
   * without the list holding either one.
   */
  codeHint: string
  /** The customer this code belongs to, and the terms it mints. */
  organizationId: string
  organizationName: string
  email: string
  planType: string
  maxMembers: number
  /** The lite seats the minted license carries: null when unlimited. Never a stand-in number. */
  maxMembersLite: Generated<number | null>
  /** True when the minted license carries unlimited lite seats; maxMembersLite is then null. */
  liteSeatsUnlimited: Generated<boolean>
  /** The Developer seats the minted license carries: null when not stated, which sells none. */
  maxMembersDeveloper: number | null
  /** How long the license this code mints runs for. */
  licenseTermDays: number
  /** Hosted services the minted license may call. */
  services: Generated<string[]>
  /** When the code stops working, which is not the term of what it mints. */
  expiresAt: Timestamp
  /** A code a customer may redeem on more than one install, for a rollout. */
  reusable: Generated<boolean>
  /**
   * Set by the redemption that won, which is what makes a single-use code
   * single use: the write itself carries the condition.
   */
  redeemedAt: Timestamp | null
  redeemedByInstanceId: string | null
  /** The license the winning redemption minted. */
  issuedLicenseId: string | null
  /** How many installs have redeemed it, which only moves on a reusable code. */
  redemptionCount: Generated<number>
  revokedAt: Timestamp | null
  revokedById: string | null
  createdById: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

/**
 * How a connected customer is invoiced (ADR-139, section 7): one payment
 * provider customer, one quarterly usage subscription in USD, and the
 * commercial terms the annual invoice was raised from.
 *
 * One row per organization. Every write through it is addressed by a stored
 * provider id, so running onboarding or a renewal twice changes nothing.
 */
export interface ConnectedBillingAccount {
  id: string
  organizationId: string
  /** The customer at the payment provider. Invoiced, never charged. */
  stripeCustomerId: string
  /**
   * The one metered subscription, and the item hosted usage is reported on.
   * Null between the customer being created and the subscription existing,
   * which is where a half-finished onboarding resumes from.
   */
  usageSubscriptionId: string | null
  usageSubscriptionItemId: string | null
  termStartsAt: Timestamp
  termEndsAt: Timestamp
  /** Prepaid usage for the term, in USD cents. The credit grants add up to it. */
  commitUsdCents: Generated<number>
  /** The seat contract, which may be in another currency than hosted usage. */
  seatCurrency: Currency
  seatRateCents: number
  seats: number
  /**
   * `us_bank_transfer` or `eu_bank_transfer` when the customer pays into a
   * virtual bank account at the provider. Null when it wires to LangWatch's
   * own account and finance marks the invoice paid.
   */
  bankTransferType: string | null
  bankTransferCountry: string | null
  billingEmail: string
  /**
   * A renewal whose credit is waiting for the last usage invoice of the old
   * term to be finalized: the new term's end and commit, and who asked.
   */
  pendingRenewal: unknown
  createdById: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/**
 * A paid credit at the payment provider, which only metered usage draws down.
 * Kept here so a grant is created once per reason and the backoffice can show
 * what the customer holds without asking the provider.
 */
export interface ConnectedCreditGrant {
  id: string
  /**
   * The `ConnectedBillingAccount` row. A plain column ref, like
   * `IssuedLicense.replacesId`: relationMode="prisma" carries no foreign keys.
   */
  accountId: string
  /** The grant at the payment provider. Unique, so a retry records one row. */
  stripeCreditGrantId: string
  amountUsdCents: number
  kind: ConnectedCreditGrantKind
  /**
   * The term this credit was granted for, and when the provider lets it lapse:
   * after the last invoice of that term has been finalized, never on the day
   * the term ends.
   */
  termEndsAt: Timestamp
  expiresAt: Timestamp
  createdAt: Generated<Timestamp>
}

/**
 * An invoice LangWatch raised for a connected customer, as the backoffice
 * reads it.
 */
export interface ConnectedInvoice {
  id: string
  /** The `ConnectedBillingAccount` row, as a plain column ref. */
  accountId: string
  stripeInvoiceId: string
  kind: ConnectedInvoiceKind
  currency: Currency
  amountCents: number
  /** The provider's own status: `draft`, `open`, `paid`, `void` or `uncollectible`. */
  status: string
  paidOutOfBandAt: Timestamp | null
  termStartsAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/**
 * Every SCIM request we could attribute, and what we answered (ADR-126).
 *
 * EVIDENCE, NOT TRUTH. A request authors nothing — an HTTP round trip that
 * was refused decided nothing — so this is a table and not an event: nothing
 * downstream derives from it, rebuilding the world from the log must not
 * depend on how often a provider retried a GET, and rows age out of a
 * retention window an event log would never have.
 *
 * NO ROW FOR A REQUEST WE CANNOT ATTRIBUTE. A token that does not verify is
 * answered 401 and recorded nowhere: we do not know whose organization to
 * file it under, and a table unauthenticated traffic can write is a table
 * anybody on the internet can fill. What answers that case instead is
 * `ScimToken.lastUsedAt` being null. The 403 IS recorded — a lapsed plan is
 * a credential we recognize.
 *
 * THE PAYLOAD RULE HOLDS. A method, a resource name, a status, a reason slug
 * and our own sentence. Never a token, never a header, never a provider's
 * raw message — the same reason `scim_apply_failed` carries a code and not a
 * message.
 */
export interface ScimRequestLog {
  id: string
  organizationId: string
  /**
   * Null when the refusal resolved an organization and stopped before the
   * connection — a lapsed plan is answered off the token alone.
   */
  connectionId: string | null
  method: string
  /**
   * The SCIM resource asked for ("Users", "Users/:id", "Groups"), never the
   * raw path: a path carries ids and query strings, and this is read by
   * people rather than matched by machines.
   */
  resource: string
  status: number
  /**
   * A stable slug, only on a refusal, so a reader branches on this and never
   * on the sentence beside it.
   */
  reason: string | null
  /** Our own short sentence, customer-safe. */
  detail: string | null
  occurredAt: Generated<Timestamp>
}

export interface ScimToken {
  id: string
  organizationId: string
  /**
   * D08: the SSO connection this token was issued for, and its ENTIRE write
   * authority — a token minted for one connection can never touch another
   * connection's people. No relation, matching `SsoConnection` itself: that
   * projection carries no foreign keys, so neither may a row pointing at it.
   *
   * Nullable only for the tokens that predate the column. They were minted
   * when "which organization" was the whole answer, so they keep the
   * organization-wide authority they were sold with; the backfill points
   * each at its organization's first connection where there is one. Minting
   * a NEW token without a connection is refused (`scim_connection_required`).
   */
  connectionId: string | null
  /**
   * UNIQUE, and that is a tenancy boundary rather than tidiness.
   *
   * The lookup that turns a bearer token into an organization is keyed on
   * this column alone — it has to be, since the token is the only thing the
   * caller presents. While every token was 32 random bytes a collision was
   * impossible and the missing constraint cost nothing. An administrator may
   * now supply a value they already hold, so two organizations CAN choose the
   * same string, and without this the lookup would answer with whichever row
   * the planner reached first: one customer's directory provisioning and
   * deleting another customer's people.
   */
  hashedToken: string
  /**
   * How `hashedToken` was derived, so the two schemes can coexist while
   * tokens minted under the old one are still in use.
   *
   * `sha256` is a bare digest, which is correct for 32 random bytes and wrong
   * for a value a person chose — a database dump of those is a wordlist away
   * from live credentials. `hmac-sha256` is keyed on the deployment's own
   * secret, so a dump on its own is inert. Every new row is the latter.
   */
  hashScheme: Generated<string>
  description: string | null
  createdAt: Generated<Timestamp>
  lastUsedAt: Timestamp | null
}

/**
 * D08 — who the directory means, per connection.
 *
 * The identity provider's own identifier for a person, which is what
 * survives their email changing: an address is not identity, people change
 * theirs, and the same person carries different identifiers on two different
 * connections. So the PAIR is the key and `externalId` alone never is.
 *
 * No relation to `SsoConnection` (that projection carries no foreign keys);
 * the `User` relation is real, so a deleted user takes their directory
 * identities with them rather than stranding a mapping that resolves to
 * nobody.
 */
export interface ScimExternalId {
  organizationId: string
  id: string
  connectionId: string
  externalId: string
  userId: string
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/** Directory profile and lifecycle are tenant facts, independent of the shared account. */
export interface ScimUserResource {
  organizationId: string
  userId: string
  userName: string
  name: string | null
  active: Generated<boolean>
  deletedAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/**
 * A connection's managed users, including resources without an externalId.
 * Scalar keys only; readers exclude users that no longer exist.
 */
export interface ScimDirectoryUser {
  organizationId: string
  connectionId: string
  userId: string
}

/**
 * D08 — one connection's directory sync, as the `scim-sync` pipeline's fold
 * writes it.
 *
 * Pure event truth: every column is fold-written, rows are never deleted,
 * and REVOKED is a tombstone rather than an absence — a dead letter that
 * vanished when a connection was torn down would be a removal nobody could
 * check afterwards.
 *
 * No relation to Organization or SsoConnection, matching `SsoConnection`
 * itself: a projection is rebuilt from events, and a foreign key would make
 * a replay depend on relational state the log does not carry.
 */
export interface ScimSyncState {
  /** The aggregate: the connection's id, which IS its sync's id. */
  id: string
  connectionId: string
  organizationId: string
  /** TOKEN_ISSUED | SYNCING | ERROR | REVOKED */
  state: string
  lastPushedAt: Timestamp | null
  /**
   * The failure standing right now: { op, errorCode, attempts, retiredAtMs,
   * userId, occurredAtMs }. A reason CODE and ids — never a token, never a
   * secret, never an internal hostname.
   */
  lastFailure: unknown
  /**
   * Every failure retired without ever being applied, newest last. The dead
   * letters an administrator has to see; retiring never drops one.
   */
  deadLetters: Generated<unknown>
  /** revoke | teardown, once the sync has ended. */
  revokedCause: string | null
  occurredAt: Timestamp
  lastEventId: string
  acceptedAt: Timestamp
  projectionVersion: string
  /**
   * Event-derived, not database-managed: `@default(now())` / `@updatedAt`
   * would make a replayed row differ from the row it rebuilds, which is
   * exactly the parity this projection has to keep.
   */
  createdAt: Timestamp
  updatedAt: Timestamp
}

/**
 * SCIM's own copy of an organization's SSO connections: a peer fold over
 * identity's connection facts (ARCHITECTURE §9), folded with identity-contract's
 * reducer, one row per connection. SCIM reads this instead of asking identity.
 * Rebuilt from identity's event log by a projection replay, so no relation:
 * a foreign key would make a replay depend on state the log does not carry.
 */
export interface ScimSsoConnectionView {
  /** The aggregate: identity's connection id. */
  id: string
  organizationId: string
  /** The connection as the reducer folded it (SsoConnectionState + LastEventOccurredAt). */
  folded: unknown
  /** The event ids of the last delivery folded in, so a redelivery folds once. */
  appliedEventIds: Generated<string[]>
  /** A row of another version reads as undecodable and is re-folded from the log. */
  projectionVersion: string
  /** Event-derived, so a replayed row equals the row it rebuilds. */
  createdAt: Timestamp
  updatedAt: Timestamp
}

/**
 * Organization's monthly seat peak, folded from its seat changes (modules/organization, ADR-002).
 * A read model, one row per organization: rebuildable by replaying the event log, so no relation.
 * Epochs in milliseconds.
 * Organization's seat-peak counts (ADR-002): one row per organization once counting opened.
 */
export interface OrganizationSeatPeak {
  /** The organization, which is also the aggregate id. */
  organizationId: string
  /** When the opening snapshot was taken, epoch milliseconds. */
  openedAt: number
  fullSeats: number
  liteSeats: number
  developerSeats: number
  /** The latest change time applied to any holder; invitations expire up to it. */
  latestChangeAt: number
  /** Changes whose `from` disagreed with the seat the fold held. */
  mismatches: number
}

/** One member's or pending invitation's seat as the seat peak last applied it (ADR-002). */
export interface OrganizationSeatPeakHolder {
  organizationId: string
  /** `member`, `invitation` or `hold` (a seat held after somebody left). */
  holderKind: string
  /** The member's or the hold's userId, or the invitation's id. */
  holderId: string
  /** `full`, `lite`, `developer` or `none`. */
  seat: string
  lastChangedAt: number
  /** Absent for a seat the opening snapshot recorded. */
  lastChangeId: string | null
  /** An invitation's expiry; null for a member or an invitation that never expires. */
  expiresAt: number | null
  /** The fold gave the invitation's seat back at `expiresAt`. */
  expiredByFold: Generated<boolean>
  /** A hold whose seat a live full invitation to its address took: that invitation's id. */
  suppressedBy: string | null
  /** The fold gave a suppressed hold its seat back when that invitation expired. */
  restoredByFold: Generated<boolean>
}

/** One calendar month's (UTC) seat peaks for an organization (ADR-002). */
export interface OrganizationSeatPeakMonth {
  organizationId: string
  /** `YYYY-MM`. */
  month: string
  fullPeak: number
  fullPeakAt: number
  litePeak: number
  litePeakAt: number
  fullAtEnd: number
  liteAtEnd: number
}

export interface Project {
  id: string
  name: string
  slug: string
  apiKey: string
  /**
   * When the legacy apiKey last authenticated a request, written at most once a
   * minute per project and process. Spec: modules/api-key/specs/request-credential-type.feature
   */
  legacyApiKeyLastUsedAt: Timestamp | null
  /**
   * Random per-project secret whose SHA-256 digest is this project's LangWatchQL
   * SQL tenant capability in ClickHouse. Deliberately NOT the API key: the
   * credential a caller authenticates with and the capability the database
   * resolves a tenant from stay decoupled, so rotating one never invalidates
   * the other, and a leak of the ClickHouse key map's input discloses nothing
   * a caller could authenticate with. Generated by the database on insert.
   */
  lwqlKey: Generated<string>
  teamId: string
  language: string
  framework: string
  /**
   * Project kind:
   * - "application": user-facing workspace where developers send their
   * own app traces. The default for everything created via the
   * onboarding/quickstart flow.
   * - "internal_governance": hidden per-org routing/tenancy artifact
   * for IngestionSource ingest data. Auto-created lazily on the
   * first IngestionSource mint via ensureHiddenGovernanceProject().
   * MUST be filtered out of every user-visible Project surface
   * (project picker, /api/v1/projects, billing exports, RBAC
   * pickers, etc.) - the regression-test invariant in
   * specs/ai-gateway/governance/ui-contract.feature codifies this.
   * - "aggregate": a project that owns no traces and reads its member
   * projects through shared project-reader grants (ADR-177). Hidden
   * from every "send traces here" picker and from ingestion; visible in
   * the project switcher to organisation admins only.
   * Stored as a free-form string (matching IngestionSource.sourceType
   * pattern) so future routing kinds can ship without schema migrations.
   */
  kind: Generated<string>
  /**
   * The scope rule an aggregate project materialises into grants. Null on
   * every other kind; decided in ADR-177. Shapes: { kind: "all-personal" } |
   * { kind: "personal-by-department", departmentId } |
   * { kind: "explicit", projectIds: string[] }.
   */
  aggregateRule: unknown
  firstMessage: Generated<boolean>
  integrated: Generated<boolean>
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  userLinkTemplate: string | null
  traceSharingEnabled: Generated<boolean>
  presenceEnabled: Generated<boolean>
  s3Endpoint: string | null
  s3AccessKeyId: string | null
  s3SecretAccessKey: string | null
  s3Bucket: string | null
  archivedAt: Timestamp | null
  isPersonal: Generated<boolean>
  ownerUserId: string | null
  /**
   * Per-feature unlock flags for personal projects only. Shape:
   * `{ evaluations: bool, datasets: bool, annotations: bool, automations: bool }`.
   * Empty `{}` (default) reads as all-false at the service layer -
   * existing rows pre-migration require no backfill. The bundle is
   * a UI/nav predicate, NOT an auth gate (the underlying tRPC
   * routers stay open even when the bundle is off; admin tooling +
   * offline migration scripts can still operate on the data).
   * Team / org-scope projects are not feature-gated - this column
   * is meaningful only when `isPersonal=true`.
   */
  personalFeatures: Generated<unknown>
  /**
   * Department assignment for spend attribution (the project / agent
   * lens, used when a trace has no principal user). Accounting-only.
   */
  departmentId: string | null
  /**
   * Per-project Langy egress allow-list (ADR-076). A nullable JSON array of
   * host patterns (`["registry.npmjs.org", "*.internal.acme.com"]`) that the
   * project's Langy workers may reach outbound. The *presence* of the list is
   * the enforcement mode: null/absent ⇒ monitor-only (watch, never block);
   * non-empty ⇒ the worker's egress adapter restricts outbound to
   * floor ∪ this list. Enforced in the agent pod's egress adapter — NOT the
   * gateway — so it lives on the project, not the VirtualKey. Resolved by
   * LangyCredentialService.getEgressAllowlist and threaded through the
   * per-request credentials envelope. See specs/langy/langy-egress-enforcement.feature.
   */
  langyEgressAllowlist: unknown
  /**
   * When a coding-agent session was last folded for this project, and when a
   * pull request was last mapped for a branch one of those sessions ran on.
   * Both are instrumentation signals, not counters: the project sidebar shows
   * its Sessions and Pull requests destinations only while the matching column
   * is inside the recency window, so a project that never sends coding-agent
   * telemetry never grows the two links, and one that stops sending loses
   * them again. Written through the throttled touch methods on
   * ProjectService, which skip the write while the stored value is still
   * recent, so a busy project costs at most one update an hour per column.
   * No index: every read is by project id, which is the primary key.
   * See specs/coding-agent/project-menu-links.feature.
   */
  lastCodingAgentSessionAt: Timestamp | null
  lastCodingAgentPullRequestAt: Timestamp | null
}

/**
 * A department is a single-valued accounting dimension, separate from
 * RBAC, for rolling up AI spend by org unit (Engineering, Marketing,
 * ...). Assigned to OrganizationUser, Team, and Project via their
 * `departmentId` column. Pure accounting: never grants or restricts
 * access. See specs/ai-gateway/governance/departments.feature.
 */
export interface Department {
  id: string
  organizationId: string
  name: string
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  archivedAt: Timestamp | null
}

/**
 * One dated person → department link (ADR-128 §13, #7882).
 *
 * `OrganizationUser.departmentId` stays the CURRENT assignment every screen
 * reads; these rows are its history, appended by the same `assignUser` write
 * so the two cannot drift. They exist for the cost reads: January's spend
 * stays with January's department, so a read resolves against the link that
 * was open on the cost's day, not against today's pointer. Members with no
 * link on a day resolve to "unassigned", never to an error.
 *
 * Same shape as IdentityMatch's dated links, same one-open-link guard: the
 * partial unique index in the migration rejects a second open row per
 * (organizationId, userId) with SQLSTATE 23505.
 */
export interface DepartmentMembershipHistory {
  id: string
  organizationId: string
  userId: string
  departmentId: string
  validFrom: Timestamp
  /** Null means the link is open. Reassignment and clearing close it. */
  validTo: Timestamp | null
  createdAt: Generated<Timestamp>
}

/**
 * Per-(organization, tool) override for which `langwatch <tool>` paths the
 * CLI may use: `allowVk` gates the gateway / virtual-key path (Path A) and
 * `allowOtelDirect` gates the direct-OTLP ingestion path (Path B). A missing
 * row resolves to the hardcoded defaults (both true for claude/codex/gemini/
 * opencode; cursor allows the gateway path only), so the policy is purely
 * additive: disabling a path is an explicit admin choice. The CLI caches the
 * resolved map at login and gates path selection on it.
 * See specs/ai-gateway/governance/cli-tool-mode-policy.feature.
 */
export interface PlatformToolPolicy {
  id: string
  organizationId: string
  toolSlug: string
  allowVk: Generated<boolean>
  allowOtelDirect: Generated<boolean>
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface OrganizationInvite {
  id: string
  email: string
  inviteCode: string
  expiration: Timestamp | null
  status: Generated<INVITE_STATUS>
  organizationId: string
  teamIds: string
  teamAssignments: unknown
  role: OrganizationUserRole
  requestedBy: string | null
  subscriptionId: string | null
  acceptedByUserId: string | null
  acceptedViaIdentifierId: string | null
  /** Invites the person who pays with the Billing grant and no seat (ADR-205). */
  billingOnly: Generated<boolean>
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface Monitor {
  id: string
  projectId: string
  experimentId: string | null
  evaluatorId: string | null
  checkType: string
  name: string
  slug: string
  executionMode: Generated<EvaluationExecutionMode>
  enabled: Generated<boolean>
  preconditions: unknown
  parameters: unknown
  mappings: unknown
  sample: Generated<number>
  level: Generated<string>
  threadIdleTimeout: number | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  autoPausedAt: Timestamp | null
  autoPausedReason: string | null
  evaluatorLabel: string | null
  createdById: string | null
  updatedById: string | null
}

export interface Cost {
  id: string
  projectId: string
  costType: CostType
  costName: string | null
  referenceType: CostReferenceType
  referenceId: string
  amount: number
  currency: string
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  extraInfo: unknown
}

/**
 * PROJECTION of the topic-clustering event stream (ADR-051): topics are
 * `topics_recorded` events; this table is written ONLY by
 * TopicModelFoldProjection (cursor: TopicModelProjection) and is
 * rebuildable by replay. Kept named `Topic` because it is the domain
 * entity every surface reads — the `*Projection` suffix is reserved for
 * cursor/read-model tables.
 */
export interface Topic {
  id: string
  projectId: string
  name: string
  parentId: string | null
  embeddings_model: string
  centroid: unknown
  p95Distance: number
  automaticallyGenerated: Generated<boolean>
  /**
   * The `topics_recorded` event that recorded this topic (its id in the
   * event log). Null means the row predates event ownership and has not
   * been seeded onto the stream yet; the projection sets it on every row
   * it writes.
   */
  lastEventId: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface Dataset {
  id: string
  projectId: string
  name: string
  slug: string
  columnTypes: unknown
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  archivedAt: Timestamp | null
  mapping: unknown
  useS3: Generated<boolean>
  s3RecordCount: number | null
  contentLayout: Generated<string>
  status: Generated<string>
  statusError: string | null
  stagingKey: string | null
  sourceStoredObjectId: string | null
  uploadFilename: string | null
  rowCount: number | null
  sizeBytes: ColumnType<string, bigint | number | string, bigint | number | string> | null
  chunkCount: number | null
  chunkOffsets: unknown
}

export interface DatasetRecord {
  id: string
  datasetId: string
  projectId: string
  entry: unknown
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/** One immutable snapshot of a dataset's columns and rows; the rows are JSONL chunks in object storage. */
export interface DatasetVersion {
  id: string
  projectId: string
  datasetId: string
  version: number
  /** A name given by hand; a named version is never thinned by retention. */
  name: string | null
  /** "autosave", "manual", "experiment_run", "restore" or "initial". */
  trigger: string
  /** SHA-256 over the columns and the rows in order; equal content saves no new version. */
  contentHash: string
  columnTypes: unknown
  rowCount: number
  sizeBytes: ColumnType<string, bigint | number | string, bigint | number | string>
  chunkCount: number
  restoredFromVersion: number | null
  authorId: string | null
  /** "user", "api_key" or "system": who saved the version when authorId is null. */
  authorKind: string | null
  /** Set when an experiment pins this version; a pinned version is never thinned. */
  pinnedAt: Timestamp | null
  createdAt: Generated<Timestamp>
}

/** An experiment run that read one dataset version; a version a run read is never thinned. */
export interface DatasetVersionRun {
  id: string
  projectId: string
  datasetId: string
  versionId: string
  runId: string
  experimentId: string | null
  createdAt: Generated<Timestamp>
}

export interface Dashboard {
  id: string
  /** The project that owns the board; only there can it be edited. */
  projectId: string
  name: string
  /**
   * Legacy board sort position. The Dashboards v2 app orders boards by a
   * member's favourites and never writes this; kept because the legacy
   * Analytics reports section and the REST reorder endpoint still write it,
   * and a rolling deploy's older image still reads it.
   */
  order: Generated<number>
  /**
   * Free text under the board's name; null until a member writes one.
   * Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
   */
  description: string | null
  /**
   * The member who created it; null for boards that predate the column.
   * Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
   */
  createdById: string | null
  /**
   * Who sees it. Only the author changes it; a board with no author stays PROJECT.
   * Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
   */
  scope: Generated<DashboardScope>
  /**
   * The owning project's organization, stamped when the board is first set to
   * ORGANIZATION; null before. No relation, per the database rule.
   * Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
   */
  organizationId: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/**
 * A member's star, in their own order within a project: a board or a From
 * LangWatch template. Personal: one member's stars never show to another. No
 * foreign key (per the database rule): the dashboard service removes a board's
 * favourites when it deletes the board, so a star never outlives its dashboard.
 */
export interface DashboardFavourite {
  id: string
  userId: string
  /**
   * A DashboardsV2 board's id (copied boards kept their Dashboard id); null when
   * the star is a template.
   */
  dashboardId: string | null
  /** The starred From LangWatch template id; null when the star is a board. */
  templateId: string | null
  projectId: string
  /** The member's drag order among their own stars in the project. */
  position: number
  createdAt: Generated<Timestamp>
}

/**
 * A Dashboards v2 board, apart from the Analytics Reports boards in Dashboard. Boards
 * made before the split were copied here with their ids, so stars and insight
 * pointers still name them. No relation, per the database rule.
 */
export interface DashboardsV2 {
  id: string
  /** The project that owns the board; only there can it be edited. */
  projectId: string
  name: string
  /** Free text under the board's name; null until a member writes one. */
  description: string | null
  /** The member who created it; null for boards that predate authorship. */
  createdById: string | null
  /** Who sees it. Only the author changes it; a board with no author stays PROJECT. */
  scope: Generated<DashboardScope>
  /**
   * The owning project's organization, stamped when the board is first set to
   * ORGANIZATION; null before.
   */
  organizationId: string | null
  /**
   * The product whose sidebar lists it: llm-ops, me, gateway or governance (ADR-002,
   * decision 8). Every board made before homes existed is llm-ops.
   */
  homeProduct: Generated<string>
  /**
   * Deprecated: unused since each widget query names its own sources (dashboards-v2 AC200,
   * 2026-10-28); dropped in a later migration, once no serving release reads it.
   */
  excludedOrigins: unknown
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/**
 * A Dashboards v2 widget: a srcdoc definition placed on a DashboardsV2 board, or
 * unplaced. Widgets made before the split were copied from CustomGraph with their
 * ids. No relation, per the database rule.
 */
export interface DashboardWidget {
  id: string
  projectId: string
  /** The DashboardsV2 board it sits on; null while unplaced. */
  dashboardId: string | null
  name: string
  /** The widget's srcdoc definition. */
  definition: unknown
  gridColumn: Generated<number>
  gridRow: Generated<number>
  colSpan: Generated<number>
  rowSpan: Generated<number>
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/**
 * One insight in its owner's inbox in a project, folded from its `insight` aggregate
 * (modules/insight). A read model: rebuildable by replaying the event log. Epochs in
 * milliseconds.
 */
export interface InsightProjection {
  /** The insight id, which is also its aggregate id. */
  id: string
  projectId: string
  title: string
  body: string
  /** bad | watch | good */
  tone: string
  topic: string | null
  validDays: number
  lwql: string | null
  sourceConversationId: string | null
  sourceMessageId: string | null
  filedByUserId: string | null
  filedAt: number
  /** Set by a later run that found the insight still true. */
  renewedAt: number | null
  createdAt: number
  updatedAt: number
  occurredAt: number
  acceptedAt: number
  lastEventId: string
  projectionVersion: string
  /**
   * The board and widget it came from. A pointer only, never a relation: the row reads the
   * same once either is deleted, so the names are kept as filed.
   */
  boardId: string | null
  boardName: string | null
  widgetId: string | null
  widgetName: string | null
  /** chat | run */
  filedVia: Generated<string>
  /** The fixed window `lwql` replays over (epoch milliseconds, half-open) and its step. */
  replayStart: number | null
  replayEnd: number | null
  replayGranularitySeconds: number | null
  /** The values in force when it was filed: `{ period, parameters }`. */
  replayContext: unknown
  /**
   * Whose insight it is: the one person who reads it. NULL on a row folded before the
   * column existed, which is owned by whoever filed it (`filedByUserId`).
   */
  ownerUserId: string | null
  /**
   * The stable key a daily run matches a repeated finding by, so it renews the insight
   * and files no twin. NULL for an insight saved from a chat and on a row folded before
   * the column existed.
   */
  findingKey: string | null
}

/**
 * One reader's own state on one insight (seen, done, kept), folded from the same stream.
 * An insight has one reader today, its owner; a later forward may add more.
 */
export interface InsightReaderProjection {
  id: string
  projectId: string
  insightId: string
  userId: string
  seenAt: number | null
  archivedAt: number | null
  keptAt: number | null
  createdAt: number
  updatedAt: number
  occurredAt: number
  acceptedAt: number
  lastEventId: string
  projectionVersion: string
}

/**
 * One person's daily insights run on one board in a project, folded from its
 * `insight_daily_schedule` aggregate (modules/insight). A read model: rebuildable by
 * replaying the event log. Epochs in milliseconds.
 */
export interface InsightDailyScheduleProjection {
  /** The schedule id, which is also its aggregate id: one per project, person and board. */
  id: string
  projectId: string
  userId: string
  /** dashboard | template. A pointer only, never a relation: a run checks the board is there. */
  boardKind: string
  boardId: string
  boardName: string
  /** on | off | undecided. */
  state: Generated<string>
  /**
   * board_deleted | skipped_days | member_removed | project_archived. Why the system turned
   * it off. NULL while it is on, and when the person turned it off.
   */
  offReason: string | null
  hour: number | null
  timezone: string | null
  maxInsights: number | null
  /** The last run that recorded an outcome. NULL until one did. */
  lastRunId: string | null
  lastRunAt: number | null
  /** filed | nothing | failed | skipped */
  lastRunOutcome: string | null
  lastRunReason: string | null
  lastRunFiled: number | null
  /** How many insights the run found still true and renewed. */
  lastRunRenewed: number | null
  lastRunConversationId: string | null
  /** The scheduled days in a row whose run ended skipped, and the slot of the newest counted. */
  consecutiveSkips: Generated<number>
  skipStreakSlot: number | null
  createdAt: number
  updatedAt: number
  occurredAt: number
  acceptedAt: number
  lastEventId: string
  projectionVersion: string
}

export interface SavedView {
  id: string
  projectId: string
  userId: string | null
  name: string
  filters: Generated<unknown>
  query: string | null
  period: unknown
  order: Generated<number>
  /**
   * Discriminator for the two storage shapes that share this table:
   * - "v1-traces-filter" (default) - legacy /messages page filter map saved views
   * - "v2-traces-lens"             - new traces v2 lens configs (columns, addons, grouping, sort, filterText) encoded inside `filters` JSON
   * Each client filters by its own kind on read so neither sees the other's rows.
   */
  kind: Generated<string>
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface CustomGraph {
  id: string
  projectId: string
  name: string
  /** The chart's definition, in the shape its `kind` names. */
  graph: unknown
  filters: unknown
  /**
   * Discriminator for the chart shapes that share this table:
   * - "builder" (default) - the chart builder's own `graph` payload
   * - "workbench_sql"     - a saved LangWatchQL workbench chart, whose
   * `graph` holds the versioned { sql, parameters, vegaLiteSpec }
   * definition (see server/analytics/saved-workbench-charts)
   * - "dashboard_srcdoc"  - frozen legacy copies of Dashboards v2 widgets, which
   * now live in DashboardWidget; nothing new writes this kind
   * Each reader filters by its own kind, so neither sees the other's rows and
   * a builder chart can never be interpreted as a workbench definition.
   */
  kind: Generated<string>
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  dashboardId: string | null
  gridColumn: Generated<number>
  gridRow: Generated<number>
  colSpan: Generated<number>
  rowSpan: Generated<number>
}

export interface BatchEvaluation {
  id: string
  experimentId: string
  projectId: string
  data: Generated<unknown>
  status: string
  score: number
  label: string | null
  passed: boolean
  details: string
  cost: number
  datasetSlug: string
  datasetId: string
  evaluation: string
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface Trigger {
  id: string
  name: string
  projectId: string
  action: TriggerAction
  triggerKind: Generated<TriggerKind>
  actionParams: unknown
  filters: unknown
  filterQuery: string | null
  lastRunAt: Generated<number>
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  active: Generated<boolean>
  pausedReason: string | null
  pausedAt: Timestamp | null
  message: string | null
  deleted: Generated<boolean>
  alertType: AlertType | null
  slackTemplateType: string | null
  slackTemplate: string | null
  emailSubjectTemplate: string | null
  emailBodyTemplate: string | null
  notificationCadence: Generated<string>
  traceDebounceMs: Generated<number>
  customGraphId: string | null
}

/**
 * One named Slack connection (ADR-093 §5a); an organization holds as many as
 * it likes, each scoped to the organization or to one project. The secret is
 * AES-256-GCM ciphertext and never returned to a client. `secretFingerprint`
 * is an HMAC of the plaintext, so one secret is one connection per scope
 * without decrypting anything. No foreign key: `scopeId` is
 * polymorphic (ADR-021), as on RetentionPolicy and CustomLLMModelCost.
 */
export interface SlackIntegration {
  id: string
  name: string
  kind: SlackIntegrationKind
  scopeType: SlackIntegrationScopeType
  scopeId: string
  organizationId: string
  botTokenEncrypted: string | null
  webhookUrlEncrypted: string | null
  secretFingerprint: string
  secretHint: string
  slackTeamId: string | null
  slackTeamName: string | null
  createdById: string
  updatedById: string
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/**
 * An automation holding a Slack connection in use (ARCHITECTURE.md §3, the
 * claims ruling): a claimed connection is never deleted. `projectId` is the
 * claimant's project; `organizationId` is the connection's tenancy anchor.
 * No foreign key: slack's service refuses to delete a claimed connection.
 */
export interface SlackConnectionClaim {
  connectionId: string
  claimantId: string
  claimantLabel: string
  organizationId: string
  projectId: string
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/**
 * Org-anchored outbound webhook endpoint (the webhook platform, enterprise
 * only). The signing secret is AES-256-GCM encrypted at rest (shared
 * encrypt/decrypt, CREDENTIALS_SECRET) and returned in plaintext exactly
 * once, at create or roll time. `enabledEvents` holds registry event types
 * or family wildcards ("gateway.*", "*"). Failure-streak columns drive the
 * 72h auto-disable. Spec: specs/webhooks/webhook-endpoints.feature.
 */
export interface WebhookEndpoint {
  id: string
  organizationId: string
  /**
   * Where this endpoint delivers. Everything above the last hop is shared:
   * same batching, same retry ladder, same delivery log, same signature over
   * the same bytes.
   */
  destinationKind: Generated<WebhookDestinationKind>
  /**
   * The receiver URL. Present on `http` endpoints and null on the others;
   * the per-kind shape is pinned by a CHECK constraint Prisma cannot model.
   */
  url: string | null
  /**
   * The Amazon SQS queue this endpoint delivers to, on `sqs` endpoints. The
   * region and the owning account are read off it rather than configured
   * separately, so they can never disagree with the queue.
   */
  sqsQueueUrl: string | null
  /**
   * The role we assume to write to the queue, and the ExternalId that proves
   * the request came from us rather than from anyone who learned the role's
   * name. The recommended credential mode.
   */
  sqsRoleArn: string | null
  sqsExternalId: string | null
  /**
   * A static key pair for the queue, encrypted at rest exactly like the
   * signing secret. Both halves are present or both are absent: a half-filled
   * pair is what the AWS SDK reads as "we have credentials" and stops looking
   * with, so the CHECK constraint refuses it.
   */
  sqsAccessKeyId: string | null
  sqsSecretAccessKeyEncrypted: string | null
  secretEncrypted: string
  /**
   * How deliveries are formatted and signed. Null: the batch envelope signed
   * `t=,v1=`. `legacy_sha256`: one raw message per POST signed `sha256=`, kept
   * only for endpoints migrated from governance anomaly destinations.
   */
  signatureScheme: string | null
  /**
   * Set only by governance's anomaly destination migration: a create that
   * repeats it in the same organization answers the endpoint it made. An
   * archived endpoint gives its key up, so a retry after cleanup creates anew.
   */
  idempotencyKey: string | null
  /**
   * The secret this endpoint signed with before the last roll. Deliveries
   * carry a second `v1=` computed from it until it expires, so a receiver
   * swaps secrets on its own schedule instead of dropping deliveries in the
   * gap. Null once the window closes or if the secret was never rolled.
   */
  previousSecretEncrypted: string | null
  /** When the previous secret stops being signed with and accepted. */
  previousSecretExpiresAt: Timestamp | null
  enabledEvents: Generated<string[]>
  status: Generated<WebhookEndpointStatus>
  /** Why the endpoint is disabled: "auto_failures_72h" or "manual". */
  disabledReason: string | null
  disabledAt: Timestamp | null
  /**
   * Start of the current consecutive-failure streak; cleared on any 2xx.
   * 72 hours of unbroken failures from this instant auto-disables.
   */
  failingSince: Timestamp | null
  /**
   * Delivery controls, each server-bounded (see webhookDeliveryControls
   * constants): envelopes per POST, coalescing window before a partial
   * batch ships, and concurrent POSTs to this receiver.
   */
  maxBatchSize: Generated<number>
  maxBatchDelayMs: Generated<number>
  maxInFlight: Generated<number>
  /**
   * Self-hosted only: deliveries accept the receiver's self-signed TLS
   * certificate. Off by default; the hosted product ignores it.
   */
  allowSelfSignedCertificate: Generated<boolean>
  lastSuccessAt: Timestamp | null
  lastFailureAt: Timestamp | null
  archivedAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/**
 * The per-ATTEMPT delivery log for BOTH webhook channels: the endpoints
 * platform's batches and the automations channel's trigger fires. One row per
 * HTTP attempt; `dispatchId` groups every attempt of one logical send (the
 * value of the X-LangWatch-Delivery-Id or X-LangWatch-Event-Id header). A slim
 * facts table: outcome, status, latency, a capped error message, and a
 * truncated copy of a failure response. Our request content (URL, headers,
 * body) is never stored. Pruned after 30 days.
 *
 * The two channels used to keep separate tables with the same columns and two
 * prune implementations. They are one table with a discriminator now, so the
 * tenancy columns of the channel that did not write a row are null.
 */
export interface WebhookEndpointDelivery {
  id: string
  channel: Generated<WebhookDeliveryChannel>
  /** Platform tenancy. Null on automations rows. */
  organizationId: string | null
  endpointId: string | null
  /** Automations tenancy. Null on platform rows. */
  projectId: string | null
  triggerId: string | null
  dispatchId: string
  /**
   * 1-based attempt and the batch's envelope count. Both null on automations
   * rows, which record neither: one trigger fire is one event, and the
   * channel's sender does not thread its attempt number into the log.
   */
  attempt: number | null
  eventCount: number | null
  responseStatus: number | null
  latencyMs: number | null
  error: string | null
  response: unknown
  outcome: WebhookDeliveryOutcome
  firedAt: Generated<Timestamp>
  createdAt: Generated<Timestamp>
}

export interface Experiment {
  id: string
  name: string | null
  type: ExperimentType
  slug: string
  projectId: string
  workflowId: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  archivedAt: Timestamp | null
  workbenchState: unknown
  workbenchVersion: Generated<number>
}

export interface ExperimentVersion {
  id: string
  experimentId: string
  projectId: string
  /**
   * The number the version is addressed and displayed by. Numbered rows hold
   * a gap-free sequence. The rolling autosave row rides on the counter
   * instead, which is what keeps it clear of that sequence.
   */
  version: number
  /**
   * `Experiment.workbenchVersion` at the moment this row was last written.
   * The rolling row is rewritten in place, so this is the only column that
   * says how recent its content is, and the version list sorts and pages by
   * it. It also joins a row to the live state: the row whose `counterVersion`
   * equals the experiment's counter holds what the workbench shows now.
   */
  counterVersion: number
  autoSaved: Generated<boolean>
  commitMessage: string | null
  authorId: string | null
  /** Who wrote this version: "user", "langy" or "api". */
  authorLabel: string
  /**
   * The run that wrote this version, when a run wrote it.
   *
   * A run writes its cells into the workbench state, which advances the
   * counter. Without this, the page that started that run reads the bump as
   * somebody else's write, refuses to save, and asks the reader to reload
   * over their own unsaved edits. Named here, the page adopts the version and
   * carries on.
   */
  runId: string | null
  state: unknown
  /** Schema version of `state`, so a later shape change can be read forward. */
  schemaVersion: Generated<string>
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface Annotation {
  id: string
  projectId: string
  traceId: string
  comment: string | null
  isThumbsUp: boolean | null
  userId: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  email: string | null
  scoreOptions: unknown
  expectedOutput: string | null
  anchorKind: string | null
  anchorId: string | null
  anchorPath: string | null
}

export interface ModelProvider {
  id: string
  name: string
  provider: string
  routingHandle: string | null
  enabled: boolean
  customKeys: unknown
  extraHeaders: unknown
  customModels: unknown
  customEmbeddingsModels: unknown
  deploymentMapping: unknown
  rateLimitRpm: number | null
  rateLimitTpm: number | null
  rateLimitRpd: number | null
  rotationPolicy: Generated<GatewayProviderRotationPolicy>
  providerConfig: unknown
  fallbackPriorityGlobal: number | null
  langySkipPermissionsModels: unknown
  healthStatus: Generated<GatewayProviderHealthStatus>
  circuitOpenedAt: Timestamp | null
  lastHealthCheckAt: Timestamp | null
  disabledAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  organizationId: string
}

export interface ModelProviderScope {
  id: string
  modelProviderId: string
  scopeType: ModelProviderScopeType
  scopeId: string
  createdAt: Generated<Timestamp>
}

export interface ModelDefaultConfig {
  id: string
  config: unknown
  authorId: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  organizationId: string
}

export interface ModelDefaultConfigScope {
  id: string
  configId: string
  scopeType: ModelDefaultScopeType
  scopeId: string
  createdAt: Generated<Timestamp>
}

/**
 * The organization's GitHub connection: a GitHub App installation mapped to a
 * LangWatch organization. Short-lived (1h) installation access tokens are
 * minted against these rows on demand, by Langy for writes (bot-authored pull
 * requests, scoped to the installation's repositories) and by pull-request
 * linkage for reads. No user token and no secret are stored here, the crown
 * jewel is the App private key, held only in the control-plane env.
 * Spec: specs/integrations/github-connection.feature.
 *
 * installationId is globally unique (one GitHub installation, one row); a
 * single LangWatch organization may own MANY installations (different GitHub
 * accounts), so the org to installation relationship is one-to-many.
 */
export interface GithubInstallation {
  id: string
  /** GitHub installation id (numeric; stored as string). Globally unique. */
  installationId: string
  organizationId: string
  /** The GitHub account (organization or user) the App is installed on. */
  accountLogin: string
  /** "Organization" or "User". */
  accountType: string
  /** GitHub numeric account id (stable across renames). */
  accountId: string
  /**
   * "all" or "selected" — whether the App can see every repo on the account
   * or only a chosen subset.
   */
  repositorySelection: string
  /**
   * Cached repository list [{ id, fullName }] when repositorySelection =
   * "selected". Null when "all" (the full set is resolved live at mint time).
   */
  repositories: unknown
  /**
   * Set when GitHub suspends the installation; cleared on unsuspend. A
   * suspended installation cannot mint tokens.
   */
  suspendedAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/**
 * A pull request discovered through the organization's GitHub connection and
 * mapped to the branch it was opened from, so a coding-agent session that
 * knows only its branch can name the pull request it produced.
 * Spec: specs/integrations/github-connection.feature.
 *
 * The stored title, state, draft and timestamp columns are SNAPSHOTS taken
 * when the mapping was last refreshed. They are fallback labels only: the
 * frontend reads live status from GitHub, so a stale row never presents
 * itself as current truth.
 */
export interface GithubPullRequest {
  id: string
  organizationId: string
  /** The GitHub host the repository lives on, e.g. "github.com". */
  repositoryHost: string
  /**
   * "owner/repo", stored LOWERCASED so a lookup matches whatever casing the
   * session reported.
   */
  repositoryFullName: string
  /** The branch the pull request was opened from. */
  headBranch: string
  prNumber: number
  htmlUrl: string
  title: string
  /** GitHub's own state string: "open" or "closed". */
  state: string
  isDraft: Generated<boolean>
  authorLogin: string | null
  prCreatedAt: Timestamp
  prClosedAt: Timestamp | null
  prMergedAt: Timestamp | null
  /**
   * GitHub's own `updated_at` for the snapshot this row holds, and the
   * freshness guard on every write to it: a snapshot older than the stored one
   * is skipped. GitHub may deliver webhooks out of order, and without the
   * guard a late `opened` lands after `closed` and reopens a merged pull
   * request. NULL means the row predates the column, so any write is accepted.
   */
  prUpdatedAt: Timestamp | null
  /** When this row was first mapped to its branch. */
  mappedAt: Generated<Timestamp>
  /** When the snapshot columns were last refreshed from GitHub. */
  lastCheckedAt: Generated<Timestamp>
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/**
 * The per-branch bookkeeping behind pull-request mapping: when a branch was
 * last looked up on GitHub, how many pull requests it resolved to, and when it
 * may be asked about again. A branch with no pull request is the common case,
 * so `notFoundAt` plus `recheckAfter` is what keeps an unmapped branch from
 * costing a GitHub call on every read.
 * Spec: specs/integrations/github-connection.feature.
 */
export interface GithubBranchPullRequestCheck {
  id: string
  organizationId: string
  repositoryHost: string
  /** "owner/repo", stored LOWERCASED, matching GithubPullRequest. */
  repositoryFullName: string
  headBranch: string
  lastCheckedAt: Timestamp
  prCount: Generated<number>
  /**
   * Set while the branch resolves to no pull request; cleared once one is
   * found.
   */
  notFoundAt: Timestamp | null
  /** The earliest instant GitHub may be asked about this branch again. */
  recheckAfter: Timestamp | null
  /** Consecutive lookups that found nothing, backing off the recheck. */
  attempts: Generated<number>
  /**
   * When a reader last asked about this branch, whether or not GitHub was
   * called.
   */
  lastRequestedAt: Timestamp
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/**
 * Rebuildable operational state folded from Langy's canonical event stream.
 * The source of truth remains the ClickHouse event log.
 */
export interface LangyConversationProjection {
  id: string
  projectId: string
  conversationId: string
  /** Owner; first-writer-wins from the first message's userId. */
  userId: string
  title: string | null
  /** derived | auto | user — governs auto-title precedence (user is sticky). */
  titleSource: LangyProjectionTitleSource
  /**
   * Who started the conversation: `interactive` for the person in the panel, `run` for a
   * module acting for them with nobody at the keyboard. A row folded before the column
   * existed reads as interactive.
   */
  origin: Generated<string>
  status: string
  isShared: Generated<boolean>
  sharedAt: number | null
  sharedById: string | null
  messageCount: Generated<number>
  lastActivityAt: number | null
  /** The turn in flight, or null when idle (busy guard reads this). */
  currentTurnId: string | null
  lastError: string | null
  /** The model the latest accepted turn ran on; reopening seeds the picker from it. */
  lastModel: string | null
  /** ADR-048 shutdown-handoff resume token (opaque, worker-authored). */
  pendingHandoffToken: string | null
  pendingHandoffTurnId: string | null
  /** Per-conversation HMAC key for worker stream frames; server-side only. */
  runToken: string | null
  archivedAt: number | null
  /** Messages CI started on a pull request since the person last wrote. */
  ciFollowUpsInARow: Generated<number>
  createdAt: number
  updatedAt: number
  /** Business time of the latest applied event. */
  occurredAt: number
  /** Event-log acceptance time. Together with lastEventId this is the cursor. */
  acceptedAt: number
  lastEventId: string
  projectionVersion: string
}

/**
 * Durable idempotency receipt for one logical Langy send. A transport retry
 * reuses the same requestId and receives the original conversation/turn ids
 * instead of minting a second message or turn.
 */
export interface LangyTurnRequest {
  id: string
  projectId: string
  userId: string
  requestId: string
  conversationId: string
  turnId: string
  status: string
  leaseOwner: string
  leaseExpiresAt: Timestamp
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/**
 * Postgres-authoritative one-turn admission for a conversation. PREPARING
 * claims are leased so a crash before the durable start event can recover;
 * COMMITTED claims are released only by the matching terminal event.
 */
export interface LangyActiveTurn {
  id: string
  projectId: string
  conversationId: string
  turnId: string
  requestId: string
  userId: string
  status: string
  leaseOwner: string
  leaseExpiresAt: Timestamp
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/** Rebuildable per-turn operational render document. */
export interface LangyConversationTurnProjection {
  id: string
  projectId: string
  conversationId: string
  turnId: string
  /**
   * LangyConversationTurnStatus. Deliberately TEXT, not a Postgres enum: this
   * is a PROJECTION — one writer (the fold), rebuildable from the event log,
   * and nothing queries or indexes on it. An enum bought no selectivity here,
   * only typo protection, and charged a migration for every new status that
   * Postgres then could not remove. The one definition lives in
   * `LANGY_CONVERSATION_TURN_STATUS`; the repository validates against it.
   */
  status: string
  /** LangyMessagePart[] — the user question that opened the turn. */
  questionParts: unknown
  /** LangyMessagePart[] — the agent's whole final answer. */
  answerParts: unknown
  /** LangyTurnToolCall[] — tool-call lifecycle audit in initiation order. */
  toolCalls: unknown
  /** LangyPlanItemData[] | null — null means the turn never kept a plan. */
  plan: unknown
  error: string | null
  startedAt: number | null
  endedAt: number | null
  createdAt: number
  updatedAt: number
  occurredAt: number
  acceptedAt: number
  lastEventId: string
  projectionVersion: string
}

/**
 * One pull request Langy opened, from the answer whose `gh pr create` printed it. Finished
 * checks look it up by repository and number to wake the conversation; the daily pull request
 * cap counts a person's rows by day. Rebuilt by a projection replay.
 */
export interface LangyPullRequest {
  id: string
  projectId: string
  /** `owner/repo`, lowercased: GitHub matches repository names without case. */
  repositoryFullName: string
  prNumber: number
  conversationId: string
  /** The conversation's owner, who opened it through Langy. */
  userId: string
  /** When the answer that opened it was recorded (epoch ms). */
  openedAt: number
  /** The agent_responded event it came from. */
  sourceEventId: string
}

/** Rebuildable operational message fact. One source event produces one row. */
export interface LangyMessageProjection {
  id: string
  projectId: string
  conversationId: string
  messageId: string
  role: string
  parts: unknown
  sourceEventId: string
  occurredAt: number
  acceptedAt: number
  createdAt: number
  updatedAt: number
}

export interface ProjectSecret {
  id: string
  projectId: string
  name: string
  encryptedValue: string
  createdById: string
  updatedById: string
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface TriggerSent {
  id: string
  triggerId: string
  traceId: string | null
  customGraphId: string | null
  projectId: string
  resolvedAt: Timestamp | null
  openIncidentKey: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface TriggerLatestEvaluation {
  triggerId: string
  projectId: string
  evaluatedAt: Timestamp
  verdict: string
  observedValue: number | null
  threshold: number | null
  operator: string | null
  timePeriodMinutes: number | null
  skipCode: string | null
  updatedAt: Timestamp
}

export interface EmailSuppression {
  id: string
  projectId: string
  email: string
  triggerId: string | null
  reason: Generated<string>
  createdAt: Generated<Timestamp>
}

export interface AnnotationScore {
  id: string
  name: string
  projectId: string
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  deletedAt: Timestamp | null
  description: string | null
  active: Generated<boolean>
  dataType: AnnotationScoreDataType | null
  options: unknown
  defaultValue: unknown
  global: Generated<boolean>
}

export interface AnnotationQueue {
  id: string
  name: string
  slug: string
  projectId: string
  description: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface AnnotationQueueMembers {
  annotationQueueId: string
  userId: string
}

export interface AnnotationQueueScores {
  annotationQueueId: string
  annotationScoreId: string
}

export interface AnnotationQueueItem {
  id: string
  annotationQueueId: string | null
  userId: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  createdByUserId: string | null
  traceId: string
  projectId: string
  doneAt: Timestamp | null
  /**
   * Set while the annotator wants this item's trace added to a dataset once
   * the queue walk is finished. Persisted rather than kept in the browser so
   * the selection survives a refresh and stays visible to teammates.
   */
  markedForDatasetAt: Timestamp | null
}

export interface ShareLink {
  id: string
  token: string
  resourceType: ShareResourceType
  resourceId: string
  threadId: string | null
  projectId: string
  userId: string | null
  visibility: Generated<ShareVisibility>
  expiresAt: Timestamp | null
  maxViews: number | null
  viewCount: Generated<number>
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface PinnedTrace {
  id: string
  projectId: string
  traceId: string
  userId: string | null
  source: Generated<PinSource>
  reason: string | null
  createdAt: Generated<Timestamp>
}

export interface TraceEditOverlay {
  id: string
  projectId: string
  traceId: string
  patch: unknown
  createdById: string | null
  updatedById: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface TraceIngestSourceBilling {
  organizationId: string
  sourceType: string
  billed: boolean
  recordedAt: Timestamp
  updatedAt: Timestamp
}

export interface RetentionPolicy {
  id: string
  organizationId: string
  scopeType: RetentionPolicyScopeType
  scopeId: string
  category: string
  retentionDays: number
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface RetentionApplyStatus {
  id: string
  organizationId: string
  scopeType: RetentionPolicyScopeType
  scopeId: string
  changedAt: number | null
  appliedAt: number | null
  raisedAt: number | null
  shortenedAt: number | null
  createdAt: number
  updatedAt: number
  occurredAt: number
  acceptedAt: number
  lastEventId: string
  projectionVersion: string
}

export interface DataPrivacyPolicy {
  id: string
  organizationId: string
  scopeType: DataPrivacyScopeType
  scopeId: string
  personalOnly: Generated<boolean>
  config: unknown
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface InstantEvalJudgeProject {
  projectId: string
  organizationId: string
  createdAt: Timestamp
}

export interface InstantEvalJudgeUsageBilling {
  organizationId: string
  usageBilled: boolean
  occurredAt: Timestamp
  fromCatchUp: boolean
}

export interface InstantEvalJudgeSpend {
  organizationId: string
  requestId: string
  spendNanoUsd: ColumnType<string, bigint | number | string, bigint | number | string>
  occurredAt: Timestamp
}

export interface CustomLLMModelCost {
  id: string
  organizationId: string
  scopeType: Generated<CustomLLMModelCostScopeType>
  scopeId: string
  projectId: string | null
  model: string
  regex: string
  inputCostPerToken: number | null
  outputCostPerToken: number | null
  cacheReadCostPerToken: number | null
  cacheCreationCostPerToken: number | null
  cacheCreation1hCostPerToken: number | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface Workflow {
  id: string
  projectId: string
  name: string
  icon: string
  description: string
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  latestVersionId: string | null
  currentVersionId: string | null
  publishedId: string | null
  publishedById: string | null
  archivedAt: Timestamp | null
  isEvaluator: Generated<boolean>
  isComponent: Generated<boolean>
  copiedFromWorkflowId: string | null
}

export interface WorkflowVersion {
  id: string
  version: string
  commitMessage: string
  authorId: string
  projectId: string
  workflowId: string
  parentId: string | null
  autoSaved: Generated<boolean>
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  dsl: unknown
}

export interface AuditLog {
  id: string
  createdAt: Generated<Timestamp>
  userId: string | null
  projectId: string | null
  organizationId: string | null
  action: string
  args: unknown
  error: string | null
  ipAddress: string | null
  userAgent: string | null
  metadata: unknown
  /**
   * WHO REALLY DID IT, when that is not `userId`.
   *
   * Under an impersonation the session's `user.id` is the SUBJECT — the
   * customer whose access an operator is borrowing — so a row naming only
   * `userId` filed the operator's act against the customer. An operator
   * granting a permanent way into somebody's organization appeared, in the
   * only durable record of it, to be that somebody.
   *
   * Null on an ordinary request, where the two would be the same person and
   * saying so twice adds nothing. ADR-118 / D06: every decision names both.
   */
  actorUserId: string | null
  targetKind: string | null
  targetId: string | null
  before: unknown
  after: unknown
  /** A producer's outbox key: a redelivered audit intent finds its row here (Alex, audit R2). */
  idempotencyKey: string | null
}

export interface LlmPromptConfig {
  id: string
  handle: string | null
  name: string
  projectId: string
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  deletedAt: Timestamp | null
  organizationId: string
  scope: Generated<PromptScope>
  copiedFromPromptId: string | null
}

export interface LlmPromptConfigVersion {
  id: string
  version: number
  commitMessage: string | null
  authorId: string | null
  configId: string
  config: unknown
  schemaVersion: string
  runtimeParameters: Generated<unknown>
  createdAt: Generated<Timestamp>
  projectId: string
}

export interface PromptTagAssignment {
  id: string
  configId: string
  versionId: string
  tagId: string
  projectId: string
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  createdById: string | null
  updatedById: string | null
}

export interface PromptTag {
  id: string
  organizationId: string
  name: string
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  createdById: string | null
  updatedById: string | null
}

export interface Analytics {
  id: string
  projectId: string
  key: string
  value: unknown
  numericValue: number | null
  stringValue: string | null
  boolValue: boolean | null
  createdAt: Generated<Timestamp>
}

export interface CustomRole {
  id: string
  organizationId: string
  name: string
  description: string | null
  permissions: unknown
  kind: Generated<string>
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface Notification {
  id: string
  organizationId: string | null
  projectId: string | null
  metadata: unknown
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  sentAt: Timestamp
}

export interface Agent {
  id: string
  projectId: string
  name: string
  type: string
  config: unknown
  workflowId: string | null
  copiedFromAgentId: string | null
  environment: string | null
  ownerUserId: string | null
  hostLabel: string | null
  identityKey: string | null
  lastSeenAt: Timestamp | null
  archivedAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  createdById: string | null
  updatedById: string | null
}

/** One immutable snapshot of an agent; the Agent row is always the latest. */
export interface AgentVersion {
  id: string
  projectId: string
  agentId: string
  version: number
  name: string
  type: string
  config: unknown
  workflowId: string | null
  commitMessage: string | null
  authorId: string | null
  /** "user", "api_key" or "system": who wrote the version when authorId is null. */
  authorKind: string | null
  createdAt: Generated<Timestamp>
}

/** A deploy label ("production", "staging") pointing at one version of an agent. */
export interface AgentVersionLabel {
  id: string
  projectId: string
  agentId: string
  label: string
  versionId: string
  createdById: string | null
  updatedById: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface Evaluator {
  id: string
  projectId: string
  name: string
  slug: string | null
  type: string
  config: unknown
  workflowId: string | null
  copiedFromEvaluatorId: string | null
  archivedAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  createdById: string | null
  updatedById: string | null
}

/** One immutable snapshot of an evaluator; the Evaluator row is always the latest. */
export interface EvaluatorVersion {
  id: string
  projectId: string
  evaluatorId: string
  version: number
  name: string
  type: string
  config: unknown
  workflowId: string | null
  commitMessage: string | null
  authorId: string | null
  /** "user", "api_key" or "system": who wrote the version when authorId is null. */
  authorKind: string | null
  createdAt: Generated<Timestamp>
}

/** A deploy label ("production", "staging") pointing at one version of an evaluator. */
export interface EvaluatorVersionLabel {
  id: string
  projectId: string
  evaluatorId: string
  label: string
  versionId: string
  createdById: string | null
  updatedById: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface Scenario {
  id: string
  projectId: string
  name: string
  situation: string
  criteria: string[]
  labels: string[]
  parameters: unknown
  simulatorModel: string | null
  judgeModel: string | null
  traceQuietPeriodMs: number | null
  traceWaitTimeoutMs: number | null
  maxTurns: number | null
  minTurns: number | null
  fields: unknown
  callerVoice: unknown
  testSuiteId: string | null
  kind: Generated<string>
  redTeam: unknown
  sourceTemplateId: string | null
  version: Generated<number>
  lastUpdatedById: string | null
  archivedAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface ScenarioVersion {
  id: string
  scenarioId: string
  /**
   * The project of the parent case, carried here so every read filters by
   * tenant without a join. The service is what keeps the pair honest: no
   * composite foreign key can, under relationMode="prisma".
   */
  projectId: string
  version: number
  authorId: string | null
  /** Who wrote this version: "user", "api", "cli" or "langy". */
  authorLabel: string
  changeDescription: string | null
  /**
   * Envelope { schemaVersion, fields, changedFields } holding the editable
   * field set as saved and the names of the fields this save changed.
   */
  snapshot: unknown
  /** Schema version of `snapshot`, so a later shape change can be read forward. */
  schemaVersion: Generated<number>
  createdAt: Generated<Timestamp>
}

export interface SimulationSuite {
  id: string
  projectId: string
  name: string
  slug: string
  description: string | null
  scenarioIds: string[]
  targets: unknown
  repeatCount: Generated<number>
  labels: string[]
  simulatorModel: string | null
  judgeModel: string | null
  traceQuietPeriodMs: number | null
  traceWaitTimeoutMs: number | null
  kind: Generated<string>
  scope: unknown
  fields: unknown
  evaluators: unknown
  templateId: string | null
  templateVersion: number | null
  archivedAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface Subscription {
  id: string
  organizationId: string
  plan: PlanTypes
  stripeSubscriptionId: string | null
  status: SubscriptionStatus
  startDate: Timestamp | null
  endDate: Timestamp | null
  lastPaymentFailedDate: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  maxMembers: number | null
  maxProjects: number | null
  maxMessagesPerMonth: number | null
  maxWorkflows: number | null
  maxRetentionDays: number | null
  maxMembersLite: number | null
  /** Developer seats the organization pays for; null falls to the plan's own cap. */
  maxMembersDeveloper: number | null
  maxTeams: number | null
  maxPrompts: number | null
  maxEvaluators: number | null
  maxScenarios: number | null
  maxAgents: number | null
  maxExperiments: number | null
  maxOnlineEvaluations: number | null
  maxDatasets: number | null
  maxDashboards: number | null
  maxCustomGraphs: number | null
  maxAutomations: number | null
  /** Seats paid for this period that stay usable after a seat decrease, until seatsKeptUntil. */
  seatsKept: number | null
  seatsKeptUntil: Timestamp | null
  /** The seat quantity the last finished seat change set the seat line to; provider events never write it. */
  seatLineQuantity: number | null
  seatLineQuantityAt: Timestamp | null
  /**
   * A Stripe seat line write recorded before it is sent and cleared once it is known to have
   * finished under the lease; one left behind is settled by the seat line repair.
   * Every seat quantity written while this record was pending.
   */
  seatLineWriteQuantities: Generated<number[]>
  /** The first write still pending; a later holder's write keeps it. */
  seatLineWriteSince: Timestamp | null
  /** The latest write sent; the repair waits a grace window after it. */
  seatLineWriteAt: Timestamp | null
  seatLineWriteAttempts: number | null
  seatLineWriteFlaggedAt: Timestamp | null
  /** The seat quantity Stripe should bill if a seat change stops after lifting the seat line; null once settled. */
  seatRepairQuantity: number | null
  /** The quantity the change lifted the line to; the repair moves the line back only while Stripe still bills it. */
  seatRepairLiftedQuantity: number | null
  seatRepairSince: Timestamp | null
  seatRepairAttempts: number | null
  /** Set once the seat line repair has failed often enough that Ops should look at it. */
  seatRepairFlaggedAt: Timestamp | null
  /** The lease a seat change or seat line repair holds while it reads and moves the seat line. */
  seatLineLeaseToken: string | null
  seatLineLeaseUntil: Timestamp | null
  /**
   * An open failed payment recovery case: the unpaid renewal invoice, day 0 (its first failed
   * charge) and the last notice sent (3, 7, 14 or 21). Cleared when the case closes.
   */
  paymentRecoveryInvoiceId: string | null
  paymentRecoveryStartedAt: Timestamp | null
  paymentRecoveryNoticeDay: number | null
  /**
   * The over-plan mark: traces recorded from overPlanSince are withheld while it is open and the
   * organization is on Free with more members than Free allows; ended once that no longer holds.
   */
  overPlanSince: Timestamp | null
  overPlanEndedAt: Timestamp | null
}

export interface Invoice {
  id: string
  subscriptionId: string
  totalAmount: number
  currency: string
  dueDate: Timestamp
  paid: Generated<boolean>
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface InvoiceItem {
  id: string
  invoiceId: string
  description: string
  amount: number
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface Group {
  id: string
  organizationId: string
  name: string
  slug: string
  externalId: string | null
  scimSource: string | null
  /**
   * D08: which SSO connection provisioned this group. Null for manually
   * created groups and for the SCIM groups that predate connection scoping.
   */
  scimConnectionId: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface GroupMembership {
  userId: string
  groupId: string
  createdAt: Generated<Timestamp>
}

/** Who authz treats as gone, from user's and identity's facts; authz never reads User (ARCHITECTURE.md, "Platform operators are a grant"). */
export interface AuthzUserStanding {
  userId: string
  deactivatedAt: Timestamp | null
  erasedAt: Timestamp | null
  standingChangedAt: Timestamp
}

/**
 * Whether an organization's permission checks run on the authorization engine: authz's own
 * switch, set when the engine migration proves the organization and pinned off by an operator's
 * rollback. No row is the legacy path. Spec: modules/authz/specs/engine-switch.feature.
 */
export interface AuthzEngineSwitch {
  organizationId: string
  onEngine: boolean
  pinned: Generated<boolean>
  switchedAt: Timestamp
  lastAppliedEventId: string | null
  updatedAt: Generated<Timestamp>
}

/** ADR-201: an organisation's braid setting. Owned by authz; read into the organisation's snapshot. */
export interface AuthzOrganizationPolicy {
  organizationId: string
  braidsOrganizationAdminsOnly: Generated<boolean>
  braidsEnabled: Generated<boolean>
  updatedByUserId: string
  updatedAt: Timestamp
}

export interface AuthzBraidWrite {
  organizationId: string
  commandId: string
  grantId: string
  readerProjectId: string
  ownerProjectId: string
  termsKey: string
  condition: unknown
  expiresAt: Timestamp | null
  replacesGrantId: string | null
  createdByUserId: string | null
  createdAt: Generated<Timestamp>
}

export interface RoleBinding {
  id: string
  organizationId: string
  userId: string | null
  groupId: string | null
  apiKeyId: string | null
  role: TeamUserRole
  customRoleId: string | null
  scopeType: RoleBindingScopeType
  scopeId: string
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface Grant {
  id: string
  organizationId: string
  principalType: GrantPrincipalType
  principalId: string | null
  roleKey: string | null
  legacyRole: string | null
  source: string
  scopeType: GrantScopeType
  scopeId: string
  token: string | null
  permission: string | null
  resourceKind: string | null
  projectId: string | null
  createdByUserId: string | null
  expiresAt: Timestamp | null
  maxViews: number | null
  condition: unknown
  replacesGrantId: string | null
  occurredAt: Timestamp
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  revokedAt: Timestamp | null
  revokedReason: string | null
}

export interface GrantUsage {
  grantId: string
  organizationId: string
  projectId: string
  viewCount: Generated<number>
  lastViewedAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface Role {
  id: string
  organizationId: string
  name: string
  description: string | null
  permissions: unknown
  kind: Generated<string>
  occurredAt: Timestamp
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  deletedAt: Timestamp | null
}

export interface ApiKey {
  id: string
  name: string
  description: string | null
  lookupId: string
  hashedSecret: string
  permissionMode: Generated<string>
  userId: string | null
  createdByUserId: string | null
  organizationId: string
  expiresAt: Timestamp | null
  lastUsedAt: Timestamp | null
  revokedAt: Timestamp | null
  /**
   * Why the key was revoked, set with `revokedAt`: "user" when a person
   * revoked it (the API-keys page, the REST API, the devices tab), "rotation"
   * when a mint replaced it, "session" when the CLI login key it was parented
   * to was revoked, "expired" when that session ran out, "offboarded" when the
   * person's membership of the organization ended. "cap" survives on rows the
   * retired per-tool cap revoked. The CLI reads it to decide whether a dead
   * ingest key may be re-minted: a key a person revoked stays dead.
   * NULL for live keys and for keys revoked before the cause was recorded.
   */
  revocationCause: string | null
  /**
   * The CLI login key of the device session that minted this ingestion key.
   * Set by the CLI's personal mint, NULL for every other key. Revoking the
   * parent (logout, the devices tab, a re-login from the same device, session
   * expiry) revokes the children with it. Column ref only, no FK and no
   * index: lookups filter by the indexed `userId` first and match the parent
   * in memory over a person's few live ingest keys.
   */
  parentApiKeyId: string | null
  /**
   * Non-null marks this key as an "ingestion key": a project-scoped,
   * ingest-only ApiKey the unified `langwatch <tool>` CLI (and template
   * installs) mint for OTLP Path B. The slug (claude_code / codex / gemini /
   * opencode / claude_cowork) is stamped as `langwatch.source` provenance on
   * every span the key authorizes. NULL for normal API / service keys.
   */
  ingestSourceType: string | null
  /**
   * Set only for ingestion keys derived from an admin IngestionTemplate
   * (e.g. claude_cowork). Drives receiver OTTL application + the
   * `langwatch.template.id` provenance stamp. NULL for unified-CLI tools
   * (which have no template) and for normal keys. Column ref only (no FK,
   * per the app-layer-integrity convention); resolved by the ingest-key service.
   */
  ingestionTemplateId: string | null
  /**
   * Human label of the CLI device session that minted this ingestion key
   * (e.g. "Rogerio's MacBook Pro"), snapshotted at mint time from the device
   * flow's `client_info`. Denormalized so the org-wide API-keys settings page
   * can show which device a key came from without resolving another user's
   * Redis-stored sessions. NULL for non-ingestion keys and for keys minted by
   * a CLI old enough to predate device metadata.
   */
  createdByDeviceLabel: string | null
  /**
   * True for a key LangWatch minted for itself (a run, the gateway's trace
   * export, a Langy session). Hides it from listings, refuses customer edits
   * and revokes, and lets an ownerless run key act as the system. Never set
   * by a customer create; rows before the column stay false.
   */
  isSystemManaged: Generated<boolean>
  /**
   * The browser session that minted a docs session key (ADR-185), so signing
   * out of that session revokes it. NULL for every other key. Unindexed:
   * sign-out filters by the indexed `userId` first.
   */
  createdByBrowserSessionId: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/** ADR-201: a project's key for signing the tags on its calls. Owned by the API key module. */
export interface SigningKey {
  id: string
  organizationId: string
  projectId: string
  name: string
  publicId: string
  secretCiphertext: string
  createdByUserId: string
  createdAt: Generated<Timestamp>
  rotatedAt: Timestamp | null
  verifiesUntil: Timestamp | null
  revokedAt: Timestamp | null
  lastUsedAt: Timestamp | null
}

export interface ProjectActiveDay {
  id: string
  projectId: string
  day: string
  createdAt: Generated<Timestamp>
}

/**
 * Nurturing's read model: which organization a project belongs to, from project's created event.
 * The per-email usage places a project and lists an organization's projects from these rows only.
 */
export interface NurturingProject {
  projectId: string
  organizationId: string
}

/**
 * retired: unread since v3.28.0. The trace-day folds count a trace with no placement copy, so no
 * retry clock reads or writes this table; an operator step may drop it once old writers are gone.
 */
export interface NurturingUncopiedProject {
  projectId: string
  firstSeenAt: Generated<Timestamp>
}

/**
 * Nurturing's own copy of each organization's members, from the owners' member facts (§9):
 * organization's signed_up and invite_accepted and auth's sso_auto_added join, member_removed
 * and member_disabled leave, member_enabled returns. Each column keeps its latest time.
 */
export interface NurturingOrganizationMember {
  organizationId: string
  userId: string
  joinedAt: Timestamp | null
  leftAt: Timestamp | null
  returnedAt: Timestamp | null
}

/**
 * When an organization was created, from organization's created event: nurturing's own copy for
 * the per-email usage's day offsets. Its own table, so an early creation never adds an unseeded
 * NurturingOrganization row.
 */
export interface NurturingOrganizationCreated {
  organizationId: string
  createdAt: Timestamp
}

/**
 * Each organization's plan, folded from billing's subscription_changed facts: nurturing's own
 * copy. The newest fact wins, and a real fact wins a tie over a catch-up; Customer.io hears a
 * fact only when it changes this row.
 */
export interface NurturingOrganizationPlan {
  organizationId: string
  plan: string | null
  occurredAt: Timestamp
  fromCatchUp: boolean
  /** The event ids folded in, so a redelivery folds once. */
  appliedEventIds: Generated<string[]>
  /** The fact that last changed the row; only it may tell Customer.io. */
  winnerEventId: string | null
}

/**
 * Each gateway virtual key, folded from gateway's vk_lifecycle facts: nurturing's own copy. The
 * creator comes from the created fact alone; the holder and purpose are each the latest non-null
 * value by the fact's time (stamped beside them); the state times keep their latest.
 */
export interface NurturingVirtualKey {
  virtualKeyId: string
  organizationId: string
  name: string
  occurredAt: Timestamp
  keyCreatedAt: Timestamp | null
  creatorUserId: string | null
  principalUserId: string | null
  principalUserIdAt: Timestamp | null
  purpose: string | null
  purposeAt: Timestamp | null
  disabledAt: Timestamp | null
  enabledAt: Timestamp | null
  revokedAt: Timestamp | null
}

/**
 * Each account nurturing erased, as a SHA-256 of its id and never the id: a gateway key fact
 * naming it, however late it lands, stores no creator or holder. Platform-wide, like the erasure.
 */
export interface NurturingErasedUser {
  userIdHash: string
  erasedAt: Generated<Timestamp>
}

/**
 * Nurturing's read model per organization: its admin and its milestone counts.
 * `seeded` marks an organization learned from the project backfill: its first_* already happened.
 */
export interface NurturingOrganization {
  organizationId: string
  adminUserId: string | null
  seeded: Generated<boolean>
  evaluationCount: Generated<number>
  simulationRunCount: Generated<number>
  codingFirstSessionAt: Timestamp | null
  codingAgentTypes: Generated<string[]>
  codingActiveDays: Generated<number>
  codingLastActiveDay: string | null
  codingTerminalTypes: Generated<string[]>
  codingRepositoryKeys: Generated<string[]>
  codingFirstDays: Generated<number[]>
  lastTraceAt: Timestamp | null
  adminAt: Timestamp | null
  updatedAt: Timestamp
}

/**
 * Nurturing's own per-project usage: a peer fold over the owners' prompt, evaluator, scenario
 * and simulation-run events (ARCHITECTURE §9), one peer fold per owner pipeline, summed per
 * organization for the per-email traits. Traces are counted in NurturingProjectTraceDayCount.
 */
export interface NurturingProjectUsageView {
  projectId: string
  /**
   * The owner pipeline whose peer fold wrote the row (prompt, evaluator, scenario, simulation).
   * Evaluation days are not kept here: the evaluation-days map and its replay write them to
   * NurturingProjectEvaluationDay.
   */
  owner: string
  /** The project's usage as the fold folded it (NurturingProjectUsageState). */
  folded: unknown
  /** The event ids of the last delivery folded in, so a redelivery folds once. */
  appliedEventIds: Generated<string[]>
  /** A row of another version reads as undecodable and is re-folded from the log. */
  projectionVersion: string
  /** Event-derived, so a replayed row equals the row it rebuilds. */
  updatedAt: Timestamp
}

/**
 * retired: unread since v3.28.0. Replaced by NurturingProjectTraceDayCount, which the trace-day
 * peer folds write and a replay rebuilds; an operator step may drop it once old writers are gone.
 */
export interface NurturingProjectTraceDay {
  projectId: string
  /** UTC `YYYY-MM-DD` of the root spans' append time. */
  day: string
  traceCount: Generated<number>
  /** Traces trace recorded failed (lw.trace.trace_failed). */
  errorCount: Generated<number>
  /** Append time of the newest root span. */
  lastTraceAt: Timestamp | null
  /** Append time of the newest trace_failed. */
  lastFailedTraceAt: Timestamp | null
}

/**
 * Nurturing's own trace counts (ARCHITECTURE §9), one row per project, UTC day of append time and
 * source: `spans` from trace's span_received peer fold (a customer root span is one trace),
 * `failures` from trace_failed's. Each fold owns its rows, so the two never overwrite each other;
 * readers sum a day's sources. Rebuilt by a projection replay. Kept 30 days.
 */
export interface NurturingProjectTraceDayCount {
  projectId: string
  /** UTC `YYYY-MM-DD` of the facts' append time. */
  day: string
  /** The peer fold that wrote the row: `spans` or `failures`. */
  source: string
  traceCount: Generated<number>
  errorCount: Generated<number>
  /** Append time of the newest root span. */
  lastTraceAt: Timestamp | null
  /** Append time of the newest trace_failed. */
  lastFailedTraceAt: Timestamp | null
  /** The event ids of the last delivery folded in, so a redelivery folds once. */
  appliedEventIds: Generated<string[]>
  /** A row of another version reads as undecodable and is left out of the read. */
  projectionVersion: string
}

/**
 * Each project and UTC day with a completed evaluation, from evaluation's lifecycle_completed,
 * written insert-ignore. Replaces the `evaluation` rows of NurturingProjectUsageView.
 */
export interface NurturingProjectEvaluationDay {
  projectId: string
  /** UTC `YYYY-MM-DD`. */
  day: string
}

/**
 * Nurturing's copy of each personal project's owner, from organization's and project's facts.
 * The first owner recorded stays: a personal project's owner never changes.
 */
export interface NurturingPersonalProject {
  projectId: string
  organizationId: string
  ownerUserId: string
}

/** Nurturing's copy of each project that received a real trace, from trace's and project's facts. */
export interface NurturingTracedProject {
  projectId: string
  organizationId: string
}

export interface BillingMeterCheckpoint {
  id: string
  organizationId: string
  billingMonth: string
  meter: Generated<string>
  lastReportedTotal: Generated<number>
  pendingReportedTotal: number | null
  consecutiveFailures: Generated<number>
  lastCountedEventId: string | null
  updatedAt: Timestamp
}

export interface VirtualKey {
  id: string
  organizationId: string
  name: string
  description: string | null
  status: Generated<VirtualKeyStatus>
  purpose: Generated<VirtualKeyPurpose>
  externalId: string | null
  metadata: Generated<unknown>
  disabledAt: Timestamp | null
  disabledReason: string | null
  expiresAt: Timestamp | null
  hashedSecret: string
  displayPrefix: string
  principalUserId: string | null
  traceProjectId: string | null
  config: Generated<unknown>
  revision: Generated<ColumnType<string, bigint | number | string, bigint | number | string>>
  previousHashedSecret: string | null
  previousSecretValidUntil: Timestamp | null
  revokedAt: Timestamp | null
  revokedById: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  createdById: string
  lastUsedAt: Timestamp | null
  routingPolicyId: string | null
  routingMode: Generated<VirtualKeyRoutingMode>
  connectServices: Generated<string[]>
  licenseTokenHash: string | null
  licenseInstanceId: string | null
  licenseExpiresAt: Timestamp | null
  licenseId: string | null
}

export interface VirtualKeyScope {
  id: string
  virtualKeyId: string
  scopeType: VirtualKeyScopeType
  scopeId: string
  createdAt: Generated<Timestamp>
}

export interface RoutingPolicy {
  id: string
  organizationId: string
  name: string
  description: string | null
  modelProviderIds: Generated<unknown>
  modelAliases: Generated<unknown>
  defaultModel: string | null
  policyRules: Generated<unknown>
  isDefault: Generated<boolean>
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  createdById: string | null
  updatedById: string | null
}

export interface RoutingPolicyScope {
  id: string
  routingPolicyId: string
  scopeType: RoutingPolicyScopeType
  scopeId: string
  createdAt: Generated<Timestamp>
}

export interface GatewayGuardrail {
  id: string
  projectId: string
  name: string
  description: string | null
  evaluatorId: string
  direction: GatewayGuardrailDirection
  failureMode: Generated<GatewayGuardrailFailureMode>
  archivedAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  createdById: string | null
  updatedById: string | null
}

/**
 * Per-platform fleet configuration that connects a closed SaaS
 * product's audit / OTel / S3 stream to LangWatch's Activity
 * Monitor (cf. specs/ai-gateway/governance/activity-monitor.feature
 * + ingestion-sources.feature, docs/ai-gateway/governance/architecture.md).
 *
 * One IngestionSource = one upstream fleet (e.g. "Miro Cowork",
 * "Acme Workato production"). The runtime ingests events on whatever
 * cadence the source supports - push for OTel/webhook/S3 (Tier C/D),
 * pull for poll-based admin APIs (Tier C). Per-source-type adapters
 * normalise platform-specific shapes into OCSF + AOS before they
 * land in trace_summaries with SourceType + SourceId tags.
 */
export interface IngestionSource {
  id: string
  organizationId: string
  /** Optional team-scope. null = org-wide. */
  teamId: string | null
  /**
   * 'otel_generic' | 'claude_cowork' | 'workato' | 'copilot_studio' |
   * 'openai_compliance' | 'claude_compliance' | 's3_custom' (extensible).
   * String not enum so adapters can be added without migrations - the
   * supported set lives in code (activity-monitor adapter registry).
   */
  sourceType: string
  /**
   * Human-friendly label for the admin UI. Unique per-org so two
   * "Cowork" sources can coexist with different names.
   */
  name: string
  description: string | null
  /**
   * HMAC bearer token the upstream operator pastes into the platform
   * (or LangWatch sends in pull mode). Stored encrypted at rest via
   * the same pepper-derived KEK as VirtualKey secrets. Auto-generated
   * on create, rotatable via the admin UI with a 24h grace window.
   */
  ingestSecretHash: string
  /**
   * Per-source-type configuration. Shape is platform-specific:
   * - otel_generic: { allowedSourceTypeLabel?: string }
   * - claude_cowork: { workspaceId: string }
   * - workato: { sharedSecretLastFour: string }
   * - copilot_studio: { tenantId, clientId, clientSecretHash, pollEverySec }
   * - openai_compliance: { bucket, prefix, roleArn, pollEverySec }
   * - claude_compliance: { workspaceApiKeyHash, pollEverySec }
   * - s3_custom: { bucket, prefix, roleArn, parserDsl, pollEverySec }
   */
  parserConfig: Generated<unknown>
  /**
   * For pull-based sources, persisted between polls so the next poll
   * resumes where the last left off (timestamp / cursor / file marker).
   */
  pollerCursor: unknown
  /**
   * Phase 10 - number of consecutive errors observed by the puller
   * worker. Reset to 0 on a successful run. Drives admin-UI banners
   * ("Microsoft authentication failed - re-authorize") + back-off
   * scheduling when the source has been broken for a while.
   * Spec: specs/ai-governance/puller-framework/puller-adapter-contract.feature
   */
  errorCount: Generated<number>
  /**
   * Phase 10 - cron schedule for pull-mode sources. Validated by the
   * adapter's pullConfig schema; mirrored here for fast worker
   * dispatch. Null for push-mode sources.
   */
  pullSchedule: string | null
  /**
   * Status flag. 'active' (default), 'disabled' (admin paused), or
   * 'awaiting_first_event' (created but never received).
   */
  status: Generated<string>
  /**
   * ADR-088 v7: optional trace DESTINATION for conversation-bearing pulls
   * (Databricks Genie). Null = don't route. Same shape as
   * VirtualKey.traceProjectId: a destination is NOT a scope row and grants
   * no access to the source; it only decides where routed conversation
   * traces land. Must be a project of this org (validated at write time).
   * Freely editable; already-routed history stays where it landed.
   * NOT the deferred cost-attribution projectId (ADR-088 Decision 4).
   */
  traceProjectId: string | null
  /**
   * The account the provider itself named when this connection was saved -
   * an OpenAI or Anthropic organization id, read once per save and compared
   * so a second connection cannot read one account's spend twice. NOT
   * derived from the administrator key: nothing worked out from a customer
   * credential is stored here or anywhere else, because the account id is
   * not a secret and cannot be turned back into a key.
   * Nullable because every connection that predates the guard has none, and
   * because Azure and Power Platform name what they read in parserConfig
   * instead. "Unknown account" has to be a storable state.
   * Spec: specs/ai-gateway/governance/ingestion-sources.feature
   */
  providerAccountId: string | null
  /**
   * Last event timestamp observed (any direction). Surfaces in the
   * admin UI as "last event 2 minutes ago".
   */
  lastEventAt: Timestamp | null
  /**
   * ADR-128 - when a pull run last reached the provider without erroring,
   * mirrored from the ingestionPullRunStatus projection. Deliberately NOT
   * lastEventAt: an empty period is a working puller, and a run that
   * delivered data and then failed is not. Drives the health badge and
   * the "no data since" cut-off, below which a day's spend is unknown
   * rather than zero.
   * Spec: specs/governance/ingestion-source-health.feature
   */
  lastSuccessAt: Timestamp | null
  /**
   * ADR-128 - how far the last run actually read, and whether it reached the
   * end. Mirrored from the ingestionPullRunStatus projection, in the same way
   * and for the same reason as lastSuccessAt above: every reader of this
   * surface selects from this row, and two of them render a LIST of sources,
   * so reaching a per-run projection would be one query per row on a screen
   * built to show many at once.
   *
   * These two are the only evidence separating a source that truncated once
   * and then finished from one that is stuck. A stuck source's run clock
   * advances on every attempt, so the clock cannot be that evidence and is
   * deliberately not stored here - only the point the run REACHED.
   *
   * Completeness is "complete" or "truncated". Null on both is unknown, which
   * is what every row written before this existed reads as, and unknown must
   * never render as either answer.
   * Spec: specs/governance/ingestion-source-health.feature
   */
  lastReadThroughAt: Timestamp | null
  lastRunCompleteness: string | null
  /**
   * ADR-088 - the window this source read while pulled cost recording was
   * off for the organization, so its spend was dropped rather than priced.
   * Both null = nothing dropped. The pull cursor advances whether or not the
   * money path is live (audit-only is a supported mode), so without this the
   * loss is unrecoverable AND invisible: those days read as no spend rather
   * than as unknown. Widened by any run that drops a priced item, cleared
   * only when a later run re-reads back across the whole window.
   * Spec: specs/governance/pulled-usage-cost-gap.feature
   */
  unpricedUsageSince: Timestamp | null
  unpricedUsageThrough: Timestamp | null
  archivedAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  createdById: string | null
}

/**
 * Admin-defined anomaly detection rule. Evaluated server-side at
 * ingest time (Option C - not in this slice yet) against
 * gateway_activity_events. When triggered, an alert is dispatched
 * to the destinations configured on the rule (Slack / SIEM / webhook
 * / PagerDuty / email).
 *
 * This is a configuration entity - the actual rule evaluation +
 * alert state lives in a future ClickHouse table
 * (gateway_anomaly_alerts) and a server-side worker. Spec:
 * specs/ai-gateway/governance/anomaly-rules.feature
 * (Alexis lane B C-UI; mocked-v0 today, real persistence is THIS slice).
 */
export interface AnomalyRule {
  id: string
  organizationId: string
  /**
   * 'organization' | 'team' | 'project' | 'source_type' | 'source'.
   * String not enum so new scopes ship without migrations.
   */
  scope: string
  /**
   * Org id / team id / project id / source-type string / source id
   * - narrows what the rule applies to.
   */
  scopeId: string
  name: string
  description: string | null
  /**
   * 'critical' | 'warning' | 'info'. Drives default destination
   * routing (e.g. critical → PagerDuty, warning → Slack only).
   */
  severity: string
  /**
   * Extensible rule taxonomy: 'spend_spike' | 'rate_limit' |
   * 'tool_mismatch' | 'after_hours' | 'unusual_model' | 'pii_leak'
   * | 'custom'. New types ship without migrations - evaluator
   * picks a per-type implementation from a code registry.
   */
  ruleType: string
  /**
   * Per-rule-type threshold shape. Examples:
   * spend_spike: { windowSec, ratioVsBaseline, minBaselineUsd }
   * rate_limit: { windowSec, requestsThreshold }
   * after_hours: { startHour, endHour, timezone, requestsThreshold }
   */
  thresholdConfig: Generated<unknown>
  /**
   * Alert routing per destination. Shape:
   * { slack?: { channelId }, siem?: { endpointUrl, authHeader },
   * webhook?: { url, authHeader }, pagerduty?: { routingKey },
   * email?: { to: string[] } }
   * Empty map = use org defaults (when defaults exist; today no-op).
   */
  destinationConfig: Generated<unknown>
  /**
   * 'active' (default - evaluator picks it up) | 'disabled' (admin
   * paused; rule rows kept for audit). 'archived' uses archivedAt
   * timestamp instead so we can filter `archivedAt IS NULL` cheaply.
   */
  status: Generated<string>
  archivedAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  createdById: string | null
}

/**
 * Persisted anomaly detection. Inserted by the anomaly-detection
 * reactor (event-sourcing pipeline; modeled on PR #3351's
 * alertTrigger pattern) when a rule's threshold trips while
 * processing a freshly-arrived ActivityEventReceived event.
 * Read by api.activityMonitor.recentAnomalies for the admin dashboard.
 *
 * Dedup: (ruleId, triggerWindowStart) - re-evaluating the same
 * window updates the existing row rather than creating a duplicate.
 * Acknowledgement / resolution are admin actions on the row.
 *
 * Spec: specs/ai-gateway/governance/anomaly-detection.feature
 */
export interface AnomalyAlert {
  id: string
  organizationId: string
  ruleId: string
  /**
   * Denormalised from rule for fast WHERE filtering on the
   * dashboard query (no join needed for the common path).
   */
  severity: string
  ruleName: string
  ruleType: string
  /** Window the trigger refers to (UTC). */
  triggerWindowStart: Timestamp
  triggerWindowEnd: Timestamp
  /** Numeric trigger context. Per-rule-type detail in `detail` JSONB. */
  triggerSpendUsd: string | null
  triggerEventCount: number | null
  /**
   * Per-rule-type extra context: {topActors: [...], baseline: ...,
   * observed: ..., timezone: ..., etc}.
   */
  detail: Generated<unknown>
  /** 'open' | 'acknowledged' | 'resolved'. */
  state: Generated<string>
  /**
   * Per-destination dispatch outcome:
   * { webhook: { lastAttemptIso, lastError?, attemptCount, statusCode? },
   * slack: ..., pagerduty: ..., siem: ..., email: ... }
   */
  destinationStatus: Generated<unknown>
  detectedAt: Generated<Timestamp>
  acknowledgedAt: Timestamp | null
  resolvedAt: Timestamp | null
}

/**
 * Org-scoped catalog of AI tools surfaced on the personal home portal
 * (/me) and managed by org admins at /governance/tool-catalog.
 *
 * Three first-class tile types via `type` discriminator:
 * - "coding_assistant" - click expands to setup helper (e.g. `langwatch claude`)
 * - "model_provider"   - click expands to inline personal-VK creation
 * - "external_tool"    - click expands to admin-attached markdown + external link
 *
 * Scoping: organization-default entries can be OVERRIDDEN at the team
 * tier (same `slug` + `scope=team` + matching `scopeId`) so admins can
 * hide or re-skin a globally-listed tool for a specific team. The
 * user-facing `list` endpoint resolves this hierarchy at read time.
 *
 * `config` is a discriminated JSON union per `type`:
 * coding_assistant: { setupCommand, setupDocsUrl, helperText? }
 * model_provider:   { providerKey, suggestedRoutingPolicyId?, defaultLabel?, projectSuggestionText? }
 * external_tool:    { descriptionMarkdown, linkUrl, ctaLabel? }
 *
 * Spec: specs/ai-governance/personal-portal/tool-catalog-*.feature
 */
export interface AiToolEntry {
  id: string
  organizationId: string
  /**
   * "organization" | "team". String not enum so new scopes ship
   * without migrations.
   */
  scope: string
  /**
   * organizationId or teamId, narrowing what scope this entry applies
   * to. Always equals organizationId when scope=organization.
   */
  scopeId: string
  /**
   * "coding_assistant" | "model_provider" | "external_tool".
   * Discriminator drives the per-tile UX shape and validates
   * the `config` JSON column at the app layer.
   */
  type: string
  displayName: string
  /**
   * Stable slug for team-overrides-org resolution (a team-scoped
   * entry with the same slug as an org-scoped one hides the org
   * default for users in that team). Server-owned: generated from
   * displayName + nanoid(6) on create, immutable on update.
   */
  slug: string
  /**
   * @deprecated Use `iconAsset` instead. Kept for back-compat with
   * pre-refactor rows; new writes should populate iconAsset.
   */
  iconKey: string | null
  /**
   * Icon source - single column, prefix-discriminated:
   * "preset:claude_code" / "preset:codex" / ...   (built-in)
   * "data:image/svg+xml;base64,..."               (admin-uploaded)
   * null                                          → fall back to
   * type-default in UI
   */
  iconAsset: string | null
  /** Admin-curated sort within the user-facing portal grid. */
  order: Generated<number>
  /**
   * When false the entry is hidden from the user-facing list but
   * remains in the admin list (so admins can re-enable without
   * losing per-entry config). Disabled is "soft-hide", archived is
   * "soft-delete" (preserves audit trail + scope binding history).
   */
  enabled: Generated<boolean>
  /**
   * Per-type config JSON. Validated via zod discriminated union
   * at the service layer (see AiToolEntryService).
   */
  config: Generated<unknown>
  archivedAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  createdById: string | null
  updatedById: string | null
}

/**
 * @deprecated Per-team scope binding for AiToolEntry. Superseded by
 * AiToolEntryDepartment; kept so the rename migration is non-destructive.
 * No new rows are written here.
 */
export interface AiToolEntryTeam {
  id: string
  entryId: string
  teamId: string
}

/**
 * Per-department scope binding for AiToolEntry. Empty set = org-wide.
 * Non-empty = entry visible only to members of those departments (an
 * OrganizationUser whose `departmentId` is in the set). Mirrors
 * AiToolEntryTeam's shape. `departmentId` is a plain column ref
 * (relationMode="prisma"), matching how Department is referenced
 * elsewhere; the service validates the id belongs to the org.
 */
export interface AiToolEntryDepartment {
  id: string
  entryId: string
  departmentId: string
}

/**
 * Admin/platform-curated catalog row for upstream OTLP normalization.
 * Owns the OTTL transform that reshapes a specific upstream tool's
 * span shape into canonical gen_ai conventions. Users connect via an
 * ingestion key (an ingest-only ApiKey carrying ingestionTemplateId) to
 * land traces in their personal project.
 *
 * `organizationId` NULL = platform-published default (always-visible
 * across every org). NOT NULL = org-authored (visible only within
 * that org). v1 ships only platform-published; org-authored authoring
 * UI lands v2 (the column ships now to avoid a future migration).
 *
 * Spec: specs/ai-gateway/governance/ingestion-templates-catalog.feature
 */
export interface IngestionTemplate {
  id: string
  organizationId: string | null
  /**
   * Stable slug - unique within (organizationId, slug) scope. Generated
   * from displayName + nanoid(6) on create for org-authored; canonical
   * values for platform-published (`claude_code` / `cursor` /
   * `claude_cowork` / `raw_otlp_advanced`).
   */
  slug: string
  /**
   * Mirrors the slug for v1 platform-published rows. Kept as a separate
   * column so a future v2 may publish two templates that share an
   * upstream sourceType (e.g. claude_code-personal vs claude_code-team)
   * without a migration. String not enum - extensibility.
   */
  sourceType: string
  displayName: string
  description: string | null
  /**
   * Icon source - single column, prefix-discriminated, mirrors
   * AiToolEntry.iconAsset:
   * `preset:<kind>`                 (built-in slot)
   * `data:image/svg+xml;base64,…`   (admin-uploaded for org-authored)
   * null                            → fall back to slug-default in UI
   */
  iconAsset: string | null
  /**
   * Discriminator for what (if any) extra credential metadata the
   * install drawer captures from the user:
   * null              → otlp_token-only (drawer auto-issues the
   * ingestion key, no input form)
   * `static_api_key`  → drawer captures the user's upstream tool API key
   * `agent_id`        → drawer captures an agent identifier (e.g.
   * Cursor agentId)
   * String not enum so v2 shapes (oauth_redirect / s3_bucket) ship
   * without a migration.
   */
  credentialSchema: string | null
  /**
   * Admin/platform-authored OTTL statements. Applied at receive time
   * AFTER the receiver re-stamps the binding's authoritative principal
   * + provenance keys (see `protectedTemplateAttributeKeys` 19-key
   * closed list). Templates write canonical OTTL outputs (gen_ai.*,
   * cost.usd, response.model) - they CANNOT write attribution or
   * provenance keys; the receiver guard restores those post-OTTL.
   */
  ottlRules: Generated<string>
  /**
   * True for the v1 platform-published default templates
   * (organizationId IS NULL). Future org-authored rows write false.
   * Drives the admin UI "READ-ONLY view + Request a custom template"
   * surface - admins consume + opt-out only in v1.
   */
  platformPublished: Generated<boolean>
  /**
   * Admin/platform enable. False = hidden from /me catalog. Existing
   * ingestion keys continue to land traces; new installs are blocked.
   */
  enabled: Generated<boolean>
  archivedAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  createdById: string | null
  updatedById: string | null
}

export interface GatewayA2ATarget {
  id: string
  organizationId: string
  projectId: string | null
  name: string
  agentName: string
  cardUrl: string
  upstreamUrl: string
  encryptedUpstreamHeaders: string | null
  restrictToAllowList: Generated<boolean>
  allowedVirtualKeyIds: Generated<string[]>
  blocked: Generated<boolean>
  blockReason: string | null
  createdById: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  archivedAt: Timestamp | null
}

export interface GatewayConnectUpstream {
  organizationId: string
  baseUrl: string
  encryptedToken: string
  instanceId: string
  updatedAt: Timestamp
}

export interface GatewayTraceExportKey {
  projectId: string
  apiKeyId: string
  encryptedToken: string
  createdAt: Generated<Timestamp>
}

export interface GatewayBudget {
  id: string
  organizationId: string
  scopeType: GatewayBudgetScopeType
  scopeId: string
  providerKey: string | null
  name: string
  description: string | null
  window: GatewayBudgetWindow
  limitUsd: string
  onBreach: Generated<GatewayBudgetBreachAction>
  timezone: string | null
  externalId: string | null
  metadata: Generated<unknown>
  spentUsd: Generated<string>
  currentPeriodStartedAt: Generated<Timestamp>
  resetsAt: Timestamp
  lastResetAt: Timestamp | null
  cycleAnchorAt: Timestamp | null
  archivedAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  createdById: string
  managedByVirtualKeyId: string | null
}

export interface GatewayBudgetBucketBoundary {
  id: string
  organizationId: string
  budgetId: string
  bucketScopeId: string
  periodStartedAt: Timestamp
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface GatewayBudgetLedger {
  id: string
  budgetId: string
  virtualKeyId: string
  modelProviderId: string | null
  providerKey: string | null
  gatewayRequestId: string
  amountUsd: string
  tokensInput: Generated<number>
  tokensOutput: Generated<number>
  tokensCacheRead: Generated<number>
  tokensCacheWrite: Generated<number>
  model: string
  durationMs: number | null
  status: Generated<GatewayBudgetLedgerStatus>
  occurredAt: Generated<Timestamp>
}

export interface GatewayChangeEvent {
  revision: Generated<ColumnType<string, bigint | number | string, bigint | number | string>>
  organizationId: string
  kind: GatewayChangeEventKind
  virtualKeyId: string | null
  budgetId: string | null
  modelProviderId: string | null
  projectId: string | null
  payload: unknown
  createdAt: Generated<Timestamp>
}

export interface GatewayCacheRule {
  id: string
  organizationId: string
  name: string
  description: string | null
  priority: Generated<number>
  enabled: Generated<boolean>
  matchers: unknown
  action: unknown
  modeEnum: GatewayCacheRuleMode
  archivedAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  createdById: string
}

export interface FeatureFlag {
  key: string
  enabled: boolean
  rules: unknown
  lastEditedBy: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/**
 * One experiment setting per (flag, subject). A USER row is that person's
 * own opt-in; an ORGANIZATION or PROJECT row is an owner's explicit policy
 * for that exact scope. No row means inherit.
 */
export interface FeatureFlagExperimentSetting {
  flagKey: string
  subjectType: FeatureFlagExperimentSubjectType
  subjectId: string
  enabled: boolean
  changedByUserId: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/** Durable state for a long-running event-sourced process. */
export interface ProcessManagerInstance {
  id: string
  processName: string
  projectId: string
  processKey: string
  tenantId: string
  userId: string | null
  state: unknown
  revision: number
  nextWakeAt: Timestamp | null
  updatedAt: Timestamp
}

/** Idempotency marker for a source event consumed by a process definition. */
export interface ProcessManagerInbox {
  id: string
  processName: string
  projectId: string
  processKey: string
  tenantId: string
  /**
   * The raw source event id, kept UNINDEXED for diagnostics. A caller composes
   * this and its length is that domain's business, so it must never carry a
   * btree constraint — see `sourceEventKey`.
   */
  sourceEventId: string
  /**
   * The fixed-width digest of `sourceEventId` the uniqueness is enforced on,
   * derived by the store (`stores/inboxKey.ts`). Postgres cannot index a btree
   * row past ~2704 bytes; keying on a digest makes the constraint the same
   * width no matter what any pipeline concatenates into an idempotency key.
   *
   * Nullable ONLY for the rolling deploy that introduces it — pods running the
   * previous build write rows without it, and a NOT NULL would fail their
   * commits outright. The store always supplies it.
   *
   * The trade, stated plainly: a NULL is not constrained by the unique index
   * (btree treats NULLs as distinct), so while an old writer is serving, the
   * one case that loses its guard is the same source event reaching two
   * different process keys at once — the store's own read-check plus its
   * per-processKey advisory lock cover everything else. A follow-up migration
   * adds NOT NULL once the fleet has cycled and closes it.
   */
  sourceEventKey: string | null
  consumedAt: Timestamp
}

/** Transactional intents emitted by a durable process transition. */
export interface ProcessManagerOutbox {
  id: string
  processName: string
  projectId: string
  processKey: string
  tenantId: string
  userId: string | null
  messageKey: string
  intentType: string
  payload: unknown
  traceCarrier: unknown
  sourceEventId: string | null
  status: Generated<ProcessManagerOutboxStatus>
  attempts: Generated<number>
  nextAttemptAt: Timestamp
  leasedUntil: Timestamp | null
  leaseToken: string | null
  dispatchedAt: Timestamp | null
  createdAt: Timestamp
  updatedAt: Timestamp
}

/**
 * Why an outbox delivery failed, one row per FAILED attempt
 * (specs/ops/dead-letter-recovery.feature). Successes write nothing, so the
 * table grows with trouble, not with traffic. Rows die with their message:
 * the cascade delete is the retention story, riding the existing sweep.
 */
export interface ProcessManagerOutboxAttempt {
  id: string
  outboxId: string
  projectId: string
  /** 1-based attempt number, as the dispatcher counts it. */
  attempt: number
  occurredAt: Timestamp
  outcome: ProcessManagerOutboxAttemptOutcome
  errorType: string
  /**
   * The safe failure diagnostic the dispatcher already computes for its
   * span — never a raw provider body.
   */
  errorMessage: string
  retryAfterMs: number | null
}

/**
 * Per-project topic clustering run status (ADR-051): the rebuildable
 * operational read model behind the settings page. One row per project;
 * folded from run_completed/run_failed events, never written directly.
 */
export interface TopicClusteringRunProjection {
  /** KSUID (`topicrun_…`), generated by the repository on first upsert. */
  id: string
  projectId: string
  lastRequestedAt: number | null
  lastRequestTrigger: string | null
  lastRunAt: number | null
  /** completed | skipped | failed */
  lastRunOutcome: string | null
  lastRunMode: string | null
  lastRunSkippedReason: string | null
  lastRunError: string | null
  lastRunErrorCode: string | null
  /** True when the customer can resolve the failure (credentials, quota). */
  lastRunErrorUserActionable: Generated<boolean>
  lastRunTracesProcessed: Generated<number>
  lastRunTopicsCount: Generated<number>
  lastRunSubtopicsCount: Generated<number>
  lastRunPages: Generated<number>
  inProgressRunId: string | null
  inProgressTraces: Generated<number>
  inProgressPages: Generated<number>
  /**
   * Business time the in-progress run opened; bounds how long the read
   * model may report it as running (mirrors the process's stale-run guard).
   */
  inProgressStartedAt: number | null
  createdAt: number
  updatedAt: number
  /** Business time of the latest applied event. */
  occurredAt: number
  /** Event-log acceptance time. Together with lastEventId this is the cursor. */
  acceptedAt: number
  lastEventId: string
  projectionVersion: string
}

/**
 * Cursor row for the topic-model projection (ADR-051). The model itself is
 * the `Topic` table, written through by the projection; this row only
 * records how far the fold has applied the event log.
 */
export interface TopicModelProjection {
  /** KSUID (`topicmodel_…`), generated by the repository on first upsert. */
  id: string
  projectId: string
  createdAt: number
  updatedAt: number
  /** Business time of the latest applied event. */
  occurredAt: number
  /** Event-log acceptance time. Together with lastEventId this is the cursor. */
  acceptedAt: number
  lastEventId: string
  projectionVersion: string
}

/**
 * Per-project topic clustering run history (audit read model, ADR-051).
 * One row per project; `runs` is the bounded newest-first entry list.
 * Rebuildable by replaying the event log.
 */
export interface TopicClusteringRunHistoryProjection {
  /** KSUID (`topicrunhist_…`), generated by the repository on first upsert. */
  id: string
  projectId: string
  runs: unknown
  createdAt: number
  updatedAt: number
  /** Business time of the latest applied event. */
  occurredAt: number
  /** Event-log acceptance time. Together with lastEventId this is the cursor. */
  acceptedAt: number
  lastEventId: string
  projectionVersion: string
}

/** Rebuildable per-source pull checkpoint and operator-facing run status. */
export interface IngestionPullRunProjection {
  id: string
  sourceId: string
  projectId: string
  enabled: Generated<boolean>
  cron: string | null
  cursor: string | null
  lastRunAt: number | null
  lastRunOutcome: string | null
  lastRunEventCount: Generated<number>
  lastRunError: string | null
  lastRunErrorCode: string | null
  consecutiveErrors: Generated<number>
  /**
   * When a run last completed without erroring. Backfilled by replay, which
   * is why the projection version moved with this column.
   */
  lastSuccessAt: number | null
  /**
   * Identity of the run whose outcome this row reflects. Outcomes from a
   * superseded run arrive with a smaller value and are ignored, so a late
   * completion cannot regress Cursor or the status fields.
   */
  lastRunScheduledFor: number | null
  /**
   * ADR-128 - how far the last run read, in epoch milliseconds, and whether
   * it reached the end. The fold writes both; the IngestionSource row mirrors
   * them for the screens. Null read-through means no run has reported one.
   * Null completeness reads as unknown, never as either answer, which is what
   * every row written before this existed holds.
   *
   * Nullable on purpose, so replay is not required to make this migration
   * safe: a row that has never seen a run carrying these stays null and every
   * reader already handles unknown.
   * Spec: specs/governance/ingestion-source-health.feature
   */
  lastReadThroughAt: number | null
  lastRunCompleteness: string | null
  /**
   * The last time somebody asked this source to enumerate its agents, and how
   * that ask ended. A listing is a different scope on the same connection as
   * the pull above: a credential that cannot enumerate a directory still pulls
   * cost perfectly, so no health, coverage or completeness reader may treat a
   * refusal here as the source failing to pull.
   *
   * Count is set ONLY when the outcome is listed, and stays null on a refusal.
   * Zero is a real answer meaning the provider returned an empty list, so a
   * zero default would manufacture that answer for a source that refused.
   * Reason is set only on a refusal and is a plain string, never an enum and
   * never a foreign key, because the log outlives the vocabulary. Status is
   * nullable even when refused, because a refusal we generated ourselves has
   * no HTTP status, and it is never rendered to a customer.
   *
   * The two `...ListingAt` columns hold the instant the ASK was accepted, not
   * the instant its answer was recorded, and the fold orders outcomes by them
   * so a superseded listing cannot overwrite the one that replaced it. The
   * name predates that job and is kept because renaming a column is a
   * migration; the fold's own field docs are the statement of record.
   */
  lastAgentsListingAt: number | null
  lastAgentsListingOutcome: string | null
  lastAgentsListingCount: number | null
  lastAgentsListingReason: string | null
  lastAgentsListingStatus: number | null
  /**
   * The same for people, in its own set of columns rather than sharing the
   * agents set with a kind discriminator. A source can be refused for people
   * and fine for agents, and one shared set holds only the most recent listing
   * of either kind, so syncing people would erase an agents refusal and the
   * agents page would go back to not knowing after briefly knowing.
   *
   * LastPeopleDirectoryCount is everyone the provider's directory named.
   * LastPeopleWithheldCount is how many of those this deployment does not hold
   * because erasure suppression removed them: a SUBSET of the directory count,
   * never an addition to it. Adding the two counts the same people twice.
   * Subtracting is the valid arithmetic, and the surviving total is that
   * subtraction rather than a third column that could disagree.
   *
   * The withheld count is safe to show as a CURRENT figure and never as a
   * series: it moving from zero to one at a known moment says an erasure
   * happened then, which on a small tenant identifies the person as surely as
   * a name would. So no trend line, no history drawer, no per-run export, and
   * never beside a per-person list.
   */
  lastPeopleListingAt: number | null
  lastPeopleListingOutcome: string | null
  lastPeopleDirectoryCount: number | null
  lastPeopleWithheldCount: number | null
  lastPeopleListingReason: string | null
  lastPeopleListingStatus: number | null
  createdAt: number
  updatedAt: number
  occurredAt: number
  acceptedAt: number
  lastEventId: string
  projectionVersion: string
}

export interface ScheduledJob {
  id: string
  projectId: string
  targetType: string
  targetId: string
  cron: string
  timezone: string
  nextRunAt: Timestamp
  lastSlot: Timestamp | null
  currentSlot: Timestamp | null
  attempts: Generated<number>
  lastError: string | null
  active: Generated<boolean>
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface BugReport {
  id: string
  createdAt: Generated<Timestamp>
  source: string
  kind: string
  title: string
  summary: string | null
  sessionData: string | null
  sessionTruncated: Generated<boolean>
  agent: string | null
  contactEmail: string | null
  cliVersion: string | null
  linkedProjectId: string | null
  metadata: unknown
}

export interface IdempotencyReceipt {
  id: string
  scopeId: string
  key: string
  claimId: string
  requestFingerprint: string
  responseStatus: number | null
  responseBody: string | null
  createdAt: Generated<Timestamp>
  heartbeatAt: Timestamp
  expiresAt: Timestamp
}

export interface GatewayRealtimeSession {
  id: string
  projectId: string
  organizationId: string
  virtualKeyId: string
  modelProviderId: string
  vendor: string
  agentId: string | null
  model: string
  requestedModel: string | null
  vendorConversationId: string | null
  status: Generated<GatewayRealtimeSessionStatus>
  mintedAt: Generated<Timestamp>
  closedAt: Timestamp | null
  closeReason: string | null
  traceId: string | null
  vendorCostRaw: unknown
  kind: string | null
  metering: string | null
  credentialExpiresAt: Timestamp | null
  transcriptionModel: string | null
  endUserId: string | null
  lastReportAt: Timestamp | null
  reportedCostNanoUsd: Generated<
    ColumnType<string, bigint | number | string, bigint | number | string>
  >
  reportCount: Generated<number>
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface GatewayRealtimeSessionReport {
  id: string
  sessionId: string
  reportKey: string
  projectId: string
  model: string
  usage: unknown
  costNanoUsd: ColumnType<string, bigint | number | string, bigint | number | string>
  createdAt: Generated<Timestamp>
}

export interface SystemMigrationTenantState {
  migrationName: string
  tenantId: string
  status: string
  report: unknown
  occurredAt: Generated<Timestamp>
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

export interface SystemMigrationEnrollment {
  organizationId: string
  migrationName: string
  enrolledByUserId: string
  createdAt: Generated<Timestamp>
}

export interface StoredObject {
  tenantId: string
  id: string
  status: string
  purpose: string
  ownerKind: string
  ownerId: string
  filename: string
  sha256: string
  sizeBytes: ColumnType<string, bigint | number | string, bigint | number | string>
  mediaType: string
  mediaTypeVerified: Generated<boolean>
  storageProvider: string | null
  storageDestinationId: string | null
  storageProviderRelativeId: string | null
  generation: Generated<number>
  audiences: unknown
  expiresAt: Timestamp | null
  availableAt: Timestamp | null
  deletedAt: Timestamp | null
  source: string
  legacyFingerprint: string | null
  createdAt: Timestamp
  updatedAt: Generated<Timestamp>
}

/**
 * A name-text seen on a provider's cost or audit rows: "the provider told us
 * someone called m.silva exists". One row per provider-scoped identity.
 *
 * This is the unit of GDPR erasure (ADR-128 §9), and deliberately so: most of
 * these people have no LangWatch login at all - contractors, seat holders who
 * never signed in, anyone whose email a provider put on a cost row - so there
 * is no platform user to key on and a user-deletion event can never fire for
 * them.
 */
export interface DiscoveredPerson {
  id: string
  organizationId: string
  /** "anthropic" | "openai" | "databricks" | ... - the provider that named them. */
  provider: string
  /**
   * What the provider calls them, verbatim. The join key onto the money rows'
   * RawActorId. Replaced by a deterministic pseudonym on erasure, which keeps
   * the row addressable for spend attribution without keeping the identifier.
   */
  rawActorId: string
  /** Name or email text as seen. Pseudonymized alongside rawActorId on erasure. */
  displayText: string
  /**
   * person | service_account - determined deterministically per ADR-128 §10
   * (a bare UUID in Databricks `executed_by` is a service account, never an
   * agent), never guessed from the shape of a name.
   */
  kind: string
  /**
   * The department the provider's directory filed this person under, verbatim.
   * An OBSERVATION, not the `Department` entity: it is free text a tenant owns
   * on their side, it exists for people who hold no LangWatch account, and
   * nothing rolls spend up by it. Spend attribution stays on `Department`, via
   * the membership the match engine proves (directoryDepartmentSync.service).
   * Null means no directory has named one - never "no department", because a
   * blank field in the tenant's directory and a person no directory lists are
   * the same null here.
   */
  department: string | null
  firstSeenAt: Timestamp
  lastSeenAt: Timestamp
  /**
   * Set when this row's identity has been erased. The row survives so its
   * spend stays attributable to *someone*; the identifier does not.
   */
  erasedAt: Timestamp | null
  /**
   * Set alongside `erasedAt` and cleared once the daily cost rows have been
   * removed and their rebuild requested. Non-null means the ClickHouse half
   * of the erasure did not finish — a crash between the delete and the
   * replay leaves days short by the erased amount, and without this marker a
   * re-run reads `erasedAt`, short-circuits, and reports a clean erasure that
   * did no work. The re-run resumes at that step instead.
   */
  moneyRowsPendingAt: Timestamp | null
  /**
   * The earliest day to rebuild from, recorded BEFORE the rows are deleted.
   * Once they are gone there is nothing left to ask which days they were on,
   * so a resume with no record of this could only replay everything or
   * nothing. Null means nothing was rebuildable — either no day carried the
   * identifier, or every one of them is older than the event log keeps.
   */
  moneyRebuildSince: string | null
  /**
   * ADR-128 §12 conflict rule: automatic linking has stopped for this
   * identity because evidence pointed at two candidates, or contradicted an
   * existing link. Stored HERE rather than on the suggestion rows, because a
   * halt on auto-linking that a recompute or a restart clears is not a halt.
   */
  suspendedAt: Timestamp | null
  /** What tripped the suspension, for the human who reviews it. */
  suspendedReason: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/**
 * The same shape as DiscoveredPerson for non-human actors - Genie spaces,
 * Copilot bots, OpenAI projects - plus whatever provider-native metadata the
 * source published. Discovering an agent grants it nothing either.
 */
export interface DiscoveredAgent {
  id: string
  organizationId: string
  provider: string
  /** Genie space / Copilot bot / OpenAI project id, verbatim. */
  rawAgentId: string
  displayText: string
  /**
   * Workspace, plan, and kin - provider-native descriptive fields only. Never
   * credentials: this row is read by admin screens, not by any pull path.
   */
  metadata: unknown
  firstSeenAt: Timestamp
  lastSeenAt: Timestamp
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/**
 * One source's presence fact for one discovered agent. An agent is gone only when every
 * source that listed it marked it gone; an agent no source listed (older sightings) never is.
 */
export interface DiscoveredAgentListing {
  id: string
  organizationId: string
  ingestionSourceId: string
  discoveredAgentId: string
  lastListedAt: Timestamp
  /** Set when this source's listing stopped naming the agent; cleared when it names it again. */
  goneAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/**
 * One agent in the organisation's governance registry, from any source governance reads
 * (enterprise/modules/governance/specs/agent-registry.feature). `entry` is the latest
 * machine-readable snapshot; `version` grows by one on every material change.
 */
export interface AgentRegistryEntry {
  id: string
  organizationId: string
  /** Where the agent came from and its id there, e.g. `langwatch:<agentId>`. */
  sourceKey: string
  source: string
  status: string
  version: number
  fingerprint: string
  entry: unknown
  firstSeenAt: Timestamp
  lastSeenAt: Timestamp | null
  removedAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/**
 * The registry's append-only change feed: one row per registered, versioned, removed or
 * tool-deprecated fact. `sequence` is the pull cursor; `deliveryRecordedAt` marks the webhook
 * fan-out as recorded in governance's outbox.
 */
export interface AgentRegistryEvent {
  id: string
  organizationId: string
  sequence: Generated<ColumnType<string, bigint | number | string, bigint | number | string>>
  type: string
  entryId: string
  fromVersion: number | null
  toVersion: number | null
  payload: unknown
  occurredAt: Timestamp
  deliveryRecordedAt: Timestamp | null
}

/**
 * An organisation's registry baseline: its first reconcile, written as already pushed so the
 * agents it already had are not announced. Every reconcile after it pushes its changes.
 */
export interface AgentRegistryBaseline {
  organizationId: string
  baselinedAt: Timestamp
}

/** An organisation's agent registry settings. No row reads as the defaults. */
export interface AgentRegistrySettings {
  organizationId: string
  /** Days without traffic after which a listed agent reads as dormant. */
  dormantAfterDays: Generated<number>
  updatedAt: Generated<Timestamp>
}

/**
 * One target of a registry agent's kill switch stop plan (enterprise/modules/governance/specs/
 * agent-kill-switch.feature): a Databricks serving endpoint, a Bedrock agent alias or a gateway
 * virtual key. A plan write replaces every row of the agent at once.
 */
export interface AgentStopPlanTarget {
  id: string
  organizationId: string
  /** The AgentRegistryEntry the plan belongs to. */
  agentId: string
  position: number
  label: string
  target: unknown
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/**
 * One agent's stop plan version: bumped by every save, so a save over a plan someone else saved
 * since it was read is refused. No row means the plan was never saved (version 0).
 */
export interface AgentStopPlan {
  organizationId: string
  /** The AgentRegistryEntry the plan belongs to. */
  agentId: string
  version: number
  updatedAt: Generated<Timestamp>
}

/**
 * An AWS credential a Bedrock kill switch target signs with: a static key pair, a role to assume,
 * or both. The secret access key and the external id are stored sealed.
 */
export interface GovernanceAwsConnection {
  id: string
  organizationId: string
  name: string
  accessKeyId: string | null
  encryptedSecretAccessKey: string | null
  roleArn: string | null
  encryptedExternalId: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/**
 * The `agent_stop` aggregate's Postgres projection: one row per kill switch stop, its status and
 * every target's outcome, folded from the stop's events.
 */
export interface AgentStop {
  id: string
  organizationId: string
  agentId: string
  agentName: string
  status: string
  severity: string
  approval: string
  trigger: string
  reason: string
  requestedByUserId: string | null
  requestedByApiKeyId: string | null
  decidedByUserId: string | null
  decidedByApiKeyId: string | null
  restoredByUserId: string | null
  restoredByApiKeyId: string | null
  targets: unknown
  requestedAt: Timestamp
  decidedAt: Timestamp | null
  quarantinedAt: Timestamp | null
  restoreRequestedAt: Timestamp | null
  restoredAt: Timestamp | null
  lastAppliedEventId: string
  lastAcceptedAt: ColumnType<string, bigint | number | string, bigint | number | string>
  lastOccurredAt: ColumnType<string, bigint | number | string, bigint | number | string>
  projectionVersion: string
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
}

/**
 * The one open stop an agent may have: the api takes this row before it sends a stop's request,
 * so two requests at once cannot both open one. A later request takes it over once the stop it
 * names has closed.
 * `sentAt` is set once the request command was sent: from then on the claim passes on only after
 * its stop closed, never on age. `decision` is taken by the one approve or deny that wins.
 * `decisionTakenAt` is set while that decision is in flight and cleared once its command was
 * sent; a decision left in flight is taken over once abandoned.
 */
export interface AgentStopClaim {
  organizationId: string
  agentId: string
  stopId: string
  claimedAt: Timestamp
  sentAt: Timestamp | null
  decision: string | null
  decisionTakenAt: Timestamp | null
}

/**
 * Whether a target was already stopped before a stop first touched it, recorded once before the
 * stop's platform call, so a retried call after a success still knows the change is the stop's own.
 */
export interface AgentStopTargetPriorState {
  organizationId: string
  stopId: string
  targetId: string
  wasStopped: boolean
  recordedAt: Timestamp
}

/**
 * The dated match table: "`user_abc123` on Anthropic is platform user #42,
 * valid from March 3", closed when it stops being true.
 *
 * Dating is what makes a re-issued email address survivable: the leaver's link
 * closes at offboarding, the new hire gets a new link, and January's spend
 * stays with January's person. A correction closes the old row and opens a new
 * one; nothing is ever silently rewritten.
 */
export interface IdentityMatch {
  id: string
  organizationId: string
  /** The provider-side identity this link is about. */
  discoveredPersonId: string
  /**
   * The platform user. Nullable, and blanked by erasure (ADR-128 §9 step 2) -
   * the row and its dates remain, so the history of the link is not rewritten
   * by the erasure either.
   */
  userId: string | null
  /**
   * directory_id | verified_email | human_confirmed. Exact verified-email
   * equality is the primary evidence; a directory id corroborates but never
   * stands alone (ADR-128 §12).
   */
  evidenceKind: string
  validFrom: Timestamp
  /** Null means the link is open. Offboarding and corrections close it. */
  validTo: Timestamp | null
  createdAt: Generated<Timestamp>
}

/**
 * Every ClickHouse `TenantId` this organization has ever written governance
 * rows under, appended on first use and never pruned.
 *
 * The live `resolveGovProjectId` lookup cannot be this translation: it filters
 * `archivedAt: null` and so returns null forever once someone archives the
 * governance project, while the write path re-reads the same project by slug
 * with no such filter and keeps landing rows under the old tenant. The
 * organization would map to zero tenants on read and one on write - a
 * permanent split-brain in which the cost screen reports "no governance data"
 * and an erasure job erases nothing and reports success.
 *
 * Reads resolve against the whole history; erasure walks all of it, because
 * personal data does not stop existing in a tenant that stopped being current.
 */
export interface GovernanceTenantHistory {
  id: string
  organizationId: string
  /** A hidden governance project id this organization has written rows under. */
  tenantId: string
  firstUsedAt: Timestamp
  lastUsedAt: Timestamp
}

/**
 * Hashes of erased identifiers - the do-not-reimport list.
 *
 * Without it the erasure is undone by the pipeline that produced it: the
 * pullers look 30 days back, so the next day's pull re-ingests the same email
 * and re-writes what we erased. Every write path carrying an identifier
 * consults this list and drops a suppressed one (ADR-128 §9 step 1).
 *
 * It stores hashes and never identifiers, so it is not itself a copy of the
 * data it exists to keep out. The stored `identifierHash` and the pseudonym
 * written in place of the erased value are the same function of the same
 * input - HMAC-SHA256(secret, identifier) - so a write path computes that digest
 * once and uses it twice: as the membership test, and, on a hit, as the
 * replacement value. That is why no original-to-pseudonym mapping table
 * exists; one would hold the erased identifier in plaintext forever.
 *
 * Scoped to the provider as well as the organization: the same string can be a
 * different person under a different provider.
 */
export interface ErasedIdentifierSuppression {
  id: string
  organizationId: string
  provider: string
  identifierHash: string
  erasedAt: Timestamp
}

/**
 * Match candidates a background job computed and a human has yet to confirm
 * (ADR-128 §12).
 *
 * Stored rather than computed when the review screen asks, and the reason is
 * measured rather than stylistic: fuzzy matching is quadratic, there is no
 * database route to it here, and at this document's own example size - 2,000
 * discovered people against 500 accounts - plain edit distance over the
 * 1,000,000 pairs blocked the event loop for 2.9 seconds. Per page load,
 * uncached, stalling every other request on the instance. A stored row also
 * makes a pending count answerable without paying for the whole sweep.
 *
 * Rows are written only by the job, which prefilters before it scores and
 * recomputes when its inputs change. Confirming one opens an `IdentityMatch`
 * and deletes the suggestion.
 *
 * Suspension deliberately does NOT live here - it is on `DiscoveredPerson`,
 * because a halt on automatic linking that the next recompute clears is not a
 * halt.
 */
export interface IdentityMatchSuggestion {
  id: string
  organizationId: string
  discoveredPersonId: string
  /**
   * The platform user this might be. Not nullable: a suggestion with no
   * candidate is not a suggestion.
   */
  userId: string
  /**
   * Similarity of the two name texts, after the prefilter, in [0, 1]. A
   * display hint and an ordering key - never a threshold anything auto-links
   * on, which is the whole distinction this table exists to hold.
   */
  score: number
  computedAt: Timestamp
}

/**
 * What a mid-term seat change of one connected license owes (ADR-141,
 * section 7): the seats added on the reissued license, invoiced prorated to
 * the end of the term.
 *
 * The row is written before the payment provider is called and completed
 * after it, so a call that failed halfway is retried by the next seat
 * invoicing pass without inventing a second invoice. One row per reissued license.
 */
export interface ConnectedSeatChange {
  id: string
  /**
   * The reissued `IssuedLicense` row that carries the new seat count. A
   * plain column ref, like `replacesId`: relationMode="prisma" carries no
   * foreign keys.
   */
  licenseId: string
  /**
   * The `ConnectedBillingAccount` row the invoice is raised against. Null
   * when the customer had no billing account (`not_onboarded`).
   */
  accountId: string | null
  /** The customer organization the license belongs to. */
  organizationId: string | null
  /** When the seats changed, which is the day the proration counts from. */
  changedAt: Timestamp
  /** Seats on the new license above the seats on the one it replaced. */
  addedSeats: number
  /** What one added seat costs for the rest of the term, in cents of `currency`. */
  unitAmountCents: number
  /** `unitAmountCents` times `addedSeats`. Zero when nothing is owed. */
  amountCents: number
  /** Null when no billing account named one (`not_onboarded`). */
  currency: Currency | null
  /** The one-off invoice the added seats were billed on. */
  stripeInvoiceId: string | null
  state: ConnectedSeatChangeState
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

/**
 * The monthly usage statement one connected customer was sent (ADR-139,
 * section 7). One row per account per month keeps the statement to one send,
 * whatever runs the tick.
 */
export interface ConnectedStatement {
  id: string
  /** The `ConnectedBillingAccount` row. A plain column ref. */
  accountId: string
  /** First day of the month the statement covers, in UTC. */
  month: Timestamp
  sentAt: Timestamp
}

/**
 * One renewed licence per paid Stripe invoice of a self-serve licence
 * subscription (specs/licensing/self-serving-license-purchase.feature). A
 * redelivered invoice event finds the row and mints nothing. A claim never
 * delivered is taken over once `claimedAt` is older than the claim lease.
 */
export interface LicenseRenewalClaim {
  stripeInvoiceId: string
  stripeSubscriptionId: string
  /** Fences completion and release to the delivery that holds the claim. */
  claimToken: string
  claimedAt: Timestamp
  /** The signed licence between minting and mailing, so a retry mails it again instead of minting another; cleared once delivered. */
  issuedLicense: unknown
  deliveredAt: Timestamp | null
  /** The `licenseId` inside the delivered licence. */
  licenseId: string | null
}

/**
 * The identity this install presents to LangWatch (ADR-139, section 2).
 *
 * Exactly one row, minted on first use and never rewritten. It survives
 * restarts, backups and hostname changes, and an install restored into
 * another database mints another one, which is what the registry refuses as
 * the wrong instance.
 *
 * It replaced an identity built from the organization name and id, which
 * leaked the customer's name in cleartext, named one organization rather than
 * one install, and could not be joined to a license row without splitting a
 * string.
 */
export interface InstanceIdentity {
  /**
   * Fixed, so this table holds one row. A second insert loses the race on the
   * primary key and re-reads the winner's id.
   */
  id: Generated<string>
  /** A random UUID. Carries nothing about the customer. */
  instanceId: string
  createdAt: Generated<Timestamp>
  /**
   * When the usage report last reached LangWatch, and why the last attempt
   * did not. Both are what the checkup page reads to say whether reporting
   * works, which nothing recorded before: the sender logged success as soon
   * as `fetch` resolved and never looked at the response.
   */
  lastReportAt: Timestamp | null
  lastReportError: string | null
  /**
   * The optional category of the usage report, switched off from the checkup
   * page. With it off the report still says which release runs and how many
   * projects the install carries, and says nothing about what anyone does
   * with it (ADR-139, section 10).
   */
  optionalMetricsOptOut: Generated<boolean>
  /**
   * Hostname, which names the customer's own network and so has a switch of
   * its own rather than travelling with its category.
   */
  hostnameOptOut: Generated<boolean>
  /**
   * Sending a failed upgrade step's diagnostics to LangWatch on its own,
   * switched off by the operator (upgrade-diagnostics-via-connect, decision 4).
   */
  upgradeDiagnosticsOptOut: Generated<boolean>
}

/**
 * One self-hosted install, as its daily usage report describes it (ADR-139,
 * section 10). The row is created by the first report that names an instance
 * id and is updated by every report after that, so the backoffice reads one
 * row per install rather than folding a history on every page load.
 *
 * The report carries no credential, so it identifies an install and never a
 * customer. `organizationId` and `issuedLicenseId` are resolved from the
 * license bound to this instance, which the install proved by presenting its
 * token to license sync. A field in the posted body never reaches them.
 */
export interface SelfHostedInstance {
  id: string
  /** The UUID the install minted for itself. The join key for everything else. */
  instanceId: string
  firstSeenAt: Generated<Timestamp>
  lastSeenAt: Timestamp
  /** What the install says about itself. */
  version: string | null
  installMethod: string | null
  chartVersion: string | null
  /** The host it answers on, when the operator left hostname reporting on. */
  hostname: string | null
  environment: string | null
  /**
   * The day the install itself says it first came up, which predates the
   * first report on an install that switched reporting on later.
   */
  installedAt: Timestamp | null
  reportSchemaVersion: number | null
  /**
   * The customer, and the license that named them. Plain column refs rather
   * than relations, like `IssuedLicense.replacesId`: relationMode="prisma"
   * carries no foreign keys, and the receiver resolves both from the license
   * registry rather than from the report.
   */
  organizationId: string | null
  issuedLicenseId: string | null
  /**
   * Domains with counts, as `{"acme.test": 14}`. Never an address. Null on an
   * install that switched the optional category off.
   */
  userEmailDomains: unknown
  /**
   * The same domains as the keys of `userEmailDomains`, flattened so an
   * operator can search for one. A JSON object's keys are the domains
   * themselves, and no query language addresses a key it does not already
   * know the name of.
   */
  userDomains: Generated<string[]>
  /**
   * What the last report carried, so the list and the drawer read one row:
   * the onboarding ladder, the usage counts and the operational block, under
   * the same keys the dictionary declares.
   */
  latestReport: unknown
  /**
   * What the last report said the operator had switched off, read from which
   * categories arrived rather than from a field the install could assert.
   */
  optionalMetricsReported: Generated<boolean>
  hostnameReported: Generated<boolean>
  /**
   * How many reports have landed, and how many fields the last one carried
   * that this release has no name for. A count above zero says the senders
   * are ahead of the receiver.
   */
  reportCount: Generated<number>
  lastUnknownFields: Generated<number>
  /**
   * The lead signals already raised for this install. A signal on this list
   * is never raised again, which is what keeps a large install from
   * reporting that it is large every morning.
   */
  raisedSignals: Generated<string[]>
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

/**
 * One upgrade diagnostics bundle a self-hosted install sent over Connect, on
 * LangWatch Cloud. The bytes live in object storage under
 * `upgrade-diagnostics/` and expire at 30 days by a lifecycle rule; this row
 * points at them and feeds Cloud admin and the Slack signal. Plain column
 * refs, no relation, as the usage report rows.
 */
export interface SelfHostedInstanceDiagnostic {
  id: string
  /** The `SelfHostedInstance.instanceId` the licence caller presented. */
  instanceId: string
  stepId: string
  stepDescription: string | null
  release: string | null
  attempt: number
  /** `operator` or `automatic`. */
  trigger: string
  byteLength: number
  objectKey: string
  receivedAt: Generated<Timestamp>
  /** When the lifecycle rule removes the bundle; the row is listed until then. */
  expiresAt: Timestamp
  /** Set on the report that posted this step and release's one Slack message. */
  slackAnnouncedAt: Timestamp | null
}

/**
 * A notice LangWatch staff send every self-hosted install in Cloud's reply to its daily usage
 * report (enterprise/modules/licensing/specs/self-hosted-notices.feature). Cloud only, and no
 * tenant: every install receives every published notice that has not expired.
 */
export interface SelfHostedNotice {
  id: string
  /** `info`, `warning` or `security`. */
  severity: string
  title: string
  body: string
  /** An https link to read more; null where the notice names none. */
  url: string | null
  /** The release that fixes what the notice is about; an install on it or newer hides it. */
  fixedIn: string | null
  /** `draft`, `published` or `withdrawn`. Only a published notice that has not expired is sent. */
  status: string
  expiresAt: Timestamp | null
  publishedAt: Timestamp | null
  /** The staff member who published it, by user id. */
  publishedBy: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

/**
 * One usage report, exactly as it arrived. Kept so a number on an install's
 * row can be traced back to the report that set it, and so a question asked
 * later can be answered over the history rather than over the latest snapshot.
 */
export interface SelfHostedInstanceReport {
  id: string
  /**
   * The `SelfHostedInstance.instanceId` this report came from. A plain column
   * ref for the same reason the instance row's refs are.
   */
  instanceId: string
  receivedAt: Generated<Timestamp>
  version: string | null
  reportSchemaVersion: number | null
  /** Fields this release had no name for. Counted, never stored by name. */
  unknownFields: Generated<number>
  /** The validated report, with every unknown field already dropped. */
  payload: unknown
}

/**
 * One browser a person subscribed to Web Push, owned by notification. A
 * person holds one row per device; the endpoint names the browser, so a
 * browser that subscribes again replaces its own row.
 */
export interface WebPushSubscription {
  id: string
  /**
   * The person the pushes are for. A plain column: the row leaves with the
   * person through notification's own subscriber on the user lifecycle.
   */
  userId: string
  /** The push service URL the browser handed out. Unique per browser. */
  endpoint: string
  /**
   * The browser's P-256 public key and auth secret, base64url, as the
   * browser reports them. Payloads are encrypted to these.
   */
  p256dh: string
  auth: string
  userAgent: string | null
  createdAt: Generated<Timestamp>
  /** When the push service last accepted a push for this browser. */
  lastSuccessAt: Timestamp | null
}

/**
 * This installation's VAPID key pair, generated on first use so Web Push works
 * without configuration. Fixed id, so the table holds one row: a second
 * insert loses the race on the primary key and reads the winner's pair.
 */
export interface WebPushVapidKey {
  id: Generated<string>
  /** The uncompressed P-256 public key, base64url. The browser subscribes with it. */
  publicKey: string
  /** The private key, base64url, encrypted with the process encryption key. */
  privateKeyEncrypted: string
  createdAt: Generated<Timestamp>
}

export interface SignInProviderMount {
  /**
   * Fixed, so this table holds one row: licensing's copy of sso's sign-in
   * provider fact (ADR-179).
   */
  id: Generated<string>
  /** The configured provider's name; null when the deployment signs in by email. */
  provider: string | null
  mounted: boolean
  /** The last sso event applied: the row's idempotency guard and its version. */
  lastAppliedEventId: string
  updatedAt: Generated<Timestamp>
}

/** An organization enrolled in the weekly digest. No row: not enrolled, so nothing sends. */
export interface DigestCohortOrganization {
  organizationId: string
  enrolledBy: string
  enrolledAt: Generated<Timestamp>
}

/** A person's opt-out of the digest. No row, or no `unsubscribedAt`: subscribed. */
export interface DigestSubscription {
  userId: string
  unsubscribedAt: Timestamp | null
  /** The send whose unsubscribe link or one-click header the person used. */
  unsubscribedFromSendId: string | null
  updatedAt: Timestamp
}

/** A project's ISO week, folded from its owners' events. The id is the fold's aggregate key. */
export interface DigestProjectWeek {
  id: string
  projectId: string
  organizationId: string
  isoWeek: string
  /** The fold state: counts, sums, a fixed-bucket latency histogram, capped model and evaluator maps. */
  state: unknown
  occurredAt: Timestamp
  lastEventId: string
  acceptedAt: Timestamp
  projectionVersion: string
  createdAt: Timestamp
  updatedAt: Timestamp
}

/** A person's ISO week of coding-agent activity in one organization, folded from the agents' facts. */
export interface DigestUserWeek {
  id: string
  userId: string
  organizationId: string
  isoWeek: string
  /** Sessions (a capped set of ids) and the agents' metric series, converged last-write-wins. */
  state: unknown
  occurredAt: Timestamp
  lastEventId: string
  acceptedAt: Timestamp
  projectionVersion: string
  createdAt: Timestamp
  updatedAt: Timestamp
}

/** One person's digest for one ISO week, accepted by our own mailer. The id is `userId:isoWeek`. */
export interface DigestSend {
  id: string
  userId: string
  organizationId: string
  isoWeek: string
  templateId: string
  sentAt: Timestamp
  createdAt: Generated<Timestamp>
}

/** The digest's one settings row. `armed` is the weekly send's master switch; it starts off. */
export interface DigestSettings {
  id: Generated<string>
  armed: Generated<boolean>
  lastPlannedWeek: string | null
  updatedAt: Timestamp
}

export interface DB {
  Account: Account
  AccountCredential: AccountCredential
  Session: Session
  UserErasure: UserErasure
  User: User
  VerificationToken: VerificationToken
  Identifier: Identifier
  IdentitySignInSwitch: IdentitySignInSwitch
  IdentifierReservation: IdentifierReservation
  IdentityProjectionCursor: IdentityProjectionCursor
  TwoFactor: TwoFactor
  SignInAttemptLock: SignInAttemptLock
  Passkey: Passkey
  MfaEnrollment: MfaEnrollment
  SsoConnection: SsoConnection
  SsoConnectionRegistrationSlot: SsoConnectionRegistrationSlot
  SsoConnectionReproofCursor: SsoConnectionReproofCursor
  SsoCredential: SsoCredential
  SsoVerifiedDomain: SsoVerifiedDomain
  SsoVerifiedDomainHolder: SsoVerifiedDomainHolder
  SsoAuthenticationActivity: SsoAuthenticationActivity
  SsoProvider: SsoProvider
  SsoBreakGlassBinding: SsoBreakGlassBinding
  SsoActivationRecoveryReservation: SsoActivationRecoveryReservation
  JoinRequest: JoinRequest
  TeamUser: TeamUser
  OrganizationUser: OrganizationUser
  SeatHold: SeatHold
  SeatRequest: SeatRequest
  OrganizationSeatLease: OrganizationSeatLease
  OrganizationAccessNotice: OrganizationAccessNotice
  Team: Team
  Organization: Organization
  IssuedLicense: IssuedLicense
  OrganizationLicense: OrganizationLicense
  ActivationCode: ActivationCode
  ConnectedBillingAccount: ConnectedBillingAccount
  ConnectedCreditGrant: ConnectedCreditGrant
  ConnectedInvoice: ConnectedInvoice
  ScimRequestLog: ScimRequestLog
  ScimToken: ScimToken
  ScimExternalId: ScimExternalId
  ScimUserResource: ScimUserResource
  ScimDirectoryUser: ScimDirectoryUser
  ScimSyncState: ScimSyncState
  ScimSsoConnectionView: ScimSsoConnectionView
  OrganizationSeatPeak: OrganizationSeatPeak
  OrganizationSeatPeakHolder: OrganizationSeatPeakHolder
  OrganizationSeatPeakMonth: OrganizationSeatPeakMonth
  Project: Project
  Department: Department
  DepartmentMembershipHistory: DepartmentMembershipHistory
  PlatformToolPolicy: PlatformToolPolicy
  OrganizationInvite: OrganizationInvite
  Monitor: Monitor
  Cost: Cost
  Topic: Topic
  Dataset: Dataset
  DatasetRecord: DatasetRecord
  DatasetVersion: DatasetVersion
  DatasetVersionRun: DatasetVersionRun
  Dashboard: Dashboard
  DashboardFavourite: DashboardFavourite
  DashboardsV2: DashboardsV2
  DashboardWidget: DashboardWidget
  InsightProjection: InsightProjection
  InsightReaderProjection: InsightReaderProjection
  InsightDailyScheduleProjection: InsightDailyScheduleProjection
  SavedView: SavedView
  CustomGraph: CustomGraph
  BatchEvaluation: BatchEvaluation
  Trigger: Trigger
  SlackIntegration: SlackIntegration
  slack_connection_claim: SlackConnectionClaim
  WebhookEndpoint: WebhookEndpoint
  WebhookEndpointDelivery: WebhookEndpointDelivery
  Experiment: Experiment
  ExperimentVersion: ExperimentVersion
  Annotation: Annotation
  ModelProvider: ModelProvider
  ModelProviderScope: ModelProviderScope
  ModelDefaultConfig: ModelDefaultConfig
  ModelDefaultConfigScope: ModelDefaultConfigScope
  GithubInstallation: GithubInstallation
  GithubPullRequest: GithubPullRequest
  GithubBranchPullRequestCheck: GithubBranchPullRequestCheck
  LangyConversationProjection: LangyConversationProjection
  LangyTurnRequest: LangyTurnRequest
  LangyActiveTurn: LangyActiveTurn
  LangyConversationTurnProjection: LangyConversationTurnProjection
  LangyPullRequest: LangyPullRequest
  LangyMessageProjection: LangyMessageProjection
  ProjectSecret: ProjectSecret
  TriggerSent: TriggerSent
  TriggerLatestEvaluation: TriggerLatestEvaluation
  EmailSuppression: EmailSuppression
  AnnotationScore: AnnotationScore
  AnnotationQueue: AnnotationQueue
  AnnotationQueueMembers: AnnotationQueueMembers
  AnnotationQueueScores: AnnotationQueueScores
  AnnotationQueueItem: AnnotationQueueItem
  ShareLink: ShareLink
  PinnedTrace: PinnedTrace
  TraceEditOverlay: TraceEditOverlay
  TraceIngestSourceBilling: TraceIngestSourceBilling
  RetentionPolicy: RetentionPolicy
  RetentionApplyStatus: RetentionApplyStatus
  DataPrivacyPolicy: DataPrivacyPolicy
  InstantEvalJudgeProject: InstantEvalJudgeProject
  InstantEvalJudgeUsageBilling: InstantEvalJudgeUsageBilling
  InstantEvalJudgeSpend: InstantEvalJudgeSpend
  CustomLLMModelCost: CustomLLMModelCost
  Workflow: Workflow
  WorkflowVersion: WorkflowVersion
  AuditLog: AuditLog
  LlmPromptConfig: LlmPromptConfig
  LlmPromptConfigVersion: LlmPromptConfigVersion
  PromptTagAssignment: PromptTagAssignment
  PromptTag: PromptTag
  Analytics: Analytics
  CustomRole: CustomRole
  Notification: Notification
  Agent: Agent
  AgentVersion: AgentVersion
  AgentVersionLabel: AgentVersionLabel
  Evaluator: Evaluator
  EvaluatorVersion: EvaluatorVersion
  EvaluatorVersionLabel: EvaluatorVersionLabel
  Scenario: Scenario
  ScenarioVersion: ScenarioVersion
  SimulationSuite: SimulationSuite
  Subscription: Subscription
  Invoice: Invoice
  InvoiceItem: InvoiceItem
  Group: Group
  GroupMembership: GroupMembership
  AuthzUserStanding: AuthzUserStanding
  AuthzEngineSwitch: AuthzEngineSwitch
  AuthzOrganizationPolicy: AuthzOrganizationPolicy
  AuthzBraidWrite: AuthzBraidWrite
  RoleBinding: RoleBinding
  Grant: Grant
  GrantUsage: GrantUsage
  Role: Role
  ApiKey: ApiKey
  SigningKey: SigningKey
  ProjectActiveDay: ProjectActiveDay
  NurturingProject: NurturingProject
  NurturingUncopiedProject: NurturingUncopiedProject
  NurturingOrganizationMember: NurturingOrganizationMember
  NurturingOrganizationCreated: NurturingOrganizationCreated
  NurturingOrganizationPlan: NurturingOrganizationPlan
  NurturingVirtualKey: NurturingVirtualKey
  NurturingErasedUser: NurturingErasedUser
  NurturingOrganization: NurturingOrganization
  NurturingProjectUsageView: NurturingProjectUsageView
  NurturingProjectTraceDay: NurturingProjectTraceDay
  NurturingProjectTraceDayCount: NurturingProjectTraceDayCount
  NurturingProjectEvaluationDay: NurturingProjectEvaluationDay
  NurturingPersonalProject: NurturingPersonalProject
  NurturingTracedProject: NurturingTracedProject
  BillingMeterCheckpoint: BillingMeterCheckpoint
  VirtualKey: VirtualKey
  VirtualKeyScope: VirtualKeyScope
  RoutingPolicy: RoutingPolicy
  RoutingPolicyScope: RoutingPolicyScope
  GatewayGuardrail: GatewayGuardrail
  IngestionSource: IngestionSource
  AnomalyRule: AnomalyRule
  AnomalyAlert: AnomalyAlert
  AiToolEntry: AiToolEntry
  AiToolEntryTeam: AiToolEntryTeam
  AiToolEntryDepartment: AiToolEntryDepartment
  IngestionTemplate: IngestionTemplate
  GatewayA2ATarget: GatewayA2ATarget
  GatewayConnectUpstream: GatewayConnectUpstream
  GatewayTraceExportKey: GatewayTraceExportKey
  GatewayBudget: GatewayBudget
  GatewayBudgetBucketBoundary: GatewayBudgetBucketBoundary
  GatewayBudgetLedger: GatewayBudgetLedger
  GatewayChangeEvent: GatewayChangeEvent
  GatewayCacheRule: GatewayCacheRule
  FeatureFlag: FeatureFlag
  FeatureFlagExperimentSetting: FeatureFlagExperimentSetting
  ProcessManagerInstance: ProcessManagerInstance
  ProcessManagerInbox: ProcessManagerInbox
  ProcessManagerOutbox: ProcessManagerOutbox
  ProcessManagerOutboxAttempt: ProcessManagerOutboxAttempt
  TopicClusteringRunProjection: TopicClusteringRunProjection
  TopicModelProjection: TopicModelProjection
  TopicClusteringRunHistoryProjection: TopicClusteringRunHistoryProjection
  IngestionPullRunProjection: IngestionPullRunProjection
  ScheduledJob: ScheduledJob
  BugReport: BugReport
  IdempotencyReceipt: IdempotencyReceipt
  GatewayRealtimeSession: GatewayRealtimeSession
  GatewayRealtimeSessionReport: GatewayRealtimeSessionReport
  SystemMigrationTenantState: SystemMigrationTenantState
  SystemMigrationEnrollment: SystemMigrationEnrollment
  StoredObject: StoredObject
  DiscoveredPerson: DiscoveredPerson
  DiscoveredAgent: DiscoveredAgent
  DiscoveredAgentListing: DiscoveredAgentListing
  AgentRegistryEntry: AgentRegistryEntry
  AgentRegistryEvent: AgentRegistryEvent
  AgentRegistryBaseline: AgentRegistryBaseline
  AgentRegistrySettings: AgentRegistrySettings
  AgentStopPlanTarget: AgentStopPlanTarget
  AgentStopPlan: AgentStopPlan
  GovernanceAwsConnection: GovernanceAwsConnection
  AgentStop: AgentStop
  AgentStopClaim: AgentStopClaim
  AgentStopTargetPriorState: AgentStopTargetPriorState
  IdentityMatch: IdentityMatch
  GovernanceTenantHistory: GovernanceTenantHistory
  ErasedIdentifierSuppression: ErasedIdentifierSuppression
  IdentityMatchSuggestion: IdentityMatchSuggestion
  ConnectedSeatChange: ConnectedSeatChange
  ConnectedStatement: ConnectedStatement
  LicenseRenewalClaim: LicenseRenewalClaim
  InstanceIdentity: InstanceIdentity
  SelfHostedInstance: SelfHostedInstance
  SelfHostedInstanceDiagnostic: SelfHostedInstanceDiagnostic
  SelfHostedNotice: SelfHostedNotice
  SelfHostedInstanceReport: SelfHostedInstanceReport
  WebPushSubscription: WebPushSubscription
  WebPushVapidKey: WebPushVapidKey
  SignInProviderMount: SignInProviderMount
  DigestCohortOrganization: DigestCohortOrganization
  DigestSubscription: DigestSubscription
  DigestProjectWeek: DigestProjectWeek
  DigestUserWeek: DigestUserWeek
  DigestSend: DigestSend
  DigestSettings: DigestSettings
}
