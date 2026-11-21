package shelf.models

/** How a membership's role came to change. */
enum class RoleChangeSource(
    val dbName: String,
) {
    MANUAL("MANUAL"),
    SSO("SSO"),
}
