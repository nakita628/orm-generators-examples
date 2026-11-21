#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// A Dashboards v2 widget: a srcdoc definition placed on a DashboardsV2 board, or
/// unplaced. Widgets made before the split were copied from CustomGraph with their
/// ids. No relation, per the database rule.
/// </summary>
public partial class DashboardWidget
{
    public string Id { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    /// <summary>
    /// The DashboardsV2 board it sits on; null while unplaced.
    /// </summary>
    public string? DashboardId { get; set; }

    public string Name { get; set; } = null!;

    /// <summary>
    /// The widget's srcdoc definition.
    /// </summary>
    public string Definition { get; set; } = null!;

    public int GridColumn { get; set; } = 0;

    public int GridRow { get; set; } = 0;

    public int ColSpan { get; set; } = 1;

    public int RowSpan { get; set; } = 1;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
