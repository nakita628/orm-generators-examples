#nullable enable

using System;

namespace Rallly.Models;

/// <summary>
/// Monthly counters for metered actions charged to a space, independent of
/// the records they produce, so deleting a poll does not give back what it
/// spent.
/// </summary>
public partial class SpaceUsage
{
    public string SpaceId { get; set; } = null!;

    public SpaceUsageResource Resource { get; set; }

    public DateOnly Period { get; set; }

    public int Count { get; set; } = 0;

    public DateTime? NearLimitWarnedAt { get; set; }

    public DateTime? AtLimitWarnedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual Space Space { get; set; } = null!;
}
