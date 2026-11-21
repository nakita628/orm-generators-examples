package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object AssetLocationTable : IdTable<String>("\"AssetLocation\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val assetId =
        reference(
            "assetId",
            AssetTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AssetLocation_assetId_fkey\"",
        )
    val locationId =
        reference(
            "locationId",
            LocationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AssetLocation_locationId_fkey\"",
        )
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AssetLocation_organizationId_fkey\"",
        )
    val quantity = integer("quantity").default(1)
    val assetKitId =
        optReference(
            "assetKitId",
            AssetKitTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AssetLocation_assetKitId_fkey\"",
        )
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"AssetLocation_pkey\"")

    init {
        index("\"AssetLocation_locationId_idx\"", false, locationId)
        index("\"AssetLocation_organizationId_idx\"", false, organizationId)
        index("\"AssetLocation_assetId_idx\"", false, assetId)
        index("\"AssetLocation_assetKitId_idx\"", false, assetKitId)
    }
}
