package langfuse.models

enum class SessionViewOwnership(
    val dbName: String,
) {
    PRIVATE("PRIVATE"),
    PROJECT("PROJECT"),
}
