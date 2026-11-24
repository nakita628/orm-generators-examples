package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * One target of a registry agent's kill switch stop plan (enterprise/modules/governance/specs/
 * agent-kill-switch.feature): a Databricks serving endpoint, a Bedrock agent alias or a gateway
 * virtual key. A plan write replaces every row of the agent at once.
 */
class AgentStopPlanTargetEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AgentStopPlanTargetEntity>(AgentStopPlanTargetTable)

    var organizationId by AgentStopPlanTargetTable.organizationId

    /** The AgentRegistryEntry the plan belongs to. */
    var agentId by AgentStopPlanTargetTable.agentId
    var position by AgentStopPlanTargetTable.position
    var label by AgentStopPlanTargetTable.label
    var target by AgentStopPlanTargetTable.target
    var createdAt by AgentStopPlanTargetTable.createdAt
    var updatedAt by AgentStopPlanTargetTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AgentStopPlanTargetTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
