#nullable enable

using System;

namespace Shelf.Models;

/// <summary>
/// *
/// * Audit trail for all quantity changes on quantity-tracked assets.
/// * Every event records who performed it, who received/returned items,
/// * which booking it relates to, and the category of change.
/// * Quantity is always positive — direction is determined by category.
/// </summary>
public partial class ConsumptionLog
{
    public string Id { get; set; } = null!;

    public virtual Asset Asset { get; set; } = null!;

    public string AssetId { get; set; } = null!;

    public ConsumptionCategory Category { get; set; }

    public int Quantity { get; set; }

    public string? Note { get; set; }

    public virtual User PerformedBy { get; set; } = null!;

    public string UserId { get; set; } = null!;

    public virtual Booking? Booking { get; set; }

    public string? BookingId { get; set; }

    public virtual TeamMember? Custodian { get; set; }

    public string? CustodianId { get; set; }

    public string? BookingAssetId { get; set; }

    public virtual Location? Location { get; set; }

    public string? LocationId { get; set; }

    public int? StockChange { get; set; }

    public DateTime CreatedAt { get; set; }
}
