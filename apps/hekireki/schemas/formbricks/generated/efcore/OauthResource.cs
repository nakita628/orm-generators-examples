#nullable enable

using System;
using System.Collections.Generic;

namespace Formbricks.Models;

/// <summary>
/// Better Auth 1.7 (ENG-2343): a protected resource the authorization server issues tokens for, with
/// its own token lifetimes, scope allow-list and claims. Replaces the flat `validAudiences` option —
/// a token is now bound to the resource it was approved for instead of to a static allow-list.
/// `identifier` is the RFC 8707 `resource` value and is the natural key other rows point at.
/// </summary>
public partial class OauthResource
{
    public string Id { get; set; } = null!;

    public string Identifier { get; set; } = null!;

    public string Name { get; set; } = null!;

    public int? AccessTokenTtl { get; set; }

    public int? RefreshTokenTtl { get; set; }

    public string? SigningAlgorithm { get; set; }

    public string? SigningKeyId { get; set; }

    public List<string>? AllowedScopes { get; set; } = new List<string>();

    public string? CustomClaims { get; set; }

    public bool? DpopBoundAccessTokensRequired { get; set; } = false;

    public bool? Disabled { get; set; } = false;

    public DateTime? CreatedAt { get; set; }

    public DateTime? UpdatedAt { get; set; }

    public int? PolicyVersion { get; set; } = 1;

    public string? Metadata { get; set; }

    public virtual ICollection<OauthClientResource> Clients { get; set; } = new List<OauthClientResource>();
}
