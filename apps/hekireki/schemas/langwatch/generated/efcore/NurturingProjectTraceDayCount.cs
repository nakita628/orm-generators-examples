#nullable enable

using System;
using System.Collections.Generic;

namespace LangWatch.Models;

/// <summary>
/// Nurturing's own trace counts (ARCHITECTURE §9), one row per project, UTC day of append time and
/// source: `spans` from trace's span_received peer fold (a customer root span is one trace),
/// `failures` from trace_failed's. Each fold owns its rows, so the two never overwrite each other;
/// readers sum a day's sources. Rebuilt by a projection replay. Kept 30 days.
/// </summary>
public partial class NurturingProjectTraceDayCount
{
    public string ProjectId { get; set; } = null!;

    /// <summary>
    /// UTC `YYYY-MM-DD` of the facts' append time.
    /// </summary>
    public string Day { get; set; } = null!;

    /// <summary>
    /// The peer fold that wrote the row: `spans` or `failures`.
    /// </summary>
    public string Source { get; set; } = null!;

    public int TraceCount { get; set; } = 0;

    public int ErrorCount { get; set; } = 0;

    /// <summary>
    /// Append time of the newest root span.
    /// </summary>
    public DateTime? LastTraceAt { get; set; }

    /// <summary>
    /// Append time of the newest trace_failed.
    /// </summary>
    public DateTime? LastFailedTraceAt { get; set; }

    /// <summary>
    /// The event ids of the last delivery folded in, so a redelivery folds once.
    /// </summary>
    public List<string>? AppliedEventIds { get; set; } = new List<string>();

    /// <summary>
    /// A row of another version reads as undecodable and is left out of the read.
    /// </summary>
    public string ProjectionVersion { get; set; } = null!;
}
