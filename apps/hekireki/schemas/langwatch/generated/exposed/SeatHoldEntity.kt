package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * A full seat still counted after the person who took it stopped holding it (removed,
 * disabled or demoted), until 30 days after they took it (specs/licensing/seat-rotation.feature).
 * One per organization and person; a return within the hold takes the seat back.
 */
class SeatHoldEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, SeatHoldEntity>(SeatHoldTable)

    var organizationId by SeatHoldTable.organizationId
    var userId by SeatHoldTable.userId

    /** Display copy for the Users list, kept because a removed person has no membership row. */
    var email by SeatHoldTable.email
    var name by SeatHoldTable.name

    /** When the person took the seat this hold keeps. */
    var seatTakenAt by SeatHoldTable.seatTakenAt

    /** seatTakenAt plus 30 days; expired holds are ignored by every count and deleted hourly. */
    var heldUntil by SeatHoldTable.heldUntil

    /** REMOVED | DISABLED | DEMOTED */
    var reason by SeatHoldTable.reason
    var createdAt by SeatHoldTable.createdAt
}
