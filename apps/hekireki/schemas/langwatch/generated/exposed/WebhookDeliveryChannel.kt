package langwatch.models

/**
 * Which webhook pipeline wrote a delivery row, and therefore which columns
 * carry its tenancy: `platform` rows anchor on organizationId + endpointId,
 * `automations` rows on projectId + triggerId.
 */
enum class WebhookDeliveryChannel(
    val dbName: String,
) {
    PLATFORM("platform"),
    AUTOMATIONS("automations"),
}
