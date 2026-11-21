package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * Organization's monthly seat peak, folded from its seat changes (modules/organization, ADR-002).
 * A read model, one row per organization: rebuildable by replaying the event log, so no relation.
 * Epochs in milliseconds.
 * Organization's seat-peak counts (ADR-002): one row per organization once counting opened.
 */
class OrganizationSeatPeakEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, OrganizationSeatPeakEntity>(OrganizationSeatPeakTable)

    /** When the opening snapshot was taken, epoch milliseconds. */
    var openedAt by OrganizationSeatPeakTable.openedAt
    var fullSeats by OrganizationSeatPeakTable.fullSeats
    var liteSeats by OrganizationSeatPeakTable.liteSeats
    var developerSeats by OrganizationSeatPeakTable.developerSeats

    /** The latest change time applied to any holder; invitations expire up to it. */
    var latestChangeAt by OrganizationSeatPeakTable.latestChangeAt

    /** Changes whose `from` disagreed with the seat the fold held. */
    var mismatches by OrganizationSeatPeakTable.mismatches
}
