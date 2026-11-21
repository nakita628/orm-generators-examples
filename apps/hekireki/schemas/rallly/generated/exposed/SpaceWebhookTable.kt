package rallly.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.TextColumnType
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object SpaceWebhookTable : IdTable<String>("space_webhooks") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val spaceId =
        reference(
            "space_id",
            SpaceTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "space_webhooks_space_id_fkey",
        )
    val url = text("url")

    /** AES-256-GCM under SECRET_PASSWORD (see @rallly/utils/encryption); the plaintext is shown once at creation */
    val secret = text("secret")

    /** Subscribed event names (`poll.closed`, ...); validated by the feature schema, not a DB enum, so the vocabulary grows without migrations */
    val events = array("events", TextColumnType()).nullable()
    val enabled = bool("enabled").default(true)

    /** The payload contract this endpoint was built against (WEBHOOK_VERSION at creation). No default: a writer that forgets it fails instead of recording a stale contract. A breaking payload change branches on this so existing endpoints keep their shape */
    val version = text("version")

    /** Fan-out position in the poll activity log: activities created after this instant are still to be delivered. Starts at creation so a new endpoint never replays history */
    val cursor = pgTimestamp("cursor", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val lastDeliveredAt = pgTimestamp("last_delivered_at", 3).nullable()
    val lastFailedAt = pgTimestamp("last_failed_at", 3).nullable()

    /** Exhausted deliveries in a row; reset by any success. The dispatcher disables the endpoint once it reaches its cap */
    val consecutiveFailures = integer("consecutive_failures").default(0)
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updated_at", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "space_webhooks_pkey")

    init {
        index("space_webhooks_space_id_idx", false, spaceId)
    }
}
