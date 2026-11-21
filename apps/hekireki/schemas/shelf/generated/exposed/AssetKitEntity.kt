package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Pivot linking an `Asset` to the `Kit` it belongs to.
 *
 * `@@unique([assetId])` enforces "at most one kit per asset" today.
 * The structure already allows multi-row-per-asset; if multi-kit
 * allocation lands later, drop the `assetId` unique constraint and
 * add a type-aware single-row trigger + sum-within-total CONSTRAINT
 * TRIGGER (parallel to the Custody / BookingAsset pivots).
 */
class AssetKitEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AssetKitEntity>(AssetKitTable)

    var asset by AssetEntity referencedOn AssetKitTable.assetId
    var assetId by AssetKitTable.assetId
    var kit by KitEntity referencedOn AssetKitTable.kitId
    var kitId by AssetKitTable.kitId
    var organization by OrganizationEntity referencedOn AssetKitTable.organizationId
    var organizationId by AssetKitTable.organizationId
    var quantity by AssetKitTable.quantity
    val assetLocations by AssetLocationEntity optionalReferrersOn AssetLocationTable.assetKitId
    var createdAt by AssetKitTable.createdAt
    var updatedAt by AssetKitTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AssetKitTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
