package documenso.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Per-recipient cached credential metadata for the configured Cloud Signature
 * Consortium (CSC) provider. Holds the TSP-validated certificate chain and the
 * symmetric-encrypted service-scope access token. One row per recipient on a
 * CSC envelope; absent for SES recipients.
 */
class CscCredentialEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, CscCredentialEntity>(CscCredentialTable)

    var providerId by CscCredentialTable.providerId
    var credentialId by CscCredentialTable.credentialId

    /**
     * TSP-validated certificate chain (length-prefixed). Nullable until the
     * service-scope OAuth callback successfully retrieves and validates it.
     */
    var certCache by CscCredentialTable.certCache

    /**
     * Algorithm metadata derived from the CSC `credentials/info` response.
     * Held as flat scalars so each is independently queryable.
     */
    var signatureAlgorithm by CscCredentialTable.signatureAlgorithm
    var keyType by CscCredentialTable.keyType
    var digestAlgorithm by CscCredentialTable.digestAlgorithm
    var keyLenBits by CscCredentialTable.keyLenBits
    var signAlgoParams by CscCredentialTable.signAlgoParams

    /**
     * Service-scope access token, symmetric-encrypted at the application layer
     * before persistence. Nullable until the service-scope OAuth callback runs.
     */
    var serviceTokenCiphertext by CscCredentialTable.serviceTokenCiphertext
    var serviceTokenExpiresAt by CscCredentialTable.serviceTokenExpiresAt
    var recipientId by CscCredentialTable.recipientId
    var recipient by RecipientEntity referencedOn CscCredentialTable.recipientId
    var createdAt by CscCredentialTable.createdAt
    var updatedAt by CscCredentialTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(CscCredentialTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
