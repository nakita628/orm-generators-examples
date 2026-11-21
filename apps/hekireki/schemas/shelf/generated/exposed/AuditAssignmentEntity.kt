package shelf.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class AuditAssignmentEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AuditAssignmentEntity>(AuditAssignmentTable)

    var auditSession by AuditSessionEntity referencedOn AuditAssignmentTable.auditSessionId
    var auditSessionId by AuditAssignmentTable.auditSessionId
    var user by UserEntity referencedOn AuditAssignmentTable.userId
    var userId by AuditAssignmentTable.userId
    var role by AuditAssignmentTable.role
    var createdAt by AuditAssignmentTable.createdAt
    var updatedAt by AuditAssignmentTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AuditAssignmentTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
