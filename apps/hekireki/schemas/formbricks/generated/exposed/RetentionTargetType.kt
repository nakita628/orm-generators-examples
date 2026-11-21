package formbricks.models

/**
 * What a `RetentionRunItem` row points at. Response deletions are grouped into one `survey` row per
 * survey, with the number of responses in `count`.
 */
enum class RetentionTargetType(
    val dbName: String,
) {
    SURVEY("survey"),
    USER("user"),
}
