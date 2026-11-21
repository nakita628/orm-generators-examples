#nullable enable

using System;

namespace Trigger.Models;

public partial class OrganizationAccessToken
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// User-provided name for the token
    /// </summary>
    public string Name { get; set; } = null!;

    /// <summary>
    /// Used to differentiate between user-generated and system-generated tokens
    /// </summary>
    public OrganizationAccessTokenType Type { get; set; } = OrganizationAccessTokenType.User;

    /// <summary>
    /// This is used to find the token in the database
    /// </summary>
    public string HashedToken { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    /// <summary>
    /// Optional expiration date for the token
    /// </summary>
    public DateTime? ExpiresAt { get; set; }

    public DateTime? RevokedAt { get; set; }

    public DateTime? LastAccessedAt { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
