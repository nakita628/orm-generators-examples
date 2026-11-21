package shelf.models

/** Lifecycle of a background asset import. */
enum class AssetImportStatus(
    val dbName: String,
) {
    DRAFT("DRAFT"),
    QUEUED("QUEUED"),
    RUNNING("RUNNING"),
    COMPLETED("COMPLETED"),
    FAILED("FAILED"),
}
