package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class AssetCustomFieldValueEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AssetCustomFieldValueEntity>(AssetCustomFieldValueTable)

    var value by AssetCustomFieldValueTable.value
    var asset by AssetEntity referencedOn AssetCustomFieldValueTable.assetId
    var assetId by AssetCustomFieldValueTable.assetId
    var customField by CustomFieldEntity referencedOn AssetCustomFieldValueTable.customFieldId
    var customFieldId by AssetCustomFieldValueTable.customFieldId
    var createdAt by AssetCustomFieldValueTable.createdAt
    var updatedAt by AssetCustomFieldValueTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AssetCustomFieldValueTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
