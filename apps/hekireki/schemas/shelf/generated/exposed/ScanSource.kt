package shelf.models

/** The surface that recorded a `Scan`. */
enum class ScanSource(
    val dbName: String,
) {
    QR_LINK("QR_LINK"),
    WEB_SCANNER("WEB_SCANNER"),
    WEB_DRAWER("WEB_DRAWER"),
    AUDIT("AUDIT"),
    COMPANION("COMPANION"),
    GPS_UPDATE("GPS_UPDATE"),
}
