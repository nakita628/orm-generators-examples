#nullable enable

using System;
using System.Collections.Generic;

namespace LangWatch.Models;

/// <summary>
/// A short code a fresh install pastes instead of a license blob (ADR-139,
/// section 5). Redeeming it mints the license the code describes and returns
/// it signed, which is what identifies a new customer without anybody emailing
/// a license key around.
///
/// The row holds a hash of the code, never the code. Every lookup hashes what
/// the caller presented first, so the stored value is a step away from anything
/// that could be presented: read access to this table is not credential access.
/// </summary>
public partial class ActivationCode
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// SHA-256 of the normalised code. Looked up on every redemption.
    /// </summary>
    public string CodeHash { get; set; } = null!;

    /// <summary>
    /// The last four characters, so an operator tells two codes apart in a list
    /// without the list holding either one.
    /// </summary>
    public string CodeHint { get; set; } = null!;

    /// <summary>
    /// The customer this code belongs to, and the terms it mints.
    /// </summary>
    public string OrganizationId { get; set; } = null!;

    public string OrganizationName { get; set; } = null!;

    public string Email { get; set; } = null!;

    public string PlanType { get; set; } = null!;

    public int MaxMembers { get; set; }

    /// <summary>
    /// The lite seats the minted license carries: null when unlimited. Never a stand-in number.
    /// </summary>
    public int? MaxMembersLite { get; set; } = 0;

    /// <summary>
    /// True when the minted license carries unlimited lite seats; maxMembersLite is then null.
    /// </summary>
    public bool LiteSeatsUnlimited { get; set; } = false;

    /// <summary>
    /// The Developer seats the minted license carries: null when not stated, which sells none.
    /// </summary>
    public int? MaxMembersDeveloper { get; set; }

    /// <summary>
    /// How long the license this code mints runs for.
    /// </summary>
    public int LicenseTermDays { get; set; }

    /// <summary>
    /// Hosted services the minted license may call.
    /// </summary>
    public List<string>? Services { get; set; } = new List<string>();

    /// <summary>
    /// When the code stops working, which is not the term of what it mints.
    /// </summary>
    public DateTime ExpiresAt { get; set; }

    /// <summary>
    /// A code a customer may redeem on more than one install, for a rollout.
    /// </summary>
    public bool Reusable { get; set; } = false;

    /// <summary>
    /// Set by the redemption that won, which is what makes a single-use code
    /// single use: the write itself carries the condition.
    /// </summary>
    public DateTime? RedeemedAt { get; set; }

    public string? RedeemedByInstanceId { get; set; }

    /// <summary>
    /// The license the winning redemption minted.
    /// </summary>
    public string? IssuedLicenseId { get; set; }

    /// <summary>
    /// How many installs have redeemed it, which only moves on a reusable code.
    /// </summary>
    public int RedemptionCount { get; set; } = 0;

    public DateTime? RevokedAt { get; set; }

    public string? RevokedById { get; set; }

    public string? CreatedById { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
