package formbricks.models

/**
 * Defines permission levels for API keys.
 * Controls what operations an API key can perform.
 */
enum class ApiKeyPermission(
    val dbName: String,
) {
    READ("read"),
    WRITE("write"),
    MANAGE("manage"),
}
