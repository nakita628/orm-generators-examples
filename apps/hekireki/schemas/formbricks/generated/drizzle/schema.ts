import {
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
import { createId } from '@paralleldrive/cuid2'

const utcDate = customType<{ data: Date; driverData: string }>({
  dataType: () => 'date',
  toDriver: (value) => value.toISOString().slice(0, 10),
  fromDriver: (value) => new Date(value),
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

/**
 * Creatives (ENG-3993): page elements the web SDK renders next to surveys. Every type lives in
 * `Creative`; `content` and `styling` are JSON validated per type by the API v3 schemas, so a new type is
 * a new payload variant, not new columns.
 */
export const creativeTypeEnum = pgEnum('CreativeType', ['feedbackButton', 'banner'])

/**
 * The author's intent only. Scheduled, Active, "Paused with its survey" and Closed after the
 * end time are derived (`apps/web/modules/ee/creatives/lib/display-status.ts` and its SQL twin); nothing
 * writes this column when a schedule ends or a survey pauses.
 */
export const creativeStatusEnum = pgEnum('CreativeStatus', [
  'draft',
  'inProgress',
  'paused',
  'completed',
])

/**
 * What makes a creative appear. A creative left with no action never shows; it never falls back to
 * every page view.
 */
export const creativeShowOnEnum = pgEnum('CreativeShowOn', ['pageView', 'action'])

/** The device type counts are split by: the values of the Show when `deviceType` operand. */
export const creativeDeviceTypeEnum = pgEnum('CreativeDeviceType', ['phone', 'tablet', 'desktop'])

export const pipelineTriggersEnum = pgEnum('PipelineTriggers', [
  'responseCreated',
  'responseUpdated',
  'responseFinished',
])

export const webhookSourceEnum = pgEnum('WebhookSource', [
  'user',
  'zapier',
  'make',
  'n8n',
  'activepieces',
])

export const contactAttributeTypeEnum = pgEnum('ContactAttributeType', ['default', 'custom'])

export const contactAttributeDataTypeEnum = pgEnum('ContactAttributeDataType', [
  'string',
  'number',
  'date',
])

/** What a reviewer decided about a response. One value today; null on `Response.outcome` means none. */
export const responseOutcomeEnum = pgEnum('ResponseOutcome', ['lowQuality'])

export const surveyStatusEnum = pgEnum('SurveyStatus', [
  'draft',
  'inProgress',
  'paused',
  'completed',
])

export const surveyAttributeFilterConditionEnum = pgEnum('SurveyAttributeFilterCondition', [
  'equals',
  'notEquals',
])

export const surveyQuotaActionEnum = pgEnum('SurveyQuotaAction', ['endSurvey', 'continueSurvey'])

export const responseQuotaLinkStatusEnum = pgEnum('ResponseQuotaLinkStatus', [
  'screenedIn',
  'screenedOut',
])

export const surveyTypeEnum = pgEnum('SurveyType', ['link', 'app'])

export const displayOptionsEnum = pgEnum('displayOptions', [
  'displayOnce',
  'displayMultiple',
  'displaySome',
  'respondMultiple',
])

/**
 * Who may read a survey. `workspace` is everyone with workspace access (the behaviour before ENG-3282);
 * `restricted` is the owner and the organization's owners and managers only.
 */
export const surveyVisibilityEnum = pgEnum('SurveyVisibility', ['restricted', 'workspace'])

export const surveyScriptModeEnum = pgEnum('SurveyScriptMode', ['add', 'replace'])

export const actionTypeEnum = pgEnum('ActionType', ['code', 'noCode'])

export const integrationTypeEnum = pgEnum('IntegrationType', [
  'googleSheets',
  'notion',
  'airtable',
  'slack',
])

export const dataMigrationStatusEnum = pgEnum('DataMigrationStatus', [
  'pending',
  'applied',
  'failed',
])

export const widgetPlacementEnum = pgEnum('WidgetPlacement', [
  'bottomLeft',
  'bottomRight',
  'topLeft',
  'topRight',
  'center',
])

export const surveyOverlayEnum = pgEnum('SurveyOverlay', ['none', 'light', 'dark'])

/**
 * The bootstrap marker's only value (ENG-2247). A one-variant enum rather than a boolean on
 * purpose: a unique index over a nullable boolean has a slot for `true` AND one for `false`, so the
 * first `false` ever written would take the second slot and break every later sign-up. There is no
 * `false` to write here, which is the same guarantee without a CHECK constraint Prisma cannot model.
 */
export const bootstrapAdminMarkerEnum = pgEnum('BootstrapAdminMarker', ['bootstrapAdmin'])

export const organizationRoleEnum = pgEnum('OrganizationRole', [
  'owner',
  'manager',
  'member',
  'billing',
  'dashboard_viewer',
])

/**
 * Defines permission levels for API keys.
 * Controls what operations an API key can perform.
 */
export const apiKeyPermissionEnum = pgEnum('ApiKeyPermission', ['read', 'write', 'manage'])

export const identityProviderEnum = pgEnum('IdentityProvider', [
  'email',
  'github',
  'google',
  'azuread',
  'openid',
  'saml',
])

export const mcpGrantAccessEnum = pgEnum('McpGrantAccess', ['all', 'organizations', 'workspaces'])

/**
 * Where an Embedded Data field's value comes from.
 * `computed` is set by survey logic, `ingested` arrives from a URL param or the SDK, and
 * `reserved` is auto-captured system metadata. Reserved fields are never stored as rows —
 * they are a static catalog in code — but the value exists so every layer shares one type.
 */
export const embeddedDataSourceEnum = pgEnum('EmbeddedDataSource', [
  'computed',
  'ingested',
  'reserved',
])

/** The declared type of an Embedded Data field's value. */
export const embeddedDataTypeEnum = pgEnum('EmbeddedDataType', [
  'string',
  'number',
  'boolean',
  'date',
])

export const teamUserRoleEnum = pgEnum('TeamUserRole', ['admin', 'contributor'])

export const workspaceTeamPermissionEnum = pgEnum('WorkspaceTeamPermission', [
  'read',
  'readWrite',
  'manage',
])

export const chartTypeEnum = pgEnum('ChartType', ['area', 'bar', 'pie', 'big_number', 'matrix'])

export const feedbackSourceTypeEnum = pgEnum('FeedbackSourceType', ['formbricks_survey', 'csv'])

export const feedbackSourceStatusEnum = pgEnum('FeedbackSourceStatus', [
  'active',
  'paused',
  'error',
])

/**
 * Which survey responses a Formbricks feedback source imports into Hub.
 * completedOnly matches the live ingestion path, which only runs on responseFinished.
 * all additionally imports partial responses — answers a respondent typed but never submitted.
 */
export const feedbackSourceImportModeEnum = pgEnum('FeedbackSourceImportMode', [
  'completedOnly',
  'all',
])

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
export const feedbackSourceElementScopeEnum = pgEnum('FeedbackSourceElementScope', [
  'all',
  'specific',
])

export const hubFieldTypeEnum = pgEnum('HubFieldType', [
  'text',
  'categorical',
  'nps',
  'csat',
  'ces',
  'rating',
  'number',
  'boolean',
  'date',
])

export const deletionCleanupKindEnum = pgEnum('DeletionCleanupKind', [
  'hubSurvey',
  'hubResponses',
  'storageFiles',
  'storageSurveyFolder',
  'hubResponseOutcome',
])

export const attachmentExportStatusEnum = pgEnum('AttachmentExportStatus', [
  'queued',
  'running',
  'completed',
  'failed',
  'expired',
])

/** Data retention (ENG-3697). The kind of data a policy governs. */
export const retentionEntityEnum = pgEnum('RetentionEntity', ['responses', 'surveys', 'members'])

/** The surveys policy's conditions. Every ticked one must hold, and they share the policy's period. */
export const retentionSurveyConditionEnum = pgEnum('RetentionSurveyCondition', [
  'noResponse',
  'noChange',
  'createdBefore',
])

/**
 * What wrote a `RetentionRun`: the nightly sweep, or the archive purge recording the deletion of surveys
 * the surveys policy archived. Only sweep runs count as the sweep having run (its lease, its warning
 * restart after a gap, where its scan resumes, the health check).
 */
export const retentionRunKindEnum = pgEnum('RetentionRunKind', ['sweep', 'purge'])

/**
 * Who made an exemption: a person (with a reason), or a restore from the archive under a surveys policy
 * whose only condition is "created before" (`restartSurveyRetention`), which carries no reason.
 */
export const retentionExemptionSourceEnum = pgEnum('RetentionExemptionSource', [
  'manual',
  'restore',
])

/**
 * What a `RetentionRunItem` row points at. Response deletions are grouped into one `survey` row per
 * survey, with the number of responses in `count`.
 */
export const retentionTargetTypeEnum = pgEnum('RetentionTargetType', ['survey', 'user'])

export const retentionRunItemActionEnum = pgEnum('RetentionRunItemAction', [
  'notified',
  'archived',
  'deactivated',
  'deleted',
  'skipped',
])

export const retentionSkipReasonEnum = pgEnum('RetentionSkipReason', [
  'exempt',
  'lastOwner',
  'otherOrganization',
  'noRecipient',
])

export const workflowStatusEnum = pgEnum('WorkflowStatus', [
  'draft',
  'enabled',
  'disabled',
  'archived',
])

export const workflowRunStatusEnum = pgEnum('WorkflowRunStatus', [
  'queued',
  'running',
  'completed',
  'failed',
  'canceled',
])

export const workflowRunLogStatusEnum = pgEnum('WorkflowRunLogStatus', [
  'pending',
  'running',
  'succeeded',
  'failed',
  'skipped',
])

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
export const creative = pgTable(
  'Creative',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    name: text('name').notNull(),
    type: creativeTypeEnum('type').notNull(),
    status: creativeStatusEnum('status').notNull().default('draft'),
    publishedAt: timestamp('publishedAt', { precision: 3 }),
    archivedAt: timestamp('archivedAt', { precision: 3 }),
    workspaceId: text('workspaceId').notNull(),
    createdBy: text('createdBy'),
    content: jsonb('content').notNull(),
    styling: jsonb('styling').notNull(),
    surveyId: text('surveyId'),
    linkUrl: text('linkUrl'),
    showOn: creativeShowOnEnum('showOn').notNull().default('pageView'),
    showConditions: jsonb('showConditions'),
    displayPercentage: doublePrecision('displayPercentage'),
    startAt: timestamp('startAt', { precision: 3 }),
    endAt: timestamp('endAt', { precision: 3 }),
    autoSelectLanguage: boolean('autoSelectLanguage'),
  },
  (table) => [
    index('Creative_workspaceId_updated_at_idx').on(table.workspaceId, table.updatedAt),
    index('Creative_workspaceId_created_at_idx').on(table.workspaceId, table.createdAt),
    index('Creative_workspaceId_name_idx').on(table.workspaceId, table.name),
    index('Creative_surveyId_idx').on(table.surveyId),
    foreignKey({
      name: 'Creative_workspaceId_fkey',
      columns: [table.workspaceId],
      foreignColumns: [workspace.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Creative_createdBy_fkey',
      columns: [table.createdBy],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Creative_surveyId_workspaceId_fkey',
      columns: [table.surveyId, table.workspaceId],
      foreignColumns: [survey.id, survey.workspaceId],
    })
      .onDelete('no action')
      .onUpdate('no action'),
  ],
)

/**
 * A creative's Impressions and Clicks on one UTC day on one device type (ENG-4000). Written
 * only by the counter flush job as `GREATEST(stored, counted)`, from Redis counters that hold the day's
 * running totals, so a repeated flush changes nothing. At most 3 rows per creative per day. Deleting
 * the creative deletes its counts.
 */
export const creativeDailyStat = pgTable(
  'CreativeDailyStat',
  {
    creativeId: text('creativeId').notNull(),
    /** UTC day. */
    date: utcDate('date').notNull(),
    deviceType: creativeDeviceTypeEnum('deviceType').notNull(),
    impressions: integer('impressions').notNull().default(0),
    clicks: integer('clicks').notNull().default(0),
  },
  (table) => [
    primaryKey({
      name: 'CreativeDailyStat_pkey',
      columns: [table.creativeId, table.date, table.deviceType],
    }),
    foreignKey({
      name: 'CreativeDailyStat_creativeId_fkey',
      columns: [table.creativeId],
      foreignColumns: [creative.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * A creative's Impressions and Clicks on one UTC day on one page (host + path), written like
 * `CreativeDailyStat`. At most 1000 pages per creative per day; the rest are counted under `(other)`.
 */
export const creativeDailyPageStat = pgTable(
  'CreativeDailyPageStat',
  {
    creativeId: text('creativeId').notNull(),
    /** UTC day. */
    date: utcDate('date').notNull(),
    /** Host + path, without query string or hash, or `(other)`. */
    page: text('page').notNull(),
    impressions: integer('impressions').notNull().default(0),
    clicks: integer('clicks').notNull().default(0),
  },
  (table) => [
    primaryKey({
      name: 'CreativeDailyPageStat_pkey',
      columns: [table.creativeId, table.date, table.page],
    }),
    foreignKey({
      name: 'CreativeDailyPageStat_creativeId_fkey',
      columns: [table.creativeId],
      foreignColumns: [creative.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/** Links a creative to the actions that show it (`showOn: action`). Mirrors `SurveyTrigger`. */
export const creativeTrigger = pgTable(
  'CreativeTrigger',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    creativeId: text('creativeId').notNull(),
    actionClassId: text('actionClassId').notNull(),
  },
  (table) => [
    index('CreativeTrigger_actionClassId_idx').on(table.actionClassId),
    uniqueIndex('CreativeTrigger_creativeId_actionClassId_key').on(
      table.creativeId,
      table.actionClassId,
    ),
    foreignKey({
      name: 'CreativeTrigger_creativeId_fkey',
      columns: [table.creativeId],
      foreignColumns: [creative.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'CreativeTrigger_actionClassId_fkey',
      columns: [table.actionClassId],
      foreignColumns: [actionClass.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/** The languages a creative's texts are written in. Mirrors `SurveyLanguage`; exactly one is the default. */
export const creativeLanguage = pgTable(
  'CreativeLanguage',
  {
    languageId: text('languageId').notNull(),
    creativeId: text('creativeId').notNull(),
    default: boolean('default').notNull().default(false),
    enabled: boolean('enabled').notNull().default(true),
  },
  (table) => [
    primaryKey({ name: 'CreativeLanguage_pkey', columns: [table.languageId, table.creativeId] }),
    index('CreativeLanguage_creativeId_idx').on(table.creativeId),
    foreignKey({
      name: 'CreativeLanguage_languageId_fkey',
      columns: [table.languageId],
      foreignColumns: [language.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'CreativeLanguage_creativeId_fkey',
      columns: [table.creativeId],
      foreignColumns: [creative.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * Metadata for an outbound email attempt. Content and recipient addresses are never stored here.
 * Identifiers are deliberately not foreign keys so erasure can scrub a row without retaining source data.
 */
export const outboundEmailAudit = pgTable(
  'OutboundEmailAudit',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => crypto.randomUUID()),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    expiresAt: timestamp('expiresAt', { precision: 3 }).notNull(),
    contentExpiresAt: timestamp('contentExpiresAt', { precision: 3 }).notNull(),
    environment: text('environment').notNull(),
    emailType: text('emailType').notNull(),
    trigger: text('trigger').notNull(),
    actorKind: text('actorKind').notNull(),
    actorUserId: text('actorUserId'),
    configuredByUserId: text('configuredByUserId'),
    actorUserHash: text('actorUserHash'),
    configuredByUserHash: text('configuredByUserHash'),
    organizationId: text('organizationId'),
    surveyId: text('surveyId'),
    responseId: text('responseId'),
    followUpId: text('followUpId'),
    workflowId: text('workflowId'),
    workflowRunId: text('workflowRunId'),
    inviteId: text('inviteId'),
    recipientMac: text('recipientMac'),
    sesMessageHash: text('sesMessageHash'),
    snapshotKey: text('snapshotKey'),
    snapshotSha256: text('snapshotSha256'),
    status: text('status').notNull().default('prepared'),
    statusRecordedAt: timestamp('statusRecordedAt', { precision: 3 }),
    auditGap: boolean('auditGap').notNull().default(false),
    erasureRequestedAt: timestamp('erasureRequestedAt', { precision: 3 }),
    erasedAt: timestamp('erasedAt', { precision: 3 }),
    contentDeleteRetryAt: timestamp('contentDeleteRetryAt', { precision: 3 }),
    contentDeleteFailures: integer('contentDeleteFailures').notNull().default(0),
  },
  (table) => [
    index('OutboundEmailAudit_environment_createdAt_id_idx').on(
      table.environment,
      table.createdAt,
      table.id,
    ),
    index('OutboundEmailAudit_organizationId_createdAt_id_idx').on(
      table.organizationId,
      table.createdAt,
      table.id,
    ),
    index('OutboundEmailAudit_surveyId_createdAt_id_idx').on(
      table.surveyId,
      table.createdAt,
      table.id,
    ),
    index('OutboundEmailAudit_responseId_createdAt_id_idx').on(
      table.responseId,
      table.createdAt,
      table.id,
    ),
    index('OutboundEmailAudit_actorUserId_createdAt_id_idx').on(
      table.actorUserId,
      table.createdAt,
      table.id,
    ),
    index('OutboundEmailAudit_configuredByUserId_createdAt_id_idx').on(
      table.configuredByUserId,
      table.createdAt,
      table.id,
    ),
    index('OutboundEmailAudit_actorUserHash_createdAt_id_idx').on(
      table.actorUserHash,
      table.createdAt,
      table.id,
    ),
    index('OutboundEmailAudit_configuredByUserHash_createdAt_id_idx').on(
      table.configuredByUserHash,
      table.createdAt,
      table.id,
    ),
    index('OutboundEmailAudit_recipientMac_createdAt_id_idx').on(
      table.recipientMac,
      table.createdAt,
      table.id,
    ),
    index('OutboundEmailAudit_sesMessageHash_idx').on(table.sesMessageHash),
    index('OutboundEmailAudit_status_createdAt_id_idx').on(table.status, table.createdAt, table.id),
    index('OutboundEmailAudit_contentExpiresAt_id_idx').on(table.contentExpiresAt, table.id),
    index('OutboundEmailAudit_expiresAt_id_idx').on(table.expiresAt, table.id),
    index('OutboundEmailAudit_erasureRequestedAt_id_idx').on(table.erasureRequestedAt, table.id),
    index('OutboundEmailAudit_contentDeleteRetryAt_id_idx').on(
      table.contentDeleteRetryAt,
      table.id,
    ),
  ],
)

/**
 * Holds an account erasure request across an in-flight email send. The address is never stored.
 * The user id is removed once deletion is confirmed; one-way recovery keys expire after 90 days.
 */
export const outboundEmailAuditErasureIntent = pgTable(
  'OutboundEmailAuditErasureIntent',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => crypto.randomUUID()),
    userId: text('userId'),
    userHash: text('userHash').notNull(),
    recipientMac: text('recipientMac').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    activatedAt: timestamp('activatedAt', { precision: 3 }),
    checkedAt: timestamp('checkedAt', { precision: 3 }),
    expiresAt: timestamp('expiresAt', { precision: 3 }).notNull(),
  },
  (table) => [
    index('OutboundEmailAuditErasureIntent_userId_idx').on(table.userId),
    index('OutboundEmailAuditErasureIntent_recipientMac_activatedAt_idx').on(
      table.recipientMac,
      table.activatedAt,
    ),
    index('OutboundEmailAuditErasureIntent_userHash_activatedAt_idx').on(
      table.userHash,
      table.activatedAt,
    ),
    index('OutboundEmailAuditErasureIntent_checkedAt_id_idx').on(table.checkedAt, table.id),
    index('OutboundEmailAuditErasureIntent_expiresAt_idx').on(table.expiresAt),
  ],
)

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
export const webhook = pgTable(
  'Webhook',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    name: text('name'),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    url: text('url').notNull(),
    source: webhookSourceEnum('source').notNull().default('user'),
    workspaceId: text('workspaceId').notNull(),
    triggers: pipelineTriggersEnum('triggers').array(),
    surveyIds: text('surveyIds').array(),
    secret: text('secret'),
  },
  (table) => [
    index('Webhook_workspaceId_idx').on(table.workspaceId),
    foreignKey({
      name: 'Webhook_workspaceId_fkey',
      columns: [table.workspaceId],
      foreignColumns: [workspace.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const contactAttribute = pgTable(
  'ContactAttribute',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    attributeKeyId: text('attributeKeyId').notNull(),
    contactId: text('contactId').notNull(),
    value: text('value').notNull(),
    valueNumber: doublePrecision('valueNumber'),
    valueDate: timestamp('valueDate', { precision: 3 }),
  },
  (table) => [
    index('ContactAttribute_attributeKeyId_value_idx').on(table.attributeKeyId, table.value),
    index('ContactAttribute_attributeKeyId_valueNumber_idx').on(
      table.attributeKeyId,
      table.valueNumber,
    ),
    index('ContactAttribute_attributeKeyId_valueDate_idx').on(
      table.attributeKeyId,
      table.valueDate,
    ),
    uniqueIndex('ContactAttribute_contactId_attributeKeyId_key').on(
      table.contactId,
      table.attributeKeyId,
    ),
    foreignKey({
      name: 'ContactAttribute_attributeKeyId_fkey',
      columns: [table.attributeKeyId],
      foreignColumns: [contactAttributeKey.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'ContactAttribute_contactId_fkey',
      columns: [table.contactId],
      foreignColumns: [contact.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const contactAttributeKey = pgTable(
  'ContactAttributeKey',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    isUnique: boolean('isUnique').notNull().default(false),
    key: text('key').notNull(),
    name: text('name'),
    description: text('description'),
    type: contactAttributeTypeEnum('type').notNull().default('custom'),
    dataType: contactAttributeDataTypeEnum('dataType').notNull().default('string'),
    workspaceId: text('workspaceId').notNull(),
  },
  (table) => [
    index('ContactAttributeKey_workspaceId_created_at_idx').on(table.workspaceId, table.createdAt),
    uniqueIndex('ContactAttributeKey_key_workspaceId_key').on(table.key, table.workspaceId),
    foreignKey({
      name: 'ContactAttributeKey_workspaceId_fkey',
      columns: [table.workspaceId],
      foreignColumns: [workspace.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const contact = pgTable(
  'Contact',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    workspaceId: text('workspaceId').notNull(),
  },
  (table) => [
    index('Contact_workspaceId_idx').on(table.workspaceId),
    foreignKey({
      name: 'Contact_workspaceId_fkey',
      columns: [table.workspaceId],
      foreignColumns: [workspace.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const response = pgTable(
  'Response',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow)
      .$onUpdate(utcNow),
    finished: boolean('finished').notNull().default(false),
    surveyId: text('surveyId').notNull(),
    contactId: text('contactId'),
    endingId: text('endingId'),
    /** [ResponseData] */
    data: jsonb('data').notNull().default({}),
    /** [ResponseVariables] */
    variables: jsonb('variables').notNull().default({}),
    /** [ResponseTtc] */
    ttc: jsonb('ttc').notNull().default({}),
    /** [ResponseMeta] */
    meta: jsonb('meta').notNull().default({}),
    /**
     * Why a value in `data` will not read back as its declared Embedded Data type — a failed coercion
     * or a truncated oversize value (ENG-1845). Server-computed on every ingest and never accepted
     * from the client. Null means no ingest boundary has written it, which is not the same claim as
     * an empty list: without this column a wrong-typed value is indistinguishable from an absent one,
     * and surfacing that in the UI (v2) would have nothing to read.
     * [ResponseIngestFlags]
     */
    ingestFlags: jsonb('ingest_flags'),
    /**
     * What the Cal.com elements reported when the respondent booked, `{ [Question ID]: booking }`. `data`
     * keeps the string "booked" for those questions, so this column is purely additive: it is omitted from
     * every read without a `select` (`PRISMA_GLOBAL_OMIT`), and only the webhook body (as
     * `calBookingsByQuestionId`) and `/api/v3/responses` read it. Server-derived from the incoming `data`,
     * never accepted as a field of its own.
     * [ResponseCalBookings]
     */
    calBookings: jsonb('cal_bookings_by_question_id'),
    /**
     * A reviewer's verdict on the response; null means none. `lowQuality` hides it from the Analysis
     * export by default and from Feedback Records, and nothing else: it is never deleted and every
     * count, list, webhook and integration still includes it (ENG-6780).
     */
    outcome: responseOutcomeEnum('outcome'),
    /**
     * What the respondent was not shown (ENG-6936, ADR 0029): the blocks and elements a show-if kept
     * from them, and per element its options, matrix rows and columns, by stable ids. Written by the
     * survey while they answer; null means unknown (older responses and bundles), read as before.
     * Omitted from reads without a `select` (PRISMA_GLOBAL_OMIT), so it never reaches a v1/v2 payload,
     * webhook or integration.
     * [ResponseNotShown]
     */
    notShown: jsonb('not_shown'),
    /** [ResponseContactAttributes] */
    contactAttributes: jsonb('contactAttributes'),
    singleUseId: text('singleUseId'),
    language: text('language'),
    displayId: text('displayId'),
  },
  (table) => [
    uniqueIndex('Response_displayId_key').on(table.displayId),
    index('Response_created_at_idx').on(table.createdAt),
    index('Response_created_at_id_idx').on(table.createdAt, table.id),
    index('Response_surveyId_created_at_idx').on(table.surveyId, table.createdAt),
    index('Response_surveyId_created_at_id_idx').on(table.surveyId, table.createdAt, table.id),
    index('Response_contactId_created_at_idx').on(table.contactId, table.createdAt),
    index('Response_contactId_created_at_id_idx').on(table.contactId, table.createdAt, table.id),
    index('Response_surveyId_finished_idx').on(table.surveyId, table.finished),
    uniqueIndex('Response_surveyId_singleUseId_key').on(table.surveyId, table.singleUseId),
    foreignKey({
      name: 'Response_surveyId_fkey',
      columns: [table.surveyId],
      foreignColumns: [survey.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Response_contactId_fkey',
      columns: [table.contactId],
      foreignColumns: [contact.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Response_displayId_fkey',
      columns: [table.displayId],
      foreignColumns: [display.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

/**
 * Represents a label that can be applied to survey responses.
 * Used for categorizing and organizing responses within a workspace.
 *
 * @property id - Unique identifier for the tag
 * @property name - Display name of the tag
 * @property responses - Survey responses tagged with this label
 * @property workspace - The workspace this tag belongs to
 */
export const tag = pgTable(
  'Tag',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    name: text('name').notNull(),
    workspaceId: text('workspaceId').notNull(),
  },
  (table) => [
    index('Tag_workspaceId_idx').on(table.workspaceId),
    uniqueIndex('Tag_workspaceId_name_key').on(table.workspaceId, table.name),
    foreignKey({
      name: 'Tag_workspaceId_fkey',
      columns: [table.workspaceId],
      foreignColumns: [workspace.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * Junction table linking tags to responses.
 * Enables many-to-many relationship between tags and responses.
 *
 * @property response - The tagged response
 * @property tag - The tag applied to the response
 */
export const tagsOnResponses = pgTable(
  'TagsOnResponses',
  { responseId: text('responseId').notNull(), tagId: text('tagId').notNull() },
  (table) => [
    primaryKey({ name: 'TagsOnResponses_pkey', columns: [table.responseId, table.tagId] }),
    foreignKey({
      name: 'TagsOnResponses_responseId_fkey',
      columns: [table.responseId],
      foreignColumns: [response.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'TagsOnResponses_tagId_fkey',
      columns: [table.tagId],
      foreignColumns: [tag.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const display = pgTable(
  'Display',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    surveyId: text('surveyId').notNull(),
    contactId: text('contactId'),
    creativeId: text('creativeId'),
  },
  (table) => [
    index('Display_surveyId_idx').on(table.surveyId),
    index('Display_contactId_created_at_idx').on(table.contactId, table.createdAt),
    foreignKey({
      name: 'Display_surveyId_fkey',
      columns: [table.surveyId],
      foreignColumns: [survey.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Display_contactId_fkey',
      columns: [table.contactId],
      foreignColumns: [contact.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const surveyTrigger = pgTable(
  'SurveyTrigger',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    surveyId: text('surveyId').notNull(),
    actionClassId: text('actionClassId').notNull(),
  },
  (table) => [
    uniqueIndex('SurveyTrigger_surveyId_actionClassId_key').on(table.surveyId, table.actionClassId),
    foreignKey({
      name: 'SurveyTrigger_surveyId_fkey',
      columns: [table.surveyId],
      foreignColumns: [survey.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'SurveyTrigger_actionClassId_fkey',
      columns: [table.actionClassId],
      foreignColumns: [actionClass.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const surveyAttributeFilter = pgTable(
  'SurveyAttributeFilter',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    attributeKeyId: text('attributeKeyId').notNull(),
    surveyId: text('surveyId').notNull(),
    condition: surveyAttributeFilterConditionEnum('condition').notNull(),
    value: text('value').notNull(),
  },
  (table) => [
    index('SurveyAttributeFilter_attributeKeyId_idx').on(table.attributeKeyId),
    uniqueIndex('SurveyAttributeFilter_surveyId_attributeKeyId_key').on(
      table.surveyId,
      table.attributeKeyId,
    ),
    foreignKey({
      name: 'SurveyAttributeFilter_attributeKeyId_fkey',
      columns: [table.attributeKeyId],
      foreignColumns: [contactAttributeKey.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'SurveyAttributeFilter_surveyId_fkey',
      columns: [table.surveyId],
      foreignColumns: [survey.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const survey = pgTable(
  'Survey',
  {
    /** [SurveyHiddenFields] */
    hiddenFields: jsonb('hiddenFields').notNull().default({ enabled: false }),
    /** [SurveyVariables] */
    variables: jsonb('variables').notNull().default([]),
    ownershipInitialized: boolean('ownershipInitialized').notNull().default(false),
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    name: text('name').notNull(),
    redirectUrl: text('redirectUrl'),
    type: surveyTypeEnum('type').notNull().default('app'),
    workspaceId: text('workspaceId').notNull(),
    createdBy: text('createdBy'),
    /**
     * ENG-3282: authorization facts, projected into SpiceDB by the survey outbox trigger. `ownerId` is
     * the durable ownership source; `createdBy` stays pure attribution and is never rewritten.
     */
    visibility: surveyVisibilityEnum('visibility').notNull().default('workspace'),
    ownerId: text('ownerId'),
    /**
     * Incremented by every accepted visibility change; fences projection work. An INSERT starts at 0
     * with `visibilityProjectedVersion` -1 (the `survey_visibility_pending` trigger): the initial
     * projection, decided from PostgreSQL facts until the projector acknowledges it.
     */
    visibilityVersion: integer('visibilityVersion').notNull().default(0),
    /** The `visibilityVersion` SpiceDB last acknowledged; -1 while none ever was. */
    visibilityProjectedVersion: integer('visibilityProjectedVersion').notNull().default(0),
    /**
     * "visibilityVersion" <> "visibilityProjectedVersion", kept by the `survey_visibility_pending`
     * trigger (a BEFORE trigger overrides whatever an INSERT or UPDATE supplies). Read-only: never put it
     * in create/update data.
     */
    visibilityPending: boolean('visibilityPending').notNull().default(false),
    visibilityChangedAt: timestamp('visibilityChangedAt', { precision: 3 }),
    visibilityChangedById: text('visibilityChangedById'),
    status: surveyStatusEnum('status').notNull().default('draft'),
    /** [SurveyWelcomeCard] */
    welcomeCard: jsonb('welcomeCard').notNull().default({ enabled: false }),
    /** [SurveyQuestions] */
    questions: jsonb('questions').notNull().default([]),
    /** [SurveyBlocks] */
    blocks: jsonb('blocks').array().default([]),
    /** [SurveyEnding] */
    endings: jsonb('endings').array().default([]),
    displayOption: displayOptionsEnum('displayOption').notNull().default('displayOnce'),
    recontactDays: integer('recontactDays'),
    displayLimit: integer('displayLimit'),
    /** [SurveyInlineTriggers] */
    inlineTriggers: jsonb('inlineTriggers'),
    autoClose: integer('autoClose'),
    keepEndingCardOpen: boolean('keepEndingCardOpen').notNull().default(false),
    autoComplete: integer('autoComplete'),
    delay: integer('delay').notNull().default(0),
    publishOn: timestamp('publishOn', { precision: 3 }),
    closeOn: timestamp('closeOn', { precision: 3 }),
    archivedAt: timestamp('archivedAt', { precision: 3 }),
    /** [SurveyClosedMessage] */
    surveyClosedMessage: jsonb('surveyClosedMessage'),
    segmentId: text('segmentId'),
    /** [SurveyWorkspaceOverwrites] */
    workspaceOverwrites: jsonb('workspaceOverwrites'),
    /** [SurveyStyling] */
    styling: jsonb('styling'),
    /** [SurveySingleUse] */
    singleUse: jsonb('singleUse').default({ enabled: false, isEncrypted: true }),
    isVerifyEmailEnabled: boolean('isVerifyEmailEnabled').notNull().default(false),
    isBackButtonHidden: boolean('isBackButtonHidden').notNull().default(false),
    isAutoProgressingEnabled: boolean('isAutoProgressingEnabled').notNull().default(false),
    isCaptureIpEnabled: boolean('isCaptureIpEnabled').notNull().default(false),
    isAnonymizeResponsesEnabled: boolean('is_anonymize_responses_enabled').notNull().default(false),
    pin: text('pin'),
    displayPercentage: numeric('displayPercentage', { precision: 65, scale: 30 }),
    showLanguageSwitch: boolean('showLanguageSwitch'),
    /**
     * Open the survey in the respondent's browser language when it matches an enabled survey
     * language. Null on surveys that predate the setting, which reads as off.
     */
    autoSelectLanguage: boolean('autoSelectLanguage'),
    /** [SurveyRecaptcha] */
    recaptcha: jsonb('recaptcha').default({ enabled: false, threshold: 0.1 }),
    /** [SurveyLinkMetadata] */
    metadata: jsonb('metadata').notNull().default({}),
    slug: text('slug'),
    customHeadScripts: text('customHeadScripts'),
    customHeadScriptsMode: surveyScriptModeEnum('customHeadScriptsMode').default('add'),
    /**
     * [CustomCssStored] Survey custom CSS (ENG-2949): source plus trusted compiled output, written only
     * by the shared custom CSS save service. null = no survey CSS. Kept outside `styling` (M2.07).
     */
    customCss: jsonb('customCss'),
    /**
     * [SurveyShowConditions] App surveys: conditions checked in the browser when a trigger fires
     * (ENG-3873). null = no conditions, which is what every existing survey keeps.
     */
    showConditions: jsonb('showConditions'),
    /**
     * [SurveyStartLogic] Rules before the first question (ENG-6932): block-rule format, run by the
     * survey on load. null = none, which is what every existing survey keeps.
     */
    startLogic: jsonb('startLogic'),
  },
  (table) => [
    uniqueIndex('Survey_slug_key').on(table.slug),
    index('Survey_workspaceId_updated_at_idx').on(table.workspaceId, table.updatedAt),
    index('Survey_workspaceId_ownerId_idx').on(table.workspaceId, table.ownerId),
    index('Survey_segmentId_idx').on(table.segmentId),
    index('Survey_status_publishOn_idx').on(table.status, table.publishOn),
    index('Survey_status_closeOn_idx').on(table.status, table.closeOn),
    index('Survey_archivedAt_idx').on(table.archivedAt),
    unique('Survey_id_workspaceId_key').on(table.id, table.workspaceId),
    foreignKey({
      name: 'Survey_workspaceId_fkey',
      columns: [table.workspaceId],
      foreignColumns: [workspace.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Survey_createdBy_fkey',
      columns: [table.createdBy],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({ name: 'Survey_ownerId_fkey', columns: [table.ownerId], foreignColumns: [user.id] })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Survey_segmentId_fkey',
      columns: [table.segmentId],
      foreignColumns: [segment.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

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
export const surveyQuota = pgTable(
  'SurveyQuota',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    surveyId: text('surveyId').notNull(),
    name: text('name').notNull(),
    limit: integer('limit').notNull(),
    /** [SurveyQuotaLogic] */
    logic: jsonb('logic').notNull().default({}),
    action: surveyQuotaActionEnum('action').notNull(),
    endingCardId: text('endingCardId'),
    countPartialSubmissions: boolean('countPartialSubmissions').notNull().default(false),
  },
  (table) => [
    uniqueIndex('SurveyQuota_surveyId_name_key').on(table.surveyId, table.name),
    foreignKey({
      name: 'SurveyQuota_surveyId_fkey',
      columns: [table.surveyId],
      foreignColumns: [survey.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * Junction table linking responses to quotas.
 * Tracks which responses counted towards specific quotas and their status.
 *
 * @property response - The response that counted towards the quota
 * @property quota - The quota this response counted towards
 * @property status - Whether the response was screened in or out
 */
export const responseQuotaLink = pgTable(
  'ResponseQuotaLink',
  {
    responseId: text('responseId').notNull(),
    quotaId: text('quotaId').notNull(),
    status: responseQuotaLinkStatusEnum('status').notNull(),
  },
  (table) => [
    primaryKey({ name: 'ResponseQuotaLink_pkey', columns: [table.responseId, table.quotaId] }),
    index('ResponseQuotaLink_quotaId_status_idx').on(table.quotaId, table.status),
    foreignKey({
      name: 'ResponseQuotaLink_responseId_fkey',
      columns: [table.responseId],
      foreignColumns: [response.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'ResponseQuotaLink_quotaId_fkey',
      columns: [table.quotaId],
      foreignColumns: [surveyQuota.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const surveyFollowUp = pgTable(
  'SurveyFollowUp',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    surveyId: text('surveyId').notNull(),
    name: text('name').notNull(),
    /** [SurveyFollowUpTrigger] */
    trigger: jsonb('trigger').notNull(),
    /** [SurveyFollowUpAction] */
    action: jsonb('action').notNull(),
  },
  (table) => [
    foreignKey({
      name: 'SurveyFollowUp_surveyId_fkey',
      columns: [table.surveyId],
      foreignColumns: [survey.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const actionClass = pgTable(
  'ActionClass',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    name: text('name').notNull(),
    description: text('description'),
    type: actionTypeEnum('type').notNull(),
    key: text('key'),
    /** [ActionClassNoCodeConfig] */
    noCodeConfig: jsonb('noCodeConfig'),
    workspaceId: text('workspaceId').notNull(),
  },
  (table) => [
    index('ActionClass_workspaceId_created_at_idx').on(table.workspaceId, table.createdAt),
    uniqueIndex('ActionClass_key_workspaceId_key').on(table.key, table.workspaceId),
    uniqueIndex('ActionClass_name_workspaceId_key').on(table.name, table.workspaceId),
    foreignKey({
      name: 'ActionClass_workspaceId_fkey',
      columns: [table.workspaceId],
      foreignColumns: [workspace.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * Represents third-party service integrations.
 * Enables data flow between Formbricks and external services.
 *
 * @property id - Unique identifier for the integration
 * @property type - The service being integrated (Google Sheets, Notion, etc.)
 * @property workspace - The workspace this integration belongs to
 * @property config - Service-specific configuration details
 */
export const integration = pgTable(
  'Integration',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    type: integrationTypeEnum('type').notNull(),
    /** [IntegrationConfig] */
    config: jsonb('config').notNull(),
    workspaceId: text('workspaceId').notNull(),
  },
  (table) => [
    index('Integration_workspaceId_idx').on(table.workspaceId),
    uniqueIndex('Integration_type_workspaceId_key').on(table.type, table.workspaceId),
    foreignKey({
      name: 'Integration_workspaceId_fkey',
      columns: [table.workspaceId],
      foreignColumns: [workspace.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const dataMigration = pgTable(
  'DataMigration',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    startedAt: timestamp('started_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    finishedAt: timestamp('finished_at', { precision: 3 }),
    name: text('name').notNull(),
    status: dataMigrationStatusEnum('status').notNull(),
  },
  (table) => [uniqueIndex('DataMigration_name_key').on(table.name)],
)

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
export const workspace = pgTable(
  'Workspace',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    name: text('name').notNull(),
    legacyEnvironmentId: text('legacyEnvironmentId'),
    organizationId: text('organizationId').notNull(),
    /** [Styling] */
    styling: jsonb('styling').notNull().default({ allowStyleOverwrite: true }),
    /** [WorkspaceConfig] */
    config: jsonb('config').notNull().default({}),
    recontactDays: integer('recontactDays').notNull().default(7),
    linkSurveyBranding: boolean('linkSurveyBranding').notNull().default(true),
    inAppSurveyBranding: boolean('inAppSurveyBranding').notNull().default(true),
    placement: widgetPlacementEnum('placement').notNull().default('bottomRight'),
    clickOutsideClose: boolean('clickOutsideClose').notNull().default(true),
    overlay: surveyOverlayEnum('overlay').notNull().default('none'),
    /** Custom overlay colour (hex, no alpha) and opacity (10-100%). null = the preset for `overlay`. */
    overlayColor: text('overlayColor'),
    overlayOpacity: integer('overlayOpacity'),
    /** [Logo] */
    logo: jsonb('logo'),
    appSetupCompleted: boolean('appSetupCompleted').notNull().default(false),
    customHeadScripts: text('customHeadScripts'),
    /** [CustomCssStored] Workspace custom CSS (ENG-2949), written only by the shared save service. */
    customCss: jsonb('customCss'),
    /**
     * [CustomCssStored] The one recoverable previous workspace CSS revision, source included. Restored by
     * submitting its source again through the current processor and authorization checks.
     */
    customCssPrevious: jsonb('customCssPrevious'),
  },
  (table) => [
    uniqueIndex('Workspace_legacyEnvironmentId_key').on(table.legacyEnvironmentId),
    uniqueIndex('Workspace_organizationId_name_key').on(table.organizationId, table.name),
    foreignKey({
      name: 'Workspace_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const organization = pgTable('Organization', {
  id: text('id')
    .primaryKey()
    .$defaultFn(() => cuid()),
  createdAt: timestamp('created_at', { precision: 3 })
    .notNull()
    .default(sql`CURRENT_TIMESTAMP`)
    .$defaultFn(utcNow),
  updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
  name: text('name').notNull(),
  /** [OrganizationWhitelabel] */
  whitelabel: jsonb('whitelabel').notNull().default({}),
  isAISmartToolsEnabled: boolean('isAISmartToolsEnabled').notNull().default(false),
  displayTimeZone: text('displayTimeZone'),
})

/**
 * Stores billing and Stripe synchronization data for an organization.
 *
 * @property organizationId - Reference to the organization
 * @property stripeCustomerId - Stripe customer ID for cloud billing
 * @property limits - Plan limits used by feature and quota checks
 * @property usageCycleAnchor - Stable anchor used to derive the current monthly usage cycle
 * @property stripe - Stripe synchronization snapshot data
 */
export const organizationBilling = pgTable(
  'OrganizationBilling',
  {
    organizationId: text('organization_id').primaryKey(),
    stripeCustomerId: text('stripe_customer_id'),
    /** [OrganizationBillingPlanLimits] */
    limits: jsonb('limits').notNull(),
    usageCycleAnchor: timestamp('usage_cycle_anchor', { precision: 3 }),
    /** [OrganizationStripeBilling] */
    stripe: jsonb('stripe'),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('OrganizationBilling_stripe_customer_id_key').on(table.stripeCustomerId),
    foreignKey({
      name: 'OrganizationBilling_organization_id_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const membership = pgTable(
  'Membership',
  {
    organizationId: text('organizationId').notNull(),
    userId: text('userId').notNull(),
    accepted: boolean('accepted').notNull().default(false),
    role: organizationRoleEnum('role').notNull().default('member'),
  },
  (table) => [
    primaryKey({ name: 'Membership_pkey', columns: [table.userId, table.organizationId] }),
    index('Membership_organizationId_idx').on(table.organizationId),
    foreignKey({
      name: 'Membership_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Membership_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const invite = pgTable(
  'Invite',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => crypto.randomUUID()),
    email: text('email').notNull(),
    name: text('name'),
    organizationId: text('organizationId').notNull(),
    creatorId: text('creatorId').notNull(),
    acceptorId: text('acceptorId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    expiresAt: timestamp('expiresAt', { precision: 3 }).notNull(),
    role: organizationRoleEnum('role').notNull().default('member'),
    teamIds: text('teamIds').array().default([]),
  },
  (table) => [
    index('Invite_email_organizationId_idx').on(table.email, table.organizationId),
    index('Invite_organizationId_idx').on(table.organizationId),
    foreignKey({
      name: 'Invite_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Invite_creatorId_fkey',
      columns: [table.creatorId],
      foreignColumns: [user.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Invite_acceptorId_fkey',
      columns: [table.acceptorId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const apiKey = pgTable(
  'ApiKey',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    createdBy: text('createdBy'),
    lastUsedAt: timestamp('lastUsedAt', { precision: 3 }),
    label: text('label').notNull(),
    hashedKey: text('hashedKey').notNull(),
    lookupHash: text('lookupHash'),
    organizationId: text('organizationId').notNull(),
    /** [OrganizationAccess] */
    organizationAccess: jsonb('organizationAccess').notNull().default({}),
  },
  (table) => [
    uniqueIndex('ApiKey_lookupHash_key').on(table.lookupHash),
    index('ApiKey_organizationId_idx').on(table.organizationId),
    foreignKey({
      name: 'ApiKey_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * Links API keys to workspaces with specific permissions.
 * Enables granular access control for API keys across workspaces.
 *
 * @property id - Unique identifier for the workspace access entry
 * @property apiKey - The associated API key
 * @property workspace - The workspace being accessed
 * @property permission - Level of access granted
 */
export const apiKeyWorkspace = pgTable(
  'ApiKeyWorkspace',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    apiKeyId: text('apiKeyId').notNull(),
    workspaceId: text('workspaceId').notNull(),
    permission: apiKeyPermissionEnum('permission').notNull(),
  },
  (table) => [
    index('ApiKeyWorkspace_workspaceId_idx').on(table.workspaceId),
    uniqueIndex('ApiKeyWorkspace_apiKeyId_workspaceId_key').on(table.apiKeyId, table.workspaceId),
    foreignKey({
      name: 'ApiKeyWorkspace_apiKeyId_fkey',
      columns: [table.apiKeyId],
      foreignColumns: [apiKey.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'ApiKeyWorkspace_workspaceId_fkey',
      columns: [table.workspaceId],
      foreignColumns: [workspace.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const account = pgTable(
  'Account',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    userId: text('userId').notNull(),
    type: text('type'),
    provider: text('provider').notNull(),
    providerAccountId: text('providerAccountId').notNull(),
    access_token: text('access_token'),
    refresh_token: text('refresh_token'),
    expires_at: integer('expires_at'),
    ext_expires_in: integer('ext_expires_in'),
    token_type: text('token_type'),
    scope: text('scope'),
    id_token: text('id_token'),
    session_state: text('session_state'),
    password: text('password'),
    accessTokenExpiresAt: timestamp('accessTokenExpiresAt', { precision: 3 }),
    refreshTokenExpiresAt: timestamp('refreshTokenExpiresAt', { precision: 3 }),
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
    issuer: text('issuer'),
  },
  (table) => [
    index('Account_userId_idx').on(table.userId),
    uniqueIndex('Account_provider_providerAccountId_key').on(
      table.provider,
      table.providerAccountId,
    ),
    uniqueIndex('Account_issuer_providerAccountId_key').on(table.issuer, table.providerAccountId),
    foreignKey({ name: 'Account_userId_fkey', columns: [table.userId], foreignColumns: [user.id] })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * Stores active authentication sessions for revocable server-side login state.
 *
 * @property sessionToken - Opaque token stored in the browser cookie
 * @property user - The Formbricks user who owns this session
 * @property expires - Hard expiry for the session
 */
export const session = pgTable(
  'Session',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    sessionToken: text('sessionToken').notNull(),
    userId: text('userId').notNull(),
    expires: timestamp('expires', { precision: 3 }).notNull(),
    ipAddress: text('ipAddress'),
    userAgent: text('userAgent'),
  },
  (table) => [
    uniqueIndex('Session_sessionToken_key').on(table.sessionToken),
    index('Session_userId_idx').on(table.userId),
    foreignKey({ name: 'Session_userId_fkey', columns: [table.userId], foreignColumns: [user.id] })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/** Stores one-time verification tokens used by Auth.js adapter flows. */
export const verificationToken = pgTable(
  'VerificationToken',
  {
    identifier: text('identifier').notNull(),
    token: text('token').notNull(),
    expires: timestamp('expires', { precision: 3 }).notNull(),
  },
  (table) => [
    uniqueIndex('VerificationToken_identifier_token_key').on(table.identifier, table.token),
  ],
)

/**
 * Better Auth (ENG-1054) two-factor secrets. Supersedes the legacy User.twoFactorSecret /
 * User.backupCodes columns (kept until cutover). `secret` and `backupCodes` are encrypted
 * at rest by Better Auth using BETTER_AUTH_SECRET.
 */
export const twoFactor = pgTable(
  'TwoFactor',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    secret: text('secret').notNull(),
    backupCodes: text('backupCodes').notNull(),
    userId: text('userId').notNull(),
    verified: boolean('verified').notNull().default(true),
    /** Better Auth 1.6.23+ two-factor brute-force lockout: consecutive failed verification attempts. */
    failedVerificationCount: integer('failedVerificationCount').notNull().default(0),
    /** Better Auth 1.6.23+ two-factor brute-force lockout: set once failedVerificationCount trips the limit. */
    lockedUntil: timestamp('lockedUntil', { precision: 3 }),
  },
  (table) => [
    uniqueIndex('TwoFactor_userId_key').on(table.userId),
    foreignKey({
      name: 'TwoFactor_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/** Better Auth JWT signing keys for OAuth/OIDC access and ID tokens. */
export const jwks = pgTable('jwks', {
  id: text('id')
    .primaryKey()
    .$defaultFn(() => cuid()),
  publicKey: text('publicKey').notNull(),
  privateKey: text('privateKey').notNull(),
  createdAt: timestamp('createdAt', { precision: 3 }).notNull(),
  expiresAt: timestamp('expiresAt', { precision: 3 }),
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
  alg: text('alg'),
  /**
   * Better Auth 1.7 (ENG-2343): the curve for EC/OKP keys. Written whenever it can be derived from
   * the generated key, which it can for the default EdDSA/Ed25519 config we run, so it is on the
   * same always-written path as `alg` above.
   */
  crv: text('crv'),
})

/** Better Auth OAuth 2.1 client registrations for MCP and future OAuth clients. */
export const oauthClient = pgTable(
  'oauthClient',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    clientId: text('clientId').notNull(),
    clientSecret: text('clientSecret'),
    disabled: boolean('disabled').default(false),
    skipConsent: boolean('skipConsent'),
    enableEndSession: boolean('enableEndSession'),
    subjectType: text('subjectType'),
    scopes: text('scopes').array().default([]),
    userId: text('userId'),
    createdAt: timestamp('createdAt', { precision: 3 }),
    updatedAt: timestamp('updatedAt', { precision: 3 }),
    name: text('name'),
    uri: text('uri'),
    icon: text('icon'),
    contacts: text('contacts').array().default([]),
    tos: text('tos'),
    policy: text('policy'),
    softwareId: text('softwareId'),
    softwareVersion: text('softwareVersion'),
    softwareStatement: text('softwareStatement'),
    redirectUris: text('redirectUris').array(),
    postLogoutRedirectUris: text('postLogoutRedirectUris').array().default([]),
    tokenEndpointAuthMethod: text('tokenEndpointAuthMethod'),
    grantTypes: text('grantTypes').array().default([]),
    responseTypes: text('responseTypes').array().default([]),
    /**
     * Unwritten by Better Auth 1.7 (it dropped both from its schema), kept deliberately: they are
     * nullable, so keeping them makes the 1.7 migration purely additive and a rollback to 1.6 needs no
     * down-migration. Drop them in a separate cleanup after a bake period.
     */
    public: boolean('public'),
    type: text('type'),
    requirePKCE: boolean('requirePKCE'),
    referenceId: text('referenceId'),
    metadata: jsonb('metadata'),
    /** Better Auth 1.7 (ENG-2343). */
    clientDiscoveryId: text('clientDiscoveryId'),
    clientCredentialsScopes: text('clientCredentialsScopes').array().default([]),
    backchannelLogoutUri: text('backchannelLogoutUri'),
    backchannelLogoutSessionRequired: boolean('backchannelLogoutSessionRequired'),
    applicationType: text('applicationType'),
    jwks: text('jwks'),
    jwksUri: text('jwksUri'),
    dpopBoundAccessTokens: boolean('dpopBoundAccessTokens').default(false),
  },
  (table) => [
    unique('oauthClient_clientId_key').on(table.clientId),
    index('oauthClient_userId_idx').on(table.userId),
    foreignKey({
      name: 'oauthClient_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * Better Auth opaque access token records. The token column stores Better Auth's hashed lookup value.
 * JWT access tokens are verified through JWKS and bounded by the configured short lifetime.
 */
export const oauthAccessToken = pgTable(
  'oauthAccessToken',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    token: text('token').notNull(),
    clientId: text('clientId').notNull(),
    sessionId: text('sessionId'),
    userId: text('userId'),
    referenceId: text('referenceId'),
    refreshId: text('refreshId'),
    expiresAt: timestamp('expiresAt', { precision: 3 }).notNull(),
    createdAt: timestamp('createdAt', { precision: 3 }).notNull(),
    scopes: text('scopes').array(),
    /**
     * Better Auth 1.7 (ENG-2343). `resources` is the RFC 8707 resource set the grant approved — the
     * binding this whole upgrade exists to get (GHSA-p2fr-6hmx-4528).
     */
    authorizationCodeId: text('authorizationCodeId'),
    resources: text('resources').array().default([]),
    requestedUserInfoClaims: text('requestedUserInfoClaims').array().default([]),
    revoked: timestamp('revoked', { precision: 3 }),
    confirmation: jsonb('confirmation'),
  },
  (table) => [
    uniqueIndex('oauthAccessToken_token_key').on(table.token),
    index('oauthAccessToken_clientId_idx').on(table.clientId),
    index('oauthAccessToken_sessionId_idx').on(table.sessionId),
    index('oauthAccessToken_userId_idx').on(table.userId),
    index('oauthAccessToken_refreshId_idx').on(table.refreshId),
    index('oauthAccessToken_authorizationCodeId_idx').on(table.authorizationCodeId),
    foreignKey({
      name: 'oauthAccessToken_clientId_fkey',
      columns: [table.clientId],
      foreignColumns: [oauthClient.clientId],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'oauthAccessToken_sessionId_fkey',
      columns: [table.sessionId],
      foreignColumns: [session.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'oauthAccessToken_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'oauthAccessToken_refreshId_fkey',
      columns: [table.refreshId],
      foreignColumns: [oauthRefreshToken.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/** Better Auth OAuth refresh tokens used by delegated MCP clients. The token column stores a hashed lookup value. */
export const oauthRefreshToken = pgTable(
  'oauthRefreshToken',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    token: text('token').notNull(),
    clientId: text('clientId').notNull(),
    sessionId: text('sessionId'),
    userId: text('userId').notNull(),
    referenceId: text('referenceId'),
    expiresAt: timestamp('expiresAt', { precision: 3 }).notNull(),
    createdAt: timestamp('createdAt', { precision: 3 }).notNull(),
    revoked: timestamp('revoked', { precision: 3 }),
    authTime: timestamp('authTime', { precision: 3 }),
    scopes: text('scopes').array(),
    /**
     * Better Auth 1.7 (ENG-2343). `resources` carries the full original grant so a refresh can narrow
     * the resource set but never widen it.
     */
    authorizationCodeId: text('authorizationCodeId'),
    resources: text('resources').array().default([]),
    requestedUserInfoClaims: text('requestedUserInfoClaims').array().default([]),
    rotatedAt: timestamp('rotatedAt', { precision: 3 }),
    rotationReplayResponse: text('rotationReplayResponse'),
    rotationReplayExpiresAt: timestamp('rotationReplayExpiresAt', { precision: 3 }),
    confirmation: jsonb('confirmation'),
  },
  (table) => [
    uniqueIndex('oauthRefreshToken_token_key').on(table.token),
    index('oauthRefreshToken_clientId_idx').on(table.clientId),
    index('oauthRefreshToken_sessionId_idx').on(table.sessionId),
    index('oauthRefreshToken_userId_idx').on(table.userId),
    index('oauthRefreshToken_authorizationCodeId_idx').on(table.authorizationCodeId),
    foreignKey({
      name: 'oauthRefreshToken_clientId_fkey',
      columns: [table.clientId],
      foreignColumns: [oauthClient.clientId],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'oauthRefreshToken_sessionId_fkey',
      columns: [table.sessionId],
      foreignColumns: [session.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'oauthRefreshToken_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/** Better Auth OAuth user consent grants for delegated clients. */
export const oauthConsent = pgTable(
  'oauthConsent',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    clientId: text('clientId').notNull(),
    userId: text('userId'),
    referenceId: text('referenceId'),
    scopes: text('scopes').array(),
    createdAt: timestamp('createdAt', { precision: 3 }).notNull(),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull(),
    /**
     * Better Auth 1.7 (ENG-2343). /authorize compares the requested resources against these; without
     * them every pre-existing connection is forced back through the consent screen.
     */
    resources: text('resources').array().default([]),
    requestedUserInfoClaims: text('requestedUserInfoClaims').array().default([]),
  },
  (table) => [
    index('oauthConsent_clientId_idx').on(table.clientId),
    index('oauthConsent_userId_idx').on(table.userId),
    foreignKey({
      name: 'oauthConsent_clientId_fkey',
      columns: [table.clientId],
      foreignColumns: [oauthClient.clientId],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'oauthConsent_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * Better Auth 1.7 (ENG-2343): a protected resource the authorization server issues tokens for, with
 * its own token lifetimes, scope allow-list and claims. Replaces the flat `validAudiences` option —
 * a token is now bound to the resource it was approved for instead of to a static allow-list.
 * `identifier` is the RFC 8707 `resource` value and is the natural key other rows point at.
 * What one user allows one MCP client to do (GLOSSARY.md "Grant", ADR-0041): its organization & workspace
 * access and an access level per permission. One per OAuth connection, keyed by user and client rather
 * than by `oauthConsent` row, because the provider can write duplicate consent rows for one connection
 * and settles them later. Every MCP request reads it; without one the request is refused. The client's
 * OAuth scopes stay on the consent and cap it.
 */
export const mcpGrant = pgTable(
  'McpGrant',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    userId: text('userId').notNull(),
    clientId: text('clientId').notNull(),
    /**
     * Which organizations and workspaces the grant reaches (ENG-7153): All access, only the organizations
     * in `organizations` (and every workspace in them, later ones included), or only the workspaces in
     * `workspaces`. Enforced in SpiceDB through `can()` (ADR-0042); a value no row backs reaches nothing.
     */
    access: mcpGrantAccessEnum('access').notNull().default('all'),
    /**
     * The access level per permission of `apps/web/modules/mcp/permissions.ts`: `{ "surveys": "write" }`.
     * A permission that is absent has No access.
     */
    permissions: jsonb('permissions').notNull().default({}),
  },
  (table) => [
    index('McpGrant_clientId_idx').on(table.clientId),
    uniqueIndex('McpGrant_userId_clientId_key').on(table.userId, table.clientId),
    foreignKey({ name: 'McpGrant_userId_fkey', columns: [table.userId], foreignColumns: [user.id] })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'McpGrant_clientId_fkey',
      columns: [table.clientId],
      foreignColumns: [oauthClient.clientId],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/** An organization a grant with `organizations` access reaches, workspaces created there later included. */
export const mcpGrantOrganization = pgTable(
  'McpGrantOrganization',
  { grantId: text('grantId').notNull(), organizationId: text('organizationId').notNull() },
  (table) => [
    primaryKey({
      name: 'McpGrantOrganization_pkey',
      columns: [table.grantId, table.organizationId],
    }),
    index('McpGrantOrganization_organizationId_idx').on(table.organizationId),
    foreignKey({
      name: 'McpGrantOrganization_grantId_fkey',
      columns: [table.grantId],
      foreignColumns: [mcpGrant.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'McpGrantOrganization_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/** A workspace a grant with `workspaces` access reaches. */
export const mcpGrantWorkspace = pgTable(
  'McpGrantWorkspace',
  { grantId: text('grantId').notNull(), workspaceId: text('workspaceId').notNull() },
  (table) => [
    primaryKey({ name: 'McpGrantWorkspace_pkey', columns: [table.grantId, table.workspaceId] }),
    index('McpGrantWorkspace_workspaceId_idx').on(table.workspaceId),
    foreignKey({
      name: 'McpGrantWorkspace_grantId_fkey',
      columns: [table.grantId],
      foreignColumns: [mcpGrant.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'McpGrantWorkspace_workspaceId_fkey',
      columns: [table.workspaceId],
      foreignColumns: [workspace.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const oauthResource = pgTable(
  'oauthResource',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    identifier: text('identifier').notNull(),
    name: text('name').notNull(),
    accessTokenTtl: integer('accessTokenTtl'),
    refreshTokenTtl: integer('refreshTokenTtl'),
    signingAlgorithm: text('signingAlgorithm'),
    signingKeyId: text('signingKeyId'),
    allowedScopes: text('allowedScopes').array().default([]),
    customClaims: jsonb('customClaims'),
    dpopBoundAccessTokensRequired: boolean('dpopBoundAccessTokensRequired').default(false),
    disabled: boolean('disabled').default(false),
    createdAt: timestamp('createdAt', { precision: 3 }),
    updatedAt: timestamp('updatedAt', { precision: 3 }),
    policyVersion: integer('policyVersion').default(1),
    metadata: jsonb('metadata'),
  },
  (table) => [unique('oauthResource_identifier_key').on(table.identifier)],
)

/**
 * Better Auth 1.7 (ENG-2343): links a client to the resources it may request tokens for. With
 * `enforcePerClientResources` (the default) a missing row means `invalid_target` at the token
 * endpoint, which is why the migration backfills one row per existing client.
 * Both sides reference natural keys, matching the other oauth child tables.
 */
export const oauthClientResource = pgTable(
  'oauthClientResource',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    clientId: text('clientId').notNull(),
    resourceId: text('resourceId').notNull(),
    metadata: jsonb('metadata'),
    createdAt: timestamp('createdAt', { precision: 3 }),
  },
  (table) => [
    index('oauthClientResource_resourceId_idx').on(table.resourceId),
    uniqueIndex('oauthClientResource_clientId_resourceId_key').on(table.clientId, table.resourceId),
    foreignKey({
      name: 'oauthClientResource_clientId_fkey',
      columns: [table.clientId],
      foreignColumns: [oauthClient.clientId],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'oauthClientResource_resourceId_fkey',
      columns: [table.resourceId],
      foreignColumns: [oauthResource.identifier],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * Better Auth 1.7 (ENG-2343): replay store for `private_key_jwt` client assertions (the row id is the
 * assertion jti). We use no private_key_jwt client authentication today, but the plugin declares the
 * table, so omitting it risks an adapter error on a code path we do not control.
 */
export const oauthClientAssertion = pgTable(
  'oauthClientAssertion',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    expiresAt: timestamp('expiresAt', { precision: 3 }).notNull(),
  },
  (table) => [index('oauthClientAssertion_expiresAt_idx').on(table.expiresAt)],
)

/**
 * Stores the active password reset token for a user.
 * Tokens are opaque, hashed at rest, revocable, and single-use.
 */
export const passwordResetToken = pgTable(
  'PasswordResetToken',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    tokenHash: text('token_hash').notNull(),
    expiresAt: timestamp('expires_at', { precision: 3 }).notNull(),
    userId: text('userId').notNull(),
  },
  (table) => [
    uniqueIndex('PasswordResetToken_token_hash_key').on(table.tokenHash),
    uniqueIndex('PasswordResetToken_userId_key').on(table.userId),
    index('PasswordResetToken_expires_at_idx').on(table.expiresAt),
    foreignKey({
      name: 'PasswordResetToken_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const user = pgTable(
  'User',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    name: text('name').notNull(),
    email: text('email').notNull(),
    emailVerified: boolean('email_verified').notNull().default(false),
    twoFactorSecret: text('twoFactorSecret'),
    twoFactorEnabled: boolean('twoFactorEnabled').notNull().default(false),
    backupCodes: text('backupCodes'),
    password: text('password'),
    identityProvider: identityProviderEnum('identityProvider').notNull().default('email'),
    identityProviderAccountId: text('identityProviderAccountId'),
    groupId: text('groupId'),
    /** [UserNotificationSettings] */
    notificationSettings: jsonb('notificationSettings').notNull().default({}),
    /** [Locale] */
    locale: text('locale').notNull().default('en-US'),
    lastLoginAt: timestamp('lastLoginAt', { precision: 3 }),
    isActive: boolean('isActive').notNull().default(true),
    /**
     * When an admin last reactivated this user. Restarts the members retention clock (ENG-3713)
     * without touching `lastLoginAt`, which the member list, usage counts and APIs report as a real
     * sign-in.
     */
    reactivatedAt: timestamp('reactivatedAt', { precision: 3 }),
    /**
     * When the user last completed a sign-in (`lastActivePlugin`, after any second factor) or renewed a
     * session (Better Auth `databaseHooks.session.update`), on the database clock and never moved
     * backwards. The members retention clock reads it: sessions are
     * deleted on sign-out and expiry, so their own timestamps can't serve as a record of activity.
     */
    lastActiveAt: timestamp('lastActiveAt', { precision: 3 }),
    isBootstrapAdmin: bootstrapAdminMarkerEnum('isBootstrapAdmin'),
  },
  (table) => [
    uniqueIndex('User_email_key').on(table.email),
    uniqueIndex('User_isBootstrapAdmin_key').on(table.isBootstrapAdmin),
  ],
)

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
export const segment = pgTable(
  'Segment',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    title: text('title').notNull(),
    description: text('description'),
    isPrivate: boolean('isPrivate').notNull().default(true),
    /** [SegmentFilter] */
    filters: jsonb('filters').notNull().default([]),
    workspaceId: text('workspaceId').notNull(),
  },
  (table) => [
    index('Segment_workspaceId_idx').on(table.workspaceId),
    uniqueIndex('Segment_workspaceId_title_key').on(table.workspaceId, table.title),
    foreignKey({
      name: 'Segment_workspaceId_fkey',
      columns: [table.workspaceId],
      foreignColumns: [workspace.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * Represents a supported language in the system.
 * Used for multilingual survey support.
 *
 * @property id - Unique identifier for the language
 * @property code - Language code (e.g., 'en-US')
 * @property alias - Optional friendly name
 * @property workspace - The workspace this language is enabled for
 */
export const language = pgTable(
  'Language',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    code: text('code').notNull(),
    alias: text('alias'),
    workspaceId: text('workspaceId').notNull(),
  },
  (table) => [
    uniqueIndex('Language_workspaceId_code_key').on(table.workspaceId, table.code),
    foreignKey({
      name: 'Language_workspaceId_fkey',
      columns: [table.workspaceId],
      foreignColumns: [workspace.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * Links surveys to their supported languages.
 * Manages which languages are available for each survey.
 *
 * @property language - The supported language
 * @property survey - The associated survey
 * @property default - Whether this is the default language
 * @property enabled - Whether this language is currently active
 */
export const surveyLanguage = pgTable(
  'SurveyLanguage',
  {
    languageId: text('languageId').notNull(),
    surveyId: text('surveyId').notNull(),
    default: boolean('default').notNull().default(false),
    enabled: boolean('enabled').notNull().default(true),
  },
  (table) => [
    primaryKey({ name: 'SurveyLanguage_pkey', columns: [table.languageId, table.surveyId] }),
    index('SurveyLanguage_surveyId_idx').on(table.surveyId),
    foreignKey({
      name: 'SurveyLanguage_languageId_fkey',
      columns: [table.languageId],
      foreignColumns: [language.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'SurveyLanguage_surveyId_fkey',
      columns: [table.surveyId],
      foreignColumns: [survey.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const embeddedData = pgTable(
  'EmbeddedData',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    key: text('key'),
    name: text('name').notNull(),
    description: text('description'),
    source: embeddedDataSourceEnum('source').notNull(),
    dataType: embeddedDataTypeEnum('dataType').notNull().default('string'),
    /** [EmbeddedDataDefaultValue] */
    defaultValue: jsonb('defaultValue'),
    locked: boolean('locked').notNull().default(false),
    surveyId: text('surveyId'),
    workspaceId: text('workspaceId').notNull(),
  },
  (table) => [
    index('EmbeddedData_surveyId_idx').on(table.surveyId),
    uniqueIndex('EmbeddedData_workspaceId_key_key').on(table.workspaceId, table.key),
    unique('EmbeddedData_id_workspaceId_key').on(table.id, table.workspaceId),
    foreignKey({
      name: 'EmbeddedData_surveyId_workspaceId_fkey',
      columns: [table.surveyId, table.workspaceId],
      foreignColumns: [survey.id, survey.workspaceId],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'EmbeddedData_workspaceId_fkey',
      columns: [table.workspaceId],
      foreignColumns: [workspace.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const surveyEmbeddedData = pgTable(
  'SurveyEmbeddedData',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    workspaceId: text('workspaceId').notNull(),
    surveyId: text('surveyId').notNull(),
    embeddedDataId: text('embeddedDataId').notNull(),
    storageKey: text('storageKey').notNull(),
    order: integer('order').notNull(),
  },
  (table) => [
    index('SurveyEmbeddedData_embeddedDataId_idx').on(table.embeddedDataId),
    index('SurveyEmbeddedData_surveyId_order_idx').on(table.surveyId, table.order),
    uniqueIndex('SurveyEmbeddedData_surveyId_embeddedDataId_key').on(
      table.surveyId,
      table.embeddedDataId,
    ),
    uniqueIndex('SurveyEmbeddedData_surveyId_storageKey_key').on(table.surveyId, table.storageKey),
    foreignKey({
      name: 'SurveyEmbeddedData_surveyId_workspaceId_fkey',
      columns: [table.surveyId, table.workspaceId],
      foreignColumns: [survey.id, survey.workspaceId],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'SurveyEmbeddedData_embeddedDataId_workspaceId_fkey',
      columns: [table.embeddedDataId, table.workspaceId],
      foreignColumns: [embeddedData.id, embeddedData.workspaceId],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const team = pgTable(
  'Team',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    name: text('name').notNull(),
    organizationId: text('organizationId').notNull(),
  },
  (table) => [
    uniqueIndex('Team_organizationId_name_key').on(table.organizationId, table.name),
    foreignKey({
      name: 'Team_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * Links users to teams with specific roles.
 * Manages team membership and permissions.
 *
 * @property team - The associated team
 * @property user - The team member
 * @property role - User's role within the team
 */
export const teamUser = pgTable(
  'TeamUser',
  {
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    teamId: text('teamId').notNull(),
    userId: text('userId').notNull(),
    role: teamUserRoleEnum('role').notNull(),
  },
  (table) => [
    primaryKey({ name: 'TeamUser_pkey', columns: [table.teamId, table.userId] }),
    index('TeamUser_userId_idx').on(table.userId),
    foreignKey({ name: 'TeamUser_teamId_fkey', columns: [table.teamId], foreignColumns: [team.id] })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({ name: 'TeamUser_userId_fkey', columns: [table.userId], foreignColumns: [user.id] })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * Defines team access to specific workspaces.
 * Manages workspace-level permissions for teams.
 *
 * @property workspace - The accessed workspace
 * @property team - The team receiving access
 * @property permission - Level of access granted
 */
export const workspaceTeam = pgTable(
  'WorkspaceTeam',
  {
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    workspaceId: text('workspaceId').notNull(),
    teamId: text('teamId').notNull(),
    permission: workspaceTeamPermissionEnum('permission').notNull().default('read'),
  },
  (table) => [
    primaryKey({ name: 'WorkspaceTeam_pkey', columns: [table.workspaceId, table.teamId] }),
    index('WorkspaceTeam_teamId_idx').on(table.teamId),
    foreignKey({
      name: 'WorkspaceTeam_workspaceId_fkey',
      columns: [table.workspaceId],
      foreignColumns: [workspace.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'WorkspaceTeam_teamId_fkey',
      columns: [table.teamId],
      foreignColumns: [team.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const chart = pgTable(
  'Chart',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    name: text('name').notNull(),
    type: chartTypeEnum('type').notNull(),
    workspaceId: text('workspaceId').notNull(),
    /** [ChartQuery] - Cube.js query configuration */
    query: jsonb('query').notNull().default({}),
    /** [ChartConfig] - Visualization configuration (colors, labels, formatting) */
    config: jsonb('config').notNull().default({}),
    createdBy: text('createdBy'),
    feedbackDirectoryId: text('feedbackDirectoryId').notNull(),
  },
  (table) => [
    index('Chart_workspaceId_created_at_idx').on(table.workspaceId, table.createdAt),
    index('Chart_feedbackDirectoryId_idx').on(table.feedbackDirectoryId),
    uniqueIndex('Chart_workspaceId_name_key').on(table.workspaceId, table.name),
    foreignKey({
      name: 'Chart_workspaceId_fkey',
      columns: [table.workspaceId],
      foreignColumns: [workspace.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Chart_createdBy_fkey',
      columns: [table.createdBy],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Chart_feedbackDirectoryId_fkey',
      columns: [table.feedbackDirectoryId],
      foreignColumns: [feedbackDirectory.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

/**
 * A workspace's display name for an imported (embedded-data) key of its Feedback Datasets, shown in
 * place of the key in chart legends, headers and the chart builder's field pickers. One per key and
 * workspace; clearing the name deletes the row, and the key shows as collected again. Built-in fields
 * are never renamed here: they keep their translated names.
 *
 * @property key - The embedded-data key exactly as collected, case included
 * @property label - The name shown for the key across the workspace
 */
export const embeddedFieldLabel = pgTable(
  'EmbeddedFieldLabel',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    workspaceId: text('workspaceId').notNull(),
    key: text('key').notNull(),
    label: text('label').notNull(),
  },
  (table) => [
    uniqueIndex('EmbeddedFieldLabel_workspaceId_key_key').on(table.workspaceId, table.key),
    foreignKey({
      name: 'EmbeddedFieldLabel_workspaceId_fkey',
      columns: [table.workspaceId],
      foreignColumns: [workspace.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const dashboard = pgTable(
  'Dashboard',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    name: text('name').notNull(),
    workspaceId: text('workspaceId').notNull(),
    createdBy: text('createdBy'),
    filters: jsonb('filters').notNull().default([]),
  },
  (table) => [
    index('Dashboard_workspaceId_created_at_idx').on(table.workspaceId, table.createdAt),
    uniqueIndex('Dashboard_workspaceId_name_key').on(table.workspaceId, table.name),
    foreignKey({
      name: 'Dashboard_workspaceId_fkey',
      columns: [table.workspaceId],
      foreignColumns: [workspace.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Dashboard_createdBy_fkey',
      columns: [table.createdBy],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

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
export const dashboardWidget = pgTable(
  'DashboardWidget',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    dashboardId: text('dashboardId').notNull(),
    chartId: text('chartId').notNull(),
    /** [WidgetLayout] - Grid layout: { x, y, w, h } */
    layout: jsonb('layout').notNull().default({ x: 0, y: 0, w: 4, h: 3 }),
    order: integer('order').notNull().default(0),
  },
  (table) => [
    index('DashboardWidget_dashboardId_order_idx').on(table.dashboardId, table.order),
    uniqueIndex('DashboardWidget_dashboardId_chartId_key').on(table.dashboardId, table.chartId),
    foreignKey({
      name: 'DashboardWidget_dashboardId_fkey',
      columns: [table.dashboardId],
      foreignColumns: [dashboard.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'DashboardWidget_chartId_fkey',
      columns: [table.chartId],
      foreignColumns: [chart.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * Work left over after a delete commits (ENG-3612, ENG-3614: deleted means deleted): Hub records and
 * storage files that can't be removed inside the database transaction. A row is written in the same
 * transaction as the delete, so a delete that commits always gets its cleanup and one that rolls back
 * queues nothing. A worker drains it with retries, never giving up, and deletes each row once done.
 */
export const deletionCleanup = pgTable(
  'DeletionCleanup',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => createId()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    kind: deletionCleanupKindEnum('kind').notNull(),
    /** Scoping and logs. Not foreign keys: the rows they name are usually gone already. */
    organizationId: text('organizationId').notNull(),
    workspaceId: text('workspaceId').notNull(),
    surveyId: text('surveyId').notNull(),
    /** Hub kinds: the organisation's feedback directories, read inside the delete's transaction. */
    tenantIds: text('tenantIds').array().default([]),
    /** `hubResponses`: the deleted responses whose records go. `hubResponseOutcome`: exactly one live response. */
    responseIds: text('responseIds').array().default([]),
    /** `storageFiles`: the object keys or URLs to delete. */
    fileKeys: text('fileKeys').array().default([]),
    attempts: integer('attempts').notNull().default(0),
    /** Not drained before this, so a late pipeline write can't slip in behind a survey's cleanup. */
    nextAttemptAt: timestamp('nextAttemptAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    /** The last failure: a short code and message, never record or file content. */
    lastError: text('lastError'),
  },
  (table) => [
    index('DeletionCleanup_nextAttemptAt_idx').on(table.nextAttemptAt),
    index('DeletionCleanup_organizationId_created_at_idx').on(
      table.organizationId,
      table.createdAt,
    ),
  ],
)

/**
 * A survey's response attachments built into one ZIP in object storage by a background job and
 * downloaded through a short-lived presigned URL, so an interrupted download can resume (ENG-6862).
 * Only the requester reads it. The reconcile job deletes the archive at `expiresAt`; deleting the
 * survey deletes it with the survey's storage prefix.
 *
 * A partial unique index in the migration allows one queued or running export per requester and
 * survey, so a double click returns the export already in flight.
 */
export const attachmentExport = pgTable(
  'AttachmentExport',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => createId()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    surveyId: text('surveyId').notNull(),
    workspaceId: text('workspaceId').notNull(),
    requestedById: text('requestedById'),
    status: attachmentExportStatusEnum('status').notNull().default('queued'),
    /** The response filter at request time, a `TResponseFilterCriteria`; null exports every response. */
    filterCriteria: jsonb('filterCriteria'),
    /** IANA zone the archive's folder names are stamped in: the org's display zone at request time. */
    timeZone: text('timeZone').notNull(),
    /**
     * Bumped on every enqueue and part of the job id, so a re-dispatch never dedupes against a job an
     * older pod dropped as unknown during a rolling upgrade.
     */
    dispatchCount: integer('dispatchCount').notNull().default(0),
    /** The pre-flight's counts, shown while the archive builds. */
    fileCount: integer('fileCount').notNull(),
    responseCount: integer('responseCount').notNull(),
    /** What the archive actually holds, once built. */
    includedFileCount: integer('includedFileCount'),
    totalBytes: bigint('totalBytes', { mode: 'bigint' }),
    truncated: boolean('truncated').notNull().default(false),
    storageKey: text('storageKey'),
    /** A short code (`stale`, `storage_upload_failed`, …), never a message or file name. */
    errorCode: text('errorCode'),
    startedAt: timestamp('startedAt', { precision: 3 }),
    finishedAt: timestamp('finishedAt', { precision: 3 }),
    expiresAt: timestamp('expiresAt', { precision: 3 }),
  },
  (table) => [
    index('AttachmentExport_surveyId_requestedById_created_at_idx').on(
      table.surveyId,
      table.requestedById,
      table.createdAt,
    ),
    index('AttachmentExport_status_created_at_idx').on(table.status, table.createdAt),
    index('AttachmentExport_expiresAt_idx').on(table.expiresAt),
    index('AttachmentExport_workspaceId_idx').on(table.workspaceId),
    foreignKey({
      name: 'AttachmentExport_surveyId_workspaceId_fkey',
      columns: [table.surveyId, table.workspaceId],
      foreignColumns: [survey.id, survey.workspaceId],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AttachmentExport_workspaceId_fkey',
      columns: [table.workspaceId],
      foreignColumns: [workspace.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AttachmentExport_requestedById_fkey',
      columns: [table.requestedById],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

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
export const feedbackSource = pgTable(
  'FeedbackSource',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    name: text('name').notNull(),
    type: feedbackSourceTypeEnum('type').notNull(),
    status: feedbackSourceStatusEnum('status').notNull().default('active'),
    /**
     * Defaults to completedOnly so a new source agrees with the live ingestion path, which only
     * fires on responseFinished. Existing rows take the default; nothing is re-imported.
     */
    importMode: feedbackSourceImportModeEnum('import_mode').notNull().default('completedOnly'),
    elementScope: feedbackSourceElementScopeEnum('elementScope').notNull().default('specific'),
    workspaceId: text('workspaceId').notNull(),
    feedbackDirectoryId: text('feedbackDirectoryId').notNull(),
    lastSyncAt: timestamp('last_sync_at', { precision: 3 }),
    createdBy: text('created_by'),
  },
  (table) => [
    index('FeedbackSource_type_idx').on(table.type),
    index('FeedbackSource_feedbackDirectoryId_workspaceId_idx').on(
      table.feedbackDirectoryId,
      table.workspaceId,
    ),
    unique('FeedbackSource_id_workspaceId_key').on(table.id, table.workspaceId),
    uniqueIndex('FeedbackSource_workspaceId_name_key').on(table.workspaceId, table.name),
    foreignKey({
      name: 'FeedbackSource_workspaceId_fkey',
      columns: [table.workspaceId],
      foreignColumns: [workspace.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'FeedbackSource_feedbackDirectoryId_fkey',
      columns: [table.feedbackDirectoryId],
      foreignColumns: [feedbackDirectory.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'FeedbackSource_created_by_fkey',
      columns: [table.createdBy],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'FeedbackSource_feedbackDirectoryId_workspaceId_fkey',
      columns: [table.feedbackDirectoryId, table.workspaceId],
      foreignColumns: [
        feedbackDirectoryWorkspace.feedbackDirectoryId,
        feedbackDirectoryWorkspace.workspaceId,
      ],
    })
      .onDelete('no action')
      .onUpdate('cascade'),
  ],
)

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
export const feedbackSourceFormbricksMapping = pgTable(
  'FeedbackSourceFormbricksMapping',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    feedbackSourceId: text('feedback_source_id').notNull(),
    workspaceId: text('workspaceId').notNull(),
    surveyId: text('surveyId').notNull(),
    elementId: text('elementId').notNull(),
    hubFieldType: hubFieldTypeEnum('hubFieldType').notNull(),
    customFieldLabel: text('custom_field_label'),
  },
  (table) => [
    index('FeedbackSourceFormbricksMapping_workspaceId_surveyId_idx').on(
      table.workspaceId,
      table.surveyId,
    ),
    index('FeedbackSourceFormbricksMapping_surveyId_idx').on(table.surveyId),
    uniqueIndex('FeedbackSourceFormbricksMapping_workspaceId_feedbackSourceId_su').on(
      table.workspaceId,
      table.feedbackSourceId,
      table.surveyId,
      table.elementId,
    ),
    foreignKey({
      name: 'FeedbackSourceFormbricksMapping_feedbackSourceId_workspaceId_fk',
      columns: [table.feedbackSourceId, table.workspaceId],
      foreignColumns: [feedbackSource.id, feedbackSource.workspaceId],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'FeedbackSourceFormbricksMapping_surveyId_workspaceId_fkey',
      columns: [table.surveyId, table.workspaceId],
      foreignColumns: [survey.id, survey.workspaceId],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const feedbackSourceFieldMapping = pgTable(
  'FeedbackSourceFieldMapping',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    feedbackSourceId: text('feedback_source_id').notNull(),
    workspaceId: text('workspaceId').notNull(),
    sourceFieldId: text('source_field_id').notNull(),
    targetFieldId: text('target_field_id').notNull(),
    staticValue: text('static_value'),
  },
  (table) => [
    uniqueIndex('FeedbackSourceFieldMapping_workspaceId_feedbackSourceId_sourceF').on(
      table.workspaceId,
      table.feedbackSourceId,
      table.sourceFieldId,
      table.targetFieldId,
    ),
    foreignKey({
      name: 'FeedbackSourceFieldMapping_feedbackSourceId_workspaceId_fkey',
      columns: [table.feedbackSourceId, table.workspaceId],
      foreignColumns: [feedbackSource.id, feedbackSource.workspaceId],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const feedbackDirectory = pgTable(
  'FeedbackDirectory',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    name: text('name').notNull(),
    isArchived: boolean('isArchived').notNull().default(false),
    organizationId: text('organizationId').notNull(),
  },
  (table) => [
    uniqueIndex('FeedbackDirectory_organizationId_name_key').on(table.organizationId, table.name),
    foreignKey({
      name: 'FeedbackDirectory_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * Links feedback directories to workspaces (workspaces).
 * Manages which workspaces can access a given directory.
 *
 * @property feedbackDirectory - The directory being accessed
 * @property workspace - The workspace receiving access
 */
export const feedbackDirectoryWorkspace = pgTable(
  'FeedbackDirectoryWorkspace',
  {
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    feedbackDirectoryId: text('feedbackDirectoryId').notNull(),
    workspaceId: text('workspaceId').notNull(),
  },
  (table) => [
    primaryKey({
      name: 'FeedbackDirectoryWorkspace_pkey',
      columns: [table.feedbackDirectoryId, table.workspaceId],
    }),
    index('FeedbackDirectoryWorkspace_workspaceId_idx').on(table.workspaceId),
    foreignKey({
      name: 'FeedbackDirectoryWorkspace_feedbackDirectoryId_fkey',
      columns: [table.feedbackDirectoryId],
      foreignColumns: [feedbackDirectory.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'FeedbackDirectoryWorkspace_workspaceId_fkey',
      columns: [table.workspaceId],
      foreignColumns: [workspace.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * ENG-3282: per-scope readiness of a SpiceDB projection. A row with `readyAt` set means the scope was
 * backfilled and audited clean, so the evaluator may enforce it. Written only by the backfill CLI.
 */
export const authzedProjectionScopeState = pgTable('AuthzedProjectionScopeState', {
  scope: text('scope').primaryKey(),
  readyAt: timestamp('readyAt', { precision: 3 }),
  readyBy: text('readyBy'),
  updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
})

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
export const authzedProjectionOutbox = pgTable('AuthzedProjectionOutbox', {
  id: text('id')
    .primaryKey()
    .$defaultFn(() => crypto.randomUUID()),
  targetType: text('targetType').notNull(),
  primaryId: text('primaryId').notNull(),
  secondaryId: text('secondaryId'),
  isRevocation: boolean('isRevocation').notNull().default(false),
  /** Total delivery attempts. Drives the retry backoff and reports how hard delivery was tried. */
  attempts: integer('attempts').notNull().default(0),
  /**
   * Non-retryable failures attributable to this event alone. The only input to dead-lettering, so
   * that a SpiceDB outage of any duration cannot dead-letter events that were never the problem.
   */
  permanentFailures: integer('permanentFailures').notNull().default(0),
  availableAt: timestamp('availableAt', { precision: 3 })
    .notNull()
    .default(sql`CURRENT_TIMESTAMP`)
    .$defaultFn(utcNow),
  leasedAt: timestamp('leasedAt', { precision: 3 }),
  leaseExpiresAt: timestamp('leaseExpiresAt', { precision: 3 }),
  leaseOwner: text('leaseOwner'),
  processedAt: timestamp('processedAt', { precision: 3 }),
  deadLetteredAt: timestamp('deadLetteredAt', { precision: 3 }),
  lastAttemptAt: timestamp('lastAttemptAt', { precision: 3 }),
  lastErrorCode: text('lastErrorCode'),
  createdAt: timestamp('createdAt', { precision: 3 })
    .notNull()
    .default(sql`CURRENT_TIMESTAMP`)
    .$defaultFn(utcNow),
  updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
})

export const onlineUpgradeState = pgTable('OnlineUpgradeState', {
  release: text('release').primaryKey(),
  targetImage: text('targetImage').notNull(),
  phase: text('phase').notNull(),
  completedPhases: text('completedPhases').array().default([]),
  surveyCursor: text('surveyCursor'),
  surveyCount: bigint('surveyCount', { mode: 'bigint' })
    .notNull()
    .default(sql`0`),
  baselineCompletedAt: timestamp('baselineCompletedAt', { precision: 3 }),
  forwardOnlyAt: timestamp('forwardOnlyAt', { precision: 3 }),
  queuePausedAt: timestamp('queuePausedAt', { precision: 3 }),
  queuePauseDeadlineAt: timestamp('queuePauseDeadlineAt', { precision: 3 }),
  queueResumedAt: timestamp('queueResumedAt', { precision: 3 }),
  activatedAt: timestamp('activatedAt', { precision: 3 }),
  updatedAt: timestamp('updatedAt', { precision: 3 })
    .notNull()
    .default(sql`CURRENT_TIMESTAMP`)
    .$defaultFn(utcNow)
    .$onUpdate(utcNow),
})

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
export const retentionPolicy = pgTable(
  'RetentionPolicy',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => createId()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    organizationId: text('organizationId').notNull(),
    entity: retentionEntityEnum('entity').notNull(),
    enabled: boolean('enabled').notNull().default(false),
    /**
     * When the current configuration took effect: the policy was switched on or unpaused, or changed
     * in a way that could bring an action earlier (a shorter period or warning, different
     * conditions). Lengthening does not move it. The warning always runs in full from here, so no
     * change can delete data on the night it is made (ENG-3614). Also the clock for members with no
     * recorded sign-in.
     */
    enabledAt: timestamp('enabledAt', { precision: 3 }),
    /**
     * How long before the action the notice goes out. The action never runs before the notice has
     * been out for this long, so the warning always runs in full.
     */
    warnDays: integer('warnDays').notNull().default(60),
    /**
     * The period, counted from the target's clock, after which the policy acts: responses are
     * deleted, surveys archived (the archive purge deletes them 30 days later), members deactivated.
     */
    periodDays: integer('periodDays').notNull(),
    /** Surveys policy only. */
    conditions: retentionSurveyConditionEnum('conditions').array().default([]),
    /**
     * The last night the sweep's time ran out before it reached this policy's organisation, licensed.
     * The organisation was due and the sweep was running, so it counts as reached for the gap that
     * restarts the warning (`RETENTION_SWEEP_GAP_MS`): a backlog of organisations never restarts warnings.
     */
    deferredAt: timestamp('deferredAt', { precision: 3 }),
    updatedById: text('updatedById'),
  },
  (table) => [
    index('RetentionPolicy_updatedById_idx').on(table.updatedById),
    uniqueIndex('RetentionPolicy_organizationId_entity_key').on(table.organizationId, table.entity),
    foreignKey({
      name: 'RetentionPolicy_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'RetentionPolicy_updatedById_fkey',
      columns: [table.updatedById],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

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
export const retentionExemption = pgTable(
  'RetentionExemption',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => createId()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    organizationId: text('organizationId').notNull(),
    entity: retentionEntityEnum('entity').notNull(),
    surveyId: text('surveyId'),
    until: timestamp('until', { precision: 3 }).notNull(),
    source: retentionExemptionSourceEnum('source').notNull().default('manual'),
    /** Required for `manual`, absent for `restore` (`RetentionExemption_reason_check`). */
    reason: text('reason'),
    createdById: text('createdById'),
    revokedAt: timestamp('revokedAt', { precision: 3 }),
    revokedById: text('revokedById'),
  },
  (table) => [
    index('RetentionExemption_organizationId_created_at_id_idx').on(
      table.organizationId,
      table.createdAt,
      table.id,
    ),
    index('RetentionExemption_surveyId_idx').on(table.surveyId),
    index('RetentionExemption_createdById_idx').on(table.createdById),
    index('RetentionExemption_revokedById_idx').on(table.revokedById),
    foreignKey({
      name: 'RetentionExemption_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'RetentionExemption_surveyId_fkey',
      columns: [table.surveyId],
      foreignColumns: [survey.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'RetentionExemption_createdById_fkey',
      columns: [table.createdById],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'RetentionExemption_revokedById_fkey',
      columns: [table.revokedById],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

/**
 * One sweep of one policy for one organisation: the History tab's rows. `finishedAt` stays null
 * while the run is in progress, or if it never finished.
 */
export const retentionRun = pgTable(
  'RetentionRun',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => createId()),
    organizationId: text('organizationId').notNull(),
    entity: retentionEntityEnum('entity').notNull(),
    kind: retentionRunKindEnum('kind').notNull().default('sweep'),
    startedAt: timestamp('startedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    finishedAt: timestamp('finishedAt', { precision: 3 }),
    notifiedCount: integer('notifiedCount').notNull().default(0),
    /** Archived surveys, or deactivated members. */
    archivedCount: integer('archivedCount').notNull().default(0),
    /** Deleted responses, or, on a purge run, deleted surveys. */
    deletedCount: integer('deletedCount').notNull().default(0),
    skippedCount: integer('skippedCount').notNull().default(0),
    /** False when the run changed nothing. History hides those by default. */
    hasChanges: boolean('hasChanges').notNull().default(false),
    /**
     * Where the run's scan of candidates stopped for want of time, so the next run resumes there instead
     * of rescanning the same first targets every night. Null when the scan reached the end.
     */
    scanCursor: text('scanCursor'),
  },
  (table) => [
    index('RetentionRun_organizationId_startedAt_id_idx').on(
      table.organizationId,
      table.startedAt,
      table.id,
    ),
    foreignKey({
      name: 'RetentionRun_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * What a run did, one row per target: the History CSV. Ids, counts, names and the recipient's email,
 * never response content. `targetId` is not a foreign key on purpose: the target is often gone, and
 * `targetName` is a snapshot taken when the row is written so the CSV still says what it was.
 */
export const retentionRunItem = pgTable(
  'RetentionRunItem',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => createId()),
    runId: text('runId').notNull(),
    targetType: retentionTargetTypeEnum('targetType').notNull(),
    targetId: text('targetId').notNull(),
    targetName: text('targetName'),
    action: retentionRunItemActionEnum('action').notNull(),
    /** 1, except for grouped response deletions, where it is the number of responses deleted. */
    count: integer('count').notNull().default(1),
    /** Who was emailed, for `notified` rows. */
    recipient: text('recipient'),
    skipReason: retentionSkipReasonEnum('skipReason'),
  },
  (table) => [
    index('RetentionRunItem_runId_id_idx').on(table.runId, table.id),
    index('RetentionRunItem_targetId_idx').on(table.targetId),
    foreignKey({
      name: 'RetentionRunItem_runId_fkey',
      columns: [table.runId],
      foreignColumns: [retentionRun.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

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
export const retentionNotice = pgTable(
  'RetentionNotice',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => createId()),
    organizationId: text('organizationId').notNull(),
    entity: retentionEntityEnum('entity').notNull(),
    surveyId: text('surveyId'),
    userId: text('userId'),
    /** When the notice was claimed for sending. A claim that never got delivered can be taken over. */
    sentAt: timestamp('sentAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    /**
     * When the notice was handed to the mail transport, or recorded without one (no SMTP). The
     * warning only counts from a delivered notice, so a crash between claim and send can't let the
     * action run without it.
     */
    deliveredAt: timestamp('deliveredAt', { precision: 3 }),
    /** Whether an email actually went out. False when SMTP isn't configured. */
    emailSent: boolean('emailSent').notNull().default(false),
    /**
     * Set by each claim; delivery is recorded only for the claim that sent the email, so a slow sender
     * can't mark a newer claim delivered.
     */
    claimToken: text('claimToken'),
    /**
     * The target's clock the notice was computed for (surveys and members). A notice counts only while
     * the target's clock is still this one, so activity between the sweep's read and its claim voids
     * it. Null for the responses reminder, which is sent once per survey and follows no one clock, and
     * for a survey notice voided for good: a person deleted or moved one of its responses, or the
     * responses policy's deletion left it computed for another clock (see `NOTICE_FOLLOWS_CLOCK`).
     */
    clockAt: timestamp('clockAt', { precision: 3 }),
  },
  (table) => [
    index('RetentionNotice_organizationId_idx').on(table.organizationId),
    uniqueIndex('RetentionNotice_surveyId_entity_key').on(table.surveyId, table.entity),
    uniqueIndex('RetentionNotice_userId_organizationId_entity_key').on(
      table.userId,
      table.organizationId,
      table.entity,
    ),
    foreignKey({
      name: 'RetentionNotice_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'RetentionNotice_surveyId_fkey',
      columns: [table.surveyId],
      foreignColumns: [survey.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'RetentionNotice_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * When a survey was last restored from the archive. A restore starts the surveys policy's "no response"
 * and "no change" periods over (`getSurveyRetentionClock`); only ever moves forward. A table of its own
 * so retention adds nothing to the survey row.
 */
export const retentionSurveyRestore = pgTable(
  'RetentionSurveyRestore',
  {
    surveyId: text('surveyId').primaryKey(),
    restoredAt: timestamp('restoredAt', { precision: 3 }).notNull(),
  },
  (table) => [
    foreignKey({
      name: 'RetentionSurveyRestore_surveyId_fkey',
      columns: [table.surveyId],
      foreignColumns: [survey.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const workflow = pgTable(
  'Workflow',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => createId()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    name: text('name').notNull(),
    description: text('description'),
    status: workflowStatusEnum('status').notNull().default('draft'),
    workspaceId: text('workspaceId').notNull(),
    createdBy: text('createdBy'),
    /** [WorkflowDefinition] */
    definition: jsonb('definition').notNull(),
    /**
     * Webhook Secret URL token (ADRs 0048, 0057): SHA-256 hash to check calls, and an encrypted copy so
     * the editor can show the URL again. Never part of the definition, its versions or its copies.
     */
    webhookTokenHash: text('webhookTokenHash'),
    webhookTokenEncrypted: text('webhookTokenEncrypted'),
  },
  (table) => [
    index('Workflow_workspaceId_status_idx').on(table.workspaceId, table.status),
    index('Workflow_workspaceId_updated_at_idx').on(table.workspaceId, table.updatedAt),
    unique('Workflow_id_workspaceId_key').on(table.id, table.workspaceId),
    foreignKey({
      name: 'Workflow_workspaceId_fkey',
      columns: [table.workspaceId],
      foreignColumns: [workspace.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Workflow_createdBy_fkey',
      columns: [table.createdBy],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const workflowVersion = pgTable(
  'WorkflowVersion',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => createId()),
    workflowId: text('workflowId').notNull(),
    workspaceId: text('workspaceId').notNull(),
    version: integer('version').notNull(),
    /** [WorkflowExecutableDefinition] */
    definition: jsonb('definition').notNull(),
    publishedAt: timestamp('publishedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    publishedBy: text('publishedBy'),
  },
  (table) => [
    index('WorkflowVersion_workflowId_workspaceId_idx').on(table.workflowId, table.workspaceId),
    index('WorkflowVersion_publishedBy_idx').on(table.publishedBy),
    uniqueIndex('WorkflowVersion_workflowId_version_key').on(table.workflowId, table.version),
    foreignKey({
      name: 'WorkflowVersion_workflowId_workspaceId_fkey',
      columns: [table.workflowId, table.workspaceId],
      foreignColumns: [workflow.id, workflow.workspaceId],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'WorkflowVersion_publishedBy_fkey',
      columns: [table.publishedBy],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const workflowRun = pgTable(
  'WorkflowRun',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => createId()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    workflowId: text('workflowId').notNull(),
    workspaceId: text('workspaceId').notNull(),
    workflowVersionId: text('workflowVersionId'),
    responseId: text('responseId'),
    status: workflowRunStatusEnum('status').notNull().default('queued'),
    triggerType: text('triggerType').notNull(),
    surveyId: text('surveyId'),
    isDryRun: boolean('isDryRun').notNull().default(false),
    idempotencyKey: text('idempotencyKey'),
    attempt: integer('attempt').notNull().default(0),
    nextAttemptAt: timestamp('nextAttemptAt', { precision: 3 }),
    lastErrorAt: timestamp('lastErrorAt', { precision: 3 }),
    /** [WorkflowTriggerRunPayload] */
    triggerPayload: jsonb('triggerPayload').notNull(),
    /** [WorkflowRunData] */
    data: jsonb('data').notNull().default({}),
    error: text('error'),
    dispatchedAt: timestamp('dispatchedAt', { precision: 3 }),
    startedAt: timestamp('startedAt', { precision: 3 }),
    finishedAt: timestamp('finishedAt', { precision: 3 }),
    /** The caller's IP for a run started by a Webhook Secret URL call; shown in run detail. */
    callerIp: text('callerIp'),
    /**
     * When the run retention job cleared the run's content (ADR 0044): payload contents, run data, step
     * input, output and errors, the run error and the caller IP. Null while the content is kept.
     */
    contentDeletedAt: timestamp('contentDeletedAt', { precision: 3 }),
    /** The retention period in days that applied when the content was cleared ("Details deleted after N days"). */
    contentRetentionDays: integer('contentRetentionDays'),
    /**
     * The draft definition a Test run executed (ADR 0049), so its graph can be drawn after the draft
     * changes. Null for live runs (they point at `workflowVersionId`) and for Test runs saved before it
     * existed. Configuration like a version's definition, not run content: retention leaves it alone.
     * [WorkflowExecutableDefinition]
     */
    testRunDefinition: jsonb('testRunDefinition'),
  },
  (table) => [
    index('WorkflowRun_workspaceId_created_at_idx').on(table.workspaceId, table.createdAt),
    index('WorkflowRun_workflowId_workspaceId_created_at_idx').on(
      table.workflowId,
      table.workspaceId,
      table.createdAt,
    ),
    index('WorkflowRun_status_created_at_idx').on(table.status, table.createdAt),
    index('WorkflowRun_responseId_idx').on(table.responseId),
    index('WorkflowRun_isDryRun_idx').on(table.isDryRun),
    index('WorkflowRun_isDryRun_created_at_idx').on(table.isDryRun, table.createdAt),
    index('WorkflowRun_nextAttemptAt_idx').on(table.nextAttemptAt),
    index('WorkflowRun_contentDeletedAt_finishedAt_idx').on(
      table.contentDeletedAt,
      table.finishedAt,
    ),
    uniqueIndex('WorkflowRun_workflowId_idempotencyKey_key').on(
      table.workflowId,
      table.idempotencyKey,
    ),
    foreignKey({
      name: 'WorkflowRun_workflowId_workspaceId_fkey',
      columns: [table.workflowId, table.workspaceId],
      foreignColumns: [workflow.id, workflow.workspaceId],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'WorkflowRun_workspaceId_fkey',
      columns: [table.workspaceId],
      foreignColumns: [workspace.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'WorkflowRun_workflowVersionId_fkey',
      columns: [table.workflowVersionId],
      foreignColumns: [workflowVersion.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'WorkflowRun_responseId_fkey',
      columns: [table.responseId],
      foreignColumns: [response.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * Per-call data kept beside a workflow (ADR 0048): one row per source, `trigger` for the Webhook
 * trigger's last payload (plus the last call's time and IP), or an HTTP request step id for its last
 * Test request response. Written on every call, so it lives here and never bumps `Workflow.updatedAt`.
 */
export const workflowSample = pgTable(
  'WorkflowSample',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => createId()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    workflowId: text('workflowId').notNull(),
    workspaceId: text('workspaceId').notNull(),
    source: text('source').notNull(),
    /** [WorkflowSampleData] */
    sample: jsonb('sample').notNull(),
    lastCalledAt: timestamp('lastCalledAt', { precision: 3 }),
    lastCallerIp: text('lastCallerIp'),
  },
  (table) => [
    index('WorkflowSample_workflowId_workspaceId_idx').on(table.workflowId, table.workspaceId),
    index('WorkflowSample_updated_at_idx').on(table.updatedAt),
    uniqueIndex('WorkflowSample_workflowId_source_key').on(table.workflowId, table.source),
    foreignKey({
      name: 'WorkflowSample_workflowId_workspaceId_fkey',
      columns: [table.workflowId, table.workspaceId],
      foreignColumns: [workflow.id, workflow.workspaceId],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const workflowRunLog = pgTable(
  'WorkflowRunLog',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => createId()),
    runId: text('runId').notNull(),
    sequence: integer('sequence').notNull(),
    stepId: text('stepId').notNull(),
    stepType: text('stepType').notNull(),
    status: workflowRunLogStatusEnum('status').notNull().default('pending'),
    /** [WorkflowRunLogInput] */
    input: jsonb('input').notNull().default({}),
    /** [WorkflowRunLogOutput] */
    output: jsonb('output').notNull().default({}),
    error: text('error'),
    startedAt: timestamp('startedAt', { precision: 3 }),
    finishedAt: timestamp('finishedAt', { precision: 3 }),
  },
  (table) => [
    index('WorkflowRunLog_runId_sequence_idx').on(table.runId, table.sequence),
    index('WorkflowRunLog_status_startedAt_idx').on(table.status, table.startedAt),
    uniqueIndex('WorkflowRunLog_runId_stepId_key').on(table.runId, table.stepId),
    foreignKey({
      name: 'WorkflowRunLog_runId_fkey',
      columns: [table.runId],
      foreignColumns: [workflowRun.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * Encrypted value of an HTTP request secret header (ENG-7168, ADR 0043). The definition and every version
 * store only `{name, secret: true, secretId}`; Replace overwrites this row in place, so the enabled
 * version uses the new value at once. Rows nothing references are deleted on Enable; all are deleted
 * with the workflow. A value is only sent to the origin it was set for (ADR 0051).
 */
export const workflowSecret = pgTable(
  'WorkflowSecret',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => createId()),
    createdAt: timestamp('created_at', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updated_at', { precision: 3 }).notNull().$onUpdate(utcNow),
    workflowId: text('workflowId').notNull(),
    workspaceId: text('workspaceId').notNull(),
    /** `symmetricEncrypt` ciphertext (AES-256-GCM). Never logged, never returned. */
    value: text('value').notNull(),
    /**
     * Origin (scheme, host, port) of the step's URL when the value was set or replaced; the value is only
     * sent there (ADR 0051). Null for a value set before host binding: held to the step's saved URL
     * until it is set again.
     */
    boundOrigin: text('boundOrigin'),
  },
  (table) => [
    index('WorkflowSecret_workflowId_workspaceId_idx').on(table.workflowId, table.workspaceId),
    foreignKey({
      name: 'WorkflowSecret_workflowId_workspaceId_fkey',
      columns: [table.workflowId, table.workspaceId],
      foreignColumns: [workflow.id, workflow.workspaceId],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const relations = defineRelations(
  {
    creative,
    creativeDailyStat,
    creativeDailyPageStat,
    creativeTrigger,
    creativeLanguage,
    outboundEmailAudit,
    outboundEmailAuditErasureIntent,
    webhook,
    contactAttribute,
    contactAttributeKey,
    contact,
    response,
    tag,
    tagsOnResponses,
    display,
    surveyTrigger,
    surveyAttributeFilter,
    survey,
    surveyQuota,
    responseQuotaLink,
    surveyFollowUp,
    actionClass,
    integration,
    dataMigration,
    workspace,
    organization,
    organizationBilling,
    membership,
    invite,
    apiKey,
    apiKeyWorkspace,
    account,
    session,
    verificationToken,
    twoFactor,
    jwks,
    oauthClient,
    oauthAccessToken,
    oauthRefreshToken,
    oauthConsent,
    mcpGrant,
    mcpGrantOrganization,
    mcpGrantWorkspace,
    oauthResource,
    oauthClientResource,
    oauthClientAssertion,
    passwordResetToken,
    user,
    segment,
    language,
    surveyLanguage,
    embeddedData,
    surveyEmbeddedData,
    team,
    teamUser,
    workspaceTeam,
    chart,
    embeddedFieldLabel,
    dashboard,
    dashboardWidget,
    deletionCleanup,
    attachmentExport,
    feedbackSource,
    feedbackSourceFormbricksMapping,
    feedbackSourceFieldMapping,
    feedbackDirectory,
    feedbackDirectoryWorkspace,
    authzedProjectionScopeState,
    authzedProjectionOutbox,
    onlineUpgradeState,
    retentionPolicy,
    retentionExemption,
    retentionRun,
    retentionRunItem,
    retentionNotice,
    retentionSurveyRestore,
    workflow,
    workflowVersion,
    workflowRun,
    workflowSample,
    workflowRunLog,
    workflowSecret,
  },
  (r) => ({
    creative: {
      workspace: r.one.workspace({
        from: r.creative.workspaceId,
        to: r.workspace.id,
        optional: false,
      }),
      creator: r.one.user({ from: r.creative.createdBy, to: r.user.id }),
      survey: r.one.survey({
        from: [r.creative.surveyId, r.creative.workspaceId],
        to: [r.survey.id, r.survey.workspaceId],
      }),
      triggers: r.many.creativeTrigger(),
      languages: r.many.creativeLanguage(),
      dailyStats: r.many.creativeDailyStat(),
      dailyPageStats: r.many.creativeDailyPageStat(),
    },
    creativeDailyStat: {
      creative: r.one.creative({
        from: r.creativeDailyStat.creativeId,
        to: r.creative.id,
        optional: false,
      }),
    },
    creativeDailyPageStat: {
      creative: r.one.creative({
        from: r.creativeDailyPageStat.creativeId,
        to: r.creative.id,
        optional: false,
      }),
    },
    creativeTrigger: {
      creative: r.one.creative({
        from: r.creativeTrigger.creativeId,
        to: r.creative.id,
        optional: false,
      }),
      actionClass: r.one.actionClass({
        from: r.creativeTrigger.actionClassId,
        to: r.actionClass.id,
        optional: false,
      }),
    },
    creativeLanguage: {
      language: r.one.language({
        from: r.creativeLanguage.languageId,
        to: r.language.id,
        optional: false,
      }),
      creative: r.one.creative({
        from: r.creativeLanguage.creativeId,
        to: r.creative.id,
        optional: false,
      }),
    },
    webhook: {
      workspace: r.one.workspace({
        from: r.webhook.workspaceId,
        to: r.workspace.id,
        optional: false,
      }),
    },
    contactAttribute: {
      attributeKey: r.one.contactAttributeKey({
        from: r.contactAttribute.attributeKeyId,
        to: r.contactAttributeKey.id,
        optional: false,
      }),
      contact: r.one.contact({
        from: r.contactAttribute.contactId,
        to: r.contact.id,
        optional: false,
      }),
    },
    contactAttributeKey: {
      workspace: r.one.workspace({
        from: r.contactAttributeKey.workspaceId,
        to: r.workspace.id,
        optional: false,
      }),
      attributes: r.many.contactAttribute(),
      attributeFilters: r.many.surveyAttributeFilter(),
    },
    contact: {
      workspace: r.one.workspace({
        from: r.contact.workspaceId,
        to: r.workspace.id,
        optional: false,
      }),
      responses: r.many.response(),
      attributes: r.many.contactAttribute(),
      displays: r.many.display(),
    },
    response: {
      survey: r.one.survey({ from: r.response.surveyId, to: r.survey.id, optional: false }),
      contact: r.one.contact({ from: r.response.contactId, to: r.contact.id }),
      tags: r.many.tagsOnResponses(),
      quotaLinks: r.many.responseQuotaLink(),
      display: r.one.display({ from: r.response.displayId, to: r.display.id }),
      workflowRuns: r.many.workflowRun(),
    },
    tag: {
      responses: r.many.tagsOnResponses(),
      workspace: r.one.workspace({ from: r.tag.workspaceId, to: r.workspace.id, optional: false }),
    },
    tagsOnResponses: {
      response: r.one.response({
        from: r.tagsOnResponses.responseId,
        to: r.response.id,
        optional: false,
      }),
      tag: r.one.tag({ from: r.tagsOnResponses.tagId, to: r.tag.id, optional: false }),
    },
    display: {
      survey: r.one.survey({ from: r.display.surveyId, to: r.survey.id, optional: false }),
      contact: r.one.contact({ from: r.display.contactId, to: r.contact.id }),
      response: r.one.response({ from: r.display.id, to: r.response.displayId }),
    },
    surveyTrigger: {
      survey: r.one.survey({ from: r.surveyTrigger.surveyId, to: r.survey.id, optional: false }),
      actionClass: r.one.actionClass({
        from: r.surveyTrigger.actionClassId,
        to: r.actionClass.id,
        optional: false,
      }),
    },
    surveyAttributeFilter: {
      attributeKey: r.one.contactAttributeKey({
        from: r.surveyAttributeFilter.attributeKeyId,
        to: r.contactAttributeKey.id,
        optional: false,
      }),
      survey: r.one.survey({
        from: r.surveyAttributeFilter.surveyId,
        to: r.survey.id,
        optional: false,
      }),
    },
    survey: {
      workspace: r.one.workspace({
        from: r.survey.workspaceId,
        to: r.workspace.id,
        optional: false,
      }),
      creator: r.one.user({ from: r.survey.createdBy, to: r.user.id, alias: 'SurveyCreator' }),
      owner: r.one.user({ from: r.survey.ownerId, to: r.user.id, alias: 'SurveyOwner' }),
      responses: r.many.response(),
      quotas: r.many.surveyQuota(),
      triggers: r.many.surveyTrigger(),
      attributeFilters: r.many.surveyAttributeFilter(),
      displays: r.many.display(),
      segment: r.one.segment({ from: r.survey.segmentId, to: r.segment.id }),
      languages: r.many.surveyLanguage(),
      embeddedData: r.many.embeddedData(),
      embeddedDataLinks: r.many.surveyEmbeddedData(),
      followUps: r.many.surveyFollowUp(),
      feedbackSourceMappings: r.many.feedbackSourceFormbricksMapping(),
      retentionExemptions: r.many.retentionExemption(),
      retentionNotices: r.many.retentionNotice(),
      retentionRestore: r.one.retentionSurveyRestore({
        from: r.survey.id,
        to: r.retentionSurveyRestore.surveyId,
      }),
      creatives: r.many.creative(),
      attachmentExports: r.many.attachmentExport(),
    },
    surveyQuota: {
      survey: r.one.survey({ from: r.surveyQuota.surveyId, to: r.survey.id, optional: false }),
      quotaLinks: r.many.responseQuotaLink(),
    },
    responseQuotaLink: {
      response: r.one.response({
        from: r.responseQuotaLink.responseId,
        to: r.response.id,
        optional: false,
      }),
      quota: r.one.surveyQuota({
        from: r.responseQuotaLink.quotaId,
        to: r.surveyQuota.id,
        optional: false,
      }),
    },
    surveyFollowUp: {
      survey: r.one.survey({ from: r.surveyFollowUp.surveyId, to: r.survey.id, optional: false }),
    },
    actionClass: {
      workspace: r.one.workspace({
        from: r.actionClass.workspaceId,
        to: r.workspace.id,
        optional: false,
      }),
      surveyTriggers: r.many.surveyTrigger(),
      creativeTriggers: r.many.creativeTrigger(),
    },
    integration: {
      workspace: r.one.workspace({
        from: r.integration.workspaceId,
        to: r.workspace.id,
        optional: false,
      }),
    },
    workspace: {
      organization: r.one.organization({
        from: r.workspace.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      languages: r.many.language(),
      embeddedData: r.many.embeddedData(),
      workspaceTeams: r.many.workspaceTeam(),
      feedbackDirectoryWorkspaces: r.many.feedbackDirectoryWorkspace(),
      charts: r.many.chart(),
      dashboards: r.many.dashboard(),
      embeddedFieldLabels: r.many.embeddedFieldLabel(),
      creatives: r.many.creative(),
      attachmentExports: r.many.attachmentExport(),
      surveys: r.many.survey(),
      contacts: r.many.contact(),
      actionClasses: r.many.actionClass(),
      contactAttributeKeys: r.many.contactAttributeKey(),
      webhooks: r.many.webhook(),
      tags: r.many.tag(),
      segments: r.many.segment(),
      integrations: r.many.integration(),
      apiKeyWorkspaces: r.many.apiKeyWorkspace(),
      feedbackSources: r.many.feedbackSource(),
      workflows: r.many.workflow(),
      workflowRuns: r.many.workflowRun(),
      mcpGrants: r.many.mcpGrantWorkspace(),
    },
    organization: {
      memberships: r.many.membership(),
      workspaces: r.many.workspace(),
      billing: r.one.organizationBilling({
        from: r.organization.id,
        to: r.organizationBilling.organizationId,
      }),
      invites: r.many.invite(),
      teams: r.many.team(),
      apiKeys: r.many.apiKey(),
      feedbackDirectories: r.many.feedbackDirectory(),
      retentionPolicies: r.many.retentionPolicy(),
      retentionExemptions: r.many.retentionExemption(),
      retentionRuns: r.many.retentionRun(),
      retentionNotices: r.many.retentionNotice(),
      mcpGrants: r.many.mcpGrantOrganization(),
    },
    organizationBilling: {
      organization: r.one.organization({
        from: r.organizationBilling.organizationId,
        to: r.organization.id,
        optional: false,
      }),
    },
    membership: {
      organization: r.one.organization({
        from: r.membership.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      user: r.one.user({ from: r.membership.userId, to: r.user.id, optional: false }),
    },
    invite: {
      organization: r.one.organization({
        from: r.invite.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      creator: r.one.user({
        from: r.invite.creatorId,
        to: r.user.id,
        optional: false,
        alias: 'inviteCreatedBy',
      }),
      acceptor: r.one.user({ from: r.invite.acceptorId, to: r.user.id, alias: 'inviteAcceptedBy' }),
    },
    apiKey: {
      organization: r.one.organization({
        from: r.apiKey.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      apiKeyWorkspaces: r.many.apiKeyWorkspace(),
    },
    apiKeyWorkspace: {
      apiKey: r.one.apiKey({ from: r.apiKeyWorkspace.apiKeyId, to: r.apiKey.id, optional: false }),
      workspace: r.one.workspace({
        from: r.apiKeyWorkspace.workspaceId,
        to: r.workspace.id,
        optional: false,
      }),
    },
    account: { user: r.one.user({ from: r.account.userId, to: r.user.id }) },
    session: {
      user: r.one.user({ from: r.session.userId, to: r.user.id, optional: false }),
      oauthAccessTokens: r.many.oauthAccessToken(),
      oauthRefreshTokens: r.many.oauthRefreshToken(),
    },
    twoFactor: { user: r.one.user({ from: r.twoFactor.userId, to: r.user.id, optional: false }) },
    oauthClient: {
      user: r.one.user({ from: r.oauthClient.userId, to: r.user.id }),
      accessTokens: r.many.oauthAccessToken(),
      refreshTokens: r.many.oauthRefreshToken(),
      consents: r.many.oauthConsent(),
      resources: r.many.oauthClientResource(),
      mcpGrants: r.many.mcpGrant(),
    },
    oauthAccessToken: {
      client: r.one.oauthClient({
        from: r.oauthAccessToken.clientId,
        to: r.oauthClient.clientId,
        optional: false,
      }),
      session: r.one.session({ from: r.oauthAccessToken.sessionId, to: r.session.id }),
      user: r.one.user({ from: r.oauthAccessToken.userId, to: r.user.id }),
      refresh: r.one.oauthRefreshToken({
        from: r.oauthAccessToken.refreshId,
        to: r.oauthRefreshToken.id,
      }),
    },
    oauthRefreshToken: {
      client: r.one.oauthClient({
        from: r.oauthRefreshToken.clientId,
        to: r.oauthClient.clientId,
        optional: false,
      }),
      session: r.one.session({ from: r.oauthRefreshToken.sessionId, to: r.session.id }),
      user: r.one.user({ from: r.oauthRefreshToken.userId, to: r.user.id, optional: false }),
      accessTokens: r.many.oauthAccessToken(),
    },
    oauthConsent: {
      client: r.one.oauthClient({
        from: r.oauthConsent.clientId,
        to: r.oauthClient.clientId,
        optional: false,
      }),
      user: r.one.user({ from: r.oauthConsent.userId, to: r.user.id }),
    },
    mcpGrant: {
      user: r.one.user({ from: r.mcpGrant.userId, to: r.user.id, optional: false }),
      client: r.one.oauthClient({
        from: r.mcpGrant.clientId,
        to: r.oauthClient.clientId,
        optional: false,
      }),
      organizations: r.many.mcpGrantOrganization(),
      workspaces: r.many.mcpGrantWorkspace(),
    },
    mcpGrantOrganization: {
      grant: r.one.mcpGrant({
        from: r.mcpGrantOrganization.grantId,
        to: r.mcpGrant.id,
        optional: false,
      }),
      organization: r.one.organization({
        from: r.mcpGrantOrganization.organizationId,
        to: r.organization.id,
        optional: false,
      }),
    },
    mcpGrantWorkspace: {
      grant: r.one.mcpGrant({
        from: r.mcpGrantWorkspace.grantId,
        to: r.mcpGrant.id,
        optional: false,
      }),
      workspace: r.one.workspace({
        from: r.mcpGrantWorkspace.workspaceId,
        to: r.workspace.id,
        optional: false,
      }),
    },
    oauthResource: { clients: r.many.oauthClientResource() },
    oauthClientResource: {
      client: r.one.oauthClient({
        from: r.oauthClientResource.clientId,
        to: r.oauthClient.clientId,
        optional: false,
      }),
      resource: r.one.oauthResource({
        from: r.oauthClientResource.resourceId,
        to: r.oauthResource.identifier,
        optional: false,
      }),
    },
    passwordResetToken: {
      user: r.one.user({ from: r.passwordResetToken.userId, to: r.user.id, optional: false }),
    },
    user: {
      memberships: r.many.membership(),
      accounts: r.many.account(),
      sessions: r.many.session(),
      twoFactor: r.one.twoFactor({ from: r.user.id, to: r.twoFactor.userId }),
      oauthClients: r.many.oauthClient(),
      oauthAccessTokens: r.many.oauthAccessToken(),
      oauthRefreshTokens: r.many.oauthRefreshToken(),
      oauthConsents: r.many.oauthConsent(),
      mcpGrants: r.many.mcpGrant(),
      passwordResetToken: r.one.passwordResetToken({
        from: r.user.id,
        to: r.passwordResetToken.userId,
      }),
      invitesCreated: r.many.invite({ alias: 'inviteCreatedBy' }),
      invitesAccepted: r.many.invite({ alias: 'inviteAcceptedBy' }),
      surveys: r.many.survey({ alias: 'SurveyCreator' }),
      ownedSurveys: r.many.survey({ alias: 'SurveyOwner' }),
      charts: r.many.chart(),
      dashboards: r.many.dashboard(),
      creatives: r.many.creative(),
      feedbackSources: r.many.feedbackSource(),
      teamUsers: r.many.teamUser(),
      workflows: r.many.workflow(),
      workflowVersionsPublished: r.many.workflowVersion(),
      retentionPoliciesUpdated: r.many.retentionPolicy(),
      exemptionsCreated: r.many.retentionExemption({ alias: 'retentionExemptionCreatedBy' }),
      exemptionsRevoked: r.many.retentionExemption({ alias: 'retentionExemptionRevokedBy' }),
      retentionNotices: r.many.retentionNotice(),
      attachmentExports: r.many.attachmentExport(),
    },
    segment: {
      workspace: r.one.workspace({
        from: r.segment.workspaceId,
        to: r.workspace.id,
        optional: false,
      }),
      surveys: r.many.survey(),
    },
    language: {
      workspace: r.one.workspace({
        from: r.language.workspaceId,
        to: r.workspace.id,
        optional: false,
      }),
      surveyLanguages: r.many.surveyLanguage(),
      creativeLanguages: r.many.creativeLanguage(),
    },
    surveyLanguage: {
      language: r.one.language({
        from: r.surveyLanguage.languageId,
        to: r.language.id,
        optional: false,
      }),
      survey: r.one.survey({ from: r.surveyLanguage.surveyId, to: r.survey.id, optional: false }),
    },
    embeddedData: {
      survey: r.one.survey({
        from: [r.embeddedData.surveyId, r.embeddedData.workspaceId],
        to: [r.survey.id, r.survey.workspaceId],
      }),
      workspace: r.one.workspace({
        from: r.embeddedData.workspaceId,
        to: r.workspace.id,
        optional: false,
      }),
      surveyLinks: r.many.surveyEmbeddedData(),
    },
    surveyEmbeddedData: {
      survey: r.one.survey({
        from: [r.surveyEmbeddedData.surveyId, r.surveyEmbeddedData.workspaceId],
        to: [r.survey.id, r.survey.workspaceId],
        optional: false,
      }),
      embeddedData: r.one.embeddedData({
        from: [r.surveyEmbeddedData.embeddedDataId, r.surveyEmbeddedData.workspaceId],
        to: [r.embeddedData.id, r.embeddedData.workspaceId],
        optional: false,
      }),
    },
    team: {
      organization: r.one.organization({
        from: r.team.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      teamUsers: r.many.teamUser(),
      workspaceTeams: r.many.workspaceTeam(),
    },
    teamUser: {
      team: r.one.team({ from: r.teamUser.teamId, to: r.team.id, optional: false }),
      user: r.one.user({ from: r.teamUser.userId, to: r.user.id, optional: false }),
    },
    workspaceTeam: {
      workspace: r.one.workspace({
        from: r.workspaceTeam.workspaceId,
        to: r.workspace.id,
        optional: false,
      }),
      team: r.one.team({ from: r.workspaceTeam.teamId, to: r.team.id, optional: false }),
    },
    chart: {
      workspace: r.one.workspace({
        from: r.chart.workspaceId,
        to: r.workspace.id,
        optional: false,
      }),
      creator: r.one.user({ from: r.chart.createdBy, to: r.user.id }),
      feedbackDirectory: r.one.feedbackDirectory({
        from: r.chart.feedbackDirectoryId,
        to: r.feedbackDirectory.id,
        optional: false,
      }),
      widgets: r.many.dashboardWidget(),
    },
    embeddedFieldLabel: {
      workspace: r.one.workspace({
        from: r.embeddedFieldLabel.workspaceId,
        to: r.workspace.id,
        optional: false,
      }),
    },
    dashboard: {
      workspace: r.one.workspace({
        from: r.dashboard.workspaceId,
        to: r.workspace.id,
        optional: false,
      }),
      creator: r.one.user({ from: r.dashboard.createdBy, to: r.user.id }),
      widgets: r.many.dashboardWidget(),
    },
    dashboardWidget: {
      dashboard: r.one.dashboard({
        from: r.dashboardWidget.dashboardId,
        to: r.dashboard.id,
        optional: false,
      }),
      chart: r.one.chart({ from: r.dashboardWidget.chartId, to: r.chart.id, optional: false }),
    },
    attachmentExport: {
      survey: r.one.survey({
        from: [r.attachmentExport.surveyId, r.attachmentExport.workspaceId],
        to: [r.survey.id, r.survey.workspaceId],
        optional: false,
      }),
      workspace: r.one.workspace({
        from: r.attachmentExport.workspaceId,
        to: r.workspace.id,
        optional: false,
      }),
      requestedBy: r.one.user({ from: r.attachmentExport.requestedById, to: r.user.id }),
    },
    feedbackSource: {
      workspace: r.one.workspace({
        from: r.feedbackSource.workspaceId,
        to: r.workspace.id,
        optional: false,
      }),
      feedbackDirectory: r.one.feedbackDirectory({
        from: r.feedbackSource.feedbackDirectoryId,
        to: r.feedbackDirectory.id,
        optional: false,
      }),
      formbricksMappings: r.many.feedbackSourceFormbricksMapping(),
      fieldMappings: r.many.feedbackSourceFieldMapping(),
      creator: r.one.user({ from: r.feedbackSource.createdBy, to: r.user.id }),
      feedbackDirectoryWorkspace: r.one.feedbackDirectoryWorkspace({
        from: [r.feedbackSource.feedbackDirectoryId, r.feedbackSource.workspaceId],
        to: [
          r.feedbackDirectoryWorkspace.feedbackDirectoryId,
          r.feedbackDirectoryWorkspace.workspaceId,
        ],
        optional: false,
      }),
    },
    feedbackSourceFormbricksMapping: {
      feedbackSource: r.one.feedbackSource({
        from: [
          r.feedbackSourceFormbricksMapping.feedbackSourceId,
          r.feedbackSourceFormbricksMapping.workspaceId,
        ],
        to: [r.feedbackSource.id, r.feedbackSource.workspaceId],
        optional: false,
      }),
      survey: r.one.survey({
        from: [
          r.feedbackSourceFormbricksMapping.surveyId,
          r.feedbackSourceFormbricksMapping.workspaceId,
        ],
        to: [r.survey.id, r.survey.workspaceId],
        optional: false,
      }),
    },
    feedbackSourceFieldMapping: {
      feedbackSource: r.one.feedbackSource({
        from: [
          r.feedbackSourceFieldMapping.feedbackSourceId,
          r.feedbackSourceFieldMapping.workspaceId,
        ],
        to: [r.feedbackSource.id, r.feedbackSource.workspaceId],
        optional: false,
      }),
    },
    feedbackDirectory: {
      organization: r.one.organization({
        from: r.feedbackDirectory.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      workspaces: r.many.feedbackDirectoryWorkspace(),
      feedbackSources: r.many.feedbackSource(),
      charts: r.many.chart(),
    },
    feedbackDirectoryWorkspace: {
      feedbackDirectory: r.one.feedbackDirectory({
        from: r.feedbackDirectoryWorkspace.feedbackDirectoryId,
        to: r.feedbackDirectory.id,
        optional: false,
      }),
      workspace: r.one.workspace({
        from: r.feedbackDirectoryWorkspace.workspaceId,
        to: r.workspace.id,
        optional: false,
      }),
      feedbackSources: r.many.feedbackSource(),
    },
    retentionPolicy: {
      organization: r.one.organization({
        from: r.retentionPolicy.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      updatedBy: r.one.user({ from: r.retentionPolicy.updatedById, to: r.user.id }),
    },
    retentionExemption: {
      organization: r.one.organization({
        from: r.retentionExemption.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      survey: r.one.survey({ from: r.retentionExemption.surveyId, to: r.survey.id }),
      createdBy: r.one.user({
        from: r.retentionExemption.createdById,
        to: r.user.id,
        alias: 'retentionExemptionCreatedBy',
      }),
      revokedBy: r.one.user({
        from: r.retentionExemption.revokedById,
        to: r.user.id,
        alias: 'retentionExemptionRevokedBy',
      }),
    },
    retentionRun: {
      organization: r.one.organization({
        from: r.retentionRun.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      items: r.many.retentionRunItem(),
    },
    retentionRunItem: {
      run: r.one.retentionRun({
        from: r.retentionRunItem.runId,
        to: r.retentionRun.id,
        optional: false,
      }),
    },
    retentionNotice: {
      organization: r.one.organization({
        from: r.retentionNotice.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      survey: r.one.survey({ from: r.retentionNotice.surveyId, to: r.survey.id }),
      user: r.one.user({ from: r.retentionNotice.userId, to: r.user.id }),
    },
    retentionSurveyRestore: {
      survey: r.one.survey({
        from: r.retentionSurveyRestore.surveyId,
        to: r.survey.id,
        optional: false,
      }),
    },
    workflow: {
      workspace: r.one.workspace({
        from: r.workflow.workspaceId,
        to: r.workspace.id,
        optional: false,
      }),
      creator: r.one.user({ from: r.workflow.createdBy, to: r.user.id }),
      versions: r.many.workflowVersion(),
      runs: r.many.workflowRun(),
      secrets: r.many.workflowSecret(),
      samples: r.many.workflowSample(),
    },
    workflowVersion: {
      workflow: r.one.workflow({
        from: [r.workflowVersion.workflowId, r.workflowVersion.workspaceId],
        to: [r.workflow.id, r.workflow.workspaceId],
        optional: false,
      }),
      publisher: r.one.user({ from: r.workflowVersion.publishedBy, to: r.user.id }),
      runs: r.many.workflowRun(),
    },
    workflowRun: {
      workflow: r.one.workflow({
        from: [r.workflowRun.workflowId, r.workflowRun.workspaceId],
        to: [r.workflow.id, r.workflow.workspaceId],
        optional: false,
      }),
      workspace: r.one.workspace({
        from: r.workflowRun.workspaceId,
        to: r.workspace.id,
        optional: false,
      }),
      workflowVersion: r.one.workflowVersion({
        from: r.workflowRun.workflowVersionId,
        to: r.workflowVersion.id,
      }),
      response: r.one.response({ from: r.workflowRun.responseId, to: r.response.id }),
      logs: r.many.workflowRunLog(),
    },
    workflowSample: {
      workflow: r.one.workflow({
        from: [r.workflowSample.workflowId, r.workflowSample.workspaceId],
        to: [r.workflow.id, r.workflow.workspaceId],
        optional: false,
      }),
    },
    workflowRunLog: {
      run: r.one.workflowRun({
        from: r.workflowRunLog.runId,
        to: r.workflowRun.id,
        optional: false,
      }),
    },
    workflowSecret: {
      workflow: r.one.workflow({
        from: [r.workflowSecret.workflowId, r.workflowSecret.workspaceId],
        to: [r.workflow.id, r.workflow.workspaceId],
        optional: false,
      }),
    },
  }),
)
