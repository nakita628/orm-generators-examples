package triggerdev.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

class OAuthRefreshTokenEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, OAuthRefreshTokenEntity>(OAuthRefreshTokenTable)

    var hashedToken by OAuthRefreshTokenTable.hashedToken
    var personalAccessToken by PersonalAccessTokenEntity referencedOn OAuthRefreshTokenTable.personalAccessTokenId
    var personalAccessTokenId by OAuthRefreshTokenTable.personalAccessTokenId
    var oauthClient by OAuthClientEntity referencedOn OAuthRefreshTokenTable.oauthClientId
    var oauthClientId by OAuthRefreshTokenTable.oauthClientId
    var scope by OAuthRefreshTokenTable.scope
    var resource by OAuthRefreshTokenTable.resource
    var usedAt by OAuthRefreshTokenTable.usedAt
    var replacedById by OAuthRefreshTokenTable.replacedById
    var idleExpiresAt by OAuthRefreshTokenTable.idleExpiresAt
    var createdAt by OAuthRefreshTokenTable.createdAt
}
