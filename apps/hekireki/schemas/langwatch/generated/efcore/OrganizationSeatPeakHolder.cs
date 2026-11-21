#nullable enable

namespace LangWatch.Models;

/// <summary>
/// One member's or pending invitation's seat as the seat peak last applied it (ADR-002).
/// </summary>
public partial class OrganizationSeatPeakHolder
{
    public string OrganizationId { get; set; } = null!;

    /// <summary>
    /// `member`, `invitation` or `hold` (a seat held after somebody left).
    /// </summary>
    public string HolderKind { get; set; } = null!;

    /// <summary>
    /// The member's or the hold's userId, or the invitation's id.
    /// </summary>
    public string HolderId { get; set; } = null!;

    /// <summary>
    /// `full`, `lite`, `developer` or `none`.
    /// </summary>
    public string Seat { get; set; } = null!;

    public double LastChangedAt { get; set; }

    /// <summary>
    /// Absent for a seat the opening snapshot recorded.
    /// </summary>
    public string? LastChangeId { get; set; }

    /// <summary>
    /// An invitation's expiry; null for a member or an invitation that never expires.
    /// </summary>
    public double? ExpiresAt { get; set; }

    /// <summary>
    /// The fold gave the invitation's seat back at `expiresAt`.
    /// </summary>
    public bool ExpiredByFold { get; set; } = false;

    /// <summary>
    /// A hold whose seat a live full invitation to its address took: that invitation's id.
    /// </summary>
    public string? SuppressedBy { get; set; }

    /// <summary>
    /// The fold gave a suppressed hold its seat back when that invitation expired.
    /// </summary>
    public bool RestoredByFold { get; set; } = false;
}
