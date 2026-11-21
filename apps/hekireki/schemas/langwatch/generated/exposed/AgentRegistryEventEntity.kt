package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * The registry's append-only change feed: one row per registered, versioned, removed or
 * tool-deprecated fact. `sequence` is the pull cursor; `deliveryRecordedAt` marks the webhook
 * fan-out as recorded in governance's outbox.
 */
class AgentRegistryEventEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AgentRegistryEventEntity>(AgentRegistryEventTable)

    var organizationId by AgentRegistryEventTable.organizationId
    var sequence by AgentRegistryEventTable.sequence
    var type by AgentRegistryEventTable.type
    var entryId by AgentRegistryEventTable.entryId
    var fromVersion by AgentRegistryEventTable.fromVersion
    var toVersion by AgentRegistryEventTable.toVersion
    var payload by AgentRegistryEventTable.payload
    var occurredAt by AgentRegistryEventTable.occurredAt
    var deliveryRecordedAt by AgentRegistryEventTable.deliveryRecordedAt
}
