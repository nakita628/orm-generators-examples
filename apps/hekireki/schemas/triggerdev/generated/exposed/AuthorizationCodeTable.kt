package triggerdev.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/** Used to generate PersonalAccessTokens, they're one-time use */
object AuthorizationCodeTable : IdTable<String>("\"AuthorizationCode\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val code = text("code").uniqueIndex("\"AuthorizationCode_code_key\"")
    val personalAccessTokenId =
        optReference(
            "personalAccessTokenId",
            PersonalAccessTokenTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AuthorizationCode_personalAccessTokenId_fkey\"",
        )

    /** Client IP of the request that minted the code, compared with the approving browser's IP. Deleted with the code once it expires. */
    val requesterIp = text("requesterIp").nullable()

    /** S256 PKCE challenge from the client that minted the code. Null for codes minted by CLIs that predate it. */
    val codeChallenge = text("codeChallenge").nullable()

    /** Loopback port the approving browser is redirected to with the approval code. Null means the user pastes it. */
    val callbackPort = integer("callbackPort").nullable()

    /** sha256 of the approval code, which only the approving browser ever sees. */
    val approvalCodeHash = text("approvalCodeHash").nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"AuthorizationCode_pkey\"")
}
