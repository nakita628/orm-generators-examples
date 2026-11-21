package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * A Dashboards v2 widget: a srcdoc definition placed on a DashboardsV2 board, or
 * unplaced. Widgets made before the split were copied from CustomGraph with their
 * ids. No relation, per the database rule.
 */
object DashboardWidgetTable : IdTable<String>("\"DashboardWidget\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val projectId = text("projectId")

    /** The DashboardsV2 board it sits on; null while unplaced. */
    val dashboardId = text("dashboardId").nullable()
    val name = text("name")

    /** The widget's srcdoc definition. */
    val definition = jsonb<String>("definition", { it }, { it })
    val gridColumn = integer("gridColumn").default(0)
    val gridRow = integer("gridRow").default(0)
    val colSpan = integer("colSpan").default(1)
    val rowSpan = integer("rowSpan").default(1)
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"DashboardWidget_pkey\"")

    init {
        index("\"DashboardWidget_projectId_dashboardId_idx\"", false, projectId, dashboardId)
    }
}
