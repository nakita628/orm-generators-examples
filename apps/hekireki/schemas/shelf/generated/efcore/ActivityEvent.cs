#nullable enable

using System;

namespace Shelf.Models;

public partial class ActivityEvent
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public DateTime OccurredAt { get; set; }

    public string? ActorUserId { get; set; }

    public string? ActorSnapshot { get; set; }

    public ActivityAction Action { get; set; }

    public ActivityEntity EntityType { get; set; }

    public string EntityId { get; set; } = null!;

    public string? AssetId { get; set; }

    public string? BookingId { get; set; }

    public string? AuditSessionId { get; set; }

    public string? AuditAssetId { get; set; }

    public string? KitId { get; set; }

    public string? LocationId { get; set; }

    public string? TeamMemberId { get; set; }

    public string? TargetUserId { get; set; }

    public string? Field { get; set; }

    public string? FromValue { get; set; }

    public string? ToValue { get; set; }

    public string? Meta { get; set; }

    public virtual Organization Organization { get; set; } = null!;
}
