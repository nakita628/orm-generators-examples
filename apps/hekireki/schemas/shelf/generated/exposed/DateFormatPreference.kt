package shelf.models

enum class DateFormatPreference(
    val dbName: String,
) {
    DD_MM_YYYY("DD_MM_YYYY"),
    MM_DD_YYYY("MM_DD_YYYY"),
    YYYY_MM_DD("YYYY_MM_DD"),
    MMM_DD_YYYY("MMM_DD_YYYY"),
    DD_MMM_YYYY("DD_MMM_YYYY"),
}
