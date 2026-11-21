package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class DashboardEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, DashboardEntity>(DashboardTable)

    /** The project that owns the board; only there can it be edited. */
    var projectId by DashboardTable.projectId
    var name by DashboardTable.name

    /**
     * Legacy board sort position. The Dashboards v2 app orders boards by a
     * member's favourites and never writes this; kept because the legacy
     * Analytics reports section and the REST reorder endpoint still write it,
     * and a rolling deploy's older image still reads it.
     */
    var order by DashboardTable.order

    /**
     * Free text under the board's name; null until a member writes one.
     * Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
     */
    var description by DashboardTable.description

    /**
     * The member who created it; null for boards that predate the column.
     * Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
     */
    var createdById by DashboardTable.createdById

    /**
     * Who sees it. Only the author changes it; a board with no author stays PROJECT.
     * Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
     */
    var scope by DashboardTable.scope

    /**
     * The owning project's organization, stamped when the board is first set to
     * ORGANIZATION; null before. No relation, per the database rule.
     * Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
     */
    var organizationId by DashboardTable.organizationId
    var createdAt by DashboardTable.createdAt
    var updatedAt by DashboardTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(DashboardTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
