#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// Whether a target was already stopped before a stop first touched it, recorded once before the
/// stop's platform call, so a retried call after a success still knows the change is the stop's own.
/// </summary>
public partial class AgentStopTargetPriorState
{
    public string OrganizationId { get; set; } = null!;

    public string StopId { get; set; } = null!;

    public string TargetId { get; set; } = null!;

    public bool WasStopped { get; set; }

    public DateTime RecordedAt { get; set; }
}
