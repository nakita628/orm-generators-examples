#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// A paid credit at the payment provider, which only metered usage draws down.
/// Kept here so a grant is created once per reason and the backoffice can show
/// what the customer holds without asking the provider.
/// </summary>
public partial class ConnectedCreditGrant
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// The `ConnectedBillingAccount` row. A plain column ref, like
    /// `IssuedLicense.replacesId`: relationMode="prisma" carries no foreign keys.
    /// </summary>
    public string AccountId { get; set; } = null!;

    /// <summary>
    /// The grant at the payment provider. Unique, so a retry records one row.
    /// </summary>
    public string StripeCreditGrantId { get; set; } = null!;

    public int AmountUsdCents { get; set; }

    public ConnectedCreditGrantKind Kind { get; set; }

    /// <summary>
    /// The term this credit was granted for, and when the provider lets it lapse:
    /// after the last invoice of that term has been finalized, never on the day
    /// the term ends.
    /// </summary>
    public DateTime TermEndsAt { get; set; }

    public DateTime ExpiresAt { get; set; }

    public DateTime CreatedAt { get; set; }
}
