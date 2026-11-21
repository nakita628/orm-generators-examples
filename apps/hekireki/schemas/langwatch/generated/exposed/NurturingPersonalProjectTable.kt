package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable

/**
 * Nurturing's copy of each personal project's owner, from organization's and project's facts.
 * The first owner recorded stays: a personal project's owner never changes.
 */
object NurturingPersonalProjectTable : IdTable<String>("\"NurturingPersonalProject\"") {
    override val id = text("projectId").entityId()
    val organizationId = text("organizationId")
    val ownerUserId = text("ownerUserId")
    override val primaryKey = PrimaryKey(id, name = "\"NurturingPersonalProject_pkey\"")
}
