package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Represents a person or user who can receive and respond to surveys.
 * Contacts are workspace-specific and can have multiple attributes and responses.
 *
 * @property id - Unique identifier for the contact
 * @property workspace - The workspace this contact belongs to
 * @property responses - Survey responses from this contact
 * @property attributes - Custom attributes associated with this contact
 * @property displays - Record of surveys shown to this contact
 */
class ContactEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, ContactEntity>(ContactTable)

    var createdAt by ContactTable.createdAt
    var updatedAt by ContactTable.updatedAt
    var workspace by WorkspaceEntity referencedOn ContactTable.workspaceId
    var workspaceId by ContactTable.workspaceId
    val responses by ResponseEntity optionalReferrersOn ResponseTable.contactId
    val attributes by ContactAttributeEntity referrersOn ContactAttributeTable.contactId
    val displays by DisplayEntity optionalReferrersOn DisplayTable.contactId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(ContactTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
