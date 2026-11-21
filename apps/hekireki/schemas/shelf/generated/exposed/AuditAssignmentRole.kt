package shelf.models

enum class AuditAssignmentRole(
    val dbName: String,
) {
    LEAD("LEAD"),
    PARTICIPANT("PARTICIPANT"),
}
