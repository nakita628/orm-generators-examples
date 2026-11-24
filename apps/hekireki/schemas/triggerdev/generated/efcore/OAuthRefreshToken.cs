#nullable enable

using System;
using System.Collections.Generic;

namespace Trigger.Models;

public partial class OAuthRefreshToken
{
    public string Id { get; set; } = null!;

    public string HashedToken { get; set; } = null!;

    public virtual PersonalAccessToken PersonalAccessToken { get; set; } = null!;

    public string PersonalAccessTokenId { get; set; } = null!;

    public virtual OAuthClient OauthClient { get; set; } = null!;

    public string OauthClientId { get; set; } = null!;

    public List<string>? Scope { get; set; }

    public string Resource { get; set; } = null!;

    public DateTime? UsedAt { get; set; }

    public string? ReplacedById { get; set; }

    public DateTime IdleExpiresAt { get; set; }

    public DateTime CreatedAt { get; set; }
}
