package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class AuditImageEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AuditImageEntity>(AuditImageTable)

    var imageUrl by AuditImageTable.imageUrl
    var thumbnailUrl by AuditImageTable.thumbnailUrl
    var description by AuditImageTable.description
    var auditSession by AuditSessionEntity referencedOn AuditImageTable.auditSessionId
    var auditSessionId by AuditImageTable.auditSessionId
    var auditAsset by AuditAssetEntity optionalReferencedOn AuditImageTable.auditAssetId
    var auditAssetId by AuditImageTable.auditAssetId
    var uploadedBy by UserEntity optionalReferencedOn AuditImageTable.uploadedById
    var uploadedById by AuditImageTable.uploadedById
    var organization by OrganizationEntity referencedOn AuditImageTable.organizationId
    var organizationId by AuditImageTable.organizationId

    /**
     * When the photo was taken, from its EXIF data (or the phone app's copy of
     * it). Null when the photo carried no usable date. EXIF can be edited, so
     * this is the best signal available, not proof.
     */
    var capturedAt by AuditImageTable.capturedAt

    /**
     * How the photo reached Shelf. Null only for photos uploaded before this
     * field existed; a phone upload with no or an unknown source is PHONE.
     */
    var captureSource by AuditImageTable.captureSource
    var createdAt by AuditImageTable.createdAt
    var updatedAt by AuditImageTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AuditImageTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
