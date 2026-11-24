package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Per-call data kept beside a workflow (ADR 0048): one row per source, `trigger` for the Webhook
 * trigger's last payload (plus the last call's time and IP), or an HTTP request step id for its last
 * Test request response. Written on every call, so it lives here and never bumps `Workflow.updatedAt`.
 */
class WorkflowSampleEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, WorkflowSampleEntity>(WorkflowSampleTable)

    var createdAt by WorkflowSampleTable.createdAt
    var updatedAt by WorkflowSampleTable.updatedAt
    var workflowId by WorkflowSampleTable.workflowId
    var workspaceId by WorkflowSampleTable.workspaceId
    var source by WorkflowSampleTable.source1

    /** [WorkflowSampleData] */
    var sample by WorkflowSampleTable.sample
    var lastCalledAt by WorkflowSampleTable.lastCalledAt
    var lastCallerIp by WorkflowSampleTable.lastCallerIp

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(WorkflowSampleTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
