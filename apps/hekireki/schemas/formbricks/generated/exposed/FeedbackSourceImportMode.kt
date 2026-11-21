package formbricks.models

/**
 * Which survey responses a Formbricks feedback source imports into Hub.
 * completedOnly matches the live ingestion path, which only runs on responseFinished.
 * all additionally imports partial responses — answers a respondent typed but never submitted.
 */
enum class FeedbackSourceImportMode(
    val dbName: String,
) {
    COMPLETED_ONLY("completedOnly"),
    ALL("all"),
}
