package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityClass

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
class AnomalyAlertEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AnomalyAlertEntity>(AnomalyAlertTable)

    var organizationId by AnomalyAlertTable.organizationId
    var ruleId by AnomalyAlertTable.ruleId

    /**
     * Denormalised from rule for fast WHERE filtering on the
     * dashboard query (no join needed for the common path).
     */
    var severity by AnomalyAlertTable.severity
    var ruleName by AnomalyAlertTable.ruleName
    var ruleType by AnomalyAlertTable.ruleType

    /** Window the trigger refers to (UTC). */
    var triggerWindowStart by AnomalyAlertTable.triggerWindowStart
    var triggerWindowEnd by AnomalyAlertTable.triggerWindowEnd

    /** Numeric trigger context. Per-rule-type detail in `detail` JSONB. */
    var triggerSpendUsd by AnomalyAlertTable.triggerSpendUsd
    var triggerEventCount by AnomalyAlertTable.triggerEventCount

    /**
     * Per-rule-type extra context: {topActors: [...], baseline: ...,
     * observed: ..., timezone: ..., etc}.
     */
    var detail by AnomalyAlertTable.detail

    /** 'open' | 'acknowledged' | 'resolved'. */
    var state by AnomalyAlertTable.state

    /**
     * Per-destination dispatch outcome:
     * { webhook: { lastAttemptIso, lastError?, attemptCount, statusCode? },
     * slack: ..., pagerduty: ..., siem: ..., email: ... }
     */
    var destinationStatus by AnomalyAlertTable.destinationStatus
    var detectedAt by AnomalyAlertTable.detectedAt
    var acknowledgedAt by AnomalyAlertTable.acknowledgedAt
    var resolvedAt by AnomalyAlertTable.resolvedAt
}
