package triggerdev.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Historical record of every run a Session has owned. Append-only —
 * rows are inserted on each `ensureRunForSession` claim, never updated.
 * Lets us reconstruct the run timeline of a chat for debugging /
 * dashboard surfaces. The relation cascades on Session delete (tied to
 * the session lifecycle) but `runId` is a plain string column with no
 * FK to TaskRun so run pruning is independent.
 */
object SessionRunTable : IdTable<String>("\"SessionRun\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val sessionId =
        reference(
            "sessionId",
            SessionTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"SessionRun_sessionId_fkey\"",
        )

    /** TaskRun.id (no FK — runs may be archived independently of session history) */
    val runId = text("runId").uniqueIndex("\"SessionRun_runId_key\"")

    /**
     * One of: "initial" | "continuation" | "upgrade" | "manual".
     * Plain string for forward-compat with future trigger reasons.
     */
    val reason = text("reason")
    val triggeredAt = pgTimestamp("triggeredAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"SessionRun_pkey\"")

    init {
        index("\"SessionRun_sessionId_idx\"", false, sessionId)
        index("\"SessionRun_sessionId_triggeredAt_id_idx\"", false, sessionId, triggeredAt, id)
    }
}
