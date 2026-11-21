package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * One agent in the organisation's governance registry, from any source governance reads
 * (enterprise/modules/governance/specs/agent-registry.feature). `entry` is the latest
 * machine-readable snapshot; `version` grows by one on every material change.
 */
object AgentRegistryEntryTable : IdTable<String>("\"AgentRegistryEntry\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val organizationId = text("organizationId")

    /** Where the agent came from and its id there, e.g. `langwatch:<agentId>`. */
    val sourceKey = text("sourceKey")
    val source1 = text("source")
    val status = text("status")
    val version = integer("version")
    val fingerprint = text("fingerprint")
    val entry = jsonb<String>("entry", { it }, { it })
    val firstSeenAt = pgTimestamp("firstSeenAt", 3)
    val lastSeenAt = pgTimestamp("lastSeenAt", 3).nullable()
    val removedAt = pgTimestamp("removedAt", 3).nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"AgentRegistryEntry_pkey\"")

    init {
        index("\"AgentRegistryEntry_organizationId_updatedAt_idx\"", false, organizationId, updatedAt)
        uniqueIndex("\"AgentRegistryEntry_organizationId_sourceKey_key\"", organizationId, sourceKey)
    }
}
