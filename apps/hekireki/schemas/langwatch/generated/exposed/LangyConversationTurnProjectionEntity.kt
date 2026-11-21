package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/** Rebuildable per-turn operational render document. */
class LangyConversationTurnProjectionEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, LangyConversationTurnProjectionEntity>(LangyConversationTurnProjectionTable)

    var projectId by LangyConversationTurnProjectionTable.projectId
    var conversationId by LangyConversationTurnProjectionTable.conversationId
    var turnId by LangyConversationTurnProjectionTable.turnId

    /**
     * LangyConversationTurnStatus. Deliberately TEXT, not a Postgres enum: this
     * is a PROJECTION — one writer (the fold), rebuildable from the event log,
     * and nothing queries or indexes on it. An enum bought no selectivity here,
     * only typo protection, and charged a migration for every new status that
     * Postgres then could not remove. The one definition lives in
     * `LANGY_CONVERSATION_TURN_STATUS`; the repository validates against it.
     */
    var status by LangyConversationTurnProjectionTable.status

    /** LangyMessagePart[] — the user question that opened the turn. */
    var questionParts by LangyConversationTurnProjectionTable.questionParts

    /** LangyMessagePart[] — the agent's whole final answer. */
    var answerParts by LangyConversationTurnProjectionTable.answerParts

    /** LangyTurnToolCall[] — tool-call lifecycle audit in initiation order. */
    var toolCalls by LangyConversationTurnProjectionTable.toolCalls

    /** LangyPlanItemData[] | null — null means the turn never kept a plan. */
    var plan by LangyConversationTurnProjectionTable.plan
    var error by LangyConversationTurnProjectionTable.error
    var startedAt by LangyConversationTurnProjectionTable.startedAt
    var endedAt by LangyConversationTurnProjectionTable.endedAt
    var createdAt by LangyConversationTurnProjectionTable.createdAt
    var updatedAt by LangyConversationTurnProjectionTable.updatedAt
    var occurredAt by LangyConversationTurnProjectionTable.occurredAt
    var acceptedAt by LangyConversationTurnProjectionTable.acceptedAt
    var lastEventId by LangyConversationTurnProjectionTable.lastEventId
    var projectionVersion by LangyConversationTurnProjectionTable.projectionVersion
}
