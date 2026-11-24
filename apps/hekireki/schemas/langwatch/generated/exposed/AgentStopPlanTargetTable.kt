package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * One target of a registry agent's kill switch stop plan (enterprise/modules/governance/specs/
 * agent-kill-switch.feature): a Databricks serving endpoint, a Bedrock agent alias or a gateway
 * virtual key. A plan write replaces every row of the agent at once.
 */
object AgentStopPlanTargetTable : IdTable<String>("\"AgentStopPlanTarget\"") {
    override val id = text("id").entityId()
    val organizationId = text("organizationId")

    /** The AgentRegistryEntry the plan belongs to. */
    val agentId = text("agentId")
    val position = integer("position")
    val label = text("label")
    val target = jsonb<String>("target", { it }, { it })
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"AgentStopPlanTarget_pkey\"")

    init {
        index("\"AgentStopPlanTarget_organizationId_agentId_idx\"", false, organizationId, agentId)
    }
}
