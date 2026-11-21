package triggerdev.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Used by the RunEngine during TaskRun execution
 * It has the required information to transactionally progress a run through states,
 * and prevent side effects like heartbeats failing a run that has progressed.
 * It is optimised for performance and is designed to be cleared at some point,
 * so there are no cascading relationships to other models.
 */
class TaskRunExecutionSnapshotEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, TaskRunExecutionSnapshotEntity>(TaskRunExecutionSnapshotTable)

    /** This should always be 2+ (V1 didn't use the run engine or snapshots) */
    var engine by TaskRunExecutionSnapshotTable.engine

    /** The execution status */
    var executionStatus by TaskRunExecutionSnapshotTable.executionStatus

    /** For debugging */
    var description by TaskRunExecutionSnapshotTable.description

    /** We store invalid snapshots as a record of the run state when we tried to move */
    var isValid by TaskRunExecutionSnapshotTable.isValid
    var error by TaskRunExecutionSnapshotTable.error

    /** The previous snapshot ID */
    var previousSnapshotId by TaskRunExecutionSnapshotTable.previousSnapshotId

    /** Run */
    var runId by TaskRunExecutionSnapshotTable.runId
    var run by TaskRunEntity referencedOn TaskRunExecutionSnapshotTable.runId
    var runStatus by TaskRunExecutionSnapshotTable.runStatus
    var batchId by TaskRunExecutionSnapshotTable.batchId

    /** This is the current run attempt number. Users can define how many attempts they want for a run. */
    var attemptNumber by TaskRunExecutionSnapshotTable.attemptNumber

    /** Environment */
    var environmentId by TaskRunExecutionSnapshotTable.environmentId
    var environmentType by TaskRunExecutionSnapshotTable.environmentType
    var projectId by TaskRunExecutionSnapshotTable.projectId
    var organizationId by TaskRunExecutionSnapshotTable.organizationId

    /** Waitpoints that have been completed for this execution */
    var completedWaitpoints by WaitpointEntity via CompletedWaitpointsTable

    /** An array of waitpoint IDs in the correct order, used for batches */
    var completedWaitpointOrder by TaskRunExecutionSnapshotTable.completedWaitpointOrder

    /** Checkpoint */
    var checkpointId by TaskRunExecutionSnapshotTable.checkpointId
    var checkpoint by TaskRunCheckpointEntity optionalReferencedOn TaskRunExecutionSnapshotTable.checkpointId

    /** Worker */
    var workerId by TaskRunExecutionSnapshotTable.workerId
    var runnerId by TaskRunExecutionSnapshotTable.runnerId
    var createdAt by TaskRunExecutionSnapshotTable.createdAt
    var updatedAt by TaskRunExecutionSnapshotTable.updatedAt
    var lastHeartbeatAt by TaskRunExecutionSnapshotTable.lastHeartbeatAt

    /** Metadata used by various systems in the run engine */
    var metadata by TaskRunExecutionSnapshotTable.metadata

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(TaskRunExecutionSnapshotTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
