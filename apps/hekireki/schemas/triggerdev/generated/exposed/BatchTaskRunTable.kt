package triggerdev.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.TextColumnType
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

object BatchTaskRunTable : IdTable<String>("\"BatchTaskRun\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val friendlyId = text("friendlyId").uniqueIndex("\"BatchTaskRun_friendlyId_key\"")
    val idempotencyKey = text("idempotencyKey").nullable()
    val idempotencyKeyExpiresAt = pgTimestamp("idempotencyKeyExpiresAt", 3).nullable()
    val status =
        pgEnum("status", "\"BatchTaskRunStatus\"", BatchTaskRunStatus.entries, BatchTaskRunStatus::dbName).default(
            BatchTaskRunStatus.PENDING,
        )
    val runtimeEnvironmentId = text("runtimeEnvironmentId")
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }

    /** Friendly IDs */
    val runIds = array("runIds", TextColumnType()).nullable().databaseDefault("'{}'") { emptyList() }
    val runCount = integer("runCount").default(0)
    val payload = text("payload").nullable()
    val payloadType = text("payloadType").default("application/json")
    val options1 = jsonb<String>("options", { it }, { it }).nullable()
    val batchVersion = text("batchVersion").default("v1")

    /** sealed is set to true once no more items can be added to the batch */
    val sealed = bool("sealed").default(false)
    val sealedAt = pgTimestamp("sealedAt", 3).nullable()

    /** this is the expected number of items in the batch */
    val expectedCount = integer("expectedCount").default(0)

    /** this is the completed number of items in the batch. once this reaches expectedCount, and the batch is sealed, the batch is considered completed */
    val completedCount = integer("completedCount").default(0)
    val completedAt = pgTimestamp("completedAt", 3).nullable()
    val resumedAt = pgTimestamp("resumedAt", 3).nullable()

    /** this is used to be able to "seal" this BatchTaskRun when all of the runs have been triggered asynchronously, and using the "parallel" processing strategy */
    val processingJobsCount = integer("processingJobsCount").default(0)
    val processingJobsExpectedCount = integer("processingJobsExpectedCount").default(0)

    /** optional token that can be used to authenticate the task run */
    val oneTimeUseToken = text("oneTimeUseToken").nullable().uniqueIndex("\"BatchTaskRun_oneTimeUseToken_key\"")

    /** When processing started (status changed to PROCESSING) */
    val processingStartedAt = pgTimestamp("processingStartedAt", 3).nullable()

    /** When processing completed (all items processed) */
    val processingCompletedAt = pgTimestamp("processingCompletedAt", 3).nullable()

    /** Count of successfully created runs */
    val successfulRunCount = integer("successfulRunCount").nullable()

    /** Count of failed run creations */
    val failedRunCount = integer("failedRunCount").nullable()
    val taskIdentifier = text("taskIdentifier").nullable()
    val checkpointEventId =
        optReference(
            "checkpointEventId",
            CheckpointRestoreEventTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"BatchTaskRun_checkpointEventId_fkey\"",
        ).uniqueIndex("\"BatchTaskRun_checkpointEventId_key\"")
    val dependentTaskAttemptId =
        optReference(
            "dependentTaskAttemptId",
            TaskRunAttemptTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"BatchTaskRun_dependentTaskAttemptId_fkey\"",
        )
    override val primaryKey = PrimaryKey(id, name = "\"BatchTaskRun_pkey\"")

    init {
        index("\"BatchTaskRun_dependentTaskAttemptId_idx\"", false, dependentTaskAttemptId)
        index(
            "\"BatchTaskRun_runtimeEnvironmentId_id_idx\"",
            false,
            functions = listOf(indexExpression("\"runtimeEnvironmentId\""), indexExpression("\"id\" DESC")),
        )
        index(
            "\"BatchTaskRun_runtimeEnvironmentId_createdAt_id_idx\"",
            false,
            functions =
                listOf(indexExpression("\"runtimeEnvironmentId\""), indexExpression("\"createdAt\" DESC"), indexExpression("\"id\" DESC")),
        )
        uniqueIndex("\"BatchTaskRun_runtimeEnvironmentId_idempotencyKey_key\"", runtimeEnvironmentId, idempotencyKey)
    }
}
