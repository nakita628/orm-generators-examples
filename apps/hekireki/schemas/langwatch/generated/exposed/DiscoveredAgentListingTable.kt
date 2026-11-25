package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * One source's presence fact for one discovered agent. An agent is gone only when every
 * source that listed it marked it gone; an agent no source listed (older sightings) never is.
 */
object DiscoveredAgentListingTable : IdTable<String>("\"DiscoveredAgentListing\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val organizationId = text("organizationId")
    val ingestionSourceId = text("ingestionSourceId")
    val discoveredAgentId = text("discoveredAgentId")
    val lastListedAt = pgTimestamp("lastListedAt", 3)

    /** Set when this source's listing stopped naming the agent; cleared when it names it again. */
    val goneAt = pgTimestamp("goneAt", 3).nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"DiscoveredAgentListing_pkey\"")

    init {
        index("\"DiscoveredAgentListing_organizationId_discoveredAgentId_idx\"", false, organizationId, discoveredAgentId)
        uniqueIndex("DiscoveredAgentListing_ingestionSourceId_discoveredAgentId_key", ingestionSourceId, discoveredAgentId)
    }
}
