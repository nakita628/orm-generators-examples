#nullable enable

using System;
using System.Collections.Generic;

namespace LangWatch.Models;

/// <summary>
/// Each organization's plan, folded from billing's subscription_changed facts: nurturing's own
/// copy. The newest fact wins, and a real fact wins a tie over a catch-up; Customer.io hears a
/// fact only when it changes this row.
/// </summary>
public partial class NurturingOrganizationPlan
{
    public string OrganizationId { get; set; } = null!;

    public string? Plan { get; set; }

    public DateTime OccurredAt { get; set; }

    public bool FromCatchUp { get; set; }

    /// <summary>
    /// The event ids folded in, so a redelivery folds once.
    /// </summary>
    public List<string>? AppliedEventIds { get; set; } = new List<string>();

    /// <summary>
    /// The fact that last changed the row; only it may tell Customer.io.
    /// </summary>
    public string? WinnerEventId { get; set; }
}
