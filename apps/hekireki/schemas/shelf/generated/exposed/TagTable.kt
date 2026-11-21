package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object TagTable : IdTable<String>("\"Tag\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val name = text("name")
    val description = text("description").nullable()
    val color = text("color").nullable()
    val useFor =
        array("useFor", PgEnumColumnType("\"TagUseFor\"", TagUseFor.entries, TagUseFor::dbName))
            .nullable()
            .databaseDefault("'{\"ASSET\"}'") { listOf(TagUseFor.ASSET) }
    val userId =
        reference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.RESTRICT,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Tag_userId_fkey\"",
        )
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Tag_organizationId_fkey\"",
        )
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"Tag_pkey\"")

    init {
        index("\"Tag_organizationId_idx\"", false, organizationId)
        index("\"Tag_userId_idx\"", false, userId)
        index("\"Tag_name_trgm_idx\"", false, functions = listOf(indexExpression("\"name\" gin_trgm_ops")), indexType = "gin")
    }
}
