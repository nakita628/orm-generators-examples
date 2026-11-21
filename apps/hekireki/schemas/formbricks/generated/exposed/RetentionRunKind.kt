package formbricks.models

/**
 * What wrote a `RetentionRun`: the nightly sweep, or the archive purge recording the deletion of surveys
 * the surveys policy archived. Only sweep runs count as the sweep having run (its lease, its warning
 * restart after a gap, where its scan resumes, the health check).
 */
enum class RetentionRunKind(
    val dbName: String,
) {
    SWEEP("sweep"),
    PURGE("purge"),
}
