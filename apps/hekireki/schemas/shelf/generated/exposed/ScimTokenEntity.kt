package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

class ScimTokenEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, ScimTokenEntity>(ScimTokenTable)

    var tokenHash by ScimTokenTable.tokenHash
    var label by ScimTokenTable.label
    var organization by OrganizationEntity referencedOn ScimTokenTable.organizationId
    var organizationId by ScimTokenTable.organizationId
    var createdById by ScimTokenTable.createdById
    var createdAt by ScimTokenTable.createdAt
    var lastUsedAt by ScimTokenTable.lastUsedAt
}
