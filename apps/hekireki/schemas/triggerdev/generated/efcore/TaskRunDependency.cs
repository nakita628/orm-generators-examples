#nullable enable

using System;

namespace Trigger.Models;

/// <summary>
/// This is used for triggerAndWait and batchTriggerAndWait. The taskRun is the child task, it points at a parent attempt or a batch
/// </summary>
public partial class TaskRunDependency
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// The child run
    /// </summary>
    public virtual TaskRun TaskRun { get; set; } = null!;

    public string TaskRunId { get; set; } = null!;

    public virtual CheckpointRestoreEvent? CheckpointEvent { get; set; }

    public string? CheckpointEventId { get; set; }

    /// <summary>
    /// An attempt that is dependent on this task run.
    /// </summary>
    public virtual TaskRunAttempt? DependentAttempt { get; set; }

    public string? DependentAttemptId { get; set; }

    /// <summary>
    /// A batch run that is dependent on this task run
    /// </summary>
    public virtual BatchTaskRun? DependentBatchRun { get; set; }

    public string? DependentBatchRunId { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public DateTime? ResumedAt { get; set; }
}
