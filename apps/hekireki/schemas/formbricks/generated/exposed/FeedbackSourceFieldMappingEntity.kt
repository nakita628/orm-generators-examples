package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * Generic field mapping for non-Formbricks feedback sources (CSV today; webhook/email/slack later).
 * Maps source fields to Hub FeedbackRecord fields.
 *
 * @property id - Unique identifier for the mapping
 * @property feedbackSource - The parent feedback source
 * @property sourceFieldId - Field path for webhook (e.g., "user.id"), column name for CSV
 * @property targetFieldId - Hub field (collected_at, field_id, value_text, etc.)
 * @property staticValue - If set, use this value instead of reading from sourceFieldId
 */
class FeedbackSourceFieldMappingEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, FeedbackSourceFieldMappingEntity>(FeedbackSourceFieldMappingTable)

    var createdAt by FeedbackSourceFieldMappingTable.createdAt
    var feedbackSourceId by FeedbackSourceFieldMappingTable.feedbackSourceId
    var workspaceId by FeedbackSourceFieldMappingTable.workspaceId
    var sourceFieldId by FeedbackSourceFieldMappingTable.sourceFieldId
    var targetFieldId by FeedbackSourceFieldMappingTable.targetFieldId
    var staticValue by FeedbackSourceFieldMappingTable.staticValue
}
