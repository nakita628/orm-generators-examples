package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable

/** Nurturing's copy of each project that received a real trace, from trace's and project's facts. */
object NurturingTracedProjectTable : IdTable<String>("\"NurturingTracedProject\"") {
    override val id = text("projectId").entityId()
    val organizationId = text("organizationId")
    override val primaryKey = PrimaryKey(id, name = "\"NurturingTracedProject_pkey\"")

    init {
        index("\"NurturingTracedProject_organizationId_idx\"", false, organizationId)
    }
}
