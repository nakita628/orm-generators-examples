package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class BookingNoteEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, BookingNoteEntity>(BookingNoteTable)

    var content by BookingNoteTable.content
    var type by BookingNoteTable.type
    var createdAt by BookingNoteTable.createdAt
    var updatedAt by BookingNoteTable.updatedAt
    var user by UserEntity optionalReferencedOn BookingNoteTable.userId
    var userId by BookingNoteTable.userId
    var booking by BookingEntity referencedOn BookingNoteTable.bookingId
    var bookingId by BookingNoteTable.bookingId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(BookingNoteTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
