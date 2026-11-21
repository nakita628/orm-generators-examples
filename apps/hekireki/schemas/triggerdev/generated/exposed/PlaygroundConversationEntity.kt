package triggerdev.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

class PlaygroundConversationEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, PlaygroundConversationEntity>(PlaygroundConversationTable)

    /** The chat session ID used by the transport */
    var chatId by PlaygroundConversationTable.chatId

    /** User-editable conversation title (auto-generated from first message) */
    var title by PlaygroundConversationTable.title

    /** Which agent this conversation is with */
    var agentSlug by PlaygroundConversationTable.agentSlug

    /** The current active run backing this conversation (null if no run yet) */
    var runId by PlaygroundConversationTable.runId
    var run by TaskRunEntity optionalReferencedOn PlaygroundConversationTable.runId

    /** The client data JSON used for this conversation */
    var clientData by PlaygroundConversationTable.clientData

    /** Accumulated UIMessages from completed turns (for resume without stream replay) */
    var messages by PlaygroundConversationTable.messages

    /** Last SSE event ID — resume from this position to avoid replaying old turns */
    var lastEventId by PlaygroundConversationTable.lastEventId
    var project by ProjectEntity referencedOn PlaygroundConversationTable.projectId
    var projectId by PlaygroundConversationTable.projectId
    var runtimeEnvironment by RuntimeEnvironmentEntity referencedOn PlaygroundConversationTable.runtimeEnvironmentId
    var runtimeEnvironmentId by PlaygroundConversationTable.runtimeEnvironmentId

    /** The user who started this conversation */
    var userId by PlaygroundConversationTable.userId
    var createdAt by PlaygroundConversationTable.createdAt
    var updatedAt by PlaygroundConversationTable.updatedAt

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(PlaygroundConversationTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
