package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
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
class ScimExternalIdEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, ScimExternalIdEntity>(ScimExternalIdTable)

    var organizationId by ScimExternalIdTable.organizationId
    var connectionId by ScimExternalIdTable.connectionId
    var externalId by ScimExternalIdTable.externalId
    var userId by ScimExternalIdTable.userId
    var createdAt by ScimExternalIdTable.createdAt
    var updatedAt by ScimExternalIdTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(ScimExternalIdTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
