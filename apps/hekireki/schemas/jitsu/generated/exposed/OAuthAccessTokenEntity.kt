package jitsu.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

class OAuthAccessTokenEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, OAuthAccessTokenEntity>(OAuthAccessTokenTable)

    /** @zod.custom(z.coerce.date()) */
    var createdAt by OAuthAccessTokenTable.createdAt

    /** @zod.custom(z.coerce.date()) */
    var lastUsed by OAuthAccessTokenTable.lastUsed
    var hash by OAuthAccessTokenTable.hash

    /** @zod.custom(z.coerce.date()) */
    var expiresAt by OAuthAccessTokenTable.expiresAt
    var refreshTokenId by OAuthAccessTokenTable.refreshTokenId
    var refreshToken by UserApiTokenEntity referencedOn OAuthAccessTokenTable.refreshTokenId
}
