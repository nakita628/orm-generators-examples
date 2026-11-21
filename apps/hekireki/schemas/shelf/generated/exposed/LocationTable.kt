package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object LocationTable : IdTable<String>("\"Location\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val name = text("name")
    val description = text("description").nullable()
    val address = text("address").nullable()
    val latitude = double("latitude").nullable()
    val longitude = double("longitude").nullable()
    val imageUrl = text("imageUrl").nullable()
    val thumbnailUrl = text("thumbnailUrl").nullable()
    val imageId =
        optReference(
            "imageId",
            ImageTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Location_imageId_fkey\"",
        ).uniqueIndex("\"Location_imageId_key\"")
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val userId =
        reference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Location_userId_fkey\"",
        )
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Location_organizationId_fkey\"",
        )
    val parentId =
        optReference(
            "parentId",
            LocationTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Location_parentId_fkey\"",
        )
    override val primaryKey = PrimaryKey(id, name = "\"Location_pkey\"")

    init {
        index("\"Location_organizationId_idx\"", false, organizationId)
        index("\"Location_userId_idx\"", false, userId)
        index("\"Location_organizationId_parentId_idx\"", false, organizationId, parentId)
        index("\"Location_name_trgm_idx\"", false, functions = listOf(indexExpression("\"name\" gin_trgm_ops")), indexType = "gin")
        index(
            "\"Location_description_trgm_idx\"",
            false,
            functions = listOf(indexExpression("\"description\" gin_trgm_ops")),
            indexType = "gin",
        )
        index("\"Location_address_trgm_idx\"", false, functions = listOf(indexExpression("\"address\" gin_trgm_ops")), indexType = "gin")
    }
}
