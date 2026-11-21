#nullable enable

using System;
using System.Collections.Generic;

namespace Trigger.Models;

/// <summary>
/// Used by the RunEngine during TaskRun execution
/// It has the required information to transactionally progress a run through states,
/// and prevent side effects like heartbeats failing a run that has progressed.
/// It is optimised for performance and is designed to be cleared at some point,
/// so there are no cascading relationships to other models.
/// </summary>
public partial class TaskRunExecutionSnapshot
{
    public string Id { get; set; } = null!;

    /// <summary>
    /// This should always be 2+ (V1 didn't use the run engine or snapshots)
    /// </summary>
    public RunEngineVersion Engine { get; set; } = RunEngineVersion.V2;

    /// <summary>
    /// The execution status
    /// </summary>
    public TaskRunExecutionStatus ExecutionStatus { get; set; }

    /// <summary>
    /// For debugging
    /// </summary>
    public string Description { get; set; } = null!;

    /// <summary>
    /// We store invalid snapshots as a record of the run state when we tried to move
    /// </summary>
    public bool IsValid { get; set; } = true;

    public string? Error { get; set; }

    /// <summary>
    /// The previous snapshot ID
    /// </summary>
    public string? PreviousSnapshotId { get; set; }

    /// <summary>
    /// Run
    /// </summary>
    public string RunId { get; set; } = null!;

    public virtual TaskRun Run { get; set; } = null!;

    public TaskRunStatus RunStatus { get; set; }

    public string? BatchId { get; set; }

    /// <summary>
    /// This is the current run attempt number. Users can define how many attempts they want for a run.
    /// </summary>
    public int? AttemptNumber { get; set; }

    /// <summary>
    /// Environment
    /// </summary>
    public string EnvironmentId { get; set; } = null!;

    public RuntimeEnvironmentType EnvironmentType { get; set; }

    public string ProjectId { get; set; } = null!;

    public string OrganizationId { get; set; } = null!;

    /// <summary>
    /// Waitpoints that have been completed for this execution
    /// </summary>
    public virtual ICollection<Waitpoint> CompletedWaitpoints { get; set; } = new List<Waitpoint>();

    /// <summary>
    /// An array of waitpoint IDs in the correct order, used for batches
    /// </summary>
    public List<string>? CompletedWaitpointOrder { get; set; }

    /// <summary>
    /// Checkpoint
    /// </summary>
    public string? CheckpointId { get; set; }

    public virtual TaskRunCheckpoint? Checkpoint { get; set; }

    /// <summary>
    /// Worker
    /// </summary>
    public string? WorkerId { get; set; }

    public string? RunnerId { get; set; }

    public DateTime CreatedAt { get; set; }

    public DateTime UpdatedAt { get; set; }

    public DateTime? LastHeartbeatAt { get; set; }

    /// <summary>
    /// Metadata used by various systems in the run engine
    /// </summary>
    public string? Metadata { get; set; }
}
