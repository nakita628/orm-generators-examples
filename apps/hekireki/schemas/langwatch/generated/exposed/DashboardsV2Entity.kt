package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * A Dashboards v2 board, apart from the Analytics Reports boards in Dashboard. Boards
 * made before the split were copied here with their ids, so stars and insight
 * pointers still name them. No relation, per the database rule.
 */
class DashboardsV2Entity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, DashboardsV2Entity>(DashboardsV2Table)

    /** The project that owns the board; only there can it be edited. */
    var projectId by DashboardsV2Table.projectId
    var name by DashboardsV2Table.name

    /** Free text under the board's name; null until a member writes one. */
    var description by DashboardsV2Table.description

    /** The member who created it; null for boards that predate authorship. */
    var createdById by DashboardsV2Table.createdById

    /** Who sees it. Only the author changes it; a board with no author stays PROJECT. */
    var scope by DashboardsV2Table.scope

    /**
     * The owning project's organization, stamped when the board is first set to
     * ORGANIZATION; null before.
     */
    var organizationId by DashboardsV2Table.organizationId

    /**
     * Deprecated: unused since each widget query names its own sources (dashboards-v2 AC200,
     * 2026-10-28); dropped in a later migration, once no serving release reads it.
     */
    var excludedOrigins by DashboardsV2Table.excludedOrigins
    var createdAt by DashboardsV2Table.createdAt
    var updatedAt by DashboardsV2Table.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(DashboardsV2Table.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
