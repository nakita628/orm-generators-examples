import type { ColumnType } from 'kysely'

export type Generated<T> =
  T extends ColumnType<infer S, infer I, infer U>
    ? ColumnType<S, I | undefined, U>
    : ColumnType<T, T | undefined, T>

export type Timestamp = ColumnType<Date, Date | string, Date | string>

export type AuthenticationMethod = 'GITHUB' | 'MAGIC_LINK' | 'GOOGLE' | 'SSO'

export type OrganizationAccessTokenType = 'USER' | 'SYSTEM'

export type OrgMemberRole = 'ADMIN' | 'MEMBER'

export type ApiKeyIssuer = 'USER' | 'SYSTEM'

export type RuntimeEnvironmentType = 'PRODUCTION' | 'STAGING' | 'DEVELOPMENT' | 'PREVIEW'

export type EnvironmentPauseSource = 'BILLING_LIMIT' | 'ADMIN_BAN'

export type ProjectVersion = 'V2' | 'V3'

export type EnvironmentKeyWrappingProvider = 'PLATFORM_KEYRING' | 'AWS_KMS'

export type SecretStoreProvider = 'DATABASE' | 'AWS_PARAM_STORE'

export type OidcSigningKeyStatus = 'STAGED' | 'ACTIVE' | 'RETIRED'

/** Each purpose is a separate key set with its own issuer, rotation and publication. */
export type OidcSigningKeyPurpose = 'EXTERNAL_OIDC' | 'INTERNAL_RUN'

export type TaskTriggerSource = 'STANDARD' | 'SCHEDULED' | 'AGENT' | 'WEBHOOK'

export type WebhookEndpointStatus = 'ACTIVE' | 'INACTIVE' | 'DELETING'

export type WebhookDeliveryStatus =
  | 'PENDING'
  | 'PROCESSING'
  | 'SUCCEEDED'
  | 'FAILED'
  | 'FILTERED'
  | 'UNMATCHED'

export type TaskRunStatus =
  | 'DELAYED'
  | 'PENDING'
  | 'PENDING_VERSION'
  | 'WAITING_FOR_DEPLOY'
  | 'DEQUEUED'
  | 'EXECUTING'
  | 'WAITING_TO_RESUME'
  | 'RETRYING_AFTER_FAILURE'
  | 'PAUSED'
  | 'CANCELED'
  | 'INTERRUPTED'
  | 'COMPLETED_SUCCESSFULLY'
  | 'COMPLETED_WITH_ERRORS'
  | 'SYSTEM_FAILURE'
  | 'CRASHED'
  | 'EXPIRED'
  | 'TIMED_OUT'

export type RunEngineVersion = 'V1' | 'V2'

export type TaskRunExecutionStatus =
  | 'RUN_CREATED'
  | 'DELAYED'
  | 'QUEUED'
  | 'QUEUED_EXECUTING'
  | 'PENDING_EXECUTING'
  | 'EXECUTING'
  | 'EXECUTING_WITH_WAITPOINTS'
  | 'SUSPENDED'
  | 'PENDING_CANCEL'
  | 'FINISHED'

export type TaskRunCheckpointType = 'DOCKER' | 'KUBERNETES' | 'COMPUTE'

export type WaitpointType = 'RUN' | 'DATETIME' | 'MANUAL' | 'BATCH'

export type WaitpointStatus = 'PENDING' | 'COMPLETED'

export type WorkerInstanceGroupType = 'MANAGED' | 'UNMANAGED'

export type WorkloadType = 'CONTAINER' | 'MICROVM'

export type TaskRunAttemptStatus =
  | 'PENDING'
  | 'EXECUTING'
  | 'PAUSED'
  | 'FAILED'
  | 'CANCELED'
  | 'COMPLETED'

export type TaskEventLevel = 'TRACE' | 'DEBUG' | 'LOG' | 'INFO' | 'WARN' | 'ERROR'

export type TaskEventKind =
  | 'UNSPECIFIED'
  | 'INTERNAL'
  | 'SERVER'
  | 'CLIENT'
  | 'PRODUCER'
  | 'CONSUMER'
  | 'UNRECOGNIZED'
  | 'LOG'

export type TaskEventStatus = 'UNSET' | 'OK' | 'ERROR' | 'UNRECOGNIZED'

export type TaskQueueType = 'VIRTUAL' | 'NAMED'

export type TaskQueueVersion = 'V1' | 'V2'

export type TaskQueueConcurrencyVersion = 'V1' | 'V2'

export type TaskQueueRole = 'QUEUE' | 'LIMIT'

export type BatchTaskRunStatus =
  | 'PENDING'
  | 'PROCESSING'
  | 'COMPLETED'
  | 'PARTIAL_FAILED'
  | 'ABORTED'

export type BatchTaskRunItemStatus = 'PENDING' | 'FAILED' | 'CANCELED' | 'COMPLETED'

export type CheckpointType = 'DOCKER' | 'KUBERNETES'

export type CheckpointRestoreEventType = 'CHECKPOINT' | 'RESTORE'

export type WorkerDeploymentType = 'MANAGED' | 'UNMANAGED' | 'V1'

export type WorkerDeploymentStatus =
  | 'PENDING'
  | 'INSTALLING'
  | 'BUILDING'
  | 'DEPLOYING'
  | 'DEPLOYED'
  | 'FAILED'
  | 'CANCELED'
  | 'TIMED_OUT'

export type ScheduleType = 'DECLARATIVE' | 'IMPERATIVE'

export type ScheduleGeneratorType = 'CRON'

export type ProjectAlertChannelType = 'EMAIL' | 'SLACK' | 'WEBHOOK'

export type ProjectAlertType =
  | 'TASK_RUN'
  | 'TASK_RUN_ATTEMPT'
  | 'DEPLOYMENT_FAILURE'
  | 'DEPLOYMENT_SUCCESS'
  | 'ERROR_GROUP'
  | 'DASHBOARD_AGENT_WATCH'

export type ProjectAlertStatus = 'PENDING' | 'SENT' | 'FAILED'

export type IntegrationService = 'SLACK' | 'VERCEL'

export type BulkActionType = 'CANCEL' | 'REPLAY'

export type BulkActionStatus = 'PENDING' | 'COMPLETED' | 'ABORTED'

export type BulkActionNotificationType = 'NONE' | 'EMAIL'

export type BulkActionItemStatus = 'PENDING' | 'COMPLETED' | 'FAILED'

export type GithubRepositorySelection = 'ALL' | 'SELECTED'

export type TrustedDeployScope = 'ANY' | 'BRANCH' | 'TAG' | 'ENVIRONMENT'

export type ImpersonationAuditLogAction = 'START' | 'STOP'

export type SupportAccessMode = 'ALLOW' | 'REQUIRES_REQUEST'

export type SupportAccessRequestStatus = 'PENDING' | 'APPROVED' | 'CANCELLED'

export type CustomerQuerySource = 'DASHBOARD' | 'API'

export type CustomerQueryScope = 'ORGANIZATION' | 'PROJECT' | 'ENVIRONMENT'

export type PlatformNotificationSurface = 'WEBAPP' | 'CLI'

export type PlatformNotificationScope = 'USER' | 'PROJECT' | 'ORGANIZATION' | 'GLOBAL'

export type ErrorGroupStatus = 'UNRESOLVED' | 'RESOLVED' | 'IGNORED'

export type DataStoreKind = 'CLICKHOUSE'

export interface User {
  id: string
  email: string
  /** Address awaiting confirmation; becomes `email` only after a link sent to it is opened */
  pendingEmail: string | null
  authenticationMethod: AuthenticationMethod
  authenticationProfile: unknown
  authenticationExtraParams: unknown
  authIdentifier: string | null
  displayName: string | null
  name: string | null
  avatarUrl: string | null
  admin: Generated<boolean>
  /** Preferences for the dashboard */
  dashboardPreferences: unknown
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  /** @deprecated */
  isOnCloudWaitlist: Generated<boolean>
  /** @deprecated */
  featureCloud: Generated<boolean>
  /** @deprecated */
  isOnHostedRepoWaitlist: Generated<boolean>
  marketingEmails: Generated<boolean>
  confirmedBasicDetails: Generated<boolean>
  referralSource: string | null
  onboardingData: unknown
  mfaEnabledAt: Timestamp | null
  mfaSecretReferenceId: string | null
  /** Hash of the last used code to prevent replay attacks */
  mfaLastUsedCode: string | null
  /**
   * Maximum session lifetime in seconds for this user. Default = 1 year (Gregorian).
   * May be further restricted by Organization.maxSessionDuration on any of the user's orgs.
   */
  sessionDuration: Generated<number>
  /**
   * Absolute timestamp at which the user's current session must end. Written
   * at login (and any time the effective duration is recomputed) and shortened
   * when an admin lowers an org cap. Read on requireUser / getUser; null is
   * treated as "no enforced deadline" (legacy sessions from before this
   * shipped). Read-on-demand has zero per-request DB cost — the value
   * piggybacks on the User fetch that requireUser already does.
   */
  nextSessionEnd: Timestamp | null
  invitationCodeId: string | null
}

export interface MfaBackupCode {
  id: string
  /** Hash of the actual code */
  code: string
  userId: string
  usedAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface InvitationCode {
  id: string
  code: string
  createdAt: Generated<Timestamp>
}

/** Used to generate PersonalAccessTokens, they're one-time use */
export interface AuthorizationCode {
  id: string
  code: string
  personalAccessTokenId: string | null
  /** Client IP of the request that minted the code, compared with the approving browser's IP. Deleted with the code once it expires. */
  requesterIp: string | null
  /** S256 PKCE challenge from the client that minted the code. Null for codes minted by CLIs that predate it. */
  codeChallenge: string | null
  /** Loopback port the approving browser is redirected to with the approval code. Null means the user pastes it. */
  callbackPort: number | null
  /** sha256 of the approval code, which only the approving browser ever sees. */
  approvalCodeHash: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface PersonalAccessToken {
  id: string
  /** If generated by the CLI this will be "cli", otherwise user-provided */
  name: string
  /** This is the token encrypted using the ENCRYPTION_KEY */
  encryptedToken: unknown
  /** This is shown in the UI, with ******** */
  obfuscatedToken: string
  /** This is used to find the token in the database */
  hashedToken: string
  userId: string
  revokedAt: Timestamp | null
  lastAccessedAt: Timestamp | null
  /** After this the token no longer authenticates. Null means it never expires. */
  expiresAt: Timestamp | null
  /**
   * Set on CLI login tokens: they stop authenticating this many seconds after their last use
   * (or creation, if never used). Null means no idle limit.
   */
  idleTimeoutSeconds: number | null
  /** When the "about to expire" email was sent. Set before sending, so it's sent at most once. */
  expiryReminderSentAt: Timestamp | null
  /** When the "has expired" email was sent. Set before sending, so it's sent at most once. */
  expiredNoticeSentAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface OrganizationAccessToken {
  id: string
  /** User-provided name for the token */
  name: string
  /** Used to differentiate between user-generated and system-generated tokens */
  type: Generated<OrganizationAccessTokenType>
  /** This is used to find the token in the database */
  hashedToken: string
  organizationId: string
  /** Optional expiration date for the token */
  expiresAt: Timestamp | null
  revokedAt: Timestamp | null
  lastAccessedAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface Organization {
  id: string
  slug: string
  title: string
  maximumExecutionTimePerRunInMs: Generated<number>
  maximumConcurrencyLimit: Generated<number>
  /** This is deprecated and will be removed in the future */
  maximumSchedulesLimit: Generated<number>
  maximumDevQueueSize: number | null
  maximumDeployedQueueSize: number | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  deletedAt: Timestamp | null
  companySize: string | null
  onboardingData: unknown
  avatar: unknown
  runsEnabled: Generated<boolean>
  v3Enabled: Generated<boolean>
  /** @deprecated */
  v2Enabled: Generated<boolean>
  /** @deprecated */
  v2MarqsEnabled: Generated<boolean>
  /** @deprecated */
  hasRequestedV3: Generated<boolean>
  apiRateLimiterConfig: unknown
  realtimeRateLimiterConfig: unknown
  batchRateLimitConfig: unknown
  batchQueueConcurrencyConfig: unknown
  /** Webhook limits from the org's plan or the back office: { maxWaitersPerEnvironment, maxWaitersPerEndpoint, concurrency }. Null uses the env defaults. */
  webhookLimitsConfig: unknown
  featureFlags: unknown
  maximumProjectCount: Generated<number>
  /**
   * Maximum session lifetime in seconds for members of this organization.
   * When set, caps each member's User.sessionDuration. The most restrictive cap across
   * all of a user's orgs (excluding nulls) wins.
   */
  maxSessionDuration: number | null
  /**
   * S2 basin that holds this org's realtime streams. Null until the
   * per-org basin has been provisioned (OSS / s2-lite installs leave
   * it null forever; reads fall back to the global basin env var).
   * Set once at provisioning time; retention is reconfigured in-place
   * when the org's plan changes.
   */
  streamBasinName: string | null
  supportAccessMode: Generated<SupportAccessMode>
  supportAccessModeUpdatedAt: Timestamp | null
  supportAccessModeUpdatedById: string | null
}

export interface OrgMember {
  id: string
  organizationId: string
  userId: string
  role: Generated<OrgMemberRole>
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface OrgMemberInvite {
  id: string
  token: string
  email: string
  role: Generated<OrgMemberRole>
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
  rbacRoleId: string | null
  organizationId: string
  inviterId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface RuntimeEnvironment {
  id: string
  slug: string
  apiKey: string
  /** Non-null prevents the root API key from being returned to users. */
  rootApiKeyHiddenAt: Timestamp | null
  /** @deprecated was for v2 */
  pkApiKey: string
  type: Generated<RuntimeEnvironmentType>
  /**
   * If true, this environment has branches and is treated differently in the dashboard/API
   * NB: this flag is NOT used for Development branches, instead (type, parentEnvironmentId) = (DEVELOPMENT, NULL) is used
   */
  isBranchableEnvironment: Generated<boolean>
  branchName: string | null
  parentEnvironmentId: string | null
  /** Only configured on the PREVIEW root. Null disables automatic archiving. */
  previewAutoArchiveAfterDays: number | null
  previewAutoArchiveExcludedBranches: Generated<string[]>
  /** Durable per-project cleanup progress. Null next-check disables scheduling. */
  previewAutoArchiveNextCheckAt: Timestamp | null
  previewAutoArchiveCursorCreatedAt: Timestamp | null
  previewAutoArchiveCursorId: string | null
  git: unknown
  /** When set API calls will fail */
  archivedAt: Timestamp | null
  /** A memorable code for the environment */
  shortcode: string
  maximumConcurrencyLimit: Generated<number>
  concurrencyLimitBurstFactor: Generated<string>
  paused: Generated<boolean>
  pauseSource: EnvironmentPauseSource | null
  autoEnableInternalSources: Generated<boolean>
  organizationId: string
  projectId: string
  orgMemberId: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  /** Allows us to customize the built-in environment variables for a specific environment, like TRIGGER_OTEL_SPAN_ATTRIBUTE_VALUE_LENGTH_LIMIT */
  builtInEnvironmentVariableOverrides: unknown
  tunnelId: string | null
  currentSessionId: string | null
}

export interface ApiKey {
  id: string
  name: string
  key_hash: string
  last_four: string
  /** Optional one-line note the creator adds to say what the key is for. */
  description: string | null
  runtime_environment_id: string
  created_by_user_id: string | null
  issued_by: Generated<ApiKeyIssuer>
  preset_id: string | null
  scopes: string[]
  /** CIDR ranges this key may be used from. Empty means any IP. */
  allowed_ip_ranges: Generated<string[]>
  last_used_at: Timestamp | null
  last_used_ip: string | null
  revoked_at: Timestamp | null
  expires_at: Timestamp | null
  updated_at: Timestamp
  created_at: Generated<Timestamp>
}

/**
 * Records of previously-valid API keys that are still accepted for authentication
 * during a grace window after rotation. Extend or end the grace period by updating `expiresAt`.
 */
export interface RevokedApiKey {
  id: string
  apiKey: string
  runtimeEnvironmentId: string
  expiresAt: Timestamp
  createdAt: Generated<Timestamp>
}

export interface Project {
  id: string
  slug: string
  name: string
  externalRef: string
  organizationId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  deletedAt: Timestamp | null
  /** Set the first time the CLI `init` command completes against this project. Drives the dev onboarding progress. */
  initializedAt: Timestamp | null
  /** Runtime used when a deployment config does not specify one. Null preserves the legacy Node 20 fallback. */
  defaultRuntime: string | null
  version: Generated<ProjectVersion>
  engine: Generated<RunEngineVersion>
  builderProjectId: string | null
  defaultWorkerGroupId: string | null
  /** The master queues they are allowed to use (impacts what they can set as default and trigger runs with) */
  allowedMasterQueues: Generated<string[]>
  buildSettings: unknown
  onboardingData: unknown
}

export interface EnvironmentEncryptionPolicy {
  environmentId: string
  revision: Generated<number>
  wrappingProvider: EnvironmentKeyWrappingProvider | null
  wrappingKeyReference: string | null
  activeKeyId: string | null
}

export interface EnvironmentEncryptionKey {
  id: string
  environmentId: string
  wrappedKey: unknown
  wrappingProvider: EnvironmentKeyWrappingProvider
  wrappingKeyReference: string
  revision: Generated<number>
  createdAt: Generated<Timestamp>
}

export interface SecretReference {
  id: string
  key: string
  provider: Generated<SecretStoreProvider>
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface SecretStore {
  key: string
  value: unknown
  version: Generated<string>
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface OidcSigningKey {
  id: string
  kid: string
  purpose: Generated<OidcSigningKeyPurpose>
  /** At most one STAGED and one ACTIVE key per purpose: enforced by a partial unique index in the migration. */
  status: OidcSigningKeyStatus
  publicJwk: unknown
  readyAt: Timestamp
  publicUntil: Timestamp | null
  generatedByUserId: string
  activatedAt: Timestamp | null
  activatedByUserId: string | null
  retiredAt: Timestamp | null
  retiredByUserId: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface DataMigration {
  id: string
  name: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  completedAt: Timestamp | null
}

export interface BackgroundWorker {
  id: string
  friendlyId: string
  engine: Generated<RunEngineVersion>
  contentHash: string
  sdkVersion: Generated<string>
  cliVersion: Generated<string>
  projectId: string
  runtimeEnvironmentId: string
  version: string
  metadata: unknown
  runtime: string | null
  runtimeVersion: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  workerGroupId: string | null
  supportsLazyAttempts: Generated<boolean>
}

export interface BackgroundWorkerFile {
  id: string
  friendlyId: string
  filePath: string
  contentHash: string
  contents: Buffer
  projectId: string
  createdAt: Generated<Timestamp>
}

export interface Prompt {
  id: string
  friendly_id: string
  slug: string
  description: string | null
  type: Generated<string>
  organizationId: string
  projectId: string
  runtimeEnvironmentId: string
  filePath: string | null
  exportName: string | null
  variableSchema: unknown
  defaultModel: string | null
  defaultConfig: unknown
  tags: Generated<string[]>
  archivedAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface PromptVersion {
  id: string
  version: number
  promptId: string
  textContent: string | null
  chatContent: unknown
  model: string | null
  config: unknown
  source: string
  commitMessage: string | null
  contentHash: string
  labels: Generated<string[]>
  workerId: string | null
  createdBy: string | null
  createdAt: Generated<Timestamp>
}

export interface BackgroundWorkerTask {
  id: string
  slug: string
  description: string | null
  friendlyId: string
  filePath: string
  exportName: string | null
  workerId: string
  projectId: string
  fileId: string | null
  runtimeEnvironmentId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  queueConfig: unknown
  retryConfig: unknown
  machineConfig: unknown
  /**
   * Gates declared on the task: other queues its runs must also hold a concurrency
   * slot in while executing. Shape: [{ queue: string, concurrencyKey?: string }]
   */
  gates: unknown
  queueId: string | null
  maxDurationInSeconds: number | null
  ttl: string | null
  triggerSource: Generated<TaskTriggerSource>
  /**
   * Extra task configuration JSON. Shape depends on triggerSource.
   * AGENT: { type: "ai-sdk-chat" }
   */
  config: unknown
  payloadSchema: unknown
}

export interface WebhookEndpoint {
  id: string
  friendlyId: string
  /**
   * URL key. CSPRNG, NOT a friendlyId: crypto.randomBytes(16) base64url.
   * Generated by the create path (deploy-sync / P2 API), NOT by @default.
   */
  opaqueId: string
  organizationId: string
  projectId: string
  runtimeEnvironmentId: string
  environmentType: RuntimeEnvironmentType
  endpointTenantId: Generated<string>
  endpointExternalRef: Generated<string>
  source: string
  /** The declared webhooks.endpoint.define id (string ref, no relation). */
  declaredId: string
  /** Subscribers of this endpoint: StoredWebhookRoutingTarget[] ({ id, type: "task" | "session", ..., filterAst }). */
  routingTargets: Generated<unknown>
  verifierArtifact: unknown
  metadata: Generated<unknown>
  secretProvisioning: Generated<string>
  /** The source's instructions for an AI agent registering this endpoint with the provider. */
  setupPrompt: string | null
  /** SecretReference.key string. Plain String, NO @relation -> no FK to SecretReference. */
  signingSecretKey: string | null
  status: Generated<WebhookEndpointStatus>
  /**
   * When an operator disabled the endpoint via the dashboard/API. Null means the declarative sync
   * owns the status: a redeploy that re-declares a previously-removed (auto-deactivated) webhook
   * reactivates it. Non-null means the operator disabled it, so the sync leaves the status alone.
   */
  manuallyDeactivatedAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface WebhookDelivery {
  id: string
  friendlyId: string
  webhookEndpointId: string
  organizationId: string
  projectId: string
  runtimeEnvironmentId: string
  environmentType: RuntimeEnvironmentType
  externalDeliveryId: string
  idempotencyKey: string
  runId: string | null
  status: Generated<WebhookDeliveryStatus>
  /** Set from the x-trigger-test ingress header; marks console/test-send deliveries so the list can filter them. */
  isTest: Generated<boolean>
  parsedEvent: unknown
  headers: unknown
  rawBodyHash: string | null
  errorMessage: string | null
  filterReason: string | null
  /** Per-target results: WebhookDeliveryTargetResult[] ({ id, type, status, reason?, runId?, error? }). */
  targets: Generated<unknown>
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  processedAt: Timestamp | null
  /** Retention class in days: the LIST partition key, stamped at ingest from the org's plan and encoded in the id. */
  retentionDays: Generated<number>
}

export interface Alert {
  id: string
  organizationId: string
  projectId: string
  environmentId: string
  capabilityKey: string
  params: unknown
  condition: unknown
  query: string | null
  expiresAt: Timestamp | null
  lifecycle: Generated<string>
  attention: Generated<string>
  observedCondition: Generated<string>
  observedConditionChangedAt: Timestamp | null
  pendingMetSince: Timestamp | null
  evaluatorState: unknown
  evaluationLeaseId: string | null
  evaluationLeaseUntil: Timestamp | null
  lastEvaluatedAt: Timestamp | null
  nextEvaluationAt: Timestamp | null
  notifiers: Generated<unknown>
  name: string | null
  createdBy: string
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  resolvedAt: Timestamp | null
  resolvedBy: string | null
  stoppedAt: Timestamp | null
  stoppedBy: string | null
  lastOccurrenceAt: Timestamp | null
  occurrencesSinceRaise: Generated<number>
  occurrencesTotal: Generated<bigint>
  groupId: string | null
}

/** RANGE-partitioned by month on `at` (see migration). */
export interface AlertEvent {
  id: string
  alertId: string
  type: string
  at: Generated<Timestamp>
  facts: Generated<unknown>
  actorId: string | null
  sourceEventId: string | null
}

export interface AlertOccurrenceSeen {
  alertId: string
  sourceEventId: string
  rowTime: Timestamp
  seenAt: Generated<Timestamp>
  expiresAt: Timestamp
}

/** RANGE-partitioned by month on `eventAt` (see migration). */
export interface AlertDelivery {
  id: string
  alertEventId: string
  eventAt: Timestamp
  notifierId: string
  notifier: unknown
  status: Generated<string>
  claimId: string | null
  claimedAt: Timestamp | null
  deliveredAt: Timestamp | null
  failureReason: string | null
  providerMetadata: unknown
  attempts: Generated<number>
  firstAttemptAt: Timestamp | null
  createdAt: Generated<Timestamp>
}

export interface PlaygroundConversation {
  id: string
  /** The chat session ID used by the transport */
  chatId: string
  /** User-editable conversation title (auto-generated from first message) */
  title: Generated<string>
  /** Which agent this conversation is with */
  agentSlug: string
  /** The current active run backing this conversation (null if no run yet) */
  runId: string | null
  /** The client data JSON used for this conversation */
  clientData: unknown
  /** Accumulated UIMessages from completed turns (for resume without stream replay) */
  messages: unknown
  /** Last SSE event ID — resume from this position to avoid replaying old turns */
  lastEventId: string | null
  projectId: string
  runtimeEnvironmentId: string
  /** The user who started this conversation */
  userId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

/**
 * Durable, typed, bidirectional I/O primitive. Owns two S2 streams (.out / .in).
 * The row is essentially static — no status, no counters, no pointers. No
 * foreign keys: project/runtimeEnvironment/organization ids are plain
 * scalar columns (matches TaskRun pattern). List-style queries are served
 * from ClickHouse, not Postgres, so only point-lookup indexes live here.
 */
export interface Session {
  id: string
  friendlyId: string
  /**
   * User-supplied identifier scoped to the environment. Used for
   * idempotent upsert and for resolving sessions via the public API.
   */
  externalId: string | null
  /** Plain string — intentionally not an enum. */
  type: string
  /** Denormalized scoping columns — no FK relations. */
  projectId: string
  runtimeEnvironmentId: string
  environmentType: RuntimeEnvironmentType
  organizationId: string
  /**
   * Task this session triggers runs against. Required — Sessions are
   * task-bound: creating a session also triggers its first run, and
   * every subsequent re-trigger uses this same identifier.
   */
  taskIdentifier: string
  /**
   * Trigger config used for every run this session schedules. Shape
   * (validated at the route layer, opaque to the DB):
   * { basePayload: object, machine?: string, queue?: string,
   * tags?: string[], maxAttempts?: number,
   * idleTimeoutInSeconds?: number }
   * `basePayload` carries the customer's client-data; runtime fields
   * (chatId, messages, trigger) are merged at trigger time.
   */
  triggerConfig: unknown
  /**
   * Whether this session was created from the Test/playground UI rather
   * than by a real chat.agent() trigger. Mirrors TaskRun.isTest.
   */
  isTest: Generated<boolean>
  tags: Generated<string[]>
  metadata: unknown
  /**
   * Live run pointer — non-FK so run deletion never cascades. Can lag
   * reality; the `.in/append` handler re-checks the snapshot status
   * before reusing it.
   */
  currentRunId: string | null
  /**
   * Monotonic counter used for optimistic locking on `currentRunId`
   * swaps. Bumped atomically alongside any update that changes
   * `currentRunId`.
   */
  currentRunVersion: Generated<number>
  /** Terminal markers — written once, never flipped back. */
  closedAt: Timestamp | null
  closedReason: string | null
  expiresAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  /**
   * S2 basin where this session's stream pair lives. Stamped at create
   * time from `Organization.streamBasinName` so reads can resolve the
   * basin without joining org. Null when the org has no per-org basin
   * (OSS, or pre-backfill); reads fall back to the global basin.
   */
  streamBasinName: string | null
  /**
   * Storage URI (with protocol prefix) for this session's chat.agent
   * snapshot blob. Set on first snapshot write. Null = pre-column session,
   * fall back to computed default path.
   */
  chatSnapshotStoragePath: string | null
}

/**
 * Historical record of every run a Session has owned. Append-only —
 * rows are inserted on each `ensureRunForSession` claim, never updated.
 * Lets us reconstruct the run timeline of a chat for debugging /
 * dashboard surfaces. The relation cascades on Session delete (tied to
 * the session lifecycle) but `runId` is a plain string column with no
 * FK to TaskRun so run pruning is independent.
 */
export interface SessionRun {
  id: string
  sessionId: string
  /** TaskRun.id (no FK — runs may be archived independently of session history) */
  runId: string
  /**
   * One of: "initial" | "continuation" | "upgrade" | "manual".
   * Plain string for forward-compat with future trigger reasons.
   */
  reason: string
  triggeredAt: Generated<Timestamp>
}

export interface TaskRun {
  id: string
  number: Generated<number>
  friendlyId: string
  engine: Generated<RunEngineVersion>
  status: Generated<TaskRunStatus>
  statusReason: string | null
  idempotencyKey: string | null
  idempotencyKeyExpiresAt: Timestamp | null
  /** Stores the user-provided key and scope: { key: string, scope: "run" | "attempt" | "global" } */
  idempotencyKeyOptions: unknown
  /** Debounce options: { key: string, delay: string, createdAt: Date } */
  debounce: unknown
  taskIdentifier: string
  isTest: Generated<boolean>
  payload: string
  payloadType: Generated<string>
  context: unknown
  traceContext: unknown
  traceId: string
  spanId: string
  runtimeEnvironmentId: string
  environmentType: RuntimeEnvironmentType | null
  projectId: string
  organizationId: string | null
  queue: string
  lockedQueueId: string | null
  /** The main queue that this run is part of */
  masterQueue: Generated<string>
  /** User-facing geo region, stamped at trigger; workerQueue is where it actually ran. */
  region: string | null
  /** @deprecated */
  secondaryMasterQueue: string | null
  /** From engine v2+ this will be defined after a run has been dequeued (starting at 1) */
  attemptNumber: number | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  /** Denormized column that holds the raw tags */
  runTags: string[]
  /** Denormalized version of the background worker task */
  taskVersion: string | null
  sdkVersion: string | null
  cliVersion: string | null
  /** startedAt marks the point at which a run is dequeued from MarQS */
  startedAt: Timestamp | null
  /** executedAt is set when the first attempt is about to execute */
  executedAt: Timestamp | null
  completedAt: Timestamp | null
  machinePreset: string | null
  usageDurationMs: Generated<number>
  costInCents: Generated<number>
  baseCostInCents: Generated<number>
  lockedAt: Timestamp | null
  lockedById: string | null
  lockedToVersionId: string | null
  /**
   * The "priority" of the run. This is just a negative offset in ms for the queue timestamp
   * E.g. a value of 60_000 would put the run into the queue 60s ago.
   */
  priorityMs: Generated<number>
  concurrencyKey: string | null
  /**
   * Gates for this run: other queues it must also hold a concurrency slot in while
   * executing. Shape: [{ queue: string, concurrencyKey?: string }]
   */
  gates: unknown
  delayUntil: Timestamp | null
  queuedAt: Timestamp | null
  ttl: string | null
  expiredAt: Timestamp | null
  maxAttempts: number | null
  lockedRetryConfig: unknown
  /** optional token that can be used to authenticate the task run */
  oneTimeUseToken: string | null
  /** Where the logs are stored */
  taskEventStore: Generated<string>
  queueTimestamp: Timestamp | null
  scheduleInstanceId: string | null
  scheduleId: string | null
  bulkActionGroupIds: Generated<string[]>
  logsDeletedAt: Timestamp | null
  replayedFromTaskRunFriendlyId: string | null
  rootTaskRunId: string | null
  parentTaskRunId: string | null
  parentTaskRunAttemptId: string | null
  batchId: string | null
  /** whether or not the task run was created because of a triggerAndWait for batchTriggerAndWait */
  resumeParentOnCompletion: Generated<boolean>
  /** The depth of this task run in the task run hierarchy */
  depth: Generated<number>
  /** The span ID of the "trigger" span in the parent task run */
  parentSpanId: string | null
  /** Holds the state of the run chain for deadlock detection */
  runChainState: unknown
  /** seed run metadata */
  seedMetadata: string | null
  seedMetadataType: Generated<string>
  /** Run metadata */
  metadata: string | null
  metadataType: Generated<string>
  metadataVersion: Generated<number>
  /** Structured annotations: triggerSource, triggerAction, rootTriggerSource, rootScheduleId */
  annotations: unknown
  /** Whether the latest attempt was a warm start. Null until first attempt starts. */
  isWarmStart: boolean | null
  /** Run output */
  output: string | null
  outputType: Generated<string>
  /** Run error */
  error: unknown
  /** Organization's billing plan type (cached for fallback when billing API fails) */
  planType: string | null
  maxDurationInSeconds: number | null
  /** The version of the realtime streams implementation used by the run */
  realtimeStreamsVersion: Generated<string>
  /** Store the stream keys that are being used by the run */
  realtimeStreams: Generated<string[]>
  /**
   * S2 basin where this run's realtime streams live. Stamped at create
   * time from `Organization.streamBasinName` so reads can resolve the
   * basin without joining org. Null when the org has no per-org basin
   * (OSS, or pre-backfill); reads fall back to the global basin.
   */
  streamBasinName: string | null
}

export interface TaskRunTemplate {
  id: string
  taskSlug: string
  triggerSource: TaskTriggerSource
  label: string
  payload: string | null
  payloadType: Generated<string>
  metadata: string | null
  metadataType: Generated<string>
  queue: string | null
  concurrencyKey: string | null
  ttlSeconds: number | null
  maxAttempts: number | null
  maxDurationSeconds: number | null
  tags: string[]
  machinePreset: string | null
  projectId: string
  organizationId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

/**
 * Used by the RunEngine during TaskRun execution
 * It has the required information to transactionally progress a run through states,
 * and prevent side effects like heartbeats failing a run that has progressed.
 * It is optimised for performance and is designed to be cleared at some point,
 * so there are no cascading relationships to other models.
 */
export interface TaskRunExecutionSnapshot {
  id: string
  /** This should always be 2+ (V1 didn't use the run engine or snapshots) */
  engine: Generated<RunEngineVersion>
  /** The execution status */
  executionStatus: TaskRunExecutionStatus
  /** For debugging */
  description: string
  /** We store invalid snapshots as a record of the run state when we tried to move */
  isValid: Generated<boolean>
  error: string | null
  /** The previous snapshot ID */
  previousSnapshotId: string | null
  /** Run */
  runId: string
  runStatus: TaskRunStatus
  batchId: string | null
  /** This is the current run attempt number. Users can define how many attempts they want for a run. */
  attemptNumber: number | null
  /** Environment */
  environmentId: string
  environmentType: RuntimeEnvironmentType
  projectId: string
  organizationId: string
  /** An array of waitpoint IDs in the correct order, used for batches */
  completedWaitpointOrder: string[]
  /** Checkpoint */
  checkpointId: string | null
  /** Worker */
  workerId: string | null
  runnerId: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  lastHeartbeatAt: Timestamp | null
  /** Metadata used by various systems in the run engine */
  metadata: unknown
}

export interface TaskRunCheckpoint {
  id: string
  friendlyId: string
  type: TaskRunCheckpointType
  location: string
  imageRef: string | null
  reason: string | null
  metadata: string | null
  projectId: string
  runtimeEnvironmentId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

/**
 * A Waitpoint blocks a run from continuing until it's completed
 * If there's a waitpoint blocking a run, it shouldn't be in the queue
 */
export interface Waitpoint {
  id: string
  friendlyId: string
  type: WaitpointType
  status: Generated<WaitpointStatus>
  completedAt: Timestamp | null
  /** If it's an Event type waitpoint, this is the event. It can also be provided for the DATETIME type */
  idempotencyKey: string
  /** If this is true then we can show it in the dashboard/return it from the SDK */
  userProvidedIdempotencyKey: boolean
  /** If there's a user provided idempotency key, this is the time it expires at */
  idempotencyKeyExpiresAt: Timestamp | null
  /**
   * If an idempotencyKey is no longer active, we store it here and generate a new one for the idempotencyKey field.
   * Clearing an idempotencyKey is useful for debounce or cancelling child runs.
   * This is a workaround because Prisma doesn't support partial indexes.
   */
  inactiveIdempotencyKey: string | null
  /** If it's a RUN type waitpoint, this is the associated run */
  completedByTaskRunId: string | null
  /**
   * If it's a DATETIME type waitpoint, this is the date.
   * If it's a MANUAL waitpoint, this can be set as the `timeout`.
   */
  completedAfter: Timestamp | null
  /** If it's a BATCH type waitpoint, this is the associated batch */
  completedByBatchId: string | null
  /** When completed, an output can be stored here */
  output: string | null
  outputType: Generated<string>
  outputIsError: Generated<boolean>
  projectId: string
  environmentId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  /**
   * Denormized column that holds the raw tags
   * Denormalized column that holds the raw tags
   */
  tags: string[]
}

export interface TaskRunWaitpoint {
  id: string
  taskRunId: string
  waitpointId: string
  projectId: string
  /** This span id is completed when the waitpoint is completed. This is used with cached runs (idempotent) */
  spanIdToComplete: string | null
  batchId: string | null
  batchIndex: number | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface WaitpointTag {
  id: string
  name: string
  environmentId: string
  projectId: string
  createdAt: Generated<Timestamp>
}

export interface FeatureFlag {
  id: string
  key: string
  value: unknown
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface WorkerInstance {
  id: string
  /** For example "worker-1" */
  name: string
  /**
   * If managed, it will default to the name, e.g. "worker-1"
   * If unmanged, it will be prefixed with the deployment ID e.g. "deploy-123-worker-1"
   */
  resourceIdentifier: string
  metadata: unknown
  workerGroupId: string
  organizationId: string | null
  projectId: string | null
  environmentId: string | null
  deploymentId: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  lastDequeueAt: Timestamp | null
  lastHeartbeatAt: Timestamp | null
}

export interface WorkerInstanceGroup {
  id: string
  type: WorkerInstanceGroupType
  /** For example "us-east-1" */
  name: string
  /**
   * If managed, it will default to the name, e.g. "us-east-1"
   * If unmanged, it will be prefixed with the project ID e.g. "project_1-us-east-1"
   */
  masterQueue: string
  /** "N. Virginia, USA", "Frankfurt, Germany", etc. Used for display purposes */
  description: string | null
  hidden: Generated<boolean>
  tokenId: string
  organizationId: string | null
  projectId: string | null
  cloudProvider: string | null
  /** "usa", "europe", etc. Used like a pseudo enum for things like flags */
  location: string | null
  staticIPs: string | null
  workloadType: Generated<WorkloadType>
  /** Geo region; container + MICROVM groups for one geo share it. Set-once once it has runs. */
  region: string | null
  /**
   * When true, runs enqueued to this worker queue may skip the intermediate queue
   * and be pushed directly to the worker queue when concurrency is available.
   */
  enableFastPath: Generated<boolean>
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface WorkerGroupToken {
  id: string
  tokenHash: string
  previousTokenHash: string | null
  previousTokenValidUntil: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface TaskRunTag {
  id: string
  name: string
  friendlyId: string
  projectId: string
  createdAt: Generated<Timestamp>
}

/** This is used for triggerAndWait and batchTriggerAndWait. The taskRun is the child task, it points at a parent attempt or a batch */
export interface TaskRunDependency {
  id: string
  taskRunId: string
  checkpointEventId: string | null
  dependentAttemptId: string | null
  dependentBatchRunId: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  resumedAt: Timestamp | null
}

/** deprecated, we hadn't included the project id in the unique constraint */
export interface TaskRunCounter {
  taskIdentifier: string
  lastNumber: Generated<number>
}

/** Used for the TaskRun number (e.g. #1,421) */
export interface TaskRunNumberCounter {
  id: string
  taskIdentifier: string
  environmentId: string
  lastNumber: Generated<number>
}

/** This is not used from engine v2+, attempts use the TaskRunExecutionSnapshot and TaskRun */
export interface TaskRunAttempt {
  id: string
  number: Generated<number>
  friendlyId: string
  taskRunId: string
  backgroundWorkerId: string
  backgroundWorkerTaskId: string
  runtimeEnvironmentId: string
  queueId: string
  status: Generated<TaskRunAttemptStatus>
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  startedAt: Timestamp | null
  completedAt: Timestamp | null
  usageDurationMs: Generated<number>
  error: unknown
  output: string | null
  outputType: Generated<string>
}

/** This is the unified otel span/log model that will eventually be replaced by clickhouse */
export interface TaskEvent {
  id: string
  /** This matches the span name for a trace event, or the log body for a log event */
  message: string
  traceId: string
  spanId: string
  parentId: string | null
  tracestate: string | null
  isError: Generated<boolean>
  isPartial: Generated<boolean>
  isCancelled: Generated<boolean>
  /** deprecated: don't use this, moving this to properties, this now uses TaskEventKind.LOG */
  isDebug: Generated<boolean>
  serviceName: string
  serviceNamespace: string
  level: Generated<TaskEventLevel>
  kind: Generated<TaskEventKind>
  status: Generated<TaskEventStatus>
  links: unknown
  events: unknown
  /** This is the time the event started in nanoseconds since the epoch */
  startTime: bigint
  /** This is the duration of the event in nanoseconds */
  duration: Generated<bigint>
  attemptId: string | null
  attemptNumber: number | null
  environmentId: string
  environmentType: RuntimeEnvironmentType
  organizationId: string
  projectId: string
  projectRef: string
  runId: string
  runIsTest: Generated<boolean>
  idempotencyKey: string | null
  taskSlug: string
  taskPath: string | null
  taskExportName: string | null
  workerId: string | null
  workerVersion: string | null
  queueId: string | null
  queueName: string | null
  batchId: string | null
  /** This represents all the span attributes available, like http.status_code, and special attributes like $style.icon, $output, $metadata.payload.userId, as it's used for searching and filtering */
  properties: unknown
  /** This represents all span attributes in the $metadata namespace, like $metadata.payload */
  metadata: unknown
  /** This represents all span attributes in the $style namespace, like $style */
  style: unknown
  /** This represents all span attributes in the $output namespace, like $output */
  output: unknown
  /** This represents the mimetype of the output, such as application/json or application/super+json */
  outputType: string | null
  payload: unknown
  payloadType: string | null
  createdAt: Generated<Timestamp>
  usageDurationMs: Generated<number>
  usageCostInCents: Generated<number>
  machinePreset: string | null
  machinePresetCpu: number | null
  machinePresetMemory: number | null
  machinePresetCentsPerMs: number | null
}

export interface TaskQueue {
  id: string
  friendlyId: string
  name: string
  type: Generated<TaskQueueType>
  version: Generated<TaskQueueVersion>
  orderableName: string | null
  /**
   * How the queue's concurrency was declared: V1 = legacy concurrencyLimit semantics
   * (per key when keyed, whole queue otherwise), V2 = the explicit perKey/total shape.
   */
  concurrencyVersion: Generated<TaskQueueConcurrencyVersion>
  /**
   * QUEUE rows are real queues; LIMIT rows back named concurrency limits and are
   * excluded from every queue-facing surface.
   */
  role: Generated<TaskQueueRole>
  projectId: string
  runtimeEnvironmentId: string
  /** Represents the current concurrency limit for the queue */
  concurrencyLimit: number | null
  /** When the concurrency limit was overridden */
  concurrencyLimitOverriddenAt: Timestamp | null
  /** Who overrode the concurrency limit (will be null if overridden via the API) */
  concurrencyLimitOverriddenBy: string | null
  /** If concurrencyLimit is overridden, this is the overridden value */
  concurrencyLimitBase: number | null
  /**
   * If the override was expressed as a percentage of the environment limit, this stores that
   * percentage (the source of truth). The absolute concurrencyLimit is materialized from it.
   * Decimal(5,2) allows fractional percentages like 12.50% (0.01–100.00).
   */
  concurrencyLimitOverridePercent: string | null
  /**
   * Caps total concurrent runs across ALL concurrencyKey values of this queue
   * (concurrencyLimit applies per key value). Null = no total cap.
   */
  totalConcurrencyLimit: number | null
  /** When the total concurrency limit was overridden */
  totalConcurrencyLimitOverriddenAt: Timestamp | null
  /** Who overrode the total concurrency limit (null when overridden via the API) */
  totalConcurrencyLimitOverriddenBy: string | null
  /** If totalConcurrencyLimit is overridden, the declared value it reverts to on reset */
  totalConcurrencyLimitBase: number | null
  rateLimit: unknown
  paused: Generated<boolean>
  /** Hidden from the dashboard queue list. Display-only: never read on trigger/dequeue paths. */
  archivedAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface BatchTaskRun {
  id: string
  friendlyId: string
  idempotencyKey: string | null
  idempotencyKeyExpiresAt: Timestamp | null
  status: Generated<BatchTaskRunStatus>
  runtimeEnvironmentId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  /** Friendly IDs */
  runIds: Generated<string[]>
  runCount: Generated<number>
  payload: string | null
  payloadType: Generated<string>
  options: unknown
  batchVersion: Generated<string>
  /** sealed is set to true once no more items can be added to the batch */
  sealed: Generated<boolean>
  sealedAt: Timestamp | null
  /** this is the expected number of items in the batch */
  expectedCount: Generated<number>
  /** this is the completed number of items in the batch. once this reaches expectedCount, and the batch is sealed, the batch is considered completed */
  completedCount: Generated<number>
  completedAt: Timestamp | null
  resumedAt: Timestamp | null
  /** this is used to be able to "seal" this BatchTaskRun when all of the runs have been triggered asynchronously, and using the "parallel" processing strategy */
  processingJobsCount: Generated<number>
  processingJobsExpectedCount: Generated<number>
  /** optional token that can be used to authenticate the task run */
  oneTimeUseToken: string | null
  /** When processing started (status changed to PROCESSING) */
  processingStartedAt: Timestamp | null
  /** When processing completed (all items processed) */
  processingCompletedAt: Timestamp | null
  /** Count of successfully created runs */
  successfulRunCount: number | null
  /** Count of failed run creations */
  failedRunCount: number | null
  taskIdentifier: string | null
  checkpointEventId: string | null
  dependentTaskAttemptId: string | null
}

/** Used in engine V1 only */
export interface BatchTaskRunItem {
  id: string
  status: Generated<BatchTaskRunItemStatus>
  batchTaskRunId: string
  taskRunId: string
  taskRunAttemptId: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  completedAt: Timestamp | null
}

/** Track individual run creation failures in batch processing (Run Engine v2) */
export interface BatchTaskRunError {
  id: string
  batchTaskRunId: string
  /** Which item in the batch (0-based index) */
  index: number
  /** The task identifier that was being triggered */
  taskIdentifier: string
  /** The payload that failed (JSON, may be truncated) */
  payload: string | null
  /** The options that were used */
  options: unknown
  /** Error message */
  error: string
  /** Error code if available */
  errorCode: string | null
  createdAt: Generated<Timestamp>
}

export interface EnvironmentVariable {
  id: string
  friendlyId: string
  key: string
  projectId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface EnvironmentVariableValue {
  id: string
  valueReferenceId: string | null
  variableId: string
  environmentId: string
  /** If true, the value is secret and cannot be revealed */
  isSecret: Generated<boolean>
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  version: Generated<number>
  lastUpdatedBy: unknown
}

export interface Checkpoint {
  id: string
  friendlyId: string
  type: CheckpointType
  location: string
  imageRef: string
  reason: string | null
  metadata: string | null
  runId: string
  attemptId: string
  attemptNumber: number | null
  projectId: string
  runtimeEnvironmentId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface CheckpointRestoreEvent {
  id: string
  type: CheckpointRestoreEventType
  reason: string | null
  metadata: string | null
  checkpointId: string
  runId: string
  attemptId: string
  projectId: string
  runtimeEnvironmentId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface WorkerDeployment {
  id: string
  contentHash: string
  friendlyId: string
  shortCode: string
  version: string
  runtime: string | null
  runtimeVersion: string | null
  /** CLI version that initiated the deploy, stamped at initialization */
  cliVersion: string | null
  imageReference: string | null
  imagePlatform: Generated<string>
  /** Key of the microVM template prebuilt from this deployment's image, set when the builder succeeds */
  templateKey: string | null
  /** Object key of the full index snapshot in the deployment index store. Set only after the upload succeeds. */
  indexObjectKey: string | null
  /** sha256 hex of the snapshot's uncompressed JSON */
  indexSha256: string | null
  indexFormatVersion: number | null
  externalBuildData: unknown
  buildServerMetadata: unknown
  /**
   * Encrypted build-time env vars for pre-bundled (fromBundle) deploys, as an
   * EncryptedSecretValue envelope. Kept for the redeploy window, then cleared by a sweep.
   */
  buildEnvVars: unknown
  status: Generated<WorkerDeploymentStatus>
  type: Generated<WorkerDeploymentType>
  projectId: string
  environmentId: string
  workerId: string | null
  triggeredById: string | null
  triggeredVia: string | null
  commitSHA: string | null
  externalId: string | null
  startedAt: Timestamp | null
  installedAt: Timestamp | null
  builtAt: Timestamp | null
  deployedAt: Timestamp | null
  failedAt: Timestamp | null
  errorData: unknown
  canceledAt: Timestamp | null
  canceledReason: string | null
  git: unknown
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface WorkerDeploymentPromotion {
  id: string
  /** This is the promotion label, e.g. "current" */
  label: string
  deploymentId: string
  environmentId: string
}

/** Schedules can be attached to tasks to trigger them on a schedule */
export interface TaskSchedule {
  id: string
  type: Generated<ScheduleType>
  /** users see this as `id`. They start with schedule_ */
  friendlyId: string
  /** a reference to a task (not a foreign key because it's across versions) */
  taskIdentifier: string
  /** can be provided and we won't create another with the same key */
  deduplicationKey: string
  userProvidedDeduplicationKey: Generated<boolean>
  /** the CRON pattern */
  generatorExpression: string
  generatorDescription: Generated<string>
  generatorType: Generated<ScheduleGeneratorType>
  /** These are IANA format string, or the default "UTC". E.g. "America/New_York" */
  timezone: Generated<string>
  windowDurationSeconds: number | null
  windowPercentage: number | null
  defaultWindowDurationSeconds: number | null
  minimumWindowDurationSeconds: number | null
  /** Can be provided by the user then accessed inside a run */
  externalId: string | null
  /** @deprecated stop writing 2026-04-30; reads moved out of code (UI now derives from cron's previous slot). Drop in follow-up. */
  lastRunTriggeredAt: Timestamp | null
  projectId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  active: Generated<boolean>
}

/** An instance links a schedule with an environment */
export interface TaskScheduleInstance {
  id: string
  taskScheduleId: string
  environmentId: string
  projectId: string
  schedulePhase: number | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  active: Generated<boolean>
  /** @deprecated stop writing 2026-04-30; engine derives from cron + exactScheduleTime. Drop in follow-up. */
  lastScheduledTimestamp: Timestamp | null
  /** @deprecated stop writing 2026-04-30; engine derives from cron + now(). Drop in follow-up. */
  nextScheduledTimestamp: Timestamp | null
}

export interface RuntimeEnvironmentSession {
  id: string
  ipAddress: string
  environmentId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  disconnectedAt: Timestamp | null
}

export interface ProjectAlertChannel {
  id: string
  friendlyId: string
  /** can be provided and we won't create another with the same key */
  deduplicationKey: string
  userProvidedDeduplicationKey: Generated<boolean>
  integrationId: string | null
  enabled: Generated<boolean>
  type: ProjectAlertChannelType
  name: string
  properties: unknown
  alertTypes: ProjectAlertType[]
  environmentTypes: Generated<RuntimeEnvironmentType[]>
  errorAlertConfig: unknown
  projectId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface ProjectAlert {
  id: string
  friendlyId: string
  projectId: string
  environmentId: string
  channelId: string
  status: Generated<ProjectAlertStatus>
  type: ProjectAlertType
  taskRunAttemptId: string | null
  taskRunId: string | null
  workerDeploymentId: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface ProjectAlertStorage {
  id: string
  projectId: string
  alertChannelId: string
  alertType: ProjectAlertType
  storageId: string
  storageData: unknown
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface OrganizationIntegration {
  id: string
  friendlyId: string
  service: IntegrationService
  /** Identifier for external, integration's organization (e.g. Vercel's team) */
  externalOrganizationId: string | null
  integrationData: unknown
  tokenReferenceId: string
  organizationId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  deletedAt: Timestamp | null
}

export interface OrganizationProjectIntegration {
  id: string
  organizationIntegrationId: string
  projectId: string
  /** Identifier for webhooks, for example Vercel's projectId */
  externalEntityId: string
  /** Save useful data like config or external entity name */
  integrationData: unknown
  /** UserId who installed the integration */
  installedBy: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  deletedAt: Timestamp | null
}

/** Bulk actions, like canceling and replaying runs */
export interface BulkActionGroup {
  id: string
  friendlyId: string
  projectId: string
  environmentId: string | null
  type: BulkActionType
  /** When the group is created it's pending. After we've processed all the items it's completed. This does not mean the associated runs are completed. */
  status: Generated<BulkActionStatus>
  /** The query that will be executed to get the runs to process (if null, we are passing run ids directly) */
  queryName: string | null
  /** The params that will be passed to the query */
  params: unknown
  /** Billing-limit dedupe key (grace hitAt or resolve job key). Null for dashboard bulk actions. */
  dedupeKey: string | null
  /** The cursor that will be passed to the query (null for the first page) */
  cursor: unknown
  /** The number of runs that have been processed successfully */
  successCount: Generated<number>
  /** The number of runs that have been failed */
  failureCount: Generated<number>
  /** The total number of runs that will be processed */
  totalCount: Generated<number>
  completionNotification: Generated<BulkActionNotificationType>
  userId: string | null
  /** The name of the bulk action */
  name: string | null
  /** The time the bulk action was completed */
  completedAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface BulkActionItem {
  id: string
  /** @deprecated not used in new BulkActions */
  friendlyId: string | null
  groupId: string
  type: BulkActionType
  /** When the item is created it's pending. After we've processed the item it's completed. This does not mean the associated runs are completed. */
  status: Generated<BulkActionItemStatus>
  sourceRunId: string
  destinationRunId: string | null
  error: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface RealtimeStreamChunk {
  id: string
  key: string
  value: string
  sequence: number
  runId: string
  createdAt: Generated<Timestamp>
}

/** This is the unified otel span/log model that will eventually be replaced by clickhouse */
export interface TaskEventPartitioned {
  id: string
  /** This matches the span name for a trace event, or the log body for a log event */
  message: string
  traceId: string
  spanId: string
  parentId: string | null
  tracestate: string | null
  isError: Generated<boolean>
  isPartial: Generated<boolean>
  isCancelled: Generated<boolean>
  serviceName: string
  serviceNamespace: string
  level: Generated<TaskEventLevel>
  kind: Generated<TaskEventKind>
  status: Generated<TaskEventStatus>
  links: unknown
  events: unknown
  /** This is the time the event started in nanoseconds since the epoch */
  startTime: bigint
  /** This is the duration of the event in nanoseconds */
  duration: Generated<bigint>
  attemptId: string | null
  attemptNumber: number | null
  environmentId: string
  environmentType: RuntimeEnvironmentType
  organizationId: string
  projectId: string
  projectRef: string
  runId: string
  runIsTest: Generated<boolean>
  idempotencyKey: string | null
  taskSlug: string
  taskPath: string | null
  taskExportName: string | null
  workerId: string | null
  workerVersion: string | null
  queueId: string | null
  queueName: string | null
  batchId: string | null
  /** This represents all the span attributes available, like http.status_code, and special attributes like $style.icon, $output, $metadata.payload.userId, as it's used for searching and filtering */
  properties: unknown
  /** This represents all span attributes in the $metadata namespace, like $metadata.payload */
  metadata: unknown
  /** This represents all span attributes in the $style namespace, like $style */
  style: unknown
  /** This represents all span attributes in the $output namespace, like $output */
  output: unknown
  /** This represents the mimetype of the output, such as application/json or application/super+json */
  outputType: string | null
  payload: unknown
  payloadType: string | null
  createdAt: Generated<Timestamp>
  usageDurationMs: Generated<number>
  usageCostInCents: Generated<number>
  machinePreset: string | null
  machinePresetCpu: number | null
  machinePresetMemory: number | null
  machinePresetCentsPerMs: number | null
}

export interface GithubAppInstallation {
  id: string
  appInstallationId: bigint
  targetId: bigint
  targetType: string
  accountHandle: string
  permissions: unknown
  repositorySelection: GithubRepositorySelection
  installedBy: string | null
  installedByUserId: string | null
  organizationId: string
  deletedAt: Timestamp | null
  suspendedAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface GithubRepository {
  id: string
  githubId: bigint
  name: string
  fullName: string
  htmlUrl: string
  private: boolean
  defaultBranch: string
  installationId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface TrustedDeployRepository {
  id: string
  projectId: string
  issuer: Generated<string>
  repositoryOwner: string
  repositoryName: string
  repositoryId: string | null
  repositoryOwnerId: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface TrustedDeployEnvironment {
  id: string
  trustedDeployRepositoryId: string
  runtimeEnvironmentId: string
  enabled: Generated<boolean>
  scope: TrustedDeployScope
  scopeValue: string | null
  additionalClaims: Generated<unknown>
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface ConnectedGithubRepository {
  id: string
  projectId: string
  repositoryId: string
  branchTracking: unknown
  previewDeploymentsEnabled: Generated<boolean>
  buildFilters: unknown
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface ImpersonationAuditLog {
  id: string
  action: ImpersonationAuditLogAction
  adminId: string
  targetId: string
  ipAddress: string | null
  createdAt: Generated<Timestamp>
}

export interface SupportAccessRequest {
  id: string
  organizationId: string
  requestedById: string | null
  reason: string
  status: Generated<SupportAccessRequestStatus>
  approvedById: string | null
  approvedAt: Timestamp | null
  expiresAt: Timestamp | null
  cancelledById: string | null
  cancelledAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface CustomerQuery {
  id: string
  /** The TSQL query text that was executed */
  query: string
  /** The scope of the query (determines which tenant IDs were used) */
  scope: CustomerQueryScope
  /** Query execution statistics from ClickHouse */
  stats: unknown
  /** AI-generated title summarizing the query */
  title: string | null
  /** Where the query originated from */
  source: Generated<CustomerQuerySource>
  organizationId: string
  projectId: string | null
  environmentId: string | null
  userId: string | null
  /**
   * Time filter settings used when the query was executed
   * Period like "7d", "24h", etc.
   */
  filterPeriod: string | null
  /** Custom start date (ISO 8601) */
  filterFrom: Timestamp | null
  /** Custom end date (ISO 8601) */
  filterTo: Timestamp | null
  createdAt: Generated<Timestamp>
}

export interface IntegrationDeployment {
  id: string
  /** For example Vercel */
  integrationName: string
  /** External ID */
  integrationDeploymentId: string
  commitSHA: string
  deploymentId: string | null
  /** External deployment status */
  status: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

/** A user-defined metrics dashboard */
export interface MetricsDashboard {
  id: string
  friendlyId: string
  title: string
  description: Generated<string>
  organizationId: string
  projectId: string | null
  ownerId: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  /**
   * JSON that defines the config, queries, layout and config of all widgets.
   * There will be a version field for the format.
   */
  layout: string
}

/** A known LLM model or model pattern for cost tracking */
export interface LlmModel {
  id: string
  friendly_id: string
  project_id: string | null
  model_name: string
  match_pattern: string
  start_date: Timestamp | null
  source: Generated<string>
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  provider: string | null
  description: string | null
  context_window: number | null
  max_output_tokens: number | null
  capabilities: Generated<string[]>
  is_hidden: Generated<boolean>
  base_model_name: string | null
  pricing_unit: string | null
}

/** A pricing tier for a model (supports volume-based or conditional pricing) */
export interface LlmPricingTier {
  id: string
  model_id: string
  name: string
  is_default: Generated<boolean>
  priority: Generated<number>
  conditions: Generated<unknown>
}

/** A price point for a usage type within a pricing tier */
export interface LlmPrice {
  id: string
  model_id: string
  pricing_tier_id: string
  usage_type: string
  price: string
}

/** Admin-created notification definitions */
export interface PlatformNotification {
  id: string
  friendlyId: string
  /** Admin-facing title for identification in admin tools */
  title: string
  /** Versioned JSON rendering content (see payload schema in spec) */
  payload: unknown
  surface: PlatformNotificationSurface
  scope: PlatformNotificationScope
  userId: string | null
  organizationId: string | null
  projectId: string | null
  /** When notification becomes active */
  startsAt: Timestamp
  /** When notification expires */
  endsAt: Timestamp
  /** CLI: auto-expire N days after user first saw it */
  cliMaxDaysAfterFirstSeen: number | null
  /** CLI: auto-expire after shown N times to user */
  cliMaxShowCount: number | null
  /** CLI: only show every N-th request (e.g. cliShowEvery=3 means show on 3rd, 6th, 9th…) */
  cliShowEvery: number | null
  /** Ordering within same scope level (higher = more important) */
  priority: Generated<number>
  /**
   * Staged-but-unpublished. While true the notification is hidden from all
   * user-facing reads regardless of startsAt/endsAt; publishing sets real
   * dates and flips this to false.
   */
  isDraft: Generated<boolean>
  /** Soft delete */
  archivedAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

/** Per-user tracking of notification views and dismissals */
export interface PlatformNotificationInteraction {
  id: string
  notificationId: string
  userId: string
  /** Set by beacon/CLI GET on first impression */
  firstSeenAt: Generated<Timestamp>
  /** Times shown (incremented per beacon/CLI view) */
  showCount: Generated<number>
  /** User dismissed in webapp */
  webappDismissedAt: Timestamp | null
  /** User clicked a link in webapp */
  webappClickedAt: Timestamp | null
  /** Auto-dismissed or expired in CLI */
  cliDismissedAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

/**
 * *
 * * Error group state is used to track when a user has interacted with an error (ignored/resolved)
 * * The actual error data is in ClickHouse.
 */
export interface ErrorGroupState {
  id: string
  organizationId: string
  projectId: string
  environmentId: string | null
  taskIdentifier: string
  errorFingerprint: string
  status: Generated<ErrorGroupStatus>
  /**
   * *
   * * Error is ignored until this date
   */
  ignoredUntil: Timestamp | null
  /**
   * *
   * * Error is ignored until this occurrence rate
   */
  ignoredUntilOccurrenceRate: number | null
  /**
   * *
   * * Error is ignored until this total occurrences
   */
  ignoredUntilTotalOccurrences: number | null
  /**
   * Total occurrence count at the time the error was ignored (from ClickHouse).
   * Used with ignoredUntilTotalOccurrences to compute occurrences since ignoring.
   */
  ignoredAtOccurrenceCount: bigint | null
  /**
   * *
   * * Error was ignored at this date
   */
  ignoredAt: Timestamp | null
  /**
   * *
   * * Reason for ignoring the error
   */
  ignoredReason: string | null
  /**
   * *
   * * User who ignored the error
   */
  ignoredByUserId: string | null
  /**
   * *
   * * Error was resolved at this date
   */
  resolvedAt: Timestamp | null
  /**
   * *
   * * Error was resolved in this version
   */
  resolvedInVersion: string | null
  /**
   * *
   * * User who resolved the error
   */
  resolvedBy: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface TaskIdentifier {
  id: string
  runtimeEnvironmentId: string
  projectId: string
  slug: string
  currentTriggerSource: Generated<TaskTriggerSource>
  currentWorkerId: string | null
  firstSeenAt: Generated<Timestamp>
  lastSeenAt: Generated<Timestamp>
  isInLatestDeployment: Generated<boolean>
}

/**
 * Defines org-scoped data store overrides (e.g. dedicated ClickHouse for HIPAA orgs).
 * Multiple organizations can share a single data store row via organizationIds.
 */
export interface OrganizationDataStore {
  id: string
  /** Human-readable unique key (e.g. "hipaa-clickhouse-us-east") */
  key: string
  /** Organization IDs that use this data store */
  organizationIds: string[]
  kind: DataStoreKind
  /** Versioned config JSON. Structure is discriminated by the top-level `version` field. */
  config: unknown
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface BackgroundWorkerToBackgroundWorkerFile {
  A: string
  B: string
}

export interface BackgroundWorkerToTaskQueue {
  A: string
  B: string
}

export interface WaitpointRunConnections {
  A: string
  B: string
}

export interface CompletedWaitpoints {
  A: string
  B: string
}

export interface DB {
  User: User
  MfaBackupCode: MfaBackupCode
  InvitationCode: InvitationCode
  AuthorizationCode: AuthorizationCode
  PersonalAccessToken: PersonalAccessToken
  OrganizationAccessToken: OrganizationAccessToken
  Organization: Organization
  OrgMember: OrgMember
  OrgMemberInvite: OrgMemberInvite
  RuntimeEnvironment: RuntimeEnvironment
  api_keys: ApiKey
  RevokedApiKey: RevokedApiKey
  Project: Project
  EnvironmentEncryptionPolicy: EnvironmentEncryptionPolicy
  EnvironmentEncryptionKey: EnvironmentEncryptionKey
  SecretReference: SecretReference
  SecretStore: SecretStore
  OidcSigningKey: OidcSigningKey
  DataMigration: DataMigration
  BackgroundWorker: BackgroundWorker
  BackgroundWorkerFile: BackgroundWorkerFile
  prompts: Prompt
  prompt_versions: PromptVersion
  BackgroundWorkerTask: BackgroundWorkerTask
  WebhookEndpoint: WebhookEndpoint
  WebhookDelivery: WebhookDelivery
  Alert: Alert
  AlertEvent: AlertEvent
  AlertOccurrenceSeen: AlertOccurrenceSeen
  AlertDelivery: AlertDelivery
  PlaygroundConversation: PlaygroundConversation
  Session: Session
  SessionRun: SessionRun
  TaskRun: TaskRun
  TaskRunTemplate: TaskRunTemplate
  TaskRunExecutionSnapshot: TaskRunExecutionSnapshot
  TaskRunCheckpoint: TaskRunCheckpoint
  Waitpoint: Waitpoint
  TaskRunWaitpoint: TaskRunWaitpoint
  WaitpointTag: WaitpointTag
  FeatureFlag: FeatureFlag
  WorkerInstance: WorkerInstance
  WorkerInstanceGroup: WorkerInstanceGroup
  WorkerGroupToken: WorkerGroupToken
  TaskRunTag: TaskRunTag
  TaskRunDependency: TaskRunDependency
  TaskRunCounter: TaskRunCounter
  TaskRunNumberCounter: TaskRunNumberCounter
  TaskRunAttempt: TaskRunAttempt
  TaskEvent: TaskEvent
  TaskQueue: TaskQueue
  BatchTaskRun: BatchTaskRun
  BatchTaskRunItem: BatchTaskRunItem
  BatchTaskRunError: BatchTaskRunError
  EnvironmentVariable: EnvironmentVariable
  EnvironmentVariableValue: EnvironmentVariableValue
  Checkpoint: Checkpoint
  CheckpointRestoreEvent: CheckpointRestoreEvent
  WorkerDeployment: WorkerDeployment
  WorkerDeploymentPromotion: WorkerDeploymentPromotion
  TaskSchedule: TaskSchedule
  TaskScheduleInstance: TaskScheduleInstance
  RuntimeEnvironmentSession: RuntimeEnvironmentSession
  ProjectAlertChannel: ProjectAlertChannel
  ProjectAlert: ProjectAlert
  ProjectAlertStorage: ProjectAlertStorage
  OrganizationIntegration: OrganizationIntegration
  OrganizationProjectIntegration: OrganizationProjectIntegration
  BulkActionGroup: BulkActionGroup
  BulkActionItem: BulkActionItem
  RealtimeStreamChunk: RealtimeStreamChunk
  TaskEventPartitioned: TaskEventPartitioned
  GithubAppInstallation: GithubAppInstallation
  GithubRepository: GithubRepository
  TrustedDeployRepository: TrustedDeployRepository
  TrustedDeployEnvironment: TrustedDeployEnvironment
  ConnectedGithubRepository: ConnectedGithubRepository
  ImpersonationAuditLog: ImpersonationAuditLog
  SupportAccessRequest: SupportAccessRequest
  CustomerQuery: CustomerQuery
  IntegrationDeployment: IntegrationDeployment
  MetricsDashboard: MetricsDashboard
  llm_models: LlmModel
  llm_pricing_tiers: LlmPricingTier
  llm_prices: LlmPrice
  PlatformNotification: PlatformNotification
  PlatformNotificationInteraction: PlatformNotificationInteraction
  ErrorGroupState: ErrorGroupState
  TaskIdentifier: TaskIdentifier
  OrganizationDataStore: OrganizationDataStore
  _BackgroundWorkerToBackgroundWorkerFile: BackgroundWorkerToBackgroundWorkerFile
  _BackgroundWorkerToTaskQueue: BackgroundWorkerToTaskQueue
  _WaitpointRunConnections: WaitpointRunConnections
  _completedWaitpoints: CompletedWaitpoints
}
