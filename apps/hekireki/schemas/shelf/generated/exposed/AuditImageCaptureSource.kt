package shelf.models

/**
 * How an audit photo reached Shelf, as reported by the client. The phone app
 * reports CAMERA only for its own camera, which cannot hand back an old
 * picture, so it is the strongest evidence short of a modified client; a
 * library photo's date is only what its file claims; the web app cannot tell
 * a camera shot from a file; PHONE is a phone app version that does not say
 * which.
 */
enum class AuditImageCaptureSource(
    val dbName: String,
) {
    CAMERA("CAMERA"),
    LIBRARY("LIBRARY"),
    WEB_UPLOAD("WEB_UPLOAD"),
    PHONE("PHONE"),
}
