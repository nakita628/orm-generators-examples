package shelf.models

/** Which step of the background import runs next (or failed). */
enum class AssetImportPhase(
    val dbName: String,
) {
    PREPARE("PREPARE"),
    ROWS("ROWS"),
    FINALIZE("FINALIZE"),
}
