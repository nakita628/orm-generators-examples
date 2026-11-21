#nullable enable

using System;
using System.Collections.Generic;

namespace Formbricks.Models;

/// <summary>
/// Better Auth OAuth refresh tokens used by delegated MCP clients. The token column stores a hashed lookup value.
/// </summary>
public partial class OauthRefreshToken
{
    public string Id { get; set; } = null!;

    public string Token { get; set; } = null!;

    public string ClientId { get; set; } = null!;

    public virtual OauthClient Client { get; set; } = null!;

    public string? SessionId { get; set; }

    public virtual Session? Session { get; set; }

    public string UserId { get; set; } = null!;

    public virtual User User { get; set; } = null!;

    public string? ReferenceId { get; set; }

    public DateTime ExpiresAt { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime? Revoked { get; set; }

    public DateTime? AuthTime { get; set; }

    public List<string>? Scopes { get; set; }

    public virtual ICollection<OauthAccessToken> AccessTokens { get; set; } = new List<OauthAccessToken>();

    /// <summary>
    /// Better Auth 1.7 (ENG-2343). `resources` carries the full original grant so a refresh can narrow
    /// the resource set but never widen it.
    /// </summary>
    public string? AuthorizationCodeId { get; set; }

    public List<string>? Resources { get; set; } = new List<string>();

    public List<string>? RequestedUserInfoClaims { get; set; } = new List<string>();

    public DateTime? RotatedAt { get; set; }

    public string? RotationReplayResponse { get; set; }

    public DateTime? RotationReplayExpiresAt { get; set; }

    public string? Confirmation { get; set; }
}
