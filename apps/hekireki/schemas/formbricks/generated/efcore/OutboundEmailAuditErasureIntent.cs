#nullable enable

using System;

namespace Formbricks.Models;

/// <summary>
/// Holds an account erasure request across an in-flight email send. The address is never stored.
/// The user id is removed once deletion is confirmed; one-way recovery keys expire after 90 days.
/// </summary>
public partial class OutboundEmailAuditErasureIntent
{
    public string Id { get; set; } = null!;

    public string? UserId { get; set; }

    public string UserHash { get; set; } = null!;

    public string RecipientMac { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime? ActivatedAt { get; set; }

    public DateTime? CheckedAt { get; set; }

    public DateTime ExpiresAt { get; set; }
}
