package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * Maps survey elements to Hub FeedbackRecords for Formbricks feedback sources.
 * Each row represents one element that will create FeedbackRecords when answered.
 *
 * @property id - Unique identifier for the mapping
 * @property feedbackSource - The parent feedback source
 * @property survey - The survey containing the element
 * @property elementId - The element ID within the survey (from blocks[].elements[].id)
 * @property hubFieldType - The field_type to use in Hub (text, nps, rating, etc.)
 * @property customFieldLabel - Optional override for the element headline as field_label in Hub
 */
class FeedbackSourceFormbricksMappingEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, FeedbackSourceFormbricksMappingEntity>(FeedbackSourceFormbricksMappingTable)

    var createdAt by FeedbackSourceFormbricksMappingTable.createdAt
    var feedbackSourceId by FeedbackSourceFormbricksMappingTable.feedbackSourceId
    var workspaceId by FeedbackSourceFormbricksMappingTable.workspaceId
    var surveyId by FeedbackSourceFormbricksMappingTable.surveyId
    var elementId by FeedbackSourceFormbricksMappingTable.elementId
    var hubFieldType by FeedbackSourceFormbricksMappingTable.hubFieldType
    var customFieldLabel by FeedbackSourceFormbricksMappingTable.customFieldLabel
}
