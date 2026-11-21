package formbricks.models

/** What a reviewer decided about a response. One value today; null on `Response.outcome` means none. */
enum class ResponseOutcome(
    val dbName: String,
) {
    LOW_QUALITY("lowQuality"),
}
