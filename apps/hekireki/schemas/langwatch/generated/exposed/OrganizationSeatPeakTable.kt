package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable

/**
 * Organization's monthly seat peak, folded from its seat changes (modules/organization, ADR-002).
 * A read model, one row per organization: rebuildable by replaying the event log, so no relation.
 * Epochs in milliseconds.
 * Organization's seat-peak counts (ADR-002): one row per organization once counting opened.
 */
object OrganizationSeatPeakTable : IdTable<String>("\"OrganizationSeatPeak\"") {
    /** The organization, which is also the aggregate id. */
    override val id = text("organizationId").entityId()

    /** When the opening snapshot was taken, epoch milliseconds. */
    val openedAt = double("openedAt")
    val fullSeats = integer("fullSeats")
    val liteSeats = integer("liteSeats")
    val developerSeats = integer("developerSeats")

    /** The latest change time applied to any holder; invitations expire up to it. */
    val latestChangeAt = double("latestChangeAt")

    /** Changes whose `from` disagreed with the seat the fold held. */
    val mismatches = integer("mismatches")
    override val primaryKey = PrimaryKey(id, name = "\"OrganizationSeatPeak_pkey\"")
}
