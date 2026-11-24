#nullable enable

using System;

namespace Formbricks.Models;

/// <summary>
/// A workspace's display name for an imported (embedded-data) key of its Feedback Datasets, shown in
/// place of the key in chart legends, headers and the chart builder's field pickers. One per key and
/// workspace; clearing the name deletes the row, and the key shows as collected again. Built-in fields
/// are never renamed here: they keep their translated names.
///
/// @property key - The embedded-data key exactly as collected, case included
/// @property label - The name shown for the key across the workspace
/// </summary>
public partial class EmbeddedFieldLabel
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual Workspace Workspace { get; set; } = null!;

    public string WorkspaceId { get; set; } = null!;

    public string Key { get; set; } = null!;

    public string Label { get; set; } = null!;
}
