package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
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
object AssetKitTable : IdTable<String>("\"AssetKit\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val assetId =
        reference(
            "assetId",
            AssetTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AssetKit_assetId_fkey\"",
        )
    val kitId =
        reference(
            "kitId",
            KitTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AssetKit_kitId_fkey\"",
        )
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AssetKit_organizationId_fkey\"",
        )
    val quantity = integer("quantity").default(1)
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"AssetKit_pkey\"")

    init {
        index("\"AssetKit_kitId_idx\"", false, kitId)
        index("\"AssetKit_organizationId_idx\"", false, organizationId)
        uniqueIndex("\"AssetKit_assetId_kitId_key\"", assetId, kitId)
    }
}
