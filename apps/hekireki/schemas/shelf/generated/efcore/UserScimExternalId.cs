#nullable enable

using System;

namespace Shelf.Models;

/// <summary>
/// *
/// * Org-scoped SCIM external ID linking a Shelf user to an identity-provider
/// * object within a specific organization.  A single Shelf user can appear in
/// * multiple organizations, each managed by a different identity provider, so the
/// * external ID must be scoped to the (user, organization) pair.
/// </summary>
public partial class UserScimExternalId
{
    public string Id { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public string ScimExternalId { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual User User { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;
}
