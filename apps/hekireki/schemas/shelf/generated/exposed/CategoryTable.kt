package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object CategoryTable : IdTable<String>("\"Category\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val name = text("name")
    val description = text("description").nullable()
    val color = text("color")
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val userId =
        reference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.RESTRICT,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Category_userId_fkey\"",
        )
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Category_organizationId_fkey\"",
        )
    override val primaryKey = PrimaryKey(id, name = "\"Category_pkey\"")

    init {
        index("\"Category_organizationId_idx\"", false, organizationId)
        index("\"Category_userId_idx\"", false, userId)
        index("\"Category_name_trgm_idx\"", false, functions = listOf(indexExpression("\"name\" gin_trgm_ops")), indexType = "gin")
    }
}
