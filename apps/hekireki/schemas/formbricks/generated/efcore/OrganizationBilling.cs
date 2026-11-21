#nullable enable

using System;

namespace Formbricks.Models;

/// <summary>
/// Stores billing and Stripe synchronization data for an organization.
///
/// @property organizationId - Reference to the organization
/// @property stripeCustomerId - Stripe customer ID for cloud billing
/// @property limits - Plan limits used by feature and quota checks
/// @property usageCycleAnchor - Stable anchor used to derive the current monthly usage cycle
/// @property stripe - Stripe synchronization snapshot data
/// </summary>
public partial class OrganizationBilling
{
    public string OrganizationId { get; set; } = null!;

    public virtual Organization Organization { get; set; } = null!;

    public string? StripeCustomerId { get; set; }

    /// <summary>
    /// [OrganizationBillingPlanLimits]
    /// </summary>
    public string Limits { get; set; } = null!;

    public DateTime? UsageCycleAnchor { get; set; }

    /// <summary>
    /// [OrganizationStripeBilling]
    /// </summary>
    public string? Stripe { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }
}
