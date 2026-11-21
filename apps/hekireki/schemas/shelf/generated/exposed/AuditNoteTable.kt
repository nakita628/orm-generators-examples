package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object AuditNoteTable : IdTable<String>("\"AuditNote\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val content = text("content")
    val type = pgEnum("type", "\"NoteType\"", NoteType.entries, NoteType::dbName).default(NoteType.COMMENT)
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val userId =
        optReference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AuditNote_userId_fkey\"",
        )
    val auditSessionId =
        reference(
            "auditSessionId",
            AuditSessionTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AuditNote_auditSessionId_fkey\"",
        )
    val auditAssetId =
        optReference(
            "auditAssetId",
            AuditAssetTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AuditNote_auditAssetId_fkey\"",
        )
    override val primaryKey = PrimaryKey(id, name = "\"AuditNote_pkey\"")

    init {
        index("\"AuditNote_auditSessionId_idx\"", false, auditSessionId)
        index("\"AuditNote_userId_idx\"", false, userId)
        index("\"AuditNote_auditAssetId_idx\"", false, auditAssetId)
    }
}
