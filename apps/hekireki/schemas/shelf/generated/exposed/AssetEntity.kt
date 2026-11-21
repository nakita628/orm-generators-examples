package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class AssetEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AssetEntity>(AssetTable)

    var title by AssetTable.title
    var description by AssetTable.description
    var mainImage by AssetTable.mainImage
    var thumbnailImage by AssetTable.thumbnailImage
    var mainImageExpiration by AssetTable.mainImageExpiration
    var status by AssetTable.status
    var valuation by AssetTable.valuation
    var availableToBook by AssetTable.availableToBook
    var sequentialId by AssetTable.sequentialId
    var archivedAt by AssetTable.archivedAt
    var type by AssetTable.type
    var assetModelId by AssetTable.assetModelId
    var assetModel by AssetModelEntity optionalReferencedOn AssetTable.assetModelId
    var quantity by AssetTable.quantity
    var minQuantity by AssetTable.minQuantity
    var consumptionType by AssetTable.consumptionType
    var unitOfMeasure by AssetTable.unitOfMeasure
    var lowStockNotifiedAt by AssetTable.lowStockNotifiedAt
    var stockLedgerStartedAt by AssetTable.stockLedgerStartedAt
    var createdAt by AssetTable.createdAt
    var updatedAt by AssetTable.updatedAt
    var user by UserEntity referencedOn AssetTable.userId
    var userId by AssetTable.userId
    var organization by OrganizationEntity referencedOn AssetTable.organizationId
    var organizationId by AssetTable.organizationId
    var category by CategoryEntity optionalReferencedOn AssetTable.categoryId
    var categoryId by AssetTable.categoryId
    val custody by CustodyEntity referrersOn CustodyTable.assetId
    val notes by NoteEntity referrersOn NoteTable.assetId
    val qrCodes by QrEntity optionalReferrersOn QrTable.assetId
    val barcodes by BarcodeEntity optionalReferrersOn BarcodeTable.assetId
    val reports by ReportFoundEntity optionalReferrersOn ReportFoundTable.assetId
    var tags by TagEntity via AssetToTagTable
    val customFields by AssetCustomFieldValueEntity referrersOn AssetCustomFieldValueTable.assetId
    val bookingAssets by BookingAssetEntity referrersOn BookingAssetTable.assetId
    val consumptionLogs by ConsumptionLogEntity referrersOn ConsumptionLogTable.assetId
    val reminders by AssetReminderEntity referrersOn AssetReminderTable.assetId
    val auditEntries by AuditAssetEntity referrersOn AuditAssetTable.assetId
    val auditScans by AuditScanEntity optionalReferrersOn AuditScanTable.assetId
    val scans by ScanEntity optionalReferrersOn ScanTable.assetId
    val assetKits by AssetKitEntity referrersOn AssetKitTable.assetId
    val assetLocations by AssetLocationEntity referrersOn AssetLocationTable.assetId
    var preferredBarcodeId by AssetTable.preferredBarcodeId
    var preferredBarcode by BarcodeEntity optionalReferencedOn AssetTable.preferredBarcodeId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AssetTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
