package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class AuditAssetEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AuditAssetEntity>(AuditAssetTable)

    var auditSession by AuditSessionEntity referencedOn AuditAssetTable.auditSessionId
    var auditSessionId by AuditAssetTable.auditSessionId
    var asset by AssetEntity referencedOn AuditAssetTable.assetId
    var assetId by AuditAssetTable.assetId
    var expected by AuditAssetTable.expected
    var status by AuditAssetTable.status
    var scannedBy by UserEntity optionalReferencedOn AuditAssetTable.scannedById
    var scannedById by AuditAssetTable.scannedById
    var scannedAt by AuditAssetTable.scannedAt
    var metadata by AuditAssetTable.metadata

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
    var expectedQuantity by AuditAssetTable.expectedQuantity

    /**
     * Units the auditor counted. Null until a quantity-tracked asset is
     * counted, and always null for INDIVIDUAL assets. A FOUND row with a count
     * below `expectedQuantity` is a shortfall: the status stays FOUND (the pool
     * was located), the gap is this number.
     *
     * Both columns are non-negative, enforced by CHECK constraints in raw SQL
     * (migration `20261104120000_add_audit_asset_quantities`).
     */
    var foundQuantity by AuditAssetTable.foundQuantity
    var createdAt by AuditAssetTable.createdAt
    var updatedAt by AuditAssetTable.updatedAt
    val scans by AuditScanEntity optionalReferrersOn AuditScanTable.auditAssetId
    val images by AuditImageEntity optionalReferrersOn AuditImageTable.auditAssetId
    val notes by AuditNoteEntity optionalReferrersOn AuditNoteTable.auditAssetId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AuditAssetTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
