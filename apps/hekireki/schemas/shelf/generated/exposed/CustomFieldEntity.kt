package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class CustomFieldEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, CustomFieldEntity>(CustomFieldTable)

    var name by CustomFieldTable.name
    var helpText by CustomFieldTable.helpText
    var required by CustomFieldTable.required
    var active by CustomFieldTable.active
    var type by CustomFieldTable.type
    var options by CustomFieldTable.options1

    /**
     * How a TEXT field's values are read. LINK requires a web link in every
     * value; any other type stays PLAIN (enforced by a CHECK constraint).
     */
    var format by CustomFieldTable.format
    var organization by OrganizationEntity referencedOn CustomFieldTable.organizationId
    var organizationId by CustomFieldTable.organizationId
    var createdBy by UserEntity referencedOn CustomFieldTable.userId
    var userId by CustomFieldTable.userId
    var categories by CategoryEntity via CategoryToCustomFieldTable
    var createdAt by CustomFieldTable.createdAt
    var updatedAt by CustomFieldTable.updatedAt
    var deletedAt by CustomFieldTable.deletedAt
    val assetCustomFieldsValues by AssetCustomFieldValueEntity referrersOn AssetCustomFieldValueTable.customFieldId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(CustomFieldTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
