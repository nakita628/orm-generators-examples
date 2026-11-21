#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// D08 — who the directory means, per connection.
///
/// The identity provider's own identifier for a person, which is what
/// survives their email changing: an address is not identity, people change
/// theirs, and the same person carries different identifiers on two different
/// connections. So the PAIR is the key and `externalId` alone never is.
///
/// No relation to `SsoConnection` (that projection carries no foreign keys);
/// the `User` relation is real, so a deleted user takes their directory
/// identities with them rather than stranding a mapping that resolves to
/// nobody.
/// </summary>
public partial class ScimExternalId
{
    public string OrganizationId { get; set; } = null!;

    public string Id { get; set; } = null!;

    public string ConnectionId { get; set; } = null!;

    public string ExternalId { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public virtual User User { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
