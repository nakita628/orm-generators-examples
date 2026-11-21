package shelf.models

/**
 * The kind of PDF a PdfTemplate lays out. Only the booking PDF has saved
 * layouts today; the check-in receipt and the audit receipt can join as new
 * kinds.
 */
enum class PdfTemplateKind(
    val dbName: String,
) {
    BOOKING("BOOKING"),
}
