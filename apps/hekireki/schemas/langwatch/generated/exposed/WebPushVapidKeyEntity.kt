package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * This installation's VAPID key pair, generated on first use so Web Push works
 * without configuration. Fixed id, so the table holds one row: a second
 * insert loses the race on the primary key and reads the winner's pair.
 */
class WebPushVapidKeyEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, WebPushVapidKeyEntity>(WebPushVapidKeyTable)

    /** The uncompressed P-256 public key, base64url. The browser subscribes with it. */
    var publicKey by WebPushVapidKeyTable.publicKey

    /** The private key, base64url, encrypted with the process encryption key. */
    var privateKeyEncrypted by WebPushVapidKeyTable.privateKeyEncrypted
    var createdAt by WebPushVapidKeyTable.createdAt
}
