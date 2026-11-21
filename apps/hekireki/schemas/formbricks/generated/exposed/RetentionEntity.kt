package formbricks.models

/** Data retention (ENG-3697). The kind of data a policy governs. */
enum class RetentionEntity(
    val dbName: String,
) {
    RESPONSES("responses"),
    SURVEYS("surveys"),
    MEMBERS("members"),
}
