package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeID
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.CompositeEntity
import org.jetbrains.exposed.v1.dao.CompositeEntityClass

/** One member's or pending invitation's seat as the seat peak last applied it (ADR-002). */
class OrganizationSeatPeakHolderEntity(
    id: EntityID<CompositeID>,
) : CompositeEntity(id) {
    companion object : CompositeEntityClass<OrganizationSeatPeakHolderEntity>(OrganizationSeatPeakHolderTable)

    val organizationId by OrganizationSeatPeakHolderTable.organizationId

    /** `member`, `invitation` or `hold` (a seat held after somebody left). */
    val holderKind by OrganizationSeatPeakHolderTable.holderKind

    /** The member's or the hold's userId, or the invitation's id. */
    val holderId by OrganizationSeatPeakHolderTable.holderId

    /** `full`, `lite`, `developer` or `none`. */
    var seat by OrganizationSeatPeakHolderTable.seat
    var lastChangedAt by OrganizationSeatPeakHolderTable.lastChangedAt

    /** Absent for a seat the opening snapshot recorded. */
    var lastChangeId by OrganizationSeatPeakHolderTable.lastChangeId

    /** An invitation's expiry; null for a member or an invitation that never expires. */
    var expiresAt by OrganizationSeatPeakHolderTable.expiresAt

    /** The fold gave the invitation's seat back at `expiresAt`. */
    var expiredByFold by OrganizationSeatPeakHolderTable.expiredByFold

    /** A hold whose seat a live full invitation to its address took: that invitation's id. */
    var suppressedBy by OrganizationSeatPeakHolderTable.suppressedBy

    /** The fold gave a suppressed hold its seat back when that invitation expired. */
    var restoredByFold by OrganizationSeatPeakHolderTable.restoredByFold
}
