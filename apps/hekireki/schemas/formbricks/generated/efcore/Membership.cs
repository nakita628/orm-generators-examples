#nullable enable

namespace Formbricks.Models;

/// <summary>
/// Links users to organizations with specific roles.
/// Manages organization membership and permissions.
/// Core model for managing user access within organizations.
///
/// @property organization - The organization the user belongs to
/// @property user - The member user
/// @property accepted - Whether the user has accepted the membership
/// @property role - User's role within the organization (owner, manager, member, billing)
/// </summary>
public partial class Membership
{
    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual User User { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public bool Accepted { get; set; } = false;

    public OrganizationRole Role { get; set; } = OrganizationRole.Member;
}
