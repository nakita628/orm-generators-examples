package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * This installation's VAPID key pair, generated on first use so Web Push works
 * without configuration. Fixed id, so the table holds one row: a second
 * insert loses the race on the primary key and reads the winner's pair.
 */
object WebPushVapidKeyTable : IdTable<String>("\"WebPushVapidKey\"") {
    override val id = text("id").default("self").entityId()

    /** The uncompressed P-256 public key, base64url. The browser subscribes with it. */
    val publicKey = text("publicKey")

    /** The private key, base64url, encrypted with the process encryption key. */
    val privateKeyEncrypted = text("privateKeyEncrypted")
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"WebPushVapidKey_pkey\"")
}
