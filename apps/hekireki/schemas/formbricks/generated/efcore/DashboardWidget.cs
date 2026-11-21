#nullable enable

using System;

namespace Formbricks.Models;

/// <summary>
/// Represents a chart widget on a dashboard.
/// Widgets are positioned using a grid layout system.
///
/// @property id - Unique identifier for the widget
/// @property dashboard - The dashboard this widget belongs to
/// @property chart - The chart displayed in this widget (chart name is used for display)
/// @property layout - Grid layout configuration (x, y, width, height)
/// @property order - Display order within the dashboard
/// </summary>
public partial class DashboardWidget
{
    public string Id { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual Dashboard Dashboard { get; set; } = null!;

    public string DashboardId { get; set; } = null!;

    public virtual Chart Chart { get; set; } = null!;

    public string ChartId { get; set; } = null!;

    /// <summary>
    /// [WidgetLayout] - Grid layout: { x, y, w, h }
    /// </summary>
    public string Layout { get; set; } = "{\"x\":0,\"y\":0,\"w\":4,\"h\":3}";

    public int Order { get; set; } = 0;
}
