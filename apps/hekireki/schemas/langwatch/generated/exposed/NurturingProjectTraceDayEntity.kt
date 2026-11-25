package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeID
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.CompositeEntity
import org.jetbrains.exposed.v1.dao.CompositeEntityClass

/**
 * retired: unread since v3.28.0. Replaced by NurturingProjectTraceDayCount, which the trace-day
 * peer folds write and a replay rebuilds; an operator step may drop it once old writers are gone.
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
