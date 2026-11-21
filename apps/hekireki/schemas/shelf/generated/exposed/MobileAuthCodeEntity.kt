package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * Native-app SSO web-delegated login — single-use session-handoff codes.
 * @see apps/webapp/app/routes/_auth+/oauth.callback.mobile.tsx
 */
class MobileAuthCodeEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, MobileAuthCodeEntity>(MobileAuthCodeTable)

    var user by UserEntity referencedOn MobileAuthCodeTable.userId
    var userId by MobileAuthCodeTable.userId
    var codeHash by MobileAuthCodeTable.codeHash
    var codeChallenge by MobileAuthCodeTable.codeChallenge
    var sessionCiphertext by MobileAuthCodeTable.sessionCiphertext
    var consumedAt by MobileAuthCodeTable.consumedAt
    var expiresAt by MobileAuthCodeTable.expiresAt
    var createdAt by MobileAuthCodeTable.createdAt
}
