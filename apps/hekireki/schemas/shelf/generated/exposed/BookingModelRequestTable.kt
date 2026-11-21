package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
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
object BookingModelRequestTable : IdTable<String>("\"BookingModelRequest\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val bookingId =
        reference(
            "bookingId",
            BookingTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"BookingModelRequest_bookingId_fkey\"",
        )
    val assetModelId =
        reference(
            "assetModelId",
            AssetModelTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"BookingModelRequest_assetModelId_fkey\"",
        )

    /**
     * Total units originally reserved. Mutated only by explicit
     * operator edits (upsert) — never decremented on materialisation.
     * This preserves the audit trail of "how many were planned".
     */
    val quantity = integer("quantity")

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
    val fulfilledQuantity = integer("fulfilledQuantity").default(0)

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
    val fulfilledAt = pgTimestamp("fulfilledAt", 3).nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"BookingModelRequest_pkey\"")

    init {
        index("\"BookingModelRequest_assetModelId_idx\"", false, assetModelId)
        uniqueIndex("\"BookingModelRequest_bookingId_assetModelId_key\"", bookingId, assetModelId)
    }
}
