package langwatch.models

/** Where a license in the registry came from. */
enum class IssuedLicenseSource(
    val dbName: String,
) {
    BACKOFFICE("BACKOFFICE"),
    PURCHASE("PURCHASE"),
    SCRIPT("SCRIPT"),
    LEGACY_IMPORT("LEGACY_IMPORT"),
}
