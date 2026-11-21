package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

class AuditScanEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AuditScanEntity>(AuditScanTable)

    var auditSession by AuditSessionEntity referencedOn AuditScanTable.auditSessionId
    var auditSessionId by AuditScanTable.auditSessionId
    var auditAsset by AuditAssetEntity optionalReferencedOn AuditScanTable.auditAssetId
    var auditAssetId by AuditScanTable.auditAssetId
    var asset by AssetEntity optionalReferencedOn AuditScanTable.assetId
    var assetId by AuditScanTable.assetId
    var scannedBy by UserEntity optionalReferencedOn AuditScanTable.scannedById
    var scannedById by AuditScanTable.scannedById
    var code by AuditScanTable.code

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
    var assetTitle by AuditScanTable.assetTitle
    var wasExpected by AuditScanTable.wasExpected
    var metadata by AuditScanTable.metadata
    var scannedAt by AuditScanTable.scannedAt
    var createdAt by AuditScanTable.createdAt
}
