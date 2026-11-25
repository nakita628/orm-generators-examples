package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * retired: unread since v3.28.0. The trace-day folds count a trace with no placement copy, so no
 * retry clock reads or writes this table; an operator step may drop it once old writers are gone.
 */
object NurturingUncopiedProjectTable : IdTable<String>("\"NurturingUncopiedProject\"") {
    override val id = text("projectId").entityId()
    val firstSeenAt = pgTimestamp("firstSeenAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"NurturingUncopiedProject_pkey\"")
}
