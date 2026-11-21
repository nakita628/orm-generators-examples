#nullable enable

using System;
using System.Collections.Generic;

namespace Shelf.Models;

/// <summary>
/// A progressive checkout SESSION: one scan batch, covering N assets via the
/// positional `assetIds` / `quantities` / `bookingAssetIds` arrays.
///
/// Records sessions and the units they claimed — NOT whether an asset is out.
/// The all-at-once checkout (`checkoutBooking`, the Check out button) flips
/// statuses and writes NO row here, so **absence proves nothing** and must never
/// be read as "was not checked out". Reading it that way is what let a single
/// later scan mark every button-checked-out asset on a booking as
/// "never checked out". `BookingAsset.checkedOutAt` answers that question.
/// </summary>
public partial class PartialBookingCheckout
{
    public string Id { get; set; } = null!;

    public List<string>? AssetIds { get; set; }

    public List<int>? Quantities { get; set; } = new List<int>();

    public List<string>? BookingAssetIds { get; set; } = new List<string>();

    public int CheckoutCount { get; set; }

    public DateTime CheckoutTimestamp { get; set; }

    public virtual Booking Booking { get; set; } = null!;

    public string BookingId { get; set; } = null!;

    public virtual User CheckedOutBy { get; set; } = null!;

    public string CheckedOutById { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
