#nullable enable

using System;
using System.Collections.Generic;

namespace LangWatch.Models;

/// <summary>
/// The registry of licenses LangWatch has issued (ADR-139). It lives on
/// LangWatch Cloud; a self-hosted install has the table and no rows.
///
/// The row never holds the token an install presents. `tokenHash` is the
/// SHA-256 of that token, so read access to this table is not credential
/// access. Status is derived: revoked, superseded, past `expiresAt`, else active.
/// </summary>
public partial class IssuedLicense
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// The `licenseId` inside the signed license (`lic-&lt;uuid&gt;`).
    /// </summary>
    public string LicenseId { get; set; } = null!;

    /// <summary>
    /// SHA-256 of the `lwl_` token. Looked up on every credential resolution.
    /// </summary>
    public string TokenHash { get; set; } = null!;

    /// <summary>
    /// The customer organization on LangWatch Cloud. Null for a license that is
    /// recorded but not linked to a customer yet, such as a self-serve purchase:
    /// it resolves to nothing until an operator links it.
    /// </summary>
    public string? OrganizationId { get; set; }

    public virtual Organization? Organization { get; set; }

    /// <summary>
    /// As signed into the license.
    /// </summary>
    public string OrganizationName { get; set; } = null!;

    public string Email { get; set; } = null!;

    public string PlanType { get; set; } = null!;

    public int MaxMembers { get; set; }

    /// <summary>
    /// The lite seats sold: null when not stated or when unlimited. Never a stand-in number.
    /// </summary>
    public int? MaxMembersLite { get; set; }

    /// <summary>
    /// True when the license sells unlimited lite seats; maxMembersLite is then null.
    /// </summary>
    public bool LiteSeatsUnlimited { get; set; } = false;

    /// <summary>
    /// The Developer seats sold: null when not stated, which sells none.
    /// </summary>
    public int? MaxMembersDeveloper { get; set; }

    public DateTime IssuedAt { get; set; }

    public DateTime ExpiresAt { get; set; }

    public IssuedLicenseSource Source { get; set; }

    /// <summary>
    /// The operator who issued or registered it. Null for the purchase flow and the script.
    /// </summary>
    public string? IssuedById { get; set; }

    public DateTime? RevokedAt { get; set; }

    public string? RevokedById { get; set; }

    public string? RevokedReason { get; set; }

    /// <summary>
    /// Set when the install first presents the license that replaced this one.
    /// </summary>
    public DateTime? SupersededAt { get; set; }

    /// <summary>
    /// The id of the IssuedLicense this one replaces, when it was reissued. A
    /// plain column ref (relationMode="prisma" cannot declare a self-relation on
    /// Postgres); the registry service validates it. Unique, so a license is
    /// replaced by one license at most.
    /// </summary>
    public string? ReplacesId { get; set; }

    /// <summary>
    /// A reissued license waiting to be delivered over sync, encrypted. Erased
    /// when the install first presents it. The only license text the registry holds.
    /// </summary>
    public string? PendingDeliveryLicense { get; set; }

    /// <summary>
    /// Hosted services this license may call, such as `instant_evals`.
    /// </summary>
    public List<string>? Services { get; set; } = new List<string>();

    /// <summary>
    /// The annual rate of one seat, in cents of `seatCurrency`. Needed to prorate added seats.
    /// </summary>
    public int? SeatRateCents { get; set; }

    public Currency? SeatCurrency { get; set; }

    /// <summary>
    /// The prepaid usage commit for the term, in USD cents.
    /// </summary>
    public int CommitUsdCents { get; set; } = 0;

    /// <summary>
    /// Whether usage may continue past the commit, up to `overageMaxUsdCents`.
    /// </summary>
    public bool OverageEnabled { get; set; } = false;

    public int? OverageMaxUsdCents { get; set; }

    /// <summary>
    /// The install this license is bound to. Set on first use, cleared by an operator.
    /// </summary>
    public string? InstanceId { get; set; }

    public DateTime? InstanceBoundAt { get; set; }

    /// <summary>
    /// The last license sync this install reported: when it arrived, the
    /// LangWatch version that sent it, and the seats in use at that moment.
    /// Null until the install syncs for the first time.
    /// </summary>
    public DateTime? LastSyncAt { get; set; }

    public string? LastSyncVersion { get; set; }

    public int? ReportedMembers { get; set; }

    public int? ReportedMembersLite { get; set; }

    /// <summary>
    /// The managed gateway key (purpose CONNECT) that hosted-service traffic of
    /// this license runs under. One per license, so spend rows name the install.
    /// Created when the license first resolves. A plain column, not a relation.
    /// </summary>
    public string? VirtualKeyId { get; set; }

    /// <summary>
    /// The seats of the license this one replaced, set only when a seat change
    /// raised a linked license. Billing reads it through licensing to invoice
    /// the added seats, so the fact commits with the license or not at all.
    /// </summary>
    public int? SeatsRaisedFrom { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
