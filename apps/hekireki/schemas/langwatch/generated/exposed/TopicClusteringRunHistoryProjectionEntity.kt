package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * Per-project topic clustering run history (audit read model, ADR-051).
 * One row per project; `runs` is the bounded newest-first entry list.
 * Rebuildable by replaying the event log.
 */
class TopicClusteringRunHistoryProjectionEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, TopicClusteringRunHistoryProjectionEntity>(TopicClusteringRunHistoryProjectionTable)

    var projectId by TopicClusteringRunHistoryProjectionTable.projectId
    var runs by TopicClusteringRunHistoryProjectionTable.runs
    var createdAt by TopicClusteringRunHistoryProjectionTable.createdAt
    var updatedAt by TopicClusteringRunHistoryProjectionTable.updatedAt

    /** Business time of the latest applied event. */
    var occurredAt by TopicClusteringRunHistoryProjectionTable.occurredAt

    /** Event-log acceptance time. Together with lastEventId this is the cursor. */
    var acceptedAt by TopicClusteringRunHistoryProjectionTable.acceptedAt
    var lastEventId by TopicClusteringRunHistoryProjectionTable.lastEventId
    var projectionVersion by TopicClusteringRunHistoryProjectionTable.projectionVersion
}
