#nullable enable

using System;
using System.Collections.Generic;

namespace Formbricks.Models;

/// <summary>
/// Stores active authentication sessions for revocable server-side login state.
///
/// @property sessionToken - Opaque token stored in the browser cookie
/// @property user - The Formbricks user who owns this session
/// @property expires - Hard expiry for the session
/// </summary>
public partial class Session
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public string SessionToken { get; set; } = null!;

    public virtual User User { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public DateTime Expires { get; set; }

    public string? IpAddress { get; set; }

    public string? UserAgent { get; set; }

    public virtual ICollection<OauthAccessToken> OauthAccessTokens { get; set; } = new List<OauthAccessToken>();

    public virtual ICollection<OauthRefreshToken> OauthRefreshTokens { get; set; } = new List<OauthRefreshToken>();
}
