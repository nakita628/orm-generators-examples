package langwatch.models

/**
 * Who can use a Slack connection (ADR-093 §5a): every project in the
 * organization, or one project. Chosen per connection.
 */
enum class SlackIntegrationScopeType(
    val dbName: String,
) {
    ORGANIZATION("ORGANIZATION"),
    PROJECT("PROJECT"),
}
