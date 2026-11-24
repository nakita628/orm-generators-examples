#nullable enable

namespace Formbricks.Models;

/// <summary>
/// An organization a grant with `organizations` access reaches, workspaces created there later included.
/// </summary>
public partial class McpGrantOrganization
{
    public string GrantId { get; set; } = null!;

    public virtual McpGrant Grant { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;
}
