import type { ColumnType } from 'kysely'

export type Generated<T> =
  T extends ColumnType<infer S, infer I, infer U>
    ? ColumnType<S, I | undefined, U>
    : ColumnType<T, T | undefined, T>

export type Timestamp = ColumnType<Date, Date | string, Date | string>

/**
 * Creatives (ENG-3993): page elements the web SDK renders next to surveys. Every type lives in
 * `Creative`; `content` and `styling` are JSON validated per type by the API v3 schemas, so a new type is
 * a new payload variant, not new columns.
 */
export type CreativeType = 'feedbackButton' | 'banner'

/**
 * The author's intent only. Scheduled, Active, "Paused with its survey" and Closed after the
 * end time are derived (`apps/web/modules/ee/creatives/lib/display-status.ts` and its SQL twin); nothing
 * writes this column when a schedule ends or a survey pauses.
 */
export type CreativeStatus = 'draft' | 'inProgress' | 'paused' | 'completed'

/**
 * What makes a creative appear. A creative left with no action never shows; it never falls back to
 * every page view.
 */
export type CreativeShowOn = 'pageView' | 'action'

/** The device type counts are split by: the values of the Show when `deviceType` operand. */
export type CreativeDeviceType = 'phone' | 'tablet' | 'desktop'

export type PipelineTriggers = 'responseCreated' | 'responseUpdated' | 'responseFinished'

export type WebhookSource = 'user' | 'zapier' | 'make' | 'n8n' | 'activepieces'

export type ContactAttributeType = 'default' | 'custom'

export type ContactAttributeDataType = 'string' | 'number' | 'date'

/** What a reviewer decided about a response. One value today; null on `Response.outcome` means none. */
export type ResponseOutcome = 'lowQuality'

export type SurveyStatus = 'draft' | 'inProgress' | 'paused' | 'completed'

export type SurveyAttributeFilterCondition = 'equals' | 'notEquals'

export type SurveyQuotaAction = 'endSurvey' | 'continueSurvey'

export type ResponseQuotaLinkStatus = 'screenedIn' | 'screenedOut'

export type SurveyType = 'link' | 'app'

export type displayOptions = 'displayOnce' | 'displayMultiple' | 'displaySome' | 'respondMultiple'

/**
 * Who may read a survey. `workspace` is everyone with workspace access (the behaviour before ENG-3282);
 * `restricted` is the owner and the organization's owners and managers only.
 */
export type SurveyVisibility = 'restricted' | 'workspace'

export type SurveyScriptMode = 'add' | 'replace'

export type ActionType = 'code' | 'noCode'

export type IntegrationType = 'googleSheets' | 'notion' | 'airtable' | 'slack'

export type DataMigrationStatus = 'pending' | 'applied' | 'failed'

export type WidgetPlacement = 'bottomLeft' | 'bottomRight' | 'topLeft' | 'topRight' | 'center'

export type SurveyOverlay = 'none' | 'light' | 'dark'

/**
 * The bootstrap marker's only value (ENG-2247). A one-variant enum rather than a boolean on
 * purpose: a unique index over a nullable boolean has a slot for `true` AND one for `false`, so the
 * first `false` ever written would take the second slot and break every later sign-up. There is no
 * `false` to write here, which is the same guarantee without a CHECK constraint Prisma cannot model.
 */
export type BootstrapAdminMarker = 'bootstrapAdmin'

export type OrganizationRole = 'owner' | 'manager' | 'member' | 'billing' | 'dashboard_viewer'

/**
 * Defines permission levels for API keys.
 * Controls what operations an API key can perform.
 */
export type ApiKeyPermission = 'read' | 'write' | 'manage'

export type IdentityProvider = 'email' | 'github' | 'google' | 'azuread' | 'openid' | 'saml'

/**
 * Where an Embedded Data field's value comes from.
 * `computed` is set by survey logic, `ingested` arrives from a URL param or the SDK, and
 * `reserved` is auto-captured system metadata. Reserved fields are never stored as rows —
 * they are a static catalog in code — but the value exists so every layer shares one type.
 */
export type EmbeddedDataSource = 'computed' | 'ingested' | 'reserved'

/** The declared type of an Embedded Data field's value. */
export type EmbeddedDataType = 'string' | 'number' | 'boolean' | 'date'

export type TeamUserRole = 'admin' | 'contributor'

export type WorkspaceTeamPermission = 'read' | 'readWrite' | 'manage'

export type ChartType = 'area' | 'bar' | 'pie' | 'big_number' | 'matrix'

export type FeedbackSourceType = 'formbricks_survey' | 'csv'

export type FeedbackSourceStatus = 'active' | 'paused' | 'error'

/**
 * Which survey responses a Formbricks feedback source imports into Hub.
 * completedOnly matches the live ingestion path, which only runs on responseFinished.
 * all additionally imports partial responses — answers a respondent typed but never submitted.
 */
export type FeedbackSourceImportMode = 'completedOnly' | 'all'

/**
 * Whether a formbricks_survey source tracks every supported element of the surveys it maps, or an
 * explicit subset the operator curated. Ignored by csv sources, which map fields instead.
 *
 * Orthogonal to FeedbackSourceImportMode: that decides which responses are imported, this decides
 * which questions a response is imported for.
 *
 * This exists because a question added to a survey after the source was connected is
 * indistinguishable, from stored state alone, from a question the operator deliberately excluded —
 * survey elements carry no timestamps. Recording the intent at write time (where the full candidate
 * set IS known) is the same approach webhooks take with `surveyIds: []` meaning "all surveys", and
 * segment survey-interaction filters with `surveyScope: "any" | "specific"`.
 *
 * Defaults to `specific` so sources that predate the column keep their current behaviour: they gain
 * no mappings on their own until the operator next saves the source, at which point the scope is
 * derived from the selection they submit.
 */
export type FeedbackSourceElementScope = 'all' | 'specific'

export type HubFieldType =
  | 'text'
  | 'categorical'
  | 'nps'
  | 'csat'
  | 'ces'
  | 'rating'
  | 'number'
  | 'boolean'
  | 'date'

export type DeletionCleanupKind =
  | 'hubSurvey'
  | 'hubResponses'
  | 'storageFiles'
  | 'storageSurveyFolder'
  | 'hubResponseOutcome'

export type AttachmentExportStatus = 'queued' | 'running' | 'completed' | 'failed' | 'expired'

/** Data retention (ENG-3697). The kind of data a policy governs. */
export type RetentionEntity = 'responses' | 'surveys' | 'members'

/** The surveys policy's conditions. Every ticked one must hold, and they share the policy's period. */
export type RetentionSurveyCondition = 'noResponse' | 'noChange' | 'createdBefore'

/**
 * What wrote a `RetentionRun`: the nightly sweep, or the archive purge recording the deletion of surveys
 * the surveys policy archived. Only sweep runs count as the sweep having run (its lease, its warning
 * restart after a gap, where its scan resumes, the health check).
 */
export type RetentionRunKind = 'sweep' | 'purge'

/**
 * Who made an exemption: a person (with a reason), or a restore from the archive under a surveys policy
 * whose only condition is "created before" (`restartSurveyRetention`), which carries no reason.
 */
export type RetentionExemptionSource = 'manual' | 'restore'

/**
 * What a `RetentionRunItem` row points at. Response deletions are grouped into one `survey` row per
 * survey, with the number of responses in `count`.
 */
export type RetentionTargetType = 'survey' | 'user'

export type RetentionRunItemAction = 'notified' | 'archived' | 'deactivated' | 'deleted' | 'skipped'

export type RetentionSkipReason = 'exempt' | 'lastOwner' | 'otherOrganization' | 'noRecipient'

export type WorkflowStatus = 'draft' | 'enabled' | 'disabled' | 'archived'

export type WorkflowRunStatus = 'queued' | 'running' | 'completed' | 'failed' | 'canceled'

export type WorkflowRunLogStatus = 'pending' | 'running' | 'succeeded' | 'failed' | 'skipped'

/**
 * A feedback button or a banner. Opens at most one Destination: a survey (`surveyId`) or a link
 * (`linkUrl`), never both.
 *
 * @property type - Fixed at creation.
 * @property content - Per-type JSON (`FeedbackButtonContent` / `BannerContent` in the v3 spec), texts keyed by locale code.
 * @property styling - Per-type JSON (`FeedbackButtonStyling` / `BannerStyling` in the v3 spec).
 * @property survey - The Destination survey, in the same workspace by construction of the composite key.
 * @property showConditions - Show when (`ZShowConditions`, shared with surveys); null always shows.
 * @property autoSelectLanguage - The browser's language picks the texts' language when the site sets none; null (never chosen, as on Survey) is off.
 */
export interface Creative {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  name: string
  type: CreativeType
  status: Generated<CreativeStatus>
  publishedAt: Timestamp | null
  archivedAt: Timestamp | null
  workspaceId: string
  createdBy: string | null
  content: unknown
  styling: unknown
  surveyId: string | null
  linkUrl: string | null
  showOn: Generated<CreativeShowOn>
  showConditions: unknown
  displayPercentage: number | null
  startAt: Timestamp | null
  endAt: Timestamp | null
  autoSelectLanguage: boolean | null
}

/**
 * A creative's Impressions and Clicks on one UTC day on one device type (ENG-4000). Written
 * only by the counter flush job as `GREATEST(stored, counted)`, from Redis counters that hold the day's
 * running totals, so a repeated flush changes nothing. At most 3 rows per creative per day. Deleting
 * the creative deletes its counts.
 */
export interface CreativeDailyStat {
  creativeId: string
  /** UTC day. */
  date: Timestamp
  deviceType: CreativeDeviceType
  impressions: Generated<number>
  clicks: Generated<number>
}

/**
 * A creative's Impressions and Clicks on one UTC day on one page (host + path), written like
 * `CreativeDailyStat`. At most 1000 pages per creative per day; the rest are counted under `(other)`.
 */
export interface CreativeDailyPageStat {
  creativeId: string
  /** UTC day. */
  date: Timestamp
  /** Host + path, without query string or hash, or `(other)`. */
  page: string
  impressions: Generated<number>
  clicks: Generated<number>
}

/** Links a creative to the actions that show it (`showOn: action`). Mirrors `SurveyTrigger`. */
export interface CreativeTrigger {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  creativeId: string
  actionClassId: string
}

/** The languages a creative's texts are written in. Mirrors `SurveyLanguage`; exactly one is the default. */
export interface CreativeLanguage {
  languageId: string
  creativeId: string
  default: Generated<boolean>
  enabled: Generated<boolean>
}

/**
 * Metadata for an outbound email attempt. Content and recipient addresses are never stored here.
 * Identifiers are deliberately not foreign keys so erasure can scrub a row without retaining source data.
 */
export interface OutboundEmailAudit {
  id: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  expiresAt: Timestamp
  contentExpiresAt: Timestamp
  environment: string
  emailType: string
  trigger: string
  actorKind: string
  actorUserId: string | null
  configuredByUserId: string | null
  actorUserHash: string | null
  configuredByUserHash: string | null
  organizationId: string | null
  surveyId: string | null
  responseId: string | null
  followUpId: string | null
  workflowId: string | null
  workflowRunId: string | null
  inviteId: string | null
  recipientMac: string | null
  sesMessageHash: string | null
  snapshotKey: string | null
  snapshotSha256: string | null
  status: Generated<string>
  statusRecordedAt: Timestamp | null
  auditGap: Generated<boolean>
  erasureRequestedAt: Timestamp | null
  erasedAt: Timestamp | null
  contentDeleteRetryAt: Timestamp | null
  contentDeleteFailures: Generated<number>
}

/**
 * Holds an account erasure request across an in-flight email send. The address is never stored.
 * The user id is removed once deletion is confirmed; one-way recovery keys expire after 90 days.
 */
export interface OutboundEmailAuditErasureIntent {
  id: string
  userId: string | null
  userHash: string
  recipientMac: string
  createdAt: Generated<Timestamp>
  activatedAt: Timestamp | null
  checkedAt: Timestamp | null
  expiresAt: Timestamp
}

/**
 * Represents a webhook endpoint for receiving survey-related events.
 * Webhooks can be configured to receive notifications about response creation, updates, and completion.
 *
 * @property id - Unique identifier for the webhook
 * @property name - Optional display name for the webhook
 * @property url - The endpoint URL where events will be sent
 * @property source - Origin of the webhook (user, zapier, make, etc.)
 * @property workspace - Associated workspace
 * @property triggers - Types of events that trigger this webhook
 * @property surveyIds - List of surveys this webhook is monitoring
 */
export interface Webhook {
  id: string
  name: string | null
  created_at: Generated<Timestamp>
  updated_at: Generated<Timestamp>
  url: string
  source: Generated<WebhookSource>
  workspaceId: string
  triggers: PipelineTriggers[]
  surveyIds: string[]
  secret: string | null
}

/**
 * Represents an attribute value associated with a contact.
 * Used to store custom properties and characteristics of contacts.
 *
 * @property id - Unique identifier for the attribute
 * @property attributeKey - Reference to the attribute definition
 * @property contact - The contact this attribute belongs to
 * @property value - The string value of the attribute (used for string type + backwards compatibility)
 * @property valueNumber - Native numeric storage for number type attributes
 * @property valueDate - Native date storage for date type attributes
 */
export interface ContactAttribute {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  attributeKeyId: string
  contactId: string
  value: string
  valueNumber: number | null
  valueDate: Timestamp | null
}

/**
 * Defines the possible attributes that can be assigned to contacts.
 * Acts as a schema for contact attributes within a workspace.
 *
 * @property id - Unique identifier for the attribute key
 * @property isUnique - Whether the attribute must have unique values across contacts
 * @property key - The attribute identifier used in the system
 * @property name - Display name for the attribute
 * @property type - Whether this is a default or custom attribute
 * @property dataType - The data type of the attribute (string, number, date)
 * @property workspace - The workspace this attribute belongs to
 */
export interface ContactAttributeKey {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  isUnique: Generated<boolean>
  key: string
  name: string | null
  description: string | null
  type: Generated<ContactAttributeType>
  dataType: Generated<ContactAttributeDataType>
  workspaceId: string
}

/**
 * Represents a person or user who can receive and respond to surveys.
 * Contacts are workspace-specific and can have multiple attributes and responses.
 *
 * @property id - Unique identifier for the contact
 * @property workspace - The workspace this contact belongs to
 * @property responses - Survey responses from this contact
 * @property attributes - Custom attributes associated with this contact
 * @property displays - Record of surveys shown to this contact
 */
export interface Contact {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  workspaceId: string
}

/**
 * Stores a user's response to a survey, including their answers and metadata.
 * Each response is linked to a specific survey and optionally to a contact.
 *
 * @property id - Unique identifier for the response
 * @property finished - Whether the survey was completed
 * @property survey - The associated survey
 * @property contact - The optional contact who provided the response
 * @property data - JSON object containing the actual response data
 * @property variables - Custom variables used in the response
 * @property ttc - Time to completion metrics
 * @property meta - Additional metadata about the response
 */
export interface Response {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Generated<Timestamp>
  finished: Generated<boolean>
  surveyId: string
  contactId: string | null
  endingId: string | null
  /** [ResponseData] */
  data: Generated<unknown>
  /** [ResponseVariables] */
  variables: Generated<unknown>
  /** [ResponseTtc] */
  ttc: Generated<unknown>
  /** [ResponseMeta] */
  meta: Generated<unknown>
  /**
   * Why a value in `data` will not read back as its declared Embedded Data type — a failed coercion
   * or a truncated oversize value (ENG-1845). Server-computed on every ingest and never accepted
   * from the client. Null means no ingest boundary has written it, which is not the same claim as
   * an empty list: without this column a wrong-typed value is indistinguishable from an absent one,
   * and surfacing that in the UI (v2) would have nothing to read.
   * [ResponseIngestFlags]
   */
  ingest_flags: unknown
  /**
   * A reviewer's verdict on the response; null means none. `lowQuality` hides it from the Analysis
   * export by default and from Feedback Records, and nothing else: it is never deleted and every
   * count, list, webhook and integration still includes it (ENG-6780).
   */
  outcome: ResponseOutcome | null
  /**
   * What the respondent was not shown (ENG-6936, ADR 0029): the blocks and elements a show-if kept
   * from them, and per element its options, matrix rows and columns, by stable ids. Written by the
   * survey while they answer; null means unknown (older responses and bundles), read as before.
   * Omitted from reads without a `select` (PRISMA_GLOBAL_OMIT), so it never reaches a v1/v2 payload,
   * webhook or integration.
   * [ResponseNotShown]
   */
  not_shown: unknown
  /** [ResponseContactAttributes] */
  contactAttributes: unknown
  singleUseId: string | null
  language: string | null
  displayId: string | null
}

/**
 * Represents a label that can be applied to survey responses.
 * Used for categorizing and organizing responses within a workspace.
 *
 * @property id - Unique identifier for the tag
 * @property name - Display name of the tag
 * @property responses - Survey responses tagged with this label
 * @property workspace - The workspace this tag belongs to
 */
export interface Tag {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  name: string
  workspaceId: string
}

/**
 * Junction table linking tags to responses.
 * Enables many-to-many relationship between tags and responses.
 *
 * @property response - The tagged response
 * @property tag - The tag applied to the response
 */
export interface TagsOnResponses {
  responseId: string
  tagId: string
}

/**
 * Records when a survey is shown to a user.
 * Tracks survey display history and response status.
 *
 * @property id - Unique identifier for the display event
 * @property survey - The survey that was displayed
 * @property contact - The contact who saw the survey
 * @property response - The associated response if one exists
 * @property creativeId - The creative that opened the survey. No foreign key, so the record survives the creative's deletion.
 */
export interface Display {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  surveyId: string
  contactId: string | null
  creativeId: string | null
}

/**
 * Links surveys to specific trigger actions.
 * Defines which user actions should cause a survey to be displayed.
 * This is the connection table between Surveys and ActionClasses that determines
 * when and under what conditions a survey should be triggered.
 *
 * @property id - Unique identifier for the trigger
 * @property survey - The survey to be triggered
 * @property actionClass - The action that triggers the survey
 * @property createdAt - When the trigger was created
 * @property updatedAt - When the trigger was last modified
 */
export interface SurveyTrigger {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  surveyId: string
  actionClassId: string
}

/**
 * Defines targeting rules for surveys based on contact attributes.
 * Used to show surveys only to contacts matching specific criteria.
 *
 * @property id - Unique identifier for the filter
 * @property attributeKey - The contact attribute to filter on
 * @property survey - The survey being filtered
 * @property condition - The comparison operator to use
 * @property value - The value to compare against
 */
export interface SurveyAttributeFilter {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  attributeKeyId: string
  surveyId: string
  condition: SurveyAttributeFilterCondition
  value: string
}

/**
 * Represents a complete survey configuration including questions, styling, and display rules.
 * Core model for the survey functionality in Formbricks.
 *
 * @property id - Unique identifier for the survey
 * @property name - Display name of the survey
 * @property type - Survey delivery method (link, web, website, app)
 * @property status - Current state of the survey (draft, active, completed, etc)
 * @property workspace - The workspace this survey belongs to
 * @property questions - JSON array containing survey questions configuration
 * @property displayOption - Rules for how often the survey can be shown
 * @property triggers - Actions that can trigger this survey
 * @property attributeFilters - Rules for targeting specific contacts
 * @property customHeadScripts - Survey-specific custom HTML scripts (self-hosted only)
 * @property customHeadScriptsMode - "add" (merge with workspace) or "replace" (override workspace)
 */
export interface Survey {
  /** [SurveyHiddenFields] */
  hiddenFields: Generated<unknown>
  /** [SurveyVariables] */
  variables: Generated<unknown>
  ownershipInitialized: Generated<boolean>
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  name: string
  redirectUrl: string | null
  type: Generated<SurveyType>
  workspaceId: string
  createdBy: string | null
  /**
   * ENG-3282: authorization facts, projected into SpiceDB by the survey outbox trigger. `ownerId` is
   * the durable ownership source; `createdBy` stays pure attribution and is never rewritten.
   */
  visibility: Generated<SurveyVisibility>
  ownerId: string | null
  /**
   * Incremented by every accepted visibility change; fences projection work. An INSERT starts at 0
   * with `visibilityProjectedVersion` -1 (the `survey_visibility_pending` trigger): the initial
   * projection, decided from PostgreSQL facts until the projector acknowledges it.
   */
  visibilityVersion: Generated<number>
  /** The `visibilityVersion` SpiceDB last acknowledged; -1 while none ever was. */
  visibilityProjectedVersion: Generated<number>
  /**
   * "visibilityVersion" <> "visibilityProjectedVersion", kept by the `survey_visibility_pending`
   * trigger (a BEFORE trigger overrides whatever an INSERT or UPDATE supplies). Read-only: never put it
   * in create/update data.
   */
  visibilityPending: Generated<boolean>
  visibilityChangedAt: Timestamp | null
  visibilityChangedById: string | null
  status: Generated<SurveyStatus>
  /** [SurveyWelcomeCard] */
  welcomeCard: Generated<unknown>
  /** [SurveyQuestions] */
  questions: Generated<unknown>
  /** [SurveyBlocks] */
  blocks: Generated<unknown[]>
  /** [SurveyEnding] */
  endings: Generated<unknown[]>
  displayOption: Generated<displayOptions>
  recontactDays: number | null
  displayLimit: number | null
  /** [SurveyInlineTriggers] */
  inlineTriggers: unknown
  autoClose: number | null
  keepEndingCardOpen: Generated<boolean>
  autoComplete: number | null
  delay: Generated<number>
  publishOn: Timestamp | null
  closeOn: Timestamp | null
  archivedAt: Timestamp | null
  /** [SurveyClosedMessage] */
  surveyClosedMessage: unknown
  segmentId: string | null
  /** [SurveyWorkspaceOverwrites] */
  workspaceOverwrites: unknown
  /** [SurveyStyling] */
  styling: unknown
  /** [SurveySingleUse] */
  singleUse: Generated<unknown>
  isVerifyEmailEnabled: Generated<boolean>
  isBackButtonHidden: Generated<boolean>
  isAutoProgressingEnabled: Generated<boolean>
  isCaptureIpEnabled: Generated<boolean>
  is_anonymize_responses_enabled: Generated<boolean>
  pin: string | null
  displayPercentage: string | null
  showLanguageSwitch: boolean | null
  /**
   * Open the survey in the respondent's browser language when it matches an enabled survey
   * language. Null on surveys that predate the setting, which reads as off.
   */
  autoSelectLanguage: boolean | null
  /** [SurveyRecaptcha] */
  recaptcha: Generated<unknown>
  /** [SurveyLinkMetadata] */
  metadata: Generated<unknown>
  slug: string | null
  customHeadScripts: string | null
  customHeadScriptsMode: Generated<SurveyScriptMode | null>
  /**
   * [CustomCssStored] Survey custom CSS (ENG-2949): source plus trusted compiled output, written only
   * by the shared custom CSS save service. null = no survey CSS. Kept outside `styling` (M2.07).
   */
  customCss: unknown
  /**
   * [SurveyShowConditions] App surveys: conditions checked in the browser when a trigger fires
   * (ENG-3873). null = no conditions, which is what every existing survey keeps.
   */
  showConditions: unknown
  /**
   * [SurveyStartLogic] Rules before the first question (ENG-6932): block-rule format, run by the
   * survey on load. null = none, which is what every existing survey keeps.
   */
  startLogic: unknown
}

/**
 * Represents a quota configuration for a survey.
 * Defines response limits and conditions for quota management.
 *
 * @property id - Unique identifier for the quota
 * @property survey - The survey this quota belongs to
 * @property name - Display name for the quota
 * @property limit - Maximum number of responses allowed
 * @property conditions - JSON object containing quota conditions
 * @property action - Action to take when quota is reached
 * @property endingCardId - Optional ending card to show when quota is reached
 */
export interface SurveyQuota {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  surveyId: string
  name: string
  limit: number
  /** [SurveyQuotaLogic] */
  logic: Generated<unknown>
  action: SurveyQuotaAction
  endingCardId: string | null
  countPartialSubmissions: Generated<boolean>
}

/**
 * Junction table linking responses to quotas.
 * Tracks which responses counted towards specific quotas and their status.
 *
 * @property response - The response that counted towards the quota
 * @property quota - The quota this response counted towards
 * @property status - Whether the response was screened in or out
 */
export interface ResponseQuotaLink {
  responseId: string
  quotaId: string
  status: ResponseQuotaLinkStatus
}

/**
 * Defines follow-up actions for survey responses.
 * Enables automated actions based on specific survey response conditions.
 *
 * @property id - Unique identifier for the follow-up
 * @property survey - The associated survey
 * @property name - Display name for the follow-up
 * @property trigger - Conditions that activate the follow-up
 * @property action - Actions to take when triggered
 */
export interface SurveyFollowUp {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  surveyId: string
  name: string
  /** [SurveyFollowUpTrigger] */
  trigger: unknown
  /** [SurveyFollowUpAction] */
  action: unknown
}

/**
 * Represents a user action that can trigger surveys.
 * Used to define points in the user journey where surveys can be shown.
 *
 * @property id - Unique identifier for the action
 * @property name - Display name of the action
 * @property type - Whether this is a code or no-code action
 * @property key - Unique identifier used in code implementation
 * @property noCodeConfig - Configuration for no-code setup
 * @property workspace - The workspace this action belongs to
 */
export interface ActionClass {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  name: string
  description: string | null
  type: ActionType
  key: string | null
  /** [ActionClassNoCodeConfig] */
  noCodeConfig: unknown
  workspaceId: string
}

/**
 * Represents third-party service integrations.
 * Enables data flow between Formbricks and external services.
 *
 * @property id - Unique identifier for the integration
 * @property type - The service being integrated (Google Sheets, Notion, etc.)
 * @property workspace - The workspace this integration belongs to
 * @property config - Service-specific configuration details
 */
export interface Integration {
  id: string
  type: IntegrationType
  /** [IntegrationConfig] */
  config: unknown
  workspaceId: string
}

/**
 * Tracks database schema migrations.
 * Used to manage and track the state of data structure changes.
 *
 * @property id - Unique identifier for the migration
 * @property name - Name of the migration
 * @property status - Current state of the migration (pending, applied, failed)
 * @property startedAt - When the migration began
 * @property finishedAt - When the migration completed
 */
export interface DataMigration {
  id: string
  started_at: Generated<Timestamp>
  finished_at: Timestamp | null
  name: string
  status: DataMigrationStatus
}

/**
 * Main grouping mechanism for resources in Formbricks.
 * Each organization can have multiple workspaces to separate different applications or products.
 *
 * @property id - Unique identifier for the workspace
 * @property name - Display name of the workspace
 * @property organization - Reference to parent organization
 * @property surveys - Collection of surveys in this workspace
 * @property styling - Workspace-wide styling configuration
 * @property config - Workspace-specific configuration
 * @property recontactDays - Default recontact delay for surveys
 * @property placement - Default widget placement for in-app surveys
 */
export interface Workspace {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  name: string
  legacyEnvironmentId: string | null
  organizationId: string
  /** [Styling] */
  styling: Generated<unknown>
  /** [WorkspaceConfig] */
  config: Generated<unknown>
  recontactDays: Generated<number>
  linkSurveyBranding: Generated<boolean>
  inAppSurveyBranding: Generated<boolean>
  placement: Generated<WidgetPlacement>
  clickOutsideClose: Generated<boolean>
  overlay: Generated<SurveyOverlay>
  /** Custom overlay colour (hex, no alpha) and opacity (10-100%). null = the preset for `overlay`. */
  overlayColor: string | null
  overlayOpacity: number | null
  /** [Logo] */
  logo: unknown
  appSetupCompleted: Generated<boolean>
  customHeadScripts: string | null
  /** [CustomCssStored] Workspace custom CSS (ENG-2949), written only by the shared save service. */
  customCss: unknown
  /**
   * [CustomCssStored] The one recoverable previous workspace CSS revision, source included. Restored by
   * submitting its source again through the current processor and authorization checks.
   */
  customCssPrevious: unknown
}

/**
 * Represents the top-level organizational hierarchy in Formbricks.
 * Self-hosting instances typically have a single organization, while Formbricks Cloud
 * supports multiple organizations with multi-tenancy.
 *
 * @property id - Unique identifier for the organization
 * @property name - Display name of the organization
 * @property memberships - User memberships within the organization
 * @property workspaces - Collection of workspaces owned by the organization
 * @property billing - JSON field containing billing information
 * @property whitelabel - Whitelabel configuration for the organization
 * @property isAISmartToolsEnabled - Controls access to AI smart tools (e.g. translations) that never touch collected data
 * @property displayTimeZone - Optional IANA time zone used for human-facing response timestamps (exports, integrations) and for the calendar days behind the analytics date filters (survey summary, charts, dashboards); null means UTC
 */
export interface Organization {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  name: string
  /** [OrganizationWhitelabel] */
  whitelabel: Generated<unknown>
  isAISmartToolsEnabled: Generated<boolean>
  displayTimeZone: string | null
}

/**
 * Stores billing and Stripe synchronization data for an organization.
 *
 * @property organizationId - Reference to the organization
 * @property stripeCustomerId - Stripe customer ID for cloud billing
 * @property limits - Plan limits used by feature and quota checks
 * @property usageCycleAnchor - Stable anchor used to derive the current monthly usage cycle
 * @property stripe - Stripe synchronization snapshot data
 */
export interface OrganizationBilling {
  organization_id: string
  stripe_customer_id: string | null
  /** [OrganizationBillingPlanLimits] */
  limits: unknown
  usage_cycle_anchor: Timestamp | null
  /** [OrganizationStripeBilling] */
  stripe: unknown
  created_at: Generated<Timestamp>
  updated_at: Timestamp
}

/**
 * Links users to organizations with specific roles.
 * Manages organization membership and permissions.
 * Core model for managing user access within organizations.
 *
 * @property organization - The organization the user belongs to
 * @property user - The member user
 * @property accepted - Whether the user has accepted the membership
 * @property role - User's role within the organization (owner, manager, member, billing)
 */
export interface Membership {
  organizationId: string
  userId: string
  accepted: Generated<boolean>
  role: Generated<OrganizationRole>
}

/**
 * Represents pending invitations to join an organization.
 * Used to manage the process of adding new users to an organization.
 * Once accepted, invites are converted into memberships.
 *
 * @property id - Unique identifier for the invite
 * @property email - Email address of the invited user
 * @property name - Optional display name for the invited user
 * @property organization - The organization sending the invite
 * @property creator - The user who created the invite
 * @property acceptor - The user who accepted the invite (if accepted)
 * @property expiresAt - When the invite becomes invalid
 * @property role - Intended role for the invited user
 * @property teamIds - Teams the user will be added to upon acceptance
 */
export interface Invite {
  id: string
  email: string
  name: string | null
  organizationId: string
  creatorId: string
  acceptorId: string | null
  createdAt: Generated<Timestamp>
  expiresAt: Timestamp
  role: Generated<OrganizationRole>
  teamIds: Generated<string[]>
}

/**
 * Represents enhanced API authentication keys with organization-level ownership.
 * Used for authenticating API requests to Formbricks with more granular permissions.
 *
 * @property id - Unique identifier for the API key
 * @property label - Optional descriptive name for the key
 * @property hashedKey - Securely stored API key
 * @property organization - The organization this key belongs to
 * @property createdBy - User ID who created this key
 * @property lastUsedAt - Timestamp of last usage
 * @property apiKeyWorkspaces - Workspaces this key has access to
 */
export interface ApiKey {
  id: string
  createdAt: Generated<Timestamp>
  createdBy: string | null
  lastUsedAt: Timestamp | null
  label: string
  hashedKey: string
  lookupHash: string | null
  organizationId: string
  /** [OrganizationAccess] */
  organizationAccess: Generated<unknown>
}

/**
 * Links API keys to workspaces with specific permissions.
 * Enables granular access control for API keys across workspaces.
 *
 * @property id - Unique identifier for the workspace access entry
 * @property apiKey - The associated API key
 * @property workspace - The workspace being accessed
 * @property permission - Level of access granted
 */
export interface ApiKeyWorkspace {
  id: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  apiKeyId: string
  workspaceId: string
  permission: ApiKeyPermission
}

/**
 * Stores third-party authentication account information.
 * Enables OAuth and other external authentication methods.
 *
 * @property id - Unique identifier for the account
 * @property user - The Formbricks user who owns this account
 * @property provider - The authentication provider (GitHub, Google, etc.)
 * @property providerAccountId - User ID from the provider
 * @property access_token - OAuth access token
 * @property refresh_token - OAuth refresh token
 */
export interface Account {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  userId: string
  type: string | null
  provider: string
  providerAccountId: string
  access_token: string | null
  refresh_token: string | null
  expires_at: number | null
  ext_expires_in: number | null
  token_type: string | null
  scope: string | null
  id_token: string | null
  session_state: string | null
  password: string | null
  accessTokenExpiresAt: Timestamp | null
  refreshTokenExpiresAt: Timestamp | null
  /**
   * Better Auth 1.7.0 (ENG-2343): the account's identity namespace. 1.7.0 keyed accounts on
   * (issuer, accountId) and filtered every lookup on it — findAccountByKey, findAccountOwnerByKey,
   * findCredentialAccount and updatePassword — so on 1.7.0 a null here locks the user out.
   * 1.7.3 dropped the field and 1.7.7 keys accounts on (provider, providerAccountId) again
   * (ENG-3880). We keep writing it (see `account.create.before` in better-auth-hooks.ts) so a
   * rollback to 1.7.0 still finds every row; drop it once that rollback is off the table.
   *
   * NULLABLE on purpose. A NOT NULL column would fail every INSERT from a 1.6 pod still serving
   * during a rolling deploy. The migration backfills every existing row; see the migration SQL for
   * the value derivation.
   */
  issuer: string | null
}

/**
 * Stores active authentication sessions for revocable server-side login state.
 *
 * @property sessionToken - Opaque token stored in the browser cookie
 * @property user - The Formbricks user who owns this session
 * @property expires - Hard expiry for the session
 */
export interface Session {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  sessionToken: string
  userId: string
  expires: Timestamp
  ipAddress: string | null
  userAgent: string | null
}

/** Stores one-time verification tokens used by Auth.js adapter flows. */
export interface VerificationToken {
  identifier: string
  token: string
  expires: Timestamp
}

/**
 * Better Auth (ENG-1054) two-factor secrets. Supersedes the legacy User.twoFactorSecret /
 * User.backupCodes columns (kept until cutover). `secret` and `backupCodes` are encrypted
 * at rest by Better Auth using BETTER_AUTH_SECRET.
 */
export interface TwoFactor {
  id: string
  secret: string
  backupCodes: string
  userId: string
  verified: Generated<boolean>
  /** Better Auth 1.6.23+ two-factor brute-force lockout: consecutive failed verification attempts. */
  failedVerificationCount: Generated<number>
  /** Better Auth 1.6.23+ two-factor brute-force lockout: set once failedVerificationCount trips the limit. */
  lockedUntil: Timestamp | null
}

/** Better Auth JWT signing keys for OAuth/OIDC access and ID tokens. */
export interface jwks {
  id: string
  publicKey: string
  privateKey: string
  createdAt: Timestamp
  expiresAt: Timestamp | null
  /**
   * Better Auth 1.7 (ENG-2343): the JWS algorithm this key signs with, so a keyring can hold more
   * than one alg and `getLatestKeyByAlg` can pick the right key per resource.
   *
   * NOT optional upstream in practice, despite being declared `required: false`: `createJwk` always
   * puts `alg` in the row it writes (defaulting to EdDSA), so without this column every key mint is
   * rejected by Prisma as an unknown argument — which breaks JWT signing, and with it the whole MCP
   * OAuth flow, on any deployment that has yet to mint a key. Nullable because rows written before
   * 1.7 have no value; upstream reads treat `alg: null` as the configured default alg, so existing
   * keys keep working untouched.
   */
  alg: string | null
  /**
   * Better Auth 1.7 (ENG-2343): the curve for EC/OKP keys. Written whenever it can be derived from
   * the generated key, which it can for the default EdDSA/Ed25519 config we run, so it is on the
   * same always-written path as `alg` above.
   */
  crv: string | null
}

/** Better Auth OAuth 2.1 client registrations for MCP and future OAuth clients. */
export interface oauthClient {
  id: string
  clientId: string
  clientSecret: string | null
  disabled: Generated<boolean | null>
  skipConsent: boolean | null
  enableEndSession: boolean | null
  subjectType: string | null
  scopes: Generated<string[]>
  userId: string | null
  createdAt: Timestamp | null
  updatedAt: Timestamp | null
  name: string | null
  uri: string | null
  icon: string | null
  contacts: Generated<string[]>
  tos: string | null
  policy: string | null
  softwareId: string | null
  softwareVersion: string | null
  softwareStatement: string | null
  redirectUris: string[]
  postLogoutRedirectUris: Generated<string[]>
  tokenEndpointAuthMethod: string | null
  grantTypes: Generated<string[]>
  responseTypes: Generated<string[]>
  /**
   * Unwritten by Better Auth 1.7 (it dropped both from its schema), kept deliberately: they are
   * nullable, so keeping them makes the 1.7 migration purely additive and a rollback to 1.6 needs no
   * down-migration. Drop them in a separate cleanup after a bake period.
   */
  public: boolean | null
  type: string | null
  requirePKCE: boolean | null
  referenceId: string | null
  metadata: unknown
  /** Better Auth 1.7 (ENG-2343). */
  clientDiscoveryId: string | null
  clientCredentialsScopes: Generated<string[]>
  backchannelLogoutUri: string | null
  backchannelLogoutSessionRequired: boolean | null
  applicationType: string | null
  jwks: string | null
  jwksUri: string | null
  dpopBoundAccessTokens: Generated<boolean | null>
}

/**
 * Better Auth opaque access token records. The token column stores Better Auth's hashed lookup value.
 * JWT access tokens are verified through JWKS and bounded by the configured short lifetime.
 */
export interface oauthAccessToken {
  id: string
  token: string
  clientId: string
  sessionId: string | null
  userId: string | null
  referenceId: string | null
  refreshId: string | null
  expiresAt: Timestamp
  createdAt: Timestamp
  scopes: string[]
  /**
   * Better Auth 1.7 (ENG-2343). `resources` is the RFC 8707 resource set the grant approved — the
   * binding this whole upgrade exists to get (GHSA-p2fr-6hmx-4528).
   */
  authorizationCodeId: string | null
  resources: Generated<string[]>
  requestedUserInfoClaims: Generated<string[]>
  revoked: Timestamp | null
  confirmation: unknown
}

/** Better Auth OAuth refresh tokens used by delegated MCP clients. The token column stores a hashed lookup value. */
export interface oauthRefreshToken {
  id: string
  token: string
  clientId: string
  sessionId: string | null
  userId: string
  referenceId: string | null
  expiresAt: Timestamp
  createdAt: Timestamp
  revoked: Timestamp | null
  authTime: Timestamp | null
  scopes: string[]
  /**
   * Better Auth 1.7 (ENG-2343). `resources` carries the full original grant so a refresh can narrow
   * the resource set but never widen it.
   */
  authorizationCodeId: string | null
  resources: Generated<string[]>
  requestedUserInfoClaims: Generated<string[]>
  rotatedAt: Timestamp | null
  rotationReplayResponse: string | null
  rotationReplayExpiresAt: Timestamp | null
  confirmation: unknown
}

/** Better Auth OAuth user consent grants for delegated clients. */
export interface oauthConsent {
  id: string
  clientId: string
  userId: string | null
  referenceId: string | null
  scopes: string[]
  createdAt: Timestamp
  updatedAt: Timestamp
  /**
   * Better Auth 1.7 (ENG-2343). /authorize compares the requested resources against these; without
   * them every pre-existing connection is forced back through the consent screen.
   */
  resources: Generated<string[]>
  requestedUserInfoClaims: Generated<string[]>
}

/**
 * Better Auth 1.7 (ENG-2343): a protected resource the authorization server issues tokens for, with
 * its own token lifetimes, scope allow-list and claims. Replaces the flat `validAudiences` option —
 * a token is now bound to the resource it was approved for instead of to a static allow-list.
 * `identifier` is the RFC 8707 `resource` value and is the natural key other rows point at.
 */
export interface oauthResource {
  id: string
  identifier: string
  name: string
  accessTokenTtl: number | null
  refreshTokenTtl: number | null
  signingAlgorithm: string | null
  signingKeyId: string | null
  allowedScopes: Generated<string[]>
  customClaims: unknown
  dpopBoundAccessTokensRequired: Generated<boolean | null>
  disabled: Generated<boolean | null>
  createdAt: Timestamp | null
  updatedAt: Timestamp | null
  policyVersion: Generated<number | null>
  metadata: unknown
}

/**
 * Better Auth 1.7 (ENG-2343): links a client to the resources it may request tokens for. With
 * `enforcePerClientResources` (the default) a missing row means `invalid_target` at the token
 * endpoint, which is why the migration backfills one row per existing client.
 * Both sides reference natural keys, matching the other oauth child tables.
 */
export interface oauthClientResource {
  id: string
  clientId: string
  resourceId: string
  metadata: unknown
  createdAt: Timestamp | null
}

/**
 * Better Auth 1.7 (ENG-2343): replay store for `private_key_jwt` client assertions (the row id is the
 * assertion jti). We use no private_key_jwt client authentication today, but the plugin declares the
 * table, so omitting it risks an adapter error on a code path we do not control.
 */
export interface oauthClientAssertion {
  id: string
  expiresAt: Timestamp
}

/**
 * Stores the active password reset token for a user.
 * Tokens are opaque, hashed at rest, revocable, and single-use.
 */
export interface PasswordResetToken {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  token_hash: string
  expires_at: Timestamp
  userId: string
}

/**
 * Represents a user in the Formbricks system.
 * Central model for user authentication and profile management.
 *
 * @property id - Unique identifier for the user
 * @property name - Display name of the user
 * @property email - User's email address
 * @property twoFactorEnabled - Whether 2FA is active
 * @property memberships - Organizations the user belongs to
 * @property notificationSettings - User's notification preferences
 */
export interface User {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  name: string
  email: string
  email_verified: Generated<boolean>
  twoFactorSecret: string | null
  twoFactorEnabled: Generated<boolean>
  backupCodes: string | null
  password: string | null
  identityProvider: Generated<IdentityProvider>
  identityProviderAccountId: string | null
  groupId: string | null
  /** [UserNotificationSettings] */
  notificationSettings: Generated<unknown>
  /** [Locale] */
  locale: Generated<string>
  lastLoginAt: Timestamp | null
  isActive: Generated<boolean>
  /**
   * When an admin last reactivated this user. Restarts the members retention clock (ENG-3713)
   * without touching `lastLoginAt`, which the member list, usage counts and APIs report as a real
   * sign-in.
   */
  reactivatedAt: Timestamp | null
  /**
   * When the user last completed a sign-in (`lastActivePlugin`, after any second factor) or renewed a
   * session (Better Auth `databaseHooks.session.update`), on the database clock and never moved
   * backwards. The members retention clock reads it: sessions are
   * deleted on sign-out and expiry, so their own timestamps can't serve as a record of activity.
   */
  lastActiveAt: Timestamp | null
  isBootstrapAdmin: BootstrapAdminMarker | null
}

/**
 * Defines a segment of contacts based on attributes.
 * Used for targeting surveys to specific user groups.
 *
 * @property id - Unique identifier for the segment
 * @property title - Display name of the segment
 * @property filters - Rules defining the segment
 * @property isPrivate - Whether the segment is private
 * @property workspace - The workspace this segment belongs to
 */
export interface Segment {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  title: string
  description: string | null
  isPrivate: Generated<boolean>
  /** [SegmentFilter] */
  filters: Generated<unknown>
  workspaceId: string
}

/**
 * Represents a supported language in the system.
 * Used for multilingual survey support.
 *
 * @property id - Unique identifier for the language
 * @property code - Language code (e.g., 'en-US')
 * @property alias - Optional friendly name
 * @property workspace - The workspace this language is enabled for
 */
export interface Language {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  code: string
  alias: string | null
  workspaceId: string
}

/**
 * Links surveys to their supported languages.
 * Manages which languages are available for each survey.
 *
 * @property language - The supported language
 * @property survey - The associated survey
 * @property default - Whether this is the default language
 * @property enabled - Whether this language is currently active
 */
export interface SurveyLanguage {
  languageId: string
  surveyId: string
  default: Generated<boolean>
  enabled: Generated<boolean>
}

/**
 * Defines one Embedded Data field: the extra context stored on a response that isn't an
 * answer to a question (brand, plan, page type, a computed score, ...).
 *
 * A field is either **local** — used by a single survey, which is how existing variables and
 * hidden fields are migrated — or **shared**, meaning it is part of the workspace library and can
 * be linked by many surveys. There is no stored flag for that: `surveyId` is the owner and `key`
 * is the library name, so a local field has `surveyId` set and `key` null, and a shared field the
 * other way round. Deriving it rather than storing a boolean removes a column that could drift
 * out of step with the two fields it stands for.
 *
 * @property key - Library name, unique per workspace and immutable. Null for local fields
 * @property name - Human-readable display label
 * @property source - Where the value comes from (computed | ingested). `reserved` is part of the
 * shared enum but never stored as a row: those fields are a static catalog in code
 * @property dataType - Declared value type, used for coercion at ingest
 * @property defaultValue - Initial value (computed) or fallback when nothing arrives (ingested)
 * @property locked - Ingested only: ignore external writes and use the default instead
 * @property survey - The owning survey for a local field; cascades so nothing is orphaned
 */
export interface EmbeddedData {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  key: string | null
  name: string
  description: string | null
  source: EmbeddedDataSource
  dataType: Generated<EmbeddedDataType>
  /** [EmbeddedDataDefaultValue] */
  defaultValue: unknown
  locked: Generated<boolean>
  surveyId: string | null
  workspaceId: string
}

/**
 * Links a survey to an Embedded Data field and records the key that field's value is stored
 * under inside that survey.
 *
 * `workspaceId` is part of both foreign keys on purpose: `surveyId` and `embeddedDataId` alone are
 * independent, so nothing would stop a survey in one workspace linking a field defined in another.
 * Referencing `(id, workspaceId)` on both sides makes a cross-workspace pair unrepresentable, and
 * puts the workspace scope every query needs on the row instead of behind a join through Survey.
 *
 * @property storageKey - The key the value lives under: a cuid for computed fields, the field
 * name for ingested ones. Migrated fields keep their original key, which is what lets existing
 * recall tokens, logic conditions and stored responses keep resolving untouched.
 *
 * @property order - Position within this survey's field list, dense from 0. Lives on the link
 * rather than on `EmbeddedData` because a shared definition sits at a different position in each
 * survey that links it. Deliberately has no default: every writer knows the position it is writing,
 * so a missing one should be a type error rather than a row that silently sorts to the front.
 */
export interface SurveyEmbeddedData {
  id: string
  workspaceId: string
  surveyId: string
  embeddedDataId: string
  storageKey: string
  order: number
}

/**
 * Represents a team within an organization.
 * Enables group-based access control and collaboration.
 *
 * @property id - Unique identifier for the team
 * @property name - Display name of the team
 * @property organization - The parent organization
 * @property teamUsers - Users who are part of this team
 * @property workspaceTeams - Workspaces this team has access to
 */
export interface Team {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  name: string
  organizationId: string
}

/**
 * Links users to teams with specific roles.
 * Manages team membership and permissions.
 *
 * @property team - The associated team
 * @property user - The team member
 * @property role - User's role within the team
 */
export interface TeamUser {
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  teamId: string
  userId: string
  role: TeamUserRole
}

/**
 * Defines team access to specific workspaces.
 * Manages workspace-level permissions for teams.
 *
 * @property workspace - The accessed workspace
 * @property team - The team receiving access
 * @property permission - Level of access granted
 */
export interface WorkspaceTeam {
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  workspaceId: string
  teamId: string
  permission: Generated<WorkspaceTeamPermission>
}

/**
 * Represents a chart/visualization that can be used in multiple dashboards.
 * Charts are reusable components that query analytics data.
 *
 * @property id - Unique identifier for the chart
 * @property name - Display name of the chart
 * @property type - Type of visualization (bar, line, pie, etc.)
 * @property workspace - The workspace this chart belongs to
 * @property query - Cube.js query configuration (JSON)
 * @property config - Chart-specific configuration (colors, labels, etc.)
 * @property createdBy - User who created the chart
 * @property dashboards - Dashboards that use this chart
 */
export interface Chart {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  name: string
  type: ChartType
  workspaceId: string
  /** [ChartQuery] - Cube.js query configuration */
  query: Generated<unknown>
  /** [ChartConfig] - Visualization configuration (colors, labels, formatting) */
  config: Generated<unknown>
  createdBy: string | null
  feedbackDirectoryId: string
}

/**
 * Represents a dashboard containing multiple charts.
 * Dashboards aggregate analytics insights at the workspace level.
 *
 * @property id - Unique identifier for the dashboard
 * @property name - Display name of the dashboard
 * @property workspace - The workspace this dashboard belongs to
 * @property widgets - Charts on this dashboard
 * @property filters - The dashboard filters pinned to its top bar, in order: `{ id, field, operator, label }`,
 * at most 10 (ADR-0026). Each viewer's picks live in their own browser, not here.
 * @property createdBy - User who created the dashboard
 */
export interface Dashboard {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  name: string
  workspaceId: string
  createdBy: string | null
  filters: Generated<unknown>
}

/**
 * Represents a chart widget on a dashboard.
 * Widgets are positioned using a grid layout system.
 *
 * @property id - Unique identifier for the widget
 * @property dashboard - The dashboard this widget belongs to
 * @property chart - The chart displayed in this widget (chart name is used for display)
 * @property layout - Grid layout configuration (x, y, width, height)
 * @property order - Display order within the dashboard
 */
export interface DashboardWidget {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  dashboardId: string
  chartId: string
  /** [WidgetLayout] - Grid layout: { x, y, w, h } */
  layout: Generated<unknown>
  order: Generated<number>
}

/**
 * Work left over after a delete commits (ENG-3612, ENG-3614: deleted means deleted): Hub records and
 * storage files that can't be removed inside the database transaction. A row is written in the same
 * transaction as the delete, so a delete that commits always gets its cleanup and one that rolls back
 * queues nothing. A worker drains it with retries, never giving up, and deletes each row once done.
 */
export interface DeletionCleanup {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  kind: DeletionCleanupKind
  /** Scoping and logs. Not foreign keys: the rows they name are usually gone already. */
  organizationId: string
  workspaceId: string
  surveyId: string
  /** Hub kinds: the organisation's feedback directories, read inside the delete's transaction. */
  tenantIds: Generated<string[]>
  /** `hubResponses`: the deleted responses whose records go. `hubResponseOutcome`: exactly one live response. */
  responseIds: Generated<string[]>
  /** `storageFiles`: the object keys or URLs to delete. */
  fileKeys: Generated<string[]>
  attempts: Generated<number>
  /** Not drained before this, so a late pipeline write can't slip in behind a survey's cleanup. */
  nextAttemptAt: Generated<Timestamp>
  /** The last failure: a short code and message, never record or file content. */
  lastError: string | null
}

/**
 * A survey's response attachments built into one ZIP in object storage by a background job and
 * downloaded through a short-lived presigned URL, so an interrupted download can resume (ENG-6862).
 * Only the requester reads it. The reconcile job deletes the archive at `expiresAt`; deleting the
 * survey deletes it with the survey's storage prefix.
 *
 * A partial unique index in the migration allows one queued or running export per requester and
 * survey, so a double click returns the export already in flight.
 */
export interface AttachmentExport {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  surveyId: string
  workspaceId: string
  requestedById: string | null
  status: Generated<AttachmentExportStatus>
  /** The response filter at request time, a `TResponseFilterCriteria`; null exports every response. */
  filterCriteria: unknown
  /** IANA zone the archive's folder names are stamped in: the org's display zone at request time. */
  timeZone: string
  /**
   * Bumped on every enqueue and part of the job id, so a re-dispatch never dedupes against a job an
   * older pod dropped as unknown during a rolling upgrade.
   */
  dispatchCount: Generated<number>
  /** The pre-flight's counts, shown while the archive builds. */
  fileCount: number
  responseCount: number
  /** What the archive actually holds, once built. */
  includedFileCount: number | null
  totalBytes: bigint | null
  truncated: Generated<boolean>
  storageKey: string | null
  /** A short code (`stale`, `storage_upload_failed`, …), never a message or file name. */
  errorCode: string | null
  startedAt: Timestamp | null
  finishedAt: Timestamp | null
  expiresAt: Timestamp | null
}

/**
 * A feedback source — connects external data into Unify Feedback for record creation.
 * Naming note: "FeedbackSource" is the Unify Feedback domain term; the word "Connector"
 * is reserved for the Workflows domain (see docs/development/standards/practices/naming-conventions.mdx).
 *
 * @property id - Unique identifier for the feedback source
 * @property name - Display name for the feedback source
 * @property type - Type of feedback source (formbricks_survey, csv, …)
 * @property status - Current state of the feedback source (active, paused, error)
 * @property workspace - The workspace this feedback source belongs to
 * @property feedbackDirectory - Hub directory the feedback source writes into
 * @property formbricksMappings - Element mappings for Formbricks feedback sources
 * @property fieldMappings - Field mappings for other feedback source types
 */
export interface FeedbackSource {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  name: string
  type: FeedbackSourceType
  status: Generated<FeedbackSourceStatus>
  /**
   * Defaults to completedOnly so a new source agrees with the live ingestion path, which only
   * fires on responseFinished. Existing rows take the default; nothing is re-imported.
   */
  import_mode: Generated<FeedbackSourceImportMode>
  elementScope: Generated<FeedbackSourceElementScope>
  workspaceId: string
  feedbackDirectoryId: string
  last_sync_at: Timestamp | null
  created_by: string | null
}

/**
 * Maps survey elements to Hub FeedbackRecords for Formbricks feedback sources.
 * Each row represents one element that will create FeedbackRecords when answered.
 *
 * @property id - Unique identifier for the mapping
 * @property feedbackSource - The parent feedback source
 * @property survey - The survey containing the element
 * @property elementId - The element ID within the survey (from blocks[].elements[].id)
 * @property hubFieldType - The field_type to use in Hub (text, nps, rating, etc.)
 * @property customFieldLabel - Optional override for the element headline as field_label in Hub
 */
export interface FeedbackSourceFormbricksMapping {
  id: string
  created_at: Generated<Timestamp>
  feedback_source_id: string
  workspaceId: string
  surveyId: string
  elementId: string
  hubFieldType: HubFieldType
  custom_field_label: string | null
}

/**
 * Generic field mapping for non-Formbricks feedback sources (CSV today; webhook/email/slack later).
 * Maps source fields to Hub FeedbackRecord fields.
 *
 * @property id - Unique identifier for the mapping
 * @property feedbackSource - The parent feedback source
 * @property sourceFieldId - Field path for webhook (e.g., "user.id"), column name for CSV
 * @property targetFieldId - Hub field (collected_at, field_id, value_text, etc.)
 * @property staticValue - If set, use this value instead of reading from sourceFieldId
 */
export interface FeedbackSourceFieldMapping {
  id: string
  created_at: Generated<Timestamp>
  feedback_source_id: string
  workspaceId: string
  source_field_id: string
  target_field_id: string
  static_value: string | null
}

/**
 * Represents a feedback directory (Hub tenant) owned by an organization.
 * Directories group feedback data and are assigned to workspaces for access control.
 *
 * @property id - Unique identifier for the directory
 * @property name - Display name of the directory
 * @property isArchived - Soft delete flag
 * @property organization - The parent organization
 * @property workspaces - Workspaces assigned to this directory
 */
export interface FeedbackDirectory {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  name: string
  isArchived: Generated<boolean>
  organizationId: string
}

/**
 * Links feedback directories to workspaces (workspaces).
 * Manages which workspaces can access a given directory.
 *
 * @property feedbackDirectory - The directory being accessed
 * @property workspace - The workspace receiving access
 */
export interface FeedbackDirectoryWorkspace {
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  feedbackDirectoryId: string
  workspaceId: string
}

/**
 * ENG-3282: per-scope readiness of a SpiceDB projection. A row with `readyAt` set means the scope was
 * backfilled and audited clean, so the evaluator may enforce it. Written only by the backfill CLI.
 */
export interface AuthzedProjectionScopeState {
  scope: string
  readyAt: Timestamp | null
  readyBy: string | null
  updatedAt: Timestamp
}

/**
 * Durable PostgreSQL queue for rebuilding SpiceDB relationships after an authorization source changes.
 * Source-table triggers insert these rows in the same transaction as the mutation. Identifiers are
 * operationally sensitive and must never be included in logs, metrics, or CLI output.
 *
 * This model declares NO indexes on purpose. All three live in
 * `migration/20260818120000_add_authzed_projection_outbox/migration.sql`, because each one is
 * partial (`WHERE "processedAt" IS NULL AND ...`) and Prisma's `@@index` has no `where`. Declaring
 * approximate non-partial copies here would be worse than declaring none: `prisma db push` would
 * see a name match with a different definition, drop and recreate them without the predicate, and
 * the migration's `CREATE INDEX IF NOT EXISTS` would then skip them — leaving a silently wrong
 * shape that carries a week of delivered rows on every hot-path lookup. Deployments are unaffected
 * either way (`prisma migrate deploy` never reads this file); a dev who runs `db push` restores
 * them by rerunning the migration, which is the same contract its triggers already live under.
 */
export interface AuthzedProjectionOutbox {
  id: string
  targetType: string
  primaryId: string
  secondaryId: string | null
  isRevocation: Generated<boolean>
  /** Total delivery attempts. Drives the retry backoff and reports how hard delivery was tried. */
  attempts: Generated<number>
  /**
   * Non-retryable failures attributable to this event alone. The only input to dead-lettering, so
   * that a SpiceDB outage of any duration cannot dead-letter events that were never the problem.
   */
  permanentFailures: Generated<number>
  availableAt: Generated<Timestamp>
  leasedAt: Timestamp | null
  leaseExpiresAt: Timestamp | null
  leaseOwner: string | null
  processedAt: Timestamp | null
  deadLetteredAt: Timestamp | null
  lastAttemptAt: Timestamp | null
  lastErrorCode: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface OnlineUpgradeState {
  release: string
  targetImage: string
  phase: string
  completedPhases: Generated<string[]>
  surveyCursor: string | null
  surveyCount: Generated<bigint>
  baselineCompletedAt: Timestamp | null
  forwardOnlyAt: Timestamp | null
  queuePausedAt: Timestamp | null
  queuePauseDeadlineAt: Timestamp | null
  queueResumedAt: Timestamp | null
  activatedAt: Timestamp | null
  updatedAt: Generated<Timestamp>
}

/**
 * One policy per organisation and kind of data: a notice `warnDays` before the action, and the action
 * `periodDays` after the target's clock. The action follows from the kind of data: responses are
 * deleted, surveys archived (the archive purge deletes them 30 days later, as it does every archived
 * survey), members deactivated. The dates themselves are never stored; they are computed by the shared
 * helper in `apps/web/modules/ee/data-retention/lib/schedule.ts`, so the UI, the emails and the sweep
 * always agree.
 *
 * The allowed ranges are validated by the API. The migration repeats the ones that guard against
 * acting early or re-notifying (a notice of at least 14 days and shorter than the period, an enabled
 * policy always has `enabledAt`) and the structural ones (conditions on the surveys policy only) as
 * CHECK constraints, which Prisma cannot express.
 */
export interface RetentionPolicy {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  organizationId: string
  entity: RetentionEntity
  enabled: Generated<boolean>
  /**
   * When the current configuration took effect: the policy was switched on or unpaused, or changed
   * in a way that could bring an action earlier (a shorter period or warning, different
   * conditions). Lengthening does not move it. The warning always runs in full from here, so no
   * change can delete data on the night it is made (ENG-3614). Also the clock for members with no
   * recorded sign-in.
   */
  enabledAt: Timestamp | null
  /**
   * How long before the action the notice goes out. The action never runs before the notice has
   * been out for this long, so the warning always runs in full.
   */
  warnDays: Generated<number>
  /**
   * The period, counted from the target's clock, after which the policy acts: responses are
   * deleted, surveys archived (the archive purge deletes them 30 days later), members deactivated.
   */
  periodDays: number
  /** Surveys policy only. */
  conditions: Generated<RetentionSurveyCondition[]>
  /**
   * The last night the sweep's time ran out before it reached this policy's organisation, licensed.
   * The organisation was due and the sweep was running, so it counts as reached for the gap that
   * restarts the warning (`RETENTION_SWEEP_GAP_MS`): a backlog of organisations never restarts warnings.
   */
  deferredAt: Timestamp | null
  updatedById: string | null
}

/**
 * Holds one survey back from one policy until `until`. Keyed by organisation + entity rather than a
 * policy id, so removing and re-adding a policy doesn't orphan its exemptions.
 *
 * Active means `revokedAt IS NULL AND until > now()`, so expiry needs no job. "One active exemption
 * per survey and policy" is the partial unique index `RetentionExemption_surveyId_entity_active_key`
 * (`WHERE "revokedAt" IS NULL`) in
 * `migration/20261007120000_add_data_retention_tables/migration.sql`; Prisma's `@@unique` has no
 * `where`, so it is deliberately not declared here (see "Indexes Prisma cannot express" in the
 * README). An expired row still holds that slot until it is revoked, so creating a new exemption
 * must close an expired one in the same transaction.
 *
 * `surveyId` is nullable and a CHECK requires exactly one target, so a later `workflowId` is one more
 * column in that CHECK. The survey's organisation can't be tied to `organizationId` in the database
 * (a survey carries its workspace, not its organisation): writers derive it from the survey.
 */
export interface RetentionExemption {
  id: string
  created_at: Generated<Timestamp>
  organizationId: string
  entity: RetentionEntity
  surveyId: string | null
  until: Timestamp
  source: Generated<RetentionExemptionSource>
  /** Required for `manual`, absent for `restore` (`RetentionExemption_reason_check`). */
  reason: string | null
  createdById: string | null
  revokedAt: Timestamp | null
  revokedById: string | null
}

/**
 * One sweep of one policy for one organisation: the History tab's rows. `finishedAt` stays null
 * while the run is in progress, or if it never finished.
 */
export interface RetentionRun {
  id: string
  organizationId: string
  entity: RetentionEntity
  kind: Generated<RetentionRunKind>
  startedAt: Generated<Timestamp>
  finishedAt: Timestamp | null
  notifiedCount: Generated<number>
  /** Archived surveys, or deactivated members. */
  archivedCount: Generated<number>
  /** Deleted responses, or, on a purge run, deleted surveys. */
  deletedCount: Generated<number>
  skippedCount: Generated<number>
  /** False when the run changed nothing. History hides those by default. */
  hasChanges: Generated<boolean>
  /**
   * Where the run's scan of candidates stopped for want of time, so the next run resumes there instead
   * of rescanning the same first targets every night. Null when the scan reached the end.
   */
  scanCursor: string | null
}

/**
 * What a run did, one row per target: the History CSV. Ids, counts, names and the recipient's email,
 * never response content. `targetId` is not a foreign key on purpose: the target is often gone, and
 * `targetName` is a snapshot taken when the row is written so the CSV still says what it was.
 */
export interface RetentionRunItem {
  id: string
  runId: string
  targetType: RetentionTargetType
  targetId: string
  targetName: string | null
  action: RetentionRunItemAction
  /** 1, except for grouped response deletions, where it is the number of responses deleted. */
  count: Generated<number>
  /** Who was emailed, for `notified` rows. */
  recipient: string | null
  skipReason: RetentionSkipReason | null
}

/**
 * The notice for one target under one policy, so each notice goes out once: a survey or a member, or
 * the one-time responses reminder for a survey. It counts only once delivered, and is void if it was
 * claimed before the policy's `enabledAt`, before the target's latest exemption ended or (surveys and
 * members) before the target's clock: then a later cycle claims it again and sends a new one.
 *
 * A CHECK in the migration ties each entity to its target: `surveys` and `responses` → `surveyId`,
 * `members` → `userId`; another makes a sent email imply a delivery time. The survey notices that still
 * hold a clock are the partial index `RetentionNotice_surveys_clock_pending_idx` (`WHERE "entity" =
 * 'surveys' AND "clockAt" IS NOT NULL`) in
 * `migration/20261012120000_retention_notice_void_on_response_delete/migration.sql`, with the triggers
 * that clear `clockAt`; Prisma's `@@index` has no `where`, so it is deliberately not declared here (see
 * "Indexes Prisma cannot express" in the README).
 */
export interface RetentionNotice {
  id: string
  organizationId: string
  entity: RetentionEntity
  surveyId: string | null
  userId: string | null
  /** When the notice was claimed for sending. A claim that never got delivered can be taken over. */
  sentAt: Generated<Timestamp>
  /**
   * When the notice was handed to the mail transport, or recorded without one (no SMTP). The
   * warning only counts from a delivered notice, so a crash between claim and send can't let the
   * action run without it.
   */
  deliveredAt: Timestamp | null
  /** Whether an email actually went out. False when SMTP isn't configured. */
  emailSent: Generated<boolean>
  /**
   * Set by each claim; delivery is recorded only for the claim that sent the email, so a slow sender
   * can't mark a newer claim delivered.
   */
  claimToken: string | null
  /**
   * The target's clock the notice was computed for (surveys and members). A notice counts only while
   * the target's clock is still this one, so activity between the sweep's read and its claim voids
   * it. Null for the responses reminder, which is sent once per survey and follows no one clock, and
   * for a survey notice voided for good: a person deleted or moved one of its responses, or the
   * responses policy's deletion left it computed for another clock (see `NOTICE_FOLLOWS_CLOCK`).
   */
  clockAt: Timestamp | null
}

/**
 * When a survey was last restored from the archive. A restore starts the surveys policy's "no response"
 * and "no change" periods over (`getSurveyRetentionClock`); only ever moves forward. A table of its own
 * so retention adds nothing to the survey row.
 */
export interface RetentionSurveyRestore {
  surveyId: string
  restoredAt: Timestamp
}

export interface Workflow {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  name: string
  description: string | null
  status: Generated<WorkflowStatus>
  workspaceId: string
  createdBy: string | null
  /** [WorkflowDefinition] */
  definition: unknown
}

export interface WorkflowVersion {
  id: string
  workflowId: string
  workspaceId: string
  version: number
  /** [WorkflowExecutableDefinition] */
  definition: unknown
  publishedAt: Generated<Timestamp>
  publishedBy: string | null
}

export interface WorkflowRun {
  id: string
  created_at: Generated<Timestamp>
  updated_at: Timestamp
  workflowId: string
  workspaceId: string
  workflowVersionId: string | null
  responseId: string | null
  status: Generated<WorkflowRunStatus>
  triggerType: string
  surveyId: string | null
  isDryRun: Generated<boolean>
  idempotencyKey: string | null
  attempt: Generated<number>
  nextAttemptAt: Timestamp | null
  lastErrorAt: Timestamp | null
  /** [WorkflowTriggerRunPayload] */
  triggerPayload: unknown
  /** [WorkflowRunData] */
  data: Generated<unknown>
  error: string | null
  dispatchedAt: Timestamp | null
  startedAt: Timestamp | null
  finishedAt: Timestamp | null
}

export interface WorkflowRunLog {
  id: string
  runId: string
  sequence: number
  stepId: string
  stepType: string
  status: Generated<WorkflowRunLogStatus>
  /** [WorkflowRunLogInput] */
  input: Generated<unknown>
  /** [WorkflowRunLogOutput] */
  output: Generated<unknown>
  error: string | null
  startedAt: Timestamp | null
  finishedAt: Timestamp | null
}

export interface DB {
  Creative: Creative
  CreativeDailyStat: CreativeDailyStat
  CreativeDailyPageStat: CreativeDailyPageStat
  CreativeTrigger: CreativeTrigger
  CreativeLanguage: CreativeLanguage
  OutboundEmailAudit: OutboundEmailAudit
  OutboundEmailAuditErasureIntent: OutboundEmailAuditErasureIntent
  Webhook: Webhook
  ContactAttribute: ContactAttribute
  ContactAttributeKey: ContactAttributeKey
  Contact: Contact
  Response: Response
  Tag: Tag
  TagsOnResponses: TagsOnResponses
  Display: Display
  SurveyTrigger: SurveyTrigger
  SurveyAttributeFilter: SurveyAttributeFilter
  Survey: Survey
  SurveyQuota: SurveyQuota
  ResponseQuotaLink: ResponseQuotaLink
  SurveyFollowUp: SurveyFollowUp
  ActionClass: ActionClass
  Integration: Integration
  DataMigration: DataMigration
  Workspace: Workspace
  Organization: Organization
  OrganizationBilling: OrganizationBilling
  Membership: Membership
  Invite: Invite
  ApiKey: ApiKey
  ApiKeyWorkspace: ApiKeyWorkspace
  Account: Account
  Session: Session
  VerificationToken: VerificationToken
  TwoFactor: TwoFactor
  jwks: jwks
  oauthClient: oauthClient
  oauthAccessToken: oauthAccessToken
  oauthRefreshToken: oauthRefreshToken
  oauthConsent: oauthConsent
  oauthResource: oauthResource
  oauthClientResource: oauthClientResource
  oauthClientAssertion: oauthClientAssertion
  PasswordResetToken: PasswordResetToken
  User: User
  Segment: Segment
  Language: Language
  SurveyLanguage: SurveyLanguage
  EmbeddedData: EmbeddedData
  SurveyEmbeddedData: SurveyEmbeddedData
  Team: Team
  TeamUser: TeamUser
  WorkspaceTeam: WorkspaceTeam
  Chart: Chart
  Dashboard: Dashboard
  DashboardWidget: DashboardWidget
  DeletionCleanup: DeletionCleanup
  AttachmentExport: AttachmentExport
  FeedbackSource: FeedbackSource
  FeedbackSourceFormbricksMapping: FeedbackSourceFormbricksMapping
  FeedbackSourceFieldMapping: FeedbackSourceFieldMapping
  FeedbackDirectory: FeedbackDirectory
  FeedbackDirectoryWorkspace: FeedbackDirectoryWorkspace
  AuthzedProjectionScopeState: AuthzedProjectionScopeState
  AuthzedProjectionOutbox: AuthzedProjectionOutbox
  OnlineUpgradeState: OnlineUpgradeState
  RetentionPolicy: RetentionPolicy
  RetentionExemption: RetentionExemption
  RetentionRun: RetentionRun
  RetentionRunItem: RetentionRunItem
  RetentionNotice: RetentionNotice
  RetentionSurveyRestore: RetentionSurveyRestore
  Workflow: Workflow
  WorkflowVersion: WorkflowVersion
  WorkflowRun: WorkflowRun
  WorkflowRunLog: WorkflowRunLog
}
