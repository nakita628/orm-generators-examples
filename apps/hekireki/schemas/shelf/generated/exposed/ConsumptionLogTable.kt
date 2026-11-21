package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * *
 * * Audit trail for all quantity changes on quantity-tracked assets.
 * * Every event records who performed it, who received/returned items,
 * * which booking it relates to, and the category of change.
 * * Quantity is always positive — direction is determined by category.
 */
object ConsumptionLogTable : IdTable<String>("\"ConsumptionLog\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val assetId =
        reference(
            "assetId",
            AssetTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"ConsumptionLog_assetId_fkey\"",
        )
    val category = pgEnum("category", "\"ConsumptionCategory\"", ConsumptionCategory.entries, ConsumptionCategory::dbName)
    val quantity = integer("quantity")
    val note = text("note").nullable()
    val userId =
        reference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.RESTRICT,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"ConsumptionLog_userId_fkey\"",
        )
    val bookingId =
        optReference(
            "bookingId",
            BookingTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"ConsumptionLog_bookingId_fkey\"",
        )
    val custodianId =
        optReference(
            "custodianId",
            TeamMemberTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"ConsumptionLog_custodianId_fkey\"",
        )
    val bookingAssetId = text("bookingAssetId").nullable()
    val locationId =
        optReference(
            "locationId",
            LocationTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"ConsumptionLog_locationId_fkey\"",
        )
    val stockChange = integer("stockChange").nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"ConsumptionLog_pkey\"")

    init {
        index("\"ConsumptionLog_assetId_createdAt_idx\"", false, assetId, createdAt)
        index("\"ConsumptionLog_userId_idx\"", false, userId)
        index("\"ConsumptionLog_bookingId_idx\"", false, bookingId)
        index("\"ConsumptionLog_custodianId_idx\"", false, custodianId)
        index("\"ConsumptionLog_bookingAssetId_idx\"", false, bookingAssetId)
        index("\"ConsumptionLog_locationId_idx\"", false, locationId)
    }
}
