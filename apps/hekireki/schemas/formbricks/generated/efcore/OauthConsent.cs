#nullable enable

using System;
using System.Collections.Generic;

namespace Formbricks.Models;

/// <summary>
/// Better Auth OAuth user consent grants for delegated clients.
/// </summary>
public partial class OauthConsent
{
    public string Id { get; set; } = null!;

    public string ClientId { get; set; } = null!;

    public virtual OauthClient Client { get; set; } = null!;

    public string? UserId { get; set; }

    public virtual User? User { get; set; }

    public string? ReferenceId { get; set; }

    public List<string>? Scopes { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    /// <summary>
    /// Better Auth 1.7 (ENG-2343). /authorize compares the requested resources against these; without
    /// them every pre-existing connection is forced back through the consent screen.
    /// </summary>
    public List<string>? Resources { get; set; } = new List<string>();

    public List<string>? RequestedUserInfoClaims { get; set; } = new List<string>();
}
