package triggerdev.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/** Used to generate PersonalAccessTokens, they're one-time use */
class AuthorizationCodeEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AuthorizationCodeEntity>(AuthorizationCodeTable)

    var code by AuthorizationCodeTable.code
    var personalAccessToken by PersonalAccessTokenEntity optionalReferencedOn AuthorizationCodeTable.personalAccessTokenId
    var personalAccessTokenId by AuthorizationCodeTable.personalAccessTokenId

    /** Client IP of the request that minted the code, compared with the approving browser's IP. Deleted with the code once it expires. */
    var requesterIp by AuthorizationCodeTable.requesterIp

    /** S256 PKCE challenge from the client that minted the code. Null for codes minted by CLIs that predate it. */
    var codeChallenge by AuthorizationCodeTable.codeChallenge

    /** Loopback port the approving browser is redirected to with the approval code. Null means the user pastes it. */
    var callbackPort by AuthorizationCodeTable.callbackPort

    /** sha256 of the approval code, which only the approving browser ever sees. */
    var approvalCodeHash by AuthorizationCodeTable.approvalCodeHash
    var createdAt by AuthorizationCodeTable.createdAt
    var updatedAt by AuthorizationCodeTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AuthorizationCodeTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
