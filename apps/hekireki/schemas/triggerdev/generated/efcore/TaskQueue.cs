#nullable enable

using System;
using System.Collections.Generic;

namespace Trigger.Models;

public partial class TaskQueue
{
    public string Id { get; set; } = null!;

    public string FriendlyId { get; set; } = null!;

    public string Name { get; set; } = null!;

    public TaskQueueType Type { get; set; } = TaskQueueType.Virtual;

    public TaskQueueVersion Version { get; set; } = TaskQueueVersion.V1;

    public string? OrderableName { get; set; }

    /// <summary>
    /// How the queue's concurrency was declared: V1 = legacy concurrencyLimit semantics
    /// (per key when keyed, whole queue otherwise), V2 = the explicit perKey/total shape.
    /// </summary>
    public TaskQueueConcurrencyVersion ConcurrencyVersion { get; set; } = TaskQueueConcurrencyVersion.V1;

    /// <summary>
    /// QUEUE rows are real queues; LIMIT rows back named concurrency limits and are
    /// excluded from every queue-facing surface.
    /// </summary>
    public TaskQueueRole Role { get; set; } = TaskQueueRole.Queue;

    public virtual Project Project { get; set; } = null!;

    public string ProjectId { get; set; } = null!;

    public virtual RuntimeEnvironment RuntimeEnvironment { get; set; } = null!;

    public string RuntimeEnvironmentId { get; set; } = null!;

    /// <summary>
    /// Represents the current concurrency limit for the queue
    /// </summary>
    public int? ConcurrencyLimit { get; set; }

    /// <summary>
    /// When the concurrency limit was overridden
    /// </summary>
    public DateTime? ConcurrencyLimitOverriddenAt { get; set; }

    /// <summary>
    /// Who overrode the concurrency limit (will be null if overridden via the API)
    /// </summary>
    public string? ConcurrencyLimitOverriddenBy { get; set; }

    /// <summary>
    /// If concurrencyLimit is overridden, this is the overridden value
    /// </summary>
    public int? ConcurrencyLimitBase { get; set; }

    /// <summary>
    /// If the override was expressed as a percentage of the environment limit, this stores that
    /// percentage (the source of truth). The absolute concurrencyLimit is materialized from it.
    /// Decimal(5,2) allows fractional percentages like 12.50% (0.01–100.00).
    /// </summary>
    public decimal? ConcurrencyLimitOverridePercent { get; set; }

    /// <summary>
    /// Caps total concurrent runs across ALL concurrencyKey values of this queue
    /// (concurrencyLimit applies per key value). Null = no total cap.
    /// </summary>
    public int? TotalConcurrencyLimit { get; set; }

    /// <summary>
    /// When the total concurrency limit was overridden
    /// </summary>
    public DateTime? TotalConcurrencyLimitOverriddenAt { get; set; }

    /// <summary>
    /// Who overrode the total concurrency limit (null when overridden via the API)
    /// </summary>
    public string? TotalConcurrencyLimitOverriddenBy { get; set; }

    /// <summary>
    /// If totalConcurrencyLimit is overridden, the declared value it reverts to on reset
    /// </summary>
    public int? TotalConcurrencyLimitBase { get; set; }

    public string? RateLimit { get; set; }

    public bool Paused { get; set; } = false;

    /// <summary>
    /// Hidden from the dashboard queue list. Display-only: never read on trigger/dequeue paths.
    /// </summary>
    public DateTime? ArchivedAt { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public virtual ICollection<TaskRunAttempt> Attempts { get; set; } = new List<TaskRunAttempt>();

    public virtual ICollection<BackgroundWorkerTask> Tasks { get; set; } = new List<BackgroundWorkerTask>();

    public virtual ICollection<BackgroundWorker> Workers { get; set; } = new List<BackgroundWorker>();
}
