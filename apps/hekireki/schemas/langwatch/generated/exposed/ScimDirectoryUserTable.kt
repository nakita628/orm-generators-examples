package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeIdTable

/**
 * A connection's managed users, including resources without an externalId.
 * Scalar keys only; readers exclude users that no longer exist.
 */
object ScimDirectoryUserTable : CompositeIdTable("\"ScimDirectoryUser\"") {
    val organizationId = text("organizationId")
    val connectionId = text("connectionId").entityId()
    val userId = text("userId").entityId()
    override val primaryKey = PrimaryKey(connectionId, userId, name = "\"ScimDirectoryUser_pkey\"")

    init {
        index("\"ScimDirectoryUser_organizationId_userId_idx\"", false, organizationId, userId)
    }
}
