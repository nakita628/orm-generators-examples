#nullable enable

using System;

namespace Shelf.Models;

public partial class AssetLocation
{
    public string Id { get; set; } = null!;

    public virtual Asset Asset { get; set; } = null!;

    public string AssetId { get; set; } = null!;

    public virtual Location Location { get; set; } = null!;

    public string LocationId { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public int Quantity { get; set; } = 1;

    public virtual AssetKit? AssetKit { get; set; }

    public string? AssetKitId { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
