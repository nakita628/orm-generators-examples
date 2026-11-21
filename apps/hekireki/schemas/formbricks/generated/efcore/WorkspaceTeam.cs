#nullable enable

using System;

namespace Formbricks.Models;

/// <summary>
/// Defines team access to specific workspaces.
/// Manages workspace-level permissions for teams.
///
/// @property workspace - The accessed workspace
/// @property team - The team receiving access
/// @property permission - Level of access granted
/// </summary>
public partial class WorkspaceTeam
{
    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public string WorkspaceId { get; set; } = null!;

    public virtual Workspace Workspace { get; set; } = null!;

    public string TeamId { get; set; } = null!;

    public virtual Team Team { get; set; } = null!;

    public WorkspaceTeamPermission Permission { get; set; } = WorkspaceTeamPermission.Read;
}
