package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * Rebuildable operational state folded from Langy's canonical event stream.
 * The source of truth remains the ClickHouse event log.
 */
class LangyConversationProjectionEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, LangyConversationProjectionEntity>(LangyConversationProjectionTable)

    var projectId by LangyConversationProjectionTable.projectId
    var conversationId by LangyConversationProjectionTable.conversationId

    /** Owner; first-writer-wins from the first message's userId. */
    var userId by LangyConversationProjectionTable.userId
    var title by LangyConversationProjectionTable.title

    /** derived | auto | user — governs auto-title precedence (user is sticky). */
    var titleSource by LangyConversationProjectionTable.titleSource

    /**
     * Who started the conversation: `interactive` for the person in the panel, `run` for a
     * module acting for them with nobody at the keyboard. A row folded before the column
     * existed reads as interactive.
     */
    var origin by LangyConversationProjectionTable.origin
    var status by LangyConversationProjectionTable.status
    var isShared by LangyConversationProjectionTable.isShared
    var sharedAt by LangyConversationProjectionTable.sharedAt
    var sharedById by LangyConversationProjectionTable.sharedById
    var messageCount by LangyConversationProjectionTable.messageCount
    var lastActivityAt by LangyConversationProjectionTable.lastActivityAt

    /** The turn in flight, or null when idle (busy guard reads this). */
    var currentTurnId by LangyConversationProjectionTable.currentTurnId
    var lastError by LangyConversationProjectionTable.lastError

    /** The model the latest accepted turn ran on; reopening seeds the picker from it. */
    var lastModel by LangyConversationProjectionTable.lastModel

    /** ADR-048 shutdown-handoff resume token (opaque, worker-authored). */
    var pendingHandoffToken by LangyConversationProjectionTable.pendingHandoffToken
    var pendingHandoffTurnId by LangyConversationProjectionTable.pendingHandoffTurnId

    /** Per-conversation HMAC key for worker stream frames; server-side only. */
    var runToken by LangyConversationProjectionTable.runToken
    var archivedAt by LangyConversationProjectionTable.archivedAt
    var createdAt by LangyConversationProjectionTable.createdAt
    var updatedAt by LangyConversationProjectionTable.updatedAt

    /** Business time of the latest applied event. */
    var occurredAt by LangyConversationProjectionTable.occurredAt

    /** Event-log acceptance time. Together with lastEventId this is the cursor. */
    var acceptedAt by LangyConversationProjectionTable.acceptedAt
    var lastEventId by LangyConversationProjectionTable.lastEventId
    var projectionVersion by LangyConversationProjectionTable.projectionVersion
}
