package shelf.models

enum class TierId(
    val dbName: String,
) {
    FREE("free"),
    TIER_1("tier_1"),
    TIER_2("tier_2"),
    CUSTOM("custom"),
}
