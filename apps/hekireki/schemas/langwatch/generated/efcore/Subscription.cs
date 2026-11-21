#nullable enable

using System;
using System.Collections.Generic;

namespace LangWatch.Models;

public partial class Subscription
{
    public string Id { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    public PlanTypes Plan { get; set; }

    public string? StripeSubscriptionId { get; set; }

    public SubscriptionStatus Status { get; set; }

    public DateTime? StartDate { get; set; }

    public DateTime? EndDate { get; set; }

    public DateTime? LastPaymentFailedDate { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public int? MaxMembers { get; set; }

    public int? MaxProjects { get; set; }

    public int? MaxMessagesPerMonth { get; set; }

    public int? MaxWorkflows { get; set; }

    public int? MaxRetentionDays { get; set; }

    public int? MaxMembersLite { get; set; }

    /// <summary>
    /// Developer seats the organization pays for; null falls to the plan's own cap.
    /// </summary>
    public int? MaxMembersDeveloper { get; set; }

    public int? MaxTeams { get; set; }

    public int? MaxPrompts { get; set; }

    public int? MaxEvaluators { get; set; }

    public int? MaxScenarios { get; set; }

    public int? MaxAgents { get; set; }

    public int? MaxExperiments { get; set; }

    public int? MaxOnlineEvaluations { get; set; }

    public int? MaxDatasets { get; set; }

    public int? MaxDashboards { get; set; }

    public int? MaxCustomGraphs { get; set; }

    public int? MaxAutomations { get; set; }

    /// <summary>
    /// Seats paid for this period that stay usable after a seat decrease, until seatsKeptUntil.
    /// </summary>
    public int? SeatsKept { get; set; }

    public DateTime? SeatsKeptUntil { get; set; }

    /// <summary>
    /// The seat quantity the last finished seat change set the seat line to; provider events never write it.
    /// </summary>
    public int? SeatLineQuantity { get; set; }

    public DateTime? SeatLineQuantityAt { get; set; }

    /// <summary>
    /// A Stripe seat line write recorded before it is sent and cleared once it is known to have
    /// finished under the lease; one left behind is settled by the seat line repair.
    /// Every seat quantity written while this record was pending.
    /// </summary>
    public List<int>? SeatLineWriteQuantities { get; set; } = new List<int>();

    /// <summary>
    /// The first write still pending; a later holder's write keeps it.
    /// </summary>
    public DateTime? SeatLineWriteSince { get; set; }

    /// <summary>
    /// The latest write sent; the repair waits a grace window after it.
    /// </summary>
    public DateTime? SeatLineWriteAt { get; set; }

    public int? SeatLineWriteAttempts { get; set; }

    public DateTime? SeatLineWriteFlaggedAt { get; set; }

    /// <summary>
    /// The seat quantity Stripe should bill if a seat change stops after lifting the seat line; null once settled.
    /// </summary>
    public int? SeatRepairQuantity { get; set; }

    /// <summary>
    /// The quantity the change lifted the line to; the repair moves the line back only while Stripe still bills it.
    /// </summary>
    public int? SeatRepairLiftedQuantity { get; set; }

    public DateTime? SeatRepairSince { get; set; }

    public int? SeatRepairAttempts { get; set; }

    /// <summary>
    /// Set once the seat line repair has failed often enough that Ops should look at it.
    /// </summary>
    public DateTime? SeatRepairFlaggedAt { get; set; }

    /// <summary>
    /// The lease a seat change or seat line repair holds while it reads and moves the seat line.
    /// </summary>
    public string? SeatLineLeaseToken { get; set; }

    public DateTime? SeatLineLeaseUntil { get; set; }

    public virtual Organization Organization { get; set; } = null!;

    public virtual ICollection<Invoice> Invoices { get; set; } = new List<Invoice>();
}
