package formbricks.models

/** The surveys policy's conditions. Every ticked one must hold, and they share the policy's period. */
enum class RetentionSurveyCondition(
    val dbName: String,
) {
    NO_RESPONSE("noResponse"),
    NO_CHANGE("noChange"),
    CREATED_BEFORE("createdBefore"),
}
