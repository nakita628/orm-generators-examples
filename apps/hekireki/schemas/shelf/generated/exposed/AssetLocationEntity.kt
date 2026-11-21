package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class AssetLocationEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AssetLocationEntity>(AssetLocationTable)

    var asset by AssetEntity referencedOn AssetLocationTable.assetId
    var assetId by AssetLocationTable.assetId
    var location by LocationEntity referencedOn AssetLocationTable.locationId
    var locationId by AssetLocationTable.locationId
    var organization by OrganizationEntity referencedOn AssetLocationTable.organizationId
    var organizationId by AssetLocationTable.organizationId
    var quantity by AssetLocationTable.quantity
    var assetKit by AssetKitEntity optionalReferencedOn AssetLocationTable.assetKitId
    var assetKitId by AssetLocationTable.assetKitId
    var createdAt by AssetLocationTable.createdAt
    var updatedAt by AssetLocationTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AssetLocationTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
