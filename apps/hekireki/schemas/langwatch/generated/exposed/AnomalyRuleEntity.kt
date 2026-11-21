package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.EntityID
import org.jetbrains.exposed.v1.dao.Entity
import org.jetbrains.exposed.v1.dao.EntityBatchUpdate
import org.jetbrains.exposed.v1.dao.EntityClass
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Admin-defined anomaly detection rule. Evaluated server-side at
 * ingest time (Option C - not in this slice yet) against
 * gateway_activity_events. When triggered, an alert is dispatched
 * to the destinations configured on the rule (Slack / SIEM / webhook
 * / PagerDuty / email).
 *
 * This is a configuration entity - the actual rule evaluation +
 * alert state lives in a future ClickHouse table
 * (gateway_anomaly_alerts) and a server-side worker. Spec:
 * specs/ai-gateway/governance/anomaly-rules.feature
 * (Alexis lane B C-UI; mocked-v0 today, real persistence is THIS slice).
 */
class AnomalyRuleEntity(
    id: EntityID<String>,
) : Entity<String>(id) {
    companion object : EntityClass<String, AnomalyRuleEntity>(AnomalyRuleTable)

    var organizationId by AnomalyRuleTable.organizationId

    /**
     * 'organization' | 'team' | 'project' | 'source_type' | 'source'.
     * String not enum so new scopes ship without migrations.
     */
    var scope by AnomalyRuleTable.scope

    /**
     * Org id / team id / project id / source-type string / source id
     * - narrows what the rule applies to.
     */
    var scopeId by AnomalyRuleTable.scopeId
    var name by AnomalyRuleTable.name
    var description by AnomalyRuleTable.description

    /**
     * 'critical' | 'warning' | 'info'. Drives default destination
     * routing (e.g. critical → PagerDuty, warning → Slack only).
     */
    var severity by AnomalyRuleTable.severity

    /**
     * Extensible rule taxonomy: 'spend_spike' | 'rate_limit' |
     * 'tool_mismatch' | 'after_hours' | 'unusual_model' | 'pii_leak'
     * | 'custom'. New types ship without migrations - evaluator
     * picks a per-type implementation from a code registry.
     */
    var ruleType by AnomalyRuleTable.ruleType

    /**
     * Per-rule-type threshold shape. Examples:
     * spend_spike: { windowSec, ratioVsBaseline, minBaselineUsd }
     * rate_limit: { windowSec, requestsThreshold }
     * after_hours: { startHour, endHour, timezone, requestsThreshold }
     */
    var thresholdConfig by AnomalyRuleTable.thresholdConfig

    /**
     * Alert routing per destination. Shape:
     * { slack?: { channelId }, siem?: { endpointUrl, authHeader },
     * webhook?: { url, authHeader }, pagerduty?: { routingKey },
     * email?: { to: string[] } }
     * Empty map = use org defaults (when defaults exist; today no-op).
     */
    var destinationConfig by AnomalyRuleTable.destinationConfig

    /**
     * 'active' (default - evaluator picks it up) | 'disabled' (admin
     * paused; rule rows kept for audit). 'archived' uses archivedAt
     * timestamp instead so we can filter `archivedAt IS NULL` cheaply.
     */
    var status by AnomalyRuleTable.status
    var archivedAt by AnomalyRuleTable.archivedAt
    var createdAt by AnomalyRuleTable.createdAt
    var updatedAt by AnomalyRuleTable.updatedAt
    var createdById by AnomalyRuleTable.createdById

    override fun flush(batch: EntityBatchUpdate?): Boolean {
        if (writeValues.isNotEmpty()) {
            if (!isWritten(AnomalyRuleTable.updatedAt)) {
                updatedAt = Instant.now().truncatedTo(ChronoUnit.MILLIS)
            }
        }
        return super.flush(batch)
    }
}
