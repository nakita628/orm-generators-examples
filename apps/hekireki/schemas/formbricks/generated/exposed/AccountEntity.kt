package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
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
class AccountEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AccountEntity>(AccountTable)

    var createdAt by AccountTable.createdAt
    var updatedAt by AccountTable.updatedAt
    var user by UserEntity referencedOn AccountTable.userId
    var userId by AccountTable.userId
    var type by AccountTable.type
    var provider by AccountTable.provider
    var providerAccountId by AccountTable.providerAccountId
    var accessToken by AccountTable.accessToken
    var refreshToken by AccountTable.refreshToken
    var expiresAt by AccountTable.expiresAt
    var extExpiresIn by AccountTable.extExpiresIn
    var tokenType by AccountTable.tokenType
    var scope by AccountTable.scope
    var idToken by AccountTable.idToken
    var sessionState by AccountTable.sessionState
    var password by AccountTable.password
    var accessTokenExpiresAt by AccountTable.accessTokenExpiresAt
    var refreshTokenExpiresAt by AccountTable.refreshTokenExpiresAt

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
    var issuer by AccountTable.issuer

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AccountTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
