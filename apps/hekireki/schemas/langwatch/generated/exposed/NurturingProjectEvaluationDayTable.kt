package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.CompositeIdTable

/**
 * Each project and UTC day with a completed evaluation, from evaluation's lifecycle_completed,
 * written insert-ignore. Replaces the `evaluation` rows of NurturingProjectUsageView.
 */
object NurturingProjectEvaluationDayTable : CompositeIdTable("\"NurturingProjectEvaluationDay\"") {
    val projectId = text("projectId").entityId()

    /** UTC `YYYY-MM-DD`. */
    val day = text("day").entityId()
    override val primaryKey = PrimaryKey(projectId, day, name = "\"NurturingProjectEvaluationDay_pkey\"")
}
