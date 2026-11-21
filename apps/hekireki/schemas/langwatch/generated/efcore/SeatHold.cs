#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// A full seat still counted after the person who took it stopped holding it (removed,
/// disabled or demoted), until 30 days after they took it (specs/licensing/seat-rotation.feature).
/// One per organization and person; a return within the hold takes the seat back.
/// </summary>
public partial class SeatHold
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public string UserId { get; set; } = null!;

    /// <summary>
    /// Display copy for the Users list, kept because a removed person has no membership row.
    /// </summary>
    public string? Email { get; set; }

    public string? Name { get; set; }

    /// <summary>
    /// When the person took the seat this hold keeps.
    /// </summary>
    public DateTime SeatTakenAt { get; set; }

    /// <summary>
    /// seatTakenAt plus 30 days; expired holds are ignored by every count and deleted hourly.
    /// </summary>
    public DateTime HeldUntil { get; set; }

    /// <summary>
    /// REMOVED | DISABLED | DEMOTED
    /// </summary>
    public string Reason { get; set; } = null!;

    public DateTime CreatedAt { get; set; }
}
