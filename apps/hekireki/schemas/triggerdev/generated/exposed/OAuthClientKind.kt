package triggerdev.models

enum class OAuthClientKind(
    val dbName: String,
) {
    PREREGISTERED("PREREGISTERED"),
    CIMD("CIMD"),
    DCR("DCR"),
}
