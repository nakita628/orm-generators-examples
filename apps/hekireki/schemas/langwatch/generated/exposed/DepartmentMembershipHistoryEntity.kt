package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * One dated person → department link (ADR-128 §13, #7882).
 *
 * `OrganizationUser.departmentId` stays the CURRENT assignment every screen
 * reads; these rows are its history, appended by the same `assignUser` write
 * so the two cannot drift. They exist for the cost reads: January's spend
 * stays with January's department, so a read resolves against the link that
 * was open on the cost's day, not against today's pointer. Members with no
 * link on a day resolve to "unassigned", never to an error.
 *
 * Same shape as IdentityMatch's dated links, same one-open-link guard: the
 * partial unique index in the migration rejects a second open row per
 * (organizationId, userId) with SQLSTATE 23505.
 */
class DepartmentMembershipHistoryEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, DepartmentMembershipHistoryEntity>(DepartmentMembershipHistoryTable)

    var organizationId by DepartmentMembershipHistoryTable.organizationId
    var userId by DepartmentMembershipHistoryTable.userId
    var departmentId by DepartmentMembershipHistoryTable.departmentId
    var validFrom by DepartmentMembershipHistoryTable.validFrom

    /** Null means the link is open. Reassignment and clearing close it. */
    var validTo by DepartmentMembershipHistoryTable.validTo
    var createdAt by DepartmentMembershipHistoryTable.createdAt
}
