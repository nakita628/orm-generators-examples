package langfuse.models

enum class GatewaySpendLimitAlertType(
    val dbName: String,
) {
    WARNING("WARNING"),
    EXCEEDED("EXCEEDED"),
}
