package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * A successful callback through one SSO connection. Identifier.lastUsedAt
 * also advances during routing and resolution, so it cannot prove that a
 * member actually authenticated through the replacement before cutover.
 */
class SsoAuthenticationActivityEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, SsoAuthenticationActivityEntity>(SsoAuthenticationActivityTable)

    var organizationId by SsoAuthenticationActivityTable.organizationId
    var connectionId by SsoAuthenticationActivityTable.connectionId
    var userId by SsoAuthenticationActivityTable.userId
    var authenticatedAt by SsoAuthenticationActivityTable.authenticatedAt

    /** The provider-side account the assertion named; null before #7633's arrival door recorded it. */
    var providerAccountId by SsoAuthenticationActivityTable.providerAccountId
}
