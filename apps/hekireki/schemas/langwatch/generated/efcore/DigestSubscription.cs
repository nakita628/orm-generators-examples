#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// A person's opt-out of the digest. No row, or no `unsubscribedAt`: subscribed.
/// </summary>
public partial class DigestSubscription
{
    public string UserId { get; set; } = null!;

    public DateTime? UnsubscribedAt { get; set; }

    /// <summary>
    /// The send whose unsubscribe link or one-click header the person used.
    /// </summary>
    public string? UnsubscribedFromSendId { get; set; }

    public DateTime UpdatedAt { get; set; }
}
