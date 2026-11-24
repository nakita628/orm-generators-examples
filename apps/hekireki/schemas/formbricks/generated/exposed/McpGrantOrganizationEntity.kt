package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeID
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.CompositeEntity
import org.jetbrains.exposed.v1.dao.CompositeEntityClass

/** An organization a grant with `organizations` access reaches, workspaces created there later included. */
class McpGrantOrganizationEntity(
    id: EntityID<CompositeID>,
) : CompositeEntity(id) {
    companion object : CompositeEntityClass<McpGrantOrganizationEntity>(McpGrantOrganizationTable)

    val grantId by McpGrantOrganizationTable.grantId
    val grant by McpGrantEntity referencedOn McpGrantOrganizationTable.grantId
    val organizationId by McpGrantOrganizationTable.organizationId
    val organization by OrganizationEntity referencedOn McpGrantOrganizationTable.organizationId
}
