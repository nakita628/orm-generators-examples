#nullable enable

namespace LangWatch.Models;

/// <summary>
/// Organization's monthly seat peak, folded from its seat changes (modules/organization, ADR-002).
/// A read model, one row per organization: rebuildable by replaying the event log, so no relation.
/// Epochs in milliseconds.
/// Organization's seat-peak counts (ADR-002): one row per organization once counting opened.
/// </summary>
public partial class OrganizationSeatPeak
{
    /// <summary>
    /// The organization, which is also the aggregate id.
    /// </summary>
    public string OrganizationId { get; set; } = null!;

    /// <summary>
    /// When the opening snapshot was taken, epoch milliseconds.
    /// </summary>
    public double OpenedAt { get; set; }

    public int FullSeats { get; set; }

    public int LiteSeats { get; set; }

    public int DeveloperSeats { get; set; }

    /// <summary>
    /// The latest change time applied to any holder; invitations expire up to it.
    /// </summary>
    public double LatestChangeAt { get; set; }

    /// <summary>
    /// Changes whose `from` disagreed with the seat the fold held.
    /// </summary>
    public int Mismatches { get; set; }
}
