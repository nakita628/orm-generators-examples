package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object QrTable : IdTable<String>("\"Qr\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val version = integer("version").default(0)
    val errorCorrection =
        pgEnum("errorCorrection", "\"ErrorCorrection\"", ErrorCorrection.entries, ErrorCorrection::dbName).default(ErrorCorrection.L)
    val assetId =
        optReference(
            "assetId",
            AssetTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Qr_assetId_fkey\"",
        )
    val kitId =
        optReference(
            "kitId",
            KitTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Qr_kitId_fkey\"",
        )
    val userId =
        optReference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Qr_userId_fkey\"",
        )
    val organizationId =
        optReference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Qr_organizationId_fkey\"",
        )
    val batchId =
        optReference(
            "batchId",
            PrintBatchTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Qr_batchId_fkey\"",
        )
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"Qr_pkey\"")

    init {
        index("\"Qr_assetId_idx\"", false, assetId)
        index("\"Qr_kitId_idx\"", false, kitId)
        index("\"Qr_userId_idx\"", false, userId)
        index("\"Qr_organizationId_idx\"", false, organizationId)
        index("\"Qr_batchId_idx\"", false, batchId)
        index("\"Qr_id_trgm_idx\"", false, functions = listOf(indexExpression("\"id\" gin_trgm_ops")), indexType = "gin")
    }
}
