package formbricks.models

/**
 * Who may read a survey. `workspace` is everyone with workspace access (the behaviour before ENG-3282);
 * `restricted` is the owner and the organization's owners and managers only.
 */
enum class SurveyVisibility(
    val dbName: String,
) {
    RESTRICTED("restricted"),
    WORKSPACE("workspace"),
}
