package shelf.models

enum class BarcodeType(
    val dbName: String,
) {
    Code128("Code128"),
    Code39("Code39"),
    DataMatrix("DataMatrix"),
    ExternalQR("ExternalQR"),
    EAN13("EAN13"),
}
