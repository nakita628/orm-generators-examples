package shelf.models

enum class AuditAssetStatus(
    val dbName: String,
) {
    PENDING("PENDING"),
    FOUND("FOUND"),
    MISSING("MISSING"),
    UNEXPECTED("UNEXPECTED"),
}
