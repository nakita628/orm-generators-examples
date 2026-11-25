#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// retired: unread since v3.28.0. Replaced by NurturingProjectTraceDayCount, which the trace-day
/// peer folds write and a replay rebuilds; an operator step may drop it once old writers are gone.
/// </summary>
public partial class NurturingProjectTraceDay
{
    public string ProjectId { get; set; } = null!;

    /// <summary>
    /// UTC `YYYY-MM-DD` of the root spans' append time.
    /// </summary>
    public string Day { get; set; } = null!;

    public int TraceCount { get; set; } = 0;

    /// <summary>
    /// Traces trace recorded failed (lw.trace.trace_failed).
    /// </summary>
    public int ErrorCount { get; set; } = 0;

    /// <summary>
    /// Append time of the newest root span.
    /// </summary>
    public DateTime? LastTraceAt { get; set; }

    /// <summary>
    /// Append time of the newest trace_failed.
    /// </summary>
    public DateTime? LastFailedTraceAt { get; set; }
}
