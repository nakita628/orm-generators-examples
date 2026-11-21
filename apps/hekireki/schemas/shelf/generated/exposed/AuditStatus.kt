package shelf.models

enum class AuditStatus(
    val dbName: String,
) {
    PENDING("PENDING"),
    ACTIVE("ACTIVE"),
    COMPLETED("COMPLETED"),
    CANCELLED("CANCELLED"),
    ARCHIVED("ARCHIVED"),
}
