package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb

/** Rebuildable per-turn operational render document. */
object LangyConversationTurnProjectionTable : IdTable<String>("\"LangyConversationTurnProjection\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val projectId = text("projectId")
    val conversationId = text("conversationId")
    val turnId = text("turnId")

    /**
     * LangyConversationTurnStatus. Deliberately TEXT, not a Postgres enum: this
     * is a PROJECTION — one writer (the fold), rebuildable from the event log,
     * and nothing queries or indexes on it. An enum bought no selectivity here,
     * only typo protection, and charged a migration for every new status that
     * Postgres then could not remove. The one definition lives in
     * `LANGY_CONVERSATION_TURN_STATUS`; the repository validates against it.
     */
    val status = text("status")

    /** LangyMessagePart[] — the user question that opened the turn. */
    val questionParts = jsonb<String>("questionParts", { it }, { it })

    /** LangyMessagePart[] — the agent's whole final answer. */
    val answerParts = jsonb<String>("answerParts", { it }, { it })

    /** LangyTurnToolCall[] — tool-call lifecycle audit in initiation order. */
    val toolCalls = jsonb<String>("toolCalls", { it }, { it })

    /** LangyPlanItemData[] | null — null means the turn never kept a plan. */
    val plan = jsonb<String>("plan", { it }, { it }).nullable()
    val error = text("error").nullable()
    val startedAt = double("startedAt").nullable()
    val endedAt = double("endedAt").nullable()
    val createdAt = double("createdAt")
    val updatedAt = double("updatedAt")
    val occurredAt = double("occurredAt")
    val acceptedAt = double("acceptedAt")
    val lastEventId = text("lastEventId")
    val projectionVersion = text("projectionVersion")
    override val primaryKey = PrimaryKey(id, name = "\"LangyConversationTurnProjection_pkey\"")

    init {
        index("LangyConversationTurnProjection_projectId_conversationId_up_idx", false, projectId, conversationId, updatedAt)
        uniqueIndex("LangyConversationTurnProjection_projectId_conversationId_tu_key", projectId, conversationId, turnId)
    }
}
