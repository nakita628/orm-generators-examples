#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// A Dashboards v2 board, apart from the Analytics Reports boards in Dashboard. Boards
/// made before the split were copied here with their ids, so stars and insight
/// pointers still name them. No relation, per the database rule.
/// </summary>
public partial class DashboardsV2
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// The project that owns the board; only there can it be edited.
    /// </summary>
    public string ProjectId { get; set; } = null!;

    public string Name { get; set; } = null!;

    /// <summary>
    /// Free text under the board's name; null until a member writes one.
    /// </summary>
    public string? Description { get; set; }

    /// <summary>
    /// The member who created it; null for boards that predate authorship.
    /// </summary>
    public string? CreatedById { get; set; }

    /// <summary>
    /// Who sees it. Only the author changes it; a board with no author stays PROJECT.
    /// </summary>
    public DashboardScope Scope { get; set; } = DashboardScope.Project;

    /// <summary>
    /// The owning project's organization, stamped when the board is first set to
    /// ORGANIZATION; null before.
    /// </summary>
    public string? OrganizationId { get; set; }

    /// <summary>
    /// The product whose sidebar lists it: llm-ops, me, gateway or governance (ADR-002,
    /// decision 8). Every board made before homes existed is llm-ops.
    /// </summary>
    public string HomeProduct { get; set; } = "llm-ops";

    /// <summary>
    /// Deprecated: unused since each widget query names its own sources (dashboards-v2 AC200,
    /// 2026-10-28); dropped in a later migration, once no serving release reads it.
    /// </summary>
    public string? ExcludedOrigins { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
