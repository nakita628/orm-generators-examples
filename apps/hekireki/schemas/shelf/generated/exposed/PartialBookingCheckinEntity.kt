package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
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
class PartialBookingCheckinEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, PartialBookingCheckinEntity>(PartialBookingCheckinTable)

    var assetIds by PartialBookingCheckinTable.assetIds
    var checkinCount by PartialBookingCheckinTable.checkinCount
    var checkinTimestamp by PartialBookingCheckinTable.checkinTimestamp
    var booking by BookingEntity referencedOn PartialBookingCheckinTable.bookingId
    var bookingId by PartialBookingCheckinTable.bookingId
    var checkedInBy by UserEntity referencedOn PartialBookingCheckinTable.checkedInById
    var checkedInById by PartialBookingCheckinTable.checkedInById
    var createdAt by PartialBookingCheckinTable.createdAt
    var updatedAt by PartialBookingCheckinTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(PartialBookingCheckinTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
