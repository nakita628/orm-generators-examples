package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.TextColumnType
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * A progressive check-in SESSION, mirroring PartialBookingCheckout.
 *
 * `assetIds` lists only FULLY reconciled assets: INDIVIDUAL always, and
 * QUANTITY_TRACKED only once its remaining units hit zero (partial returns
 * live in `ConsumptionLog`). Completing a booking outright writes no row here,
 * so the same warning applies — absence is not evidence of anything.
 * `BookingAsset.checkedInAt` is the per-slice fact.
 */
object PartialBookingCheckinTable : IdTable<String>("\"PartialBookingCheckin\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val assetIds = array("assetIds", TextColumnType()).nullable()
    val checkinCount = integer("checkinCount")
    val checkinTimestamp =
        pgTimestamptz("checkinTimestamp", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val bookingId =
        reference(
            "bookingId",
            BookingTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"PartialBookingCheckin_bookingId_fkey\"",
        )
    val checkedInById =
        reference(
            "checkedInById",
            UserTable,
            onDelete = ReferenceOption.RESTRICT,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"PartialBookingCheckin_checkedInById_fkey\"",
        )
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"PartialBookingCheckin_pkey\"")

    init {
        index("\"PartialBookingCheckin_bookingId_idx\"", false, bookingId)
        index("\"PartialBookingCheckin_checkedInById_idx\"", false, checkedInById)
        index("\"PartialBookingCheckin_checkinTimestamp_idx\"", false, checkinTimestamp)
        index("\"PartialBookingCheckin_bookingId_checkinTimestamp_idx\"", false, bookingId, checkinTimestamp)
    }
}
