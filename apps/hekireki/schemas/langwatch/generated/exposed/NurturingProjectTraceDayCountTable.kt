package langwatch.models

import org.jetbrains.exposed.v1.core.TextColumnType
import org.jetbrains.exposed.v1.core.dao.id.CompositeIdTable

/**
 * Nurturing's own trace counts (ARCHITECTURE §9), one row per project, UTC day of append time and
 * source: `spans` from trace's span_received peer fold (a customer root span is one trace),
 * `failures` from trace_failed's. Each fold owns its rows, so the two never overwrite each other;
 * readers sum a day's sources. Rebuilt by a projection replay. Kept 30 days.
 */
object NurturingProjectTraceDayCountTable : CompositeIdTable("\"NurturingProjectTraceDayCount\"") {
    val projectId = text("projectId").entityId()

    /** UTC `YYYY-MM-DD` of the facts' append time. */
    val day = text("day").entityId()

    /** The peer fold that wrote the row: `spans` or `failures`. */
    val source1 = text("source").entityId()
    val traceCount = integer("traceCount").default(0)
    val errorCount = integer("errorCount").default(0)

    /** Append time of the newest root span. */
    val lastTraceAt = pgTimestamp("lastTraceAt", 3).nullable()

    /** Append time of the newest trace_failed. */
    val lastFailedTraceAt = pgTimestamp("lastFailedTraceAt", 3).nullable()

    /** The event ids of the last delivery folded in, so a redelivery folds once. */
    val appliedEventIds = array("appliedEventIds", TextColumnType()).nullable().databaseDefault("'{}'") { emptyList() }

    /** A row of another version reads as undecodable and is left out of the read. */
    val projectionVersion = text("projectionVersion")
    override val primaryKey = PrimaryKey(projectId, day, source1, name = "\"NurturingProjectTraceDayCount_pkey\"")

    init {
        index("\"NurturingProjectTraceDayCount_day_idx\"", false, day)
    }
}
