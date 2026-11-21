package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object UpdateTable : IdTable<String>("\"Update\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val title = text("title")
    val content = text("content")
    val url = text("url").nullable()
    val imageUrl = text("imageUrl").nullable()
    val publishDate = pgTimestamp("publishDate", 3)
    val status = pgEnum("status", "\"UpdateStatus\"", UpdateStatus.entries, UpdateStatus::dbName).default(UpdateStatus.DRAFT)
    val targetRoles =
        array("targetRoles", PgEnumColumnType("\"OrganizationRoles\"", OrganizationRoles.entries, OrganizationRoles::dbName)).nullable()
    val clickCount = integer("clickCount").default(0)
    val viewCount = integer("viewCount").default(0)
    val createdById =
        reference(
            "createdById",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Update_createdById_fkey\"",
        )
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"Update_pkey\"")

    init {
        index("\"Update_status_publishDate_idx\"", false, status, publishDate)
        index("\"Update_publishDate_idx\"", false, publishDate)
        index("\"Update_createdById_idx\"", false, createdById)
    }
}
