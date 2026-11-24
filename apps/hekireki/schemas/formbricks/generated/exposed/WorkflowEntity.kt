package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class WorkflowEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, WorkflowEntity>(WorkflowTable)

    var createdAt by WorkflowTable.createdAt
    var updatedAt by WorkflowTable.updatedAt
    var name by WorkflowTable.name
    var description by WorkflowTable.description
    var status by WorkflowTable.status
    var workspace by WorkspaceEntity referencedOn WorkflowTable.workspaceId
    var workspaceId by WorkflowTable.workspaceId
    var creator by UserEntity optionalReferencedOn WorkflowTable.createdBy
    var createdBy by WorkflowTable.createdBy

    /** [WorkflowDefinition] */
    var definition by WorkflowTable.definition

    /**
     * Webhook Secret URL token (ADRs 0048, 0057): SHA-256 hash to check calls, and an encrypted copy so
     * the editor can show the URL again. Never part of the definition, its versions or its copies.
     */
    var webhookTokenHash by WorkflowTable.webhookTokenHash
    var webhookTokenEncrypted by WorkflowTable.webhookTokenEncrypted

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(WorkflowTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
