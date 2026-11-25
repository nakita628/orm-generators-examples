package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeID
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.CompositeEntity
import org.jetbrains.exposed.v1.dao.CompositeEntityClass
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import java.time.Instant
import java.time.temporal.ChronoUnit

class OrganizationUserEntity(
    id: EntityID<CompositeID>,
) : CompositeEntity(id) {
    companion object : CompositeEntityClass<OrganizationUserEntity>(OrganizationUserTable)

    val userId by OrganizationUserTable.userId
    val organizationId by OrganizationUserTable.organizationId
    var role by OrganizationUserTable.role
    var createdAt by OrganizationUserTable.createdAt
    var updatedAt by OrganizationUserTable.updatedAt

    /**
     * Department assignment for spend attribution (the people lens).
     * Single-valued, org-scoped, accounting-only (never gates access).
     * Plain column ref (relationMode="prisma"); the service validates the
     * id belongs to this org. See departments.feature.
     */
    var departmentId by OrganizationUserTable.departmentId

    /**
     * A Billing-only member (ADR-205): role EXTERNAL, the Billing grant, no team and
     * no seat. Any role change clears it. See billing-only-invite.feature.
     */
    var billingOnly by OrganizationUserTable.billingOnly

    /**
     * Set when an admin disables this membership to stay within the licensed
     * seat count. The person loses access to this organization only, keeps
     * their role, department and everything they did, and frees their seat.
     * Reversible: re-enabling restores access if a seat is available.
     * Seat counting excludes disabled memberships. See
     * seat-reconciliation.feature.
     */
    var disabledAt by OrganizationUserTable.disabledAt

    /**
     * Stable grant intent for an unfinished automatic SSO admission.
     * Cleared after the authoritative grant is observed; never inferred from missing access.
     */
    var pendingSsoGrantId by OrganizationUserTable.pendingSsoGrantId

    /**
     * A new generation for each membership lifetime. Stamped USER grant facts
     * must match the current active membership before they can be projected;
     * removing and rejoining therefore cannot revive a delayed old attach.
     */
    var membershipStamp by OrganizationUserTable.membershipStamp

    /**
     * When this membership started holding the full seat it holds now; null means createdAt.
     * Anchors the 30-day seat hold (specs/licensing/seat-rotation.feature).
     */
    var fullSeatSince by OrganizationUserTable.fullSeatSince

    /**
     * Set when admission wrote this membership disabled because no seat was free: the person
     * is waiting for one. Cleared when the membership is enabled.
     */
    var seatPendingSince by OrganizationUserTable.seatPendingSince

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(OrganizationUserTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
