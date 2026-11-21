package shelf.models

/** * Determines how an asset is tracked throughout the system. Immutable after creation. */
enum class AssetType(
    val dbName: String,
) {
    INDIVIDUAL("INDIVIDUAL"),
    QUANTITY_TRACKED("QUANTITY_TRACKED"),
}
