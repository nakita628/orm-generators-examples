package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Represents a chart/visualization that can be used in multiple dashboards.
 * Charts are reusable components that query analytics data.
 *
 * @property id - Unique identifier for the chart
 * @property name - Display name of the chart
 * @property type - Type of visualization (bar, line, pie, etc.)
 * @property workspace - The workspace this chart belongs to
 * @property query - Cube.js query configuration (JSON)
 * @property config - Chart-specific configuration (colors, labels, etc.)
 * @property createdBy - User who created the chart
 * @property dashboards - Dashboards that use this chart
 */
class ChartEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, ChartEntity>(ChartTable)

    var createdAt by ChartTable.createdAt
    var updatedAt by ChartTable.updatedAt
    var name by ChartTable.name
    var type by ChartTable.type
    var workspace by WorkspaceEntity referencedOn ChartTable.workspaceId
    var workspaceId by ChartTable.workspaceId

    /** [ChartQuery] - Cube.js query configuration */
    var query by ChartTable.query

    /** [ChartConfig] - Visualization configuration (colors, labels, formatting) */
    var config by ChartTable.config
    var creator by UserEntity optionalReferencedOn ChartTable.createdBy
    var createdBy by ChartTable.createdBy
    var feedbackDirectory by FeedbackDirectoryEntity referencedOn ChartTable.feedbackDirectoryId
    var feedbackDirectoryId by ChartTable.feedbackDirectoryId
    val widgets by DashboardWidgetEntity referrersOn DashboardWidgetTable.chartId

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(ChartTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
