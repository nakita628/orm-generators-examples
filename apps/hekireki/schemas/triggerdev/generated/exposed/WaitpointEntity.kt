package triggerdev.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * A Waitpoint blocks a run from continuing until it's completed
 * If there's a waitpoint blocking a run, it shouldn't be in the queue
 */
class WaitpointEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, WaitpointEntity>(WaitpointTable)

    var friendlyId by WaitpointTable.friendlyId
    var type by WaitpointTable.type
    var status by WaitpointTable.status
    var completedAt by WaitpointTable.completedAt

    /** If it's an Event type waitpoint, this is the event. It can also be provided for the DATETIME type */
    var idempotencyKey by WaitpointTable.idempotencyKey

    /** If this is true then we can show it in the dashboard/return it from the SDK */
    var userProvidedIdempotencyKey by WaitpointTable.userProvidedIdempotencyKey

    /** If there's a user provided idempotency key, this is the time it expires at */
    var idempotencyKeyExpiresAt by WaitpointTable.idempotencyKeyExpiresAt

    /**
     * If an idempotencyKey is no longer active, we store it here and generate a new one for the idempotencyKey field.
     * Clearing an idempotencyKey is useful for debounce or cancelling child runs.
     * This is a workaround because Prisma doesn't support partial indexes.
     */
    var inactiveIdempotencyKey by WaitpointTable.inactiveIdempotencyKey

    /** If it's a RUN type waitpoint, this is the associated run */
    var completedByTaskRunId by WaitpointTable.completedByTaskRunId
    var completedByTaskRun by TaskRunEntity optionalReferencedOn WaitpointTable.completedByTaskRunId

    /**
     * If it's a DATETIME type waitpoint, this is the date.
     * If it's a MANUAL waitpoint, this can be set as the `timeout`.
     */
    var completedAfter by WaitpointTable.completedAfter

    /** If it's a BATCH type waitpoint, this is the associated batch */
    var completedByBatchId by WaitpointTable.completedByBatchId
    var completedByBatch by BatchTaskRunEntity optionalReferencedOn WaitpointTable.completedByBatchId

    /** The runs this waitpoint is blocking */
    val blockingTaskRuns by TaskRunWaitpointEntity referrersOn TaskRunWaitpointTable.waitpointId

    /** All runs that have ever been blocked by this waitpoint, used for display purposes */
    var connectedRuns by TaskRunEntity via WaitpointRunConnectionsTable

    /** When a waitpoint is complete */
    var completedExecutionSnapshots by TaskRunExecutionSnapshotEntity via CompletedWaitpointsTable

    /** When completed, an output can be stored here */
    var output by WaitpointTable.output
    var outputType by WaitpointTable.outputType
    var outputIsError by WaitpointTable.outputIsError
    var project by ProjectEntity referencedOn WaitpointTable.projectId
    var projectId by WaitpointTable.projectId
    var environment by RuntimeEnvironmentEntity referencedOn WaitpointTable.environmentId
    var environmentId by WaitpointTable.environmentId
    var createdAt by WaitpointTable.createdAt
    var updatedAt by WaitpointTable.updatedAt

    /**
     * Denormized column that holds the raw tags
     * Denormalized column that holds the raw tags
     */
    var tags by WaitpointTable.tags

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(WaitpointTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
