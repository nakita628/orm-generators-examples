package teable.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

object DomainEventOutboxTable : IdTable<String>("domain_event_outbox") {
    override val id = text("id").entityId()
    val baseId = text("base_id")
    val tableId = text("table_id").nullable()
    val messageName = text("message_name")
    val schemaVersion = integer("schema_version")
    val aggregateId = text("aggregate_id").nullable()
    val payload = jsonb<String>("payload", { it }, { it })
    val payloadBytes = integer("payload_bytes")
    val catalogGeneration = integer("catalog_generation")
    val requiredConsumers = jsonb<String>("required_consumers", { it }, { it })
    val bindingId = text("binding_id").nullable()
    val storageEpoch = integer("storage_epoch").nullable()
    val unpublished = bool("unpublished").default(true)
    val settled = text("settled").nullable()
    val settledAt = pgTimestamp("settled_at", 3).nullable()
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "domain_event_outbox_pkey")

    init {
        index("domain_event_outbox_base_id_idx", false, baseId)
    }
}
