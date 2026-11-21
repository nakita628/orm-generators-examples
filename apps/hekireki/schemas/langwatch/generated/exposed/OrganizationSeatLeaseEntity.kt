package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * Held by the one seat-taking write running on an organization (seat-requests.feature). A
 * row past expiresAt is free to take; a holder deletes its own row when its write is done.
 */
class OrganizationSeatLeaseEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, OrganizationSeatLeaseEntity>(OrganizationSeatLeaseTable)

    /** Random per hold, so a release deletes only its own lease. */
    var holder by OrganizationSeatLeaseTable.holder
    var expiresAt by OrganizationSeatLeaseTable.expiresAt
    var createdAt by OrganizationSeatLeaseTable.createdAt
}
