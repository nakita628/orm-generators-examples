package shelf.models

enum class ActivityEntity(
    val dbName: String,
) {
    ASSET("ASSET"),
    BOOKING("BOOKING"),
    AUDIT("AUDIT"),
    KIT("KIT"),
    LOCATION("LOCATION"),
    TEAM_MEMBER("TEAM_MEMBER"),
    CUSTODY("CUSTODY"),
    USER("USER"),
    ORGANIZATION("ORGANIZATION"),
}
