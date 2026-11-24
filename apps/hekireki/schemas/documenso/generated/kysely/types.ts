import type { ColumnType } from 'kysely'

export type Generated<T> =
  T extends ColumnType<infer S, infer I, infer U>
    ? ColumnType<S, I | undefined, U>
    : ColumnType<T, T | undefined, T>

export type Timestamp = ColumnType<Date, Date | string, Date | string>

export type IdentityProvider = 'DOCUMENSO' | 'GOOGLE' | 'OIDC'

export type Role = 'ADMIN' | 'USER'

export type UserSecurityAuditLogType =
  | 'ACCOUNT_PROFILE_UPDATE'
  | 'ACCOUNT_SSO_LINK'
  | 'ACCOUNT_SSO_UNLINK'
  | 'ORGANISATION_SSO_LINK'
  | 'ORGANISATION_SSO_UNLINK'
  | 'AUTH_2FA_DISABLE'
  | 'AUTH_2FA_ENABLE'
  | 'PASSKEY_CREATED'
  | 'PASSKEY_DELETED'
  | 'PASSKEY_UPDATED'
  | 'PASSWORD_RESET'
  | 'PASSWORD_UPDATE'
  | 'SESSION_REVOKED'
  | 'SIGN_OUT'
  | 'SIGN_IN'
  | 'SIGN_IN_FAIL'
  | 'SIGN_IN_2FA_FAIL'
  | 'SIGN_IN_PASSKEY_FAIL'

export type WebhookTriggerEvents =
  | 'DOCUMENT_CREATED'
  | 'DOCUMENT_SENT'
  | 'DOCUMENT_OPENED'
  | 'DOCUMENT_SIGNED'
  | 'DOCUMENT_COMPLETED'
  | 'DOCUMENT_REJECTED'
  | 'DOCUMENT_CANCELLED'
  | 'RECIPIENT_EXPIRED'
  | 'DOCUMENT_RECIPIENT_COMPLETED'
  | 'DOCUMENT_REMINDER_SENT'
  | 'TEMPLATE_CREATED'
  | 'TEMPLATE_UPDATED'
  | 'TEMPLATE_DELETED'
  | 'TEMPLATE_USED'

export type WebhookCallStatus = 'SUCCESS' | 'FAILED'

export type ApiTokenAlgorithm = 'SHA512'

export type SubscriptionStatus = 'ACTIVE' | 'PAST_DUE' | 'INACTIVE'

export type DocumentStatus = 'DRAFT' | 'PENDING' | 'COMPLETED' | 'REJECTED' | 'CANCELLED'

export type DocumentSource = 'DOCUMENT' | 'TEMPLATE' | 'TEMPLATE_DIRECT_LINK'

export type DocumentVisibility = 'EVERYONE' | 'MANAGER_AND_ABOVE' | 'ADMIN'

export type FolderType = 'DOCUMENT' | 'TEMPLATE'

export type EnvelopeType = 'DOCUMENT' | 'TEMPLATE'

export type DocumentDataType = 'S3_PATH' | 'BYTES' | 'BYTES_64'

export type DocumentSigningOrder = 'PARALLEL' | 'SEQUENTIAL'

export type DocumentDistributionMethod = 'EMAIL' | 'NONE'

export type ReadStatus = 'NOT_OPENED' | 'OPENED'

export type SendStatus = 'NOT_SENT' | 'SENT'

export type SigningStatus = 'NOT_SIGNED' | 'SIGNED' | 'REJECTED'

export type RecipientRole = 'CC' | 'SIGNER' | 'VIEWER' | 'APPROVER' | 'ASSISTANT'

export type FieldType =
  | 'SIGNATURE'
  | 'FREE_SIGNATURE'
  | 'INITIALS'
  | 'NAME'
  | 'EMAIL'
  | 'DATE'
  | 'TEXT'
  | 'NUMBER'
  | 'RADIO'
  | 'CHECKBOX'
  | 'DROPDOWN'

export type OrganisationType = 'PERSONAL' | 'ORGANISATION'

export type OrganisationGroupType = 'INTERNAL_ORGANISATION' | 'INTERNAL_TEAM' | 'CUSTOM'

export type OrganisationMemberRole = 'ADMIN' | 'MANAGER' | 'MEMBER'

export type TeamMemberRole = 'ADMIN' | 'MANAGER' | 'MEMBER'

export type OrganisationMemberInviteStatus = 'ACCEPTED' | 'PENDING' | 'DECLINED'

export type TemplateType = 'PUBLIC' | 'PRIVATE' | 'ORGANISATION'

export type BackgroundJobStatus = 'PENDING' | 'PROCESSING' | 'COMPLETED' | 'FAILED'

export type BackgroundJobTaskStatus = 'PENDING' | 'COMPLETED' | 'FAILED'

export type EmailDomainStatus = 'PENDING' | 'ACTIVE'

export type EmailTransportType = 'SMTP_AUTH' | 'SMTP_API' | 'RESEND' | 'MAILCHANNELS'

export interface User {
  id: Generated<number>
  name: string | null
  email: string
  emailVerified: Timestamp | null
  password: string | null
  source: string | null
  signature: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  lastSignedIn: Generated<Timestamp>
  roles: Generated<ColumnType<string, Role[], Role[]>>
  identityProvider: Generated<IdentityProvider>
  avatarImageId: string | null
  disabled: Generated<boolean>
  twoFactorSecret: string | null
  twoFactorEnabled: Generated<boolean>
  twoFactorBackupCodes: string | null
}

export interface TeamProfile {
  id: string
  enabled: Generated<boolean>
  teamId: number
  bio: string | null
}

export interface UserSecurityAuditLog {
  id: Generated<number>
  userId: number
  createdAt: Generated<Timestamp>
  type: UserSecurityAuditLogType
  userAgent: string | null
  ipAddress: string | null
}

export interface PasswordResetToken {
  id: Generated<number>
  token: string
  createdAt: Generated<Timestamp>
  expiry: Timestamp
  userId: number
}

export interface Passkey {
  id: string
  userId: number
  name: string
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  lastUsedAt: Timestamp | null
  /** @zod.custom.use(z.instanceof(Uint8Array)) */
  credentialId: Buffer
  /** @zod.custom.use(z.instanceof(Uint8Array)) */
  credentialPublicKey: Buffer
  counter: ColumnType<string, bigint | number | string, bigint | number | string>
  credentialDeviceType: string
  credentialBackedUp: boolean
  transports: string[]
}

export interface AnonymousVerificationToken {
  id: string
  token: string
  expiresAt: Timestamp
  createdAt: Generated<Timestamp>
}

export interface VerificationToken {
  id: Generated<number>
  secondaryId: string
  identifier: string
  token: string
  completed: Generated<boolean>
  expires: Timestamp
  createdAt: Generated<Timestamp>
  metadata: unknown
  userId: number
}

export interface Webhook {
  id: string
  webhookUrl: string
  eventTriggers: ColumnType<string, WebhookTriggerEvents[], WebhookTriggerEvents[]>
  secret: string | null
  enabled: Generated<boolean>
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  userId: number
  teamId: number
}

export interface WebhookCall {
  id: string
  status: WebhookCallStatus
  url: string
  event: WebhookTriggerEvents
  requestBody: unknown
  responseCode: number
  responseHeaders: unknown
  responseBody: unknown
  createdAt: Generated<Timestamp>
  webhookId: string
}

export interface ApiToken {
  id: Generated<number>
  name: string
  token: string
  algorithm: Generated<ApiTokenAlgorithm>
  expires: Timestamp | null
  createdAt: Generated<Timestamp>
  lastUsedAt: Timestamp | null
  userId: number | null
  teamId: number
}

export interface Subscription {
  id: Generated<number>
  status: Generated<SubscriptionStatus>
  planId: string
  priceId: string
  periodEnd: Timestamp | null
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  cancelAtPeriodEnd: Generated<boolean>
  customerId: string
  organisationId: string
}

/** @zod.import(["import { ZClaimFlagsSchema, ZRateLimitArraySchema } from '@documenso/lib/types/subscription';"]) */
export interface SubscriptionClaim {
  id: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  name: string
  locked: Generated<boolean>
  teamCount: number
  memberCount: number
  envelopeItemCount: number
  recipientCount: number
  envelopeContentCount: number
  envelopeContentImageCount: number
  /** [ClaimFlags] @zod.custom.use(ZClaimFlagsSchema) */
  flags: unknown
  /** [RateLimitArray] @zod.custom.use(ZRateLimitArraySchema) */
  documentRateLimits: unknown
  documentQuota: number | null
  /** [RateLimitArray] @zod.custom.use(ZRateLimitArraySchema) */
  emailRateLimits: unknown
  emailQuota: number | null
  /** [RateLimitArray] @zod.custom.use(ZRateLimitArraySchema) */
  apiRateLimits: unknown
  apiQuota: number | null
  emailTransportId: string | null
}

/** @zod.import(["import { ZClaimFlagsSchema, ZRateLimitArraySchema } from '@documenso/lib/types/subscription';"]) */
export interface OrganisationClaim {
  id: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  originalSubscriptionClaimId: string | null
  teamCount: number
  memberCount: number
  envelopeItemCount: number
  recipientCount: number
  envelopeContentCount: number
  envelopeContentImageCount: number
  /** [ClaimFlags] @zod.custom.use(ZClaimFlagsSchema) */
  flags: unknown
  /** [RateLimitArray] @zod.custom.use(ZRateLimitArraySchema) */
  documentRateLimits: unknown
  documentQuota: number | null
  /** [RateLimitArray] @zod.custom.use(ZRateLimitArraySchema) */
  emailRateLimits: unknown
  emailQuota: number | null
  /** [RateLimitArray] @zod.custom.use(ZRateLimitArraySchema) */
  apiRateLimits: unknown
  apiQuota: number | null
  emailTransportId: string | null
}

export interface OrganisationMonthlyStat {
  id: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  organisationId: string
  /** UTC calendar month in `YYYY-MM` form, e.g. "2026-05". */
  period: string
  documentCount: Generated<number>
  emailCount: Generated<number>
  apiCount: Generated<number>
  emailReports: Generated<number>
}

export interface Account {
  id: string
  createdAt: Generated<Timestamp>
  userId: number
  type: string
  provider: string
  providerAccountId: string
  refresh_token: string | null
  access_token: string | null
  expires_at: number | null
  created_at: number | null
  ext_expires_in: number | null
  token_type: string | null
  scope: string | null
  id_token: string | null
  session_state: string | null
  password: string | null
}

export interface Session {
  id: string
  sessionToken: string
  userId: number
  ipAddress: string | null
  userAgent: string | null
  expiresAt: Timestamp
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  isEmbed: Generated<boolean>
}

export interface AuthHandoff {
  id: string
  type: string
  payload: string
  createdAt: Generated<Timestamp>
  expiresAt: Timestamp
}

export interface Folder {
  id: string
  name: string
  userId: number
  teamId: number
  pinned: Generated<boolean>
  parentId: string | null
  createdAt: Generated<Timestamp>
  updatedAt: Generated<Timestamp>
  visibility: Generated<DocumentVisibility>
  type: FolderType
}

/** @zod.import(["import { ZDocumentAuthOptionsSchema } from '@documenso/lib/types/document-auth';", "import { ZDocumentFormValuesSchema } from '@documenso/lib/types/document-form-values';", "import { ZSignatureLevelSchema } from '@documenso/lib/types/signature-level';"]) */
export interface Envelope {
  id: string
  secondaryId: string
  /** @zod.string.describe("A custom external ID you can use to identify the document.") */
  externalId: string | null
  type: EnvelopeType
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  completedAt: Timestamp | null
  deletedAt: Timestamp | null
  title: string
  status: Generated<DocumentStatus>
  source: DocumentSource
  /** @zod.string.describe("The token for viewing the document using the QR code on the certificate.") */
  qrToken: string | null
  /** [SignatureLevel] @zod.custom.use(ZSignatureLevelSchema) */
  signatureLevel: string
  internalVersion: number
  useLegacyFieldInsertion: Generated<boolean>
  /** [DocumentAuthOptions] @zod.custom.use(ZDocumentAuthOptionsSchema) */
  authOptions: unknown
  /** [DocumentFormValues] @zod.custom.use(ZDocumentFormValuesSchema) */
  formValues: unknown
  visibility: Generated<DocumentVisibility>
  templateType: Generated<TemplateType>
  publicTitle: Generated<string>
  publicDescription: Generated<string>
  templateId: number | null
  /** @zod.number.describe("The ID of the user that created this document.") */
  userId: number
  teamId: number
  folderId: string | null
  documentMetaId: string
}

export interface EnvelopeItem {
  id: string
  title: string
  order: number
  documentDataId: string
  envelopeId: string
}

export interface DocumentAuditLog {
  id: string
  envelopeId: string | null
  createdAt: Generated<Timestamp>
  type: string
  data: unknown
  name: string | null
  email: string | null
  userId: number | null
  userAgent: string | null
  ipAddress: string | null
}

export interface DocumentData {
  id: string
  type: DocumentDataType
  data: string
  initialData: string
}

/** @zod.import(["import { ZDocumentEmailSettingsSchema } from '@documenso/lib/types/document-email';", "import { ZEnvelopeExpirationPeriod as ZEnvelopeExpirationPeriodSchema } from '@documenso/lib/constants/envelope-expiration';", "import { ZEnvelopeReminderSettings as ZEnvelopeReminderSettingsSchema } from '@documenso/lib/constants/envelope-reminder';"]) */
export interface DocumentMeta {
  id: string
  subject: string | null
  message: string | null
  timezone: Generated<string | null>
  dateFormat: Generated<string | null>
  redirectUrl: string | null
  signingOrder: Generated<DocumentSigningOrder>
  allowDictateNextSigner: Generated<boolean>
  typedSignatureEnabled: Generated<boolean>
  uploadSignatureEnabled: Generated<boolean>
  drawSignatureEnabled: Generated<boolean>
  language: Generated<string>
  distributionMethod: Generated<DocumentDistributionMethod>
  /** [DocumentEmailSettings] @zod.custom.use(ZDocumentEmailSettingsSchema) */
  emailSettings: unknown
  emailReplyTo: string | null
  emailId: string | null
  /** [EnvelopeExpirationPeriod] @zod.custom.use(ZEnvelopeExpirationPeriodSchema) */
  envelopeExpirationPeriod: unknown
  /** [EnvelopeReminderSettings] @zod.custom.use(ZEnvelopeReminderSettingsSchema) */
  reminderSettings: unknown
}

/** @zod.import(["import { ZEnvelopeAttachmentTypeSchema } from '@documenso/lib/types/envelope-attachment';"]) */
export interface EnvelopeAttachment {
  id: string
  /** [EnvelopeAttachmentType] @zod.custom.use(ZEnvelopeAttachmentTypeSchema) */
  type: string
  label: string
  data: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  envelopeId: string
}

/** @zod.import(["import { ZRecipientAuthOptionsSchema } from '@documenso/lib/types/document-auth';"]) */
export interface Recipient {
  id: Generated<number>
  envelopeId: string
  email: string
  name: Generated<string>
  token: string
  documentDeletedAt: Timestamp | null
  expired: Timestamp | null
  expiresAt: Timestamp | null
  expirationNotifiedAt: Timestamp | null
  sentAt: Timestamp | null
  signedAt: Timestamp | null
  lastReminderSentAt: Timestamp | null
  nextReminderAt: Timestamp | null
  reminderCount: Generated<number>
  /** [RecipientAuthOptions] @zod.custom.use(ZRecipientAuthOptionsSchema) */
  authOptions: unknown
  /** @zod.number.describe("The order in which the recipient should sign the document. Only works if the document is set to sequential signing.") */
  signingOrder: number | null
  rejectionReason: string | null
  role: Generated<RecipientRole>
  readStatus: Generated<ReadStatus>
  signingStatus: Generated<SigningStatus>
  sendStatus: Generated<SendStatus>
}

/** @zod.import(["import { ZFieldMetaNotOptionalSchema } from '@documenso/lib/types/field-meta';"]) */
export interface Field {
  id: Generated<number>
  secondaryId: string
  envelopeId: string
  envelopeItemId: string
  recipientId: number
  type: FieldType
  /** @zod.number.describe("The page number of the field on the document. Starts from 1.") */
  page: number
  positionX: Generated<string>
  positionY: Generated<string>
  width: Generated<string>
  height: Generated<string>
  customText: string
  inserted: boolean
  /** [FieldMeta] @zod.custom.use(ZFieldMetaNotOptionalSchema) */
  fieldMeta: unknown
}

/** @zod.import(["import { ZEnvelopeContentMetaSchema } from '@documenso/lib/types/envelope-content-meta';"]) */
export interface EnvelopeContent {
  id: string
  envelopeId: string
  envelopeItemId: string
  /** [EnvelopeContentMeta] @zod.custom.use(ZEnvelopeContentMetaSchema) */
  contentMeta: unknown
  /**
   * Any number of contents may point at the same data content, e.g. a
   * document created from a template shares the template's images.
   */
  dataContentId: string | null
}

/**
 * Generic stored data (e.g. an uploaded image) which other models reference.
 * The kind of data and its kind specific details live in `metadata`.
 * @zod.import(["import { ZDataContentMetaSchema } from '@documenso/lib/types/data-content-meta';"])
 */
export interface DataContent {
  id: string
  /** How `data` is stored, e.g. an S3 path or inline bytes. */
  type: DocumentDataType
  data: string
  /** [DataContentMeta] @zod.custom.use(ZDataContentMetaSchema) */
  metadata: unknown
}

export interface Signature {
  id: Generated<number>
  created: Generated<Timestamp>
  recipientId: number
  fieldId: number
  signatureImageAsBase64: string | null
  typedSignature: string | null
}

/**
 * Per-recipient cached credential metadata for the configured Cloud Signature
 * Consortium (CSC) provider. Holds the TSP-validated certificate chain and the
 * symmetric-encrypted service-scope access token. One row per recipient on a
 * CSC envelope; absent for SES recipients.
 */
export interface CscCredential {
  id: string
  providerId: string
  credentialId: string
  /**
   * TSP-validated certificate chain (length-prefixed). Nullable until the
   * service-scope OAuth callback successfully retrieves and validates it.
   */
  certCache: Buffer | null
  /**
   * Algorithm metadata derived from the CSC `credentials/info` response.
   * Held as flat scalars so each is independently queryable.
   */
  signatureAlgorithm: string
  keyType: string
  digestAlgorithm: string
  keyLenBits: number | null
  signAlgoParams: string | null
  /**
   * Service-scope access token, symmetric-encrypted at the application layer
   * before persistence. Nullable until the service-scope OAuth callback runs.
   */
  serviceTokenCiphertext: Buffer | null
  serviceTokenExpiresAt: Timestamp | null
  recipientId: number
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

/**
 * Per-recipient transient sign-time state held across the credential-scope
 * OAuth round-trip. The `recipientId @unique` constraint caps each recipient
 * to at most one in-flight session — re-clicking Sign UPSERTs the row.
 * Cleaned up transitively when the Recipient is deleted.
 * @zod.import(["import { ZCscSessionItemsSchema } from '@documenso/lib/types/csc-session';"])
 */
export interface CscSession {
  id: string
  /**
   * Denormalised envelope id for ergonomic lookups. No FK relation — cleanup
   * flows through the Recipient cascade.
   */
  envelopeId: string
  /**
   * Pinned at prep time; stable across overlay re-renders so the captured
   * signedAttrs digest matches the bytes the TSP signs.
   */
  signingTime: Timestamp
  /**
   * Contract between prep and sign: an array of
   * `{ envelopeItemId, documentDataId, hashB64, ordinal }` per envelope item.
   * `documentDataId` pins the immutable bytes that produced the captured
   * digest so sign-time can re-derive the same hash against the same content.
   * [CscSessionItems] @zod.custom.use(ZCscSessionItemsSchema)
   */
  itemsJson: unknown
  /**
   * Populated by the credential-scope OAuth callback. Both nullable at row
   * creation; set together when the SAD is exchanged.
   */
  encryptedSad: Buffer | null
  sadExpiresAt: Timestamp | null
  createdAt: Generated<Timestamp>
  recipientId: number
}

export interface DocumentShareLink {
  id: Generated<number>
  email: string
  slug: string
  envelopeId: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
}

export interface Organisation {
  id: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  type: OrganisationType
  name: string
  url: string
  avatarImageId: string | null
  customerId: string | null
  organisationClaimId: string
  ownerUserId: number
  organisationGlobalSettingsId: string
  organisationAuthenticationPortalId: string
}

export interface OrganisationMember {
  id: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  userId: number
  organisationId: string
}

export interface OrganisationMemberInvite {
  id: string
  createdAt: Generated<Timestamp>
  email: string
  token: string
  status: Generated<OrganisationMemberInviteStatus>
  organisationId: string
  organisationRole: OrganisationMemberRole
}

export interface OrganisationGroup {
  id: string
  name: string | null
  type: OrganisationGroupType
  organisationRole: OrganisationMemberRole
  organisationId: string
}

export interface OrganisationGroupMember {
  id: string
  groupId: string
  organisationMemberId: string
}

export interface TeamGroup {
  id: string
  organisationGroupId: string
  teamRole: TeamMemberRole
  teamId: number
}

/** @zod.import(["import { ZDocumentEmailSettingsSchema } from '@documenso/lib/types/document-email';", "import { ZDefaultRecipientsSchema } from '@documenso/lib/types/default-recipients';", "import { ZEnvelopeExpirationPeriod as ZEnvelopeExpirationPeriodSchema } from '@documenso/lib/constants/envelope-expiration';", "import { ZEnvelopeReminderSettings as ZEnvelopeReminderSettingsSchema } from '@documenso/lib/constants/envelope-reminder';", "import { ZCssVarsSchema } from '@documenso/lib/types/css-vars';"]) */
export interface OrganisationGlobalSettings {
  id: string
  documentVisibility: Generated<DocumentVisibility>
  documentLanguage: Generated<string>
  includeSenderDetails: Generated<boolean>
  includeSigningCertificate: Generated<boolean>
  includeAuditLog: Generated<boolean>
  documentTimezone: string | null
  documentDateFormat: Generated<string>
  delegateDocumentOwnership: Generated<boolean>
  typedSignatureEnabled: Generated<boolean>
  uploadSignatureEnabled: Generated<boolean>
  drawSignatureEnabled: Generated<boolean>
  /** [DefaultRecipient[]] @zod.custom.use(ZDefaultRecipientsSchema) */
  defaultRecipients: unknown
  emailId: string | null
  emailReplyTo: string | null
  /** [DocumentEmailSettings] @zod.custom.use(ZDocumentEmailSettingsSchema) */
  emailDocumentSettings: unknown
  brandingEnabled: Generated<boolean>
  brandingLogo: Generated<string>
  brandingUrl: Generated<string>
  brandingCompanyDetails: Generated<string>
  /** [TCssVarsSchema] @zod.custom.use(ZCssVarsSchema) */
  brandingColors: unknown
  brandingCss: Generated<string>
  /** [EnvelopeExpirationPeriod] @zod.custom.use(ZEnvelopeExpirationPeriodSchema) */
  envelopeExpirationPeriod: unknown
  /** [EnvelopeReminderSettings] @zod.custom.use(ZEnvelopeReminderSettingsSchema) */
  reminderSettings: unknown
  aiFeaturesEnabled: Generated<boolean>
}

/** @zod.import(["import { ZDocumentEmailSettingsSchema } from '@documenso/lib/types/document-email';", "import { ZDefaultRecipientsSchema } from '@documenso/lib/types/default-recipients';", "import { ZEnvelopeExpirationPeriod as ZEnvelopeExpirationPeriodSchema } from '@documenso/lib/constants/envelope-expiration';", "import { ZEnvelopeReminderSettings as ZEnvelopeReminderSettingsSchema } from '@documenso/lib/constants/envelope-reminder';", "import { ZCssVarsSchema } from '@documenso/lib/types/css-vars';"]) */
export interface TeamGlobalSettings {
  id: string
  documentVisibility: DocumentVisibility | null
  documentLanguage: string | null
  documentTimezone: string | null
  documentDateFormat: string | null
  delegateDocumentOwnership: boolean | null
  includeSenderDetails: boolean | null
  includeSigningCertificate: boolean | null
  includeAuditLog: boolean | null
  typedSignatureEnabled: boolean | null
  uploadSignatureEnabled: boolean | null
  drawSignatureEnabled: boolean | null
  /** [DefaultRecipient[]] @zod.custom.use(ZDefaultRecipientsSchema) */
  defaultRecipients: unknown
  emailId: string | null
  emailReplyTo: string | null
  /** [DocumentEmailSettingsNullable] @zod.custom.use(ZDocumentEmailSettingsSchema) */
  emailDocumentSettings: unknown
  brandingEnabled: boolean | null
  brandingLogo: string | null
  brandingUrl: string | null
  brandingCompanyDetails: string | null
  /** [TCssVarsSchema] @zod.custom.use(ZCssVarsSchema) */
  brandingColors: unknown
  brandingCss: string | null
  /** [EnvelopeExpirationPeriod] @zod.custom.use(ZEnvelopeExpirationPeriodSchema) */
  envelopeExpirationPeriod: unknown
  /** [EnvelopeReminderSettings] @zod.custom.use(ZEnvelopeReminderSettingsSchema) */
  reminderSettings: unknown
  aiFeaturesEnabled: boolean | null
}

export interface Team {
  id: Generated<number>
  name: string
  url: string
  createdAt: Generated<Timestamp>
  avatarImageId: string | null
  organisationId: string
  teamGlobalSettingsId: string
}

export interface TeamEmail {
  teamId: number
  createdAt: Generated<Timestamp>
  name: string
  email: string
}

export interface TeamEmailVerification {
  teamId: number
  name: string
  email: string
  token: string
  completed: Generated<boolean>
  expiresAt: Timestamp
  createdAt: Generated<Timestamp>
}

export interface TemplateDirectLink {
  id: string
  envelopeId: string
  token: string
  createdAt: Generated<Timestamp>
  enabled: boolean
  directTemplateRecipientId: number
}

export interface SiteSettings {
  id: string
  enabled: Generated<boolean>
  data: unknown
  lastModifiedByUserId: number | null
  lastModifiedAt: Generated<Timestamp>
}

export interface BackgroundJob {
  id: string
  status: Generated<BackgroundJobStatus>
  payload: unknown
  retried: Generated<number>
  maxRetries: Generated<number>
  jobId: string
  name: string
  version: string
  submittedAt: Generated<Timestamp>
  updatedAt: Timestamp
  completedAt: Timestamp | null
  lastRetriedAt: Timestamp | null
}

export interface BackgroundJobTask {
  id: string
  name: string
  status: Generated<BackgroundJobTaskStatus>
  result: unknown
  retried: Generated<number>
  maxRetries: Generated<number>
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  completedAt: Timestamp | null
  jobId: string
}

export interface AvatarImage {
  id: string
  bytes: string
}

export interface EmailDomain {
  id: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  status: Generated<EmailDomainStatus>
  selector: string
  domain: string
  publicKey: string
  privateKey: string
  lastVerifiedAt: Timestamp | null
  organisationId: string
}

export interface OrganisationEmail {
  id: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  email: string
  emailName: string
  emailDomainId: string
  organisationId: string
}

export interface EmailTransport {
  id: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  name: string
  type: EmailTransportType
  fromName: string
  fromAddress: string
  config: string
}

export interface OrganisationAuthenticationPortal {
  id: string
  enabled: Generated<boolean>
  clientId: Generated<string>
  clientSecret: Generated<string>
  wellKnownUrl: Generated<string>
  defaultOrganisationRole: Generated<OrganisationMemberRole>
  autoProvisionUsers: Generated<boolean>
  allowedDomains: Generated<string[]>
  allowPersonalOrganisations: Generated<boolean>
}

export interface Counter {
  id: string
  value: number
}

/**
 * `key` is sha256(organisationId:lower(recipientEmail):clientUserId) so neither
 * the recipient email nor the embedder's identifier is stored in plaintext.
 */
export interface EmbedSignature {
  key: string
  createdAt: Generated<Timestamp>
  updatedAt: Timestamp
  organisationId: string
  /**
   * Last writer. Not a relation so the row survives document deletion and
   * remains useful for abuse tracing.
   */
  recipientId: number
  signature: string
}

export interface RateLimit {
  key: string
  action: string
  bucket: Timestamp
  count: Generated<number>
  createdAt: Generated<Timestamp>
}

export interface DB {
  User: User
  TeamProfile: TeamProfile
  UserSecurityAuditLog: UserSecurityAuditLog
  PasswordResetToken: PasswordResetToken
  Passkey: Passkey
  AnonymousVerificationToken: AnonymousVerificationToken
  VerificationToken: VerificationToken
  Webhook: Webhook
  WebhookCall: WebhookCall
  ApiToken: ApiToken
  Subscription: Subscription
  SubscriptionClaim: SubscriptionClaim
  OrganisationClaim: OrganisationClaim
  OrganisationMonthlyStat: OrganisationMonthlyStat
  Account: Account
  Session: Session
  AuthHandoff: AuthHandoff
  Folder: Folder
  Envelope: Envelope
  EnvelopeItem: EnvelopeItem
  DocumentAuditLog: DocumentAuditLog
  DocumentData: DocumentData
  DocumentMeta: DocumentMeta
  EnvelopeAttachment: EnvelopeAttachment
  Recipient: Recipient
  Field: Field
  EnvelopeContent: EnvelopeContent
  DataContent: DataContent
  Signature: Signature
  CscCredential: CscCredential
  CscSession: CscSession
  DocumentShareLink: DocumentShareLink
  Organisation: Organisation
  OrganisationMember: OrganisationMember
  OrganisationMemberInvite: OrganisationMemberInvite
  OrganisationGroup: OrganisationGroup
  OrganisationGroupMember: OrganisationGroupMember
  TeamGroup: TeamGroup
  OrganisationGlobalSettings: OrganisationGlobalSettings
  TeamGlobalSettings: TeamGlobalSettings
  Team: Team
  TeamEmail: TeamEmail
  TeamEmailVerification: TeamEmailVerification
  TemplateDirectLink: TemplateDirectLink
  SiteSettings: SiteSettings
  BackgroundJob: BackgroundJob
  BackgroundJobTask: BackgroundJobTask
  AvatarImage: AvatarImage
  EmailDomain: EmailDomain
  OrganisationEmail: OrganisationEmail
  EmailTransport: EmailTransport
  OrganisationAuthenticationPortal: OrganisationAuthenticationPortal
  Counter: Counter
  EmbedSignature: EmbedSignature
  RateLimit: RateLimit
}
