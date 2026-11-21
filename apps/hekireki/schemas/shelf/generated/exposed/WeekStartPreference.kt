package shelf.models

enum class WeekStartPreference(
    val dbName: String,
) {
    MONDAY("MONDAY"),
    SUNDAY("SUNDAY"),
    SATURDAY("SATURDAY"),
}
