package formbricks.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * A feedback source — connects external data into Unify Feedback for record creation.
 * Naming note: "FeedbackSource" is the Unify Feedback domain term; the word "Connector"
 * is reserved for the Workflows domain (see docs/development/standards/practices/naming-conventions.mdx).
 *
 * @property id - Unique identifier for the feedback source
 * @property name - Display name for the feedback source
 * @property type - Type of feedback source (formbricks_survey, csv, …)
 * @property status - Current state of the feedback source (active, paused, error)
 * @property workspace - The workspace this feedback source belongs to
 * @property feedbackDirectory - Hub directory the feedback source writes into
 * @property formbricksMappings - Element mappings for Formbricks feedback sources
 * @property fieldMappings - Field mappings for other feedback source types
 */
class FeedbackSourceEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, FeedbackSourceEntity>(FeedbackSourceTable)

    var createdAt by FeedbackSourceTable.createdAt
    var updatedAt by FeedbackSourceTable.updatedAt
    var name by FeedbackSourceTable.name
    var type by FeedbackSourceTable.type
    var status by FeedbackSourceTable.status

    /**
     * Defaults to completedOnly so a new source agrees with the live ingestion path, which only
     * fires on responseFinished. Existing rows take the default; nothing is re-imported.
     */
    var importMode by FeedbackSourceTable.importMode
    var elementScope by FeedbackSourceTable.elementScope
    var workspaceId by FeedbackSourceTable.workspaceId
    var workspace by WorkspaceEntity referencedOn FeedbackSourceTable.workspaceId
    var feedbackDirectoryId by FeedbackSourceTable.feedbackDirectoryId
    var feedbackDirectory by FeedbackDirectoryEntity referencedOn FeedbackSourceTable.feedbackDirectoryId
    var lastSyncAt by FeedbackSourceTable.lastSyncAt
    var createdBy by FeedbackSourceTable.createdBy
    var creator by UserEntity optionalReferencedOn FeedbackSourceTable.createdBy
    var feedbackDirectoryWorkspace by FeedbackDirectoryWorkspaceEntity referencedOn FeedbackSourceTable

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(FeedbackSourceTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
