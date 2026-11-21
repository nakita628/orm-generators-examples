#nullable enable

using System;
using System.Collections.Generic;

namespace Jitsu.Models;

public partial class OAuthClient
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// @zod.custom(z.coerce.date())
    /// </summary>
    public DateTime CreatedAt { get; set; }

    /// <summary>
    /// @zod.custom(z.coerce.date())
    /// </summary>
    public DateTime UpdatedAt { get; set; }

    public string ClientSecretHash { get; set; } = null!;

    public string Name { get; set; } = null!;

    public List<string>? RedirectUris { get; set; }

    public virtual ICollection<UserApiToken> Tokens { get; set; } = new List<UserApiToken>();
}
