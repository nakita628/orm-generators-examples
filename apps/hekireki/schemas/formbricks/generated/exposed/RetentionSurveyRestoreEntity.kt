package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * When a survey was last restored from the archive. A restore starts the surveys policy's "no response"
 * and "no change" periods over (`getSurveyRetentionClock`); only ever moves forward. A table of its own
 * so retention adds nothing to the survey row.
 */
class RetentionSurveyRestoreEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, RetentionSurveyRestoreEntity>(RetentionSurveyRestoreTable)

    val survey by SurveyEntity referencedOn RetentionSurveyRestoreTable.id
    var restoredAt by RetentionSurveyRestoreTable.restoredAt
}
