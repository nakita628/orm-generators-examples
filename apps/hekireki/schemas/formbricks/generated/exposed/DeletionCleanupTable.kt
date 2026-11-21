package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.TextColumnType
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Work left over after a delete commits (ENG-3612, ENG-3614: deleted means deleted): Hub records and
 * storage files that can't be removed inside the database transaction. A row is written in the same
 * transaction as the delete, so a delete that commits always gets its cleanup and one that rolls back
 * queues nothing. A worker drains it with retries, never giving up, and deletes each row once done.
 */
object DeletionCleanupTable : IdTable<String>("\"DeletionCleanup\"") {
    override val id = text("id").clientDefault { CUID.randomCUID2().toString() }.entityId()
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updated_at", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val kind = pgEnum("kind", "\"DeletionCleanupKind\"", DeletionCleanupKind.entries, DeletionCleanupKind::dbName)

    /** Scoping and logs. Not foreign keys: the rows they name are usually gone already. */
    val organizationId = text("organizationId")
    val workspaceId = text("workspaceId")
    val surveyId = text("surveyId")

    /** Hub kinds: the organisation's feedback directories, read inside the delete's transaction. */
    val tenantIds = array("tenantIds", TextColumnType()).nullable().databaseDefault("'{}'") { emptyList() }

    /** `hubResponses`: the deleted responses whose records go. `hubResponseOutcome`: exactly one live response. */
    val responseIds = array("responseIds", TextColumnType()).nullable().databaseDefault("'{}'") { emptyList() }

    /** `storageFiles`: the object keys or URLs to delete. */
    val fileKeys = array("fileKeys", TextColumnType()).nullable().databaseDefault("'{}'") { emptyList() }
    val attempts = integer("attempts").default(0)

    /** Not drained before this, so a late pipeline write can't slip in behind a survey's cleanup. */
    val nextAttemptAt =
        pgTimestamp("nextAttemptAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }

    /** The last failure: a short code and message, never record or file content. */
    val lastError = text("lastError").nullable()
    override val primaryKey = PrimaryKey(id, name = "\"DeletionCleanup_pkey\"")

    init {
        index("\"DeletionCleanup_nextAttemptAt_idx\"", false, nextAttemptAt)
        index("\"DeletionCleanup_organizationId_created_at_idx\"", false, organizationId, createdAt)
    }
}
