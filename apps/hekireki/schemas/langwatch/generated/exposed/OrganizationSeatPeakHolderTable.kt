package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeIdTable

/** One member's or pending invitation's seat as the seat peak last applied it (ADR-002). */
object OrganizationSeatPeakHolderTable : CompositeIdTable("\"OrganizationSeatPeakHolder\"") {
    val organizationId = text("organizationId").entityId()

    /** `member`, `invitation` or `hold` (a seat held after somebody left). */
    val holderKind = text("holderKind").entityId()

    /** The member's or the hold's userId, or the invitation's id. */
    val holderId = text("holderId").entityId()

    /** `full`, `lite`, `developer` or `none`. */
    val seat = text("seat")
    val lastChangedAt = double("lastChangedAt")

    /** Absent for a seat the opening snapshot recorded. */
    val lastChangeId = text("lastChangeId").nullable()

    /** An invitation's expiry; null for a member or an invitation that never expires. */
    val expiresAt = double("expiresAt").nullable()

    /** The fold gave the invitation's seat back at `expiresAt`. */
    val expiredByFold = bool("expiredByFold").default(false)

    /** A hold whose seat a live full invitation to its address took: that invitation's id. */
    val suppressedBy = text("suppressedBy").nullable()

    /** The fold gave a suppressed hold its seat back when that invitation expired. */
    val restoredByFold = bool("restoredByFold").default(false)
    override val primaryKey = PrimaryKey(organizationId, holderKind, holderId, name = "\"OrganizationSeatPeakHolder_pkey\"")

    init {
        index("OrganizationSeatPeakHolder_organizationId_holderKind_seat_e_idx", false, organizationId, holderKind, seat, expiresAt)
        index("\"OrganizationSeatPeakHolder_organizationId_suppressedBy_idx\"", false, organizationId, suppressedBy)
    }
}
