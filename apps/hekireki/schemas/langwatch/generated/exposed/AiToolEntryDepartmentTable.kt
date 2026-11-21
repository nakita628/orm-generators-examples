package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable

/**
 * Per-department scope binding for AiToolEntry. Empty set = org-wide.
 * Non-empty = entry visible only to members of those departments (an
 * OrganizationUser whose `departmentId` is in the set). Mirrors
 * AiToolEntryTeam's shape. `departmentId` is a plain column ref
 * (relationMode="prisma"), matching how Department is referenced
 * elsewhere; the service validates the id belongs to the org.
 */
object AiToolEntryDepartmentTable : IdTable<String>("\"AiToolEntryDepartment\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val entryId = text("entryId")
    val departmentId = text("departmentId")
    override val primaryKey = PrimaryKey(id, name = "\"AiToolEntryDepartment_pkey\"")

    init {
        index("\"AiToolEntryDepartment_departmentId_idx\"", false, departmentId)
        uniqueIndex("\"AiToolEntryDepartment_entryId_departmentId_key\"", entryId, departmentId)
    }
}
