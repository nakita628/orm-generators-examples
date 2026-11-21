package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * A member's star, in their own order within a project: a board or a From
 * LangWatch template. Personal: one member's stars never show to another. No
 * foreign key (per the database rule): the dashboard service removes a board's
 * favourites when it deletes the board, so a star never outlives its dashboard.
 */
class DashboardFavouriteEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, DashboardFavouriteEntity>(DashboardFavouriteTable)

    var userId by DashboardFavouriteTable.userId

    /**
     * A DashboardsV2 board's id (copied boards kept their Dashboard id); null when
     * the star is a template.
     */
    var dashboardId by DashboardFavouriteTable.dashboardId

    /** The starred From LangWatch template id; null when the star is a board. */
    var templateId by DashboardFavouriteTable.templateId
    var projectId by DashboardFavouriteTable.projectId

    /** The member's drag order among their own stars in the project. */
    var position by DashboardFavouriteTable.position
    var createdAt by DashboardFavouriteTable.createdAt
}
