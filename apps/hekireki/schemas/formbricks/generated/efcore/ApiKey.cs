#nullable enable

using System;
using System.Collections.Generic;

namespace Formbricks.Models;

/// <summary>
/// Represents enhanced API authentication keys with organization-level ownership.
/// Used for authenticating API requests to Formbricks with more granular permissions.
///
/// @property id - Unique identifier for the API key
/// @property label - Optional descriptive name for the key
/// @property hashedKey - Securely stored API key
/// @property organization - The organization this key belongs to
/// @property createdBy - User ID who created this key
/// @property lastUsedAt - Timestamp of last usage
/// @property apiKeyWorkspaces - Workspaces this key has access to
/// </summary>
public partial class ApiKey
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public string? CreatedBy { get; set; }

    public DateTime? LastUsedAt { get; set; }

    public string Label { get; set; } = null!;

    public string HashedKey { get; set; } = null!;

    public string? LookupHash { get; set; }

    public string OrganizationId { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    public virtual ICollection<ApiKeyWorkspace> ApiKeyWorkspaces { get; set; } = new List<ApiKeyWorkspace>();

    /// <summary>
    /// [OrganizationAccess]
    /// </summary>
    public string OrganizationAccess { get; set; } = "{}";
}
