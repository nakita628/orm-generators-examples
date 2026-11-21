package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * Nurturing's read model: which organization a project belongs to, from project's created event.
 * The per-email usage places a project and lists an organization's projects from these rows only.
 */
class NurturingProjectEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, NurturingProjectEntity>(NurturingProjectTable)

    var organizationId by NurturingProjectTable.organizationId
}
