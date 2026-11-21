package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * *
 * * A template/grouping entity for assets.
 * * Groups individual assets of the same type (e.g., "Dell Latitude 5550")
 * * and provides default field values when creating new assets from the model.
 */
object AssetModelTable : IdTable<String>("\"AssetModel\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val name = text("name")
    val description = text("description").nullable()

    /**
     * Public URL (bucket `files`) of the model's cover image. Rendered by any
     * asset of this model that has no image of its own. Public — and therefore
     * never-expiring and cacheable — like Location.imageUrl; see
     * resolveAssetImage() for the cascade.
     */
    val image = text("image").nullable()

    /**
     * Public URL of the 108px thumbnail generated alongside `image`. Used at
     * list sizes so inheriting assets never render the full-size original.
     */
    val thumbnailImage = text("thumbnailImage").nullable()
    val defaultCategoryId =
        optReference(
            "defaultCategoryId",
            CategoryTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AssetModel_defaultCategoryId_fkey\"",
        )
    val defaultValuation = double("defaultValuation").nullable()
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AssetModel_organizationId_fkey\"",
        )
    val userId =
        reference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AssetModel_userId_fkey\"",
        )
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"AssetModel_pkey\"")

    init {
        index("\"AssetModel_organizationId_name_idx\"", false, organizationId, name)
        index("\"AssetModel_userId_idx\"", false, userId)
        index("\"AssetModel_name_trgm_idx\"", false, functions = listOf(indexExpression("\"name\" gin_trgm_ops")), indexType = "gin")
    }
}
