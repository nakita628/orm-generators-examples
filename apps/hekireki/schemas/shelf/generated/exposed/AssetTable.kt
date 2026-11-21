package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object AssetTable : IdTable<String>("\"Asset\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val title = text("title")
    val description = text("description").nullable()
    val mainImage = text("mainImage").nullable()
    val thumbnailImage = text("thumbnailImage").nullable()
    val mainImageExpiration = pgTimestamp("mainImageExpiration", 3).nullable()
    val status = pgEnum("status", "\"AssetStatus\"", AssetStatus.entries, AssetStatus::dbName).default(AssetStatus.AVAILABLE)
    val valuation = double("value").nullable()
    val availableToBook = bool("availableToBook").default(true)
    val sequentialId = text("sequentialId").nullable()
    val archivedAt = pgTimestamptz("archivedAt", 3).nullable()
    val type = pgEnum("type", "\"AssetType\"", AssetType.entries, AssetType::dbName).default(AssetType.INDIVIDUAL)
    val assetModelId =
        optReference(
            "assetModelId",
            AssetModelTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Asset_assetModelId_fkey\"",
        )
    val quantity = integer("quantity").nullable()
    val minQuantity = integer("minQuantity").nullable()
    val consumptionType = pgEnum("consumptionType", "\"ConsumptionType\"", ConsumptionType.entries, ConsumptionType::dbName).nullable()
    val unitOfMeasure = text("unitOfMeasure").nullable()
    val lowStockNotifiedAt = pgTimestamp("lowStockNotifiedAt", 3).nullable()
    val stockLedgerStartedAt = pgTimestamp("stockLedgerStartedAt", 3).nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val userId =
        reference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Asset_userId_fkey\"",
        )
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Asset_organizationId_fkey\"",
        )
    val categoryId =
        optReference(
            "categoryId",
            CategoryTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Asset_categoryId_fkey\"",
        )
    val preferredBarcodeId =
        optReference(
            "preferredBarcodeId",
            BarcodeTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Asset_preferredBarcodeId_fkey\"",
        ).uniqueIndex("\"Asset_preferredBarcodeId_key\"")
    override val primaryKey = PrimaryKey(id, name = "\"Asset_pkey\"")

    init {
        index(
            "\"Asset_title_description_idx\"",
            false,
            functions = listOf(indexExpression("\"title\" gin_trgm_ops"), indexExpression("\"description\" gin_trgm_ops")),
            indexType = "gin",
        )
        index(
            "\"Asset_sequentialId_trgm_idx\"",
            false,
            functions = listOf(indexExpression("\"sequentialId\" gin_trgm_ops")),
            indexType = "gin",
        )
        index("\"Asset_organizationId_compound_idx\"", false, organizationId, title, status, availableToBook)
        index("\"Asset_status_organizationId_idx\"", false, status, organizationId)
        index("\"Asset_organizationId_archivedAt_idx\"", false, organizationId, archivedAt)
        index("\"Asset_createdAt_organizationId_idx\"", false, createdAt, organizationId)
        index(
            "\"Asset_organizationId_createdAt_id_idx\"",
            false,
            functions = listOf(indexExpression("\"organizationId\""), indexExpression("\"createdAt\" DESC"), indexExpression("\"id\"")),
        )
        index("\"Asset_valuation_organizationId_idx\"", false, valuation, organizationId)
        index("\"Asset_categoryId_organizationId_idx\"", false, categoryId, organizationId)
        index("\"Asset_assetModelId_organizationId_idx\"", false, assetModelId, organizationId)
        index("\"Asset_sequentialId_idx\"", false, sequentialId)
        index("\"Asset_userId_idx\"", false, userId)
        uniqueIndex("\"Asset_organizationId_sequentialId_key\"", organizationId, sequentialId)
    }
}
