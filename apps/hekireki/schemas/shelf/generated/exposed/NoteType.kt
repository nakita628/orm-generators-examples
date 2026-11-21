package shelf.models

enum class NoteType(
    val dbName: String,
) {
    COMMENT("COMMENT"),
    UPDATE("UPDATE"),
}
