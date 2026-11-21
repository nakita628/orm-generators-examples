package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/** Nurturing's copy of each project that received a real trace, from trace's and project's facts. */
class NurturingTracedProjectEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, NurturingTracedProjectEntity>(NurturingTracedProjectTable)

    var organizationId by NurturingTracedProjectTable.organizationId
}
