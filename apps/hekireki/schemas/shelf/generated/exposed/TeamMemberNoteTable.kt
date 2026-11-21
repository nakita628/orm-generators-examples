package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object TeamMemberNoteTable : IdTable<String>("\"TeamMemberNote\"") {
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
            fkName = "\"TeamMemberNote_userId_fkey\"",
        )
    val teamMemberId =
        reference(
            "teamMemberId",
            TeamMemberTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"TeamMemberNote_teamMemberId_fkey\"",
        )
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"TeamMemberNote_organizationId_fkey\"",
        )
    override val primaryKey = PrimaryKey(id, name = "\"TeamMemberNote_pkey\"")

    init {
        index("\"TeamMemberNote_teamMemberId_idx\"", false, teamMemberId)
        index("\"TeamMemberNote_userId_idx\"", false, userId)
        index("\"TeamMemberNote_organizationId_idx\"", false, organizationId)
        index("\"TeamMemberNote_teamMemberId_organizationId_idx\"", false, teamMemberId, organizationId)
    }
}
