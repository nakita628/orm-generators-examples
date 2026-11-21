package langwatch.models

/**
 * How a Slack connection posts: a bot token over the Web API (any channel it
 * is invited to, every block renders) or an incoming webhook (one channel).
 */
enum class SlackIntegrationKind(
    val dbName: String,
) {
    BOT("BOT"),
    INCOMING_WEBHOOK("INCOMING_WEBHOOK"),
}
