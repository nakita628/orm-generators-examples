package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * *
 * * Audit trail for all quantity changes on quantity-tracked assets.
 * * Every event records who performed it, who received/returned items,
 * * which booking it relates to, and the category of change.
 * * Quantity is always positive — direction is determined by category.
 */
class ConsumptionLogEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, ConsumptionLogEntity>(ConsumptionLogTable)

    var asset by AssetEntity referencedOn ConsumptionLogTable.assetId
    var assetId by ConsumptionLogTable.assetId
    var category by ConsumptionLogTable.category
    var quantity by ConsumptionLogTable.quantity
    var note by ConsumptionLogTable.note
    var performedBy by UserEntity referencedOn ConsumptionLogTable.userId
    var userId by ConsumptionLogTable.userId
    var booking by BookingEntity optionalReferencedOn ConsumptionLogTable.bookingId
    var bookingId by ConsumptionLogTable.bookingId
    var custodian by TeamMemberEntity optionalReferencedOn ConsumptionLogTable.custodianId
    var custodianId by ConsumptionLogTable.custodianId
    var bookingAssetId by ConsumptionLogTable.bookingAssetId
    var location by LocationEntity optionalReferencedOn ConsumptionLogTable.locationId
    var locationId by ConsumptionLogTable.locationId
    var stockChange by ConsumptionLogTable.stockChange
    var createdAt by ConsumptionLogTable.createdAt
}
