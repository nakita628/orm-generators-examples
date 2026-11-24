import {
  type PgTableExtraConfigValue,
  bigint,
  boolean,
  customType,
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
} from 'drizzle-orm/pg-core'
import { defineRelations, sql } from 'drizzle-orm'
import cuid from 'cuid'

const bytea = customType<{ data: Uint8Array }>({
  dataType: () => 'bytea',
})

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

export const authenticationMethodEnum = pgEnum('AuthenticationMethod', [
  'GITHUB',
  'MAGIC_LINK',
  'GOOGLE',
  'SSO',
])

export const oauthClientKindEnum = pgEnum('OAuthClientKind', ['PREREGISTERED', 'CIMD', 'DCR'])

export const organizationAccessTokenTypeEnum = pgEnum('OrganizationAccessTokenType', [
  'USER',
  'SYSTEM',
])

export const orgMemberRoleEnum = pgEnum('OrgMemberRole', ['ADMIN', 'MEMBER'])

export const apiKeyIssuerEnum = pgEnum('ApiKeyIssuer', ['USER', 'SYSTEM'])

export const runtimeEnvironmentTypeEnum = pgEnum('RuntimeEnvironmentType', [
  'PRODUCTION',
  'STAGING',
  'DEVELOPMENT',
  'PREVIEW',
])

export const environmentPauseSourceEnum = pgEnum('EnvironmentPauseSource', [
  'BILLING_LIMIT',
  'ADMIN_BAN',
])

export const projectVersionEnum = pgEnum('ProjectVersion', ['V2', 'V3'])

export const environmentKeyWrappingProviderEnum = pgEnum('EnvironmentKeyWrappingProvider', [
  'PLATFORM_KEYRING',
  'AWS_KMS',
])

export const secretStoreProviderEnum = pgEnum('SecretStoreProvider', [
  'DATABASE',
  'AWS_PARAM_STORE',
])

export const oidcSigningKeyStatusEnum = pgEnum('OidcSigningKeyStatus', [
  'STAGED',
  'ACTIVE',
  'RETIRED',
])

/** Each purpose is a separate key set with its own issuer, rotation and publication. */
export const oidcSigningKeyPurposeEnum = pgEnum('OidcSigningKeyPurpose', [
  'EXTERNAL_OIDC',
  'INTERNAL_RUN',
])

export const taskTriggerSourceEnum = pgEnum('TaskTriggerSource', [
  'STANDARD',
  'SCHEDULED',
  'AGENT',
  'WEBHOOK',
])

export const taskExecutionModeEnum = pgEnum('TaskExecutionMode', ['STANDARD', 'SANDBOX'])

export const webhookEndpointStatusEnum = pgEnum('WebhookEndpointStatus', [
  'ACTIVE',
  'INACTIVE',
  'DELETING',
])

export const webhookDeliveryStatusEnum = pgEnum('WebhookDeliveryStatus', [
  'PENDING',
  'PROCESSING',
  'SUCCEEDED',
  'FAILED',
  'FILTERED',
  'UNMATCHED',
])

export const taskRunStatusEnum = pgEnum('TaskRunStatus', [
  'DELAYED',
  'PENDING',
  'PENDING_VERSION',
  'WAITING_FOR_DEPLOY',
  'DEQUEUED',
  'EXECUTING',
  'WAITING_TO_RESUME',
  'RETRYING_AFTER_FAILURE',
  'PAUSED',
  'CANCELED',
  'INTERRUPTED',
  'COMPLETED_SUCCESSFULLY',
  'COMPLETED_WITH_ERRORS',
  'SYSTEM_FAILURE',
  'CRASHED',
  'EXPIRED',
  'TIMED_OUT',
])

export const runEngineVersionEnum = pgEnum('RunEngineVersion', ['V1', 'V2'])

export const taskRunExecutionStatusEnum = pgEnum('TaskRunExecutionStatus', [
  'RUN_CREATED',
  'DELAYED',
  'QUEUED',
  'QUEUED_EXECUTING',
  'PENDING_EXECUTING',
  'EXECUTING',
  'EXECUTING_WITH_WAITPOINTS',
  'SUSPENDED',
  'PENDING_CANCEL',
  'FINISHED',
])

export const taskRunCheckpointTypeEnum = pgEnum('TaskRunCheckpointType', [
  'DOCKER',
  'KUBERNETES',
  'COMPUTE',
])

export const waitpointTypeEnum = pgEnum('WaitpointType', ['RUN', 'DATETIME', 'MANUAL', 'BATCH'])

export const waitpointStatusEnum = pgEnum('WaitpointStatus', ['PENDING', 'COMPLETED'])

export const workerInstanceGroupTypeEnum = pgEnum('WorkerInstanceGroupType', [
  'MANAGED',
  'UNMANAGED',
])

export const workloadTypeEnum = pgEnum('WorkloadType', ['CONTAINER', 'MICROVM'])

export const taskRunAttemptStatusEnum = pgEnum('TaskRunAttemptStatus', [
  'PENDING',
  'EXECUTING',
  'PAUSED',
  'FAILED',
  'CANCELED',
  'COMPLETED',
])

export const taskEventLevelEnum = pgEnum('TaskEventLevel', [
  'TRACE',
  'DEBUG',
  'LOG',
  'INFO',
  'WARN',
  'ERROR',
])

export const taskEventKindEnum = pgEnum('TaskEventKind', [
  'UNSPECIFIED',
  'INTERNAL',
  'SERVER',
  'CLIENT',
  'PRODUCER',
  'CONSUMER',
  'UNRECOGNIZED',
  'LOG',
])

export const taskEventStatusEnum = pgEnum('TaskEventStatus', [
  'UNSET',
  'OK',
  'ERROR',
  'UNRECOGNIZED',
])

export const taskQueueTypeEnum = pgEnum('TaskQueueType', ['VIRTUAL', 'NAMED'])

export const taskQueueVersionEnum = pgEnum('TaskQueueVersion', ['V1', 'V2'])

export const taskQueueConcurrencyVersionEnum = pgEnum('TaskQueueConcurrencyVersion', ['V1', 'V2'])

export const taskQueueRoleEnum = pgEnum('TaskQueueRole', ['QUEUE', 'LIMIT'])

export const batchTaskRunStatusEnum = pgEnum('BatchTaskRunStatus', [
  'PENDING',
  'PROCESSING',
  'COMPLETED',
  'PARTIAL_FAILED',
  'ABORTED',
])

export const batchTaskRunItemStatusEnum = pgEnum('BatchTaskRunItemStatus', [
  'PENDING',
  'FAILED',
  'CANCELED',
  'COMPLETED',
])

export const environmentVariableIgnoreReasonEnum = pgEnum('EnvironmentVariableIgnoreReason', [
  'EMPTY',
  'MISSING',
  'OIDC',
])

export const checkpointTypeEnum = pgEnum('CheckpointType', ['DOCKER', 'KUBERNETES'])

export const checkpointRestoreEventTypeEnum = pgEnum('CheckpointRestoreEventType', [
  'CHECKPOINT',
  'RESTORE',
])

export const workerDeploymentTypeEnum = pgEnum('WorkerDeploymentType', [
  'MANAGED',
  'UNMANAGED',
  'V1',
])

export const deploymentInitiatedByCredentialEnum = pgEnum('DeploymentInitiatedByCredential', [
  'ROOT_API_KEY',
  'SCOPED_API_KEY',
  'PERSONAL_ACCESS_TOKEN',
  'ORGANIZATION_ACCESS_TOKEN',
  'TRUSTED_PUBLISHING',
])

export const workerDeploymentStatusEnum = pgEnum('WorkerDeploymentStatus', [
  'PENDING',
  'INSTALLING',
  'BUILDING',
  'DEPLOYING',
  'DEPLOYED',
  'FAILED',
  'CANCELED',
  'TIMED_OUT',
])

export const scheduleTypeEnum = pgEnum('ScheduleType', ['DECLARATIVE', 'IMPERATIVE'])

export const scheduleGeneratorTypeEnum = pgEnum('ScheduleGeneratorType', ['CRON'])

export const projectAlertChannelTypeEnum = pgEnum('ProjectAlertChannelType', [
  'EMAIL',
  'SLACK',
  'WEBHOOK',
])

export const projectAlertTypeEnum = pgEnum('ProjectAlertType', [
  'TASK_RUN',
  'TASK_RUN_ATTEMPT',
  'DEPLOYMENT_FAILURE',
  'DEPLOYMENT_SUCCESS',
  'ERROR_GROUP',
  'DASHBOARD_AGENT_WATCH',
])

export const projectAlertStatusEnum = pgEnum('ProjectAlertStatus', ['PENDING', 'SENT', 'FAILED'])

export const integrationServiceEnum = pgEnum('IntegrationService', ['SLACK', 'VERCEL'])

export const bulkActionTypeEnum = pgEnum('BulkActionType', ['CANCEL', 'REPLAY'])

export const bulkActionStatusEnum = pgEnum('BulkActionStatus', ['PENDING', 'COMPLETED', 'ABORTED'])

export const bulkActionNotificationTypeEnum = pgEnum('BulkActionNotificationType', [
  'NONE',
  'EMAIL',
])

export const bulkActionItemStatusEnum = pgEnum('BulkActionItemStatus', [
  'PENDING',
  'COMPLETED',
  'FAILED',
])

export const githubRepositorySelectionEnum = pgEnum('GithubRepositorySelection', [
  'ALL',
  'SELECTED',
])

export const trustedDeployScopeEnum = pgEnum('TrustedDeployScope', [
  'ANY',
  'BRANCH',
  'TAG',
  'ENVIRONMENT',
])

export const impersonationAuditLogActionEnum = pgEnum('ImpersonationAuditLogAction', [
  'START',
  'STOP',
])

export const supportAccessModeEnum = pgEnum('SupportAccessMode', ['ALLOW', 'REQUIRES_REQUEST'])

export const supportAccessRequestStatusEnum = pgEnum('SupportAccessRequestStatus', [
  'PENDING',
  'APPROVED',
  'CANCELLED',
])

export const customerQuerySourceEnum = pgEnum('CustomerQuerySource', ['DASHBOARD', 'API'])

export const customerQueryScopeEnum = pgEnum('CustomerQueryScope', [
  'ORGANIZATION',
  'PROJECT',
  'ENVIRONMENT',
])

export const platformNotificationSurfaceEnum = pgEnum('PlatformNotificationSurface', [
  'WEBAPP',
  'CLI',
])

export const platformNotificationScopeEnum = pgEnum('PlatformNotificationScope', [
  'USER',
  'PROJECT',
  'ORGANIZATION',
  'GLOBAL',
])

export const errorGroupStatusEnum = pgEnum('ErrorGroupStatus', [
  'UNRESOLVED',
  'RESOLVED',
  'IGNORED',
])

export const dataStoreKindEnum = pgEnum('DataStoreKind', ['CLICKHOUSE'])

export const user = pgTable(
  'User',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    email: text('email').notNull(),
    /** Address awaiting confirmation; becomes `email` only after a link sent to it is opened */
    pendingEmail: text('pendingEmail'),
    authenticationMethod: authenticationMethodEnum('authenticationMethod').notNull(),
    authenticationProfile: jsonb('authenticationProfile'),
    authenticationExtraParams: jsonb('authenticationExtraParams'),
    authIdentifier: text('authIdentifier'),
    displayName: text('displayName'),
    name: text('name'),
    avatarUrl: text('avatarUrl'),
    admin: boolean('admin').notNull().default(false),
    /** Preferences for the dashboard */
    dashboardPreferences: jsonb('dashboardPreferences'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    /** @deprecated */
    isOnCloudWaitlist: boolean('isOnCloudWaitlist').notNull().default(false),
    /** @deprecated */
    featureCloud: boolean('featureCloud').notNull().default(false),
    /** @deprecated */
    isOnHostedRepoWaitlist: boolean('isOnHostedRepoWaitlist').notNull().default(false),
    marketingEmails: boolean('marketingEmails').notNull().default(true),
    confirmedBasicDetails: boolean('confirmedBasicDetails').notNull().default(false),
    referralSource: text('referralSource'),
    onboardingData: jsonb('onboardingData'),
    mfaEnabledAt: timestamp('mfaEnabledAt', { precision: 3 }),
    mfaSecretReferenceId: text('mfaSecretReferenceId'),
    /** Hash of the last used code to prevent replay attacks */
    mfaLastUsedCode: text('mfaLastUsedCode'),
    /**
     * Maximum session lifetime in seconds for this user. Default = 1 year (Gregorian).
     * May be further restricted by Organization.maxSessionDuration on any of the user's orgs.
     */
    sessionDuration: integer('sessionDuration').notNull().default(31556952),
    /**
     * Absolute timestamp at which the user's current session must end. Written
     * at login (and any time the effective duration is recomputed) and shortened
     * when an admin lowers an org cap. Read on requireUser / getUser; null is
     * treated as "no enforced deadline" (legacy sessions from before this
     * shipped). Read-on-demand has zero per-request DB cost — the value
     * piggybacks on the User fetch that requireUser already does.
     */
    nextSessionEnd: timestamp('nextSessionEnd', { precision: 3 }),
    invitationCodeId: text('invitationCodeId'),
  },
  (table) => [
    uniqueIndex('User_email_key').on(table.email),
    uniqueIndex('User_authIdentifier_key').on(table.authIdentifier),
    foreignKey({
      name: 'User_mfaSecretReferenceId_fkey',
      columns: [table.mfaSecretReferenceId],
      foreignColumns: [secretReference.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'User_invitationCodeId_fkey',
      columns: [table.invitationCodeId],
      foreignColumns: [invitationCode.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const mfaBackupCode = pgTable(
  'MfaBackupCode',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    /** Hash of the actual code */
    code: text('code').notNull(),
    userId: text('userId').notNull(),
    usedAt: timestamp('usedAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('MfaBackupCode_userId_code_key').on(table.userId, table.code),
    foreignKey({
      name: 'MfaBackupCode_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

export const invitationCode = pgTable(
  'InvitationCode',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    code: text('code').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [uniqueIndex('InvitationCode_code_key').on(table.code)],
)

/** Used to generate PersonalAccessTokens, they're one-time use */
export const authorizationCode = pgTable(
  'AuthorizationCode',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    code: text('code').notNull(),
    personalAccessTokenId: text('personalAccessTokenId'),
    /** Client IP of the request that minted the code, compared with the approving browser's IP. Deleted with the code once it expires. */
    requesterIp: text('requesterIp'),
    /** S256 PKCE challenge from the client that minted the code. Null for codes minted by CLIs that predate it. */
    codeChallenge: text('codeChallenge'),
    /** Loopback port the approving browser is redirected to with the approval code. Null means the user pastes it. */
    callbackPort: integer('callbackPort'),
    /** sha256 of the approval code, which only the approving browser ever sees. */
    approvalCodeHash: text('approvalCodeHash'),
    oauthClientId: text('oauthClientId'),
    redirectUri: text('redirectUri'),
    scope: text('scope').array().default([]),
    resource: text('resource'),
    approvedByUserId: text('approvedByUserId'),
    /** The user who started the OAuth authorization. Only they can approve it. */
    requestedByUserId: text('requestedByUserId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('AuthorizationCode_code_key').on(table.code),
    index('AuthorizationCode_oauthClientId_idx').on(table.oauthClientId),
    foreignKey({
      name: 'AuthorizationCode_personalAccessTokenId_fkey',
      columns: [table.personalAccessTokenId],
      foreignColumns: [personalAccessToken.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AuthorizationCode_oauthClientId_fkey',
      columns: [table.oauthClientId],
      foreignColumns: [oAuthClient.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const personalAccessToken = pgTable(
  'PersonalAccessToken',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    /** If generated by the CLI this will be "cli", otherwise user-provided */
    name: text('name').notNull(),
    /** This is the token encrypted using the ENCRYPTION_KEY */
    encryptedToken: jsonb('encryptedToken').notNull(),
    /** This is shown in the UI, with ******** */
    obfuscatedToken: text('obfuscatedToken').notNull(),
    /** This is used to find the token in the database */
    hashedToken: text('hashedToken').notNull(),
    userId: text('userId').notNull(),
    revokedAt: timestamp('revokedAt', { precision: 3 }),
    lastAccessedAt: timestamp('lastAccessedAt', { precision: 3 }),
    /** After this the token no longer authenticates. Null means it never expires. */
    expiresAt: timestamp('expiresAt', { precision: 3 }),
    /**
     * Set on CLI login tokens: they stop authenticating this many seconds after their last use
     * (or creation, if never used). Null means no idle limit.
     */
    idleTimeoutSeconds: integer('idleTimeoutSeconds'),
    /** When the "about to expire" email was sent. Set before sending, so it's sent at most once. */
    expiryReminderSentAt: timestamp('expiryReminderSentAt', { precision: 3 }),
    /** When the "has expired" email was sent. Set before sending, so it's sent at most once. */
    expiredNoticeSentAt: timestamp('expiredNoticeSentAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    oauthClientId: text('oauthClientId'),
  },
  (table) => [
    uniqueIndex('PersonalAccessToken_hashedToken_key').on(table.hashedToken),
    index('PersonalAccessToken_userId_idx').on(table.userId),
    index('PersonalAccessToken_oauthClientId_userId_idx').on(table.oauthClientId, table.userId),
    foreignKey({
      name: 'PersonalAccessToken_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'PersonalAccessToken_oauthClientId_fkey',
      columns: [table.oauthClientId],
      foreignColumns: [oAuthClient.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

export const oAuthClient = pgTable(
  'OAuthClient',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    kind: oauthClientKindEnum('kind').notNull(),
    clientId: text('clientId').notNull(),
    metadataUrl: text('metadataUrl'),
    displayName: text('displayName'),
    selfReportedName: text('selfReportedName'),
    redirectUris: text('redirectUris').array(),
    dedupeKey: text('dedupeKey'),
    metadataExpiresAt: timestamp('metadataExpiresAt', { precision: 3 }),
    lastUsedAt: timestamp('lastUsedAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('OAuthClient_clientId_key').on(table.clientId),
    uniqueIndex('OAuthClient_dedupeKey_key').on(table.dedupeKey),
    index('OAuthClient_kind_lastUsedAt_idx').on(table.kind, table.lastUsedAt),
  ],
)

export const oAuthRefreshToken = pgTable(
  'OAuthRefreshToken',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    hashedToken: text('hashedToken').notNull(),
    personalAccessTokenId: text('personalAccessTokenId').notNull(),
    oauthClientId: text('oauthClientId').notNull(),
    scope: text('scope').array(),
    resource: text('resource').notNull(),
    usedAt: timestamp('usedAt', { precision: 3 }),
    replacedById: text('replacedById'),
    idleExpiresAt: timestamp('idleExpiresAt', { precision: 3 }).notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    uniqueIndex('OAuthRefreshToken_hashedToken_key').on(table.hashedToken),
    index('OAuthRefreshToken_personalAccessTokenId_idx').on(table.personalAccessTokenId),
    index('OAuthRefreshToken_oauthClientId_idx').on(table.oauthClientId),
    index('OAuthRefreshToken_idleExpiresAt_idx').on(table.idleExpiresAt),
    foreignKey({
      name: 'OAuthRefreshToken_personalAccessTokenId_fkey',
      columns: [table.personalAccessTokenId],
      foreignColumns: [personalAccessToken.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'OAuthRefreshToken_oauthClientId_fkey',
      columns: [table.oauthClientId],
      foreignColumns: [oAuthClient.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const organizationAccessToken = pgTable(
  'OrganizationAccessToken',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    /** User-provided name for the token */
    name: text('name').notNull(),
    /** Used to differentiate between user-generated and system-generated tokens */
    type: organizationAccessTokenTypeEnum('type').notNull().default('USER'),
    /** This is used to find the token in the database */
    hashedToken: text('hashedToken').notNull(),
    organizationId: text('organizationId').notNull(),
    /** Optional expiration date for the token */
    expiresAt: timestamp('expiresAt', { precision: 3 }),
    revokedAt: timestamp('revokedAt', { precision: 3 }),
    lastAccessedAt: timestamp('lastAccessedAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('OrganizationAccessToken_hashedToken_key').on(table.hashedToken),
    index('OrganizationAccessToken_organizationId_type_createdAt_idx').on(
      table.organizationId,
      table.type,
      table.createdAt,
    ),
    foreignKey({
      name: 'OrganizationAccessToken_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

export const organization = pgTable(
  'Organization',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    slug: text('slug').notNull(),
    title: text('title').notNull(),
    maximumExecutionTimePerRunInMs: integer('maximumExecutionTimePerRunInMs')
      .notNull()
      .default(900000),
    maximumConcurrencyLimit: integer('maximumConcurrencyLimit').notNull().default(10),
    /** This is deprecated and will be removed in the future */
    maximumSchedulesLimit: integer('maximumSchedulesLimit').notNull().default(5),
    maximumDevQueueSize: integer('maximumDevQueueSize'),
    maximumDeployedQueueSize: integer('maximumDeployedQueueSize'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    deletedAt: timestamp('deletedAt', { precision: 3 }),
    companySize: text('companySize'),
    onboardingData: jsonb('onboardingData'),
    avatar: jsonb('avatar'),
    runsEnabled: boolean('runsEnabled').notNull().default(true),
    isActivated: boolean('v3Enabled').notNull().default(false),
    /** @deprecated */
    v2Enabled: boolean('v2Enabled').notNull().default(false),
    /** @deprecated */
    v2MarqsEnabled: boolean('v2MarqsEnabled').notNull().default(false),
    /** @deprecated */
    hasRequestedV3: boolean('hasRequestedV3').notNull().default(false),
    apiRateLimiterConfig: jsonb('apiRateLimiterConfig'),
    realtimeRateLimiterConfig: jsonb('realtimeRateLimiterConfig'),
    batchRateLimitConfig: jsonb('batchRateLimitConfig'),
    batchQueueConcurrencyConfig: jsonb('batchQueueConcurrencyConfig'),
    /** Webhook limits from the org's plan or the back office: { maxWaitersPerEnvironment, maxWaitersPerEndpoint, concurrency }. Null uses the env defaults. */
    webhookLimitsConfig: jsonb('webhookLimitsConfig'),
    featureFlags: jsonb('featureFlags'),
    maximumProjectCount: integer('maximumProjectCount').notNull().default(25),
    /**
     * Maximum session lifetime in seconds for members of this organization.
     * When set, caps each member's User.sessionDuration. The most restrictive cap across
     * all of a user's orgs (excluding nulls) wins.
     */
    maxSessionDuration: integer('maxSessionDuration'),
    /**
     * S2 basin that holds this org's realtime streams. Null until the
     * per-org basin has been provisioned (OSS / s2-lite installs leave
     * it null forever; reads fall back to the global basin env var).
     * Set once at provisioning time; retention is reconfigured in-place
     * when the org's plan changes.
     */
    streamBasinName: text('streamBasinName'),
    supportAccessMode: supportAccessModeEnum('supportAccessMode').notNull().default('ALLOW'),
    supportAccessModeUpdatedAt: timestamp('supportAccessModeUpdatedAt', { precision: 3 }),
    supportAccessModeUpdatedById: text('supportAccessModeUpdatedById'),
    /** When true, members without MFA are locked out of the dashboard and CLI until they set it up. */
    mfaRequired: boolean('mfaRequired').notNull().default(false),
    /** Dashboard actions that need an MFA check within the step-up window (see MfaStepUpAction in the webapp). */
    mfaStepUpActions: text('mfaStepUpActions').array().default([]),
    mfaStepUpWindowMinutes: integer('mfaStepUpWindowMinutes').notNull().default(5),
  },
  (table) => [uniqueIndex('Organization_slug_key').on(table.slug)],
)

export const orgMember = pgTable(
  'OrgMember',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    organizationId: text('organizationId').notNull(),
    userId: text('userId').notNull(),
    role: orgMemberRoleEnum('role').notNull().default('MEMBER'),
    /** Who invited them (a User id, kept if that user is deleted). Null if not invited or joined before this was recorded. */
    invitedById: text('invitedById'),
    /** How they joined: created, invite, sso_jit, directory_sync or manual. Null if joined before this was recorded. */
    joinMethod: text('joinMethod'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('OrgMember_userId_idx').on(table.userId),
    uniqueIndex('OrgMember_organizationId_userId_key').on(table.organizationId, table.userId),
    foreignKey({
      name: 'OrgMember_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'OrgMember_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const orgMemberInvite = pgTable(
  'OrgMemberInvite',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    token: text('token')
      .notNull()
      .$defaultFn(() => cuid()),
    email: text('email').notNull(),
    role: orgMemberRoleEnum('role').notNull().default('MEMBER'),
    /**
     * Optional RBAC role to assign on invite acceptance. When set, the
     * accept-invite flow calls the loaded RBAC plugin's setUserRole with
     * this id after creating the OrgMember. Null = legacy behaviour, the
     * runtime fallback derives the role from `role` above.
     *
     * Plain text (not an FK) — the RBAC plugin's RbacRole table lives on
     * a separate schema (Drizzle, not Prisma) so we can't model the FK
     * here. Validation happens at write time (action) and read time
     * (acceptInvite).
     */
    rbacRoleId: text('rbacRoleId'),
    organizationId: text('organizationId').notNull(),
    inviterId: text('inviterId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('OrgMemberInvite_token_key').on(table.token),
    uniqueIndex('OrgMemberInvite_organizationId_email_key').on(table.organizationId, table.email),
    foreignKey({
      name: 'OrgMemberInvite_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'OrgMemberInvite_inviterId_fkey',
      columns: [table.inviterId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const runtimeEnvironment = pgTable(
  'RuntimeEnvironment',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    slug: text('slug').notNull(),
    apiKey: text('apiKey').notNull(),
    /** Non-null prevents the root API key from being returned to users. */
    rootApiKeyHiddenAt: timestamp('rootApiKeyHiddenAt', { precision: 3 }),
    /** Non-null disables automatic root key injection into new deployed attempts. */
    rootKeyInjectionDisabledAt: timestamp('rootKeyInjectionDisabledAt', { precision: 3 }),
    /** @deprecated was for v2 */
    pkApiKey: text('pkApiKey').notNull(),
    type: runtimeEnvironmentTypeEnum('type').notNull().default('DEVELOPMENT'),
    /**
     * If true, this environment has branches and is treated differently in the dashboard/API
     * NB: this flag is NOT used for Development branches, instead (type, parentEnvironmentId) = (DEVELOPMENT, NULL) is used
     */
    isBranchableEnvironment: boolean('isBranchableEnvironment').notNull().default(false),
    branchName: text('branchName'),
    parentEnvironmentId: text('parentEnvironmentId'),
    /** Only configured on the PREVIEW root. Null disables automatic archiving. */
    previewAutoArchiveAfterDays: integer('previewAutoArchiveAfterDays'),
    previewAutoArchiveExcludedBranches: text('previewAutoArchiveExcludedBranches')
      .array()
      .default([]),
    /** Durable per-project cleanup progress. Null next-check disables scheduling. */
    previewAutoArchiveNextCheckAt: timestamp('previewAutoArchiveNextCheckAt', { precision: 3 }),
    previewAutoArchiveCursorCreatedAt: timestamp('previewAutoArchiveCursorCreatedAt', {
      precision: 3,
    }),
    previewAutoArchiveCursorId: text('previewAutoArchiveCursorId'),
    git: jsonb('git'),
    /** When set API calls will fail */
    archivedAt: timestamp('archivedAt', { precision: 3 }),
    /** A memorable code for the environment */
    shortcode: text('shortcode').notNull(),
    maximumConcurrencyLimit: integer('maximumConcurrencyLimit').notNull().default(5),
    concurrencyLimitBurstFactor: numeric('concurrencyLimitBurstFactor', { precision: 4, scale: 2 })
      .notNull()
      .default('2'),
    paused: boolean('paused').notNull().default(false),
    pauseSource: environmentPauseSourceEnum('pauseSource'),
    autoEnableInternalSources: boolean('autoEnableInternalSources').notNull().default(true),
    organizationId: text('organizationId').notNull(),
    projectId: text('projectId').notNull(),
    orgMemberId: text('orgMemberId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    /** Allows us to customize the built-in environment variables for a specific environment, like TRIGGER_OTEL_SPAN_ATTRIBUTE_VALUE_LENGTH_LIMIT */
    builtInEnvironmentVariableOverrides: jsonb('builtInEnvironmentVariableOverrides'),
    tunnelId: text('tunnelId'),
    currentSessionId: text('currentSessionId'),
  },
  (table): PgTableExtraConfigValue[] => [
    uniqueIndex('RuntimeEnvironment_apiKey_key').on(table.apiKey),
    uniqueIndex('RuntimeEnvironment_pkApiKey_key').on(table.pkApiKey),
    index('RuntimeEnvironment_parentEnvironmentId_idx').on(table.parentEnvironmentId),
    index('RuntimeEnvironment_projectId_idx').on(table.projectId),
    index('RuntimeEnvironment_organizationId_idx').on(table.organizationId),
    uniqueIndex('RuntimeEnvironment_projectId_slug_orgMemberId_key').on(
      table.projectId,
      table.slug,
      table.orgMemberId,
    ),
    uniqueIndex('RuntimeEnvironment_projectId_shortcode_key').on(table.projectId, table.shortcode),
    foreignKey({
      name: 'RuntimeEnvironment_parentEnvironmentId_fkey',
      columns: [table.parentEnvironmentId],
      foreignColumns: [table.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'RuntimeEnvironment_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'RuntimeEnvironment_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'RuntimeEnvironment_orgMemberId_fkey',
      columns: [table.orgMemberId],
      foreignColumns: [orgMember.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'RuntimeEnvironment_currentSessionId_fkey',
      columns: [table.currentSessionId],
      foreignColumns: [runtimeEnvironmentSession.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const apiKeys = pgTable(
  'api_keys',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    name: text('name').notNull(),
    keyHash: text('key_hash').notNull(),
    lastFour: text('last_four').notNull(),
    /** Optional one-line note the creator adds to say what the key is for. */
    description: text('description'),
    runtimeEnvironmentId: text('runtime_environment_id').notNull(),
    createdByUserId: text('created_by_user_id'),
    issuedBy: apiKeyIssuerEnum('issued_by').notNull().default('USER'),
    presetId: text('preset_id'),
    scopes: text('scopes').array(),
    /** CIDR ranges this key may be used from. Empty means any IP. */
    allowedIpRanges: text('allowed_ip_ranges').array().default([]),
    lastUsedAt: timestamp('last_used_at', { precision: 3 }),
    lastUsedIp: text('last_used_ip'),
    revokedAt: timestamp('revoked_at', { precision: 3 }),
    expiresAt: timestamp('expires_at', { precision: 3 }),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    uniqueIndex('api_keys_key_hash_key').on(table.keyHash),
    index('api_keys_runtime_environment_id_revoked_at_created_at_idx').on(
      table.runtimeEnvironmentId,
      table.revokedAt,
      table.createdAt,
    ),
    index('api_keys_env_issued_by_revoked_at_created_at_idx').on(
      table.runtimeEnvironmentId,
      table.issuedBy,
      table.revokedAt,
      table.createdAt,
    ),
    foreignKey({
      name: 'api_keys_runtime_environment_id_fkey',
      columns: [table.runtimeEnvironmentId],
      foreignColumns: [runtimeEnvironment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'api_keys_created_by_user_id_fkey',
      columns: [table.createdByUserId],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

/**
 * Records of previously-valid API keys that are still accepted for authentication
 * during a grace window after rotation. Extend or end the grace period by updating `expiresAt`.
 */
export const revokedApiKey = pgTable(
  'RevokedApiKey',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    apiKey: text('apiKey').notNull(),
    runtimeEnvironmentId: text('runtimeEnvironmentId').notNull(),
    expiresAt: timestamp('expiresAt', { precision: 3 }).notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('RevokedApiKey_apiKey_idx').on(table.apiKey),
    index('RevokedApiKey_runtimeEnvironmentId_idx').on(table.runtimeEnvironmentId),
    foreignKey({
      name: 'RevokedApiKey_runtimeEnvironmentId_fkey',
      columns: [table.runtimeEnvironmentId],
      foreignColumns: [runtimeEnvironment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const project = pgTable(
  'Project',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    slug: text('slug').notNull(),
    name: text('name').notNull(),
    externalRef: text('externalRef').notNull(),
    organizationId: text('organizationId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    deletedAt: timestamp('deletedAt', { precision: 3 }),
    /** Set the first time the CLI `init` command completes against this project. Drives the dev onboarding progress. */
    initializedAt: timestamp('initializedAt', { precision: 3 }),
    /** Runtime used when a deployment config does not specify one. Null preserves the legacy Node 20 fallback. */
    defaultRuntime: text('defaultRuntime'),
    version: projectVersionEnum('version').notNull().default('V2'),
    engine: runEngineVersionEnum('engine').notNull().default('V1'),
    builderProjectId: text('builderProjectId'),
    defaultWorkerGroupId: text('defaultWorkerGroupId'),
    /** Newly triggered deployed runs require the secret proxy (pinned in their annotations). */
    secretProxyEnabled: boolean('secretProxyEnabled').notNull().default(false),
    /** The master queues they are allowed to use (impacts what they can set as default and trigger runs with) */
    allowedWorkerQueues: text('allowedMasterQueues').array().default([]),
    buildSettings: jsonb('buildSettings'),
    onboardingData: jsonb('onboardingData'),
  },
  (table): PgTableExtraConfigValue[] => [
    uniqueIndex('Project_slug_key').on(table.slug),
    uniqueIndex('Project_externalRef_key').on(table.externalRef),
    foreignKey({
      name: 'Project_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Project_defaultWorkerGroupId_fkey',
      columns: [table.defaultWorkerGroupId],
      foreignColumns: [workerInstanceGroup.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const environmentEncryptionPolicy = pgTable(
  'EnvironmentEncryptionPolicy',
  {
    environmentId: text('environmentId').primaryKey(),
    revision: integer('revision').notNull().default(0),
    wrappingProvider: environmentKeyWrappingProviderEnum('wrappingProvider'),
    wrappingKeyReference: text('wrappingKeyReference'),
    activeKeyId: text('activeKeyId'),
  },
  (table) => [
    foreignKey({
      name: 'EnvironmentEncryptionPolicy_environmentId_fkey',
      columns: [table.environmentId],
      foreignColumns: [runtimeEnvironment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'EnvironmentEncryptionPolicy_activeKeyId_environmentId_fkey',
      columns: [table.activeKeyId, table.environmentId],
      foreignColumns: [environmentEncryptionKey.id, environmentEncryptionKey.environmentId],
    })
      .onDelete('no action')
      .onUpdate('no action'),
  ],
)

export const environmentEncryptionKey = pgTable(
  'EnvironmentEncryptionKey',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    environmentId: text('environmentId').notNull(),
    wrappedKey: jsonb('wrappedKey').notNull(),
    wrappingProvider: environmentKeyWrappingProviderEnum('wrappingProvider').notNull(),
    wrappingKeyReference: text('wrappingKeyReference').notNull(),
    revision: integer('revision').notNull().default(0),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('EnvironmentEncryptionKey_environmentId_idx').on(table.environmentId),
    unique('EnvironmentEncryptionKey_id_environmentId_key').on(table.id, table.environmentId),
    foreignKey({
      name: 'EnvironmentEncryptionKey_environmentId_fkey',
      columns: [table.environmentId],
      foreignColumns: [runtimeEnvironment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const secretReference = pgTable(
  'SecretReference',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    key: text('key').notNull(),
    provider: secretStoreProviderEnum('provider').notNull().default('DATABASE'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [uniqueIndex('SecretReference_key_key').on(table.key)],
)

export const secretStore = pgTable(
  'SecretStore',
  {
    key: text('key').notNull(),
    value: jsonb('value').notNull(),
    version: text('version').notNull().default('1'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('SecretStore_key_key').on(table.key),
    index('SecretStore_key_idx').on(table.key),
  ],
)

export const oidcSigningKey = pgTable(
  'OidcSigningKey',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    kid: text('kid').notNull(),
    purpose: oidcSigningKeyPurposeEnum('purpose').notNull().default('EXTERNAL_OIDC'),
    /** At most one STAGED and one ACTIVE key per purpose: enforced by a partial unique index in the migration. */
    status: oidcSigningKeyStatusEnum('status').notNull(),
    publicJwk: jsonb('publicJwk').notNull(),
    readyAt: timestamp('readyAt', { precision: 3 }).notNull(),
    publicUntil: timestamp('publicUntil', { precision: 3 }),
    generatedByUserId: text('generatedByUserId').notNull(),
    activatedAt: timestamp('activatedAt', { precision: 3 }),
    activatedByUserId: text('activatedByUserId'),
    retiredAt: timestamp('retiredAt', { precision: 3 }),
    retiredByUserId: text('retiredByUserId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [uniqueIndex('OidcSigningKey_kid_key').on(table.kid)],
)

export const dataMigration = pgTable(
  'DataMigration',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    name: text('name').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    completedAt: timestamp('completedAt', { precision: 3 }),
  },
  (table) => [uniqueIndex('DataMigration_name_key').on(table.name)],
)

export const backgroundWorker = pgTable(
  'BackgroundWorker',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    friendlyId: text('friendlyId').notNull(),
    engine: runEngineVersionEnum('engine').notNull().default('V1'),
    contentHash: text('contentHash').notNull(),
    sdkVersion: text('sdkVersion').notNull().default('unknown'),
    cliVersion: text('cliVersion').notNull().default('unknown'),
    projectId: text('projectId').notNull(),
    runtimeEnvironmentId: text('runtimeEnvironmentId').notNull(),
    version: text('version').notNull(),
    metadata: jsonb('metadata').notNull(),
    runtime: text('runtime'),
    runtimeVersion: text('runtimeVersion'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    workerGroupId: text('workerGroupId'),
    supportsLazyAttempts: boolean('supportsLazyAttempts').notNull().default(false),
  },
  (table) => [
    uniqueIndex('BackgroundWorker_friendlyId_key').on(table.friendlyId),
    index('BackgroundWorker_runtimeEnvironmentId_idx').on(table.runtimeEnvironmentId),
    index('BackgroundWorker_runtimeEnvironmentId_createdAt_idx').on(
      table.runtimeEnvironmentId,
      table.createdAt,
    ),
    uniqueIndex('BackgroundWorker_projectId_runtimeEnvironmentId_version_key').on(
      table.projectId,
      table.runtimeEnvironmentId,
      table.version,
    ),
    foreignKey({
      name: 'BackgroundWorker_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'BackgroundWorker_runtimeEnvironmentId_fkey',
      columns: [table.runtimeEnvironmentId],
      foreignColumns: [runtimeEnvironment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'BackgroundWorker_workerGroupId_fkey',
      columns: [table.workerGroupId],
      foreignColumns: [workerInstanceGroup.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const backgroundWorkerFile = pgTable(
  'BackgroundWorkerFile',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    friendlyId: text('friendlyId').notNull(),
    filePath: text('filePath').notNull(),
    contentHash: text('contentHash').notNull(),
    contents: bytea('contents').notNull(),
    projectId: text('projectId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    uniqueIndex('BackgroundWorkerFile_friendlyId_key').on(table.friendlyId),
    uniqueIndex('BackgroundWorkerFile_projectId_contentHash_key').on(
      table.projectId,
      table.contentHash,
    ),
    foreignKey({
      name: 'BackgroundWorkerFile_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const prompts = pgTable(
  'prompts',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    friendlyId: text('friendly_id').notNull(),
    slug: text('slug').notNull(),
    description: text('description'),
    type: text('type').notNull().default('text'),
    organizationId: text('organizationId').notNull(),
    projectId: text('projectId').notNull(),
    runtimeEnvironmentId: text('runtimeEnvironmentId').notNull(),
    filePath: text('filePath'),
    exportName: text('exportName'),
    variableSchema: jsonb('variableSchema'),
    defaultModel: text('defaultModel'),
    defaultConfig: jsonb('defaultConfig'),
    tags: text('tags').array().default([]),
    archivedAt: timestamp('archivedAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('prompts_friendly_id_key').on(table.friendlyId),
    index('prompts_projectId_idx').on(table.projectId),
    index('prompts_runtimeEnvironmentId_idx').on(table.runtimeEnvironmentId),
    uniqueIndex('prompts_projectId_runtimeEnvironmentId_slug_key').on(
      table.projectId,
      table.runtimeEnvironmentId,
      table.slug,
    ),
    foreignKey({
      name: 'prompts_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'prompts_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'prompts_runtimeEnvironmentId_fkey',
      columns: [table.runtimeEnvironmentId],
      foreignColumns: [runtimeEnvironment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const promptVersions = pgTable(
  'prompt_versions',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    version: integer('version').notNull(),
    promptId: text('promptId').notNull(),
    textContent: text('textContent'),
    chatContent: jsonb('chatContent'),
    model: text('model'),
    config: jsonb('config'),
    source: text('source').notNull(),
    commitMessage: text('commitMessage'),
    contentHash: text('contentHash').notNull(),
    labels: text('labels').array().default([]),
    workerId: text('workerId'),
    createdBy: text('createdBy'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('prompt_versions_promptId_idx').on(table.promptId),
    index('prompt_versions_contentHash_idx').on(table.contentHash),
    uniqueIndex('prompt_versions_promptId_version_key').on(table.promptId, table.version),
    foreignKey({
      name: 'prompt_versions_promptId_fkey',
      columns: [table.promptId],
      foreignColumns: [prompts.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'prompt_versions_workerId_fkey',
      columns: [table.workerId],
      foreignColumns: [backgroundWorker.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const backgroundWorkerTask = pgTable(
  'BackgroundWorkerTask',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    slug: text('slug').notNull(),
    description: text('description'),
    friendlyId: text('friendlyId').notNull(),
    filePath: text('filePath').notNull(),
    exportName: text('exportName'),
    workerId: text('workerId').notNull(),
    projectId: text('projectId').notNull(),
    fileId: text('fileId'),
    runtimeEnvironmentId: text('runtimeEnvironmentId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    queueConfig: jsonb('queueConfig'),
    retryConfig: jsonb('retryConfig'),
    machineConfig: jsonb('machineConfig'),
    /**
     * Gates declared on the task: other queues its runs must also hold a concurrency
     * slot in while executing. Shape: [{ queue: string, concurrencyKey?: string }]
     */
    gates: jsonb('gates'),
    queueId: text('queueId'),
    maxDurationInSeconds: integer('maxDurationInSeconds'),
    ttl: text('ttl'),
    triggerSource: taskTriggerSourceEnum('triggerSource').notNull().default('STANDARD'),
    /**
     * Extra task configuration JSON. Shape depends on triggerSource.
     * AGENT: { type: "ai-sdk-chat" }
     */
    config: jsonb('config'),
    executionMode: taskExecutionModeEnum('executionMode').notNull().default('STANDARD'),
    /** RunExecutionPolicy from @trigger.dev/core. Null means the standard policy. */
    executionPolicy: jsonb('executionPolicy'),
    payloadSchema: jsonb('payloadSchema'),
  },
  (table) => [
    uniqueIndex('BackgroundWorkerTask_friendlyId_key').on(table.friendlyId),
    index('BackgroundWorkerTask_projectId_slug_idx').on(table.projectId, table.slug),
    index('BackgroundWorkerTask_projectId_slug_createdAt_idx').on(
      table.projectId,
      table.slug,
      table.createdAt,
    ),
    index('BackgroundWorkerTask_runtimeEnvironmentId_projectId_idx').on(
      table.runtimeEnvironmentId,
      table.projectId,
    ),
    index('BackgroundWorkerTask_runtimeEnvironmentId_slug_triggerSourc_idx').on(
      table.runtimeEnvironmentId,
      table.slug,
      table.triggerSource,
    ),
    uniqueIndex('BackgroundWorkerTask_workerId_slug_key').on(table.workerId, table.slug),
    foreignKey({
      name: 'BackgroundWorkerTask_workerId_fkey',
      columns: [table.workerId],
      foreignColumns: [backgroundWorker.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'BackgroundWorkerTask_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'BackgroundWorkerTask_fileId_fkey',
      columns: [table.fileId],
      foreignColumns: [backgroundWorkerFile.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'BackgroundWorkerTask_runtimeEnvironmentId_fkey',
      columns: [table.runtimeEnvironmentId],
      foreignColumns: [runtimeEnvironment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'BackgroundWorkerTask_queueId_fkey',
      columns: [table.queueId],
      foreignColumns: [taskQueue.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const webhookEndpoint = pgTable(
  'WebhookEndpoint',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    friendlyId: text('friendlyId').notNull(),
    /**
     * URL key. CSPRNG, NOT a friendlyId: crypto.randomBytes(16) base64url.
     * Generated by the create path (deploy-sync / P2 API), NOT by @default.
     */
    opaqueId: text('opaqueId').notNull(),
    organizationId: text('organizationId').notNull(),
    projectId: text('projectId').notNull(),
    runtimeEnvironmentId: text('runtimeEnvironmentId').notNull(),
    environmentType: runtimeEnvironmentTypeEnum('environmentType').notNull(),
    endpointTenantId: text('endpointTenantId').notNull().default(''),
    endpointExternalRef: text('endpointExternalRef').notNull().default(''),
    source: text('source').notNull(),
    /** The declared webhooks.endpoint.define id (string ref, no relation). */
    declaredId: text('declaredId').notNull(),
    /** Subscribers of this endpoint: StoredWebhookRoutingTarget[] ({ id, type: "task" | "session", ..., filterAst }). */
    routingTargets: jsonb('routingTargets').notNull().default([]),
    verifierArtifact: jsonb('verifierArtifact').notNull(),
    metadata: jsonb('metadata').notNull().default({}),
    secretProvisioning: text('secretProvisioning').notNull().default('either'),
    /** The source's instructions for an AI agent registering this endpoint with the provider. */
    setupPrompt: text('setupPrompt'),
    /** SecretReference.key string. Plain String, NO @relation -> no FK to SecretReference. */
    signingSecretKey: text('signingSecretKey'),
    status: webhookEndpointStatusEnum('status').notNull().default('ACTIVE'),
    /**
     * When an operator disabled the endpoint via the dashboard/API. Null means the declarative sync
     * owns the status: a redeploy that re-declares a previously-removed (auto-deactivated) webhook
     * reactivates it. Non-null means the operator disabled it, so the sync leaves the status alone.
     */
    manuallyDeactivatedAt: timestamp('manuallyDeactivatedAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('WebhookEndpoint_friendlyId_key').on(table.friendlyId),
    uniqueIndex('WebhookEndpoint_opaqueId_key').on(table.opaqueId),
    index('WebhookEndpoint_runtimeEnvironmentId_source_idx').on(
      table.runtimeEnvironmentId,
      table.source,
    ),
    index('WebhookEndpoint_projectId_createdAt_idx').on(table.projectId, table.createdAt),
    uniqueIndex('WebhookEndpoint_runtimeEnvironmentId_declaredId_endpointTen_key').on(
      table.runtimeEnvironmentId,
      table.declaredId,
      table.endpointTenantId,
      table.endpointExternalRef,
    ),
  ],
)

export const webhookDelivery = pgTable(
  'WebhookDelivery',
  {
    id: text('id')
      .notNull()
      .$defaultFn(() => cuid()),
    friendlyId: text('friendlyId').notNull(),
    webhookEndpointId: text('webhookEndpointId').notNull(),
    organizationId: text('organizationId').notNull(),
    projectId: text('projectId').notNull(),
    runtimeEnvironmentId: text('runtimeEnvironmentId').notNull(),
    environmentType: runtimeEnvironmentTypeEnum('environmentType').notNull(),
    externalDeliveryId: text('externalDeliveryId').notNull(),
    idempotencyKey: text('idempotencyKey').notNull(),
    runId: text('runId'),
    status: webhookDeliveryStatusEnum('status').notNull().default('PENDING'),
    /** Set from the x-trigger-test ingress header; marks console/test-send deliveries so the list can filter them. */
    isTest: boolean('isTest').notNull().default(false),
    parsedEvent: jsonb('parsedEvent'),
    headers: jsonb('headers'),
    rawBodyHash: text('rawBodyHash'),
    errorMessage: text('errorMessage'),
    filterReason: text('filterReason'),
    /** Per-target results: WebhookDeliveryTargetResult[] ({ id, type, status, reason?, runId?, error? }). */
    targets: jsonb('targets').notNull().default([]),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    processedAt: timestamp('processedAt', { precision: 3 }),
    /** Retention class in days: the LIST partition key, stamped at ingest from the org's plan and encoded in the id. */
    retentionDays: integer('retentionDays').notNull().default(30),
  },
  (table) => [
    primaryKey({
      name: 'WebhookDelivery_pkey',
      columns: [table.id, table.createdAt, table.retentionDays],
    }),
    index('WebhookDelivery_webhookEndpointId_createdAt_idx').on(
      table.webhookEndpointId,
      table.createdAt,
    ),
  ],
)

export const alert = pgTable(
  'Alert',
  {
    id: text('id').primaryKey(),
    organizationId: text('organizationId').notNull(),
    projectId: text('projectId').notNull(),
    environmentId: text('environmentId').notNull(),
    capabilityKey: text('capabilityKey').notNull(),
    params: jsonb('params').notNull(),
    condition: jsonb('condition').notNull(),
    query: text('query'),
    expiresAt: timestamp('expiresAt', { withTimezone: true, precision: 6 }),
    lifecycle: text('lifecycle').notNull().default('active'),
    attention: text('attention').notNull().default('clear'),
    observedCondition: text('observedCondition').notNull().default('unknown'),
    observedConditionChangedAt: timestamp('observedConditionChangedAt', {
      withTimezone: true,
      precision: 6,
    }),
    pendingMetSince: timestamp('pendingMetSince', { withTimezone: true, precision: 6 }),
    evaluatorState: jsonb('evaluatorState'),
    evaluationLeaseId: text('evaluationLeaseId'),
    evaluationLeaseUntil: timestamp('evaluationLeaseUntil', { withTimezone: true, precision: 6 }),
    lastEvaluatedAt: timestamp('lastEvaluatedAt', { withTimezone: true, precision: 6 }),
    nextEvaluationAt: timestamp('nextEvaluationAt', { withTimezone: true, precision: 6 }),
    notifiers: jsonb('notifiers').notNull().default([]),
    name: text('name'),
    createdBy: text('createdBy').notNull(),
    createdAt: timestamp('createdAt', { withTimezone: true, precision: 6 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { withTimezone: true, precision: 6 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    resolvedAt: timestamp('resolvedAt', { withTimezone: true, precision: 6 }),
    resolvedBy: text('resolvedBy'),
    stoppedAt: timestamp('stoppedAt', { withTimezone: true, precision: 6 }),
    stoppedBy: text('stoppedBy'),
    lastOccurrenceAt: timestamp('lastOccurrenceAt', { withTimezone: true, precision: 6 }),
    occurrencesSinceRaise: integer('occurrencesSinceRaise').notNull().default(0),
    occurrencesTotal: bigint('occurrencesTotal', { mode: 'bigint' })
      .notNull()
      .default(sql`0`),
    groupId: text('groupId'),
  },
  (table) => [
    index('Alert_environmentId_lifecycle_idx').on(table.environmentId, table.lifecycle),
    index('Alert_lifecycle_evaluationLeaseUntil_idx').on(
      table.lifecycle,
      table.evaluationLeaseUntil,
    ),
    index('Alert_projectId_idx').on(table.projectId),
    index('Alert_organizationId_idx').on(table.organizationId),
    index('Alert_groupId_idx').on(table.groupId),
  ],
)

/** RANGE-partitioned by month on `at` (see migration). */
export const alertEvent = pgTable(
  'AlertEvent',
  {
    id: text('id').notNull(),
    alertId: text('alertId').notNull(),
    type: text('type').notNull(),
    at: timestamp('at', { withTimezone: true, precision: 6 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    facts: jsonb('facts').notNull().default({}),
    actorId: text('actorId'),
    sourceEventId: text('sourceEventId'),
  },
  (table) => [
    primaryKey({ name: 'AlertEvent_pkey', columns: [table.id, table.at] }),
    index('AlertEvent_alertId_at_idx').on(table.alertId, table.at),
  ],
)

export const alertOccurrenceSeen = pgTable(
  'AlertOccurrenceSeen',
  {
    alertId: text('alertId').notNull(),
    sourceEventId: text('sourceEventId').notNull(),
    rowTime: timestamp('rowTime', { withTimezone: true, precision: 6 }).notNull(),
    seenAt: timestamp('seenAt', { withTimezone: true, precision: 6 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    expiresAt: timestamp('expiresAt', { withTimezone: true, precision: 6 }).notNull(),
  },
  (table) => [
    primaryKey({ name: 'AlertOccurrenceSeen_pkey', columns: [table.alertId, table.sourceEventId] }),
    index('AlertOccurrenceSeen_expiresAt_idx').on(table.expiresAt),
  ],
)

/** RANGE-partitioned by month on `eventAt` (see migration). */
export const alertDelivery = pgTable(
  'AlertDelivery',
  {
    id: text('id').notNull(),
    alertEventId: text('alertEventId').notNull(),
    eventAt: timestamp('eventAt', { withTimezone: true, precision: 6 }).notNull(),
    notifierId: text('notifierId').notNull(),
    notifier: jsonb('notifier').notNull(),
    status: text('status').notNull().default('pending'),
    claimId: text('claimId'),
    claimedAt: timestamp('claimedAt', { withTimezone: true, precision: 6 }),
    deliveredAt: timestamp('deliveredAt', { withTimezone: true, precision: 6 }),
    failureReason: text('failureReason'),
    providerMetadata: jsonb('providerMetadata'),
    attempts: integer('attempts').notNull().default(0),
    firstAttemptAt: timestamp('firstAttemptAt', { withTimezone: true, precision: 6 }),
    createdAt: timestamp('createdAt', { withTimezone: true, precision: 6 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    primaryKey({ name: 'AlertDelivery_pkey', columns: [table.id, table.eventAt] }),
    index('AlertDelivery_status_createdAt_idx').on(table.status, table.createdAt),
    index('AlertDelivery_status_claimedAt_idx').on(table.status, table.claimedAt),
    uniqueIndex('AlertDelivery_alertEventId_notifierId_eventAt_key').on(
      table.alertEventId,
      table.notifierId,
      table.eventAt,
    ),
  ],
)

export const playgroundConversation = pgTable(
  'PlaygroundConversation',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    /** The chat session ID used by the transport */
    chatId: text('chatId').notNull(),
    /** User-editable conversation title (auto-generated from first message) */
    title: text('title').notNull().default('New conversation'),
    /** Which agent this conversation is with */
    agentSlug: text('agentSlug').notNull(),
    /** The current active run backing this conversation (null if no run yet) */
    runId: text('runId'),
    /** The client data JSON used for this conversation */
    clientData: jsonb('clientData'),
    /** Accumulated UIMessages from completed turns (for resume without stream replay) */
    messages: jsonb('messages'),
    /** Last SSE event ID — resume from this position to avoid replaying old turns */
    lastEventId: text('lastEventId'),
    projectId: text('projectId').notNull(),
    runtimeEnvironmentId: text('runtimeEnvironmentId').notNull(),
    /** The user who started this conversation */
    userId: text('userId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('PlaygroundConversation_runtimeEnvironmentId_agentSlug_updat_idx').on(
      table.runtimeEnvironmentId,
      table.agentSlug,
      table.updatedAt,
    ),
    index('PlaygroundConversation_userId_runtimeEnvironmentId_idx').on(
      table.userId,
      table.runtimeEnvironmentId,
    ),
    uniqueIndex('PlaygroundConversation_chatId_runtimeEnvironmentId_key').on(
      table.chatId,
      table.runtimeEnvironmentId,
    ),
    foreignKey({
      name: 'PlaygroundConversation_runId_fkey',
      columns: [table.runId],
      foreignColumns: [taskRun.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'PlaygroundConversation_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'PlaygroundConversation_runtimeEnvironmentId_fkey',
      columns: [table.runtimeEnvironmentId],
      foreignColumns: [runtimeEnvironment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * Durable, typed, bidirectional I/O primitive. Owns two S2 streams (.out / .in).
 * The row is essentially static — no status, no counters, no pointers. No
 * foreign keys: project/runtimeEnvironment/organization ids are plain
 * scalar columns (matches TaskRun pattern). List-style queries are served
 * from ClickHouse, not Postgres, so only point-lookup indexes live here.
 */
export const session = pgTable(
  'Session',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    friendlyId: text('friendlyId').notNull(),
    /**
     * User-supplied identifier scoped to the environment. Used for
     * idempotent upsert and for resolving sessions via the public API.
     */
    externalId: text('externalId'),
    /** Plain string — intentionally not an enum. */
    type: text('type').notNull(),
    /** Denormalized scoping columns — no FK relations. */
    projectId: text('projectId').notNull(),
    runtimeEnvironmentId: text('runtimeEnvironmentId').notNull(),
    environmentType: runtimeEnvironmentTypeEnum('environmentType').notNull(),
    organizationId: text('organizationId').notNull(),
    /**
     * Task this session triggers runs against. Required — Sessions are
     * task-bound: creating a session also triggers its first run, and
     * every subsequent re-trigger uses this same identifier.
     */
    taskIdentifier: text('taskIdentifier').notNull(),
    /**
     * Trigger config used for every run this session schedules. Shape
     * (validated at the route layer, opaque to the DB):
     * { basePayload: object, machine?: string, queue?: string,
     * tags?: string[], maxAttempts?: number,
     * idleTimeoutInSeconds?: number }
     * `basePayload` carries the customer's client-data; runtime fields
     * (chatId, messages, trigger) are merged at trigger time.
     */
    triggerConfig: jsonb('triggerConfig').notNull(),
    /**
     * Whether this session was created from the Test/playground UI rather
     * than by a real chat.agent() trigger. Mirrors TaskRun.isTest.
     */
    isTest: boolean('isTest').notNull().default(false),
    tags: text('tags').array().default([]),
    metadata: jsonb('metadata'),
    /**
     * Live run pointer — non-FK so run deletion never cascades. Can lag
     * reality; the `.in/append` handler re-checks the snapshot status
     * before reusing it.
     */
    currentRunId: text('currentRunId'),
    /**
     * Monotonic counter used for optimistic locking on `currentRunId`
     * swaps. Bumped atomically alongside any update that changes
     * `currentRunId`.
     */
    currentRunVersion: integer('currentRunVersion').notNull().default(0),
    /** Terminal markers — written once, never flipped back. */
    closedAt: timestamp('closedAt', { precision: 3 }),
    closedReason: text('closedReason'),
    expiresAt: timestamp('expiresAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    /**
     * S2 basin where this session's stream pair lives. Stamped at create
     * time from `Organization.streamBasinName` so reads can resolve the
     * basin without joining org. Null when the org has no per-org basin
     * (OSS, or pre-backfill); reads fall back to the global basin.
     */
    streamBasinName: text('streamBasinName'),
    /**
     * Storage URI (with protocol prefix) for this session's chat.agent
     * snapshot blob. Set on first snapshot write. Null = pre-column session,
     * fall back to computed default path.
     */
    chatSnapshotStoragePath: text('chatSnapshotStoragePath'),
  },
  (table) => [
    uniqueIndex('Session_friendlyId_key').on(table.friendlyId),
    index('Session_expiresAt_idx').on(table.expiresAt),
    index('Session_currentRunId_idx').on(table.currentRunId),
    uniqueIndex('Session_runtimeEnvironmentId_externalId_key').on(
      table.runtimeEnvironmentId,
      table.externalId,
    ),
  ],
)

/**
 * Historical record of every run a Session has owned. Append-only —
 * rows are inserted on each `ensureRunForSession` claim, never updated.
 * Lets us reconstruct the run timeline of a chat for debugging /
 * dashboard surfaces. The relation cascades on Session delete (tied to
 * the session lifecycle) but `runId` is a plain string column with no
 * FK to TaskRun so run pruning is independent.
 */
export const sessionRun = pgTable(
  'SessionRun',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    sessionId: text('sessionId').notNull(),
    /** TaskRun.id (no FK — runs may be archived independently of session history) */
    runId: text('runId').notNull(),
    /**
     * One of: "initial" | "continuation" | "upgrade" | "manual".
     * Plain string for forward-compat with future trigger reasons.
     */
    reason: text('reason').notNull(),
    triggeredAt: timestamp('triggeredAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    uniqueIndex('SessionRun_runId_key').on(table.runId),
    index('SessionRun_sessionId_idx').on(table.sessionId),
    index('SessionRun_sessionId_triggeredAt_id_idx').on(
      table.sessionId,
      table.triggeredAt,
      table.id,
    ),
    foreignKey({
      name: 'SessionRun_sessionId_fkey',
      columns: [table.sessionId],
      foreignColumns: [session.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const taskRun = pgTable(
  'TaskRun',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    number: integer('number').notNull().default(0),
    friendlyId: text('friendlyId').notNull(),
    engine: runEngineVersionEnum('engine').notNull().default('V1'),
    status: taskRunStatusEnum('status').notNull().default('PENDING'),
    statusReason: text('statusReason'),
    idempotencyKey: text('idempotencyKey'),
    idempotencyKeyExpiresAt: timestamp('idempotencyKeyExpiresAt', { precision: 3 }),
    /** Stores the user-provided key and scope: { key: string, scope: "run" | "attempt" | "global" } */
    idempotencyKeyOptions: jsonb('idempotencyKeyOptions'),
    /** Debounce options: { key: string, delay: string, createdAt: Date } */
    debounce: jsonb('debounce'),
    taskIdentifier: text('taskIdentifier').notNull(),
    isTest: boolean('isTest').notNull().default(false),
    payload: text('payload').notNull(),
    payloadType: text('payloadType').notNull().default('application/json'),
    context: jsonb('context'),
    traceContext: jsonb('traceContext'),
    traceId: text('traceId').notNull(),
    spanId: text('spanId').notNull(),
    runtimeEnvironmentId: text('runtimeEnvironmentId').notNull(),
    environmentType: runtimeEnvironmentTypeEnum('environmentType'),
    projectId: text('projectId').notNull(),
    organizationId: text('organizationId'),
    queue: text('queue').notNull(),
    lockedQueueId: text('lockedQueueId'),
    /** The main queue that this run is part of */
    workerQueue: text('masterQueue').notNull().default('main'),
    /** User-facing geo region, stamped at trigger; workerQueue is where it actually ran. */
    region: text('region'),
    /** @deprecated */
    secondaryMasterQueue: text('secondaryMasterQueue'),
    /** From engine v2+ this will be defined after a run has been dequeued (starting at 1) */
    attemptNumber: integer('attemptNumber'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    /** Denormized column that holds the raw tags */
    runTags: text('runTags').array(),
    /** Denormalized version of the background worker task */
    taskVersion: text('taskVersion'),
    sdkVersion: text('sdkVersion'),
    cliVersion: text('cliVersion'),
    /** startedAt marks the point at which a run is dequeued from MarQS */
    startedAt: timestamp('startedAt', { precision: 3 }),
    /** executedAt is set when the first attempt is about to execute */
    executedAt: timestamp('executedAt', { precision: 3 }),
    completedAt: timestamp('completedAt', { precision: 3 }),
    machinePreset: text('machinePreset'),
    usageDurationMs: integer('usageDurationMs').notNull().default(0),
    costInCents: doublePrecision('costInCents').notNull().default(0),
    baseCostInCents: doublePrecision('baseCostInCents').notNull().default(0),
    lockedAt: timestamp('lockedAt', { precision: 3 }),
    lockedById: text('lockedById'),
    lockedToVersionId: text('lockedToVersionId'),
    /**
     * The "priority" of the run. This is just a negative offset in ms for the queue timestamp
     * E.g. a value of 60_000 would put the run into the queue 60s ago.
     */
    priorityMs: integer('priorityMs').notNull().default(0),
    concurrencyKey: text('concurrencyKey'),
    /**
     * Gates for this run: other queues it must also hold a concurrency slot in while
     * executing. Shape: [{ queue: string, concurrencyKey?: string }]
     */
    gates: jsonb('gates'),
    delayUntil: timestamp('delayUntil', { precision: 3 }),
    queuedAt: timestamp('queuedAt', { precision: 3 }),
    ttl: text('ttl'),
    expiredAt: timestamp('expiredAt', { precision: 3 }),
    maxAttempts: integer('maxAttempts'),
    lockedRetryConfig: jsonb('lockedRetryConfig'),
    /** optional token that can be used to authenticate the task run */
    oneTimeUseToken: text('oneTimeUseToken'),
    /** Where the logs are stored */
    taskEventStore: text('taskEventStore').notNull().default('taskEvent'),
    queueTimestamp: timestamp('queueTimestamp', { precision: 3 }),
    scheduleInstanceId: text('scheduleInstanceId'),
    scheduleId: text('scheduleId'),
    bulkActionGroupIds: text('bulkActionGroupIds').array().default([]),
    logsDeletedAt: timestamp('logsDeletedAt', { precision: 3 }),
    replayedFromTaskRunFriendlyId: text('replayedFromTaskRunFriendlyId'),
    rootTaskRunId: text('rootTaskRunId'),
    parentTaskRunId: text('parentTaskRunId'),
    parentTaskRunAttemptId: text('parentTaskRunAttemptId'),
    batchId: text('batchId'),
    /** whether or not the task run was created because of a triggerAndWait for batchTriggerAndWait */
    resumeParentOnCompletion: boolean('resumeParentOnCompletion').notNull().default(false),
    /** The depth of this task run in the task run hierarchy */
    depth: integer('depth').notNull().default(0),
    /** The span ID of the "trigger" span in the parent task run */
    parentSpanId: text('parentSpanId'),
    /** Holds the state of the run chain for deadlock detection */
    runChainState: jsonb('runChainState'),
    /** seed run metadata */
    seedMetadata: text('seedMetadata'),
    seedMetadataType: text('seedMetadataType').notNull().default('application/json'),
    /** Run metadata */
    metadata: text('metadata'),
    metadataType: text('metadataType').notNull().default('application/json'),
    metadataVersion: integer('metadataVersion').notNull().default(1),
    /** Structured annotations: triggerSource, triggerAction, rootTriggerSource, rootScheduleId */
    annotations: jsonb('annotations'),
    /** Whether the latest attempt was a warm start. Null until first attempt starts. */
    isWarmStart: boolean('isWarmStart'),
    /** Run output */
    output: text('output'),
    outputType: text('outputType').notNull().default('application/json'),
    /** Run error */
    error: jsonb('error'),
    /** Organization's billing plan type (cached for fallback when billing API fails) */
    planType: text('planType'),
    maxDurationInSeconds: integer('maxDurationInSeconds'),
    /** The version of the realtime streams implementation used by the run */
    realtimeStreamsVersion: text('realtimeStreamsVersion').notNull().default('v1'),
    /** Store the stream keys that are being used by the run */
    realtimeStreams: text('realtimeStreams').array().default([]),
    /**
     * S2 basin where this run's realtime streams live. Stamped at create
     * time from `Organization.streamBasinName` so reads can resolve the
     * basin without joining org. Null when the org has no per-org basin
     * (OSS, or pre-backfill); reads fall back to the global basin.
     */
    streamBasinName: text('streamBasinName'),
  },
  (table): PgTableExtraConfigValue[] => [
    uniqueIndex('TaskRun_friendlyId_key').on(table.friendlyId),
    index('TaskRun_parentTaskRunId_idx').on(table.parentTaskRunId),
    index('TaskRun_spanId_idx').on(table.spanId),
    index('TaskRun_parentSpanId_idx').on(table.parentSpanId),
    index('TaskRun_runTags_idx').on(table.runTags),
    index('TaskRun_runtimeEnvironmentId_batchId_idx').on(table.runtimeEnvironmentId, table.batchId),
    index('TaskRun_runtimeEnvironmentId_createdAt_idx').on(
      table.runtimeEnvironmentId,
      table.createdAt,
    ),
    index('TaskRun_createdAt_idx').on(table.createdAt),
    uniqueIndex('TaskRun_oneTimeUseToken_key').on(table.oneTimeUseToken),
    uniqueIndex('TaskRun_runtimeEnvironmentId_taskIdentifier_idempotencyKey_key').on(
      table.runtimeEnvironmentId,
      table.taskIdentifier,
      table.idempotencyKey,
    ),
    foreignKey({
      name: 'TaskRun_runtimeEnvironmentId_fkey',
      columns: [table.runtimeEnvironmentId],
      foreignColumns: [runtimeEnvironment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'TaskRun_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'TaskRun_lockedById_fkey',
      columns: [table.lockedById],
      foreignColumns: [backgroundWorkerTask.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'TaskRun_lockedToVersionId_fkey',
      columns: [table.lockedToVersionId],
      foreignColumns: [backgroundWorker.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'TaskRun_rootTaskRunId_fkey',
      columns: [table.rootTaskRunId],
      foreignColumns: [table.id],
    })
      .onDelete('set null')
      .onUpdate('no action'),
    foreignKey({
      name: 'TaskRun_parentTaskRunId_fkey',
      columns: [table.parentTaskRunId],
      foreignColumns: [table.id],
    })
      .onDelete('set null')
      .onUpdate('no action'),
    foreignKey({
      name: 'TaskRun_parentTaskRunAttemptId_fkey',
      columns: [table.parentTaskRunAttemptId],
      foreignColumns: [taskRunAttempt.id],
    })
      .onDelete('set null')
      .onUpdate('no action'),
    foreignKey({
      name: 'TaskRun_batchId_fkey',
      columns: [table.batchId],
      foreignColumns: [batchTaskRun.id],
    })
      .onDelete('set null')
      .onUpdate('no action'),
  ],
)

export const taskRunTemplate = pgTable(
  'TaskRunTemplate',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    taskSlug: text('taskSlug').notNull(),
    triggerSource: taskTriggerSourceEnum('triggerSource').notNull(),
    label: text('label').notNull(),
    payload: text('payload'),
    payloadType: text('payloadType').notNull().default('application/json'),
    metadata: text('metadata'),
    metadataType: text('metadataType').notNull().default('application/json'),
    queue: text('queue'),
    concurrencyKey: text('concurrencyKey'),
    ttlSeconds: integer('ttlSeconds'),
    maxAttempts: integer('maxAttempts'),
    maxDurationSeconds: integer('maxDurationSeconds'),
    tags: text('tags').array(),
    machinePreset: text('machinePreset'),
    projectId: text('projectId').notNull(),
    organizationId: text('organizationId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('TaskRunTemplate_projectId_taskSlug_triggerSource_createdAt_idx').on(
      table.projectId,
      table.taskSlug,
      table.triggerSource,
      table.createdAt,
    ),
  ],
)

/**
 * Used by the RunEngine during TaskRun execution
 * It has the required information to transactionally progress a run through states,
 * and prevent side effects like heartbeats failing a run that has progressed.
 * It is optimised for performance and is designed to be cleared at some point,
 * so there are no cascading relationships to other models.
 */
export const taskRunExecutionSnapshot = pgTable(
  'TaskRunExecutionSnapshot',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    /** This should always be 2+ (V1 didn't use the run engine or snapshots) */
    engine: runEngineVersionEnum('engine').notNull().default('V2'),
    /** The execution status */
    executionStatus: taskRunExecutionStatusEnum('executionStatus').notNull(),
    /** For debugging */
    description: text('description').notNull(),
    /** We store invalid snapshots as a record of the run state when we tried to move */
    isValid: boolean('isValid').notNull().default(true),
    error: text('error'),
    /** The previous snapshot ID */
    previousSnapshotId: text('previousSnapshotId'),
    /** Run */
    runId: text('runId').notNull(),
    runStatus: taskRunStatusEnum('runStatus').notNull(),
    batchId: text('batchId'),
    /** This is the current run attempt number. Users can define how many attempts they want for a run. */
    attemptNumber: integer('attemptNumber'),
    /** Environment */
    environmentId: text('environmentId').notNull(),
    environmentType: runtimeEnvironmentTypeEnum('environmentType').notNull(),
    projectId: text('projectId').notNull(),
    organizationId: text('organizationId').notNull(),
    /** An array of waitpoint IDs in the correct order, used for batches */
    completedWaitpointOrder: text('completedWaitpointOrder').array(),
    /** Checkpoint */
    checkpointId: text('checkpointId'),
    /** Worker */
    workerId: text('workerId'),
    runnerId: text('runnerId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    lastHeartbeatAt: timestamp('lastHeartbeatAt', { precision: 3 }),
    /** Metadata used by various systems in the run engine */
    metadata: jsonb('metadata'),
  },
  (table) => [
    index('TaskRunExecutionSnapshot_runId_isValid_createdAt_idx').on(
      table.runId,
      table.isValid,
      table.createdAt,
    ),
    foreignKey({
      name: 'TaskRunExecutionSnapshot_runId_fkey',
      columns: [table.runId],
      foreignColumns: [taskRun.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'TaskRunExecutionSnapshot_checkpointId_fkey',
      columns: [table.checkpointId],
      foreignColumns: [taskRunCheckpoint.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const taskRunCheckpoint = pgTable(
  'TaskRunCheckpoint',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    friendlyId: text('friendlyId').notNull(),
    type: taskRunCheckpointTypeEnum('type').notNull(),
    location: text('location').notNull(),
    imageRef: text('imageRef'),
    reason: text('reason'),
    metadata: text('metadata'),
    projectId: text('projectId').notNull(),
    runtimeEnvironmentId: text('runtimeEnvironmentId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('TaskRunCheckpoint_friendlyId_key').on(table.friendlyId),
    foreignKey({
      name: 'TaskRunCheckpoint_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'TaskRunCheckpoint_runtimeEnvironmentId_fkey',
      columns: [table.runtimeEnvironmentId],
      foreignColumns: [runtimeEnvironment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * A Waitpoint blocks a run from continuing until it's completed
 * If there's a waitpoint blocking a run, it shouldn't be in the queue
 */
export const waitpoint = pgTable(
  'Waitpoint',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    friendlyId: text('friendlyId').notNull(),
    type: waitpointTypeEnum('type').notNull(),
    status: waitpointStatusEnum('status').notNull().default('PENDING'),
    completedAt: timestamp('completedAt', { precision: 3 }),
    /** If it's an Event type waitpoint, this is the event. It can also be provided for the DATETIME type */
    idempotencyKey: text('idempotencyKey').notNull(),
    /** If this is true then we can show it in the dashboard/return it from the SDK */
    userProvidedIdempotencyKey: boolean('userProvidedIdempotencyKey').notNull(),
    /** If there's a user provided idempotency key, this is the time it expires at */
    idempotencyKeyExpiresAt: timestamp('idempotencyKeyExpiresAt', { precision: 3 }),
    /**
     * If an idempotencyKey is no longer active, we store it here and generate a new one for the idempotencyKey field.
     * Clearing an idempotencyKey is useful for debounce or cancelling child runs.
     * This is a workaround because Prisma doesn't support partial indexes.
     */
    inactiveIdempotencyKey: text('inactiveIdempotencyKey'),
    /** If it's a RUN type waitpoint, this is the associated run */
    completedByTaskRunId: text('completedByTaskRunId'),
    /**
     * If it's a DATETIME type waitpoint, this is the date.
     * If it's a MANUAL waitpoint, this can be set as the `timeout`.
     */
    completedAfter: timestamp('completedAfter', { precision: 3 }),
    /** If it's a BATCH type waitpoint, this is the associated batch */
    completedByBatchId: text('completedByBatchId'),
    /** When completed, an output can be stored here */
    output: text('output'),
    outputType: text('outputType').notNull().default('application/json'),
    outputIsError: boolean('outputIsError').notNull().default(false),
    projectId: text('projectId').notNull(),
    environmentId: text('environmentId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    /**
     * Denormized column that holds the raw tags
     * Denormalized column that holds the raw tags
     */
    tags: text('tags').array(),
  },
  (table) => [
    uniqueIndex('Waitpoint_friendlyId_key').on(table.friendlyId),
    uniqueIndex('Waitpoint_completedByTaskRunId_key').on(table.completedByTaskRunId),
    index('Waitpoint_completedByBatchId_idx').on(table.completedByBatchId),
    index('Waitpoint_environmentId_type_createdAt_idx').on(
      table.environmentId,
      table.type,
      table.createdAt,
    ),
    index('Waitpoint_environmentId_type_status_idx').on(
      table.environmentId,
      table.type,
      table.status,
    ),
    index('Waitpoint_environmentId_type_id_idx').on(table.environmentId, table.type, table.id),
    uniqueIndex('Waitpoint_environmentId_idempotencyKey_key').on(
      table.environmentId,
      table.idempotencyKey,
    ),
    foreignKey({
      name: 'Waitpoint_completedByTaskRunId_fkey',
      columns: [table.completedByTaskRunId],
      foreignColumns: [taskRun.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Waitpoint_completedByBatchId_fkey',
      columns: [table.completedByBatchId],
      foreignColumns: [batchTaskRun.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Waitpoint_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Waitpoint_environmentId_fkey',
      columns: [table.environmentId],
      foreignColumns: [runtimeEnvironment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const taskRunWaitpoint = pgTable(
  'TaskRunWaitpoint',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    taskRunId: text('taskRunId').notNull(),
    waitpointId: text('waitpointId').notNull(),
    projectId: text('projectId').notNull(),
    /** This span id is completed when the waitpoint is completed. This is used with cached runs (idempotent) */
    spanIdToComplete: text('spanIdToComplete'),
    batchId: text('batchId'),
    batchIndex: integer('batchIndex'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('TaskRunWaitpoint_taskRunId_idx').on(table.taskRunId),
    index('TaskRunWaitpoint_waitpointId_idx').on(table.waitpointId),
    uniqueIndex('TaskRunWaitpoint_taskRunId_waitpointId_batchIndex_key').on(
      table.taskRunId,
      table.waitpointId,
      table.batchIndex,
    ),
    foreignKey({
      name: 'TaskRunWaitpoint_taskRunId_fkey',
      columns: [table.taskRunId],
      foreignColumns: [taskRun.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'TaskRunWaitpoint_waitpointId_fkey',
      columns: [table.waitpointId],
      foreignColumns: [waitpoint.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'TaskRunWaitpoint_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'TaskRunWaitpoint_batchId_fkey',
      columns: [table.batchId],
      foreignColumns: [batchTaskRun.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const waitpointTag = pgTable(
  'WaitpointTag',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    name: text('name').notNull(),
    environmentId: text('environmentId').notNull(),
    projectId: text('projectId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    uniqueIndex('WaitpointTag_environmentId_name_key').on(table.environmentId, table.name),
    foreignKey({
      name: 'WaitpointTag_environmentId_fkey',
      columns: [table.environmentId],
      foreignColumns: [runtimeEnvironment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'WaitpointTag_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const featureFlag = pgTable(
  'FeatureFlag',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    key: text('key').notNull(),
    value: jsonb('value'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [uniqueIndex('FeatureFlag_key_key').on(table.key)],
)

export const workerInstance = pgTable(
  'WorkerInstance',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    /** For example "worker-1" */
    name: text('name').notNull(),
    /**
     * If managed, it will default to the name, e.g. "worker-1"
     * If unmanged, it will be prefixed with the deployment ID e.g. "deploy-123-worker-1"
     */
    resourceIdentifier: text('resourceIdentifier').notNull(),
    metadata: jsonb('metadata'),
    workerGroupId: text('workerGroupId').notNull(),
    organizationId: text('organizationId'),
    projectId: text('projectId'),
    environmentId: text('environmentId'),
    deploymentId: text('deploymentId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    lastDequeueAt: timestamp('lastDequeueAt', { precision: 3 }),
    lastHeartbeatAt: timestamp('lastHeartbeatAt', { precision: 3 }),
  },
  (table) => [
    uniqueIndex('WorkerInstance_workerGroupId_resourceIdentifier_key').on(
      table.workerGroupId,
      table.resourceIdentifier,
    ),
    foreignKey({
      name: 'WorkerInstance_workerGroupId_fkey',
      columns: [table.workerGroupId],
      foreignColumns: [workerInstanceGroup.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'WorkerInstance_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'WorkerInstance_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'WorkerInstance_environmentId_fkey',
      columns: [table.environmentId],
      foreignColumns: [runtimeEnvironment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'WorkerInstance_deploymentId_fkey',
      columns: [table.deploymentId],
      foreignColumns: [workerDeployment.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const workerInstanceGroup = pgTable(
  'WorkerInstanceGroup',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    type: workerInstanceGroupTypeEnum('type').notNull(),
    /** For example "us-east-1" */
    name: text('name').notNull(),
    /**
     * If managed, it will default to the name, e.g. "us-east-1"
     * If unmanged, it will be prefixed with the project ID e.g. "project_1-us-east-1"
     */
    masterQueue: text('masterQueue').notNull(),
    /** "N. Virginia, USA", "Frankfurt, Germany", etc. Used for display purposes */
    description: text('description'),
    hidden: boolean('hidden').notNull().default(false),
    tokenId: text('tokenId').notNull(),
    organizationId: text('organizationId'),
    projectId: text('projectId'),
    cloudProvider: text('cloudProvider'),
    /** "usa", "europe", etc. Used like a pseudo enum for things like flags */
    location: text('location'),
    staticIPs: text('staticIPs'),
    workloadType: workloadTypeEnum('workloadType').notNull().default('CONTAINER'),
    /** Geo region; container + MICROVM groups for one geo share it. Set-once once it has runs. */
    region: text('region'),
    /**
     * When true, runs enqueued to this worker queue may skip the intermediate queue
     * and be pushed directly to the worker queue when concurrency is available.
     */
    enableFastPath: boolean('enableFastPath').notNull().default(false),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table): PgTableExtraConfigValue[] => [
    uniqueIndex('WorkerInstanceGroup_masterQueue_key').on(table.masterQueue),
    uniqueIndex('WorkerInstanceGroup_tokenId_key').on(table.tokenId),
    foreignKey({
      name: 'WorkerInstanceGroup_tokenId_fkey',
      columns: [table.tokenId],
      foreignColumns: [workerGroupToken.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'WorkerInstanceGroup_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'WorkerInstanceGroup_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const workerGroupToken = pgTable(
  'WorkerGroupToken',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    tokenHash: text('tokenHash').notNull(),
    previousTokenHash: text('previousTokenHash'),
    previousTokenValidUntil: timestamp('previousTokenValidUntil', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    workerGroupId: text('workerGroupId'),
  },
  (table): PgTableExtraConfigValue[] => [
    uniqueIndex('WorkerGroupToken_tokenHash_key').on(table.tokenHash),
    uniqueIndex('WorkerGroupToken_previousTokenHash_key').on(table.previousTokenHash),
    index('WorkerGroupToken_workerGroupId_idx').on(table.workerGroupId),
    foreignKey({
      name: 'WorkerGroupToken_workerGroupId_fkey',
      columns: [table.workerGroupId],
      foreignColumns: [workerInstanceGroup.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const taskRunTag = pgTable(
  'TaskRunTag',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    name: text('name').notNull(),
    friendlyId: text('friendlyId').notNull(),
    projectId: text('projectId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    uniqueIndex('TaskRunTag_friendlyId_key').on(table.friendlyId),
    index('TaskRunTag_name_id_idx').on(table.name, table.id),
    uniqueIndex('TaskRunTag_projectId_name_key').on(table.projectId, table.name),
    foreignKey({
      name: 'TaskRunTag_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/** This is used for triggerAndWait and batchTriggerAndWait. The taskRun is the child task, it points at a parent attempt or a batch */
export const taskRunDependency = pgTable(
  'TaskRunDependency',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    taskRunId: text('taskRunId').notNull(),
    checkpointEventId: text('checkpointEventId'),
    dependentAttemptId: text('dependentAttemptId'),
    dependentBatchRunId: text('dependentBatchRunId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    resumedAt: timestamp('resumedAt', { precision: 3 }),
  },
  (table) => [
    uniqueIndex('TaskRunDependency_taskRunId_key').on(table.taskRunId),
    uniqueIndex('TaskRunDependency_checkpointEventId_key').on(table.checkpointEventId),
    index('TaskRunDependency_dependentAttemptId_idx').on(table.dependentAttemptId),
    index('TaskRunDependency_dependentBatchRunId_idx').on(table.dependentBatchRunId),
    foreignKey({
      name: 'TaskRunDependency_taskRunId_fkey',
      columns: [table.taskRunId],
      foreignColumns: [taskRun.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'TaskRunDependency_checkpointEventId_fkey',
      columns: [table.checkpointEventId],
      foreignColumns: [checkpointRestoreEvent.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'TaskRunDependency_dependentAttemptId_fkey',
      columns: [table.dependentAttemptId],
      foreignColumns: [taskRunAttempt.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'TaskRunDependency_dependentBatchRunId_fkey',
      columns: [table.dependentBatchRunId],
      foreignColumns: [batchTaskRun.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

/** deprecated, we hadn't included the project id in the unique constraint */
export const taskRunCounter = pgTable('TaskRunCounter', {
  taskIdentifier: text('taskIdentifier').primaryKey(),
  lastNumber: integer('lastNumber').notNull().default(0),
})

/** Used for the TaskRun number (e.g. #1,421) */
export const taskRunNumberCounter = pgTable(
  'TaskRunNumberCounter',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    taskIdentifier: text('taskIdentifier').notNull(),
    environmentId: text('environmentId').notNull(),
    lastNumber: integer('lastNumber').notNull().default(0),
  },
  (table) => [
    uniqueIndex('TaskRunNumberCounter_taskIdentifier_environmentId_key').on(
      table.taskIdentifier,
      table.environmentId,
    ),
    foreignKey({
      name: 'TaskRunNumberCounter_environmentId_fkey',
      columns: [table.environmentId],
      foreignColumns: [runtimeEnvironment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/** This is not used from engine v2+, attempts use the TaskRunExecutionSnapshot and TaskRun */
export const taskRunAttempt = pgTable(
  'TaskRunAttempt',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    number: integer('number').notNull().default(0),
    friendlyId: text('friendlyId').notNull(),
    taskRunId: text('taskRunId').notNull(),
    backgroundWorkerId: text('backgroundWorkerId').notNull(),
    backgroundWorkerTaskId: text('backgroundWorkerTaskId').notNull(),
    runtimeEnvironmentId: text('runtimeEnvironmentId').notNull(),
    queueId: text('queueId').notNull(),
    status: taskRunAttemptStatusEnum('status').notNull().default('PENDING'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    startedAt: timestamp('startedAt', { precision: 3 }),
    completedAt: timestamp('completedAt', { precision: 3 }),
    usageDurationMs: integer('usageDurationMs').notNull().default(0),
    error: jsonb('error'),
    output: text('output'),
    outputType: text('outputType').notNull().default('application/json'),
  },
  (table): PgTableExtraConfigValue[] => [
    uniqueIndex('TaskRunAttempt_friendlyId_key').on(table.friendlyId),
    index('TaskRunAttempt_taskRunId_idx').on(table.taskRunId),
    uniqueIndex('TaskRunAttempt_taskRunId_number_key').on(table.taskRunId, table.number),
    foreignKey({
      name: 'TaskRunAttempt_taskRunId_fkey',
      columns: [table.taskRunId],
      foreignColumns: [taskRun.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'TaskRunAttempt_backgroundWorkerId_fkey',
      columns: [table.backgroundWorkerId],
      foreignColumns: [backgroundWorker.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'TaskRunAttempt_backgroundWorkerTaskId_fkey',
      columns: [table.backgroundWorkerTaskId],
      foreignColumns: [backgroundWorkerTask.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'TaskRunAttempt_runtimeEnvironmentId_fkey',
      columns: [table.runtimeEnvironmentId],
      foreignColumns: [runtimeEnvironment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'TaskRunAttempt_queueId_fkey',
      columns: [table.queueId],
      foreignColumns: [taskQueue.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/** This is the unified otel span/log model that will eventually be replaced by clickhouse */
export const taskEvent = pgTable(
  'TaskEvent',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    /** This matches the span name for a trace event, or the log body for a log event */
    message: text('message').notNull(),
    traceId: text('traceId').notNull(),
    spanId: text('spanId').notNull(),
    parentId: text('parentId'),
    tracestate: text('tracestate'),
    isError: boolean('isError').notNull().default(false),
    isPartial: boolean('isPartial').notNull().default(false),
    isCancelled: boolean('isCancelled').notNull().default(false),
    /** deprecated: don't use this, moving this to properties, this now uses TaskEventKind.LOG */
    isDebug: boolean('isDebug').notNull().default(false),
    serviceName: text('serviceName').notNull(),
    serviceNamespace: text('serviceNamespace').notNull(),
    level: taskEventLevelEnum('level').notNull().default('TRACE'),
    kind: taskEventKindEnum('kind').notNull().default('INTERNAL'),
    status: taskEventStatusEnum('status').notNull().default('UNSET'),
    links: jsonb('links'),
    events: jsonb('events'),
    /** This is the time the event started in nanoseconds since the epoch */
    startTime: bigint('startTime', { mode: 'bigint' }).notNull(),
    /** This is the duration of the event in nanoseconds */
    duration: bigint('duration', { mode: 'bigint' })
      .notNull()
      .default(sql`0`),
    attemptId: text('attemptId'),
    attemptNumber: integer('attemptNumber'),
    environmentId: text('environmentId').notNull(),
    environmentType: runtimeEnvironmentTypeEnum('environmentType').notNull(),
    organizationId: text('organizationId').notNull(),
    projectId: text('projectId').notNull(),
    projectRef: text('projectRef').notNull(),
    runId: text('runId').notNull(),
    runIsTest: boolean('runIsTest').notNull().default(false),
    idempotencyKey: text('idempotencyKey'),
    taskSlug: text('taskSlug').notNull(),
    taskPath: text('taskPath'),
    taskExportName: text('taskExportName'),
    workerId: text('workerId'),
    workerVersion: text('workerVersion'),
    queueId: text('queueId'),
    queueName: text('queueName'),
    batchId: text('batchId'),
    /** This represents all the span attributes available, like http.status_code, and special attributes like $style.icon, $output, $metadata.payload.userId, as it's used for searching and filtering */
    properties: jsonb('properties').notNull(),
    /** This represents all span attributes in the $metadata namespace, like $metadata.payload */
    metadata: jsonb('metadata'),
    /** This represents all span attributes in the $style namespace, like $style */
    style: jsonb('style'),
    /** This represents all span attributes in the $output namespace, like $output */
    output: jsonb('output'),
    /** This represents the mimetype of the output, such as application/json or application/super+json */
    outputType: text('outputType'),
    payload: jsonb('payload'),
    payloadType: text('payloadType'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    usageDurationMs: integer('usageDurationMs').notNull().default(0),
    usageCostInCents: doublePrecision('usageCostInCents').notNull().default(0),
    machinePreset: text('machinePreset'),
    machinePresetCpu: doublePrecision('machinePresetCpu'),
    machinePresetMemory: doublePrecision('machinePresetMemory'),
    machinePresetCentsPerMs: doublePrecision('machinePresetCentsPerMs'),
  },
  (table) => [
    index('TaskEvent_traceId_idx').on(table.traceId),
    index('TaskEvent_spanId_idx').on(table.spanId),
    index('TaskEvent_runId_idx').on(table.runId),
  ],
)

export const taskQueue = pgTable(
  'TaskQueue',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    friendlyId: text('friendlyId').notNull(),
    name: text('name').notNull(),
    type: taskQueueTypeEnum('type').notNull().default('VIRTUAL'),
    version: taskQueueVersionEnum('version').notNull().default('V1'),
    orderableName: text('orderableName'),
    /**
     * How the queue's concurrency was declared: V1 = legacy concurrencyLimit semantics
     * (per key when keyed, whole queue otherwise), V2 = the explicit perKey/total shape.
     */
    concurrencyVersion: taskQueueConcurrencyVersionEnum('concurrencyVersion')
      .notNull()
      .default('V1'),
    /**
     * QUEUE rows are real queues; LIMIT rows back named concurrency limits and are
     * excluded from every queue-facing surface.
     */
    role: taskQueueRoleEnum('role').notNull().default('QUEUE'),
    projectId: text('projectId').notNull(),
    runtimeEnvironmentId: text('runtimeEnvironmentId').notNull(),
    /** Represents the current concurrency limit for the queue */
    concurrencyLimit: integer('concurrencyLimit'),
    /** When the concurrency limit was overridden */
    concurrencyLimitOverriddenAt: timestamp('concurrencyLimitOverriddenAt', { precision: 3 }),
    /** Who overrode the concurrency limit (will be null if overridden via the API) */
    concurrencyLimitOverriddenBy: text('concurrencyLimitOverriddenBy'),
    /** If concurrencyLimit is overridden, this is the overridden value */
    concurrencyLimitBase: integer('concurrencyLimitBase'),
    /**
     * If the override was expressed as a percentage of the environment limit, this stores that
     * percentage (the source of truth). The absolute concurrencyLimit is materialized from it.
     * Decimal(5,2) allows fractional percentages like 12.50% (0.01–100.00).
     */
    concurrencyLimitOverridePercent: numeric('concurrencyLimitOverridePercent', {
      precision: 5,
      scale: 2,
    }),
    /**
     * Caps total concurrent runs across ALL concurrencyKey values of this queue
     * (concurrencyLimit applies per key value). Null = no total cap.
     */
    totalConcurrencyLimit: integer('totalConcurrencyLimit'),
    /** When the total concurrency limit was overridden */
    totalConcurrencyLimitOverriddenAt: timestamp('totalConcurrencyLimitOverriddenAt', {
      precision: 3,
    }),
    /** Who overrode the total concurrency limit (null when overridden via the API) */
    totalConcurrencyLimitOverriddenBy: text('totalConcurrencyLimitOverriddenBy'),
    /** If totalConcurrencyLimit is overridden, the declared value it reverts to on reset */
    totalConcurrencyLimitBase: integer('totalConcurrencyLimitBase'),
    rateLimit: jsonb('rateLimit'),
    paused: boolean('paused').notNull().default(false),
    /** Hidden from the dashboard queue list. Display-only: never read on trigger/dequeue paths. */
    archivedAt: timestamp('archivedAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('TaskQueue_friendlyId_key').on(table.friendlyId),
    uniqueIndex('TaskQueue_runtimeEnvironmentId_name_key').on(
      table.runtimeEnvironmentId,
      table.name,
    ),
    foreignKey({
      name: 'TaskQueue_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'TaskQueue_runtimeEnvironmentId_fkey',
      columns: [table.runtimeEnvironmentId],
      foreignColumns: [runtimeEnvironment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const batchTaskRun = pgTable(
  'BatchTaskRun',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    friendlyId: text('friendlyId').notNull(),
    idempotencyKey: text('idempotencyKey'),
    idempotencyKeyExpiresAt: timestamp('idempotencyKeyExpiresAt', { precision: 3 }),
    status: batchTaskRunStatusEnum('status').notNull().default('PENDING'),
    runtimeEnvironmentId: text('runtimeEnvironmentId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    /** Friendly IDs */
    runIds: text('runIds').array().default([]),
    runCount: integer('runCount').notNull().default(0),
    payload: text('payload'),
    payloadType: text('payloadType').notNull().default('application/json'),
    options: jsonb('options'),
    batchVersion: text('batchVersion').notNull().default('v1'),
    /** sealed is set to true once no more items can be added to the batch */
    sealed: boolean('sealed').notNull().default(false),
    sealedAt: timestamp('sealedAt', { precision: 3 }),
    /** this is the expected number of items in the batch */
    expectedCount: integer('expectedCount').notNull().default(0),
    /** this is the completed number of items in the batch. once this reaches expectedCount, and the batch is sealed, the batch is considered completed */
    completedCount: integer('completedCount').notNull().default(0),
    completedAt: timestamp('completedAt', { precision: 3 }),
    resumedAt: timestamp('resumedAt', { precision: 3 }),
    /** this is used to be able to "seal" this BatchTaskRun when all of the runs have been triggered asynchronously, and using the "parallel" processing strategy */
    processingJobsCount: integer('processingJobsCount').notNull().default(0),
    processingJobsExpectedCount: integer('processingJobsExpectedCount').notNull().default(0),
    /** optional token that can be used to authenticate the task run */
    oneTimeUseToken: text('oneTimeUseToken'),
    /** When processing started (status changed to PROCESSING) */
    processingStartedAt: timestamp('processingStartedAt', { precision: 3 }),
    /** When processing completed (all items processed) */
    processingCompletedAt: timestamp('processingCompletedAt', { precision: 3 }),
    /** Count of successfully created runs */
    successfulRunCount: integer('successfulRunCount'),
    /** Count of failed run creations */
    failedRunCount: integer('failedRunCount'),
    taskIdentifier: text('taskIdentifier'),
    checkpointEventId: text('checkpointEventId'),
    dependentTaskAttemptId: text('dependentTaskAttemptId'),
  },
  (table): PgTableExtraConfigValue[] => [
    uniqueIndex('BatchTaskRun_friendlyId_key').on(table.friendlyId),
    uniqueIndex('BatchTaskRun_checkpointEventId_key').on(table.checkpointEventId),
    index('BatchTaskRun_dependentTaskAttemptId_idx').on(table.dependentTaskAttemptId),
    index('BatchTaskRun_runtimeEnvironmentId_id_idx').on(table.runtimeEnvironmentId, table.id),
    index('BatchTaskRun_runtimeEnvironmentId_createdAt_id_idx').on(
      table.runtimeEnvironmentId,
      table.createdAt,
      table.id,
    ),
    uniqueIndex('BatchTaskRun_oneTimeUseToken_key').on(table.oneTimeUseToken),
    uniqueIndex('BatchTaskRun_runtimeEnvironmentId_idempotencyKey_key').on(
      table.runtimeEnvironmentId,
      table.idempotencyKey,
    ),
    foreignKey({
      name: 'BatchTaskRun_checkpointEventId_fkey',
      columns: [table.checkpointEventId],
      foreignColumns: [checkpointRestoreEvent.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'BatchTaskRun_dependentTaskAttemptId_fkey',
      columns: [table.dependentTaskAttemptId],
      foreignColumns: [taskRunAttempt.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/** Used in engine V1 only */
export const batchTaskRunItem = pgTable(
  'BatchTaskRunItem',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    status: batchTaskRunItemStatusEnum('status').notNull().default('PENDING'),
    batchTaskRunId: text('batchTaskRunId').notNull(),
    taskRunId: text('taskRunId').notNull(),
    taskRunAttemptId: text('taskRunAttemptId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    completedAt: timestamp('completedAt', { precision: 3 }),
  },
  (table) => [
    index('idx_batchtaskrunitem_taskrunattempt').on(table.taskRunAttemptId),
    index('idx_batchtaskrunitem_taskrun').on(table.taskRunId),
    uniqueIndex('BatchTaskRunItem_batchTaskRunId_taskRunId_key').on(
      table.batchTaskRunId,
      table.taskRunId,
    ),
    foreignKey({
      name: 'BatchTaskRunItem_batchTaskRunId_fkey',
      columns: [table.batchTaskRunId],
      foreignColumns: [batchTaskRun.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'BatchTaskRunItem_taskRunId_fkey',
      columns: [table.taskRunId],
      foreignColumns: [taskRun.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'BatchTaskRunItem_taskRunAttemptId_fkey',
      columns: [table.taskRunAttemptId],
      foreignColumns: [taskRunAttempt.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

/** Track individual run creation failures in batch processing (Run Engine v2) */
export const batchTaskRunError = pgTable(
  'BatchTaskRunError',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    batchTaskRunId: text('batchTaskRunId').notNull(),
    /** Which item in the batch (0-based index) */
    index: integer('index').notNull(),
    /** The task identifier that was being triggered */
    taskIdentifier: text('taskIdentifier').notNull(),
    /** The payload that failed (JSON, may be truncated) */
    payload: text('payload'),
    /** The options that were used */
    options: jsonb('options'),
    /** Error message */
    error: text('error').notNull(),
    /** Error code if available */
    errorCode: text('errorCode'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('BatchTaskRunError_batchTaskRunId_idx').on(table.batchTaskRunId),
    uniqueIndex('BatchTaskRunError_batchTaskRunId_index_key').on(table.batchTaskRunId, table.index),
    foreignKey({
      name: 'BatchTaskRunError_batchTaskRunId_fkey',
      columns: [table.batchTaskRunId],
      foreignColumns: [batchTaskRun.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const environmentVariable = pgTable(
  'EnvironmentVariable',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    friendlyId: text('friendlyId').notNull(),
    key: text('key').notNull(),
    projectId: text('projectId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('EnvironmentVariable_friendlyId_key').on(table.friendlyId),
    uniqueIndex('EnvironmentVariable_projectId_key_key').on(table.projectId, table.key),
    foreignKey({
      name: 'EnvironmentVariable_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const environmentVariableValue = pgTable(
  'EnvironmentVariableValue',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    valueReferenceId: text('valueReferenceId'),
    variableId: text('variableId').notNull(),
    environmentId: text('environmentId').notNull(),
    /** If true, the value is secret and cannot be revealed */
    isSecret: boolean('isSecret').notNull().default(false),
    isProxyOnly: boolean('isProxyOnly').notNull().default(false),
    proxyConfig: jsonb('proxyConfig'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    version: integer('version').notNull().default(1),
    lastUpdatedBy: jsonb('lastUpdatedBy'),
  },
  (table) => [
    index('EnvironmentVariableValue_environmentId_idx').on(table.environmentId),
    index('EnvironmentVariableValue_valueReferenceId_idx').on(table.valueReferenceId),
    index('EnvironmentVariableValue_environmentId_updatedAt_variableId_idx').on(
      table.environmentId,
      table.updatedAt,
      table.variableId,
    ),
    uniqueIndex('EnvironmentVariableValue_variableId_environmentId_key').on(
      table.variableId,
      table.environmentId,
    ),
    foreignKey({
      name: 'EnvironmentVariableValue_valueReferenceId_fkey',
      columns: [table.valueReferenceId],
      foreignColumns: [secretReference.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'EnvironmentVariableValue_variableId_fkey',
      columns: [table.variableId],
      foreignColumns: [environmentVariable.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'EnvironmentVariableValue_environmentId_fkey',
      columns: [table.environmentId],
      foreignColumns: [runtimeEnvironment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * A reason the environment variables "Needs attention" page shows a key in one environment, which
 * the project's team chose to ignore. An ignored EMPTY goes when the value is next saved.
 */
export const environmentVariableIgnore = pgTable(
  'EnvironmentVariableIgnore',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    projectId: text('projectId').notNull(),
    environmentId: text('environmentId').notNull(),
    variableId: text('variableId').notNull(),
    reason: environmentVariableIgnoreReasonEnum('reason').notNull(),
    createdById: text('createdById'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('EnvironmentVariableIgnore_variableId_idx').on(table.variableId),
    index('EnvironmentVariableIgnore_projectId_idx').on(table.projectId),
    index('EnvironmentVariableIgnore_createdById_idx').on(table.createdById),
    uniqueIndex('EnvironmentVariableIgnore_environmentId_variableId_reason_key').on(
      table.environmentId,
      table.variableId,
      table.reason,
    ),
    foreignKey({
      name: 'EnvironmentVariableIgnore_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'EnvironmentVariableIgnore_environmentId_fkey',
      columns: [table.environmentId],
      foreignColumns: [runtimeEnvironment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'EnvironmentVariableIgnore_variableId_fkey',
      columns: [table.variableId],
      foreignColumns: [environmentVariable.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'EnvironmentVariableIgnore_createdById_fkey',
      columns: [table.createdById],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const checkpoint = pgTable(
  'Checkpoint',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    friendlyId: text('friendlyId').notNull(),
    type: checkpointTypeEnum('type').notNull(),
    location: text('location').notNull(),
    imageRef: text('imageRef').notNull(),
    reason: text('reason'),
    metadata: text('metadata'),
    runId: text('runId').notNull(),
    attemptId: text('attemptId').notNull(),
    attemptNumber: integer('attemptNumber'),
    projectId: text('projectId').notNull(),
    runtimeEnvironmentId: text('runtimeEnvironmentId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table): PgTableExtraConfigValue[] => [
    uniqueIndex('Checkpoint_friendlyId_key').on(table.friendlyId),
    index('Checkpoint_attemptId_idx').on(table.attemptId),
    index('Checkpoint_runId_idx').on(table.runId),
    foreignKey({
      name: 'Checkpoint_runId_fkey',
      columns: [table.runId],
      foreignColumns: [taskRun.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Checkpoint_attemptId_fkey',
      columns: [table.attemptId],
      foreignColumns: [taskRunAttempt.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Checkpoint_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Checkpoint_runtimeEnvironmentId_fkey',
      columns: [table.runtimeEnvironmentId],
      foreignColumns: [runtimeEnvironment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const checkpointRestoreEvent = pgTable(
  'CheckpointRestoreEvent',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    type: checkpointRestoreEventTypeEnum('type').notNull(),
    reason: text('reason'),
    metadata: text('metadata'),
    checkpointId: text('checkpointId').notNull(),
    runId: text('runId').notNull(),
    attemptId: text('attemptId').notNull(),
    projectId: text('projectId').notNull(),
    runtimeEnvironmentId: text('runtimeEnvironmentId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table): PgTableExtraConfigValue[] => [
    index('CheckpointRestoreEvent_checkpointId_idx').on(table.checkpointId),
    index('CheckpointRestoreEvent_runId_idx').on(table.runId),
    foreignKey({
      name: 'CheckpointRestoreEvent_checkpointId_fkey',
      columns: [table.checkpointId],
      foreignColumns: [checkpoint.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'CheckpointRestoreEvent_runId_fkey',
      columns: [table.runId],
      foreignColumns: [taskRun.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'CheckpointRestoreEvent_attemptId_fkey',
      columns: [table.attemptId],
      foreignColumns: [taskRunAttempt.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'CheckpointRestoreEvent_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'CheckpointRestoreEvent_runtimeEnvironmentId_fkey',
      columns: [table.runtimeEnvironmentId],
      foreignColumns: [runtimeEnvironment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const workerDeployment = pgTable(
  'WorkerDeployment',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    contentHash: text('contentHash').notNull(),
    friendlyId: text('friendlyId').notNull(),
    shortCode: text('shortCode').notNull(),
    version: text('version').notNull(),
    runtime: text('runtime'),
    runtimeVersion: text('runtimeVersion'),
    /** CLI version that initiated the deploy, stamped at initialization or worker upload */
    cliVersion: text('cliVersion'),
    imageReference: text('imageReference'),
    imagePlatform: text('imagePlatform').notNull().default('linux/amd64'),
    /** Key of the microVM template prebuilt from this deployment's image, set when the builder succeeds */
    templateKey: text('templateKey'),
    /** Object key of the full index snapshot in the deployment index store. Set only after the upload succeeds. */
    indexObjectKey: text('indexObjectKey'),
    /** sha256 hex of the snapshot's uncompressed JSON */
    indexSha256: text('indexSha256'),
    indexFormatVersion: integer('indexFormatVersion'),
    externalBuildData: jsonb('externalBuildData'),
    buildServerMetadata: jsonb('buildServerMetadata'),
    /**
     * Encrypted build-time env vars for pre-bundled (fromBundle) deploys, as an
     * EncryptedSecretValue envelope. Kept for the redeploy window, then cleared by a sweep.
     */
    buildEnvVars: jsonb('buildEnvVars'),
    status: workerDeploymentStatusEnum('status').notNull().default('PENDING'),
    type: workerDeploymentTypeEnum('type').notNull().default('V1'),
    projectId: text('projectId').notNull(),
    environmentId: text('environmentId').notNull(),
    workerId: text('workerId'),
    triggeredById: text('triggeredById'),
    triggeredVia: text('triggeredVia'),
    commitSHA: text('commitSHA'),
    externalId: text('externalId'),
    /** The kind of credential that started the deployment. Null when unknown. */
    initiatedByCredential: deploymentInitiatedByCredentialEnum('initiatedByCredential'),
    startedAt: timestamp('startedAt', { precision: 3 }),
    installedAt: timestamp('installedAt', { precision: 3 }),
    builtAt: timestamp('builtAt', { precision: 3 }),
    deployedAt: timestamp('deployedAt', { precision: 3 }),
    failedAt: timestamp('failedAt', { precision: 3 }),
    errorData: jsonb('errorData'),
    canceledAt: timestamp('canceledAt', { precision: 3 }),
    canceledReason: text('canceledReason'),
    /** The newer GitHub build that superseded this one while it was still queued */
    supersededById: text('supersededById'),
    git: jsonb('git'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('WorkerDeployment_friendlyId_key').on(table.friendlyId),
    uniqueIndex('WorkerDeployment_workerId_key').on(table.workerId),
    index('WorkerDeployment_commitSHA_idx').on(table.commitSHA),
    index('WorkerDeployment_environmentId_createdAt_idx').on(table.environmentId, table.createdAt),
    index('WorkerDeployment_environmentId_status_id_idx').on(
      table.environmentId,
      table.status,
      table.id,
    ),
    index('WorkerDeployment_environmentId_externalId_idx').on(
      table.environmentId,
      table.externalId,
    ),
    uniqueIndex('WorkerDeployment_projectId_shortCode_key').on(table.projectId, table.shortCode),
    uniqueIndex('WorkerDeployment_environmentId_version_key').on(
      table.environmentId,
      table.version,
    ),
    foreignKey({
      name: 'WorkerDeployment_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'WorkerDeployment_environmentId_fkey',
      columns: [table.environmentId],
      foreignColumns: [runtimeEnvironment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'WorkerDeployment_workerId_fkey',
      columns: [table.workerId],
      foreignColumns: [backgroundWorker.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'WorkerDeployment_triggeredById_fkey',
      columns: [table.triggeredById],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const workerDeploymentPromotion = pgTable(
  'WorkerDeploymentPromotion',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    /** This is the promotion label, e.g. "current" */
    label: text('label').notNull(),
    deploymentId: text('deploymentId').notNull(),
    environmentId: text('environmentId').notNull(),
  },
  (table) => [
    uniqueIndex('WorkerDeploymentPromotion_environmentId_label_key').on(
      table.environmentId,
      table.label,
    ),
    foreignKey({
      name: 'WorkerDeploymentPromotion_deploymentId_fkey',
      columns: [table.deploymentId],
      foreignColumns: [workerDeployment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'WorkerDeploymentPromotion_environmentId_fkey',
      columns: [table.environmentId],
      foreignColumns: [runtimeEnvironment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/** Schedules can be attached to tasks to trigger them on a schedule */
export const taskSchedule = pgTable(
  'TaskSchedule',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    type: scheduleTypeEnum('type').notNull().default('IMPERATIVE'),
    /** users see this as `id`. They start with schedule_ */
    friendlyId: text('friendlyId').notNull(),
    /** a reference to a task (not a foreign key because it's across versions) */
    taskIdentifier: text('taskIdentifier').notNull(),
    /** can be provided and we won't create another with the same key */
    deduplicationKey: text('deduplicationKey')
      .notNull()
      .$defaultFn(() => cuid()),
    userProvidedDeduplicationKey: boolean('userProvidedDeduplicationKey').notNull().default(false),
    /** the CRON pattern */
    generatorExpression: text('generatorExpression').notNull(),
    generatorDescription: text('generatorDescription').notNull().default(''),
    generatorType: scheduleGeneratorTypeEnum('generatorType').notNull().default('CRON'),
    /** These are IANA format string, or the default "UTC". E.g. "America/New_York" */
    timezone: text('timezone').notNull().default('UTC'),
    windowDurationSeconds: integer('windowDurationSeconds'),
    windowPercentage: integer('windowPercentage'),
    defaultWindowDurationSeconds: integer('defaultWindowDurationSeconds'),
    minimumWindowDurationSeconds: integer('minimumWindowDurationSeconds'),
    /** Can be provided by the user then accessed inside a run */
    externalId: text('externalId'),
    /** @deprecated stop writing 2026-04-30; reads moved out of code (UI now derives from cron's previous slot). Drop in follow-up. */
    lastRunTriggeredAt: timestamp('lastRunTriggeredAt', { precision: 3 }),
    projectId: text('projectId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    active: boolean('active').notNull().default(true),
  },
  (table) => [
    uniqueIndex('TaskSchedule_friendlyId_key').on(table.friendlyId),
    index('TaskSchedule_projectId_idx').on(table.projectId),
    index('TaskSchedule_projectId_createdAt_idx').on(table.projectId, table.createdAt),
    uniqueIndex('TaskSchedule_projectId_deduplicationKey_key').on(
      table.projectId,
      table.deduplicationKey,
    ),
    foreignKey({
      name: 'TaskSchedule_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/** An instance links a schedule with an environment */
export const taskScheduleInstance = pgTable(
  'TaskScheduleInstance',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    taskScheduleId: text('taskScheduleId').notNull(),
    environmentId: text('environmentId').notNull(),
    projectId: text('projectId').notNull(),
    schedulePhase: integer('schedulePhase'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    active: boolean('active').notNull().default(true),
    /** @deprecated stop writing 2026-04-30; engine derives from cron + exactScheduleTime. Drop in follow-up. */
    lastScheduledTimestamp: timestamp('lastScheduledTimestamp', { precision: 3 }),
    /** @deprecated stop writing 2026-04-30; engine derives from cron + now(). Drop in follow-up. */
    nextScheduledTimestamp: timestamp('nextScheduledTimestamp', { precision: 3 }),
  },
  (table) => [
    index('TaskScheduleInstance_environmentId_idx').on(table.environmentId),
    index('TaskScheduleInstance_projectId_active_idx').on(table.projectId, table.active),
    uniqueIndex('TaskScheduleInstance_taskScheduleId_environmentId_key').on(
      table.taskScheduleId,
      table.environmentId,
    ),
    foreignKey({
      name: 'TaskScheduleInstance_taskScheduleId_fkey',
      columns: [table.taskScheduleId],
      foreignColumns: [taskSchedule.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'TaskScheduleInstance_environmentId_fkey',
      columns: [table.environmentId],
      foreignColumns: [runtimeEnvironment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'TaskScheduleInstance_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const runtimeEnvironmentSession = pgTable(
  'RuntimeEnvironmentSession',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    ipAddress: text('ipAddress').notNull(),
    environmentId: text('environmentId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    disconnectedAt: timestamp('disconnectedAt', { precision: 3 }),
  },
  (table): PgTableExtraConfigValue[] => [
    foreignKey({
      name: 'RuntimeEnvironmentSession_environmentId_fkey',
      columns: [table.environmentId],
      foreignColumns: [runtimeEnvironment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const projectAlertChannel = pgTable(
  'ProjectAlertChannel',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    friendlyId: text('friendlyId').notNull(),
    /** can be provided and we won't create another with the same key */
    deduplicationKey: text('deduplicationKey')
      .notNull()
      .$defaultFn(() => cuid()),
    userProvidedDeduplicationKey: boolean('userProvidedDeduplicationKey').notNull().default(false),
    integrationId: text('integrationId'),
    enabled: boolean('enabled').notNull().default(true),
    type: projectAlertChannelTypeEnum('type').notNull(),
    name: text('name').notNull(),
    properties: jsonb('properties').notNull(),
    alertTypes: projectAlertTypeEnum('alertTypes').array(),
    environmentTypes: runtimeEnvironmentTypeEnum('environmentTypes')
      .array()
      .default(['STAGING', 'PRODUCTION']),
    errorAlertConfig: jsonb('errorAlertConfig'),
    projectId: text('projectId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('ProjectAlertChannel_friendlyId_key').on(table.friendlyId),
    uniqueIndex('ProjectAlertChannel_projectId_deduplicationKey_key').on(
      table.projectId,
      table.deduplicationKey,
    ),
    foreignKey({
      name: 'ProjectAlertChannel_integrationId_fkey',
      columns: [table.integrationId],
      foreignColumns: [organizationIntegration.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'ProjectAlertChannel_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const projectAlert = pgTable(
  'ProjectAlert',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    friendlyId: text('friendlyId').notNull(),
    projectId: text('projectId').notNull(),
    environmentId: text('environmentId').notNull(),
    channelId: text('channelId').notNull(),
    status: projectAlertStatusEnum('status').notNull().default('PENDING'),
    type: projectAlertTypeEnum('type').notNull(),
    taskRunAttemptId: text('taskRunAttemptId'),
    taskRunId: text('taskRunId'),
    workerDeploymentId: text('workerDeploymentId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('ProjectAlert_friendlyId_key').on(table.friendlyId),
    index('ProjectAlert_channelId_idx').on(table.channelId),
    foreignKey({
      name: 'ProjectAlert_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'ProjectAlert_environmentId_fkey',
      columns: [table.environmentId],
      foreignColumns: [runtimeEnvironment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'ProjectAlert_channelId_fkey',
      columns: [table.channelId],
      foreignColumns: [projectAlertChannel.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'ProjectAlert_workerDeploymentId_fkey',
      columns: [table.workerDeploymentId],
      foreignColumns: [workerDeployment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const projectAlertStorage = pgTable(
  'ProjectAlertStorage',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    projectId: text('projectId').notNull(),
    alertChannelId: text('alertChannelId').notNull(),
    alertType: projectAlertTypeEnum('alertType').notNull(),
    storageId: text('storageId').notNull(),
    storageData: jsonb('storageData').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('ProjectAlertStorage_alertChannelId_alertType_storageId_idx').on(
      table.alertChannelId,
      table.alertType,
      table.storageId,
    ),
    foreignKey({
      name: 'ProjectAlertStorage_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'ProjectAlertStorage_alertChannelId_fkey',
      columns: [table.alertChannelId],
      foreignColumns: [projectAlertChannel.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const organizationIntegration = pgTable(
  'OrganizationIntegration',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    friendlyId: text('friendlyId').notNull(),
    service: integrationServiceEnum('service').notNull(),
    /** Identifier for external, integration's organization (e.g. Vercel's team) */
    externalOrganizationId: text('externalOrganizationId'),
    integrationData: jsonb('integrationData').notNull(),
    tokenReferenceId: text('tokenReferenceId').notNull(),
    organizationId: text('organizationId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    deletedAt: timestamp('deletedAt', { precision: 3 }),
  },
  (table) => [
    uniqueIndex('OrganizationIntegration_friendlyId_key').on(table.friendlyId),
    index('OrganizationIntegration_externalOrganizationId_idx').on(table.externalOrganizationId),
    foreignKey({
      name: 'OrganizationIntegration_tokenReferenceId_fkey',
      columns: [table.tokenReferenceId],
      foreignColumns: [secretReference.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'OrganizationIntegration_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const organizationProjectIntegration = pgTable(
  'OrganizationProjectIntegration',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    organizationIntegrationId: text('organizationIntegrationId').notNull(),
    projectId: text('projectId').notNull(),
    /** Identifier for webhooks, for example Vercel's projectId */
    externalEntityId: text('externalEntityId').notNull(),
    /** Save useful data like config or external entity name */
    integrationData: jsonb('integrationData').notNull(),
    /** UserId who installed the integration */
    installedBy: text('installedBy'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    deletedAt: timestamp('deletedAt', { precision: 3 }),
  },
  (table) => [
    index('OrganizationProjectIntegration_projectId_idx').on(table.projectId),
    index('OrganizationProjectIntegration_projectId_organizationIntegr_idx').on(
      table.projectId,
      table.organizationIntegrationId,
    ),
    index('OrganizationProjectIntegration_externalEntityId_idx').on(table.externalEntityId),
    foreignKey({
      name: 'OrganizationProjectIntegration_organizationIntegrationId_fkey',
      columns: [table.organizationIntegrationId],
      foreignColumns: [organizationIntegration.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'OrganizationProjectIntegration_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/** Bulk actions, like canceling and replaying runs */
export const bulkActionGroup = pgTable(
  'BulkActionGroup',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    friendlyId: text('friendlyId').notNull(),
    projectId: text('projectId').notNull(),
    environmentId: text('environmentId'),
    type: bulkActionTypeEnum('type').notNull(),
    /** When the group is created it's pending. After we've processed all the items it's completed. This does not mean the associated runs are completed. */
    status: bulkActionStatusEnum('status').notNull().default('PENDING'),
    /** The query that will be executed to get the runs to process (if null, we are passing run ids directly) */
    queryName: text('queryName'),
    /** The params that will be passed to the query */
    params: jsonb('params'),
    /** Billing-limit dedupe key (grace hitAt or resolve job key). Null for dashboard bulk actions. */
    dedupeKey: text('dedupeKey'),
    /** The cursor that will be passed to the query (null for the first page) */
    cursor: jsonb('cursor'),
    /** The number of runs that have been processed successfully */
    successCount: integer('successCount').notNull().default(0),
    /** The number of runs that have been failed */
    failureCount: integer('failureCount').notNull().default(0),
    /** The total number of runs that will be processed */
    totalCount: integer('totalCount').notNull().default(0),
    completionNotification: bulkActionNotificationTypeEnum('completionNotification')
      .notNull()
      .default('NONE'),
    userId: text('userId'),
    /** The name of the bulk action */
    name: text('name'),
    /** The time the bulk action was completed */
    completedAt: timestamp('completedAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('BulkActionGroup_friendlyId_key').on(table.friendlyId),
    index('BulkActionGroup_environmentId_createdAt_idx').on(table.environmentId, table.createdAt),
    index('BulkActionGroup_environmentId_type_dedupeKey_idx').on(
      table.environmentId,
      table.type,
      table.dedupeKey,
    ),
    index('BulkActionGroup_environmentId_status_type_idx').on(
      table.environmentId,
      table.status,
      table.type,
    ),
    foreignKey({
      name: 'BulkActionGroup_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'BulkActionGroup_environmentId_fkey',
      columns: [table.environmentId],
      foreignColumns: [runtimeEnvironment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'BulkActionGroup_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const bulkActionItem = pgTable(
  'BulkActionItem',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    /** @deprecated not used in new BulkActions */
    friendlyId: text('friendlyId'),
    groupId: text('groupId').notNull(),
    type: bulkActionTypeEnum('type').notNull(),
    /** When the item is created it's pending. After we've processed the item it's completed. This does not mean the associated runs are completed. */
    status: bulkActionItemStatusEnum('status').notNull().default('PENDING'),
    sourceRunId: text('sourceRunId').notNull(),
    destinationRunId: text('destinationRunId'),
    error: text('error'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('BulkActionItem_groupId_createdAt_idx').on(table.groupId, table.createdAt),
    foreignKey({
      name: 'BulkActionItem_sourceRunId_fkey',
      columns: [table.sourceRunId],
      foreignColumns: [taskRun.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'BulkActionItem_destinationRunId_fkey',
      columns: [table.destinationRunId],
      foreignColumns: [taskRun.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const realtimeStreamChunk = pgTable(
  'RealtimeStreamChunk',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    key: text('key').notNull(),
    value: text('value').notNull(),
    sequence: integer('sequence').notNull(),
    runId: text('runId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('RealtimeStreamChunk_runId_idx').on(table.runId),
    index('RealtimeStreamChunk_createdAt_idx').on(table.createdAt),
  ],
)

/** This is the unified otel span/log model that will eventually be replaced by clickhouse */
export const taskEventPartitioned = pgTable(
  'TaskEventPartitioned',
  {
    id: text('id')
      .notNull()
      .$defaultFn(() => cuid()),
    /** This matches the span name for a trace event, or the log body for a log event */
    message: text('message').notNull(),
    traceId: text('traceId').notNull(),
    spanId: text('spanId').notNull(),
    parentId: text('parentId'),
    tracestate: text('tracestate'),
    isError: boolean('isError').notNull().default(false),
    isPartial: boolean('isPartial').notNull().default(false),
    isCancelled: boolean('isCancelled').notNull().default(false),
    serviceName: text('serviceName').notNull(),
    serviceNamespace: text('serviceNamespace').notNull(),
    level: taskEventLevelEnum('level').notNull().default('TRACE'),
    kind: taskEventKindEnum('kind').notNull().default('INTERNAL'),
    status: taskEventStatusEnum('status').notNull().default('UNSET'),
    links: jsonb('links'),
    events: jsonb('events'),
    /** This is the time the event started in nanoseconds since the epoch */
    startTime: bigint('startTime', { mode: 'bigint' }).notNull(),
    /** This is the duration of the event in nanoseconds */
    duration: bigint('duration', { mode: 'bigint' })
      .notNull()
      .default(sql`0`),
    attemptId: text('attemptId'),
    attemptNumber: integer('attemptNumber'),
    environmentId: text('environmentId').notNull(),
    environmentType: runtimeEnvironmentTypeEnum('environmentType').notNull(),
    organizationId: text('organizationId').notNull(),
    projectId: text('projectId').notNull(),
    projectRef: text('projectRef').notNull(),
    runId: text('runId').notNull(),
    runIsTest: boolean('runIsTest').notNull().default(false),
    idempotencyKey: text('idempotencyKey'),
    taskSlug: text('taskSlug').notNull(),
    taskPath: text('taskPath'),
    taskExportName: text('taskExportName'),
    workerId: text('workerId'),
    workerVersion: text('workerVersion'),
    queueId: text('queueId'),
    queueName: text('queueName'),
    batchId: text('batchId'),
    /** This represents all the span attributes available, like http.status_code, and special attributes like $style.icon, $output, $metadata.payload.userId, as it's used for searching and filtering */
    properties: jsonb('properties').notNull(),
    /** This represents all span attributes in the $metadata namespace, like $metadata.payload */
    metadata: jsonb('metadata'),
    /** This represents all span attributes in the $style namespace, like $style */
    style: jsonb('style'),
    /** This represents all span attributes in the $output namespace, like $output */
    output: jsonb('output'),
    /** This represents the mimetype of the output, such as application/json or application/super+json */
    outputType: text('outputType'),
    payload: jsonb('payload'),
    payloadType: text('payloadType'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    usageDurationMs: integer('usageDurationMs').notNull().default(0),
    usageCostInCents: doublePrecision('usageCostInCents').notNull().default(0),
    machinePreset: text('machinePreset'),
    machinePresetCpu: doublePrecision('machinePresetCpu'),
    machinePresetMemory: doublePrecision('machinePresetMemory'),
    machinePresetCentsPerMs: doublePrecision('machinePresetCentsPerMs'),
  },
  (table) => [
    primaryKey({ name: 'TaskEventPartitioned_pkey', columns: [table.id, table.createdAt] }),
    index('TaskEventPartitioned_traceId_idx').on(table.traceId),
    index('TaskEventPartitioned_spanId_idx').on(table.spanId),
    index('TaskEventPartitioned_runId_idx').on(table.runId),
  ],
)

export const githubAppInstallation = pgTable(
  'GithubAppInstallation',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    appInstallationId: bigint('appInstallationId', { mode: 'bigint' }).notNull(),
    targetId: bigint('targetId', { mode: 'bigint' }).notNull(),
    targetType: text('targetType').notNull(),
    accountHandle: text('accountHandle').notNull(),
    permissions: jsonb('permissions'),
    repositorySelection: githubRepositorySelectionEnum('repositorySelection').notNull(),
    installedBy: text('installedBy'),
    installedByUserId: text('installedByUserId'),
    organizationId: text('organizationId').notNull(),
    deletedAt: timestamp('deletedAt', { precision: 3 }),
    suspendedAt: timestamp('suspendedAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('GithubAppInstallation_appInstallationId_key').on(table.appInstallationId),
    index('GithubAppInstallation_organizationId_idx').on(table.organizationId),
    foreignKey({
      name: 'GithubAppInstallation_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

export const githubRepository = pgTable(
  'GithubRepository',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    githubId: bigint('githubId', { mode: 'bigint' }).notNull(),
    name: text('name').notNull(),
    fullName: text('fullName').notNull(),
    htmlUrl: text('htmlUrl').notNull(),
    private: boolean('private').notNull(),
    defaultBranch: text('defaultBranch').notNull(),
    installationId: text('installationId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('GithubRepository_installationId_idx').on(table.installationId),
    uniqueIndex('GithubRepository_installationId_githubId_key').on(
      table.installationId,
      table.githubId,
    ),
    foreignKey({
      name: 'GithubRepository_installationId_fkey',
      columns: [table.installationId],
      foreignColumns: [githubAppInstallation.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const trustedDeployRepository = pgTable(
  'TrustedDeployRepository',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    projectId: text('projectId').notNull(),
    issuer: text('issuer').notNull().default('https://token.actions.githubusercontent.com'),
    repositoryOwner: text('repositoryOwner').notNull(),
    repositoryName: text('repositoryName').notNull(),
    repositoryId: text('repositoryId'),
    repositoryOwnerId: text('repositoryOwnerId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('TrustedDeployRepository_projectId_key').on(table.projectId),
    foreignKey({
      name: 'TrustedDeployRepository_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const trustedDeployEnvironment = pgTable(
  'TrustedDeployEnvironment',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    trustedDeployRepositoryId: text('trustedDeployRepositoryId').notNull(),
    runtimeEnvironmentId: text('runtimeEnvironmentId').notNull(),
    enabled: boolean('enabled').notNull().default(false),
    scope: trustedDeployScopeEnum('scope').notNull(),
    scopeValue: text('scopeValue'),
    additionalClaims: jsonb('additionalClaims').notNull().default({}),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('TrustedDeployEnvironment_runtimeEnvironmentId_key').on(table.runtimeEnvironmentId),
    index('TrustedDeployEnvironment_trustedDeployRepositoryId_idx').on(
      table.trustedDeployRepositoryId,
    ),
    foreignKey({
      name: 'TrustedDeployEnvironment_trustedDeployRepositoryId_fkey',
      columns: [table.trustedDeployRepositoryId],
      foreignColumns: [trustedDeployRepository.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'TrustedDeployEnvironment_runtimeEnvironmentId_fkey',
      columns: [table.runtimeEnvironmentId],
      foreignColumns: [runtimeEnvironment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const connectedGithubRepository = pgTable(
  'ConnectedGithubRepository',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    projectId: text('projectId').notNull(),
    repositoryId: text('repositoryId').notNull(),
    branchTracking: jsonb('branchTracking').notNull(),
    previewDeploymentsEnabled: boolean('previewDeploymentsEnabled').notNull().default(true),
    /** Skip queued builds once a newer commit is queued for the same environment or preview branch */
    supersedeProductionBuilds: boolean('supersedeProductionBuilds').notNull().default(false),
    supersedeStagingBuilds: boolean('supersedeStagingBuilds').notNull().default(true),
    supersedePreviewBuilds: boolean('supersedePreviewBuilds').notNull().default(true),
    buildFilters: jsonb('buildFilters'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('ConnectedGithubRepository_repositoryId_idx').on(table.repositoryId),
    uniqueIndex('ConnectedGithubRepository_projectId_key').on(table.projectId),
    foreignKey({
      name: 'ConnectedGithubRepository_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'ConnectedGithubRepository_repositoryId_fkey',
      columns: [table.repositoryId],
      foreignColumns: [githubRepository.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const impersonationAuditLog = pgTable(
  'ImpersonationAuditLog',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    action: impersonationAuditLogActionEnum('action').notNull(),
    adminId: text('adminId').notNull(),
    targetId: text('targetId').notNull(),
    ipAddress: text('ipAddress'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('ImpersonationAuditLog_adminId_idx').on(table.adminId),
    index('ImpersonationAuditLog_targetId_idx').on(table.targetId),
    index('ImpersonationAuditLog_createdAt_idx').on(table.createdAt),
    foreignKey({
      name: 'ImpersonationAuditLog_adminId_fkey',
      columns: [table.adminId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'ImpersonationAuditLog_targetId_fkey',
      columns: [table.targetId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const supportAccessRequest = pgTable(
  'SupportAccessRequest',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    organizationId: text('organizationId').notNull(),
    requestedById: text('requestedById'),
    reason: text('reason').notNull(),
    status: supportAccessRequestStatusEnum('status').notNull().default('PENDING'),
    approvedById: text('approvedById'),
    approvedAt: timestamp('approvedAt', { precision: 3 }),
    expiresAt: timestamp('expiresAt', { precision: 3 }),
    cancelledById: text('cancelledById'),
    cancelledAt: timestamp('cancelledAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('SupportAccessRequest_organizationId_status_expiresAt_idx').on(
      table.organizationId,
      table.status,
      table.expiresAt,
    ),
    index('SupportAccessRequest_requestedById_idx').on(table.requestedById),
    index('SupportAccessRequest_approvedById_idx').on(table.approvedById),
    foreignKey({
      name: 'SupportAccessRequest_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'SupportAccessRequest_requestedById_fkey',
      columns: [table.requestedById],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'SupportAccessRequest_approvedById_fkey',
      columns: [table.approvedById],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const customerQuery = pgTable(
  'CustomerQuery',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    /** The TSQL query text that was executed */
    query: text('query').notNull(),
    /** The scope of the query (determines which tenant IDs were used) */
    scope: customerQueryScopeEnum('scope').notNull(),
    /** Query execution statistics from ClickHouse */
    stats: jsonb('stats').notNull(),
    /** AI-generated title summarizing the query */
    title: text('title'),
    /** Where the query originated from */
    source: customerQuerySourceEnum('source').notNull().default('DASHBOARD'),
    organizationId: text('organizationId').notNull(),
    projectId: text('projectId'),
    environmentId: text('environmentId'),
    userId: text('userId'),
    /**
     * Time filter settings used when the query was executed
     * Period like "7d", "24h", etc.
     */
    filterPeriod: text('filterPeriod'),
    /** Custom start date (ISO 8601) */
    filterFrom: timestamp('filterFrom', { precision: 3 }),
    /** Custom end date (ISO 8601) */
    filterTo: timestamp('filterTo', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('CustomerQuery_organizationId_createdAt_idx').on(table.organizationId, table.createdAt),
    index('CustomerQuery_createdAt_idx').on(table.createdAt),
    foreignKey({
      name: 'CustomerQuery_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'CustomerQuery_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'CustomerQuery_environmentId_fkey',
      columns: [table.environmentId],
      foreignColumns: [runtimeEnvironment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'CustomerQuery_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const integrationDeployment = pgTable(
  'IntegrationDeployment',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    /** For example Vercel */
    integrationName: text('integrationName').notNull(),
    /** External ID */
    integrationDeploymentId: text('integrationDeploymentId').notNull(),
    commitSHA: text('commitSHA').notNull(),
    deploymentId: text('deploymentId'),
    /** External deployment status */
    status: text('status'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('IntegrationDeployment_commitSHA_idx').on(table.commitSHA),
    index('IntegrationDeployment_deploymentId_idx').on(table.deploymentId),
    foreignKey({
      name: 'IntegrationDeployment_deploymentId_fkey',
      columns: [table.deploymentId],
      foreignColumns: [workerDeployment.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

/** A user-defined metrics dashboard */
export const metricsDashboard = pgTable(
  'MetricsDashboard',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    friendlyId: text('friendlyId').notNull(),
    title: text('title').notNull(),
    description: text('description').notNull().default(''),
    organizationId: text('organizationId').notNull(),
    projectId: text('projectId'),
    ownerId: text('ownerId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    /**
     * JSON that defines the config, queries, layout and config of all widgets.
     * There will be a version field for the format.
     */
    layout: text('layout').notNull(),
  },
  (table) => [
    uniqueIndex('MetricsDashboard_friendlyId_key').on(table.friendlyId),
    index('MetricsDashboard_projectId_createdAt_idx').on(table.projectId, table.createdAt),
    foreignKey({
      name: 'MetricsDashboard_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'MetricsDashboard_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'MetricsDashboard_ownerId_fkey',
      columns: [table.ownerId],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

/** A known LLM model or model pattern for cost tracking */
export const llmModels = pgTable(
  'llm_models',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    friendlyId: text('friendly_id').notNull(),
    projectId: text('project_id'),
    modelName: text('model_name').notNull(),
    matchPattern: text('match_pattern').notNull(),
    startDate: timestamp('start_date', { precision: 3 }),
    source: text('source').notNull().default('default'),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    provider: text('provider'),
    description: text('description'),
    contextWindow: integer('context_window'),
    maxOutputTokens: integer('max_output_tokens'),
    capabilities: text('capabilities').array().default([]),
    isHidden: boolean('is_hidden').notNull().default(false),
    baseModelName: text('base_model_name'),
    pricingUnit: text('pricing_unit'),
  },
  (table) => [
    uniqueIndex('llm_models_friendly_id_key').on(table.friendlyId),
    index('llm_models_project_id_idx').on(table.projectId),
    uniqueIndex('llm_models_project_id_model_name_start_date_key').on(
      table.projectId,
      table.modelName,
      table.startDate,
    ),
    foreignKey({
      name: 'llm_models_project_id_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/** A pricing tier for a model (supports volume-based or conditional pricing) */
export const llmPricingTiers = pgTable(
  'llm_pricing_tiers',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    modelId: text('model_id').notNull(),
    name: text('name').notNull(),
    isDefault: boolean('is_default').notNull().default(true),
    priority: integer('priority').notNull().default(0),
    conditions: jsonb('conditions').notNull().default([]),
  },
  (table) => [
    uniqueIndex('llm_pricing_tiers_model_id_priority_key').on(table.modelId, table.priority),
    uniqueIndex('llm_pricing_tiers_model_id_name_key').on(table.modelId, table.name),
    foreignKey({
      name: 'llm_pricing_tiers_model_id_fkey',
      columns: [table.modelId],
      foreignColumns: [llmModels.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/** A price point for a usage type within a pricing tier */
export const llmPrices = pgTable(
  'llm_prices',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    modelId: text('model_id').notNull(),
    pricingTierId: text('pricing_tier_id').notNull(),
    usageType: text('usage_type').notNull(),
    price: numeric('price', { precision: 20, scale: 12 }).notNull(),
  },
  (table) => [
    uniqueIndex('llm_prices_model_id_usage_type_pricing_tier_id_key').on(
      table.modelId,
      table.usageType,
      table.pricingTierId,
    ),
    foreignKey({
      name: 'llm_prices_model_id_fkey',
      columns: [table.modelId],
      foreignColumns: [llmModels.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'llm_prices_pricing_tier_id_fkey',
      columns: [table.pricingTierId],
      foreignColumns: [llmPricingTiers.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/** Admin-created notification definitions */
export const platformNotification = pgTable(
  'PlatformNotification',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    friendlyId: text('friendlyId')
      .notNull()
      .$defaultFn(() => cuid()),
    /** Admin-facing title for identification in admin tools */
    title: text('title').notNull(),
    /** Versioned JSON rendering content (see payload schema in spec) */
    payload: jsonb('payload').notNull(),
    surface: platformNotificationSurfaceEnum('surface').notNull(),
    scope: platformNotificationScopeEnum('scope').notNull(),
    userId: text('userId'),
    organizationId: text('organizationId'),
    projectId: text('projectId'),
    /** When notification becomes active */
    startsAt: timestamp('startsAt', { precision: 3 }).notNull(),
    /** When notification expires */
    endsAt: timestamp('endsAt', { precision: 3 }).notNull(),
    /** CLI: auto-expire N days after user first saw it */
    cliMaxDaysAfterFirstSeen: integer('cliMaxDaysAfterFirstSeen'),
    /** CLI: auto-expire after shown N times to user */
    cliMaxShowCount: integer('cliMaxShowCount'),
    /** CLI: only show every N-th request (e.g. cliShowEvery=3 means show on 3rd, 6th, 9th…) */
    cliShowEvery: integer('cliShowEvery'),
    /** Ordering within same scope level (higher = more important) */
    priority: integer('priority').notNull().default(0),
    /**
     * Staged-but-unpublished. While true the notification is hidden from all
     * user-facing reads regardless of startsAt/endsAt; publishing sets real
     * dates and flips this to false.
     */
    isDraft: boolean('isDraft').notNull().default(false),
    /** Soft delete */
    archivedAt: timestamp('archivedAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('PlatformNotification_friendlyId_key').on(table.friendlyId),
    index('PlatformNotification_surface_scope_startsAt_idx').on(
      table.surface,
      table.scope,
      table.startsAt,
    ),
    foreignKey({
      name: 'PlatformNotification_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'PlatformNotification_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'PlatformNotification_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/** Per-user tracking of notification views and dismissals */
export const platformNotificationInteraction = pgTable(
  'PlatformNotificationInteraction',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    notificationId: text('notificationId').notNull(),
    userId: text('userId').notNull(),
    /** Set by beacon/CLI GET on first impression */
    firstSeenAt: timestamp('firstSeenAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    /** Times shown (incremented per beacon/CLI view) */
    showCount: integer('showCount').notNull().default(1),
    /** User dismissed in webapp */
    webappDismissedAt: timestamp('webappDismissedAt', { precision: 3 }),
    /** User clicked a link in webapp */
    webappClickedAt: timestamp('webappClickedAt', { precision: 3 }),
    /** Auto-dismissed or expired in CLI */
    cliDismissedAt: timestamp('cliDismissedAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('PlatformNotificationInteraction_notificationId_userId_key').on(
      table.notificationId,
      table.userId,
    ),
    foreignKey({
      name: 'PlatformNotificationInteraction_notificationId_fkey',
      columns: [table.notificationId],
      foreignColumns: [platformNotification.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'PlatformNotificationInteraction_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * *
 * * Error group state is used to track when a user has interacted with an error (ignored/resolved)
 * * The actual error data is in ClickHouse.
 */
export const errorGroupState = pgTable(
  'ErrorGroupState',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    organizationId: text('organizationId').notNull(),
    projectId: text('projectId').notNull(),
    environmentId: text('environmentId'),
    taskIdentifier: text('taskIdentifier').notNull(),
    errorFingerprint: text('errorFingerprint').notNull(),
    status: errorGroupStatusEnum('status').notNull().default('UNRESOLVED'),
    /**
     * *
     * * Error is ignored until this date
     */
    ignoredUntil: timestamp('ignoredUntil', { precision: 3 }),
    /**
     * *
     * * Error is ignored until this occurrence rate
     */
    ignoredUntilOccurrenceRate: integer('ignoredUntilOccurrenceRate'),
    /**
     * *
     * * Error is ignored until this total occurrences
     */
    ignoredUntilTotalOccurrences: integer('ignoredUntilTotalOccurrences'),
    /**
     * Total occurrence count at the time the error was ignored (from ClickHouse).
     * Used with ignoredUntilTotalOccurrences to compute occurrences since ignoring.
     */
    ignoredAtOccurrenceCount: bigint('ignoredAtOccurrenceCount', { mode: 'bigint' }),
    /**
     * *
     * * Error was ignored at this date
     */
    ignoredAt: timestamp('ignoredAt', { precision: 3 }),
    /**
     * *
     * * Reason for ignoring the error
     */
    ignoredReason: text('ignoredReason'),
    /**
     * *
     * * User who ignored the error
     */
    ignoredByUserId: text('ignoredByUserId'),
    /**
     * *
     * * Error was resolved at this date
     */
    resolvedAt: timestamp('resolvedAt', { precision: 3 }),
    /**
     * *
     * * Error was resolved in this version
     */
    resolvedInVersion: text('resolvedInVersion'),
    /**
     * *
     * * User who resolved the error
     */
    resolvedBy: text('resolvedBy'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('ErrorGroupState_environmentId_status_idx').on(table.environmentId, table.status),
    uniqueIndex('ErrorGroupState_environmentId_taskIdentifier_errorFingerpri_key').on(
      table.environmentId,
      table.taskIdentifier,
      table.errorFingerprint,
    ),
    foreignKey({
      name: 'ErrorGroupState_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'ErrorGroupState_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'ErrorGroupState_environmentId_fkey',
      columns: [table.environmentId],
      foreignColumns: [runtimeEnvironment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const taskIdentifier = pgTable(
  'TaskIdentifier',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    runtimeEnvironmentId: text('runtimeEnvironmentId').notNull(),
    projectId: text('projectId').notNull(),
    slug: text('slug').notNull(),
    currentTriggerSource: taskTriggerSourceEnum('currentTriggerSource')
      .notNull()
      .default('STANDARD'),
    currentWorkerId: text('currentWorkerId'),
    firstSeenAt: timestamp('firstSeenAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    lastSeenAt: timestamp('lastSeenAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    isInLatestDeployment: boolean('isInLatestDeployment').notNull().default(true),
  },
  (table) => [
    index('TaskIdentifier_runtimeEnvironmentId_isInLatestDeployment_idx').on(
      table.runtimeEnvironmentId,
      table.isInLatestDeployment,
    ),
    uniqueIndex('TaskIdentifier_runtimeEnvironmentId_slug_key').on(
      table.runtimeEnvironmentId,
      table.slug,
    ),
    foreignKey({
      name: 'TaskIdentifier_runtimeEnvironmentId_fkey',
      columns: [table.runtimeEnvironmentId],
      foreignColumns: [runtimeEnvironment.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'TaskIdentifier_projectId_fkey',
      columns: [table.projectId],
      foreignColumns: [project.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'TaskIdentifier_currentWorkerId_fkey',
      columns: [table.currentWorkerId],
      foreignColumns: [backgroundWorker.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

/**
 * Defines org-scoped data store overrides (e.g. dedicated ClickHouse for HIPAA orgs).
 * Multiple organizations can share a single data store row via organizationIds.
 */
export const organizationDataStore = pgTable(
  'OrganizationDataStore',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    /** Human-readable unique key (e.g. "hipaa-clickhouse-us-east") */
    key: text('key').notNull(),
    /** Organization IDs that use this data store */
    organizationIds: text('organizationIds').array(),
    kind: dataStoreKindEnum('kind').notNull(),
    /** Versioned config JSON. Structure is discriminated by the top-level `version` field. */
    config: jsonb('config').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('OrganizationDataStore_key_key').on(table.key),
    index('OrganizationDataStore_kind_idx').on(table.kind),
  ],
)

export const backgroundWorkerToBackgroundWorkerFile = pgTable(
  '_BackgroundWorkerToBackgroundWorkerFile',
  { A: text('A').notNull(), B: text('B').notNull() },
  (table) => [
    primaryKey({
      name: '_BackgroundWorkerToBackgroundWorkerFile_AB_pkey',
      columns: [table.A, table.B],
    }),
    index('_BackgroundWorkerToBackgroundWorkerFile_B_index').on(table.B),
    foreignKey({
      name: '_BackgroundWorkerToBackgroundWorkerFile_A_fkey',
      columns: [table.A],
      foreignColumns: [backgroundWorker.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: '_BackgroundWorkerToBackgroundWorkerFile_B_fkey',
      columns: [table.B],
      foreignColumns: [backgroundWorkerFile.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const backgroundWorkerToTaskQueue = pgTable(
  '_BackgroundWorkerToTaskQueue',
  { A: text('A').notNull(), B: text('B').notNull() },
  (table) => [
    primaryKey({ name: '_BackgroundWorkerToTaskQueue_AB_pkey', columns: [table.A, table.B] }),
    index('_BackgroundWorkerToTaskQueue_B_index').on(table.B),
    foreignKey({
      name: '_BackgroundWorkerToTaskQueue_A_fkey',
      columns: [table.A],
      foreignColumns: [backgroundWorker.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: '_BackgroundWorkerToTaskQueue_B_fkey',
      columns: [table.B],
      foreignColumns: [taskQueue.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const waitpointRunConnections = pgTable(
  '_WaitpointRunConnections',
  { A: text('A').notNull(), B: text('B').notNull() },
  (table) => [
    primaryKey({ name: '_WaitpointRunConnections_AB_pkey', columns: [table.A, table.B] }),
    index('_WaitpointRunConnections_B_index').on(table.B),
    foreignKey({
      name: '_WaitpointRunConnections_A_fkey',
      columns: [table.A],
      foreignColumns: [taskRun.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: '_WaitpointRunConnections_B_fkey',
      columns: [table.B],
      foreignColumns: [waitpoint.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const completedWaitpoints = pgTable(
  '_completedWaitpoints',
  { A: text('A').notNull(), B: text('B').notNull() },
  (table) => [
    primaryKey({ name: '_completedWaitpoints_AB_pkey', columns: [table.A, table.B] }),
    index('_completedWaitpoints_B_index').on(table.B),
    foreignKey({
      name: '_completedWaitpoints_A_fkey',
      columns: [table.A],
      foreignColumns: [taskRunExecutionSnapshot.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: '_completedWaitpoints_B_fkey',
      columns: [table.B],
      foreignColumns: [waitpoint.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const relations = defineRelations(
  {
    user,
    mfaBackupCode,
    invitationCode,
    authorizationCode,
    personalAccessToken,
    oAuthClient,
    oAuthRefreshToken,
    organizationAccessToken,
    organization,
    orgMember,
    orgMemberInvite,
    runtimeEnvironment,
    apiKeys,
    revokedApiKey,
    project,
    environmentEncryptionPolicy,
    environmentEncryptionKey,
    secretReference,
    secretStore,
    oidcSigningKey,
    dataMigration,
    backgroundWorker,
    backgroundWorkerFile,
    prompts,
    promptVersions,
    backgroundWorkerTask,
    webhookEndpoint,
    webhookDelivery,
    alert,
    alertEvent,
    alertOccurrenceSeen,
    alertDelivery,
    playgroundConversation,
    session,
    sessionRun,
    taskRun,
    taskRunTemplate,
    taskRunExecutionSnapshot,
    taskRunCheckpoint,
    waitpoint,
    taskRunWaitpoint,
    waitpointTag,
    featureFlag,
    workerInstance,
    workerInstanceGroup,
    workerGroupToken,
    taskRunTag,
    taskRunDependency,
    taskRunCounter,
    taskRunNumberCounter,
    taskRunAttempt,
    taskEvent,
    taskQueue,
    batchTaskRun,
    batchTaskRunItem,
    batchTaskRunError,
    environmentVariable,
    environmentVariableValue,
    environmentVariableIgnore,
    checkpoint,
    checkpointRestoreEvent,
    workerDeployment,
    workerDeploymentPromotion,
    taskSchedule,
    taskScheduleInstance,
    runtimeEnvironmentSession,
    projectAlertChannel,
    projectAlert,
    projectAlertStorage,
    organizationIntegration,
    organizationProjectIntegration,
    bulkActionGroup,
    bulkActionItem,
    realtimeStreamChunk,
    taskEventPartitioned,
    githubAppInstallation,
    githubRepository,
    trustedDeployRepository,
    trustedDeployEnvironment,
    connectedGithubRepository,
    impersonationAuditLog,
    supportAccessRequest,
    customerQuery,
    integrationDeployment,
    metricsDashboard,
    llmModels,
    llmPricingTiers,
    llmPrices,
    platformNotification,
    platformNotificationInteraction,
    errorGroupState,
    taskIdentifier,
    organizationDataStore,
    backgroundWorkerToBackgroundWorkerFile,
    backgroundWorkerToTaskQueue,
    waitpointRunConnections,
    completedWaitpoints,
  },
  (r) => ({
    user: {
      orgMemberships: r.many.orgMember(),
      sentInvites: r.many.orgMemberInvite(),
      mfaSecretReference: r.one.secretReference({
        from: r.user.mfaSecretReferenceId,
        to: r.secretReference.id,
      }),
      invitationCode: r.one.invitationCode({
        from: r.user.invitationCodeId,
        to: r.invitationCode.id,
      }),
      personalAccessTokens: r.many.personalAccessToken(),
      createdApiKeys: r.many.apiKeys(),
      deployments: r.many.workerDeployment(),
      backupCodes: r.many.mfaBackupCode(),
      bulkActions: r.many.bulkActionGroup(),
      impersonationsPerformed: r.many.impersonationAuditLog({ alias: 'ImpersonationAdmin' }),
      impersonationsReceived: r.many.impersonationAuditLog({ alias: 'ImpersonationTarget' }),
      supportAccessRequestsMade: r.many.supportAccessRequest({ alias: 'SupportAccessRequester' }),
      supportAccessRequestsApproved: r.many.supportAccessRequest({
        alias: 'SupportAccessApprover',
      }),
      customerQueries: r.many.customerQuery(),
      metricsDashboards: r.many.metricsDashboard(),
      platformNotifications: r.many.platformNotification(),
      platformNotificationInteractions: r.many.platformNotificationInteraction(),
      environmentVariableIgnores: r.many.environmentVariableIgnore(),
    },
    mfaBackupCode: {
      user: r.one.user({ from: r.mfaBackupCode.userId, to: r.user.id, optional: false }),
    },
    invitationCode: { users: r.many.user() },
    authorizationCode: {
      personalAccessToken: r.one.personalAccessToken({
        from: r.authorizationCode.personalAccessTokenId,
        to: r.personalAccessToken.id,
      }),
      oauthClient: r.one.oAuthClient({
        from: r.authorizationCode.oauthClientId,
        to: r.oAuthClient.id,
      }),
    },
    personalAccessToken: {
      user: r.one.user({ from: r.personalAccessToken.userId, to: r.user.id, optional: false }),
      authorizationCodes: r.many.authorizationCode(),
      oauthClient: r.one.oAuthClient({
        from: r.personalAccessToken.oauthClientId,
        to: r.oAuthClient.id,
      }),
      oauthRefreshTokens: r.many.oAuthRefreshToken(),
    },
    oAuthClient: {
      authorizationCodes: r.many.authorizationCode(),
      personalAccessTokens: r.many.personalAccessToken(),
      refreshTokens: r.many.oAuthRefreshToken(),
    },
    oAuthRefreshToken: {
      personalAccessToken: r.one.personalAccessToken({
        from: r.oAuthRefreshToken.personalAccessTokenId,
        to: r.personalAccessToken.id,
        optional: false,
      }),
      oauthClient: r.one.oAuthClient({
        from: r.oAuthRefreshToken.oauthClientId,
        to: r.oAuthClient.id,
        optional: false,
      }),
    },
    organizationAccessToken: {
      organization: r.one.organization({
        from: r.organizationAccessToken.organizationId,
        to: r.organization.id,
        optional: false,
      }),
    },
    organization: {
      environments: r.many.runtimeEnvironment(),
      projects: r.many.project(),
      members: r.many.orgMember(),
      invites: r.many.orgMemberInvite(),
      organizationIntegrations: r.many.organizationIntegration(),
      organizationAccessTokens: r.many.organizationAccessToken(),
      workerGroups: r.many.workerInstanceGroup(),
      workerInstances: r.many.workerInstance(),
      githubAppInstallations: r.many.githubAppInstallation(),
      customerQueries: r.many.customerQuery(),
      metricsDashboards: r.many.metricsDashboard(),
      prompts: r.many.prompts(),
      platformNotifications: r.many.platformNotification(),
      errorGroupStates: r.many.errorGroupState(),
      supportAccessRequests: r.many.supportAccessRequest(),
    },
    orgMember: {
      organization: r.one.organization({
        from: r.orgMember.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      user: r.one.user({ from: r.orgMember.userId, to: r.user.id, optional: false }),
      environments: r.many.runtimeEnvironment(),
    },
    orgMemberInvite: {
      organization: r.one.organization({
        from: r.orgMemberInvite.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      inviter: r.one.user({ from: r.orgMemberInvite.inviterId, to: r.user.id, optional: false }),
    },
    runtimeEnvironment: {
      encryptionPolicy: r.one.environmentEncryptionPolicy({
        from: r.runtimeEnvironment.id,
        to: r.environmentEncryptionPolicy.environmentId,
      }),
      encryptionKeys: r.many.environmentEncryptionKey(),
      parentEnvironment: r.one.runtimeEnvironment({
        from: r.runtimeEnvironment.parentEnvironmentId,
        to: r.runtimeEnvironment.id,
        alias: 'parentEnvironment',
      }),
      childEnvironments: r.many.runtimeEnvironment({ alias: 'parentEnvironment' }),
      organization: r.one.organization({
        from: r.runtimeEnvironment.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      project: r.one.project({
        from: r.runtimeEnvironment.projectId,
        to: r.project.id,
        optional: false,
      }),
      orgMember: r.one.orgMember({ from: r.runtimeEnvironment.orgMemberId, to: r.orgMember.id }),
      backgroundWorkers: r.many.backgroundWorker(),
      backgroundWorkerTasks: r.many.backgroundWorkerTask(),
      taskRuns: r.many.taskRun(),
      taskQueues: r.many.taskQueue(),
      environmentVariableValues: r.many.environmentVariableValue(),
      environmentVariableIgnores: r.many.environmentVariableIgnore(),
      checkpoints: r.many.checkpoint(),
      workerDeployments: r.many.workerDeployment(),
      workerDeploymentPromotions: r.many.workerDeploymentPromotion(),
      taskRunAttempts: r.many.taskRunAttempt(),
      CheckpointRestoreEvent: r.many.checkpointRestoreEvent(),
      taskScheduleInstances: r.many.taskScheduleInstance(),
      alerts: r.many.projectAlert(),
      sessions: r.many.runtimeEnvironmentSession({
        alias: 'RuntimeEnvironmentToRuntimeEnvironmentSession',
      }),
      currentSession: r.one.runtimeEnvironmentSession({
        from: r.runtimeEnvironment.currentSessionId,
        to: r.runtimeEnvironmentSession.id,
        alias: 'currentSession',
      }),
      taskRunNumberCounter: r.many.taskRunNumberCounter(),
      taskRunCheckpoints: r.many.taskRunCheckpoint(),
      waitpoints: r.many.waitpoint(),
      workerInstances: r.many.workerInstance(),
      waitpointTags: r.many.waitpointTag(),
      BulkActionGroup: r.many.bulkActionGroup(),
      customerQueries: r.many.customerQuery(),
      prompts: r.many.prompts(),
      playgroundConversations: r.many.playgroundConversation(),
      errorGroupStates: r.many.errorGroupState(),
      taskIdentifiers: r.many.taskIdentifier(),
      apiKeys: r.many.apiKeys(),
      revokedApiKeys: r.many.revokedApiKey(),
      trustedDeployGrant: r.one.trustedDeployEnvironment({
        from: r.runtimeEnvironment.id,
        to: r.trustedDeployEnvironment.runtimeEnvironmentId,
      }),
    },
    apiKeys: {
      runtimeEnvironment: r.one.runtimeEnvironment({
        from: r.apiKeys.runtimeEnvironmentId,
        to: r.runtimeEnvironment.id,
        optional: false,
      }),
      createdBy: r.one.user({ from: r.apiKeys.createdByUserId, to: r.user.id }),
    },
    revokedApiKey: {
      runtimeEnvironment: r.one.runtimeEnvironment({
        from: r.revokedApiKey.runtimeEnvironmentId,
        to: r.runtimeEnvironment.id,
        optional: false,
      }),
    },
    project: {
      organization: r.one.organization({
        from: r.project.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      workerGroups: r.many.workerInstanceGroup({ alias: 'ProjectToWorkerInstanceGroup' }),
      workers: r.many.workerInstance(),
      defaultWorkerGroup: r.one.workerInstanceGroup({
        from: r.project.defaultWorkerGroupId,
        to: r.workerInstanceGroup.id,
        alias: 'ProjectDefaultWorkerGroup',
      }),
      environments: r.many.runtimeEnvironment(),
      backgroundWorkers: r.many.backgroundWorker(),
      backgroundWorkerTasks: r.many.backgroundWorkerTask(),
      taskRuns: r.many.taskRun(),
      runTags: r.many.taskRunTag(),
      taskQueues: r.many.taskQueue(),
      environmentVariables: r.many.environmentVariable(),
      environmentVariableIgnores: r.many.environmentVariableIgnore(),
      checkpoints: r.many.checkpoint(),
      WorkerDeployment: r.many.workerDeployment(),
      CheckpointRestoreEvent: r.many.checkpointRestoreEvent(),
      taskSchedules: r.many.taskSchedule(),
      alertChannels: r.many.projectAlertChannel(),
      alerts: r.many.projectAlert(),
      alertStorages: r.many.projectAlertStorage(),
      bulkActionGroups: r.many.bulkActionGroup(),
      BackgroundWorkerFile: r.many.backgroundWorkerFile(),
      waitpoints: r.many.waitpoint(),
      taskRunWaitpoints: r.many.taskRunWaitpoint(),
      taskRunCheckpoints: r.many.taskRunCheckpoint(),
      waitpointTags: r.many.waitpointTag(),
      connectedGithubRepository: r.one.connectedGithubRepository({
        from: r.project.id,
        to: r.connectedGithubRepository.projectId,
      }),
      trustedDeployRepository: r.one.trustedDeployRepository({
        from: r.project.id,
        to: r.trustedDeployRepository.projectId,
      }),
      organizationProjectIntegration: r.many.organizationProjectIntegration(),
      customerQueries: r.many.customerQuery(),
      playgroundConversations: r.many.playgroundConversation(),
      taskScheduleInstances: r.many.taskScheduleInstance(),
      metricsDashboards: r.many.metricsDashboard(),
      llmModels: r.many.llmModels(),
      prompts: r.many.prompts(),
      platformNotifications: r.many.platformNotification(),
      errorGroupStates: r.many.errorGroupState(),
      taskIdentifiers: r.many.taskIdentifier(),
    },
    environmentEncryptionPolicy: {
      environment: r.one.runtimeEnvironment({
        from: r.environmentEncryptionPolicy.environmentId,
        to: r.runtimeEnvironment.id,
        optional: false,
      }),
      activeKey: r.one.environmentEncryptionKey({
        from: [
          r.environmentEncryptionPolicy.activeKeyId,
          r.environmentEncryptionPolicy.environmentId,
        ],
        to: [r.environmentEncryptionKey.id, r.environmentEncryptionKey.environmentId],
      }),
    },
    environmentEncryptionKey: {
      environment: r.one.runtimeEnvironment({
        from: r.environmentEncryptionKey.environmentId,
        to: r.runtimeEnvironment.id,
        optional: false,
      }),
      activeFor: r.many.environmentEncryptionPolicy(),
    },
    secretReference: {
      environmentVariableValues: r.many.environmentVariableValue(),
      OrganizationIntegration: r.many.organizationIntegration(),
      User: r.many.user(),
    },
    backgroundWorker: {
      project: r.one.project({
        from: r.backgroundWorker.projectId,
        to: r.project.id,
        optional: false,
      }),
      runtimeEnvironment: r.one.runtimeEnvironment({
        from: r.backgroundWorker.runtimeEnvironmentId,
        to: r.runtimeEnvironment.id,
        optional: false,
      }),
      tasks: r.many.backgroundWorkerTask(),
      attempts: r.many.taskRunAttempt(),
      lockedRuns: r.many.taskRun(),
      files: r.many.backgroundWorkerFile({
        from: r.backgroundWorker.id.through(r.backgroundWorkerToBackgroundWorkerFile.A),
        to: r.backgroundWorkerFile.id.through(r.backgroundWorkerToBackgroundWorkerFile.B),
      }),
      queues: r.many.taskQueue({
        from: r.backgroundWorker.id.through(r.backgroundWorkerToTaskQueue.A),
        to: r.taskQueue.id.through(r.backgroundWorkerToTaskQueue.B),
      }),
      promptVersions: r.many.promptVersions(),
      deployment: r.one.workerDeployment({
        from: r.backgroundWorker.id,
        to: r.workerDeployment.workerId,
      }),
      workerGroup: r.one.workerInstanceGroup({
        from: r.backgroundWorker.workerGroupId,
        to: r.workerInstanceGroup.id,
      }),
      taskIdentifiers: r.many.taskIdentifier(),
    },
    backgroundWorkerFile: {
      project: r.one.project({
        from: r.backgroundWorkerFile.projectId,
        to: r.project.id,
        optional: false,
      }),
      backgroundWorkers: r.many.backgroundWorker({
        from: r.backgroundWorkerFile.id.through(r.backgroundWorkerToBackgroundWorkerFile.B),
        to: r.backgroundWorker.id.through(r.backgroundWorkerToBackgroundWorkerFile.A),
      }),
      tasks: r.many.backgroundWorkerTask(),
    },
    prompts: {
      organization: r.one.organization({
        from: r.prompts.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      project: r.one.project({ from: r.prompts.projectId, to: r.project.id, optional: false }),
      runtimeEnvironment: r.one.runtimeEnvironment({
        from: r.prompts.runtimeEnvironmentId,
        to: r.runtimeEnvironment.id,
        optional: false,
      }),
      versions: r.many.promptVersions(),
    },
    promptVersions: {
      prompt: r.one.prompts({ from: r.promptVersions.promptId, to: r.prompts.id, optional: false }),
      worker: r.one.backgroundWorker({
        from: r.promptVersions.workerId,
        to: r.backgroundWorker.id,
      }),
    },
    backgroundWorkerTask: {
      worker: r.one.backgroundWorker({
        from: r.backgroundWorkerTask.workerId,
        to: r.backgroundWorker.id,
        optional: false,
      }),
      project: r.one.project({
        from: r.backgroundWorkerTask.projectId,
        to: r.project.id,
        optional: false,
      }),
      file: r.one.backgroundWorkerFile({
        from: r.backgroundWorkerTask.fileId,
        to: r.backgroundWorkerFile.id,
      }),
      runtimeEnvironment: r.one.runtimeEnvironment({
        from: r.backgroundWorkerTask.runtimeEnvironmentId,
        to: r.runtimeEnvironment.id,
        optional: false,
      }),
      attempts: r.many.taskRunAttempt(),
      runs: r.many.taskRun(),
      queue: r.one.taskQueue({ from: r.backgroundWorkerTask.queueId, to: r.taskQueue.id }),
    },
    playgroundConversation: {
      run: r.one.taskRun({ from: r.playgroundConversation.runId, to: r.taskRun.id }),
      project: r.one.project({
        from: r.playgroundConversation.projectId,
        to: r.project.id,
        optional: false,
      }),
      runtimeEnvironment: r.one.runtimeEnvironment({
        from: r.playgroundConversation.runtimeEnvironmentId,
        to: r.runtimeEnvironment.id,
        optional: false,
      }),
    },
    session: { runs: r.many.sessionRun() },
    sessionRun: {
      session: r.one.session({ from: r.sessionRun.sessionId, to: r.session.id, optional: false }),
    },
    taskRun: {
      runtimeEnvironment: r.one.runtimeEnvironment({
        from: r.taskRun.runtimeEnvironmentId,
        to: r.runtimeEnvironment.id,
        optional: false,
      }),
      project: r.one.project({ from: r.taskRun.projectId, to: r.project.id, optional: false }),
      attempts: r.many.taskRunAttempt({ alias: 'attempts' }),
      checkpoints: r.many.checkpoint(),
      lockedBy: r.one.backgroundWorkerTask({
        from: r.taskRun.lockedById,
        to: r.backgroundWorkerTask.id,
      }),
      lockedToVersion: r.one.backgroundWorker({
        from: r.taskRun.lockedToVersionId,
        to: r.backgroundWorker.id,
      }),
      associatedWaitpoint: r.one.waitpoint({
        from: r.taskRun.id,
        to: r.waitpoint.completedByTaskRunId,
        alias: 'CompletingRun',
      }),
      blockedByWaitpoints: r.many.taskRunWaitpoint(),
      connectedWaitpoints: r.many.waitpoint({
        from: r.taskRun.id.through(r.waitpointRunConnections.A),
        to: r.waitpoint.id.through(r.waitpointRunConnections.B),
      }),
      batchItems: r.many.batchTaskRunItem(),
      dependency: r.one.taskRunDependency({
        from: r.taskRun.id,
        to: r.taskRunDependency.taskRunId,
      }),
      CheckpointRestoreEvent: r.many.checkpointRestoreEvent(),
      executionSnapshots: r.many.taskRunExecutionSnapshot(),
      sourceBulkActionItems: r.many.bulkActionItem({ alias: 'SourceActionItemRun' }),
      destinationBulkActionItems: r.many.bulkActionItem({ alias: 'DestinationActionItemRun' }),
      rootTaskRun: r.one.taskRun({
        from: r.taskRun.rootTaskRunId,
        to: r.taskRun.id,
        alias: 'TaskRootRun',
      }),
      descendantRuns: r.many.taskRun({ alias: 'TaskRootRun' }),
      parentTaskRun: r.one.taskRun({
        from: r.taskRun.parentTaskRunId,
        to: r.taskRun.id,
        alias: 'TaskParentRun',
      }),
      childRuns: r.many.taskRun({ alias: 'TaskParentRun' }),
      parentTaskRunAttempt: r.one.taskRunAttempt({
        from: r.taskRun.parentTaskRunAttemptId,
        to: r.taskRunAttempt.id,
        alias: 'TaskParentRunAttempt',
      }),
      batch: r.one.batchTaskRun({ from: r.taskRun.batchId, to: r.batchTaskRun.id }),
      playgroundConversations: r.many.playgroundConversation(),
    },
    taskRunExecutionSnapshot: {
      run: r.one.taskRun({
        from: r.taskRunExecutionSnapshot.runId,
        to: r.taskRun.id,
        optional: false,
      }),
      completedWaitpoints: r.many.waitpoint({
        from: r.taskRunExecutionSnapshot.id.through(r.completedWaitpoints.A),
        to: r.waitpoint.id.through(r.completedWaitpoints.B),
      }),
      checkpoint: r.one.taskRunCheckpoint({
        from: r.taskRunExecutionSnapshot.checkpointId,
        to: r.taskRunCheckpoint.id,
      }),
    },
    taskRunCheckpoint: {
      project: r.one.project({
        from: r.taskRunCheckpoint.projectId,
        to: r.project.id,
        optional: false,
      }),
      runtimeEnvironment: r.one.runtimeEnvironment({
        from: r.taskRunCheckpoint.runtimeEnvironmentId,
        to: r.runtimeEnvironment.id,
        optional: false,
      }),
      executionSnapshot: r.many.taskRunExecutionSnapshot(),
    },
    waitpoint: {
      completedByTaskRun: r.one.taskRun({
        from: r.waitpoint.completedByTaskRunId,
        to: r.taskRun.id,
        alias: 'CompletingRun',
      }),
      completedByBatch: r.one.batchTaskRun({
        from: r.waitpoint.completedByBatchId,
        to: r.batchTaskRun.id,
      }),
      blockingTaskRuns: r.many.taskRunWaitpoint(),
      connectedRuns: r.many.taskRun({
        from: r.waitpoint.id.through(r.waitpointRunConnections.B),
        to: r.taskRun.id.through(r.waitpointRunConnections.A),
      }),
      completedExecutionSnapshots: r.many.taskRunExecutionSnapshot({
        from: r.waitpoint.id.through(r.completedWaitpoints.B),
        to: r.taskRunExecutionSnapshot.id.through(r.completedWaitpoints.A),
      }),
      project: r.one.project({ from: r.waitpoint.projectId, to: r.project.id, optional: false }),
      environment: r.one.runtimeEnvironment({
        from: r.waitpoint.environmentId,
        to: r.runtimeEnvironment.id,
        optional: false,
      }),
    },
    taskRunWaitpoint: {
      taskRun: r.one.taskRun({
        from: r.taskRunWaitpoint.taskRunId,
        to: r.taskRun.id,
        optional: false,
      }),
      waitpoint: r.one.waitpoint({
        from: r.taskRunWaitpoint.waitpointId,
        to: r.waitpoint.id,
        optional: false,
      }),
      project: r.one.project({
        from: r.taskRunWaitpoint.projectId,
        to: r.project.id,
        optional: false,
      }),
      batch: r.one.batchTaskRun({ from: r.taskRunWaitpoint.batchId, to: r.batchTaskRun.id }),
    },
    waitpointTag: {
      environment: r.one.runtimeEnvironment({
        from: r.waitpointTag.environmentId,
        to: r.runtimeEnvironment.id,
        optional: false,
      }),
      project: r.one.project({ from: r.waitpointTag.projectId, to: r.project.id, optional: false }),
    },
    workerInstance: {
      workerGroup: r.one.workerInstanceGroup({
        from: r.workerInstance.workerGroupId,
        to: r.workerInstanceGroup.id,
        optional: false,
      }),
      organization: r.one.organization({
        from: r.workerInstance.organizationId,
        to: r.organization.id,
      }),
      project: r.one.project({ from: r.workerInstance.projectId, to: r.project.id }),
      environment: r.one.runtimeEnvironment({
        from: r.workerInstance.environmentId,
        to: r.runtimeEnvironment.id,
      }),
      deployment: r.one.workerDeployment({
        from: r.workerInstance.deploymentId,
        to: r.workerDeployment.id,
      }),
    },
    workerInstanceGroup: {
      token: r.one.workerGroupToken({
        from: r.workerInstanceGroup.tokenId,
        to: r.workerGroupToken.id,
        optional: false,
        alias: 'WorkerGroupPrimaryToken',
      }),
      childTokens: r.many.workerGroupToken({ alias: 'WorkerGroupChildTokens' }),
      workers: r.many.workerInstance(),
      backgroundWorkers: r.many.backgroundWorker(),
      defaultForProjects: r.many.project({ alias: 'ProjectDefaultWorkerGroup' }),
      organization: r.one.organization({
        from: r.workerInstanceGroup.organizationId,
        to: r.organization.id,
      }),
      project: r.one.project({
        from: r.workerInstanceGroup.projectId,
        to: r.project.id,
        alias: 'ProjectToWorkerInstanceGroup',
      }),
    },
    workerGroupToken: {
      workerGroup: r.one.workerInstanceGroup({
        from: r.workerGroupToken.id,
        to: r.workerInstanceGroup.tokenId,
        alias: 'WorkerGroupPrimaryToken',
      }),
      childOf: r.one.workerInstanceGroup({
        from: r.workerGroupToken.workerGroupId,
        to: r.workerInstanceGroup.id,
        alias: 'WorkerGroupChildTokens',
      }),
    },
    taskRunTag: {
      project: r.one.project({ from: r.taskRunTag.projectId, to: r.project.id, optional: false }),
    },
    taskRunDependency: {
      taskRun: r.one.taskRun({
        from: r.taskRunDependency.taskRunId,
        to: r.taskRun.id,
        optional: false,
      }),
      checkpointEvent: r.one.checkpointRestoreEvent({
        from: r.taskRunDependency.checkpointEventId,
        to: r.checkpointRestoreEvent.id,
      }),
      dependentAttempt: r.one.taskRunAttempt({
        from: r.taskRunDependency.dependentAttemptId,
        to: r.taskRunAttempt.id,
      }),
      dependentBatchRun: r.one.batchTaskRun({
        from: r.taskRunDependency.dependentBatchRunId,
        to: r.batchTaskRun.id,
      }),
    },
    taskRunNumberCounter: {
      environment: r.one.runtimeEnvironment({
        from: r.taskRunNumberCounter.environmentId,
        to: r.runtimeEnvironment.id,
        optional: false,
      }),
    },
    taskRunAttempt: {
      taskRun: r.one.taskRun({
        from: r.taskRunAttempt.taskRunId,
        to: r.taskRun.id,
        optional: false,
        alias: 'attempts',
      }),
      backgroundWorker: r.one.backgroundWorker({
        from: r.taskRunAttempt.backgroundWorkerId,
        to: r.backgroundWorker.id,
        optional: false,
      }),
      backgroundWorkerTask: r.one.backgroundWorkerTask({
        from: r.taskRunAttempt.backgroundWorkerTaskId,
        to: r.backgroundWorkerTask.id,
        optional: false,
      }),
      runtimeEnvironment: r.one.runtimeEnvironment({
        from: r.taskRunAttempt.runtimeEnvironmentId,
        to: r.runtimeEnvironment.id,
        optional: false,
      }),
      queue: r.one.taskQueue({
        from: r.taskRunAttempt.queueId,
        to: r.taskQueue.id,
        optional: false,
      }),
      dependencies: r.many.taskRunDependency(),
      batchDependencies: r.many.batchTaskRun(),
      checkpoints: r.many.checkpoint(),
      batchTaskRunItems: r.many.batchTaskRunItem(),
      CheckpointRestoreEvent: r.many.checkpointRestoreEvent(),
      childRuns: r.many.taskRun({ alias: 'TaskParentRunAttempt' }),
    },
    taskQueue: {
      project: r.one.project({ from: r.taskQueue.projectId, to: r.project.id, optional: false }),
      runtimeEnvironment: r.one.runtimeEnvironment({
        from: r.taskQueue.runtimeEnvironmentId,
        to: r.runtimeEnvironment.id,
        optional: false,
      }),
      attempts: r.many.taskRunAttempt(),
      tasks: r.many.backgroundWorkerTask(),
      workers: r.many.backgroundWorker({
        from: r.taskQueue.id.through(r.backgroundWorkerToTaskQueue.B),
        to: r.backgroundWorker.id.through(r.backgroundWorkerToTaskQueue.A),
      }),
    },
    batchTaskRun: {
      runs: r.many.taskRun(),
      runsBlocked: r.many.taskRunWaitpoint(),
      waitpoints: r.many.waitpoint(),
      errors: r.many.batchTaskRunError(),
      items: r.many.batchTaskRunItem(),
      checkpointEvent: r.one.checkpointRestoreEvent({
        from: r.batchTaskRun.checkpointEventId,
        to: r.checkpointRestoreEvent.id,
      }),
      dependentTaskAttempt: r.one.taskRunAttempt({
        from: r.batchTaskRun.dependentTaskAttemptId,
        to: r.taskRunAttempt.id,
      }),
      runDependencies: r.many.taskRunDependency(),
    },
    batchTaskRunItem: {
      batchTaskRun: r.one.batchTaskRun({
        from: r.batchTaskRunItem.batchTaskRunId,
        to: r.batchTaskRun.id,
        optional: false,
      }),
      taskRun: r.one.taskRun({
        from: r.batchTaskRunItem.taskRunId,
        to: r.taskRun.id,
        optional: false,
      }),
      taskRunAttempt: r.one.taskRunAttempt({
        from: r.batchTaskRunItem.taskRunAttemptId,
        to: r.taskRunAttempt.id,
      }),
    },
    batchTaskRunError: {
      batchTaskRun: r.one.batchTaskRun({
        from: r.batchTaskRunError.batchTaskRunId,
        to: r.batchTaskRun.id,
        optional: false,
      }),
    },
    environmentVariable: {
      project: r.one.project({
        from: r.environmentVariable.projectId,
        to: r.project.id,
        optional: false,
      }),
      values: r.many.environmentVariableValue(),
      ignores: r.many.environmentVariableIgnore(),
    },
    environmentVariableValue: {
      valueReference: r.one.secretReference({
        from: r.environmentVariableValue.valueReferenceId,
        to: r.secretReference.id,
      }),
      variable: r.one.environmentVariable({
        from: r.environmentVariableValue.variableId,
        to: r.environmentVariable.id,
        optional: false,
      }),
      environment: r.one.runtimeEnvironment({
        from: r.environmentVariableValue.environmentId,
        to: r.runtimeEnvironment.id,
        optional: false,
      }),
    },
    environmentVariableIgnore: {
      project: r.one.project({
        from: r.environmentVariableIgnore.projectId,
        to: r.project.id,
        optional: false,
      }),
      environment: r.one.runtimeEnvironment({
        from: r.environmentVariableIgnore.environmentId,
        to: r.runtimeEnvironment.id,
        optional: false,
      }),
      variable: r.one.environmentVariable({
        from: r.environmentVariableIgnore.variableId,
        to: r.environmentVariable.id,
        optional: false,
      }),
      createdBy: r.one.user({ from: r.environmentVariableIgnore.createdById, to: r.user.id }),
    },
    checkpoint: {
      events: r.many.checkpointRestoreEvent(),
      run: r.one.taskRun({ from: r.checkpoint.runId, to: r.taskRun.id, optional: false }),
      attempt: r.one.taskRunAttempt({
        from: r.checkpoint.attemptId,
        to: r.taskRunAttempt.id,
        optional: false,
      }),
      project: r.one.project({ from: r.checkpoint.projectId, to: r.project.id, optional: false }),
      runtimeEnvironment: r.one.runtimeEnvironment({
        from: r.checkpoint.runtimeEnvironmentId,
        to: r.runtimeEnvironment.id,
        optional: false,
      }),
    },
    checkpointRestoreEvent: {
      checkpoint: r.one.checkpoint({
        from: r.checkpointRestoreEvent.checkpointId,
        to: r.checkpoint.id,
        optional: false,
      }),
      run: r.one.taskRun({
        from: r.checkpointRestoreEvent.runId,
        to: r.taskRun.id,
        optional: false,
      }),
      attempt: r.one.taskRunAttempt({
        from: r.checkpointRestoreEvent.attemptId,
        to: r.taskRunAttempt.id,
        optional: false,
      }),
      project: r.one.project({
        from: r.checkpointRestoreEvent.projectId,
        to: r.project.id,
        optional: false,
      }),
      runtimeEnvironment: r.one.runtimeEnvironment({
        from: r.checkpointRestoreEvent.runtimeEnvironmentId,
        to: r.runtimeEnvironment.id,
        optional: false,
      }),
      taskRunDependency: r.one.taskRunDependency({
        from: r.checkpointRestoreEvent.id,
        to: r.taskRunDependency.checkpointEventId,
      }),
      batchTaskRunDependency: r.one.batchTaskRun({
        from: r.checkpointRestoreEvent.id,
        to: r.batchTaskRun.checkpointEventId,
      }),
    },
    workerDeployment: {
      project: r.one.project({
        from: r.workerDeployment.projectId,
        to: r.project.id,
        optional: false,
      }),
      environment: r.one.runtimeEnvironment({
        from: r.workerDeployment.environmentId,
        to: r.runtimeEnvironment.id,
        optional: false,
      }),
      worker: r.one.backgroundWorker({
        from: r.workerDeployment.workerId,
        to: r.backgroundWorker.id,
      }),
      triggeredBy: r.one.user({ from: r.workerDeployment.triggeredById, to: r.user.id }),
      promotions: r.many.workerDeploymentPromotion(),
      alerts: r.many.projectAlert(),
      workerInstance: r.many.workerInstance(),
      integrationDeployments: r.many.integrationDeployment(),
    },
    workerDeploymentPromotion: {
      deployment: r.one.workerDeployment({
        from: r.workerDeploymentPromotion.deploymentId,
        to: r.workerDeployment.id,
        optional: false,
      }),
      environment: r.one.runtimeEnvironment({
        from: r.workerDeploymentPromotion.environmentId,
        to: r.runtimeEnvironment.id,
        optional: false,
      }),
    },
    taskSchedule: {
      instances: r.many.taskScheduleInstance(),
      project: r.one.project({ from: r.taskSchedule.projectId, to: r.project.id, optional: false }),
    },
    taskScheduleInstance: {
      taskSchedule: r.one.taskSchedule({
        from: r.taskScheduleInstance.taskScheduleId,
        to: r.taskSchedule.id,
        optional: false,
      }),
      environment: r.one.runtimeEnvironment({
        from: r.taskScheduleInstance.environmentId,
        to: r.runtimeEnvironment.id,
        optional: false,
      }),
      project: r.one.project({
        from: r.taskScheduleInstance.projectId,
        to: r.project.id,
        optional: false,
      }),
    },
    runtimeEnvironmentSession: {
      environment: r.one.runtimeEnvironment({
        from: r.runtimeEnvironmentSession.environmentId,
        to: r.runtimeEnvironment.id,
        optional: false,
        alias: 'RuntimeEnvironmentToRuntimeEnvironmentSession',
      }),
      currentEnvironments: r.many.runtimeEnvironment({ alias: 'currentSession' }),
    },
    projectAlertChannel: {
      integration: r.one.organizationIntegration({
        from: r.projectAlertChannel.integrationId,
        to: r.organizationIntegration.id,
      }),
      project: r.one.project({
        from: r.projectAlertChannel.projectId,
        to: r.project.id,
        optional: false,
      }),
      alerts: r.many.projectAlert(),
      alertStorages: r.many.projectAlertStorage(),
    },
    projectAlert: {
      project: r.one.project({ from: r.projectAlert.projectId, to: r.project.id, optional: false }),
      environment: r.one.runtimeEnvironment({
        from: r.projectAlert.environmentId,
        to: r.runtimeEnvironment.id,
        optional: false,
      }),
      channel: r.one.projectAlertChannel({
        from: r.projectAlert.channelId,
        to: r.projectAlertChannel.id,
        optional: false,
      }),
      workerDeployment: r.one.workerDeployment({
        from: r.projectAlert.workerDeploymentId,
        to: r.workerDeployment.id,
      }),
    },
    projectAlertStorage: {
      project: r.one.project({
        from: r.projectAlertStorage.projectId,
        to: r.project.id,
        optional: false,
      }),
      alertChannel: r.one.projectAlertChannel({
        from: r.projectAlertStorage.alertChannelId,
        to: r.projectAlertChannel.id,
        optional: false,
      }),
    },
    organizationIntegration: {
      tokenReference: r.one.secretReference({
        from: r.organizationIntegration.tokenReferenceId,
        to: r.secretReference.id,
        optional: false,
      }),
      organization: r.one.organization({
        from: r.organizationIntegration.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      alertChannels: r.many.projectAlertChannel(),
      organizationProjectIntegration: r.many.organizationProjectIntegration(),
    },
    organizationProjectIntegration: {
      organizationIntegration: r.one.organizationIntegration({
        from: r.organizationProjectIntegration.organizationIntegrationId,
        to: r.organizationIntegration.id,
        optional: false,
      }),
      project: r.one.project({
        from: r.organizationProjectIntegration.projectId,
        to: r.project.id,
        optional: false,
      }),
    },
    bulkActionGroup: {
      project: r.one.project({
        from: r.bulkActionGroup.projectId,
        to: r.project.id,
        optional: false,
      }),
      environment: r.one.runtimeEnvironment({
        from: r.bulkActionGroup.environmentId,
        to: r.runtimeEnvironment.id,
      }),
      user: r.one.user({ from: r.bulkActionGroup.userId, to: r.user.id }),
    },
    bulkActionItem: {
      sourceRun: r.one.taskRun({
        from: r.bulkActionItem.sourceRunId,
        to: r.taskRun.id,
        optional: false,
        alias: 'SourceActionItemRun',
      }),
      destinationRun: r.one.taskRun({
        from: r.bulkActionItem.destinationRunId,
        to: r.taskRun.id,
        alias: 'DestinationActionItemRun',
      }),
    },
    githubAppInstallation: {
      organization: r.one.organization({
        from: r.githubAppInstallation.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      repositories: r.many.githubRepository(),
    },
    githubRepository: {
      installation: r.one.githubAppInstallation({
        from: r.githubRepository.installationId,
        to: r.githubAppInstallation.id,
        optional: false,
      }),
      ConnectedGithubRepository: r.many.connectedGithubRepository(),
    },
    trustedDeployRepository: {
      project: r.one.project({
        from: r.trustedDeployRepository.projectId,
        to: r.project.id,
        optional: false,
      }),
      environments: r.many.trustedDeployEnvironment(),
    },
    trustedDeployEnvironment: {
      trustedDeployRepository: r.one.trustedDeployRepository({
        from: r.trustedDeployEnvironment.trustedDeployRepositoryId,
        to: r.trustedDeployRepository.id,
        optional: false,
      }),
      runtimeEnvironment: r.one.runtimeEnvironment({
        from: r.trustedDeployEnvironment.runtimeEnvironmentId,
        to: r.runtimeEnvironment.id,
        optional: false,
      }),
    },
    connectedGithubRepository: {
      project: r.one.project({
        from: r.connectedGithubRepository.projectId,
        to: r.project.id,
        optional: false,
      }),
      repository: r.one.githubRepository({
        from: r.connectedGithubRepository.repositoryId,
        to: r.githubRepository.id,
        optional: false,
      }),
    },
    impersonationAuditLog: {
      admin: r.one.user({
        from: r.impersonationAuditLog.adminId,
        to: r.user.id,
        optional: false,
        alias: 'ImpersonationAdmin',
      }),
      target: r.one.user({
        from: r.impersonationAuditLog.targetId,
        to: r.user.id,
        optional: false,
        alias: 'ImpersonationTarget',
      }),
    },
    supportAccessRequest: {
      organization: r.one.organization({
        from: r.supportAccessRequest.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      requestedBy: r.one.user({
        from: r.supportAccessRequest.requestedById,
        to: r.user.id,
        alias: 'SupportAccessRequester',
      }),
      approvedBy: r.one.user({
        from: r.supportAccessRequest.approvedById,
        to: r.user.id,
        alias: 'SupportAccessApprover',
      }),
    },
    customerQuery: {
      organization: r.one.organization({
        from: r.customerQuery.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      project: r.one.project({ from: r.customerQuery.projectId, to: r.project.id }),
      environment: r.one.runtimeEnvironment({
        from: r.customerQuery.environmentId,
        to: r.runtimeEnvironment.id,
      }),
      user: r.one.user({ from: r.customerQuery.userId, to: r.user.id }),
    },
    integrationDeployment: {
      workerDeployment: r.one.workerDeployment({
        from: r.integrationDeployment.deploymentId,
        to: r.workerDeployment.id,
      }),
    },
    metricsDashboard: {
      organization: r.one.organization({
        from: r.metricsDashboard.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      project: r.one.project({ from: r.metricsDashboard.projectId, to: r.project.id }),
      owner: r.one.user({ from: r.metricsDashboard.ownerId, to: r.user.id }),
    },
    llmModels: {
      project: r.one.project({ from: r.llmModels.projectId, to: r.project.id }),
      pricingTiers: r.many.llmPricingTiers(),
      prices: r.many.llmPrices(),
    },
    llmPricingTiers: {
      model: r.one.llmModels({
        from: r.llmPricingTiers.modelId,
        to: r.llmModels.id,
        optional: false,
      }),
      prices: r.many.llmPrices(),
    },
    llmPrices: {
      model: r.one.llmModels({ from: r.llmPrices.modelId, to: r.llmModels.id, optional: false }),
      pricingTier: r.one.llmPricingTiers({
        from: r.llmPrices.pricingTierId,
        to: r.llmPricingTiers.id,
        optional: false,
      }),
    },
    platformNotification: {
      user: r.one.user({ from: r.platformNotification.userId, to: r.user.id }),
      organization: r.one.organization({
        from: r.platformNotification.organizationId,
        to: r.organization.id,
      }),
      project: r.one.project({ from: r.platformNotification.projectId, to: r.project.id }),
      interactions: r.many.platformNotificationInteraction(),
    },
    platformNotificationInteraction: {
      notification: r.one.platformNotification({
        from: r.platformNotificationInteraction.notificationId,
        to: r.platformNotification.id,
        optional: false,
      }),
      user: r.one.user({
        from: r.platformNotificationInteraction.userId,
        to: r.user.id,
        optional: false,
      }),
    },
    errorGroupState: {
      organization: r.one.organization({
        from: r.errorGroupState.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      project: r.one.project({
        from: r.errorGroupState.projectId,
        to: r.project.id,
        optional: false,
      }),
      environment: r.one.runtimeEnvironment({
        from: r.errorGroupState.environmentId,
        to: r.runtimeEnvironment.id,
      }),
    },
    taskIdentifier: {
      runtimeEnvironment: r.one.runtimeEnvironment({
        from: r.taskIdentifier.runtimeEnvironmentId,
        to: r.runtimeEnvironment.id,
        optional: false,
      }),
      project: r.one.project({
        from: r.taskIdentifier.projectId,
        to: r.project.id,
        optional: false,
      }),
      currentWorker: r.one.backgroundWorker({
        from: r.taskIdentifier.currentWorkerId,
        to: r.backgroundWorker.id,
      }),
    },
    backgroundWorkerToBackgroundWorkerFile: {
      backgroundWorker: r.one.backgroundWorker({
        from: r.backgroundWorkerToBackgroundWorkerFile.A,
        to: r.backgroundWorker.id,
        optional: false,
      }),
      backgroundWorkerFile: r.one.backgroundWorkerFile({
        from: r.backgroundWorkerToBackgroundWorkerFile.B,
        to: r.backgroundWorkerFile.id,
        optional: false,
      }),
    },
    backgroundWorkerToTaskQueue: {
      backgroundWorker: r.one.backgroundWorker({
        from: r.backgroundWorkerToTaskQueue.A,
        to: r.backgroundWorker.id,
        optional: false,
      }),
      taskQueue: r.one.taskQueue({
        from: r.backgroundWorkerToTaskQueue.B,
        to: r.taskQueue.id,
        optional: false,
      }),
    },
    waitpointRunConnections: {
      taskRun: r.one.taskRun({
        from: r.waitpointRunConnections.A,
        to: r.taskRun.id,
        optional: false,
      }),
      waitpoint: r.one.waitpoint({
        from: r.waitpointRunConnections.B,
        to: r.waitpoint.id,
        optional: false,
      }),
    },
    completedWaitpoints: {
      taskRunExecutionSnapshot: r.one.taskRunExecutionSnapshot({
        from: r.completedWaitpoints.A,
        to: r.taskRunExecutionSnapshot.id,
        optional: false,
      }),
      waitpoint: r.one.waitpoint({
        from: r.completedWaitpoints.B,
        to: r.waitpoint.id,
        optional: false,
      }),
    },
  }),
)
