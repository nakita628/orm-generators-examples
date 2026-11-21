#nullable enable

using System;
using System.Collections.Generic;

namespace Shelf.Models;

/// <summary>
/// Pivot linking an `Asset` to the `Kit` it belongs to.
///
/// `@@unique([assetId])` enforces "at most one kit per asset" today.
/// The structure already allows multi-row-per-asset; if multi-kit
/// allocation lands later, drop the `assetId` unique constraint and
/// add a type-aware single-row trigger + sum-within-total CONSTRAINT
/// TRIGGER (parallel to the Custody / BookingAsset pivots).
/// </summary>
public partial class AssetKit
{
    public string Id { get; set; } = null!;

    public virtual Asset Asset { get; set; } = null!;

    public string AssetId { get; set; } = null!;

    public virtual Kit Kit { get; set; } = null!;

    public string KitId { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public int Quantity { get; set; } = 1;

    public virtual ICollection<AssetLocation> AssetLocations { get; set; } = new List<AssetLocation>();

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
