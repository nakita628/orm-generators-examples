package shelf.models

/**
 * The format of a TEXT custom field. Unlike the type, it can change after the
 * field exists.
 */
enum class CustomFieldFormat(
    val dbName: String,
) {
    PLAIN("PLAIN"),
    LINK("LINK"),
}
