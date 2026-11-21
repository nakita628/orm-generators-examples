package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.TextColumnType
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object CustomFieldTable : IdTable<String>("\"CustomField\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val name = text("name")
    val helpText = text("helpText").nullable()
    val required = bool("required").default(false)
    val active = bool("active").default(true)
    val type = pgEnum("type", "\"CustomFieldType\"", CustomFieldType.entries, CustomFieldType::dbName).default(CustomFieldType.TEXT)
    val options1 = array("options", TextColumnType()).nullable()

    /**
     * How a TEXT field's values are read. LINK requires a web link in every
     * value; any other type stays PLAIN (enforced by a CHECK constraint).
     */
    val format =
        pgEnum("format", "\"CustomFieldFormat\"", CustomFieldFormat.entries, CustomFieldFormat::dbName).default(CustomFieldFormat.PLAIN)
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"CustomField_organizationId_fkey\"",
        )
    val userId =
        reference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.RESTRICT,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"CustomField_userId_fkey\"",
        )
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val deletedAt = pgTimestamp("deletedAt", 3).nullable()
    override val primaryKey = PrimaryKey(id, name = "\"CustomField_pkey\"")

    init {
        index("\"CustomField_organizationId_idx\"", false, organizationId)
        index("\"CustomField_userId_idx\"", false, userId)
        index("\"CustomField_organizationId_deletedAt_idx\"", false, organizationId, deletedAt)
    }
}
