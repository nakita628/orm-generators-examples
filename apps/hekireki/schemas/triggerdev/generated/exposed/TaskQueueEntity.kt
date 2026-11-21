package triggerdev.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class TaskQueueEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, TaskQueueEntity>(TaskQueueTable)

    var friendlyId by TaskQueueTable.friendlyId
    var name by TaskQueueTable.name
    var type by TaskQueueTable.type
    var version by TaskQueueTable.version
    var orderableName by TaskQueueTable.orderableName

    /**
     * How the queue's concurrency was declared: V1 = legacy concurrencyLimit semantics
     * (per key when keyed, whole queue otherwise), V2 = the explicit perKey/total shape.
     */
    var concurrencyVersion by TaskQueueTable.concurrencyVersion

    /**
     * QUEUE rows are real queues; LIMIT rows back named concurrency limits and are
     * excluded from every queue-facing surface.
     */
    var role by TaskQueueTable.role
    var project by ProjectEntity referencedOn TaskQueueTable.projectId
    var projectId by TaskQueueTable.projectId
    var runtimeEnvironment by RuntimeEnvironmentEntity referencedOn TaskQueueTable.runtimeEnvironmentId
    var runtimeEnvironmentId by TaskQueueTable.runtimeEnvironmentId

    /** Represents the current concurrency limit for the queue */
    var concurrencyLimit by TaskQueueTable.concurrencyLimit

    /** When the concurrency limit was overridden */
    var concurrencyLimitOverriddenAt by TaskQueueTable.concurrencyLimitOverriddenAt

    /** Who overrode the concurrency limit (will be null if overridden via the API) */
    var concurrencyLimitOverriddenBy by TaskQueueTable.concurrencyLimitOverriddenBy

    /** If concurrencyLimit is overridden, this is the overridden value */
    var concurrencyLimitBase by TaskQueueTable.concurrencyLimitBase

    /**
     * If the override was expressed as a percentage of the environment limit, this stores that
     * percentage (the source of truth). The absolute concurrencyLimit is materialized from it.
     * Decimal(5,2) allows fractional percentages like 12.50% (0.01–100.00).
     */
    var concurrencyLimitOverridePercent by TaskQueueTable.concurrencyLimitOverridePercent

    /**
     * Caps total concurrent runs across ALL concurrencyKey values of this queue
     * (concurrencyLimit applies per key value). Null = no total cap.
     */
    var totalConcurrencyLimit by TaskQueueTable.totalConcurrencyLimit

    /** When the total concurrency limit was overridden */
    var totalConcurrencyLimitOverriddenAt by TaskQueueTable.totalConcurrencyLimitOverriddenAt

    /** Who overrode the total concurrency limit (null when overridden via the API) */
    var totalConcurrencyLimitOverriddenBy by TaskQueueTable.totalConcurrencyLimitOverriddenBy

    /** If totalConcurrencyLimit is overridden, the declared value it reverts to on reset */
    var totalConcurrencyLimitBase by TaskQueueTable.totalConcurrencyLimitBase
    var rateLimit by TaskQueueTable.rateLimit
    var paused by TaskQueueTable.paused

    /** Hidden from the dashboard queue list. Display-only: never read on trigger/dequeue paths. */
    var archivedAt by TaskQueueTable.archivedAt
    var createdAt by TaskQueueTable.createdAt
    var updatedAt by TaskQueueTable.updatedAt
    val attempts by TaskRunAttemptEntity referrersOn TaskRunAttemptTable.queueId
    val tasks by BackgroundWorkerTaskEntity optionalReferrersOn BackgroundWorkerTaskTable.queueId
    var workers by BackgroundWorkerEntity via BackgroundWorkerToTaskQueueTable

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(TaskQueueTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
