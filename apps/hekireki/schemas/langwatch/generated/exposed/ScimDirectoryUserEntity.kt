package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeID
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.CompositeEntity
import org.jetbrains.exposed.v1.dao.CompositeEntityClass

/**
 * A connection's managed users, including resources without an externalId.
 * Scalar keys only; readers exclude users that no longer exist.
 */
class ScimDirectoryUserEntity(
    id: EntityID<CompositeID>,
) : CompositeEntity(id) {
    companion object : CompositeEntityClass<ScimDirectoryUserEntity>(ScimDirectoryUserTable)

    var organizationId by ScimDirectoryUserTable.organizationId
    val connectionId by ScimDirectoryUserTable.connectionId
    val userId by ScimDirectoryUserTable.userId
}
