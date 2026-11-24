package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeIdTable

/**
 * Whether a target was already stopped before a stop first touched it, recorded once before the
 * stop's platform call, so a retried call after a success still knows the change is the stop's own.
 */
object AgentStopTargetPriorStateTable : CompositeIdTable("\"AgentStopTargetPriorState\"") {
    val organizationId = text("organizationId").entityId()
    val stopId = text("stopId").entityId()
    val targetId = text("targetId").entityId()
    val wasStopped = bool("wasStopped")
    val recordedAt = pgTimestamp("recordedAt", 3)
    override val primaryKey = PrimaryKey(organizationId, stopId, targetId, name = "\"AgentStopTargetPriorState_pkey\"")
}
