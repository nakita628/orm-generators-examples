package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeID
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.CompositeEntity
import org.jetbrains.exposed.v1.dao.CompositeEntityClass

/**
 * Nurturing's own per-project usage: a peer fold over the owners' prompt, evaluator, scenario
 * and simulation-run events (ARCHITECTURE §9), one peer fold per owner pipeline, summed per
 * organization for the per-email traits. Traces are counted in NurturingProjectTraceDayCount.
 */
class NurturingProjectUsageViewEntity(
    id: EntityID<CompositeID>,
) : CompositeEntity(id) {
    companion object : CompositeEntityClass<NurturingProjectUsageViewEntity>(NurturingProjectUsageViewTable)

    val projectId by NurturingProjectUsageViewTable.projectId

    /**
     * The owner pipeline whose peer fold wrote the row (prompt, evaluator, scenario, simulation).
     * Evaluation days are not kept here: the evaluation-days map and its replay write them to
     * NurturingProjectEvaluationDay.
     */
    val owner by NurturingProjectUsageViewTable.owner

    /** The project's usage as the fold folded it (NurturingProjectUsageState). */
    var folded by NurturingProjectUsageViewTable.folded

    /** The event ids of the last delivery folded in, so a redelivery folds once. */
    var appliedEventIds by NurturingProjectUsageViewTable.appliedEventIds

    /** A row of another version reads as undecodable and is re-folded from the log. */
    var projectionVersion by NurturingProjectUsageViewTable.projectionVersion

    /** Event-derived, so a replayed row equals the row it rebuilds. */
    var updatedAt by NurturingProjectUsageViewTable.updatedAt
}
