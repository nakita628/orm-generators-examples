#nullable enable

using System;
using System.Collections.Generic;

namespace Shelf.Models;

/// <summary>
/// A progressive check-in SESSION, mirroring PartialBookingCheckout.
///
/// `assetIds` lists only FULLY reconciled assets: INDIVIDUAL always, and
/// QUANTITY_TRACKED only once its remaining units hit zero (partial returns
/// live in `ConsumptionLog`). Completing a booking outright writes no row here,
/// so the same warning applies — absence is not evidence of anything.
/// `BookingAsset.checkedInAt` is the per-slice fact.
/// </summary>
public partial class PartialBookingCheckin
{
    public string Id { get; set; } = null!;

    public List<string>? AssetIds { get; set; }

    public int CheckinCount { get; set; }

    public DateTime CheckinTimestamp { get; set; }

    public virtual Booking Booking { get; set; } = null!;

    public string BookingId { get; set; } = null!;

    public virtual User CheckedInBy { get; set; } = null!;

    public string CheckedInById { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
