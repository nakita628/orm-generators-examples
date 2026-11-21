#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// One browser a person subscribed to Web Push, owned by notification. A
/// person holds one row per device; the endpoint names the browser, so a
/// browser that subscribes again replaces its own row.
/// </summary>
public partial class WebPushSubscription
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// The person the pushes are for. A plain column: the row leaves with the
    /// person through notification's own subscriber on the user lifecycle.
    /// </summary>
    public string UserId { get; set; } = null!;

    /// <summary>
    /// The push service URL the browser handed out. Unique per browser.
    /// </summary>
    public string Endpoint { get; set; } = null!;

    /// <summary>
    /// The browser's P-256 public key and auth secret, base64url, as the
    /// browser reports them. Payloads are encrypted to these.
    /// </summary>
    public string P256dh { get; set; } = null!;

    public string Auth { get; set; } = null!;

    public string? UserAgent { get; set; }

    public DateTime CreatedAt { get; set; }

    /// <summary>
    /// When the push service last accepted a push for this browser.
    /// </summary>
    public DateTime? LastSuccessAt { get; set; }
}
