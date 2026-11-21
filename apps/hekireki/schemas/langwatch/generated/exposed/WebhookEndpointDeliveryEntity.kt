package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

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
class WebhookEndpointDeliveryEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, WebhookEndpointDeliveryEntity>(WebhookEndpointDeliveryTable)

    var channel by WebhookEndpointDeliveryTable.channel

    /** Platform tenancy. Null on automations rows. */
    var organizationId by WebhookEndpointDeliveryTable.organizationId
    var endpointId by WebhookEndpointDeliveryTable.endpointId

    /** Automations tenancy. Null on platform rows. */
    var projectId by WebhookEndpointDeliveryTable.projectId
    var triggerId by WebhookEndpointDeliveryTable.triggerId
    var dispatchId by WebhookEndpointDeliveryTable.dispatchId

    /**
     * 1-based attempt and the batch's envelope count. Both null on automations
     * rows, which record neither: one trigger fire is one event, and the
     * channel's sender does not thread its attempt number into the log.
     */
    var attempt by WebhookEndpointDeliveryTable.attempt
    var eventCount by WebhookEndpointDeliveryTable.eventCount
    var responseStatus by WebhookEndpointDeliveryTable.responseStatus
    var latencyMs by WebhookEndpointDeliveryTable.latencyMs
    var error by WebhookEndpointDeliveryTable.error
    var response by WebhookEndpointDeliveryTable.response
    var outcome by WebhookEndpointDeliveryTable.outcome
    var firedAt by WebhookEndpointDeliveryTable.firedAt
    var createdAt by WebhookEndpointDeliveryTable.createdAt
}
