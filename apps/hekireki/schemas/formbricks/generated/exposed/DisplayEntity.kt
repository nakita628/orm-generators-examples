package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Records when a survey is shown to a user.
 * Tracks survey display history and response status.
 *
 * @property id - Unique identifier for the display event
 * @property survey - The survey that was displayed
 * @property contact - The contact who saw the survey
 * @property response - The associated response if one exists
 * @property creativeId - The creative that opened the survey. No foreign key, so the record survives the creative's deletion.
 */
class DisplayEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, DisplayEntity>(DisplayTable)

    var createdAt by DisplayTable.createdAt
    var updatedAt by DisplayTable.updatedAt
    var survey by SurveyEntity referencedOn DisplayTable.surveyId
    var surveyId by DisplayTable.surveyId
    var contact by ContactEntity optionalReferencedOn DisplayTable.contactId
    var contactId by DisplayTable.contactId
    val response by ResponseEntity optionalBackReferencedOn ResponseTable.displayId
    var creativeId by DisplayTable.creativeId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(DisplayTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
