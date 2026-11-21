package shelf.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

object AuditAssetTable : IdTable<String>("\"AuditAsset\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val auditSessionId =
        reference(
            "auditSessionId",
            AuditSessionTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AuditAsset_auditSessionId_fkey\"",
        )
    val assetId =
        reference(
            "assetId",
            AssetTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AuditAsset_assetId_fkey\"",
        )
    val expected = bool("expected").default(true)
    val status =
        pgEnum("status", "\"AuditAssetStatus\"", AuditAssetStatus.entries, AuditAssetStatus::dbName).default(AuditAssetStatus.PENDING)
    val scannedById =
        optReference(
            "scannedById",
            UserTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"AuditAsset_scannedById_fkey\"",
        )
    val scannedAt = pgTimestamp("scannedAt", 3).nullable()
    val metadata = jsonb<String>("metadata", { it }, { it }).nullable()

    /**
     * Units the audit expects to find, snapshotted when the row is created: the
     * units in the audit's scope (placed at its locations, held in its kits, or
     * held by its person; at most `Asset.quantity`), or `Asset.quantity` when
     * none are in scope or the audit has no scope. Null for INDIVIDUAL assets,
     * which count as one.
     *
     * A snapshot, not a live read: stock that moves during the audit must not
     * move the target the auditor is counting against.
     */
    val expectedQuantity = integer("expectedQuantity").nullable()

    /**
     * Units the auditor counted. Null until a quantity-tracked asset is
     * counted, and always null for INDIVIDUAL assets. A FOUND row with a count
     * below `expectedQuantity` is a shortfall: the status stays FOUND (the pool
     * was located), the gap is this number.
     *
     * Both columns are non-negative, enforced by CHECK constraints in raw SQL
     * (migration `20261104120000_add_audit_asset_quantities`).
     */
    val foundQuantity = integer("foundQuantity").nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"AuditAsset_pkey\"")

    init {
        index("\"AuditAsset_status_idx\"", false, status)
        index("\"AuditAsset_scannedById_idx\"", false, scannedById)
        index("\"AuditAsset_assetId_idx\"", false, assetId)
        uniqueIndex("\"AuditAsset_auditSessionId_assetId_key\"", auditSessionId, assetId)
    }
}
