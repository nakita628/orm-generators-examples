package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Book-by-Model: intent row attached to a booking for a quantity of
 * assets of a given AssetModel, without yet picking specific assets.
 * Concrete `BookingAsset` rows are only created at
 * scan-to-assign time; when the last unit is assigned, the request
 * row is deleted. `quantity` is the number of units still needing
 * assignment (starts at the full requested total, decrements on scan).
 *
 * Keeping model requests separate from `BookingAsset` means the ~40
 * code paths that assume `BookingAsset.assetId` points to a real
 * asset stay untouched — model-only bookings become concrete the
 * moment they're checked out.
 */
class BookingModelRequestEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, BookingModelRequestEntity>(BookingModelRequestTable)

    var booking by BookingEntity referencedOn BookingModelRequestTable.bookingId
    var bookingId by BookingModelRequestTable.bookingId
    var assetModel by AssetModelEntity referencedOn BookingModelRequestTable.assetModelId
    var assetModelId by BookingModelRequestTable.assetModelId

    /**
     * Total units originally reserved. Mutated only by explicit
     * operator edits (upsert) — never decremented on materialisation.
     * This preserves the audit trail of "how many were planned".
     */
    var quantity by BookingModelRequestTable.quantity

    /**
     * Units already assigned to concrete `BookingAsset` rows via
     * `materializeModelRequestForAsset`. Scans increment it.
     *
     * It DOES decrease: `removeAssets` (booking service) decrements it by the
     * number of removed `BookingAsset` rows whose `bookingModelRequestId`
     * points at this request, floored at 0, and clears `fulfilledAt` when the
     * result drops below `quantity`.
     *
     * Counted from that column, never from a shared `assetModelId`: an asset
     * added directly alongside a reservation shares the model without having
     * discharged anything, so grouping by model re-opened reservations that
     * were still correctly satisfied.
     */
    var fulfilledQuantity by BookingModelRequestTable.fulfilledQuantity

    /**
     * Timestamp of the moment the reservation became fully fulfilled — set
     * when `fulfilledQuantity` catches up to `quantity`, and preserved (not
     * re-stamped) if an unchanged quantity is saved again. Cleared back to
     * null when outstanding units reappear: the operator edits `quantity`
     * upwards, or `removeAssets` drops `fulfilledQuantity` below it.
     *
     * The invariant that makes this the primary "is this outstanding?" filter:
     * non-null IFF `fulfilledQuantity === quantity`, so `fulfilledAt IS NULL`
     * means units remain to materialise. `fulfilledQuantity` can never exceed
     * `quantity` (the claim in `materializeModelRequestForAsset` is a single
     * statement guarded by `WHERE "fulfilledQuantity" < "quantity"`; upsert
     * rejects a shrink below it), so any decrement re-opens the row.
     */
    var fulfilledAt by BookingModelRequestTable.fulfilledAt
    var createdAt by BookingModelRequestTable.createdAt
    var updatedAt by BookingModelRequestTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(BookingModelRequestTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
