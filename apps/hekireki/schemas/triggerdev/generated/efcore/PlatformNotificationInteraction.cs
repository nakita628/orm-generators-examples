#nullable enable

using System;

namespace Trigger.Models;

/// <summary>
/// Per-user tracking of notification views and dismissals
/// </summary>
public partial class PlatformNotificationInteraction
{
    public string Id { get; set; } = null!;

    public virtual PlatformNotification Notification { get; set; } = null!;

    public string NotificationId { get; set; } = null!;

    public virtual User User { get; set; } = null!;

    public string UserId { get; set; } = null!;

    /// <summary>
    /// Set by beacon/CLI GET on first impression
    /// </summary>
    public DateTime FirstSeenAt { get; set; }

    /// <summary>
    /// Times shown (incremented per beacon/CLI view)
    /// </summary>
    public int ShowCount { get; set; } = 1;

    /// <summary>
    /// User dismissed in webapp
    /// </summary>
    public DateTime? WebappDismissedAt { get; set; }

    /// <summary>
    /// User clicked a link in webapp
    /// </summary>
    public DateTime? WebappClickedAt { get; set; }

    /// <summary>
    /// Auto-dismissed or expired in CLI
    /// </summary>
    public DateTime? CliDismissedAt { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
