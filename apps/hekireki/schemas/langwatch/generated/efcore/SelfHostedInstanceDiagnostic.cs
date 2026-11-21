#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// One upgrade diagnostics bundle a self-hosted install sent over Connect, on
/// LangWatch Cloud. The bytes live in object storage under
/// `upgrade-diagnostics/` and expire at 30 days by a lifecycle rule; this row
/// points at them and feeds Cloud admin and the Slack signal. Plain column
/// refs, no relation, as the usage report rows.
/// </summary>
public partial class SelfHostedInstanceDiagnostic
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// The `SelfHostedInstance.instanceId` the licence caller presented.
    /// </summary>
    public string InstanceId { get; set; } = null!;

    public string StepId { get; set; } = null!;

    public string? StepDescription { get; set; }

    public string? Release { get; set; }

    public int Attempt { get; set; }

    /// <summary>
    /// `operator` or `automatic`.
    /// </summary>
    public string Trigger { get; set; } = null!;

    public int ByteLength { get; set; }

    public string ObjectKey { get; set; } = null!;

    public DateTime ReceivedAt { get; set; }

    /// <summary>
    /// When the lifecycle rule removes the bundle; the row is listed until then.
    /// </summary>
    public DateTime ExpiresAt { get; set; }

    /// <summary>
    /// Set on the report that posted this step and release's one Slack message.
    /// </summary>
    public DateTime? SlackAnnouncedAt { get; set; }
}
