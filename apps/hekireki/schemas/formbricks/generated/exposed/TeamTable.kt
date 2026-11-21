package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Represents a team within an organization.
 * Enables group-based access control and collaboration.
 *
 * @property id - Unique identifier for the team
 * @property name - Display name of the team
 * @property organization - The parent organization
 * @property teamUsers - Users who are part of this team
 * @property workspaceTeams - Workspaces this team has access to
 */
object TeamTable : IdTable<String>("\"Team\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updated_at", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val name = text("name")
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Team_organizationId_fkey\"",
        )
    override val primaryKey = PrimaryKey(id, name = "\"Team_pkey\"")

    init {
        uniqueIndex("\"Team_organizationId_name_key\"", organizationId, name)
    }
}
