package formbricks.models

import io.github.thibaultmeyer.cuid.CUID
import org.jetbrains.exposed.v1.core.ReferenceOption
import org.jetbrains.exposed.v1.core.dao.id.IdTable
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * A workspace's display name for an imported (embedded-data) key of its Feedback Datasets, shown in
 * place of the key in chart legends, headers and the chart builder's field pickers. One per key and
 * workspace; clearing the name deletes the row, and the key shows as collected again. Built-in fields
 * are never renamed here: they keep their translated names.
 *
 * @property key - The embedded-data key exactly as collected, case included
 * @property label - The name shown for the key across the workspace
 */
object EmbeddedFieldLabelTable : IdTable<String>("\"EmbeddedFieldLabel\"") {
    override val id = text("id").clientDefault { CUID.randomCUID1().toString() }.entityId()
    val createdAt = pgTimestamp("created_at", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updated_at", 3).clientDefault { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val workspaceId =
        reference(
            "workspaceId",
            WorkspaceTable,
            onDelete = ReferenceOption.CASCADE,
            onUpdate = ReferenceOption.CASCADE,
            fkName = "\"EmbeddedFieldLabel_workspaceId_fkey\"",
        )
    val key = text("key")
    val label = text("label")
    override val primaryKey = PrimaryKey(id, name = "\"EmbeddedFieldLabel_pkey\"")

    init {
        uniqueIndex("\"EmbeddedFieldLabel_workspaceId_key_key\"", workspaceId, key)
    }
}
