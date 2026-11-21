package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * Nurturing's copy of each personal project's owner, from organization's and project's facts.
 * The first owner recorded stays: a personal project's owner never changes.
 */
class NurturingPersonalProjectEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, NurturingPersonalProjectEntity>(NurturingPersonalProjectTable)

    var organizationId by NurturingPersonalProjectTable.organizationId
    var ownerUserId by NurturingPersonalProjectTable.ownerUserId
}
