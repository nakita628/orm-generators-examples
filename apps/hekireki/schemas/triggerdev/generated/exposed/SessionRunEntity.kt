package triggerdev.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * Historical record of every run a Session has owned. Append-only —
 * rows are inserted on each `ensureRunForSession` claim, never updated.
 * Lets us reconstruct the run timeline of a chat for debugging /
 * dashboard surfaces. The relation cascades on Session delete (tied to
 * the session lifecycle) but `runId` is a plain string column with no
 * FK to TaskRun so run pruning is independent.
 */
class SessionRunEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, SessionRunEntity>(SessionRunTable)

    var sessionId by SessionRunTable.sessionId

    /** TaskRun.id (no FK — runs may be archived independently of session history) */
    var runId by SessionRunTable.runId

    /**
     * One of: "initial" | "continuation" | "upgrade" | "manual".
     * Plain string for forward-compat with future trigger reasons.
     */
    var reason by SessionRunTable.reason
    var triggeredAt by SessionRunTable.triggeredAt
    var session by SessionEntity referencedOn SessionRunTable.sessionId
}
