package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class CustomGraphEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, CustomGraphEntity>(CustomGraphTable)

    var projectId by CustomGraphTable.projectId
    var name by CustomGraphTable.name

    /** The chart's definition, in the shape its `kind` names. */
    var graph by CustomGraphTable.graph
    var filters by CustomGraphTable.filters

    /**
     * Discriminator for the chart shapes that share this table:
     * - "builder" (default) - the chart builder's own `graph` payload
     * - "workbench_sql"     - a saved LangWatchQL workbench chart, whose
     * `graph` holds the versioned { sql, parameters, vegaLiteSpec }
     * definition (see server/analytics/saved-workbench-charts)
     * - "dashboard_srcdoc"  - frozen legacy copies of Dashboards v2 widgets, which
     * now live in DashboardWidget; nothing new writes this kind
     * Each reader filters by its own kind, so neither sees the other's rows and
     * a builder chart can never be interpreted as a workbench definition.
     */
    var kind by CustomGraphTable.kind
    var createdAt by CustomGraphTable.createdAt
    var updatedAt by CustomGraphTable.updatedAt
    var dashboardId by CustomGraphTable.dashboardId
    var gridColumn by CustomGraphTable.gridColumn
    var gridRow by CustomGraphTable.gridRow
    var colSpan by CustomGraphTable.colSpan
    var rowSpan by CustomGraphTable.rowSpan

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(CustomGraphTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
