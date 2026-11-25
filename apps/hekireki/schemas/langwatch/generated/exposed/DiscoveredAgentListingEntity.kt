package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * One source's presence fact for one discovered agent. An agent is gone only when every
 * source that listed it marked it gone; an agent no source listed (older sightings) never is.
 */
class DiscoveredAgentListingEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, DiscoveredAgentListingEntity>(DiscoveredAgentListingTable)

    var organizationId by DiscoveredAgentListingTable.organizationId
    var ingestionSourceId by DiscoveredAgentListingTable.ingestionSourceId
    var discoveredAgentId by DiscoveredAgentListingTable.discoveredAgentId
    var lastListedAt by DiscoveredAgentListingTable.lastListedAt

    /** Set when this source's listing stopped naming the agent; cleared when it names it again. */
    var goneAt by DiscoveredAgentListingTable.goneAt
    var createdAt by DiscoveredAgentListingTable.createdAt
    var updatedAt by DiscoveredAgentListingTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(DiscoveredAgentListingTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
