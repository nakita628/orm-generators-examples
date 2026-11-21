package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Defines targeting rules for surveys based on contact attributes.
 * Used to show surveys only to contacts matching specific criteria.
 *
 * @property id - Unique identifier for the filter
 * @property attributeKey - The contact attribute to filter on
 * @property survey - The survey being filtered
 * @property condition - The comparison operator to use
 * @property value - The value to compare against
 */
class SurveyAttributeFilterEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, SurveyAttributeFilterEntity>(SurveyAttributeFilterTable)

    var createdAt by SurveyAttributeFilterTable.createdAt
    var updatedAt by SurveyAttributeFilterTable.updatedAt
    var attributeKey by ContactAttributeKeyEntity referencedOn SurveyAttributeFilterTable.attributeKeyId
    var attributeKeyId by SurveyAttributeFilterTable.attributeKeyId
    var survey by SurveyEntity referencedOn SurveyAttributeFilterTable.surveyId
    var surveyId by SurveyAttributeFilterTable.surveyId
    var condition by SurveyAttributeFilterTable.condition
    var value by SurveyAttributeFilterTable.value

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(SurveyAttributeFilterTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
