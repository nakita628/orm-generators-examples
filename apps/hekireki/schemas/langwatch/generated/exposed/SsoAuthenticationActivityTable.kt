package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable

/**
 * A successful callback through one SSO connection. Identifier.lastUsedAt
 * also advances during routing and resolution, so it cannot prove that a
 * member actually authenticated through the replacement before cutover.
 */
object SsoAuthenticationActivityTable : IdTable<String>("\"SsoAuthenticationActivity\"") {
    override val id = text("id").entityId()
    val organizationId = text("organizationId")
    val connectionId = text("connectionId")
    val userId = text("userId")
    val authenticatedAt = pgTimestamp("authenticatedAt", 3)

    /** The provider-side account the assertion named; null before #7633's arrival door recorded it. */
    val providerAccountId = text("providerAccountId").nullable()
    override val primaryKey = PrimaryKey(id, name = "\"SsoAuthenticationActivity_pkey\"")

    init {
        index("SsoAuthenticationActivity_organizationId_connectionId_authe_idx", false, organizationId, connectionId, authenticatedAt)
        index("\"SsoAuthenticationActivity_organizationId_userId_idx\"", false, organizationId, userId)
    }
}
