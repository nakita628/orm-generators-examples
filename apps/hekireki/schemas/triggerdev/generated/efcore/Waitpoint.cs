#nullable enable

using System;
using System.Collections.Generic;

namespace Trigger.Models;

/// <summary>
/// A Waitpoint blocks a run from continuing until it's completed
/// If there's a waitpoint blocking a run, it shouldn't be in the queue
/// </summary>
public partial class Waitpoint
{
    public string Id { get; set; } = null!;

    public string FriendlyId { get; set; } = null!;

    public WaitpointType Type { get; set; }

    public WaitpointStatus Status { get; set; } = WaitpointStatus.Pending;

    public DateTime? CompletedAt { get; set; }

    /// <summary>
    /// If it's an Event type waitpoint, this is the event. It can also be provided for the DATETIME type
    /// </summary>
    public string IdempotencyKey { get; set; } = null!;

    /// <summary>
    /// If this is true then we can show it in the dashboard/return it from the SDK
    /// </summary>
    public bool UserProvidedIdempotencyKey { get; set; }

    /// <summary>
    /// If there's a user provided idempotency key, this is the time it expires at
    /// </summary>
    public DateTime? IdempotencyKeyExpiresAt { get; set; }

    /// <summary>
    /// If an idempotencyKey is no longer active, we store it here and generate a new one for the idempotencyKey field.
    /// Clearing an idempotencyKey is useful for debounce or cancelling child runs.
    /// This is a workaround because Prisma doesn't support partial indexes.
    /// </summary>
    public string? InactiveIdempotencyKey { get; set; }

    /// <summary>
    /// If it's a RUN type waitpoint, this is the associated run
    /// </summary>
    public string? CompletedByTaskRunId { get; set; }

    public virtual TaskRun? CompletedByTaskRun { get; set; }

    /// <summary>
    /// If it's a DATETIME type waitpoint, this is the date.
    /// If it's a MANUAL waitpoint, this can be set as the `timeout`.
    /// </summary>
    public DateTime? CompletedAfter { get; set; }

    /// <summary>
    /// If it's a BATCH type waitpoint, this is the associated batch
    /// </summary>
    public string? CompletedByBatchId { get; set; }

    public virtual BatchTaskRun? CompletedByBatch { get; set; }

    /// <summary>
    /// The runs this waitpoint is blocking
    /// </summary>
    public virtual ICollection<TaskRunWaitpoint> BlockingTaskRuns { get; set; } = new List<TaskRunWaitpoint>();

    /// <summary>
    /// All runs that have ever been blocked by this waitpoint, used for display purposes
    /// </summary>
    public virtual ICollection<TaskRun> ConnectedRuns { get; set; } = new List<TaskRun>();

    /// <summary>
    /// When a waitpoint is complete
    /// </summary>
    public virtual ICollection<TaskRunExecutionSnapshot> CompletedExecutionSnapshots { get; set; } = new List<TaskRunExecutionSnapshot>();

    /// <summary>
    /// When completed, an output can be stored here
    /// </summary>
    public string? Output { get; set; }

    public string OutputType { get; set; } = "application/json";

    public bool OutputIsError { get; set; } = false;

    public virtual Project Project { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    public virtual RuntimeEnvironment Environment { get; set; } = null!;

    public string EnvironmentId { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    /// <summary>
    /// Denormized column that holds the raw tags
    /// Denormalized column that holds the raw tags
    /// </summary>
    public List<string>? Tags { get; set; }
}
