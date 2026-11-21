package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class BarcodeEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, BarcodeEntity>(BarcodeTable)

    var value by BarcodeTable.value
    var type by BarcodeTable.type
    var asset by AssetEntity optionalReferencedOn BarcodeTable.assetId
    var assetId by BarcodeTable.assetId
    val preferredByAsset by AssetEntity optionalBackReferencedOn AssetTable.preferredBarcodeId
    val scans by ScanEntity optionalReferrersOn ScanTable.barcodeId
    var kit by KitEntity optionalReferencedOn BarcodeTable.kitId
    var kitId by BarcodeTable.kitId
    var organization by OrganizationEntity referencedOn BarcodeTable.organizationId
    var organizationId by BarcodeTable.organizationId
    var createdAt by BarcodeTable.createdAt
    var updatedAt by BarcodeTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(BarcodeTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
