#nullable enable

using System;

namespace LangWatch.Models;

public partial class SavedView
{
    public string Id { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    public virtual Project Project { get; set; } = null!;

    public string? UserId { get; set; }

    public virtual User? User { get; set; }

    public string Name { get; set; } = null!;

    public string Filters { get; set; } = "{}";

    public string? Query { get; set; }

    public string? Period { get; set; }

    public int Order { get; set; } = 0;

    /// <summary>
    /// Discriminator for the two storage shapes that share this table:
    /// - "v1-traces-filter" (default) - legacy /messages page filter map saved views
    /// - "v2-traces-lens"             - new traces v2 lens configs (columns, addons, grouping, sort, filterText) encoded inside `filters` JSON
    /// Each client filters by its own kind on read so neither sees the other's rows.
    /// </summary>
    public string Kind { get; set; } = "v1-traces-filter";

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
