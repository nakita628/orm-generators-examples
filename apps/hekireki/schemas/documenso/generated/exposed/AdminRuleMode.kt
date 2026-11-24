package documenso.models

enum class AdminRuleMode(
    val dbName: String,
) {
    ENFORCE("ENFORCE"),
    MONITOR("MONITOR"),
    DISABLED("DISABLED"),
}
