#nullable enable

using System;
using System.Collections.Generic;

namespace Formbricks.Models;

/// <summary>
/// Better Auth opaque access token records. The token column stores Better Auth's hashed lookup value.
/// JWT access tokens are verified through JWKS and bounded by the configured short lifetime.
/// </summary>
public partial class OauthAccessToken
{
    public string Id { get; set; } = null!;

    public string Token { get; set; } = null!;

    public string ClientId { get; set; } = null!;

    public virtual OauthClient Client { get; set; } = null!;

    public string? SessionId { get; set; }

    public virtual Session? Session { get; set; }

    public string? UserId { get; set; }

    public virtual User? User { get; set; }

    public string? ReferenceId { get; set; }

    public string? RefreshId { get; set; }

    public virtual OauthRefreshToken? Refresh { get; set; }

    public DateTime ExpiresAt { get; set; }

    public DateTime CreatedAt { get; set; }

    public List<string>? Scopes { get; set; }

    /// <summary>
    /// Better Auth 1.7 (ENG-2343). `resources` is the RFC 8707 resource set the grant approved — the
    /// binding this whole upgrade exists to get (GHSA-p2fr-6hmx-4528).
    /// </summary>
    public string? AuthorizationCodeId { get; set; }

    public List<string>? Resources { get; set; } = new List<string>();

    public List<string>? RequestedUserInfoClaims { get; set; } = new List<string>();

    public DateTime? Revoked { get; set; }

    public string? Confirmation { get; set; }
}
