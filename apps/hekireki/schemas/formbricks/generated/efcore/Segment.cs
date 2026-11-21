#nullable enable

using System;
using System.Collections.Generic;

namespace Formbricks.Models;

/// <summary>
/// Defines a segment of contacts based on attributes.
/// Used for targeting surveys to specific user groups.
///
/// @property id - Unique identifier for the segment
/// @property title - Display name of the segment
/// @property filters - Rules defining the segment
/// @property isPrivate - Whether the segment is private
/// @property workspace - The workspace this segment belongs to
/// </summary>
public partial class Segment
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public string Title { get; set; } = null!;

    public string? Description { get; set; }

    public bool IsPrivate { get; set; } = true;

    /// <summary>
    /// [SegmentFilter]
    /// </summary>
    public string Filters { get; set; } = "[]";

    public virtual Workspace Workspace { get; set; } = null!;

    public string WorkspaceId { get; set; } = null!;

    public virtual ICollection<Survey> Surveys { get; set; } = new List<Survey>();
}
