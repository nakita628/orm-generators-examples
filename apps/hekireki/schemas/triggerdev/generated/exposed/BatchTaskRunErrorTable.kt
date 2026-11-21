package triggerdev.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/** Track individual run creation failures in batch processing (Run Engine v2) */
object BatchTaskRunErrorTable : IdTable<String>("\"BatchTaskRunError\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val batchTaskRunId =
        reference(
            "batchTaskRunId",
            BatchTaskRunTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"BatchTaskRunError_batchTaskRunId_fkey\"",
        )

    /** Which item in the batch (0-based index) */
    val index = integer("index")

    /** The task identifier that was being triggered */
    val taskIdentifier = text("taskIdentifier")

    /** The payload that failed (JSON, may be truncated) */
    val payload = text("payload").nullable()

    /** The options that were used */
    val options1 = jsonb<String>("options", { it }, { it }).nullable()

    /** Error message */
    val error = text("error")

    /** Error code if available */
    val errorCode = text("errorCode").nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"BatchTaskRunError_pkey\"")

    init {
        index("\"BatchTaskRunError_batchTaskRunId_idx\"", false, batchTaskRunId)
        uniqueIndex("\"BatchTaskRunError_batchTaskRunId_index_key\"", batchTaskRunId, index)
    }
}
