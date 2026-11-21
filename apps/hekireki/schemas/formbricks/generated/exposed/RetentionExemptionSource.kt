package formbricks.models

/**
 * Who made an exemption: a person (with a reason), or a restore from the archive under a surveys policy
 * whose only condition is "created before" (`restartSurveyRetention`), which carries no reason.
 */
enum class RetentionExemptionSource(
    val dbName: String,
) {
    MANUAL("manual"),
    RESTORE("restore"),
}
