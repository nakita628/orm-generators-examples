package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class OrganizationEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, OrganizationEntity>(OrganizationTable)

    var name by OrganizationTable.name
    var type by OrganizationTable.type
    var owner by UserEntity referencedOn OrganizationTable.userId
    var userId by OrganizationTable.userId
    var currency by OrganizationTable.currency
    val members by TeamMemberEntity referrersOn TeamMemberTable.organizationId
    val assets by AssetEntity referrersOn AssetTable.organizationId
    val assetKits by AssetKitEntity referrersOn AssetKitTable.organizationId
    val assetLocations by AssetLocationEntity referrersOn AssetLocationTable.organizationId
    val scans by ScanEntity referrersOn ScanTable.organizationId
    var image by ImageEntity optionalReferencedOn OrganizationTable.imageId
    var imageId by OrganizationTable.imageId
    var enabledSso by OrganizationTable.enabledSso
    var ssoDetailsId by OrganizationTable.ssoDetailsId
    var ssoDetails by SsoDetailsEntity optionalReferencedOn OrganizationTable.ssoDetailsId
    val workingHours by WorkingHoursEntity optionalBackReferencedOn WorkingHoursTable.organizationId
    val bookingSettings by BookingSettingsEntity optionalBackReferencedOn BookingSettingsTable.organizationId
    var selfServiceCanSeeCustody by OrganizationTable.selfServiceCanSeeCustody
    var selfServiceCanSeeBookings by OrganizationTable.selfServiceCanSeeBookings
    var baseUserCanSeeCustody by OrganizationTable.baseUserCanSeeCustody
    var baseUserCanSeeBookings by OrganizationTable.baseUserCanSeeBookings
    var managersCanEditAssets by OrganizationTable.managersCanEditAssets
    var barcodesEnabled by OrganizationTable.barcodesEnabled
    var barcodesEnabledAt by OrganizationTable.barcodesEnabledAt
    var usedBarcodeTrial by OrganizationTable.usedBarcodeTrial
    var auditsEnabled by OrganizationTable.auditsEnabled
    var auditsEnabledAt by OrganizationTable.auditsEnabledAt
    var usedAuditTrial by OrganizationTable.usedAuditTrial
    var workspaceDisabled by OrganizationTable.workspaceDisabled
    val locations by LocationEntity referrersOn LocationTable.organizationId
    val categories by CategoryEntity referrersOn CategoryTable.organizationId
    val tags by TagEntity referrersOn TagTable.organizationId
    val qrCodes by QrEntity optionalReferrersOn QrTable.organizationId
    val barcodes by BarcodeEntity referrersOn BarcodeTable.organizationId
    val invites by InviteEntity referrersOn InviteTable.organizationId
    val userOrganizations by UserOrganizationEntity referrersOn UserOrganizationTable.organizationId
    var createdAt by OrganizationTable.createdAt
    var updatedAt by OrganizationTable.updatedAt
    val customFields by CustomFieldEntity referrersOn CustomFieldTable.organizationId
    val images by ImageEntity referrersOn ImageTable.ownerOrgId
    val bookings by BookingEntity referrersOn BookingTable.organizationId
    val kits by KitEntity referrersOn KitTable.organizationId
    val assetModels by AssetModelEntity referrersOn AssetModelTable.organizationId
    val assetIndexSettings by AssetIndexSettingsEntity referrersOn AssetIndexSettingsTable.organizationId
    val assetReminders by AssetReminderEntity referrersOn AssetReminderTable.organizationId
    val assetFilterPresets by AssetFilterPresetEntity referrersOn AssetFilterPresetTable.organizationId
    val auditSessions by AuditSessionEntity referrersOn AuditSessionTable.organizationId
    val auditImages by AuditImageEntity referrersOn AuditImageTable.organizationId
    val assetImports by AssetImportEntity referrersOn AssetImportTable.organizationId
    val linkPreviews by LinkPreviewEntity referrersOn LinkPreviewTable.organizationId
    val pdfTemplates by PdfTemplateEntity referrersOn PdfTemplateTable.organizationId
    val usersWithLastSelected by UserEntity optionalReferrersOn UserTable.lastSelectedOrganizationId
    val scimTokens by ScimTokenEntity referrersOn ScimTokenTable.organizationId
    val scimExternalIds by UserScimExternalIdEntity referrersOn UserScimExternalIdTable.organizationId
    val roleChangeLogs by RoleChangeLogEntity referrersOn RoleChangeLogTable.organizationId
    val teamMemberNotes by TeamMemberNoteEntity referrersOn TeamMemberNoteTable.organizationId
    val activityEvents by ActivityEventEntity referrersOn ActivityEventTable.organizationId
    var hasSequentialIdsMigrated by OrganizationTable.hasSequentialIdsMigrated
    var qrIdDisplayPreference by OrganizationTable.qrIdDisplayPreference
    var showShelfBranding by OrganizationTable.showShelfBranding
    var showQrCodesOnPdfs by OrganizationTable.showQrCodesOnPdfs
    var customEmailFooter by OrganizationTable.customEmailFooter

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(OrganizationTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
