package triggerdev.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.TextColumnType
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Used by the RunEngine during TaskRun execution
 * It has the required information to transactionally progress a run through states,
 * and prevent side effects like heartbeats failing a run that has progressed.
 * It is optimised for performance and is designed to be cleared at some point,
 * so there are no cascading relationships to other models.
 */
object TaskRunExecutionSnapshotTable : IdTable<String>("\"TaskRunExecutionSnapshot\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()

    /** This should always be 2+ (V1 didn't use the run engine or snapshots) */
    val engine = pgEnum("engine", "\"RunEngineVersion\"", RunEngineVersion.entries, RunEngineVersion::dbName).default(RunEngineVersion.V2)

    /** The execution status */
    val executionStatus =
        pgEnum("executionStatus", "\"TaskRunExecutionStatus\"", TaskRunExecutionStatus.entries, TaskRunExecutionStatus::dbName)

    /** For debugging */
    val description = text("description")

    /** We store invalid snapshots as a record of the run state when we tried to move */
    val isValid = bool("isValid").default(true)
    val error = text("error").nullable()

    /** The previous snapshot ID */
    val previousSnapshotId = text("previousSnapshotId").nullable()

    /** Run */
    val runId =
        reference(
            "runId",
            TaskRunTable,
            onDelete = ReferenceOption.RESTRICT,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"TaskRunExecutionSnapshot_runId_fkey\"",
        )
    val runStatus = pgEnum("runStatus", "\"TaskRunStatus\"", TaskRunStatus.entries, TaskRunStatus::dbName)
    val batchId = text("batchId").nullable()

    /** This is the current run attempt number. Users can define how many attempts they want for a run. */
    val attemptNumber = integer("attemptNumber").nullable()

    /** Environment */
    val environmentId = text("environmentId")
    val environmentType =
        pgEnum("environmentType", "\"RuntimeEnvironmentType\"", RuntimeEnvironmentType.entries, RuntimeEnvironmentType::dbName)
    val projectId = text("projectId")
    val organizationId = text("organizationId")

    /** An array of waitpoint IDs in the correct order, used for batches */
    val completedWaitpointOrder = array("completedWaitpointOrder", TextColumnType()).nullable()

    /** Checkpoint */
    val checkpointId =
        optReference(
            "checkpointId",
            TaskRunCheckpointTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"TaskRunExecutionSnapshot_checkpointId_fkey\"",
        )

    /** Worker */
    val workerId = text("workerId").nullable()
    val runnerId = text("runnerId").nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val lastHeartbeatAt = pgTimestamp("lastHeartbeatAt", 3).nullable()

    /** Metadata used by various systems in the run engine */
    val metadata = jsonb<String>("metadata", { it }, { it }).nullable()
    override val primaryKey = PrimaryKey(id, name = "\"TaskRunExecutionSnapshot_pkey\"")

    init {
        index(
            "\"TaskRunExecutionSnapshot_runId_isValid_createdAt_idx\"",
            false,
            functions = listOf(indexExpression("\"runId\""), indexExpression("\"isValid\""), indexExpression("\"createdAt\" DESC")),
        )
    }
}
