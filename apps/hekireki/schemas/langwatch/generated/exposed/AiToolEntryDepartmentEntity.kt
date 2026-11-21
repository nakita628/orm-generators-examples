package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * Per-department scope binding for AiToolEntry. Empty set = org-wide.
 * Non-empty = entry visible only to members of those departments (an
 * OrganizationUser whose `departmentId` is in the set). Mirrors
 * AiToolEntryTeam's shape. `departmentId` is a plain column ref
 * (relationMode="prisma"), matching how Department is referenced
 * elsewhere; the service validates the id belongs to the org.
 */
class AiToolEntryDepartmentEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AiToolEntryDepartmentEntity>(AiToolEntryDepartmentTable)

    var entryId by AiToolEntryDepartmentTable.entryId
    var departmentId by AiToolEntryDepartmentTable.departmentId
}
