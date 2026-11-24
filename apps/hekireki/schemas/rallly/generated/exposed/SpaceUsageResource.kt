package rallly.models

enum class SpaceUsageResource(
    val dbName: String,
) {
    API_POLL("api_poll"),
    EMAIL_INVITE("email_invite"),
}
