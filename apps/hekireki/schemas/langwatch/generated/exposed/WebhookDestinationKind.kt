package langwatch.models

/**
 * The last hop a webhook endpoint delivers over. Lowercase because these
 * values are the wire spelling: the same strings appear in the REST DTO,
 * both SDKs and the CLI.
 */
enum class WebhookDestinationKind(
    val dbName: String,
) {
    HTTP("http"),
    SQS("sqs"),
}
