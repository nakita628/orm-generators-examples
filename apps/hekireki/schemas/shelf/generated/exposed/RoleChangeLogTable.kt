package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object RoleChangeLogTable : IdTable<String>("\"RoleChangeLog\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val previousRole = pgEnum("previousRole", "\"OrganizationRoles\"", OrganizationRoles.entries, OrganizationRoles::dbName)
    val newRole = pgEnum("newRole", "\"OrganizationRoles\"", OrganizationRoles.entries, OrganizationRoles::dbName)

    /** How the change was made. SSO rows name the member themself as `changedBy`. */
    val source1 =
        pgEnum("source", "\"RoleChangeSource\"", RoleChangeSource.entries, RoleChangeSource::dbName).default(RoleChangeSource.MANUAL)
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val userId =
        reference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.RESTRICT,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"RoleChangeLog_userId_fkey\"",
        )
    val changedById =
        reference(
            "changedById",
            UserTable,
            onDelete = ReferenceOption.RESTRICT,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"RoleChangeLog_changedById_fkey\"",
        )
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.RESTRICT,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"RoleChangeLog_organizationId_fkey\"",
        )
    override val primaryKey = PrimaryKey(id, name = "\"RoleChangeLog_pkey\"")

    init {
        index("\"RoleChangeLog_userId_organizationId_idx\"", false, userId, organizationId)
        index("\"RoleChangeLog_organizationId_createdAt_idx\"", false, organizationId, createdAt)
    }
}
