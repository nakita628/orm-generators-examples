#nullable enable

using System;

namespace Trigger.Models;

/// <summary>
/// Records of previously-valid API keys that are still accepted for authentication
/// during a grace window after rotation. Extend or end the grace period by updating `expiresAt`.
/// </summary>
public partial class RevokedApiKey
{
    public string Id { get; set; } = null!;

    public string ApiKey { get; set; } = null!;

    public virtual RuntimeEnvironment RuntimeEnvironment { get; set; } = null!;

    public string RuntimeEnvironmentId { get; set; } = null!;

    public DateTime ExpiresAt { get; set; }

    public DateTime CreatedAt { get; set; }
}
