package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * A project nurturing saw traces of with no placement copy, and when it first did: every worker
 * retries its trace batches only within one window from then, across restarts. Deleted when the
 * copy lands.
 */
object NurturingUncopiedProjectTable : IdTable<String>("\"NurturingUncopiedProject\"") {
    override val id = text("projectId").entityId()
    val firstSeenAt = pgTimestamp("firstSeenAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"NurturingUncopiedProject_pkey\"")
}
