package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * Cursor row for the topic-model projection (ADR-051). The model itself is
 * the `Topic` table, written through by the projection; this row only
 * records how far the fold has applied the event log.
 */
class TopicModelProjectionEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, TopicModelProjectionEntity>(TopicModelProjectionTable)

    var projectId by TopicModelProjectionTable.projectId
    var createdAt by TopicModelProjectionTable.createdAt
    var updatedAt by TopicModelProjectionTable.updatedAt

    /** Business time of the latest applied event. */
    var occurredAt by TopicModelProjectionTable.occurredAt

    /** Event-log acceptance time. Together with lastEventId this is the cursor. */
    var acceptedAt by TopicModelProjectionTable.acceptedAt
    var lastEventId by TopicModelProjectionTable.lastEventId
    var projectionVersion by TopicModelProjectionTable.projectionVersion
}
