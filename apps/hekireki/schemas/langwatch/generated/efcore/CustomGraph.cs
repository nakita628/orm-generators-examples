#nullable enable

using System;
using System.Collections.Generic;

namespace LangWatch.Models;

public partial class CustomGraph
{
    public string Id { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    public virtual Project Project { get; set; } = null!;

    public string Name { get; set; } = null!;

    /// <summary>
    /// The chart's definition, in the shape its `kind` names.
    /// </summary>
    public string Graph { get; set; } = null!;

    public string? Filters { get; set; }

    /// <summary>
    /// Discriminator for the chart shapes that share this table:
    /// - "builder" (default) - the chart builder's own `graph` payload
    /// - "workbench_sql"     - a saved LangWatchQL workbench chart, whose
    /// `graph` holds the versioned { sql, parameters, vegaLiteSpec }
    /// definition (see server/analytics/saved-workbench-charts)
    /// - "dashboard_srcdoc"  - frozen legacy copies of Dashboards v2 widgets, which
    /// now live in DashboardWidget; nothing new writes this kind
    /// Each reader filters by its own kind, so neither sees the other's rows and
    /// a builder chart can never be interpreted as a workbench definition.
    /// </summary>
    public string Kind { get; set; } = "builder";

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public string? DashboardId { get; set; }

    public virtual Dashboard? Dashboard { get; set; }

    public int GridColumn { get; set; } = 0;

    public int GridRow { get; set; } = 0;

    public int ColSpan { get; set; } = 1;

    public int RowSpan { get; set; } = 1;

    public virtual Trigger? Trigger { get; set; }

    public virtual ICollection<TriggerSent> TriggerSent { get; set; } = new List<TriggerSent>();
}
