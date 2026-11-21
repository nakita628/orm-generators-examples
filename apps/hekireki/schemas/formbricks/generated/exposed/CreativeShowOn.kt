package formbricks.models

/**
 * What makes a creative appear. A creative left with no action never shows; it never falls back to
 * every page view.
 */
enum class CreativeShowOn(
    val dbName: String,
) {
    PAGE_VIEW("pageView"),
    ACTION("action"),
}
