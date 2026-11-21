package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Represents a feedback directory (Hub tenant) owned by an organization.
 * Directories group feedback data and are assigned to workspaces for access control.
 *
 * @property id - Unique identifier for the directory
 * @property name - Display name of the directory
 * @property isArchived - Soft delete flag
 * @property organization - The parent organization
 * @property workspaces - Workspaces assigned to this directory
 */
object FeedbackDirectoryTable : IdTable<String>("\"FeedbackDirectory\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updated_at", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val name = text("name")
    val isArchived = bool("isArchived").default(false)
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"FeedbackDirectory_organizationId_fkey\"",
        )
    override val primaryKey = PrimaryKey(id, name = "\"FeedbackDirectory_pkey\"")

    init {
        uniqueIndex("\"FeedbackDirectory_organizationId_name_key\"", organizationId, name)
    }
}
