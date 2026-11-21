package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * A full seat still counted after the person who took it stopped holding it (removed,
 * disabled or demoted), until 30 days after they took it (specs/licensing/seat-rotation.feature).
 * One per organization and person; a return within the hold takes the seat back.
 */
object SeatHoldTable : IdTable<String>("\"SeatHold\"") {
    override val id = text("id").entityId()
    val organizationId = text("organizationId")
    val userId = text("userId")

    /** Display copy for the Users list, kept because a removed person has no membership row. */
    val email = text("email").nullable()
    val name = text("name").nullable()

    /** When the person took the seat this hold keeps. */
    val seatTakenAt = pgTimestamp("seatTakenAt", 3)

    /** seatTakenAt plus 30 days; expired holds are ignored by every count and deleted hourly. */
    val heldUntil = pgTimestamp("heldUntil", 3)

    /** REMOVED | DISABLED | DEMOTED */
    val reason = text("reason")
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"SeatHold_pkey\"")

    init {
        index("\"SeatHold_heldUntil_idx\"", false, heldUntil)
        uniqueIndex("\"SeatHold_organizationId_userId_key\"", organizationId, userId)
    }
}
