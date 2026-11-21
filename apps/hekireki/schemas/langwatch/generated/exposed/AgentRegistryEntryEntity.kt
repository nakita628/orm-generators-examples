package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * One agent in the organisation's governance registry, from any source governance reads
 * (enterprise/modules/governance/specs/agent-registry.feature). `entry` is the latest
 * machine-readable snapshot; `version` grows by one on every material change.
 */
class AgentRegistryEntryEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AgentRegistryEntryEntity>(AgentRegistryEntryTable)

    var organizationId by AgentRegistryEntryTable.organizationId

    /** Where the agent came from and its id there, e.g. `langwatch:<agentId>`. */
    var sourceKey by AgentRegistryEntryTable.sourceKey
    var source by AgentRegistryEntryTable.source1
    var status by AgentRegistryEntryTable.status
    var version by AgentRegistryEntryTable.version
    var fingerprint by AgentRegistryEntryTable.fingerprint
    var entry by AgentRegistryEntryTable.entry
    var firstSeenAt by AgentRegistryEntryTable.firstSeenAt
    var lastSeenAt by AgentRegistryEntryTable.lastSeenAt
    var removedAt by AgentRegistryEntryTable.removedAt
    var createdAt by AgentRegistryEntryTable.createdAt
    var updatedAt by AgentRegistryEntryTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AgentRegistryEntryTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
