package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class CategoryEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, CategoryEntity>(CategoryTable)

    var name by CategoryTable.name
    var description by CategoryTable.description
    var color by CategoryTable.color
    var createdAt by CategoryTable.createdAt
    var updatedAt by CategoryTable.updatedAt
    val assets by AssetEntity optionalReferrersOn AssetTable.categoryId
    val kits by KitEntity optionalReferrersOn KitTable.categoryId
    var user by UserEntity referencedOn CategoryTable.userId
    var userId by CategoryTable.userId
    var customFields by CustomFieldEntity via CategoryToCustomFieldTable
    val assetModelDefaults by AssetModelEntity optionalReferrersOn AssetModelTable.defaultCategoryId
    var organization by OrganizationEntity referencedOn CategoryTable.organizationId
    var organizationId by CategoryTable.organizationId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(CategoryTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
