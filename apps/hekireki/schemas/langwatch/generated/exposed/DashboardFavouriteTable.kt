package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * A member's star, in their own order within a project: a board or a From
 * LangWatch template. Personal: one member's stars never show to another. No
 * foreign key (per the database rule): the dashboard service removes a board's
 * favourites when it deletes the board, so a star never outlives its dashboard.
 */
object DashboardFavouriteTable : IdTable<String>("\"DashboardFavourite\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val userId = text("userId")

    /**
     * A DashboardsV2 board's id (copied boards kept their Dashboard id); null when
     * the star is a template.
     */
    val dashboardId = text("dashboardId").nullable()

    /** The starred From LangWatch template id; null when the star is a board. */
    val templateId = text("templateId").nullable()
    val projectId = text("projectId")

    /** The member's drag order among their own stars in the project. */
    val position = integer("position")
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    override val primaryKey = PrimaryKey(id, name = "\"DashboardFavourite_pkey\"")

    init {
        index("\"DashboardFavourite_userId_projectId_position_idx\"", false, userId, projectId, position)
        index("\"DashboardFavourite_dashboardId_idx\"", false, dashboardId)
        uniqueIndex("\"DashboardFavourite_userId_dashboardId_key\"", userId, dashboardId)
        uniqueIndex("\"DashboardFavourite_userId_projectId_templateId_key\"", userId, projectId, templateId)
    }
}
