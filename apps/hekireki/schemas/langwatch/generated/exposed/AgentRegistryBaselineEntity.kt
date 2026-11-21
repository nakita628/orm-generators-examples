package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * An organisation's registry baseline: its first reconcile, written as already pushed so the
 * agents it already had are not announced. Every reconcile after it pushes its changes.
 */
class AgentRegistryBaselineEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AgentRegistryBaselineEntity>(AgentRegistryBaselineTable)

    var baselinedAt by AgentRegistryBaselineTable.baselinedAt
}
