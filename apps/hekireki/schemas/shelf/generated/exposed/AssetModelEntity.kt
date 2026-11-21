package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * *
 * * A template/grouping entity for assets.
 * * Groups individual assets of the same type (e.g., "Dell Latitude 5550")
 * * and provides default field values when creating new assets from the model.
 */
class AssetModelEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AssetModelEntity>(AssetModelTable)

    var name by AssetModelTable.name
    var description by AssetModelTable.description

    /**
     * Public URL (bucket `files`) of the model's cover image. Rendered by any
     * asset of this model that has no image of its own. Public — and therefore
     * never-expiring and cacheable — like Location.imageUrl; see
     * resolveAssetImage() for the cascade.
     */
    var image by AssetModelTable.image

    /**
     * Public URL of the 108px thumbnail generated alongside `image`. Used at
     * list sizes so inheriting assets never render the full-size original.
     */
    var thumbnailImage by AssetModelTable.thumbnailImage
    var defaultCategoryId by AssetModelTable.defaultCategoryId
    var defaultCategory by CategoryEntity optionalReferencedOn AssetModelTable.defaultCategoryId
    var defaultValuation by AssetModelTable.defaultValuation

    /**
     * Default custom field values. Like the category and valuation defaults,
     * each one fills the matching field only when it is empty, at the moment
     * an asset joins this model. Nothing stays linked afterwards.
     */
    val customFieldDefaults by AssetModelCustomFieldDefaultEntity referrersOn AssetModelCustomFieldDefaultTable.assetModelId
    val assets by AssetEntity optionalReferrersOn AssetTable.assetModelId
    val modelRequests by BookingModelRequestEntity referrersOn BookingModelRequestTable.assetModelId
    var organization by OrganizationEntity referencedOn AssetModelTable.organizationId
    var organizationId by AssetModelTable.organizationId
    var createdBy by UserEntity referencedOn AssetModelTable.userId
    var userId by AssetModelTable.userId
    var createdAt by AssetModelTable.createdAt
    var updatedAt by AssetModelTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AssetModelTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
