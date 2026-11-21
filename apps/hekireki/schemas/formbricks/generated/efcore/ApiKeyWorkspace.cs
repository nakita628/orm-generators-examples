#nullable enable

using System;

namespace Formbricks.Models;

/// <summary>
/// Links API keys to workspaces with specific permissions.
/// Enables granular access control for API keys across workspaces.
///
/// @property id - Unique identifier for the workspace access entry
/// @property apiKey - The associated API key
/// @property workspace - The workspace being accessed
/// @property permission - Level of access granted
/// </summary>
public partial class ApiKeyWorkspace
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public string ApiKeyId { get; set; } = null!;

    public virtual ApiKey ApiKey { get; set; } = null!;

    public virtual Workspace Workspace { get; set; } = null!;

    public string WorkspaceId { get; set; } = null!;

    public ApiKeyPermission Permission { get; set; }
}
