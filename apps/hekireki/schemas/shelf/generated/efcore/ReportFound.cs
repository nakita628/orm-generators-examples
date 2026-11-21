#nullable enable

using System;

namespace Shelf.Models;

public partial class ReportFound
{
    public string Id { get; set; } = null!;

    public string Email { get; set; } = null!;

    public string Content { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual Asset? Asset { get; set; }

    public string? AssetId { get; set; }

    public virtual Kit? Kit { get; set; }

    public string? KitId { get; set; }
}
