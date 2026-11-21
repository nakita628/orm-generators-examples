package langwatch.models

import org.jetbrains.exposed.v1.core.dao.id.IdTable
import org.jetbrains.exposed.v1.json.jsonb
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
object AnomalyRuleTable : IdTable<String>("\"AnomalyRule\"") {
    override val id = text("id").clientDefault { nanoid(21) }.entityId()
    val organizationId = text("organizationId")

    /**
     * 'organization' | 'team' | 'project' | 'source_type' | 'source'.
     * String not enum so new scopes ship without migrations.
     */
    val scope = text("scope")

    /**
     * Org id / team id / project id / source-type string / source id
     * - narrows what the rule applies to.
     */
    val scopeId = text("scopeId")
    val name = text("name")
    val description = text("description").nullable()

    /**
     * 'critical' | 'warning' | 'info'. Drives default destination
     * routing (e.g. critical → PagerDuty, warning → Slack only).
     */
    val severity = text("severity")

    /**
     * Extensible rule taxonomy: 'spend_spike' | 'rate_limit' |
     * 'tool_mismatch' | 'after_hours' | 'unusual_model' | 'pii_leak'
     * | 'custom'. New types ship without migrations - evaluator
     * picks a per-type implementation from a code registry.
     */
    val ruleType = text("ruleType")

    /**
     * Per-rule-type threshold shape. Examples:
     * spend_spike: { windowSec, ratioVsBaseline, minBaselineUsd }
     * rate_limit: { windowSec, requestsThreshold }
     * after_hours: { startHour, endHour, timezone, requestsThreshold }
     */
    val thresholdConfig = jsonb<String>("thresholdConfig", { it }, { it }).default("{}")

    /**
     * Alert routing per destination. Shape:
     * { slack?: { channelId }, siem?: { endpointUrl, authHeader },
     * webhook?: { url, authHeader }, pagerduty?: { routingKey },
     * email?: { to: string[] } }
     * Empty map = use org defaults (when defaults exist; today no-op).
     */
    val destinationConfig = jsonb<String>("destinationConfig", { it }, { it }).default("{}")

    /**
     * 'active' (default - evaluator picks it up) | 'disabled' (admin
     * paused; rule rows kept for audit). 'archived' uses archivedAt
     * timestamp instead so we can filter `archivedAt IS NULL` cheaply.
     */
    val status = text("status").default("active")
    val archivedAt = pgTimestamp("archivedAt", 3).nullable()
    val createdAt = pgTimestamp("createdAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val updatedAt = pgTimestamp("updatedAt", 3).databaseDefault("CURRENT_TIMESTAMP") { Instant.now().truncatedTo(ChronoUnit.MILLIS) }
    val createdById = text("createdById").nullable()
    override val primaryKey = PrimaryKey(id, name = "\"AnomalyRule_pkey\"")

    init {
        index("\"AnomalyRule_organizationId_idx\"", false, organizationId)
        index("\"AnomalyRule_scope_scopeId_idx\"", false, scope, scopeId)
        index("\"AnomalyRule_ruleType_idx\"", false, ruleType)
        index("\"AnomalyRule_status_idx\"", false, status)
        uniqueIndex("\"AnomalyRule_organizationId_name_key\"", organizationId, name)
    }
}
