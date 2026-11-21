package langwatch.models

/** How the invoicing of one seat change ended. */
enum class ConnectedSeatChangeState(
    val dbName: String,
) {
    INTENT("intent"),
    INVOICED("invoiced"),
    NOTHING_TO_INVOICE("nothing_to_invoice"),
    NOT_ONBOARDED("not_onboarded"),
}
