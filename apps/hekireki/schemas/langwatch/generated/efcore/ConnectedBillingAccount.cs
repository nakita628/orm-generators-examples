#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// How a connected customer is invoiced (ADR-139, section 7): one payment
/// provider customer, one quarterly usage subscription in USD, and the
/// commercial terms the annual invoice was raised from.
///
/// One row per organization. Every write through it is addressed by a stored
/// provider id, so running onboarding or a renewal twice changes nothing.
/// </summary>
public partial class ConnectedBillingAccount
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    /// <summary>
    /// The customer at the payment provider. Invoiced, never charged.
    /// </summary>
    public string StripeCustomerId { get; set; } = null!;

    /// <summary>
    /// The one metered subscription, and the item hosted usage is reported on.
    /// Null between the customer being created and the subscription existing,
    /// which is where a half-finished onboarding resumes from.
    /// </summary>
    public string? UsageSubscriptionId { get; set; }

    public string? UsageSubscriptionItemId { get; set; }

    public DateTime TermStartsAt { get; set; }

    public DateTime TermEndsAt { get; set; }

    /// <summary>
    /// Prepaid usage for the term, in USD cents. The credit grants add up to it.
    /// </summary>
    public int CommitUsdCents { get; set; } = 0;

    /// <summary>
    /// The seat contract, which may be in another currency than hosted usage.
    /// </summary>
    public Currency SeatCurrency { get; set; }

    public int SeatRateCents { get; set; }

    public int Seats { get; set; }

    /// <summary>
    /// `us_bank_transfer` or `eu_bank_transfer` when the customer pays into a
    /// virtual bank account at the provider. Null when it wires to LangWatch's
    /// own account and finance marks the invoice paid.
    /// </summary>
    public string? BankTransferType { get; set; }

    public string? BankTransferCountry { get; set; }

    public string BillingEmail { get; set; } = null!;

    /// <summary>
    /// A renewal whose credit is waiting for the last usage invoice of the old
    /// term to be finalized: the new term's end and commit, and who asked.
    /// </summary>
    public string? PendingRenewal { get; set; }

    public string? CreatedById { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
