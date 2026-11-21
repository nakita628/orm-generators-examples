package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object UserOrganizationTable : IdTable<String>("\"UserOrganization\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val userId =
        reference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"UserOrganization_userId_fkey\"",
        )
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"UserOrganization_organizationId_fkey\"",
        )
    val roles = array("roles", PgEnumColumnType("\"OrganizationRoles\"", OrganizationRoles.entries, OrganizationRoles::dbName)).nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }

    /**
     * Secret, unguessable token for this member's iCal calendar-feed subscription
     * in this workspace. Null until the user generates a feed; regenerating it
     * rotates the value (invalidating any previously-shared subscription URL).
     * `@unique` doubles as the lookup index for the public feed route.
     */
    val calendarTokenId = text("calendarTokenId").nullable().uniqueIndex("\"UserOrganization_calendarTokenId_key\"")
    override val primaryKey = PrimaryKey(id, name = "\"UserOrganization_pkey\"")

    init {
        index("\"UserOrganization_organizationId_idx\"", false, organizationId)
        uniqueIndex("\"UserOrganization_userId_organizationId_key\"", userId, organizationId)
    }
}
