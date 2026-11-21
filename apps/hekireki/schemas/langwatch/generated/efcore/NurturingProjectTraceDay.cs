#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// Nurturing's own trace counts from trace's span_received and trace_failed (ARCHITECTURE §9), a
/// root span being one trace, one row per project and UTC day of append time, incremented in place
/// by each batch. A redelivered batch can count a trace twice; accepted (issue 9378). Kept 30 days.
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
