#nullable enable

using System;

namespace Calcom.Models;

public partial class OrganizationOnboarding
{
    public string Id { get; set; } = null!;

    public virtual User CreatedBy { get; set; } = null!;

    public int CreatedById { get; set; }

    public DateTime CreatedAt { get; set; }

    public string OrgOwnerEmail { get; set; } = null!;

    public string? Error { get; set; }

    public DateTime UpdatedAt { get; set; }

    public int? OrganizationId { get; set; }

    public virtual Team? Organization { get; set; }

    public BillingPeriod BillingPeriod { get; set; }

    public double PricePerSeat { get; set; }

    public int Seats { get; set; }

    public bool IsPlatform { get; set; } = false;

    public string Name { get; set; } = null!;

    public string Slug { get; set; } = null!;

    public string? Logo { get; set; }

    public string? Bio { get; set; }

    public string? BrandColor { get; set; }

    public string? BannerUrl { get; set; }

    public bool IsDomainConfigured { get; set; } = false;

    public string? StripeCustomerId { get; set; }

    public string? StripeSubscriptionId { get; set; }

    public string? StripeSubscriptionItemId { get; set; }

    /// <summary>
    /// @zod.import(["import { orgOnboardingInvitedMembersSchema } from '../../zod-utils'"]).custom.use(orgOnboardingInvitedMembersSchema)
    /// </summary>
    public string InvitedMembers { get; set; } = "[]";

    /// <summary>
    /// @zod.import(["import { orgOnboardingTeamsSchema } from '../../zod-utils'"]).custom.use(orgOnboardingTeamsSchema)
    /// </summary>
    public string Teams { get; set; } = "[]";

    public bool IsComplete { get; set; } = false;
}
