package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

/**
 * One person's daily insights run on one board in a project, folded from its
 * `insight_daily_schedule` aggregate (modules/insight). A read model: rebuildable by
 * replaying the event log. Epochs in milliseconds.
 */
class InsightDailyScheduleProjectionEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, InsightDailyScheduleProjectionEntity>(InsightDailyScheduleProjectionTable)

    var projectId by InsightDailyScheduleProjectionTable.projectId
    var userId by InsightDailyScheduleProjectionTable.userId

    /** dashboard | template. A pointer only, never a relation: a run checks the board is there. */
    var boardKind by InsightDailyScheduleProjectionTable.boardKind
    var boardId by InsightDailyScheduleProjectionTable.boardId
    var boardName by InsightDailyScheduleProjectionTable.boardName

    /** on | off | undecided. */
    var state by InsightDailyScheduleProjectionTable.state

    /**
     * board_deleted | skipped_days | member_removed | project_archived. Why the system turned
     * it off. NULL while it is on, and when the person turned it off.
     */
    var offReason by InsightDailyScheduleProjectionTable.offReason
    var hour by InsightDailyScheduleProjectionTable.hour
    var timezone by InsightDailyScheduleProjectionTable.timezone
    var maxInsights by InsightDailyScheduleProjectionTable.maxInsights

    /** The last run that recorded an outcome. NULL until one did. */
    var lastRunId by InsightDailyScheduleProjectionTable.lastRunId
    var lastRunAt by InsightDailyScheduleProjectionTable.lastRunAt

    /** filed | nothing | failed | skipped */
    var lastRunOutcome by InsightDailyScheduleProjectionTable.lastRunOutcome
    var lastRunReason by InsightDailyScheduleProjectionTable.lastRunReason
    var lastRunFiled by InsightDailyScheduleProjectionTable.lastRunFiled

    /** How many insights the run found still true and renewed. */
    var lastRunRenewed by InsightDailyScheduleProjectionTable.lastRunRenewed
    var lastRunConversationId by InsightDailyScheduleProjectionTable.lastRunConversationId

    /** The scheduled days in a row whose run ended skipped, and the slot of the newest counted. */
    var consecutiveSkips by InsightDailyScheduleProjectionTable.consecutiveSkips
    var skipStreakSlot by InsightDailyScheduleProjectionTable.skipStreakSlot
    var createdAt by InsightDailyScheduleProjectionTable.createdAt
    var updatedAt by InsightDailyScheduleProjectionTable.updatedAt
    var occurredAt by InsightDailyScheduleProjectionTable.occurredAt
    var acceptedAt by InsightDailyScheduleProjectionTable.acceptedAt
    var lastEventId by InsightDailyScheduleProjectionTable.lastEventId
    var projectionVersion by InsightDailyScheduleProjectionTable.projectionVersion
}
