package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeID
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.CompositeEntity
import org.jetbrains.exposed.v1.dao.CompositeEntityClass

/**
 * Nurturing's own trace counts from trace's span_received and trace_failed (ARCHITECTURE §9), a
 * root span being one trace, one row per project and UTC day of append time, incremented in place
 * by each batch. A redelivered batch can count a trace twice; accepted (issue 9378). Kept 30 days.
 */
class NurturingProjectTraceDayEntity(
    id: EntityID<CompositeID>,
) : CompositeEntity(id) {
    companion object : CompositeEntityClass<NurturingProjectTraceDayEntity>(NurturingProjectTraceDayTable)

    val projectId by NurturingProjectTraceDayTable.projectId

    /** UTC `YYYY-MM-DD` of the root spans' append time. */
    val day by NurturingProjectTraceDayTable.day
    var traceCount by NurturingProjectTraceDayTable.traceCount

    /** Traces trace recorded failed (lw.trace.trace_failed). */
    var errorCount by NurturingProjectTraceDayTable.errorCount

    /** Append time of the newest root span. */
    var lastTraceAt by NurturingProjectTraceDayTable.lastTraceAt

    /** Append time of the newest trace_failed. */
    var lastFailedTraceAt by NurturingProjectTraceDayTable.lastFailedTraceAt
}
