import {
  type PgTableExtraConfigValue,
  bigint,
  bigserial,
  boolean,
  doublePrecision,
  foreignKey,
  index,
  integer,
  jsonb,
  numeric,
  pgEnum,
  pgTable,
  primaryKey,
  text,
  timestamp,
  unique,
  uniqueIndex,
  varchar,
} from 'drizzle-orm/pg-core'
import { defineRelations, sql } from 'drizzle-orm'
import { nanoid } from 'nanoid'
import cuid from 'cuid'

const utcNow = (() => {
  let now: Date | undefined
  return () => {
    if (now === undefined) {
      now = new Date()
      queueMicrotask(() => {
        now = undefined
      })
    }
    return now
  }
})()

export const ssoConnectionMigrationPhaseEnum = pgEnum('SsoConnectionMigrationPhase', [
  'SETUP',
  'GRACE_LEGACY',
  'GRACE_DIRECT',
  'FINALIZING',
  'FINALIZED',
])

export const teamUserRoleEnum = pgEnum('TeamUserRole', ['ADMIN', 'MEMBER', 'VIEWER', 'CUSTOM'])

export const roleBindingScopeTypeEnum = pgEnum('RoleBindingScopeType', [
  'ORGANIZATION',
  'TEAM',
  'PROJECT',
])

export const grantScopeTypeEnum = pgEnum('GrantScopeType', [
  'ORGANIZATION',
  'TEAM',
  'PROJECT',
  'RESOURCE',
  'PLATFORM',
])

export const grantPrincipalTypeEnum = pgEnum('GrantPrincipalType', [
  'USER',
  'API_KEY',
  'GROUP',
  'TEAM',
  'ORGANIZATION',
  'PROJECT',
  'ANYONE',
])

export const organizationUserRoleEnum = pgEnum('OrganizationUserRole', [
  'ADMIN',
  'MEMBER',
  'EXTERNAL',
  'DEVELOPER',
])

/** Where a license in the registry came from. */
export const issuedLicenseSourceEnum = pgEnum('IssuedLicenseSource', [
  'BACKOFFICE',
  'PURCHASE',
  'SCRIPT',
  'LEGACY_IMPORT',
])

/** Why a credit was granted. The amount is always prepaid, never promotional. */
export const connectedCreditGrantKindEnum = pgEnum('ConnectedCreditGrantKind', [
  'commit',
  'added',
  'renewal',
])

/** What an invoice of a connected customer is for. */
export const connectedInvoiceKindEnum = pgEnum('ConnectedInvoiceKind', [
  'annual',
  'seat_change',
  'usage',
])

export const planTypesEnum = pgEnum('PlanTypes', [
  'FREE',
  'PRO',
  'GROWTH',
  'GROWTH_SEAT_EUR_MONTHLY',
  'GROWTH_SEAT_EUR_ANNUAL',
  'GROWTH_SEAT_USD_MONTHLY',
  'GROWTH_SEAT_USD_ANNUAL',
  'ENTERPRISE',
  'LAUNCH',
  'ACCELERATE',
  'LAUNCH_ANNUAL',
  'ACCELERATE_ANNUAL',
])

export const pricingModelEnum = pgEnum('PricingModel', ['TIERED', 'SEAT_EVENT'])

/** Organization signup intent (ADR-038): which product the org came for. */
export const organizationIntentEnum = pgEnum('OrganizationIntent', ['AGENT_GOVERNANCE', 'LLM_OPS'])

export const subscriptionStatusEnum = pgEnum('SubscriptionStatus', [
  'PENDING',
  'FAILED',
  'ACTIVE',
  'CANCELLED',
])

export const currencyEnum = pgEnum('Currency', ['USD', 'EUR'])

export const inviteStatusEnum = pgEnum('INVITE_STATUS', [
  'PENDING',
  'ACCEPTED',
  'WAITING_APPROVAL',
  'PAYMENT_PENDING',
  'REVOKED',
])

export const evaluationExecutionModeEnum = pgEnum('EvaluationExecutionMode', [
  'ON_MESSAGE',
  'AS_GUARDRAIL',
  'MANUALLY',
])

export const costTypeEnum = pgEnum('CostType', [
  'TRACE_CHECK',
  'GUARDRAIL',
  'CLUSTERING',
  'BATCH_EVALUATION',
])

export const costReferenceTypeEnum = pgEnum('CostReferenceType', [
  'CHECK',
  'TRACE',
  'PROJECT',
  'BATCH',
])

/** Who sees a dashboard, named as the prompt scope is, with one more level. */
export const dashboardScopeEnum = pgEnum('DashboardScope', ['PRIVATE', 'PROJECT', 'ORGANIZATION'])

export const triggerActionEnum = pgEnum('TriggerAction', [
  'SEND_EMAIL',
  'ADD_TO_DATASET',
  'ADD_TO_ANNOTATION_QUEUE',
  'SEND_SLACK_MESSAGE',
  'SEND_WEBHOOK',
])

export const triggerKindEnum = pgEnum('TriggerKind', ['AUTOMATION', 'ALERT', 'REPORT'])

export const alertTypeEnum = pgEnum('AlertType', ['CRITICAL', 'WARNING', 'INFO'])

/**
 * Who can use a Slack connection (ADR-093 §5a): every project in the
 * organization, or one project. Chosen per connection.
 */
export const slackIntegrationScopeTypeEnum = pgEnum('SlackIntegrationScopeType', [
  'ORGANIZATION',
  'PROJECT',
])

/**
 * How a Slack connection posts: a bot token over the Web API (any channel it
 * is invited to, every block renders) or an incoming webhook (one channel).
 */
export const slackIntegrationKindEnum = pgEnum('SlackIntegrationKind', ['BOT', 'INCOMING_WEBHOOK'])

export const webhookDeliveryOutcomeEnum = pgEnum('WebhookDeliveryOutcome', [
  'success',
  'retryable',
  'terminal',
  'pending',
])

/**
 * The last hop a webhook endpoint delivers over. Lowercase because these
 * values are the wire spelling: the same strings appear in the REST DTO,
 * both SDKs and the CLI.
 */
export const webhookDestinationKindEnum = pgEnum('WebhookDestinationKind', ['http', 'sqs'])

export const webhookEndpointStatusEnum = pgEnum('WebhookEndpointStatus', ['ACTIVE', 'DISABLED'])

/**
 * Which webhook pipeline wrote a delivery row, and therefore which columns
 * carry its tenancy: `platform` rows anchor on organizationId + endpointId,
 * `automations` rows on projectId + triggerId.
 */
export const webhookDeliveryChannelEnum = pgEnum('WebhookDeliveryChannel', [
  'platform',
  'automations',
])

export const experimentTypeEnum = pgEnum('ExperimentType', [
  'DSPY',
  'BATCH_EVALUATION',
  'BATCH_EVALUATION_V2',
  'EVALUATIONS_V3',
])

export const modelProviderScopeTypeEnum = pgEnum('ModelProviderScopeType', [
  'ORGANIZATION',
  'TEAM',
  'PROJECT',
])

export const modelDefaultScopeTypeEnum = pgEnum('ModelDefaultScopeType', [
  'ORGANIZATION',
  'TEAM',
  'PROJECT',
])

export const langyProjectionTitleSourceEnum = pgEnum('LangyProjectionTitleSource', [
  'derived',
  'auto',
  'user',
])

export const annotationScoreDataTypeEnum = pgEnum('AnnotationScoreDataType', [
  'CATEGORICAL',
  'BOOLEAN',
  'LIKERT',
  'OPTION',
  'CHECKBOX',
])

export const shareResourceTypeEnum = pgEnum('ShareResourceType', ['TRACE', 'THREAD'])

export const shareVisibilityEnum = pgEnum('ShareVisibility', ['PUBLIC', 'ORGANIZATION', 'PROJECT'])

export const pinSourceEnum = pgEnum('PinSource', ['manual', 'share'])

export const retentionPolicyScopeTypeEnum = pgEnum('RetentionPolicyScopeType', [
  'ORGANIZATION',
  'TEAM',
  'PROJECT',
])

export const dataPrivacyScopeTypeEnum = pgEnum('DataPrivacyScopeType', [
  'ORGANIZATION',
  'DEPARTMENT',
  'TEAM',
  'PROJECT',
])

export const customLlmmodelCostScopeTypeEnum = pgEnum('CustomLLMModelCostScopeType', [
  'ORGANIZATION',
  'TEAM',
  'PROJECT',
])

export const promptScopeEnum = pgEnum('PromptScope', ['ORGANIZATION', 'PROJECT'])

export const virtualKeyStatusEnum = pgEnum('VirtualKeyStatus', ['ACTIVE', 'DISABLED', 'REVOKED'])

export const virtualKeyPurposeEnum = pgEnum('VirtualKeyPurpose', ['USER', 'LANGY', 'CONNECT'])

export const gatewayProviderRotationPolicyEnum = pgEnum('GatewayProviderRotationPolicy', ['MANUAL'])

export const gatewayProviderHealthStatusEnum = pgEnum('GatewayProviderHealthStatus', [
  'UNKNOWN',
  'HEALTHY',
  'DEGRADED',
  'CIRCUIT_OPEN',
])

export const gatewayBudgetScopeTypeEnum = pgEnum('GatewayBudgetScopeType', [
  'ORGANIZATION',
  'TEAM',
  'PROJECT',
  'VIRTUAL_KEY',
  'PRINCIPAL',
  'GROUP',
  'ATTRIBUTED_USER',
])

export const gatewayBudgetWindowEnum = pgEnum('GatewayBudgetWindow', [
  'MINUTE',
  'HOUR',
  'DAY',
  'WEEK',
  'MONTH',
  'TOTAL',
  'MANUAL',
])

export const gatewayBudgetBreachActionEnum = pgEnum('GatewayBudgetBreachAction', ['BLOCK', 'WARN'])

export const gatewayBudgetLedgerStatusEnum = pgEnum('GatewayBudgetLedgerStatus', [
  'SUCCESS',
  'PROVIDER_ERROR',
  'BLOCKED_BY_GUARDRAIL',
  'CANCELLED',
])

export const gatewayChangeEventKindEnum = pgEnum('GatewayChangeEventKind', [
  'VK_CREATED',
  'VK_CONFIG_UPDATED',
  'VK_REVOKED',
  'VK_ROTATED',
  'VK_DISABLED',
  'VK_ENABLED',
  'BUDGET_CREATED',
  'BUDGET_UPDATED',
  'BUDGET_DELETED',
  'MODEL_PROVIDER_UPDATED',
  'CACHE_RULE_CREATED',
  'CACHE_RULE_UPDATED',
  'CACHE_RULE_DELETED',
  'ROUTING_POLICY_UPDATED',
  'ROUTING_POLICY_DELETED',
  'A2A_TARGET_CREATED',
  'A2A_TARGET_UPDATED',
  'A2A_TARGET_DELETED',
])

export const gatewayCacheRuleModeEnum = pgEnum('GatewayCacheRuleMode', [
  'RESPECT',
  'FORCE',
  'DISABLE',
])

export const virtualKeyRoutingModeEnum = pgEnum('VirtualKeyRoutingMode', [
  'NONE',
  'FALLBACK_ALL',
  'POLICY',
])

export const virtualKeyScopeTypeEnum = pgEnum('VirtualKeyScopeType', [
  'ORGANIZATION',
  'TEAM',
  'PROJECT',
])

export const routingPolicyScopeTypeEnum = pgEnum('RoutingPolicyScopeType', [
  'ORGANIZATION',
  'TEAM',
  'PROJECT',
])

export const gatewayGuardrailDirectionEnum = pgEnum('GatewayGuardrailDirection', [
  'PRE',
  'POST',
  'STREAM_CHUNK',
])

export const gatewayGuardrailFailureModeEnum = pgEnum('GatewayGuardrailFailureMode', [
  'FAIL_OPEN',
  'FAIL_CLOSED',
])

export const featureFlagExperimentSubjectTypeEnum = pgEnum('FeatureFlagExperimentSubjectType', [
  'USER',
  'ORGANIZATION',
  'PROJECT',
])

export const processManagerOutboxStatusEnum = pgEnum('ProcessManagerOutboxStatus', [
  'pending',
  'dispatched',
  'dead',
  'discarded',
])

export const processManagerOutboxAttemptOutcomeEnum = pgEnum('ProcessManagerOutboxAttemptOutcome', [
  'retry_scheduled',
  'dead',
])

export const gatewayRealtimeSessionStatusEnum = pgEnum('GatewayRealtimeSessionStatus', [
  'OPEN',
  'CLOSED',
  'FAILED',
  'EXPIRED',
])

/** How the invoicing of one seat change ended. */
export const connectedSeatChangeStateEnum = pgEnum('ConnectedSeatChangeState', [
  'intent',
  'invoiced',
  'nothing_to_invoice',
  'not_onboarded',
])

export const account = pgTable(
  'Account',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    userId: text('userId').notNull(),
    type: text('type').notNull().default('oauth'),
    provider: text('provider').notNull(),
    /**
     * WHO asserted `providerAccountId`. Better Auth keys accounts by
     * `(provider, providerAccountId)`; auth's account create hook writes this
     * (modules/auth/specs/account-issuer.feature). Nullable only so the column
     * could be added to a populated table; the migration backfilled every row.
     */
    issuer: text('issuer'),
    providerAccountId: text('providerAccountId').notNull(),
    refresh_token: text('refresh_token'),
    access_token: text('access_token'),
    expires_at: timestamp('expires_at', { precision: 3 }),
    ext_expires_in: integer('ext_expires_in'),
    token_type: text('token_type'),
    scope: text('scope'),
    id_token: text('id_token'),
    session_state: text('session_state'),
    password: text('password'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('Account_userId_idx').on(table.userId),
    index('Account_issuer_providerAccountId_idx').on(table.issuer, table.providerAccountId),
    uniqueIndex('Account_provider_providerAccountId_key').on(
      table.provider,
      table.providerAccountId,
    ),
    foreignKey({ name: 'Account_userId_fkey', columns: [table.userId], foreignColumns: [user.id] })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const accountCredential = pgTable(
  'AccountCredential',
  {
    id: text('id').primaryKey(),
    userId: text('userId').notNull(),
    provider: text('provider').notNull(),
    password: text('password'),
    accessToken: text('accessToken'),
    refreshToken: text('refreshToken'),
    idToken: text('idToken'),
    accessTokenExpiresAt: timestamp('accessTokenExpiresAt', { precision: 3 }),
    refreshTokenExpiresAt: timestamp('refreshTokenExpiresAt', { precision: 3 }),
    scope: text('scope'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('AccountCredential_userId_idx').on(table.userId),
    foreignKey({
      name: 'AccountCredential_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const session = pgTable(
  'Session',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    sessionToken: text('sessionToken').notNull(),
    userId: text('userId').notNull(),
    expires: timestamp('expires', { precision: 3 }).notNull(),
    ipAddress: text('ipAddress'),
    userAgent: text('userAgent'),
    /**
     * Which of the user's sign-in methods minted this session (D06). NULL on
     * every session that predates the column and on any mint we could not
     * attribute - an ordinary value, not an error. Nothing reads it to decide
     * whether the session is valid, so no session was ended to add it.
     */
    identifierId: text('identifierId'),
    /**
     * What the sign-in that minted this session PROVED, as RFC 8176 method
     * references (D06). Empty on every session minted before the column, and
     * on any ordinary password sign-in. An organization that requires a second
     * factor reads this when a member reaches its data; it is never read at
     * mint and never ends a session.
     */
    amr: text('amr').array().default([]),
    /**
     * Minted while the account had never signed in: the session its sign-up
     * started, which skips the secure-account offer. False on every later one
     * and on every session that predates the column (specs/identity/passkeys.feature).
     */
    startedBySignUp: boolean('startedBySignUp').notNull().default(false),
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
    impersonating: jsonb('impersonating'),
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
    actorUserId: text('actorUserId'),
    subjectUserId: text('subjectUserId'),
    /**
     * Why the operator started it. Required to start one
     * (specs/features/backoffice-user-impersonation-reason.feature) and kept
     * beside both people rather than only in the audit log, so a request under
     * impersonation carries its own justification.
     */
    impersonationReason: text('impersonationReason'),
    /**
     * When the borrowed access lapses. A session outlives its impersonation
     * window: past this the claims are ignored and the operator is themselves
     * again, with nothing ended.
     */
    impersonationExpiresAt: timestamp('impersonationExpiresAt', { precision: 3 }),
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
    lastSeenAt: timestamp('lastSeenAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('Session_sessionToken_key').on(table.sessionToken),
    index('Session_userId_idx').on(table.userId),
    index('Session_identifierId_idx').on(table.identifierId),
    foreignKey({ name: 'Session_userId_fkey', columns: [table.userId], foreignColumns: [user.id] })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * One row per erased account, holding only its creation time: no id, name or address. It keeps
 * the install's first-created account knowable after an erasure deletes that row.
 * Spec: modules/user/specs/user.feature
 */
export const userErasure = pgTable(
  'UserErasure',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    accountCreatedAt: timestamp('accountCreatedAt', { precision: 3 }).notNull(),
    erasedAt: timestamp('erasedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [index('UserErasure_accountCreatedAt_idx').on(table.accountCreatedAt)],
)

export const user = pgTable(
  'User',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    name: text('name'),
    email: text('email'),
    emailVerified: boolean('emailVerified').notNull().default(false),
    signupConfirmationPending: boolean('signupConfirmationPending').notNull().default(false),
    passkeySignupClaimHash: text('passkeySignupClaimHash'),
    image: text('image'),
    pendingSsoSetup: boolean('pendingSsoSetup').notNull().default(false),
    userHashKey: text('userHashKey'),
    /**
     * Owned by the two-factor plugin (D06), which is why it is a plain
     * boolean rather than a state: the plugin flips it when a setup is
     * confirmed and when it is turned off, and challenges every sign-in while
     * it is true. The lifecycle - and the history - lives on MfaEnrollment.
     */
    twoFactorEnabled: boolean('twoFactorEnabled').notNull().default(false),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    lastLoginAt: timestamp('lastLoginAt', { precision: 3 }),
    deactivatedAt: timestamp('deactivatedAt', { precision: 3 }),
    lastHomePath: text('lastHomePath'),
    tracesExplorerTourDismissedAt: timestamp('tracesExplorerTourDismissedAt', { precision: 3 }),
    /**
     * How Langy should reach this person's code when a task needs a change
     * (ADR-129). "github" means they chose the pull request path and asked to
     * be remembered; NULL means ask. A shared local folder is never stored
     * here, because a folder has to be shared again each time.
     * Spec: specs/langy/langy-code-access.feature
     */
    langyCodeAccessPreference: text('langyCodeAccessPreference'),
    /**
     * When this person last said "not now" to being offered a passkey
     * (ADR-120). On the account rather than in browser storage for two
     * reasons: a per-browser record forgets on a new device, which is exactly
     * where the offer should be MORE eager, and an interval nobody can enforce
     * is not an interval. NULL means never asked.
     */
    passkeyNudgeDismissedAt: timestamp('passkeyNudgeDismissedAt', { precision: 3 }),
    /**
     * The domains this person has said "no thanks" to being offered their
     * colleagues on (D12). Per domain rather than one flag: somebody who
     * changes jobs and verifies a second address should be offered their new
     * employer's organization, not silenced forever by a decision they made
     * about the old one. Empty means never dismissed anything.
     */
    joinOfferDismissedDomains: text('joinOfferDismissedDomains').array().default([]),
    /**
     * This person's answer about browser notifications, one entry per topic:
     * { "langy": "enabled" | "declined" }. A topic that is absent was never
     * answered. The browser's own permission is not stored here.
     * Spec: specs/langy/langy-notifications.feature
     */
    notificationPreferences: jsonb('notificationPreferences').notNull().default({}),
    /**
     * When an operator marked this person's password as no longer trusted. A
     * password sign-in opens no session while it is set; any write of a new
     * password hash clears it. Spec: modules/auth/specs/forced-password-reset.feature
     */
    passwordResetRequiredAt: timestamp('passwordResetRequiredAt', { precision: 3 }),
  },
  (table) => [uniqueIndex('User_email_key').on(table.email)],
)

export const verificationToken = pgTable(
  'VerificationToken',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    identifier: text('identifier').notNull(),
    /** Not unique alone: better-auth stores plain values here (a user id, a counter). See WEB-19201. */
    token: text('token').notNull(),
    expires: timestamp('expires', { precision: 3 }).notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('VerificationToken_identifier_token_key').on(table.identifier, table.token),
  ],
)

export const identifier = pgTable(
  'Identifier',
  {
    id: text('id').primaryKey(),
    userId: text('userId').notNull(),
    provider: text('provider').notNull(),
    value: text('value'),
    domain: text('domain'),
    identifierHash: text('identifierHash'),
    accountId: text('accountId'),
    providerId: text('providerId'),
    issuer: text('issuer'),
    providerAccountId: text('providerAccountId'),
    state: text('state').notNull(),
    connectionId: text('connectionId'),
    verifiedAt: timestamp('verifiedAt', { precision: 3 }),
    attachedAt: timestamp('attachedAt', { precision: 3 }).notNull(),
    detachedAt: timestamp('detachedAt', { precision: 3 }),
    lastUsedAt: timestamp('lastUsedAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('Identifier_userId_idx').on(table.userId),
    index('Identifier_identifierHash_idx').on(table.identifierHash),
    index('Identifier_value_idx').on(table.value),
    index('Identifier_domain_idx').on(table.domain),
    index('Identifier_provider_providerAccountId_idx').on(table.provider, table.providerAccountId),
    index('Identifier_issuer_providerAccountId_idx').on(table.issuer, table.providerAccountId),
    index('Identifier_accountId_idx').on(table.accountId),
  ],
)

/**
 * Whether a user signs in on the identifier path: identity's own switch, set when the
 * identifier backfill proves the user and pinned off by an operator's rollback. No row is the
 * legacy path. Spec: modules/identity/specs/sign-in-switch.feature. Column refs only, no FK.
 */
export const identitySignInSwitch = pgTable('IdentitySignInSwitch', {
  userId: text('userId').primaryKey(),
  onIdentifiers: boolean('onIdentifiers').notNull(),
  pinned: boolean('pinned').notNull().default(false),
  switchedAt: timestamp('switchedAt', { precision: 3 }).notNull(),
  lastAppliedEventId: text('lastAppliedEventId'),
  updatedAt: timestamp('updatedAt', { precision: 3 })
    .notNull()
    .default(sql`CURRENT_TIMESTAMP`)
    .$defaultFn(utcNow)
    .$onUpdate(utcNow),
})

export const identifierReservation = pgTable(
  'IdentifierReservation',
  {
    normalizedValue: text('normalizedValue').primaryKey(),
    userId: text('userId').notNull(),
    identifierId: text('identifierId').notNull(),
    commandId: text('commandId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('IdentifierReservation_userId_idx').on(table.userId),
    index('IdentifierReservation_identifierId_idx').on(table.identifierId),
  ],
)

export const identityProjectionCursor = pgTable('IdentityProjectionCursor', {
  userId: text('userId').primaryKey(),
  lastEventId: text('lastEventId').notNull(),
  acceptedAt: timestamp('acceptedAt', { precision: 3 }).notNull(),
  occurredAt: timestamp('occurredAt', { precision: 3 }).notNull(),
  projectionVersion: text('projectionVersion').notNull(),
  createdAt: timestamp('createdAt', { precision: 3 })
    .notNull()
    .default(sql`CURRENT_TIMESTAMP`)
    .$defaultFn(utcNow),
  updatedAt: timestamp('updatedAt', { precision: 3 })
    .notNull()
    .default(sql`CURRENT_TIMESTAMP`)
    .$defaultFn(utcNow)
    .$onUpdate(utcNow),
})

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
export const twoFactor = pgTable(
  'TwoFactor',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    userId: text('userId').notNull(),
    secret: text('secret').notNull(),
    backupCodes: text('backupCodes').notNull(),
    verified: boolean('verified').notNull().default(true),
    failedVerificationCount: integer('failedVerificationCount').notNull().default(0),
    lockedUntil: timestamp('lockedUntil', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('TwoFactor_userId_idx').on(table.userId),
    index('TwoFactor_secret_idx').on(table.secret),
    foreignKey({
      name: 'TwoFactor_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const signInAttemptLock = pgTable(
  'SignInAttemptLock',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    /** Keyed hash of the normalised address. Never the address itself. */
    identifierHash: text('identifierHash').notNull(),
    /**
     * Consecutive failures since the last success. Reset by a success, not by
     * time - "five consecutive", not five ever.
     */
    failedCount: integer('failedCount').notNull().default(0),
    /** When the current lock lifts. Null when not locked. */
    lockedUntil: timestamp('lockedUntil', { precision: 3 }),
    /**
     * Locks taken since the last success. At five this stops lifting by
     * itself and becomes an incident - see `heldForReview`.
     */
    consecutiveLockouts: integer('consecutiveLockouts').notNull().default(0),
    /**
     * Held past its lock because the count of locks reached the ceiling the
     * control names. Cleared by an administrator, or by the owner proving
     * mailbox control through a password reset - the second is what keeps a
     * held account from being a denial of service anyone can inflict by
     * typing a colleague's address wrong often enough.
     */
    heldForReview: boolean('heldForReview').notNull().default(false),
    /**
     * Whose account this resolved to at the time, when it resolved to one.
     * Null for an address with no account - those rows are the reason this is
     * nullable, and the reason this table is not hung off User.
     */
    userId: text('userId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('SignInAttemptLock_identifierHash_key').on(table.identifierHash),
    index('SignInAttemptLock_lockedUntil_idx').on(table.lockedUntil),
    index('SignInAttemptLock_userId_idx').on(table.userId),
  ],
)

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
export const passkey = pgTable(
  'Passkey',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    name: text('name'),
    publicKey: text('publicKey').notNull(),
    userId: text('userId').notNull(),
    credentialID: text('credentialID').notNull(),
    counter: integer('counter').notNull().default(0),
    deviceType: text('deviceType').notNull(),
    backedUp: boolean('backedUp').notNull().default(false),
    transports: text('transports'),
    aaguid: text('aaguid'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('Passkey_userId_idx').on(table.userId),
    index('Passkey_credentialID_idx').on(table.credentialID),
    foreignKey({ name: 'Passkey_userId_fkey', columns: [table.userId], foreignColumns: [user.id] })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const mfaEnrollment = pgTable(
  'MfaEnrollment',
  {
    /**
     * The user. Also the aggregate id and the tenant - one identity history
     * per person (ADR-029 §4).
     */
    userId: text('userId').primaryKey(),
    /** The current enrollment, or NULL for somebody who never started one. */
    enrollmentId: text('enrollmentId'),
    /** totp. Never sms. */
    method: text('method'),
    /** NONE | PENDING | ENABLED | EXPIRED | DISABLED */
    state: text('state').notNull(),
    enrolledAt: timestamp('enrolledAt', { precision: 3 }),
    confirmedAt: timestamp('confirmedAt', { precision: 3 }),
    expiredAt: timestamp('expiredAt', { precision: 3 }),
    disabledAt: timestamp('disabledAt', { precision: 3 }),
    /** password+totp | org-admin - who ended it, and on whose authority. */
    disabledVia: text('disabledVia'),
    backupCodeCount: integer('backupCodeCount').notNull().default(0),
    /** Positions already spent, ascending. Never codes. */
    consumedBackupCodeIndexes: integer('consumedBackupCodeIndexes').array().default([]),
    failedCount: integer('failedCount').notNull().default(0),
    occurredAt: timestamp('occurredAt', { precision: 3 }).notNull(),
    lastEventId: text('lastEventId').notNull(),
    acceptedAt: timestamp('acceptedAt', { precision: 3 }).notNull(),
    projectionVersion: text('projectionVersion').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [index('MfaEnrollment_state_idx').on(table.state)],
)

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
export const ssoConnection = pgTable(
  'SsoConnection',
  {
    id: text('id').primaryKey(),
    organizationId: text('organizationId').notNull(),
    /** oidc | saml */
    type: text('type').notNull(),
    /**
     * DRAFT | CLAIMED | APPROVED | REJECTED | DISCARDED |
     * VERIFICATION_PENDING | VERIFIED | ACTIVE | SUSPENDED |
     * TEARDOWN_PENDING | TORN_DOWN
     */
    state: text('state').notNull(),
    claimedDomains: text('claimedDomains').array(),
    /**
     * [{ domain, state, claimedAtMs, claimedByActorId, decidedAtMs,
     * decidedByActorId, authority, waitedMs, note }] - every claim this
     * connection has made and how long each waited to be decided. The tier-3
     * approval queue is a read of this column, and the epic's Open Q2 wants
     * that latency measured from the day the queue exists.
     */
    domainClaims: jsonb('domainClaims').notNull().default([]),
    approvedDomains: text('approvedDomains').array(),
    /** The only domains that ever route, and only while state is ACTIVE. */
    verifiedDomains: text('verifiedDomains').array(),
    /**
     * The subset of `verifiedDomains` whose published record stayed missing
     * through its grace window (ADR-123). They still ROUTE - everyone already
     * here signs in unchanged - and they stop vouching for anybody NEW: no
     * provisioning on first sign-in, no walking in by domain. A column rather
     * than a read of the JSON below, because both questions are asked on a
     * sign-in path and have to be one indexed predicate.
     */
    lapsedDomains: text('lapsedDomains').array(),
    /**
     * [{ domain, method, actorId, verifiedAtMs, proofState, firstAbsentAtMs,
     * graceEndsAtMs }] - what proved each verified domain, who, and whether
     * that evidence is still there. Kept on the row rather than only in the
     * log so that an operator-attested domain can never be read as one the
     * customer proved, and a wavering one never as one nothing is doubting.
     */
    domainVerifications: jsonb('domainVerifications').notNull().default([]),
    /** { domain, method, tokenHash } while a ceremony is in flight. */
    pendingVerification: jsonb('pendingVerification'),
    idpMetadata: jsonb('idpMetadata').notNull(),
    /**
     * admit | request | refuse - what this connection does with somebody who
     * signs in through it and is not a member yet (ADR-117 3). Stated at
     * registration and changed by the setup journey, so it is never absent
     * and no reader has to decide what absence means. This was two fields
     * for a while - a boolean `allowsJit` beside it, kept in step by the
     * fold - and they disagreed about the middle answer.
     */
    arrivalPolicy: text('arrivalPolicy').notNull().default('refuse'),
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
    allowsJit: boolean('allowsJit').notNull().default(false),
    /**
     * When somebody CHOSE the policy above, and null while the registration
     * default stands. A different fact from the policy rather than a second
     * copy of it: a connection always has a behaviour, and separately
     * somebody has or has not decided on it. Going live waits for the
     * deciding - "turn everybody away" is a decision too.
     */
    arrivalPolicyDecidedAt: timestamp('arrivalPolicyDecidedAt', { precision: 3 }),
    /** self-serve | legacy-grandfathered */
    source: text('source').notNull(),
    testLoginAccountId: text('testLoginAccountId'),
    /**
     * The grandfathered connection this direct connection replaces. Null for
     * ordinary connections and for the grandfathered predecessor itself.
     *
     * Partial uniqueness for nonterminal replacements is installed by the
     * flattened migration; Prisma cannot express that predicate here.
     */
    replacesConnectionId: text('replacesConnectionId'),
    /**
     * SETUP | GRACE_LEGACY | GRACE_DIRECT | FINALIZING | FINALIZED.
     * Null means this connection is not a legacy-to-direct replacement.
     */
    migrationPhase: ssoConnectionMigrationPhaseEnum('migrationPhase'),
    graceStartedAt: timestamp('graceStartedAt', { precision: 3 }),
    /**
     * When an administrator last selected the normal route. This is audit
     * evidence only: routing is derived from `migrationPhase`, never recency.
     */
    routeChangedAt: timestamp('routeChangedAt', { precision: 3 }),
    finalizationRequestedAt: timestamp('finalizationRequestedAt', { precision: 3 }),
    finalizedAt: timestamp('finalizedAt', { precision: 3 }),
    /** { domain, note } - the last ops rejection, read back on a re-claim. */
    rejection: jsonb('rejection'),
    createdBy: text('createdBy'),
    /** While TEARDOWN_PENDING: when the process manager's wake completes it. */
    tearDownAfter: timestamp('tearDownAfter', { precision: 3 }),
    occurredAt: timestamp('occurredAt', { precision: 3 }).notNull(),
    lastEventId: text('lastEventId').notNull(),
    acceptedAt: timestamp('acceptedAt', { precision: 3 }).notNull(),
    projectionVersion: text('projectionVersion').notNull(),
    /**
     * Event-derived, not database-managed: `@default(now())` / `@updatedAt`
     * would make a replayed row differ from the row it rebuilds, which is
     * exactly the parity this projection has to keep.
     */
    createdAt: timestamp('createdAt', { precision: 3 }).notNull(),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull(),
  },
  (table) => [
    index('SsoConnection_organizationId_idx').on(table.organizationId),
    index('SsoConnection_organizationId_migrationPhase_idx').on(
      table.organizationId,
      table.migrationPhase,
    ),
    index('SsoConnection_state_idx').on(table.state),
    index('SsoConnection_verifiedDomains_idx').on(table.verifiedDomains),
    index('SsoConnection_lapsedDomains_idx').on(table.lapsedDomains),
  ],
)

/**
 * The one-statement pre-event registration lock. `SsoConnection` remains
 * event truth; this row only decides which command may append that truth.
 * A slot whose projected connection is terminal may be atomically replaced.
 */
export const ssoConnectionRegistrationSlot = pgTable(
  'SsoConnectionRegistrationSlot',
  {
    organizationId: text('organizationId').notNull(),
    /** legacy | direct */
    kind: text('kind').notNull(),
    connectionId: text('connectionId').notNull(),
    replacesConnectionId: text('replacesConnectionId'),
    commandId: text('commandId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    primaryKey({
      name: 'SsoConnectionRegistrationSlot_pkey',
      columns: [table.organizationId, table.kind],
    }),
    uniqueIndex('SsoConnectionRegistrationSlot_connectionId_key').on(table.connectionId),
  ],
)

/**
 * Operational round-robin position for ADR-123's domain re-proof sweep.
 *
 * This is deliberately not a column on `SsoConnection`: that row is a
 * whole-row event projection and its fold is its only writer, while a
 * healthy re-read states no fact. Keeping the scheduling cursor beside the
 * projection preserves both rules and still lets a fixed-size sweep revisit
 * every connection, including one whose proof has started wavering.
 */
export const ssoConnectionReproofCursor = pgTable(
  'SsoConnectionReproofCursor',
  {
    connectionId: text('connectionId').primaryKey(),
    lastReproofAt: timestamp('lastReproofAt', { precision: 3 }).notNull(),
  },
  (table) => [
    index('SsoConnectionReproofCursor_lastReproofAt_idx').on(table.lastReproofAt),
    foreignKey({
      name: 'SsoConnectionReproofCursor_connectionId_fkey',
      columns: [table.connectionId],
      foreignColumns: [ssoConnection.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const ssoCredential = pgTable(
  'SsoCredential',
  {
    id: text('id').primaryKey(),
    organizationId: text('organizationId').notNull(),
    /**
     * The connection whose registration minted this. Not a relation, matching
     * `SsoConnection` itself: a projection rebuilt from events must not depend
     * on relational state the log does not carry.
     */
    connectionId: text('connectionId').notNull(),
    /** oidc-client-id | oidc-client-secret | saml-idp-config */
    kind: text('kind').notNull(),
    ciphertext: text('ciphertext').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('SsoCredential_organizationId_connectionId_idx').on(
      table.organizationId,
      table.connectionId,
    ),
  ],
)

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
export const ssoVerifiedDomain = pgTable(
  'SsoVerifiedDomain',
  { domain: text('domain').primaryKey(), organizationId: text('organizationId').notNull() },
  (table) => [
    index('SsoVerifiedDomain_organizationId_idx').on(table.organizationId),
    unique('SsoVerifiedDomain_domain_organizationId_key').on(table.domain, table.organizationId),
  ],
)

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
export const ssoVerifiedDomainHolder = pgTable(
  'SsoVerifiedDomainHolder',
  {
    domain: text('domain').notNull(),
    connectionId: text('connectionId').notNull(),
    organizationId: text('organizationId').notNull(),
  },
  (table) => [
    primaryKey({
      name: 'SsoVerifiedDomainHolder_pkey',
      columns: [table.domain, table.connectionId],
    }),
    index('SsoVerifiedDomainHolder_connectionId_idx').on(table.connectionId),
    index('SsoVerifiedDomainHolder_organizationId_idx').on(table.organizationId),
    foreignKey({
      name: 'SsoVerifiedDomainHolder_domain_organizationId_fkey',
      columns: [table.domain, table.organizationId],
      foreignColumns: [ssoVerifiedDomain.domain, ssoVerifiedDomain.organizationId],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * A successful callback through one SSO connection. Identifier.lastUsedAt
 * also advances during routing and resolution, so it cannot prove that a
 * member actually authenticated through the replacement before cutover.
 */
export const ssoAuthenticationActivity = pgTable(
  'SsoAuthenticationActivity',
  {
    id: text('id').primaryKey(),
    organizationId: text('organizationId').notNull(),
    connectionId: text('connectionId').notNull(),
    userId: text('userId').notNull(),
    authenticatedAt: timestamp('authenticatedAt', { precision: 3 }).notNull(),
    /** The provider-side account the assertion named; null before #7633's arrival door recorded it. */
    providerAccountId: text('providerAccountId'),
  },
  (table) => [
    index('SsoAuthenticationActivity_organizationId_connectionId_authe_idx').on(
      table.organizationId,
      table.connectionId,
      table.authenticatedAt,
    ),
    index('SsoAuthenticationActivity_organizationId_userId_idx').on(
      table.organizationId,
      table.userId,
    ),
  ],
)

export const ssoProvider = pgTable(
  'SsoProvider',
  {
    id: text('id').primaryKey(),
    issuer: text('issuer').notNull(),
    oidcConfig: text('oidcConfig'),
    samlConfig: text('samlConfig'),
    userId: text('userId'),
    providerId: text('providerId').notNull(),
    organizationId: text('organizationId'),
    domain: text('domain').notNull(),
  },
  (table) => [
    uniqueIndex('SsoProvider_providerId_key').on(table.providerId),
    index('SsoProvider_organizationId_idx').on(table.organizationId),
    index('SsoProvider_domain_idx').on(table.domain),
  ],
)

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
export const ssoBreakGlassBinding = pgTable(
  'SsoBreakGlassBinding',
  {
    id: text('id').primaryKey(),
    organizationId: text('organizationId').notNull(),
    /** Who holds the way in. */
    userId: text('userId').notNull(),
    /** Who granted it. A way back in is never self-served. */
    grantedByUserId: text('grantedByUserId').notNull(),
    grantedAt: timestamp('grantedAt', { precision: 3 }).notNull(),
    expiresAt: timestamp('expiresAt', { precision: 3 }).notNull(),
    /** Set when a renewal replaced this row; the row itself is never edited. */
    supersededAt: timestamp('supersededAt', { precision: 3 }),
    /** The binding this one renewed, so the previous end date stays readable. */
    renewedFromId: text('renewedFromId'),
    /**
     * Which of the fourteen, seven and one day warnings have been sent, so a
     * sweep that runs twice in a day never says the same thing twice.
     */
    warnedDays: integer('warnedDays').array(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('SsoBreakGlassBinding_organizationId_expiresAt_idx').on(
      table.organizationId,
      table.expiresAt,
    ),
    index('SsoBreakGlassBinding_userId_idx').on(table.userId),
  ],
)

/**
 * A durable claim on one live recovery path while activation is between its
 * guard and projected ACTIVE fact. It has no expiry: only projection of the
 * matching activation, or an explicit cancellation, may consume it.
 */
export const ssoActivationRecoveryReservation = pgTable(
  'SsoActivationRecoveryReservation',
  {
    commandId: text('commandId').primaryKey(),
    organizationId: text('organizationId').notNull(),
    connectionId: text('connectionId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 }).notNull(),
  },
  (table) => [
    index('SsoActivationRecoveryReservation_organizationId_idx').on(table.organizationId),
    uniqueIndex('SsoActivationRecoveryReservation_organizationId_connectionI_key').on(
      table.organizationId,
      table.connectionId,
    ),
  ],
)

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
export const joinRequest = pgTable(
  'JoinRequest',
  {
    id: text('id').primaryKey(),
    userId: text('userId').notNull(),
    organizationId: text('organizationId').notNull(),
    /** The normalized domain the match was made on. */
    domain: text('domain').notNull(),
    /** PENDING | APPROVED | REJECTED | EXPIRED | WITHDRAWN */
    state: text('state').notNull(),
    /** How the domain was matched: `verified-identifier-domain` or `sso-connection-domain`. */
    matchedVia: text('matchedVia').notNull(),
    /**
     * Where the request was made: `web` or `cli` (ADR-171 v6). A request from
     * the terminal lands a Developer seat whatever the joiner seat says.
     */
    origin: text('origin').notNull().default('web'),
    /** While PENDING: when the day-14 wake lapses it. Null once resolved. */
    expiresAt: timestamp('expiresAt', { precision: 3 }),
    resolvedAt: timestamp('resolvedAt', { precision: 3 }),
    /**
     * user | policy | invite - who ended it. `policy` is domain auto-join's
     * own principal, deliberately not the one SSO auto-join uses.
     */
    resolvedByType: text('resolvedByType'),
    /** The admin's user id, `domain-auto`, or the invitation id. */
    resolvedById: text('resolvedById'),
    /** user | invite-accepted, on a withdrawal. */
    withdrawalCause: text('withdrawalCause'),
    /** The SSO connection a single sign-on arrival was admitted through; null otherwise. */
    connectionId: text('connectionId'),
    occurredAt: timestamp('occurredAt', { precision: 3 }).notNull(),
    lastEventId: text('lastEventId').notNull(),
    acceptedAt: timestamp('acceptedAt', { precision: 3 }).notNull(),
    projectionVersion: text('projectionVersion').notNull(),
    /**
     * Event-derived, not database-managed: `@default(now())` / `@updatedAt`
     * would make a replayed row differ from the row it rebuilds, which is
     * exactly the parity this projection has to keep.
     */
    createdAt: timestamp('createdAt', { precision: 3 }).notNull(),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull(),
  },
  (table) => [
    index('JoinRequest_organizationId_state_idx').on(table.organizationId, table.state),
    index('JoinRequest_userId_state_idx').on(table.userId, table.state),
    index('JoinRequest_userId_organizationId_idx').on(table.userId, table.organizationId),
    index('JoinRequest_domain_idx').on(table.domain),
  ],
)

export const teamUser = pgTable(
  'TeamUser',
  {
    userId: text('userId').notNull(),
    teamId: text('teamId').notNull(),
    role: teamUserRoleEnum('role').notNull(),
    assignedRoleId: text('assignedRoleId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    primaryKey({ name: 'TeamUser_pkey', columns: [table.userId, table.teamId] }),
    index('TeamUser_teamId_idx').on(table.teamId),
    index('TeamUser_userId_idx').on(table.userId),
    index('TeamUser_assignedRoleId_idx').on(table.assignedRoleId),
    foreignKey({
      name: 'TeamUser_assignedRoleId_fkey',
      columns: [table.assignedRoleId],
      foreignColumns: [customRole.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({ name: 'TeamUser_userId_fkey', columns: [table.userId], foreignColumns: [user.id] })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({ name: 'TeamUser_teamId_fkey', columns: [table.teamId], foreignColumns: [team.id] })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

export const organizationUser = pgTable(
  'OrganizationUser',
  {
    userId: text('userId').notNull(),
    organizationId: text('organizationId').notNull(),
    role: organizationUserRoleEnum('role').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    /**
     * Department assignment for spend attribution (the people lens).
     * Single-valued, org-scoped, accounting-only (never gates access).
     * Plain column ref (relationMode="prisma"); the service validates the
     * id belongs to this org. See departments.feature.
     */
    departmentId: text('departmentId'),
    /**
     * Set when an admin disables this membership to stay within the licensed
     * seat count. The person loses access to this organization only, keeps
     * their role, department and everything they did, and frees their seat.
     * Reversible: re-enabling restores access if a seat is available.
     * Seat counting excludes disabled memberships. See
     * seat-reconciliation.feature.
     */
    disabledAt: timestamp('disabledAt', { precision: 3 }),
    /**
     * Stable grant intent for an unfinished automatic SSO admission.
     * Cleared after the authoritative grant is observed; never inferred from missing access.
     */
    pendingSsoGrantId: text('pendingSsoGrantId'),
    /**
     * A new generation for each membership lifetime. Stamped USER grant facts
     * must match the current active membership before they can be projected;
     * removing and rejoining therefore cannot revive a delayed old attach.
     */
    membershipStamp: text('membershipStamp')
      .notNull()
      .default(sql`gen_random_uuid()::text`),
    /**
     * When this membership started holding the full seat it holds now; null means createdAt.
     * Anchors the 30-day seat hold (specs/licensing/seat-rotation.feature).
     */
    fullSeatSince: timestamp('fullSeatSince', { precision: 3 }),
    /**
     * Set when admission wrote this membership disabled because no seat was free: the person
     * is waiting for one. Cleared when the membership is enabled.
     */
    seatPendingSince: timestamp('seatPendingSince', { precision: 3 }),
  },
  (table) => [
    primaryKey({ name: 'OrganizationUser_pkey', columns: [table.userId, table.organizationId] }),
    index('OrganizationUser_organizationId_idx').on(table.organizationId),
    index('OrganizationUser_userId_idx').on(table.userId),
    index('OrganizationUser_departmentId_idx').on(table.departmentId),
    foreignKey({
      name: 'OrganizationUser_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'OrganizationUser_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

/**
 * A full seat still counted after the person who took it stopped holding it (removed,
 * disabled or demoted), until 30 days after they took it (specs/licensing/seat-rotation.feature).
 * One per organization and person; a return within the hold takes the seat back.
 */
export const seatHold = pgTable(
  'SeatHold',
  {
    id: text('id').primaryKey(),
    organizationId: text('organizationId').notNull(),
    userId: text('userId').notNull(),
    /** Display copy for the Users list, kept because a removed person has no membership row. */
    email: text('email'),
    name: text('name'),
    /** When the person took the seat this hold keeps. */
    seatTakenAt: timestamp('seatTakenAt', { precision: 3 }).notNull(),
    /** seatTakenAt plus 30 days; expired holds are ignored by every count and deleted hourly. */
    heldUntil: timestamp('heldUntil', { precision: 3 }).notNull(),
    /** REMOVED | DISABLED | DEMOTED */
    reason: text('reason').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('SeatHold_heldUntil_idx').on(table.heldUntil),
    uniqueIndex('SeatHold_organizationId_userId_key').on(table.organizationId, table.userId),
  ],
)

/**
 * A person asked for a full seat the organization had no room for (modules/organization/specs/
 * seat-requests.feature). One row per organization and address; reopening reuses the row.
 */
export const seatRequest = pgTable(
  'SeatRequest',
  {
    id: text('id').primaryKey(),
    organizationId: text('organizationId').notNull(),
    /** Lowercased address of the person asking. */
    email: text('email').notNull(),
    /** Null when only the address is known (a sign-in the identity provider refused). */
    userId: text('userId'),
    name: text('name'),
    /** OPEN | APPROVED | DECLINED */
    status: text('status').notNull(),
    /** pending | lite | idp_refused: how the person was left when the request opened. */
    source: text('source').notNull(),
    /** The identity provider also refused this person's sign-in. */
    idpRefused: boolean('idpRefused').notNull().default(false),
    /** okta | entra | google | onelogin | jumpcloud | ping | auth0 | oidc | saml */
    idpKind: text('idpKind'),
    /** When sales were told about this request; reset when it reopens. */
    salesNotifiedAt: timestamp('salesNotifiedAt', { precision: 3 }),
    openedAt: timestamp('openedAt', { precision: 3 }).notNull(),
    resolvedAt: timestamp('resolvedAt', { precision: 3 }),
    resolvedByUserId: text('resolvedByUserId'),
    /**
     * Set when an approval claims the request, cleared once its seat is given; an approval
     * still pending long after is one whose seat never landed, and is reopened.
     */
    seatPendingSince: timestamp('seatPendingSince', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('SeatRequest_organizationId_status_idx').on(table.organizationId, table.status),
    index('SeatRequest_userId_idx').on(table.userId),
    uniqueIndex('SeatRequest_organizationId_email_key').on(table.organizationId, table.email),
  ],
)

/**
 * Held by the one seat-taking write running on an organization (seat-requests.feature). A
 * row past expiresAt is free to take; a holder deletes its own row when its write is done.
 */
export const organizationSeatLease = pgTable('OrganizationSeatLease', {
  organizationId: text('organizationId').primaryKey(),
  /** Random per hold, so a release deletes only its own lease. */
  holder: text('holder').notNull(),
  expiresAt: timestamp('expiresAt', { precision: 3 }).notNull(),
  createdAt: timestamp('createdAt', { precision: 3 })
    .notNull()
    .default(sql`CURRENT_TIMESTAMP`)
    .$defaultFn(utcNow),
})

/**
 * The mail log for seat requests and refused sign-ins: at most one row per organization,
 * kind, address and UTC day, each mailed at once.
 */
export const organizationAccessNotice = pgTable(
  'OrganizationAccessNotice',
  {
    id: text('id').primaryKey(),
    organizationId: text('organizationId').notNull(),
    /** seat_request | idp_refused */
    kind: text('kind').notNull(),
    email: text('email').notNull(),
    name: text('name'),
    idp: text('idp'),
    /** UTC day the notice was recorded on, YYYY-MM-DD. */
    dayKey: text('dayKey').notNull(),
    /** Whether the seat request it announces was also refused by the identity provider. */
    idpRefused: boolean('idpRefused').notNull().default(false),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    /** Set once the notice was handed to the mailer; null if a send never started. */
    mailedAt: timestamp('mailedAt', { precision: 3 }),
  },
  (table) => [
    uniqueIndex('OrganizationAccessNotice_organizationId_kind_email_dayKey_key').on(
      table.organizationId,
      table.kind,
      table.email,
      table.dayKey,
    ),
  ],
)

export const team = pgTable(
  'Team',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    name: text('name').notNull(),
    slug: text('slug').notNull(),
    organizationId: text('organizationId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    archivedAt: timestamp('archivedAt', { precision: 3 }),
    isPersonal: boolean('isPersonal').notNull().default(false),
    ownerUserId: varchar('ownerUserId', { length: 255 }),
    /**
     * Department assignment inherited as the default for members and
     * projects that have none of their own. Accounting-only.
     */
    departmentId: text('departmentId'),
  },
  (table) => [
    uniqueIndex('Team_slug_key').on(table.slug),
    index('Team_organizationId_idx').on(table.organizationId),
    index('Team_isPersonal_idx').on(table.isPersonal),
    index('Team_departmentId_idx').on(table.departmentId),
    foreignKey({
      name: 'Team_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Team_ownerUserId_fkey',
      columns: [table.ownerUserId],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const organization = pgTable(
  'Organization',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    name: text('name').notNull(),
    /**
     * `customer` for every organization a person founds or an operator provisions; `system` for
     * the install's one hidden organization the Ops checkup canaries run in, never listed to a person.
     * Spec: modules/ops/specs/checkup-canary-project.feature.
     */
    kind: text('kind').notNull().default('customer'),
    phoneNumber: text('phoneNumber'),
    slug: text('slug').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    usageSpendingMaxLimit: integer('usageSpendingMaxLimit'),
    /**
     * The largest file a dataset image or file cell accepts for this
     * organization, in mebibytes. Null means the default. Set by an operator
     * from the back office only; every other dataset size limit scales with it.
     */
    datasetAttachmentMaxMb: integer('datasetAttachmentMaxMb'),
    /**
     * The label of this organization's `CLICKHOUSE_URL__<label>__<id>` route, recorded by the
     * organization module the first time it is seen. Once set, a missing route refuses the
     * organization instead of serving it from the shared ClickHouse. Null: never private.
     * Spec: specs/private-dataplane/clickhouse-routing.feature.
     */
    privateClickHouseLabel: text('privateClickHouseLabel'),
    /** The same for `DATAPLANE_S3__<label>__<id>`. Spec: specs/private-dataplane/s3-routing.feature. */
    privateS3Label: text('privateS3Label'),
    /** When either label was first recorded. */
    privateDataplaneSeenAt: timestamp('privateDataplaneSeenAt', { precision: 3 }),
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
    maxSessionDurationDays: integer('maxSessionDurationDays').notNull().default(0),
    /**
     * Every member of this organization can prove a second factor (D06). A
     * MEMBERSHIP CONDITION, not a session policy: it is asked when a member
     * reaches this organization's data, it holds the ones who cannot prove
     * one at an enrollment gate for THIS organization alone, and it ends no
     * session, ever. Turning it on signs nobody out.
     * Spec: specs/identity/mfa-and-session-shape.feature.
     */
    mfaRequired: boolean('mfaRequired').notNull().default(false),
    /**
     * Whether members may connect the documentation playground to their own access with a
     * short-lived docs session key (ADR-185, amendment of 2026-11-13). On by default; turned off,
     * no key is minted and live ones stop authenticating. The pasted-key path is untouched.
     * Spec: modules/api-key/specs/docs-session-keys.feature.
     */
    docsPlaygroundSessionsEnabled: boolean('docsPlaygroundSessionsEnabled').notNull().default(true),
    /**
     * When the docs playground was last turned off, written in the same update as the switch. A
     * docs session key minted at or before it never authenticates again, nor any key minted under
     * one, so turning it back on revives none of them (ADR-185, amendment of 2026-11-14).
     */
    docsPlaygroundSessionsDisabledAt: timestamp('docsPlaygroundSessionsDisabledAt', {
      precision: 3,
    }),
    /**
     * How many consecutive failed sign-ins lock an account (GAC-09).
     * 0 = never lock, which is the default and what every organization that
     * has not asked for this keeps. Counts EVERY credential together - a
     * password, a one-time code, a backup code, a passkey - because five per
     * factor hands an attacker one budget per factor.
     * Spec: specs/identity/org-account-lockout.feature.
     */
    lockoutAfterFailedAttempts: integer('lockoutAfterFailedAttempts').notNull().default(0),
    /** How long a lock lasts, in minutes. Ignored while the threshold is 0. */
    lockoutMinutes: integer('lockoutMinutes').notNull().default(30),
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
    sessionIdleTimeoutMinutes: integer('sessionIdleTimeoutMinutes').notNull().default(1440),
    /**
     * The ceiling on a browser session from the moment it was minted, in
     * minutes, which no amount of activity extends. 0 = no ceiling, the
     * default. Offered beside the idle timeout because a session kept alive
     * forever by a polling page is not really bounded.
     */
    sessionMaxLifetimeMinutes: integer('sessionMaxLifetimeMinutes').notNull().default(0),
    signupData: jsonb('signupData'),
    signedDPA: boolean('signedDPA').notNull().default(false),
    elasticsearchNodeUrl: text('elasticsearchNodeUrl'),
    elasticsearchApiKey: text('elasticsearchApiKey'),
    useCustomElasticsearch: boolean('useCustomElasticsearch').notNull().default(false),
    s3Endpoint: text('s3Endpoint'),
    s3AccessKeyId: text('s3AccessKeyId'),
    s3SecretAccessKey: text('s3SecretAccessKey'),
    s3Bucket: text('s3Bucket'),
    useCustomS3: boolean('useCustomS3').notNull().default(false),
    sentPlanLimitAlert: timestamp('sentPlanLimitAlert', { precision: 3 }),
    ssoDomain: text('ssoDomain'),
    ssoProvider: text('ssoProvider'),
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
    domainJoin: text('domainJoin').notNull().default('request'),
    /**
     * The organisation role handed to a person admitted without an invitation:
     * a self-service domain join (`domainJoin`) or an SSO-admitted login.
     * `MEMBER` (default, a Full seat) or `DEVELOPER` (ADR-171). Set on the
     * organisation settings page only; the join flow never asks, and an
     * invitation always names its own role.
     */
    joinerRole: organizationUserRoleEnum('joinerRole').notNull().default('MEMBER'),
    /**
     * The domains an administrator NAMED when turning automatic joining on.
     * Empty means automatic joining admits nobody whatever `domainJoin` says -
     * which is what keeps a setting left on from admitting a domain nobody
     * ever agreed to.
     */
    joinDomains: text('joinDomains').array(),
    presenceEnabled: boolean('presenceEnabled').notNull().default(true),
    /**
     * Org-level trace-sharing kill switch. Off = trace share links cannot be
     * created in ANY project of this org, and existing links stop resolving —
     * the org toggle is the global switch, the per-project
     * `Project.traceSharingEnabled` scopes the kill to a single project (mirrors
     * `presenceEnabled`). Effective sharing = org AND project. ADR-039.
     */
    traceSharingEnabled: boolean('traceSharingEnabled').notNull().default(true),
    /**
     * Org-configurable contact surfaced to users in CLI "contact your
     * admin" messages and the in-app budget-exceeded banner. Free text:
     * accepts an email, a URL (internal ticketing system), or any other
     * instruction. When null, callers fall back to the first ADMIN
     * member's email (`resolveSupportContact` helper).
     */
    supportContact: text('supportContact'),
    /**
     * Signup intent, editable later as the org "Primary use" setting.
     * AGENT_GOVERNANCE lands "/" on /me, LLM_OPS on the project home.
     * NULL means intent unset (every pre-fork org and any org created
     * outside onboarding) — the home resolver keeps legacy behavior.
     * ADR-038.
     */
    primaryIntent: organizationIntentEnum('primaryIntent'),
    promoCode: text('promoCode'),
    stripeCustomerId: text('stripeCustomerId'),
    currency: currencyEnum('currency').notNull().default('EUR'),
    pricingModel: pricingModelEnum('pricingModel').notNull().default('TIERED'),
    license: text('license'),
    licenseExpiresAt: timestamp('licenseExpiresAt', { precision: 3 }),
    licenseLastValidatedAt: timestamp('licenseLastValidatedAt', { precision: 3 }),
    /**
     * A customer that runs LangWatch self-hosted. The hosted usage, budgets and
     * invoices of its connected installs attach to this organization (ADR-139).
     */
    selfHostedCustomer: boolean('selfHostedCustomer').notNull().default(false),
    /**
     * LangWatch-hosted services this organization has switched OFF, named the
     * way its license names them ("instant_evals"). A service the license is
     * entitled to is on unless it is listed here, so a customer who bought
     * hosted judging gets it without finding a settings page first, and an
     * administrator who switched one off keeps it off across license changes
     * (ADR-139, section 2). A license that names no service reaches nothing,
     * whatever this column holds.
     */
    connectServicesDisabled: text('connectServicesDisabled').array().default([]),
    /**
     * When a member of this organization switched Instant Evals on from the
     * search bar, which is the organization's consent to have judged text sent
     * to the judge's provider. Null until then: a self-serve organization is
     * never opted in on its behalf, and an enterprise one is switched on by an
     * operator through the release flag after asking, never here.
     */
    instantEvalsEnabledAt: timestamp('instantEvalsEnabledAt', { precision: 3 }),
    /**
     * The member who switched them on, so the audit trail names a person and
     * not just a moment.
     */
    instantEvalsEnabledByUserId: text('instantEvalsEnabledByUserId'),
    /** When the last license sync succeeded. */
    connectLastSyncAt: timestamp('connectLastSyncAt', { precision: 3 }),
    /**
     * The code of the failure the last sync ended on, cleared on the next
     * success. Settings, Connect renders the copy for it.
     */
    connectLastSyncError: text('connectLastSyncError'),
  },
  (table) => [
    uniqueIndex('Organization_slug_key').on(table.slug),
    uniqueIndex('Organization_ssoDomain_key').on(table.ssoDomain),
    uniqueIndex('Organization_stripeCustomerId_key').on(table.stripeCustomerId),
    index('Organization_ssoProvider_idx').on(table.ssoProvider),
  ],
)

/**
 * The registry of licenses LangWatch has issued (ADR-139). It lives on
 * LangWatch Cloud; a self-hosted install has the table and no rows.
 *
 * The row never holds the token an install presents. `tokenHash` is the
 * SHA-256 of that token, so read access to this table is not credential
 * access. Status is derived: revoked, superseded, past `expiresAt`, else active.
 */
export const issuedLicense = pgTable(
  'IssuedLicense',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    /** The `licenseId` inside the signed license (`lic-<uuid>`). */
    licenseId: text('licenseId').notNull(),
    /** SHA-256 of the `lwl_` token. Looked up on every credential resolution. */
    tokenHash: text('tokenHash').notNull(),
    /**
     * The customer organization on LangWatch Cloud. Null for a license that is
     * recorded but not linked to a customer yet, such as a self-serve purchase:
     * it resolves to nothing until an operator links it.
     */
    organizationId: text('organizationId'),
    /** As signed into the license. */
    organizationName: text('organizationName').notNull(),
    email: text('email').notNull(),
    planType: text('planType').notNull(),
    maxMembers: integer('maxMembers').notNull(),
    /** The lite seats sold: null when not stated or when unlimited. Never a stand-in number. */
    maxMembersLite: integer('maxMembersLite'),
    /** True when the license sells unlimited lite seats; maxMembersLite is then null. */
    liteSeatsUnlimited: boolean('liteSeatsUnlimited').notNull().default(false),
    /** The Developer seats sold: null when not stated, which sells none. */
    maxMembersDeveloper: integer('maxMembersDeveloper'),
    issuedAt: timestamp('issuedAt', { precision: 3 }).notNull(),
    expiresAt: timestamp('expiresAt', { precision: 3 }).notNull(),
    source: issuedLicenseSourceEnum('source').notNull(),
    /** The operator who issued or registered it. Null for the purchase flow and the script. */
    issuedById: text('issuedById'),
    revokedAt: timestamp('revokedAt', { precision: 3 }),
    revokedById: text('revokedById'),
    revokedReason: text('revokedReason'),
    /** Set when the install first presents the license that replaced this one. */
    supersededAt: timestamp('supersededAt', { precision: 3 }),
    /**
     * The id of the IssuedLicense this one replaces, when it was reissued. A
     * plain column ref (relationMode="prisma" cannot declare a self-relation on
     * Postgres); the registry service validates it. Unique, so a license is
     * replaced by one license at most.
     */
    replacesId: text('replacesId'),
    /**
     * A reissued license waiting to be delivered over sync, encrypted. Erased
     * when the install first presents it. The only license text the registry holds.
     */
    pendingDeliveryLicense: text('pendingDeliveryLicense'),
    /** Hosted services this license may call, such as `instant_evals`. */
    services: text('services').array().default([]),
    /** The annual rate of one seat, in cents of `seatCurrency`. Needed to prorate added seats. */
    seatRateCents: integer('seatRateCents'),
    seatCurrency: currencyEnum('seatCurrency'),
    /** The prepaid usage commit for the term, in USD cents. */
    commitUsdCents: integer('commitUsdCents').notNull().default(0),
    /** Whether usage may continue past the commit, up to `overageMaxUsdCents`. */
    overageEnabled: boolean('overageEnabled').notNull().default(false),
    overageMaxUsdCents: integer('overageMaxUsdCents'),
    /** The install this license is bound to. Set on first use, cleared by an operator. */
    instanceId: text('instanceId'),
    instanceBoundAt: timestamp('instanceBoundAt', { precision: 3 }),
    /**
     * The last license sync this install reported: when it arrived, the
     * LangWatch version that sent it, and the seats in use at that moment.
     * Null until the install syncs for the first time.
     */
    lastSyncAt: timestamp('lastSyncAt', { precision: 3 }),
    lastSyncVersion: text('lastSyncVersion'),
    reportedMembers: integer('reportedMembers'),
    reportedMembersLite: integer('reportedMembersLite'),
    /**
     * The managed gateway key (purpose CONNECT) that hosted-service traffic of
     * this license runs under. One per license, so spend rows name the install.
     * Created when the license first resolves. A plain column, not a relation.
     */
    virtualKeyId: text('virtualKeyId'),
    /**
     * The seats of the license this one replaced, set only when a seat change
     * raised a linked license. Billing reads it through licensing to invoice
     * the added seats, so the fact commits with the license or not at all.
     */
    seatsRaisedFrom: integer('seatsRaisedFrom'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('IssuedLicense_licenseId_key').on(table.licenseId),
    uniqueIndex('IssuedLicense_tokenHash_key').on(table.tokenHash),
    uniqueIndex('IssuedLicense_replacesId_key').on(table.replacesId),
    uniqueIndex('IssuedLicense_virtualKeyId_key').on(table.virtualKeyId),
    index('IssuedLicense_organizationId_idx').on(table.organizationId),
    index('IssuedLicense_expiresAt_idx').on(table.expiresAt),
    index('IssuedLicense_seatsRaisedFrom_idx').on(table.seatsRaisedFrom),
    foreignKey({
      name: 'IssuedLicense_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

/**
 * The licence an organization activated and its dates, owned by licensing
 * (round 37 D6). A row whose key is null is a cleared licence. An organization
 * with no row is read from its Organization columns until the copy step runs.
 */
export const organizationLicense = pgTable('OrganizationLicense', {
  organizationId: text('organizationId').primaryKey(),
  licenseKey: text('licenseKey'),
  expiresAt: timestamp('expiresAt', { precision: 3 }),
  validatedAt: timestamp('validatedAt', { precision: 3 }),
  updatedAt: timestamp('updatedAt', { precision: 3 })
    .notNull()
    .default(sql`CURRENT_TIMESTAMP`)
    .$defaultFn(utcNow)
    .$onUpdate(utcNow),
})

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
export const activationCode = pgTable(
  'ActivationCode',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    /** SHA-256 of the normalised code. Looked up on every redemption. */
    codeHash: text('codeHash').notNull(),
    /**
     * The last four characters, so an operator tells two codes apart in a list
     * without the list holding either one.
     */
    codeHint: text('codeHint').notNull(),
    /** The customer this code belongs to, and the terms it mints. */
    organizationId: text('organizationId').notNull(),
    organizationName: text('organizationName').notNull(),
    email: text('email').notNull(),
    planType: text('planType').notNull(),
    maxMembers: integer('maxMembers').notNull(),
    /** The lite seats the minted license carries: null when unlimited. Never a stand-in number. */
    maxMembersLite: integer('maxMembersLite').default(0),
    /** True when the minted license carries unlimited lite seats; maxMembersLite is then null. */
    liteSeatsUnlimited: boolean('liteSeatsUnlimited').notNull().default(false),
    /** The Developer seats the minted license carries: null when not stated, which sells none. */
    maxMembersDeveloper: integer('maxMembersDeveloper'),
    /** How long the license this code mints runs for. */
    licenseTermDays: integer('licenseTermDays').notNull(),
    /** Hosted services the minted license may call. */
    services: text('services').array().default([]),
    /** When the code stops working, which is not the term of what it mints. */
    expiresAt: timestamp('expiresAt', { precision: 3 }).notNull(),
    /** A code a customer may redeem on more than one install, for a rollout. */
    reusable: boolean('reusable').notNull().default(false),
    /**
     * Set by the redemption that won, which is what makes a single-use code
     * single use: the write itself carries the condition.
     */
    redeemedAt: timestamp('redeemedAt', { precision: 3 }),
    redeemedByInstanceId: text('redeemedByInstanceId'),
    /** The license the winning redemption minted. */
    issuedLicenseId: text('issuedLicenseId'),
    /** How many installs have redeemed it, which only moves on a reusable code. */
    redemptionCount: integer('redemptionCount').notNull().default(0),
    revokedAt: timestamp('revokedAt', { precision: 3 }),
    revokedById: text('revokedById'),
    createdById: text('createdById'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('ActivationCode_codeHash_key').on(table.codeHash),
    index('ActivationCode_organizationId_idx').on(table.organizationId),
    index('ActivationCode_expiresAt_idx').on(table.expiresAt),
  ],
)

/**
 * How a connected customer is invoiced (ADR-139, section 7): one payment
 * provider customer, one quarterly usage subscription in USD, and the
 * commercial terms the annual invoice was raised from.
 *
 * One row per organization. Every write through it is addressed by a stored
 * provider id, so running onboarding or a renewal twice changes nothing.
 */
export const connectedBillingAccount = pgTable(
  'ConnectedBillingAccount',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    /** The customer at the payment provider. Invoiced, never charged. */
    stripeCustomerId: text('stripeCustomerId').notNull(),
    /**
     * The one metered subscription, and the item hosted usage is reported on.
     * Null between the customer being created and the subscription existing,
     * which is where a half-finished onboarding resumes from.
     */
    usageSubscriptionId: text('usageSubscriptionId'),
    usageSubscriptionItemId: text('usageSubscriptionItemId'),
    termStartsAt: timestamp('termStartsAt', { precision: 3 }).notNull(),
    termEndsAt: timestamp('termEndsAt', { precision: 3 }).notNull(),
    /** Prepaid usage for the term, in USD cents. The credit grants add up to it. */
    commitUsdCents: integer('commitUsdCents').notNull().default(0),
    /** The seat contract, which may be in another currency than hosted usage. */
    seatCurrency: currencyEnum('seatCurrency').notNull(),
    seatRateCents: integer('seatRateCents').notNull(),
    seats: integer('seats').notNull(),
    /**
     * `us_bank_transfer` or `eu_bank_transfer` when the customer pays into a
     * virtual bank account at the provider. Null when it wires to LangWatch's
     * own account and finance marks the invoice paid.
     */
    bankTransferType: text('bankTransferType'),
    bankTransferCountry: text('bankTransferCountry'),
    billingEmail: text('billingEmail').notNull(),
    /**
     * A renewal whose credit is waiting for the last usage invoice of the old
     * term to be finalized: the new term's end and commit, and who asked.
     */
    pendingRenewal: jsonb('pendingRenewal'),
    createdById: text('createdById'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('ConnectedBillingAccount_organizationId_key').on(table.organizationId),
    uniqueIndex('ConnectedBillingAccount_stripeCustomerId_key').on(table.stripeCustomerId),
    foreignKey({
      name: 'ConnectedBillingAccount_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

/**
 * A paid credit at the payment provider, which only metered usage draws down.
 * Kept here so a grant is created once per reason and the backoffice can show
 * what the customer holds without asking the provider.
 */
export const connectedCreditGrant = pgTable(
  'ConnectedCreditGrant',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    /**
     * The `ConnectedBillingAccount` row. A plain column ref, like
     * `IssuedLicense.replacesId`: relationMode="prisma" carries no foreign keys.
     */
    accountId: text('accountId').notNull(),
    /** The grant at the payment provider. Unique, so a retry records one row. */
    stripeCreditGrantId: text('stripeCreditGrantId').notNull(),
    amountUsdCents: integer('amountUsdCents').notNull(),
    kind: connectedCreditGrantKindEnum('kind').notNull(),
    /**
     * The term this credit was granted for, and when the provider lets it lapse:
     * after the last invoice of that term has been finalized, never on the day
     * the term ends.
     */
    termEndsAt: timestamp('termEndsAt', { precision: 3 }).notNull(),
    expiresAt: timestamp('expiresAt', { precision: 3 }).notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    uniqueIndex('ConnectedCreditGrant_stripeCreditGrantId_key').on(table.stripeCreditGrantId),
    index('ConnectedCreditGrant_accountId_idx').on(table.accountId),
  ],
)

/**
 * An invoice LangWatch raised for a connected customer, as the backoffice
 * reads it.
 */
export const connectedInvoice = pgTable(
  'ConnectedInvoice',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    /** The `ConnectedBillingAccount` row, as a plain column ref. */
    accountId: text('accountId').notNull(),
    stripeInvoiceId: text('stripeInvoiceId').notNull(),
    kind: connectedInvoiceKindEnum('kind').notNull(),
    currency: currencyEnum('currency').notNull(),
    amountCents: integer('amountCents').notNull(),
    /** The provider's own status: `draft`, `open`, `paid`, `void` or `uncollectible`. */
    status: text('status').notNull(),
    paidOutOfBandAt: timestamp('paidOutOfBandAt', { precision: 3 }),
    termStartsAt: timestamp('termStartsAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('ConnectedInvoice_stripeInvoiceId_key').on(table.stripeInvoiceId),
    index('ConnectedInvoice_accountId_idx').on(table.accountId),
  ],
)

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
export const scimRequestLog = pgTable(
  'ScimRequestLog',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    /**
     * Null when the refusal resolved an organization and stopped before the
     * connection — a lapsed plan is answered off the token alone.
     */
    connectionId: text('connectionId'),
    method: text('method').notNull(),
    /**
     * The SCIM resource asked for ("Users", "Users/:id", "Groups"), never the
     * raw path: a path carries ids and query strings, and this is read by
     * people rather than matched by machines.
     */
    resource: text('resource').notNull(),
    status: integer('status').notNull(),
    /**
     * A stable slug, only on a refusal, so a reader branches on this and never
     * on the sentence beside it.
     */
    reason: text('reason'),
    /** Our own short sentence, customer-safe. */
    detail: text('detail'),
    occurredAt: timestamp('occurredAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('ScimRequestLog_organizationId_occurredAt_idx').on(
      table.organizationId,
      table.occurredAt,
    ),
    index('ScimRequestLog_connectionId_occurredAt_idx').on(table.connectionId, table.occurredAt),
    index('ScimRequestLog_occurredAt_idx').on(table.occurredAt),
    foreignKey({
      name: 'ScimRequestLog_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const scimToken = pgTable(
  'ScimToken',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
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
    connectionId: text('connectionId'),
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
    hashedToken: text('hashedToken').notNull(),
    /**
     * How `hashedToken` was derived, so the two schemes can coexist while
     * tokens minted under the old one are still in use.
     *
     * `sha256` is a bare digest, which is correct for 32 random bytes and wrong
     * for a value a person chose — a database dump of those is a wordlist away
     * from live credentials. `hmac-sha256` is keyed on the deployment's own
     * secret, so a dump on its own is inert. Every new row is the latter.
     */
    hashScheme: text('hashScheme').notNull().default('sha256'),
    description: text('description'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    lastUsedAt: timestamp('lastUsedAt', { precision: 3 }),
  },
  (table) => [
    uniqueIndex('ScimToken_hashedToken_key').on(table.hashedToken),
    index('ScimToken_organizationId_idx').on(table.organizationId),
    index('ScimToken_connectionId_idx').on(table.connectionId),
    foreignKey({
      name: 'ScimToken_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const scimExternalId = pgTable(
  'ScimExternalId',
  {
    organizationId: text('organizationId').notNull(),
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    connectionId: text('connectionId').notNull(),
    externalId: text('externalId').notNull(),
    userId: text('userId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('ScimExternalId_organizationId_userId_idx').on(table.organizationId, table.userId),
    index('ScimExternalId_userId_idx').on(table.userId),
    index('ScimExternalId_connectionId_idx').on(table.connectionId),
    uniqueIndex('ScimExternalId_connectionId_externalId_key').on(
      table.connectionId,
      table.externalId,
    ),
    foreignKey({
      name: 'ScimExternalId_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/** Directory profile and lifecycle are tenant facts, independent of the shared account. */
export const scimUserResource = pgTable(
  'ScimUserResource',
  {
    organizationId: text('organizationId').notNull(),
    userId: text('userId').notNull(),
    userName: text('userName').notNull(),
    name: text('name'),
    active: boolean('active').notNull().default(true),
    deletedAt: timestamp('deletedAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    primaryKey({ name: 'ScimUserResource_pkey', columns: [table.organizationId, table.userId] }),
    index('ScimUserResource_organizationId_userName_idx').on(table.organizationId, table.userName),
    index('ScimUserResource_userId_idx').on(table.userId),
    foreignKey({
      name: 'ScimUserResource_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * A connection's managed users, including resources without an externalId.
 * Scalar keys only; readers exclude users that no longer exist.
 */
export const scimDirectoryUser = pgTable(
  'ScimDirectoryUser',
  {
    organizationId: text('organizationId').notNull(),
    connectionId: text('connectionId').notNull(),
    userId: text('userId').notNull(),
  },
  (table) => [
    primaryKey({ name: 'ScimDirectoryUser_pkey', columns: [table.connectionId, table.userId] }),
    index('ScimDirectoryUser_organizationId_userId_idx').on(table.organizationId, table.userId),
  ],
)

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
export const scimSyncState = pgTable(
  'ScimSyncState',
  {
    /** The aggregate: the connection's id, which IS its sync's id. */
    id: text('id').primaryKey(),
    connectionId: text('connectionId').notNull(),
    organizationId: text('organizationId').notNull(),
    /** TOKEN_ISSUED | SYNCING | ERROR | REVOKED */
    state: text('state').notNull(),
    lastPushedAt: timestamp('lastPushedAt', { precision: 3 }),
    /**
     * The failure standing right now: { op, errorCode, attempts, retiredAtMs,
     * userId, occurredAtMs }. A reason CODE and ids — never a token, never a
     * secret, never an internal hostname.
     */
    lastFailure: jsonb('lastFailure'),
    /**
     * Every failure retired without ever being applied, newest last. The dead
     * letters an administrator has to see; retiring never drops one.
     */
    deadLetters: jsonb('deadLetters').notNull().default([]),
    /** revoke | teardown, once the sync has ended. */
    revokedCause: text('revokedCause'),
    occurredAt: timestamp('occurredAt', { precision: 3 }).notNull(),
    lastEventId: text('lastEventId').notNull(),
    acceptedAt: timestamp('acceptedAt', { precision: 3 }).notNull(),
    projectionVersion: text('projectionVersion').notNull(),
    /**
     * Event-derived, not database-managed: `@default(now())` / `@updatedAt`
     * would make a replayed row differ from the row it rebuilds, which is
     * exactly the parity this projection has to keep.
     */
    createdAt: timestamp('createdAt', { precision: 3 }).notNull(),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull(),
  },
  (table) => [
    index('ScimSyncState_organizationId_idx').on(table.organizationId),
    index('ScimSyncState_state_idx').on(table.state),
    index('ScimSyncState_connectionId_idx').on(table.connectionId),
  ],
)

/**
 * SCIM's own copy of an organization's SSO connections: a peer fold over
 * identity's connection facts (ARCHITECTURE §9), folded with identity-contract's
 * reducer, one row per connection. SCIM reads this instead of asking identity.
 * Rebuilt from identity's event log by a projection replay, so no relation:
 * a foreign key would make a replay depend on state the log does not carry.
 */
export const scimSsoConnectionView = pgTable(
  'ScimSsoConnectionView',
  {
    /** The aggregate: identity's connection id. */
    id: text('id').primaryKey(),
    organizationId: text('organizationId').notNull(),
    /** The connection as the reducer folded it (SsoConnectionState + LastEventOccurredAt). */
    folded: jsonb('folded').notNull(),
    /** The event ids of the last delivery folded in, so a redelivery folds once. */
    appliedEventIds: text('appliedEventIds').array().default([]),
    /** A row of another version reads as undecodable and is re-folded from the log. */
    projectionVersion: text('projectionVersion').notNull(),
    /** Event-derived, so a replayed row equals the row it rebuilds. */
    createdAt: timestamp('createdAt', { precision: 3 }).notNull(),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull(),
  },
  (table) => [
    index('ScimSsoConnectionView_organizationId_createdAt_idx').on(
      table.organizationId,
      table.createdAt,
    ),
  ],
)

/**
 * Organization's monthly seat peak, folded from its seat changes (modules/organization, ADR-002).
 * A read model, one row per organization: rebuildable by replaying the event log, so no relation.
 * Epochs in milliseconds.
 * Organization's seat-peak counts (ADR-002): one row per organization once counting opened.
 */
export const organizationSeatPeak = pgTable('OrganizationSeatPeak', {
  /** The organization, which is also the aggregate id. */
  organizationId: text('organizationId').primaryKey(),
  /** When the opening snapshot was taken, epoch milliseconds. */
  openedAt: doublePrecision('openedAt').notNull(),
  fullSeats: integer('fullSeats').notNull(),
  liteSeats: integer('liteSeats').notNull(),
  developerSeats: integer('developerSeats').notNull(),
  /** The latest change time applied to any holder; invitations expire up to it. */
  latestChangeAt: doublePrecision('latestChangeAt').notNull(),
  /** Changes whose `from` disagreed with the seat the fold held. */
  mismatches: integer('mismatches').notNull(),
})

/** One member's or pending invitation's seat as the seat peak last applied it (ADR-002). */
export const organizationSeatPeakHolder = pgTable(
  'OrganizationSeatPeakHolder',
  {
    organizationId: text('organizationId').notNull(),
    /** `member`, `invitation` or `hold` (a seat held after somebody left). */
    holderKind: text('holderKind').notNull(),
    /** The member's or the hold's userId, or the invitation's id. */
    holderId: text('holderId').notNull(),
    /** `full`, `lite`, `developer` or `none`. */
    seat: text('seat').notNull(),
    lastChangedAt: doublePrecision('lastChangedAt').notNull(),
    /** Absent for a seat the opening snapshot recorded. */
    lastChangeId: text('lastChangeId'),
    /** An invitation's expiry; null for a member or an invitation that never expires. */
    expiresAt: doublePrecision('expiresAt'),
    /** The fold gave the invitation's seat back at `expiresAt`. */
    expiredByFold: boolean('expiredByFold').notNull().default(false),
    /** A hold whose seat a live full invitation to its address took: that invitation's id. */
    suppressedBy: text('suppressedBy'),
    /** The fold gave a suppressed hold its seat back when that invitation expired. */
    restoredByFold: boolean('restoredByFold').notNull().default(false),
  },
  (table) => [
    primaryKey({
      name: 'OrganizationSeatPeakHolder_pkey',
      columns: [table.organizationId, table.holderKind, table.holderId],
    }),
    index('OrganizationSeatPeakHolder_organizationId_holderKind_seat_e_idx').on(
      table.organizationId,
      table.holderKind,
      table.seat,
      table.expiresAt,
    ),
    index('OrganizationSeatPeakHolder_organizationId_suppressedBy_idx').on(
      table.organizationId,
      table.suppressedBy,
    ),
  ],
)

/** One calendar month's (UTC) seat peaks for an organization (ADR-002). */
export const organizationSeatPeakMonth = pgTable(
  'OrganizationSeatPeakMonth',
  {
    organizationId: text('organizationId').notNull(),
    /** `YYYY-MM`. */
    month: text('month').notNull(),
    fullPeak: integer('fullPeak').notNull(),
    fullPeakAt: doublePrecision('fullPeakAt').notNull(),
    litePeak: integer('litePeak').notNull(),
    litePeakAt: doublePrecision('litePeakAt').notNull(),
    fullAtEnd: integer('fullAtEnd').notNull(),
    liteAtEnd: integer('liteAtEnd').notNull(),
  },
  (table) => [
    primaryKey({
      name: 'OrganizationSeatPeakMonth_pkey',
      columns: [table.organizationId, table.month],
    }),
  ],
)

export const project = pgTable(
  'Project',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    name: text('name').notNull(),
    slug: text('slug').notNull(),
    apiKey: text('apiKey').notNull(),
    /**
     * When the legacy apiKey last authenticated a request, written at most once a
     * minute per project and process. Spec: modules/api-key/specs/request-credential-type.feature
     */
    legacyApiKeyLastUsedAt: timestamp('legacyApiKeyLastUsedAt', { precision: 3 }),
    /**
     * Random per-project secret whose SHA-256 digest is this project's LangWatchQL
     * SQL tenant capability in ClickHouse. Deliberately NOT the API key: the
     * credential a caller authenticates with and the capability the database
     * resolves a tenant from stay decoupled, so rotating one never invalidates
     * the other, and a leak of the ClickHouse key map's input discloses nothing
     * a caller could authenticate with. Generated by the database on insert.
     */
    lwqlKey: text('lwqlKey')
      .notNull()
      .default(sql`(gen_random_uuid())::text`),
    teamId: text('teamId').notNull(),
    language: text('language').notNull(),
    framework: text('framework').notNull(),
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
    kind: text('kind').notNull().default('application'),
    /**
     * The scope rule an aggregate project materialises into grants. Null on
     * every other kind; decided in ADR-177. Shapes: { kind: "all-personal" } |
     * { kind: "personal-by-department", departmentId } |
     * { kind: "explicit", projectIds: string[] }.
     */
    aggregateRule: jsonb('aggregateRule'),
    firstMessage: boolean('firstMessage').notNull().default(false),
    integrated: boolean('integrated').notNull().default(false),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    userLinkTemplate: text('userLinkTemplate'),
    traceSharingEnabled: boolean('traceSharingEnabled').notNull().default(true),
    presenceEnabled: boolean('presenceEnabled').notNull().default(true),
    s3Endpoint: text('s3Endpoint'),
    s3AccessKeyId: text('s3AccessKeyId'),
    s3SecretAccessKey: text('s3SecretAccessKey'),
    s3Bucket: text('s3Bucket'),
    archivedAt: timestamp('archivedAt', { precision: 3 }),
    isPersonal: boolean('isPersonal').notNull().default(false),
    ownerUserId: varchar('ownerUserId', { length: 255 }),
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
    personalFeatures: jsonb('personalFeatures').notNull().default({}),
    /**
     * Department assignment for spend attribution (the project / agent
     * lens, used when a trace has no principal user). Accounting-only.
     */
    departmentId: text('departmentId'),
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
    langyEgressAllowlist: jsonb('langyEgressAllowlist'),
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
    lastCodingAgentSessionAt: timestamp('lastCodingAgentSessionAt', { precision: 3 }),
    lastCodingAgentPullRequestAt: timestamp('lastCodingAgentPullRequestAt', { precision: 3 }),
  },
  (table) => [
    uniqueIndex('Project_slug_key').on(table.slug),
    uniqueIndex('Project_apiKey_key').on(table.apiKey),
    uniqueIndex('Project_lwqlKey_key').on(table.lwqlKey),
    index('Project_teamId_idx').on(table.teamId),
    index('Project_isPersonal_idx').on(table.isPersonal),
    index('Project_ownerUserId_idx').on(table.ownerUserId),
    index('Project_teamId_kind_idx').on(table.teamId, table.kind),
    index('Project_departmentId_idx').on(table.departmentId),
    foreignKey({ name: 'Project_teamId_fkey', columns: [table.teamId], foreignColumns: [team.id] })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Project_ownerUserId_fkey',
      columns: [table.ownerUserId],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

/**
 * A department is a single-valued accounting dimension, separate from
 * RBAC, for rolling up AI spend by org unit (Engineering, Marketing,
 * ...). Assigned to OrganizationUser, Team, and Project via their
 * `departmentId` column. Pure accounting: never grants or restricts
 * access. See specs/ai-gateway/governance/departments.feature.
 */
export const department = pgTable(
  'Department',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    name: text('name').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    archivedAt: timestamp('archivedAt', { precision: 3 }),
  },
  (table) => [index('Department_organizationId_idx').on(table.organizationId)],
)

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
export const departmentMembershipHistory = pgTable(
  'DepartmentMembershipHistory',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    userId: text('userId').notNull(),
    departmentId: text('departmentId').notNull(),
    validFrom: timestamp('validFrom', { precision: 3 }).notNull(),
    /** Null means the link is open. Reassignment and clearing close it. */
    validTo: timestamp('validTo', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('DepartmentMembershipHistory_organizationId_userId_validFrom_idx').on(
      table.organizationId,
      table.userId,
      table.validFrom,
    ),
  ],
)

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
export const platformToolPolicy = pgTable(
  'PlatformToolPolicy',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    toolSlug: text('toolSlug').notNull(),
    allowVk: boolean('allowVk').notNull().default(true),
    allowOtelDirect: boolean('allowOtelDirect').notNull().default(true),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('PlatformToolPolicy_organizationId_idx').on(table.organizationId),
    uniqueIndex('PlatformToolPolicy_organizationId_toolSlug_key').on(
      table.organizationId,
      table.toolSlug,
    ),
  ],
)

export const organizationInvite = pgTable(
  'OrganizationInvite',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    email: text('email').notNull(),
    inviteCode: text('inviteCode').notNull(),
    expiration: timestamp('expiration', { precision: 3 }),
    status: inviteStatusEnum('status').notNull().default('PENDING'),
    organizationId: text('organizationId').notNull(),
    teamIds: text('teamIds').notNull(),
    teamAssignments: jsonb('teamAssignments'),
    role: organizationUserRoleEnum('role').notNull(),
    requestedBy: text('requestedBy'),
    subscriptionId: text('subscriptionId'),
    acceptedByUserId: text('acceptedByUserId'),
    acceptedViaIdentifierId: text('acceptedViaIdentifierId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('OrganizationInvite_inviteCode_key').on(table.inviteCode),
    index('OrganizationInvite_organizationId_idx').on(table.organizationId),
    index('OrganizationInvite_requestedBy_idx').on(table.requestedBy),
    foreignKey({
      name: 'OrganizationInvite_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'OrganizationInvite_requestedBy_fkey',
      columns: [table.requestedBy],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const monitor = pgTable(
  'Monitor',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    projectId: text('projectId').notNull(),
    experimentId: text('experimentId'),
    evaluatorId: text('evaluatorId'),
    checkType: text('checkType').notNull(),
    name: text('name').notNull(),
    slug: text('slug').notNull(),
    executionMode: evaluationExecutionModeEnum('executionMode').notNull().default('ON_MESSAGE'),
    enabled: boolean('enabled').notNull().default(true),
    preconditions: jsonb('preconditions').notNull(),
    parameters: jsonb('parameters').notNull(),
    mappings: jsonb('mappings'),
    sample: doublePrecision('sample').notNull().default(1),
    level: text('level').notNull().default('trace'),
    threadIdleTimeout: integer('threadIdleTimeout'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    autoPausedAt: timestamp('autoPausedAt', { precision: 3 }),
    autoPausedReason: text('autoPausedReason'),
    evaluatorLabel: text('evaluatorLabel'),
    createdById: text('createdById'),
    updatedById: text('updatedById'),
  },
  (table) => [
    uniqueIndex('Monitor_experimentId_key').on(table.experimentId),
    index('Monitor_projectId_idx').on(table.projectId),
    index('Monitor_evaluatorId_idx').on(table.evaluatorId),
    uniqueIndex('Monitor_projectId_slug_key').on(table.projectId, table.slug),
    foreignKey({
      name: 'Monitor_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Monitor_experimentId_fkey',
      columns: [table.experimentId],
      foreignColumns: [experiment.id],
    })
      .onDelete('restrict')
      .onUpdate('restrict'),
    foreignKey({
      name: 'Monitor_evaluatorId_fkey',
      columns: [table.evaluatorId],
      foreignColumns: [evaluator.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const cost = pgTable(
  'Cost',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    projectId: text('projectId').notNull(),
    costType: costTypeEnum('costType').notNull(),
    costName: text('costName'),
    referenceType: costReferenceTypeEnum('referenceType').notNull(),
    referenceId: text('referenceId').notNull(),
    amount: doublePrecision('amount').notNull(),
    currency: text('currency').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    extraInfo: jsonb('extraInfo'),
  },
  (table) => [
    index('Cost_referenceType_referenceId_idx').on(table.referenceType, table.referenceId),
    index('Cost_costType_idx').on(table.costType),
    index('Cost_projectId_idx').on(table.projectId),
    index('Cost_projectId_createdAt_idx').on(table.projectId, table.createdAt),
    foreignKey({
      name: 'Cost_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

/**
 * PROJECTION of the topic-clustering event stream (ADR-051): topics are
 * `topics_recorded` events; this table is written ONLY by
 * TopicModelFoldProjection (cursor: TopicModelProjection) and is
 * rebuildable by replay. Kept named `Topic` because it is the domain
 * entity every surface reads — the `*Projection` suffix is reserved for
 * cursor/read-model tables.
 */
export const topic = pgTable(
  'Topic',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    projectId: text('projectId').notNull(),
    name: text('name').notNull(),
    parentId: text('parentId'),
    embeddings_model: text('embeddings_model').notNull(),
    centroid: jsonb('centroid').notNull(),
    p95Distance: doublePrecision('p95Distance').notNull(),
    automaticallyGenerated: boolean('automaticallyGenerated').notNull().default(true),
    /**
     * The `topics_recorded` event that recorded this topic (its id in the
     * event log). Null means the row predates event ownership and has not
     * been seeded onto the stream yet; the projection sets it on every row
     * it writes.
     */
    lastEventId: text('lastEventId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('Topic_parentId_idx').on(table.parentId),
    index('Topic_projectId_idx').on(table.projectId),
    foreignKey({
      name: 'Topic_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Topic_parentId_fkey',
      columns: [table.parentId],
      foreignColumns: [table.id],
    })
      .onDelete('restrict')
      .onUpdate('restrict'),
  ],
)

export const dataset = pgTable(
  'Dataset',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    projectId: text('projectId').notNull(),
    name: text('name').notNull(),
    slug: text('slug').notNull(),
    columnTypes: jsonb('columnTypes').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    archivedAt: timestamp('archivedAt', { precision: 3 }),
    mapping: jsonb('mapping'),
    useS3: boolean('useS3').notNull().default(false),
    s3RecordCount: integer('s3RecordCount'),
    contentLayout: text('contentLayout').notNull().default('postgres'),
    status: text('status').notNull().default('ready'),
    statusError: text('statusError'),
    stagingKey: text('stagingKey'),
    sourceStoredObjectId: text('sourceStoredObjectId'),
    uploadFilename: text('uploadFilename'),
    rowCount: integer('rowCount'),
    sizeBytes: bigint('sizeBytes', { mode: 'bigint' }),
    chunkCount: integer('chunkCount'),
    chunkOffsets: jsonb('chunkOffsets'),
  },
  (table) => [
    uniqueIndex('Dataset_projectId_slug_key').on(table.projectId, table.slug),
    foreignKey({
      name: 'Dataset_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

export const datasetRecord = pgTable(
  'DatasetRecord',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    datasetId: text('datasetId').notNull(),
    projectId: text('projectId').notNull(),
    entry: jsonb('entry').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('DatasetRecord_datasetId_idx').on(table.datasetId),
    index('DatasetRecord_projectId_idx').on(table.projectId),
    index('DatasetRecord_createdAt_idx').on(table.createdAt),
    index('DatasetRecord_projectId_datasetId_idx').on(table.projectId, table.datasetId),
    foreignKey({
      name: 'DatasetRecord_datasetId_fkey',
      columns: [table.datasetId],
      foreignColumns: [dataset.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'DatasetRecord_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

/** One immutable snapshot of a dataset's columns and rows; the rows are JSONL chunks in object storage. */
export const datasetVersion = pgTable(
  'DatasetVersion',
  {
    id: text('id').primaryKey(),
    projectId: text('projectId').notNull(),
    datasetId: text('datasetId').notNull(),
    version: integer('version').notNull(),
    /** A name given by hand; a named version is never thinned by retention. */
    name: text('name'),
    /** "autosave", "manual", "experiment_run", "restore" or "initial". */
    trigger: text('trigger').notNull(),
    /** SHA-256 over the columns and the rows in order; equal content saves no new version. */
    contentHash: text('contentHash').notNull(),
    columnTypes: jsonb('columnTypes').notNull(),
    rowCount: integer('rowCount').notNull(),
    sizeBytes: bigint('sizeBytes', { mode: 'bigint' }).notNull(),
    chunkCount: integer('chunkCount').notNull(),
    restoredFromVersion: integer('restoredFromVersion'),
    authorId: text('authorId'),
    /** "user", "api_key" or "system": who saved the version when authorId is null. */
    authorKind: text('authorKind'),
    /** Set when an experiment pins this version; a pinned version is never thinned. */
    pinnedAt: timestamp('pinnedAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('DatasetVersion_projectId_datasetId_idx').on(table.projectId, table.datasetId),
    uniqueIndex('DatasetVersion_datasetId_version_key').on(table.datasetId, table.version),
  ],
)

/** An experiment run that read one dataset version; a version a run read is never thinned. */
export const datasetVersionRun = pgTable(
  'DatasetVersionRun',
  {
    id: text('id').primaryKey(),
    projectId: text('projectId').notNull(),
    datasetId: text('datasetId').notNull(),
    versionId: text('versionId').notNull(),
    runId: text('runId').notNull(),
    experimentId: text('experimentId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('DatasetVersionRun_versionId_idx').on(table.versionId),
    uniqueIndex('DatasetVersionRun_projectId_runId_datasetId_key').on(
      table.projectId,
      table.runId,
      table.datasetId,
    ),
  ],
)

export const dashboard = pgTable(
  'Dashboard',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    /** The project that owns the board; only there can it be edited. */
    projectId: text('projectId').notNull(),
    name: text('name').notNull(),
    /**
     * Legacy board sort position. The Dashboards v2 app orders boards by a
     * member's favourites and never writes this; kept because the legacy
     * Analytics reports section and the REST reorder endpoint still write it,
     * and a rolling deploy's older image still reads it.
     */
    order: integer('order').notNull().default(0),
    /**
     * Free text under the board's name; null until a member writes one.
     * Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
     */
    description: text('description'),
    /**
     * The member who created it; null for boards that predate the column.
     * Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
     */
    createdById: text('createdById'),
    /**
     * Who sees it. Only the author changes it; a board with no author stays PROJECT.
     * Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
     */
    scope: dashboardScopeEnum('scope').notNull().default('PROJECT'),
    /**
     * The owning project's organization, stamped when the board is first set to
     * ORGANIZATION; null before. No relation, per the database rule.
     * Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
     */
    organizationId: text('organizationId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('Dashboard_projectId_idx').on(table.projectId),
    index('Dashboard_projectId_order_idx').on(table.projectId, table.order),
    index('Dashboard_organizationId_scope_idx').on(table.organizationId, table.scope),
    foreignKey({
      name: 'Dashboard_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

/**
 * A member's star, in their own order within a project: a board or a From
 * LangWatch template. Personal: one member's stars never show to another. No
 * foreign key (per the database rule): the dashboard service removes a board's
 * favourites when it deletes the board, so a star never outlives its dashboard.
 */
export const dashboardFavourite = pgTable(
  'DashboardFavourite',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    userId: text('userId').notNull(),
    /**
     * A DashboardsV2 board's id (copied boards kept their Dashboard id); null when
     * the star is a template.
     */
    dashboardId: text('dashboardId'),
    /** The starred From LangWatch template id; null when the star is a board. */
    templateId: text('templateId'),
    projectId: text('projectId').notNull(),
    /** The member's drag order among their own stars in the project. */
    position: integer('position').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('DashboardFavourite_userId_projectId_position_idx').on(
      table.userId,
      table.projectId,
      table.position,
    ),
    index('DashboardFavourite_dashboardId_idx').on(table.dashboardId),
    uniqueIndex('DashboardFavourite_userId_dashboardId_key').on(table.userId, table.dashboardId),
    uniqueIndex('DashboardFavourite_userId_projectId_templateId_key').on(
      table.userId,
      table.projectId,
      table.templateId,
    ),
  ],
)

/**
 * A Dashboards v2 board, apart from the Analytics Reports boards in Dashboard. Boards
 * made before the split were copied here with their ids, so stars and insight
 * pointers still name them. No relation, per the database rule.
 */
export const dashboardsV2 = pgTable(
  'DashboardsV2',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    /** The project that owns the board; only there can it be edited. */
    projectId: text('projectId').notNull(),
    name: text('name').notNull(),
    /** Free text under the board's name; null until a member writes one. */
    description: text('description'),
    /** The member who created it; null for boards that predate authorship. */
    createdById: text('createdById'),
    /** Who sees it. Only the author changes it; a board with no author stays PROJECT. */
    scope: dashboardScopeEnum('scope').notNull().default('PROJECT'),
    /**
     * The owning project's organization, stamped when the board is first set to
     * ORGANIZATION; null before.
     */
    organizationId: text('organizationId'),
    /**
     * Deprecated: unused since each widget query names its own sources (dashboards-v2 AC200,
     * 2026-10-28); dropped in a later migration, once no serving release reads it.
     */
    excludedOrigins: jsonb('excludedOrigins'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('DashboardsV2_projectId_idx').on(table.projectId),
    index('DashboardsV2_organizationId_scope_idx').on(table.organizationId, table.scope),
  ],
)

/**
 * A Dashboards v2 widget: a srcdoc definition placed on a DashboardsV2 board, or
 * unplaced. Widgets made before the split were copied from CustomGraph with their
 * ids. No relation, per the database rule.
 */
export const dashboardWidget = pgTable(
  'DashboardWidget',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    projectId: text('projectId').notNull(),
    /** The DashboardsV2 board it sits on; null while unplaced. */
    dashboardId: text('dashboardId'),
    name: text('name').notNull(),
    /** The widget's srcdoc definition. */
    definition: jsonb('definition').notNull(),
    gridColumn: integer('gridColumn').notNull().default(0),
    gridRow: integer('gridRow').notNull().default(0),
    colSpan: integer('colSpan').notNull().default(1),
    rowSpan: integer('rowSpan').notNull().default(1),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('DashboardWidget_projectId_dashboardId_idx').on(table.projectId, table.dashboardId),
  ],
)

/**
 * One insight in its owner's inbox in a project, folded from its `insight` aggregate
 * (modules/insight). A read model: rebuildable by replaying the event log. Epochs in
 * milliseconds.
 */
export const insightProjection = pgTable(
  'InsightProjection',
  {
    /** The insight id, which is also its aggregate id. */
    id: text('id').primaryKey(),
    projectId: text('projectId').notNull(),
    title: text('title').notNull(),
    body: text('body').notNull(),
    /** bad | watch | good */
    tone: text('tone').notNull(),
    topic: text('topic'),
    validDays: integer('validDays').notNull(),
    lwql: text('lwql'),
    sourceConversationId: text('sourceConversationId'),
    sourceMessageId: text('sourceMessageId'),
    filedByUserId: text('filedByUserId'),
    filedAt: doublePrecision('filedAt').notNull(),
    /** Set by a later run that found the insight still true. */
    renewedAt: doublePrecision('renewedAt'),
    createdAt: doublePrecision('createdAt').notNull(),
    updatedAt: doublePrecision('updatedAt').notNull(),
    occurredAt: doublePrecision('occurredAt').notNull(),
    acceptedAt: doublePrecision('acceptedAt').notNull(),
    lastEventId: text('lastEventId').notNull(),
    projectionVersion: text('projectionVersion').notNull(),
    /**
     * The board and widget it came from. A pointer only, never a relation: the row reads the
     * same once either is deleted, so the names are kept as filed.
     */
    boardId: text('boardId'),
    boardName: text('boardName'),
    widgetId: text('widgetId'),
    widgetName: text('widgetName'),
    /** chat | run */
    filedVia: text('filedVia').notNull().default('chat'),
    /** The fixed window `lwql` replays over (epoch milliseconds, half-open) and its step. */
    replayStart: doublePrecision('replayStart'),
    replayEnd: doublePrecision('replayEnd'),
    replayGranularitySeconds: integer('replayGranularitySeconds'),
    /** The values in force when it was filed: `{ period, parameters }`. */
    replayContext: jsonb('replayContext'),
    /**
     * Whose insight it is: the one person who reads it. NULL on a row folded before the
     * column existed, which is owned by whoever filed it (`filedByUserId`).
     */
    ownerUserId: text('ownerUserId'),
    /**
     * The stable key a daily run matches a repeated finding by, so it renews the insight
     * and files no twin. NULL for an insight saved from a chat and on a row folded before
     * the column existed.
     */
    findingKey: text('findingKey'),
  },
  (table) => [
    index('InsightProjection_projectId_filedAt_idx').on(table.projectId, table.filedAt),
    index('InsightProjection_projectId_ownerUserId_filedAt_idx').on(
      table.projectId,
      table.ownerUserId,
      table.filedAt,
    ),
  ],
)

/**
 * One reader's own state on one insight (seen, done, kept), folded from the same stream.
 * An insight has one reader today, its owner; a later forward may add more.
 */
export const insightReaderProjection = pgTable(
  'InsightReaderProjection',
  {
    id: text('id').primaryKey(),
    projectId: text('projectId').notNull(),
    insightId: text('insightId').notNull(),
    userId: text('userId').notNull(),
    seenAt: doublePrecision('seenAt'),
    archivedAt: doublePrecision('archivedAt'),
    keptAt: doublePrecision('keptAt'),
    createdAt: doublePrecision('createdAt').notNull(),
    updatedAt: doublePrecision('updatedAt').notNull(),
    occurredAt: doublePrecision('occurredAt').notNull(),
    acceptedAt: doublePrecision('acceptedAt').notNull(),
    lastEventId: text('lastEventId').notNull(),
    projectionVersion: text('projectionVersion').notNull(),
  },
  (table) => [
    index('InsightReaderProjection_projectId_userId_idx').on(table.projectId, table.userId),
    uniqueIndex('InsightReaderProjection_projectId_insightId_userId_key').on(
      table.projectId,
      table.insightId,
      table.userId,
    ),
  ],
)

/**
 * One person's daily insights run on one board in a project, folded from its
 * `insight_daily_schedule` aggregate (modules/insight). A read model: rebuildable by
 * replaying the event log. Epochs in milliseconds.
 */
export const insightDailyScheduleProjection = pgTable(
  'InsightDailyScheduleProjection',
  {
    /** The schedule id, which is also its aggregate id: one per project, person and board. */
    id: text('id').primaryKey(),
    projectId: text('projectId').notNull(),
    userId: text('userId').notNull(),
    /** dashboard | template. A pointer only, never a relation: a run checks the board is there. */
    boardKind: text('boardKind').notNull(),
    boardId: text('boardId').notNull(),
    boardName: text('boardName').notNull(),
    /** on | off | undecided. */
    state: text('state').notNull().default('undecided'),
    /**
     * board_deleted | skipped_days | member_removed | project_archived. Why the system turned
     * it off. NULL while it is on, and when the person turned it off.
     */
    offReason: text('offReason'),
    hour: integer('hour'),
    timezone: text('timezone'),
    maxInsights: integer('maxInsights'),
    /** The last run that recorded an outcome. NULL until one did. */
    lastRunId: text('lastRunId'),
    lastRunAt: doublePrecision('lastRunAt'),
    /** filed | nothing | failed | skipped */
    lastRunOutcome: text('lastRunOutcome'),
    lastRunReason: text('lastRunReason'),
    lastRunFiled: integer('lastRunFiled'),
    /** How many insights the run found still true and renewed. */
    lastRunRenewed: integer('lastRunRenewed'),
    lastRunConversationId: text('lastRunConversationId'),
    /** The scheduled days in a row whose run ended skipped, and the slot of the newest counted. */
    consecutiveSkips: integer('consecutiveSkips').notNull().default(0),
    skipStreakSlot: doublePrecision('skipStreakSlot'),
    createdAt: doublePrecision('createdAt').notNull(),
    updatedAt: doublePrecision('updatedAt').notNull(),
    occurredAt: doublePrecision('occurredAt').notNull(),
    acceptedAt: doublePrecision('acceptedAt').notNull(),
    lastEventId: text('lastEventId').notNull(),
    projectionVersion: text('projectionVersion').notNull(),
  },
  (table) => [
    uniqueIndex('InsightDailySchedule_projectId_userId_board_key').on(
      table.projectId,
      table.userId,
      table.boardKind,
      table.boardId,
    ),
  ],
)

export const savedView = pgTable(
  'SavedView',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    projectId: text('projectId').notNull(),
    userId: text('userId'),
    name: text('name').notNull(),
    filters: jsonb('filters').notNull().default({}),
    query: text('query'),
    period: jsonb('period'),
    order: integer('order').notNull().default(0),
    /**
     * Discriminator for the two storage shapes that share this table:
     * - "v1-traces-filter" (default) - legacy /messages page filter map saved views
     * - "v2-traces-lens"             - new traces v2 lens configs (columns, addons, grouping, sort, filterText) encoded inside `filters` JSON
     * Each client filters by its own kind on read so neither sees the other's rows.
     */
    kind: text('kind').notNull().default('v1-traces-filter'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('SavedView_projectId_idx').on(table.projectId),
    index('SavedView_projectId_order_idx').on(table.projectId, table.order),
    index('SavedView_userId_idx').on(table.userId),
    index('SavedView_projectId_kind_order_idx').on(table.projectId, table.kind, table.order),
    foreignKey({
      name: 'SavedView_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'SavedView_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const customGraph = pgTable(
  'CustomGraph',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    projectId: text('projectId').notNull(),
    name: text('name').notNull(),
    /** The chart's definition, in the shape its `kind` names. */
    graph: jsonb('graph').notNull(),
    filters: jsonb('filters'),
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
    kind: text('kind').notNull().default('builder'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    dashboardId: text('dashboardId'),
    gridColumn: integer('gridColumn').notNull().default(0),
    gridRow: integer('gridRow').notNull().default(0),
    colSpan: integer('colSpan').notNull().default(1),
    rowSpan: integer('rowSpan').notNull().default(1),
  },
  (table) => [
    index('CustomGraph_projectId_idx').on(table.projectId),
    index('CustomGraph_dashboardId_idx').on(table.dashboardId),
    index('CustomGraph_projectId_kind_idx').on(table.projectId, table.kind),
    foreignKey({
      name: 'CustomGraph_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'CustomGraph_dashboardId_fkey',
      columns: [table.dashboardId],
      foreignColumns: [dashboard.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const batchEvaluation = pgTable(
  'BatchEvaluation',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    experimentId: text('experimentId').notNull(),
    projectId: text('projectId').notNull(),
    data: jsonb('data').notNull().default({}),
    status: text('status').notNull(),
    score: doublePrecision('score').notNull(),
    label: text('label'),
    passed: boolean('passed').notNull(),
    details: text('details').notNull(),
    cost: doublePrecision('cost').notNull(),
    datasetSlug: text('datasetSlug').notNull(),
    datasetId: text('datasetId').notNull(),
    evaluation: text('evaluation').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('BatchEvaluation_projectId_idx').on(table.projectId),
    index('BatchEvaluation_datasetId_idx').on(table.datasetId),
    index('BatchEvaluation_experimentId_idx').on(table.experimentId),
    uniqueIndex('BatchEvaluation_id_key').on(table.id),
    foreignKey({
      name: 'BatchEvaluation_experimentId_fkey',
      columns: [table.experimentId],
      foreignColumns: [experiment.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'BatchEvaluation_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'BatchEvaluation_datasetId_fkey',
      columns: [table.datasetId],
      foreignColumns: [dataset.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

export const trigger = pgTable(
  'Trigger',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    name: text('name').notNull(),
    projectId: text('projectId').notNull(),
    action: triggerActionEnum('action').notNull(),
    triggerKind: triggerKindEnum('triggerKind').notNull().default('AUTOMATION'),
    actionParams: jsonb('actionParams').notNull(),
    filters: jsonb('filters').notNull(),
    filterQuery: text('filterQuery'),
    lastRunAt: doublePrecision('lastRunAt').notNull().default(0),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    active: boolean('active').notNull().default(true),
    pausedReason: text('pausedReason'),
    pausedAt: timestamp('pausedAt', { withTimezone: true, precision: 3 }),
    message: text('message'),
    deleted: boolean('deleted').notNull().default(false),
    alertType: alertTypeEnum('alertType'),
    slackTemplateType: text('slackTemplateType'),
    slackTemplate: text('slackTemplate'),
    emailSubjectTemplate: text('emailSubjectTemplate'),
    emailBodyTemplate: text('emailBodyTemplate'),
    notificationCadence: text('notificationCadence').notNull().default('immediate'),
    traceDebounceMs: integer('traceDebounceMs').notNull().default(30000),
    customGraphId: text('customGraphId'),
  },
  (table) => [
    uniqueIndex('Trigger_customGraphId_key').on(table.customGraphId),
    index('Trigger_projectId_idx').on(table.projectId),
    foreignKey({
      name: 'Trigger_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Trigger_customGraphId_fkey',
      columns: [table.customGraphId],
      foreignColumns: [customGraph.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * One named Slack connection (ADR-093 §5a); an organization holds as many as
 * it likes, each scoped to the organization or to one project. The secret is
 * AES-256-GCM ciphertext and never returned to a client. `secretFingerprint`
 * is an HMAC of the plaintext, so one secret is one connection per scope
 * without decrypting anything. No foreign key: `scopeId` is
 * polymorphic (ADR-021), as on RetentionPolicy and CustomLLMModelCost.
 */
export const slackIntegration = pgTable(
  'SlackIntegration',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    name: text('name').notNull(),
    kind: slackIntegrationKindEnum('kind').notNull(),
    scopeType: slackIntegrationScopeTypeEnum('scopeType').notNull(),
    scopeId: text('scopeId').notNull(),
    organizationId: text('organizationId').notNull(),
    botTokenEncrypted: text('botTokenEncrypted'),
    webhookUrlEncrypted: text('webhookUrlEncrypted'),
    secretFingerprint: text('secretFingerprint').notNull(),
    secretHint: text('secretHint').notNull(),
    slackTeamId: text('slackTeamId'),
    slackTeamName: text('slackTeamName'),
    createdById: text('createdById').notNull(),
    updatedById: text('updatedById').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('SlackIntegration_scopeType_scopeId_idx').on(table.scopeType, table.scopeId),
    uniqueIndex('SlackIntegration_organizationId_scopeType_scopeId_secretFin_key').on(
      table.organizationId,
      table.scopeType,
      table.scopeId,
      table.secretFingerprint,
    ),
  ],
)

/**
 * An automation holding a Slack connection in use (ARCHITECTURE.md §3, the
 * claims ruling): a claimed connection is never deleted. `projectId` is the
 * claimant's project; `organizationId` is the connection's tenancy anchor.
 * No foreign key: slack's service refuses to delete a claimed connection.
 */
export const slackConnectionClaim = pgTable(
  'slack_connection_claim',
  {
    connectionId: text('connectionId').notNull(),
    claimantId: text('claimantId').notNull(),
    claimantLabel: text('claimantLabel').notNull(),
    organizationId: text('organizationId').notNull(),
    projectId: text('projectId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    primaryKey({
      name: 'slack_connection_claim_pkey',
      columns: [table.connectionId, table.claimantId],
    }),
    index('slack_connection_claim_organizationId_idx').on(table.organizationId),
  ],
)

/**
 * Org-anchored outbound webhook endpoint (the webhook platform, enterprise
 * only). The signing secret is AES-256-GCM encrypted at rest (shared
 * encrypt/decrypt, CREDENTIALS_SECRET) and returned in plaintext exactly
 * once, at create or roll time. `enabledEvents` holds registry event types
 * or family wildcards ("gateway.*", "*"). Failure-streak columns drive the
 * 72h auto-disable. Spec: specs/webhooks/webhook-endpoints.feature.
 */
export const webhookEndpoint = pgTable(
  'WebhookEndpoint',
  {
    id: text('id').primaryKey(),
    organizationId: text('organizationId').notNull(),
    /**
     * Where this endpoint delivers. Everything above the last hop is shared:
     * same batching, same retry ladder, same delivery log, same signature over
     * the same bytes.
     */
    destinationKind: webhookDestinationKindEnum('destinationKind').notNull().default('http'),
    /**
     * The receiver URL. Present on `http` endpoints and null on the others;
     * the per-kind shape is pinned by a CHECK constraint Prisma cannot model.
     */
    url: text('url'),
    /**
     * The Amazon SQS queue this endpoint delivers to, on `sqs` endpoints. The
     * region and the owning account are read off it rather than configured
     * separately, so they can never disagree with the queue.
     */
    sqsQueueUrl: text('sqsQueueUrl'),
    /**
     * The role we assume to write to the queue, and the ExternalId that proves
     * the request came from us rather than from anyone who learned the role's
     * name. The recommended credential mode.
     */
    sqsRoleArn: text('sqsRoleArn'),
    sqsExternalId: text('sqsExternalId'),
    /**
     * A static key pair for the queue, encrypted at rest exactly like the
     * signing secret. Both halves are present or both are absent: a half-filled
     * pair is what the AWS SDK reads as "we have credentials" and stops looking
     * with, so the CHECK constraint refuses it.
     */
    sqsAccessKeyId: text('sqsAccessKeyId'),
    sqsSecretAccessKeyEncrypted: text('sqsSecretAccessKeyEncrypted'),
    secretEncrypted: text('secretEncrypted').notNull(),
    /**
     * How deliveries are formatted and signed. Null: the batch envelope signed
     * `t=,v1=`. `legacy_sha256`: one raw message per POST signed `sha256=`, kept
     * only for endpoints migrated from governance anomaly destinations.
     */
    signatureScheme: text('signatureScheme'),
    /**
     * Set only by governance's anomaly destination migration: a create that
     * repeats it in the same organization answers the endpoint it made. An
     * archived endpoint gives its key up, so a retry after cleanup creates anew.
     */
    idempotencyKey: text('idempotencyKey'),
    /**
     * The secret this endpoint signed with before the last roll. Deliveries
     * carry a second `v1=` computed from it until it expires, so a receiver
     * swaps secrets on its own schedule instead of dropping deliveries in the
     * gap. Null once the window closes or if the secret was never rolled.
     */
    previousSecretEncrypted: text('previousSecretEncrypted'),
    /** When the previous secret stops being signed with and accepted. */
    previousSecretExpiresAt: timestamp('previousSecretExpiresAt', { precision: 3 }),
    enabledEvents: text('enabledEvents').array().default([]),
    status: webhookEndpointStatusEnum('status').notNull().default('ACTIVE'),
    /** Why the endpoint is disabled: "auto_failures_72h" or "manual". */
    disabledReason: text('disabledReason'),
    disabledAt: timestamp('disabledAt', { precision: 3 }),
    /**
     * Start of the current consecutive-failure streak; cleared on any 2xx.
     * 72 hours of unbroken failures from this instant auto-disables.
     */
    failingSince: timestamp('failingSince', { precision: 3 }),
    /**
     * Delivery controls, each server-bounded (see webhookDeliveryControls
     * constants): envelopes per POST, coalescing window before a partial
     * batch ships, and concurrent POSTs to this receiver.
     */
    maxBatchSize: integer('maxBatchSize').notNull().default(100),
    maxBatchDelayMs: integer('maxBatchDelayMs').notNull().default(250),
    maxInFlight: integer('maxInFlight').notNull().default(4),
    /**
     * Self-hosted only: deliveries accept the receiver's self-signed TLS
     * certificate. Off by default; the hosted product ignores it.
     */
    allowSelfSignedCertificate: boolean('allowSelfSignedCertificate').notNull().default(false),
    lastSuccessAt: timestamp('lastSuccessAt', { precision: 3 }),
    lastFailureAt: timestamp('lastFailureAt', { precision: 3 }),
    archivedAt: timestamp('archivedAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('WebhookEndpoint_organizationId_status_idx').on(table.organizationId, table.status),
    uniqueIndex('WebhookEndpoint_organizationId_idempotencyKey_key').on(
      table.organizationId,
      table.idempotencyKey,
    ),
    foreignKey({
      name: 'WebhookEndpoint_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const webhookEndpointDelivery = pgTable(
  'WebhookEndpointDelivery',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    channel: webhookDeliveryChannelEnum('channel').notNull().default('platform'),
    /** Platform tenancy. Null on automations rows. */
    organizationId: text('organizationId'),
    endpointId: text('endpointId'),
    /** Automations tenancy. Null on platform rows. */
    projectId: text('projectId'),
    triggerId: text('triggerId'),
    dispatchId: text('dispatchId').notNull(),
    /**
     * 1-based attempt and the batch's envelope count. Both null on automations
     * rows, which record neither: one trigger fire is one event, and the
     * channel's sender does not thread its attempt number into the log.
     */
    attempt: integer('attempt'),
    eventCount: integer('eventCount'),
    responseStatus: integer('responseStatus'),
    latencyMs: integer('latencyMs'),
    error: text('error'),
    response: jsonb('response'),
    outcome: webhookDeliveryOutcomeEnum('outcome').notNull(),
    firedAt: timestamp('firedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('WebhookEndpointDelivery_endpointId_firedAt_idx').on(table.endpointId, table.firedAt),
    index('WebhookEndpointDelivery_organizationId_firedAt_idx').on(
      table.organizationId,
      table.firedAt,
    ),
    index('WebhookEndpointDelivery_firedAt_idx').on(table.firedAt),
    index('WebhookEndpointDelivery_endpointId_dispatchId_idx').on(
      table.endpointId,
      table.dispatchId,
    ),
    index('WebhookEndpointDelivery_projectId_triggerId_firedAt_idx').on(
      table.projectId,
      table.triggerId,
      table.firedAt,
    ),
    index('WebhookEndpointDelivery_projectId_triggerId_dispatchId_idx').on(
      table.projectId,
      table.triggerId,
      table.dispatchId,
    ),
    foreignKey({
      name: 'WebhookEndpointDelivery_endpointId_fkey',
      columns: [table.endpointId],
      foreignColumns: [webhookEndpoint.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'WebhookEndpointDelivery_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'WebhookEndpointDelivery_triggerId_fkey',
      columns: [table.triggerId],
      foreignColumns: [trigger.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const experiment = pgTable(
  'Experiment',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    name: text('name'),
    type: experimentTypeEnum('type').notNull(),
    slug: text('slug').notNull(),
    projectId: text('projectId').notNull(),
    workflowId: text('workflowId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    archivedAt: timestamp('archivedAt', { precision: 3 }),
    workbenchState: jsonb('workbenchState'),
    workbenchVersion: integer('workbenchVersion').notNull().default(0),
  },
  (table) => [
    index('Experiment_projectId_idx').on(table.projectId),
    index('Experiment_workflowId_idx').on(table.workflowId),
    index('Experiment_projectId_archivedAt_idx').on(table.projectId, table.archivedAt),
    uniqueIndex('Experiment_projectId_slug_key').on(table.projectId, table.slug),
    foreignKey({
      name: 'Experiment_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Experiment_workflowId_fkey',
      columns: [table.workflowId],
      foreignColumns: [workflow.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const experimentVersion = pgTable(
  'ExperimentVersion',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    experimentId: text('experimentId').notNull(),
    projectId: text('projectId').notNull(),
    /**
     * The number the version is addressed and displayed by. Numbered rows hold
     * a gap-free sequence. The rolling autosave row rides on the counter
     * instead, which is what keeps it clear of that sequence.
     */
    version: integer('version').notNull(),
    /**
     * `Experiment.workbenchVersion` at the moment this row was last written.
     * The rolling row is rewritten in place, so this is the only column that
     * says how recent its content is, and the version list sorts and pages by
     * it. It also joins a row to the live state: the row whose `counterVersion`
     * equals the experiment's counter holds what the workbench shows now.
     */
    counterVersion: integer('counterVersion').notNull(),
    autoSaved: boolean('autoSaved').notNull().default(false),
    commitMessage: text('commitMessage'),
    authorId: text('authorId'),
    /** Who wrote this version: "user", "langy" or "api". */
    authorLabel: text('authorLabel').notNull(),
    /**
     * The run that wrote this version, when a run wrote it.
     *
     * A run writes its cells into the workbench state, which advances the
     * counter. Without this, the page that started that run reads the bump as
     * somebody else's write, refuses to save, and asks the reader to reload
     * over their own unsaved edits. Named here, the page adopts the version and
     * carries on.
     */
    runId: text('runId'),
    state: jsonb('state').notNull(),
    /** Schema version of `state`, so a later shape change can be read forward. */
    schemaVersion: text('schemaVersion').notNull().default('1'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('ExperimentVersion_projectId_experimentId_autoSaved_idx').on(
      table.projectId,
      table.experimentId,
      table.autoSaved,
    ),
    index('ExperimentVersion_projectId_experimentId_version_idx').on(
      table.projectId,
      table.experimentId,
      table.version,
    ),
    index('ExperimentVersion_projectId_experimentId_counterVersion_idx').on(
      table.projectId,
      table.experimentId,
      table.counterVersion,
    ),
    index('ExperimentVersion_createdAt_idx').on(table.createdAt),
    uniqueIndex('ExperimentVersion_experimentId_version_key').on(table.experimentId, table.version),
  ],
)

export const annotation = pgTable(
  'Annotation',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    projectId: text('projectId').notNull(),
    traceId: text('traceId').notNull(),
    comment: text('comment'),
    isThumbsUp: boolean('isThumbsUp'),
    userId: text('userId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    email: text('email'),
    scoreOptions: jsonb('scoreOptions'),
    expectedOutput: text('expectedOutput'),
    anchorKind: text('anchorKind'),
    anchorId: text('anchorId'),
    anchorPath: text('anchorPath'),
  },
  (table) => [
    index('Annotation_projectId_idx').on(table.projectId),
    index('Annotation_traceId_idx').on(table.traceId),
    index('Annotation_userId_idx').on(table.userId),
    index('Annotation_projectId_traceId_idx').on(table.projectId, table.traceId),
    index('Annotation_projectId_traceId_anchorKind_idx').on(
      table.projectId,
      table.traceId,
      table.anchorKind,
    ),
    foreignKey({
      name: 'Annotation_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Annotation_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const modelProvider = pgTable(
  'ModelProvider',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    name: text('name').notNull(),
    provider: text('provider').notNull(),
    routingHandle: text('routingHandle'),
    enabled: boolean('enabled').notNull(),
    customKeys: jsonb('customKeys'),
    extraHeaders: jsonb('extraHeaders'),
    customModels: jsonb('customModels'),
    customEmbeddingsModels: jsonb('customEmbeddingsModels'),
    deploymentMapping: jsonb('deploymentMapping'),
    rateLimitRpm: integer('rateLimitRpm'),
    rateLimitTpm: integer('rateLimitTpm'),
    rateLimitRpd: integer('rateLimitRpd'),
    rotationPolicy: gatewayProviderRotationPolicyEnum('rotationPolicy').notNull().default('MANUAL'),
    providerConfig: jsonb('providerConfig'),
    fallbackPriorityGlobal: integer('fallbackPriorityGlobal'),
    langySkipPermissionsModels: jsonb('langySkipPermissionsModels'),
    healthStatus: gatewayProviderHealthStatusEnum('healthStatus').notNull().default('UNKNOWN'),
    circuitOpenedAt: timestamp('circuitOpenedAt', { precision: 3 }),
    lastHealthCheckAt: timestamp('lastHealthCheckAt', { precision: 3 }),
    disabledAt: timestamp('disabledAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    organizationId: text('organizationId').notNull(),
  },
  (table) => [
    index('ModelProvider_healthStatus_idx').on(table.healthStatus),
    index('ModelProvider_fallbackPriorityGlobal_idx').on(table.fallbackPriorityGlobal),
    index('ModelProvider_organizationId_idx').on(table.organizationId),
    uniqueIndex('ModelProvider_organizationId_routingHandle_key').on(
      table.organizationId,
      table.routingHandle,
    ),
  ],
)

export const modelProviderScope = pgTable(
  'ModelProviderScope',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    modelProviderId: text('modelProviderId').notNull(),
    scopeType: modelProviderScopeTypeEnum('scopeType').notNull(),
    scopeId: text('scopeId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('ModelProviderScope_scopeType_scopeId_idx').on(table.scopeType, table.scopeId),
    index('ModelProviderScope_modelProviderId_idx').on(table.modelProviderId),
    uniqueIndex('ModelProviderScope_modelProviderId_scopeType_scopeId_key').on(
      table.modelProviderId,
      table.scopeType,
      table.scopeId,
    ),
    foreignKey({
      name: 'ModelProviderScope_modelProviderId_fkey',
      columns: [table.modelProviderId],
      foreignColumns: [modelProvider.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const modelDefaultConfig = pgTable(
  'ModelDefaultConfig',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    config: jsonb('config').notNull(),
    authorId: text('authorId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    organizationId: text('organizationId').notNull(),
  },
  (table) => [index('ModelDefaultConfig_organizationId_idx').on(table.organizationId)],
)

export const modelDefaultConfigScope = pgTable(
  'ModelDefaultConfigScope',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    configId: text('configId').notNull(),
    scopeType: modelDefaultScopeTypeEnum('scopeType').notNull(),
    scopeId: text('scopeId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('ModelDefaultConfigScope_scopeType_scopeId_idx').on(table.scopeType, table.scopeId),
    index('ModelDefaultConfigScope_configId_idx').on(table.configId),
    uniqueIndex('ModelDefaultConfigScope_configId_scopeType_scopeId_key').on(
      table.configId,
      table.scopeType,
      table.scopeId,
    ),
    foreignKey({
      name: 'ModelDefaultConfigScope_configId_fkey',
      columns: [table.configId],
      foreignColumns: [modelDefaultConfig.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const githubInstallation = pgTable(
  'GithubInstallation',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    /** GitHub installation id (numeric; stored as string). Globally unique. */
    installationId: text('installationId').notNull(),
    organizationId: text('organizationId').notNull(),
    /** The GitHub account (organization or user) the App is installed on. */
    accountLogin: text('accountLogin').notNull(),
    /** "Organization" or "User". */
    accountType: text('accountType').notNull(),
    /** GitHub numeric account id (stable across renames). */
    accountId: text('accountId').notNull(),
    /**
     * "all" or "selected" — whether the App can see every repo on the account
     * or only a chosen subset.
     */
    repositorySelection: text('repositorySelection').notNull(),
    /**
     * Cached repository list [{ id, fullName }] when repositorySelection =
     * "selected". Null when "all" (the full set is resolved live at mint time).
     */
    repositories: jsonb('repositories'),
    /**
     * Set when GitHub suspends the installation; cleared on unsuspend. A
     * suspended installation cannot mint tokens.
     */
    suspendedAt: timestamp('suspendedAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('GithubInstallation_installationId_key').on(table.installationId),
    index('GithubInstallation_organizationId_idx').on(table.organizationId),
    foreignKey({
      name: 'GithubInstallation_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const githubPullRequest = pgTable(
  'GithubPullRequest',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    /** The GitHub host the repository lives on, e.g. "github.com". */
    repositoryHost: text('repositoryHost').notNull(),
    /**
     * "owner/repo", stored LOWERCASED so a lookup matches whatever casing the
     * session reported.
     */
    repositoryFullName: text('repositoryFullName').notNull(),
    /** The branch the pull request was opened from. */
    headBranch: text('headBranch').notNull(),
    prNumber: integer('prNumber').notNull(),
    htmlUrl: text('htmlUrl').notNull(),
    title: text('title').notNull(),
    /** GitHub's own state string: "open" or "closed". */
    state: text('state').notNull(),
    isDraft: boolean('isDraft').notNull().default(false),
    authorLogin: text('authorLogin'),
    prCreatedAt: timestamp('prCreatedAt', { precision: 3 }).notNull(),
    prClosedAt: timestamp('prClosedAt', { precision: 3 }),
    prMergedAt: timestamp('prMergedAt', { precision: 3 }),
    /**
     * GitHub's own `updated_at` for the snapshot this row holds, and the
     * freshness guard on every write to it: a snapshot older than the stored one
     * is skipped. GitHub may deliver webhooks out of order, and without the
     * guard a late `opened` lands after `closed` and reopens a merged pull
     * request. NULL means the row predates the column, so any write is accepted.
     */
    prUpdatedAt: timestamp('prUpdatedAt', { precision: 3 }),
    /** When this row was first mapped to its branch. */
    mappedAt: timestamp('mappedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    /** When the snapshot columns were last refreshed from GitHub. */
    lastCheckedAt: timestamp('lastCheckedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('GithubPullRequest_organizationId_repositoryHost_repositoryF_idx').on(
      table.organizationId,
      table.repositoryHost,
      table.repositoryFullName,
      table.headBranch,
    ),
    uniqueIndex('GithubPullRequest_organizationId_repositoryHost_repositoryF_key').on(
      table.organizationId,
      table.repositoryHost,
      table.repositoryFullName,
      table.prNumber,
    ),
    foreignKey({
      name: 'GithubPullRequest_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * The per-branch bookkeeping behind pull-request mapping: when a branch was
 * last looked up on GitHub, how many pull requests it resolved to, and when it
 * may be asked about again. A branch with no pull request is the common case,
 * so `notFoundAt` plus `recheckAfter` is what keeps an unmapped branch from
 * costing a GitHub call on every read.
 * Spec: specs/integrations/github-connection.feature.
 */
export const githubBranchPullRequestCheck = pgTable(
  'GithubBranchPullRequestCheck',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    repositoryHost: text('repositoryHost').notNull(),
    /** "owner/repo", stored LOWERCASED, matching GithubPullRequest. */
    repositoryFullName: text('repositoryFullName').notNull(),
    headBranch: text('headBranch').notNull(),
    lastCheckedAt: timestamp('lastCheckedAt', { precision: 3 }).notNull(),
    prCount: integer('prCount').notNull().default(0),
    /**
     * Set while the branch resolves to no pull request; cleared once one is
     * found.
     */
    notFoundAt: timestamp('notFoundAt', { precision: 3 }),
    /** The earliest instant GitHub may be asked about this branch again. */
    recheckAfter: timestamp('recheckAfter', { precision: 3 }),
    /** Consecutive lookups that found nothing, backing off the recheck. */
    attempts: integer('attempts').notNull().default(0),
    /**
     * When a reader last asked about this branch, whether or not GitHub was
     * called.
     */
    lastRequestedAt: timestamp('lastRequestedAt', { precision: 3 }).notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('GithubBranchPullRequestCheck_recheckAfter_lastRequestedAt_n_idx').on(
      table.recheckAfter,
      table.lastRequestedAt,
      table.notFoundAt,
    ),
    uniqueIndex('GithubBranchPullRequestCheck_organizationId_repositoryHost__key').on(
      table.organizationId,
      table.repositoryHost,
      table.repositoryFullName,
      table.headBranch,
    ),
    foreignKey({
      name: 'GithubBranchPullRequestCheck_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * Rebuildable operational state folded from Langy's canonical event stream.
 * The source of truth remains the ClickHouse event log.
 */
export const langyConversationProjection = pgTable(
  'LangyConversationProjection',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    projectId: text('projectId').notNull(),
    ConversationId: text('conversationId').notNull(),
    /** Owner; first-writer-wins from the first message's userId. */
    UserId: text('userId').notNull(),
    Title: text('title'),
    /** derived | auto | user — governs auto-title precedence (user is sticky). */
    TitleSource: langyProjectionTitleSourceEnum('titleSource').notNull(),
    /**
     * Who started the conversation: `interactive` for the person in the panel, `run` for a
     * module acting for them with nobody at the keyboard. A row folded before the column
     * existed reads as interactive.
     */
    Origin: text('origin').notNull().default('interactive'),
    Status: text('status').notNull(),
    IsShared: boolean('isShared').notNull().default(false),
    SharedAt: doublePrecision('sharedAt'),
    SharedById: text('sharedById'),
    MessageCount: integer('messageCount').notNull().default(0),
    LastActivityAt: doublePrecision('lastActivityAt'),
    /** The turn in flight, or null when idle (busy guard reads this). */
    CurrentTurnId: text('currentTurnId'),
    LastError: text('lastError'),
    /** The model the latest accepted turn ran on; reopening seeds the picker from it. */
    LastModel: text('lastModel'),
    /** ADR-048 shutdown-handoff resume token (opaque, worker-authored). */
    PendingHandoffToken: text('pendingHandoffToken'),
    PendingHandoffTurnId: text('pendingHandoffTurnId'),
    /** Per-conversation HMAC key for worker stream frames; server-side only. */
    RunToken: text('runToken'),
    ArchivedAt: doublePrecision('archivedAt'),
    /** Messages CI started on a pull request since the person last wrote. */
    CiFollowUpsInARow: integer('ciFollowUpsInARow').notNull().default(0),
    CreatedAt: doublePrecision('createdAt').notNull(),
    UpdatedAt: doublePrecision('updatedAt').notNull(),
    /** Business time of the latest applied event. */
    OccurredAt: doublePrecision('occurredAt').notNull(),
    /** Event-log acceptance time. Together with lastEventId this is the cursor. */
    AcceptedAt: doublePrecision('acceptedAt').notNull(),
    LastEventId: text('lastEventId').notNull(),
    ProjectionVersion: text('projectionVersion').notNull(),
  },
  (table) => [
    index('LangyConversationProjection_projectId_userId_conversationId_idx').on(
      table.projectId,
      table.UserId,
      table.ConversationId,
    ),
    index('LangyConversationProjection_projectId_userId_lastActivityAt_idx').on(
      table.projectId,
      table.UserId,
      table.LastActivityAt,
    ),
    uniqueIndex('LangyConversationProjection_projectId_conversationId_key').on(
      table.projectId,
      table.ConversationId,
    ),
  ],
)

/**
 * Durable idempotency receipt for one logical Langy send. A transport retry
 * reuses the same requestId and receives the original conversation/turn ids
 * instead of minting a second message or turn.
 */
export const langyTurnRequest = pgTable(
  'LangyTurnRequest',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    projectId: text('projectId').notNull(),
    userId: text('userId').notNull(),
    requestId: text('requestId').notNull(),
    conversationId: text('conversationId').notNull(),
    turnId: text('turnId').notNull(),
    status: text('status').notNull(),
    leaseOwner: text('leaseOwner').notNull(),
    leaseExpiresAt: timestamp('leaseExpiresAt', { precision: 3 }).notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('LangyTurnRequest_projectId_conversationId_idx').on(
      table.projectId,
      table.conversationId,
    ),
    uniqueIndex('LangyTurnRequest_projectId_userId_requestId_key').on(
      table.projectId,
      table.userId,
      table.requestId,
    ),
  ],
)

/**
 * Postgres-authoritative one-turn admission for a conversation. PREPARING
 * claims are leased so a crash before the durable start event can recover;
 * COMMITTED claims are released only by the matching terminal event.
 */
export const langyActiveTurn = pgTable(
  'LangyActiveTurn',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    projectId: text('projectId').notNull(),
    conversationId: text('conversationId').notNull(),
    turnId: text('turnId').notNull(),
    requestId: text('requestId').notNull(),
    userId: text('userId').notNull(),
    status: text('status').notNull(),
    leaseOwner: text('leaseOwner').notNull(),
    leaseExpiresAt: timestamp('leaseExpiresAt', { precision: 3 }).notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('LangyActiveTurn_projectId_userId_idx').on(table.projectId, table.userId),
    uniqueIndex('LangyActiveTurn_projectId_conversationId_key').on(
      table.projectId,
      table.conversationId,
    ),
  ],
)

/** Rebuildable per-turn operational render document. */
export const langyConversationTurnProjection = pgTable(
  'LangyConversationTurnProjection',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    projectId: text('projectId').notNull(),
    ConversationId: text('conversationId').notNull(),
    TurnId: text('turnId').notNull(),
    /**
     * LangyConversationTurnStatus. Deliberately TEXT, not a Postgres enum: this
     * is a PROJECTION — one writer (the fold), rebuildable from the event log,
     * and nothing queries or indexes on it. An enum bought no selectivity here,
     * only typo protection, and charged a migration for every new status that
     * Postgres then could not remove. The one definition lives in
     * `LANGY_CONVERSATION_TURN_STATUS`; the repository validates against it.
     */
    Status: text('status').notNull(),
    /** LangyMessagePart[] — the user question that opened the turn. */
    QuestionParts: jsonb('questionParts').notNull(),
    /** LangyMessagePart[] — the agent's whole final answer. */
    AnswerParts: jsonb('answerParts').notNull(),
    /** LangyTurnToolCall[] — tool-call lifecycle audit in initiation order. */
    ToolCalls: jsonb('toolCalls').notNull(),
    /** LangyPlanItemData[] | null — null means the turn never kept a plan. */
    Plan: jsonb('plan'),
    Error: text('error'),
    StartedAt: doublePrecision('startedAt'),
    EndedAt: doublePrecision('endedAt'),
    CreatedAt: doublePrecision('createdAt').notNull(),
    UpdatedAt: doublePrecision('updatedAt').notNull(),
    OccurredAt: doublePrecision('occurredAt').notNull(),
    AcceptedAt: doublePrecision('acceptedAt').notNull(),
    LastEventId: text('lastEventId').notNull(),
    ProjectionVersion: text('projectionVersion').notNull(),
  },
  (table) => [
    index('LangyConversationTurnProjection_projectId_conversationId_up_idx').on(
      table.projectId,
      table.ConversationId,
      table.UpdatedAt,
    ),
    uniqueIndex('LangyConversationTurnProjection_projectId_conversationId_tu_key').on(
      table.projectId,
      table.ConversationId,
      table.TurnId,
    ),
  ],
)

/**
 * One pull request Langy opened, from the answer whose `gh pr create` printed it. Finished
 * checks look it up by repository and number to wake the conversation; the daily pull request
 * cap counts a person's rows by day. Rebuilt by a projection replay.
 */
export const langyPullRequest = pgTable(
  'LangyPullRequest',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    projectId: text('projectId').notNull(),
    /** `owner/repo`, lowercased: GitHub matches repository names without case. */
    RepositoryFullName: text('repositoryFullName').notNull(),
    PrNumber: integer('prNumber').notNull(),
    ConversationId: text('conversationId').notNull(),
    /** The conversation's owner, who opened it through Langy. */
    UserId: text('userId').notNull(),
    /** When the answer that opened it was recorded (epoch ms). */
    OpenedAt: doublePrecision('openedAt').notNull(),
    /** The agent_responded event it came from. */
    SourceEventId: text('sourceEventId').notNull(),
  },
  (table) => [
    index('LangyPullRequest_userId_openedAt_idx').on(table.UserId, table.OpenedAt),
    uniqueIndex('LangyPullRequest_repositoryFullName_prNumber_key').on(
      table.RepositoryFullName,
      table.PrNumber,
    ),
  ],
)

/** Rebuildable operational message fact. One source event produces one row. */
export const langyMessageProjection = pgTable(
  'LangyMessageProjection',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    projectId: text('projectId').notNull(),
    ConversationId: text('conversationId').notNull(),
    MessageId: text('messageId').notNull(),
    Role: text('role').notNull(),
    Parts: jsonb('parts').notNull(),
    SourceEventId: text('sourceEventId').notNull(),
    OccurredAt: doublePrecision('occurredAt').notNull(),
    AcceptedAt: doublePrecision('acceptedAt').notNull(),
    CreatedAt: doublePrecision('createdAt').notNull(),
    UpdatedAt: doublePrecision('updatedAt').notNull(),
  },
  (table) => [
    index('LangyMessageProjection_projectId_conversationId_createdAt_idx').on(
      table.projectId,
      table.ConversationId,
      table.CreatedAt,
    ),
    uniqueIndex('LangyMessageProjection_projectId_conversationId_messageId_key').on(
      table.projectId,
      table.ConversationId,
      table.MessageId,
    ),
    uniqueIndex('LangyMessageProjection_projectId_sourceEventId_key').on(
      table.projectId,
      table.SourceEventId,
    ),
  ],
)

export const projectSecret = pgTable(
  'ProjectSecret',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    projectId: text('projectId').notNull(),
    name: text('name').notNull(),
    encryptedValue: text('encryptedValue').notNull(),
    createdById: text('createdById').notNull(),
    updatedById: text('updatedById').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('ProjectSecret_projectId_idx').on(table.projectId),
    uniqueIndex('ProjectSecret_projectId_name_key').on(table.projectId, table.name),
    foreignKey({
      name: 'ProjectSecret_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'ProjectSecret_createdById_fkey',
      columns: [table.createdById],
      foreignColumns: [user.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'ProjectSecret_updatedById_fkey',
      columns: [table.updatedById],
      foreignColumns: [user.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

export const triggerSent = pgTable(
  'TriggerSent',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    triggerId: text('triggerId').notNull(),
    traceId: text('traceId'),
    customGraphId: text('customGraphId'),
    projectId: text('projectId').notNull(),
    resolvedAt: timestamp('resolvedAt', { precision: 3 }),
    openIncidentKey: text('openIncidentKey'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('TriggerSent_openIncidentKey_key').on(table.openIncidentKey),
    index('TriggerSent_triggerId_idx').on(table.triggerId),
    index('TriggerSent_projectId_idx').on(table.projectId),
    index('TriggerSent_customGraphId_idx').on(table.customGraphId),
    index('TriggerSent_resolvedAt_idx').on(table.resolvedAt),
    index('TriggerSent_projectId_triggerId_createdAt_id_idx').on(
      table.projectId,
      table.triggerId,
      table.createdAt,
      table.id,
    ),
    uniqueIndex('TriggerSent_triggerId_traceId_key').on(table.triggerId, table.traceId),
    foreignKey({
      name: 'TriggerSent_customGraphId_fkey',
      columns: [table.customGraphId],
      foreignColumns: [customGraph.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'TriggerSent_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'TriggerSent_triggerId_fkey',
      columns: [table.triggerId],
      foreignColumns: [trigger.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

export const triggerLatestEvaluation = pgTable(
  'TriggerLatestEvaluation',
  {
    triggerId: text('triggerId').primaryKey(),
    projectId: text('projectId').notNull(),
    evaluatedAt: timestamp('evaluatedAt', { withTimezone: true, precision: 3 }).notNull(),
    verdict: text('verdict').notNull(),
    observedValue: doublePrecision('observedValue'),
    threshold: doublePrecision('threshold'),
    operator: text('operator'),
    timePeriodMinutes: integer('timePeriodMinutes'),
    skipCode: text('skipCode'),
    updatedAt: timestamp('updatedAt', { withTimezone: true, precision: 3 })
      .notNull()
      .$onUpdate(utcNow),
  },
  (table) => [
    index('TriggerLatestEvaluation_projectId_idx').on(table.projectId),
    foreignKey({
      name: 'TriggerLatestEvaluation_triggerId_fkey',
      columns: [table.triggerId],
      foreignColumns: [trigger.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const emailSuppression = pgTable(
  'EmailSuppression',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    projectId: text('projectId').notNull(),
    email: text('email').notNull(),
    triggerId: text('triggerId'),
    reason: text('reason').notNull().default('unsubscribe'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('EmailSuppression_projectId_idx').on(table.projectId),
    uniqueIndex('EmailSuppression_projectId_email_triggerId_key').on(
      table.projectId,
      table.email,
      table.triggerId,
    ),
    foreignKey({
      name: 'EmailSuppression_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

export const annotationScore = pgTable(
  'AnnotationScore',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    name: text('name').notNull(),
    projectId: text('projectId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    deletedAt: timestamp('deletedAt', { precision: 3 }),
    description: text('description'),
    active: boolean('active').notNull().default(true),
    dataType: annotationScoreDataTypeEnum('dataType'),
    options: jsonb('options'),
    defaultValue: jsonb('defaultValue'),
    global: boolean('global').notNull().default(false),
  },
  (table) => [
    index('AnnotationScore_projectId_idx').on(table.projectId),
    foreignKey({
      name: 'AnnotationScore_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

export const annotationQueue = pgTable(
  'AnnotationQueue',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    name: text('name').notNull(),
    slug: text('slug').notNull(),
    projectId: text('projectId').notNull(),
    description: text('description'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('AnnotationQueue_projectId_idx').on(table.projectId),
    uniqueIndex('AnnotationQueue_projectId_slug_key').on(table.projectId, table.slug),
    foreignKey({
      name: 'AnnotationQueue_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

export const annotationQueueMembers = pgTable(
  'AnnotationQueueMembers',
  { annotationQueueId: text('annotationQueueId').notNull(), userId: text('userId').notNull() },
  (table) => [
    primaryKey({
      name: 'AnnotationQueueMembers_pkey',
      columns: [table.annotationQueueId, table.userId],
    }),
    index('AnnotationQueueMembers_annotationQueueId_idx').on(table.annotationQueueId),
    index('AnnotationQueueMembers_userId_idx').on(table.userId),
    foreignKey({
      name: 'AnnotationQueueMembers_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AnnotationQueueMembers_annotationQueueId_fkey',
      columns: [table.annotationQueueId],
      foreignColumns: [annotationQueue.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

export const annotationQueueScores = pgTable(
  'AnnotationQueueScores',
  {
    annotationQueueId: text('annotationQueueId').notNull(),
    annotationScoreId: text('annotationScoreId').notNull(),
  },
  (table) => [
    primaryKey({
      name: 'AnnotationQueueScores_pkey',
      columns: [table.annotationQueueId, table.annotationScoreId],
    }),
    index('AnnotationQueueScores_annotationQueueId_idx').on(table.annotationQueueId),
    index('AnnotationQueueScores_annotationScoreId_idx').on(table.annotationScoreId),
    foreignKey({
      name: 'AnnotationQueueScores_annotationScoreId_fkey',
      columns: [table.annotationScoreId],
      foreignColumns: [annotationScore.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AnnotationQueueScores_annotationQueueId_fkey',
      columns: [table.annotationQueueId],
      foreignColumns: [annotationQueue.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

export const annotationQueueItem = pgTable(
  'AnnotationQueueItem',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    annotationQueueId: text('annotationQueueId'),
    userId: text('userId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    createdByUserId: text('createdByUserId'),
    traceId: text('traceId').notNull(),
    projectId: text('projectId').notNull(),
    doneAt: timestamp('doneAt', { precision: 3 }),
    /**
     * Set while the annotator wants this item's trace added to a dataset once
     * the queue walk is finished. Persisted rather than kept in the browser so
     * the selection survives a refresh and stays visible to teammates.
     */
    markedForDatasetAt: timestamp('markedForDatasetAt', { precision: 3 }),
  },
  (table) => [
    index('AnnotationQueueItem_annotationQueueId_idx').on(table.annotationQueueId),
    index('AnnotationQueueItem_userId_idx').on(table.userId),
    index('AnnotationQueueItem_projectId_idx').on(table.projectId),
    index('AnnotationQueueItem_projectId_doneAt_idx').on(table.projectId, table.doneAt),
    index('AnnotationQueueItem_projectId_userId_doneAt_idx').on(
      table.projectId,
      table.userId,
      table.doneAt,
    ),
    index('AnnotationQueueItem_annotationQueueId_doneAt_idx').on(
      table.annotationQueueId,
      table.doneAt,
    ),
    uniqueIndex('AnnotationQueueItem_traceId_userId_projectId_key').on(
      table.traceId,
      table.userId,
      table.projectId,
    ),
    uniqueIndex('AnnotationQueueItem_traceId_annotationQueueId_projectId_key').on(
      table.traceId,
      table.annotationQueueId,
      table.projectId,
    ),
    foreignKey({
      name: 'AnnotationQueueItem_annotationQueueId_fkey',
      columns: [table.annotationQueueId],
      foreignColumns: [annotationQueue.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AnnotationQueueItem_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AnnotationQueueItem_createdByUserId_fkey',
      columns: [table.createdByUserId],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AnnotationQueueItem_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

export const shareLink = pgTable(
  'ShareLink',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    token: text('token').notNull(),
    resourceType: shareResourceTypeEnum('resourceType').notNull(),
    resourceId: text('resourceId').notNull(),
    threadId: text('threadId'),
    projectId: text('projectId').notNull(),
    userId: text('userId'),
    visibility: shareVisibilityEnum('visibility').notNull().default('PUBLIC'),
    expiresAt: timestamp('expiresAt', { precision: 3 }),
    maxViews: integer('maxViews'),
    viewCount: integer('viewCount').notNull().default(0),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('ShareLink_token_key').on(table.token),
    index('ShareLink_projectId_resourceType_resourceId_idx').on(
      table.projectId,
      table.resourceType,
      table.resourceId,
    ),
    index('ShareLink_userId_idx').on(table.userId),
    foreignKey({
      name: 'ShareLink_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'ShareLink_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const pinnedTrace = pgTable(
  'PinnedTrace',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    projectId: text('projectId').notNull(),
    traceId: text('traceId').notNull(),
    userId: text('userId'),
    source: pinSourceEnum('source').notNull().default('manual'),
    reason: text('reason'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('PinnedTrace_projectId_idx').on(table.projectId),
    uniqueIndex('PinnedTrace_projectId_traceId_key').on(table.projectId, table.traceId),
    foreignKey({
      name: 'PinnedTrace_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'PinnedTrace_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const traceEditOverlay = pgTable(
  'TraceEditOverlay',
  {
    id: text('id').primaryKey(),
    projectId: text('projectId').notNull(),
    traceId: text('traceId').notNull(),
    patch: jsonb('patch').notNull(),
    createdById: text('createdById'),
    updatedById: text('updatedById'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('TraceEditOverlay_projectId_idx').on(table.projectId),
    uniqueIndex('TraceEditOverlay_projectId_traceId_key').on(table.projectId, table.traceId),
    foreignKey({
      name: 'TraceEditOverlay_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'TraceEditOverlay_createdById_fkey',
      columns: [table.createdById],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'TraceEditOverlay_updatedById_fkey',
      columns: [table.updatedById],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const traceIngestSourceBilling = pgTable(
  'TraceIngestSourceBilling',
  {
    organizationId: text('organizationId').notNull(),
    sourceType: text('sourceType').notNull(),
    billed: boolean('billed').notNull(),
    recordedAt: timestamp('recordedAt', { precision: 3 }).notNull(),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    primaryKey({
      name: 'TraceIngestSourceBilling_pkey',
      columns: [table.organizationId, table.sourceType],
    }),
  ],
)

export const retentionPolicy = pgTable(
  'RetentionPolicy',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    scopeType: retentionPolicyScopeTypeEnum('scopeType').notNull(),
    scopeId: text('scopeId').notNull(),
    category: text('category').notNull(),
    retentionDays: integer('retentionDays').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('RetentionPolicy_organizationId_idx').on(table.organizationId),
    index('RetentionPolicy_scopeType_scopeId_idx').on(table.scopeType, table.scopeId),
    uniqueIndex('RetentionPolicy_scopeType_scopeId_category_key').on(
      table.scopeType,
      table.scopeId,
      table.category,
    ),
  ],
)

export const retentionApplyStatus = pgTable(
  'RetentionApplyStatus',
  {
    id: text('id').primaryKey(),
    organizationId: text('organizationId').notNull(),
    scopeType: retentionPolicyScopeTypeEnum('scopeType').notNull(),
    scopeId: text('scopeId').notNull(),
    changedAt: doublePrecision('changedAt'),
    appliedAt: doublePrecision('appliedAt'),
    raisedAt: doublePrecision('raisedAt'),
    shortenedAt: doublePrecision('shortenedAt'),
    createdAt: doublePrecision('createdAt').notNull(),
    updatedAt: doublePrecision('updatedAt').notNull(),
    occurredAt: doublePrecision('occurredAt').notNull(),
    acceptedAt: doublePrecision('acceptedAt').notNull(),
    lastEventId: text('lastEventId').notNull(),
    projectionVersion: text('projectionVersion').notNull(),
  },
  (table) => [
    uniqueIndex('RetentionApplyStatus_organizationId_scopeType_scopeId_key').on(
      table.organizationId,
      table.scopeType,
      table.scopeId,
    ),
  ],
)

export const dataPrivacyPolicy = pgTable(
  'DataPrivacyPolicy',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    scopeType: dataPrivacyScopeTypeEnum('scopeType').notNull(),
    scopeId: text('scopeId').notNull(),
    personalOnly: boolean('personalOnly').notNull().default(false),
    config: jsonb('config').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('DataPrivacyPolicy_organizationId_idx').on(table.organizationId),
    index('DataPrivacyPolicy_scopeType_scopeId_idx').on(table.scopeType, table.scopeId),
    uniqueIndex('DataPrivacyPolicy_scopeType_scopeId_personalOnly_key').on(
      table.scopeType,
      table.scopeId,
      table.personalOnly,
    ),
  ],
)

export const instantEvalJudgeProject = pgTable('InstantEvalJudgeProject', {
  projectId: text('projectId').primaryKey(),
  organizationId: text('organizationId').notNull(),
  createdAt: timestamp('createdAt', { precision: 3 }).notNull(),
})

export const instantEvalJudgeUsageBilling = pgTable('InstantEvalJudgeUsageBilling', {
  organizationId: text('organizationId').primaryKey(),
  usageBilled: boolean('usageBilled').notNull(),
  occurredAt: timestamp('occurredAt', { precision: 3 }).notNull(),
  fromCatchUp: boolean('fromCatchUp').notNull(),
})

export const instantEvalJudgeSpend = pgTable(
  'InstantEvalJudgeSpend',
  {
    organizationId: text('organizationId').notNull(),
    requestId: text('requestId').notNull(),
    spendNanoUsd: bigint('spendNanoUsd', { mode: 'bigint' }).notNull(),
    occurredAt: timestamp('occurredAt', { precision: 3 }).notNull(),
  },
  (table) => [
    primaryKey({
      name: 'InstantEvalJudgeSpend_pkey',
      columns: [table.organizationId, table.requestId],
    }),
  ],
)

export const customLLMModelCost = pgTable(
  'CustomLLMModelCost',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    scopeType: customLlmmodelCostScopeTypeEnum('scopeType').notNull().default('PROJECT'),
    scopeId: text('scopeId').notNull(),
    projectId: text('projectId'),
    model: text('model').notNull(),
    regex: text('regex').notNull(),
    inputCostPerToken: doublePrecision('inputCostPerToken'),
    outputCostPerToken: doublePrecision('outputCostPerToken'),
    cacheReadCostPerToken: doublePrecision('cacheReadCostPerToken'),
    cacheCreationCostPerToken: doublePrecision('cacheCreationCostPerToken'),
    cacheCreation1hCostPerToken: doublePrecision('cacheCreation1hCostPerToken'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('CustomLLMModelCost_organizationId_idx').on(table.organizationId),
    index('CustomLLMModelCost_scopeType_scopeId_idx').on(table.scopeType, table.scopeId),
    index('CustomLLMModelCost_projectId_idx').on(table.projectId),
  ],
)

export const workflow = pgTable(
  'Workflow',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    projectId: text('projectId').notNull(),
    name: text('name').notNull(),
    icon: text('icon').notNull(),
    description: text('description').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    latestVersionId: text('latestVersionId'),
    currentVersionId: text('currentVersionId'),
    publishedId: text('publishedId'),
    publishedById: text('publishedById'),
    archivedAt: timestamp('archivedAt', { precision: 3 }),
    isEvaluator: boolean('isEvaluator').notNull().default(false),
    isComponent: boolean('isComponent').notNull().default(false),
    copiedFromWorkflowId: text('copiedFromWorkflowId'),
  },
  (table): PgTableExtraConfigValue[] => [
    index('Workflow_projectId_idx').on(table.projectId),
    index('Workflow_publishedById_idx').on(table.publishedById),
    index('Workflow_latestVersionId_idx').on(table.latestVersionId),
    index('Workflow_currentVersionId_idx').on(table.currentVersionId),
    index('Workflow_copiedFromWorkflowId_idx').on(table.copiedFromWorkflowId),
    foreignKey({
      name: 'Workflow_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Workflow_latestVersionId_fkey',
      columns: [table.latestVersionId],
      foreignColumns: [workflowVersion.id],
    })
      .onDelete('restrict')
      .onUpdate('restrict'),
    foreignKey({
      name: 'Workflow_currentVersionId_fkey',
      columns: [table.currentVersionId],
      foreignColumns: [workflowVersion.id],
    })
      .onDelete('restrict')
      .onUpdate('restrict'),
    foreignKey({
      name: 'Workflow_publishedById_fkey',
      columns: [table.publishedById],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Workflow_copiedFromWorkflowId_fkey',
      columns: [table.copiedFromWorkflowId],
      foreignColumns: [table.id],
    })
      .onDelete('restrict')
      .onUpdate('restrict'),
  ],
)

export const workflowVersion = pgTable(
  'WorkflowVersion',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    version: text('version').notNull(),
    commitMessage: text('commitMessage').notNull(),
    authorId: text('authorId').notNull(),
    projectId: text('projectId').notNull(),
    workflowId: text('workflowId').notNull(),
    parentId: text('parentId'),
    autoSaved: boolean('autoSaved').notNull().default(false),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    dsl: jsonb('dsl').notNull(),
  },
  (table): PgTableExtraConfigValue[] => [
    index('WorkflowVersion_workflowId_idx').on(table.workflowId),
    index('WorkflowVersion_projectId_idx').on(table.projectId),
    index('WorkflowVersion_parentId_idx').on(table.parentId),
    index('WorkflowVersion_authorId_idx').on(table.authorId),
    index('WorkflowVersion_createdAt_idx').on(table.createdAt),
    uniqueIndex('WorkflowVersion_workflowId_version_key').on(table.workflowId, table.version),
    foreignKey({
      name: 'WorkflowVersion_authorId_fkey',
      columns: [table.authorId],
      foreignColumns: [user.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'WorkflowVersion_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'WorkflowVersion_workflowId_fkey',
      columns: [table.workflowId],
      foreignColumns: [workflow.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'WorkflowVersion_parentId_fkey',
      columns: [table.parentId],
      foreignColumns: [table.id],
    })
      .onDelete('restrict')
      .onUpdate('restrict'),
  ],
)

export const auditLog = pgTable(
  'AuditLog',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    userId: text('userId'),
    projectId: text('projectId'),
    organizationId: text('organizationId'),
    action: text('action').notNull(),
    args: jsonb('args'),
    error: text('error'),
    ipAddress: text('ipAddress'),
    userAgent: text('userAgent'),
    metadata: jsonb('metadata'),
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
    actorUserId: text('actorUserId'),
    targetKind: text('targetKind'),
    targetId: text('targetId'),
    before: jsonb('before'),
    after: jsonb('after'),
    /** A producer's outbox key: a redelivered audit intent finds its row here (Alex, audit R2). */
    idempotencyKey: text('idempotencyKey'),
  },
  (table) => [
    uniqueIndex('AuditLog_idempotencyKey_key').on(table.idempotencyKey),
    index('AuditLog_createdAt_idx').on(table.createdAt),
    index('AuditLog_userId_idx').on(table.userId),
    index('AuditLog_actorUserId_idx').on(table.actorUserId),
    index('AuditLog_projectId_idx').on(table.projectId),
    index('AuditLog_action_idx').on(table.action),
    index('AuditLog_ipAddress_idx').on(table.ipAddress),
    index('AuditLog_organizationId_createdAt_idx').on(table.organizationId, table.createdAt),
    index('AuditLog_targetKind_targetId_idx').on(table.targetKind, table.targetId),
  ],
)

export const llmPromptConfig = pgTable(
  'LlmPromptConfig',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    handle: text('handle'),
    name: text('name').notNull(),
    projectId: text('projectId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    deletedAt: timestamp('deletedAt', { precision: 3 }),
    organizationId: text('organizationId').notNull(),
    scope: promptScopeEnum('scope').notNull().default('PROJECT'),
    copiedFromPromptId: text('copiedFromPromptId'),
  },
  (table) => [
    index('LlmPromptConfig_projectId_idx').on(table.projectId),
    index('LlmPromptConfig_organizationId_scope_idx').on(table.organizationId, table.scope),
    index('LlmPromptConfig_copiedFromPromptId_idx').on(table.copiedFromPromptId),
    uniqueIndex('LlmPromptConfig_handle_key').on(table.handle),
    foreignKey({
      name: 'LlmPromptConfig_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'LlmPromptConfig_copiedFromPromptId_fkey',
      columns: [table.copiedFromPromptId],
      foreignColumns: [table.id],
    })
      .onDelete('restrict')
      .onUpdate('restrict'),
  ],
)

export const llmPromptConfigVersion = pgTable(
  'LlmPromptConfigVersion',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    version: integer('version').notNull(),
    commitMessage: text('commitMessage'),
    authorId: text('authorId'),
    configId: text('configId').notNull(),
    configData: jsonb('config').notNull(),
    schemaVersion: text('schemaVersion').notNull(),
    runtimeParameters: jsonb('runtimeParameters').notNull().default({}),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    projectId: text('projectId').notNull(),
  },
  (table) => [
    index('LlmPromptConfigVersion_configId_idx').on(table.configId),
    index('LlmPromptConfigVersion_authorId_idx').on(table.authorId),
    index('LlmPromptConfigVersion_createdAt_idx').on(table.createdAt),
    uniqueIndex('LlmPromptConfigVersion_configId_version_key').on(table.configId, table.version),
    foreignKey({
      name: 'LlmPromptConfigVersion_authorId_fkey',
      columns: [table.authorId],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'LlmPromptConfigVersion_configId_fkey',
      columns: [table.configId],
      foreignColumns: [llmPromptConfig.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const promptTagAssignment = pgTable(
  'PromptTagAssignment',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    configId: text('configId').notNull(),
    versionId: text('versionId').notNull(),
    tagId: text('tagId').notNull(),
    projectId: text('projectId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    createdById: text('createdById'),
    updatedById: text('updatedById'),
  },
  (table) => [
    index('PromptTagAssignment_configId_idx').on(table.configId),
    index('PromptTagAssignment_versionId_idx').on(table.versionId),
    index('PromptTagAssignment_projectId_idx').on(table.projectId),
    index('PromptTagAssignment_tagId_idx').on(table.tagId),
    uniqueIndex('PromptTagAssignment_configId_tagId_key').on(table.configId, table.tagId),
    foreignKey({
      name: 'PromptTagAssignment_configId_fkey',
      columns: [table.configId],
      foreignColumns: [llmPromptConfig.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'PromptTagAssignment_versionId_fkey',
      columns: [table.versionId],
      foreignColumns: [llmPromptConfigVersion.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'PromptTagAssignment_tagId_fkey',
      columns: [table.tagId],
      foreignColumns: [promptTag.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'PromptTagAssignment_createdById_fkey',
      columns: [table.createdById],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'PromptTagAssignment_updatedById_fkey',
      columns: [table.updatedById],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const promptTag = pgTable(
  'PromptTag',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    name: text('name').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    createdById: text('createdById'),
    updatedById: text('updatedById'),
  },
  (table) => [
    index('PromptTag_organizationId_idx').on(table.organizationId),
    uniqueIndex('PromptTag_organizationId_name_key').on(table.organizationId, table.name),
    foreignKey({
      name: 'PromptTag_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'PromptTag_createdById_fkey',
      columns: [table.createdById],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'PromptTag_updatedById_fkey',
      columns: [table.updatedById],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const analytics = pgTable(
  'Analytics',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    projectId: text('projectId').notNull(),
    key: text('key').notNull(),
    value: jsonb('value'),
    numericValue: doublePrecision('numericValue'),
    stringValue: text('stringValue'),
    boolValue: boolean('boolValue'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('Analytics_createdAt_key_projectId_numericValue_idx').on(
      table.createdAt,
      table.key,
      table.projectId,
      table.numericValue,
    ),
    index('Analytics_projectId_idx').on(table.projectId),
    foreignKey({
      name: 'Analytics_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

export const customRole = pgTable(
  'CustomRole',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    name: text('name').notNull(),
    description: text('description'),
    permissions: jsonb('permissions').notNull(),
    kind: text('kind').notNull().default('custom'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('CustomRole_organizationId_idx').on(table.organizationId),
    uniqueIndex('CustomRole_organizationId_name_key').on(table.organizationId, table.name),
    foreignKey({
      name: 'CustomRole_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

export const notification = pgTable(
  'Notification',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId'),
    projectId: text('projectId'),
    metadata: jsonb('metadata').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    sentAt: timestamp('sentAt', { precision: 3 }).notNull(),
  },
  (table) => [
    index('Notification_organizationId_idx').on(table.organizationId),
    index('Notification_projectId_idx').on(table.projectId),
    index('Notification_createdAt_idx').on(table.createdAt),
    index('Notification_sentAt_idx').on(table.sentAt),
    foreignKey({
      name: 'Notification_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Notification_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const agent = pgTable(
  'Agent',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    projectId: text('projectId').notNull(),
    name: text('name').notNull(),
    type: text('type').notNull(),
    config: jsonb('config').notNull(),
    workflowId: text('workflowId'),
    copiedFromAgentId: text('copiedFromAgentId'),
    environment: text('environment'),
    ownerUserId: text('ownerUserId'),
    hostLabel: text('hostLabel'),
    identityKey: text('identityKey'),
    lastSeenAt: timestamp('lastSeenAt', { precision: 3 }),
    archivedAt: timestamp('archivedAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    createdById: text('createdById'),
    updatedById: text('updatedById'),
  },
  (table) => [
    index('Agent_projectId_idx').on(table.projectId),
    index('Agent_workflowId_idx').on(table.workflowId),
    index('Agent_copiedFromAgentId_idx').on(table.copiedFromAgentId),
    index('Agent_projectId_lastSeenAt_idx').on(table.projectId, table.lastSeenAt),
    uniqueIndex('Agent_projectId_identityKey_key').on(table.projectId, table.identityKey),
    foreignKey({
      name: 'Agent_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Agent_workflowId_fkey',
      columns: [table.workflowId],
      foreignColumns: [workflow.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Agent_copiedFromAgentId_fkey',
      columns: [table.copiedFromAgentId],
      foreignColumns: [table.id],
    })
      .onDelete('restrict')
      .onUpdate('restrict'),
  ],
)

/** One immutable snapshot of an agent; the Agent row is always the latest. */
export const agentVersion = pgTable(
  'AgentVersion',
  {
    id: text('id').primaryKey(),
    projectId: text('projectId').notNull(),
    agentId: text('agentId').notNull(),
    version: integer('version').notNull(),
    name: text('name').notNull(),
    type: text('type').notNull(),
    config: jsonb('config').notNull(),
    workflowId: text('workflowId'),
    commitMessage: text('commitMessage'),
    authorId: text('authorId'),
    /** "user", "api_key" or "system": who wrote the version when authorId is null. */
    authorKind: text('authorKind'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('AgentVersion_projectId_idx').on(table.projectId),
    uniqueIndex('AgentVersion_agentId_version_key').on(table.agentId, table.version),
  ],
)

/** A deploy label ("production", "staging") pointing at one version of an agent. */
export const agentVersionLabel = pgTable(
  'AgentVersionLabel',
  {
    id: text('id').primaryKey(),
    projectId: text('projectId').notNull(),
    agentId: text('agentId').notNull(),
    label: text('label').notNull(),
    versionId: text('versionId').notNull(),
    createdById: text('createdById'),
    updatedById: text('updatedById'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('AgentVersionLabel_projectId_idx').on(table.projectId),
    index('AgentVersionLabel_versionId_idx').on(table.versionId),
    uniqueIndex('AgentVersionLabel_agentId_label_key').on(table.agentId, table.label),
  ],
)

export const evaluator = pgTable(
  'Evaluator',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    projectId: text('projectId').notNull(),
    name: text('name').notNull(),
    slug: text('slug'),
    type: text('type').notNull(),
    config: jsonb('config').notNull(),
    workflowId: text('workflowId'),
    copiedFromEvaluatorId: text('copiedFromEvaluatorId'),
    archivedAt: timestamp('archivedAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    createdById: text('createdById'),
    updatedById: text('updatedById'),
  },
  (table) => [
    index('Evaluator_projectId_idx').on(table.projectId),
    index('Evaluator_workflowId_idx').on(table.workflowId),
    index('Evaluator_copiedFromEvaluatorId_idx').on(table.copiedFromEvaluatorId),
    uniqueIndex('Evaluator_projectId_slug_key').on(table.projectId, table.slug),
    foreignKey({
      name: 'Evaluator_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Evaluator_workflowId_fkey',
      columns: [table.workflowId],
      foreignColumns: [workflow.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Evaluator_copiedFromEvaluatorId_fkey',
      columns: [table.copiedFromEvaluatorId],
      foreignColumns: [table.id],
    })
      .onDelete('restrict')
      .onUpdate('restrict'),
  ],
)

/** One immutable snapshot of an evaluator; the Evaluator row is always the latest. */
export const evaluatorVersion = pgTable(
  'EvaluatorVersion',
  {
    id: text('id').primaryKey(),
    projectId: text('projectId').notNull(),
    evaluatorId: text('evaluatorId').notNull(),
    version: integer('version').notNull(),
    name: text('name').notNull(),
    type: text('type').notNull(),
    config: jsonb('config').notNull(),
    workflowId: text('workflowId'),
    commitMessage: text('commitMessage'),
    authorId: text('authorId'),
    /** "user", "api_key" or "system": who wrote the version when authorId is null. */
    authorKind: text('authorKind'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('EvaluatorVersion_projectId_idx').on(table.projectId),
    uniqueIndex('EvaluatorVersion_evaluatorId_version_key').on(table.evaluatorId, table.version),
  ],
)

/** A deploy label ("production", "staging") pointing at one version of an evaluator. */
export const evaluatorVersionLabel = pgTable(
  'EvaluatorVersionLabel',
  {
    id: text('id').primaryKey(),
    projectId: text('projectId').notNull(),
    evaluatorId: text('evaluatorId').notNull(),
    label: text('label').notNull(),
    versionId: text('versionId').notNull(),
    createdById: text('createdById'),
    updatedById: text('updatedById'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('EvaluatorVersionLabel_projectId_idx').on(table.projectId),
    index('EvaluatorVersionLabel_versionId_idx').on(table.versionId),
    uniqueIndex('EvaluatorVersionLabel_evaluatorId_label_key').on(table.evaluatorId, table.label),
  ],
)

export const scenario = pgTable(
  'Scenario',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    projectId: text('projectId').notNull(),
    name: text('name').notNull(),
    situation: text('situation').notNull(),
    criteria: text('criteria').array(),
    labels: text('labels').array(),
    parameters: jsonb('parameters'),
    simulatorModel: text('simulatorModel'),
    judgeModel: text('judgeModel'),
    traceQuietPeriodMs: integer('traceQuietPeriodMs'),
    traceWaitTimeoutMs: integer('traceWaitTimeoutMs'),
    maxTurns: integer('maxTurns'),
    minTurns: integer('minTurns'),
    fields: jsonb('fields'),
    callerVoice: jsonb('callerVoice'),
    testSuiteId: text('testSuiteId'),
    kind: text('kind').notNull().default('standard'),
    redTeam: jsonb('redTeam'),
    sourceTemplateId: text('sourceTemplateId'),
    version: integer('version').notNull().default(1),
    lastUpdatedById: text('lastUpdatedById'),
    archivedAt: timestamp('archivedAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('Scenario_projectId_idx').on(table.projectId),
    index('Scenario_projectId_testSuiteId_idx').on(table.projectId, table.testSuiteId),
    foreignKey({
      name: 'Scenario_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Scenario_testSuiteId_fkey',
      columns: [table.testSuiteId],
      foreignColumns: [simulationSuite.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Scenario_lastUpdatedById_fkey',
      columns: [table.lastUpdatedById],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const scenarioVersion = pgTable(
  'ScenarioVersion',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    scenarioId: text('scenarioId').notNull(),
    /**
     * The project of the parent case, carried here so every read filters by
     * tenant without a join. The service is what keeps the pair honest: no
     * composite foreign key can, under relationMode="prisma".
     */
    projectId: text('projectId').notNull(),
    version: integer('version').notNull(),
    authorId: text('authorId'),
    /** Who wrote this version: "user", "api", "cli" or "langy". */
    authorLabel: text('authorLabel').notNull(),
    changeDescription: text('changeDescription'),
    /**
     * Envelope { schemaVersion, fields, changedFields } holding the editable
     * field set as saved and the names of the fields this save changed.
     */
    snapshot: jsonb('snapshot').notNull(),
    /** Schema version of `snapshot`, so a later shape change can be read forward. */
    schemaVersion: integer('schemaVersion').notNull().default(1),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('ScenarioVersion_projectId_scenarioId_version_idx').on(
      table.projectId,
      table.scenarioId,
      table.version,
    ),
    index('ScenarioVersion_createdAt_idx').on(table.createdAt),
    uniqueIndex('ScenarioVersion_scenarioId_version_key').on(table.scenarioId, table.version),
    foreignKey({
      name: 'ScenarioVersion_scenarioId_fkey',
      columns: [table.scenarioId],
      foreignColumns: [scenario.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const simulationSuite = pgTable(
  'SimulationSuite',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    projectId: text('projectId').notNull(),
    name: text('name').notNull(),
    slug: text('slug').notNull(),
    description: text('description'),
    scenarioIds: text('scenarioIds').array(),
    targets: jsonb('targets').notNull(),
    repeatCount: integer('repeatCount').notNull().default(1),
    labels: text('labels').array(),
    simulatorModel: text('simulatorModel'),
    judgeModel: text('judgeModel'),
    traceQuietPeriodMs: integer('traceQuietPeriodMs'),
    traceWaitTimeoutMs: integer('traceWaitTimeoutMs'),
    kind: text('kind').notNull().default('run_plan'),
    scope: jsonb('scope'),
    fields: jsonb('fields'),
    evaluators: jsonb('evaluators'),
    templateId: text('templateId'),
    templateVersion: integer('templateVersion'),
    archivedAt: timestamp('archivedAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('SimulationSuite_projectId_idx').on(table.projectId),
    index('SimulationSuite_projectId_kind_idx').on(table.projectId, table.kind),
    uniqueIndex('SimulationSuite_projectId_slug_key').on(table.projectId, table.slug),
    foreignKey({
      name: 'SimulationSuite_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

export const subscription = pgTable(
  'Subscription',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    plan: planTypesEnum('plan').notNull(),
    stripeSubscriptionId: text('stripeSubscriptionId'),
    status: subscriptionStatusEnum('status').notNull(),
    startDate: timestamp('startDate', { precision: 3 }),
    endDate: timestamp('endDate', { precision: 3 }),
    lastPaymentFailedDate: timestamp('lastPaymentFailedDate', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    maxMembers: integer('maxMembers'),
    maxProjects: integer('maxProjects'),
    maxMessagesPerMonth: integer('maxMessagesPerMonth'),
    maxWorkflows: integer('maxWorkflows'),
    maxRetentionDays: integer('maxRetentionDays'),
    maxMembersLite: integer('maxMembersLite'),
    /** Developer seats the organization pays for; null falls to the plan's own cap. */
    maxMembersDeveloper: integer('maxMembersDeveloper'),
    maxTeams: integer('maxTeams'),
    maxPrompts: integer('maxPrompts'),
    maxEvaluators: integer('maxEvaluators'),
    maxScenarios: integer('maxScenarios'),
    maxAgents: integer('maxAgents'),
    maxExperiments: integer('maxExperiments'),
    maxOnlineEvaluations: integer('maxOnlineEvaluations'),
    maxDatasets: integer('maxDatasets'),
    maxDashboards: integer('maxDashboards'),
    maxCustomGraphs: integer('maxCustomGraphs'),
    maxAutomations: integer('maxAutomations'),
    /** Seats paid for this period that stay usable after a seat decrease, until seatsKeptUntil. */
    seatsKept: integer('seatsKept'),
    seatsKeptUntil: timestamp('seatsKeptUntil', { precision: 3 }),
    /** The seat quantity the last finished seat change set the seat line to; provider events never write it. */
    seatLineQuantity: integer('seatLineQuantity'),
    seatLineQuantityAt: timestamp('seatLineQuantityAt', { precision: 3 }),
    /**
     * A Stripe seat line write recorded before it is sent and cleared once it is known to have
     * finished under the lease; one left behind is settled by the seat line repair.
     * Every seat quantity written while this record was pending.
     */
    seatLineWriteQuantities: integer('seatLineWriteQuantities').array().default([]),
    /** The first write still pending; a later holder's write keeps it. */
    seatLineWriteSince: timestamp('seatLineWriteSince', { precision: 3 }),
    /** The latest write sent; the repair waits a grace window after it. */
    seatLineWriteAt: timestamp('seatLineWriteAt', { precision: 3 }),
    seatLineWriteAttempts: integer('seatLineWriteAttempts'),
    seatLineWriteFlaggedAt: timestamp('seatLineWriteFlaggedAt', { precision: 3 }),
    /** The seat quantity Stripe should bill if a seat change stops after lifting the seat line; null once settled. */
    seatRepairQuantity: integer('seatRepairQuantity'),
    /** The quantity the change lifted the line to; the repair moves the line back only while Stripe still bills it. */
    seatRepairLiftedQuantity: integer('seatRepairLiftedQuantity'),
    seatRepairSince: timestamp('seatRepairSince', { precision: 3 }),
    seatRepairAttempts: integer('seatRepairAttempts'),
    /** Set once the seat line repair has failed often enough that Ops should look at it. */
    seatRepairFlaggedAt: timestamp('seatRepairFlaggedAt', { precision: 3 }),
    /** The lease a seat change or seat line repair holds while it reads and moves the seat line. */
    seatLineLeaseToken: text('seatLineLeaseToken'),
    seatLineLeaseUntil: timestamp('seatLineLeaseUntil', { precision: 3 }),
    /**
     * An open failed payment recovery case: the unpaid renewal invoice, day 0 (its first failed
     * charge) and the last notice sent (3, 7, 14 or 21). Cleared when the case closes.
     */
    paymentRecoveryInvoiceId: text('paymentRecoveryInvoiceId'),
    paymentRecoveryStartedAt: timestamp('paymentRecoveryStartedAt', { precision: 3 }),
    paymentRecoveryNoticeDay: integer('paymentRecoveryNoticeDay'),
    /**
     * The over-plan mark: traces recorded from overPlanSince are withheld while it is open and the
     * organization is on Free with more members than Free allows; ended once that no longer holds.
     */
    overPlanSince: timestamp('overPlanSince', { precision: 3 }),
    overPlanEndedAt: timestamp('overPlanEndedAt', { precision: 3 }),
  },
  (table) => [
    uniqueIndex('Subscription_stripeSubscriptionId_key').on(table.stripeSubscriptionId),
    index('Subscription_organizationId_idx').on(table.organizationId),
    foreignKey({
      name: 'Subscription_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

export const invoice = pgTable(
  'Invoice',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    subscriptionId: text('subscriptionId').notNull(),
    totalAmount: doublePrecision('totalAmount').notNull(),
    currency: text('currency').notNull(),
    dueDate: timestamp('dueDate', { precision: 3 }).notNull(),
    paid: boolean('paid').notNull().default(false),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('Invoice_subscriptionId_idx').on(table.subscriptionId),
    foreignKey({
      name: 'Invoice_subscriptionId_fkey',
      columns: [table.subscriptionId],
      foreignColumns: [subscription.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

export const invoiceItem = pgTable(
  'InvoiceItem',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    invoiceId: text('invoiceId').notNull(),
    description: text('description').notNull(),
    amount: doublePrecision('amount').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('InvoiceItem_invoiceId_idx').on(table.invoiceId),
    foreignKey({
      name: 'InvoiceItem_invoiceId_fkey',
      columns: [table.invoiceId],
      foreignColumns: [invoice.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

export const group = pgTable(
  'Group',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    name: text('name').notNull(),
    slug: text('slug').notNull(),
    externalId: text('externalId'),
    scimSource: text('scimSource'),
    /**
     * D08: which SSO connection provisioned this group. Null for manually
     * created groups and for the SCIM groups that predate connection scoping.
     */
    scimConnectionId: text('scimConnectionId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('Group_organizationId_idx').on(table.organizationId),
    index('Group_scimConnectionId_idx').on(table.scimConnectionId),
    uniqueIndex('Group_organizationId_slug_key').on(table.organizationId, table.slug),
    uniqueIndex('Group_scimConnectionId_externalId_key').on(
      table.scimConnectionId,
      table.externalId,
    ),
    foreignKey({
      name: 'Group_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const groupMembership = pgTable(
  'GroupMembership',
  {
    userId: text('userId').notNull(),
    groupId: text('groupId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    primaryKey({ name: 'GroupMembership_pkey', columns: [table.userId, table.groupId] }),
    index('GroupMembership_userId_idx').on(table.userId),
    index('GroupMembership_groupId_idx').on(table.groupId),
    foreignKey({
      name: 'GroupMembership_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'GroupMembership_groupId_fkey',
      columns: [table.groupId],
      foreignColumns: [group.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/** Who authz treats as gone, from user's and identity's facts; authz never reads User (ARCHITECTURE.md, "Platform operators are a grant"). */
export const authzUserStanding = pgTable('AuthzUserStanding', {
  userId: text('userId').primaryKey(),
  deactivatedAt: timestamp('deactivatedAt', { precision: 3 }),
  erasedAt: timestamp('erasedAt', { precision: 3 }),
  standingChangedAt: timestamp('standingChangedAt', { precision: 3 }).notNull(),
})

/**
 * Whether an organization's permission checks run on the authorization engine: authz's own
 * switch, set when the engine migration proves the organization and pinned off by an operator's
 * rollback. No row is the legacy path. Spec: modules/authz/specs/engine-switch.feature.
 */
export const authzEngineSwitch = pgTable('AuthzEngineSwitch', {
  organizationId: text('organizationId').primaryKey(),
  onEngine: boolean('onEngine').notNull(),
  pinned: boolean('pinned').notNull().default(false),
  switchedAt: timestamp('switchedAt', { precision: 3 }).notNull(),
  lastAppliedEventId: text('lastAppliedEventId'),
  updatedAt: timestamp('updatedAt', { precision: 3 })
    .notNull()
    .default(sql`CURRENT_TIMESTAMP`)
    .$defaultFn(utcNow)
    .$onUpdate(utcNow),
})

/** ADR-201: an organisation's braid setting. Owned by authz; read into the organisation's snapshot. */
export const authzOrganizationPolicy = pgTable('AuthzOrganizationPolicy', {
  organizationId: text('organizationId').primaryKey(),
  braidsOrganizationAdminsOnly: boolean('braidsOrganizationAdminsOnly').notNull().default(false),
  braidsEnabled: boolean('braidsEnabled').notNull().default(false),
  updatedByUserId: text('updatedByUserId').notNull(),
  updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
})

export const authzBraidWrite = pgTable(
  'AuthzBraidWrite',
  {
    organizationId: text('organizationId').notNull(),
    commandId: text('commandId').notNull(),
    grantId: text('grantId').notNull(),
    readerProjectId: text('readerProjectId').notNull(),
    ownerProjectId: text('ownerProjectId').notNull(),
    termsKey: text('termsKey').notNull(),
    condition: jsonb('condition').notNull(),
    expiresAt: timestamp('expiresAt', { precision: 3 }),
    replacesGrantId: text('replacesGrantId'),
    createdByUserId: text('createdByUserId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    primaryKey({ name: 'AuthzBraidWrite_pkey', columns: [table.organizationId, table.commandId] }),
    index('AuthzBraidWrite_organizationId_readerProjectId_createdAt_idx').on(
      table.organizationId,
      table.readerProjectId,
      table.createdAt,
    ),
  ],
)

export const roleBinding = pgTable(
  'RoleBinding',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    userId: text('userId'),
    groupId: text('groupId'),
    apiKeyId: text('apiKeyId'),
    role: teamUserRoleEnum('role').notNull(),
    customRoleId: text('customRoleId'),
    scopeType: roleBindingScopeTypeEnum('scopeType').notNull(),
    scopeId: text('scopeId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('RoleBinding_organizationId_idx').on(table.organizationId),
    index('RoleBinding_userId_idx').on(table.userId),
    index('RoleBinding_groupId_idx').on(table.groupId),
    index('RoleBinding_apiKeyId_idx').on(table.apiKeyId),
    index('RoleBinding_scopeType_scopeId_idx').on(table.scopeType, table.scopeId),
    foreignKey({
      name: 'RoleBinding_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'RoleBinding_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'RoleBinding_groupId_fkey',
      columns: [table.groupId],
      foreignColumns: [group.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'RoleBinding_apiKeyId_fkey',
      columns: [table.apiKeyId],
      foreignColumns: [apiKey.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'RoleBinding_customRoleId_fkey',
      columns: [table.customRoleId],
      foreignColumns: [customRole.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const grant = pgTable(
  'Grant',
  {
    id: text('id').primaryKey(),
    organizationId: text('organizationId').notNull(),
    principalType: grantPrincipalTypeEnum('principalType').notNull(),
    principalId: text('principalId'),
    roleKey: text('roleKey'),
    legacyRole: text('legacyRole'),
    source: text('source').notNull(),
    scopeType: grantScopeTypeEnum('scopeType').notNull(),
    scopeId: text('scopeId').notNull(),
    token: text('token'),
    permission: text('permission'),
    resourceKind: text('resourceKind'),
    projectId: text('projectId'),
    createdByUserId: text('createdByUserId'),
    expiresAt: timestamp('expiresAt', { precision: 3 }),
    maxViews: integer('maxViews'),
    condition: jsonb('condition'),
    replacesGrantId: text('replacesGrantId'),
    occurredAt: timestamp('occurredAt', { precision: 3 }).notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    revokedAt: timestamp('revokedAt', { precision: 3 }),
    revokedReason: text('revokedReason'),
  },
  (table) => [
    uniqueIndex('Grant_token_key').on(table.token),
    index('Grant_organizationId_scopeType_scopeId_idx').on(
      table.organizationId,
      table.scopeType,
      table.scopeId,
    ),
    index('Grant_projectId_resourceKind_scopeId_idx').on(
      table.projectId,
      table.resourceKind,
      table.scopeId,
    ),
  ],
)

export const grantUsage = pgTable(
  'GrantUsage',
  {
    grantId: text('grantId').primaryKey(),
    organizationId: text('organizationId').notNull(),
    projectId: text('projectId').notNull(),
    viewCount: integer('viewCount').notNull().default(0),
    lastViewedAt: timestamp('lastViewedAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('GrantUsage_organizationId_idx').on(table.organizationId),
    index('GrantUsage_projectId_idx').on(table.projectId),
  ],
)

export const role = pgTable(
  'Role',
  {
    id: text('id').primaryKey(),
    organizationId: text('organizationId').notNull(),
    name: text('name').notNull(),
    description: text('description'),
    permissions: jsonb('permissions').notNull(),
    kind: text('kind').notNull().default('custom'),
    occurredAt: timestamp('occurredAt', { precision: 3 }).notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    deletedAt: timestamp('deletedAt', { precision: 3 }),
  },
  (table) => [index('Role_organizationId_idx').on(table.organizationId)],
)

export const apiKey = pgTable(
  'ApiKey',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    name: text('name').notNull(),
    description: text('description'),
    lookupId: text('lookupId').notNull(),
    hashedSecret: text('hashedSecret').notNull(),
    permissionMode: text('permissionMode').notNull().default('all'),
    userId: text('userId'),
    createdByUserId: text('createdByUserId'),
    organizationId: text('organizationId').notNull(),
    expiresAt: timestamp('expiresAt', { precision: 3 }),
    lastUsedAt: timestamp('lastUsedAt', { precision: 3 }),
    revokedAt: timestamp('revokedAt', { precision: 3 }),
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
    revocationCause: text('revocationCause'),
    /**
     * The CLI login key of the device session that minted this ingestion key.
     * Set by the CLI's personal mint, NULL for every other key. Revoking the
     * parent (logout, the devices tab, a re-login from the same device, session
     * expiry) revokes the children with it. Column ref only, no FK and no
     * index: lookups filter by the indexed `userId` first and match the parent
     * in memory over a person's few live ingest keys.
     */
    parentApiKeyId: text('parentApiKeyId'),
    /**
     * Non-null marks this key as an "ingestion key": a project-scoped,
     * ingest-only ApiKey the unified `langwatch <tool>` CLI (and template
     * installs) mint for OTLP Path B. The slug (claude_code / codex / gemini /
     * opencode / claude_cowork) is stamped as `langwatch.source` provenance on
     * every span the key authorizes. NULL for normal API / service keys.
     */
    ingestSourceType: text('ingestSourceType'),
    /**
     * Set only for ingestion keys derived from an admin IngestionTemplate
     * (e.g. claude_cowork). Drives receiver OTTL application + the
     * `langwatch.template.id` provenance stamp. NULL for unified-CLI tools
     * (which have no template) and for normal keys. Column ref only (no FK,
     * per the app-layer-integrity convention); resolved by the ingest-key service.
     */
    ingestionTemplateId: text('ingestionTemplateId'),
    /**
     * Human label of the CLI device session that minted this ingestion key
     * (e.g. "Rogerio's MacBook Pro"), snapshotted at mint time from the device
     * flow's `client_info`. Denormalized so the org-wide API-keys settings page
     * can show which device a key came from without resolving another user's
     * Redis-stored sessions. NULL for non-ingestion keys and for keys minted by
     * a CLI old enough to predate device metadata.
     */
    createdByDeviceLabel: text('createdByDeviceLabel'),
    /**
     * True for a key LangWatch minted for itself (a run, the gateway's trace
     * export, a Langy session). Hides it from listings, refuses customer edits
     * and revokes, and lets an ownerless run key act as the system. Never set
     * by a customer create; rows before the column stay false.
     */
    isSystemManaged: boolean('isSystemManaged').notNull().default(false),
    /**
     * The browser session that minted a docs session key (ADR-185), so signing
     * out of that session revokes it. NULL for every other key. Unindexed:
     * sign-out filters by the indexed `userId` first.
     */
    createdByBrowserSessionId: text('createdByBrowserSessionId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('ApiKey_lookupId_key').on(table.lookupId),
    index('ApiKey_userId_idx').on(table.userId),
    index('ApiKey_organizationId_idx').on(table.organizationId),
    index('ApiKey_lookupId_idx').on(table.lookupId),
    index('ApiKey_ingestSourceType_idx').on(table.ingestSourceType),
    index('ApiKey_ingestionTemplateId_idx').on(table.ingestionTemplateId),
    foreignKey({ name: 'ApiKey_userId_fkey', columns: [table.userId], foreignColumns: [user.id] })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'ApiKey_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

/** ADR-201: a project's key for signing the tags on its calls. Owned by the API key module. */
export const signingKey = pgTable(
  'SigningKey',
  {
    id: text('id').primaryKey(),
    organizationId: text('organizationId').notNull(),
    projectId: text('projectId').notNull(),
    name: text('name').notNull(),
    publicId: text('publicId').notNull(),
    secretCiphertext: text('secretCiphertext').notNull(),
    createdByUserId: text('createdByUserId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    rotatedAt: timestamp('rotatedAt', { precision: 3 }),
    verifiesUntil: timestamp('verifiesUntil', { precision: 3 }),
    revokedAt: timestamp('revokedAt', { precision: 3 }),
    lastUsedAt: timestamp('lastUsedAt', { precision: 3 }),
  },
  (table) => [
    uniqueIndex('SigningKey_publicId_key').on(table.publicId),
    index('SigningKey_projectId_idx').on(table.projectId),
    index('SigningKey_organizationId_idx').on(table.organizationId),
  ],
)

export const projectActiveDay = pgTable(
  'ProjectActiveDay',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    projectId: text('projectId').notNull(),
    day: text('day').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [uniqueIndex('ProjectActiveDay_projectId_day_key').on(table.projectId, table.day)],
)

/**
 * Nurturing's read model: which organization a project belongs to, from project's created event.
 * The per-email usage places a project and lists an organization's projects from these rows only.
 */
export const nurturingProject = pgTable(
  'NurturingProject',
  { projectId: text('projectId').primaryKey(), organizationId: text('organizationId').notNull() },
  (table) => [index('NurturingProject_organizationId_idx').on(table.organizationId)],
)

/**
 * A project nurturing saw traces of with no placement copy, and when it first did: every worker
 * retries its trace batches only within one window from then, across restarts. Deleted when the
 * copy lands.
 */
export const nurturingUncopiedProject = pgTable('NurturingUncopiedProject', {
  projectId: text('projectId').primaryKey(),
  firstSeenAt: timestamp('firstSeenAt', { precision: 3 })
    .notNull()
    .default(sql`CURRENT_TIMESTAMP`)
    .$defaultFn(utcNow),
})

/**
 * Nurturing's own copy of each organization's members, from the owners' member facts (§9):
 * organization's signed_up and invite_accepted and auth's sso_auto_added join, member_removed
 * and member_disabled leave, member_enabled returns. Each column keeps its latest time.
 */
export const nurturingOrganizationMember = pgTable(
  'NurturingOrganizationMember',
  {
    organizationId: text('organizationId').notNull(),
    userId: text('userId').notNull(),
    joinedAt: timestamp('joinedAt', { precision: 3 }),
    leftAt: timestamp('leftAt', { precision: 3 }),
    returnedAt: timestamp('returnedAt', { precision: 3 }),
  },
  (table) => [
    primaryKey({
      name: 'NurturingOrganizationMember_pkey',
      columns: [table.organizationId, table.userId],
    }),
    index('NurturingOrganizationMember_userId_idx').on(table.userId),
  ],
)

/**
 * When an organization was created, from organization's created event: nurturing's own copy for
 * the per-email usage's day offsets. Its own table, so an early creation never adds an unseeded
 * NurturingOrganization row.
 */
export const nurturingOrganizationCreated = pgTable(
  'NurturingOrganizationCreated',
  {
    organizationId: text('organizationId').primaryKey(),
    createdAt: timestamp('createdAt', { precision: 3 }).notNull(),
  },
  (table) => [index('NurturingOrganizationCreated_createdAt_idx').on(table.createdAt)],
)

/**
 * Nurturing's read model per organization: its admin and its milestone counts.
 * `seeded` marks an organization learned from the project backfill: its first_* already happened.
 */
export const nurturingOrganization = pgTable('NurturingOrganization', {
  organizationId: text('organizationId').primaryKey(),
  adminUserId: text('adminUserId'),
  seeded: boolean('seeded').notNull().default(false),
  evaluationCount: integer('evaluationCount').notNull().default(0),
  simulationRunCount: integer('simulationRunCount').notNull().default(0),
  codingFirstSessionAt: timestamp('codingFirstSessionAt', { precision: 3 }),
  codingAgentTypes: text('codingAgentTypes').array().default([]),
  codingActiveDays: integer('codingActiveDays').notNull().default(0),
  codingLastActiveDay: text('codingLastActiveDay'),
  codingTerminalTypes: text('codingTerminalTypes').array().default([]),
  codingRepositoryKeys: text('codingRepositoryKeys').array().default([]),
  codingFirstDays: integer('codingFirstDays').array().default([]),
  lastTraceAt: timestamp('lastTraceAt', { precision: 3 }),
  updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
})

/**
 * Nurturing's own per-project usage: a peer fold over the owners' prompt, evaluator, scenario
 * and simulation-run events (ARCHITECTURE §9), one peer fold per owner pipeline, summed per
 * organization for the per-email traits. Traces are counted in NurturingProjectTraceDay.
 */
export const nurturingProjectUsageView = pgTable(
  'NurturingProjectUsageView',
  {
    projectId: text('projectId').notNull(),
    /**
     * The owner pipeline whose peer fold wrote the row (prompt, evaluator, scenario, simulation),
     * or `evaluation`: the project's evaluation days, written by a subscriber, at most once a day.
     */
    owner: text('owner').notNull(),
    /** The project's usage as the fold folded it (NurturingProjectUsageState). */
    folded: jsonb('folded').notNull(),
    /** The event ids of the last delivery folded in, so a redelivery folds once. */
    appliedEventIds: text('appliedEventIds').array().default([]),
    /** A row of another version reads as undecodable and is re-folded from the log. */
    projectionVersion: text('projectionVersion').notNull(),
    /** Event-derived, so a replayed row equals the row it rebuilds. */
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull(),
  },
  (table) => [
    primaryKey({ name: 'NurturingProjectUsageView_pkey', columns: [table.projectId, table.owner] }),
  ],
)

/**
 * Nurturing's own trace counts from trace's span_received and trace_failed (ARCHITECTURE §9), a
 * root span being one trace, one row per project and UTC day of append time, incremented in place
 * by each batch. A redelivered batch can count a trace twice; accepted (issue 9378). Kept 30 days.
 */
export const nurturingProjectTraceDay = pgTable(
  'NurturingProjectTraceDay',
  {
    projectId: text('projectId').notNull(),
    /** UTC `YYYY-MM-DD` of the root spans' append time. */
    day: text('day').notNull(),
    traceCount: integer('traceCount').notNull().default(0),
    /** Traces trace recorded failed (lw.trace.trace_failed). */
    errorCount: integer('errorCount').notNull().default(0),
    /** Append time of the newest root span. */
    lastTraceAt: timestamp('lastTraceAt', { precision: 3 }),
    /** Append time of the newest trace_failed. */
    lastFailedTraceAt: timestamp('lastFailedTraceAt', { precision: 3 }),
  },
  (table) => [
    primaryKey({ name: 'NurturingProjectTraceDay_pkey', columns: [table.projectId, table.day] }),
  ],
)

/**
 * Nurturing's copy of each personal project's owner, from organization's and project's facts.
 * The first owner recorded stays: a personal project's owner never changes.
 */
export const nurturingPersonalProject = pgTable('NurturingPersonalProject', {
  projectId: text('projectId').primaryKey(),
  organizationId: text('organizationId').notNull(),
  ownerUserId: text('ownerUserId').notNull(),
})

/** Nurturing's copy of each project that received a real trace, from trace's and project's facts. */
export const nurturingTracedProject = pgTable(
  'NurturingTracedProject',
  { projectId: text('projectId').primaryKey(), organizationId: text('organizationId').notNull() },
  (table) => [index('NurturingTracedProject_organizationId_idx').on(table.organizationId)],
)

export const billingMeterCheckpoint = pgTable(
  'BillingMeterCheckpoint',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    organizationId: text('organizationId').notNull(),
    billingMonth: text('billingMonth').notNull(),
    meter: text('meter').notNull().default('langwatch_billable_events'),
    lastReportedTotal: integer('lastReportedTotal').notNull().default(0),
    pendingReportedTotal: integer('pendingReportedTotal'),
    consecutiveFailures: integer('consecutiveFailures').notNull().default(0),
    lastCountedEventId: text('lastCountedEventId'),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('BillingMeterCheckpoint_organizationId_billingMonth_meter_key').on(
      table.organizationId,
      table.billingMonth,
      table.meter,
    ),
  ],
)

export const virtualKey = pgTable(
  'VirtualKey',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    name: text('name').notNull(),
    description: text('description'),
    status: virtualKeyStatusEnum('status').notNull().default('ACTIVE'),
    purpose: virtualKeyPurposeEnum('purpose').notNull().default('USER'),
    externalId: text('externalId'),
    metadata: jsonb('metadata').notNull().default({}),
    disabledAt: timestamp('disabledAt', { precision: 3 }),
    disabledReason: text('disabledReason'),
    expiresAt: timestamp('expiresAt', { precision: 3 }),
    hashedSecret: text('hashedSecret').notNull(),
    displayPrefix: text('displayPrefix').notNull(),
    principalUserId: text('principalUserId'),
    traceProjectId: text('traceProjectId'),
    config: jsonb('config').notNull().default({}),
    revision: bigint('revision', { mode: 'bigint' })
      .notNull()
      .default(sql`0`),
    previousHashedSecret: text('previousHashedSecret'),
    previousSecretValidUntil: timestamp('previousSecretValidUntil', { precision: 3 }),
    revokedAt: timestamp('revokedAt', { precision: 3 }),
    revokedById: text('revokedById'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    createdById: text('createdById').notNull(),
    lastUsedAt: timestamp('lastUsedAt', { precision: 3 }),
    routingPolicyId: text('routingPolicyId'),
    routingMode: virtualKeyRoutingModeEnum('routingMode').notNull().default('NONE'),
    connectServices: text('connectServices').array().default([]),
    licenseTokenHash: text('licenseTokenHash'),
    licenseInstanceId: text('licenseInstanceId'),
    licenseExpiresAt: timestamp('licenseExpiresAt', { precision: 3 }),
    licenseId: text('licenseId'),
  },
  (table) => [
    uniqueIndex('VirtualKey_licenseTokenHash_key').on(table.licenseTokenHash),
    index('VirtualKey_organizationId_idx').on(table.organizationId),
    index('VirtualKey_displayPrefix_idx').on(table.displayPrefix),
    index('VirtualKey_status_idx').on(table.status),
    index('VirtualKey_principalUserId_idx').on(table.principalUserId),
    index('VirtualKey_revision_idx').on(table.revision),
    index('VirtualKey_routingPolicyId_idx').on(table.routingPolicyId),
    index('VirtualKey_organizationId_purpose_idx').on(table.organizationId, table.purpose),
    uniqueIndex('VirtualKey_hashedSecret_key').on(table.hashedSecret),
    uniqueIndex('VirtualKey_organizationId_externalId_key').on(
      table.organizationId,
      table.externalId,
    ),
    foreignKey({
      name: 'VirtualKey_principalUserId_fkey',
      columns: [table.principalUserId],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'VirtualKey_routingPolicyId_fkey',
      columns: [table.routingPolicyId],
      foreignColumns: [routingPolicy.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const virtualKeyScope = pgTable(
  'VirtualKeyScope',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    virtualKeyId: text('virtualKeyId').notNull(),
    scopeType: virtualKeyScopeTypeEnum('scopeType').notNull(),
    scopeId: text('scopeId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('VirtualKeyScope_scopeType_scopeId_idx').on(table.scopeType, table.scopeId),
    index('VirtualKeyScope_virtualKeyId_idx').on(table.virtualKeyId),
    uniqueIndex('VirtualKeyScope_virtualKeyId_scopeType_scopeId_key').on(
      table.virtualKeyId,
      table.scopeType,
      table.scopeId,
    ),
    foreignKey({
      name: 'VirtualKeyScope_virtualKeyId_fkey',
      columns: [table.virtualKeyId],
      foreignColumns: [virtualKey.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const routingPolicy = pgTable(
  'RoutingPolicy',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    name: text('name').notNull(),
    description: text('description'),
    modelProviderIds: jsonb('modelProviderIds').notNull().default([]),
    modelAliases: jsonb('modelAliases').notNull().default({}),
    defaultModel: text('defaultModel'),
    policyRules: jsonb('policyRules').notNull().default({}),
    isDefault: boolean('isDefault').notNull().default(false),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    createdById: text('createdById'),
    updatedById: text('updatedById'),
  },
  (table) => [
    index('RoutingPolicy_organizationId_idx').on(table.organizationId),
    foreignKey({
      name: 'RoutingPolicy_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'RoutingPolicy_createdById_fkey',
      columns: [table.createdById],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'RoutingPolicy_updatedById_fkey',
      columns: [table.updatedById],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const routingPolicyScope = pgTable(
  'RoutingPolicyScope',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    routingPolicyId: text('routingPolicyId').notNull(),
    scopeType: routingPolicyScopeTypeEnum('scopeType').notNull(),
    scopeId: text('scopeId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('RoutingPolicyScope_scopeType_scopeId_idx').on(table.scopeType, table.scopeId),
    index('RoutingPolicyScope_routingPolicyId_idx').on(table.routingPolicyId),
    uniqueIndex('RoutingPolicyScope_routingPolicyId_scopeType_scopeId_key').on(
      table.routingPolicyId,
      table.scopeType,
      table.scopeId,
    ),
    foreignKey({
      name: 'RoutingPolicyScope_routingPolicyId_fkey',
      columns: [table.routingPolicyId],
      foreignColumns: [routingPolicy.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const gatewayGuardrail = pgTable(
  'GatewayGuardrail',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    projectId: text('projectId').notNull(),
    name: text('name').notNull(),
    description: text('description'),
    evaluatorId: text('evaluatorId').notNull(),
    direction: gatewayGuardrailDirectionEnum('direction').notNull(),
    failureMode: gatewayGuardrailFailureModeEnum('failureMode').notNull().default('FAIL_CLOSED'),
    archivedAt: timestamp('archivedAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    createdById: text('createdById'),
    updatedById: text('updatedById'),
  },
  (table) => [
    index('GatewayGuardrail_projectId_idx').on(table.projectId),
    index('GatewayGuardrail_evaluatorId_idx').on(table.evaluatorId),
    index('GatewayGuardrail_projectId_archivedAt_idx').on(table.projectId, table.archivedAt),
    foreignKey({
      name: 'GatewayGuardrail_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'GatewayGuardrail_evaluatorId_fkey',
      columns: [table.evaluatorId],
      foreignColumns: [evaluator.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'GatewayGuardrail_createdById_fkey',
      columns: [table.createdById],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'GatewayGuardrail_updatedById_fkey',
      columns: [table.updatedById],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

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
export const ingestionSource = pgTable(
  'IngestionSource',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    /** Optional team-scope. null = org-wide. */
    teamId: text('teamId'),
    /**
     * 'otel_generic' | 'claude_cowork' | 'workato' | 'copilot_studio' |
     * 'openai_compliance' | 'claude_compliance' | 's3_custom' (extensible).
     * String not enum so adapters can be added without migrations - the
     * supported set lives in code (activity-monitor adapter registry).
     */
    sourceType: text('sourceType').notNull(),
    /**
     * Human-friendly label for the admin UI. Unique per-org so two
     * "Cowork" sources can coexist with different names.
     */
    name: text('name').notNull(),
    description: text('description'),
    /**
     * HMAC bearer token the upstream operator pastes into the platform
     * (or LangWatch sends in pull mode). Stored encrypted at rest via
     * the same pepper-derived KEK as VirtualKey secrets. Auto-generated
     * on create, rotatable via the admin UI with a 24h grace window.
     */
    ingestSecretHash: text('ingestSecretHash').notNull(),
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
    parserConfig: jsonb('parserConfig').notNull().default({}),
    /**
     * For pull-based sources, persisted between polls so the next poll
     * resumes where the last left off (timestamp / cursor / file marker).
     */
    pollerCursor: jsonb('pollerCursor'),
    /**
     * Phase 10 - number of consecutive errors observed by the puller
     * worker. Reset to 0 on a successful run. Drives admin-UI banners
     * ("Microsoft authentication failed - re-authorize") + back-off
     * scheduling when the source has been broken for a while.
     * Spec: specs/ai-governance/puller-framework/puller-adapter-contract.feature
     */
    errorCount: integer('errorCount').notNull().default(0),
    /**
     * Phase 10 - cron schedule for pull-mode sources. Validated by the
     * adapter's pullConfig schema; mirrored here for fast worker
     * dispatch. Null for push-mode sources.
     */
    pullSchedule: text('pullSchedule'),
    /**
     * Status flag. 'active' (default), 'disabled' (admin paused), or
     * 'awaiting_first_event' (created but never received).
     */
    status: text('status').notNull().default('awaiting_first_event'),
    /**
     * ADR-088 v7: optional trace DESTINATION for conversation-bearing pulls
     * (Databricks Genie). Null = don't route. Same shape as
     * VirtualKey.traceProjectId: a destination is NOT a scope row and grants
     * no access to the source; it only decides where routed conversation
     * traces land. Must be a project of this org (validated at write time).
     * Freely editable; already-routed history stays where it landed.
     * NOT the deferred cost-attribution projectId (ADR-088 Decision 4).
     */
    traceProjectId: text('traceProjectId'),
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
    providerAccountId: text('providerAccountId'),
    /**
     * Last event timestamp observed (any direction). Surfaces in the
     * admin UI as "last event 2 minutes ago".
     */
    lastEventAt: timestamp('lastEventAt', { precision: 3 }),
    /**
     * ADR-128 - when a pull run last reached the provider without erroring,
     * mirrored from the ingestionPullRunStatus projection. Deliberately NOT
     * lastEventAt: an empty period is a working puller, and a run that
     * delivered data and then failed is not. Drives the health badge and
     * the "no data since" cut-off, below which a day's spend is unknown
     * rather than zero.
     * Spec: specs/governance/ingestion-source-health.feature
     */
    lastSuccessAt: timestamp('lastSuccessAt', { precision: 3 }),
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
    lastReadThroughAt: timestamp('lastReadThroughAt', { precision: 3 }),
    lastRunCompleteness: text('lastRunCompleteness'),
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
    unpricedUsageSince: timestamp('unpricedUsageSince', { precision: 3 }),
    unpricedUsageThrough: timestamp('unpricedUsageThrough', { precision: 3 }),
    archivedAt: timestamp('archivedAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    createdById: text('createdById'),
  },
  (table) => [
    index('IngestionSource_organizationId_idx').on(table.organizationId),
    index('IngestionSource_teamId_idx').on(table.teamId),
    index('IngestionSource_sourceType_idx').on(table.sourceType),
    uniqueIndex('IngestionSource_organizationId_name_key').on(table.organizationId, table.name),
    foreignKey({
      name: 'IngestionSource_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'IngestionSource_teamId_fkey',
      columns: [table.teamId],
      foreignColumns: [team.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'IngestionSource_createdById_fkey',
      columns: [table.createdById],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

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
export const anomalyRule = pgTable(
  'AnomalyRule',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    /**
     * 'organization' | 'team' | 'project' | 'source_type' | 'source'.
     * String not enum so new scopes ship without migrations.
     */
    scope: text('scope').notNull(),
    /**
     * Org id / team id / project id / source-type string / source id
     * - narrows what the rule applies to.
     */
    scopeId: text('scopeId').notNull(),
    name: text('name').notNull(),
    description: text('description'),
    /**
     * 'critical' | 'warning' | 'info'. Drives default destination
     * routing (e.g. critical → PagerDuty, warning → Slack only).
     */
    severity: text('severity').notNull(),
    /**
     * Extensible rule taxonomy: 'spend_spike' | 'rate_limit' |
     * 'tool_mismatch' | 'after_hours' | 'unusual_model' | 'pii_leak'
     * | 'custom'. New types ship without migrations - evaluator
     * picks a per-type implementation from a code registry.
     */
    ruleType: text('ruleType').notNull(),
    /**
     * Per-rule-type threshold shape. Examples:
     * spend_spike: { windowSec, ratioVsBaseline, minBaselineUsd }
     * rate_limit: { windowSec, requestsThreshold }
     * after_hours: { startHour, endHour, timezone, requestsThreshold }
     */
    thresholdConfig: jsonb('thresholdConfig').notNull().default({}),
    /**
     * Alert routing per destination. Shape:
     * { slack?: { channelId }, siem?: { endpointUrl, authHeader },
     * webhook?: { url, authHeader }, pagerduty?: { routingKey },
     * email?: { to: string[] } }
     * Empty map = use org defaults (when defaults exist; today no-op).
     */
    destinationConfig: jsonb('destinationConfig').notNull().default({}),
    /**
     * 'active' (default - evaluator picks it up) | 'disabled' (admin
     * paused; rule rows kept for audit). 'archived' uses archivedAt
     * timestamp instead so we can filter `archivedAt IS NULL` cheaply.
     */
    status: text('status').notNull().default('active'),
    archivedAt: timestamp('archivedAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    createdById: text('createdById'),
  },
  (table) => [
    index('AnomalyRule_organizationId_idx').on(table.organizationId),
    index('AnomalyRule_scope_scopeId_idx').on(table.scope, table.scopeId),
    index('AnomalyRule_ruleType_idx').on(table.ruleType),
    index('AnomalyRule_status_idx').on(table.status),
    uniqueIndex('AnomalyRule_organizationId_name_key').on(table.organizationId, table.name),
    foreignKey({
      name: 'AnomalyRule_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AnomalyRule_createdById_fkey',
      columns: [table.createdById],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

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
export const anomalyAlert = pgTable(
  'AnomalyAlert',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    ruleId: text('ruleId').notNull(),
    /**
     * Denormalised from rule for fast WHERE filtering on the
     * dashboard query (no join needed for the common path).
     */
    severity: text('severity').notNull(),
    ruleName: text('ruleName').notNull(),
    ruleType: text('ruleType').notNull(),
    /** Window the trigger refers to (UTC). */
    triggerWindowStart: timestamp('triggerWindowStart', { precision: 3 }).notNull(),
    triggerWindowEnd: timestamp('triggerWindowEnd', { precision: 3 }).notNull(),
    /** Numeric trigger context. Per-rule-type detail in `detail` JSONB. */
    triggerSpendUsd: numeric('triggerSpendUsd', { precision: 18, scale: 6 }),
    triggerEventCount: integer('triggerEventCount'),
    /**
     * Per-rule-type extra context: {topActors: [...], baseline: ...,
     * observed: ..., timezone: ..., etc}.
     */
    detail: jsonb('detail').notNull().default({}),
    /** 'open' | 'acknowledged' | 'resolved'. */
    state: text('state').notNull().default('open'),
    /**
     * Per-destination dispatch outcome:
     * { webhook: { lastAttemptIso, lastError?, attemptCount, statusCode? },
     * slack: ..., pagerduty: ..., siem: ..., email: ... }
     */
    destinationStatus: jsonb('destinationStatus').notNull().default({}),
    detectedAt: timestamp('detectedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    acknowledgedAt: timestamp('acknowledgedAt', { precision: 3 }),
    resolvedAt: timestamp('resolvedAt', { precision: 3 }),
  },
  (table) => [
    index('AnomalyAlert_organizationId_detectedAt_idx').on(table.organizationId, table.detectedAt),
    index('AnomalyAlert_ruleId_idx').on(table.ruleId),
    index('AnomalyAlert_state_idx').on(table.state),
    uniqueIndex('AnomalyAlert_ruleId_triggerWindowStart_key').on(
      table.ruleId,
      table.triggerWindowStart,
    ),
    foreignKey({
      name: 'AnomalyAlert_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AnomalyAlert_ruleId_fkey',
      columns: [table.ruleId],
      foreignColumns: [anomalyRule.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const aiToolEntry = pgTable(
  'AiToolEntry',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    /**
     * "organization" | "team". String not enum so new scopes ship
     * without migrations.
     */
    scope: text('scope').notNull(),
    /**
     * organizationId or teamId, narrowing what scope this entry applies
     * to. Always equals organizationId when scope=organization.
     */
    scopeId: text('scopeId').notNull(),
    /**
     * "coding_assistant" | "model_provider" | "external_tool".
     * Discriminator drives the per-tile UX shape and validates
     * the `config` JSON column at the app layer.
     */
    type: text('type').notNull(),
    displayName: text('displayName').notNull(),
    /**
     * Stable slug for team-overrides-org resolution (a team-scoped
     * entry with the same slug as an org-scoped one hides the org
     * default for users in that team). Server-owned: generated from
     * displayName + nanoid(6) on create, immutable on update.
     */
    slug: text('slug').notNull(),
    /**
     * @deprecated Use `iconAsset` instead. Kept for back-compat with
     * pre-refactor rows; new writes should populate iconAsset.
     */
    iconKey: text('iconKey'),
    /**
     * Icon source - single column, prefix-discriminated:
     * "preset:claude_code" / "preset:codex" / ...   (built-in)
     * "data:image/svg+xml;base64,..."               (admin-uploaded)
     * null                                          → fall back to
     * type-default in UI
     */
    iconAsset: text('iconAsset'),
    /** Admin-curated sort within the user-facing portal grid. */
    order: integer('order').notNull().default(0),
    /**
     * When false the entry is hidden from the user-facing list but
     * remains in the admin list (so admins can re-enable without
     * losing per-entry config). Disabled is "soft-hide", archived is
     * "soft-delete" (preserves audit trail + scope binding history).
     */
    enabled: boolean('enabled').notNull().default(true),
    /**
     * Per-type config JSON. Validated via zod discriminated union
     * at the service layer (see AiToolEntryService).
     */
    config: jsonb('config').notNull().default({}),
    archivedAt: timestamp('archivedAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    createdById: text('createdById'),
    updatedById: text('updatedById'),
  },
  (table) => [
    index('AiToolEntry_organizationId_scope_scopeId_idx').on(
      table.organizationId,
      table.scope,
      table.scopeId,
    ),
    index('AiToolEntry_organizationId_enabled_archivedAt_idx').on(
      table.organizationId,
      table.enabled,
      table.archivedAt,
    ),
    index('AiToolEntry_organizationId_type_idx').on(table.organizationId, table.type),
    foreignKey({
      name: 'AiToolEntry_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * @deprecated Per-team scope binding for AiToolEntry. Superseded by
 * AiToolEntryDepartment; kept so the rename migration is non-destructive.
 * No new rows are written here.
 */
export const aiToolEntryTeam = pgTable(
  'AiToolEntryTeam',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    entryId: text('entryId').notNull(),
    teamId: text('teamId').notNull(),
  },
  (table) => [
    index('AiToolEntryTeam_teamId_idx').on(table.teamId),
    uniqueIndex('AiToolEntryTeam_entryId_teamId_key').on(table.entryId, table.teamId),
    foreignKey({
      name: 'AiToolEntryTeam_entryId_fkey',
      columns: [table.entryId],
      foreignColumns: [aiToolEntry.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AiToolEntryTeam_teamId_fkey',
      columns: [table.teamId],
      foreignColumns: [team.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * Per-department scope binding for AiToolEntry. Empty set = org-wide.
 * Non-empty = entry visible only to members of those departments (an
 * OrganizationUser whose `departmentId` is in the set). Mirrors
 * AiToolEntryTeam's shape. `departmentId` is a plain column ref
 * (relationMode="prisma"), matching how Department is referenced
 * elsewhere; the service validates the id belongs to the org.
 */
export const aiToolEntryDepartment = pgTable(
  'AiToolEntryDepartment',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    entryId: text('entryId').notNull(),
    departmentId: text('departmentId').notNull(),
  },
  (table) => [
    index('AiToolEntryDepartment_departmentId_idx').on(table.departmentId),
    uniqueIndex('AiToolEntryDepartment_entryId_departmentId_key').on(
      table.entryId,
      table.departmentId,
    ),
    foreignKey({
      name: 'AiToolEntryDepartment_entryId_fkey',
      columns: [table.entryId],
      foreignColumns: [aiToolEntry.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const ingestionTemplate = pgTable(
  'IngestionTemplate',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId'),
    /**
     * Stable slug - unique within (organizationId, slug) scope. Generated
     * from displayName + nanoid(6) on create for org-authored; canonical
     * values for platform-published (`claude_code` / `cursor` /
     * `claude_cowork` / `raw_otlp_advanced`).
     */
    slug: text('slug').notNull(),
    /**
     * Mirrors the slug for v1 platform-published rows. Kept as a separate
     * column so a future v2 may publish two templates that share an
     * upstream sourceType (e.g. claude_code-personal vs claude_code-team)
     * without a migration. String not enum - extensibility.
     */
    sourceType: text('sourceType').notNull(),
    displayName: text('displayName').notNull(),
    description: text('description'),
    /**
     * Icon source - single column, prefix-discriminated, mirrors
     * AiToolEntry.iconAsset:
     * `preset:<kind>`                 (built-in slot)
     * `data:image/svg+xml;base64,…`   (admin-uploaded for org-authored)
     * null                            → fall back to slug-default in UI
     */
    iconAsset: text('iconAsset'),
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
    credentialSchema: text('credentialSchema'),
    /**
     * Admin/platform-authored OTTL statements. Applied at receive time
     * AFTER the receiver re-stamps the binding's authoritative principal
     * + provenance keys (see `protectedTemplateAttributeKeys` 19-key
     * closed list). Templates write canonical OTTL outputs (gen_ai.*,
     * cost.usd, response.model) - they CANNOT write attribution or
     * provenance keys; the receiver guard restores those post-OTTL.
     */
    ottlRules: text('ottlRules').notNull().default(''),
    /**
     * True for the v1 platform-published default templates
     * (organizationId IS NULL). Future org-authored rows write false.
     * Drives the admin UI "READ-ONLY view + Request a custom template"
     * surface - admins consume + opt-out only in v1.
     */
    platformPublished: boolean('platformPublished').notNull().default(false),
    /**
     * Admin/platform enable. False = hidden from /me catalog. Existing
     * ingestion keys continue to land traces; new installs are blocked.
     */
    enabled: boolean('enabled').notNull().default(true),
    archivedAt: timestamp('archivedAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    createdById: text('createdById'),
    updatedById: text('updatedById'),
  },
  (table) => [
    index('IngestionTemplate_organizationId_archivedAt_idx').on(
      table.organizationId,
      table.archivedAt,
    ),
    index('IngestionTemplate_sourceType_idx').on(table.sourceType),
    uniqueIndex('IngestionTemplate_organizationId_slug_key').on(table.organizationId, table.slug),
    foreignKey({
      name: 'IngestionTemplate_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const gatewayA2ATarget = pgTable(
  'GatewayA2ATarget',
  {
    id: text('id').primaryKey(),
    organizationId: text('organizationId').notNull(),
    projectId: text('projectId'),
    name: text('name').notNull(),
    agentName: text('agentName').notNull(),
    cardUrl: text('cardUrl').notNull(),
    upstreamUrl: text('upstreamUrl').notNull(),
    encryptedUpstreamHeaders: text('encryptedUpstreamHeaders'),
    restrictToAllowList: boolean('restrictToAllowList').notNull().default(false),
    allowedVirtualKeyIds: text('allowedVirtualKeyIds').array().default([]),
    blocked: boolean('blocked').notNull().default(false),
    blockReason: text('blockReason'),
    createdById: text('createdById').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    archivedAt: timestamp('archivedAt', { precision: 3 }),
  },
  (table) => [
    index('GatewayA2ATarget_organizationId_projectId_idx').on(
      table.organizationId,
      table.projectId,
    ),
  ],
)

export const gatewayConnectUpstream = pgTable('GatewayConnectUpstream', {
  organizationId: text('organizationId').primaryKey(),
  baseUrl: text('baseUrl').notNull(),
  encryptedToken: text('encryptedToken').notNull(),
  instanceId: text('instanceId').notNull(),
  updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
})

export const gatewayTraceExportKey = pgTable('GatewayTraceExportKey', {
  projectId: text('projectId').primaryKey(),
  apiKeyId: text('apiKeyId').notNull(),
  encryptedToken: text('encryptedToken').notNull(),
  createdAt: timestamp('createdAt', { precision: 3 })
    .notNull()
    .default(sql`CURRENT_TIMESTAMP`)
    .$defaultFn(utcNow),
})

export const gatewayBudget = pgTable(
  'GatewayBudget',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    scopeType: gatewayBudgetScopeTypeEnum('scopeType').notNull(),
    scopeId: text('scopeId').notNull(),
    providerKey: text('providerKey'),
    name: text('name').notNull(),
    description: text('description'),
    window: gatewayBudgetWindowEnum('window').notNull(),
    limitUsd: numeric('limitUsd', { precision: 18, scale: 6 }).notNull(),
    onBreach: gatewayBudgetBreachActionEnum('onBreach').notNull().default('BLOCK'),
    timezone: text('timezone'),
    externalId: text('externalId'),
    metadata: jsonb('metadata').notNull().default({}),
    spentUsd: numeric('spentUsd', { precision: 18, scale: 6 }).notNull().default('0'),
    currentPeriodStartedAt: timestamp('currentPeriodStartedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    resetsAt: timestamp('resetsAt', { precision: 3 }).notNull(),
    lastResetAt: timestamp('lastResetAt', { precision: 3 }),
    cycleAnchorAt: timestamp('cycleAnchorAt', { precision: 3 }),
    archivedAt: timestamp('archivedAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    createdById: text('createdById').notNull(),
    managedByVirtualKeyId: text('managedByVirtualKeyId'),
  },
  (table) => [
    index('GatewayBudget_organizationId_idx').on(table.organizationId),
    index('GatewayBudget_scopeType_scopeId_providerKey_idx').on(
      table.scopeType,
      table.scopeId,
      table.providerKey,
    ),
    index('GatewayBudget_managedByVirtualKeyId_idx').on(table.managedByVirtualKeyId),
    index('GatewayBudget_resetsAt_idx').on(table.resetsAt),
    uniqueIndex('GatewayBudget_organizationId_externalId_key').on(
      table.organizationId,
      table.externalId,
    ),
    foreignKey({
      name: 'GatewayBudget_createdById_fkey',
      columns: [table.createdById],
      foreignColumns: [user.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

export const gatewayBudgetBucketBoundary = pgTable(
  'GatewayBudgetBucketBoundary',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    budgetId: text('budgetId').notNull(),
    bucketScopeId: text('bucketScopeId').notNull(),
    periodStartedAt: timestamp('periodStartedAt', { precision: 3 }).notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('GatewayBudgetBucketBoundary_organizationId_idx').on(table.organizationId),
    uniqueIndex('GatewayBudgetBucketBoundary_budgetId_bucketScopeId_key').on(
      table.budgetId,
      table.bucketScopeId,
    ),
    foreignKey({
      name: 'GatewayBudgetBucketBoundary_budgetId_fkey',
      columns: [table.budgetId],
      foreignColumns: [gatewayBudget.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const gatewayBudgetLedger = pgTable(
  'GatewayBudgetLedger',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    budgetId: text('budgetId').notNull(),
    virtualKeyId: text('virtualKeyId').notNull(),
    modelProviderId: text('modelProviderId'),
    providerKey: text('providerKey'),
    gatewayRequestId: text('gatewayRequestId').notNull(),
    amountUsd: numeric('amountUsd', { precision: 18, scale: 6 }).notNull(),
    tokensInput: integer('tokensInput').notNull().default(0),
    tokensOutput: integer('tokensOutput').notNull().default(0),
    tokensCacheRead: integer('tokensCacheRead').notNull().default(0),
    tokensCacheWrite: integer('tokensCacheWrite').notNull().default(0),
    model: text('model').notNull(),
    durationMs: integer('durationMs'),
    status: gatewayBudgetLedgerStatusEnum('status').notNull().default('SUCCESS'),
    occurredAt: timestamp('occurredAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('GatewayBudgetLedger_virtualKeyId_idx').on(table.virtualKeyId),
    index('GatewayBudgetLedger_gatewayRequestId_idx').on(table.gatewayRequestId),
    index('GatewayBudgetLedger_occurredAt_idx').on(table.occurredAt),
    uniqueIndex('GatewayBudgetLedger_budgetId_gatewayRequestId_key').on(
      table.budgetId,
      table.gatewayRequestId,
    ),
    foreignKey({
      name: 'GatewayBudgetLedger_budgetId_fkey',
      columns: [table.budgetId],
      foreignColumns: [gatewayBudget.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'GatewayBudgetLedger_virtualKeyId_fkey',
      columns: [table.virtualKeyId],
      foreignColumns: [virtualKey.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const gatewayChangeEvent = pgTable(
  'GatewayChangeEvent',
  {
    revision: bigserial('revision', { mode: 'bigint' }).primaryKey(),
    organizationId: text('organizationId').notNull(),
    kind: gatewayChangeEventKindEnum('kind').notNull(),
    virtualKeyId: text('virtualKeyId'),
    budgetId: text('budgetId'),
    modelProviderId: text('modelProviderId'),
    projectId: text('projectId'),
    payload: jsonb('payload'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('GatewayChangeEvent_organizationId_revision_idx').on(
      table.organizationId,
      table.revision,
    ),
    index('GatewayChangeEvent_projectId_revision_idx').on(table.projectId, table.revision),
    index('GatewayChangeEvent_virtualKeyId_idx').on(table.virtualKeyId),
    index('GatewayChangeEvent_budgetId_idx').on(table.budgetId),
    foreignKey({
      name: 'GatewayChangeEvent_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'GatewayChangeEvent_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const gatewayCacheRule = pgTable(
  'GatewayCacheRule',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    name: text('name').notNull(),
    description: text('description'),
    priority: integer('priority').notNull().default(100),
    enabled: boolean('enabled').notNull().default(true),
    matchers: jsonb('matchers').notNull(),
    action: jsonb('action').notNull(),
    modeEnum: gatewayCacheRuleModeEnum('modeEnum').notNull(),
    archivedAt: timestamp('archivedAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    createdById: text('createdById').notNull(),
  },
  (table) => [
    index('GatewayCacheRule_organizationId_archivedAt_priority_idx').on(
      table.organizationId,
      table.archivedAt,
      table.priority,
    ),
    index('GatewayCacheRule_organizationId_enabled_priority_idx').on(
      table.organizationId,
      table.enabled,
      table.priority,
    ),
  ],
)

export const featureFlag = pgTable('FeatureFlag', {
  key: text('key').primaryKey(),
  enabled: boolean('enabled').notNull(),
  rules: jsonb('rules'),
  lastEditedBy: text('lastEditedBy'),
  createdAt: timestamp('createdAt', { precision: 3 })
    .notNull()
    .default(sql`CURRENT_TIMESTAMP`)
    .$defaultFn(utcNow),
  updatedAt: timestamp('updatedAt', { precision: 3 })
    .notNull()
    .default(sql`CURRENT_TIMESTAMP`)
    .$defaultFn(utcNow)
    .$onUpdate(utcNow),
})

/**
 * One experiment setting per (flag, subject). A USER row is that person's
 * own opt-in; an ORGANIZATION or PROJECT row is an owner's explicit policy
 * for that exact scope. No row means inherit.
 */
export const featureFlagExperimentSetting = pgTable(
  'FeatureFlagExperimentSetting',
  {
    flagKey: text('flagKey').notNull(),
    subjectType: featureFlagExperimentSubjectTypeEnum('subjectType').notNull(),
    subjectId: text('subjectId').notNull(),
    enabled: boolean('enabled').notNull(),
    changedByUserId: text('changedByUserId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    primaryKey({
      name: 'FeatureFlagExperimentSetting_pkey',
      columns: [table.flagKey, table.subjectType, table.subjectId],
    }),
    index('FeatureFlagExperimentSetting_subjectType_subjectId_idx').on(
      table.subjectType,
      table.subjectId,
    ),
  ],
)

/** Durable state for a long-running event-sourced process. */
export const processManagerInstance = pgTable(
  'ProcessManagerInstance',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    processName: text('processName').notNull(),
    projectId: text('projectId').notNull(),
    processKey: text('processKey').notNull(),
    tenantId: text('tenantId').notNull(),
    userId: text('userId'),
    state: jsonb('state').notNull(),
    revision: integer('revision').notNull(),
    nextWakeAt: timestamp('nextWakeAt', { withTimezone: true, precision: 3 }),
    updatedAt: timestamp('updatedAt', { withTimezone: true, precision: 3 }).notNull(),
  },
  (table) => [
    index('ProcessManagerInstance_nextWakeAt_idx').on(table.nextWakeAt),
    uniqueIndex('ProcessManagerInstance_processName_projectId_processKey_key').on(
      table.processName,
      table.projectId,
      table.processKey,
    ),
  ],
)

/** Idempotency marker for a source event consumed by a process definition. */
export const processManagerInbox = pgTable(
  'ProcessManagerInbox',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    processName: text('processName').notNull(),
    projectId: text('projectId').notNull(),
    processKey: text('processKey').notNull(),
    tenantId: text('tenantId').notNull(),
    /**
     * The raw source event id, kept UNINDEXED for diagnostics. A caller composes
     * this and its length is that domain's business, so it must never carry a
     * btree constraint — see `sourceEventKey`.
     */
    sourceEventId: text('sourceEventId').notNull(),
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
    sourceEventKey: text('sourceEventKey'),
    consumedAt: timestamp('consumedAt', { withTimezone: true, precision: 3 }).notNull(),
  },
  (table) => [
    index('ProcessManagerInbox_processName_projectId_processKey_consum_idx').on(
      table.processName,
      table.projectId,
      table.processKey,
      table.consumedAt,
    ),
    index('ProcessManagerInbox_consumedAt_idx').on(table.consumedAt),
    uniqueIndex('ProcessManagerInbox_processName_projectId_sourceEventKey_key').on(
      table.processName,
      table.projectId,
      table.sourceEventKey,
    ),
  ],
)

/** Transactional intents emitted by a durable process transition. */
export const processManagerOutbox = pgTable(
  'ProcessManagerOutbox',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    processName: text('processName').notNull(),
    projectId: text('projectId').notNull(),
    processKey: text('processKey').notNull(),
    tenantId: text('tenantId').notNull(),
    userId: text('userId'),
    messageKey: text('messageKey').notNull(),
    intentType: text('intentType').notNull(),
    payload: jsonb('payload').notNull(),
    traceCarrier: jsonb('traceCarrier').notNull(),
    sourceEventId: text('sourceEventId'),
    status: processManagerOutboxStatusEnum('status').notNull().default('pending'),
    attempts: integer('attempts').notNull().default(0),
    nextAttemptAt: timestamp('nextAttemptAt', { withTimezone: true, precision: 3 }).notNull(),
    leasedUntil: timestamp('leasedUntil', { withTimezone: true, precision: 3 }),
    leaseToken: text('leaseToken'),
    dispatchedAt: timestamp('dispatchedAt', { withTimezone: true, precision: 3 }),
    createdAt: timestamp('createdAt', { withTimezone: true, precision: 3 }).notNull(),
    updatedAt: timestamp('updatedAt', { withTimezone: true, precision: 3 }).notNull(),
  },
  (table) => [
    uniqueIndex('ProcessManagerOutbox_leaseToken_key').on(table.leaseToken),
    index('ProcessManagerOutbox_processName_projectId_processKey_creat_idx').on(
      table.processName,
      table.projectId,
      table.processKey,
      table.createdAt,
    ),
    index('ProcessManagerOutbox_status_nextAttemptAt_leasedUntil_idx').on(
      table.status,
      table.nextAttemptAt,
      table.leasedUntil,
    ),
    index('ProcessManagerOutbox_status_processName_nextAttemptAt_idx').on(
      table.status,
      table.processName,
      table.nextAttemptAt,
    ),
    uniqueIndex('ProcessManagerOutbox_processName_projectId_messageKey_key').on(
      table.processName,
      table.projectId,
      table.messageKey,
    ),
  ],
)

/**
 * Why an outbox delivery failed, one row per FAILED attempt
 * (specs/ops/dead-letter-recovery.feature). Successes write nothing, so the
 * table grows with trouble, not with traffic. Rows die with their message:
 * the cascade delete is the retention story, riding the existing sweep.
 */
export const processManagerOutboxAttempt = pgTable(
  'ProcessManagerOutboxAttempt',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    outboxId: text('outboxId').notNull(),
    projectId: text('projectId').notNull(),
    /** 1-based attempt number, as the dispatcher counts it. */
    attempt: integer('attempt').notNull(),
    occurredAt: timestamp('occurredAt', { withTimezone: true, precision: 3 }).notNull(),
    outcome: processManagerOutboxAttemptOutcomeEnum('outcome').notNull(),
    errorType: text('errorType').notNull(),
    /**
     * The safe failure diagnostic the dispatcher already computes for its
     * span — never a raw provider body.
     */
    errorMessage: text('errorMessage').notNull(),
    retryAfterMs: integer('retryAfterMs'),
  },
  (table) => [
    index('ProcessManagerOutboxAttempt_outboxId_attempt_idx').on(table.outboxId, table.attempt),
    foreignKey({
      name: 'ProcessManagerOutboxAttempt_outboxId_fkey',
      columns: [table.outboxId],
      foreignColumns: [processManagerOutbox.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * Per-project topic clustering run status (ADR-051): the rebuildable
 * operational read model behind the settings page. One row per project;
 * folded from run_completed/run_failed events, never written directly.
 */
export const topicClusteringRunProjection = pgTable(
  'TopicClusteringRunProjection',
  {
    /** KSUID (`topicrun_…`), generated by the repository on first upsert. */
    id: text('id').primaryKey(),
    projectId: text('projectId').notNull(),
    LastRequestedAt: doublePrecision('lastRequestedAt'),
    LastRequestTrigger: text('lastRequestTrigger'),
    LastRunAt: doublePrecision('lastRunAt'),
    /** completed | skipped | failed */
    LastRunOutcome: text('lastRunOutcome'),
    LastRunMode: text('lastRunMode'),
    LastRunSkippedReason: text('lastRunSkippedReason'),
    LastRunError: text('lastRunError'),
    LastRunErrorCode: text('lastRunErrorCode'),
    /** True when the customer can resolve the failure (credentials, quota). */
    LastRunErrorUserActionable: boolean('lastRunErrorUserActionable').notNull().default(false),
    LastRunTracesProcessed: integer('lastRunTracesProcessed').notNull().default(0),
    LastRunTopicsCount: integer('lastRunTopicsCount').notNull().default(0),
    LastRunSubtopicsCount: integer('lastRunSubtopicsCount').notNull().default(0),
    LastRunPages: integer('lastRunPages').notNull().default(0),
    InProgressRunId: text('inProgressRunId'),
    InProgressTraces: integer('inProgressTraces').notNull().default(0),
    InProgressPages: integer('inProgressPages').notNull().default(0),
    /**
     * Business time the in-progress run opened; bounds how long the read
     * model may report it as running (mirrors the process's stale-run guard).
     */
    InProgressStartedAt: doublePrecision('inProgressStartedAt'),
    CreatedAt: doublePrecision('createdAt').notNull(),
    UpdatedAt: doublePrecision('updatedAt').notNull(),
    /** Business time of the latest applied event. */
    OccurredAt: doublePrecision('occurredAt').notNull(),
    /** Event-log acceptance time. Together with lastEventId this is the cursor. */
    AcceptedAt: doublePrecision('acceptedAt').notNull(),
    LastEventId: text('lastEventId').notNull(),
    ProjectionVersion: text('projectionVersion').notNull(),
  },
  (table) => [uniqueIndex('TopicClusteringRunProjection_projectId_key').on(table.projectId)],
)

/**
 * Cursor row for the topic-model projection (ADR-051). The model itself is
 * the `Topic` table, written through by the projection; this row only
 * records how far the fold has applied the event log.
 */
export const topicModelProjection = pgTable(
  'TopicModelProjection',
  {
    /** KSUID (`topicmodel_…`), generated by the repository on first upsert. */
    id: text('id').primaryKey(),
    projectId: text('projectId').notNull(),
    CreatedAt: doublePrecision('createdAt').notNull(),
    UpdatedAt: doublePrecision('updatedAt').notNull(),
    /** Business time of the latest applied event. */
    OccurredAt: doublePrecision('occurredAt').notNull(),
    /** Event-log acceptance time. Together with lastEventId this is the cursor. */
    AcceptedAt: doublePrecision('acceptedAt').notNull(),
    LastEventId: text('lastEventId').notNull(),
    ProjectionVersion: text('projectionVersion').notNull(),
  },
  (table) => [uniqueIndex('TopicModelProjection_projectId_key').on(table.projectId)],
)

/**
 * Per-project topic clustering run history (audit read model, ADR-051).
 * One row per project; `runs` is the bounded newest-first entry list.
 * Rebuildable by replaying the event log.
 */
export const topicClusteringRunHistoryProjection = pgTable(
  'TopicClusteringRunHistoryProjection',
  {
    /** KSUID (`topicrunhist_…`), generated by the repository on first upsert. */
    id: text('id').primaryKey(),
    projectId: text('projectId').notNull(),
    Runs: jsonb('runs').notNull(),
    CreatedAt: doublePrecision('createdAt').notNull(),
    UpdatedAt: doublePrecision('updatedAt').notNull(),
    /** Business time of the latest applied event. */
    OccurredAt: doublePrecision('occurredAt').notNull(),
    /** Event-log acceptance time. Together with lastEventId this is the cursor. */
    AcceptedAt: doublePrecision('acceptedAt').notNull(),
    LastEventId: text('lastEventId').notNull(),
    ProjectionVersion: text('projectionVersion').notNull(),
  },
  (table) => [uniqueIndex('TopicClusteringRunHistoryProjection_projectId_key').on(table.projectId)],
)

/** Rebuildable per-source pull checkpoint and operator-facing run status. */
export const ingestionPullRunProjection = pgTable(
  'IngestionPullRunProjection',
  {
    id: text('id').primaryKey(),
    sourceId: text('sourceId').notNull(),
    projectId: text('projectId').notNull(),
    Enabled: boolean('enabled').notNull().default(false),
    Cron: text('cron'),
    Cursor: text('cursor'),
    LastRunAt: doublePrecision('lastRunAt'),
    LastRunOutcome: text('lastRunOutcome'),
    LastRunEventCount: integer('lastRunEventCount').notNull().default(0),
    LastRunError: text('lastRunError'),
    LastRunErrorCode: text('lastRunErrorCode'),
    ConsecutiveErrors: integer('consecutiveErrors').notNull().default(0),
    /**
     * When a run last completed without erroring. Backfilled by replay, which
     * is why the projection version moved with this column.
     */
    LastSuccessAt: doublePrecision('lastSuccessAt'),
    /**
     * Identity of the run whose outcome this row reflects. Outcomes from a
     * superseded run arrive with a smaller value and are ignored, so a late
     * completion cannot regress Cursor or the status fields.
     */
    LastRunScheduledFor: doublePrecision('lastRunScheduledFor'),
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
    LastReadThroughAt: doublePrecision('lastReadThroughAt'),
    LastRunCompleteness: text('lastRunCompleteness'),
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
    LastAgentsListingAt: doublePrecision('lastAgentsListingAt'),
    LastAgentsListingOutcome: text('lastAgentsListingOutcome'),
    LastAgentsListingCount: integer('lastAgentsListingCount'),
    LastAgentsListingReason: text('lastAgentsListingReason'),
    LastAgentsListingStatus: integer('lastAgentsListingStatus'),
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
    LastPeopleListingAt: doublePrecision('lastPeopleListingAt'),
    LastPeopleListingOutcome: text('lastPeopleListingOutcome'),
    LastPeopleDirectoryCount: integer('lastPeopleDirectoryCount'),
    LastPeopleWithheldCount: integer('lastPeopleWithheldCount'),
    LastPeopleListingReason: text('lastPeopleListingReason'),
    LastPeopleListingStatus: integer('lastPeopleListingStatus'),
    CreatedAt: doublePrecision('createdAt').notNull(),
    UpdatedAt: doublePrecision('updatedAt').notNull(),
    OccurredAt: doublePrecision('occurredAt').notNull(),
    AcceptedAt: doublePrecision('acceptedAt').notNull(),
    LastEventId: text('lastEventId').notNull(),
    ProjectionVersion: text('projectionVersion').notNull(),
  },
  (table) => [
    uniqueIndex('IngestionPullRunProjection_sourceId_key').on(table.sourceId),
    index('IngestionPullRunProjection_projectId_idx').on(table.projectId),
  ],
)

export const scheduledJob = pgTable(
  'ScheduledJob',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    projectId: text('projectId').notNull(),
    targetType: text('targetType').notNull(),
    targetId: text('targetId').notNull(),
    cron: text('cron').notNull(),
    timezone: text('timezone').notNull(),
    nextRunAt: timestamp('nextRunAt', { precision: 3 }).notNull(),
    lastSlot: timestamp('lastSlot', { precision: 3 }),
    currentSlot: timestamp('currentSlot', { precision: 3 }),
    attempts: integer('attempts').notNull().default(0),
    lastError: text('lastError'),
    active: boolean('active').notNull().default(true),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('ScheduledJob_active_nextRunAt_idx').on(table.active, table.nextRunAt),
    index('ScheduledJob_projectId_idx').on(table.projectId),
    uniqueIndex('ScheduledJob_targetType_targetId_key').on(table.targetType, table.targetId),
  ],
)

export const bugReport = pgTable(
  'BugReport',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    createdAt: timestamp('createdAt', { withTimezone: true, precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    source: text('source').notNull(),
    kind: text('kind').notNull(),
    title: text('title').notNull(),
    summary: text('summary'),
    sessionData: text('sessionData'),
    sessionTruncated: boolean('sessionTruncated').notNull().default(false),
    agent: text('agent'),
    contactEmail: text('contactEmail'),
    cliVersion: text('cliVersion'),
    linkedProjectId: text('linkedProjectId'),
    metadata: jsonb('metadata'),
  },
  (table) => [index('BugReport_createdAt_idx').on(table.createdAt)],
)

export const idempotencyReceipt = pgTable(
  'IdempotencyReceipt',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    scopeId: text('scopeId').notNull(),
    key: text('key').notNull(),
    claimId: text('claimId').notNull(),
    requestFingerprint: text('requestFingerprint').notNull(),
    responseStatus: integer('responseStatus'),
    responseBody: text('responseBody'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    heartbeatAt: timestamp('heartbeatAt', { precision: 3 }).notNull(),
    expiresAt: timestamp('expiresAt', { precision: 3 }).notNull(),
  },
  (table) => [
    index('IdempotencyReceipt_expiresAt_idx').on(table.expiresAt),
    uniqueIndex('IdempotencyReceipt_scopeId_key_key').on(table.scopeId, table.key),
  ],
)

export const gatewayRealtimeSession = pgTable(
  'GatewayRealtimeSession',
  {
    id: text('id').primaryKey(),
    projectId: text('projectId').notNull(),
    organizationId: text('organizationId').notNull(),
    virtualKeyId: text('virtualKeyId').notNull(),
    modelProviderId: text('modelProviderId').notNull(),
    vendor: text('vendor').notNull(),
    agentId: text('agentId'),
    model: text('model').notNull(),
    requestedModel: text('requestedModel'),
    vendorConversationId: text('vendorConversationId'),
    status: gatewayRealtimeSessionStatusEnum('status').notNull().default('OPEN'),
    mintedAt: timestamp('mintedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    closedAt: timestamp('closedAt', { precision: 3 }),
    closeReason: text('closeReason'),
    traceId: text('traceId'),
    vendorCostRaw: jsonb('vendorCostRaw'),
    kind: text('kind'),
    metering: text('metering'),
    credentialExpiresAt: timestamp('credentialExpiresAt', { precision: 3 }),
    transcriptionModel: text('transcriptionModel'),
    endUserId: text('endUserId'),
    lastReportAt: timestamp('lastReportAt', { precision: 3 }),
    reportedCostNanoUsd: bigint('reportedCostNanoUsd', { mode: 'bigint' })
      .notNull()
      .default(sql`0`),
    reportCount: integer('reportCount').notNull().default(0),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('GatewayRealtimeSession_virtualKeyId_status_mintedAt_idx').on(
      table.virtualKeyId,
      table.status,
      table.mintedAt,
    ),
    index('GatewayRealtimeSession_projectId_status_mintedAt_idx').on(
      table.projectId,
      table.status,
      table.mintedAt,
    ),
    index('GatewayRealtimeSession_status_metering_mintedAt_idx').on(
      table.status,
      table.metering,
      table.mintedAt,
    ),
    uniqueIndex('GatewayRealtimeSession_vendor_vendorConversationId_key').on(
      table.vendor,
      table.vendorConversationId,
    ),
  ],
)

export const gatewayRealtimeSessionReport = pgTable(
  'GatewayRealtimeSessionReport',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    sessionId: text('sessionId').notNull(),
    reportKey: text('reportKey').notNull(),
    projectId: text('projectId').notNull(),
    model: text('model').notNull(),
    usage: jsonb('usage').notNull(),
    costNanoUsd: bigint('costNanoUsd', { mode: 'bigint' }).notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('GatewayRealtimeSessionReport_projectId_idx').on(table.projectId),
    uniqueIndex('GatewayRealtimeSessionReport_sessionId_reportKey_key').on(
      table.sessionId,
      table.reportKey,
    ),
  ],
)

export const systemMigrationTenantState = pgTable(
  'SystemMigrationTenantState',
  {
    migrationName: text('migrationName').notNull(),
    tenantId: text('tenantId').notNull(),
    status: text('status').notNull(),
    report: jsonb('report'),
    occurredAt: timestamp('occurredAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    primaryKey({
      name: 'SystemMigrationTenantState_pkey',
      columns: [table.migrationName, table.tenantId],
    }),
    index('SystemMigrationTenantState_migrationName_status_idx').on(
      table.migrationName,
      table.status,
    ),
  ],
)

export const systemMigrationEnrollment = pgTable(
  'SystemMigrationEnrollment',
  {
    organizationId: text('organizationId').notNull(),
    migrationName: text('migrationName').notNull(),
    enrolledByUserId: text('enrolledByUserId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    primaryKey({
      name: 'SystemMigrationEnrollment_pkey',
      columns: [table.organizationId, table.migrationName],
    }),
    index('SystemMigrationEnrollment_migrationName_idx').on(table.migrationName),
  ],
)

export const storedObject = pgTable(
  'StoredObject',
  {
    tenantId: text('tenantId').notNull(),
    id: text('id').notNull(),
    status: text('status').notNull(),
    purpose: text('purpose').notNull(),
    ownerKind: text('ownerKind').notNull(),
    ownerId: text('ownerId').notNull(),
    filename: text('filename').notNull(),
    sha256: text('sha256').notNull(),
    sizeBytes: bigint('sizeBytes', { mode: 'bigint' }).notNull(),
    mediaType: text('mediaType').notNull(),
    mediaTypeVerified: boolean('mediaTypeVerified').notNull().default(false),
    storageProvider: text('storageProvider'),
    storageDestinationId: text('storageDestinationId'),
    storageProviderRelativeId: text('storageProviderRelativeId'),
    generation: integer('generation').notNull().default(0),
    audiences: jsonb('audiences').notNull(),
    expiresAt: timestamp('expiresAt', { precision: 3 }),
    availableAt: timestamp('availableAt', { precision: 3 }),
    deletedAt: timestamp('deletedAt', { precision: 3 }),
    source: text('source').notNull(),
    legacyFingerprint: text('legacyFingerprint'),
    createdAt: timestamp('createdAt', { precision: 3 }).notNull(),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    primaryKey({ name: 'StoredObject_pkey', columns: [table.tenantId, table.id] }),
    index('StoredObject_tenantId_status_id_idx').on(table.tenantId, table.status, table.id),
    index('StoredObject_ownerKind_ownerId_status_idx').on(
      table.ownerKind,
      table.ownerId,
      table.status,
    ),
    index('StoredObject_status_expiresAt_idx').on(table.status, table.expiresAt),
  ],
)

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
export const discoveredPerson = pgTable(
  'DiscoveredPerson',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    /** "anthropic" | "openai" | "databricks" | ... - the provider that named them. */
    provider: text('provider').notNull(),
    /**
     * What the provider calls them, verbatim. The join key onto the money rows'
     * RawActorId. Replaced by a deterministic pseudonym on erasure, which keeps
     * the row addressable for spend attribution without keeping the identifier.
     */
    rawActorId: text('rawActorId').notNull(),
    /** Name or email text as seen. Pseudonymized alongside rawActorId on erasure. */
    displayText: text('displayText').notNull(),
    /**
     * person | service_account - determined deterministically per ADR-128 §10
     * (a bare UUID in Databricks `executed_by` is a service account, never an
     * agent), never guessed from the shape of a name.
     */
    kind: text('kind').notNull(),
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
    department: text('department'),
    firstSeenAt: timestamp('firstSeenAt', { precision: 3 }).notNull(),
    lastSeenAt: timestamp('lastSeenAt', { precision: 3 }).notNull(),
    /**
     * Set when this row's identity has been erased. The row survives so its
     * spend stays attributable to *someone*; the identifier does not.
     */
    erasedAt: timestamp('erasedAt', { precision: 3 }),
    /**
     * Set alongside `erasedAt` and cleared once the daily cost rows have been
     * removed and their rebuild requested. Non-null means the ClickHouse half
     * of the erasure did not finish — a crash between the delete and the
     * replay leaves days short by the erased amount, and without this marker a
     * re-run reads `erasedAt`, short-circuits, and reports a clean erasure that
     * did no work. The re-run resumes at that step instead.
     */
    moneyRowsPendingAt: timestamp('moneyRowsPendingAt', { precision: 3 }),
    /**
     * The earliest day to rebuild from, recorded BEFORE the rows are deleted.
     * Once they are gone there is nothing left to ask which days they were on,
     * so a resume with no record of this could only replay everything or
     * nothing. Null means nothing was rebuildable — either no day carried the
     * identifier, or every one of them is older than the event log keeps.
     */
    moneyRebuildSince: text('moneyRebuildSince'),
    /**
     * ADR-128 §12 conflict rule: automatic linking has stopped for this
     * identity because evidence pointed at two candidates, or contradicted an
     * existing link. Stored HERE rather than on the suggestion rows, because a
     * halt on auto-linking that a recompute or a restart clears is not a halt.
     */
    suspendedAt: timestamp('suspendedAt', { precision: 3 }),
    /** What tripped the suspension, for the human who reviews it. */
    suspendedReason: text('suspendedReason'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('DiscoveredPerson_organizationId_lastSeenAt_idx').on(
      table.organizationId,
      table.lastSeenAt,
    ),
    uniqueIndex('DiscoveredPerson_organizationId_provider_rawActorId_key').on(
      table.organizationId,
      table.provider,
      table.rawActorId,
    ),
  ],
)

/**
 * The same shape as DiscoveredPerson for non-human actors - Genie spaces,
 * Copilot bots, OpenAI projects - plus whatever provider-native metadata the
 * source published. Discovering an agent grants it nothing either.
 */
export const discoveredAgent = pgTable(
  'DiscoveredAgent',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    provider: text('provider').notNull(),
    /** Genie space / Copilot bot / OpenAI project id, verbatim. */
    rawAgentId: text('rawAgentId').notNull(),
    displayText: text('displayText').notNull(),
    /**
     * Workspace, plan, and kin - provider-native descriptive fields only. Never
     * credentials: this row is read by admin screens, not by any pull path.
     */
    metadata: jsonb('metadata').notNull(),
    firstSeenAt: timestamp('firstSeenAt', { precision: 3 }).notNull(),
    lastSeenAt: timestamp('lastSeenAt', { precision: 3 }).notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('DiscoveredAgent_organizationId_lastSeenAt_idx').on(
      table.organizationId,
      table.lastSeenAt,
    ),
    uniqueIndex('DiscoveredAgent_organizationId_provider_rawAgentId_key').on(
      table.organizationId,
      table.provider,
      table.rawAgentId,
    ),
  ],
)

/**
 * One agent in the organisation's governance registry, from any source governance reads
 * (enterprise/modules/governance/specs/agent-registry.feature). `entry` is the latest
 * machine-readable snapshot; `version` grows by one on every material change.
 */
export const agentRegistryEntry = pgTable(
  'AgentRegistryEntry',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    /** Where the agent came from and its id there, e.g. `langwatch:<agentId>`. */
    sourceKey: text('sourceKey').notNull(),
    source: text('source').notNull(),
    status: text('status').notNull(),
    version: integer('version').notNull(),
    fingerprint: text('fingerprint').notNull(),
    entry: jsonb('entry').notNull(),
    firstSeenAt: timestamp('firstSeenAt', { precision: 3 }).notNull(),
    lastSeenAt: timestamp('lastSeenAt', { precision: 3 }),
    removedAt: timestamp('removedAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('AgentRegistryEntry_organizationId_updatedAt_idx').on(
      table.organizationId,
      table.updatedAt,
    ),
    uniqueIndex('AgentRegistryEntry_organizationId_sourceKey_key').on(
      table.organizationId,
      table.sourceKey,
    ),
  ],
)

/**
 * The registry's append-only change feed: one row per registered, versioned, removed or
 * tool-deprecated fact. `sequence` is the pull cursor; `deliveryRecordedAt` marks the webhook
 * fan-out as recorded in governance's outbox.
 */
export const agentRegistryEvent = pgTable(
  'AgentRegistryEvent',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    sequence: bigserial('sequence', { mode: 'bigint' }),
    type: text('type').notNull(),
    entryId: text('entryId').notNull(),
    fromVersion: integer('fromVersion'),
    toVersion: integer('toVersion'),
    payload: jsonb('payload').notNull(),
    occurredAt: timestamp('occurredAt', { precision: 3 }).notNull(),
    deliveryRecordedAt: timestamp('deliveryRecordedAt', { precision: 3 }),
  },
  (table) => [
    index('AgentRegistryEvent_organizationId_deliveryRecordedAt_idx').on(
      table.organizationId,
      table.deliveryRecordedAt,
    ),
    index('AgentRegistryEvent_organizationId_entryId_idx').on(table.organizationId, table.entryId),
    uniqueIndex('AgentRegistryEvent_organizationId_sequence_key').on(
      table.organizationId,
      table.sequence,
    ),
  ],
)

/**
 * An organisation's registry baseline: its first reconcile, written as already pushed so the
 * agents it already had are not announced. Every reconcile after it pushes its changes.
 */
export const agentRegistryBaseline = pgTable('AgentRegistryBaseline', {
  organizationId: text('organizationId').primaryKey(),
  baselinedAt: timestamp('baselinedAt', { precision: 3 }).notNull(),
})

/**
 * One target of a registry agent's kill switch stop plan (enterprise/modules/governance/specs/
 * agent-kill-switch.feature): a Databricks serving endpoint, a Bedrock agent alias or a gateway
 * virtual key. A plan write replaces every row of the agent at once.
 */
export const agentStopPlanTarget = pgTable(
  'AgentStopPlanTarget',
  {
    id: text('id').primaryKey(),
    organizationId: text('organizationId').notNull(),
    /** The AgentRegistryEntry the plan belongs to. */
    agentId: text('agentId').notNull(),
    position: integer('position').notNull(),
    label: text('label').notNull(),
    target: jsonb('target').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('AgentStopPlanTarget_organizationId_agentId_idx').on(table.organizationId, table.agentId),
  ],
)

/**
 * One agent's stop plan version: bumped by every save, so a save over a plan someone else saved
 * since it was read is refused. No row means the plan was never saved (version 0).
 */
export const agentStopPlan = pgTable(
  'AgentStopPlan',
  {
    organizationId: text('organizationId').notNull(),
    /** The AgentRegistryEntry the plan belongs to. */
    agentId: text('agentId').notNull(),
    version: integer('version').notNull(),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    primaryKey({ name: 'AgentStopPlan_pkey', columns: [table.organizationId, table.agentId] }),
  ],
)

/**
 * An AWS credential a Bedrock kill switch target signs with: a static key pair, a role to assume,
 * or both. The secret access key and the external id are stored sealed.
 */
export const governanceAwsConnection = pgTable(
  'GovernanceAwsConnection',
  {
    id: text('id').primaryKey(),
    organizationId: text('organizationId').notNull(),
    name: text('name').notNull(),
    accessKeyId: text('accessKeyId'),
    encryptedSecretAccessKey: text('encryptedSecretAccessKey'),
    roleArn: text('roleArn'),
    encryptedExternalId: text('encryptedExternalId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [index('GovernanceAwsConnection_organizationId_idx').on(table.organizationId)],
)

/**
 * The `agent_stop` aggregate's Postgres projection: one row per kill switch stop, its status and
 * every target's outcome, folded from the stop's events.
 */
export const agentStop = pgTable(
  'AgentStop',
  {
    id: text('id').primaryKey(),
    organizationId: text('organizationId').notNull(),
    agentId: text('agentId').notNull(),
    agentName: text('agentName').notNull(),
    status: text('status').notNull(),
    severity: text('severity').notNull(),
    approval: text('approval').notNull(),
    trigger: text('trigger').notNull(),
    reason: text('reason').notNull(),
    requestedByUserId: text('requestedByUserId'),
    requestedByApiKeyId: text('requestedByApiKeyId'),
    decidedByUserId: text('decidedByUserId'),
    decidedByApiKeyId: text('decidedByApiKeyId'),
    restoredByUserId: text('restoredByUserId'),
    restoredByApiKeyId: text('restoredByApiKeyId'),
    targets: jsonb('targets').notNull(),
    requestedAt: timestamp('requestedAt', { precision: 3 }).notNull(),
    decidedAt: timestamp('decidedAt', { precision: 3 }),
    quarantinedAt: timestamp('quarantinedAt', { precision: 3 }),
    restoreRequestedAt: timestamp('restoreRequestedAt', { precision: 3 }),
    restoredAt: timestamp('restoredAt', { precision: 3 }),
    lastAppliedEventId: text('lastAppliedEventId').notNull(),
    lastAcceptedAt: bigint('lastAcceptedAt', { mode: 'bigint' }).notNull(),
    lastOccurredAt: bigint('lastOccurredAt', { mode: 'bigint' }).notNull(),
    projectionVersion: text('projectionVersion').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
  },
  (table) => [
    index('AgentStop_organizationId_agentId_idx').on(table.organizationId, table.agentId),
    index('AgentStop_organizationId_status_idx').on(table.organizationId, table.status),
  ],
)

/**
 * The one open stop an agent may have: the api takes this row before it sends a stop's request,
 * so two requests at once cannot both open one. A later request takes it over once the stop it
 * names has closed.
 * `sentAt` is set once the request command was sent: from then on the claim passes on only after
 * its stop closed, never on age. `decision` is taken by the one approve or deny that wins.
 * `decisionTakenAt` is set while that decision is in flight and cleared once its command was
 * sent; a decision left in flight is taken over once abandoned.
 */
export const agentStopClaim = pgTable(
  'AgentStopClaim',
  {
    organizationId: text('organizationId').notNull(),
    agentId: text('agentId').notNull(),
    stopId: text('stopId').notNull(),
    claimedAt: timestamp('claimedAt', { precision: 3 }).notNull(),
    sentAt: timestamp('sentAt', { precision: 3 }),
    decision: text('decision'),
    decisionTakenAt: timestamp('decisionTakenAt', { precision: 3 }),
  },
  (table) => [
    primaryKey({ name: 'AgentStopClaim_pkey', columns: [table.organizationId, table.agentId] }),
  ],
)

/**
 * Whether a target was already stopped before a stop first touched it, recorded once before the
 * stop's platform call, so a retried call after a success still knows the change is the stop's own.
 */
export const agentStopTargetPriorState = pgTable(
  'AgentStopTargetPriorState',
  {
    organizationId: text('organizationId').notNull(),
    stopId: text('stopId').notNull(),
    targetId: text('targetId').notNull(),
    wasStopped: boolean('wasStopped').notNull(),
    recordedAt: timestamp('recordedAt', { precision: 3 }).notNull(),
  },
  (table) => [
    primaryKey({
      name: 'AgentStopTargetPriorState_pkey',
      columns: [table.organizationId, table.stopId, table.targetId],
    }),
  ],
)

/**
 * The dated match table: "`user_abc123` on Anthropic is platform user #42,
 * valid from March 3", closed when it stops being true.
 *
 * Dating is what makes a re-issued email address survivable: the leaver's link
 * closes at offboarding, the new hire gets a new link, and January's spend
 * stays with January's person. A correction closes the old row and opens a new
 * one; nothing is ever silently rewritten.
 */
export const identityMatch = pgTable(
  'IdentityMatch',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    /** The provider-side identity this link is about. */
    discoveredPersonId: text('discoveredPersonId').notNull(),
    /**
     * The platform user. Nullable, and blanked by erasure (ADR-128 §9 step 2) -
     * the row and its dates remain, so the history of the link is not rewritten
     * by the erasure either.
     */
    userId: text('userId'),
    /**
     * directory_id | verified_email | human_confirmed. Exact verified-email
     * equality is the primary evidence; a directory id corroborates but never
     * stands alone (ADR-128 §12).
     */
    evidenceKind: text('evidenceKind').notNull(),
    validFrom: timestamp('validFrom', { precision: 3 }).notNull(),
    /** Null means the link is open. Offboarding and corrections close it. */
    validTo: timestamp('validTo', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('IdentityMatch_organizationId_discoveredPersonId_idx').on(
      table.organizationId,
      table.discoveredPersonId,
    ),
    index('IdentityMatch_organizationId_userId_idx').on(table.organizationId, table.userId),
  ],
)

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
export const governanceTenantHistory = pgTable(
  'GovernanceTenantHistory',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    /** A hidden governance project id this organization has written rows under. */
    tenantId: text('tenantId').notNull(),
    firstUsedAt: timestamp('firstUsedAt', { precision: 3 }).notNull(),
    lastUsedAt: timestamp('lastUsedAt', { precision: 3 }).notNull(),
  },
  (table) => [
    index('GovernanceTenantHistory_organizationId_idx').on(table.organizationId),
    uniqueIndex('GovernanceTenantHistory_organizationId_tenantId_key').on(
      table.organizationId,
      table.tenantId,
    ),
  ],
)

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
export const erasedIdentifierSuppression = pgTable(
  'ErasedIdentifierSuppression',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    provider: text('provider').notNull(),
    identifierHash: text('identifierHash').notNull(),
    erasedAt: timestamp('erasedAt', { precision: 3 }).notNull(),
  },
  (table) => [
    index('ErasedIdentifierSuppression_organizationId_idx').on(table.organizationId),
    uniqueIndex('ErasedIdentifierSuppression_organizationId_provider_identif_key').on(
      table.organizationId,
      table.provider,
      table.identifierHash,
    ),
  ],
)

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
export const identityMatchSuggestion = pgTable(
  'IdentityMatchSuggestion',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    organizationId: text('organizationId').notNull(),
    discoveredPersonId: text('discoveredPersonId').notNull(),
    /**
     * The platform user this might be. Not nullable: a suggestion with no
     * candidate is not a suggestion.
     */
    userId: text('userId').notNull(),
    /**
     * Similarity of the two name texts, after the prefilter, in [0, 1]. A
     * display hint and an ordering key - never a threshold anything auto-links
     * on, which is the whole distinction this table exists to hold.
     */
    score: doublePrecision('score').notNull(),
    computedAt: timestamp('computedAt', { precision: 3 }).notNull(),
  },
  (table) => [
    index('IdentityMatchSuggestion_organizationId_score_idx').on(table.organizationId, table.score),
    uniqueIndex('IdentityMatchSuggestion_organizationId_discoveredPersonId_u_key').on(
      table.organizationId,
      table.discoveredPersonId,
      table.userId,
    ),
  ],
)

/**
 * What a mid-term seat change of one connected license owes (ADR-141,
 * section 7): the seats added on the reissued license, invoiced prorated to
 * the end of the term.
 *
 * The row is written before the payment provider is called and completed
 * after it, so a call that failed halfway is retried by the next seat
 * invoicing pass without inventing a second invoice. One row per reissued license.
 */
export const connectedSeatChange = pgTable(
  'ConnectedSeatChange',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    /**
     * The reissued `IssuedLicense` row that carries the new seat count. A
     * plain column ref, like `replacesId`: relationMode="prisma" carries no
     * foreign keys.
     */
    licenseId: text('licenseId').notNull(),
    /**
     * The `ConnectedBillingAccount` row the invoice is raised against. Null
     * when the customer had no billing account (`not_onboarded`).
     */
    accountId: text('accountId'),
    /** The customer organization the license belongs to. */
    organizationId: text('organizationId'),
    /** When the seats changed, which is the day the proration counts from. */
    changedAt: timestamp('changedAt', { precision: 3 }).notNull(),
    /** Seats on the new license above the seats on the one it replaced. */
    addedSeats: integer('addedSeats').notNull(),
    /** What one added seat costs for the rest of the term, in cents of `currency`. */
    unitAmountCents: integer('unitAmountCents').notNull(),
    /** `unitAmountCents` times `addedSeats`. Zero when nothing is owed. */
    amountCents: integer('amountCents').notNull(),
    /** Null when no billing account named one (`not_onboarded`). */
    currency: currencyEnum('currency'),
    /** The one-off invoice the added seats were billed on. */
    stripeInvoiceId: text('stripeInvoiceId'),
    state: connectedSeatChangeStateEnum('state').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('ConnectedSeatChange_licenseId_key').on(table.licenseId),
    index('ConnectedSeatChange_accountId_idx').on(table.accountId),
    index('ConnectedSeatChange_organizationId_idx').on(table.organizationId),
    index('ConnectedSeatChange_state_idx').on(table.state),
  ],
)

/**
 * The monthly usage statement one connected customer was sent (ADR-139,
 * section 7). One row per account per month keeps the statement to one send,
 * whatever runs the tick.
 */
export const connectedStatement = pgTable(
  'ConnectedStatement',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    /** The `ConnectedBillingAccount` row. A plain column ref. */
    accountId: text('accountId').notNull(),
    /** First day of the month the statement covers, in UTC. */
    month: timestamp('month', { precision: 3 }).notNull(),
    sentAt: timestamp('sentAt', { precision: 3 }).notNull(),
  },
  (table) => [
    index('ConnectedStatement_accountId_idx').on(table.accountId),
    uniqueIndex('ConnectedStatement_accountId_month_key').on(table.accountId, table.month),
  ],
)

/**
 * One renewed licence per paid Stripe invoice of a self-serve licence
 * subscription (specs/licensing/self-serving-license-purchase.feature). A
 * redelivered invoice event finds the row and mints nothing. A claim never
 * delivered is taken over once `claimedAt` is older than the claim lease.
 */
export const licenseRenewalClaim = pgTable('LicenseRenewalClaim', {
  stripeInvoiceId: text('stripeInvoiceId').primaryKey(),
  stripeSubscriptionId: text('stripeSubscriptionId').notNull(),
  /** Fences completion and release to the delivery that holds the claim. */
  claimToken: text('claimToken').notNull(),
  claimedAt: timestamp('claimedAt', { precision: 3 }).notNull(),
  /** The signed licence between minting and mailing, so a retry mails it again instead of minting another; cleared once delivered. */
  issuedLicense: jsonb('issuedLicense'),
  deliveredAt: timestamp('deliveredAt', { precision: 3 }),
  /** The `licenseId` inside the delivered licence. */
  licenseId: text('licenseId'),
})

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
export const instanceIdentity = pgTable(
  'InstanceIdentity',
  {
    /**
     * Fixed, so this table holds one row. A second insert loses the race on the
     * primary key and re-reads the winner's id.
     */
    id: text('id').primaryKey().default('self'),
    /** A random UUID. Carries nothing about the customer. */
    instanceId: text('instanceId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    /**
     * When the usage report last reached LangWatch, and why the last attempt
     * did not. Both are what the checkup page reads to say whether reporting
     * works, which nothing recorded before: the sender logged success as soon
     * as `fetch` resolved and never looked at the response.
     */
    lastReportAt: timestamp('lastReportAt', { precision: 3 }),
    lastReportError: text('lastReportError'),
    /**
     * The optional category of the usage report, switched off from the checkup
     * page. With it off the report still says which release runs and how many
     * projects the install carries, and says nothing about what anyone does
     * with it (ADR-139, section 10).
     */
    optionalMetricsOptOut: boolean('optionalMetricsOptOut').notNull().default(false),
    /**
     * Hostname, which names the customer's own network and so has a switch of
     * its own rather than travelling with its category.
     */
    hostnameOptOut: boolean('hostnameOptOut').notNull().default(false),
    /**
     * Sending a failed upgrade step's diagnostics to LangWatch on its own,
     * switched off by the operator (upgrade-diagnostics-via-connect, decision 4).
     */
    upgradeDiagnosticsOptOut: boolean('upgradeDiagnosticsOptOut').notNull().default(false),
  },
  (table) => [uniqueIndex('InstanceIdentity_instanceId_key').on(table.instanceId)],
)

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
export const selfHostedInstance = pgTable(
  'SelfHostedInstance',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    /** The UUID the install minted for itself. The join key for everything else. */
    instanceId: text('instanceId').notNull(),
    firstSeenAt: timestamp('firstSeenAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    lastSeenAt: timestamp('lastSeenAt', { precision: 3 }).notNull(),
    /** What the install says about itself. */
    version: text('version'),
    installMethod: text('installMethod'),
    chartVersion: text('chartVersion'),
    /** The host it answers on, when the operator left hostname reporting on. */
    hostname: text('hostname'),
    environment: text('environment'),
    /**
     * The day the install itself says it first came up, which predates the
     * first report on an install that switched reporting on later.
     */
    installedAt: timestamp('installedAt', { precision: 3 }),
    reportSchemaVersion: integer('reportSchemaVersion'),
    /**
     * The customer, and the license that named them. Plain column refs rather
     * than relations, like `IssuedLicense.replacesId`: relationMode="prisma"
     * carries no foreign keys, and the receiver resolves both from the license
     * registry rather than from the report.
     */
    organizationId: text('organizationId'),
    issuedLicenseId: text('issuedLicenseId'),
    /**
     * Domains with counts, as `{"acme.test": 14}`. Never an address. Null on an
     * install that switched the optional category off.
     */
    userEmailDomains: jsonb('userEmailDomains'),
    /**
     * The same domains as the keys of `userEmailDomains`, flattened so an
     * operator can search for one. A JSON object's keys are the domains
     * themselves, and no query language addresses a key it does not already
     * know the name of.
     */
    userDomains: text('userDomains').array().default([]),
    /**
     * What the last report carried, so the list and the drawer read one row:
     * the onboarding ladder, the usage counts and the operational block, under
     * the same keys the dictionary declares.
     */
    latestReport: jsonb('latestReport'),
    /**
     * What the last report said the operator had switched off, read from which
     * categories arrived rather than from a field the install could assert.
     */
    optionalMetricsReported: boolean('optionalMetricsReported').notNull().default(true),
    hostnameReported: boolean('hostnameReported').notNull().default(true),
    /**
     * How many reports have landed, and how many fields the last one carried
     * that this release has no name for. A count above zero says the senders
     * are ahead of the receiver.
     */
    reportCount: integer('reportCount').notNull().default(0),
    lastUnknownFields: integer('lastUnknownFields').notNull().default(0),
    /**
     * The lead signals already raised for this install. A signal on this list
     * is never raised again, which is what keeps a large install from
     * reporting that it is large every morning.
     */
    raisedSignals: text('raisedSignals').array().default([]),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('SelfHostedInstance_instanceId_key').on(table.instanceId),
    index('SelfHostedInstance_lastSeenAt_idx').on(table.lastSeenAt),
    index('SelfHostedInstance_organizationId_idx').on(table.organizationId),
    index('SelfHostedInstance_issuedLicenseId_idx').on(table.issuedLicenseId),
  ],
)

/**
 * One upgrade diagnostics bundle a self-hosted install sent over Connect, on
 * LangWatch Cloud. The bytes live in object storage under
 * `upgrade-diagnostics/` and expire at 30 days by a lifecycle rule; this row
 * points at them and feeds Cloud admin and the Slack signal. Plain column
 * refs, no relation, as the usage report rows.
 */
export const selfHostedInstanceDiagnostic = pgTable(
  'SelfHostedInstanceDiagnostic',
  {
    id: text('id').primaryKey(),
    /** The `SelfHostedInstance.instanceId` the licence caller presented. */
    instanceId: text('instanceId').notNull(),
    stepId: text('stepId').notNull(),
    stepDescription: text('stepDescription'),
    release: text('release'),
    attempt: integer('attempt').notNull(),
    /** `operator` or `automatic`. */
    trigger: text('trigger').notNull(),
    byteLength: integer('byteLength').notNull(),
    objectKey: text('objectKey').notNull(),
    receivedAt: timestamp('receivedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    /** When the lifecycle rule removes the bundle; the row is listed until then. */
    expiresAt: timestamp('expiresAt', { precision: 3 }).notNull(),
    /** Set on the report that posted this step and release's one Slack message. */
    slackAnnouncedAt: timestamp('slackAnnouncedAt', { precision: 3 }),
  },
  (table) => [
    index('SelfHostedInstanceDiagnostic_instanceId_receivedAt_idx').on(
      table.instanceId,
      table.receivedAt,
    ),
  ],
)

/**
 * A notice LangWatch staff send every self-hosted install in Cloud's reply to its daily usage
 * report (enterprise/modules/licensing/specs/self-hosted-notices.feature). Cloud only, and no
 * tenant: every install receives every published notice that has not expired.
 */
export const selfHostedNotice = pgTable('SelfHostedNotice', {
  id: text('id').primaryKey(),
  /** `info`, `warning` or `security`. */
  severity: text('severity').notNull(),
  title: text('title').notNull(),
  body: text('body').notNull(),
  /** An https link to read more; null where the notice names none. */
  url: text('url'),
  /** The release that fixes what the notice is about; an install on it or newer hides it. */
  fixedIn: text('fixedIn'),
  /** `draft`, `published` or `withdrawn`. Only a published notice that has not expired is sent. */
  status: text('status').notNull(),
  expiresAt: timestamp('expiresAt', { precision: 3 }),
  publishedAt: timestamp('publishedAt', { precision: 3 }),
  /** The staff member who published it, by user id. */
  publishedBy: text('publishedBy'),
  createdAt: timestamp('createdAt', { precision: 3 })
    .notNull()
    .default(sql`CURRENT_TIMESTAMP`)
    .$defaultFn(utcNow),
  updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
})

/**
 * One usage report, exactly as it arrived. Kept so a number on an install's
 * row can be traced back to the report that set it, and so a question asked
 * later can be answered over the history rather than over the latest snapshot.
 */
export const selfHostedInstanceReport = pgTable(
  'SelfHostedInstanceReport',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    /**
     * The `SelfHostedInstance.instanceId` this report came from. A plain column
     * ref for the same reason the instance row's refs are.
     */
    instanceId: text('instanceId').notNull(),
    receivedAt: timestamp('receivedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    version: text('version'),
    reportSchemaVersion: integer('reportSchemaVersion'),
    /** Fields this release had no name for. Counted, never stored by name. */
    unknownFields: integer('unknownFields').notNull().default(0),
    /** The validated report, with every unknown field already dropped. */
    payload: jsonb('payload').notNull(),
  },
  (table) => [
    index('SelfHostedInstanceReport_instanceId_receivedAt_idx').on(
      table.instanceId,
      table.receivedAt,
    ),
    index('SelfHostedInstanceReport_receivedAt_idx').on(table.receivedAt),
  ],
)

/**
 * One browser a person subscribed to Web Push, owned by notification. A
 * person holds one row per device; the endpoint names the browser, so a
 * browser that subscribes again replaces its own row.
 */
export const webPushSubscription = pgTable(
  'WebPushSubscription',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => nanoid()),
    /**
     * The person the pushes are for. A plain column: the row leaves with the
     * person through notification's own subscriber on the user lifecycle.
     */
    userId: text('userId').notNull(),
    /** The push service URL the browser handed out. Unique per browser. */
    endpoint: text('endpoint').notNull(),
    /**
     * The browser's P-256 public key and auth secret, base64url, as the
     * browser reports them. Payloads are encrypted to these.
     */
    p256dh: text('p256dh').notNull(),
    auth: text('auth').notNull(),
    userAgent: text('userAgent'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    /** When the push service last accepted a push for this browser. */
    lastSuccessAt: timestamp('lastSuccessAt', { precision: 3 }),
  },
  (table) => [
    uniqueIndex('WebPushSubscription_endpoint_key').on(table.endpoint),
    index('WebPushSubscription_userId_idx').on(table.userId),
  ],
)

/**
 * This installation's VAPID key pair, generated on first use so Web Push works
 * without configuration. Fixed id, so the table holds one row: a second
 * insert loses the race on the primary key and reads the winner's pair.
 */
export const webPushVapidKey = pgTable('WebPushVapidKey', {
  id: text('id').primaryKey().default('self'),
  /** The uncompressed P-256 public key, base64url. The browser subscribes with it. */
  publicKey: text('publicKey').notNull(),
  /** The private key, base64url, encrypted with the process encryption key. */
  privateKeyEncrypted: text('privateKeyEncrypted').notNull(),
  createdAt: timestamp('createdAt', { precision: 3 })
    .notNull()
    .default(sql`CURRENT_TIMESTAMP`)
    .$defaultFn(utcNow),
})

export const signInProviderMount = pgTable('SignInProviderMount', {
  /**
   * Fixed, so this table holds one row: licensing's copy of sso's sign-in
   * provider fact (ADR-179).
   */
  id: text('id').primaryKey().default('self'),
  /** The configured provider's name; null when the deployment signs in by email. */
  provider: text('provider'),
  mounted: boolean('mounted').notNull(),
  /** The last sso event applied: the row's idempotency guard and its version. */
  lastAppliedEventId: text('lastAppliedEventId').notNull(),
  updatedAt: timestamp('updatedAt', { precision: 3 })
    .notNull()
    .default(sql`CURRENT_TIMESTAMP`)
    .$defaultFn(utcNow)
    .$onUpdate(utcNow),
})

export const relations = defineRelations(
  {
    account,
    accountCredential,
    session,
    userErasure,
    user,
    verificationToken,
    identifier,
    identitySignInSwitch,
    identifierReservation,
    identityProjectionCursor,
    twoFactor,
    signInAttemptLock,
    passkey,
    mfaEnrollment,
    ssoConnection,
    ssoConnectionRegistrationSlot,
    ssoConnectionReproofCursor,
    ssoCredential,
    ssoVerifiedDomain,
    ssoVerifiedDomainHolder,
    ssoAuthenticationActivity,
    ssoProvider,
    ssoBreakGlassBinding,
    ssoActivationRecoveryReservation,
    joinRequest,
    teamUser,
    organizationUser,
    seatHold,
    seatRequest,
    organizationSeatLease,
    organizationAccessNotice,
    team,
    organization,
    issuedLicense,
    organizationLicense,
    activationCode,
    connectedBillingAccount,
    connectedCreditGrant,
    connectedInvoice,
    scimRequestLog,
    scimToken,
    scimExternalId,
    scimUserResource,
    scimDirectoryUser,
    scimSyncState,
    scimSsoConnectionView,
    organizationSeatPeak,
    organizationSeatPeakHolder,
    organizationSeatPeakMonth,
    project,
    department,
    departmentMembershipHistory,
    platformToolPolicy,
    organizationInvite,
    monitor,
    cost,
    topic,
    dataset,
    datasetRecord,
    datasetVersion,
    datasetVersionRun,
    dashboard,
    dashboardFavourite,
    dashboardsV2,
    dashboardWidget,
    insightProjection,
    insightReaderProjection,
    insightDailyScheduleProjection,
    savedView,
    customGraph,
    batchEvaluation,
    trigger,
    slackIntegration,
    slackConnectionClaim,
    webhookEndpoint,
    webhookEndpointDelivery,
    experiment,
    experimentVersion,
    annotation,
    modelProvider,
    modelProviderScope,
    modelDefaultConfig,
    modelDefaultConfigScope,
    githubInstallation,
    githubPullRequest,
    githubBranchPullRequestCheck,
    langyConversationProjection,
    langyTurnRequest,
    langyActiveTurn,
    langyConversationTurnProjection,
    langyPullRequest,
    langyMessageProjection,
    projectSecret,
    triggerSent,
    triggerLatestEvaluation,
    emailSuppression,
    annotationScore,
    annotationQueue,
    annotationQueueMembers,
    annotationQueueScores,
    annotationQueueItem,
    shareLink,
    pinnedTrace,
    traceEditOverlay,
    traceIngestSourceBilling,
    retentionPolicy,
    retentionApplyStatus,
    dataPrivacyPolicy,
    instantEvalJudgeProject,
    instantEvalJudgeUsageBilling,
    instantEvalJudgeSpend,
    customLLMModelCost,
    workflow,
    workflowVersion,
    auditLog,
    llmPromptConfig,
    llmPromptConfigVersion,
    promptTagAssignment,
    promptTag,
    analytics,
    customRole,
    notification,
    agent,
    agentVersion,
    agentVersionLabel,
    evaluator,
    evaluatorVersion,
    evaluatorVersionLabel,
    scenario,
    scenarioVersion,
    simulationSuite,
    subscription,
    invoice,
    invoiceItem,
    group,
    groupMembership,
    authzUserStanding,
    authzEngineSwitch,
    authzOrganizationPolicy,
    authzBraidWrite,
    roleBinding,
    grant,
    grantUsage,
    role,
    apiKey,
    signingKey,
    projectActiveDay,
    nurturingProject,
    nurturingUncopiedProject,
    nurturingOrganizationMember,
    nurturingOrganizationCreated,
    nurturingOrganization,
    nurturingProjectUsageView,
    nurturingProjectTraceDay,
    nurturingPersonalProject,
    nurturingTracedProject,
    billingMeterCheckpoint,
    virtualKey,
    virtualKeyScope,
    routingPolicy,
    routingPolicyScope,
    gatewayGuardrail,
    ingestionSource,
    anomalyRule,
    anomalyAlert,
    aiToolEntry,
    aiToolEntryTeam,
    aiToolEntryDepartment,
    ingestionTemplate,
    gatewayA2ATarget,
    gatewayConnectUpstream,
    gatewayTraceExportKey,
    gatewayBudget,
    gatewayBudgetBucketBoundary,
    gatewayBudgetLedger,
    gatewayChangeEvent,
    gatewayCacheRule,
    featureFlag,
    featureFlagExperimentSetting,
    processManagerInstance,
    processManagerInbox,
    processManagerOutbox,
    processManagerOutboxAttempt,
    topicClusteringRunProjection,
    topicModelProjection,
    topicClusteringRunHistoryProjection,
    ingestionPullRunProjection,
    scheduledJob,
    bugReport,
    idempotencyReceipt,
    gatewayRealtimeSession,
    gatewayRealtimeSessionReport,
    systemMigrationTenantState,
    systemMigrationEnrollment,
    storedObject,
    discoveredPerson,
    discoveredAgent,
    agentRegistryEntry,
    agentRegistryEvent,
    agentRegistryBaseline,
    agentStopPlanTarget,
    agentStopPlan,
    governanceAwsConnection,
    agentStop,
    agentStopClaim,
    agentStopTargetPriorState,
    identityMatch,
    governanceTenantHistory,
    erasedIdentifierSuppression,
    identityMatchSuggestion,
    connectedSeatChange,
    connectedStatement,
    licenseRenewalClaim,
    instanceIdentity,
    selfHostedInstance,
    selfHostedInstanceDiagnostic,
    selfHostedNotice,
    selfHostedInstanceReport,
    webPushSubscription,
    webPushVapidKey,
    signInProviderMount,
  },
  (r) => ({
    account: { user: r.one.user({ from: r.account.userId, to: r.user.id, optional: false }) },
    accountCredential: {
      user: r.one.user({ from: r.accountCredential.userId, to: r.user.id, optional: false }),
    },
    session: { user: r.one.user({ from: r.session.userId, to: r.user.id, optional: false }) },
    user: {
      accounts: r.many.account(),
      accountCredentials: r.many.accountCredential(),
      sessions: r.many.session(),
      twoFactors: r.many.twoFactor(),
      passkeys: r.many.passkey(),
      teamMemberships: r.many.teamUser(),
      orgMemberships: r.many.organizationUser(),
      scimUserResources: r.many.scimUserResource(),
      Annotation: r.many.annotation(),
      shareLinks: r.many.shareLink(),
      Workflow: r.many.workflow(),
      WorkflowVersion: r.many.workflowVersion(),
      annotationQueues: r.many.annotationQueueMembers(),
      assignedQueueItems: r.many.annotationQueueItem({ alias: 'AssignedTo' }),
      createdQueueItems: r.many.annotationQueueItem({ alias: 'CreatedBy' }),
      llmPromptConfigVersions: r.many.llmPromptConfigVersion(),
      scenarios: r.many.scenario(),
      inviteRequests: r.many.organizationInvite(),
      secretsCreated: r.many.projectSecret({ alias: 'secretCreatedBy' }),
      secretsUpdated: r.many.projectSecret({ alias: 'secretUpdatedBy' }),
      savedViews: r.many.savedView(),
      versionTagsCreated: r.many.promptTagAssignment({ alias: 'versionTagCreatedBy' }),
      versionTagsUpdated: r.many.promptTagAssignment({ alias: 'versionTagUpdatedBy' }),
      promptTagsCreated: r.many.promptTag({ alias: 'promptTagCreatedBy' }),
      promptTagsUpdated: r.many.promptTag({ alias: 'promptTagUpdatedBy' }),
      groupMemberships: r.many.groupMembership(),
      roleBindings: r.many.roleBinding(),
      scimExternalIds: r.many.scimExternalId(),
      apiKeys: r.many.apiKey(),
      gatewayBudgetsCreated: r.many.gatewayBudget(),
      principalVirtualKeys: r.many.virtualKey(),
      personalTeams: r.many.team(),
      personalProjects: r.many.project(),
      routingPoliciesCreated: r.many.routingPolicy({ alias: 'RoutingPolicyCreatedBy' }),
      routingPoliciesUpdated: r.many.routingPolicy({ alias: 'RoutingPolicyUpdatedBy' }),
      gatewayGuardrailsCreated: r.many.gatewayGuardrail({ alias: 'GatewayGuardrailCreatedBy' }),
      gatewayGuardrailsUpdated: r.many.gatewayGuardrail({ alias: 'GatewayGuardrailUpdatedBy' }),
      ingestionSourcesCreated: r.many.ingestionSource(),
      anomalyRulesCreated: r.many.anomalyRule(),
      pinnedTraces: r.many.pinnedTrace(),
      traceEditOverlaysCreated: r.many.traceEditOverlay({ alias: 'TraceEditOverlayCreatedBy' }),
      traceEditOverlaysUpdated: r.many.traceEditOverlay({ alias: 'TraceEditOverlayUpdatedBy' }),
    },
    twoFactor: { user: r.one.user({ from: r.twoFactor.userId, to: r.user.id, optional: false }) },
    passkey: { user: r.one.user({ from: r.passkey.userId, to: r.user.id, optional: false }) },
    ssoConnection: {
      reproofCursor: r.one.ssoConnectionReproofCursor({
        from: r.ssoConnection.id,
        to: r.ssoConnectionReproofCursor.connectionId,
      }),
    },
    ssoConnectionReproofCursor: {
      connection: r.one.ssoConnection({
        from: r.ssoConnectionReproofCursor.connectionId,
        to: r.ssoConnection.id,
        optional: false,
      }),
    },
    ssoVerifiedDomain: { holders: r.many.ssoVerifiedDomainHolder() },
    ssoVerifiedDomainHolder: {
      ownership: r.one.ssoVerifiedDomain({
        from: [r.ssoVerifiedDomainHolder.domain, r.ssoVerifiedDomainHolder.organizationId],
        to: [r.ssoVerifiedDomain.domain, r.ssoVerifiedDomain.organizationId],
        optional: false,
      }),
    },
    teamUser: {
      assignedRole: r.one.customRole({ from: r.teamUser.assignedRoleId, to: r.customRole.id }),
      user: r.one.user({ from: r.teamUser.userId, to: r.user.id, optional: false }),
      team: r.one.team({ from: r.teamUser.teamId, to: r.team.id, optional: false }),
    },
    organizationUser: {
      user: r.one.user({ from: r.organizationUser.userId, to: r.user.id, optional: false }),
      organization: r.one.organization({
        from: r.organizationUser.organizationId,
        to: r.organization.id,
        optional: false,
      }),
    },
    team: {
      members: r.many.teamUser(),
      organization: r.one.organization({
        from: r.team.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      projects: r.many.project(),
      ownerUser: r.one.user({ from: r.team.ownerUserId, to: r.user.id }),
      ingestionSources: r.many.ingestionSource(),
      aiToolEntryTeams: r.many.aiToolEntryTeam(),
    },
    organization: {
      members: r.many.organizationUser(),
      teams: r.many.team(),
      webhookEndpoints: r.many.webhookEndpoint(),
      OrganizationInvite: r.many.organizationInvite(),
      CustomRoles: r.many.customRole(),
      notifications: r.many.notification(),
      subscriptions: r.many.subscription(),
      scimTokens: r.many.scimToken(),
      scimRequestLog: r.many.scimRequestLog(),
      promptTags: r.many.promptTag(),
      groups: r.many.group(),
      roleBindings: r.many.roleBinding(),
      apiKeys: r.many.apiKey(),
      issuedLicenses: r.many.issuedLicense(),
      connectedBilling: r.one.connectedBillingAccount({
        from: r.organization.id,
        to: r.connectedBillingAccount.organizationId,
      }),
      gatewayChangeEvents: r.many.gatewayChangeEvent(),
      routingPolicies: r.many.routingPolicy(),
      ingestionSources: r.many.ingestionSource(),
      anomalyRules: r.many.anomalyRule(),
      anomalyAlerts: r.many.anomalyAlert(),
      aiToolEntries: r.many.aiToolEntry(),
      ingestionTemplates: r.many.ingestionTemplate(),
      githubInstallations: r.many.githubInstallation(),
      githubPullRequests: r.many.githubPullRequest(),
      githubBranchPrChecks: r.many.githubBranchPullRequestCheck(),
    },
    issuedLicense: {
      organization: r.one.organization({
        from: r.issuedLicense.organizationId,
        to: r.organization.id,
      }),
    },
    connectedBillingAccount: {
      organization: r.one.organization({
        from: r.connectedBillingAccount.organizationId,
        to: r.organization.id,
        optional: false,
      }),
    },
    scimRequestLog: {
      organization: r.one.organization({
        from: r.scimRequestLog.organizationId,
        to: r.organization.id,
        optional: false,
      }),
    },
    scimToken: {
      organization: r.one.organization({
        from: r.scimToken.organizationId,
        to: r.organization.id,
        optional: false,
      }),
    },
    scimExternalId: {
      user: r.one.user({ from: r.scimExternalId.userId, to: r.user.id, optional: false }),
    },
    scimUserResource: {
      user: r.one.user({ from: r.scimUserResource.userId, to: r.user.id, optional: false }),
    },
    project: {
      team: r.one.team({ from: r.project.teamId, to: r.team.id, optional: false }),
      checks: r.many.monitor(),
      costs: r.many.cost(),
      topics: r.many.topic(),
      datasets: r.many.dataset(),
      datasetRecords: r.many.datasetRecord(),
      customGraphs: r.many.customGraph(),
      dashboards: r.many.dashboard(),
      batchEvaluations: r.many.batchEvaluation(),
      triggers: r.many.trigger(),
      experiments: r.many.experiment(),
      annotations: r.many.annotation(),
      projectSecrets: r.many.projectSecret(),
      TriggerSent: r.many.triggerSent(),
      webhookDeliveries: r.many.webhookEndpointDelivery(),
      annotationScores: r.many.annotationScore(),
      shareLinks: r.many.shareLink(),
      workflows: r.many.workflow(),
      WorkflowVersion: r.many.workflowVersion(),
      AnnotationQueue: r.many.annotationQueue(),
      AnnotationQueueItem: r.many.annotationQueueItem(),
      llmPromptConfigs: r.many.llmPromptConfig(),
      analytics: r.many.analytics(),
      notifications: r.many.notification(),
      agents: r.many.agent(),
      evaluators: r.many.evaluator(),
      scenarios: r.many.scenario(),
      suiteConfigurations: r.many.simulationSuite(),
      savedViews: r.many.savedView(),
      gatewayGuardrails: r.many.gatewayGuardrail(),
      gatewayChangeEvents: r.many.gatewayChangeEvent(),
      pinnedTraces: r.many.pinnedTrace(),
      traceEditOverlays: r.many.traceEditOverlay(),
      emailSuppressions: r.many.emailSuppression(),
      ownerUser: r.one.user({ from: r.project.ownerUserId, to: r.user.id }),
    },
    organizationInvite: {
      organization: r.one.organization({
        from: r.organizationInvite.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      requestedByUser: r.one.user({ from: r.organizationInvite.requestedBy, to: r.user.id }),
    },
    monitor: {
      project: r.one.project({ from: r.monitor.projectId, to: r.project.id, optional: false }),
      experiment: r.one.experiment({ from: r.monitor.experimentId, to: r.experiment.id }),
      evaluator: r.one.evaluator({ from: r.monitor.evaluatorId, to: r.evaluator.id }),
    },
    cost: { project: r.one.project({ from: r.cost.projectId, to: r.project.id, optional: false }) },
    topic: {
      project: r.one.project({ from: r.topic.projectId, to: r.project.id, optional: false }),
      parent: r.one.topic({ from: r.topic.parentId, to: r.topic.id, alias: 'Subtopics' }),
      subtopics: r.many.topic({ alias: 'Subtopics' }),
    },
    dataset: {
      project: r.one.project({ from: r.dataset.projectId, to: r.project.id, optional: false }),
      datasetRecords: r.many.datasetRecord(),
      batchEvaluations: r.many.batchEvaluation(),
    },
    datasetRecord: {
      dataset: r.one.dataset({
        from: r.datasetRecord.datasetId,
        to: r.dataset.id,
        optional: false,
      }),
      project: r.one.project({
        from: r.datasetRecord.projectId,
        to: r.project.id,
        optional: false,
      }),
    },
    dashboard: {
      project: r.one.project({ from: r.dashboard.projectId, to: r.project.id, optional: false }),
      graphs: r.many.customGraph(),
    },
    savedView: {
      project: r.one.project({ from: r.savedView.projectId, to: r.project.id, optional: false }),
      user: r.one.user({ from: r.savedView.userId, to: r.user.id }),
    },
    customGraph: {
      project: r.one.project({ from: r.customGraph.projectId, to: r.project.id, optional: false }),
      dashboard: r.one.dashboard({ from: r.customGraph.dashboardId, to: r.dashboard.id }),
      trigger: r.one.trigger({ from: r.customGraph.id, to: r.trigger.customGraphId }),
      TriggerSent: r.many.triggerSent(),
    },
    batchEvaluation: {
      experiment: r.one.experiment({
        from: r.batchEvaluation.experimentId,
        to: r.experiment.id,
        optional: false,
      }),
      project: r.one.project({
        from: r.batchEvaluation.projectId,
        to: r.project.id,
        optional: false,
      }),
      dataset: r.one.dataset({
        from: r.batchEvaluation.datasetId,
        to: r.dataset.id,
        optional: false,
      }),
    },
    trigger: {
      project: r.one.project({ from: r.trigger.projectId, to: r.project.id, optional: false }),
      TriggerSent: r.many.triggerSent(),
      webhookDeliveries: r.many.webhookEndpointDelivery(),
      latestEvaluation: r.one.triggerLatestEvaluation({
        from: r.trigger.id,
        to: r.triggerLatestEvaluation.triggerId,
      }),
      customGraph: r.one.customGraph({ from: r.trigger.customGraphId, to: r.customGraph.id }),
    },
    webhookEndpoint: {
      organization: r.one.organization({
        from: r.webhookEndpoint.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      deliveries: r.many.webhookEndpointDelivery(),
    },
    webhookEndpointDelivery: {
      endpoint: r.one.webhookEndpoint({
        from: r.webhookEndpointDelivery.endpointId,
        to: r.webhookEndpoint.id,
      }),
      project: r.one.project({ from: r.webhookEndpointDelivery.projectId, to: r.project.id }),
      trigger: r.one.trigger({ from: r.webhookEndpointDelivery.triggerId, to: r.trigger.id }),
    },
    experiment: {
      project: r.one.project({ from: r.experiment.projectId, to: r.project.id, optional: false }),
      workflow: r.one.workflow({ from: r.experiment.workflowId, to: r.workflow.id }),
      batchEvaluations: r.many.batchEvaluation(),
      monitor: r.one.monitor({ from: r.experiment.id, to: r.monitor.experimentId }),
    },
    annotation: {
      project: r.one.project({ from: r.annotation.projectId, to: r.project.id, optional: false }),
      user: r.one.user({ from: r.annotation.userId, to: r.user.id }),
    },
    modelProvider: { scopes: r.many.modelProviderScope() },
    modelProviderScope: {
      modelProvider: r.one.modelProvider({
        from: r.modelProviderScope.modelProviderId,
        to: r.modelProvider.id,
        optional: false,
      }),
    },
    modelDefaultConfig: { scopes: r.many.modelDefaultConfigScope() },
    modelDefaultConfigScope: {
      config: r.one.modelDefaultConfig({
        from: r.modelDefaultConfigScope.configId,
        to: r.modelDefaultConfig.id,
        optional: false,
      }),
    },
    githubInstallation: {
      organization: r.one.organization({
        from: r.githubInstallation.organizationId,
        to: r.organization.id,
        optional: false,
      }),
    },
    githubPullRequest: {
      organization: r.one.organization({
        from: r.githubPullRequest.organizationId,
        to: r.organization.id,
        optional: false,
      }),
    },
    githubBranchPullRequestCheck: {
      organization: r.one.organization({
        from: r.githubBranchPullRequestCheck.organizationId,
        to: r.organization.id,
        optional: false,
      }),
    },
    projectSecret: {
      project: r.one.project({
        from: r.projectSecret.projectId,
        to: r.project.id,
        optional: false,
      }),
      createdBy: r.one.user({
        from: r.projectSecret.createdById,
        to: r.user.id,
        optional: false,
        alias: 'secretCreatedBy',
      }),
      updatedBy: r.one.user({
        from: r.projectSecret.updatedById,
        to: r.user.id,
        optional: false,
        alias: 'secretUpdatedBy',
      }),
    },
    triggerSent: {
      customGraph: r.one.customGraph({ from: r.triggerSent.customGraphId, to: r.customGraph.id }),
      project: r.one.project({ from: r.triggerSent.projectId, to: r.project.id, optional: false }),
      trigger: r.one.trigger({ from: r.triggerSent.triggerId, to: r.trigger.id, optional: false }),
    },
    triggerLatestEvaluation: {
      trigger: r.one.trigger({
        from: r.triggerLatestEvaluation.triggerId,
        to: r.trigger.id,
        optional: false,
      }),
    },
    emailSuppression: {
      project: r.one.project({
        from: r.emailSuppression.projectId,
        to: r.project.id,
        optional: false,
      }),
    },
    annotationScore: {
      project: r.one.project({
        from: r.annotationScore.projectId,
        to: r.project.id,
        optional: false,
      }),
      AnnotationQueueScores: r.many.annotationQueueScores(),
    },
    annotationQueue: {
      project: r.one.project({
        from: r.annotationQueue.projectId,
        to: r.project.id,
        optional: false,
      }),
      members: r.many.annotationQueueMembers(),
      AnnotationQueueScores: r.many.annotationQueueScores(),
      AnnotationQueueItems: r.many.annotationQueueItem(),
    },
    annotationQueueMembers: {
      user: r.one.user({ from: r.annotationQueueMembers.userId, to: r.user.id, optional: false }),
      annotationQueue: r.one.annotationQueue({
        from: r.annotationQueueMembers.annotationQueueId,
        to: r.annotationQueue.id,
        optional: false,
      }),
    },
    annotationQueueScores: {
      annotationScore: r.one.annotationScore({
        from: r.annotationQueueScores.annotationScoreId,
        to: r.annotationScore.id,
        optional: false,
      }),
      annotationQueue: r.one.annotationQueue({
        from: r.annotationQueueScores.annotationQueueId,
        to: r.annotationQueue.id,
        optional: false,
      }),
    },
    annotationQueueItem: {
      annotationQueue: r.one.annotationQueue({
        from: r.annotationQueueItem.annotationQueueId,
        to: r.annotationQueue.id,
      }),
      user: r.one.user({ from: r.annotationQueueItem.userId, to: r.user.id, alias: 'AssignedTo' }),
      createdByUser: r.one.user({
        from: r.annotationQueueItem.createdByUserId,
        to: r.user.id,
        alias: 'CreatedBy',
      }),
      project: r.one.project({
        from: r.annotationQueueItem.projectId,
        to: r.project.id,
        optional: false,
      }),
    },
    shareLink: {
      project: r.one.project({ from: r.shareLink.projectId, to: r.project.id, optional: false }),
      user: r.one.user({ from: r.shareLink.userId, to: r.user.id }),
    },
    pinnedTrace: {
      project: r.one.project({ from: r.pinnedTrace.projectId, to: r.project.id, optional: false }),
      user: r.one.user({ from: r.pinnedTrace.userId, to: r.user.id }),
    },
    traceEditOverlay: {
      project: r.one.project({
        from: r.traceEditOverlay.projectId,
        to: r.project.id,
        optional: false,
      }),
      createdBy: r.one.user({
        from: r.traceEditOverlay.createdById,
        to: r.user.id,
        alias: 'TraceEditOverlayCreatedBy',
      }),
      updatedBy: r.one.user({
        from: r.traceEditOverlay.updatedById,
        to: r.user.id,
        alias: 'TraceEditOverlayUpdatedBy',
      }),
    },
    workflow: {
      project: r.one.project({ from: r.workflow.projectId, to: r.project.id, optional: false }),
      latestVersion: r.one.workflowVersion({
        from: r.workflow.latestVersionId,
        to: r.workflowVersion.id,
        alias: 'LatestVersion',
      }),
      currentVersion: r.one.workflowVersion({
        from: r.workflow.currentVersionId,
        to: r.workflowVersion.id,
        alias: 'CurrentVersion',
      }),
      versions: r.many.workflowVersion({ alias: 'Versions' }),
      publishedBy: r.one.user({ from: r.workflow.publishedById, to: r.user.id }),
      experiments: r.many.experiment(),
      copiedFrom: r.one.workflow({
        from: r.workflow.copiedFromWorkflowId,
        to: r.workflow.id,
        alias: 'WorkflowCopies',
      }),
      copiedWorkflows: r.many.workflow({ alias: 'WorkflowCopies' }),
      agents: r.many.agent(),
      evaluators: r.many.evaluator(),
    },
    workflowVersion: {
      author: r.one.user({ from: r.workflowVersion.authorId, to: r.user.id, optional: false }),
      project: r.one.project({
        from: r.workflowVersion.projectId,
        to: r.project.id,
        optional: false,
      }),
      workflow: r.one.workflow({
        from: r.workflowVersion.workflowId,
        to: r.workflow.id,
        optional: false,
        alias: 'Versions',
      }),
      parent: r.one.workflowVersion({
        from: r.workflowVersion.parentId,
        to: r.workflowVersion.id,
        alias: 'WorkflowVersionHistory',
      }),
      children: r.many.workflowVersion({ alias: 'WorkflowVersionHistory' }),
      WorkflowAsLatestVersion: r.many.workflow({ alias: 'LatestVersion' }),
      WorkflowAsCurrentVersion: r.many.workflow({ alias: 'CurrentVersion' }),
    },
    llmPromptConfig: {
      project: r.one.project({
        from: r.llmPromptConfig.projectId,
        to: r.project.id,
        optional: false,
      }),
      copiedFrom: r.one.llmPromptConfig({
        from: r.llmPromptConfig.copiedFromPromptId,
        to: r.llmPromptConfig.id,
        alias: 'PromptCopies',
      }),
      copiedPrompts: r.many.llmPromptConfig({ alias: 'PromptCopies' }),
      versions: r.many.llmPromptConfigVersion(),
      versionTags: r.many.promptTagAssignment(),
    },
    llmPromptConfigVersion: {
      author: r.one.user({ from: r.llmPromptConfigVersion.authorId, to: r.user.id }),
      config: r.one.llmPromptConfig({
        from: r.llmPromptConfigVersion.configId,
        to: r.llmPromptConfig.id,
        optional: false,
      }),
      versionTags: r.many.promptTagAssignment(),
    },
    promptTagAssignment: {
      config: r.one.llmPromptConfig({
        from: r.promptTagAssignment.configId,
        to: r.llmPromptConfig.id,
        optional: false,
      }),
      version: r.one.llmPromptConfigVersion({
        from: r.promptTagAssignment.versionId,
        to: r.llmPromptConfigVersion.id,
        optional: false,
      }),
      promptTag: r.one.promptTag({
        from: r.promptTagAssignment.tagId,
        to: r.promptTag.id,
        optional: false,
      }),
      createdBy: r.one.user({
        from: r.promptTagAssignment.createdById,
        to: r.user.id,
        alias: 'versionTagCreatedBy',
      }),
      updatedBy: r.one.user({
        from: r.promptTagAssignment.updatedById,
        to: r.user.id,
        alias: 'versionTagUpdatedBy',
      }),
    },
    promptTag: {
      organization: r.one.organization({
        from: r.promptTag.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      assignments: r.many.promptTagAssignment(),
      createdBy: r.one.user({
        from: r.promptTag.createdById,
        to: r.user.id,
        alias: 'promptTagCreatedBy',
      }),
      updatedBy: r.one.user({
        from: r.promptTag.updatedById,
        to: r.user.id,
        alias: 'promptTagUpdatedBy',
      }),
    },
    analytics: {
      project: r.one.project({ from: r.analytics.projectId, to: r.project.id, optional: false }),
    },
    customRole: {
      organization: r.one.organization({
        from: r.customRole.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      assignedUsers: r.many.teamUser(),
      roleBindings: r.many.roleBinding(),
    },
    notification: {
      organization: r.one.organization({
        from: r.notification.organizationId,
        to: r.organization.id,
      }),
      project: r.one.project({ from: r.notification.projectId, to: r.project.id }),
    },
    agent: {
      project: r.one.project({ from: r.agent.projectId, to: r.project.id, optional: false }),
      workflow: r.one.workflow({ from: r.agent.workflowId, to: r.workflow.id }),
      copiedFrom: r.one.agent({
        from: r.agent.copiedFromAgentId,
        to: r.agent.id,
        alias: 'AgentCopies',
      }),
      copiedAgents: r.many.agent({ alias: 'AgentCopies' }),
    },
    evaluator: {
      project: r.one.project({ from: r.evaluator.projectId, to: r.project.id, optional: false }),
      workflow: r.one.workflow({ from: r.evaluator.workflowId, to: r.workflow.id }),
      copiedFrom: r.one.evaluator({
        from: r.evaluator.copiedFromEvaluatorId,
        to: r.evaluator.id,
        alias: 'EvaluatorCopies',
      }),
      copiedEvaluators: r.many.evaluator({ alias: 'EvaluatorCopies' }),
      monitors: r.many.monitor(),
      gatewayGuardrails: r.many.gatewayGuardrail(),
    },
    scenario: {
      project: r.one.project({ from: r.scenario.projectId, to: r.project.id, optional: false }),
      testSuite: r.one.simulationSuite({ from: r.scenario.testSuiteId, to: r.simulationSuite.id }),
      versions: r.many.scenarioVersion(),
      lastUpdatedBy: r.one.user({ from: r.scenario.lastUpdatedById, to: r.user.id }),
    },
    scenarioVersion: {
      scenario: r.one.scenario({
        from: r.scenarioVersion.scenarioId,
        to: r.scenario.id,
        optional: false,
      }),
    },
    simulationSuite: {
      project: r.one.project({
        from: r.simulationSuite.projectId,
        to: r.project.id,
        optional: false,
      }),
      testSuiteScenarios: r.many.scenario(),
    },
    subscription: {
      organization: r.one.organization({
        from: r.subscription.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      invoices: r.many.invoice(),
    },
    invoice: {
      subscription: r.one.subscription({
        from: r.invoice.subscriptionId,
        to: r.subscription.id,
        optional: false,
      }),
      lineItems: r.many.invoiceItem(),
    },
    invoiceItem: {
      invoice: r.one.invoice({ from: r.invoiceItem.invoiceId, to: r.invoice.id, optional: false }),
    },
    group: {
      organization: r.one.organization({
        from: r.group.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      members: r.many.groupMembership(),
      roleBindings: r.many.roleBinding(),
    },
    groupMembership: {
      user: r.one.user({ from: r.groupMembership.userId, to: r.user.id, optional: false }),
      group: r.one.group({ from: r.groupMembership.groupId, to: r.group.id, optional: false }),
    },
    roleBinding: {
      organization: r.one.organization({
        from: r.roleBinding.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      user: r.one.user({ from: r.roleBinding.userId, to: r.user.id }),
      group: r.one.group({ from: r.roleBinding.groupId, to: r.group.id }),
      apiKey: r.one.apiKey({ from: r.roleBinding.apiKeyId, to: r.apiKey.id }),
      customRole: r.one.customRole({ from: r.roleBinding.customRoleId, to: r.customRole.id }),
    },
    apiKey: {
      user: r.one.user({ from: r.apiKey.userId, to: r.user.id }),
      organization: r.one.organization({
        from: r.apiKey.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      roleBindings: r.many.roleBinding(),
    },
    virtualKey: {
      principalUser: r.one.user({ from: r.virtualKey.principalUserId, to: r.user.id }),
      scopes: r.many.virtualKeyScope(),
      ledgerEntries: r.many.gatewayBudgetLedger(),
      routingPolicy: r.one.routingPolicy({
        from: r.virtualKey.routingPolicyId,
        to: r.routingPolicy.id,
      }),
    },
    virtualKeyScope: {
      virtualKey: r.one.virtualKey({
        from: r.virtualKeyScope.virtualKeyId,
        to: r.virtualKey.id,
        optional: false,
      }),
    },
    routingPolicy: {
      organization: r.one.organization({
        from: r.routingPolicy.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      createdBy: r.one.user({
        from: r.routingPolicy.createdById,
        to: r.user.id,
        alias: 'RoutingPolicyCreatedBy',
      }),
      updatedBy: r.one.user({
        from: r.routingPolicy.updatedById,
        to: r.user.id,
        alias: 'RoutingPolicyUpdatedBy',
      }),
      virtualKeys: r.many.virtualKey(),
      scopes: r.many.routingPolicyScope(),
    },
    routingPolicyScope: {
      routingPolicy: r.one.routingPolicy({
        from: r.routingPolicyScope.routingPolicyId,
        to: r.routingPolicy.id,
        optional: false,
      }),
    },
    gatewayGuardrail: {
      project: r.one.project({
        from: r.gatewayGuardrail.projectId,
        to: r.project.id,
        optional: false,
      }),
      evaluator: r.one.evaluator({
        from: r.gatewayGuardrail.evaluatorId,
        to: r.evaluator.id,
        optional: false,
      }),
      createdBy: r.one.user({
        from: r.gatewayGuardrail.createdById,
        to: r.user.id,
        alias: 'GatewayGuardrailCreatedBy',
      }),
      updatedBy: r.one.user({
        from: r.gatewayGuardrail.updatedById,
        to: r.user.id,
        alias: 'GatewayGuardrailUpdatedBy',
      }),
    },
    ingestionSource: {
      organization: r.one.organization({
        from: r.ingestionSource.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      team: r.one.team({ from: r.ingestionSource.teamId, to: r.team.id }),
      createdBy: r.one.user({ from: r.ingestionSource.createdById, to: r.user.id }),
    },
    anomalyRule: {
      organization: r.one.organization({
        from: r.anomalyRule.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      createdBy: r.one.user({ from: r.anomalyRule.createdById, to: r.user.id }),
      alerts: r.many.anomalyAlert(),
    },
    anomalyAlert: {
      organization: r.one.organization({
        from: r.anomalyAlert.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      rule: r.one.anomalyRule({
        from: r.anomalyAlert.ruleId,
        to: r.anomalyRule.id,
        optional: false,
      }),
    },
    aiToolEntry: {
      organization: r.one.organization({
        from: r.aiToolEntry.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      teams: r.many.aiToolEntryTeam(),
      departments: r.many.aiToolEntryDepartment(),
    },
    aiToolEntryTeam: {
      entry: r.one.aiToolEntry({
        from: r.aiToolEntryTeam.entryId,
        to: r.aiToolEntry.id,
        optional: false,
      }),
      team: r.one.team({ from: r.aiToolEntryTeam.teamId, to: r.team.id, optional: false }),
    },
    aiToolEntryDepartment: {
      entry: r.one.aiToolEntry({
        from: r.aiToolEntryDepartment.entryId,
        to: r.aiToolEntry.id,
        optional: false,
      }),
    },
    ingestionTemplate: {
      organization: r.one.organization({
        from: r.ingestionTemplate.organizationId,
        to: r.organization.id,
      }),
    },
    gatewayBudget: {
      createdBy: r.one.user({ from: r.gatewayBudget.createdById, to: r.user.id, optional: false }),
      ledgerEntries: r.many.gatewayBudgetLedger(),
      bucketBoundaries: r.many.gatewayBudgetBucketBoundary(),
    },
    gatewayBudgetBucketBoundary: {
      budget: r.one.gatewayBudget({
        from: r.gatewayBudgetBucketBoundary.budgetId,
        to: r.gatewayBudget.id,
        optional: false,
      }),
    },
    gatewayBudgetLedger: {
      budget: r.one.gatewayBudget({
        from: r.gatewayBudgetLedger.budgetId,
        to: r.gatewayBudget.id,
        optional: false,
      }),
      virtualKey: r.one.virtualKey({
        from: r.gatewayBudgetLedger.virtualKeyId,
        to: r.virtualKey.id,
        optional: false,
      }),
    },
    gatewayChangeEvent: {
      organization: r.one.organization({
        from: r.gatewayChangeEvent.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      project: r.one.project({ from: r.gatewayChangeEvent.projectId, to: r.project.id }),
    },
    processManagerOutbox: { attemptLog: r.many.processManagerOutboxAttempt() },
    processManagerOutboxAttempt: {
      outbox: r.one.processManagerOutbox({
        from: r.processManagerOutboxAttempt.outboxId,
        to: r.processManagerOutbox.id,
        optional: false,
      }),
    },
  }),
)
