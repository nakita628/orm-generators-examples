package langwatch.models

/** Organization signup intent (ADR-038): which product the org came for. */
enum class OrganizationIntent(
    val dbName: String,
) {
    AGENT_GOVERNANCE("AGENT_GOVERNANCE"),
    LLM_OPS("LLM_OPS"),
}
