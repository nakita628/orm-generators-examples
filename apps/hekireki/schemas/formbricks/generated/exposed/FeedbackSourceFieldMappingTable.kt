package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

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
object FeedbackSourceFieldMappingTable : IdTable<String>("\"FeedbackSourceFieldMapping\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val feedbackSourceId = text("feedback_source_id")
    val workspaceId = text("workspaceId")
    val sourceFieldId = text("source_field_id")
    val targetFieldId = text("target_field_id")
    val staticValue = text("static_value").nullable()
    override val primaryKey = PrimaryKey(id, name = "\"FeedbackSourceFieldMapping_pkey\"")

    init {
        uniqueIndex(
            "FeedbackSourceFieldMapping_workspaceId_feedbackSourceId_sourceF",
            workspaceId,
            feedbackSourceId,
            sourceFieldId,
            targetFieldId,
        )
        foreignKey(
            feedbackSourceId to FeedbackSourceTable.id,
            workspaceId to FeedbackSourceTable.workspaceId,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            name = "\"FeedbackSourceFieldMapping_feedbackSourceId_workspaceId_fkey\"",
        )
    }
}
