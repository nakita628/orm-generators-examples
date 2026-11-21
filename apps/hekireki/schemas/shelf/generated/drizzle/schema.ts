import {
  type PgTableExtraConfigValue,
  boolean,
  customType,
  doublePrecision,
  foreignKey,
  index,
  integer,
  jsonb,
  pgEnum,
  pgTable,
  primaryKey,
  text,
  timestamp,
  uniqueIndex,
  varchar,
} from 'drizzle-orm/pg-core'
import { defineRelations, sql } from 'drizzle-orm'
import cuid from 'cuid'

const bytea = customType<{ data: Uint8Array }>({
  dataType: () => 'bytea',
})

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

export const assetStatusEnum = pgEnum('AssetStatus', ['AVAILABLE', 'IN_CUSTODY', 'CHECKED_OUT'])

/** * Determines how an asset is tracked throughout the system. Immutable after creation. */
export const assetTypeEnum = pgEnum('AssetType', ['INDIVIDUAL', 'QUANTITY_TRACKED'])

/** * Controls what happens after checkout for quantity-tracked assets. */
export const consumptionTypeEnum = pgEnum('ConsumptionType', ['ONE_WAY', 'TWO_WAY'])

/** * Classifies the type of quantity change in a ConsumptionLog entry. */
export const consumptionCategoryEnum = pgEnum('ConsumptionCategory', [
  'CHECKOUT',
  'RETURN',
  'RESTOCK',
  'ADJUSTMENT',
  'LOSS',
  'CONSUME',
  'DAMAGE',
  'MOVE',
  'INITIAL',
  'OPENING_BALANCE',
])

/**
 * The kind of PDF a PdfTemplate lays out. Only the booking PDF has saved
 * layouts today; the check-in receipt and the audit receipt can join as new
 * kinds.
 */
export const pdfTemplateKindEnum = pgEnum('PdfTemplateKind', ['BOOKING'])

export const assetIndexModeEnum = pgEnum('AssetIndexMode', ['SIMPLE', 'ADVANCED'])

export const tagUseForEnum = pgEnum('TagUseFor', ['ASSET', 'BOOKING'])

export const noteTypeEnum = pgEnum('NoteType', ['COMMENT', 'UPDATE'])

export const barcodeTypeEnum = pgEnum('BarcodeType', [
  'Code128',
  'Code39',
  'DataMatrix',
  'ExternalQR',
  'EAN13',
])

export const errorCorrectionEnum = pgEnum('ErrorCorrection', ['L', 'M', 'Q', 'H'])

/** The kind of code a `Scan` recorded. */
export const scanCodeTypeEnum = pgEnum('ScanCodeType', ['QR', 'BARCODE', 'SAM_ID'])

/** The surface that recorded a `Scan`. */
export const scanSourceEnum = pgEnum('ScanSource', [
  'QR_LINK',
  'WEB_SCANNER',
  'WEB_DRAWER',
  'AUDIT',
  'COMPANION',
  'GPS_UPDATE',
])

export const rolesEnum = pgEnum('Roles', ['USER', 'ADMIN'])

export const organizationTypeEnum = pgEnum('OrganizationType', ['PERSONAL', 'TEAM'])

export const qrIdDisplayPreferenceEnum = pgEnum('QrIdDisplayPreference', [
  'QR_ID',
  'SAM_ID',
  'Code128',
  'Code39',
  'DataMatrix',
  'ExternalQR',
  'EAN13',
])

export const dateFormatPreferenceEnum = pgEnum('DateFormatPreference', [
  'DD_MM_YYYY',
  'MM_DD_YYYY',
  'YYYY_MM_DD',
  'MMM_DD_YYYY',
  'DD_MMM_YYYY',
])

export const timeFormatPreferenceEnum = pgEnum('TimeFormatPreference', ['H12', 'H24'])

export const weekStartPreferenceEnum = pgEnum('WeekStartPreference', [
  'MONDAY',
  'SUNDAY',
  'SATURDAY',
])

export const organizationRolesEnum = pgEnum('OrganizationRoles', [
  'ADMIN',
  'BASE',
  'OWNER',
  'SELF_SERVICE',
  'MANAGER',
])

export const tierIdEnum = pgEnum('TierId', ['free', 'tier_1', 'tier_2', 'custom'])

export const customFieldTypeEnum = pgEnum('CustomFieldType', [
  'TEXT',
  'OPTION',
  'BOOLEAN',
  'DATE',
  'MULTILINE_TEXT',
  'AMOUNT',
  'NUMBER',
])

/**
 * The format of a TEXT custom field. Unlike the type, it can change after the
 * field exists.
 */
export const customFieldFormatEnum = pgEnum('CustomFieldFormat', ['PLAIN', 'LINK'])

export const currencyEnum = pgEnum('Currency', [
  'AED',
  'AFN',
  'ALL',
  'AMD',
  'ANG',
  'AOA',
  'ARS',
  'AUD',
  'AWG',
  'AZN',
  'BAM',
  'BBD',
  'BDT',
  'BGN',
  'BHD',
  'BIF',
  'BMD',
  'BND',
  'BOB',
  'BRL',
  'BSD',
  'BTN',
  'BWP',
  'BYN',
  'BZD',
  'CAD',
  'CDF',
  'CHF',
  'CLP',
  'CNY',
  'COP',
  'CRC',
  'CUP',
  'CVE',
  'CZK',
  'DJF',
  'DKK',
  'DOP',
  'DZD',
  'EGP',
  'ERN',
  'ETB',
  'EUR',
  'FJD',
  'FKP',
  'GBP',
  'GEL',
  'GHS',
  'GIP',
  'GMD',
  'GNF',
  'GTQ',
  'GYD',
  'HKD',
  'HNL',
  'HTG',
  'HUF',
  'IDR',
  'ILS',
  'INR',
  'IQD',
  'IRR',
  'ISK',
  'JMD',
  'JOD',
  'JPY',
  'KES',
  'KGS',
  'KHR',
  'KMF',
  'KPW',
  'KRW',
  'KWD',
  'KYD',
  'KZT',
  'LAK',
  'LBP',
  'LKR',
  'LRD',
  'LSL',
  'LYD',
  'MAD',
  'MDL',
  'MGA',
  'MKD',
  'MMK',
  'MNT',
  'MOP',
  'MRU',
  'MUR',
  'MVR',
  'MWK',
  'MXN',
  'MYR',
  'MZN',
  'NAD',
  'NGN',
  'NIO',
  'NOK',
  'NPR',
  'NZD',
  'OMR',
  'PAB',
  'PEN',
  'PGK',
  'PHP',
  'PKR',
  'PLN',
  'PYG',
  'QAR',
  'RON',
  'RSD',
  'RUB',
  'RWF',
  'SAR',
  'SBD',
  'SCR',
  'SDG',
  'SEK',
  'SGD',
  'SHP',
  'SLE',
  'SOS',
  'SRD',
  'SSP',
  'STN',
  'SVC',
  'SYP',
  'SZL',
  'THB',
  'TJS',
  'TMT',
  'TND',
  'TOP',
  'TRY',
  'TTD',
  'TWD',
  'TZS',
  'UAH',
  'UGX',
  'USD',
  'UYU',
  'UZS',
  'VES',
  'VND',
  'VUV',
  'WST',
  'XAF',
  'XCD',
  'XOF',
  'XPF',
  'YER',
  'ZAR',
  'ZMW',
  'ZWL',
])

export const inviteStatusesEnum = pgEnum('InviteStatuses', [
  'PENDING',
  'ACCEPTED',
  'REJECTED',
  'INVALIDATED',
])

export const bookingStatusEnum = pgEnum('BookingStatus', [
  'DRAFT',
  'RESERVED',
  'ONGOING',
  'OVERDUE',
  'COMPLETE',
  'ARCHIVED',
  'CANCELLED',
])

/** The operating system a companion-app push token was issued for. */
export const devicePlatformEnum = pgEnum('DevicePlatform', ['IOS', 'ANDROID'])

export const kitStatusEnum = pgEnum('KitStatus', ['AVAILABLE', 'IN_CUSTODY', 'CHECKED_OUT'])

export const reminderRecurrenceUnitEnum = pgEnum('ReminderRecurrenceUnit', [
  'DAY',
  'WEEK',
  'MONTH',
  'YEAR',
])

export const updateStatusEnum = pgEnum('UpdateStatus', ['DRAFT', 'PUBLISHED'])

export const auditStatusEnum = pgEnum('AuditStatus', [
  'PENDING',
  'ACTIVE',
  'COMPLETED',
  'CANCELLED',
  'ARCHIVED',
])

export const auditAssetStatusEnum = pgEnum('AuditAssetStatus', [
  'PENDING',
  'FOUND',
  'MISSING',
  'UNEXPECTED',
])

export const auditAssignmentRoleEnum = pgEnum('AuditAssignmentRole', ['LEAD', 'PARTICIPANT'])

/** Lifecycle of a background asset import. */
export const assetImportStatusEnum = pgEnum('AssetImportStatus', [
  'DRAFT',
  'QUEUED',
  'RUNNING',
  'COMPLETED',
  'FAILED',
])

/** Which step of the background import runs next (or failed). */
export const assetImportPhaseEnum = pgEnum('AssetImportPhase', ['PREPARE', 'ROWS', 'FINALIZE'])

/** How the last fetch of a link's page went. */
export const linkPreviewOutcomeEnum = pgEnum('LinkPreviewOutcome', [
  'TITLE',
  'NO_TITLE',
  'SIGN_IN',
  'BLOCKED',
  'UNREACHABLE',
])

/**
 * How an audit photo reached Shelf, as reported by the client. The phone app
 * reports CAMERA only for its own camera, which cannot hand back an old
 * picture, so it is the strongest evidence short of a modified client; a
 * library photo's date is only what its file claims; the web app cannot tell
 * a camera shot from a file; PHONE is a phone app version that does not say
 * which.
 */
export const auditImageCaptureSourceEnum = pgEnum('AuditImageCaptureSource', [
  'CAMERA',
  'LIBRARY',
  'WEB_UPLOAD',
  'PHONE',
])

/** How a membership's role came to change. */
export const roleChangeSourceEnum = pgEnum('RoleChangeSource', ['MANUAL', 'SSO'])

export const activityEntityEnum = pgEnum('ActivityEntity', [
  'ASSET',
  'BOOKING',
  'AUDIT',
  'KIT',
  'LOCATION',
  'TEAM_MEMBER',
  'CUSTODY',
  'USER',
  'ORGANIZATION',
])

export const activityActionEnum = pgEnum('ActivityAction', [
  'ASSET_CREATED',
  'ASSET_NAME_CHANGED',
  'ASSET_DESCRIPTION_CHANGED',
  'ASSET_CATEGORY_CHANGED',
  'ASSET_MODEL_CHANGED',
  'ASSET_KIT_CHANGED',
  'ASSET_LOCATION_CHANGED',
  'ASSET_TAGS_CHANGED',
  'ASSET_STATUS_CHANGED',
  'ASSET_VALUATION_CHANGED',
  'ASSET_QUANTITY_CHANGED',
  'ASSET_MIN_QUANTITY_CHANGED',
  'ASSET_CUSTOM_FIELD_CHANGED',
  'ASSET_PREFERRED_BARCODE_CHANGED',
  'ASSET_AVAILABLE_TO_BOOK_CHANGED',
  'ASSET_ARCHIVED',
  'ASSET_UNARCHIVED',
  'ASSET_DELETED',
  'CUSTODY_ASSIGNED',
  'CUSTODY_RELEASED',
  'BOOKING_CREATED',
  'BOOKING_STATUS_CHANGED',
  'BOOKING_DATES_CHANGED',
  'BOOKING_ASSETS_ADDED',
  'BOOKING_ASSETS_REMOVED',
  'BOOKING_CHECKED_OUT',
  'BOOKING_CHECKED_IN',
  'BOOKING_PARTIAL_CHECKIN',
  'BOOKING_PARTIAL_CHECKOUT',
  'BOOKING_CANCELLED',
  'BOOKING_ARCHIVED',
  'BOOKING_MODEL_REQUESTED',
  'BOOKING_MODEL_REQUEST_CHANGED',
  'BOOKING_MODEL_REQUEST_REMOVED',
  'BOOKING_MODEL_REQUEST_FULFILLED',
  'AUDIT_CREATED',
  'AUDIT_STARTED',
  'AUDIT_ASSETS_ADDED',
  'AUDIT_ASSETS_REMOVED',
  'AUDIT_ASSET_SCANNED',
  'AUDIT_ASSET_SCAN_REMOVED',
  'AUDIT_DUE_DATE_CHANGED',
  'AUDIT_ASSIGNEE_ADDED',
  'AUDIT_ASSIGNEE_REMOVED',
  'AUDIT_UPDATED',
  'AUDIT_COMPLETED',
  'AUDIT_CANCELLED',
  'AUDIT_ARCHIVED',
  'LOCATION_CREATED',
  'LOCATION_UPDATED',
  'KIT_CREATED',
  'KIT_UPDATED',
  'ORGANIZATION_QR_ID_DISPLAY_PREFERENCE_CHANGED',
])

export const image = pgTable(
  'Image',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    contentType: text('contentType').notNull(),
    altText: text('altText'),
    blob: bytea('blob').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    ownerOrgId: text('ownerOrgId').notNull(),
    userId: text('userId').notNull(),
  },
  (table): PgTableExtraConfigValue[] => [
    index('Image_ownerOrgId_idx').on(table.ownerOrgId),
    index('Image_userId_idx').on(table.userId),
    foreignKey({
      name: 'Image_ownerOrgId_fkey',
      columns: [table.ownerOrgId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({ name: 'Image_userId_fkey', columns: [table.userId], foreignColumns: [user.id] })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

export const user = pgTable(
  'User',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    email: text('email').notNull(),
    username: text('username')
      .notNull()
      .$defaultFn(() => cuid()),
    firstName: text('firstName'),
    lastName: text('lastName'),
    displayName: text('displayName'),
    profilePicture: text('profilePicture'),
    usedFreeTrial: boolean('usedFreeTrial').notNull().default(false),
    onboarded: boolean('onboarded').notNull().default(false),
    customerId: text('customerId'),
    sso: boolean('sso').notNull().default(false),
    createdWithInvite: boolean('createdWithInvite').notNull().default(false),
    skipSubscriptionCheck: boolean('skipSubscriptionCheck').notNull().default(false),
    hasUnpaidInvoice: boolean('hasUnpaidInvoice').notNull().default(false),
    warnForNoPaymentMethod: boolean('warnForNoPaymentMethod').notNull().default(false),
    lastSelectedOrganizationId: text('lastSelectedOrganizationId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    deletedAt: timestamp('deletedAt', { precision: 3 }),
    lastMobileActiveAt: timestamp('lastMobileActiveAt', { precision: 3 }),
    dateFormat: dateFormatPreferenceEnum('dateFormat'),
    timeFormat: timeFormatPreferenceEnum('timeFormat'),
    weekStart: weekStartPreferenceEnum('weekStart'),
    timeZone: text('timeZone'),
    tierId: tierIdEnum('tierId').notNull().default('free'),
    referralSource: text('referralSource'),
  },
  (table): PgTableExtraConfigValue[] => [
    uniqueIndex('User_email_key').on(table.email),
    uniqueIndex('User_username_key').on(table.username),
    uniqueIndex('User_customerId_key').on(table.customerId),
    index('User_firstName_lastName_idx').on(table.firstName, table.lastName),
    index('User_tierId_idx').on(table.tierId),
    index('User_lastSelectedOrganizationId_idx').on(table.lastSelectedOrganizationId),
    index('User_lastMobileActiveAt_idx').on(table.lastMobileActiveAt),
    uniqueIndex('User_email_username_key').on(table.email, table.username),
    foreignKey({
      name: 'User_lastSelectedOrganizationId_fkey',
      columns: [table.lastSelectedOrganizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({ name: 'User_tierId_fkey', columns: [table.tierId], foreignColumns: [tier.id] })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

export const userContact = pgTable(
  'UserContact',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    phone: text('phone'),
    street: text('street'),
    city: text('city'),
    stateProvince: text('stateProvince'),
    zipPostalCode: text('zipPostalCode'),
    countryRegion: text('countryRegion'),
    userId: text('userId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('UserContact_userId_key').on(table.userId),
    index('UserContact_userId_idx').on(table.userId),
    index('UserContact_phone_idx').on(table.phone),
    index('UserContact_city_stateProvince_idx').on(table.city, table.stateProvince),
    index('UserContact_countryRegion_idx').on(table.countryRegion),
    index('UserContact_zipPostalCode_idx').on(table.zipPostalCode),
    index('UserContact_city_countryRegion_idx').on(table.city, table.countryRegion),
    foreignKey({
      name: 'UserContact_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const userBusinessIntel = pgTable(
  'UserBusinessIntel',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    howDidYouHearAboutUs: text('howDidYouHearAboutUs'),
    jobTitle: text('jobTitle'),
    teamSize: text('teamSize'),
    companyName: text('companyName'),
    primaryUseCase: text('primaryUseCase'),
    currentSolution: text('currentSolution'),
    timeline: text('timeline'),
    signupPlan: text('signupPlan'),
    signupTrial: boolean('signupTrial'),
    utmSource: text('utmSource'),
    utmMedium: text('utmMedium'),
    utmCampaign: text('utmCampaign'),
    utmContent: text('utmContent'),
    userId: text('userId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('UserBusinessIntel_userId_key').on(table.userId),
    index('UserBusinessIntel_userId_idx').on(table.userId),
    index('UserBusinessIntel_companyName_idx').on(table.companyName),
    index('UserBusinessIntel_jobTitle_idx').on(table.jobTitle),
    index('UserBusinessIntel_teamSize_idx').on(table.teamSize),
    index('UserBusinessIntel_utmSource_idx').on(table.utmSource),
    index('UserBusinessIntel_utmCampaign_idx').on(table.utmCampaign),
    foreignKey({
      name: 'UserBusinessIntel_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const asset = pgTable(
  'Asset',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    title: text('title').notNull(),
    description: text('description'),
    mainImage: text('mainImage'),
    thumbnailImage: text('thumbnailImage'),
    mainImageExpiration: timestamp('mainImageExpiration', { precision: 3 }),
    status: assetStatusEnum('status').notNull().default('AVAILABLE'),
    valuation: doublePrecision('value'),
    availableToBook: boolean('availableToBook').notNull().default(true),
    sequentialId: text('sequentialId'),
    archivedAt: timestamp('archivedAt', { withTimezone: true, precision: 3 }),
    type: assetTypeEnum('type').notNull().default('INDIVIDUAL'),
    assetModelId: text('assetModelId'),
    quantity: integer('quantity'),
    minQuantity: integer('minQuantity'),
    consumptionType: consumptionTypeEnum('consumptionType'),
    unitOfMeasure: text('unitOfMeasure'),
    lowStockNotifiedAt: timestamp('lowStockNotifiedAt', { precision: 3 }),
    stockLedgerStartedAt: timestamp('stockLedgerStartedAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    userId: text('userId').notNull(),
    organizationId: text('organizationId').notNull(),
    categoryId: text('categoryId'),
    preferredBarcodeId: text('preferredBarcodeId'),
  },
  (table): PgTableExtraConfigValue[] => [
    uniqueIndex('Asset_preferredBarcodeId_key').on(table.preferredBarcodeId),
    index('Asset_title_description_idx').on(table.title, table.description),
    index('Asset_sequentialId_trgm_idx').on(table.sequentialId),
    index('Asset_organizationId_compound_idx').on(
      table.organizationId,
      table.title,
      table.status,
      table.availableToBook,
    ),
    index('Asset_status_organizationId_idx').on(table.status, table.organizationId),
    index('Asset_organizationId_archivedAt_idx').on(table.organizationId, table.archivedAt),
    index('Asset_createdAt_organizationId_idx').on(table.createdAt, table.organizationId),
    index('Asset_organizationId_createdAt_id_idx').on(
      table.organizationId,
      table.createdAt,
      table.id,
    ),
    index('Asset_valuation_organizationId_idx').on(table.valuation, table.organizationId),
    index('Asset_categoryId_organizationId_idx').on(table.categoryId, table.organizationId),
    index('Asset_assetModelId_organizationId_idx').on(table.assetModelId, table.organizationId),
    index('Asset_sequentialId_idx').on(table.sequentialId),
    index('Asset_userId_idx').on(table.userId),
    uniqueIndex('Asset_organizationId_sequentialId_key').on(
      table.organizationId,
      table.sequentialId,
    ),
    foreignKey({
      name: 'Asset_assetModelId_fkey',
      columns: [table.assetModelId],
      foreignColumns: [assetModel.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({ name: 'Asset_userId_fkey', columns: [table.userId], foreignColumns: [user.id] })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Asset_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Asset_categoryId_fkey',
      columns: [table.categoryId],
      foreignColumns: [category.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Asset_preferredBarcodeId_fkey',
      columns: [table.preferredBarcodeId],
      foreignColumns: [barcode.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const assetFilterPreset = pgTable(
  'AssetFilterPreset',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    organizationId: text('organizationId').notNull(),
    ownerId: text('ownerId').notNull(),
    name: text('name').notNull(),
    query: text('query').notNull(),
    starred: boolean('starred').notNull().default(false),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('asset_filter_presets_owner_lookup_idx').on(table.organizationId, table.ownerId),
    uniqueIndex('asset_filter_presets_owner_name_unique').on(
      table.organizationId,
      table.ownerId,
      table.name,
    ),
    foreignKey({
      name: 'AssetFilterPreset_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AssetFilterPreset_ownerId_fkey',
      columns: [table.ownerId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const assetIndexSettings = pgTable(
  'AssetIndexSettings',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    userId: text('userId').notNull(),
    organizationId: text('organizationId').notNull(),
    mode: assetIndexModeEnum('mode').notNull().default('SIMPLE'),
    columns: jsonb('columns')
      .notNull()
      .default([
        { name: 'id', visible: true, position: 0 },
        { name: 'status', visible: true, position: 1 },
        { name: 'description', visible: true, position: 2 },
        { name: 'valuation', visible: true, position: 3 },
        { name: 'createdAt', visible: true, position: 4 },
        { name: 'category', visible: true, position: 5 },
        { name: 'tags', visible: true, position: 6 },
        { name: 'location', visible: true, position: 7 },
        { name: 'kit', visible: true, position: 8 },
        { name: 'custody', visible: true, position: 9 },
      ]),
    freezeColumn: boolean('freezeColumn').notNull().default(true),
    showAssetImage: boolean('showAssetImage').notNull().default(true),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('AssetIndexSettings_organizationId_idx').on(table.organizationId),
    uniqueIndex('AssetIndexSettings_userId_organizationId_key').on(
      table.userId,
      table.organizationId,
    ),
    foreignKey({
      name: 'AssetIndexSettings_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AssetIndexSettings_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * A saved layout for a generated PDF, kept per team workspace under a name:
 * which columns, rows and header lines print, and in what order the rows
 * come. A row is a layout, not a generated file; the person generating a PDF
 * picks a template by name.
 */
export const pdfTemplate = pgTable(
  'PdfTemplate',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    name: text('name').notNull(),
    kind: pdfTemplateKindEnum('kind').notNull().default('BOOKING'),
    /**
     * The layout's switches. Read through `parseBookingTemplateConfig`,
     * which fills any key a stored row lacks from the Checklist template's
     * values.
     */
    config: jsonb('config').notNull().default({}),
    /**
     * The template the PDF dialog offers first. One per workspace and kind,
     * enforced by a partial unique index in raw SQL (migration
     * 20261103120000_add_pdf_templates). Names are also unique per workspace
     * and kind ignoring case, through an expression index in the same
     * migration.
     */
    isDefault: boolean('isDefault').notNull().default(false),
    organizationId: text('organizationId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('PdfTemplate_organizationId_kind_name_key').on(
      table.organizationId,
      table.kind,
      table.name,
    ),
    foreignKey({
      name: 'PdfTemplate_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const category = pgTable(
  'Category',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    name: text('name').notNull(),
    description: text('description'),
    color: text('color').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    userId: text('userId').notNull(),
    organizationId: text('organizationId').notNull(),
  },
  (table) => [
    index('Category_organizationId_idx').on(table.organizationId),
    index('Category_userId_idx').on(table.userId),
    index('Category_name_trgm_idx').on(table.name),
    foreignKey({ name: 'Category_userId_fkey', columns: [table.userId], foreignColumns: [user.id] })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Category_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * *
 * * A template/grouping entity for assets.
 * * Groups individual assets of the same type (e.g., "Dell Latitude 5550")
 * * and provides default field values when creating new assets from the model.
 */
export const assetModel = pgTable(
  'AssetModel',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    name: text('name').notNull(),
    description: text('description'),
    /**
     * Public URL (bucket `files`) of the model's cover image. Rendered by any
     * asset of this model that has no image of its own. Public — and therefore
     * never-expiring and cacheable — like Location.imageUrl; see
     * resolveAssetImage() for the cascade.
     */
    image: text('image'),
    /**
     * Public URL of the 108px thumbnail generated alongside `image`. Used at
     * list sizes so inheriting assets never render the full-size original.
     */
    thumbnailImage: text('thumbnailImage'),
    defaultCategoryId: text('defaultCategoryId'),
    defaultValuation: doublePrecision('defaultValuation'),
    organizationId: text('organizationId').notNull(),
    userId: text('userId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('AssetModel_organizationId_name_idx').on(table.organizationId, table.name),
    index('AssetModel_userId_idx').on(table.userId),
    index('AssetModel_name_trgm_idx').on(table.name),
    foreignKey({
      name: 'AssetModel_defaultCategoryId_fkey',
      columns: [table.defaultCategoryId],
      foreignColumns: [category.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AssetModel_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AssetModel_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * Book-by-Model: intent row attached to a booking for a quantity of
 * assets of a given AssetModel, without yet picking specific assets.
 * Concrete `BookingAsset` rows are only created at
 * scan-to-assign time; when the last unit is assigned, the request
 * row is deleted. `quantity` is the number of units still needing
 * assignment (starts at the full requested total, decrements on scan).
 *
 * Keeping model requests separate from `BookingAsset` means the ~40
 * code paths that assume `BookingAsset.assetId` points to a real
 * asset stay untouched — model-only bookings become concrete the
 * moment they're checked out.
 */
export const bookingModelRequest = pgTable(
  'BookingModelRequest',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    bookingId: text('bookingId').notNull(),
    assetModelId: text('assetModelId').notNull(),
    /**
     * Total units originally reserved. Mutated only by explicit
     * operator edits (upsert) — never decremented on materialisation.
     * This preserves the audit trail of "how many were planned".
     */
    quantity: integer('quantity').notNull(),
    /**
     * Units already assigned to concrete `BookingAsset` rows via
     * `materializeModelRequestForAsset`. Scans increment it.
     *
     * It DOES decrease: `removeAssets` (booking service) decrements it by the
     * number of removed `BookingAsset` rows whose `bookingModelRequestId`
     * points at this request, floored at 0, and clears `fulfilledAt` when the
     * result drops below `quantity`.
     *
     * Counted from that column, never from a shared `assetModelId`: an asset
     * added directly alongside a reservation shares the model without having
     * discharged anything, so grouping by model re-opened reservations that
     * were still correctly satisfied.
     */
    fulfilledQuantity: integer('fulfilledQuantity').notNull().default(0),
    /**
     * Timestamp of the moment the reservation became fully fulfilled — set
     * when `fulfilledQuantity` catches up to `quantity`, and preserved (not
     * re-stamped) if an unchanged quantity is saved again. Cleared back to
     * null when outstanding units reappear: the operator edits `quantity`
     * upwards, or `removeAssets` drops `fulfilledQuantity` below it.
     *
     * The invariant that makes this the primary "is this outstanding?" filter:
     * non-null IFF `fulfilledQuantity === quantity`, so `fulfilledAt IS NULL`
     * means units remain to materialise. `fulfilledQuantity` can never exceed
     * `quantity` (the claim in `materializeModelRequestForAsset` is a single
     * statement guarded by `WHERE "fulfilledQuantity" < "quantity"`; upsert
     * rejects a shrink below it), so any decrement re-opens the row.
     */
    fulfilledAt: timestamp('fulfilledAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('BookingModelRequest_assetModelId_idx').on(table.assetModelId),
    uniqueIndex('BookingModelRequest_bookingId_assetModelId_key').on(
      table.bookingId,
      table.assetModelId,
    ),
    foreignKey({
      name: 'BookingModelRequest_bookingId_fkey',
      columns: [table.bookingId],
      foreignColumns: [booking.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'BookingModelRequest_assetModelId_fkey',
      columns: [table.assetModelId],
      foreignColumns: [assetModel.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const tag = pgTable(
  'Tag',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    name: text('name').notNull(),
    description: text('description'),
    color: text('color'),
    useFor: tagUseForEnum('useFor').array().default(['ASSET']),
    userId: text('userId').notNull(),
    organizationId: text('organizationId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('Tag_organizationId_idx').on(table.organizationId),
    index('Tag_userId_idx').on(table.userId),
    index('Tag_name_trgm_idx').on(table.name),
    foreignKey({ name: 'Tag_userId_fkey', columns: [table.userId], foreignColumns: [user.id] })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Tag_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const note = pgTable(
  'Note',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    content: text('content').notNull(),
    type: noteTypeEnum('type').notNull().default('COMMENT'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    userId: text('userId'),
    assetId: text('assetId').notNull(),
  },
  (table) => [
    index('Note_assetId_idx').on(table.assetId),
    index('Note_userId_idx').on(table.userId),
    foreignKey({ name: 'Note_userId_fkey', columns: [table.userId], foreignColumns: [user.id] })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({ name: 'Note_assetId_fkey', columns: [table.assetId], foreignColumns: [asset.id] })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const bookingNote = pgTable(
  'BookingNote',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    content: text('content').notNull(),
    type: noteTypeEnum('type').notNull().default('COMMENT'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    userId: text('userId'),
    bookingId: text('bookingId').notNull(),
  },
  (table) => [
    index('BookingNote_bookingId_idx').on(table.bookingId),
    index('BookingNote_userId_idx').on(table.userId),
    foreignKey({
      name: 'BookingNote_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'BookingNote_bookingId_fkey',
      columns: [table.bookingId],
      foreignColumns: [booking.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const locationNote = pgTable(
  'LocationNote',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    content: text('content').notNull(),
    type: noteTypeEnum('type').notNull().default('COMMENT'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    userId: text('userId'),
    locationId: text('locationId').notNull(),
  },
  (table) => [
    index('LocationNote_locationId_idx').on(table.locationId),
    index('LocationNote_userId_idx').on(table.userId),
    foreignKey({
      name: 'LocationNote_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'LocationNote_locationId_fkey',
      columns: [table.locationId],
      foreignColumns: [location.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const qr = pgTable(
  'Qr',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    version: integer('version').notNull().default(0),
    errorCorrection: errorCorrectionEnum('errorCorrection').notNull().default('L'),
    assetId: text('assetId'),
    kitId: text('kitId'),
    userId: text('userId'),
    organizationId: text('organizationId'),
    batchId: text('batchId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('Qr_assetId_idx').on(table.assetId),
    index('Qr_kitId_idx').on(table.kitId),
    index('Qr_userId_idx').on(table.userId),
    index('Qr_organizationId_idx').on(table.organizationId),
    index('Qr_batchId_idx').on(table.batchId),
    index('Qr_id_trgm_idx').on(table.id),
    foreignKey({ name: 'Qr_assetId_fkey', columns: [table.assetId], foreignColumns: [asset.id] })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({ name: 'Qr_kitId_fkey', columns: [table.kitId], foreignColumns: [kit.id] })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({ name: 'Qr_userId_fkey', columns: [table.userId], foreignColumns: [user.id] })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Qr_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Qr_batchId_fkey',
      columns: [table.batchId],
      foreignColumns: [printBatch.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const barcode = pgTable(
  'Barcode',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    value: text('value').notNull(),
    type: barcodeTypeEnum('type').notNull().default('Code128'),
    assetId: text('assetId'),
    kitId: text('kitId'),
    organizationId: text('organizationId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table): PgTableExtraConfigValue[] => [
    index('Barcode_organizationId_value_idx').on(table.organizationId, table.value),
    index('Barcode_assetId_idx').on(table.assetId),
    index('Barcode_kitId_idx').on(table.kitId),
    index('Barcode_organizationId_idx').on(table.organizationId),
    index('Barcode_value_trgm_idx').on(table.value),
    uniqueIndex('Barcode_organizationId_value_key').on(table.organizationId, table.value),
    foreignKey({
      name: 'Barcode_assetId_fkey',
      columns: [table.assetId],
      foreignColumns: [asset.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({ name: 'Barcode_kitId_fkey', columns: [table.kitId], foreignColumns: [kit.id] })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Barcode_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const printBatch = pgTable(
  'PrintBatch',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    name: text('name').notNull(),
    printed: boolean('printed').notNull().default(false),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [uniqueIndex('PrintBatch_name_key').on(table.name)],
)

export const reportFound = pgTable(
  'ReportFound',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    email: text('email').notNull(),
    content: text('content').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    assetId: text('assetId'),
    kitId: text('kitId'),
  },
  (table) => [
    index('ReportFound_assetId_idx').on(table.assetId),
    index('ReportFound_kitId_idx').on(table.kitId),
    foreignKey({
      name: 'ReportFound_assetId_fkey',
      columns: [table.assetId],
      foreignColumns: [asset.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({ name: 'ReportFound_kitId_fkey', columns: [table.kitId], foreignColumns: [kit.id] })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const scan = pgTable(
  'Scan',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    latitude: text('latitude'),
    longitude: text('longitude'),
    userAgent: text('userAgent'),
    userId: text('userId'),
    qrId: text('qrId'),
    rawQrId: text('rawQrId'),
    codeType: scanCodeTypeEnum('codeType').notNull().default('QR'),
    code: text('code'),
    barcodeId: text('barcodeId'),
    assetId: text('assetId'),
    kitId: text('kitId'),
    organizationId: text('organizationId').notNull(),
    source: scanSourceEnum('source'),
    manuallyGenerated: boolean('manuallyGenerated').notNull().default(false),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('Scan_qrId_idx').on(table.qrId),
    index('Scan_userId_idx').on(table.userId),
    index('Scan_rawQrId_createdAt_idx').on(table.rawQrId, table.createdAt),
    index('Scan_assetId_createdAt_idx').on(table.assetId, table.createdAt),
    index('Scan_kitId_createdAt_idx').on(table.kitId, table.createdAt),
    index('Scan_barcodeId_idx').on(table.barcodeId),
    index('Scan_organizationId_idx').on(table.organizationId),
    foreignKey({ name: 'Scan_userId_fkey', columns: [table.userId], foreignColumns: [user.id] })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({ name: 'Scan_qrId_fkey', columns: [table.qrId], foreignColumns: [qr.id] })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Scan_barcodeId_fkey',
      columns: [table.barcodeId],
      foreignColumns: [barcode.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({ name: 'Scan_assetId_fkey', columns: [table.assetId], foreignColumns: [asset.id] })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({ name: 'Scan_kitId_fkey', columns: [table.kitId], foreignColumns: [kit.id] })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Scan_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const location = pgTable(
  'Location',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    name: text('name').notNull(),
    description: text('description'),
    address: text('address'),
    latitude: doublePrecision('latitude'),
    longitude: doublePrecision('longitude'),
    imageUrl: text('imageUrl'),
    thumbnailUrl: text('thumbnailUrl'),
    imageId: text('imageId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    userId: text('userId').notNull(),
    organizationId: text('organizationId').notNull(),
    parentId: text('parentId'),
  },
  (table) => [
    uniqueIndex('Location_imageId_key').on(table.imageId),
    index('Location_organizationId_idx').on(table.organizationId),
    index('Location_userId_idx').on(table.userId),
    index('Location_organizationId_parentId_idx').on(table.organizationId, table.parentId),
    index('Location_name_trgm_idx').on(table.name),
    index('Location_description_trgm_idx').on(table.description),
    index('Location_address_trgm_idx').on(table.address),
    foreignKey({
      name: 'Location_imageId_fkey',
      columns: [table.imageId],
      foreignColumns: [image.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({ name: 'Location_userId_fkey', columns: [table.userId], foreignColumns: [user.id] })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Location_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Location_parentId_fkey',
      columns: [table.parentId],
      foreignColumns: [table.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const role = pgTable(
  'Role',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    name: rolesEnum('name').notNull().default('USER'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [uniqueIndex('Role_name_key').on(table.name)],
)

export const teamMember = pgTable(
  'TeamMember',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    name: text('name').notNull(),
    organizationId: text('organizationId').notNull(),
    userId: text('userId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    deletedAt: timestamp('deletedAt', { precision: 3 }),
  },
  (table) => [
    index('TeamMember_name_idx').on(table.name),
    index('TeamMember_organizationId_idx').on(table.organizationId),
    index('TeamMember_userId_idx').on(table.userId),
    foreignKey({
      name: 'TeamMember_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'TeamMember_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const custody = pgTable(
  'Custody',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    teamMemberId: text('teamMemberId').notNull(),
    assetId: text('assetId').notNull(),
    kitCustodyId: text('kitCustodyId'),
    quantity: integer('quantity').notNull().default(1),
    locationId: text('locationId'),
    sourceUnknown: boolean('sourceUnknown').notNull().default(false),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('Custody_assetId_teamMemberId_idx').on(table.assetId, table.teamMemberId),
    index('Custody_teamMemberId_idx').on(table.teamMemberId),
    index('Custody_kitCustodyId_idx').on(table.kitCustodyId),
    index('Custody_locationId_idx').on(table.locationId),
    foreignKey({
      name: 'Custody_teamMemberId_fkey',
      columns: [table.teamMemberId],
      foreignColumns: [teamMember.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Custody_assetId_fkey',
      columns: [table.assetId],
      foreignColumns: [asset.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Custody_kitCustodyId_fkey',
      columns: [table.kitCustodyId],
      foreignColumns: [kitCustody.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Custody_locationId_fkey',
      columns: [table.locationId],
      foreignColumns: [location.id],
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
      .$defaultFn(() => cuid()),
    name: text('name').notNull().default('Personal'),
    type: organizationTypeEnum('type').notNull().default('PERSONAL'),
    userId: text('userId').notNull(),
    currency: currencyEnum('currency').notNull().default('USD'),
    imageId: text('imageId'),
    enabledSso: boolean('enabledSso').notNull().default(false),
    ssoDetailsId: text('ssoDetailsId'),
    selfServiceCanSeeCustody: boolean('selfServiceCanSeeCustody').notNull().default(false),
    selfServiceCanSeeBookings: boolean('selfServiceCanSeeBookings').notNull().default(false),
    baseUserCanSeeCustody: boolean('baseUserCanSeeCustody').notNull().default(false),
    baseUserCanSeeBookings: boolean('baseUserCanSeeBookings').notNull().default(false),
    managersCanEditAssets: boolean('managersCanEditAssets').notNull().default(false),
    barcodesEnabled: boolean('barcodesEnabled').notNull().default(false),
    barcodesEnabledAt: timestamp('barcodesEnabledAt', { precision: 3 }),
    usedBarcodeTrial: boolean('usedBarcodeTrial').notNull().default(false),
    auditsEnabled: boolean('auditsEnabled').notNull().default(false),
    auditsEnabledAt: timestamp('auditsEnabledAt', { precision: 3 }),
    usedAuditTrial: boolean('usedAuditTrial').notNull().default(false),
    workspaceDisabled: boolean('workspaceDisabled').notNull().default(false),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    hasSequentialIdsMigrated: boolean('hasSequentialIdsMigrated').notNull().default(false),
    qrIdDisplayPreference: qrIdDisplayPreferenceEnum('qrIdDisplayPreference')
      .notNull()
      .default('QR_ID'),
    showShelfBranding: boolean('showShelfBranding').notNull().default(true),
    showQrCodesOnPdfs: boolean('showQrCodesOnPdfs').notNull().default(true),
    customEmailFooter: varchar('customEmailFooter', { length: 500 }),
  },
  (table): PgTableExtraConfigValue[] => [
    uniqueIndex('Organization_imageId_key').on(table.imageId),
    index('Organization_userId_idx').on(table.userId),
    index('Organization_ssoDetailsId_idx').on(table.ssoDetailsId),
    foreignKey({
      name: 'Organization_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Organization_imageId_fkey',
      columns: [table.imageId],
      foreignColumns: [image.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Organization_ssoDetailsId_fkey',
      columns: [table.ssoDetailsId],
      foreignColumns: [ssoDetails.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const userOrganization = pgTable(
  'UserOrganization',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    userId: text('userId').notNull(),
    organizationId: text('organizationId').notNull(),
    roles: organizationRolesEnum('roles').array(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    /**
     * Secret, unguessable token for this member's iCal calendar-feed subscription
     * in this workspace. Null until the user generates a feed; regenerating it
     * rotates the value (invalidating any previously-shared subscription URL).
     * `@unique` doubles as the lookup index for the public feed route.
     */
    calendarTokenId: text('calendarTokenId'),
  },
  (table) => [
    uniqueIndex('UserOrganization_calendarTokenId_key').on(table.calendarTokenId),
    index('UserOrganization_organizationId_idx').on(table.organizationId),
    uniqueIndex('UserOrganization_userId_organizationId_key').on(
      table.userId,
      table.organizationId,
    ),
    foreignKey({
      name: 'UserOrganization_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'UserOrganization_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const ssoDetails = pgTable('SsoDetails', {
  id: text('id')
    .primaryKey()
    .$defaultFn(() => cuid()),
  domain: text('domain').notNull(),
  baseUserGroupId: text('baseUserGroupId'),
  selfServiceGroupId: text('selfServiceGroupId'),
  adminGroupId: text('adminGroupId'),
  managerGroupId: text('managerGroupId'),
  requireSsoLogin: boolean('requireSsoLogin').notNull().default(true),
  createdAt: timestamp('createdAt', { precision: 3 })
    .notNull()
    .default(sql`CURRENT_TIMESTAMP`)
    .$defaultFn(utcNow),
  updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
})

export const scimToken = pgTable(
  'ScimToken',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    tokenHash: text('tokenHash').notNull(),
    label: text('label').notNull(),
    organizationId: text('organizationId').notNull(),
    createdById: text('createdById').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    lastUsedAt: timestamp('lastUsedAt', { precision: 3 }),
  },
  (table) => [
    uniqueIndex('ScimToken_tokenHash_key').on(table.tokenHash),
    index('ScimToken_organizationId_idx').on(table.organizationId),
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
 * *
 * * Org-scoped SCIM external ID linking a Shelf user to an identity-provider
 * * object within a specific organization.  A single Shelf user can appear in
 * * multiple organizations, each managed by a different identity provider, so the
 * * external ID must be scoped to the (user, organization) pair.
 */
export const userScimExternalId = pgTable(
  'UserScimExternalId',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    userId: text('userId').notNull(),
    organizationId: text('organizationId').notNull(),
    scimExternalId: text('scimExternalId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('UserScimExternalId_organizationId_idx').on(table.organizationId),
    uniqueIndex('UserScimExternalId_userId_organizationId_key').on(
      table.userId,
      table.organizationId,
    ),
    uniqueIndex('UserScimExternalId_organizationId_scimExternalId_key').on(
      table.organizationId,
      table.scimExternalId,
    ),
    foreignKey({
      name: 'UserScimExternalId_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'UserScimExternalId_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const tier = pgTable(
  'Tier',
  {
    id: tierIdEnum('id').primaryKey(),
    name: text('name').notNull(),
    tierLimitId: tierIdEnum('tierLimitId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('Tier_tierLimitId_key').on(table.tierLimitId),
    foreignKey({
      name: 'Tier_tierLimitId_fkey',
      columns: [table.tierLimitId],
      foreignColumns: [tierLimit.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const tierLimit = pgTable('TierLimit', {
  id: tierIdEnum('id').primaryKey(),
  canImportAssets: boolean('canImportAssets').notNull().default(false),
  canExportAssets: boolean('canExportAssets').notNull().default(false),
  canImportNRM: boolean('canImportNRM').notNull().default(false),
  canHideShelfBranding: boolean('canHideShelfBranding').notNull().default(false),
  canUseReports: boolean('canUseReports').notNull().default(false),
  canUseRecurringReminders: boolean('canUseRecurringReminders').notNull().default(false),
  canUsePushNotifications: boolean('canUsePushNotifications').notNull().default(false),
  maxCustomFields: integer('maxCustomFields').notNull().default(0),
  maxOrganizations: integer('maxOrganizations').notNull().default(1),
  createdAt: timestamp('createdAt', { precision: 3 })
    .notNull()
    .default(sql`CURRENT_TIMESTAMP`)
    .$defaultFn(utcNow),
  updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
})

export const customTierLimit = pgTable(
  'CustomTierLimit',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    userId: text('userId'),
    canImportAssets: boolean('canImportAssets').notNull().default(true),
    canExportAssets: boolean('canExportAssets').notNull().default(true),
    canImportNRM: boolean('canImportNRM').notNull().default(true),
    canHideShelfBranding: boolean('canHideShelfBranding').notNull().default(true),
    canUseReports: boolean('canUseReports').notNull().default(true),
    canUseRecurringReminders: boolean('canUseRecurringReminders').notNull().default(true),
    canUsePushNotifications: boolean('canUsePushNotifications').notNull().default(true),
    maxCustomFields: integer('maxCustomFields').notNull().default(1000),
    maxOrganizations: integer('maxOrganizations').notNull().default(1),
    isEnterprise: boolean('isEnterprise').notNull().default(false),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('CustomTierLimit_userId_key').on(table.userId),
    foreignKey({
      name: 'CustomTierLimit_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const customField = pgTable(
  'CustomField',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    name: text('name').notNull(),
    helpText: text('helpText'),
    required: boolean('required').notNull().default(false),
    active: boolean('active').notNull().default(true),
    type: customFieldTypeEnum('type').notNull().default('TEXT'),
    options: text('options').array(),
    /**
     * How a TEXT field's values are read. LINK requires a web link in every
     * value; any other type stays PLAIN (enforced by a CHECK constraint).
     */
    format: customFieldFormatEnum('format').notNull().default('PLAIN'),
    organizationId: text('organizationId').notNull(),
    userId: text('userId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    deletedAt: timestamp('deletedAt', { precision: 3 }),
  },
  (table) => [
    index('CustomField_organizationId_idx').on(table.organizationId),
    index('CustomField_userId_idx').on(table.userId),
    index('CustomField_organizationId_deletedAt_idx').on(table.organizationId, table.deletedAt),
    foreignKey({
      name: 'CustomField_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'CustomField_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

export const assetCustomFieldValue = pgTable(
  'AssetCustomFieldValue',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    value: jsonb('value').notNull(),
    assetId: text('assetId').notNull(),
    customFieldId: text('customFieldId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('AssetCustomFieldValue_lookup_idx').on(table.assetId, table.customFieldId),
    index('AssetCustomFieldValue_customFieldId_idx').on(table.customFieldId),
    foreignKey({
      name: 'AssetCustomFieldValue_assetId_fkey',
      columns: [table.assetId],
      foreignColumns: [asset.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AssetCustomFieldValue_customFieldId_fkey',
      columns: [table.customFieldId],
      foreignColumns: [customField.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * A custom field value an AssetModel gives the assets that join it.
 *
 * `value` has the exact shape of `AssetCustomFieldValue.value`, so filling an
 * asset is a plain copy. Rows cascade away with their model or field.
 */
export const assetModelCustomFieldDefault = pgTable(
  'AssetModelCustomFieldDefault',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    value: jsonb('value').notNull(),
    assetModelId: text('assetModelId').notNull(),
    customFieldId: text('customFieldId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('AssetModelCustomFieldDefault_customFieldId_idx').on(table.customFieldId),
    uniqueIndex('AssetModelCustomFieldDefault_assetModelId_customFieldId_key').on(
      table.assetModelId,
      table.customFieldId,
    ),
    foreignKey({
      name: 'AssetModelCustomFieldDefault_assetModelId_fkey',
      columns: [table.assetModelId],
      foreignColumns: [assetModel.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const invite = pgTable(
  'Invite',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    inviterId: text('inviterId').notNull(),
    organizationId: text('organizationId').notNull(),
    inviteeUserId: text('inviteeUserId'),
    teamMemberId: text('teamMemberId').notNull(),
    inviteeEmail: text('inviteeEmail').notNull(),
    status: inviteStatusesEnum('status').notNull().default('PENDING'),
    inviteCode: text('inviteCode').notNull(),
    roles: organizationRolesEnum('roles').array(),
    inviteMessage: varchar('inviteMessage', { length: 1000 }),
    expiresAt: timestamp('expiresAt', { precision: 3 }).notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('Invite_inviteeUserId_idx').on(table.inviteeUserId),
    index('Invite_inviterId_idx').on(table.inviterId),
    index('Invite_organizationId_idx').on(table.organizationId),
    index('Invite_teamMemberId_idx').on(table.teamMemberId),
    foreignKey({
      name: 'Invite_inviterId_fkey',
      columns: [table.inviterId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Invite_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Invite_inviteeUserId_fkey',
      columns: [table.inviteeUserId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Invite_teamMemberId_fkey',
      columns: [table.teamMemberId],
      foreignColumns: [teamMember.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

export const announcement = pgTable('Announcement', {
  id: text('id')
    .primaryKey()
    .$defaultFn(() => cuid()),
  name: text('name').notNull(),
  content: text('content').notNull(),
  link: text('link'),
  linkText: text('linkText'),
  published: boolean('published').notNull().default(false),
  createdAt: timestamp('createdAt', { precision: 3 })
    .notNull()
    .default(sql`CURRENT_TIMESTAMP`)
    .$defaultFn(utcNow),
  updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
})

/**
 * *
 * * Audit trail for all quantity changes on quantity-tracked assets.
 * * Every event records who performed it, who received/returned items,
 * * which booking it relates to, and the category of change.
 * * Quantity is always positive — direction is determined by category.
 */
export const consumptionLog = pgTable(
  'ConsumptionLog',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    assetId: text('assetId').notNull(),
    category: consumptionCategoryEnum('category').notNull(),
    quantity: integer('quantity').notNull(),
    note: text('note'),
    userId: text('userId').notNull(),
    bookingId: text('bookingId'),
    custodianId: text('custodianId'),
    bookingAssetId: text('bookingAssetId'),
    locationId: text('locationId'),
    stockChange: integer('stockChange'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('ConsumptionLog_assetId_createdAt_idx').on(table.assetId, table.createdAt),
    index('ConsumptionLog_userId_idx').on(table.userId),
    index('ConsumptionLog_bookingId_idx').on(table.bookingId),
    index('ConsumptionLog_custodianId_idx').on(table.custodianId),
    index('ConsumptionLog_bookingAssetId_idx').on(table.bookingAssetId),
    index('ConsumptionLog_locationId_idx').on(table.locationId),
    foreignKey({
      name: 'ConsumptionLog_assetId_fkey',
      columns: [table.assetId],
      foreignColumns: [asset.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'ConsumptionLog_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'ConsumptionLog_bookingId_fkey',
      columns: [table.bookingId],
      foreignColumns: [booking.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'ConsumptionLog_custodianId_fkey',
      columns: [table.custodianId],
      foreignColumns: [teamMember.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'ConsumptionLog_locationId_fkey',
      columns: [table.locationId],
      foreignColumns: [location.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

/**
 * *
 * * Explicit pivot table between Booking and Asset.
 * * Replaces the implicit M2M (_AssetToBooking) via a rename migration in Phase 3.
 */
export const bookingAsset = pgTable(
  'BookingAsset',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    bookingId: text('bookingId').notNull(),
    assetId: text('assetId').notNull(),
    quantity: integer('quantity').notNull().default(1),
    bookingModelRequestId: text('bookingModelRequestId'),
    assetKitId: text('assetKitId'),
    sourceKitId: text('sourceKitId'),
    checkedOutAt: timestamp('checkedOutAt', { withTimezone: true, precision: 3 }),
    checkedOutById: text('checkedOutById'),
    checkedInAt: timestamp('checkedInAt', { withTimezone: true, precision: 3 }),
    checkedInById: text('checkedInById'),
    checkedOutQuantity: integer('checkedOutQuantity').notNull().default(0),
    sourceLocationId: text('sourceLocationId'),
  },
  (table) => [
    index('BookingAsset_assetId_idx').on(table.assetId),
    index('BookingAsset_assetKitId_idx').on(table.assetKitId),
    index('BookingAsset_sourceKitId_idx').on(table.sourceKitId),
    index('BookingAsset_sourceLocationId_idx').on(table.sourceLocationId),
    foreignKey({
      name: 'BookingAsset_bookingId_fkey',
      columns: [table.bookingId],
      foreignColumns: [booking.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'BookingAsset_assetId_fkey',
      columns: [table.assetId],
      foreignColumns: [asset.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const booking = pgTable(
  'Booking',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    name: text('name').notNull(),
    status: bookingStatusEnum('status').notNull().default('DRAFT'),
    description: text('description').default(''),
    activeSchedulerReference: text('activeSchedulerReference'),
    creatorId: text('creatorId').notNull(),
    custodianUserId: text('custodianUserId'),
    custodianTeamMemberId: text('custodianTeamMemberId'),
    organizationId: text('organizationId').notNull(),
    createdAt: timestamp('createdAt', { withTimezone: true, precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { withTimezone: true, precision: 3 })
      .notNull()
      .$onUpdate(utcNow),
    from: timestamp('from', { withTimezone: true, precision: 3 }).notNull(),
    to: timestamp('to', { withTimezone: true, precision: 3 }).notNull(),
    originalFrom: timestamp('originalFrom', { withTimezone: true, precision: 3 }),
    originalTo: timestamp('originalTo', { withTimezone: true, precision: 3 }),
    autoArchivedAt: timestamp('autoArchivedAt', { withTimezone: true, precision: 3 }),
    archivedWithoutCheckin: boolean('archivedWithoutCheckin').notNull().default(false),
    cancellationReason: text('cancellationReason'),
  },
  (table) => [
    index('Booking_creatorId_idx').on(table.creatorId),
    index('Booking_custodianTeamMemberId_idx').on(table.custodianTeamMemberId),
    index('Booking_custodianUserId_idx').on(table.custodianUserId),
    index('Booking_organizationId_idx').on(table.organizationId),
    index('Booking_status_idx').on(table.status),
    index('Booking_organizationId_from_id_idx').on(table.organizationId, table.from, table.id),
    foreignKey({
      name: 'Booking_creatorId_fkey',
      columns: [table.creatorId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Booking_custodianUserId_fkey',
      columns: [table.custodianUserId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Booking_custodianTeamMemberId_fkey',
      columns: [table.custodianTeamMemberId],
      foreignColumns: [teamMember.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Booking_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const bookingSettings = pgTable(
  'BookingSettings',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    bufferStartTime: integer('bufferStartTime').notNull().default(0),
    tagsRequired: boolean('tagsRequired').notNull().default(false),
    maxBookingLength: integer('maxBookingLength'),
    maxBookingLengthSkipClosedDays: boolean('maxBookingLengthSkipClosedDays')
      .notNull()
      .default(false),
    autoArchiveBookings: boolean('autoArchiveBookings').notNull().default(false),
    autoArchiveDays: integer('autoArchiveDays').notNull().default(2),
    autoArchiveExpiredReservations: boolean('autoArchiveExpiredReservations')
      .notNull()
      .default(false),
    requireExplicitCheckinForAdmin: boolean('requireExplicitCheckinForAdmin')
      .notNull()
      .default(false),
    requireExplicitCheckinForSelfService: boolean('requireExplicitCheckinForSelfService')
      .notNull()
      .default(false),
    requireExplicitCheckoutForAdmin: boolean('requireExplicitCheckoutForAdmin')
      .notNull()
      .default(false),
    requireExplicitCheckoutForSelfService: boolean('requireExplicitCheckoutForSelfService')
      .notNull()
      .default(false),
    countKitsAsSingleUnit: boolean('countKitsAsSingleUnit').notNull().default(false),
    notifyBookingCreator: boolean('notifyBookingCreator').notNull().default(true),
    notifyAdminsOnNewBooking: boolean('notifyAdminsOnNewBooking').notNull().default(true),
    organizationId: text('organizationId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('BookingSettings_organizationId_key').on(table.organizationId),
    index('BookingSettings_organizationId_idx').on(table.organizationId),
    foreignKey({
      name: 'BookingSettings_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * A progressive check-in SESSION, mirroring PartialBookingCheckout.
 *
 * `assetIds` lists only FULLY reconciled assets: INDIVIDUAL always, and
 * QUANTITY_TRACKED only once its remaining units hit zero (partial returns
 * live in `ConsumptionLog`). Completing a booking outright writes no row here,
 * so the same warning applies — absence is not evidence of anything.
 * `BookingAsset.checkedInAt` is the per-slice fact.
 */
export const partialBookingCheckin = pgTable(
  'PartialBookingCheckin',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    assetIds: text('assetIds').array(),
    checkinCount: integer('checkinCount').notNull(),
    checkinTimestamp: timestamp('checkinTimestamp', { withTimezone: true, precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    bookingId: text('bookingId').notNull(),
    checkedInById: text('checkedInById').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('PartialBookingCheckin_bookingId_idx').on(table.bookingId),
    index('PartialBookingCheckin_checkedInById_idx').on(table.checkedInById),
    index('PartialBookingCheckin_checkinTimestamp_idx').on(table.checkinTimestamp),
    index('PartialBookingCheckin_bookingId_checkinTimestamp_idx').on(
      table.bookingId,
      table.checkinTimestamp,
    ),
    foreignKey({
      name: 'PartialBookingCheckin_bookingId_fkey',
      columns: [table.bookingId],
      foreignColumns: [booking.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'PartialBookingCheckin_checkedInById_fkey',
      columns: [table.checkedInById],
      foreignColumns: [user.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

/**
 * A progressive checkout SESSION: one scan batch, covering N assets via the
 * positional `assetIds` / `quantities` / `bookingAssetIds` arrays.
 *
 * Records sessions and the units they claimed — NOT whether an asset is out.
 * The all-at-once checkout (`checkoutBooking`, the Check out button) flips
 * statuses and writes NO row here, so **absence proves nothing** and must never
 * be read as "was not checked out". Reading it that way is what let a single
 * later scan mark every button-checked-out asset on a booking as
 * "never checked out". `BookingAsset.checkedOutAt` answers that question.
 */
export const partialBookingCheckout = pgTable(
  'PartialBookingCheckout',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    assetIds: text('assetIds').array(),
    quantities: integer('quantities').array().default([]),
    bookingAssetIds: text('bookingAssetIds').array().default([]),
    checkoutCount: integer('checkoutCount').notNull(),
    checkoutTimestamp: timestamp('checkoutTimestamp', { withTimezone: true, precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    bookingId: text('bookingId').notNull(),
    checkedOutById: text('checkedOutById').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('PartialBookingCheckout_bookingId_idx').on(table.bookingId),
    index('PartialBookingCheckout_checkedOutById_idx').on(table.checkedOutById),
    index('PartialBookingCheckout_checkoutTimestamp_idx').on(table.checkoutTimestamp),
    index('PartialBookingCheckout_bookingId_checkoutTimestamp_idx').on(
      table.bookingId,
      table.checkoutTimestamp,
    ),
    foreignKey({
      name: 'PartialBookingCheckout_bookingId_fkey',
      columns: [table.bookingId],
      foreignColumns: [booking.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'PartialBookingCheckout_checkedOutById_fkey',
      columns: [table.checkedOutById],
      foreignColumns: [user.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

/**
 * Native-app SSO web-delegated login — single-use session-handoff codes.
 * @see apps/webapp/app/routes/_auth+/oauth.callback.mobile.tsx
 */
export const mobileAuthCode = pgTable(
  'MobileAuthCode',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    userId: text('userId').notNull(),
    codeHash: text('codeHash').notNull(),
    codeChallenge: text('codeChallenge'),
    sessionCiphertext: text('sessionCiphertext'),
    consumedAt: timestamp('consumedAt', { precision: 3 }),
    expiresAt: timestamp('expiresAt', { precision: 3 }).notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    uniqueIndex('MobileAuthCode_codeHash_key').on(table.codeHash),
    index('MobileAuthCode_userId_idx').on(table.userId),
    index('MobileAuthCode_expiresAt_idx').on(table.expiresAt),
    foreignKey({
      name: 'MobileAuthCode_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * One Expo push token per installed companion app. A token identifies a
 * device, not a person: when someone else signs in on the same phone, the
 * row moves to them. Pushes go only to rows where `disabledAt` is null.
 * @see apps/webapp/app/modules/push/service.server.ts
 */
export const pushToken = pgTable(
  'PushToken',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    token: text('token').notNull(),
    platform: devicePlatformEnum('platform').notNull(),
    deviceName: text('deviceName'),
    appVersion: text('appVersion'),
    lastSeenAt: timestamp('lastSeenAt', { precision: 3 }).notNull(),
    disabledAt: timestamp('disabledAt', { precision: 3 }),
    userId: text('userId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('PushToken_token_key').on(table.token),
    index('PushToken_userId_idx').on(table.userId),
    foreignKey({
      name: 'PushToken_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * A person's push notification switches. A missing row means every switch
 * is on, so no backfill is needed. The three category columns are named
 * after `PushCategory` in apps/webapp/app/modules/push/types.ts.
 */
export const notificationPreference = pgTable(
  'NotificationPreference',
  {
    userId: text('userId').primaryKey(),
    pushEnabled: boolean('pushEnabled').notNull().default(true),
    bookings: boolean('bookings').notNull().default(true),
    audits: boolean('audits').notNull().default(true),
    stockAndReminders: boolean('stockAndReminders').notNull().default(true),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    foreignKey({
      name: 'NotificationPreference_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const kit = pgTable(
  'Kit',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    name: text('name').notNull(),
    description: text('description'),
    status: kitStatusEnum('status').notNull().default('AVAILABLE'),
    image: text('image'),
    imageExpiration: timestamp('imageExpiration', { precision: 3 }),
    organizationId: text('organizationId').notNull(),
    createdById: text('createdById').notNull(),
    categoryId: text('categoryId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    locationId: text('locationId'),
  },
  (table) => [
    index('Kit_createdById_idx').on(table.createdById),
    index('Kit_organizationId_idx').on(table.organizationId),
    index('Kit_categoryId_organizationId_idx').on(table.categoryId, table.organizationId),
    index('Kit_categoryId_organizationId_createdAt_idx').on(
      table.categoryId,
      table.organizationId,
      table.createdAt,
    ),
    index('Kit_categoryId_organizationId_name_idx').on(
      table.categoryId,
      table.organizationId,
      table.name,
    ),
    index('Kit_categoryId_organizationId_status_idx').on(
      table.categoryId,
      table.organizationId,
      table.status,
    ),
    foreignKey({
      name: 'Kit_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Kit_createdById_fkey',
      columns: [table.createdById],
      foreignColumns: [user.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Kit_categoryId_fkey',
      columns: [table.categoryId],
      foreignColumns: [category.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'Kit_locationId_fkey',
      columns: [table.locationId],
      foreignColumns: [location.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const kitCustody = pgTable(
  'KitCustody',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    custodianId: text('custodianId').notNull(),
    kitId: text('kitId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('KitCustody_kitId_key').on(table.kitId),
    index('KitCustody_custodianId_idx').on(table.custodianId),
    foreignKey({
      name: 'KitCustody_custodianId_fkey',
      columns: [table.custodianId],
      foreignColumns: [teamMember.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({ name: 'KitCustody_kitId_fkey', columns: [table.kitId], foreignColumns: [kit.id] })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * Pivot linking an `Asset` to the `Kit` it belongs to.
 *
 * `@@unique([assetId])` enforces "at most one kit per asset" today.
 * The structure already allows multi-row-per-asset; if multi-kit
 * allocation lands later, drop the `assetId` unique constraint and
 * add a type-aware single-row trigger + sum-within-total CONSTRAINT
 * TRIGGER (parallel to the Custody / BookingAsset pivots).
 */
export const assetKit = pgTable(
  'AssetKit',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    assetId: text('assetId').notNull(),
    kitId: text('kitId').notNull(),
    organizationId: text('organizationId').notNull(),
    quantity: integer('quantity').notNull().default(1),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('AssetKit_kitId_idx').on(table.kitId),
    index('AssetKit_organizationId_idx').on(table.organizationId),
    uniqueIndex('AssetKit_assetId_kitId_key').on(table.assetId, table.kitId),
    foreignKey({
      name: 'AssetKit_assetId_fkey',
      columns: [table.assetId],
      foreignColumns: [asset.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({ name: 'AssetKit_kitId_fkey', columns: [table.kitId], foreignColumns: [kit.id] })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AssetKit_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const assetLocation = pgTable(
  'AssetLocation',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    assetId: text('assetId').notNull(),
    locationId: text('locationId').notNull(),
    organizationId: text('organizationId').notNull(),
    quantity: integer('quantity').notNull().default(1),
    assetKitId: text('assetKitId'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('AssetLocation_locationId_idx').on(table.locationId),
    index('AssetLocation_organizationId_idx').on(table.organizationId),
    index('AssetLocation_assetId_idx').on(table.assetId),
    index('AssetLocation_assetKitId_idx').on(table.assetKitId),
    foreignKey({
      name: 'AssetLocation_assetId_fkey',
      columns: [table.assetId],
      foreignColumns: [asset.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AssetLocation_locationId_fkey',
      columns: [table.locationId],
      foreignColumns: [location.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AssetLocation_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AssetLocation_assetKitId_fkey',
      columns: [table.assetKitId],
      foreignColumns: [assetKit.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const assetReminder = pgTable(
  'AssetReminder',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    name: text('name').notNull(),
    message: text('message').notNull(),
    alertDateTime: timestamp('alertDateTime', { precision: 3 }).notNull(),
    activeSchedulerReference: text('activeSchedulerReference'),
    recurrenceUnit: reminderRecurrenceUnitEnum('recurrenceUnit'),
    recurrenceInterval: integer('recurrenceInterval'),
    recurrenceTimezone: text('recurrenceTimezone'),
    recurrenceEndsAt: timestamp('recurrenceEndsAt', { precision: 3 }),
    organizationId: text('organizationId').notNull(),
    assetId: text('assetId').notNull(),
    createdById: text('createdById').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('AssetReminder_assetId_alertDateTime_idx').on(table.assetId, table.alertDateTime),
    index('AssetReminder_name_message_idx').on(table.name, table.message),
    index('AssetReminder_organizationId_alertDateTime_assetId_idx').on(
      table.organizationId,
      table.alertDateTime,
      table.assetId,
    ),
    index('AssetReminder_alertDateTime_activeSchedulerReference_idx').on(
      table.alertDateTime,
      table.activeSchedulerReference,
    ),
    index('AssetReminder_createdById_idx').on(table.createdById),
    foreignKey({
      name: 'AssetReminder_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AssetReminder_assetId_fkey',
      columns: [table.assetId],
      foreignColumns: [asset.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AssetReminder_createdById_fkey',
      columns: [table.createdById],
      foreignColumns: [user.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

export const workingHours = pgTable(
  'WorkingHours',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    enabled: boolean('enabled').notNull().default(false),
    weeklySchedule: jsonb('weeklySchedule')
      .notNull()
      .default({
        '0': { isOpen: false },
        '1': { isOpen: true, openTime: '09:00', closeTime: '17:00' },
        '2': { isOpen: true, openTime: '09:00', closeTime: '17:00' },
        '3': { isOpen: true, openTime: '09:00', closeTime: '17:00' },
        '4': { isOpen: true, openTime: '09:00', closeTime: '17:00' },
        '5': { isOpen: true, openTime: '09:00', closeTime: '17:00' },
        '6': { isOpen: false },
      }),
    organizationId: text('organizationId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    uniqueIndex('WorkingHours_organizationId_key').on(table.organizationId),
    index('WorkingHours_organizationId_idx').on(table.organizationId),
    foreignKey({
      name: 'WorkingHours_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const workingHoursOverride = pgTable(
  'WorkingHoursOverride',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    date: utcDate('date').notNull(),
    isOpen: boolean('isOpen').notNull().default(false),
    openTime: text('openTime'),
    closeTime: text('closeTime'),
    reason: text('reason'),
    workingHoursId: text('workingHoursId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('WorkingHoursOverride_workingHoursId_date_idx').on(table.workingHoursId, table.date),
    index('WorkingHoursOverride_date_isOpen_idx').on(table.date, table.isOpen),
    uniqueIndex('WorkingHoursOverride_workingHoursId_date_key').on(
      table.workingHoursId,
      table.date,
    ),
    foreignKey({
      name: 'WorkingHoursOverride_workingHoursId_fkey',
      columns: [table.workingHoursId],
      foreignColumns: [workingHours.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const update = pgTable(
  'Update',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    title: text('title').notNull(),
    content: text('content').notNull(),
    url: text('url'),
    imageUrl: text('imageUrl'),
    publishDate: timestamp('publishDate', { precision: 3 }).notNull(),
    status: updateStatusEnum('status').notNull().default('DRAFT'),
    targetRoles: organizationRolesEnum('targetRoles').array(),
    clickCount: integer('clickCount').notNull().default(0),
    viewCount: integer('viewCount').notNull().default(0),
    createdById: text('createdById').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('Update_status_publishDate_idx').on(table.status, table.publishDate),
    index('Update_publishDate_idx').on(table.publishDate),
    index('Update_createdById_idx').on(table.createdById),
    foreignKey({
      name: 'Update_createdById_fkey',
      columns: [table.createdById],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const userUpdateRead = pgTable(
  'UserUpdateRead',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    userId: text('userId').notNull(),
    updateId: text('updateId').notNull(),
    readAt: timestamp('readAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('UserUpdateRead_userId_idx').on(table.userId),
    index('UserUpdateRead_updateId_idx').on(table.updateId),
    index('UserUpdateRead_readAt_idx').on(table.readAt),
    uniqueIndex('UserUpdateRead_userId_updateId_key').on(table.userId, table.updateId),
    foreignKey({
      name: 'UserUpdateRead_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'UserUpdateRead_updateId_fkey',
      columns: [table.updateId],
      foreignColumns: [update.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const auditNote = pgTable(
  'AuditNote',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    content: text('content').notNull(),
    type: noteTypeEnum('type').notNull().default('COMMENT'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    userId: text('userId'),
    auditSessionId: text('auditSessionId').notNull(),
    auditAssetId: text('auditAssetId'),
  },
  (table) => [
    index('AuditNote_auditSessionId_idx').on(table.auditSessionId),
    index('AuditNote_userId_idx').on(table.userId),
    index('AuditNote_auditAssetId_idx').on(table.auditAssetId),
    foreignKey({
      name: 'AuditNote_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AuditNote_auditSessionId_fkey',
      columns: [table.auditSessionId],
      foreignColumns: [auditSession.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AuditNote_auditAssetId_fkey',
      columns: [table.auditAssetId],
      foreignColumns: [auditAsset.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const teamMemberNote = pgTable(
  'TeamMemberNote',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    content: text('content').notNull(),
    type: noteTypeEnum('type').notNull().default('COMMENT'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    userId: text('userId'),
    teamMemberId: text('teamMemberId').notNull(),
    organizationId: text('organizationId').notNull(),
  },
  (table) => [
    index('TeamMemberNote_teamMemberId_idx').on(table.teamMemberId),
    index('TeamMemberNote_userId_idx').on(table.userId),
    index('TeamMemberNote_organizationId_idx').on(table.organizationId),
    index('TeamMemberNote_teamMemberId_organizationId_idx').on(
      table.teamMemberId,
      table.organizationId,
    ),
    foreignKey({
      name: 'TeamMemberNote_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'TeamMemberNote_teamMemberId_fkey',
      columns: [table.teamMemberId],
      foreignColumns: [teamMember.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'TeamMemberNote_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const auditSession = pgTable(
  'AuditSession',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    name: text('name').notNull(),
    description: text('description'),
    targetId: text('targetId'),
    status: auditStatusEnum('status').notNull().default('PENDING'),
    scopeMeta: jsonb('scopeMeta'),
    expectedAssetCount: integer('expectedAssetCount').notNull().default(0),
    foundAssetCount: integer('foundAssetCount').notNull().default(0),
    missingAssetCount: integer('missingAssetCount').notNull().default(0),
    unexpectedAssetCount: integer('unexpectedAssetCount').notNull().default(0),
    startedAt: timestamp('startedAt', { precision: 3 }),
    dueDate: timestamp('dueDate', { precision: 3 }),
    completedAt: timestamp('completedAt', { precision: 3 }),
    cancelledAt: timestamp('cancelledAt', { precision: 3 }),
    activeSchedulerReference: text('activeSchedulerReference'),
    createdById: text('createdById').notNull(),
    organizationId: text('organizationId').notNull(),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('AuditSession_organizationId_status_idx').on(table.organizationId, table.status),
    index('AuditSession_createdById_idx').on(table.createdById),
    index('AuditSession_status_createdAt_idx').on(table.status, table.createdAt),
    foreignKey({
      name: 'AuditSession_createdById_fkey',
      columns: [table.createdById],
      foreignColumns: [user.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AuditSession_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * One CSV content import. Holds the parsed rows from preview until the import
 * completes, so the job imports exactly the file the user previewed.
 * A partial unique index on (organizationId) WHERE status IN (QUEUED, RUNNING) allows one
 * active import per workspace. Prisma cannot express it; it lives in migration
 * 20261015120100_asset_import_one_active_per_workspace.
 */
export const assetImport = pgTable(
  'AssetImport',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    status: assetImportStatusEnum('status').notNull().default('DRAFT'),
    phase: assetImportPhaseEnum('phase').notNull().default('PREPARE'),
    fileName: text('fileName'),
    totalRows: integer('totalRows').notNull(),
    /** Rows committed so far. Advanced in the same transaction as each chunk. */
    processedRows: integer('processedRows').notNull().default(0),
    /** Parsed rows (`StoredImportRow[]`). Cleared when the import completes. */
    rows: jsonb('rows'),
    /** Name to id maps written by the PREPARE step (`ResolvedImportEntities`). */
    resolved: jsonb('resolved'),
    canUseBarcodes: boolean('canUseBarcodes').notNull().default(false),
    /** User-facing reason the import failed. Null unless status is FAILED. */
    errorMessage: text('errorMessage'),
    /** Rows imported without a conflicting piece (`ImportWarning[]`), capped at 500. */
    warnings: jsonb('warnings'),
    /** Exact number of warnings, including any beyond the stored cap. */
    warningCount: integer('warningCount').notNull().default(0),
    /** A step holds the import until this time. Null when no step is running. */
    lockedUntil: timestamp('lockedUntil', { precision: 3 }),
    confirmedAt: timestamp('confirmedAt', { precision: 3 }),
    startedAt: timestamp('startedAt', { precision: 3 }),
    completedAt: timestamp('completedAt', { precision: 3 }),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    createdById: text('createdById').notNull(),
    organizationId: text('organizationId').notNull(),
  },
  (table) => [
    index('AssetImport_organizationId_status_idx').on(table.organizationId, table.status),
    index('AssetImport_createdById_organizationId_status_idx').on(
      table.createdById,
      table.organizationId,
      table.status,
    ),
    foreignKey({
      name: 'AssetImport_createdById_fkey',
      columns: [table.createdById],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AssetImport_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

/**
 * What Shelf knows about one web link used in a workspace's custom fields,
 * keyed by its `LinkKey` from @shelf/links. Two writers on disjoint columns:
 * the link preview job (fetch columns) and people (customTitle).
 */
export const linkPreview = pgTable(
  'LinkPreview',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    /**
     * `LinkKey` from @shelf/links, never a raw field value. At most 2048
     * characters (CHECK constraint in the migration).
     */
    url: text('url').notNull(),
    pageTitle: text('pageTitle'),
    /** Null until the first fetch finishes. */
    outcome: linkPreviewOutcomeEnum('outcome'),
    attempts: integer('attempts').notNull().default(0),
    fetchedAt: timestamp('fetchedAt', { precision: 3 }),
    /**
     * When the link may be fetched next; null means never again. A claim moves
     * it a few minutes ahead, which is also the lease on a fetch in flight.
     */
    nextFetchAt: timestamp('nextFetchAt', { precision: 3 })
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    /** The name a person gave the link in this workspace; wins over pageTitle. */
    customTitle: text('customTitle'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
    organizationId: text('organizationId').notNull(),
  },
  (table) => [
    uniqueIndex('LinkPreview_organizationId_url_key').on(table.organizationId, table.url),
    foreignKey({
      name: 'LinkPreview_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const auditAssignment = pgTable(
  'AuditAssignment',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    auditSessionId: text('auditSessionId').notNull(),
    userId: text('userId').notNull(),
    role: auditAssignmentRoleEnum('role'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('AuditAssignment_userId_idx').on(table.userId),
    uniqueIndex('AuditAssignment_auditSessionId_userId_key').on(table.auditSessionId, table.userId),
    foreignKey({
      name: 'AuditAssignment_auditSessionId_fkey',
      columns: [table.auditSessionId],
      foreignColumns: [auditSession.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AuditAssignment_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const auditAsset = pgTable(
  'AuditAsset',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    auditSessionId: text('auditSessionId').notNull(),
    assetId: text('assetId').notNull(),
    expected: boolean('expected').notNull().default(true),
    status: auditAssetStatusEnum('status').notNull().default('PENDING'),
    scannedById: text('scannedById'),
    scannedAt: timestamp('scannedAt', { precision: 3 }),
    metadata: jsonb('metadata'),
    /**
     * Units the audit expects to find, snapshotted when the row is created: the
     * units in the audit's scope (placed at its locations, held in its kits, or
     * held by its person; at most `Asset.quantity`), or `Asset.quantity` when
     * none are in scope or the audit has no scope. Null for INDIVIDUAL assets,
     * which count as one.
     *
     * A snapshot, not a live read: stock that moves during the audit must not
     * move the target the auditor is counting against.
     */
    expectedQuantity: integer('expectedQuantity'),
    /**
     * Units the auditor counted. Null until a quantity-tracked asset is
     * counted, and always null for INDIVIDUAL assets. A FOUND row with a count
     * below `expectedQuantity` is a shortfall: the status stays FOUND (the pool
     * was located), the gap is this number.
     *
     * Both columns are non-negative, enforced by CHECK constraints in raw SQL
     * (migration `20261104120000_add_audit_asset_quantities`).
     */
    foundQuantity: integer('foundQuantity'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('AuditAsset_status_idx').on(table.status),
    index('AuditAsset_scannedById_idx').on(table.scannedById),
    index('AuditAsset_assetId_idx').on(table.assetId),
    uniqueIndex('AuditAsset_auditSessionId_assetId_key').on(table.auditSessionId, table.assetId),
    foreignKey({
      name: 'AuditAsset_auditSessionId_fkey',
      columns: [table.auditSessionId],
      foreignColumns: [auditSession.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AuditAsset_assetId_fkey',
      columns: [table.assetId],
      foreignColumns: [asset.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AuditAsset_scannedById_fkey',
      columns: [table.scannedById],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const auditScan = pgTable(
  'AuditScan',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    auditSessionId: text('auditSessionId').notNull(),
    auditAssetId: text('auditAssetId'),
    assetId: text('assetId'),
    scannedById: text('scannedById'),
    code: text('code'),
    /**
     * What was scanned, recorded BY VALUE at scan time.
     *
     * `asset` and `auditAsset` are both SetNull/Cascade, so deleting an asset
     * leaves this row pointing at nothing: no name, and no way to tell whether
     * it had been expected. An audit is a historical record, and a record whose
     * entries can silently lose their meaning is not one — so the two
     * descriptive facts are snapshotted here rather than only referenced.
     *
     * Readers prefer the LIVE asset when it still exists (a rename should show
     * the current name) and fall back to these.
     *
     * Null means the snapshot could not be taken: scans whose asset had already
     * been deleted before these columns were backfilled. Nothing can recover
     * those, and they degrade to the scanned `code`, which also survives. Every
     * other row carries a snapshot, so null is never merely "old".
     */
    assetTitle: text('assetTitle'),
    wasExpected: boolean('wasExpected'),
    metadata: jsonb('metadata'),
    scannedAt: timestamp('scannedAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
  },
  (table) => [
    index('AuditScan_auditSessionId_scannedAt_idx').on(table.auditSessionId, table.scannedAt),
    index('AuditScan_auditAssetId_idx').on(table.auditAssetId),
    index('AuditScan_assetId_idx').on(table.assetId),
    foreignKey({
      name: 'AuditScan_auditSessionId_fkey',
      columns: [table.auditSessionId],
      foreignColumns: [auditSession.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AuditScan_auditAssetId_fkey',
      columns: [table.auditAssetId],
      foreignColumns: [auditAsset.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AuditScan_assetId_fkey',
      columns: [table.assetId],
      foreignColumns: [asset.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AuditScan_scannedById_fkey',
      columns: [table.scannedById],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
  ],
)

export const auditImage = pgTable(
  'AuditImage',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    imageUrl: text('imageUrl').notNull(),
    thumbnailUrl: text('thumbnailUrl'),
    description: text('description'),
    auditSessionId: text('auditSessionId').notNull(),
    auditAssetId: text('auditAssetId'),
    uploadedById: text('uploadedById'),
    organizationId: text('organizationId').notNull(),
    /**
     * When the photo was taken, from its EXIF data (or the phone app's copy of
     * it). Null when the photo carried no usable date. EXIF can be edited, so
     * this is the best signal available, not proof.
     */
    capturedAt: timestamp('capturedAt', { precision: 3 }),
    /**
     * How the photo reached Shelf. Null only for photos uploaded before this
     * field existed; a phone upload with no or an unknown source is PHONE.
     */
    captureSource: auditImageCaptureSourceEnum('captureSource'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    updatedAt: timestamp('updatedAt', { precision: 3 }).notNull().$onUpdate(utcNow),
  },
  (table) => [
    index('AuditImage_auditSessionId_idx').on(table.auditSessionId),
    index('AuditImage_auditAssetId_idx').on(table.auditAssetId),
    index('AuditImage_organizationId_idx').on(table.organizationId),
    index('AuditImage_uploadedById_idx').on(table.uploadedById),
    foreignKey({
      name: 'AuditImage_auditSessionId_fkey',
      columns: [table.auditSessionId],
      foreignColumns: [auditSession.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AuditImage_auditAssetId_fkey',
      columns: [table.auditAssetId],
      foreignColumns: [auditAsset.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AuditImage_uploadedById_fkey',
      columns: [table.uploadedById],
      foreignColumns: [user.id],
    })
      .onDelete('set null')
      .onUpdate('cascade'),
    foreignKey({
      name: 'AuditImage_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const roleChangeLog = pgTable(
  'RoleChangeLog',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    previousRole: organizationRolesEnum('previousRole').notNull(),
    newRole: organizationRolesEnum('newRole').notNull(),
    /** How the change was made. SSO rows name the member themself as `changedBy`. */
    source: roleChangeSourceEnum('source').notNull().default('MANUAL'),
    createdAt: timestamp('createdAt', { precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    userId: text('userId').notNull(),
    changedById: text('changedById').notNull(),
    organizationId: text('organizationId').notNull(),
  },
  (table) => [
    index('RoleChangeLog_userId_organizationId_idx').on(table.userId, table.organizationId),
    index('RoleChangeLog_organizationId_createdAt_idx').on(table.organizationId, table.createdAt),
    foreignKey({
      name: 'RoleChangeLog_userId_fkey',
      columns: [table.userId],
      foreignColumns: [user.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'RoleChangeLog_changedById_fkey',
      columns: [table.changedById],
      foreignColumns: [user.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
    foreignKey({
      name: 'RoleChangeLog_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('restrict')
      .onUpdate('cascade'),
  ],
)

export const activityEvent = pgTable(
  'ActivityEvent',
  {
    id: text('id')
      .primaryKey()
      .$defaultFn(() => cuid()),
    organizationId: text('organizationId').notNull(),
    occurredAt: timestamp('occurredAt', { withTimezone: true, precision: 3 })
      .notNull()
      .default(sql`CURRENT_TIMESTAMP`)
      .$defaultFn(utcNow),
    actorUserId: text('actorUserId'),
    actorSnapshot: jsonb('actorSnapshot'),
    action: activityActionEnum('action').notNull(),
    entityType: activityEntityEnum('entityType').notNull(),
    entityId: text('entityId').notNull(),
    assetId: text('assetId'),
    bookingId: text('bookingId'),
    auditSessionId: text('auditSessionId'),
    auditAssetId: text('auditAssetId'),
    kitId: text('kitId'),
    locationId: text('locationId'),
    teamMemberId: text('teamMemberId'),
    targetUserId: text('targetUserId'),
    field: text('field'),
    fromValue: jsonb('fromValue'),
    toValue: jsonb('toValue'),
    meta: jsonb('meta'),
  },
  (table) => [
    index('ActivityEvent_organizationId_occurredAt_idx').on(table.organizationId, table.occurredAt),
    index('ActivityEvent_organizationId_action_occurredAt_idx').on(
      table.organizationId,
      table.action,
      table.occurredAt,
    ),
    index('ActivityEvent_organizationId_entityType_entityId_occurredAt_idx').on(
      table.organizationId,
      table.entityType,
      table.entityId,
      table.occurredAt,
    ),
    index('ActivityEvent_actorUserId_occurredAt_idx').on(table.actorUserId, table.occurredAt),
    index('ActivityEvent_assetId_occurredAt_idx').on(table.assetId, table.occurredAt),
    index('ActivityEvent_bookingId_occurredAt_idx').on(table.bookingId, table.occurredAt),
    index('ActivityEvent_auditSessionId_occurredAt_idx').on(table.auditSessionId, table.occurredAt),
    foreignKey({
      name: 'ActivityEvent_organizationId_fkey',
      columns: [table.organizationId],
      foreignColumns: [organization.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const roleToUser = pgTable(
  '_RoleToUser',
  { A: text('A').notNull(), B: text('B').notNull() },
  (table) => [
    primaryKey({ name: '_RoleToUser_AB_pkey', columns: [table.A, table.B] }),
    index('_RoleToUser_B_index').on(table.B),
    foreignKey({ name: '_RoleToUser_A_fkey', columns: [table.A], foreignColumns: [role.id] })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({ name: '_RoleToUser_B_fkey', columns: [table.B], foreignColumns: [user.id] })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const assetToTag = pgTable(
  '_AssetToTag',
  { A: text('A').notNull(), B: text('B').notNull() },
  (table) => [
    primaryKey({ name: '_AssetToTag_AB_pkey', columns: [table.A, table.B] }),
    index('_AssetToTag_B_index').on(table.B),
    foreignKey({ name: '_AssetToTag_A_fkey', columns: [table.A], foreignColumns: [asset.id] })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({ name: '_AssetToTag_B_fkey', columns: [table.B], foreignColumns: [tag.id] })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const categoryToCustomField = pgTable(
  '_CategoryToCustomField',
  { A: text('A').notNull(), B: text('B').notNull() },
  (table) => [
    primaryKey({ name: '_CategoryToCustomField_AB_pkey', columns: [table.A, table.B] }),
    index('_CategoryToCustomField_B_index').on(table.B),
    foreignKey({
      name: '_CategoryToCustomField_A_fkey',
      columns: [table.A],
      foreignColumns: [category.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: '_CategoryToCustomField_B_fkey',
      columns: [table.B],
      foreignColumns: [customField.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const bookingToTag = pgTable(
  '_BookingToTag',
  { A: text('A').notNull(), B: text('B').notNull() },
  (table) => [
    primaryKey({ name: '_BookingToTag_AB_pkey', columns: [table.A, table.B] }),
    index('_BookingToTag_B_index').on(table.B),
    foreignKey({ name: '_BookingToTag_A_fkey', columns: [table.A], foreignColumns: [booking.id] })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({ name: '_BookingToTag_B_fkey', columns: [table.B], foreignColumns: [tag.id] })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const assetReminderToTeamMember = pgTable(
  '_AssetReminderToTeamMember',
  { A: text('A').notNull(), B: text('B').notNull() },
  (table) => [
    primaryKey({ name: '_AssetReminderToTeamMember_AB_pkey', columns: [table.A, table.B] }),
    index('_AssetReminderToTeamMember_B_index').on(table.B),
    foreignKey({
      name: '_AssetReminderToTeamMember_A_fkey',
      columns: [table.A],
      foreignColumns: [assetReminder.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: '_AssetReminderToTeamMember_B_fkey',
      columns: [table.B],
      foreignColumns: [teamMember.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const bookingSettingsAlwaysNotify = pgTable(
  '_BookingSettingsAlwaysNotify',
  { A: text('A').notNull(), B: text('B').notNull() },
  (table) => [
    primaryKey({ name: '_BookingSettingsAlwaysNotify_AB_pkey', columns: [table.A, table.B] }),
    index('_BookingSettingsAlwaysNotify_B_index').on(table.B),
    foreignKey({
      name: '_BookingSettingsAlwaysNotify_A_fkey',
      columns: [table.A],
      foreignColumns: [bookingSettings.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: '_BookingSettingsAlwaysNotify_B_fkey',
      columns: [table.B],
      foreignColumns: [teamMember.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const bookingNotificationRecipients = pgTable(
  '_BookingNotificationRecipients',
  { A: text('A').notNull(), B: text('B').notNull() },
  (table) => [
    primaryKey({ name: '_BookingNotificationRecipients_AB_pkey', columns: [table.A, table.B] }),
    index('_BookingNotificationRecipients_B_index').on(table.B),
    foreignKey({
      name: '_BookingNotificationRecipients_A_fkey',
      columns: [table.A],
      foreignColumns: [booking.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
    foreignKey({
      name: '_BookingNotificationRecipients_B_fkey',
      columns: [table.B],
      foreignColumns: [teamMember.id],
    })
      .onDelete('cascade')
      .onUpdate('cascade'),
  ],
)

export const relations = defineRelations(
  {
    image,
    user,
    userContact,
    userBusinessIntel,
    asset,
    assetFilterPreset,
    assetIndexSettings,
    pdfTemplate,
    category,
    assetModel,
    bookingModelRequest,
    tag,
    note,
    bookingNote,
    locationNote,
    qr,
    barcode,
    printBatch,
    reportFound,
    scan,
    location,
    role,
    teamMember,
    custody,
    organization,
    userOrganization,
    ssoDetails,
    scimToken,
    userScimExternalId,
    tier,
    tierLimit,
    customTierLimit,
    customField,
    assetCustomFieldValue,
    assetModelCustomFieldDefault,
    invite,
    announcement,
    consumptionLog,
    bookingAsset,
    booking,
    bookingSettings,
    partialBookingCheckin,
    partialBookingCheckout,
    mobileAuthCode,
    pushToken,
    notificationPreference,
    kit,
    kitCustody,
    assetKit,
    assetLocation,
    assetReminder,
    workingHours,
    workingHoursOverride,
    update,
    userUpdateRead,
    auditNote,
    teamMemberNote,
    auditSession,
    assetImport,
    linkPreview,
    auditAssignment,
    auditAsset,
    auditScan,
    auditImage,
    roleChangeLog,
    activityEvent,
    roleToUser,
    assetToTag,
    categoryToCustomField,
    bookingToTag,
    assetReminderToTeamMember,
    bookingSettingsAlwaysNotify,
    bookingNotificationRecipients,
  },
  (r) => ({
    image: {
      location: r.one.location({ from: r.image.id, to: r.location.imageId }),
      organization: r.one.organization({
        from: r.image.id,
        to: r.organization.imageId,
        alias: 'orgImage',
      }),
      ownerOrg: r.one.organization({
        from: r.image.ownerOrgId,
        to: r.organization.id,
        optional: false,
        alias: 'owner',
      }),
      user: r.one.user({ from: r.image.userId, to: r.user.id, optional: false }),
    },
    user: {
      lastSelectedOrganization: r.one.organization({
        from: r.user.lastSelectedOrganizationId,
        to: r.organization.id,
        alias: 'lastSelectedOrg',
      }),
      assets: r.many.asset(),
      categories: r.many.category(),
      notes: r.many.note(),
      bookingNotes: r.many.bookingNote(),
      auditNotes: r.many.auditNote(),
      locationNotes: r.many.locationNote(),
      qrCodes: r.many.qr(),
      scans: r.many.scan(),
      tags: r.many.tag(),
      roles: r.many.role({
        from: r.user.id.through(r.roleToUser.B),
        to: r.role.id.through(r.roleToUser.A),
      }),
      locations: r.many.location(),
      images: r.many.image(),
      organizations: r.many.organization({ alias: 'OrganizationToUser' }),
      customFields: r.many.customField(),
      sentInvites: r.many.invite({ alias: 'inviter' }),
      receivedInvites: r.many.invite({ alias: 'invitee' }),
      teamMembers: r.many.teamMember(),
      userOrganizations: r.many.userOrganization(),
      bookings: r.many.booking({ alias: 'creator' }),
      custodies: r.many.booking({ alias: 'custodian' }),
      createdKits: r.many.kit(),
      tier: r.one.tier({ from: r.user.tierId, to: r.tier.id, optional: false }),
      assetReminders: r.many.assetReminder(),
      contact: r.one.userContact({ from: r.user.id, to: r.userContact.userId }),
      businessIntel: r.one.userBusinessIntel({ from: r.user.id, to: r.userBusinessIntel.userId }),
      auditsCreated: r.many.auditSession(),
      assetImports: r.many.assetImport(),
      auditAssignments: r.many.auditAssignment(),
      auditAssetsScanned: r.many.auditAsset(),
      auditScans: r.many.auditScan(),
      auditImagesUploaded: r.many.auditImage(),
      assetIndexSettings: r.many.assetIndexSettings(),
      partialCheckins: r.many.partialBookingCheckin(),
      partialCheckouts: r.many.partialBookingCheckout(),
      mobileAuthCodes: r.many.mobileAuthCode(),
      pushTokens: r.many.pushToken(),
      notificationPreference: r.one.notificationPreference({
        from: r.user.id,
        to: r.notificationPreference.userId,
      }),
      createdUpdates: r.many.update(),
      updateReads: r.many.userUpdateRead(),
      customTierLimit: r.one.customTierLimit({ from: r.user.id, to: r.customTierLimit.userId }),
      assetFilterPresets: r.many.assetFilterPreset(),
      assetModels: r.many.assetModel(),
      consumptionLogs: r.many.consumptionLog(),
      scimExternalIds: r.many.userScimExternalId(),
      roleChangesReceived: r.many.roleChangeLog({ alias: 'roleChangeTarget' }),
      roleChangesPerformed: r.many.roleChangeLog({ alias: 'roleChangeActor' }),
      teamMemberNotesAuthored: r.many.teamMemberNote(),
    },
    userContact: {
      user: r.one.user({ from: r.userContact.userId, to: r.user.id, optional: false }),
    },
    userBusinessIntel: {
      user: r.one.user({ from: r.userBusinessIntel.userId, to: r.user.id, optional: false }),
    },
    asset: {
      assetModel: r.one.assetModel({ from: r.asset.assetModelId, to: r.assetModel.id }),
      user: r.one.user({ from: r.asset.userId, to: r.user.id, optional: false }),
      organization: r.one.organization({
        from: r.asset.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      category: r.one.category({ from: r.asset.categoryId, to: r.category.id }),
      custody: r.many.custody(),
      notes: r.many.note(),
      qrCodes: r.many.qr(),
      barcodes: r.many.barcode({ alias: 'AssetBarcodes' }),
      reports: r.many.reportFound(),
      tags: r.many.tag({
        from: r.asset.id.through(r.assetToTag.A),
        to: r.tag.id.through(r.assetToTag.B),
      }),
      customFields: r.many.assetCustomFieldValue(),
      bookingAssets: r.many.bookingAsset(),
      consumptionLogs: r.many.consumptionLog(),
      reminders: r.many.assetReminder(),
      auditEntries: r.many.auditAsset(),
      auditScans: r.many.auditScan(),
      scans: r.many.scan(),
      assetKits: r.many.assetKit(),
      assetLocations: r.many.assetLocation(),
      preferredBarcode: r.one.barcode({
        from: r.asset.preferredBarcodeId,
        to: r.barcode.id,
        alias: 'AssetPreferredBarcode',
      }),
    },
    assetFilterPreset: {
      organization: r.one.organization({
        from: r.assetFilterPreset.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      owner: r.one.user({ from: r.assetFilterPreset.ownerId, to: r.user.id, optional: false }),
    },
    assetIndexSettings: {
      user: r.one.user({ from: r.assetIndexSettings.userId, to: r.user.id, optional: false }),
      organization: r.one.organization({
        from: r.assetIndexSettings.organizationId,
        to: r.organization.id,
        optional: false,
      }),
    },
    pdfTemplate: {
      organization: r.one.organization({
        from: r.pdfTemplate.organizationId,
        to: r.organization.id,
        optional: false,
      }),
    },
    category: {
      assets: r.many.asset(),
      kits: r.many.kit(),
      user: r.one.user({ from: r.category.userId, to: r.user.id, optional: false }),
      customFields: r.many.customField({
        from: r.category.id.through(r.categoryToCustomField.A),
        to: r.customField.id.through(r.categoryToCustomField.B),
      }),
      assetModelDefaults: r.many.assetModel(),
      organization: r.one.organization({
        from: r.category.organizationId,
        to: r.organization.id,
        optional: false,
      }),
    },
    assetModel: {
      defaultCategory: r.one.category({ from: r.assetModel.defaultCategoryId, to: r.category.id }),
      customFieldDefaults: r.many.assetModelCustomFieldDefault(),
      assets: r.many.asset(),
      modelRequests: r.many.bookingModelRequest(),
      organization: r.one.organization({
        from: r.assetModel.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      createdBy: r.one.user({ from: r.assetModel.userId, to: r.user.id, optional: false }),
    },
    bookingModelRequest: {
      booking: r.one.booking({
        from: r.bookingModelRequest.bookingId,
        to: r.booking.id,
        optional: false,
      }),
      assetModel: r.one.assetModel({
        from: r.bookingModelRequest.assetModelId,
        to: r.assetModel.id,
        optional: false,
      }),
    },
    tag: {
      assets: r.many.asset({
        from: r.tag.id.through(r.assetToTag.B),
        to: r.asset.id.through(r.assetToTag.A),
      }),
      user: r.one.user({ from: r.tag.userId, to: r.user.id, optional: false }),
      organization: r.one.organization({
        from: r.tag.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      bookings: r.many.booking({
        from: r.tag.id.through(r.bookingToTag.B),
        to: r.booking.id.through(r.bookingToTag.A),
      }),
    },
    note: {
      user: r.one.user({ from: r.note.userId, to: r.user.id }),
      asset: r.one.asset({ from: r.note.assetId, to: r.asset.id, optional: false }),
    },
    bookingNote: {
      user: r.one.user({ from: r.bookingNote.userId, to: r.user.id }),
      booking: r.one.booking({ from: r.bookingNote.bookingId, to: r.booking.id, optional: false }),
    },
    locationNote: {
      user: r.one.user({ from: r.locationNote.userId, to: r.user.id }),
      location: r.one.location({
        from: r.locationNote.locationId,
        to: r.location.id,
        optional: false,
      }),
    },
    qr: {
      asset: r.one.asset({ from: r.qr.assetId, to: r.asset.id }),
      kit: r.one.kit({ from: r.qr.kitId, to: r.kit.id }),
      user: r.one.user({ from: r.qr.userId, to: r.user.id }),
      organization: r.one.organization({ from: r.qr.organizationId, to: r.organization.id }),
      batch: r.one.printBatch({ from: r.qr.batchId, to: r.printBatch.id }),
      scans: r.many.scan(),
    },
    barcode: {
      asset: r.one.asset({ from: r.barcode.assetId, to: r.asset.id, alias: 'AssetBarcodes' }),
      preferredByAsset: r.one.asset({
        from: r.barcode.id,
        to: r.asset.preferredBarcodeId,
        alias: 'AssetPreferredBarcode',
      }),
      scans: r.many.scan(),
      kit: r.one.kit({ from: r.barcode.kitId, to: r.kit.id }),
      organization: r.one.organization({
        from: r.barcode.organizationId,
        to: r.organization.id,
        optional: false,
      }),
    },
    printBatch: { qrCodes: r.many.qr() },
    reportFound: {
      asset: r.one.asset({ from: r.reportFound.assetId, to: r.asset.id }),
      kit: r.one.kit({ from: r.reportFound.kitId, to: r.kit.id }),
    },
    scan: {
      user: r.one.user({ from: r.scan.userId, to: r.user.id }),
      qr: r.one.qr({ from: r.scan.qrId, to: r.qr.id }),
      barcode: r.one.barcode({ from: r.scan.barcodeId, to: r.barcode.id }),
      asset: r.one.asset({ from: r.scan.assetId, to: r.asset.id }),
      kit: r.one.kit({ from: r.scan.kitId, to: r.kit.id }),
      organization: r.one.organization({
        from: r.scan.organizationId,
        to: r.organization.id,
        optional: false,
      }),
    },
    location: {
      image: r.one.image({ from: r.location.imageId, to: r.image.id }),
      user: r.one.user({ from: r.location.userId, to: r.user.id, optional: false }),
      organization: r.one.organization({
        from: r.location.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      assetLocations: r.many.assetLocation(),
      kits: r.many.kit(),
      notes: r.many.locationNote(),
      custodies: r.many.custody(),
      consumptionLogs: r.many.consumptionLog(),
      parent: r.one.location({
        from: r.location.parentId,
        to: r.location.id,
        alias: 'LocationHierarchy',
      }),
      children: r.many.location({ alias: 'LocationHierarchy' }),
    },
    role: {
      users: r.many.user({
        from: r.role.id.through(r.roleToUser.A),
        to: r.user.id.through(r.roleToUser.B),
      }),
    },
    teamMember: {
      organization: r.one.organization({
        from: r.teamMember.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      custodies: r.many.custody(),
      receivedInvites: r.many.invite(),
      user: r.one.user({ from: r.teamMember.userId, to: r.user.id }),
      kitCustodies: r.many.kitCustody(),
      bookings: r.many.booking({ alias: 'BookingToTeamMember' }),
      assetReminders: r.many.assetReminder({
        from: r.teamMember.id.through(r.assetReminderToTeamMember.B),
        to: r.assetReminder.id.through(r.assetReminderToTeamMember.A),
      }),
      consumptionLogs: r.many.consumptionLog(),
      teamMemberNotes: r.many.teamMemberNote(),
      alwaysNotifyBookingSettings: r.many.bookingSettings({
        from: r.teamMember.id.through(r.bookingSettingsAlwaysNotify.B),
        to: r.bookingSettings.id.through(r.bookingSettingsAlwaysNotify.A),
      }),
      bookingNotifications: r.many.booking({
        from: r.teamMember.id.through(r.bookingNotificationRecipients.B),
        to: r.booking.id.through(r.bookingNotificationRecipients.A),
      }),
    },
    custody: {
      custodian: r.one.teamMember({
        from: r.custody.teamMemberId,
        to: r.teamMember.id,
        optional: false,
      }),
      asset: r.one.asset({ from: r.custody.assetId, to: r.asset.id, optional: false }),
      kitCustody: r.one.kitCustody({ from: r.custody.kitCustodyId, to: r.kitCustody.id }),
      location: r.one.location({ from: r.custody.locationId, to: r.location.id }),
    },
    organization: {
      owner: r.one.user({
        from: r.organization.userId,
        to: r.user.id,
        optional: false,
        alias: 'OrganizationToUser',
      }),
      members: r.many.teamMember(),
      assets: r.many.asset(),
      assetKits: r.many.assetKit(),
      assetLocations: r.many.assetLocation(),
      scans: r.many.scan(),
      image: r.one.image({ from: r.organization.imageId, to: r.image.id, alias: 'orgImage' }),
      ssoDetails: r.one.ssoDetails({ from: r.organization.ssoDetailsId, to: r.ssoDetails.id }),
      workingHours: r.one.workingHours({
        from: r.organization.id,
        to: r.workingHours.organizationId,
      }),
      bookingSettings: r.one.bookingSettings({
        from: r.organization.id,
        to: r.bookingSettings.organizationId,
      }),
      locations: r.many.location(),
      categories: r.many.category(),
      tags: r.many.tag(),
      qrCodes: r.many.qr(),
      barcodes: r.many.barcode(),
      invites: r.many.invite(),
      userOrganizations: r.many.userOrganization(),
      customFields: r.many.customField(),
      images: r.many.image({ alias: 'owner' }),
      bookings: r.many.booking(),
      kits: r.many.kit(),
      assetModels: r.many.assetModel(),
      assetIndexSettings: r.many.assetIndexSettings(),
      assetReminders: r.many.assetReminder(),
      assetFilterPresets: r.many.assetFilterPreset(),
      auditSessions: r.many.auditSession(),
      auditImages: r.many.auditImage(),
      assetImports: r.many.assetImport(),
      linkPreviews: r.many.linkPreview(),
      pdfTemplates: r.many.pdfTemplate(),
      usersWithLastSelected: r.many.user({ alias: 'lastSelectedOrg' }),
      scimTokens: r.many.scimToken(),
      scimExternalIds: r.many.userScimExternalId(),
      roleChangeLogs: r.many.roleChangeLog(),
      teamMemberNotes: r.many.teamMemberNote(),
      activityEvents: r.many.activityEvent(),
    },
    userOrganization: {
      user: r.one.user({ from: r.userOrganization.userId, to: r.user.id, optional: false }),
      organization: r.one.organization({
        from: r.userOrganization.organizationId,
        to: r.organization.id,
        optional: false,
      }),
    },
    ssoDetails: { organizations: r.many.organization() },
    scimToken: {
      organization: r.one.organization({
        from: r.scimToken.organizationId,
        to: r.organization.id,
        optional: false,
      }),
    },
    userScimExternalId: {
      user: r.one.user({ from: r.userScimExternalId.userId, to: r.user.id, optional: false }),
      organization: r.one.organization({
        from: r.userScimExternalId.organizationId,
        to: r.organization.id,
        optional: false,
      }),
    },
    tier: {
      subscribers: r.many.user(),
      tierLimit: r.one.tierLimit({ from: r.tier.tierLimitId, to: r.tierLimit.id }),
    },
    tierLimit: { tier: r.one.tier({ from: r.tierLimit.id, to: r.tier.tierLimitId }) },
    customTierLimit: { user: r.one.user({ from: r.customTierLimit.userId, to: r.user.id }) },
    customField: {
      organization: r.one.organization({
        from: r.customField.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      createdBy: r.one.user({ from: r.customField.userId, to: r.user.id, optional: false }),
      categories: r.many.category({
        from: r.customField.id.through(r.categoryToCustomField.B),
        to: r.category.id.through(r.categoryToCustomField.A),
      }),
      assetCustomFieldsValues: r.many.assetCustomFieldValue(),
    },
    assetCustomFieldValue: {
      asset: r.one.asset({
        from: r.assetCustomFieldValue.assetId,
        to: r.asset.id,
        optional: false,
      }),
      customField: r.one.customField({
        from: r.assetCustomFieldValue.customFieldId,
        to: r.customField.id,
        optional: false,
      }),
    },
    assetModelCustomFieldDefault: {
      assetModel: r.one.assetModel({
        from: r.assetModelCustomFieldDefault.assetModelId,
        to: r.assetModel.id,
        optional: false,
      }),
    },
    invite: {
      inviter: r.one.user({
        from: r.invite.inviterId,
        to: r.user.id,
        optional: false,
        alias: 'inviter',
      }),
      organization: r.one.organization({
        from: r.invite.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      inviteeUser: r.one.user({ from: r.invite.inviteeUserId, to: r.user.id, alias: 'invitee' }),
      inviteeTeamMember: r.one.teamMember({
        from: r.invite.teamMemberId,
        to: r.teamMember.id,
        optional: false,
      }),
    },
    consumptionLog: {
      asset: r.one.asset({ from: r.consumptionLog.assetId, to: r.asset.id, optional: false }),
      performedBy: r.one.user({ from: r.consumptionLog.userId, to: r.user.id, optional: false }),
      booking: r.one.booking({ from: r.consumptionLog.bookingId, to: r.booking.id }),
      custodian: r.one.teamMember({ from: r.consumptionLog.custodianId, to: r.teamMember.id }),
      location: r.one.location({ from: r.consumptionLog.locationId, to: r.location.id }),
    },
    bookingAsset: {
      booking: r.one.booking({ from: r.bookingAsset.bookingId, to: r.booking.id, optional: false }),
      asset: r.one.asset({ from: r.bookingAsset.assetId, to: r.asset.id, optional: false }),
    },
    booking: {
      creator: r.one.user({
        from: r.booking.creatorId,
        to: r.user.id,
        optional: false,
        alias: 'creator',
      }),
      custodianUser: r.one.user({
        from: r.booking.custodianUserId,
        to: r.user.id,
        alias: 'custodian',
      }),
      custodianTeamMember: r.one.teamMember({
        from: r.booking.custodianTeamMemberId,
        to: r.teamMember.id,
        alias: 'BookingToTeamMember',
      }),
      organization: r.one.organization({
        from: r.booking.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      bookingAssets: r.many.bookingAsset(),
      consumptionLogs: r.many.consumptionLog(),
      notes: r.many.bookingNote(),
      modelRequests: r.many.bookingModelRequest(),
      tags: r.many.tag({
        from: r.booking.id.through(r.bookingToTag.A),
        to: r.tag.id.through(r.bookingToTag.B),
      }),
      partialCheckins: r.many.partialBookingCheckin(),
      partialCheckouts: r.many.partialBookingCheckout(),
      notificationRecipients: r.many.teamMember({
        from: r.booking.id.through(r.bookingNotificationRecipients.A),
        to: r.teamMember.id.through(r.bookingNotificationRecipients.B),
      }),
    },
    bookingSettings: {
      alwaysNotifyTeamMembers: r.many.teamMember({
        from: r.bookingSettings.id.through(r.bookingSettingsAlwaysNotify.A),
        to: r.teamMember.id.through(r.bookingSettingsAlwaysNotify.B),
      }),
      organization: r.one.organization({
        from: r.bookingSettings.organizationId,
        to: r.organization.id,
        optional: false,
      }),
    },
    partialBookingCheckin: {
      booking: r.one.booking({
        from: r.partialBookingCheckin.bookingId,
        to: r.booking.id,
        optional: false,
      }),
      checkedInBy: r.one.user({
        from: r.partialBookingCheckin.checkedInById,
        to: r.user.id,
        optional: false,
      }),
    },
    partialBookingCheckout: {
      booking: r.one.booking({
        from: r.partialBookingCheckout.bookingId,
        to: r.booking.id,
        optional: false,
      }),
      checkedOutBy: r.one.user({
        from: r.partialBookingCheckout.checkedOutById,
        to: r.user.id,
        optional: false,
      }),
    },
    mobileAuthCode: {
      user: r.one.user({ from: r.mobileAuthCode.userId, to: r.user.id, optional: false }),
    },
    pushToken: { user: r.one.user({ from: r.pushToken.userId, to: r.user.id, optional: false }) },
    notificationPreference: {
      user: r.one.user({ from: r.notificationPreference.userId, to: r.user.id, optional: false }),
    },
    kit: {
      assetKits: r.many.assetKit(),
      custody: r.one.kitCustody({ from: r.kit.id, to: r.kitCustody.kitId }),
      qrCodes: r.many.qr(),
      barcodes: r.many.barcode(),
      reports: r.many.reportFound(),
      scans: r.many.scan(),
      organization: r.one.organization({
        from: r.kit.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      createdBy: r.one.user({ from: r.kit.createdById, to: r.user.id, optional: false }),
      category: r.one.category({ from: r.kit.categoryId, to: r.category.id }),
      location: r.one.location({ from: r.kit.locationId, to: r.location.id }),
    },
    kitCustody: {
      custodian: r.one.teamMember({
        from: r.kitCustody.custodianId,
        to: r.teamMember.id,
        optional: false,
      }),
      kit: r.one.kit({ from: r.kitCustody.kitId, to: r.kit.id, optional: false }),
      inheritedCustody: r.many.custody(),
    },
    assetKit: {
      asset: r.one.asset({ from: r.assetKit.assetId, to: r.asset.id, optional: false }),
      kit: r.one.kit({ from: r.assetKit.kitId, to: r.kit.id, optional: false }),
      organization: r.one.organization({
        from: r.assetKit.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      assetLocations: r.many.assetLocation(),
    },
    assetLocation: {
      asset: r.one.asset({ from: r.assetLocation.assetId, to: r.asset.id, optional: false }),
      location: r.one.location({
        from: r.assetLocation.locationId,
        to: r.location.id,
        optional: false,
      }),
      organization: r.one.organization({
        from: r.assetLocation.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      assetKit: r.one.assetKit({ from: r.assetLocation.assetKitId, to: r.assetKit.id }),
    },
    assetReminder: {
      organization: r.one.organization({
        from: r.assetReminder.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      asset: r.one.asset({ from: r.assetReminder.assetId, to: r.asset.id, optional: false }),
      teamMembers: r.many.teamMember({
        from: r.assetReminder.id.through(r.assetReminderToTeamMember.A),
        to: r.teamMember.id.through(r.assetReminderToTeamMember.B),
      }),
      createdBy: r.one.user({ from: r.assetReminder.createdById, to: r.user.id, optional: false }),
    },
    workingHours: {
      organization: r.one.organization({
        from: r.workingHours.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      overrides: r.many.workingHoursOverride(),
    },
    workingHoursOverride: {
      workingHours: r.one.workingHours({
        from: r.workingHoursOverride.workingHoursId,
        to: r.workingHours.id,
        optional: false,
      }),
    },
    update: {
      userReads: r.many.userUpdateRead(),
      createdBy: r.one.user({ from: r.update.createdById, to: r.user.id, optional: false }),
    },
    userUpdateRead: {
      user: r.one.user({ from: r.userUpdateRead.userId, to: r.user.id, optional: false }),
      update: r.one.update({ from: r.userUpdateRead.updateId, to: r.update.id, optional: false }),
    },
    auditNote: {
      user: r.one.user({ from: r.auditNote.userId, to: r.user.id }),
      auditSession: r.one.auditSession({
        from: r.auditNote.auditSessionId,
        to: r.auditSession.id,
        optional: false,
      }),
      auditAsset: r.one.auditAsset({ from: r.auditNote.auditAssetId, to: r.auditAsset.id }),
    },
    teamMemberNote: {
      user: r.one.user({ from: r.teamMemberNote.userId, to: r.user.id }),
      teamMember: r.one.teamMember({
        from: r.teamMemberNote.teamMemberId,
        to: r.teamMember.id,
        optional: false,
      }),
      organization: r.one.organization({
        from: r.teamMemberNote.organizationId,
        to: r.organization.id,
        optional: false,
      }),
    },
    auditSession: {
      createdBy: r.one.user({ from: r.auditSession.createdById, to: r.user.id, optional: false }),
      organization: r.one.organization({
        from: r.auditSession.organizationId,
        to: r.organization.id,
        optional: false,
      }),
      assignments: r.many.auditAssignment(),
      assets: r.many.auditAsset(),
      scans: r.many.auditScan(),
      notes: r.many.auditNote(),
      images: r.many.auditImage(),
    },
    assetImport: {
      createdBy: r.one.user({ from: r.assetImport.createdById, to: r.user.id, optional: false }),
      organization: r.one.organization({
        from: r.assetImport.organizationId,
        to: r.organization.id,
        optional: false,
      }),
    },
    linkPreview: {
      organization: r.one.organization({
        from: r.linkPreview.organizationId,
        to: r.organization.id,
        optional: false,
      }),
    },
    auditAssignment: {
      auditSession: r.one.auditSession({
        from: r.auditAssignment.auditSessionId,
        to: r.auditSession.id,
        optional: false,
      }),
      user: r.one.user({ from: r.auditAssignment.userId, to: r.user.id, optional: false }),
    },
    auditAsset: {
      auditSession: r.one.auditSession({
        from: r.auditAsset.auditSessionId,
        to: r.auditSession.id,
        optional: false,
      }),
      asset: r.one.asset({ from: r.auditAsset.assetId, to: r.asset.id, optional: false }),
      scannedBy: r.one.user({ from: r.auditAsset.scannedById, to: r.user.id }),
      scans: r.many.auditScan(),
      images: r.many.auditImage(),
      notes: r.many.auditNote(),
    },
    auditScan: {
      auditSession: r.one.auditSession({
        from: r.auditScan.auditSessionId,
        to: r.auditSession.id,
        optional: false,
      }),
      auditAsset: r.one.auditAsset({ from: r.auditScan.auditAssetId, to: r.auditAsset.id }),
      asset: r.one.asset({ from: r.auditScan.assetId, to: r.asset.id }),
      scannedBy: r.one.user({ from: r.auditScan.scannedById, to: r.user.id }),
    },
    auditImage: {
      auditSession: r.one.auditSession({
        from: r.auditImage.auditSessionId,
        to: r.auditSession.id,
        optional: false,
      }),
      auditAsset: r.one.auditAsset({ from: r.auditImage.auditAssetId, to: r.auditAsset.id }),
      uploadedBy: r.one.user({ from: r.auditImage.uploadedById, to: r.user.id }),
      organization: r.one.organization({
        from: r.auditImage.organizationId,
        to: r.organization.id,
        optional: false,
      }),
    },
    roleChangeLog: {
      user: r.one.user({
        from: r.roleChangeLog.userId,
        to: r.user.id,
        optional: false,
        alias: 'roleChangeTarget',
      }),
      changedBy: r.one.user({
        from: r.roleChangeLog.changedById,
        to: r.user.id,
        optional: false,
        alias: 'roleChangeActor',
      }),
      organization: r.one.organization({
        from: r.roleChangeLog.organizationId,
        to: r.organization.id,
        optional: false,
      }),
    },
    activityEvent: {
      organization: r.one.organization({
        from: r.activityEvent.organizationId,
        to: r.organization.id,
        optional: false,
      }),
    },
    roleToUser: {
      role: r.one.role({ from: r.roleToUser.A, to: r.role.id, optional: false }),
      user: r.one.user({ from: r.roleToUser.B, to: r.user.id, optional: false }),
    },
    assetToTag: {
      asset: r.one.asset({ from: r.assetToTag.A, to: r.asset.id, optional: false }),
      tag: r.one.tag({ from: r.assetToTag.B, to: r.tag.id, optional: false }),
    },
    categoryToCustomField: {
      category: r.one.category({
        from: r.categoryToCustomField.A,
        to: r.category.id,
        optional: false,
      }),
      customField: r.one.customField({
        from: r.categoryToCustomField.B,
        to: r.customField.id,
        optional: false,
      }),
    },
    bookingToTag: {
      booking: r.one.booking({ from: r.bookingToTag.A, to: r.booking.id, optional: false }),
      tag: r.one.tag({ from: r.bookingToTag.B, to: r.tag.id, optional: false }),
    },
    assetReminderToTeamMember: {
      assetReminder: r.one.assetReminder({
        from: r.assetReminderToTeamMember.A,
        to: r.assetReminder.id,
        optional: false,
      }),
      teamMember: r.one.teamMember({
        from: r.assetReminderToTeamMember.B,
        to: r.teamMember.id,
        optional: false,
      }),
    },
    bookingSettingsAlwaysNotify: {
      bookingSettings: r.one.bookingSettings({
        from: r.bookingSettingsAlwaysNotify.A,
        to: r.bookingSettings.id,
        optional: false,
      }),
      teamMember: r.one.teamMember({
        from: r.bookingSettingsAlwaysNotify.B,
        to: r.teamMember.id,
        optional: false,
      }),
    },
    bookingNotificationRecipients: {
      booking: r.one.booking({
        from: r.bookingNotificationRecipients.A,
        to: r.booking.id,
        optional: false,
      }),
      teamMember: r.one.teamMember({
        from: r.bookingNotificationRecipients.B,
        to: r.teamMember.id,
        optional: false,
      }),
    },
  }),
)
