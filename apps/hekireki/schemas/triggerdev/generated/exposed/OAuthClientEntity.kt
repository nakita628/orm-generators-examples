package triggerdev.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class OAuthClientEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, OAuthClientEntity>(OAuthClientTable)

    var kind by OAuthClientTable.kind
    var clientId by OAuthClientTable.clientId
    var metadataUrl by OAuthClientTable.metadataUrl
    var displayName by OAuthClientTable.displayName
    var selfReportedName by OAuthClientTable.selfReportedName
    var redirectUris by OAuthClientTable.redirectUris
    var dedupeKey by OAuthClientTable.dedupeKey
    var metadataExpiresAt by OAuthClientTable.metadataExpiresAt
    var lastUsedAt by OAuthClientTable.lastUsedAt
    var createdAt by OAuthClientTable.createdAt
    var updatedAt by OAuthClientTable.updatedAt
    val authorizationCodes by AuthorizationCodeEntity optionalReferrersOn AuthorizationCodeTable.oauthClientId
    val personalAccessTokens by PersonalAccessTokenEntity optionalReferrersOn PersonalAccessTokenTable.oauthClientId
    val refreshTokens by OAuthRefreshTokenEntity referrersOn OAuthRefreshTokenTable.oauthClientId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(OAuthClientTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
