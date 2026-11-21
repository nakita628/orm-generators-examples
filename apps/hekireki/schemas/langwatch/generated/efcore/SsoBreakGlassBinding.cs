#nullable enable

using System;
using System.Collections.Generic;

namespace LangWatch.Models;

/// <summary>
/// A way back in (D05 - see specs/identity/sso-onboarding-tiers.feature).
///
/// One named person who may still sign in without the identity provider
/// while their organization's sign-in belongs to one, and the date on which
/// that stops. Activation asks whether such a person exists, because the
/// failure it prevents is the only one that cannot be recovered from inside
/// the product: a misconfigured connection goes live, every member is handed
/// to an identity provider that will not authenticate them, and nobody is
/// left who can turn it off.
///
/// Rows are immutable except for `supersededAt` and `warnedDays`. A renewal
/// writes a NEW row naming the one it replaced, so the date a way in
/// previously ended is still readable in the history - which is what makes a
/// renewal auditable rather than a column somebody moved.
///
/// No relation to Organization or User, matching the identity heads above:
/// the binding is the record of who could get in, and a cascade would erase
/// exactly the history a lockout post-mortem reads.
/// </summary>
public partial class SsoBreakGlassBinding
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    /// <summary>
    /// Who holds the way in.
    /// </summary>
    public string UserId { get; set; } = null!;

    /// <summary>
    /// Who granted it. A way back in is never self-served.
    /// </summary>
    public string GrantedByUserId { get; set; } = null!;

    public DateTime GrantedAt { get; set; }

    public DateTime ExpiresAt { get; set; }

    /// <summary>
    /// Set when a renewal replaced this row; the row itself is never edited.
    /// </summary>
    public DateTime? SupersededAt { get; set; }

    /// <summary>
    /// The binding this one renewed, so the previous end date stays readable.
    /// </summary>
    public string? RenewedFromId { get; set; }

    /// <summary>
    /// Which of the fourteen, seven and one day warnings have been sent, so a
    /// sweep that runs twice in a day never says the same thing twice.
    /// </summary>
    public List<int>? WarnedDays { get; set; }

    public DateTime CreatedAt { get; set; }
}
