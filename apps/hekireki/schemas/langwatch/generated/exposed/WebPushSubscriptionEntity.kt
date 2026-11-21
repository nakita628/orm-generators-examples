package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * One browser a person subscribed to Web Push, owned by notification. A
 * person holds one row per device; the endpoint names the browser, so a
 * browser that subscribes again replaces its own row.
 */
class WebPushSubscriptionEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, WebPushSubscriptionEntity>(WebPushSubscriptionTable)

    /**
     * The person the pushes are for. A plain column: the row leaves with the
     * person through notification's own subscriber on the user lifecycle.
     */
    var userId by WebPushSubscriptionTable.userId

    /** The push service URL the browser handed out. Unique per browser. */
    var endpoint by WebPushSubscriptionTable.endpoint

    /**
     * The browser's P-256 public key and auth secret, base64url, as the
     * browser reports them. Payloads are encrypted to these.
     */
    var p256dh by WebPushSubscriptionTable.p256dh
    var auth by WebPushSubscriptionTable.auth
    var userAgent by WebPushSubscriptionTable.userAgent
    var createdAt by WebPushSubscriptionTable.createdAt

    /** When the push service last accepted a push for this browser. */
    var lastSuccessAt by WebPushSubscriptionTable.lastSuccessAt
}
