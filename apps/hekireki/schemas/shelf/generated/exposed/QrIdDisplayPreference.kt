package shelf.models

enum class QrIdDisplayPreference(
    val dbName: String,
) {
    QR_ID("QR_ID"),
    SAM_ID("SAM_ID"),
    Code128("Code128"),
    Code39("Code39"),
    DataMatrix("DataMatrix"),
    ExternalQR("ExternalQR"),
    EAN13("EAN13"),
}
