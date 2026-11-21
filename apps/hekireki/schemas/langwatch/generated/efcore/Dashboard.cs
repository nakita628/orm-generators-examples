#nullable enable

using System;
using System.Collections.Generic;

namespace LangWatch.Models;

public partial class Dashboard
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// The project that owns the board; only there can it be edited.
    /// </summary>
    public string ProjectId { get; set; } = null!;

    public virtual Project Project { get; set; } = null!;

    public string Name { get; set; } = null!;

    /// <summary>
    /// Legacy board sort position. The Dashboards v2 app orders boards by a
    /// member's favourites and never writes this; kept because the legacy
    /// Analytics reports section and the REST reorder endpoint still write it,
    /// and a rolling deploy's older image still reads it.
    /// </summary>
    public int Order { get; set; } = 0;

    /// <summary>
    /// Free text under the board's name; null until a member writes one.
    /// Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
    /// </summary>
    public string? Description { get; set; }

    /// <summary>
    /// The member who created it; null for boards that predate the column.
    /// Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
    /// </summary>
    public string? CreatedById { get; set; }

    /// <summary>
    /// Who sees it. Only the author changes it; a board with no author stays PROJECT.
    /// Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
    /// </summary>
    public DashboardScope Scope { get; set; } = DashboardScope.Project;

    /// <summary>
    /// The owning project's organization, stamped when the board is first set to
    /// ORGANIZATION; null before. No relation, per the database rule.
    /// Retiring: Dashboards v2 boards live in DashboardsV2, so no new write sets it.
    /// </summary>
    public string? OrganizationId { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual ICollection<CustomGraph> Graphs { get; set; } = new List<CustomGraph>();
}
