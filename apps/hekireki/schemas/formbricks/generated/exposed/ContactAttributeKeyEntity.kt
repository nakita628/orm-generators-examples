package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Defines the possible attributes that can be assigned to contacts.
 * Acts as a schema for contact attributes within a workspace.
 *
 * @property id - Unique identifier for the attribute key
 * @property isUnique - Whether the attribute must have unique values across contacts
 * @property key - The attribute identifier used in the system
 * @property name - Display name for the attribute
 * @property type - Whether this is a default or custom attribute
 * @property dataType - The data type of the attribute (string, number, date)
 * @property workspace - The workspace this attribute belongs to
 */
class ContactAttributeKeyEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, ContactAttributeKeyEntity>(ContactAttributeKeyTable)

    var createdAt by ContactAttributeKeyTable.createdAt
    var updatedAt by ContactAttributeKeyTable.updatedAt
    var isUnique by ContactAttributeKeyTable.isUnique
    var key by ContactAttributeKeyTable.key
    var name by ContactAttributeKeyTable.name
    var description by ContactAttributeKeyTable.description
    var type by ContactAttributeKeyTable.type
    var dataType by ContactAttributeKeyTable.dataType
    var workspace by WorkspaceEntity referencedOn ContactAttributeKeyTable.workspaceId
    var workspaceId by ContactAttributeKeyTable.workspaceId
    val attributes by ContactAttributeEntity referrersOn ContactAttributeTable.attributeKeyId
    val attributeFilters by SurveyAttributeFilterEntity referrersOn SurveyAttributeFilterTable.attributeKeyId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(ContactAttributeKeyTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
