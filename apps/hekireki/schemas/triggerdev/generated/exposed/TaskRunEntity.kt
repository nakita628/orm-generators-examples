package triggerdev.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class TaskRunEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, TaskRunEntity>(TaskRunTable)

    var number by TaskRunTable.number
    var friendlyId by TaskRunTable.friendlyId
    var engine by TaskRunTable.engine
    var status by TaskRunTable.status
    var statusReason by TaskRunTable.statusReason
    var idempotencyKey by TaskRunTable.idempotencyKey
    var idempotencyKeyExpiresAt by TaskRunTable.idempotencyKeyExpiresAt

    /** Stores the user-provided key and scope: { key: string, scope: "run" | "attempt" | "global" } */
    var idempotencyKeyOptions by TaskRunTable.idempotencyKeyOptions

    /** Debounce options: { key: string, delay: string, createdAt: Date } */
    var debounce by TaskRunTable.debounce
    var taskIdentifier by TaskRunTable.taskIdentifier
    var isTest by TaskRunTable.isTest
    var payload by TaskRunTable.payload
    var payloadType by TaskRunTable.payloadType
    var context by TaskRunTable.context
    var traceContext by TaskRunTable.traceContext
    var traceId by TaskRunTable.traceId
    var spanId by TaskRunTable.spanId
    var runtimeEnvironment by RuntimeEnvironmentEntity referencedOn TaskRunTable.runtimeEnvironmentId
    var runtimeEnvironmentId by TaskRunTable.runtimeEnvironmentId
    var environmentType by TaskRunTable.environmentType
    var project by ProjectEntity referencedOn TaskRunTable.projectId
    var projectId by TaskRunTable.projectId
    var organizationId by TaskRunTable.organizationId
    var queue by TaskRunTable.queue
    var lockedQueueId by TaskRunTable.lockedQueueId

    /** The main queue that this run is part of */
    var workerQueue by TaskRunTable.workerQueue

    /** User-facing geo region, stamped at trigger; workerQueue is where it actually ran. */
    var region by TaskRunTable.region

    /** @deprecated */
    var secondaryMasterQueue by TaskRunTable.secondaryMasterQueue

    /** From engine v2+ this will be defined after a run has been dequeued (starting at 1) */
    var attemptNumber by TaskRunTable.attemptNumber
    var createdAt by TaskRunTable.createdAt
    var updatedAt by TaskRunTable.updatedAt
    val attempts by TaskRunAttemptEntity referrersOn TaskRunAttemptTable.taskRunId

    /** Denormized column that holds the raw tags */
    var runTags by TaskRunTable.runTags

    /** Denormalized version of the background worker task */
    var taskVersion by TaskRunTable.taskVersion
    var sdkVersion by TaskRunTable.sdkVersion
    var cliVersion by TaskRunTable.cliVersion
    val checkpoints by CheckpointEntity referrersOn CheckpointTable.runId

    /** startedAt marks the point at which a run is dequeued from MarQS */
    var startedAt by TaskRunTable.startedAt

    /** executedAt is set when the first attempt is about to execute */
    var executedAt by TaskRunTable.executedAt
    var completedAt by TaskRunTable.completedAt
    var machinePreset by TaskRunTable.machinePreset
    var usageDurationMs by TaskRunTable.usageDurationMs
    var costInCents by TaskRunTable.costInCents
    var baseCostInCents by TaskRunTable.baseCostInCents
    var lockedAt by TaskRunTable.lockedAt
    var lockedBy by BackgroundWorkerTaskEntity optionalReferencedOn TaskRunTable.lockedById
    var lockedById by TaskRunTable.lockedById
    var lockedToVersion by BackgroundWorkerEntity optionalReferencedOn TaskRunTable.lockedToVersionId
    var lockedToVersionId by TaskRunTable.lockedToVersionId

    /**
     * The "priority" of the run. This is just a negative offset in ms for the queue timestamp
     * E.g. a value of 60_000 would put the run into the queue 60s ago.
     */
    var priorityMs by TaskRunTable.priorityMs
    var concurrencyKey by TaskRunTable.concurrencyKey

    /**
     * Gates for this run: other queues it must also hold a concurrency slot in while
     * executing. Shape: [{ queue: string, concurrencyKey?: string }]
     */
    var gates by TaskRunTable.gates
    var delayUntil by TaskRunTable.delayUntil
    var queuedAt by TaskRunTable.queuedAt
    var ttl by TaskRunTable.ttl
    var expiredAt by TaskRunTable.expiredAt
    var maxAttempts by TaskRunTable.maxAttempts
    var lockedRetryConfig by TaskRunTable.lockedRetryConfig

    /** optional token that can be used to authenticate the task run */
    var oneTimeUseToken by TaskRunTable.oneTimeUseToken

    /** When this run is finished, the waitpoint will be marked as completed */
    val associatedWaitpoint by WaitpointEntity optionalBackReferencedOn WaitpointTable.completedByTaskRunId

    /** If there are any blocked waitpoints, the run won't be executed */
    val blockedByWaitpoints by TaskRunWaitpointEntity referrersOn TaskRunWaitpointTable.taskRunId

    /** All waitpoints that blocked this run at some point, used for display purposes */
    var connectedWaitpoints by WaitpointEntity via WaitpointRunConnectionsTable

    /** Where the logs are stored */
    var taskEventStore by TaskRunTable.taskEventStore
    var queueTimestamp by TaskRunTable.queueTimestamp
    val batchItems by BatchTaskRunItemEntity referrersOn BatchTaskRunItemTable.taskRunId
    val dependency by TaskRunDependencyEntity optionalBackReferencedOn TaskRunDependencyTable.taskRunId
    val checkpointRestoreEvent by CheckpointRestoreEventEntity referrersOn CheckpointRestoreEventTable.runId
    val executionSnapshots by TaskRunExecutionSnapshotEntity referrersOn TaskRunExecutionSnapshotTable.runId
    var scheduleInstanceId by TaskRunTable.scheduleInstanceId
    var scheduleId by TaskRunTable.scheduleId
    val sourceBulkActionItems by BulkActionItemEntity referrersOn BulkActionItemTable.sourceRunId
    val destinationBulkActionItems by BulkActionItemEntity optionalReferrersOn BulkActionItemTable.destinationRunId
    var bulkActionGroupIds by TaskRunTable.bulkActionGroupIds
    var logsDeletedAt by TaskRunTable.logsDeletedAt
    var replayedFromTaskRunFriendlyId by TaskRunTable.replayedFromTaskRunFriendlyId

    /** This represents the original task that that was triggered outside of a Trigger.dev task */
    var rootTaskRun by TaskRunEntity optionalReferencedOn TaskRunTable.rootTaskRunId
    var rootTaskRunId by TaskRunTable.rootTaskRunId

    /** The root run will have a list of all the descendant runs, children, grand children, etc. */
    val descendantRuns by TaskRunEntity optionalReferrersOn TaskRunTable.rootTaskRunId

    /** The immediate parent run of this task run */
    var parentTaskRun by TaskRunEntity optionalReferencedOn TaskRunTable.parentTaskRunId
    var parentTaskRunId by TaskRunTable.parentTaskRunId

    /** The immediate child runs of this task run */
    val childRuns by TaskRunEntity optionalReferrersOn TaskRunTable.parentTaskRunId

    /** The immediate parent attempt of this task run */
    var parentTaskRunAttempt by TaskRunAttemptEntity optionalReferencedOn TaskRunTable.parentTaskRunAttemptId
    var parentTaskRunAttemptId by TaskRunTable.parentTaskRunAttemptId

    /** The batch run that this task run is a part of */
    var batch by BatchTaskRunEntity optionalReferencedOn TaskRunTable.batchId
    var batchId by TaskRunTable.batchId

    /** whether or not the task run was created because of a triggerAndWait for batchTriggerAndWait */
    var resumeParentOnCompletion by TaskRunTable.resumeParentOnCompletion

    /** The depth of this task run in the task run hierarchy */
    var depth by TaskRunTable.depth

    /** The span ID of the "trigger" span in the parent task run */
    var parentSpanId by TaskRunTable.parentSpanId

    /** Holds the state of the run chain for deadlock detection */
    var runChainState by TaskRunTable.runChainState

    /** seed run metadata */
    var seedMetadata by TaskRunTable.seedMetadata
    var seedMetadataType by TaskRunTable.seedMetadataType

    /** Run metadata */
    var metadata by TaskRunTable.metadata
    var metadataType by TaskRunTable.metadataType
    var metadataVersion by TaskRunTable.metadataVersion

    /** Structured annotations: triggerSource, triggerAction, rootTriggerSource, rootScheduleId */
    var annotations by TaskRunTable.annotations

    /** Whether the latest attempt was a warm start. Null until first attempt starts. */
    var isWarmStart by TaskRunTable.isWarmStart

    /** Run output */
    var output by TaskRunTable.output
    var outputType by TaskRunTable.outputType

    /** Run error */
    var error by TaskRunTable.error

    /** Organization's billing plan type (cached for fallback when billing API fails) */
    var planType by TaskRunTable.planType
    var maxDurationInSeconds by TaskRunTable.maxDurationInSeconds

    /** The version of the realtime streams implementation used by the run */
    var realtimeStreamsVersion by TaskRunTable.realtimeStreamsVersion

    /** Store the stream keys that are being used by the run */
    var realtimeStreams by TaskRunTable.realtimeStreams

    /**
     * S2 basin where this run's realtime streams live. Stamped at create
     * time from `Organization.streamBasinName` so reads can resolve the
     * basin without joining org. Null when the org has no per-org basin
     * (OSS, or pre-backfill); reads fall back to the global basin.
     */
    var streamBasinName by TaskRunTable.streamBasinName
    val playgroundConversations by PlaygroundConversationEntity optionalReferrersOn PlaygroundConversationTable.runId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(TaskRunTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
