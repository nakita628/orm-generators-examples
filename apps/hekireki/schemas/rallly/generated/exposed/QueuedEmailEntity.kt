package rallly.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * An email waiting to be sent by the queue, one row per recipient. The
 * email is built when it is sent, from the row its subject points at.
 */
class QueuedEmailEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, QueuedEmailEntity>(QueuedEmailTable)

    /**
     * Which handler builds the email. A string rather than an enum so a new
     * kind needs no migration; the claim only takes kinds the build knows.
     */
    var kind by QueuedEmailTable.kind

    /**
     * The row the email is about, interpreted by the kind. Not a foreign key:
     * kinds point at different tables, and a subject deleted before the send
     * makes its handler skip the email.
     */
    var subjectId by QueuedEmailTable.subjectId

    /**
     * Groups the emails one action queued, so the action can send its own
     * first batch at once and its progress can be read back.
     */
    var batchId by QueuedEmailTable.batchId

    /**
     * The account whose action queued the email. Banning it skips whatever
     * is still pending.
     */
    var userId by QueuedEmailTable.userId
    var status by QueuedEmailTable.status

    /** Not claimed before this time. Unset sends at the next run. */
    var sendAfter by QueuedEmailTable.sendAfter
    var attempts by QueuedEmailTable.attempts
    var claimedAt by QueuedEmailTable.claimedAt
    var sentAt by QueuedEmailTable.sentAt
    var lastError by QueuedEmailTable.lastError
    var createdAt by QueuedEmailTable.createdAt
    var updatedAt by QueuedEmailTable.updatedAt
    var user by UserEntity referencedOn QueuedEmailTable.userId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(QueuedEmailTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
