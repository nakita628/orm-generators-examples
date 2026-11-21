package rallly.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class SpaceWebhookEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, SpaceWebhookEntity>(SpaceWebhookTable)

    var spaceId by SpaceWebhookTable.spaceId
    var url by SpaceWebhookTable.url

    /** AES-256-GCM under SECRET_PASSWORD (see @rallly/utils/encryption); the plaintext is shown once at creation */
    var secret by SpaceWebhookTable.secret

    /** Subscribed event names (`poll.closed`, ...); validated by the feature schema, not a DB enum, so the vocabulary grows without migrations */
    var events by SpaceWebhookTable.events
    var enabled by SpaceWebhookTable.enabled

    /** The payload contract this endpoint was built against (WEBHOOK_VERSION at creation). No default: a writer that forgets it fails instead of recording a stale contract. A breaking payload change branches on this so existing endpoints keep their shape */
    var version by SpaceWebhookTable.version

    /** Fan-out position in the poll activity log: activities created after this instant are still to be delivered. Starts at creation so a new endpoint never replays history */
    var cursor by SpaceWebhookTable.cursor
    var lastDeliveredAt by SpaceWebhookTable.lastDeliveredAt
    var lastFailedAt by SpaceWebhookTable.lastFailedAt

    /** Exhausted deliveries in a row; reset by any success. The dispatcher disables the endpoint once it reaches its cap */
    var consecutiveFailures by SpaceWebhookTable.consecutiveFailures
    var createdAt by SpaceWebhookTable.createdAt
    var updatedAt by SpaceWebhookTable.updatedAt
    var space by SpaceEntity referencedOn SpaceWebhookTable.spaceId
    val deliveries by WebhookDeliveryEntity referrersOn WebhookDeliveryTable.webhookId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(SpaceWebhookTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
