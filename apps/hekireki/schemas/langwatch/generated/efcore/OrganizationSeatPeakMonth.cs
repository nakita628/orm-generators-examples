#nullable enable

namespace LangWatch.Models;

/// <summary>
/// One calendar month's (UTC) seat peaks for an organization (ADR-002).
/// </summary>
public partial class OrganizationSeatPeakMonth
{
    public string OrganizationId { get; set; } = null!;

    /// <summary>
    /// `YYYY-MM`.
    /// </summary>
    public string Month { get; set; } = null!;

    public int FullPeak { get; set; }

    public double FullPeakAt { get; set; }

    public int LitePeak { get; set; }

    public double LitePeakAt { get; set; }

    public int FullAtEnd { get; set; }

    public int LiteAtEnd { get; set; }
}
