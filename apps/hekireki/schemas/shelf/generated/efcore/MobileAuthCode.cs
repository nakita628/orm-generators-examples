#nullable enable

using System;

namespace Shelf.Models;

/// <summary>
/// Native-app SSO web-delegated login — single-use session-handoff codes.
/// @see apps/webapp/app/routes/_auth+/oauth.callback.mobile.tsx
/// </summary>
public partial class MobileAuthCode
{
    public string Id { get; set; } = null!;

    public virtual User User { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public string CodeHash { get; set; } = null!;

    public string? CodeChallenge { get; set; }

    public string? SessionCiphertext { get; set; }

    public DateTime? ConsumedAt { get; set; }

    public DateTime ExpiresAt { get; set; }

    public DateTime CreatedAt { get; set; }
}
