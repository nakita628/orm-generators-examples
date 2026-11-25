package langfuse.models

enum class GatewayGuardrailType(
    val dbName: String,
) {
    REGEX("REGEX"),
    BEDROCK("BEDROCK"),
}
