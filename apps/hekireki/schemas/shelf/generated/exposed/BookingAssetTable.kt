package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable

/**
 * *
 * * Explicit pivot table between Booking and Asset.
 * * Replaces the implicit M2M (_AssetToBooking) via a rename migration in Phase 3.
 */
object BookingAssetTable : IdTable<String>("\"BookingAsset\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val bookingId =
        reference(
            "bookingId",
            BookingTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"BookingAsset_bookingId_fkey\"",
        )
    val assetId =
        reference(
            "assetId",
            AssetTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"BookingAsset_assetId_fkey\"",
        )
    val quantity = integer("quantity").default(1)
    val bookingModelRequestId = text("bookingModelRequestId").nullable()
    val assetKitId = text("assetKitId").nullable()
    val sourceKitId = text("sourceKitId").nullable()
    val checkedOutAt = pgTimestamptz("checkedOutAt", 3).nullable()
    val checkedOutById = text("checkedOutById").nullable()
    val checkedInAt = pgTimestamptz("checkedInAt", 3).nullable()
    val checkedInById = text("checkedInById").nullable()
    val checkedOutQuantity = integer("checkedOutQuantity").default(0)
    val sourceLocationId = text("sourceLocationId").nullable()
    override val primaryKey = PrimaryKey(id, name = "\"BookingAsset_pkey\"")

    init {
        index("\"BookingAsset_assetId_idx\"", false, assetId)
        index("\"BookingAsset_assetKitId_idx\"", false, assetKitId)
        index("\"BookingAsset_sourceKitId_idx\"", false, sourceKitId)
        index("\"BookingAsset_sourceLocationId_idx\"", false, sourceLocationId)
    }
}
