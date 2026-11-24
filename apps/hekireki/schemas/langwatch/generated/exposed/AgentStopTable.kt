package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * The `agent_stop` aggregate's Postgres projection: one row per kill switch stop, its status and
 * every target's outcome, folded from the stop's events.
 */
object AgentStopTable : IdTable<String>("\"AgentStop\"") {
    override val id = text("id").entityId()
    val organizationId = text("organizationId")
    val agentId = text("agentId")
    val agentName = text("agentName")
    val status = text("status")
    val severity = text("severity")
    val approval = text("approval")
    val trigger = text("trigger")
    val reason = text("reason")
    val requestedByUserId = text("requestedByUserId").nullable()
    val requestedByApiKeyId = text("requestedByApiKeyId").nullable()
    val decidedByUserId = text("decidedByUserId").nullable()
    val decidedByApiKeyId = text("decidedByApiKeyId").nullable()
    val restoredByUserId = text("restoredByUserId").nullable()
    val restoredByApiKeyId = text("restoredByApiKeyId").nullable()
    val targets = jsonb<String>("targets", { it }, { it })
    val requestedAt = pgTimestamp("requestedAt", 3)
    val decidedAt = pgTimestamp("decidedAt", 3).nullable()
    val quarantinedAt = pgTimestamp("quarantinedAt", 3).nullable()
    val restoreRequestedAt = pgTimestamp("restoreRequestedAt", 3).nullable()
    val restoredAt = pgTimestamp("restoredAt", 3).nullable()
    val lastAppliedEventId = text("lastAppliedEventId")
    val lastAcceptedAt = long("lastAcceptedAt")
    val lastOccurredAt = long("lastOccurredAt")
    val projectionVersion = text("projectionVersion")
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"AgentStop_pkey\"")

    init {
        index("\"AgentStop_organizationId_agentId_idx\"", false, organizationId, agentId)
        index("\"AgentStop_organizationId_status_idx\"", false, organizationId, status)
    }
}
