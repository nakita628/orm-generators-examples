package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * One browser a person subscribed to Web Push, owned by notification. A
 * person holds one row per device; the endpoint names the browser, so a
 * browser that subscribes again replaces its own row.
 */
object WebPushSubscriptionTable : IdTable<String>("\"WebPushSubscription\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()

    /**
     * The person the pushes are for. A plain column: the row leaves with the
     * person through notification's own subscriber on the user lifecycle.
     */
    val userId = text("userId")

    /** The push service URL the browser handed out. Unique per browser. */
    val endpoint = text("endpoint").uniqueIndex("\"WebPushSubscription_endpoint_key\"")

    /**
     * The browser's P-256 public key and auth secret, base64url, as the
     * browser reports them. Payloads are encrypted to these.
     */
    val p256dh = text("p256dh")
    val auth = text("auth")
    val userAgent = text("userAgent").nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }

    /** When the push service last accepted a push for this browser. */
    val lastSuccessAt = pgTimestamp("lastSuccessAt", 3).nullable()
    override val primaryKey = PrimaryKey(id, name = "\"WebPushSubscription_pkey\"")

    init {
        index("\"WebPushSubscription_userId_idx\"", false, userId)
    }
}
