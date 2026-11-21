package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object ImageTable : IdTable<String>("\"Image\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val contentType = text("contentType")
    val altText = text("altText").nullable()
    val blob = binary("blob")
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val ownerOrgId =
        reference(
            "ownerOrgId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Image_ownerOrgId_fkey\"",
        )
    val userId =
        reference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.RESTRICT,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Image_userId_fkey\"",
        )
    override val primaryKey = PrimaryKey(id, name = "\"Image_pkey\"")

    init {
        index("\"Image_ownerOrgId_idx\"", false, ownerOrgId)
        index("\"Image_userId_idx\"", false, userId)
    }
}
