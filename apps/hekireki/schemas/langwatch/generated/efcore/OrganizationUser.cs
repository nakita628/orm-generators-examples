#nullable enable

using System;

namespace LangWatch.Models;

public partial class OrganizationUser
{
    public string UserId { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public OrganizationUserRole Role { get; set; }

    public virtual User User { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    /// <summary>
    /// Department assignment for spend attribution (the people lens).
    /// Single-valued, org-scoped, accounting-only (never gates access).
    /// Plain column ref (relationMode="prisma"); the service validates the
    /// id belongs to this org. See departments.feature.
    /// </summary>
    public string? DepartmentId { get; set; }

    /// <summary>
    /// A Billing-only member (ADR-205): role EXTERNAL, the Billing grant, no team and
    /// no seat. Any role change clears it. See billing-only-invite.feature.
    /// </summary>
    public bool BillingOnly { get; set; } = false;

    /// <summary>
    /// Set when an admin disables this membership to stay within the licensed
    /// seat count. The person loses access to this organization only, keeps
    /// their role, department and everything they did, and frees their seat.
    /// Reversible: re-enabling restores access if a seat is available.
    /// Seat counting excludes disabled memberships. See
    /// seat-reconciliation.feature.
    /// </summary>
    public DateTime? DisabledAt { get; set; }

    /// <summary>
    /// Stable grant intent for an unfinished automatic SSO admission.
    /// Cleared after the authoritative grant is observed; never inferred from missing access.
    /// </summary>
    public string? PendingSsoGrantId { get; set; }

    /// <summary>
    /// A new generation for each membership lifetime. Stamped USER grant facts
    /// must match the current active membership before they can be projected;
    /// removing and rejoining therefore cannot revive a delayed old attach.
    /// </summary>
    public string MembershipStamp { get; set; } = null!;

    /// <summary>
    /// When this membership started holding the full seat it holds now; null means createdAt.
    /// Anchors the 30-day seat hold (specs/licensing/seat-rotation.feature).
    /// </summary>
    public DateTime? FullSeatSince { get; set; }

    /// <summary>
    /// Set when admission wrote this membership disabled because no seat was free: the person
    /// is waiting for one. Cleared when the membership is enabled.
    /// </summary>
    public DateTime? SeatPendingSince { get; set; }
}
