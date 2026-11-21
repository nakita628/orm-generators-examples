#nullable enable

using System;
using System.Collections.Generic;

namespace Formbricks.Models;

/// <summary>
/// Represents a chart/visualization that can be used in multiple dashboards.
/// Charts are reusable components that query analytics data.
///
/// @property id - Unique identifier for the chart
/// @property name - Display name of the chart
/// @property type - Type of visualization (bar, line, pie, etc.)
/// @property workspace - The workspace this chart belongs to
/// @property query - Cube.js query configuration (JSON)
/// @property config - Chart-specific configuration (colors, labels, etc.)
/// @property createdBy - User who created the chart
/// @property dashboards - Dashboards that use this chart
/// </summary>
public partial class Chart
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public string Name { get; set; } = null!;

    public ChartType Type { get; set; }

    public virtual Workspace Workspace { get; set; } = null!;

    public string WorkspaceId { get; set; } = null!;

    /// <summary>
    /// [ChartQuery] - Cube.js query configuration
    /// </summary>
    public string Query { get; set; } = "{}";

    /// <summary>
    /// [ChartConfig] - Visualization configuration (colors, labels, formatting)
    /// </summary>
    public string Config { get; set; } = "{}";

    public virtual User? Creator { get; set; }

    public string? CreatedBy { get; set; }

    public virtual FeedbackDirectory FeedbackDirectory { get; set; } = null!;

    public string FeedbackDirectoryId { get; set; } = null!;

    public virtual ICollection<DashboardWidget> Widgets { get; set; } = new List<DashboardWidget>();
}
