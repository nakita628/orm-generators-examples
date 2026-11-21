package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * *
 * * Org-scoped SCIM external ID linking a Shelf user to an identity-provider
 * * object within a specific organization.  A single Shelf user can appear in
 * * multiple organizations, each managed by a different identity provider, so the
 * * external ID must be scoped to the (user, organization) pair.
 */
class UserScimExternalIdEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, UserScimExternalIdEntity>(UserScimExternalIdTable)

    var userId by UserScimExternalIdTable.userId
    var organizationId by UserScimExternalIdTable.organizationId
    var scimExternalId by UserScimExternalIdTable.scimExternalId
    var createdAt by UserScimExternalIdTable.createdAt
    var updatedAt by UserScimExternalIdTable.updatedAt
    var user by UserEntity referencedOn UserScimExternalIdTable.userId
    var organization by OrganizationEntity referencedOn UserScimExternalIdTable.organizationId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(UserScimExternalIdTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
