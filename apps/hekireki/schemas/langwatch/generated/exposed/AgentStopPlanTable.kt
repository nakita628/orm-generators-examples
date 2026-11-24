package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeIdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * One agent's stop plan version: bumped by every save, so a save over a plan someone else saved
 * since it was read is refused. No row means the plan was never saved (version 0).
 */
object AgentStopPlanTable : CompositeIdTable("\"AgentStopPlan\"") {
    val organizationId = text("organizationId").entityId()

    /** The AgentRegistryEntry the plan belongs to. */
    val agentId = text("agentId").entityId()
    val version = integer("version")
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(organizationId, agentId, name = "\"AgentStopPlan_pkey\"")
}
