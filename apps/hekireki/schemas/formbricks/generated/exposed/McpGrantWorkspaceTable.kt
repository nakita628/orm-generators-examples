package formbricks.models

import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.CompositeIdTable

/** A workspace a grant with `workspaces` access reaches. */
object McpGrantWorkspaceTable : CompositeIdTable("\"McpGrantWorkspace\"") {
    val grantId =
        reference(
            "grantId",
            McpGrantTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"McpGrantWorkspace_grantId_fkey\"",
        )
    val workspaceId =
        reference(
            "workspaceId",
            WorkspaceTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"McpGrantWorkspace_workspaceId_fkey\"",
        )
    override val primaryKey = PrimaryKey(grantId, workspaceId, name = "\"McpGrantWorkspace_pkey\"")

    init {
        addIdColumn(grantId)
        addIdColumn(workspaceId)
        index("\"McpGrantWorkspace_workspaceId_idx\"", false, workspaceId)
    }
}
