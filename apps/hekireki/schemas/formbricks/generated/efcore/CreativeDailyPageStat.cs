#nullable enable

using System;

namespace Formbricks.Models;

/// <summary>
/// A creative's Impressions and Clicks on one UTC day on one page (host + path), written like
/// `CreativeDailyStat`. At most 1000 pages per creative per day; the rest are counted under `(other)`.
/// </summary>
public partial class CreativeDailyPageStat
{
    public virtual Creative Creative { get; set; } = null!;

    public string CreativeId { get; set; } = null!;

    /// <summary>
    /// UTC day.
    /// </summary>
    public DateOnly Date { get; set; }

    /// <summary>
    /// Host + path, without query string or hash, or `(other)`.
    /// </summary>
    public string Page { get; set; } = null!;

    public int Impressions { get; set; } = 0;

    public int Clicks { get; set; } = 0;
}
