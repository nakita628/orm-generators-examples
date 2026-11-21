package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb

/**
 * One insight in its owner's inbox in a project, folded from its `insight` aggregate
 * (modules/insight). A read model: rebuildable by replaying the event log. Epochs in
 * milliseconds.
 */
object InsightProjectionTable : IdTable<String>("\"InsightProjection\"") {
    /** The insight id, which is also its aggregate id. */
    override val id = text("id").entityId()
    val projectId = text("projectId")
    val title = text("title")
    val body = text("body")

    /** bad | watch | good */
    val tone = text("tone")
    val topic = text("topic").nullable()
    val validDays = integer("validDays")
    val lwql = text("lwql").nullable()
    val sourceConversationId = text("sourceConversationId").nullable()
    val sourceMessageId = text("sourceMessageId").nullable()
    val filedByUserId = text("filedByUserId").nullable()
    val filedAt = double("filedAt")

    /** Set by a later run that found the insight still true. */
    val renewedAt = double("renewedAt").nullable()
    val createdAt = double("createdAt")
    val updatedAt = double("updatedAt")
    val occurredAt = double("occurredAt")
    val acceptedAt = double("acceptedAt")
    val lastEventId = text("lastEventId")
    val projectionVersion = text("projectionVersion")

    /**
     * The board and widget it came from. A pointer only, never a relation: the row reads the
     * same once either is deleted, so the names are kept as filed.
     */
    val boardId = text("boardId").nullable()
    val boardName = text("boardName").nullable()
    val widgetId = text("widgetId").nullable()
    val widgetName = text("widgetName").nullable()

    /** chat | run */
    val filedVia = text("filedVia").default("chat")

    /** The fixed window `lwql` replays over (epoch milliseconds, half-open) and its step. */
    val replayStart = double("replayStart").nullable()
    val replayEnd = double("replayEnd").nullable()
    val replayGranularitySeconds = integer("replayGranularitySeconds").nullable()

    /** The values in force when it was filed: `{ period, parameters }`. */
    val replayContext = jsonb<String>("replayContext", { it }, { it }).nullable()

    /**
     * Whose insight it is: the one person who reads it. NULL on a row folded before the
     * column existed, which is owned by whoever filed it (`filedByUserId`).
     */
    val ownerUserId = text("ownerUserId").nullable()

    /**
     * The stable key a daily run matches a repeated finding by, so it renews the insight
     * and files no twin. NULL for an insight saved from a chat and on a row folded before
     * the column existed.
     */
    val findingKey = text("findingKey").nullable()
    override val primaryKey = PrimaryKey(id, name = "\"InsightProjection_pkey\"")

    init {
        index("\"InsightProjection_projectId_filedAt_idx\"", false, projectId, filedAt)
        index("\"InsightProjection_projectId_ownerUserId_filedAt_idx\"", false, projectId, ownerUserId, filedAt)
    }
}
