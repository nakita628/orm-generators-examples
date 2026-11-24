#nullable enable

using System;

namespace Langfuse.Models;

/// <summary>
/// One notification per limit, subject, alert type, threshold, counting start and window, so the gateway may report a limit any number of times.
/// </summary>
public partial class GatewaySpendLimitAlert
{
    public string Id { get; set; } = null!;

    public string SpendLimitId { get; set; } = null!;

    public virtual GatewaySpendLimit SpendLimit { get; set; } = null!;

    public string SubjectId { get; set; } = "";

    public GatewaySpendLimitAlertType Type { get; set; }

    public int ThresholdPercent { get; set; } = 100;

    public DateTime CountingStartedAt { get; set; }

    public DateTime WindowStart { get; set; }

    public DateTime CreatedAt { get; set; }
}
