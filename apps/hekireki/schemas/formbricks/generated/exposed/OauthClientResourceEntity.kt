package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * Better Auth 1.7 (ENG-2343): links a client to the resources it may request tokens for. With
 * `enforcePerClientResources` (the default) a missing row means `invalid_target` at the token
 * endpoint, which is why the migration backfills one row per existing client.
 * Both sides reference natural keys, matching the other oauth child tables.
 */
class OauthClientResourceEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, OauthClientResourceEntity>(OauthClientResourceTable)

    var clientId by OauthClientResourceTable.clientId
    var client by OauthClientEntity referencedOn OauthClientResourceTable.clientId
    var resourceId by OauthClientResourceTable.resourceId
    var resource by OauthResourceEntity referencedOn OauthClientResourceTable.resourceId
    var metadata by OauthClientResourceTable.metadata
    var createdAt by OauthClientResourceTable.createdAt
}
