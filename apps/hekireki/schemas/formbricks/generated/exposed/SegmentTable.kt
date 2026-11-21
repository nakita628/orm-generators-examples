package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Defines a segment of contacts based on attributes.
 * Used for targeting surveys to specific user groups.
 *
 * @property id - Unique identifier for the segment
 * @property title - Display name of the segment
 * @property filters - Rules defining the segment
 * @property isPrivate - Whether the segment is private
 * @property workspace - The workspace this segment belongs to
 */
object SegmentTable : IdTable<String>("\"Segment\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updated_at", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val title = text("title")
    val description = text("description").nullable()
    val isPrivate = bool("isPrivate").default(true)

    /** [SegmentFilter] */
    val filters = jsonb<String>("filters", { it }, { it }).default("[]")
    val workspaceId =
        reference(
            "workspaceId",
            WorkspaceTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Segment_workspaceId_fkey\"",
        )
    override val primaryKey = PrimaryKey(id, name = "\"Segment_pkey\"")

    init {
        index("\"Segment_workspaceId_idx\"", false, workspaceId)
        uniqueIndex("\"Segment_workspaceId_title_key\"", workspaceId, title)
    }
}
