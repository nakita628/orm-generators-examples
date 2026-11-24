package teable.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object DomainEventDeliveryTable : IdTable<String>("domain_event_delivery") {
    override val id = text("id").entityId()
    val eventId = text("event_id")
    val consumerId = text("consumer_id")
    val status = text("status")
    val attempts = integer("attempts").default(0)
    val maxAttempts = integer("max_attempts").default(12)
    val leaseToken = text("lease_token").nullable()
    val leaseExpiresAt = pgTimestamptz("lease_expires_at", 3).nullable()
    val nextAttemptAt =
        pgTimestamptz("next_attempt_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val lastError = text("last_error").nullable()
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "domain_event_delivery_pkey")

    init {
        uniqueIndex("domain_event_delivery_event_id_consumer_id_key", eventId, consumerId)
    }
}
