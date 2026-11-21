package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * A progressive checkout SESSION: one scan batch, covering N assets via the
 * positional `assetIds` / `quantities` / `bookingAssetIds` arrays.
 *
 * Records sessions and the units they claimed — NOT whether an asset is out.
 * The all-at-once checkout (`checkoutBooking`, the Check out button) flips
 * statuses and writes NO row here, so **absence proves nothing** and must never
 * be read as "was not checked out". Reading it that way is what let a single
 * later scan mark every button-checked-out asset on a booking as
 * "never checked out". `BookingAsset.checkedOutAt` answers that question.
 */
class PartialBookingCheckoutEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, PartialBookingCheckoutEntity>(PartialBookingCheckoutTable)

    var assetIds by PartialBookingCheckoutTable.assetIds
    var quantities by PartialBookingCheckoutTable.quantities
    var bookingAssetIds by PartialBookingCheckoutTable.bookingAssetIds
    var checkoutCount by PartialBookingCheckoutTable.checkoutCount
    var checkoutTimestamp by PartialBookingCheckoutTable.checkoutTimestamp
    var booking by BookingEntity referencedOn PartialBookingCheckoutTable.bookingId
    var bookingId by PartialBookingCheckoutTable.bookingId
    var checkedOutBy by UserEntity referencedOn PartialBookingCheckoutTable.checkedOutById
    var checkedOutById by PartialBookingCheckoutTable.checkedOutById
    var createdAt by PartialBookingCheckoutTable.createdAt
    var updatedAt by PartialBookingCheckoutTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(PartialBookingCheckoutTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
