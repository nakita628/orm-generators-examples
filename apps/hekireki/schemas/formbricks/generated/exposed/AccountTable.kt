package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Stores third-party authentication account information.
 * Enables OAuth and other external authentication methods.
 *
 * @property id - Unique identifier for the account
 * @property user - The Formbricks user who owns this account
 * @property provider - The authentication provider (GitHub, Google, etc.)
 * @property providerAccountId - User ID from the provider
 * @property access_token - OAuth access token
 * @property refresh_token - OAuth refresh token
 */
object AccountTable : IdTable<String>("\"Account\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updated_at", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val userId =
        reference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Account_userId_fkey\"",
        )
    val type = text("type").nullable()
    val provider = text("provider")
    val providerAccountId = text("providerAccountId")
    val accessToken = text("access_token").nullable()
    val refreshToken = text("refresh_token").nullable()
    val expiresAt = integer("expires_at").nullable()
    val extExpiresIn = integer("ext_expires_in").nullable()
    val tokenType = text("token_type").nullable()
    val scope = text("scope").nullable()
    val idToken = text("id_token").nullable()
    val sessionState = text("session_state").nullable()
    val password = text("password").nullable()
    val accessTokenExpiresAt = pgTimestamp("accessTokenExpiresAt", 3).nullable()
    val refreshTokenExpiresAt = pgTimestamp("refreshTokenExpiresAt", 3).nullable()

    /**
     * Better Auth 1.7.0 (ENG-2343): the account's identity namespace. 1.7.0 keyed accounts on
     * (issuer, accountId) and filtered every lookup on it — findAccountByKey, findAccountOwnerByKey,
     * findCredentialAccount and updatePassword — so on 1.7.0 a null here locks the user out.
     * 1.7.3 dropped the field and 1.7.7 keys accounts on (provider, providerAccountId) again
     * (ENG-3880). We keep writing it (see `account.create.before` in better-auth-hooks.ts) so a
     * rollback to 1.7.0 still finds every row; drop it once that rollback is off the table.
     *
     * NULLABLE on purpose. A NOT NULL column would fail every INSERT from a 1.6 pod still serving
     * during a rolling deploy. The migration backfills every existing row; see the migration SQL for
     * the value derivation.
     */
    val issuer = text("issuer").nullable()
    override val primaryKey = PrimaryKey(id, name = "\"Account_pkey\"")

    init {
        index("\"Account_userId_idx\"", false, userId)
        uniqueIndex("\"Account_provider_providerAccountId_key\"", provider, providerAccountId)
        uniqueIndex("\"Account_issuer_providerAccountId_key\"", issuer, providerAccountId)
    }
}
