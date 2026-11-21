package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object ScanTable : IdTable<String>("\"Scan\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val latitude = text("latitude").nullable()
    val longitude = text("longitude").nullable()
    val userAgent = text("userAgent").nullable()
    val userId =
        optReference(
            "userId",
            UserTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Scan_userId_fkey\"",
        )
    val qrId =
        optReference(
            "qrId",
            QrTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Scan_qrId_fkey\"",
        )
    val rawQrId = text("rawQrId").nullable()
    val codeType = pgEnum("codeType", "\"ScanCodeType\"", ScanCodeType.entries, ScanCodeType::dbName).default(ScanCodeType.QR)
    val code = text("code").nullable()
    val barcodeId =
        optReference(
            "barcodeId",
            BarcodeTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Scan_barcodeId_fkey\"",
        )
    val assetId =
        optReference(
            "assetId",
            AssetTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Scan_assetId_fkey\"",
        )
    val kitId =
        optReference(
            "kitId",
            KitTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Scan_kitId_fkey\"",
        )
    val organizationId =
        reference(
            "organizationId",
            OrganizationTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Scan_organizationId_fkey\"",
        )
    val source1 = pgEnum("source", "\"ScanSource\"", ScanSource.entries, ScanSource::dbName).nullable()
    val manuallyGenerated = bool("manuallyGenerated").default(false)
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"Scan_pkey\"")

    init {
        index("\"Scan_qrId_idx\"", false, qrId)
        index("\"Scan_userId_idx\"", false, userId)
        index("\"Scan_rawQrId_createdAt_idx\"", false, rawQrId, createdAt)
        index("\"Scan_assetId_createdAt_idx\"", false, assetId, createdAt)
        index("\"Scan_kitId_createdAt_idx\"", false, kitId, createdAt)
        index("\"Scan_barcodeId_idx\"", false, barcodeId)
        index("\"Scan_organizationId_idx\"", false, organizationId)
    }
}
