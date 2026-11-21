#nullable enable

using System;
using System.Collections.Generic;

namespace LangWatch.Models;

/// <summary>
/// Admin-defined anomaly detection rule. Evaluated server-side at
/// ingest time (Option C - not in this slice yet) against
/// gateway_activity_events. When triggered, an alert is dispatched
/// to the destinations configured on the rule (Slack / SIEM / webhook
/// / PagerDuty / email).
///
/// This is a configuration entity - the actual rule evaluation +
/// alert state lives in a future ClickHouse table
/// (gateway_anomaly_alerts) and a server-side worker. Spec:
/// specs/ai-gateway/governance/anomaly-rules.feature
/// (Alexis lane B C-UI; mocked-v0 today, real persistence is THIS slice).
/// </summary>
public partial class AnomalyRule
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    /// <summary>
    /// 'organization' | 'team' | 'project' | 'source_type' | 'source'.
    /// String not enum so new scopes ship without migrations.
    /// </summary>
    public string Scope { get; set; } = null!;

    /// <summary>
    /// Org id / team id / project id / source-type string / source id
    /// - narrows what the rule applies to.
    /// </summary>
    public string ScopeId { get; set; } = null!;

    public string Name { get; set; } = null!;

    public string? Description { get; set; }

    /// <summary>
    /// 'critical' | 'warning' | 'info'. Drives default destination
    /// routing (e.g. critical → PagerDuty, warning → Slack only).
    /// </summary>
    public string Severity { get; set; } = null!;

    /// <summary>
    /// Extensible rule taxonomy: 'spend_spike' | 'rate_limit' |
    /// 'tool_mismatch' | 'after_hours' | 'unusual_model' | 'pii_leak'
    /// | 'custom'. New types ship without migrations - evaluator
    /// picks a per-type implementation from a code registry.
    /// </summary>
    public string RuleType { get; set; } = null!;

    /// <summary>
    /// Per-rule-type threshold shape. Examples:
    /// spend_spike: { windowSec, ratioVsBaseline, minBaselineUsd }
    /// rate_limit: { windowSec, requestsThreshold }
    /// after_hours: { startHour, endHour, timezone, requestsThreshold }
    /// </summary>
    public string ThresholdConfig { get; set; } = "{}";

    /// <summary>
    /// Alert routing per destination. Shape:
    /// { slack?: { channelId }, siem?: { endpointUrl, authHeader },
    /// webhook?: { url, authHeader }, pagerduty?: { routingKey },
    /// email?: { to: string[] } }
    /// Empty map = use org defaults (when defaults exist; today no-op).
    /// </summary>
    public string DestinationConfig { get; set; } = "{}";

    /// <summary>
    /// 'active' (default - evaluator picks it up) | 'disabled' (admin
    /// paused; rule rows kept for audit). 'archived' uses archivedAt
    /// timestamp instead so we can filter `archivedAt IS NULL` cheaply.
    /// </summary>
    public string Status { get; set; } = "active";

    public DateTime? ArchivedAt { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public string? CreatedById { get; set; }

    public virtual User? CreatedBy { get; set; }

    public virtual ICollection<AnomalyAlert> Alerts { get; set; } = new List<AnomalyAlert>();
}
