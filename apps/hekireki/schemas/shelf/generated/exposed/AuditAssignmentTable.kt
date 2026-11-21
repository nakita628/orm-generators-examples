package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object AuditAssignmentTable : IdTable<String>("\"AuditAssignment\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val auditSessionId =
        reference(
            "auditSessionId",
            AuditSessionTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AuditAssignment_auditSessionId_fkey\"",
        )
    val userId =
        reference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AuditAssignment_userId_fkey\"",
        )
    val role = pgEnum("role", "\"AuditAssignmentRole\"", AuditAssignmentRole.entries, AuditAssignmentRole::dbName).nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"AuditAssignment_pkey\"")

    init {
        index("\"AuditAssignment_userId_idx\"", false, userId)
        uniqueIndex("\"AuditAssignment_auditSessionId_userId_key\"", auditSessionId, userId)
    }
}
