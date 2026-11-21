package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * When an organization was created, from organization's created event: nurturing's own copy for
 * the per-email usage's day offsets. Its own table, so an early creation never adds an unseeded
 * NurturingOrganization row.
 */
class NurturingOrganizationCreatedEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, NurturingOrganizationCreatedEntity>(NurturingOrganizationCreatedTable)

    var createdAt by NurturingOrganizationCreatedTable.createdAt
}
