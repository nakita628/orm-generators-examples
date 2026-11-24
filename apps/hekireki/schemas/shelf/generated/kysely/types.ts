import type { ColumnType } from 'kysely'

export type Generated<T> =
  T extends ColumnType<infer S, infer I, infer U>
    ? ColumnType<S, I | undefined, U>
    : ColumnType<T, T | undefined, T>

export type Timestamp = ColumnType<Date, Date | string, Date | string>

export type AssetStatus = 'AVAILABLE' | 'IN_CUSTODY' | 'CHECKED_OUT'

/** * Determines how an asset is tracked throughout the system. Immutable after creation. */
export type AssetType = 'INDIVIDUAL' | 'QUANTITY_TRACKED'

/** * Controls what happens after checkout for quantity-tracked assets. */
export type ConsumptionType = 'ONE_WAY' | 'TWO_WAY'

/** * Classifies the type of quantity change in a ConsumptionLog entry. */
export type ConsumptionCategory =
  | 'CHECKOUT'
  | 'RETURN'
  | 'RESTOCK'
  | 'ADJUSTMENT'
  | 'LOSS'
  | 'CONSUME'
  | 'DAMAGE'
  | 'MOVE'
  | 'INITIAL'
  | 'OPENING_BALANCE'

/**
 * The kind of PDF a PdfTemplate lays out. Only the booking PDF has saved
 * layouts today; the check-in receipt and the audit receipt can join as new
 * kinds.
 */
export type PdfTemplateKind = 'BOOKING'

export type AssetIndexMode = 'SIMPLE' | 'ADVANCED'

export type TagUseFor = 'ASSET' | 'BOOKING'

export type NoteType = 'COMMENT' | 'UPDATE'

export type BarcodeType = 'Code128' | 'Code39' | 'DataMatrix' | 'ExternalQR' | 'EAN13'

export type ErrorCorrection = 'L' | 'M' | 'Q' | 'H'

/** The kind of code a `Scan` recorded. */
export type ScanCodeType = 'QR' | 'BARCODE' | 'SAM_ID'

/** The surface that recorded a `Scan`. */
export type ScanSource =
  | 'QR_LINK'
  | 'WEB_SCANNER'
  | 'WEB_DRAWER'
  | 'AUDIT'
  | 'COMPANION'
  | 'GPS_UPDATE'

export type Roles = 'USER' | 'ADMIN'

export type OrganizationType = 'PERSONAL' | 'TEAM'

export type QrIdDisplayPreference =
  | 'QR_ID'
  | 'SAM_ID'
  | 'Code128'
  | 'Code39'
  | 'DataMatrix'
  | 'ExternalQR'
  | 'EAN13'

export type DateFormatPreference =
  | 'DD_MM_YYYY'
  | 'MM_DD_YYYY'
  | 'YYYY_MM_DD'
  | 'MMM_DD_YYYY'
  | 'DD_MMM_YYYY'

export type TimeFormatPreference = 'H12' | 'H24'

export type WeekStartPreference = 'MONDAY' | 'SUNDAY' | 'SATURDAY'

export type OrganizationRoles = 'ADMIN' | 'BASE' | 'OWNER' | 'SELF_SERVICE' | 'MANAGER'

export type TierId = 'free' | 'tier_1' | 'tier_2' | 'custom'

export type CustomFieldType =
  | 'TEXT'
  | 'OPTION'
  | 'BOOLEAN'
  | 'DATE'
  | 'MULTILINE_TEXT'
  | 'AMOUNT'
  | 'NUMBER'

/**
 * The format of a TEXT custom field. Unlike the type, it can change after the
 * field exists.
 */
export type CustomFieldFormat = 'PLAIN' | 'LINK'

export type Currency =
  | 'AED'
  | 'AFN'
  | 'ALL'
  | 'AMD'
  | 'ANG'
  | 'AOA'
  | 'ARS'
  | 'AUD'
  | 'AWG'
  | 'AZN'
  | 'BAM'
  | 'BBD'
  | 'BDT'
  | 'BGN'
  | 'BHD'
  | 'BIF'
  | 'BMD'
  | 'BND'
  | 'BOB'
  | 'BRL'
  | 'BSD'
  | 'BTN'
  | 'BWP'
  | 'BYN'
  | 'BZD'
  | 'CAD'
  | 'CDF'
  | 'CHF'
  | 'CLP'
  | 'CNY'
  | 'COP'
  | 'CRC'
  | 'CUP'
  | 'CVE'
  | 'CZK'
  | 'DJF'
  | 'DKK'
  | 'DOP'
  | 'DZD'
  | 'EGP'
  | 'ERN'
  | 'ETB'
  | 'EUR'
  | 'FJD'
  | 'FKP'
  | 'GBP'
  | 'GEL'
  | 'GHS'
  | 'GIP'
  | 'GMD'
  | 'GNF'
  | 'GTQ'
  | 'GYD'
  | 'HKD'
  | 'HNL'
  | 'HTG'
  | 'HUF'
  | 'IDR'
  | 'ILS'
  | 'INR'
  | 'IQD'
  | 'IRR'
  | 'ISK'
  | 'JMD'
  | 'JOD'
  | 'JPY'
  | 'KES'
  | 'KGS'
  | 'KHR'
  | 'KMF'
  | 'KPW'
  | 'KRW'
  | 'KWD'
  | 'KYD'
  | 'KZT'
  | 'LAK'
  | 'LBP'
  | 'LKR'
  | 'LRD'
  | 'LSL'
  | 'LYD'
  | 'MAD'
  | 'MDL'
  | 'MGA'
  | 'MKD'
  | 'MMK'
  | 'MNT'
  | 'MOP'
  | 'MRU'
  | 'MUR'
  | 'MVR'
  | 'MWK'
  | 'MXN'
  | 'MYR'
  | 'MZN'
  | 'NAD'
  | 'NGN'
  | 'NIO'
  | 'NOK'
  | 'NPR'
  | 'NZD'
  | 'OMR'
  | 'PAB'
  | 'PEN'
  | 'PGK'
  | 'PHP'
  | 'PKR'
  | 'PLN'
  | 'PYG'
  | 'QAR'
  | 'RON'
  | 'RSD'
  | 'RUB'
  | 'RWF'
  | 'SAR'
  | 'SBD'
  | 'SCR'
  | 'SDG'
  | 'SEK'
  | 'SGD'
  | 'SHP'
  | 'SLE'
  | 'SOS'
  | 'SRD'
  | 'SSP'
  | 'STN'
  | 'SVC'
  | 'SYP'
  | 'SZL'
  | 'THB'
  | 'TJS'
  | 'TMT'
  | 'TND'
  | 'TOP'
  | 'TRY'
  | 'TTD'
  | 'TWD'
  | 'TZS'
  | 'UAH'
  | 'UGX'
  | 'USD'
  | 'UYU'
  | 'UZS'
  | 'VES'
  | 'VND'
  | 'VUV'
  | 'WST'
  | 'XAF'
  | 'XCD'
  | 'XOF'
  | 'XPF'
  | 'YER'
  | 'ZAR'
  | 'ZMW'
  | 'ZWL'

export type InviteStatuses = 'PENDING' | 'ACCEPTED' | 'REJECTED' | 'INVALIDATED'

export type BookingStatus =
  | 'DRAFT'
  | 'RESERVED'
  | 'ONGOING'
  | 'OVERDUE'
  | 'COMPLETE'
  | 'ARCHIVED'
  | 'CANCELLED'

/** The operating system a companion-app push token was issued for. */
export type DevicePlatform = 'IOS' | 'ANDROID'

export type KitStatus = 'AVAILABLE' | 'IN_CUSTODY' | 'CHECKED_OUT'

export type ReminderRecurrenceUnit = 'DAY' | 'WEEK' | 'MONTH' | 'YEAR'

export type UpdateStatus = 'DRAFT' | 'PUBLISHED'

export type AuditStatus = 'PENDING' | 'ACTIVE' | 'COMPLETED' | 'CANCELLED' | 'ARCHIVED'

export type AuditAssetStatus = 'PENDING' | 'FOUND' | 'MISSING' | 'UNEXPECTED'

export type AuditAssignmentRole = 'LEAD' | 'PARTICIPANT'

/** Lifecycle of a background asset import. */
export type AssetImportStatus = 'DRAFT' | 'QUEUED' | 'RUNNING' | 'COMPLETED' | 'FAILED'

/** Which step of the background import runs next (or failed). */
export type AssetImportPhase = 'PREPARE' | 'ROWS' | 'FINALIZE'

/** How the last fetch of a link's page went. */
export type LinkPreviewOutcome = 'TITLE' | 'NO_TITLE' | 'SIGN_IN' | 'BLOCKED' | 'UNREACHABLE'

/**
 * How an audit photo reached Shelf, as reported by the client. The phone app
 * reports CAMERA only for its own camera, which cannot hand back an old
 * picture, so it is the strongest evidence short of a modified client; a
 * library photo's date is only what its file claims; the web app cannot tell
 * a camera shot from a file; PHONE is a phone app version that does not say
 * which.
 */
export type AuditImageCaptureSource = 'CAMERA' | 'LIBRARY' | 'WEB_UPLOAD' | 'PHONE'

/** How a membership's role came to change. */
export type RoleChangeSource = 'MANUAL' | 'SSO'

export type ActivityEntity =
  | 'ASSET'
  | 'BOOKING'
  | 'AUDIT'
  | 'KIT'
  | 'LOCATION'
  | 'TEAM_MEMBER'
  | 'CUSTODY'
  | 'USER'
  | 'ORGANIZATION'

export type ActivityAction =
  | 'ASSET_CREATED'
  | 'ASSET_NAME_CHANGED'
  | 'ASSET_DESCRIPTION_CHANGED'
  | 'ASSET_CATEGORY_CHANGED'
  | 'ASSET_MODEL_CHANGED'
  | 'ASSET_KIT_CHANGED'
  | 'ASSET_LOCATION_CHANGED'
  | 'ASSET_TAGS_CHANGED'
  | 'ASSET_STATUS_CHANGED'
  | 'ASSET_VALUATION_CHANGED'
  | 'ASSET_QUANTITY_CHANGED'
  | 'ASSET_MIN_QUANTITY_CHANGED'
  | 'ASSET_CUSTOM_FIELD_CHANGED'
  | 'ASSET_PREFERRED_BARCODE_CHANGED'
  | 'ASSET_AVAILABLE_TO_BOOK_CHANGED'
  | 'ASSET_ARCHIVED'
  | 'ASSET_UNARCHIVED'
  | 'ASSET_DELETED'
  | 'CUSTODY_ASSIGNED'
  | 'CUSTODY_RELEASED'
  | 'BOOKING_CREATED'
  | 'BOOKING_STATUS_CHANGED'
  | 'BOOKING_DATES_CHANGED'
  | 'BOOKING_ASSETS_ADDED'
  | 'BOOKING_ASSETS_REMOVED'
  | 'BOOKING_CHECKED_OUT'
  | 'BOOKING_CHECKED_IN'
  | 'BOOKING_PARTIAL_CHECKIN'
  | 'BOOKING_PARTIAL_CHECKOUT'
  | 'BOOKING_CANCELLED'
  | 'BOOKING_ARCHIVED'
  | 'BOOKING_APPROVED'
  | 'BOOKING_MODEL_REQUESTED'
  | 'BOOKING_MODEL_REQUEST_CHANGED'
  | 'BOOKING_MODEL_REQUEST_REMOVED'
  | 'BOOKING_MODEL_REQUEST_FULFILLED'
  | 'AUDIT_CREATED'
  | 'AUDIT_STARTED'
  | 'AUDIT_ASSETS_ADDED'
  | 'AUDIT_ASSETS_REMOVED'
  | 'AUDIT_ASSET_SCANNED'
  | 'AUDIT_ASSET_SCAN_REMOVED'
  | 'AUDIT_DUE_DATE_CHANGED'
  | 'AUDIT_ASSIGNEE_ADDED'
  | 'AUDIT_ASSIGNEE_REMOVED'
  | 'AUDIT_UPDATED'
  | 'AUDIT_COMPLETED'
  | 'AUDIT_CANCELLED'
  | 'AUDIT_ARCHIVED'
  | 'LOCATION_CREATED'
  | 'LOCATION_UPDATED'
  | 'KIT_CREATED'
  | 'KIT_UPDATED'
  | 'ORGANIZATION_QR_ID_DISPLAY_PREFERENCE_CHANGED'

export interface Image {
  id: string
  contentType: string
  altText: string | null
  blob: Buffer
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  ownerOrgId: string
  userId: string
}

export interface User {
  id: string
  email: string
  username: string
  firstName: string | null
  lastName: string | null
  displayName: string | null
  profilePicture: string | null
  usedFreeTrial: Generated<boolean>
  onboarded: Generated<boolean>
  customerId: string | null
  sso: Generated<boolean>
  createdWithInvite: Generated<boolean>
  skipSubscriptionCheck: Generated<boolean>
  hasUnpaidInvoice: Generated<boolean>
  warnForNoPaymentMethod: Generated<boolean>
  lastSelectedOrganizationId: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  deletedAt: Timestamp | null
  lastMobileActiveAt: Timestamp | null
  dateFormat: DateFormatPreference | null
  timeFormat: TimeFormatPreference | null
  weekStart: WeekStartPreference | null
  timeZone: string | null
  tierId: Generated<TierId>
  referralSource: string | null
}

export interface UserContact {
  id: string
  phone: string | null
  street: string | null
  city: string | null
  stateProvince: string | null
  zipPostalCode: string | null
  countryRegion: string | null
  userId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface UserBusinessIntel {
  id: string
  howDidYouHearAboutUs: string | null
  jobTitle: string | null
  teamSize: string | null
  companyName: string | null
  primaryUseCase: string | null
  currentSolution: string | null
  timeline: string | null
  signupPlan: string | null
  signupTrial: boolean | null
  utmSource: string | null
  utmMedium: string | null
  utmCampaign: string | null
  utmContent: string | null
  userId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface Asset {
  id: string
  title: string
  description: string | null
  mainImage: string | null
  thumbnailImage: string | null
  mainImageExpiration: Timestamp | null
  status: Generated<AssetStatus>
  value: number | null
  availableToBook: Generated<boolean>
  sequentialId: string | null
  archivedAt: Timestamp | null
  type: Generated<AssetType>
  assetModelId: string | null
  quantity: number | null
  minQuantity: number | null
  consumptionType: ConsumptionType | null
  unitOfMeasure: string | null
  lowStockNotifiedAt: Timestamp | null
  stockLedgerStartedAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  userId: string
  organizationId: string
  categoryId: string | null
  preferredBarcodeId: string | null
}

export interface AssetFilterPreset {
  id: string
  organizationId: string
  ownerId: string
  name: string
  query: string
  starred: Generated<boolean>
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface AssetIndexSettings {
  id: string
  userId: string
  organizationId: string
  mode: Generated<AssetIndexMode>
  columns: Generated<unknown>
  freezeColumn: Generated<boolean>
  showAssetImage: Generated<boolean>
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

/**
 * A saved layout for a generated PDF, kept per team workspace under a name:
 * which columns, rows and header lines print, and in what order the rows
 * come. A row is a layout, not a generated file; the person generating a PDF
 * picks a template by name.
 */
export interface PdfTemplate {
  id: string
  name: string
  kind: Generated<PdfTemplateKind>
  /**
   * The layout's switches. Read through `parseBookingTemplateConfig`,
   * which fills any key a stored row lacks from the Checklist template's
   * values.
   */
  config: Generated<unknown>
  /**
   * The template the PDF dialog offers first. One per workspace and kind,
   * enforced by a partial unique index in raw SQL (migration
   * 20261103120000_add_pdf_templates). Names are also unique per workspace
   * and kind ignoring case, through an expression index in the same
   * migration.
   */
  isDefault: Generated<boolean>
  organizationId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface Category {
  id: string
  name: string
  description: string | null
  color: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  userId: string
  organizationId: string
}

/**
 * *
 * * A template/grouping entity for assets.
 * * Groups individual assets of the same type (e.g., "Dell Latitude 5550")
 * * and provides default field values when creating new assets from the model.
 */
export interface AssetModel {
  id: string
  name: string
  description: string | null
  /**
   * Public URL (bucket `files`) of the model's cover image. Rendered by any
   * asset of this model that has no image of its own. Public — and therefore
   * never-expiring and cacheable — like Location.imageUrl; see
   * resolveAssetImage() for the cascade.
   */
  image: string | null
  /**
   * Public URL of the 108px thumbnail generated alongside `image`. Used at
   * list sizes so inheriting assets never render the full-size original.
   */
  thumbnailImage: string | null
  defaultCategoryId: string | null
  defaultValuation: number | null
  organizationId: string
  userId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

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
export interface BookingModelRequest {
  id: string
  bookingId: string
  assetModelId: string
  /**
   * Total units originally reserved. Mutated only by explicit
   * operator edits (upsert) — never decremented on materialisation.
   * This preserves the audit trail of "how many were planned".
   */
  quantity: number
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
  fulfilledQuantity: Generated<number>
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
  fulfilledAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface Tag {
  id: string
  name: string
  description: string | null
  color: string | null
  useFor: Generated<ColumnType<string, TagUseFor[], TagUseFor[]>>
  userId: string
  organizationId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface Note {
  id: string
  content: string
  type: Generated<NoteType>
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  userId: string | null
  assetId: string
}

export interface BookingNote {
  id: string
  content: string
  type: Generated<NoteType>
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  userId: string | null
  bookingId: string
}

export interface LocationNote {
  id: string
  content: string
  type: Generated<NoteType>
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  userId: string | null
  locationId: string
}

export interface Qr {
  id: string
  version: Generated<number>
  errorCorrection: Generated<ErrorCorrection>
  assetId: string | null
  kitId: string | null
  userId: string | null
  organizationId: string | null
  batchId: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface Barcode {
  id: string
  value: string
  type: Generated<BarcodeType>
  assetId: string | null
  kitId: string | null
  organizationId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface PrintBatch {
  id: string
  name: string
  printed: Generated<boolean>
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface ReportFound {
  id: string
  email: string
  content: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  assetId: string | null
  kitId: string | null
}

export interface Scan {
  id: string
  latitude: string | null
  longitude: string | null
  userAgent: string | null
  userId: string | null
  qrId: string | null
  rawQrId: string | null
  codeType: Generated<ScanCodeType>
  code: string | null
  barcodeId: string | null
  assetId: string | null
  kitId: string | null
  organizationId: string
  source: ScanSource | null
  manuallyGenerated: Generated<boolean>
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface Location {
  id: string
  name: string
  description: string | null
  address: string | null
  latitude: number | null
  longitude: number | null
  imageUrl: string | null
  thumbnailUrl: string | null
  imageId: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  userId: string
  organizationId: string
  parentId: string | null
}

export interface Role {
  id: string
  name: Generated<Roles>
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface TeamMember {
  id: string
  name: string
  organizationId: string
  userId: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  deletedAt: Timestamp | null
}

export interface Custody {
  id: string
  teamMemberId: string
  assetId: string
  kitCustodyId: string | null
  quantity: Generated<number>
  locationId: string | null
  sourceUnknown: Generated<boolean>
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface Organization {
  id: string
  name: Generated<string>
  type: Generated<OrganizationType>
  userId: string
  currency: Generated<Currency>
  imageId: string | null
  enabledSso: Generated<boolean>
  ssoDetailsId: string | null
  selfServiceCanSeeCustody: Generated<boolean>
  selfServiceCanSeeBookings: Generated<boolean>
  baseUserCanSeeCustody: Generated<boolean>
  baseUserCanSeeBookings: Generated<boolean>
  managersCanEditAssets: Generated<boolean>
  barcodesEnabled: Generated<boolean>
  barcodesEnabledAt: Timestamp | null
  usedBarcodeTrial: Generated<boolean>
  auditsEnabled: Generated<boolean>
  auditsEnabledAt: Timestamp | null
  usedAuditTrial: Generated<boolean>
  workspaceDisabled: Generated<boolean>
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  hasSequentialIdsMigrated: Generated<boolean>
  qrIdDisplayPreference: Generated<QrIdDisplayPreference>
  showShelfBranding: Generated<boolean>
  showQrCodesOnPdfs: Generated<boolean>
  customEmailFooter: string | null
}

export interface UserOrganization {
  id: string
  userId: string
  organizationId: string
  roles: ColumnType<string, OrganizationRoles[], OrganizationRoles[]>
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  /**
   * Secret, unguessable token for this member's iCal calendar-feed subscription
   * in this workspace. Null until the user generates a feed; regenerating it
   * rotates the value (invalidating any previously-shared subscription URL).
   * `@unique` doubles as the lookup index for the public feed route.
   */
  calendarTokenId: string | null
}

export interface SsoDetails {
  id: string
  domain: string
  baseUserGroupId: string | null
  selfServiceGroupId: string | null
  adminGroupId: string | null
  managerGroupId: string | null
  requireSsoLogin: Generated<boolean>
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface ScimToken {
  id: string
  tokenHash: string
  label: string
  organizationId: string
  createdById: string
  createdAt: Generated<Timestamp>
  lastUsedAt: Timestamp | null
}

/**
 * *
 * * Org-scoped SCIM external ID linking a Shelf user to an identity-provider
 * * object within a specific organization.  A single Shelf user can appear in
 * * multiple organizations, each managed by a different identity provider, so the
 * * external ID must be scoped to the (user, organization) pair.
 */
export interface UserScimExternalId {
  id: string
  userId: string
  organizationId: string
  scimExternalId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface Tier {
  id: TierId
  name: string
  tierLimitId: TierId | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface TierLimit {
  id: TierId
  canImportAssets: Generated<boolean>
  canExportAssets: Generated<boolean>
  canImportNRM: Generated<boolean>
  canHideShelfBranding: Generated<boolean>
  canUseReports: Generated<boolean>
  canUseRecurringReminders: Generated<boolean>
  canUsePushNotifications: Generated<boolean>
  maxCustomFields: Generated<number>
  maxOrganizations: Generated<number>
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface CustomTierLimit {
  id: string
  userId: string | null
  canImportAssets: Generated<boolean>
  canExportAssets: Generated<boolean>
  canImportNRM: Generated<boolean>
  canHideShelfBranding: Generated<boolean>
  canUseReports: Generated<boolean>
  canUseRecurringReminders: Generated<boolean>
  canUsePushNotifications: Generated<boolean>
  maxCustomFields: Generated<number>
  maxOrganizations: Generated<number>
  isEnterprise: Generated<boolean>
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface CustomField {
  id: string
  name: string
  helpText: string | null
  required: Generated<boolean>
  active: Generated<boolean>
  type: Generated<CustomFieldType>
  options: string[]
  /**
   * How a TEXT field's values are read. LINK requires a web link in every
   * value; any other type stays PLAIN (enforced by a CHECK constraint).
   */
  format: Generated<CustomFieldFormat>
  organizationId: string
  userId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  deletedAt: Timestamp | null
}

export interface AssetCustomFieldValue {
  id: string
  value: unknown
  assetId: string
  customFieldId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

/**
 * A custom field value an AssetModel gives the assets that join it.
 *
 * `value` has the exact shape of `AssetCustomFieldValue.value`, so filling an
 * asset is a plain copy. Rows cascade away with their model or field.
 */
export interface AssetModelCustomFieldDefault {
  id: string
  value: unknown
  assetModelId: string
  customFieldId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface Invite {
  id: string
  inviterId: string
  organizationId: string
  inviteeUserId: string | null
  teamMemberId: string
  inviteeEmail: string
  status: Generated<InviteStatuses>
  inviteCode: string
  roles: ColumnType<string, OrganizationRoles[], OrganizationRoles[]>
  inviteMessage: string | null
  expiresAt: Timestamp
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface Announcement {
  id: string
  name: string
  content: string
  link: string | null
  linkText: string | null
  published: Generated<boolean>
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

/**
 * *
 * * Audit trail for all quantity changes on quantity-tracked assets.
 * * Every event records who performed it, who received/returned items,
 * * which booking it relates to, and the category of change.
 * * Quantity is always positive — direction is determined by category.
 */
export interface ConsumptionLog {
  id: string
  assetId: string
  category: ConsumptionCategory
  quantity: number
  note: string | null
  userId: string
  bookingId: string | null
  custodianId: string | null
  bookingAssetId: string | null
  locationId: string | null
  stockChange: number | null
  createdAt: Generated<Timestamp>
}

/**
 * *
 * * Explicit pivot table between Booking and Asset.
 * * Replaces the implicit M2M (_AssetToBooking) via a rename migration in Phase 3.
 */
export interface BookingAsset {
  id: string
  bookingId: string
  assetId: string
  quantity: Generated<number>
  bookingModelRequestId: string | null
  assetKitId: string | null
  sourceKitId: string | null
  checkedOutAt: Timestamp | null
  checkedOutById: string | null
  checkedInAt: Timestamp | null
  checkedInById: string | null
  checkedOutQuantity: Generated<number>
  sourceLocationId: string | null
}

export interface Booking {
  id: string
  name: string
  status: Generated<BookingStatus>
  description: Generated<string | null>
  activeSchedulerReference: string | null
  creatorId: string
  custodianUserId: string | null
  custodianTeamMemberId: string | null
  organizationId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  from: Timestamp
  to: Timestamp
  originalFrom: Timestamp | null
  originalTo: Timestamp | null
  autoArchivedAt: Timestamp | null
  archivedWithoutCheckin: Generated<boolean>
  cancellationReason: string | null
  approvedAt: Timestamp | null
  approvedById: string | null
}

export interface BookingSettings {
  id: string
  bufferStartTime: Generated<number>
  tagsRequired: Generated<boolean>
  maxBookingLength: number | null
  maxBookingLengthSkipClosedDays: Generated<boolean>
  autoArchiveBookings: Generated<boolean>
  autoArchiveDays: Generated<number>
  autoArchiveExpiredReservations: Generated<boolean>
  requireApprovalForBase: Generated<boolean>
  requireApprovalForSelfService: Generated<boolean>
  requireExplicitCheckinForAdmin: Generated<boolean>
  requireExplicitCheckinForSelfService: Generated<boolean>
  requireExplicitCheckoutForAdmin: Generated<boolean>
  requireExplicitCheckoutForSelfService: Generated<boolean>
  countKitsAsSingleUnit: Generated<boolean>
  notifyBookingCreator: Generated<boolean>
  notifyAdminsOnNewBooking: Generated<boolean>
  organizationId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

/**
 * A progressive check-in SESSION, mirroring PartialBookingCheckout.
 *
 * `assetIds` lists only FULLY reconciled assets: INDIVIDUAL always, and
 * QUANTITY_TRACKED only once its remaining units hit zero (partial returns
 * live in `ConsumptionLog`). Completing a booking outright writes no row here,
 * so the same warning applies — absence is not evidence of anything.
 * `BookingAsset.checkedInAt` is the per-slice fact.
 */
export interface PartialBookingCheckin {
  id: string
  assetIds: string[]
  checkinCount: number
  checkinTimestamp: Generated<Timestamp>
  bookingId: string
  checkedInById: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

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
export interface PartialBookingCheckout {
  id: string
  assetIds: string[]
  quantities: Generated<number[]>
  bookingAssetIds: Generated<string[]>
  checkoutCount: number
  checkoutTimestamp: Generated<Timestamp>
  bookingId: string
  checkedOutById: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

/**
 * Native-app SSO web-delegated login — single-use session-handoff codes.
 * @see apps/webapp/app/routes/_auth+/oauth.callback.mobile.tsx
 */
export interface MobileAuthCode {
  id: string
  userId: string
  codeHash: string
  codeChallenge: string | null
  sessionCiphertext: string | null
  consumedAt: Timestamp | null
  expiresAt: Timestamp
  createdAt: Generated<Timestamp>
}

/**
 * One Expo push token per installed companion app. A token identifies a
 * device, not a person: when someone else signs in on the same phone, the
 * row moves to them. Pushes go only to rows where `disabledAt` is null.
 * @see apps/webapp/app/modules/push/service.server.ts
 */
export interface PushToken {
  id: string
  token: string
  platform: DevicePlatform
  deviceName: string | null
  appVersion: string | null
  lastSeenAt: Timestamp
  disabledAt: Timestamp | null
  userId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

/**
 * A person's push notification switches. A missing row means every switch
 * is on, so no backfill is needed. The three category columns are named
 * after `PushCategory` in apps/webapp/app/modules/push/types.ts.
 */
export interface NotificationPreference {
  userId: string
  pushEnabled: Generated<boolean>
  bookings: Generated<boolean>
  audits: Generated<boolean>
  stockAndReminders: Generated<boolean>
  updatedAt: Timestamp
}

export interface Kit {
  id: string
  name: string
  description: string | null
  status: Generated<KitStatus>
  image: string | null
  imageExpiration: Timestamp | null
  organizationId: string
  createdById: string
  categoryId: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  locationId: string | null
}

export interface KitCustody {
  id: string
  custodianId: string
  kitId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

/**
 * Pivot linking an `Asset` to the `Kit` it belongs to.
 *
 * `@@unique([assetId])` enforces "at most one kit per asset" today.
 * The structure already allows multi-row-per-asset; if multi-kit
 * allocation lands later, drop the `assetId` unique constraint and
 * add a type-aware single-row trigger + sum-within-total CONSTRAINT
 * TRIGGER (parallel to the Custody / BookingAsset pivots).
 */
export interface AssetKit {
  id: string
  assetId: string
  kitId: string
  organizationId: string
  quantity: Generated<number>
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface AssetLocation {
  id: string
  assetId: string
  locationId: string
  organizationId: string
  quantity: Generated<number>
  assetKitId: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface AssetReminder {
  id: string
  name: string
  message: string
  alertDateTime: Timestamp
  activeSchedulerReference: string | null
  recurrenceUnit: ReminderRecurrenceUnit | null
  recurrenceInterval: number | null
  recurrenceTimezone: string | null
  recurrenceEndsAt: Timestamp | null
  organizationId: string
  assetId: string
  createdById: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface WorkingHours {
  id: string
  enabled: Generated<boolean>
  weeklySchedule: Generated<unknown>
  organizationId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface WorkingHoursOverride {
  id: string
  date: Timestamp
  isOpen: Generated<boolean>
  openTime: string | null
  closeTime: string | null
  reason: string | null
  workingHoursId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface Update {
  id: string
  title: string
  content: string
  url: string | null
  imageUrl: string | null
  publishDate: Timestamp
  status: Generated<UpdateStatus>
  targetRoles: ColumnType<string, OrganizationRoles[], OrganizationRoles[]>
  clickCount: Generated<number>
  viewCount: Generated<number>
  createdById: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface UserUpdateRead {
  id: string
  userId: string
  updateId: string
  readAt: Generated<Timestamp>
}

export interface AuditNote {
  id: string
  content: string
  type: Generated<NoteType>
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  userId: string | null
  auditSessionId: string
  auditAssetId: string | null
}

export interface TeamMemberNote {
  id: string
  content: string
  type: Generated<NoteType>
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  userId: string | null
  teamMemberId: string
  organizationId: string
}

export interface AuditSession {
  id: string
  name: string
  description: string | null
  targetId: string | null
  status: Generated<AuditStatus>
  scopeMeta: unknown
  expectedAssetCount: Generated<number>
  foundAssetCount: Generated<number>
  missingAssetCount: Generated<number>
  unexpectedAssetCount: Generated<number>
  startedAt: Timestamp | null
  dueDate: Timestamp | null
  completedAt: Timestamp | null
  cancelledAt: Timestamp | null
  activeSchedulerReference: string | null
  createdById: string
  organizationId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

/**
 * One CSV content import. Holds the parsed rows from preview until the import
 * completes, so the job imports exactly the file the user previewed.
 * A partial unique index on (organizationId) WHERE status IN (QUEUED, RUNNING) allows one
 * active import per workspace. Prisma cannot express it; it lives in migration
 * 20261015120100_asset_import_one_active_per_workspace.
 */
export interface AssetImport {
  id: string
  status: Generated<AssetImportStatus>
  phase: Generated<AssetImportPhase>
  fileName: string | null
  totalRows: number
  /** Rows committed so far. Advanced in the same transaction as each chunk. */
  processedRows: Generated<number>
  /** Parsed rows (`StoredImportRow[]`). Cleared when the import completes. */
  rows: unknown
  /** Name to id maps written by the PREPARE step (`ResolvedImportEntities`). */
  resolved: unknown
  canUseBarcodes: Generated<boolean>
  /** User-facing reason the import failed. Null unless status is FAILED. */
  errorMessage: string | null
  /** Rows imported without a conflicting piece (`ImportWarning[]`), capped at 500. */
  warnings: unknown
  /** Exact number of warnings, including any beyond the stored cap. */
  warningCount: Generated<number>
  /** A step holds the import until this time. Null when no step is running. */
  lockedUntil: Timestamp | null
  confirmedAt: Timestamp | null
  startedAt: Timestamp | null
  completedAt: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  createdById: string
  organizationId: string
}

/**
 * What Shelf knows about one web link used in a workspace's custom fields,
 * keyed by its `LinkKey` from @shelf/links. Two writers on disjoint columns:
 * the link preview job (fetch columns) and people (customTitle).
 */
export interface LinkPreview {
  id: string
  /**
   * `LinkKey` from @shelf/links, never a raw field value. At most 2048
   * characters (CHECK constraint in the migration).
   */
  url: string
  pageTitle: string | null
  /** Null until the first fetch finishes. */
  outcome: LinkPreviewOutcome | null
  attempts: Generated<number>
  fetchedAt: Timestamp | null
  /**
   * When the link may be fetched next; null means never again. A claim moves
   * it a few minutes ahead, which is also the lease on a fetch in flight.
   */
  nextFetchAt: Generated<Timestamp | null>
  /** The name a person gave the link in this workspace; wins over pageTitle. */
  customTitle: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  organizationId: string
}

export interface AuditAssignment {
  id: string
  auditSessionId: string
  userId: string
  role: AuditAssignmentRole | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface AuditAsset {
  id: string
  auditSessionId: string
  assetId: string
  expected: Generated<boolean>
  status: Generated<AuditAssetStatus>
  scannedById: string | null
  scannedAt: Timestamp | null
  metadata: unknown
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
  expectedQuantity: number | null
  /**
   * Units the auditor counted. Null until a quantity-tracked asset is
   * counted, and always null for INDIVIDUAL assets. A FOUND row with a count
   * below `expectedQuantity` is a shortfall: the status stays FOUND (the pool
   * was located), the gap is this number.
   *
   * Both columns are non-negative, enforced by CHECK constraints in raw SQL
   * (migration `20261104120000_add_audit_asset_quantities`).
   */
  foundQuantity: number | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface AuditScan {
  id: string
  auditSessionId: string
  auditAssetId: string | null
  assetId: string | null
  scannedById: string | null
  code: string | null
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
  assetTitle: string | null
  wasExpected: boolean | null
  metadata: unknown
  scannedAt: Generated<Timestamp>
  createdAt: Generated<Timestamp>
}

export interface AuditImage {
  id: string
  imageUrl: string
  thumbnailUrl: string | null
  description: string | null
  auditSessionId: string
  auditAssetId: string | null
  uploadedById: string | null
  organizationId: string
  /**
   * When the photo was taken, from its EXIF data (or the phone app's copy of
   * it). Null when the photo carried no usable date. EXIF can be edited, so
   * this is the best signal available, not proof.
   */
  capturedAt: Timestamp | null
  /**
   * How the photo reached Shelf. Null only for photos uploaded before this
   * field existed; a phone upload with no or an unknown source is PHONE.
   */
  captureSource: AuditImageCaptureSource | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface RoleChangeLog {
  id: string
  previousRole: OrganizationRoles
  newRole: OrganizationRoles
  /** How the change was made. SSO rows name the member themself as `changedBy`. */
  source: Generated<RoleChangeSource>
  createdAt: Generated<Timestamp>
  userId: string
  changedById: string
  organizationId: string
}

export interface ActivityEvent {
  id: string
  organizationId: string
  occurredAt: Generated<Timestamp>
  actorUserId: string | null
  actorSnapshot: unknown
  action: ActivityAction
  entityType: ActivityEntity
  entityId: string
  assetId: string | null
  bookingId: string | null
  auditSessionId: string | null
  auditAssetId: string | null
  kitId: string | null
  locationId: string | null
  teamMemberId: string | null
  targetUserId: string | null
  field: string | null
  fromValue: unknown
  toValue: unknown
  meta: unknown
}

export interface RoleToUser {
  A: string
  B: string
}

export interface AssetToTag {
  A: string
  B: string
}

export interface CategoryToCustomField {
  A: string
  B: string
}

export interface BookingToTag {
  A: string
  B: string
}

export interface AssetReminderToTeamMember {
  A: string
  B: string
}

export interface BookingSettingsAlwaysNotify {
  A: string
  B: string
}

export interface BookingNotificationRecipients {
  A: string
  B: string
}

export interface DB {
  Image: Image
  User: User
  UserContact: UserContact
  UserBusinessIntel: UserBusinessIntel
  Asset: Asset
  AssetFilterPreset: AssetFilterPreset
  AssetIndexSettings: AssetIndexSettings
  PdfTemplate: PdfTemplate
  Category: Category
  AssetModel: AssetModel
  BookingModelRequest: BookingModelRequest
  Tag: Tag
  Note: Note
  BookingNote: BookingNote
  LocationNote: LocationNote
  Qr: Qr
  Barcode: Barcode
  PrintBatch: PrintBatch
  ReportFound: ReportFound
  Scan: Scan
  Location: Location
  Role: Role
  TeamMember: TeamMember
  Custody: Custody
  Organization: Organization
  UserOrganization: UserOrganization
  SsoDetails: SsoDetails
  ScimToken: ScimToken
  UserScimExternalId: UserScimExternalId
  Tier: Tier
  TierLimit: TierLimit
  CustomTierLimit: CustomTierLimit
  CustomField: CustomField
  AssetCustomFieldValue: AssetCustomFieldValue
  AssetModelCustomFieldDefault: AssetModelCustomFieldDefault
  Invite: Invite
  Announcement: Announcement
  ConsumptionLog: ConsumptionLog
  BookingAsset: BookingAsset
  Booking: Booking
  BookingSettings: BookingSettings
  PartialBookingCheckin: PartialBookingCheckin
  PartialBookingCheckout: PartialBookingCheckout
  MobileAuthCode: MobileAuthCode
  PushToken: PushToken
  NotificationPreference: NotificationPreference
  Kit: Kit
  KitCustody: KitCustody
  AssetKit: AssetKit
  AssetLocation: AssetLocation
  AssetReminder: AssetReminder
  WorkingHours: WorkingHours
  WorkingHoursOverride: WorkingHoursOverride
  Update: Update
  UserUpdateRead: UserUpdateRead
  AuditNote: AuditNote
  TeamMemberNote: TeamMemberNote
  AuditSession: AuditSession
  AssetImport: AssetImport
  LinkPreview: LinkPreview
  AuditAssignment: AuditAssignment
  AuditAsset: AuditAsset
  AuditScan: AuditScan
  AuditImage: AuditImage
  RoleChangeLog: RoleChangeLog
  ActivityEvent: ActivityEvent
  _RoleToUser: RoleToUser
  _AssetToTag: AssetToTag
  _CategoryToCustomField: CategoryToCustomField
  _BookingToTag: BookingToTag
  _AssetReminderToTeamMember: AssetReminderToTeamMember
  _BookingSettingsAlwaysNotify: BookingSettingsAlwaysNotify
  _BookingNotificationRecipients: BookingNotificationRecipients
}
