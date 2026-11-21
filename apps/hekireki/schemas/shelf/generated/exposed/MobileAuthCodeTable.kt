package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Native-app SSO web-delegated login — single-use session-handoff codes.
 * @see apps/webapp/app/routes/_auth+/oauth.callback.mobile.tsx
 */
object MobileAuthCodeTable : IdTable<String>("\"MobileAuthCode\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val userId =
        reference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"MobileAuthCode_userId_fkey\"",
        )
    val codeHash = text("codeHash").uniqueIndex("\"MobileAuthCode_codeHash_key\"")
    val codeChallenge = text("codeChallenge").nullable()
    val sessionCiphertext = text("sessionCiphertext").nullable()
    val consumedAt = pgTimestamp("consumedAt", 3).nullable()
    val expiresAt = pgTimestamp("expiresAt", 3)
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"MobileAuthCode_pkey\"")

    init {
        index("\"MobileAuthCode_userId_idx\"", false, userId)
        index("\"MobileAuthCode_expiresAt_idx\"", false, expiresAt)
    }
}
