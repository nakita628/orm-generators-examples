package langfuse.models

enum class GatewaySpendLimitScope(
    val dbName: String,
) {
    API_KEY("API_KEY"),
    USER("USER"),
    API_KEY_DEFAULT("API_KEY_DEFAULT"),
    USER_DEFAULT("USER_DEFAULT"),
    ORGANIZATION("ORGANIZATION"),
}
