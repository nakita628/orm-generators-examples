#nullable enable

using System;

namespace Shelf.Models;

public partial class RoleChangeLog
{
    public string Id { get; set; } = null!;

    public OrganizationRoles PreviousRole { get; set; }

    public OrganizationRoles NewRole { get; set; }

    /// <summary>
    /// How the change was made. SSO rows name the member themself as `changedBy`.
    /// </summary>
    public RoleChangeSource Source { get; set; } = RoleChangeSource.Manual;

    public DateTime CreatedAt { get; set; }

    public virtual User User { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public virtual User ChangedBy { get; set; } = null!;

    public string ChangedById { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;
}
