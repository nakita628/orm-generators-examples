#nullable enable

using System;
using System.Collections.Generic;

namespace Trigger.Models;

public partial class TaskRun
{
    public string Id { get; set; } = null!;

    public int Number { get; set; } = 0;

    public string FriendlyId { get; set; } = null!;

    public RunEngineVersion Engine { get; set; } = RunEngineVersion.V1;

    public TaskRunStatus Status { get; set; } = TaskRunStatus.Pending;

    public string? StatusReason { get; set; }

    public string? IdempotencyKey { get; set; }

    public DateTime? IdempotencyKeyExpiresAt { get; set; }

    /// <summary>
    /// Stores the user-provided key and scope: { key: string, scope: "run" | "attempt" | "global" }
    /// </summary>
    public string? IdempotencyKeyOptions { get; set; }

    /// <summary>
    /// Debounce options: { key: string, delay: string, createdAt: Date }
    /// </summary>
    public string? Debounce { get; set; }

    public string TaskIdentifier { get; set; } = null!;

    public bool IsTest { get; set; } = false;

    public string Payload { get; set; } = null!;

    public string PayloadType { get; set; } = "application/json";

    public string? Context { get; set; }

    public string? TraceContext { get; set; }

    public string TraceId { get; set; } = null!;

    public string SpanId { get; set; } = null!;

    public virtual RuntimeEnvironment RuntimeEnvironment { get; set; } = null!;

    public string RuntimeEnvironmentId { get; set; } = null!;

    public RuntimeEnvironmentType? EnvironmentType { get; set; }

    public virtual Project Project { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    public string? OrganizationId { get; set; }

    public string Queue { get; set; } = null!;

    public string? LockedQueueId { get; set; }

    /// <summary>
    /// The main queue that this run is part of
    /// </summary>
    public string WorkerQueue { get; set; } = "main";

    /// <summary>
    /// User-facing geo region, stamped at trigger; workerQueue is where it actually ran.
    /// </summary>
    public string? Region { get; set; }

    /// <summary>
    /// @deprecated
    /// </summary>
    public string? SecondaryMasterQueue { get; set; }

    /// <summary>
    /// From engine v2+ this will be defined after a run has been dequeued (starting at 1)
    /// </summary>
    public int? AttemptNumber { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual ICollection<TaskRunAttempt> Attempts { get; set; } = new List<TaskRunAttempt>();

    /// <summary>
    /// Denormized column that holds the raw tags
    /// </summary>
    public List<string>? RunTags { get; set; }

    /// <summary>
    /// Denormalized version of the background worker task
    /// </summary>
    public string? TaskVersion { get; set; }

    public string? SdkVersion { get; set; }

    public string? CliVersion { get; set; }

    public virtual ICollection<Checkpoint> Checkpoints { get; set; } = new List<Checkpoint>();

    /// <summary>
    /// startedAt marks the point at which a run is dequeued from MarQS
    /// </summary>
    public DateTime? StartedAt { get; set; }

    /// <summary>
    /// executedAt is set when the first attempt is about to execute
    /// </summary>
    public DateTime? ExecutedAt { get; set; }

    public DateTime? CompletedAt { get; set; }

    public string? MachinePreset { get; set; }

    public int UsageDurationMs { get; set; } = 0;

    public double CostInCents { get; set; } = 0.0;

    public double BaseCostInCents { get; set; } = 0.0;

    public DateTime? LockedAt { get; set; }

    public virtual BackgroundWorkerTask? LockedBy { get; set; }

    public string? LockedById { get; set; }

    public virtual BackgroundWorker? LockedToVersion { get; set; }

    public string? LockedToVersionId { get; set; }

    /// <summary>
    /// The "priority" of the run. This is just a negative offset in ms for the queue timestamp
    /// E.g. a value of 60_000 would put the run into the queue 60s ago.
    /// </summary>
    public int PriorityMs { get; set; } = 0;

    public string? ConcurrencyKey { get; set; }

    /// <summary>
    /// Gates for this run: other queues it must also hold a concurrency slot in while
    /// executing. Shape: [{ queue: string, concurrencyKey?: string }]
    /// </summary>
    public string? Gates { get; set; }

    public DateTime? DelayUntil { get; set; }

    public DateTime? QueuedAt { get; set; }

    public string? Ttl { get; set; }

    public DateTime? ExpiredAt { get; set; }

    public int? MaxAttempts { get; set; }

    public string? LockedRetryConfig { get; set; }

    /// <summary>
    /// optional token that can be used to authenticate the task run
    /// </summary>
    public string? OneTimeUseToken { get; set; }

    /// <summary>
    /// When this run is finished, the waitpoint will be marked as completed
    /// </summary>
    public virtual Waitpoint? AssociatedWaitpoint { get; set; }

    /// <summary>
    /// If there are any blocked waitpoints, the run won't be executed
    /// </summary>
    public virtual ICollection<TaskRunWaitpoint> BlockedByWaitpoints { get; set; } = new List<TaskRunWaitpoint>();

    /// <summary>
    /// All waitpoints that blocked this run at some point, used for display purposes
    /// </summary>
    public virtual ICollection<Waitpoint> ConnectedWaitpoints { get; set; } = new List<Waitpoint>();

    /// <summary>
    /// Where the logs are stored
    /// </summary>
    public string TaskEventStore { get; set; } = "taskEvent";

    public DateTime? QueueTimestamp { get; set; }

    public virtual ICollection<BatchTaskRunItem> BatchItems { get; set; } = new List<BatchTaskRunItem>();

    public virtual TaskRunDependency? Dependency { get; set; }

    public virtual ICollection<CheckpointRestoreEvent> CheckpointRestoreEvent { get; set; } = new List<CheckpointRestoreEvent>();

    public virtual ICollection<TaskRunExecutionSnapshot> ExecutionSnapshots { get; set; } = new List<TaskRunExecutionSnapshot>();

    public string? ScheduleInstanceId { get; set; }

    public string? ScheduleId { get; set; }

    public virtual ICollection<BulkActionItem> SourceBulkActionItems { get; set; } = new List<BulkActionItem>();

    public virtual ICollection<BulkActionItem> DestinationBulkActionItems { get; set; } = new List<BulkActionItem>();

    public List<string>? BulkActionGroupIds { get; set; } = new List<string>();

    public DateTime? LogsDeletedAt { get; set; }

    public string? ReplayedFromTaskRunFriendlyId { get; set; }

    /// <summary>
    /// This represents the original task that that was triggered outside of a Trigger.dev task
    /// </summary>
    public virtual TaskRun? RootTaskRun { get; set; }

    public string? RootTaskRunId { get; set; }

    /// <summary>
    /// The root run will have a list of all the descendant runs, children, grand children, etc.
    /// </summary>
    public virtual ICollection<TaskRun> DescendantRuns { get; set; } = new List<TaskRun>();

    /// <summary>
    /// The immediate parent run of this task run
    /// </summary>
    public virtual TaskRun? ParentTaskRun { get; set; }

    public string? ParentTaskRunId { get; set; }

    /// <summary>
    /// The immediate child runs of this task run
    /// </summary>
    public virtual ICollection<TaskRun> ChildRuns { get; set; } = new List<TaskRun>();

    /// <summary>
    /// The immediate parent attempt of this task run
    /// </summary>
    public virtual TaskRunAttempt? ParentTaskRunAttempt { get; set; }

    public string? ParentTaskRunAttemptId { get; set; }

    /// <summary>
    /// The batch run that this task run is a part of
    /// </summary>
    public virtual BatchTaskRun? Batch { get; set; }

    public string? BatchId { get; set; }

    /// <summary>
    /// whether or not the task run was created because of a triggerAndWait for batchTriggerAndWait
    /// </summary>
    public bool ResumeParentOnCompletion { get; set; } = false;

    /// <summary>
    /// The depth of this task run in the task run hierarchy
    /// </summary>
    public int Depth { get; set; } = 0;

    /// <summary>
    /// The span ID of the "trigger" span in the parent task run
    /// </summary>
    public string? ParentSpanId { get; set; }

    /// <summary>
    /// Holds the state of the run chain for deadlock detection
    /// </summary>
    public string? RunChainState { get; set; }

    /// <summary>
    /// seed run metadata
    /// </summary>
    public string? SeedMetadata { get; set; }

    public string SeedMetadataType { get; set; } = "application/json";

    /// <summary>
    /// Run metadata
    /// </summary>
    public string? Metadata { get; set; }

    public string MetadataType { get; set; } = "application/json";

    public int MetadataVersion { get; set; } = 1;

    /// <summary>
    /// Structured annotations: triggerSource, triggerAction, rootTriggerSource, rootScheduleId
    /// </summary>
    public string? Annotations { get; set; }

    /// <summary>
    /// Whether the latest attempt was a warm start. Null until first attempt starts.
    /// </summary>
    public bool? IsWarmStart { get; set; }

    /// <summary>
    /// Run output
    /// </summary>
    public string? Output { get; set; }

    public string OutputType { get; set; } = "application/json";

    /// <summary>
    /// Run error
    /// </summary>
    public string? Error { get; set; }

    /// <summary>
    /// Organization's billing plan type (cached for fallback when billing API fails)
    /// </summary>
    public string? PlanType { get; set; }

    public int? MaxDurationInSeconds { get; set; }

    /// <summary>
    /// The version of the realtime streams implementation used by the run
    /// </summary>
    public string RealtimeStreamsVersion { get; set; } = "v1";

    /// <summary>
    /// Store the stream keys that are being used by the run
    /// </summary>
    public List<string>? RealtimeStreams { get; set; } = new List<string>();

    /// <summary>
    /// S2 basin where this run's realtime streams live. Stamped at create
    /// time from `Organization.streamBasinName` so reads can resolve the
    /// basin without joining org. Null when the org has no per-org basin
    /// (OSS, or pre-backfill); reads fall back to the global basin.
    /// </summary>
    public string? StreamBasinName { get; set; }

    public virtual ICollection<PlaygroundConversation> PlaygroundConversations { get; set; } = new List<PlaygroundConversation>();
}
