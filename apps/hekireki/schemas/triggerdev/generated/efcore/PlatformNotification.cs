#nullable enable

using System;
using System.Collections.Generic;

namespace Trigger.Models;

/// <summary>
/// Admin-created notification definitions
/// </summary>
public partial class PlatformNotification
{
    public string Id { get; set; } = null!;

    public string FriendlyId { get; set; } = null!;

    /// <summary>
    /// Admin-facing title for identification in admin tools
    /// </summary>
    public string Title { get; set; } = null!;

    /// <summary>
    /// Versioned JSON rendering content (see payload schema in spec)
    /// </summary>
    public string Payload { get; set; } = null!;

    public PlatformNotificationSurface Surface { get; set; }

    public PlatformNotificationScope Scope { get; set; }

    /// <summary>
    /// Set when scope = USER
    /// </summary>
    public virtual User? User { get; set; }

    public string? UserId { get; set; }

    /// <summary>
    /// Set when scope = ORGANIZATION
    /// </summary>
    public virtual Organization? Organization { get; set; }

    public string? OrganizationId { get; set; }

    /// <summary>
    /// Set when scope = PROJECT
    /// </summary>
    public virtual Project? Project { get; set; }

    public string? ProjectId { get; set; }

    /// <summary>
    /// When notification becomes active
    /// </summary>
    public DateTime StartsAt { get; set; }

    /// <summary>
    /// When notification expires
    /// </summary>
    public DateTime EndsAt { get; set; }

    /// <summary>
    /// CLI: auto-expire N days after user first saw it
    /// </summary>
    public int? CliMaxDaysAfterFirstSeen { get; set; }

    /// <summary>
    /// CLI: auto-expire after shown N times to user
    /// </summary>
    public int? CliMaxShowCount { get; set; }

    /// <summary>
    /// CLI: only show every N-th request (e.g. cliShowEvery=3 means show on 3rd, 6th, 9th…)
    /// </summary>
    public int? CliShowEvery { get; set; }

    /// <summary>
    /// Ordering within same scope level (higher = more important)
    /// </summary>
    public int Priority { get; set; } = 0;

    /// <summary>
    /// Staged-but-unpublished. While true the notification is hidden from all
    /// user-facing reads regardless of startsAt/endsAt; publishing sets real
    /// dates and flips this to false.
    /// </summary>
    public bool IsDraft { get; set; } = false;

    /// <summary>
    /// Soft delete
    /// </summary>
    public DateTime? ArchivedAt { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual ICollection<PlatformNotificationInteraction> Interactions { get; set; } = new List<PlatformNotificationInteraction>();
}
