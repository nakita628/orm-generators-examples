package formbricks.models

/**
 * The author's intent only. Scheduled, Active, "Paused with its survey" and Closed after the
 * end time are derived (`apps/web/modules/ee/creatives/lib/display-status.ts` and its SQL twin); nothing
 * writes this column when a schedule ends or a survey pauses.
 */
enum class CreativeStatus(
    val dbName: String,
) {
    DRAFT("draft"),
    IN_PROGRESS("inProgress"),
    PAUSED("paused"),
    COMPLETED("completed"),
}
