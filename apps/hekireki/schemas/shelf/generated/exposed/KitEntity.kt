package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class KitEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, KitEntity>(KitTable)

    var name by KitTable.name
    var description by KitTable.description
    var status by KitTable.status
    var image by KitTable.image
    var imageExpiration by KitTable.imageExpiration
    val assetKits by AssetKitEntity referrersOn AssetKitTable.kitId
    val custody by KitCustodyEntity optionalBackReferencedOn KitCustodyTable.kitId
    val qrCodes by QrEntity optionalReferrersOn QrTable.kitId
    val barcodes by BarcodeEntity optionalReferrersOn BarcodeTable.kitId
    val reports by ReportFoundEntity optionalReferrersOn ReportFoundTable.kitId
    val scans by ScanEntity optionalReferrersOn ScanTable.kitId
    var organization by OrganizationEntity referencedOn KitTable.organizationId
    var organizationId by KitTable.organizationId
    var createdBy by UserEntity referencedOn KitTable.createdById
    var createdById by KitTable.createdById
    var category by CategoryEntity optionalReferencedOn KitTable.categoryId
    var categoryId by KitTable.categoryId
    var createdAt by KitTable.createdAt
    var updatedAt by KitTable.updatedAt
    var location by LocationEntity optionalReferencedOn KitTable.locationId
    var locationId by KitTable.locationId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(KitTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
