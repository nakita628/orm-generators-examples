#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// An invoice LangWatch raised for a connected customer, as the backoffice
/// reads it.
/// </summary>
public partial class ConnectedInvoice
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// The `ConnectedBillingAccount` row, as a plain column ref.
    /// </summary>
    public string AccountId { get; set; } = null!;

    public string StripeInvoiceId { get; set; } = null!;

    public ConnectedInvoiceKind Kind { get; set; }

    public Currency Currency { get; set; }

    public int AmountCents { get; set; }

    /// <summary>
    /// The provider's own status: `draft`, `open`, `paid`, `void` or `uncollectible`.
    /// </summary>
    public string Status { get; set; } = null!;

    public DateTime? PaidOutOfBandAt { get; set; }

    public DateTime? TermStartsAt { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
