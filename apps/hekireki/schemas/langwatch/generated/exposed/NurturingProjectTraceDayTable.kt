package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeIdTable

/**
 * Nurturing's own trace counts from trace's span_received and trace_failed (ARCHITECTURE §9), a
 * root span being one trace, one row per project and UTC day of append time, incremented in place
 * by each batch. A redelivered batch can count a trace twice; accepted (issue 9378). Kept 30 days.
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
