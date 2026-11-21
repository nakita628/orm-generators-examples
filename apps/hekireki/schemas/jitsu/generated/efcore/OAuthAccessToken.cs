#nullable enable

using System;

namespace Jitsu.Models;

public partial class OAuthAccessToken
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// @zod.custom(z.coerce.date())
    /// </summary>
    public DateTime CreatedAt { get; set; }

    /// <summary>
    /// @zod.custom(z.coerce.date())
    /// </summary>
    public DateTime? LastUsed { get; set; }

    public string Hash { get; set; } = null!;

    /// <summary>
    /// @zod.custom(z.coerce.date())
    /// </summary>
    public DateTime ExpiresAt { get; set; }

    public string RefreshTokenId { get; set; } = null!;

    public virtual UserApiToken RefreshToken { get; set; } = null!;
}
