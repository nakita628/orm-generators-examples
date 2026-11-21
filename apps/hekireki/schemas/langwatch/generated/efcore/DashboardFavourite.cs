#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// A member's star, in their own order within a project: a board or a From
/// LangWatch template. Personal: one member's stars never show to another. No
/// foreign key (per the database rule): the dashboard service removes a board's
/// favourites when it deletes the board, so a star never outlives its dashboard.
/// </summary>
public partial class DashboardFavourite
{
    public string Id { get; set; } = null!;

    public string UserId { get; set; } = null!;

    /// <summary>
    /// A DashboardsV2 board's id (copied boards kept their Dashboard id); null when
    /// the star is a template.
    /// </summary>
    public string? DashboardId { get; set; }

    /// <summary>
    /// The starred From LangWatch template id; null when the star is a board.
    /// </summary>
    public string? TemplateId { get; set; }

    public string ProjectId { get; set; } = null!;

    /// <summary>
    /// The member's drag order among their own stars in the project.
    /// </summary>
    public int Position { get; set; }

    public DateTime CreatedAt { get; set; }
}
