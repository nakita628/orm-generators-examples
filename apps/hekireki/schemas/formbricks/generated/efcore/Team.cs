#nullable enable

using System;
using System.Collections.Generic;

namespace Formbricks.Models;

/// <summary>
/// Represents a team within an organization.
/// Enables group-based access control and collaboration.
///
/// @property id - Unique identifier for the team
/// @property name - Display name of the team
/// @property organization - The parent organization
/// @property teamUsers - Users who are part of this team
/// @property workspaceTeams - Workspaces this team has access to
/// </summary>
public partial class Team
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public string Name { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    public virtual ICollection<TeamUser> TeamUsers { get; set; } = new List<TeamUser>();

    public virtual ICollection<WorkspaceTeam> WorkspaceTeams { get; set; } = new List<WorkspaceTeam>();
}
