package documenso.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Per-recipient cached credential metadata for the configured Cloud Signature
 * Consortium (CSC) provider. Holds the TSP-validated certificate chain and the
 * symmetric-encrypted service-scope access token. One row per recipient on a
 * CSC envelope; absent for SES recipients.
 */
object CscCredentialTable : IdTable<String>("\"CscCredential\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val providerId = text("providerId")
    val credentialId = text("credentialId")

    /**
     * TSP-validated certificate chain (length-prefixed). Nullable until the
     * service-scope OAuth callback successfully retrieves and validates it.
     */
    val certCache = binary("certCache").nullable()

    /**
     * Algorithm metadata derived from the CSC `credentials/info` response.
     * Held as flat scalars so each is independently queryable.
     */
    val signatureAlgorithm = text("signatureAlgorithm")
    val keyType = text("keyType")
    val digestAlgorithm = text("digestAlgorithm")
    val keyLenBits = integer("keyLenBits").nullable()
    val signAlgoParams = text("signAlgoParams").nullable()

    /**
     * Service-scope access token, symmetric-encrypted at the application layer
     * before persistence. Nullable until the service-scope OAuth callback runs.
     */
    val serviceTokenCiphertext = binary("serviceTokenCiphertext").nullable()
    val serviceTokenExpiresAt = pgTimestamp("serviceTokenExpiresAt", 3).nullable()
    val recipientId =
        reference(
            "recipientId",
            RecipientTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"CscCredential_recipientId_fkey\"",
        ).uniqueIndex("\"CscCredential_recipientId_key\"")
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"CscCredential_pkey\"")
}
