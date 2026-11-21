package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Represents an attribute value associated with a contact.
 * Used to store custom properties and characteristics of contacts.
 *
 * @property id - Unique identifier for the attribute
 * @property attributeKey - Reference to the attribute definition
 * @property contact - The contact this attribute belongs to
 * @property value - The string value of the attribute (used for string type + backwards compatibility)
 * @property valueNumber - Native numeric storage for number type attributes
 * @property valueDate - Native date storage for date type attributes
 */
class ContactAttributeEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, ContactAttributeEntity>(ContactAttributeTable)

    var createdAt by ContactAttributeTable.createdAt
    var updatedAt by ContactAttributeTable.updatedAt
    var attributeKey by ContactAttributeKeyEntity referencedOn ContactAttributeTable.attributeKeyId
    var attributeKeyId by ContactAttributeTable.attributeKeyId
    var contact by ContactEntity referencedOn ContactAttributeTable.contactId
    var contactId by ContactAttributeTable.contactId
    var value by ContactAttributeTable.value
    var valueNumber by ContactAttributeTable.valueNumber
    var valueDate by ContactAttributeTable.valueDate

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(ContactAttributeTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
