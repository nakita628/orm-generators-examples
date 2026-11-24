#nullable enable

using System;
using System.Collections.Generic;

namespace Trigger.Models;

public partial class BackgroundWorkerTask
{
    public string Id { get; set; } = null!;

    public string Slug { get; set; } = null!;

    public string? Description { get; set; }

    public string FriendlyId { get; set; } = null!;

    public string FilePath { get; set; } = null!;

    public string? ExportName { get; set; }

    public virtual BackgroundWorker Worker { get; set; } = null!;

    public string WorkerId { get; set; } = null!;

    public virtual Project Project { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    public virtual BackgroundWorkerFile? File { get; set; }

    public string? FileId { get; set; }

    public virtual RuntimeEnvironment RuntimeEnvironment { get; set; } = null!;

    public string RuntimeEnvironmentId { get; set; } = null!;

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual ICollection<TaskRunAttempt> Attempts { get; set; } = new List<TaskRunAttempt>();

    public virtual ICollection<TaskRun> Runs { get; set; } = new List<TaskRun>();

    public string? QueueConfig { get; set; }

    public string? RetryConfig { get; set; }

    public string? MachineConfig { get; set; }

    /// <summary>
    /// Gates declared on the task: other queues its runs must also hold a concurrency
    /// slot in while executing. Shape: [{ queue: string, concurrencyKey?: string }]
    /// </summary>
    public string? Gates { get; set; }

    public string? QueueId { get; set; }

    public virtual TaskQueue? Queue { get; set; }

    public int? MaxDurationInSeconds { get; set; }

    public string? Ttl { get; set; }

    public TaskTriggerSource TriggerSource { get; set; } = TaskTriggerSource.Standard;

    /// <summary>
    /// Extra task configuration JSON. Shape depends on triggerSource.
    /// AGENT: { type: "ai-sdk-chat" }
    /// </summary>
    public string? Config { get; set; }

    public TaskExecutionMode ExecutionMode { get; set; } = TaskExecutionMode.Standard;

    /// <summary>
    /// RunExecutionPolicy from @trigger.dev/core. Null means the standard policy.
    /// </summary>
    public string? ExecutionPolicy { get; set; }

    public string? PayloadSchema { get; set; }
}
