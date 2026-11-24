package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * An AWS credential a Bedrock kill switch target signs with: a static key pair, a role to assume,
 * or both. The secret access key and the external id are stored sealed.
 */
object GovernanceAwsConnectionTable : IdTable<String>("\"GovernanceAwsConnection\"") {
    override val id = text("id").entityId()
    val organizationId = text("organizationId")
    val name = text("name")
    val accessKeyId = text("accessKeyId").nullable()
    val encryptedSecretAccessKey = text("encryptedSecretAccessKey").nullable()
    val roleArn = text("roleArn").nullable()
    val encryptedExternalId = text("encryptedExternalId").nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"GovernanceAwsConnection_pkey\"")

    init {
        index("\"GovernanceAwsConnection_organizationId_idx\"", false, organizationId)
    }
}
