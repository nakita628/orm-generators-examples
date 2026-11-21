package shelf.models

enum class ReminderRecurrenceUnit(
    val dbName: String,
) {
    DAY("DAY"),
    WEEK("WEEK"),
    MONTH("MONTH"),
    YEAR("YEAR"),
}
