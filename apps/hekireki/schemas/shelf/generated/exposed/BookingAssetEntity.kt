package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * *
 * * Explicit pivot table between Booking and Asset.
 * * Replaces the implicit M2M (_AssetToBooking) via a rename migration in Phase 3.
 */
class BookingAssetEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, BookingAssetEntity>(BookingAssetTable)

    var booking by BookingEntity referencedOn BookingAssetTable.bookingId
    var bookingId by BookingAssetTable.bookingId
    var asset by AssetEntity referencedOn BookingAssetTable.assetId
    var assetId by BookingAssetTable.assetId
    var quantity by BookingAssetTable.quantity
    var bookingModelRequestId by BookingAssetTable.bookingModelRequestId
    var assetKitId by BookingAssetTable.assetKitId
    var sourceKitId by BookingAssetTable.sourceKitId
    var checkedOutAt by BookingAssetTable.checkedOutAt
    var checkedOutById by BookingAssetTable.checkedOutById
    var checkedInAt by BookingAssetTable.checkedInAt
    var checkedInById by BookingAssetTable.checkedInById
    var checkedOutQuantity by BookingAssetTable.checkedOutQuantity
    var sourceLocationId by BookingAssetTable.sourceLocationId
}
