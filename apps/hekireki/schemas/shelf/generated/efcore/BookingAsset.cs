#nullable enable

using System;

namespace Shelf.Models;

/// <summary>
/// *
/// * Explicit pivot table between Booking and Asset.
/// * Replaces the implicit M2M (_AssetToBooking) via a rename migration in Phase 3.
/// </summary>
public partial class BookingAsset
{
    public string Id { get; set; } = null!;

    public virtual Booking Booking { get; set; } = null!;

    public string BookingId { get; set; } = null!;

    public virtual Asset Asset { get; set; } = null!;

    public string AssetId { get; set; } = null!;

    public int Quantity { get; set; } = 1;

    public string? BookingModelRequestId { get; set; }

    public string? AssetKitId { get; set; }

    public string? SourceKitId { get; set; }

    public DateTime? CheckedOutAt { get; set; }

    public string? CheckedOutById { get; set; }

    public DateTime? CheckedInAt { get; set; }

    public string? CheckedInById { get; set; }

    public int CheckedOutQuantity { get; set; } = 0;

    public string? SourceLocationId { get; set; }
}
