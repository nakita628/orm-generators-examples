package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable

/**
 * Nurturing's read model: which organization a project belongs to, from project's created event.
 * The per-email usage places a project and lists an organization's projects from these rows only.
 */
object NurturingProjectTable : IdTable<String>("\"NurturingProject\"") {
    override val id = text("projectId").entityId()
    val organizationId = text("organizationId")
    override val primaryKey = PrimaryKey(id, name = "\"NurturingProject_pkey\"")

    init {
        index("\"NurturingProject_organizationId_idx\"", false, organizationId)
    }
}
