package triggerdev.models

enum class TaskExecutionMode(
    val dbName: String,
) {
    STANDARD("STANDARD"),
    SANDBOX("SANDBOX"),
}
