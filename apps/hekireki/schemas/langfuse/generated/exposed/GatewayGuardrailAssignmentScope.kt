package langfuse.models

enum class GatewayGuardrailAssignmentScope(
    val dbName: String,
) {
    ORGANIZATION("ORGANIZATION"),
    API_KEY("API_KEY"),
    USER("USER"),
}
