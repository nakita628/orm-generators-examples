package langfuse.models

enum class GatewaySpendLimitInterval(
    val dbName: String,
) {
    DAILY("DAILY"),
    WEEKLY("WEEKLY"),
    MONTHLY("MONTHLY"),
    TOTAL("TOTAL"),
}
