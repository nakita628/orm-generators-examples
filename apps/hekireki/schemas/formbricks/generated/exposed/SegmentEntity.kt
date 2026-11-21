package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
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
class SegmentEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, SegmentEntity>(SegmentTable)

    var createdAt by SegmentTable.createdAt
    var updatedAt by SegmentTable.updatedAt
    var title by SegmentTable.title
    var description by SegmentTable.description
    var isPrivate by SegmentTable.isPrivate

    /** [SegmentFilter] */
    var filters by SegmentTable.filters
    var workspace by WorkspaceEntity referencedOn SegmentTable.workspaceId
    var workspaceId by SegmentTable.workspaceId
    val surveys by SurveyEntity optionalReferrersOn SurveyTable.segmentId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(SegmentTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
