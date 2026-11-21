#nullable enable

using System;

namespace Trigger.Models;

/// <summary>
/// A user-defined metrics dashboard
/// </summary>
public partial class MetricsDashboard
{
    public string Id { get; set; } = null!;

    public string FriendlyId { get; set; } = null!;

    public string Title { get; set; } = null!;

    public string Description { get; set; } = "";

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual Project? Project { get; set; }

    public string? ProjectId { get; set; }

    /// <summary>
    /// Who created the dashboard
    /// </summary>
    public virtual User? Owner { get; set; }

    public string? OwnerId { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    /// <summary>
    /// JSON that defines the config, queries, layout and config of all widgets.
    /// There will be a version field for the format.
    /// </summary>
    public string Layout { get; set; } = null!;
}
