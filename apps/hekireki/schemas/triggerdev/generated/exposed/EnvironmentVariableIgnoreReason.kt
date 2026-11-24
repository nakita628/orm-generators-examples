package triggerdev.models

enum class EnvironmentVariableIgnoreReason(
    val dbName: String,
) {
    EMPTY("EMPTY"),
    MISSING("MISSING"),
    OIDC("OIDC"),
}
