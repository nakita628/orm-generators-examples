#nullable enable

using System;
using System.Collections.Generic;

namespace LangWatch.Models;

/// <summary>
/// Nurturing's own per-project usage: a peer fold over the owners' prompt, evaluator, scenario
/// and simulation-run events (ARCHITECTURE §9), one peer fold per owner pipeline, summed per
/// organization for the per-email traits. Traces are counted in NurturingProjectTraceDay.
/// </summary>
public partial class NurturingProjectUsageView
{
    public string ProjectId { get; set; } = null!;

    /// <summary>
    /// The owner pipeline whose peer fold wrote the row (prompt, evaluator, scenario, simulation),
    /// or `evaluation`: the project's evaluation days, written by a subscriber, at most once a day.
    /// </summary>
    public string Owner { get; set; } = null!;

    /// <summary>
    /// The project's usage as the fold folded it (NurturingProjectUsageState).
    /// </summary>
    public string Folded { get; set; } = null!;

    /// <summary>
    /// The event ids of the last delivery folded in, so a redelivery folds once.
    /// </summary>
    public List<string>? AppliedEventIds { get; set; } = new List<string>();

    /// <summary>
    /// A row of another version reads as undecodable and is re-folded from the log.
    /// </summary>
    public string ProjectionVersion { get; set; } = null!;

    /// <summary>
    /// Event-derived, so a replayed row equals the row it rebuilds.
    /// </summary>
    public DateTime UpdatedAt { get; set; }
}
