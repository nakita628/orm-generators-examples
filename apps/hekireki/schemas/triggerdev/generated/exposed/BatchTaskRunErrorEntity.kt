package triggerdev.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/** Track individual run creation failures in batch processing (Run Engine v2) */
class BatchTaskRunErrorEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, BatchTaskRunErrorEntity>(BatchTaskRunErrorTable)

    var batchTaskRun by BatchTaskRunEntity referencedOn BatchTaskRunErrorTable.batchTaskRunId
    var batchTaskRunId by BatchTaskRunErrorTable.batchTaskRunId

    /** Which item in the batch (0-based index) */
    var index by BatchTaskRunErrorTable.index

    /** The task identifier that was being triggered */
    var taskIdentifier by BatchTaskRunErrorTable.taskIdentifier

    /** The payload that failed (JSON, may be truncated) */
    var payload by BatchTaskRunErrorTable.payload

    /** The options that were used */
    var options by BatchTaskRunErrorTable.options1

    /** Error message */
    var error by BatchTaskRunErrorTable.error

    /** Error code if available */
    var errorCode by BatchTaskRunErrorTable.errorCode
    var createdAt by BatchTaskRunErrorTable.createdAt
}
