package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * A Dashboards v2 board, apart from the Analytics Reports boards in Dashboard. Boards
 * made before the split were copied here with their ids, so stars and insight
 * pointers still name them. No relation, per the database rule.
 */
object DashboardsV2Table : IdTable<String>("\"DashboardsV2\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()

    /** The project that owns the board; only there can it be edited. */
    val projectId = text("projectId")
    val name = text("name")

    /** Free text under the board's name; null until a member writes one. */
    val description = text("description").nullable()

    /** The member who created it; null for boards that predate authorship. */
    val createdById = text("createdById").nullable()

    /** Who sees it. Only the author changes it; a board with no author stays PROJECT. */
    val scope = pgEnum("scope", "\"DashboardScope\"", DashboardScope.entries, DashboardScope::dbName).default(DashboardScope.PROJECT)

    /**
     * The owning project's organization, stamped when the board is first set to
     * ORGANIZATION; null before.
     */
    val organizationId = text("organizationId").nullable()

    /**
     * Deprecated: unused since each widget query names its own sources (dashboards-v2 AC200,
     * 2026-10-28); dropped in a later migration, once no serving release reads it.
     */
    val excludedOrigins = jsonb<String>("excludedOrigins", { it }, { it }).nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"DashboardsV2_pkey\"")

    init {
        index("\"DashboardsV2_projectId_idx\"", false, projectId)
        index("\"DashboardsV2_organizationId_scope_idx\"", false, organizationId, scope)
    }
}
