package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Encrypted value of an HTTP request secret header (ENG-7168, ADR 0043). The definition and every version
 * store only `{name, secret: true, secretId}`; Replace overwrites this row in place, so the enabled
 * version uses the new value at once. Rows nothing references are deleted on Enable; all are deleted
 * with the workflow. A value is only sent to the origin it was set for (ADR 0051).
 */
class WorkflowSecretEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, WorkflowSecretEntity>(WorkflowSecretTable)

    var createdAt by WorkflowSecretTable.createdAt
    var updatedAt by WorkflowSecretTable.updatedAt
    var workflowId by WorkflowSecretTable.workflowId
    var workspaceId by WorkflowSecretTable.workspaceId

    /** `symmetricEncrypt` ciphertext (AES-256-GCM). Never logged, never returned. */
    var value by WorkflowSecretTable.value

    /**
     * Origin (scheme, host, port) of the step's URL when the value was set or replaced; the value is only
     * sent there (ADR 0051). Null for a value set before host binding: held to the step's saved URL
     * until it is set again.
     */
    var boundOrigin by WorkflowSecretTable.boundOrigin

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(WorkflowSecretTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
