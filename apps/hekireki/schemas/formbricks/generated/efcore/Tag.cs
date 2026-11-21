#nullable enable

using System;
using System.Collections.Generic;

namespace Formbricks.Models;

/// <summary>
/// Represents a label that can be applied to survey responses.
/// Used for categorizing and organizing responses within a workspace.
///
/// @property id - Unique identifier for the tag
/// @property name - Display name of the tag
/// @property responses - Survey responses tagged with this label
/// @property workspace - The workspace this tag belongs to
/// </summary>
public partial class Tag
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public string Name { get; set; } = null!;

    public virtual ICollection<TagsOnResponses> Responses { get; set; } = new List<TagsOnResponses>();

    public virtual Workspace Workspace { get; set; } = null!;

    public string WorkspaceId { get; set; } = null!;
}
