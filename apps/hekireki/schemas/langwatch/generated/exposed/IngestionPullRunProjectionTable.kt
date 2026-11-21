package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable

/** Rebuildable per-source pull checkpoint and operator-facing run status. */
object IngestionPullRunProjectionTable : IdTable<String>("\"IngestionPullRunProjection\"") {
    override val id = text("id").entityId()
    val sourceId = text("sourceId").uniqueIndex("\"IngestionPullRunProjection_sourceId_key\"")
    val projectId = text("projectId")
    val enabled = bool("enabled").default(false)
    val cron = text("cron").nullable()
    val cursor = text("cursor").nullable()
    val lastRunAt = double("lastRunAt").nullable()
    val lastRunOutcome = text("lastRunOutcome").nullable()
    val lastRunEventCount = integer("lastRunEventCount").default(0)
    val lastRunError = text("lastRunError").nullable()
    val lastRunErrorCode = text("lastRunErrorCode").nullable()
    val consecutiveErrors = integer("consecutiveErrors").default(0)

    /**
     * When a run last completed without erroring. Backfilled by replay, which
     * is why the projection version moved with this column.
     */
    val lastSuccessAt = double("lastSuccessAt").nullable()

    /**
     * Identity of the run whose outcome this row reflects. Outcomes from a
     * superseded run arrive with a smaller value and are ignored, so a late
     * completion cannot regress Cursor or the status fields.
     */
    val lastRunScheduledFor = double("lastRunScheduledFor").nullable()

    /**
     * ADR-128 - how far the last run read, in epoch milliseconds, and whether
     * it reached the end. The fold writes both; the IngestionSource row mirrors
     * them for the screens. Null read-through means no run has reported one.
     * Null completeness reads as unknown, never as either answer, which is what
     * every row written before this existed holds.
     *
     * Nullable on purpose, so replay is not required to make this migration
     * safe: a row that has never seen a run carrying these stays null and every
     * reader already handles unknown.
     * Spec: specs/governance/ingestion-source-health.feature
     */
    val lastReadThroughAt = double("lastReadThroughAt").nullable()
    val lastRunCompleteness = text("lastRunCompleteness").nullable()

    /**
     * The last time somebody asked this source to enumerate its agents, and how
     * that ask ended. A listing is a different scope on the same connection as
     * the pull above: a credential that cannot enumerate a directory still pulls
     * cost perfectly, so no health, coverage or completeness reader may treat a
     * refusal here as the source failing to pull.
     *
     * Count is set ONLY when the outcome is listed, and stays null on a refusal.
     * Zero is a real answer meaning the provider returned an empty list, so a
     * zero default would manufacture that answer for a source that refused.
     * Reason is set only on a refusal and is a plain string, never an enum and
     * never a foreign key, because the log outlives the vocabulary. Status is
     * nullable even when refused, because a refusal we generated ourselves has
     * no HTTP status, and it is never rendered to a customer.
     *
     * The two `...ListingAt` columns hold the instant the ASK was accepted, not
     * the instant its answer was recorded, and the fold orders outcomes by them
     * so a superseded listing cannot overwrite the one that replaced it. The
     * name predates that job and is kept because renaming a column is a
     * migration; the fold's own field docs are the statement of record.
     */
    val lastAgentsListingAt = double("lastAgentsListingAt").nullable()
    val lastAgentsListingOutcome = text("lastAgentsListingOutcome").nullable()
    val lastAgentsListingCount = integer("lastAgentsListingCount").nullable()
    val lastAgentsListingReason = text("lastAgentsListingReason").nullable()
    val lastAgentsListingStatus = integer("lastAgentsListingStatus").nullable()

    /**
     * The same for people, in its own set of columns rather than sharing the
     * agents set with a kind discriminator. A source can be refused for people
     * and fine for agents, and one shared set holds only the most recent listing
     * of either kind, so syncing people would erase an agents refusal and the
     * agents page would go back to not knowing after briefly knowing.
     *
     * LastPeopleDirectoryCount is everyone the provider's directory named.
     * LastPeopleWithheldCount is how many of those this deployment does not hold
     * because erasure suppression removed them: a SUBSET of the directory count,
     * never an addition to it. Adding the two counts the same people twice.
     * Subtracting is the valid arithmetic, and the surviving total is that
     * subtraction rather than a third column that could disagree.
     *
     * The withheld count is safe to show as a CURRENT figure and never as a
     * series: it moving from zero to one at a known moment says an erasure
     * happened then, which on a small tenant identifies the person as surely as
     * a name would. So no trend line, no history drawer, no per-run export, and
     * never beside a per-person list.
     */
    val lastPeopleListingAt = double("lastPeopleListingAt").nullable()
    val lastPeopleListingOutcome = text("lastPeopleListingOutcome").nullable()
    val lastPeopleDirectoryCount = integer("lastPeopleDirectoryCount").nullable()
    val lastPeopleWithheldCount = integer("lastPeopleWithheldCount").nullable()
    val lastPeopleListingReason = text("lastPeopleListingReason").nullable()
    val lastPeopleListingStatus = integer("lastPeopleListingStatus").nullable()
    val createdAt = double("createdAt")
    val updatedAt = double("updatedAt")
    val occurredAt = double("occurredAt")
    val acceptedAt = double("acceptedAt")
    val lastEventId = text("lastEventId")
    val projectionVersion = text("projectionVersion")
    override val primaryKey = PrimaryKey(id, name = "\"IngestionPullRunProjection_pkey\"")

    init {
        index("\"IngestionPullRunProjection_projectId_idx\"", false, projectId)
    }
}
