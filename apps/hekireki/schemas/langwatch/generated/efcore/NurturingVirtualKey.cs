#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// Each gateway virtual key, folded from gateway's vk_lifecycle facts: nurturing's own copy. The
/// creator comes from the created fact alone; the holder and purpose are each the latest non-null
/// value by the fact's time (stamped beside them); the state times keep their latest.
/// </summary>
public partial class NurturingVirtualKey
{
    public string VirtualKeyId { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public string Name { get; set; } = null!;

    public DateTime OccurredAt { get; set; }

    public DateTime? KeyCreatedAt { get; set; }

    public string? CreatorUserId { get; set; }

    public string? PrincipalUserId { get; set; }

    public DateTime? PrincipalUserIdAt { get; set; }

    public string? Purpose { get; set; }

    public DateTime? PurposeAt { get; set; }

    public DateTime? DisabledAt { get; set; }

    public DateTime? EnabledAt { get; set; }

    public DateTime? RevokedAt { get; set; }
}
