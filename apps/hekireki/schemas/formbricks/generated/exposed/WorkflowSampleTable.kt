package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Per-call data kept beside a workflow (ADR 0048): one row per source, `trigger` for the Webhook
 * trigger's last payload (plus the last call's time and IP), or an HTTP request step id for its last
 * Test request response. Written on every call, so it lives here and never bumps `Workflow.updatedAt`.
 */
object WorkflowSampleTable : IdTable<String>("\"WorkflowSample\"") {
    override val id = text("id").clientDefault { CUID.randomCUID2().toString() }.entityId()
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updated_at", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val workflowId = text("workflowId")
    val workspaceId = text("workspaceId")
    val source1 = text("source")

    /** [WorkflowSampleData] */
    val sample = jsonb<String>("sample", { it }, { it })
    val lastCalledAt = pgTimestamp("lastCalledAt", 3).nullable()
    val lastCallerIp = text("lastCallerIp").nullable()
    override val primaryKey = PrimaryKey(id, name = "\"WorkflowSample_pkey\"")

    init {
        index("\"WorkflowSample_workflowId_workspaceId_idx\"", false, workflowId, workspaceId)
        index("\"WorkflowSample_updated_at_idx\"", false, updatedAt)
        uniqueIndex("\"WorkflowSample_workflowId_source_key\"", workflowId, source1)
        foreignKey(
            workflowId to WorkflowTable.id,
            workspaceId to WorkflowTable.workspaceId,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            name = "\"WorkflowSample_workflowId_workspaceId_fkey\"",
        )
    }
}
