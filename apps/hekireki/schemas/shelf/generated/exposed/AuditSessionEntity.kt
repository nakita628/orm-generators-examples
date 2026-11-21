package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class AuditSessionEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AuditSessionEntity>(AuditSessionTable)

    var name by AuditSessionTable.name
    var description by AuditSessionTable.description
    var targetId by AuditSessionTable.targetId
    var status by AuditSessionTable.status
    var scopeMeta by AuditSessionTable.scopeMeta
    var expectedAssetCount by AuditSessionTable.expectedAssetCount
    var foundAssetCount by AuditSessionTable.foundAssetCount
    var missingAssetCount by AuditSessionTable.missingAssetCount
    var unexpectedAssetCount by AuditSessionTable.unexpectedAssetCount
    var startedAt by AuditSessionTable.startedAt
    var dueDate by AuditSessionTable.dueDate
    var completedAt by AuditSessionTable.completedAt
    var cancelledAt by AuditSessionTable.cancelledAt
    var activeSchedulerReference by AuditSessionTable.activeSchedulerReference
    var createdBy by UserEntity referencedOn AuditSessionTable.createdById
    var createdById by AuditSessionTable.createdById
    var organization by OrganizationEntity referencedOn AuditSessionTable.organizationId
    var organizationId by AuditSessionTable.organizationId
    val assignments by AuditAssignmentEntity referrersOn AuditAssignmentTable.auditSessionId
    val assets by AuditAssetEntity referrersOn AuditAssetTable.auditSessionId
    val scans by AuditScanEntity referrersOn AuditScanTable.auditSessionId
    val notes by AuditNoteEntity referrersOn AuditNoteTable.auditSessionId
    val images by AuditImageEntity referrersOn AuditImageTable.auditSessionId
    var createdAt by AuditSessionTable.createdAt
    var updatedAt by AuditSessionTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AuditSessionTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
