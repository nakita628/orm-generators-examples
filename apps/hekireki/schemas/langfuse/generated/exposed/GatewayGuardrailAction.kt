package langfuse.models

enum class GatewayGuardrailAction(
    val dbName: String,
) {
    BLOCK("BLOCK"),
    LOG("LOG"),
    REDACT("REDACT"),
}
