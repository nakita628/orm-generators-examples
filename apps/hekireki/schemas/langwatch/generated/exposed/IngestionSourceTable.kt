package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Per-platform fleet configuration that connects a closed SaaS
 * product's audit / OTel / S3 stream to LangWatch's Activity
 * Monitor (cf. specs/ai-gateway/governance/activity-monitor.feature
 * + ingestion-sources.feature, docs/ai-gateway/governance/architecture.md).
 *
 * One IngestionSource = one upstream fleet (e.g. "Miro Cowork",
 * "Acme Workato production"). The runtime ingests events on whatever
 * cadence the source supports - push for OTel/webhook/S3 (Tier C/D),
 * pull for poll-based admin APIs (Tier C). Per-source-type adapters
 * normalise platform-specific shapes into OCSF + AOS before they
 * land in trace_summaries with SourceType + SourceId tags.
 */
object IngestionSourceTable : IdTable<String>("\"IngestionSource\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val organizationId = text("organizationId")

    /** Optional team-scope. null = org-wide. */
    val teamId = text("teamId").nullable()

    /**
     * 'otel_generic' | 'claude_cowork' | 'workato' | 'copilot_studio' |
     * 'openai_compliance' | 'claude_compliance' | 's3_custom' (extensible).
     * String not enum so adapters can be added without migrations - the
     * supported set lives in code (activity-monitor adapter registry).
     */
    val sourceType = text("sourceType")

    /**
     * Human-friendly label for the admin UI. Unique per-org so two
     * "Cowork" sources can coexist with different names.
     */
    val name = text("name")
    val description = text("description").nullable()

    /**
     * HMAC bearer token the upstream operator pastes into the platform
     * (or LangWatch sends in pull mode). Stored encrypted at rest via
     * the same pepper-derived KEK as VirtualKey secrets. Auto-generated
     * on create, rotatable via the admin UI with a 24h grace window.
     */
    val ingestSecretHash = text("ingestSecretHash")

    /**
     * Per-source-type configuration. Shape is platform-specific:
     * - otel_generic: { allowedSourceTypeLabel?: string }
     * - claude_cowork: { workspaceId: string }
     * - workato: { sharedSecretLastFour: string }
     * - copilot_studio: { tenantId, clientId, clientSecretHash, pollEverySec }
     * - openai_compliance: { bucket, prefix, roleArn, pollEverySec }
     * - claude_compliance: { workspaceApiKeyHash, pollEverySec }
     * - s3_custom: { bucket, prefix, roleArn, parserDsl, pollEverySec }
     */
    val parserConfig = jsonb<String>("parserConfig", { it }, { it }).default("{}")

    /**
     * For pull-based sources, persisted between polls so the next poll
     * resumes where the last left off (timestamp / cursor / file marker).
     */
    val pollerCursor = jsonb<String>("pollerCursor", { it }, { it }).nullable()

    /**
     * Phase 10 - number of consecutive errors observed by the puller
     * worker. Reset to 0 on a successful run. Drives admin-UI banners
     * ("Microsoft authentication failed - re-authorize") + back-off
     * scheduling when the source has been broken for a while.
     * Spec: specs/ai-governance/puller-framework/puller-adapter-contract.feature
     */
    val errorCount = integer("errorCount").default(0)

    /**
     * Phase 10 - cron schedule for pull-mode sources. Validated by the
     * adapter's pullConfig schema; mirrored here for fast worker
     * dispatch. Null for push-mode sources.
     */
    val pullSchedule = text("pullSchedule").nullable()

    /**
     * Status flag. 'active' (default), 'disabled' (admin paused), or
     * 'awaiting_first_event' (created but never received).
     */
    val status = text("status").default("awaiting_first_event")

    /**
     * ADR-088 v7: optional trace DESTINATION for conversation-bearing pulls
     * (Databricks Genie). Null = don't route. Same shape as
     * VirtualKey.traceProjectId: a destination is NOT a scope row and grants
     * no access to the source; it only decides where routed conversation
     * traces land. Must be a project of this org (validated at write time).
     * Freely editable; already-routed history stays where it landed.
     * NOT the deferred cost-attribution projectId (ADR-088 Decision 4).
     */
    val traceProjectId = text("traceProjectId").nullable()

    /**
     * The account the provider itself named when this connection was saved -
     * an OpenAI or Anthropic organization id, read once per save and compared
     * so a second connection cannot read one account's spend twice. NOT
     * derived from the administrator key: nothing worked out from a customer
     * credential is stored here or anywhere else, because the account id is
     * not a secret and cannot be turned back into a key.
     * Nullable because every connection that predates the guard has none, and
     * because Azure and Power Platform name what they read in parserConfig
     * instead. "Unknown account" has to be a storable state.
     * Spec: specs/ai-gateway/governance/ingestion-sources.feature
     */
    val providerAccountId = text("providerAccountId").nullable()

    /**
     * Last event timestamp observed (any direction). Surfaces in the
     * admin UI as "last event 2 minutes ago".
     */
    val lastEventAt = pgTimestamp("lastEventAt", 3).nullable()

    /**
     * ADR-128 - when a pull run last reached the provider without erroring,
     * mirrored from the ingestionPullRunStatus projection. Deliberately NOT
     * lastEventAt: an empty period is a working puller, and a run that
     * delivered data and then failed is not. Drives the health badge and
     * the "no data since" cut-off, below which a day's spend is unknown
     * rather than zero.
     * Spec: specs/governance/ingestion-source-health.feature
     */
    val lastSuccessAt = pgTimestamp("lastSuccessAt", 3).nullable()

    /**
     * ADR-128 - how far the last run actually read, and whether it reached the
     * end. Mirrored from the ingestionPullRunStatus projection, in the same way
     * and for the same reason as lastSuccessAt above: every reader of this
     * surface selects from this row, and two of them render a LIST of sources,
     * so reaching a per-run projection would be one query per row on a screen
     * built to show many at once.
     *
     * These two are the only evidence separating a source that truncated once
     * and then finished from one that is stuck. A stuck source's run clock
     * advances on every attempt, so the clock cannot be that evidence and is
     * deliberately not stored here - only the point the run REACHED.
     *
     * Completeness is "complete" or "truncated". Null on both is unknown, which
     * is what every row written before this existed reads as, and unknown must
     * never render as either answer.
     * Spec: specs/governance/ingestion-source-health.feature
     */
    val lastReadThroughAt = pgTimestamp("lastReadThroughAt", 3).nullable()
    val lastRunCompleteness = text("lastRunCompleteness").nullable()

    /**
     * ADR-088 - the window this source read while pulled cost recording was
     * off for the organization, so its spend was dropped rather than priced.
     * Both null = nothing dropped. The pull cursor advances whether or not the
     * money path is live (audit-only is a supported mode), so without this the
     * loss is unrecoverable AND invisible: those days read as no spend rather
     * than as unknown. Widened by any run that drops a priced item, cleared
     * only when a later run re-reads back across the whole window.
     * Spec: specs/governance/pulled-usage-cost-gap.feature
     */
    val unpricedUsageSince = pgTimestamp("unpricedUsageSince", 3).nullable()
    val unpricedUsageThrough = pgTimestamp("unpricedUsageThrough", 3).nullable()
    val archivedAt = pgTimestamp("archivedAt", 3).nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val createdById = text("createdById").nullable()
    override val primaryKey = PrimaryKey(id, name = "\"IngestionSource_pkey\"")

    init {
        index("\"IngestionSource_organizationId_idx\"", false, organizationId)
        index("\"IngestionSource_teamId_idx\"", false, teamId)
        index("\"IngestionSource_sourceType_idx\"", false, sourceType)
        uniqueIndex("\"IngestionSource_organizationId_name_key\"", organizationId, name)
    }
}
