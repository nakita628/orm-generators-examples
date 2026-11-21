package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object BarcodeTable : IdTable<String>("\"Barcode\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val value = text("value")
    val type = pgEnum("type", "\"BarcodeType\"", BarcodeType.entries, BarcodeType::dbName).default(BarcodeType.Code128)
    val assetId =
        optReference(
            "assetId",
            AssetTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Barcode_assetId_fkey\"",
        )
    val kitId =
        optReference(
            "kitId",
            KitTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Barcode_kitId_fkey\"",
        )
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Barcode_organizationId_fkey\"",
        )
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"Barcode_pkey\"")

    init {
        index("\"Barcode_organizationId_value_idx\"", false, organizationId, value)
        index("\"Barcode_assetId_idx\"", false, assetId)
        index("\"Barcode_kitId_idx\"", false, kitId)
        index("\"Barcode_organizationId_idx\"", false, organizationId)
        index("\"Barcode_value_trgm_idx\"", false, functions = listOf(indexExpression("\"value\" gin_trgm_ops")), indexType = "gin")
        uniqueIndex("\"Barcode_organizationId_value_key\"", organizationId, value)
    }
}
