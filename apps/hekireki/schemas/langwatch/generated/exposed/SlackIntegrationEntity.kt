package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * One named Slack connection (ADR-093 §5a); an organization holds as many as
 * it likes, each scoped to the organization or to one project. The secret is
 * AES-256-GCM ciphertext and never returned to a client. `secretFingerprint`
 * is an HMAC of the plaintext, so one secret is one connection per scope
 * without decrypting anything. No foreign key: `scopeId` is
 * polymorphic (ADR-021), as on RetentionPolicy and CustomLLMModelCost.
 */
class SlackIntegrationEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, SlackIntegrationEntity>(SlackIntegrationTable)

    var name by SlackIntegrationTable.name
    var kind by SlackIntegrationTable.kind
    var scopeType by SlackIntegrationTable.scopeType
    var scopeId by SlackIntegrationTable.scopeId
    var organizationId by SlackIntegrationTable.organizationId
    var botTokenEncrypted by SlackIntegrationTable.botTokenEncrypted
    var webhookUrlEncrypted by SlackIntegrationTable.webhookUrlEncrypted
    var secretFingerprint by SlackIntegrationTable.secretFingerprint
    var secretHint by SlackIntegrationTable.secretHint
    var slackTeamId by SlackIntegrationTable.slackTeamId
    var slackTeamName by SlackIntegrationTable.slackTeamName
    var createdById by SlackIntegrationTable.createdById
    var updatedById by SlackIntegrationTable.updatedById
    var createdAt by SlackIntegrationTable.createdAt
    var updatedAt by SlackIntegrationTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(SlackIntegrationTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
