#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// What a mid-term seat change of one connected license owes (ADR-141,
/// section 7): the seats added on the reissued license, invoiced prorated to
/// the end of the term.
///
/// The row is written before the payment provider is called and completed
/// after it, so a call that failed halfway is retried by the next seat
/// invoicing pass without inventing a second invoice. One row per reissued license.
/// </summary>
public partial class ConnectedSeatChange
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// The reissued `IssuedLicense` row that carries the new seat count. A
    /// plain column ref, like `replacesId`: relationMode="prisma" carries no
    /// foreign keys.
    /// </summary>
    public string LicenseId { get; set; } = null!;

    /// <summary>
    /// The `ConnectedBillingAccount` row the invoice is raised against. Null
    /// when the customer had no billing account (`not_onboarded`).
    /// </summary>
    public string? AccountId { get; set; }

    /// <summary>
    /// The customer organization the license belongs to.
    /// </summary>
    public string? OrganizationId { get; set; }

    /// <summary>
    /// When the seats changed, which is the day the proration counts from.
    /// </summary>
    public DateTime ChangedAt { get; set; }

    /// <summary>
    /// Seats on the new license above the seats on the one it replaced.
    /// </summary>
    public int AddedSeats { get; set; }

    /// <summary>
    /// What one added seat costs for the rest of the term, in cents of `currency`.
    /// </summary>
    public int UnitAmountCents { get; set; }

    /// <summary>
    /// `unitAmountCents` times `addedSeats`. Zero when nothing is owed.
    /// </summary>
    public int AmountCents { get; set; }

    /// <summary>
    /// Null when no billing account named one (`not_onboarded`).
    /// </summary>
    public Currency? Currency { get; set; }

    /// <summary>
    /// The one-off invoice the added seats were billed on.
    /// </summary>
    public string? StripeInvoiceId { get; set; }

    public ConnectedSeatChangeState State { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
