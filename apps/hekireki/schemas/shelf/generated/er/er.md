```mermaid
erDiagram
    Organization ||--o{ Image : "(id) - (ownerOrgId)"
    User ||--o{ Image : "(id) - (userId)"
    Organization |o--o{ User : "(id) - (lastSelectedOrganizationId)"
    Tier ||--o{ User : "(id) - (tierId)"
    User ||--o| UserContact : "(id) - (userId)"
    User ||--o| UserBusinessIntel : "(id) - (userId)"
    AssetModel |o--o{ Asset : "(id) - (assetModelId)"
    User ||--o{ Asset : "(id) - (userId)"
    Organization ||--o{ Asset : "(id) - (organizationId)"
    Category |o--o{ Asset : "(id) - (categoryId)"
    Barcode |o--o| Asset : "(id) - (preferredBarcodeId)"
    Organization ||--o{ AssetFilterPreset : "(id) - (organizationId)"
    User ||--o{ AssetFilterPreset : "(id) - (ownerId)"
    User ||--o{ AssetIndexSettings : "(id) - (userId)"
    Organization ||--o{ AssetIndexSettings : "(id) - (organizationId)"
    Organization ||--o{ PdfTemplate : "(id) - (organizationId)"
    User ||--o{ Category : "(id) - (userId)"
    Organization ||--o{ Category : "(id) - (organizationId)"
    Category |o--o{ AssetModel : "(id) - (defaultCategoryId)"
    Organization ||--o{ AssetModel : "(id) - (organizationId)"
    User ||--o{ AssetModel : "(id) - (userId)"
    Booking ||--o{ BookingModelRequest : "(id) - (bookingId)"
    AssetModel ||--o{ BookingModelRequest : "(id) - (assetModelId)"
    User ||--o{ Tag : "(id) - (userId)"
    Organization ||--o{ Tag : "(id) - (organizationId)"
    User |o--o{ Note : "(id) - (userId)"
    Asset ||--o{ Note : "(id) - (assetId)"
    User |o--o{ BookingNote : "(id) - (userId)"
    Booking ||--o{ BookingNote : "(id) - (bookingId)"
    User |o--o{ LocationNote : "(id) - (userId)"
    Location ||--o{ LocationNote : "(id) - (locationId)"
    Asset |o--o{ Qr : "(id) - (assetId)"
    Kit |o--o{ Qr : "(id) - (kitId)"
    User |o--o{ Qr : "(id) - (userId)"
    Organization |o--o{ Qr : "(id) - (organizationId)"
    PrintBatch |o--o{ Qr : "(id) - (batchId)"
    Asset |o--o{ Barcode : "(id) - (assetId)"
    Kit |o--o{ Barcode : "(id) - (kitId)"
    Organization ||--o{ Barcode : "(id) - (organizationId)"
    Asset |o--o{ ReportFound : "(id) - (assetId)"
    Kit |o--o{ ReportFound : "(id) - (kitId)"
    User |o--o{ Scan : "(id) - (userId)"
    Qr |o--o{ Scan : "(id) - (qrId)"
    Barcode |o--o{ Scan : "(id) - (barcodeId)"
    Asset |o--o{ Scan : "(id) - (assetId)"
    Kit |o--o{ Scan : "(id) - (kitId)"
    Organization ||--o{ Scan : "(id) - (organizationId)"
    Image |o--o| Location : "(id) - (imageId)"
    User ||--o{ Location : "(id) - (userId)"
    Organization ||--o{ Location : "(id) - (organizationId)"
    Location |o--o{ Location : "(id) - (parentId)"
    Organization ||--o{ TeamMember : "(id) - (organizationId)"
    User |o--o{ TeamMember : "(id) - (userId)"
    TeamMember ||--o{ Custody : "(id) - (teamMemberId)"
    Asset ||--o{ Custody : "(id) - (assetId)"
    KitCustody |o--o{ Custody : "(id) - (kitCustodyId)"
    Location |o--o{ Custody : "(id) - (locationId)"
    User ||--o{ Organization : "(id) - (userId)"
    Image |o--o| Organization : "(id) - (imageId)"
    SsoDetails |o--o{ Organization : "(id) - (ssoDetailsId)"
    User ||--o{ UserOrganization : "(id) - (userId)"
    Organization ||--o{ UserOrganization : "(id) - (organizationId)"
    Organization ||--o{ ScimToken : "(id) - (organizationId)"
    User ||--o{ UserScimExternalId : "(id) - (userId)"
    Organization ||--o{ UserScimExternalId : "(id) - (organizationId)"
    TierLimit |o--o| Tier : "(id) - (tierLimitId)"
    User |o--o| CustomTierLimit : "(id) - (userId)"
    Organization ||--o{ CustomField : "(id) - (organizationId)"
    User ||--o{ CustomField : "(id) - (userId)"
    Asset ||--o{ AssetCustomFieldValue : "(id) - (assetId)"
    CustomField ||--o{ AssetCustomFieldValue : "(id) - (customFieldId)"
    AssetModel ||--o{ AssetModelCustomFieldDefault : "(id) - (assetModelId)"
    User ||--o{ Invite : "(id) - (inviterId)"
    Organization ||--o{ Invite : "(id) - (organizationId)"
    User |o--o{ Invite : "(id) - (inviteeUserId)"
    TeamMember ||--o{ Invite : "(id) - (teamMemberId)"
    Asset ||--o{ ConsumptionLog : "(id) - (assetId)"
    User ||--o{ ConsumptionLog : "(id) - (userId)"
    Booking |o--o{ ConsumptionLog : "(id) - (bookingId)"
    TeamMember |o--o{ ConsumptionLog : "(id) - (custodianId)"
    Location |o--o{ ConsumptionLog : "(id) - (locationId)"
    Booking ||--o{ BookingAsset : "(id) - (bookingId)"
    Asset ||--o{ BookingAsset : "(id) - (assetId)"
    User ||--o{ Booking : "(id) - (creatorId)"
    User |o--o{ Booking : "(id) - (custodianUserId)"
    TeamMember |o--o{ Booking : "(id) - (custodianTeamMemberId)"
    Organization ||--o{ Booking : "(id) - (organizationId)"
    User |o--o{ Booking : "(id) - (approvedById)"
    Organization ||--o| BookingSettings : "(id) - (organizationId)"
    Booking ||--o{ PartialBookingCheckin : "(id) - (bookingId)"
    User ||--o{ PartialBookingCheckin : "(id) - (checkedInById)"
    Booking ||--o{ PartialBookingCheckout : "(id) - (bookingId)"
    User ||--o{ PartialBookingCheckout : "(id) - (checkedOutById)"
    User ||--o{ MobileAuthCode : "(id) - (userId)"
    User ||--o{ PushToken : "(id) - (userId)"
    User ||--o| NotificationPreference : "(id) - (userId)"
    Organization ||--o{ Kit : "(id) - (organizationId)"
    User ||--o{ Kit : "(id) - (createdById)"
    Category |o--o{ Kit : "(id) - (categoryId)"
    Location |o--o{ Kit : "(id) - (locationId)"
    TeamMember ||--o{ KitCustody : "(id) - (custodianId)"
    Kit ||--o| KitCustody : "(id) - (kitId)"
    Asset ||--o{ AssetKit : "(id) - (assetId)"
    Kit ||--o{ AssetKit : "(id) - (kitId)"
    Organization ||--o{ AssetKit : "(id) - (organizationId)"
    Asset ||--o{ AssetLocation : "(id) - (assetId)"
    Location ||--o{ AssetLocation : "(id) - (locationId)"
    Organization ||--o{ AssetLocation : "(id) - (organizationId)"
    AssetKit |o--o{ AssetLocation : "(id) - (assetKitId)"
    Organization ||--o{ AssetReminder : "(id) - (organizationId)"
    Asset ||--o{ AssetReminder : "(id) - (assetId)"
    User ||--o{ AssetReminder : "(id) - (createdById)"
    Organization ||--o| WorkingHours : "(id) - (organizationId)"
    WorkingHours ||--o{ WorkingHoursOverride : "(id) - (workingHoursId)"
    User ||--o{ Update : "(id) - (createdById)"
    User ||--o{ UserUpdateRead : "(id) - (userId)"
    Update ||--o{ UserUpdateRead : "(id) - (updateId)"
    User |o--o{ AuditNote : "(id) - (userId)"
    AuditSession ||--o{ AuditNote : "(id) - (auditSessionId)"
    AuditAsset |o--o{ AuditNote : "(id) - (auditAssetId)"
    User |o--o{ TeamMemberNote : "(id) - (userId)"
    TeamMember ||--o{ TeamMemberNote : "(id) - (teamMemberId)"
    Organization ||--o{ TeamMemberNote : "(id) - (organizationId)"
    User ||--o{ AuditSession : "(id) - (createdById)"
    Organization ||--o{ AuditSession : "(id) - (organizationId)"
    User ||--o{ AssetImport : "(id) - (createdById)"
    Organization ||--o{ AssetImport : "(id) - (organizationId)"
    Organization ||--o{ LinkPreview : "(id) - (organizationId)"
    AuditSession ||--o{ AuditAssignment : "(id) - (auditSessionId)"
    User ||--o{ AuditAssignment : "(id) - (userId)"
    AuditSession ||--o{ AuditAsset : "(id) - (auditSessionId)"
    Asset ||--o{ AuditAsset : "(id) - (assetId)"
    User |o--o{ AuditAsset : "(id) - (scannedById)"
    AuditSession ||--o{ AuditScan : "(id) - (auditSessionId)"
    AuditAsset |o--o{ AuditScan : "(id) - (auditAssetId)"
    Asset |o--o{ AuditScan : "(id) - (assetId)"
    User |o--o{ AuditScan : "(id) - (scannedById)"
    AuditSession ||--o{ AuditImage : "(id) - (auditSessionId)"
    AuditAsset |o--o{ AuditImage : "(id) - (auditAssetId)"
    User |o--o{ AuditImage : "(id) - (uploadedById)"
    Organization ||--o{ AuditImage : "(id) - (organizationId)"
    User ||--o{ RoleChangeLog : "(id) - (userId)"
    User ||--o{ RoleChangeLog : "(id) - (changedById)"
    Organization ||--o{ RoleChangeLog : "(id) - (organizationId)"
    Organization ||--o{ ActivityEvent : "(id) - (organizationId)"
    Asset }o--o{ Tag : "(tags) - (assets)"
    Category }o--o{ CustomField : "(customFields) - (categories)"
    Role }o--o{ User : "(users) - (roles)"
    Booking }o--o{ Tag : "(tags) - (bookings)"
    Booking }o--o{ TeamMember : "(notificationRecipients) - (bookingNotifications)"
    BookingSettings }o--o{ TeamMember : "(alwaysNotifyTeamMembers) - (alwaysNotifyBookingSettings)"
    AssetReminder }o--o{ TeamMember : "(teamMembers) - (assetReminders)"
    Image {
        string id PK
        string contentType
        string altText
        bytes blob
        datetime createdAt
        datetime updatedAt
        string ownerOrgId FK
        string userId FK
    }
    User {
        string id PK
        string email UK
        string username UK
        string firstName
        string lastName
        string displayName
        string profilePicture
        boolean usedFreeTrial
        boolean onboarded
        string customerId UK
        boolean sso
        boolean createdWithInvite
        boolean skipSubscriptionCheck
        boolean hasUnpaidInvoice
        boolean warnForNoPaymentMethod
        string lastSelectedOrganizationId FK
        datetime createdAt
        datetime updatedAt
        datetime deletedAt
        datetime lastMobileActiveAt
        dateformatpreference dateFormat
        timeformatpreference timeFormat
        weekstartpreference weekStart
        string timeZone
        tierid tierId FK
        string referralSource
    }
    UserContact {
        string id PK
        string phone
        string street
        string city
        string stateProvince
        string zipPostalCode
        string countryRegion
        string userId FK, UK
        datetime createdAt
        datetime updatedAt
    }
    UserBusinessIntel {
        string id PK
        string howDidYouHearAboutUs
        string jobTitle
        string teamSize
        string companyName
        string primaryUseCase
        string currentSolution
        string timeline
        string signupPlan
        boolean signupTrial
        string utmSource
        string utmMedium
        string utmCampaign
        string utmContent
        string userId FK, UK
        datetime createdAt
        datetime updatedAt
    }
    Asset {
        string id PK
        string title
        string description
        string mainImage
        string thumbnailImage
        datetime mainImageExpiration
        assetstatus status
        float valuation
        boolean availableToBook
        string sequentialId UK
        datetime archivedAt
        assettype type
        string assetModelId FK
        int quantity
        int minQuantity
        consumptiontype consumptionType
        string unitOfMeasure
        datetime lowStockNotifiedAt
        datetime stockLedgerStartedAt
        datetime createdAt
        datetime updatedAt
        string userId FK
        string organizationId FK, UK
        string categoryId FK
        string preferredBarcodeId FK, UK
    }
    AssetFilterPreset {
        string id PK
        string organizationId FK, UK
        string ownerId FK, UK
        string name UK
        string query
        boolean starred
        datetime createdAt
        datetime updatedAt
    }
    AssetIndexSettings {
        string id PK
        string userId FK, UK
        string organizationId FK, UK
        assetindexmode mode
        json columns
        boolean freezeColumn
        boolean showAssetImage
        datetime createdAt
        datetime updatedAt
    }
    PdfTemplate {
        string id PK
        string name UK
        pdftemplatekind kind UK
        json config "The layout's switches. Read through `parseBookingTemplateConfig`, which fills any key a stored row lacks from the Checklist template's values."
        boolean isDefault "The template the PDF dialog offers first. One per workspace and kind, enforced by a partial unique index in raw SQL (migration 20261103120000_add_pdf_templates). Names are also unique per workspace and kind ignoring case, through an expression index in the same migration."
        string organizationId FK, UK
        datetime createdAt
        datetime updatedAt
    }
    Category {
        string id PK
        string name
        string description
        string color
        datetime createdAt
        datetime updatedAt
        string userId FK
        string organizationId FK
    }
    AssetModel {
        string id PK
        string name
        string description
        string image "Public URL (bucket `files`) of the model's cover image. Rendered by any asset of this model that has no image of its own. Public — and therefore never-expiring and cacheable — like Location.imageUrl; see resolveAssetImage() for the cascade."
        string thumbnailImage "Public URL of the 108px thumbnail generated alongside `image`. Used at list sizes so inheriting assets never render the full-size original."
        string defaultCategoryId FK
        float defaultValuation
        string organizationId FK
        string userId FK
        datetime createdAt
        datetime updatedAt
    }
    BookingModelRequest {
        string id PK
        string bookingId FK, UK
        string assetModelId FK, UK
        int quantity "Total units originally reserved. Mutated only by explicit operator edits (upsert) — never decremented on materialisation. This preserves the audit trail of #quot;how many were planned#quot;."
        int fulfilledQuantity "Units already assigned to concrete `BookingAsset` rows via `materializeModelRequestForAsset`. Scans increment it.  It DOES decrease: `removeAssets` (booking service) decrements it by the number of removed `BookingAsset` rows whose `bookingModelRequestId` points at this request, floored at 0, and clears `fulfilledAt` when the result drops below `quantity`.  Counted from that column, never from a shared `assetModelId`: an asset added directly alongside a reservation shares the model without having discharged anything, so grouping by model re-opened reservations that were still correctly satisfied."
        datetime fulfilledAt "Timestamp of the moment the reservation became fully fulfilled — set when `fulfilledQuantity` catches up to `quantity`, and preserved (not re-stamped) if an unchanged quantity is saved again. Cleared back to null when outstanding units reappear: the operator edits `quantity` upwards, or `removeAssets` drops `fulfilledQuantity` below it.  The invariant that makes this the primary #quot;is this outstanding?#quot; filter: non-null IFF `fulfilledQuantity === quantity`, so `fulfilledAt IS NULL` means units remain to materialise. `fulfilledQuantity` can never exceed `quantity` (the claim in `materializeModelRequestForAsset` is a single statement guarded by `WHERE #quot;fulfilledQuantity#quot; < #quot;quantity#quot;`; upsert rejects a shrink below it), so any decrement re-opens the row."
        datetime createdAt
        datetime updatedAt
    }
    Tag {
        string id PK
        string name
        string description
        string color
        tagusefor[] useFor
        string userId FK
        string organizationId FK
        datetime createdAt
        datetime updatedAt
    }
    Note {
        string id PK
        string content
        notetype type
        datetime createdAt
        datetime updatedAt
        string userId FK
        string assetId FK
    }
    BookingNote {
        string id PK
        string content
        notetype type
        datetime createdAt
        datetime updatedAt
        string userId FK
        string bookingId FK
    }
    LocationNote {
        string id PK
        string content
        notetype type
        datetime createdAt
        datetime updatedAt
        string userId FK
        string locationId FK
    }
    Qr {
        string id PK
        int version
        errorcorrection errorCorrection
        string assetId FK
        string kitId FK
        string userId FK
        string organizationId FK
        string batchId FK
        datetime createdAt
        datetime updatedAt
    }
    Barcode {
        string id PK
        string value UK
        barcodetype type
        string assetId FK
        string kitId FK
        string organizationId FK, UK
        datetime createdAt
        datetime updatedAt
    }
    PrintBatch {
        string id PK
        string name UK
        boolean printed
        datetime createdAt
        datetime updatedAt
    }
    ReportFound {
        string id PK
        string email
        string content
        datetime createdAt
        datetime updatedAt
        string assetId FK
        string kitId FK
    }
    Scan {
        string id PK
        string latitude
        string longitude
        string userAgent
        string userId FK
        string qrId FK
        string rawQrId
        scancodetype codeType
        string code
        string barcodeId FK
        string assetId FK
        string kitId FK
        string organizationId FK
        scansource source
        boolean manuallyGenerated
        datetime createdAt
        datetime updatedAt
    }
    Location {
        string id PK
        string name
        string description
        string address
        float latitude
        float longitude
        string imageUrl
        string thumbnailUrl
        string imageId FK, UK
        datetime createdAt
        datetime updatedAt
        string userId FK
        string organizationId FK
        string parentId FK
    }
    Role {
        string id PK
        roles name UK
        datetime createdAt
        datetime updatedAt
    }
    TeamMember {
        string id PK
        string name
        string organizationId FK
        string userId FK
        datetime createdAt
        datetime updatedAt
        datetime deletedAt
    }
    Custody {
        string id PK
        string teamMemberId FK
        string assetId FK
        string kitCustodyId FK
        int quantity
        string locationId FK
        boolean sourceUnknown
        datetime createdAt
        datetime updatedAt
    }
    Organization {
        string id PK
        string name
        organizationtype type
        string userId FK
        currency currency
        string imageId FK, UK
        boolean enabledSso
        string ssoDetailsId FK
        boolean selfServiceCanSeeCustody
        boolean selfServiceCanSeeBookings
        boolean baseUserCanSeeCustody
        boolean baseUserCanSeeBookings
        boolean managersCanEditAssets
        boolean barcodesEnabled
        datetime barcodesEnabledAt
        boolean usedBarcodeTrial
        boolean auditsEnabled
        datetime auditsEnabledAt
        boolean usedAuditTrial
        boolean workspaceDisabled
        datetime createdAt
        datetime updatedAt
        boolean hasSequentialIdsMigrated
        qriddisplaypreference qrIdDisplayPreference
        boolean showShelfBranding
        boolean showQrCodesOnPdfs
        string customEmailFooter
    }
    UserOrganization {
        string id PK
        string userId FK, UK
        string organizationId FK, UK
        organizationroles[] roles
        datetime createdAt
        datetime updatedAt
        string calendarTokenId UK "Secret, unguessable token for this member's iCal calendar-feed subscription in this workspace. Null until the user generates a feed; regenerating it rotates the value (invalidating any previously-shared subscription URL). `@unique` doubles as the lookup index for the public feed route."
    }
    SsoDetails {
        string id PK
        string domain
        string baseUserGroupId
        string selfServiceGroupId
        string adminGroupId
        string managerGroupId
        boolean requireSsoLogin
        datetime createdAt
        datetime updatedAt
    }
    ScimToken {
        string id PK
        string tokenHash UK
        string label
        string organizationId FK
        string createdById
        datetime createdAt
        datetime lastUsedAt
    }
    UserScimExternalId {
        string id PK
        string userId FK, UK
        string organizationId FK, UK
        string scimExternalId UK
        datetime createdAt
        datetime updatedAt
    }
    Tier {
        tierid id PK
        string name
        tierid tierLimitId FK, UK
        datetime createdAt
        datetime updatedAt
    }
    TierLimit {
        tierid id PK
        boolean canImportAssets
        boolean canExportAssets
        boolean canImportNRM
        boolean canHideShelfBranding
        boolean canUseReports
        boolean canUseRecurringReminders
        boolean canUsePushNotifications
        int maxCustomFields
        int maxOrganizations
        datetime createdAt
        datetime updatedAt
    }
    CustomTierLimit {
        string id PK
        string userId FK, UK
        boolean canImportAssets
        boolean canExportAssets
        boolean canImportNRM
        boolean canHideShelfBranding
        boolean canUseReports
        boolean canUseRecurringReminders
        boolean canUsePushNotifications
        int maxCustomFields
        int maxOrganizations
        boolean isEnterprise
        datetime createdAt
        datetime updatedAt
    }
    CustomField {
        string id PK
        string name
        string helpText
        boolean required
        boolean active
        customfieldtype type
        string[] options
        customfieldformat format "How a TEXT field's values are read. LINK requires a web link in every value; any other type stays PLAIN (enforced by a CHECK constraint)."
        string organizationId FK
        string userId FK
        datetime createdAt
        datetime updatedAt
        datetime deletedAt
    }
    AssetCustomFieldValue {
        string id PK
        json value
        string assetId FK
        string customFieldId FK
        datetime createdAt
        datetime updatedAt
    }
    AssetModelCustomFieldDefault {
        string id PK
        json value
        string assetModelId FK, UK
        string customFieldId UK
        datetime createdAt
        datetime updatedAt
    }
    Invite {
        string id PK
        string inviterId FK
        string organizationId FK
        string inviteeUserId FK
        string teamMemberId FK
        string inviteeEmail
        invitestatuses status
        string inviteCode
        organizationroles[] roles
        string inviteMessage
        datetime expiresAt
        datetime createdAt
        datetime updatedAt
    }
    Announcement {
        string id PK
        string name
        string content
        string link
        string linkText
        boolean published
        datetime createdAt
        datetime updatedAt
    }
    ConsumptionLog {
        string id PK
        string assetId FK
        consumptioncategory category
        int quantity
        string note
        string userId FK
        string bookingId FK
        string custodianId FK
        string bookingAssetId
        string locationId FK
        int stockChange
        datetime createdAt
    }
    BookingAsset {
        string id PK
        string bookingId FK
        string assetId FK
        int quantity
        string bookingModelRequestId
        string assetKitId
        string sourceKitId
        datetime checkedOutAt
        string checkedOutById
        datetime checkedInAt
        string checkedInById
        int checkedOutQuantity
        string sourceLocationId
    }
    Booking {
        string id PK
        string name
        bookingstatus status
        string description
        string activeSchedulerReference
        string creatorId FK
        string custodianUserId FK
        string custodianTeamMemberId FK
        string organizationId FK
        datetime createdAt
        datetime updatedAt
        datetime from
        datetime to
        datetime originalFrom
        datetime originalTo
        datetime autoArchivedAt
        boolean archivedWithoutCheckin
        string cancellationReason
        datetime approvedAt
        string approvedById FK
    }
    BookingSettings {
        string id PK
        int bufferStartTime
        boolean tagsRequired
        int maxBookingLength
        boolean maxBookingLengthSkipClosedDays
        boolean autoArchiveBookings
        int autoArchiveDays
        boolean autoArchiveExpiredReservations
        boolean requireApprovalForBase
        boolean requireApprovalForSelfService
        boolean requireExplicitCheckinForAdmin
        boolean requireExplicitCheckinForSelfService
        boolean requireExplicitCheckoutForAdmin
        boolean requireExplicitCheckoutForSelfService
        boolean countKitsAsSingleUnit
        boolean notifyBookingCreator
        boolean notifyAdminsOnNewBooking
        string organizationId FK, UK
        datetime createdAt
        datetime updatedAt
    }
    PartialBookingCheckin {
        string id PK
        string[] assetIds
        int checkinCount
        datetime checkinTimestamp
        string bookingId FK
        string checkedInById FK
        datetime createdAt
        datetime updatedAt
    }
    PartialBookingCheckout {
        string id PK
        string[] assetIds
        int[] quantities
        string[] bookingAssetIds
        int checkoutCount
        datetime checkoutTimestamp
        string bookingId FK
        string checkedOutById FK
        datetime createdAt
        datetime updatedAt
    }
    MobileAuthCode {
        string id PK
        string userId FK
        string codeHash UK
        string codeChallenge
        string sessionCiphertext
        datetime consumedAt
        datetime expiresAt
        datetime createdAt
    }
    PushToken {
        string id PK
        string token UK
        deviceplatform platform
        string deviceName
        string appVersion
        datetime lastSeenAt
        datetime disabledAt
        string userId FK
        datetime createdAt
        datetime updatedAt
    }
    NotificationPreference {
        string userId PK, FK
        boolean pushEnabled
        boolean bookings
        boolean audits
        boolean stockAndReminders
        datetime updatedAt
    }
    Kit {
        string id PK
        string name
        string description
        kitstatus status
        string image
        datetime imageExpiration
        string organizationId FK
        string createdById FK
        string categoryId FK
        datetime createdAt
        datetime updatedAt
        string locationId FK
    }
    KitCustody {
        string id PK
        string custodianId FK
        string kitId FK, UK
        datetime createdAt
        datetime updatedAt
    }
    AssetKit {
        string id PK
        string assetId FK, UK
        string kitId FK, UK
        string organizationId FK
        int quantity
        datetime createdAt
        datetime updatedAt
    }
    AssetLocation {
        string id PK
        string assetId FK
        string locationId FK
        string organizationId FK
        int quantity
        string assetKitId FK
        datetime createdAt
        datetime updatedAt
    }
    AssetReminder {
        string id PK
        string name
        string message
        datetime alertDateTime
        string activeSchedulerReference
        reminderrecurrenceunit recurrenceUnit
        int recurrenceInterval
        string recurrenceTimezone
        datetime recurrenceEndsAt
        string organizationId FK
        string assetId FK
        string createdById FK
        datetime createdAt
        datetime updatedAt
    }
    WorkingHours {
        string id PK
        boolean enabled
        json weeklySchedule
        string organizationId FK, UK
        datetime createdAt
        datetime updatedAt
    }
    WorkingHoursOverride {
        string id PK
        datetime date UK
        boolean isOpen
        string openTime
        string closeTime
        string reason
        string workingHoursId FK, UK
        datetime createdAt
        datetime updatedAt
    }
    Update {
        string id PK
        string title
        string content
        string url
        string imageUrl
        datetime publishDate
        updatestatus status
        organizationroles[] targetRoles
        int clickCount
        int viewCount
        string createdById FK
        datetime createdAt
        datetime updatedAt
    }
    UserUpdateRead {
        string id PK
        string userId FK, UK
        string updateId FK, UK
        datetime readAt
    }
    AuditNote {
        string id PK
        string content
        notetype type
        datetime createdAt
        datetime updatedAt
        string userId FK
        string auditSessionId FK
        string auditAssetId FK
    }
    TeamMemberNote {
        string id PK
        string content
        notetype type
        datetime createdAt
        datetime updatedAt
        string userId FK
        string teamMemberId FK
        string organizationId FK
    }
    AuditSession {
        string id PK
        string name
        string description
        string targetId
        auditstatus status
        json scopeMeta
        int expectedAssetCount
        int foundAssetCount
        int missingAssetCount
        int unexpectedAssetCount
        datetime startedAt
        datetime dueDate
        datetime completedAt
        datetime cancelledAt
        string activeSchedulerReference
        string createdById FK
        string organizationId FK
        datetime createdAt
        datetime updatedAt
    }
    AssetImport {
        string id PK
        assetimportstatus status
        assetimportphase phase
        string fileName
        int totalRows
        int processedRows "Rows committed so far. Advanced in the same transaction as each chunk."
        json rows "Parsed rows (`StoredImportRow[]`). Cleared when the import completes."
        json resolved "Name to id maps written by the PREPARE step (`ResolvedImportEntities`)."
        boolean canUseBarcodes
        string errorMessage "User-facing reason the import failed. Null unless status is FAILED."
        json warnings "Rows imported without a conflicting piece (`ImportWarning[]`), capped at 500."
        int warningCount "Exact number of warnings, including any beyond the stored cap."
        datetime lockedUntil "A step holds the import until this time. Null when no step is running."
        datetime confirmedAt
        datetime startedAt
        datetime completedAt
        datetime createdAt
        datetime updatedAt
        string createdById FK
        string organizationId FK
    }
    LinkPreview {
        string id PK
        string url UK "`LinkKey` from @shelf/links, never a raw field value. At most 2048 characters (CHECK constraint in the migration)."
        string pageTitle
        linkpreviewoutcome outcome "Null until the first fetch finishes."
        int attempts
        datetime fetchedAt
        datetime nextFetchAt "When the link may be fetched next; null means never again. A claim moves it a few minutes ahead, which is also the lease on a fetch in flight."
        string customTitle "The name a person gave the link in this workspace; wins over pageTitle."
        datetime createdAt
        datetime updatedAt
        string organizationId FK, UK
    }
    AuditAssignment {
        string id PK
        string auditSessionId FK, UK
        string userId FK, UK
        auditassignmentrole role
        datetime createdAt
        datetime updatedAt
    }
    AuditAsset {
        string id PK
        string auditSessionId FK, UK
        string assetId FK, UK
        boolean expected
        auditassetstatus status
        string scannedById FK
        datetime scannedAt
        json metadata
        int expectedQuantity "Units the audit expects to find, snapshotted when the row is created: the units in the audit's scope (placed at its locations, held in its kits, or held by its person; at most `Asset.quantity`), or `Asset.quantity` when none are in scope or the audit has no scope. Null for INDIVIDUAL assets, which count as one.  A snapshot, not a live read: stock that moves during the audit must not move the target the auditor is counting against."
        int foundQuantity "Units the auditor counted. Null until a quantity-tracked asset is counted, and always null for INDIVIDUAL assets. A FOUND row with a count below `expectedQuantity` is a shortfall: the status stays FOUND (the pool was located), the gap is this number.  Both columns are non-negative, enforced by CHECK constraints in raw SQL (migration `20261104120000_add_audit_asset_quantities`)."
        datetime createdAt
        datetime updatedAt
    }
    AuditScan {
        string id PK
        string auditSessionId FK
        string auditAssetId FK
        string assetId FK
        string scannedById FK
        string code
        string assetTitle "What was scanned, recorded BY VALUE at scan time.  `asset` and `auditAsset` are both SetNull/Cascade, so deleting an asset leaves this row pointing at nothing: no name, and no way to tell whether it had been expected. An audit is a historical record, and a record whose entries can silently lose their meaning is not one — so the two descriptive facts are snapshotted here rather than only referenced.  Readers prefer the LIVE asset when it still exists (a rename should show the current name) and fall back to these.  Null means the snapshot could not be taken: scans whose asset had already been deleted before these columns were backfilled. Nothing can recover those, and they degrade to the scanned `code`, which also survives. Every other row carries a snapshot, so null is never merely #quot;old#quot;."
        boolean wasExpected
        json metadata
        datetime scannedAt
        datetime createdAt
    }
    AuditImage {
        string id PK
        string imageUrl
        string thumbnailUrl
        string description
        string auditSessionId FK
        string auditAssetId FK
        string uploadedById FK
        string organizationId FK
        datetime capturedAt "When the photo was taken, from its EXIF data (or the phone app's copy of it). Null when the photo carried no usable date. EXIF can be edited, so this is the best signal available, not proof."
        auditimagecapturesource captureSource "How the photo reached Shelf. Null only for photos uploaded before this field existed; a phone upload with no or an unknown source is PHONE."
        datetime createdAt
        datetime updatedAt
    }
    RoleChangeLog {
        string id PK
        organizationroles previousRole
        organizationroles newRole
        rolechangesource source "How the change was made. SSO rows name the member themself as `changedBy`."
        datetime createdAt
        string userId FK
        string changedById FK
        string organizationId FK
    }
    ActivityEvent {
        string id PK
        string organizationId FK
        datetime occurredAt
        string actorUserId
        json actorSnapshot
        activityaction action
        activityentity entityType
        string entityId
        string assetId
        string bookingId
        string auditSessionId
        string auditAssetId
        string kitId
        string locationId
        string teamMemberId
        string targetUserId
        string field
        json fromValue
        json toValue
        json meta
    }
```