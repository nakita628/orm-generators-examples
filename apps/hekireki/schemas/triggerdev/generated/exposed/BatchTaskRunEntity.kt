package triggerdev.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class BatchTaskRunEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, BatchTaskRunEntity>(BatchTaskRunTable)

    var friendlyId by BatchTaskRunTable.friendlyId
    var idempotencyKey by BatchTaskRunTable.idempotencyKey
    var idempotencyKeyExpiresAt by BatchTaskRunTable.idempotencyKeyExpiresAt
    var status by BatchTaskRunTable.status
    var runtimeEnvironmentId by BatchTaskRunTable.runtimeEnvironmentId

    /** This only includes new runs, not idempotent runs. */
    val runs by TaskRunEntity optionalReferrersOn TaskRunTable.batchId
    var createdAt by BatchTaskRunTable.createdAt
    var updatedAt by BatchTaskRunTable.updatedAt

    /** Friendly IDs */
    var runIds by BatchTaskRunTable.runIds
    var runCount by BatchTaskRunTable.runCount
    var payload by BatchTaskRunTable.payload
    var payloadType by BatchTaskRunTable.payloadType
    var options by BatchTaskRunTable.options1
    var batchVersion by BatchTaskRunTable.batchVersion

    /** Specific run blockers, */
    val runsBlocked by TaskRunWaitpointEntity optionalReferrersOn TaskRunWaitpointTable.batchId

    /**
     * Waitpoints that are blocked by this batch.
     * When a Batch is created it blocks execution of the associated parent run (for andWait)
     */
    val waitpoints by WaitpointEntity optionalReferrersOn WaitpointTable.completedByBatchId

    /** sealed is set to true once no more items can be added to the batch */
    var sealed by BatchTaskRunTable.sealed
    var sealedAt by BatchTaskRunTable.sealedAt

    /** this is the expected number of items in the batch */
    var expectedCount by BatchTaskRunTable.expectedCount

    /** this is the completed number of items in the batch. once this reaches expectedCount, and the batch is sealed, the batch is considered completed */
    var completedCount by BatchTaskRunTable.completedCount
    var completedAt by BatchTaskRunTable.completedAt
    var resumedAt by BatchTaskRunTable.resumedAt

    /** this is used to be able to "seal" this BatchTaskRun when all of the runs have been triggered asynchronously, and using the "parallel" processing strategy */
    var processingJobsCount by BatchTaskRunTable.processingJobsCount
    var processingJobsExpectedCount by BatchTaskRunTable.processingJobsExpectedCount

    /** optional token that can be used to authenticate the task run */
    var oneTimeUseToken by BatchTaskRunTable.oneTimeUseToken

    /** When processing started (status changed to PROCESSING) */
    var processingStartedAt by BatchTaskRunTable.processingStartedAt

    /** When processing completed (all items processed) */
    var processingCompletedAt by BatchTaskRunTable.processingCompletedAt

    /** Count of successfully created runs */
    var successfulRunCount by BatchTaskRunTable.successfulRunCount

    /** Count of failed run creations */
    var failedRunCount by BatchTaskRunTable.failedRunCount

    /** Detailed failure records */
    val errors by BatchTaskRunErrorEntity referrersOn BatchTaskRunErrorTable.batchTaskRunId

    /** all the below properties are engine v1 only */
    val items by BatchTaskRunItemEntity referrersOn BatchTaskRunItemTable.batchTaskRunId
    var taskIdentifier by BatchTaskRunTable.taskIdentifier
    var checkpointEvent by CheckpointRestoreEventEntity optionalReferencedOn BatchTaskRunTable.checkpointEventId
    var checkpointEventId by BatchTaskRunTable.checkpointEventId
    var dependentTaskAttempt by TaskRunAttemptEntity optionalReferencedOn BatchTaskRunTable.dependentTaskAttemptId
    var dependentTaskAttemptId by BatchTaskRunTable.dependentTaskAttemptId
    val runDependencies by TaskRunDependencyEntity optionalReferrersOn TaskRunDependencyTable.dependentBatchRunId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(BatchTaskRunTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
