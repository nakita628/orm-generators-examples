package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * The same shape as DiscoveredPerson for non-human actors - Genie spaces,
 * Copilot bots, OpenAI projects - plus whatever provider-native metadata the
 * source published. Discovering an agent grants it nothing either.
 */
class DiscoveredAgentEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, DiscoveredAgentEntity>(DiscoveredAgentTable)

    var organizationId by DiscoveredAgentTable.organizationId
    var provider by DiscoveredAgentTable.provider

    /** Genie space / Copilot bot / OpenAI project id, verbatim. */
    var rawAgentId by DiscoveredAgentTable.rawAgentId
    var displayText by DiscoveredAgentTable.displayText

    /**
     * Workspace, plan, and kin - provider-native descriptive fields only. Never
     * credentials: this row is read by admin screens, not by any pull path.
     */
    var metadata by DiscoveredAgentTable.metadata
    var firstSeenAt by DiscoveredAgentTable.firstSeenAt
    var lastSeenAt by DiscoveredAgentTable.lastSeenAt
    var createdAt by DiscoveredAgentTable.createdAt
    var updatedAt by DiscoveredAgentTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(DiscoveredAgentTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
