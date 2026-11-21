package shelf.models

enum class BookingStatus(
    val dbName: String,
) {
    DRAFT("DRAFT"),
    RESERVED("RESERVED"),
    ONGOING("ONGOING"),
    OVERDUE("OVERDUE"),
    COMPLETE("COMPLETE"),
    ARCHIVED("ARCHIVED"),
    CANCELLED("CANCELLED"),
}
