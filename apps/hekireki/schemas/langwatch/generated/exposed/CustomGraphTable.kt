package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

object CustomGraphTable : IdTable<String>("\"CustomGraph\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val projectId = text("projectId")
    val name = text("name")

    /** The chart's definition, in the shape its `kind` names. */
    val graph = jsonb<String>("graph", { it }, { it })
    val filters = jsonb<String>("filters", { it }, { it }).nullable()

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
    val kind = text("kind").default("builder")
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val dashboardId = text("dashboardId").nullable()
    val gridColumn = integer("gridColumn").default(0)
    val gridRow = integer("gridRow").default(0)
    val colSpan = integer("colSpan").default(1)
    val rowSpan = integer("rowSpan").default(1)
    override val primaryKey = PrimaryKey(id, name = "\"CustomGraph_pkey\"")

    init {
        index("\"CustomGraph_projectId_idx\"", false, projectId)
        index("\"CustomGraph_dashboardId_idx\"", false, dashboardId)
        index("\"CustomGraph_projectId_kind_idx\"", false, projectId, kind)
    }
}
