package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Represents a chart widget on a dashboard.
 * Widgets are positioned using a grid layout system.
 *
 * @property id - Unique identifier for the widget
 * @property dashboard - The dashboard this widget belongs to
 * @property chart - The chart displayed in this widget (chart name is used for display)
 * @property layout - Grid layout configuration (x, y, width, height)
 * @property order - Display order within the dashboard
 */
class DashboardWidgetEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, DashboardWidgetEntity>(DashboardWidgetTable)

    var createdAt by DashboardWidgetTable.createdAt
    var updatedAt by DashboardWidgetTable.updatedAt
    var dashboard by DashboardEntity referencedOn DashboardWidgetTable.dashboardId
    var dashboardId by DashboardWidgetTable.dashboardId
    var chart by ChartEntity referencedOn DashboardWidgetTable.chartId
    var chartId by DashboardWidgetTable.chartId

    /** [WidgetLayout] - Grid layout: { x, y, w, h } */
    var layout by DashboardWidgetTable.layout
    var order by DashboardWidgetTable.order

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(DashboardWidgetTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
