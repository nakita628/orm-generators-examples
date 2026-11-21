package shelf.models

enum class KitStatus(
    val dbName: String,
) {
    AVAILABLE("AVAILABLE"),
    IN_CUSTODY("IN_CUSTODY"),
    CHECKED_OUT("CHECKED_OUT"),
}
