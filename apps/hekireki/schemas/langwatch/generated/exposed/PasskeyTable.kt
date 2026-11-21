package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * The passkey plugin's own table (D07). Protocol state, like TwoFactor: the
 * public key and the signature counter are the plugin's to read and write.
 * A passkey also becomes an `Identifier` row with provider `passkey`, which
 * is what makes it a sign-in method like any other - this table is the
 * credential, that one is the identity.
 *
 * `aaguid` names the authenticator MODEL, never a device or a person, and is
 * only present at registration. Privacy-preserving platforms report all
 * zeroes, so a label derived from it is best-effort by construction.
 */
object PasskeyTable : IdTable<String>("\"Passkey\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val name = text("name").nullable()
    val publicKey = text("publicKey")
    val userId = text("userId")
    val credentialID = text("credentialID")
    val counter = integer("counter").default(0)
    val deviceType = text("deviceType")
    val backedUp = bool("backedUp").default(false)
    val transports = text("transports").nullable()
    val aaguid = text("aaguid").nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"Passkey_pkey\"")

    init {
        index("\"Passkey_userId_idx\"", false, userId)
        index("\"Passkey_credentialID_idx\"", false, credentialID)
    }
}
