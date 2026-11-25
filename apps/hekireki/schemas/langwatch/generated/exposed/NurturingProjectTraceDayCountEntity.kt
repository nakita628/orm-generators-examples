package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeID
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.CompositeEntity
import org.jetbrains.exposed.v1.dao.CompositeEntityClass

/**
 * Nurturing's own trace counts (ARCHITECTURE §9), one row per project, UTC day of append time and
 * source: `spans` from trace's span_received peer fold (a customer root span is one trace),
 * `failures` from trace_failed's. Each fold owns its rows, so the two never overwrite each other;
 * readers sum a day's sources. Rebuilt by a projection replay. Kept 30 days.
 */
class NurturingProjectTraceDayCountEntity(
    id: EntityID<CompositeID>,
) : CompositeEntity(id) {
    companion object : CompositeEntityClass<NurturingProjectTraceDayCountEntity>(NurturingProjectTraceDayCountTable)

    val projectId by NurturingProjectTraceDayCountTable.projectId

    /** UTC `YYYY-MM-DD` of the facts' append time. */
    val day by NurturingProjectTraceDayCountTable.day

    /** The peer fold that wrote the row: `spans` or `failures`. */
    val source by NurturingProjectTraceDayCountTable.source1
    var traceCount by NurturingProjectTraceDayCountTable.traceCount
    var errorCount by NurturingProjectTraceDayCountTable.errorCount

    /** Append time of the newest root span. */
    var lastTraceAt by NurturingProjectTraceDayCountTable.lastTraceAt

    /** Append time of the newest trace_failed. */
    var lastFailedTraceAt by NurturingProjectTraceDayCountTable.lastFailedTraceAt

    /** The event ids of the last delivery folded in, so a redelivery folds once. */
    var appliedEventIds by NurturingProjectTraceDayCountTable.appliedEventIds

    /** A row of another version reads as undecodable and is left out of the read. */
    var projectionVersion by NurturingProjectTraceDayCountTable.projectionVersion
}
