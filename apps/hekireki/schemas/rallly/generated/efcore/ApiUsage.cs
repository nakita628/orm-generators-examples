#nullable enable

using System;

namespace Rallly.Models;

/// <summary>
/// Superseded by SpaceUsage (resource api_poll). Unused; kept so old
/// deployments keep working during the switch, then dropped.
/// </summary>
public partial class ApiUsage
{
    public string SpaceId { get; set; } = null!;

    public ApiUsageResource Resource { get; set; }

    public DateOnly Period { get; set; }

    public int Count { get; set; } = 0;

    public DateTime? NearLimitWarnedAt { get; set; }

    public DateTime? AtLimitWarnedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual Space Space { get; set; } = null!;
}
