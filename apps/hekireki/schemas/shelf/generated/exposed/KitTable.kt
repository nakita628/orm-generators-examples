package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object KitTable : IdTable<String>("\"Kit\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val name = text("name")
    val description = text("description").nullable()
    val status = pgEnum("status", "\"KitStatus\"", KitStatus.entries, KitStatus::dbName).default(KitStatus.AVAILABLE)
    val image = text("image").nullable()
    val imageExpiration = pgTimestamp("imageExpiration", 3).nullable()
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Kit_organizationId_fkey\"",
        )
    val createdById =
        reference(
            "createdById",
            UserTable,
            onDelete = ReferenceOption.RESTRICT,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Kit_createdById_fkey\"",
        )
    val categoryId =
        optReference(
            "categoryId",
            CategoryTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Kit_categoryId_fkey\"",
        )
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val locationId =
        optReference(
            "locationId",
            LocationTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Kit_locationId_fkey\"",
        )
    override val primaryKey = PrimaryKey(id, name = "\"Kit_pkey\"")

    init {
        index("\"Kit_createdById_idx\"", false, createdById)
        index("\"Kit_organizationId_idx\"", false, organizationId)
        index("\"Kit_categoryId_organizationId_idx\"", false, categoryId, organizationId)
        index("\"Kit_categoryId_organizationId_createdAt_idx\"", false, categoryId, organizationId, createdAt)
        index("\"Kit_categoryId_organizationId_name_idx\"", false, categoryId, organizationId, name)
        index("\"Kit_categoryId_organizationId_status_idx\"", false, categoryId, organizationId, status)
    }
}
