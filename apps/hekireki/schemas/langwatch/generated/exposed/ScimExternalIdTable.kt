package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * D08 — who the directory means, per connection.
 *
 * The identity provider's own identifier for a person, which is what
 * survives their email changing: an address is not identity, people change
 * theirs, and the same person carries different identifiers on two different
 * connections. So the PAIR is the key and `externalId` alone never is.
 *
 * No relation to `SsoConnection` (that projection carries no foreign keys);
 * the `User` relation is real, so a deleted user takes their directory
 * identities with them rather than stranding a mapping that resolves to
 * nobody.
 */
object ScimExternalIdTable : IdTable<String>("\"ScimExternalId\"") {
    val organizationId = text("organizationId")
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val connectionId = text("connectionId")
    val externalId = text("externalId")
    val userId = text("userId")
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"ScimExternalId_pkey\"")

    init {
        index("\"ScimExternalId_organizationId_userId_idx\"", false, organizationId, userId)
        index("\"ScimExternalId_userId_idx\"", false, userId)
        index("\"ScimExternalId_connectionId_idx\"", false, connectionId)
        uniqueIndex("\"ScimExternalId_connectionId_externalId_key\"", connectionId, externalId)
    }
}
