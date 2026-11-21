package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class AssetFilterPresetEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AssetFilterPresetEntity>(AssetFilterPresetTable)

    var organizationId by AssetFilterPresetTable.organizationId
    var ownerId by AssetFilterPresetTable.ownerId
    var name by AssetFilterPresetTable.name
    var query by AssetFilterPresetTable.query
    var starred by AssetFilterPresetTable.starred
    var createdAt by AssetFilterPresetTable.createdAt
    var updatedAt by AssetFilterPresetTable.updatedAt
    var organization by OrganizationEntity referencedOn AssetFilterPresetTable.organizationId
    var owner by UserEntity referencedOn AssetFilterPresetTable.ownerId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AssetFilterPresetTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
