package shelf.models

enum class UpdateStatus(
    val dbName: String,
) {
    DRAFT("DRAFT"),
    PUBLISHED("PUBLISHED"),
}
