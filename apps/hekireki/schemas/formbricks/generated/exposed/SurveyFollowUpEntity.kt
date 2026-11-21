package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Defines follow-up actions for survey responses.
 * Enables automated actions based on specific survey response conditions.
 *
 * @property id - Unique identifier for the follow-up
 * @property survey - The associated survey
 * @property name - Display name for the follow-up
 * @property trigger - Conditions that activate the follow-up
 * @property action - Actions to take when triggered
 */
class SurveyFollowUpEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, SurveyFollowUpEntity>(SurveyFollowUpTable)

    var createdAt by SurveyFollowUpTable.createdAt
    var updatedAt by SurveyFollowUpTable.updatedAt
    var survey by SurveyEntity referencedOn SurveyFollowUpTable.surveyId
    var surveyId by SurveyFollowUpTable.surveyId
    var name by SurveyFollowUpTable.name

    /** [SurveyFollowUpTrigger] */
    var trigger by SurveyFollowUpTable.trigger

    /** [SurveyFollowUpAction] */
    var action by SurveyFollowUpTable.action

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(SurveyFollowUpTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
