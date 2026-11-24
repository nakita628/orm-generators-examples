package langfuse.models

enum class ModelScope(
    val dbName: String,
) {
    LANGFUSE("LANGFUSE"),
    ORGANIZATION("ORGANIZATION"),
    PROJECT("PROJECT"),
}
