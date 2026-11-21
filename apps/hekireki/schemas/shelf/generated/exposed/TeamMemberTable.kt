package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object TeamMemberTable : IdTable<String>("\"TeamMember\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val name = text("name")
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"TeamMember_organizationId_fkey\"",
        )
    val userId =
        optReference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"TeamMember_userId_fkey\"",
        )
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val deletedAt = pgTimestamp("deletedAt", 3).nullable()
    override val primaryKey = PrimaryKey(id, name = "\"TeamMember_pkey\"")

    init {
        index("\"TeamMember_name_idx\"", false, functions = listOf(indexExpression("\"name\" gin_trgm_ops")), indexType = "gin")
        index("\"TeamMember_organizationId_idx\"", false, organizationId)
        index("\"TeamMember_userId_idx\"", false, userId)
    }
}
