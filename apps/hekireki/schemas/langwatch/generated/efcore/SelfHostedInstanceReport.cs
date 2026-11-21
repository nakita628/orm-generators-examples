#nullable enable

using System;

namespace LangWatch.Models;

/// <summary>
/// One usage report, exactly as it arrived. Kept so a number on an install's
/// row can be traced back to the report that set it, and so a question asked
/// later can be answered over the history rather than over the latest snapshot.
/// </summary>
public partial class SelfHostedInstanceReport
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// The `SelfHostedInstance.instanceId` this report came from. A plain column
    /// ref for the same reason the instance row's refs are.
    /// </summary>
    public string InstanceId { get; set; } = null!;

    public DateTime ReceivedAt { get; set; }

    public string? Version { get; set; }

    public int? ReportSchemaVersion { get; set; }

    /// <summary>
    /// Fields this release had no name for. Counted, never stored by name.
    /// </summary>
    public int UnknownFields { get; set; } = 0;

    /// <summary>
    /// The validated report, with every unknown field already dropped.
    /// </summary>
    public string Payload { get; set; } = null!;
}
