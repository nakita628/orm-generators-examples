#nullable enable

using System;
using System.Collections.Generic;

namespace Formbricks.Models;

/// <summary>
/// Better Auth OAuth 2.1 client registrations for MCP and future OAuth clients.
/// </summary>
public partial class OauthClient
{
    public string Id { get; set; } = null!;

    public string ClientId { get; set; } = null!;

    public string? ClientSecret { get; set; }

    public bool? Disabled { get; set; } = false;

    public bool? SkipConsent { get; set; }

    public bool? EnableEndSession { get; set; }

    public string? SubjectType { get; set; }

    public List<string>? Scopes { get; set; } = new List<string>();

    public string? UserId { get; set; }

    public virtual User? User { get; set; }

    public DateTime? CreatedAt { get; set; }

    public DateTime? UpdatedAt { get; set; }

    public string? Name { get; set; }

    public string? Uri { get; set; }

    public string? Icon { get; set; }

    public List<string>? Contacts { get; set; } = new List<string>();

    public string? Tos { get; set; }

    public string? Policy { get; set; }

    public string? SoftwareId { get; set; }

    public string? SoftwareVersion { get; set; }

    public string? SoftwareStatement { get; set; }

    public List<string>? RedirectUris { get; set; }

    public List<string>? PostLogoutRedirectUris { get; set; } = new List<string>();

    public string? TokenEndpointAuthMethod { get; set; }

    public List<string>? GrantTypes { get; set; } = new List<string>();

    public List<string>? ResponseTypes { get; set; } = new List<string>();

    /// <summary>
    /// Unwritten by Better Auth 1.7 (it dropped both from its schema), kept deliberately: they are
    /// nullable, so keeping them makes the 1.7 migration purely additive and a rollback to 1.6 needs no
    /// down-migration. Drop them in a separate cleanup after a bake period.
    /// </summary>
    public bool? Public { get; set; }

    public string? Type { get; set; }

    public bool? RequirePKCE { get; set; }

    public string? ReferenceId { get; set; }

    public string? Metadata { get; set; }

    /// <summary>
    /// Better Auth 1.7 (ENG-2343).
    /// </summary>
    public string? ClientDiscoveryId { get; set; }

    public List<string>? ClientCredentialsScopes { get; set; } = new List<string>();

    public string? BackchannelLogoutUri { get; set; }

    public bool? BackchannelLogoutSessionRequired { get; set; }

    public string? ApplicationType { get; set; }

    public string? Jwks { get; set; }

    public string? JwksUri { get; set; }

    public bool? DpopBoundAccessTokens { get; set; } = false;

    public virtual ICollection<OauthAccessToken> AccessTokens { get; set; } = new List<OauthAccessToken>();

    public virtual ICollection<OauthRefreshToken> RefreshTokens { get; set; } = new List<OauthRefreshToken>();

    public virtual ICollection<OauthConsent> Consents { get; set; } = new List<OauthConsent>();

    public virtual ICollection<OauthClientResource> Resources { get; set; } = new List<OauthClientResource>();

    public virtual ICollection<McpGrant> McpGrants { get; set; } = new List<McpGrant>();
}
