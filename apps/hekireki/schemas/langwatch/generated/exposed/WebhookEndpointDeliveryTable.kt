package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * The per-ATTEMPT delivery log for BOTH webhook channels: the endpoints
 * platform's batches and the automations channel's trigger fires. One row per
 * HTTP attempt; `dispatchId` groups every attempt of one logical send (the
 * value of the X-LangWatch-Delivery-Id or X-LangWatch-Event-Id header). A slim
 * facts table: outcome, status, latency, a capped error message, and a
 * truncated copy of a failure response. Our request content (URL, headers,
 * body) is never stored. Pruned after 30 days.
 *
 * The two channels used to keep separate tables with the same columns and two
 * prune implementations. They are one table with a discriminator now, so the
 * tenancy columns of the channel that did not write a row are null.
 */
object WebhookEndpointDeliveryTable : IdTable<String>("\"WebhookEndpointDelivery\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val channel =
        pgEnum("channel", "\"WebhookDeliveryChannel\"", WebhookDeliveryChannel.entries, WebhookDeliveryChannel::dbName).default(
            WebhookDeliveryChannel.PLATFORM,
        )

    /** Platform tenancy. Null on automations rows. */
    val organizationId = text("organizationId").nullable()
    val endpointId = text("endpointId").nullable()

    /** Automations tenancy. Null on platform rows. */
    val projectId = text("projectId").nullable()
    val triggerId = text("triggerId").nullable()
    val dispatchId = text("dispatchId")

    /**
     * 1-based attempt and the batch's envelope count. Both null on automations
     * rows, which record neither: one trigger fire is one event, and the
     * channel's sender does not thread its attempt number into the log.
     */
    val attempt = integer("attempt").nullable()
    val eventCount = integer("eventCount").nullable()
    val responseStatus = integer("responseStatus").nullable()
    val latencyMs = integer("latencyMs").nullable()
    val error = text("error").nullable()
    val response = jsonb<String>("response", { it }, { it }).nullable()
    val outcome = pgEnum("outcome", "\"WebhookDeliveryOutcome\"", WebhookDeliveryOutcome.entries, WebhookDeliveryOutcome::dbName)
    val firedAt = pgTimestamp("firedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"WebhookEndpointDelivery_pkey\"")

    init {
        index("\"WebhookEndpointDelivery_endpointId_firedAt_idx\"", false, endpointId, firedAt)
        index("\"WebhookEndpointDelivery_organizationId_firedAt_idx\"", false, organizationId, firedAt)
        index("\"WebhookEndpointDelivery_firedAt_idx\"", false, firedAt)
        index("\"WebhookEndpointDelivery_endpointId_dispatchId_idx\"", false, endpointId, dispatchId)
        index("\"WebhookEndpointDelivery_projectId_triggerId_firedAt_idx\"", false, projectId, triggerId, firedAt)
        index("\"WebhookEndpointDelivery_projectId_triggerId_dispatchId_idx\"", false, projectId, triggerId, dispatchId)
    }
}
