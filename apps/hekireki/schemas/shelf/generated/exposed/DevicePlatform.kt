package shelf.models

/** The operating system a companion-app push token was issued for. */
enum class DevicePlatform(
    val dbName: String,
) {
    IOS("IOS"),
    ANDROID("ANDROID"),
}
