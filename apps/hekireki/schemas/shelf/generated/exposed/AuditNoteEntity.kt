package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class AuditNoteEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AuditNoteEntity>(AuditNoteTable)

    var content by AuditNoteTable.content
    var type by AuditNoteTable.type
    var createdAt by AuditNoteTable.createdAt
    var updatedAt by AuditNoteTable.updatedAt
    var user by UserEntity optionalReferencedOn AuditNoteTable.userId
    var userId by AuditNoteTable.userId
    var auditSession by AuditSessionEntity referencedOn AuditNoteTable.auditSessionId
    var auditSessionId by AuditNoteTable.auditSessionId
    var auditAsset by AuditAssetEntity optionalReferencedOn AuditNoteTable.auditAssetId
    var auditAssetId by AuditNoteTable.auditAssetId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AuditNoteTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
