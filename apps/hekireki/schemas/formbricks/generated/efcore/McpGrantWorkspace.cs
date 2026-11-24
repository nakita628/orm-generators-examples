#nullable enable

namespace Formbricks.Models;

/// <summary>
/// A workspace a grant with `workspaces` access reaches.
/// </summary>
public partial class McpGrantWorkspace
{
    public string GrantId { get; set; } = null!;

    public virtual McpGrant Grant { get; set; } = null!;

    public string WorkspaceId { get; set; } = null!;

    public virtual Workspace Workspace { get; set; } = null!;
}
