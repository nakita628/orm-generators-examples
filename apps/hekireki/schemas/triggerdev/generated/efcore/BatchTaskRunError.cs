#nullable enable

using System;

namespace Trigger.Models;

/// <summary>
/// Track individual run creation failures in batch processing (Run Engine v2)
/// </summary>
public partial class BatchTaskRunError
{
    public string Id { get; set; } = null!;

    public virtual BatchTaskRun BatchTaskRun { get; set; } = null!;

    public string BatchTaskRunId { get; set; } = null!;

    /// <summary>
    /// Which item in the batch (0-based index)
    /// </summary>
    public int Index { get; set; }

    /// <summary>
    /// The task identifier that was being triggered
    /// </summary>
    public string TaskIdentifier { get; set; } = null!;

    /// <summary>
    /// The payload that failed (JSON, may be truncated)
    /// </summary>
    public string? Payload { get; set; }

    /// <summary>
    /// The options that were used
    /// </summary>
    public string? Options { get; set; }

    /// <summary>
    /// Error message
    /// </summary>
    public string Error { get; set; } = null!;

    /// <summary>
    /// Error code if available
    /// </summary>
    public string? ErrorCode { get; set; }

    public DateTime CreatedAt { get; set; }
}
