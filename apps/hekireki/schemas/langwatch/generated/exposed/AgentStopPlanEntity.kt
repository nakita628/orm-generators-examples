package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeID
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.CompositeEntity
import org.jetbrains.exposed.v1.dao.CompositeEntityClass
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * One agent's stop plan version: bumped by every save, so a save over a plan someone else saved
 * since it was read is refused. No row means the plan was never saved (version 0).
 */
class AgentStopPlanEntity(
    id: EntityID<CompositeID>,
) : CompositeEntity(id) {
    companion object : CompositeEntityClass<AgentStopPlanEntity>(AgentStopPlanTable)

    val organizationId by AgentStopPlanTable.organizationId

    /** The AgentRegistryEntry the plan belongs to. */
    val agentId by AgentStopPlanTable.agentId
    var version by AgentStopPlanTable.version
    var updatedAt by AgentStopPlanTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AgentStopPlanTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
