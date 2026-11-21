#nullable enable

using System;
using System.Collections.Generic;

namespace Trigger.Models;

public partial class ApiKey
{
    public string Id { get; set; } = null!;

    public string Name { get; set; } = null!;

    public string KeyHash { get; set; } = null!;

    public string LastFour { get; set; } = null!;

    /// <summary>
    /// Optional one-line note the creator adds to say what the key is for.
    /// </summary>
    public string? Description { get; set; }

    public virtual RuntimeEnvironment RuntimeEnvironment { get; set; } = null!;

    public string RuntimeEnvironmentId { get; set; } = null!;

    public virtual User? CreatedBy { get; set; }

    public string? CreatedByUserId { get; set; }

    public ApiKeyIssuer IssuedBy { get; set; } = ApiKeyIssuer.User;

    public string? PresetId { get; set; }

    public List<string>? Scopes { get; set; }

    /// <summary>
    /// CIDR ranges this key may be used from. Empty means any IP.
    /// </summary>
    public List<string>? AllowedIpRanges { get; set; } = new List<string>();

    public DateTime? LastUsedAt { get; set; }

    public string? LastUsedIp { get; set; }

    public DateTime? RevokedAt { get; set; }

    public DateTime? ExpiresAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public DateTime CreatedAt { get; set; }
}
