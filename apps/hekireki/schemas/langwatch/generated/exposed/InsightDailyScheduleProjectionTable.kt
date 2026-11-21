package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable

/**
 * One person's daily insights run on one board in a project, folded from its
 * `insight_daily_schedule` aggregate (modules/insight). A read model: rebuildable by
 * replaying the event log. Epochs in milliseconds.
 */
object InsightDailyScheduleProjectionTable : IdTable<String>("\"InsightDailyScheduleProjection\"") {
    /** The schedule id, which is also its aggregate id: one per project, person and board. */
    override val id = text("id").entityId()
    val projectId = text("projectId")
    val userId = text("userId")

    /** dashboard | template. A pointer only, never a relation: a run checks the board is there. */
    val boardKind = text("boardKind")
    val boardId = text("boardId")
    val boardName = text("boardName")

    /** on | off | undecided. */
    val state = text("state").default("undecided")

    /**
     * board_deleted | skipped_days | member_removed | project_archived. Why the system turned
     * it off. NULL while it is on, and when the person turned it off.
     */
    val offReason = text("offReason").nullable()
    val hour = integer("hour").nullable()
    val timezone = text("timezone").nullable()
    val maxInsights = integer("maxInsights").nullable()

    /** The last run that recorded an outcome. NULL until one did. */
    val lastRunId = text("lastRunId").nullable()
    val lastRunAt = double("lastRunAt").nullable()

    /** filed | nothing | failed | skipped */
    val lastRunOutcome = text("lastRunOutcome").nullable()
    val lastRunReason = text("lastRunReason").nullable()
    val lastRunFiled = integer("lastRunFiled").nullable()

    /** How many insights the run found still true and renewed. */
    val lastRunRenewed = integer("lastRunRenewed").nullable()
    val lastRunConversationId = text("lastRunConversationId").nullable()

    /** The scheduled days in a row whose run ended skipped, and the slot of the newest counted. */
    val consecutiveSkips = integer("consecutiveSkips").default(0)
    val skipStreakSlot = double("skipStreakSlot").nullable()
    val createdAt = double("createdAt")
    val updatedAt = double("updatedAt")
    val occurredAt = double("occurredAt")
    val acceptedAt = double("acceptedAt")
    val lastEventId = text("lastEventId")
    val projectionVersion = text("projectionVersion")
    override val primaryKey = PrimaryKey(id, name = "\"InsightDailyScheduleProjection_pkey\"")

    init {
        uniqueIndex("\"InsightDailySchedule_projectId_userId_board_key\"", projectId, userId, boardKind, boardId)
    }
}
