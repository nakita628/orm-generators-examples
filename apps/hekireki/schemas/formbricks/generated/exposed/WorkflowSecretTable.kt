package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Encrypted value of an HTTP request secret header (ENG-7168, ADR 0043). The definition and every version
 * store only `{name, secret: true, secretId}`; Replace overwrites this row in place, so the enabled
 * version uses the new value at once. Rows nothing references are deleted on Enable; all are deleted
 * with the workflow. A value is only sent to the origin it was set for (ADR 0051).
 */
object WorkflowSecretTable : IdTable<String>("\"WorkflowSecret\"") {
    override val id = text("id").clientDefault { CUID.randomCUID2().toString() }.entityId()
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updated_at", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val workflowId = text("workflowId")
    val workspaceId = text("workspaceId")

    /** `symmetricEncrypt` ciphertext (AES-256-GCM). Never logged, never returned. */
    val value = text("value")

    /**
     * Origin (scheme, host, port) of the step's URL when the value was set or replaced; the value is only
     * sent there (ADR 0051). Null for a value set before host binding: held to the step's saved URL
     * until it is set again.
     */
    val boundOrigin = text("boundOrigin").nullable()
    override val primaryKey = PrimaryKey(id, name = "\"WorkflowSecret_pkey\"")

    init {
        index("\"WorkflowSecret_workflowId_workspaceId_idx\"", false, workflowId, workspaceId)
        foreignKey(
            workflowId to WorkflowTable.id,
            workspaceId to WorkflowTable.workspaceId,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            name = "\"WorkflowSecret_workflowId_workspaceId_fkey\"",
        )
    }
}
