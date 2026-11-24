package triggerdev.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.TextColumnType
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object OAuthRefreshTokenTable : IdTable<String>("\"OAuthRefreshToken\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val hashedToken = text("hashedToken").uniqueIndex("\"OAuthRefreshToken_hashedToken_key\"")
    val personalAccessTokenId =
        reference(
            "personalAccessTokenId",
            PersonalAccessTokenTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"OAuthRefreshToken_personalAccessTokenId_fkey\"",
        )
    val oauthClientId =
        reference(
            "oauthClientId",
            OAuthClientTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"OAuthRefreshToken_oauthClientId_fkey\"",
        )
    val scope = array("scope", TextColumnType()).nullable()
    val resource = text("resource")
    val usedAt = pgTimestamp("usedAt", 3).nullable()
    val replacedById = text("replacedById").nullable()
    val idleExpiresAt = pgTimestamp("idleExpiresAt", 3)
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"OAuthRefreshToken_pkey\"")

    init {
        index("\"OAuthRefreshToken_personalAccessTokenId_idx\"", false, personalAccessTokenId)
        index("\"OAuthRefreshToken_oauthClientId_idx\"", false, oauthClientId)
        index("\"OAuthRefreshToken_idleExpiresAt_idx\"", false, idleExpiresAt)
    }
}
