package triggerdev.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.TextColumnType
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * A Waitpoint blocks a run from continuing until it's completed
 * If there's a waitpoint blocking a run, it shouldn't be in the queue
 */
object WaitpointTable : IdTable<String>("\"Waitpoint\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val friendlyId = text("friendlyId").uniqueIndex("\"Waitpoint_friendlyId_key\"")
    val type = pgEnum("type", "\"WaitpointType\"", WaitpointType.entries, WaitpointType::dbName)
    val status = pgEnum("status", "\"WaitpointStatus\"", WaitpointStatus.entries, WaitpointStatus::dbName).default(WaitpointStatus.PENDING)
    val completedAt = pgTimestamp("completedAt", 3).nullable()

    /** If it's an Event type waitpoint, this is the event. It can also be provided for the DATETIME type */
    val idempotencyKey = text("idempotencyKey")

    /** If this is true then we can show it in the dashboard/return it from the SDK */
    val userProvidedIdempotencyKey = bool("userProvidedIdempotencyKey")

    /** If there's a user provided idempotency key, this is the time it expires at */
    val idempotencyKeyExpiresAt = pgTimestamp("idempotencyKeyExpiresAt", 3).nullable()

    /**
     * If an idempotencyKey is no longer active, we store it here and generate a new one for the idempotencyKey field.
     * Clearing an idempotencyKey is useful for debounce or cancelling child runs.
     * This is a workaround because Prisma doesn't support partial indexes.
     */
    val inactiveIdempotencyKey = text("inactiveIdempotencyKey").nullable()

    /** If it's a RUN type waitpoint, this is the associated run */
    val completedByTaskRunId =
        optReference(
            "completedByTaskRunId",
            TaskRunTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Waitpoint_completedByTaskRunId_fkey\"",
        ).uniqueIndex("\"Waitpoint_completedByTaskRunId_key\"")

    /**
     * If it's a DATETIME type waitpoint, this is the date.
     * If it's a MANUAL waitpoint, this can be set as the `timeout`.
     */
    val completedAfter = pgTimestamp("completedAfter", 3).nullable()

    /** If it's a BATCH type waitpoint, this is the associated batch */
    val completedByBatchId =
        optReference(
            "completedByBatchId",
            BatchTaskRunTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Waitpoint_completedByBatchId_fkey\"",
        )

    /** When completed, an output can be stored here */
    val output = text("output").nullable()
    val outputType = text("outputType").default("application/json")
    val outputIsError = bool("outputIsError").default(false)
    val projectId =
        reference(
            "projectId",
            ProjectTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Waitpoint_projectId_fkey\"",
        )
    val environmentId =
        reference(
            "environmentId",
            RuntimeEnvironmentTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Waitpoint_environmentId_fkey\"",
        )
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }

    /**
     * Denormized column that holds the raw tags
     * Denormalized column that holds the raw tags
     */
    val tags = array("tags", TextColumnType()).nullable()
    override val primaryKey = PrimaryKey(id, name = "\"Waitpoint_pkey\"")

    init {
        index("\"Waitpoint_completedByBatchId_idx\"", false, completedByBatchId)
        index(
            "\"Waitpoint_environmentId_type_createdAt_idx\"",
            false,
            functions = listOf(indexExpression("\"environmentId\""), indexExpression("\"type\""), indexExpression("\"createdAt\" DESC")),
        )
        index("\"Waitpoint_environmentId_type_status_idx\"", false, environmentId, type, status)
        index(
            "\"Waitpoint_environmentId_type_id_idx\"",
            false,
            functions = listOf(indexExpression("\"environmentId\""), indexExpression("\"type\""), indexExpression("\"id\" DESC")),
        )
        uniqueIndex("\"Waitpoint_environmentId_idempotencyKey_key\"", environmentId, idempotencyKey)
    }
}
