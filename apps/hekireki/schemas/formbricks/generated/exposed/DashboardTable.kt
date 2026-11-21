package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Represents a dashboard containing multiple charts.
 * Dashboards aggregate analytics insights at the workspace level.
 *
 * @property id - Unique identifier for the dashboard
 * @property name - Display name of the dashboard
 * @property workspace - The workspace this dashboard belongs to
 * @property widgets - Charts on this dashboard
 * @property filters - The dashboard filters pinned to its top bar, in order: `{ id, field, operator, label }`,
 * at most 10 (ADR-0026). Each viewer's picks live in their own browser, not here.
 * @property createdBy - User who created the dashboard
 */
object DashboardTable : IdTable<String>("\"Dashboard\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updated_at", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val name = text("name")
    val workspaceId =
        reference(
            "workspaceId",
            WorkspaceTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Dashboard_workspaceId_fkey\"",
        )
    val createdBy =
        optReference(
            "createdBy",
            UserTable,
            onDelete = ReferenceOption.SET_NULL,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"Dashboard_createdBy_fkey\"",
        )
    val filters = jsonb<String>("filters", { it }, { it }).default("[]")
    override val primaryKey = PrimaryKey(id, name = "\"Dashboard_pkey\"")

    init {
        index("\"Dashboard_workspaceId_created_at_idx\"", false, workspaceId, createdAt)
        uniqueIndex("\"Dashboard_workspaceId_name_key\"", workspaceId, name)
    }
}
