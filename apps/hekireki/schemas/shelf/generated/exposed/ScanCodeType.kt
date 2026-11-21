package shelf.models

/** The kind of code a `Scan` recorded. */
enum class ScanCodeType(
    val dbName: String,
) {
    QR("QR"),
    BARCODE("BARCODE"),
    SAM_ID("SAM_ID"),
}
