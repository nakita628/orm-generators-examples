package teable.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeIdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object DomainEventInboxTable : CompositeIdTable("domain_event_inbox") {
    val consumerId = text("consumer_id").entityId()
    val eventId = text("event_id").entityId()
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(consumerId, eventId, name = "domain_event_inbox_pkey")

    init {
        index("domain_event_inbox_event_id_idx", false, eventId)
    }
}
