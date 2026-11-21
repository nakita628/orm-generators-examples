#nullable enable

using System;
using System.Collections.Generic;

namespace Trigger.Models;

public partial class BatchTaskRun
{
    public string Id { get; set; } = null!;

    public string FriendlyId { get; set; } = null!;

    public string? IdempotencyKey { get; set; }

    public DateTime? IdempotencyKeyExpiresAt { get; set; }

    public BatchTaskRunStatus Status { get; set; } = BatchTaskRunStatus.Pending;

    public string RuntimeEnvironmentId { get; set; } = null!;

    /// <summary>
    /// This only includes new runs, not idempotent runs.
    /// </summary>
    public virtual ICollection<TaskRun> Runs { get; set; } = new List<TaskRun>();

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    /// <summary>
    /// Friendly IDs
    /// </summary>
    public List<string>? RunIds { get; set; } = new List<string>();

    public int RunCount { get; set; } = 0;

    public string? Payload { get; set; }

    public string PayloadType { get; set; } = "application/json";

    public string? Options { get; set; }

    public string BatchVersion { get; set; } = "v1";

    /// <summary>
    /// Specific run blockers,
    /// </summary>
    public virtual ICollection<TaskRunWaitpoint> RunsBlocked { get; set; } = new List<TaskRunWaitpoint>();

    /// <summary>
    /// Waitpoints that are blocked by this batch.
    /// When a Batch is created it blocks execution of the associated parent run (for andWait)
    /// </summary>
    public virtual ICollection<Waitpoint> Waitpoints { get; set; } = new List<Waitpoint>();

    /// <summary>
    /// sealed is set to true once no more items can be added to the batch
    /// </summary>
    public bool Sealed { get; set; } = false;

    public DateTime? SealedAt { get; set; }

    /// <summary>
    /// this is the expected number of items in the batch
    /// </summary>
    public int ExpectedCount { get; set; } = 0;

    /// <summary>
    /// this is the completed number of items in the batch. once this reaches expectedCount, and the batch is sealed, the batch is considered completed
    /// </summary>
    public int CompletedCount { get; set; } = 0;

    public DateTime? CompletedAt { get; set; }

    public DateTime? ResumedAt { get; set; }

    /// <summary>
    /// this is used to be able to "seal" this BatchTaskRun when all of the runs have been triggered asynchronously, and using the "parallel" processing strategy
    /// </summary>
    public int ProcessingJobsCount { get; set; } = 0;

    public int ProcessingJobsExpectedCount { get; set; } = 0;

    /// <summary>
    /// optional token that can be used to authenticate the task run
    /// </summary>
    public string? OneTimeUseToken { get; set; }

    /// <summary>
    /// When processing started (status changed to PROCESSING)
    /// </summary>
    public DateTime? ProcessingStartedAt { get; set; }

    /// <summary>
    /// When processing completed (all items processed)
    /// </summary>
    public DateTime? ProcessingCompletedAt { get; set; }

    /// <summary>
    /// Count of successfully created runs
    /// </summary>
    public int? SuccessfulRunCount { get; set; }

    /// <summary>
    /// Count of failed run creations
    /// </summary>
    public int? FailedRunCount { get; set; }

    /// <summary>
    /// Detailed failure records
    /// </summary>
    public virtual ICollection<BatchTaskRunError> Errors { get; set; } = new List<BatchTaskRunError>();

    /// <summary>
    /// all the below properties are engine v1 only
    /// </summary>
    public virtual ICollection<BatchTaskRunItem> Items { get; set; } = new List<BatchTaskRunItem>();

    public string? TaskIdentifier { get; set; }

    public virtual CheckpointRestoreEvent? CheckpointEvent { get; set; }

    public string? CheckpointEventId { get; set; }

    public virtual TaskRunAttempt? DependentTaskAttempt { get; set; }

    public string? DependentTaskAttemptId { get; set; }

    public virtual ICollection<TaskRunDependency> RunDependencies { get; set; } = new List<TaskRunDependency>();
}
