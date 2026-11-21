#nullable enable

using System;

namespace Formbricks.Models;

/// <summary>
/// Better Auth 1.7 (ENG-2343): links a client to the resources it may request tokens for. With
/// `enforcePerClientResources` (the default) a missing row means `invalid_target` at the token
/// endpoint, which is why the migration backfills one row per existing client.
/// Both sides reference natural keys, matching the other oauth child tables.
/// </summary>
public partial class OauthClientResource
{
    public string Id { get; set; } = null!;

    public string ClientId { get; set; } = null!;

    public virtual OauthClient Client { get; set; } = null!;

    public string ResourceId { get; set; } = null!;

    public virtual OauthResource Resource { get; set; } = null!;

    public string? Metadata { get; set; }

    public DateTime? CreatedAt { get; set; }
}
