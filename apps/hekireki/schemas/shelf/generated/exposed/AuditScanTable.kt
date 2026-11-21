package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

object AuditScanTable : IdTable<String>("\"AuditScan\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val auditSessionId =
        reference(
            "auditSessionId",
            AuditSessionTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AuditScan_auditSessionId_fkey\"",
        )
    val auditAssetId =
        optReference(
            "auditAssetId",
            AuditAssetTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AuditScan_auditAssetId_fkey\"",
        )
    val assetId =
        optReference(
            "assetId",
            AssetTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AuditScan_assetId_fkey\"",
        )
    val scannedById =
        optReference(
            "scannedById",
            UserTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AuditScan_scannedById_fkey\"",
        )
    val code = text("code").nullable()

    /**
     * What was scanned, recorded BY VALUE at scan time.
     *
     * `asset` and `auditAsset` are both SetNull/Cascade, so deleting an asset
     * leaves this row pointing at nothing: no name, and no way to tell whether
     * it had been expected. An audit is a historical record, and a record whose
     * entries can silently lose their meaning is not one — so the two
     * descriptive facts are snapshotted here rather than only referenced.
     *
     * Readers prefer the LIVE asset when it still exists (a rename should show
     * the current name) and fall back to these.
     *
     * Null means the snapshot could not be taken: scans whose asset had already
     * been deleted before these columns were backfilled. Nothing can recover
     * those, and they degrade to the scanned `code`, which also survives. Every
     * other row carries a snapshot, so null is never merely "old".
     */
    val assetTitle = text("assetTitle").nullable()
    val wasExpected = bool("wasExpected").nullable()
    val metadata = jsonb<String>("metadata", { it }, { it }).nullable()
    val scannedAt = pgTimestamp("scannedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"AuditScan_pkey\"")

    init {
        index("\"AuditScan_auditSessionId_scannedAt_idx\"", false, auditSessionId, scannedAt)
        index("\"AuditScan_auditAssetId_idx\"", false, auditAssetId)
        index("\"AuditScan_assetId_idx\"", false, assetId)
    }
}
