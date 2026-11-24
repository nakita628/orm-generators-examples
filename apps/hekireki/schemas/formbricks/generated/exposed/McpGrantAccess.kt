package formbricks.models

enum class McpGrantAccess(
    val dbName: String,
) {
    ALL("all"),
    ORGANIZATIONS("organizations"),
    WORKSPACES("workspaces"),
}
