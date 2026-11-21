package langwatch.models

/** Who sees a dashboard, named as the prompt scope is, with one more level. */
enum class DashboardScope(
    val dbName: String,
) {
    PRIVATE("PRIVATE"),
    PROJECT("PROJECT"),
    ORGANIZATION("ORGANIZATION"),
}
