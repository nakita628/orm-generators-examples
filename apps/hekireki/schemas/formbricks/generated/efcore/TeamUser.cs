#nullable enable

using System;

namespace Formbricks.Models;

/// <summary>
/// Links users to teams with specific roles.
/// Manages team membership and permissions.
///
/// @property team - The associated team
/// @property user - The team member
/// @property role - User's role within the team
/// </summary>
public partial class TeamUser
{
    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public string TeamId { get; set; } = null!;

    public virtual Team Team { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public virtual User User { get; set; } = null!;

    public TeamUserRole Role { get; set; }
}
