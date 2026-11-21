package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class LocationEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, LocationEntity>(LocationTable)

    var name by LocationTable.name
    var description by LocationTable.description
    var address by LocationTable.address
    var latitude by LocationTable.latitude
    var longitude by LocationTable.longitude
    var imageUrl by LocationTable.imageUrl
    var thumbnailUrl by LocationTable.thumbnailUrl
    var image by ImageEntity optionalReferencedOn LocationTable.imageId
    var imageId by LocationTable.imageId
    var createdAt by LocationTable.createdAt
    var updatedAt by LocationTable.updatedAt
    var user by UserEntity referencedOn LocationTable.userId
    var userId by LocationTable.userId
    var organization by OrganizationEntity referencedOn LocationTable.organizationId
    var organizationId by LocationTable.organizationId
    val assetLocations by AssetLocationEntity referrersOn AssetLocationTable.locationId
    val kits by KitEntity optionalReferrersOn KitTable.locationId
    val notes by LocationNoteEntity referrersOn LocationNoteTable.locationId
    val custodies by CustodyEntity optionalReferrersOn CustodyTable.locationId
    val consumptionLogs by ConsumptionLogEntity optionalReferrersOn ConsumptionLogTable.locationId
    var parentId by LocationTable.parentId
    var parent by LocationEntity optionalReferencedOn LocationTable.parentId
    val children by LocationEntity optionalReferrersOn LocationTable.parentId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(LocationTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
