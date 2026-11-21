package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

object DashboardTable : IdTable<String>("\"Dashboard\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()

    /** The project that owns the board; only there can it be edited. */
    val projectId = text("projectId")
    val name = text("name")

    /**
     * Legacy board sort position. The Dashboards v2 app orders boards by a
     * member's favourites and never writes this; kept because the legacy
     * Analytics reports section and the REST reorder endpoint still write it,
     * and a rolling deploy's older image still reads it.
     */
    val order = integer("order").default(0)

    /**
     * Free text under the board's name; null until a member writes one.
     * Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
     */
    val description = text("description").nullable()

    /**
     * The member who created it; null for boards that predate the column.
     * Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
     */
    val createdById = text("createdById").nullable()

    /**
     * Who sees it. Only the author changes it; a board with no author stays PROJECT.
     * Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
     */
    val scope = pgEnum("scope", "\"DashboardScope\"", DashboardScope.entries, DashboardScope::dbName).default(DashboardScope.PROJECT)

    /**
     * The owning project's organization, stamped when the board is first set to
     * ORGANIZATION; null before. No relation, per the database rule.
     * Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
     */
    val organizationId = text("organizationId").nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"Dashboard_pkey\"")

    init {
        index("\"Dashboard_projectId_idx\"", false, projectId)
        index("\"Dashboard_projectId_order_idx\"", false, projectId, order)
        index("\"Dashboard_organizationId_scope_idx\"", false, organizationId, scope)
    }
}
