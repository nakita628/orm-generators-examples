package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable

/**
 * Rebuildable operational state folded from Langy's canonical event stream.
 * The source of truth remains the ClickHouse event log.
 */
object LangyConversationProjectionTable : IdTable<String>("\"LangyConversationProjection\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val projectId = text("projectId")
    val conversationId = text("conversationId")

    /** Owner; first-writer-wins from the first message's userId. */
    val userId = text("userId")
    val title = text("title").nullable()

    /** derived | auto | user — governs auto-title precedence (user is sticky). */
    val titleSource =
        pgEnum("titleSource", "\"LangyProjectionTitleSource\"", LangyProjectionTitleSource.entries, LangyProjectionTitleSource::dbName)

    /**
     * Who started the conversation: `interactive` for the person in the panel, `run` for a
     * module acting for them with nobody at the keyboard. A row folded before the column
     * existed reads as interactive.
     */
    val origin = text("origin").default("interactive")
    val status = text("status")
    val isShared = bool("isShared").default(false)
    val sharedAt = double("sharedAt").nullable()
    val sharedById = text("sharedById").nullable()
    val messageCount = integer("messageCount").default(0)
    val lastActivityAt = double("lastActivityAt").nullable()

    /** The turn in flight, or null when idle (busy guard reads this). */
    val currentTurnId = text("currentTurnId").nullable()
    val lastError = text("lastError").nullable()

    /** The model the latest accepted turn ran on; reopening seeds the picker from it. */
    val lastModel = text("lastModel").nullable()

    /** ADR-048 shutdown-handoff resume token (opaque, worker-authored). */
    val pendingHandoffToken = text("pendingHandoffToken").nullable()
    val pendingHandoffTurnId = text("pendingHandoffTurnId").nullable()

    /** Per-conversation HMAC key for worker stream frames; server-side only. */
    val runToken = text("runToken").nullable()
    val archivedAt = double("archivedAt").nullable()

    /** Messages CI started on a pull request since the person last wrote. */
    val ciFollowUpsInARow = integer("ciFollowUpsInARow").default(0)
    val createdAt = double("createdAt")
    val updatedAt = double("updatedAt")

    /** Business time of the latest applied event. */
    val occurredAt = double("occurredAt")

    /** Event-log acceptance time. Together with lastEventId this is the cursor. */
    val acceptedAt = double("acceptedAt")
    val lastEventId = text("lastEventId")
    val projectionVersion = text("projectionVersion")
    override val primaryKey = PrimaryKey(id, name = "\"LangyConversationProjection_pkey\"")

    init {
        index("LangyConversationProjection_projectId_userId_conversationId_idx", false, projectId, userId, conversationId)
        index("LangyConversationProjection_projectId_userId_lastActivityAt_idx", false, projectId, userId, lastActivityAt)
        uniqueIndex("\"LangyConversationProjection_projectId_conversationId_key\"", projectId, conversationId)
    }
}
