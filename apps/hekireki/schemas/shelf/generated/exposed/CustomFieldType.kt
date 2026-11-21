package shelf.models

enum class CustomFieldType(
    val dbName: String,
) {
    TEXT("TEXT"),
    OPTION("OPTION"),
    BOOLEAN("BOOLEAN"),
    DATE("DATE"),
    MULTILINE_TEXT("MULTILINE_TEXT"),
    AMOUNT("AMOUNT"),
    NUMBER("NUMBER"),
}
