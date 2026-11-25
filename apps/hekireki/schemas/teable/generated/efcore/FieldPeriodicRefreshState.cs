#nullable enable

using System;

namespace Teable.Models;

public partial class FieldPeriodicRefreshState
{
    public string FieldId { get; set; } = null!;

    public string TableId { get; set; } = null!;

    public string BaseId { get; set; } = null!;

    public bool Enabled { get; set; } = true;

    public string TimeZone { get; set; } = "UTC";

    public bool ContainsNowPrecision { get; set; } = false;

    public DateTime? LastRefreshedAt { get; set; }

    public string? LastRefreshedLocalDate { get; set; }

    public string? LastRangeHash { get; set; }

    public string LastStatus { get; set; } = "idle";

    public string? LastError { get; set; }

    public string? PendingTaskId { get; set; }

    public string? PendingLocalDate { get; set; }

    public string? PendingRangeHash { get; set; }

    public DateTime? PendingStartedAt { get; set; }

    public DateTime CreatedTime { get; set; }

    public DateTime? LastModifiedTime { get; set; }

    public virtual Field Field { get; set; } = null!;
}
