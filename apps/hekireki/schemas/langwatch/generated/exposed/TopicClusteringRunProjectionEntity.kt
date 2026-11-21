package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * Per-project topic clustering run status (ADR-051): the rebuildable
 * operational read model behind the settings page. One row per project;
 * folded from run_completed/run_failed events, never written directly.
 */
class TopicClusteringRunProjectionEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, TopicClusteringRunProjectionEntity>(TopicClusteringRunProjectionTable)

    var projectId by TopicClusteringRunProjectionTable.projectId
    var lastRequestedAt by TopicClusteringRunProjectionTable.lastRequestedAt
    var lastRequestTrigger by TopicClusteringRunProjectionTable.lastRequestTrigger
    var lastRunAt by TopicClusteringRunProjectionTable.lastRunAt

    /** completed | skipped | failed */
    var lastRunOutcome by TopicClusteringRunProjectionTable.lastRunOutcome
    var lastRunMode by TopicClusteringRunProjectionTable.lastRunMode
    var lastRunSkippedReason by TopicClusteringRunProjectionTable.lastRunSkippedReason
    var lastRunError by TopicClusteringRunProjectionTable.lastRunError
    var lastRunErrorCode by TopicClusteringRunProjectionTable.lastRunErrorCode

    /** True when the customer can resolve the failure (credentials, quota). */
    var lastRunErrorUserActionable by TopicClusteringRunProjectionTable.lastRunErrorUserActionable
    var lastRunTracesProcessed by TopicClusteringRunProjectionTable.lastRunTracesProcessed
    var lastRunTopicsCount by TopicClusteringRunProjectionTable.lastRunTopicsCount
    var lastRunSubtopicsCount by TopicClusteringRunProjectionTable.lastRunSubtopicsCount
    var lastRunPages by TopicClusteringRunProjectionTable.lastRunPages
    var inProgressRunId by TopicClusteringRunProjectionTable.inProgressRunId
    var inProgressTraces by TopicClusteringRunProjectionTable.inProgressTraces
    var inProgressPages by TopicClusteringRunProjectionTable.inProgressPages

    /**
     * Business time the in-progress run opened; bounds how long the read
     * model may report it as running (mirrors the process's stale-run guard).
     */
    var inProgressStartedAt by TopicClusteringRunProjectionTable.inProgressStartedAt
    var createdAt by TopicClusteringRunProjectionTable.createdAt
    var updatedAt by TopicClusteringRunProjectionTable.updatedAt

    /** Business time of the latest applied event. */
    var occurredAt by TopicClusteringRunProjectionTable.occurredAt

    /** Event-log acceptance time. Together with lastEventId this is the cursor. */
    var acceptedAt by TopicClusteringRunProjectionTable.acceptedAt
    var lastEventId by TopicClusteringRunProjectionTable.lastEventId
    var projectionVersion by TopicClusteringRunProjectionTable.projectionVersion
}
