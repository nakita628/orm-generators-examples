package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
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
class EmbeddedFieldLabelEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, EmbeddedFieldLabelEntity>(EmbeddedFieldLabelTable)

    var createdAt by EmbeddedFieldLabelTable.createdAt
    var updatedAt by EmbeddedFieldLabelTable.updatedAt
    var workspace by WorkspaceEntity referencedOn EmbeddedFieldLabelTable.workspaceId
    var workspaceId by EmbeddedFieldLabelTable.workspaceId
    var key by EmbeddedFieldLabelTable.key
    var label by EmbeddedFieldLabelTable.label

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(EmbeddedFieldLabelTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
