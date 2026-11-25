package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeID
import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.CompositeEntity
import org.jetbrains.exposed.v1.dao.CompositeEntityClass

/**
 * Each project and UTC day with a completed evaluation, from evaluation's lifecycle_completed,
 * written insert-ignore. Replaces the `evaluation` rows of NurturingProjectUsageView.
 */
class NurturingProjectEvaluationDayEntity(
    id: EntityID<CompositeID>,
) : CompositeEntity(id) {
    companion object : CompositeEntityClass<NurturingProjectEvaluationDayEntity>(NurturingProjectEvaluationDayTable)

    val projectId by NurturingProjectEvaluationDayTable.projectId

    /** UTC `YYYY-MM-DD`. */
    val day by NurturingProjectEvaluationDayTable.day
}
