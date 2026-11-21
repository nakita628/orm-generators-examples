package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable

/**
 * An organisation's registry baseline: its first reconcile, written as already pushed so the
 * agents it already had are not announced. Every reconcile after it pushes its changes.
 */
object AgentRegistryBaselineTable : IdTable<String>("\"AgentRegistryBaseline\"") {
    override val id = text("organizationId").entityId()
    val baselinedAt = pgTimestamp("baselinedAt", 3)
    override val primaryKey = PrimaryKey(id, name = "\"AgentRegistryBaseline_pkey\"")
}
