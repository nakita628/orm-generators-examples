package formbricks.models

/** The declared type of an Embedded Data field's value. */
enum class EmbeddedDataType(
    val dbName: String,
) {
    STRING("string"),
    NUMBER("number"),
    BOOLEAN("boolean"),
    DATE("date"),
}
