package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb

/**
 * The registry's append-only change feed: one row per registered, versioned, removed or
 * tool-deprecated fact. `sequence` is the pull cursor; `deliveryRecordedAt` marks the webhook
 * fan-out as recorded in governance's outbox.
 */
object AgentRegistryEventTable : IdTable<String>("\"AgentRegistryEvent\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val organizationId = text("organizationId")
    val sequence = long("sequence").autoIncrement()
    val type = text("type")
    val entryId = text("entryId")
    val fromVersion = integer("fromVersion").nullable()
    val toVersion = integer("toVersion").nullable()
    val payload = jsonb<String>("payload", { it }, { it })
    val occurredAt = pgTimestamp("occurredAt", 3)
    val deliveryRecordedAt = pgTimestamp("deliveryRecordedAt", 3).nullable()
    override val primaryKey = PrimaryKey(id, name = "\"AgentRegistryEvent_pkey\"")

    init {
        index("\"AgentRegistryEvent_organizationId_deliveryRecordedAt_idx\"", false, organizationId, deliveryRecordedAt)
        index("\"AgentRegistryEvent_organizationId_entryId_idx\"", false, organizationId, entryId)
        uniqueIndex("\"AgentRegistryEvent_organizationId_sequence_key\"", organizationId, sequence)
    }
}
