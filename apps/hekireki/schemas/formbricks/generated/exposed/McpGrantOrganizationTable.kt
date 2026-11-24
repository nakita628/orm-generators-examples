package formbricks.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.CompositeIdTable

/** An organization a grant with `organizations` access reaches, workspaces created there later included. */
object McpGrantOrganizationTable : CompositeIdTable("\"McpGrantOrganization\"") {
    val grantId =
        reference(
            "grantId",
            McpGrantTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"McpGrantOrganization_grantId_fkey\"",
        )
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"McpGrantOrganization_organizationId_fkey\"",
        )
    override val primaryKey = PrimaryKey(grantId, organizationId, name = "\"McpGrantOrganization_pkey\"")

    init {
        addIdColumn(grantId)
        addIdColumn(organizationId)
        index("\"McpGrantOrganization_organizationId_idx\"", false, organizationId)
    }
}
