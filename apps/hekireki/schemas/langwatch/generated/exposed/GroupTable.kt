package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object GroupTable : IdTable<String>("\"Group\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val organizationId = text("organizationId")
    val name = text("name")
    val slug = text("slug")
    val externalId = text("externalId").nullable()
    val scimSource = text("scimSource").nullable()

    /**
     * D08: which SSO connection provisioned this group. Null for manually
     * created groups and for the SCIM groups that predate connection scoping.
     */
    val scimConnectionId = text("scimConnectionId").nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"Group_pkey\"")

    init {
        index("\"Group_organizationId_idx\"", false, organizationId)
        index("\"Group_scimConnectionId_idx\"", false, scimConnectionId)
        uniqueIndex("\"Group_organizationId_slug_key\"", organizationId, slug)
        uniqueIndex("\"Group_scimConnectionId_externalId_key\"", scimConnectionId, externalId)
    }
}
