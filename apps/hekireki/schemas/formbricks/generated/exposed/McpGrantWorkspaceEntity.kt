package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeID
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.CompositeEntity
import org.jetbrains.exposed.v1.dao.CompositeEntityClass

/** A workspace a grant with `workspaces` access reaches. */
class McpGrantWorkspaceEntity(
    id: EntityID<CompositeID>,
) : CompositeEntity(id) {
    companion object : CompositeEntityClass<McpGrantWorkspaceEntity>(McpGrantWorkspaceTable)

    val grantId by McpGrantWorkspaceTable.grantId
    val grant by McpGrantEntity referencedOn McpGrantWorkspaceTable.grantId
    val workspaceId by McpGrantWorkspaceTable.workspaceId
    val workspace by WorkspaceEntity referencedOn McpGrantWorkspaceTable.workspaceId
}
