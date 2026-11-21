package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * *
 * * Org-scoped SCIM external ID linking a Shelf user to an identity-provider
 * * object within a specific organization.  A single Shelf user can appear in
 * * multiple organizations, each managed by a different identity provider, so the
 * * external ID must be scoped to the (user, organization) pair.
 */
object UserScimExternalIdTable : IdTable<String>("\"UserScimExternalId\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val userId =
        reference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"UserScimExternalId_userId_fkey\"",
        )
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"UserScimExternalId_organizationId_fkey\"",
        )
    val scimExternalId = text("scimExternalId")
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"UserScimExternalId_pkey\"")

    init {
        index("\"UserScimExternalId_organizationId_idx\"", false, organizationId)
        uniqueIndex("\"UserScimExternalId_userId_organizationId_key\"", userId, organizationId)
        uniqueIndex("\"UserScimExternalId_organizationId_scimExternalId_key\"", organizationId, scimExternalId)
    }
}
