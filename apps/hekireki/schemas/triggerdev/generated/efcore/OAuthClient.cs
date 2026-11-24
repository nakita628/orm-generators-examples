#nullable enable

using System;
using System.Collections.Generic;

namespace Trigger.Models;

public partial class OAuthClient
{
    public string Id { get; set; } = null!;

    public OAuthClientKind Kind { get; set; }

    public string ClientId { get; set; } = null!;

    public string? MetadataUrl { get; set; }

    public string? DisplayName { get; set; }

    public string? SelfReportedName { get; set; }

    public List<string>? RedirectUris { get; set; }

    public string? DedupeKey { get; set; }

    public DateTime? MetadataExpiresAt { get; set; }

    public DateTime? LastUsedAt { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual ICollection<AuthorizationCode> AuthorizationCodes { get; set; } = new List<AuthorizationCode>();

    public virtual ICollection<PersonalAccessToken> PersonalAccessTokens { get; set; } = new List<PersonalAccessToken>();

    public virtual ICollection<OAuthRefreshToken> RefreshTokens { get; set; } = new List<OAuthRefreshToken>();
}
