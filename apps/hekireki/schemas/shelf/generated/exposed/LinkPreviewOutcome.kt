package shelf.models

/** How the last fetch of a link's page went. */
enum class LinkPreviewOutcome(
    val dbName: String,
) {
    TITLE("TITLE"),
    NO_TITLE("NO_TITLE"),
    SIGN_IN("SIGN_IN"),
    BLOCKED("BLOCKED"),
    UNREACHABLE("UNREACHABLE"),
}
