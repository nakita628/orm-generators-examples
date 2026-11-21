#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// Persisted anomaly detection. Inserted by the anomaly-detection
/// reactor (event-sourcing pipeline; modeled on PR #3351's
/// alertTrigger pattern) when a rule's threshold trips while
/// processing a freshly-arrived ActivityEventReceived event.
/// Read by api.activityMonitor.recentAnomalies for the admin dashboard.
///
/// Dedup: (ruleId, triggerWindowStart) - re-evaluating the same
/// window updates the existing row rather than creating a duplicate.
/// Acknowledgement / resolution are admin actions on the row.
///
/// Spec: specs/ai-gateway/governance/anomaly-detection.feature
/// </summary>
public partial class AnomalyAlert
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    public string RuleId { get; set; } = null!;

    public virtual AnomalyRule Rule { get; set; } = null!;

    /// <summary>
    /// Denormalised from rule for fast WHERE filtering on the
    /// dashboard query (no join needed for the common path).
    /// </summary>
    public string Severity { get; set; } = null!;

    public string RuleName { get; set; } = null!;

    public string RuleType { get; set; } = null!;

    /// <summary>
    /// Window the trigger refers to (UTC).
    /// </summary>
    public DateTime TriggerWindowStart { get; set; }

    public DateTime TriggerWindowEnd { get; set; }

    /// <summary>
    /// Numeric trigger context. Per-rule-type detail in `detail` JSONB.
    /// </summary>
    public decimal? TriggerSpendUsd { get; set; }

    public int? TriggerEventCount { get; set; }

    /// <summary>
    /// Per-rule-type extra context: {topActors: [...], baseline: ...,
    /// observed: ..., timezone: ..., etc}.
    /// </summary>
    public string Detail { get; set; } = "{}";

    /// <summary>
    /// 'open' | 'acknowledged' | 'resolved'.
    /// </summary>
    public string State { get; set; } = "open";

    /// <summary>
    /// Per-destination dispatch outcome:
    /// { webhook: { lastAttemptIso, lastError?, attemptCount, statusCode? },
    /// slack: ..., pagerduty: ..., siem: ..., email: ... }
    /// </summary>
    public string DestinationStatus { get; set; } = "{}";

    public DateTime DetectedAt { get; set; }

    public DateTime? AcknowledgedAt { get; set; }

    public DateTime? ResolvedAt { get; set; }
}
