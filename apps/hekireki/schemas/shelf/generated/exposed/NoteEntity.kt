package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class NoteEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, NoteEntity>(NoteTable)

    var content by NoteTable.content
    var type by NoteTable.type
    var createdAt by NoteTable.createdAt
    var updatedAt by NoteTable.updatedAt
    var user by UserEntity optionalReferencedOn NoteTable.userId
    var userId by NoteTable.userId
    var asset by AssetEntity referencedOn NoteTable.assetId
    var assetId by NoteTable.assetId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(NoteTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
