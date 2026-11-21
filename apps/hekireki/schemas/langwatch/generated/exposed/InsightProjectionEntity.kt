package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * One insight in its owner's inbox in a project, folded from its `insight` aggregate
 * (modules/insight). A read model: rebuildable by replaying the event log. Epochs in
 * milliseconds.
 */
class InsightProjectionEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, InsightProjectionEntity>(InsightProjectionTable)

    var projectId by InsightProjectionTable.projectId
    var title by InsightProjectionTable.title
    var body by InsightProjectionTable.body

    /** bad | watch | good */
    var tone by InsightProjectionTable.tone
    var topic by InsightProjectionTable.topic
    var validDays by InsightProjectionTable.validDays
    var lwql by InsightProjectionTable.lwql
    var sourceConversationId by InsightProjectionTable.sourceConversationId
    var sourceMessageId by InsightProjectionTable.sourceMessageId
    var filedByUserId by InsightProjectionTable.filedByUserId
    var filedAt by InsightProjectionTable.filedAt

    /** Set by a later run that found the insight still true. */
    var renewedAt by InsightProjectionTable.renewedAt
    var createdAt by InsightProjectionTable.createdAt
    var updatedAt by InsightProjectionTable.updatedAt
    var occurredAt by InsightProjectionTable.occurredAt
    var acceptedAt by InsightProjectionTable.acceptedAt
    var lastEventId by InsightProjectionTable.lastEventId
    var projectionVersion by InsightProjectionTable.projectionVersion

    /**
     * The board and widget it came from. A pointer only, never a relation: the row reads the
     * same once either is deleted, so the names are kept as filed.
     */
    var boardId by InsightProjectionTable.boardId
    var boardName by InsightProjectionTable.boardName
    var widgetId by InsightProjectionTable.widgetId
    var widgetName by InsightProjectionTable.widgetName

    /** chat | run */
    var filedVia by InsightProjectionTable.filedVia

    /** The fixed window `lwql` replays over (epoch milliseconds, half-open) and its step. */
    var replayStart by InsightProjectionTable.replayStart
    var replayEnd by InsightProjectionTable.replayEnd
    var replayGranularitySeconds by InsightProjectionTable.replayGranularitySeconds

    /** The values in force when it was filed: `{ period, parameters }`. */
    var replayContext by InsightProjectionTable.replayContext

    /**
     * Whose insight it is: the one person who reads it. NULL on a row folded before the
     * column existed, which is owned by whoever filed it (`filedByUserId`).
     */
    var ownerUserId by InsightProjectionTable.ownerUserId

    /**
     * The stable key a daily run matches a repeated finding by, so it renews the insight
     * and files no twin. NULL for an insight saved from a chat and on a row folded before
     * the column existed.
     */
    var findingKey by InsightProjectionTable.findingKey
}
