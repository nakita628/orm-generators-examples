#nullable enable

using System;
using System.Collections.Generic;

namespace Formbricks.Models;

/// <summary>
/// Represents pending invitations to join an organization.
/// Used to manage the process of adding new users to an organization.
/// Once accepted, invites are converted into memberships.
///
/// @property id - Unique identifier for the invite
/// @property email - Email address of the invited user
/// @property name - Optional display name for the invited user
/// @property organization - The organization sending the invite
/// @property creator - The user who created the invite
/// @property acceptor - The user who accepted the invite (if accepted)
/// @property expiresAt - When the invite becomes invalid
/// @property role - Intended role for the invited user
/// @property teamIds - Teams the user will be added to upon acceptance
/// </summary>
public partial class Invite
{
    public string Id { get; set; } = null!;

    public string Email { get; set; } = null!;

    public string? Name { get; set; }

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual User Creator { get; set; } = null!;

    public string CreatorId { get; set; } = null!;

    public virtual User? Acceptor { get; set; }

    public string? AcceptorId { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime ExpiresAt { get; set; }

    public OrganizationRole Role { get; set; } = OrganizationRole.Member;

    public List<string>? TeamIds { get; set; } = new List<string>();
}
