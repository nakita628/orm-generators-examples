package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * The `agent_stop` aggregate's Postgres projection: one row per kill switch stop, its status and
 * every target's outcome, folded from the stop's events.
 */
class AgentStopEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AgentStopEntity>(AgentStopTable)

    var organizationId by AgentStopTable.organizationId
    var agentId by AgentStopTable.agentId
    var agentName by AgentStopTable.agentName
    var status by AgentStopTable.status
    var severity by AgentStopTable.severity
    var approval by AgentStopTable.approval
    var trigger by AgentStopTable.trigger
    var reason by AgentStopTable.reason
    var requestedByUserId by AgentStopTable.requestedByUserId
    var requestedByApiKeyId by AgentStopTable.requestedByApiKeyId
    var decidedByUserId by AgentStopTable.decidedByUserId
    var decidedByApiKeyId by AgentStopTable.decidedByApiKeyId
    var restoredByUserId by AgentStopTable.restoredByUserId
    var restoredByApiKeyId by AgentStopTable.restoredByApiKeyId
    var targets by AgentStopTable.targets
    var requestedAt by AgentStopTable.requestedAt
    var decidedAt by AgentStopTable.decidedAt
    var quarantinedAt by AgentStopTable.quarantinedAt
    var restoreRequestedAt by AgentStopTable.restoreRequestedAt
    var restoredAt by AgentStopTable.restoredAt
    var lastAppliedEventId by AgentStopTable.lastAppliedEventId
    var lastAcceptedAt by AgentStopTable.lastAcceptedAt
    var lastOccurredAt by AgentStopTable.lastOccurredAt
    var projectionVersion by AgentStopTable.projectionVersion
    var createdAt by AgentStopTable.createdAt
    var updatedAt by AgentStopTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AgentStopTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
