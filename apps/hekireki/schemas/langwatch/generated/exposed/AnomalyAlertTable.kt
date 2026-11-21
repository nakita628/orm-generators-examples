package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Persisted anomaly detection. Inserted by the anomaly-detection
 * reactor (event-sourcing pipeline; modeled on PR #3351's
 * alertTrigger pattern) when a rule's threshold trips while
 * processing a freshly-arrived ActivityEventReceived event.
 * Read by api.activityMonitor.recentAnomalies for the admin dashboard.
 *
 * Dedup: (ruleId, triggerWindowStart) - re-evaluating the same
 * window updates the existing row rather than creating a duplicate.
 * Acknowledgement / resolution are admin actions on the row.
 *
 * Spec: specs/ai-gateway/governance/anomaly-detection.feature
 */
object AnomalyAlertTable : IdTable<String>("\"AnomalyAlert\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val organizationId = text("organizationId")
    val ruleId = text("ruleId")

    /**
     * Denormalised from rule for fast WHERE filtering on the
     * dashboard query (no join needed for the common path).
     */
    val severity = text("severity")
    val ruleName = text("ruleName")
    val ruleType = text("ruleType")

    /** Window the trigger refers to (UTC). */
    val triggerWindowStart = pgTimestamp("triggerWindowStart", 3)
    val triggerWindowEnd = pgTimestamp("triggerWindowEnd", 3)

    /** Numeric trigger context. Per-rule-type detail in `detail` JSONB. */
    val triggerSpendUsd = decimal("triggerSpendUsd", 18, 6).nullable()
    val triggerEventCount = integer("triggerEventCount").nullable()

    /**
     * Per-rule-type extra context: {topActors: [...], baseline: ...,
     * observed: ..., timezone: ..., etc}.
     */
    val detail = jsonb<String>("detail", { it }, { it }).default("{}")

    /** 'open' | 'acknowledged' | 'resolved'. */
    val state = text("state").default("open")

    /**
     * Per-destination dispatch outcome:
     * { webhook: { lastAttemptIso, lastError?, attemptCount, statusCode? },
     * slack: ..., pagerduty: ..., siem: ..., email: ... }
     */
    val destinationStatus = jsonb<String>("destinationStatus", { it }, { it }).default("{}")
    val detectedAt = pgTimestamp("detectedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val acknowledgedAt = pgTimestamp("acknowledgedAt", 3).nullable()
    val resolvedAt = pgTimestamp("resolvedAt", 3).nullable()
    override val primaryKey = PrimaryKey(id, name = "\"AnomalyAlert_pkey\"")

    init {
        index("\"AnomalyAlert_organizationId_detectedAt_idx\"", false, organizationId, detectedAt)
        index("\"AnomalyAlert_ruleId_idx\"", false, ruleId)
        index("\"AnomalyAlert_state_idx\"", false, state)
        uniqueIndex("\"AnomalyAlert_ruleId_triggerWindowStart_key\"", ruleId, triggerWindowStart)
    }
}
