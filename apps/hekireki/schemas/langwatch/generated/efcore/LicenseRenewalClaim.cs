#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// One renewed licence per paid Stripe invoice of a self-serve licence
/// subscription (specs/licensing/self-serving-license-purchase.feature). A
/// redelivered invoice event finds the row and mints nothing. A claim never
/// delivered is taken over once `claimedAt` is older than the claim lease.
/// </summary>
public partial class LicenseRenewalClaim
{
    public string StripeInvoiceId { get; set; } = null!;

    public string StripeSubscriptionId { get; set; } = null!;

    /// <summary>
    /// Fences completion and release to the delivery that holds the claim.
    /// </summary>
    public string ClaimToken { get; set; } = null!;

    public DateTime ClaimedAt { get; set; }

    /// <summary>
    /// The signed licence between minting and mailing, so a retry mails it again instead of minting another; cleared once delivered.
    /// </summary>
    public string? IssuedLicense { get; set; }

    public DateTime? DeliveredAt { get; set; }

    /// <summary>
    /// The `licenseId` inside the delivered licence.
    /// </summary>
    public string? LicenseId { get; set; }
}
