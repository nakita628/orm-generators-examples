package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Work left over after a delete commits (ENG-3612, ENG-3614: deleted means deleted): Hub records and
 * storage files that can't be removed inside the database transaction. A row is written in the same
 * transaction as the delete, so a delete that commits always gets its cleanup and one that rolls back
 * queues nothing. A worker drains it with retries, never giving up, and deletes each row once done.
 */
class DeletionCleanupEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, DeletionCleanupEntity>(DeletionCleanupTable)

    var createdAt by DeletionCleanupTable.createdAt
    var updatedAt by DeletionCleanupTable.updatedAt
    var kind by DeletionCleanupTable.kind

    /** Scoping and logs. Not foreign keys: the rows they name are usually gone already. */
    var organizationId by DeletionCleanupTable.organizationId
    var workspaceId by DeletionCleanupTable.workspaceId
    var surveyId by DeletionCleanupTable.surveyId

    /** Hub kinds: the organisation's feedback directories, read inside the delete's transaction. */
    var tenantIds by DeletionCleanupTable.tenantIds

    /** `hubResponses`: the deleted responses whose records go. `hubResponseOutcome`: exactly one live response. */
    var responseIds by DeletionCleanupTable.responseIds

    /** `storageFiles`: the object keys or URLs to delete. */
    var fileKeys by DeletionCleanupTable.fileKeys
    var attempts by DeletionCleanupTable.attempts

    /** Not drained before this, so a late pipeline write can't slip in behind a survey's cleanup. */
    var nextAttemptAt by DeletionCleanupTable.nextAttemptAt

    /** The last failure: a short code and message, never record or file content. */
    var lastError by DeletionCleanupTable.lastError

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(DeletionCleanupTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
