package langwatch.models

import org.jetbrains.exposed.v1.core.TextColumnType
import org.jetbrains.exposed.v1.core.dao.id.CompositeIdTable
import org.jetbrains.exposed.v1.json.jsonb

/**
 * Nurturing's own per-project usage: a peer fold over the owners' prompt, evaluator, scenario
 * and simulation-run events (ARCHITECTURE §9), one peer fold per owner pipeline, summed per
 * organization for the per-email traits. Traces are counted in NurturingProjectTraceDay.
 */
object NurturingProjectUsageViewTable : CompositeIdTable("\"NurturingProjectUsageView\"") {
    val projectId = text("projectId").entityId()

    /**
     * The owner pipeline whose peer fold wrote the row (prompt, evaluator, scenario, simulation),
     * or `evaluation`: the project's evaluation days, written by a subscriber, at most once a day.
     */
    val owner = text("owner").entityId()

    /** The project's usage as the fold folded it (NurturingProjectUsageState). */
    val folded = jsonb<String>("folded", { it }, { it })

    /** The event ids of the last delivery folded in, so a redelivery folds once. */
    val appliedEventIds = array("appliedEventIds", TextColumnType()).nullable().databaseDefault("'{}'") { emptyList() }

    /** A row of another version reads as undecodable and is re-folded from the log. */
    val projectionVersion = text("projectionVersion")

    /** Event-derived, so a replayed row equals the row it rebuilds. */
    val updatedAt = pgTimestamp("updatedAt", 3)
    override val primaryKey = PrimaryKey(projectId, owner, name = "\"NurturingProjectUsageView_pkey\"")
}
