package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class AssetIndexSettingsEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AssetIndexSettingsEntity>(AssetIndexSettingsTable)

    var user by UserEntity referencedOn AssetIndexSettingsTable.userId
    var userId by AssetIndexSettingsTable.userId
    var organizationId by AssetIndexSettingsTable.organizationId
    var organization by OrganizationEntity referencedOn AssetIndexSettingsTable.organizationId
    var mode by AssetIndexSettingsTable.mode
    var columns by AssetIndexSettingsTable.columns1
    var freezeColumn by AssetIndexSettingsTable.freezeColumn
    var showAssetImage by AssetIndexSettingsTable.showAssetImage
    var createdAt by AssetIndexSettingsTable.createdAt
    var updatedAt by AssetIndexSettingsTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AssetIndexSettingsTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
