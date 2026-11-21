package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class QrEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, QrEntity>(QrTable)

    var version by QrTable.version
    var errorCorrection by QrTable.errorCorrection
    var asset by AssetEntity optionalReferencedOn QrTable.assetId
    var assetId by QrTable.assetId
    var kit by KitEntity optionalReferencedOn QrTable.kitId
    var kitId by QrTable.kitId
    var user by UserEntity optionalReferencedOn QrTable.userId
    var userId by QrTable.userId
    var organization by OrganizationEntity optionalReferencedOn QrTable.organizationId
    var organizationId by QrTable.organizationId
    var batch by PrintBatchEntity optionalReferencedOn QrTable.batchId
    var batchId by QrTable.batchId
    var createdAt by QrTable.createdAt
    var updatedAt by QrTable.updatedAt
    val scans by ScanEntity optionalReferrersOn ScanTable.qrId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(QrTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
