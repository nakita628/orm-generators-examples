package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class LocationNoteEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, LocationNoteEntity>(LocationNoteTable)

    var content by LocationNoteTable.content
    var type by LocationNoteTable.type
    var createdAt by LocationNoteTable.createdAt
    var updatedAt by LocationNoteTable.updatedAt
    var user by UserEntity optionalReferencedOn LocationNoteTable.userId
    var userId by LocationNoteTable.userId
    var location by LocationEntity referencedOn LocationNoteTable.locationId
    var locationId by LocationNoteTable.locationId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(LocationNoteTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
