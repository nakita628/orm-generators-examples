package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * A custom field value an AssetModel gives the assets that join it.
 *
 * `value` has the exact shape of `AssetCustomFieldValue.value`, so filling an
 * asset is a plain copy. Rows cascade away with their model or field.
 */
class AssetModelCustomFieldDefaultEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AssetModelCustomFieldDefaultEntity>(AssetModelCustomFieldDefaultTable)

    var value by AssetModelCustomFieldDefaultTable.value
    var assetModel by AssetModelEntity referencedOn AssetModelCustomFieldDefaultTable.assetModelId
    var assetModelId by AssetModelCustomFieldDefaultTable.assetModelId
    var customFieldId by AssetModelCustomFieldDefaultTable.customFieldId
    var createdAt by AssetModelCustomFieldDefaultTable.createdAt
    var updatedAt by AssetModelCustomFieldDefaultTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AssetModelCustomFieldDefaultTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
