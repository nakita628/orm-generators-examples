package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class UserEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, UserEntity>(UserTable)

    var email by UserTable.email
    var username by UserTable.username
    var firstName by UserTable.firstName
    var lastName by UserTable.lastName
    var displayName by UserTable.displayName
    var profilePicture by UserTable.profilePicture
    var usedFreeTrial by UserTable.usedFreeTrial
    var onboarded by UserTable.onboarded
    var customerId by UserTable.customerId
    var sso by UserTable.sso
    var createdWithInvite by UserTable.createdWithInvite
    var skipSubscriptionCheck by UserTable.skipSubscriptionCheck
    var hasUnpaidInvoice by UserTable.hasUnpaidInvoice
    var warnForNoPaymentMethod by UserTable.warnForNoPaymentMethod
    var lastSelectedOrganization by OrganizationEntity optionalReferencedOn UserTable.lastSelectedOrganizationId
    var lastSelectedOrganizationId by UserTable.lastSelectedOrganizationId
    var createdAt by UserTable.createdAt
    var updatedAt by UserTable.updatedAt
    var deletedAt by UserTable.deletedAt
    var lastMobileActiveAt by UserTable.lastMobileActiveAt
    var dateFormat by UserTable.dateFormat
    var timeFormat by UserTable.timeFormat
    var weekStart by UserTable.weekStart
    var timeZone by UserTable.timeZone
    val assets by AssetEntity referrersOn AssetTable.userId
    val categories by CategoryEntity referrersOn CategoryTable.userId
    val notes by NoteEntity optionalReferrersOn NoteTable.userId
    val bookingNotes by BookingNoteEntity optionalReferrersOn BookingNoteTable.userId
    val auditNotes by AuditNoteEntity optionalReferrersOn AuditNoteTable.userId
    val locationNotes by LocationNoteEntity optionalReferrersOn LocationNoteTable.userId
    val qrCodes by QrEntity optionalReferrersOn QrTable.userId
    val scans by ScanEntity optionalReferrersOn ScanTable.userId
    val tags by TagEntity referrersOn TagTable.userId
    var roles by RoleEntity via RoleToUserTable
    val locations by LocationEntity referrersOn LocationTable.userId
    val images by ImageEntity referrersOn ImageTable.userId
    val organizations by OrganizationEntity referrersOn OrganizationTable.userId
    val customFields by CustomFieldEntity referrersOn CustomFieldTable.userId
    val sentInvites by InviteEntity referrersOn InviteTable.inviterId
    val receivedInvites by InviteEntity optionalReferrersOn InviteTable.inviteeUserId
    val teamMembers by TeamMemberEntity optionalReferrersOn TeamMemberTable.userId
    val userOrganizations by UserOrganizationEntity referrersOn UserOrganizationTable.userId
    val bookings by BookingEntity referrersOn BookingTable.creatorId
    val custodies by BookingEntity optionalReferrersOn BookingTable.custodianUserId
    val approvedBookings by BookingEntity optionalReferrersOn BookingTable.approvedById
    val createdKits by KitEntity referrersOn KitTable.createdById
    var tierId by UserTable.tierId
    var tier by TierEntity referencedOn UserTable.tierId
    val assetReminders by AssetReminderEntity referrersOn AssetReminderTable.createdById
    val contact by UserContactEntity optionalBackReferencedOn UserContactTable.userId
    val businessIntel by UserBusinessIntelEntity optionalBackReferencedOn UserBusinessIntelTable.userId
    val auditsCreated by AuditSessionEntity referrersOn AuditSessionTable.createdById
    val assetImports by AssetImportEntity referrersOn AssetImportTable.createdById
    val auditAssignments by AuditAssignmentEntity referrersOn AuditAssignmentTable.userId
    val auditAssetsScanned by AuditAssetEntity optionalReferrersOn AuditAssetTable.scannedById
    val auditScans by AuditScanEntity optionalReferrersOn AuditScanTable.scannedById
    val auditImagesUploaded by AuditImageEntity optionalReferrersOn AuditImageTable.uploadedById
    var referralSource by UserTable.referralSource
    val assetIndexSettings by AssetIndexSettingsEntity referrersOn AssetIndexSettingsTable.userId
    val partialCheckins by PartialBookingCheckinEntity referrersOn PartialBookingCheckinTable.checkedInById
    val partialCheckouts by PartialBookingCheckoutEntity referrersOn PartialBookingCheckoutTable.checkedOutById
    val mobileAuthCodes by MobileAuthCodeEntity referrersOn MobileAuthCodeTable.userId
    val pushTokens by PushTokenEntity referrersOn PushTokenTable.userId
    val notificationPreference by NotificationPreferenceEntity optionalBackReferencedOn NotificationPreferenceTable.id
    val createdUpdates by UpdateEntity referrersOn UpdateTable.createdById
    val updateReads by UserUpdateReadEntity referrersOn UserUpdateReadTable.userId
    val customTierLimit by CustomTierLimitEntity optionalBackReferencedOn CustomTierLimitTable.userId
    val assetFilterPresets by AssetFilterPresetEntity referrersOn AssetFilterPresetTable.ownerId
    val assetModels by AssetModelEntity referrersOn AssetModelTable.userId
    val consumptionLogs by ConsumptionLogEntity referrersOn ConsumptionLogTable.userId
    val scimExternalIds by UserScimExternalIdEntity referrersOn UserScimExternalIdTable.userId
    val roleChangesReceived by RoleChangeLogEntity referrersOn RoleChangeLogTable.userId
    val roleChangesPerformed by RoleChangeLogEntity referrersOn RoleChangeLogTable.changedById
    val teamMemberNotesAuthored by TeamMemberNoteEntity optionalReferrersOn TeamMemberNoteTable.userId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(UserTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
