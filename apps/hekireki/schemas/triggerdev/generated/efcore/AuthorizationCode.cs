#nullable enable

using System;
using System.Collections.Generic;

namespace Trigger.Models;

/// <summary>
/// Used to generate PersonalAccessTokens, they're one-time use
/// </summary>
public partial class AuthorizationCode
{
    public string Id { get; set; } = null!;

    public string Code { get; set; } = null!;

    public virtual PersonalAccessToken? PersonalAccessToken { get; set; }

    public string? PersonalAccessTokenId { get; set; }

    /// <summary>
    /// Client IP of the request that minted the code, compared with the approving browser's IP. Deleted with the code once it expires.
    /// </summary>
    public string? RequesterIp { get; set; }

    /// <summary>
    /// S256 PKCE challenge from the client that minted the code. Null for codes minted by CLIs that predate it.
    /// </summary>
    public string? CodeChallenge { get; set; }

    /// <summary>
    /// Loopback port the approving browser is redirected to with the approval code. Null means the user pastes it.
    /// </summary>
    public int? CallbackPort { get; set; }

    /// <summary>
    /// sha256 of the approval code, which only the approving browser ever sees.
    /// </summary>
    public string? ApprovalCodeHash { get; set; }

    public virtual OAuthClient? OauthClient { get; set; }

    public string? OauthClientId { get; set; }

    public string? RedirectUri { get; set; }

    public List<string>? Scope { get; set; } = new List<string>();

    public string? Resource { get; set; }

    public string? ApprovedByUserId { get; set; }

    /// <summary>
    /// The user who started the OAuth authorization. Only they can approve it.
    /// </summary>
    public string? RequestedByUserId { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
