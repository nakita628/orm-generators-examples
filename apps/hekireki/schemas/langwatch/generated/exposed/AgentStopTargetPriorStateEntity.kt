package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeID
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.CompositeEntity
import org.jetbrains.exposed.v1.dao.CompositeEntityClass

/**
 * Whether a target was already stopped before a stop first touched it, recorded once before the
 * stop's platform call, so a retried call after a success still knows the change is the stop's own.
 */
class AgentStopTargetPriorStateEntity(
    id: EntityID<CompositeID>,
) : CompositeEntity(id) {
    companion object : CompositeEntityClass<AgentStopTargetPriorStateEntity>(AgentStopTargetPriorStateTable)

    val organizationId by AgentStopTargetPriorStateTable.organizationId
    val stopId by AgentStopTargetPriorStateTable.stopId
    val targetId by AgentStopTargetPriorStateTable.targetId
    var wasStopped by AgentStopTargetPriorStateTable.wasStopped
    var recordedAt by AgentStopTargetPriorStateTable.recordedAt
}
