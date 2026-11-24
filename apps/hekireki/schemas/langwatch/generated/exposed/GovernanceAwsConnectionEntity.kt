package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * An AWS credential a Bedrock kill switch target signs with: a static key pair, a role to assume,
 * or both. The secret access key and the external id are stored sealed.
 */
class GovernanceAwsConnectionEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, GovernanceAwsConnectionEntity>(GovernanceAwsConnectionTable)

    var organizationId by GovernanceAwsConnectionTable.organizationId
    var name by GovernanceAwsConnectionTable.name
    var accessKeyId by GovernanceAwsConnectionTable.accessKeyId
    var encryptedSecretAccessKey by GovernanceAwsConnectionTable.encryptedSecretAccessKey
    var roleArn by GovernanceAwsConnectionTable.roleArn
    var encryptedExternalId by GovernanceAwsConnectionTable.encryptedExternalId
    var createdAt by GovernanceAwsConnectionTable.createdAt
    var updatedAt by GovernanceAwsConnectionTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(GovernanceAwsConnectionTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
