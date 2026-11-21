package langwatch.models

/** What an invoice of a connected customer is for. */
enum class ConnectedInvoiceKind(
    val dbName: String,
) {
    ANNUAL("annual"),
    SEAT_CHANGE("seat_change"),
    USAGE("usage"),
}
