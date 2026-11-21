package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable

/**
 * When an organization was created, from organization's created event: nurturing's own copy for
 * the per-email usage's day offsets. Its own table, so an early creation never adds an unseeded
 * NurturingOrganization row.
 */
object NurturingOrganizationCreatedTable : IdTable<String>("\"NurturingOrganizationCreated\"") {
    override val id = text("organizationId").entityId()
    val createdAt = pgTimestamp("createdAt", 3)
    override val primaryKey = PrimaryKey(id, name = "\"NurturingOrganizationCreated_pkey\"")

    init {
        index("\"NurturingOrganizationCreated_createdAt_idx\"", false, createdAt)
    }
}
