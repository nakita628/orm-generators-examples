#nullable enable

using System;

namespace Formbricks.Models;

/// <summary>
/// A creative's Impressions and Clicks on one UTC day on one device type (ENG-4000). Written
/// only by the counter flush job as `GREATEST(stored, counted)`, from Redis counters that hold the day's
/// running totals, so a repeated flush changes nothing. At most 3 rows per creative per day. Deleting
/// the creative deletes its counts.
/// </summary>
public partial class CreativeDailyStat
{
    public virtual Creative Creative { get; set; } = null!;

    public string CreativeId { get; set; } = null!;

    /// <summary>
    /// UTC day.
    /// </summary>
    public DateOnly Date { get; set; }

    public CreativeDeviceType DeviceType { get; set; }

    public int Impressions { get; set; } = 0;

    public int Clicks { get; set; } = 0;
}
