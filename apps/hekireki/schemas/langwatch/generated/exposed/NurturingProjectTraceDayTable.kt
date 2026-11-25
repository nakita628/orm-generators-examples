package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeIdTable

/**
 * retired: unread since v3.28.0. Replaced by NurturingProjectTraceDayCount, which the trace-day
 * peer folds write and a replay rebuilds; an operator step may drop it once old writers are gone.
 */
object NurturingProjectTraceDayTable : CompositeIdTable("\"NurturingProjectTraceDay\"") {
    val projectId = text("projectId").entityId()

    /** UTC `YYYY-MM-DD` of the root spans' append time. */
    val day = text("day").entityId()
    val traceCount = integer("traceCount").default(0)

    /** Traces trace recorded failed (lw.trace.trace_failed). */
    val errorCount = integer("errorCount").default(0)

    /** Append time of the newest root span. */
    val lastTraceAt = pgTimestamp("lastTraceAt", 3).nullable()

    /** Append time of the newest trace_failed. */
    val lastFailedTraceAt = pgTimestamp("lastFailedTraceAt", 3).nullable()
    override val primaryKey = PrimaryKey(projectId, day, name = "\"NurturingProjectTraceDay_pkey\"")
}
